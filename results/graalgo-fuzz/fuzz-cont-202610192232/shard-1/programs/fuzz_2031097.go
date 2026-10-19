package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
func genFn0() bool {
    var v6 int32 = -2147483648
    var v7 uint8 = 127
    _ = v6
    _ = v7
    return true
}
func genFn1(p0 int) uint16 {
    fmt.Println(p0)
    if ("hello" == "a") {
        if false {
            var v8 int32 = 42
            _ = v8
        } else {
            p0 += p0
            var v9 int16 = 1
            p0 = (p0 + (p0 ^ p0))
            _ = v9
        }
        var v10 uint64 = 42
        v10 -= 1
        _ = v10
    }
    _ = p0
    return 32767
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = true
    var v3 uint32 = 0
    v4 := make(map[int]int)
    v5 := make(map[uint32]string)
    v11 := GenS1{F0: 32768}
    var v12 int16 = 42
    v1[9223372036854775807] = ""
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := v1[(-(fi13))]
        var v15 string = "a"
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := []int{4294967295}
    v17 := GenS1{F0: 9223372036854775807}
    if (-1 < -2147483648) {
        v18 := GenS2{F0: (v3 * (v3 ^ (^(v3))))}
        _ = v18
    }
    _ = len(v16)
    v12 = (v12 - (v12 >> ((1) & 63)))
    if (!v2) {
        var v19 int64 = 953
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
