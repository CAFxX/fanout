package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint64 }
func genFn0() uint {
    switch 640 {
    case -260:
        if false {
            var v5 bool = (false || true)
            fmt.Println(v5)
            fmt.Println(v5)
            _ = v5
        } else {
            switch 2147483647 {
            case -39:
                var v6 uint16 = 235
                _ = v6
            case 2147483647:
                for fi7 := 0; fi7 < 3; fi7++ {
                    for fi8 := 0; fi8 < 17; fi8++ {
                        var v9 uint = 18446744073709551615
                        var v10 uint = (v9 | v9)
                        var v11 uint = 255
                        v11 = (v11 | (v10 >> ((255) & 63)))
                        _ = fi8
                        _ = v10
                        _ = v11
                        _ = v9
                    }
                    if (!(false || (((4294967295 >= 462) && true) || ((666 > 32767) || true)))) {
                        var v12 uint = 295
                        fmt.Println(v12)
                        v12 += (v12 % ((v12) | 1))
                        _ = v12
                    }
                    _ = fi7
                }
            case 9223372036854775807:
                var v13 uint16 = 42
                _ = v13
            default:
                var v14 int = 275
                _ = v14
            }
        }
    case 327:
        switch 2147483647 {
        case -287:
            var v15 uint = 4294967295
            _ = v15
        default:
            var v16 string = "a"
            _ = v16
        }
    default:
        var v17 int64 = -1
        _ = v17
    }
    var v18 int32 = -2147483648
    var v19 int16 = -1
    var v20 float32 = 0.0
    _ = v18
    _ = v19
    _ = v20
    return 876
}
func genFn1(p0 int16, p1 bool, p2 uint) uint {
    var v21 uint32 = 514
    switch "hello" {
    case "hello":
        v22 := []int{186}
        _ = v22
    case "":
        var v23 uint8 = 0
        _ = v23
    }
    v24 := GenS1{F0: 0, F1: 0}
    var v25 uint = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v24
    _ = v25
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{}
    var v3 float64 = 3.14159
    v4 := []string{"a"}
    switch -9223372036854775808 {
    case 4294967295:
        switch (("\n\t" + "foo bar") + ("foo bar" + ("" + "hello"))) {
        case "foo bar":
            for fi26 := 0; fi26 < 5; fi26++ {
                v3 = (float64(1e10) - v3)
                v3 -= 0.0
                _ = fi26
            }
        }
    case 128:
        v3 += 0.1
    case 181:
        var v27 uint64 = 1
        _ = v27
    default:
        delete(v1, 42)
    }
    var v28 uint8 = 2
    var v29 uint = genFn0()
    var v30 int64 = 9223372036854775807
    v28 += (v28 % (((v28 & v28)) | 1))
    var v31 float64 = v3
    v30 = -9223372036854775808
    var v32 int8 = 1
    fmt.Println(v28, v3, v3)
    var v33 uint32 = 954
    v34 := GenS2{F0: 510}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
