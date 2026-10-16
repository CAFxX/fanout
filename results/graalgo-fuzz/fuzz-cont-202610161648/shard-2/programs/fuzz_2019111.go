package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() int {
    var v7 int32 = -984
    v8 := GenS1{F0: 65535}
    v9 := []int{32767, 127, -1}
    v10 := GenS1{F0: 65535}
    v11 := GenS1{F0: 2}
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 42
    v3 := []uint32{42}
    var v4 int64 = 0
    var v5 uint16 = 65535
    var v6 int = -946
    v12 := GenS1{F0: 2147483648}
    delete(v1, 2147483647)
    var v13 uint64 = 724
    var v14 int64 = (-(v4))
    var v15 int16 = 0
    if false {
        for fi16 := 0; fi16 < 7; fi16++ {
            if len(v3) > 0 {
                v3[0] = 42
            }
            v13 = v13
            _ = fi16
        }
    } else {
        v15 = (v15 ^ (-(v15)))
    }
    v4 += (v4 % ((v4) | 1))
    v17 := v1[42]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
