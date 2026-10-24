package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int8 }
func genFn0(p0 uint8, p1 uint32, p2 int32) uint {
    p2 += 1
    var v5 int16 = 0
    p2 += p2
    var v6 uint8 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 bool, p1 int8, p2 int32) uint16 {
    p1 -= (p1 ^ 94)
    v7 := GenS1{F0: 42, F1: 42, F2: p1}
    p0 = p0
    if p0 {
        v8 := []int{127, 127}
        if p0 {
            switch 127 {
            case -1:
                for fi9 := 0; fi9 < 12; fi9++ {
                    var v10 uint8 = 1
                    _ = len(v8)
                    v11 := GenS1{F0: 2147483648, F1: 18446744073709551615, F2: -1}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case 127:
                p1 += (p1 / ((p1) | 1))
            case 642:
                var v12 string = "foo bar"
                _ = v12
            default:
                var v13 uint32 = 65535
                _ = v13
            }
        } else {
            fmt.Println(p0, p1)
        }
        _ = v8
    } else {
        p1 -= p1
        var v14 uint32 = 65535
        _ = v14
    }
    var v15 int = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 1
    var v3 int8 = 0
    v4 := make(map[string]uint32)
    v2 += v2
    delete(v1, 42)
    if false {
        v16 := []int{1, 0, -9223372036854775808}
        fmt.Println(v3, v2)
        fmt.Println(v2)
        _ = v16
    }
    v17 := GenS1{F0: 2147483647, F1: 18446744073709551615, F2: (v3 + (v3 << ((genFn0(0, 4294967295, -340)) & 63)))}
    v3 = 127
    v18 := []int{1, -106, -9223372036854775808}
    if true {
        fmt.Println(v3, v2)
        v19 := GenS1{F0: 0, F1: 476, F2: (v3 - -1)}
        v20 := GenS1{F0: 149, F1: 2, F2: (^(v3))}
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: 4294967295, F1: 354, F2: (v3 << ((65535) & 63))}
    var v22 uint = 23
    var v23 uint64 = 18446744073709551615
    v4["a"] = 968
    var v24 int64 = 9223372036854775807
    v4["foo bar"] = 953
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := v4[""]
        if true {
            var v27 float32 = 3.14159
            _ = v27
        }
        switch fi25 {
        case 42:
            v23 = uint64(v26)
        }
        _ = fi25
        _ = v26
    }
    var v28 uint32 = 2147483647
    switch -9223372036854775808 {
    case 127:
        v24 = v2
    case 272:
        switch -1 {
        case -917:
            var v29 int8 = v3
            _ = v29
        case 9223372036854775807:
            var v30 string = "hello"
            _ = v30
        default:
            fmt.Println(v23, v2)
        }
    case 925:
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    default:
        var v31 uint64 = 330
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
}
