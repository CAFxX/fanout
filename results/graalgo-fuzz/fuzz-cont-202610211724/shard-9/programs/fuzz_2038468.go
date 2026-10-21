package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 string; F2 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint8) uint {
    var v4 string = ("foo bar" + "hello")
    var v5 int = -75
    fmt.Println(p0, p0, p0)
    _ = p0
    _ = v4
    _ = v5
    return 286
}
func genFn1() int8 {
    var v6 uint16 = 852
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(v6)
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 uint64 = 2
            if (0 > genFn0(uint8(v6))) {
                var v10 int16 = 32767
                fmt.Println(v6, v6)
                var v11 string = ""
                _ = v10
                _ = v11
            } else {
                if (1 != 128) {
                    var v12 int = fi7
                    _ = v12
                } else {
                    if false {
                        v9 = 985
                        switch fi8 {
                        case 127:
                            var v13 int32 = 0
                            _ = v13
                        default:
                            v14 := []int{fi8, -9223372036854775808, -9223372036854775808}
                            _ = v14
                        }
                        switch -2147483648 {
                        case -449:
                            v6 -= 222
                        }
                    }
                    var v15 uint16 = (v6 * (v6 * (v6 & v6)))
                    _ = v15
                }
            }
            v16 := []int{(fi7 * (fi8 | (fi8 - (fi8 * fi7)))), (fi8 << ((4294967295) & 63))}
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            _ = fi8
            _ = v16
            _ = v9
        }
        _ = fi7
    }
    v17 := GenS1{F0: genFn0(127)}
    v18 := []int{829}
    _ = v17
    _ = v18
    _ = v6
    return 1
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]bool)
    var v3 uint = 894
    var v19 int64 = 42
    switch "" {
    case "a":
        v1 = 100
    case "\n\t":
        if false {
            v20 := GenS2{F0: "a", F1: "\n\t", F2: genFn1()}
            v21 := GenS3{F0: -812}
            if false {
                fmt.Println(v1)
            }
            _ = v20
            _ = v21
        }
    }
    var v22 float32 = -2.25
    var v23 float64 = 0.0
    v24 := GenS4{F0: 4294967295, F1: 341}
    v25 := GenS5{F0: 4294967295}
    for fi26 := 0; fi26 < 19; fi26++ {
        v27 := GenS4{F0: 42, F1: 0}
        for fi28 := 0; fi28 < 17; fi28++ {
            if true {
                v29 := GenS6{F0: v19}
                v3 -= genFn0(174)
                v30 := GenS5{F0: 264}
                _ = v29
                _ = v30
            }
            v1 += -88
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    fmt.Println(v3, v3, v22)
    v23 = (3.14159 - v23)
    v22 -= 0.1
    v1 -= genFn1()
    v23 = (((v23 + 0.0) + v23) + (v23 - v23))
    switch "a" {
    case "a":
        var v31 uint64 = 18446744073709551615
        _ = v31
    case "\n\t":
        var v32 uint64 = 448
        _ = v32
    case "":
        fmt.Println(v1, v3)
    }
    v3 += (^(v3))
    var v33 uint64 = 221
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v33
}
