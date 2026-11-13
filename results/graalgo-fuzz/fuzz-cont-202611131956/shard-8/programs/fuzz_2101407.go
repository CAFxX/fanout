package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint32; F2 bool }
type GenS4 struct { F0 int }
func genFn0(p0 int16, p1 int32, p2 bool) int {
    p2 = (p0 != p0)
    var v6 int16 = p0
    fmt.Println(p0, p1, p1)
    var v7 int = -913
    p0 = (p0 | -618)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 uint64) uint64 {
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS1{F0: 822, F1: 1}
            var v11 uint16 = 42
            if ((32767 > 171) && ((fi8 & 128) == (fi9 + -2147483648))) {
                v11 = 42
                p0 = 940
                var v12 string = ("foo bar" + "foo bar")
                _ = v12
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        if false {
            if false {
                v13 := GenS2{F0: 0, F1: 42}
                for fi14 := 0; fi14 < 20; fi14++ {
                    if true {
                        v15 := []int{int(p0), genFn0(1000, 0, false)}
                        _ = v15
                    }
                    p0 += (^(p0))
                    _ = fi14
                }
                var v16 uint32 = 0
                _ = v13
                _ = v16
            }
            switch ("foo bar" + "\n\t") {
            case "":
                if (!false) {
                    var v17 string = ("a" + ("a" + "\n\t"))
                    _ = v17
                }
            default:
                var v18 int16 = 22
                _ = v18
            }
        } else {
            var v19 uint32 = 2147483648
            if true {
                switch "" {
                case "foo bar":
                    p0 = p0
                default:
                    v20 := []int{(fi8 & (-(fi8)))}
                    _ = v20
                }
            } else {
                if (1 >= 1) {
                    v21 := GenS1{F0: (fi8 * fi8), F1: 200}
                    v22 := []int{fi8, genFn0(32767, -439, false)}
                    _ = v21
                    _ = v22
                } else {
                    v23 := []int{128, (fi8 >> ((0) & 63))}
                    v24 := GenS2{F0: -32768, F1: 32767}
                    v25 := GenS1{F0: (fi8 % (((fi8 / ((genFn0(1, 1, (983 != 1))) | 1))) | 1)), F1: 127}
                    _ = v23
                    _ = v24
                    _ = v25
                }
                if true {
                    p0 += (p0 << ((255) & 63))
                    if ((^(fi8)) > 4294967295) {
                        p0 = (^(p0))
                    }
                    fmt.Println(v19, v19, p0)
                } else {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        if true {
                            v27 := []int{}
                            if true {
                                var v28 uint64 = p0
                                v19 += (v19 - v19)
                                var v29 int64 = -9223372036854775808
                                _ = v28
                                _ = v29
                            } else {
                                v30 := GenS2{F0: -99, F1: 451}
                                _ = v30
                            }
                            var v31 uint32 = (v19 >> ((18446744073709551615) & 63))
                            var v32 uint16 = uint16(p0)
                            _ = v27
                            _ = v31
                            _ = v32
                        }
                        _ = fi26
                    }
                }
            }
            var v33 uint8 = 42
            _ = v19
            _ = v33
        }
        fmt.Println(p0)
        _ = fi8
    }
    p0 -= p0
    v34 := GenS3{F0: 9223372036854775807, F1: 4294967295, F2: false}
    var v35 int16 = 1000
    var v36 int = 0
    _ = p0
    _ = v34
    _ = v35
    _ = v36
    return p0
}

func main() {
    v1 := []uint32{2, 2147483648, 816, 1}
    var v2 int64 = 1
    var v3 int64 = 0
    var v4 uint16 = 32767
    var v5 uint16 = 15
    var v37 string = ((("a" + "foo bar") + "a") + "a")
    v38 := GenS4{F0: int(v3)}
    v39 := GenS3{F0: (v2 * v3), F1: 421, F2: true}
    var v40 int32 = int32(v3)
    switch (v37 + v37) {
    case "a":
        v4 -= (v4 + v5)
    case "":
        v1 = append(v1, 65535)
    default:
        switch 0 {
        case -164:
            var v41 string = v37
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 15; fi42++ {
        var v43 int16 = 389
        _ = fi42
        _ = v43
    }
    var v44 uint32 = 2147483648
    var v45 uint8 = 2
    var v46 bool = (127 > 1)
    v44 = 456
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
