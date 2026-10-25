package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int }
type GenS2 struct { F0 uint16; F1 uint32; F2 int }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int8; F2 string }
func genFn0(p0 bool, p1 uint, p2 uint16) int {
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int64 = -1
            _ = fi8
            _ = v9
        }
        var v10 uint = p1
        _ = fi7
        _ = v10
    }
    var v11 uint32 = 2147483648
    if (-352 != -625) {
        var v12 string = "foo bar"
        p2 += p2
        p0 = (!true)
        _ = v12
    }
    v13 := []int{4294967295}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v13
    return 128
}

func main() {
    v1 := []string{}
    var v2 float64 = -2.25
    var v3 int = 9223372036854775807
    var v4 uint16 = 42
    var v5 int16 = 1
    v6 := make(map[int]string)
    v14 := GenS1{F0: 52, F1: 255, F2: 4294967295}
    v15 := GenS2{F0: v4, F1: 1, F2: (v3 * (v3 >> ((18446744073709551615) & 63)))}
    var v16 float64 = 0.1
    delete(v6, -696)
    _ = len(v6)
    v17 := v6[(v3 * (^(v3)))]
    v5 -= (v5 ^ (v5 | v5))
    v18 := GenS3{F0: 211, F1: 1}
    for fi19 := 0; fi19 < 14; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            if (((true && (!("\n\t" >= "foo bar"))) && (42 >= 0)) && (!(v17 != (((v17 + (v17 + ("a" + v17))) + "a") + "hello")))) {
                v5 += v5
                var v21 float64 = 0.0
                v22 := GenS2{F0: v4, F1: 4294967295, F2: fi20}
                switch (fi19 >> ((354) & 63)) {
                case 2147483647:
                    if (!(!(!(-1 < 9223372036854775807)))) {
                        var v23 uint16 = (v4 ^ 255)
                        v24 := GenS1{F0: 40, F1: 0, F2: 1}
                        v25 := GenS4{F0: v23, F1: 9223372036854775807, F2: 42}
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                case 128:
                    var v26 int16 = v5
                    _ = v26
                case 1:
                    delete(v6, 127)
                }
                var v27 uint = 0
                _ = v21
                _ = v22
                _ = v27
            } else {
                v28 := GenS4{F0: v4, F1: 9223372036854775807, F2: 553}
                if (!true) {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    var v29 int8 = 1
                    _ = v29
                } else {
                    v3 = (^(fi20))
                    var v30 uint = 497
                    switch "hello" {
                    case "\n\t":
                        v31 := GenS3{F0: int32(v5), F1: 0}
                        v3 -= v3
                        var v32 uint64 = 18446744073709551615
                        delete(v6, (fi20 | genFn0((!false), (v30 | (v30 * v30)), v4)))
                        _ = v31
                        _ = v32
                    case "hello":
                        v33 := GenS2{F0: v4, F1: 65535, F2: 2147483647}
                        v3 += fi19
                        fmt.Println(v30, v4, v16)
                        v16 = float64(1e10)
                        _ = v33
                    }
                    if true {
                        var v34 int64 = -1
                        _ = v34
                    }
                    v35 := GenS5{F0: 42, F1: 1, F2: v17}
                    _ = v30
                    _ = v35
                }
                _ = v28
            }
            _ = fi20
        }
        if false {
            fmt.Println(v5)
        } else {
            for fi36 := 0; fi36 < 6; fi36++ {
                if false {
                    var v37 int = (-(v3))
                    _ = v37
                }
                _ = fi36
            }
        }
        _ = fi19
    }
    v38 := GenS5{F0: 4, F1: 1, F2: v17}
    var v39 uint16 = v4
    var v40 uint8 = 0
    v39 = (v4 & v4)
    var v41 int = (v3 % (((^(v3))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
