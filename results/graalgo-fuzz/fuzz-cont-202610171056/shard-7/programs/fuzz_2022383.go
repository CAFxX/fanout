package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS3 struct { F0 int16; F1 int32; F2 int16 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int, p1 int) bool {
    p1 = -2147483648
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(p1, fi7, p0)
        p1 = p1
        _ = fi7
    }
    switch "foo bar" {
    case "foo bar":
        p1 -= p0
    case "hello":
        if (1 > 44) {
            var v8 uint16 = 32767
            if false {
                v9 := GenS1{F0: -299, F1: -1, F2: 1}
                var v10 uint = 18446744073709551615
                var v11 uint8 = 100
                _ = v10
                _ = v11
                _ = v9
            } else {
                if (!((p1 / (((p1 | (-(p1)))) | 1)) > 631)) {
                    v12 := []int{4294967295, (p1 | (p0 & (p0 - 0)))}
                    switch "" {
                    case "":
                        v13 := GenS2{F0: 37, F1: 990, F2: 226}
                        _ = v13
                    default:
                        var v14 int64 = 9223372036854775807
                        _ = v14
                    }
                    _ = v12
                } else {
                    v15 := []int{p0, -9223372036854775808}
                    p0 = p0
                    v15 = append(v15, (-(p0)))
                    _ = v15
                }
                var v16 int64 = -1
                var v17 int = p1
                _ = v16
                _ = v17
            }
            var v18 uint16 = v8
            _ = v18
            _ = v8
        } else {
            switch ("" + "a") {
            case "\n\t":
                v19 := []int{-9223372036854775808, (p1 << ((4294967295) & 63)), p0}
                _ = v19
            }
        }
    default:
        p1 = (p1 % (((p1 >> ((1) & 63))) | 1))
    }
    var v20 int64 = 799
    _ = p0
    _ = p1
    _ = v20
    return (v20 <= (v20 - -9223372036854775808))
}
func genFn1(p0 uint32, p1 int32) int {
    var v21 uint16 = 32767
    switch ("" + "a") {
    case "":
        if true {
            fmt.Println(v21, v21, v21)
            if (234 > 2) {
                p1 += p1
                var v22 int = int(p0)
                _ = v22
            } else {
                v23 := GenS3{F0: 615, F1: (p1 << ((42) & 63)), F2: 42}
                v24 := GenS4{F0: -745, F1: "\n\t"}
                _ = v23
                _ = v24
            }
            v25 := GenS4{F0: 32767, F1: "hello"}
            _ = v25
        }
    case "hello":
        var v26 float64 = -2.25
        _ = v26
    case "a":
        v27 := []int{-1}
        _ = v27
    default:
        switch 2147483647 {
        case -1:
            p0 -= p0
        case 4294967295:
            var v28 uint16 = (v21 & v21)
            _ = v28
        case 32768:
            var v29 int32 = (p1 - (-(p1)))
            _ = v29
        default:
            var v30 int8 = 42
            _ = v30
        }
    }
    _ = p0
    _ = p1
    _ = v21
    return 9223372036854775807
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint16 = 32767
    v3 := []uint32{1, 55, 4294967295, 2147483647}
    v4 := make(map[string]bool)
    v5 := []uint32{2, 2147483648}
    v6 := []string{}
    v2 += 0
    if ((32767 != 0) || ((true || true) || false)) {
        for fi31 := 0; fi31 < 11; fi31++ {
            v1 = (-(v1))
            for fi32 := 0; fi32 < 1; fi32++ {
                v2 = v2
                v33 := GenS4{F0: (fi31 % ((fi31) | 1)), F1: "foo bar"}
                _ = fi32
                _ = v33
            }
            v2 = 255
            _ = fi31
        }
    }
    var v34 int32 = 2147483647
    var v35 uint32 = 128
    v36 := GenS5{F0: (v35 >> ((22) & 63)), F1: 135}
    v37 := v4[("hello" + ("a" + ("foo bar" + "foo bar")))]
    v38 := GenS5{F0: 829, F1: 110}
    v35 = v35
    v39 := GenS2{F0: 128, F1: 0, F2: 113}
    for fi40 := 0; fi40 < 17; fi40++ {
        for fi41 := 0; fi41 < 8; fi41++ {
            v2 -= 189
            _ = fi41
        }
        var v42 float32 = ((1e10 - 1.5) - 0.1)
        for fi43 := 0; fi43 < 15; fi43++ {
            var v44 float64 = (3.14159 - (3.14159 / ((((0.0 + ((3.14159 * 1e10) / 3.14159)) * 1e10) + 0.0) - 1.5)))
            _ = fi43
            _ = v44
        }
        _ = fi40
        _ = v42
    }
    var v45 bool = (v37 || v37)
    var v46 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
