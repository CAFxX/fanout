package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint; F1 int8; F2 int32 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint32, p1 int, p2 uint16) uint {
    fmt.Println(p0)
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 string) uint64 {
    var v5 bool = (false || false)
    v6 := []int{}
    var v7 float64 = 0.0
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint16 = 65535
    var v2 string = ("foo bar" + "foo bar")
    var v3 int = 654
    v4 := make(map[int]string)
    var v8 uint = 1
    var v9 uint = (v8 << (((^(v8))) & 63))
    v10 := []int{v3, 32767}
    switch "a" {
    case "\n\t":
        if true {
            if true {
                if true {
                    var v11 int8 = -1
                    _ = v11
                } else {
                    v12 := GenS1{F0: -2147483648, F1: (^(v3)), F2: 127}
                    _ = v12
                }
                var v13 int8 = -51
                _ = v13
            }
            if (true && ((42 < 42) || (v8 < 4294967295))) {
                _ = len(v10)
                v14 := GenS1{F0: -1, F1: (-(v3)), F2: -128}
                v1 -= v1
                _ = v14
            } else {
                v15 := GenS2{F0: (v1 % ((1) | 1)), F1: 0}
                v1 = (^(v1))
                _ = v15
            }
            if false {
                if true {
                    v1 += 65535
                    v16 := GenS1{F0: -18, F1: (v3 * (v3 << (((v9 + 1)) & 63))), F2: 42}
                    _ = v16
                }
                var v17 float32 = ((0.1 / 1.0) - 1.5)
                for fi18 := 0; fi18 < 20; fi18++ {
                    if false {
                        var v19 float32 = v17
                        var v20 uint32 = 2
                        var v21 int = (fi18 >> (((v8 % (((v9 << ((1) & 63))) | 1))) & 63))
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    v22 := v4[-37]
                    _ = fi18
                    _ = v22
                }
                _ = v17
            } else {
                var v23 uint16 = (v1 >> ((v9) & 63))
                var v24 uint = v9
                _ = v23
                _ = v24
            }
        } else {
            var v25 uint = (v8 * genFn0(2, -159, v1))
            fmt.Println(v9, v1, v1)
            delete(v4, (v3 ^ 703))
            _ = v25
        }
    }
    v8 += v8
    for fi26 := 0; fi26 < 20; fi26++ {
        v27 := GenS3{F0: (^(v9)), F1: genFn0(0, (fi26 & v3), v1)}
        v8 += v8
        var v28 int32 = 984
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 uint32 = 1
    v30 := GenS3{F0: (v9 ^ genFn0(v29, 42, v1)), F1: genFn0((v29 / ((2147483647) | 1)), (v3 >> (((v8 * uint(v29))) & 63)), (v1 >> (((^(v9))) & 63)))}
    v31 := GenS4{F0: (v9 * v8), F1: -111, F2: 2147483647}
    fmt.Println(v29, v9)
    var v32 uint32 = 2
    switch (("a" + v2) + "a") {
    case "a":
        var v33 uint16 = (^(v1))
        _ = v33
    default:
        v32 += 203
    }
    if false {
        if true {
            fmt.Println(v8, v1, v9)
            v34 := GenS5{F0: (185 < 213)}
            _ = v34
        } else {
            v35 := GenS1{F0: 0, F1: v3, F2: -128}
            _ = v35
        }
        v9 = (v9 % ((v9) | 1))
    }
    if (true && (((((((v2 + (v2 + ("a" + v2))) + v2) + "foo bar") + v2) + (v2 + (v2 + (("hello" + ("foo bar" + (v2 + "\n\t"))) + (v2 + v2))))) + v2) > v2)) {
        v29 -= v32
        for fi36 := 0; fi36 < 3; fi36++ {
            if false {
                var v37 uint16 = (v1 * v1)
                _ = v37
            }
            _ = fi36
        }
    } else {
        fmt.Println(v2, v29)
    }
    v38 := GenS2{F0: v1, F1: 255}
    delete(v4, v3)
    if ((v8 != (v9 % ((v9) | 1))) || (1e10 <= -2.25)) {
        v1 -= 32767
        for fi39 := 0; fi39 < 4; fi39++ {
            switch v2 {
            case "\n\t":
                switch fi39 {
                case -9223372036854775808:
                    var v40 int16 = 0
                    _ = v40
                case 9223372036854775807:
                    fmt.Println(v9, v8, v29)
                case 32768:
                    delete(v4, 0)
                default:
                    var v41 bool = true
                    _ = v41
                }
            case "":
                for fi42 := 0; fi42 < 10; fi42++ {
                    var v43 uint = genFn0((v29 + (v29 ^ v32)), 284, v1)
                    fmt.Println(v9, v2, v29)
                    var v44 bool = false
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            default:
                var v45 float64 = 0.1
                _ = v45
            }
            _ = fi39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v38
    _ = v4
    _ = v8
    _ = v9
}
