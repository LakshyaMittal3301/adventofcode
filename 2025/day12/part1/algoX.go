package main

type Node struct {
	up, down, left, right, column *Node
	rowNum, colNum, nodeCount     int
}

func NewNode() *Node {
	return &Node{
		up:        nil,
		down:      nil,
		left:      nil,
		right:     nil,
		column:    nil,
		rowNum:    0,
		colNum:    0,
		nodeCount: 0,
	}
}

type DLX struct {
	root         *Node
	currSolution []int
	Solution     []int
	HasSolution  bool
}

func (d *DLX) getMinColumn() *Node {
	minCol := d.root.right
	currCol := d.root.right
	for currCol != d.root {
		if currCol.nodeCount < minCol.nodeCount {
			minCol = currCol
		}
		currCol = currCol.right
	}
	return minCol
}

func (d *DLX) cover(colNode *Node) {
	colNode.left.right = colNode.right
	colNode.right.left = colNode.left
	for rowNode := colNode.down; rowNode != colNode; rowNode = rowNode.down {
		for rightNode := rowNode.right; rightNode != rowNode; rightNode = rightNode.right {
			rightNode.up.down = rightNode.down
			rightNode.down.up = rightNode.up

			rightNode.column.nodeCount -= 1
		}
	}
}

func (d *DLX) uncover(colNode *Node) {
	for rowNode := colNode.up; rowNode != colNode; rowNode = rowNode.up {
		for leftNode := rowNode.left; leftNode != rowNode; leftNode = leftNode.left {
			leftNode.up.down = leftNode
			leftNode.down.up = leftNode

			leftNode.column.nodeCount++
		}
	}
	colNode.left.right = colNode
	colNode.right.left = colNode
}

func (d *DLX) Search() {
	if d.HasSolution {
		return
	}
	if d.root.right == d.root {
		d.saveSolution()
	}

	colNode := d.getMinColumn()

	d.cover(colNode)

	for rowNode := colNode.down; rowNode != colNode; rowNode = rowNode.down {
		d.currSolution = append(d.currSolution, rowNode.rowNum)
		for rightNode := rowNode.right; rightNode != rowNode; rightNode = rightNode.right {
			d.cover(rightNode.column)
		}

		d.Search()

		d.currSolution = d.currSolution[:len(d.currSolution)-1]

		for leftNode := rowNode.left; leftNode != rowNode; leftNode = leftNode.left {
			d.uncover(leftNode.column)
		}
	}

	d.uncover(colNode)
}

func (d *DLX) saveSolution() {
	d.Solution = append(d.Solution, d.currSolution...)
	d.HasSolution = true
}

func NewDLX(matrix [][]int) *DLX {
	n := len(matrix)
	m := len(matrix[0])
	root := NewNode()
	root.up, root.down, root.left, root.right = root, root, root, root
	headerNodes := make([](*Node), 0)
	headerNodes = append(headerNodes, root)
	for i := range m {
		node := NewNode()
		node.up, node.down = node, node
		node.left = headerNodes[len(headerNodes)-1]
		node.right = headerNodes[0]

		node.left.right = node
		node.right.left = node

		node.colNum = i
		node.column = node
		headerNodes = append(headerNodes, node)
	}

	headerNodes = headerNodes[1:]

	for i := range n {
		var prevNode *Node
		for j := range m {
			if matrix[i][j] == 0 {
				continue
			}

			node := NewNode()
			node.down = headerNodes[j]
			node.up = node.down.up

			node.down.up = node
			node.up.down = node

			node.rowNum = i
			node.colNum = j
			node.column = headerNodes[j]
			headerNodes[j].nodeCount += 1
			if prevNode == nil {
				node.left, node.right = node, node
			} else {
				node.left = prevNode
				node.right = node.left.right

				node.left.right = node
				node.right.left = node
			}
			prevNode = node
		}
	}
	return &DLX{
		root:         root,
		HasSolution:  false,
		currSolution: make([]int, 0),
		Solution:     make([]int, 0),
	}
}
