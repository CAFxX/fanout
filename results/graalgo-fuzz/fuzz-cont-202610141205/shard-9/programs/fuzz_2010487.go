package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int32; F2 int8 }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32 }
func genFn0() uint16 {
    v4 := []int{4294967295}
    v5 := GenS1{F0: true, F1: 42}
    if false {
        var v6 uint64 = 0
        for fi7 := 0; fi7 < 17; fi7++ {
            v6 = 1
            v6 += (v6 * 0)
            _ = fi7
        }
        switch "hello" {
        case "foo bar":
            v6 = (v6 - 42)
        case "":
            var v8 bool = false
            _ = v8
        }
        _ = v6
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v9 := GenS2{F0: 374}
    _ = v4
    _ = v5
    _ = v9
    return 255
}

func main() {
    var v1 float64 = -2.25
    var v2 uint64 = 69
    v3 := make(map[uint32]bool)
    var v10 float32 = (1e10 * -2.25)
    var v11 uint8 = 2
    var v12 uint64 = v2
    switch "" {
    case "a":
        var v13 bool = (v2 == v12)
        _ = v13
    default:
        fmt.Println(v11, v2)
    }
    var v14 uint8 = (v11 * 251)
    var v15 int8 = 1
    var v16 uint32 = 262
    v10 = v10
    v17 := []int{1}
    var v18 bool = false
    v3[v16] = (((3.14159 - (v1 + float64((v10 + 0.0)))) + float64(-2.25)) == v1)
    v2 += v2
    if (!(-736 > 42)) {
        if ((!((!(!v18)) || false)) || (!v18)) {
            var v19 uint64 = 42
            v20 := GenS3{F0: "foo bar", F1: 596, F2: (v15 * (-(v15)))}
            v15 = (v15 | int8(v16))
            _ = v19
            _ = v20
        } else {
            v21 := GenS4{F0: -9223372036854775808, F1: 1, F2: (509 != -2147483648)}
            switch 259 {
            case 32767:
                if v18 {
                    var v22 int32 = -1
                    _ = v22
                }
            case 1:
                var v23 int64 = -58
                _ = v23
            default:
                var v24 int = -1
                _ = v24
            }
            _ = v21
        }
        v25 := GenS5{F0: 2}
        _ = len(v17)
        _ = v25
    } else {
        fmt.Println(v11, v10, v18)
        fmt.Println(v12)
        v10 += 1.5
    }
    v12 += v2
    v17 = append(v17, 0)
    v26 := GenS6{F0: (v16 + (^(v16)))}
    for fi27 := 0; fi27 < 12; fi27++ {
        v11 = v14
        fmt.Println(v11, v11)
        fmt.Println(v16, v2)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
}
