package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int8; F1 int16; F2 int64 }
func genFn0(p0 uint32, p1 int32, p2 string) int32 {
    var v7 float64 = (0.1 * (3.14159 + 1e10))
    v8 := GenS1{F0: 42}
    var v9 int32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1() uint8 {
    v10 := []int{2147483647}
    v11 := GenS1{F0: 807}
    _ = v10
    _ = v11
    return 255
}

func main() {
    var v1 int32 = -41
    v2 := []string{""}
    v3 := []uint32{65535, 42, 1, 767}
    v4 := make(map[string]bool)
    v5 := make(map[uint32]bool)
    v6 := []uint32{42, 2147483648, 4294967295}
    var v12 bool = (0 < 1000)
    _ = len(v4)
    switch -9223372036854775808 {
    case 827:
        v13 := v5[156]
        _ = v13
    case -292:
        v4[""] = false
    case -2147483648:
        var v14 int64 = 42
        _ = v14
    default:
        v15 := GenS2{F0: 42, F1: (-(v1))}
        _ = v15
    }
    v12 = (224 >= 0)
    var v16 bool = v12
    switch "hello" {
    case "a":
        if v16 {
            fmt.Println(v12, v1)
        }
    }
    _ = len(v4)
    v17 := GenS2{F0: 2, F1: -1}
    switch ("foo bar" + ("foo bar" + "a")) {
    case "a":
        v18 := GenS3{F0: 42, F1: "\n\t"}
        _ = v18
    case "":
        for fi19 := 0; fi19 < 13; fi19++ {
            switch "foo bar" {
            case "foo bar":
                v20 := GenS1{F0: 1000}
                _ = v20
            case "a":
                switch -704 {
                case 9223372036854775807:
                    var v21 int32 = v1
                    _ = v21
                case 1:
                    v16 = v16
                default:
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                }
            }
            _ = fi19
        }
    case "hello":
        var v22 uint8 = genFn1()
        _ = v22
    default:
        if (!(("foo bar" > "") || true)) {
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 int16 = 1000
                v12 = v12
                if (127 <= genFn1()) {
                    for fi25 := 0; fi25 < 12; fi25++ {
                        switch "hello" {
                        case "a":
                            v26 := v5[65535]
                            _ = v26
                        case "\n\t":
                            v12 = (2 == 65535)
                        case "":
                            v27 := GenS2{F0: 91, F1: v1}
                            _ = v27
                        default:
                            v12 = (v12 && (!(65535 < 42)))
                        }
                        _ = fi25
                    }
                } else {
                    var v28 uint8 = 127
                    fmt.Println(v16, v12)
                    _ = v28
                }
                _ = fi23
                _ = v24
            }
            fmt.Println(v12, v12)
            v16 = true
        } else {
            fmt.Println(v12, v1)
            v16 = v16
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        v12 = v12
        var v30 float64 = (1.5 / 1.0)
        _ = fi29
        _ = v30
    }
    v31 := GenS4{F0: -128, F1: -502, F2: 572}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
