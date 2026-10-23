package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int64) int64 {
    p0 = (p0 * (p0 - (p0 & p0)))
    var v7 bool = ((("hello" + ("foo bar" + ("\n\t" + (("\n\t" + "") + "a")))) + "\n\t") < ("" + "foo bar"))
    switch 42 {
    case 42:
        v8 := GenS1{F0: -1}
        _ = v8
    case 128:
        v9 := GenS1{F0: 0}
        _ = v9
    case 942:
        var v10 uint8 = 249
        _ = v10
    default:
        switch 4294967295 {
        case -385:
            v11 := GenS1{F0: 42}
            _ = v11
        case -509:
            switch -9223372036854775808 {
            case 2147483647:
                var v12 uint8 = 255
                _ = v12
            case -425:
                var v13 int64 = (p0 << ((18446744073709551615) & 63))
                _ = v13
            }
        case -2147483648:
            switch -2147483648 {
            case 4294967295:
                for fi14 := 0; fi14 < 15; fi14++ {
                    var v15 uint16 = 32767
                    if false {
                        var v16 int16 = 1
                        v17 := GenS1{F0: 1}
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
            case -736:
                v18 := []int{698}
                _ = v18
            case -2147483648:
                v7 = true
            default:
                var v19 int64 = -9223372036854775808
                _ = v19
            }
        default:
            var v20 bool = ((v7 && v7) || ((!v7) && (1 == (p0 % ((0) | 1)))))
            _ = v20
        }
    }
    _ = p0
    _ = v7
    return p0
}
func genFn1(p0 int16, p1 uint, p2 int) uint32 {
    v21 := GenS1{F0: 0}
    var v22 int = 128
    p2 = v22
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return 2
}

func main() {
    var v1 uint64 = 267
    var v2 uint = 0
    var v3 string = "foo bar"
    v4 := make(map[string]uint32)
    v5 := make(map[string]uint32)
    v6 := make(map[int]int)
    var v23 uint8 = 47
    for fi24 := 0; fi24 < 13; fi24++ {
        v1 -= v1
        _ = fi24
    }
    var v25 int16 = 42
    var v26 int = -2147483648
    var v27 int32 = -2147483648
    fmt.Println(v3, v27, v25)
    v23 = (v23 * (v23 >> (((^(v2))) & 63)))
    var v28 uint32 = genFn1((-(v25)), v2, (v26 & v26))
    var v29 int64 = -1
    var v30 float32 = (float32((((1.5 + 0.0) / (-2.25 + 0.1)) - (3.14159 - ((3.14159 * 1e10) - (1e10 - ((-2.25 * 3.14159) + 1.5)))))) * (-2.25 * 1e10))
    v2 = v2
    var v31 int32 = 2147483647
    fmt.Println(v3, v1)
    fmt.Println(v2, v31, v25)
    v32 := []int{(-(v26)), -488, (v26 & 42)}
    switch v3 {
    case "hello":
        var v33 uint64 = (v1 / (((v1 % ((v1) | 1))) | 1))
        _ = v33
    case "":
        v34 := v6[32767]
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
