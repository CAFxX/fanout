package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint, p1 int8) bool {
    v5 := GenS1{F0: 0}
    var v6 uint32 = 65535
    if false {
        v6 = 42
    } else {
        var v7 int32 = 0
        switch 128 {
        case 684:
            v8 := []int{128, -534}
            _ = v8
        }
        _ = v7
    }
    v9 := []int{4294967295}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v9
    return true
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    v3 := []string{}
    v4 := make(map[uint32]bool)
    var v10 uint8 = 4
    var v11 int64 = 0
    v12 := GenS2{F0: 18446744073709551615}
    var v13 int32 = 2147483647
    if true {
        v11 += (v11 | v11)
    }
    var v14 bool = true
    var v15 bool = (("\n\t" + "hello") < "")
    v15 = false
    var v16 uint32 = 731
    v17 := GenS2{F0: 89}
    for fi18 := 0; fi18 < 8; fi18++ {
        v11 -= (v11 << ((42) & 63))
        _ = fi18
    }
    if (42 != 566) {
        var v19 int32 = v13
        var v20 int32 = v13
        if true {
            fmt.Println(v10, v20, v20)
            delete(v4, (v16 ^ v16))
            fmt.Println(v14)
        } else {
            fmt.Println(v11, v20)
        }
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 int = (fi21 / ((-1) | 1))
            v22 -= int(v10)
            v11 = -1
            _ = fi21
            _ = v22
        }
        v13 -= 253
    }
    var v23 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
