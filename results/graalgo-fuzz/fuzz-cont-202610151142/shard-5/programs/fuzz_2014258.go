package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int16, p1 uint32, p2 int64) uint {
    if (!(2 <= 65535)) {
        for fi5 := 0; fi5 < 5; fi5++ {
            p2 = (p2 - -44)
            v6 := []int{int(p1), fi5, -9223372036854775808, (fi5 & fi5)}
            v7 := GenS1{F0: 9223372036854775807}
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: -126}
        _ = v8
    }
    var v9 bool = true
    p2 = p2
    fmt.Println(p0, v9)
    var v10 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]uint32)
    v3 := []uint32{4294967295, 4294967295}
    var v4 uint16 = 861
    v4 -= 32767
    fmt.Println(v1, v4)
    v11 := GenS1{F0: 128}
    v12 := GenS1{F0: -852}
    if false {
        v1 = v1
        v4 = 350
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        if (9 >= 1) {
            switch -9223372036854775808 {
            case 127:
                for fi13 := 0; fi13 < 10; fi13++ {
                    fmt.Println(fi13, v1, v1)
                    _ = fi13
                }
            default:
                v4 = (v4 % ((0) | 1))
            }
        } else {
            var v14 uint64 = 42
            v4 = 32767
            _ = v14
        }
    }
    var v15 float64 = float64(float32(-2.25))
    fmt.Println(v1, v1, v4)
    var v16 int32 = v1
    var v17 int8 = -32
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
