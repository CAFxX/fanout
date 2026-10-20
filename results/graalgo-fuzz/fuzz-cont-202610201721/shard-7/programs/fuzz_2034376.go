package main

import (
    "fmt"
)

func genFn0(p0 uint64) uint16 {
    v5 := []int{}
    var v6 uint32 = 4294967295
    _ = p0
    _ = v5
    _ = v6
    return uint16(p0)
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 2
    v3 := []uint32{v2, (v2 | (v2 << ((2) & 63))), v2, v2}
    v4 := make(map[int]int)
    if false {
        if (-447 != -578) {
            var v7 int8 = 19
            _ = v7
        } else {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 bool = true
                _ = fi8
                _ = v9
            }
            var v10 uint8 = 255
            _ = v10
        }
        var v11 int16 = 1000
        _ = v11
    }
    var v12 bool = true
    fmt.Println(v12)
    var v13 uint16 = 255
    fmt.Println(v12, v2, v12)
    if (!(v12 || (v13 >= v13))) {
        var v14 uint = 4294967295
        _ = v14
    }
    if ((-208 <= -1) && (1 > (v13 | (v13 + 65535)))) {
        v13 = 786
        var v15 string = "hello"
        var v16 int = 0
        _ = v15
        _ = v16
    }
    v12 = v12
    if v12 {
        v2 -= 1
        for fi17 := 0; fi17 < 10; fi17++ {
            if (!false) {
                var v18 uint8 = uint8(v13)
                _ = v18
            }
            for fi19 := 0; fi19 < 16; fi19++ {
                v12 = true
                _ = fi19
            }
            _ = fi17
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
