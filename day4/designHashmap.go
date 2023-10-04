package day4

type MyHashMap struct {
    arr []*Node
    size int
}

type Node struct{
    key int
    value int
    next *Node
}

func Constructor() MyHashMap {
    return MyHashMap {
        arr : make([]*Node,10000),
        size : 10000,
    }
}


func (this *MyHashMap) Put(key int, value int)  {
    idx := key % this.size
    newNode := Node{
        key : key,
        value : value,
        next : nil,
    }

    if this.arr[idx] != nil {
        cur := this.arr[idx]
        prev := this.arr[idx]

        for cur != nil {
            if cur.key == key {
                cur.value = value
                return
            }
            prev = cur
            cur = cur.next
        }
        prev.next = &newNode
        return
    }

    this.arr[idx] = &newNode
}


func (this *MyHashMap) Get(key int) int {
    idx := key % this.size

    if this.arr[idx] == nil {
        return -1
    }

    cur := this.arr[idx]

    for cur != nil {
        if cur.key == key {
            return cur.value
        }
        cur = cur.next
    }

    return -1
}


func (this *MyHashMap) Remove(key int)  {
    idx := key % this.size

    if this.arr[idx] == nil {
        return
    }

    cur := this.arr[idx]

    if cur.key == key {
        cur = cur.next
        this.arr[idx] = cur
        return
    }
    prev := this.arr[idx]

    for cur != nil {
        if cur.key == key {
            prev.next = cur.next
            break
        }
        prev = cur
        cur = cur.next
    }
}


