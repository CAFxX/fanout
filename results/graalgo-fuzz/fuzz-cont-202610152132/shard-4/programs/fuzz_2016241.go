package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }
func genFn0(p0 uint16, p1 uint8, p2 uint64) int {
    p2 = 501
    fmt.Println(p1)
    p2 += (^(p2))
    if (!((!true) && ((p0 | p0) < 832))) {
        p2 += p2
    } else {
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 string = "hello"
                v9 := []int{(fi7 - 114), 2147483647, (fi7 % ((fi6) | 1))}
                _ = fi7
                _ = v8
                _ = v9
            }
            p1 -= (^(p1))
            if (-174 != 9223372036854775807) {
                v10 := GenS1{F0: (p1 ^ (p1 | uint8(p2))), F1: 2147483647, F2: (fi6 - 428)}
                _ = v10
            }
            _ = fi6
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            switch "\n\t" {
            case "a":
                switch (fi11 - int(p2)) {
                case 128:
                    switch "hello" {
                    case "":
                        var v12 int16 = -1
                        _ = v12
                    default:
                        if ((p2 - (p2 * uint64(p1))) >= (p2 / ((uint64(p0)) | 1))) {
                            var v13 uint16 = p0
                            var v14 float32 = 1e10
                            v13 = (v13 >> ((4294967295) & 63))
                            _ = v13
                            _ = v14
                        }
                    }
                case -307:
                    var v15 int16 = 44
                    _ = v15
                default:
                    p1 -= (p1 + (p1 / (((p1 + (p1 ^ p1))) | 1)))
                }
            case "\n\t":
                v16 := []int{-949, fi11}
                _ = v16
            case "hello":
                p2 = 928
            }
            p0 = p0
            for fi17 := 0; fi17 < 17; fi17++ {
                if (1e10 == 1.5) {
                    p0 = 1
                }
                _ = fi17
            }
            _ = fi11
        }
        var v18 uint64 = p2
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}
func genFn1(p0 int8, p1 int8) uint16 {
    p0 = p0
    var v19 float32 = ((0.1 - 3.14159) + 0.1)
    switch ("a" + ("foo bar" + "")) {
    case "hello":
        p1 = 0
    case "":
        var v20 bool = (!true)
        _ = v20
    case "foo bar":
        fmt.Println(p1, v19)
    default:
        p1 = (p0 & p0)
    }
    v21 := []int{int(p0), -9223372036854775808}
    _ = p0
    _ = p1
    _ = v19
    _ = v21
    return 2
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = 42
    var v3 bool = true
    var v4 int8 = 42
    var v5 float32 = (1.5 * (3.14159 * (0.0 * float32(-2.25))))
    var v22 int64 = 1
    v1 = v1
    v23 := GenS1{F0: 0, F1: 737, F2: 128}
    var v24 int64 = (v22 >> ((18446744073709551615) & 63))
    var v25 int64 = (v22 & v22)
    v25 = 0
    for fi26 := 0; fi26 < 19; fi26++ {
        v27 := GenS1{F0: 192, F1: 4294967295, F2: (fi26 << ((255) & 63))}
        _ = fi26
        _ = v27
    }
    var v28 float64 = 3.14159
    if ((1 < 128) || (0 >= 4294967295)) {
        var v29 int = genFn0(558, 2, 42)
        v30 := []int{}
        var v31 float32 = 0.1
        _ = v29
        _ = v30
        _ = v31
    }
    var v32 string = ((("\n\t" + v1) + (("hello" + v1) + v1)) + v1)
    var v33 bool = v3
    v34 := []int{42, 2147483647, genFn0(42, 0, 966), -1}
    if len(v34) >= 2 {
        v34 = v34[1:len(v34)]
    }
    _ = len(v34)
    var v35 int8 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
