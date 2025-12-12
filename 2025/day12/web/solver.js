// Client-side DLX solver for 3x3 tile fitting with optional empty cells.

const DEFAULT_MAX_AREA = 100;
const DEFAULT_MAX_PIECES = 6;
const DEFAULT_MAX_QTY = 50;
const DEFAULT_ITERATION_LIMIT = 2_000_000;

class Node {
  constructor() {
    this.up = this;
    this.down = this;
    this.left = this;
    this.right = this;
    this.column = this;
    this.rowNum = 0;
    this.colNum = 0;
    this.nodeCount = 0;
  }
}

class DLX {
  constructor(rows, numCols, iterationLimit) {
    this.currSolution = [];
    this.solution = [];
    this.hasSolution = false;
    this.iterationLimit = iterationLimit;
    this.iterations = 0;

    const root = new Node();
    root.up = root;
    root.down = root;
    root.left = root;
    root.right = root;

    const headers = [];
    headers.push(root);
    for (let i = 0; i < numCols; i++) {
      const node = new Node();
      node.up = node;
      node.down = node;
      node.left = headers[headers.length - 1];
      node.right = headers[0];

      node.left.right = node;
      node.right.left = node;

      node.colNum = i;
      node.column = node;
      headers.push(node);
    }

    headers.shift(); // remove root
    for (let rowIdx = 0; rowIdx < rows.length; rowIdx++) {
      const cols = rows[rowIdx];
      let prev = null;
      for (const col of cols) {
        const node = new Node();
        const header = headers[col];
        node.down = header;
        node.up = header.up;
        header.up.down = node;
        header.up = node;

        node.rowNum = rowIdx;
        node.colNum = col;
        node.column = header;
        header.nodeCount += 1;
        if (prev === null) {
          node.left = node;
          node.right = node;
        } else {
          node.left = prev;
          node.right = prev.right;
          prev.right.left = node;
          prev.right = node;
        }
        prev = node;
      }
    }

    this.root = root;
  }

  getMinColumn() {
    let min = this.root.right;
    let curr = this.root.right;
    while (curr !== this.root) {
      if (curr.nodeCount < min.nodeCount) {
        min = curr;
      }
      curr = curr.right;
    }
    return min;
  }

  cover(col) {
    col.left.right = col.right;
    col.right.left = col.left;
    for (let row = col.down; row !== col; row = row.down) {
      for (let right = row.right; right !== row; right = right.right) {
        right.up.down = right.down;
        right.down.up = right.up;
        right.column.nodeCount -= 1;
      }
    }
  }

  uncover(col) {
    for (let row = col.up; row !== col; row = row.up) {
      for (let left = row.left; left !== row; left = left.left) {
        left.up.down = left;
        left.down.up = left;
        left.column.nodeCount += 1;
      }
    }
    col.left.right = col;
    col.right.left = col;
  }

  saveSolution() {
    this.solution.push(...this.currSolution);
    this.hasSolution = true;
  }

  search() {
    if (this.hasSolution) return;
    if (this.iterations > this.iterationLimit) return;
    if (this.root.right === this.root) {
      this.saveSolution();
      return;
    }

    const col = this.getMinColumn();
    if (col.nodeCount === 0) return;

    this.cover(col);
    for (let row = col.down; row !== col; row = row.down) {
      this.currSolution.push(row.rowNum);
      for (let right = row.right; right !== row; right = right.right) {
        this.cover(right.column);
      }

      this.iterations += 1;
      if (this.iterations > this.iterationLimit) return;

      this.search();
      if (this.hasSolution) return;

      this.currSolution.pop();
      for (let left = row.left; left !== row; left = left.left) {
        this.uncover(left.column);
      }
    }
    this.uncover(col);
  }
}

function rotate(shape) {
  const res = [
    [false, false, false],
    [false, false, false],
    [false, false, false],
  ];
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      res[i][j] = shape[j][i];
    }
  }
  return res;
}

function flip(shape) {
  const res = [
    [false, false, false],
    [false, false, false],
    [false, false, false],
  ];
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      res[i][2 - j] = shape[i][j];
    }
  }
  return res;
}

function isSameShape(a, b) {
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      if (a[i][j] !== b[i][j]) return false;
    }
  }
  return true;
}

function uniqueVariants(shape) {
  const variants = [];
  let current = shape;
  for (let r = 0; r < 4; r++) {
    const rotShape = current;
    if (!variants.some((s) => isSameShape(s, rotShape))) {
      variants.push(rotShape);
    }
    const flipped = flip(rotShape);
    if (!variants.some((s) => isSameShape(s, flipped))) {
      variants.push(flipped);
    }
    current = rotate(rotShape);
  }
  return variants;
}

function normalizeShape(shape) {
  if (shape.length !== 3) throw new Error("shape must be 3 rows");
  const res = [
    [false, false, false],
    [false, false, false],
    [false, false, false],
  ];
  for (let i = 0; i < 3; i++) {
    if (shape[i].length !== 3) throw new Error("shape rows must be length 3");
    for (let j = 0; j < 3; j++) {
      res[i][j] = Boolean(shape[i][j]);
    }
  }
  return res;
}

function findCellsForShape(shape, baseRow, baseCol, width) {
  const cells = [];
  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      if (shape[i][j]) {
        const row = baseRow + i;
        const col = baseCol + j;
        cells.push(row * width + col);
      }
    }
  }
  return cells;
}

function buildMatrix(height, width, pieces) {
  const totalCells = height * width;
  const maxPieceId = Math.max(...pieces.map((p) => p.id));
  const idOffset = Array(maxPieceId + 1).fill(0);
  let totalPieces = 0;
  for (const piece of pieces) {
    idOffset[piece.id] = totalPieces;
    totalPieces += piece.qty;
  }

  const totalColumns = totalCells + totalPieces;
  const rows = [];
  const meta = [];

  for (let cell = 0; cell < totalCells; cell++) {
    rows.push([cell]);
    meta.push({ kind: "empty", cell });
  }

  for (const piece of pieces) {
    const variants = uniqueVariants(piece.shape);
    for (let v = 0; v < variants.length; v++) {
      const variant = variants[v];
      for (let r = 0; r <= height - 3; r++) {
        for (let c = 0; c <= width - 3; c++) {
          const cells = findCellsForShape(variant, r, c, width);
          if (cells.some((idx) => idx >= totalCells || idx < 0)) continue;
          for (let copy = 0; copy < piece.qty; copy++) {
            const row = [];
            row.push(...cells);
            row.push(totalCells + idOffset[piece.id] + copy);
            rows.push(row);
            meta.push({
              kind: "piece",
              pieceId: piece.id,
              variantId: v,
              originRow: r,
              originCol: c,
              cells,
            });
          }
        }
      }
    }
  }
  return { rows, totalColumns, meta };
}

function validateConfig(config, options) {
  const { height, width, pieces } = config;
  const maxArea = config.maxArea ?? DEFAULT_MAX_AREA;
  const maxPieces = options.maxPieces ?? DEFAULT_MAX_PIECES;
  const maxQtyPerPiece = options.maxQtyPerPiece ?? DEFAULT_MAX_QTY;

  if (height <= 0 || width <= 0) {
    throw new Error("height and width must be positive");
  }
  if (height * width > maxArea) {
    throw new Error(`grid area exceeds limit (${maxArea})`);
  }
  if (pieces.length > maxPieces) {
    throw new Error(`too many pieces (max ${maxPieces})`);
  }
  for (const p of pieces) {
    if (p.qty < 0) throw new Error("quantity cannot be negative");
    if (p.qty > maxQtyPerPiece) {
      throw new Error(`quantity for piece ${p.id} exceeds limit ${maxQtyPerPiece}`);
    }
    p.shape = normalizeShape(p.shape);
  }
}

export function solve(config, options = {}) {
  validateConfig(config, options);
  const iterationLimit = options.iterationLimit ?? DEFAULT_ITERATION_LIMIT;
  const { rows, totalColumns, meta } = buildMatrix(
    config.height,
    config.width,
    config.pieces
  );
  const dlx = new DLX(rows, totalColumns, iterationLimit);
  dlx.search();
  if (!dlx.hasSolution) {
    return { hasSolution: false, placements: [] };
  }
  const placements = [];
  for (const rowIdx of dlx.solution) {
    const info = meta[rowIdx];
    if (info && info.kind === "piece") {
      placements.push({
        pieceId: info.pieceId,
        variantId: info.variantId,
        originRow: info.originRow,
        originCol: info.originCol,
        cells: info.cells.slice(),
      });
    }
  }
  return { hasSolution: true, placements };
}
