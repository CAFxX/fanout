package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint8 }
func genFn0(p0 uint8, p1 int, p2 uint64) uint8 {
    var v6 int16 = -78
    v7 := GenS1{F0: 255, F1: v6, F2: 1}
    var v8 uint8 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (v8 >> ((0) & 63))
}

func main() {
    var v1 bool = (-358 <= 0)
    v2 := make(map[uint32]bool)
    var v3 uint = 432
    var v4 uint8 = 0
    v5 := make(map[string]bool)
    if ((1 >= 557) || (-2147483648 > 995)) {
        if v1 {
            var v9 uint16 = 1
            var v10 bool = true
            delete(v5, ("\n\t" + "a"))
            _ = v10
            _ = v9
        } else {
            v11 := v2[1]
            v12 := []int{4294967295}
            var v13 int32 = 42
            _ = v11
            _ = v12
            _ = v13
        }
    }
    v14 := []int{0, -1}
    if v1 {
        var v15 int32 = 42
        v3 += (v3 + v3)
        _ = v15
    } else {
        if v1 {
            var v16 uint8 = v4
            var v17 int32 = 42
            var v18 uint64 = 1
            _ = v16
            _ = v17
            _ = v18
        } else {
            var v19 string = "a"
            v19 = v19
            _ = v19
        }
        v1 = true
    }
    v1 = (!(-1 <= -13))
    var v20 int = 32768
    for fi21 := 0; fi21 < 15; fi21++ {
        fmt.Println(v4, v3)
        _ = fi21
    }
    v22 := GenS1{F0: (v3 * (v3 % (((v3 + v3)) | 1))), F1: -32768, F2: v4}
    if ((-746 != -403) && v1) {
        fmt.Println(v3, v3, v1)
        v3 = 0
    } else {
        v4 += (v4 % ((v4) | 1))
        fmt.Println(v20, v1, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
