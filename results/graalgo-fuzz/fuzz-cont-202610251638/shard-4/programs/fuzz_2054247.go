package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint8, p1 bool, p2 bool) int64 {
    fmt.Println(p0)
    v7 := GenS1{F0: 42, F1: -987, F2: 32}
    p2 = (0 < 18446744073709551615)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v4 int8 = 0
    v5 := make(map[uint32]uint32)
    var v6 uint16 = 32767
    var v8 uint = 18446744073709551615
    switch 128 {
    case -9223372036854775808:
        var v9 int16 = -32768
        _ = v9
    case -2147483648:
        v6 -= v6
    default:
        delete(v2, "hello")
    }
    var v10 uint32 = uint32(v6)
    fmt.Println(v6, v4)
    v11 := GenS1{F0: 1, F1: -1, F2: 186}
    if (((v6 % (((v6 ^ v6)) | 1)) >= 1) && true) {
        var v12 uint8 = 128
        if true {
            v8 -= (^(v8))
            delete(v3, v10)
            switch 947 {
            case 4294967295:
                var v13 float32 = (1e10 * 1e10)
                _ = v13
            }
        }
        v14 := GenS1{F0: (v10 & v10), F1: 0, F2: (v12 | v12)}
        _ = v12
        _ = v14
    } else {
        v15 := []int{128, 1, 0}
        fmt.Println(v10)
        _ = v15
    }
    var v16 bool = true
    fmt.Println(v6, v8)
    var v17 int = -2147483648
    var v18 int32 = 0
    v19 := GenS1{F0: v10, F1: genFn0(0, (-9223372036854775808 >= int(v8)), true), F2: 9}
    switch "" {
    case "foo bar":
        v20 := GenS1{F0: (v10 ^ (v10 / ((uint32(v6)) | 1))), F1: 0, F2: 0}
        _ = v20
    }
    fmt.Println(v8)
    for fi21 := 0; fi21 < 18; fi21++ {
        v22 := GenS2{F0: 128}
        var v23 string = "foo bar"
        _ = fi21
        _ = v22
        _ = v23
    }
    v10 = 686
    v24 := GenS1{F0: v10, F1: -1, F2: 42}
    fmt.Println(v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
