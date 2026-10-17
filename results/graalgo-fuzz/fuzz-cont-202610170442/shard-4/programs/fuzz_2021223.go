package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int) uint {
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{fi6, p0}
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(fi6, p0)
            var v9 uint = 18446744073709551615
            _ = fi8
            _ = v9
        }
        var v10 int8 = 127
        _ = fi6
        _ = v10
        _ = v7
    }
    var v11 int8 = 0
    var v12 uint8 = 128
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := []int{42, p0, p0, fi13}
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = v11
    _ = v12
    return 1
}
func genFn1(p0 uint) int64 {
    p0 = p0
    var v15 int8 = 42
    v16 := []int{}
    fmt.Println(v15, v15)
    _ = p0
    _ = v15
    _ = v16
    return 9223372036854775807
}

func main() {
    var v1 int = -899
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    var v4 uint8 = 45
    var v5 uint = 255
    if (1.5 >= 1e10) {
        v5 = v5
        var v17 string = "\n\t"
        _ = v17
    } else {
        var v18 float32 = 3.14159
        fmt.Println(v18, v1)
        _ = v18
    }
    v4 += (v4 + (^(v4)))
    v4 += v4
    if ((((3.14159 > (1.5 / float32((1e10 * (3.14159 + (((1.5 + (0.0 + 1.5)) / 1e10) - (0.1 + (1e10 / 1e10)))))))) && (1 > -128)) || true) && true) {
        var v19 string = ""
        _ = v19
    }
    switch ("" + "a") {
    case "a":
        v20 := GenS1{F0: v5, F1: 255, F2: 50}
        _ = v20
    case "foo bar":
        for fi21 := 0; fi21 < 5; fi21++ {
            for fi22 := 0; fi22 < 10; fi22++ {
                switch "foo bar" {
                case "a":
                    v1 = 2147483647
                }
                _ = fi22
            }
            var v23 int8 = 100
            v23 -= v23
            _ = fi21
            _ = v23
        }
    case "\n\t":
        v5 = (v5 | 18446744073709551615)
    default:
        var v24 uint8 = 1
        _ = v24
    }
    v25 := GenS1{F0: (^(v5)), F1: v4, F2: 48}
    if (!(928 <= 0)) {
        var v26 bool = (((0.0 < 3.14159) || (false && true)) || false)
        _ = v26
    } else {
        if true {
            v1 -= v1
        }
        for fi27 := 0; fi27 < 18; fi27++ {
            if (!false) {
                if false {
                    switch -586 {
                    case 254:
                        v28 := GenS1{F0: (v5 / (((^(v5))) | 1)), F1: (v4 >> (((^(v5))) & 63)), F2: -31}
                        _ = v28
                    case 2147483647:
                        var v29 uint8 = v4
                        _ = v29
                    default:
                        v30 := GenS2{F0: 42, F1: 42, F2: (v4 % (((v4 % ((139) | 1))) | 1))}
                        _ = v30
                    }
                } else {
                    fmt.Println(v5, v5, v5)
                }
                switch (fi27 + 275) {
                case 405:
                    var v31 int8 = 106
                    _ = v31
                default:
                    v32 := []int{v1, (v1 << (((v5 + (v5 ^ (v5 - (^(v5)))))) & 63)), -2147483648, v1}
                    _ = v32
                }
                var v33 uint8 = 147
                _ = v33
            } else {
                var v34 uint8 = v4
                _ = v34
            }
            if true {
                var v35 bool = (91 >= 42)
                v36 := v3[1]
                v37 := GenS1{F0: (v5 % (((v5 << ((genFn0((fi27 ^ 370))) & 63))) | 1)), F1: v4, F2: 103}
                _ = v35
                _ = v36
                _ = v37
            }
            _ = fi27
        }
        if (!(940 > 2)) {
            for fi38 := 0; fi38 < 7; fi38++ {
                for fi39 := 0; fi39 < 17; fi39++ {
                    var v40 uint16 = 882
                    _ = fi39
                    _ = v40
                }
                _ = fi38
            }
            v5 -= (v5 * (v5 / (((v5 >> ((genFn0((v1 & 2147483647))) & 63))) | 1)))
            var v41 int = (v1 ^ (v1 * v1))
            _ = v41
        } else {
            var v42 int8 = 127
            v43 := []int{}
            _ = v42
            _ = v43
        }
    }
    v44 := GenS3{F0: (v4 % ((v4) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
