package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 int32; F1 string }
func genFn0(p0 uint8) int32 {
    for fi7 := 0; fi7 < 5; fi7++ {
        switch fi7 {
        case 32768:
            v8 := []int{fi7, 32768, 127}
            _ = v8
        default:
            fmt.Println(p0)
        }
        v9 := []int{363, (-(fi7)), (fi7 % ((fi7) | 1)), (fi7 | 9223372036854775807)}
        _ = fi7
        _ = v9
    }
    v10 := GenS1{F0: 388, F1: 111, F2: -1}
    v11 := GenS2{F0: 264, F1: 18446744073709551615, F2: 2}
    _ = p0
    _ = v10
    _ = v11
    return -119
}
func genFn1() uint {
    v12 := GenS3{F0: 0, F1: genFn0(255), F2: 1}
    var v13 int32 = -1
    v14 := []int{127, -221, 32767, -661}
    var v15 int64 = 42
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 0
}

func main() {
    var v1 float64 = -2.25
    var v2 int64 = 42
    v3 := []string{(("foo bar" + "\n\t") + ("" + "hello")), ("\n\t" + ("" + ("\n\t" + ("\n\t" + ("\n\t" + "")))))}
    var v4 uint = 472
    var v5 float64 = v1
    var v6 bool = (("\n\t" + ("\n\t" + "")) < "foo bar")
    v16 := GenS3{F0: 0, F1: 0, F2: 175}
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS3{F0: 174, F1: 2147483647, F2: 65535}
            fmt.Println(fi18, v5)
            _ = len(v3)
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 int8 = 127
            fmt.Println(v6, v1, v21)
            v21 = v21
            v1 -= v1
            _ = fi20
            _ = v21
        }
        v6 = ((4294967295 > 2) || (0.1 >= v5))
        _ = fi17
    }
    v22 := GenS4{F0: 0, F1: v2}
    fmt.Println(v6, v2)
    switch "" {
    case "foo bar":
        for fi23 := 0; fi23 < 1; fi23++ {
            for fi24 := 0; fi24 < 16; fi24++ {
                if v6 {
                    v25 := GenS3{F0: 255, F1: -1, F2: 2}
                    if false {
                        _ = len(v3)
                    } else {
                        v26 := GenS5{F0: 1, F1: "\n\t"}
                        v3 = append(v3, "foo bar")
                        _ = v26
                    }
                    _ = v25
                }
                _ = len(v3)
                v2 = (v2 ^ v2)
                _ = fi24
            }
            _ = fi23
        }
    case "a":
        fmt.Println(v6, v4, v5)
    }
    v4 = (v4 + v4)
    v4 += 189
    var v27 bool = v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
