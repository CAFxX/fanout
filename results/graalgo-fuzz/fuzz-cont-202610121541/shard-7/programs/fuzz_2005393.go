package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int64 }
func genFn0(p0 uint64) int8 {
    var v6 bool = false
    v7 := []int{301, -1, 4294967295}
    p0 += p0
    fmt.Println(v6, p0)
    if (1 == 0) {
        var v8 int64 = 335
        for fi9 := 0; fi9 < 8; fi9++ {
            p0 += (p0 / ((p0) | 1))
            v10 := GenS1{F0: v6, F1: 1, F2: 2147483648}
            _ = fi9
            _ = v10
        }
        _ = len(v7)
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return 27
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]string)
    var v3 int8 = 42
    var v4 uint64 = 2
    var v5 uint = 2
    fmt.Println(v4, v1, v3)
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS2{F0: 174}
        var v13 int = (fi11 << ((v5) & 63))
        _ = fi11
        _ = v12
        _ = v13
    }
    v4 = (v4 & v4)
    fmt.Println(v5, v1)
    var v14 uint64 = v4
    v15 := []int{128, 0}
    _ = len(v15)
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 string = ((((v1 + (v1 + ("foo bar" + v1))) + ((((v1 + (((v1 + (v1 + v1)) + "a") + (v1 + v1))) + "a") + v1) + "\n\t")) + "\n\t") + v1)
        var v18 uint16 = 65535
        _ = fi16
        _ = v17
        _ = v18
    }
    v3 -= v3
    var v19 string = v1
    v20 := GenS3{F0: 0, F1: (v4 * (v14 ^ (v14 + 0))), F2: int64(v5)}
    v4 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
