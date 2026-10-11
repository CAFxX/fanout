package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int32; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint16) uint {
    for fi4 := 0; fi4 < 1; fi4++ {
        p0 -= (p0 - p0)
        v5 := GenS1{F0: 849}
        _ = fi4
        _ = v5
    }
    if (32767 <= 9223372036854775807) {
        if false {
            for fi6 := 0; fi6 < 10; fi6++ {
                fmt.Println(fi6)
                switch fi6 {
                case 32767:
                    p0 -= 442
                }
                if ((0.0 * 3.14159) <= 0.1) {
                    var v7 float64 = 3.14159
                    _ = v7
                }
                _ = fi6
            }
            p0 += p0
            fmt.Println(p0)
        } else {
            switch 2147483647 {
            case 608:
                switch 1 {
                case -1:
                    for fi8 := 0; fi8 < 14; fi8++ {
                        switch "foo bar" {
                        case "\n\t":
                            for fi9 := 0; fi9 < 6; fi9++ {
                                v10 := GenS1{F0: -367}
                                switch (fi8 >> ((42) & 63)) {
                                case 128:
                                    var v11 string = ((("hello" + "") + "\n\t") + "foo bar")
                                    _ = v11
                                case 201:
                                    var v12 uint16 = p0
                                    _ = v12
                                default:
                                    var v13 int32 = -2147483648
                                    _ = v13
                                }
                                _ = fi9
                                _ = v10
                            }
                        case "hello":
                            p0 -= p0
                        case "":
                            p0 = p0
                        }
                        _ = fi8
                    }
                case 9223372036854775807:
                    var v14 uint8 = 0
                    _ = v14
                }
            }
        }
    } else {
        fmt.Println(p0)
    }
    p0 -= p0
    if true {
        if false {
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS1{F0: 548}
                p0 -= p0
                _ = fi15
                _ = v16
            }
            if false {
                if false {
                    v17 := []int{32767, 0}
                    p0 += (p0 | (p0 << ((2) & 63)))
                    p0 += 1
                    _ = v17
                } else {
                    v18 := GenS2{F0: 18446744073709551615, F1: 963, F2: "a"}
                    v19 := []int{0, 1, 2147483647}
                    _ = v18
                    _ = v19
                }
            } else {
                switch "" {
                case "a":
                    if (p0 == (p0 % (((^(p0))) | 1))) {
                        v20 := GenS2{F0: 1, F1: 0, F2: ((("\n\t" + "hello") + "a") + "hello")}
                        _ = v20
                    }
                default:
                    for fi21 := 0; fi21 < 19; fi21++ {
                        p0 = (p0 & 65535)
                        _ = fi21
                    }
                }
                v22 := []int{}
                _ = v22
            }
            if (int(p0) <= -954) {
                switch 42 {
                case -683:
                    for fi23 := 0; fi23 < 20; fi23++ {
                        fmt.Println(p0, p0)
                        var v24 uint32 = 1
                        _ = fi23
                        _ = v24
                    }
                case 9223372036854775807:
                    fmt.Println(p0)
                }
            } else {
                v25 := GenS3{F0: 937}
                _ = v25
            }
        }
        var v26 float32 = 0.1
        v27 := GenS4{F0: 468}
        _ = v26
        _ = v27
    }
    fmt.Println(p0)
    _ = p0
    return 42
}
func genFn1() bool {
    for fi28 := 0; fi28 < 1; fi28++ {
        v29 := []int{}
        _ = fi28
        _ = v29
    }
    var v30 uint64 = 647
    v31 := []int{-382, 127}
    if (0 > 42) {
        var v32 int16 = 42
        _ = v32
    }
    _ = v30
    _ = v31
    return ((42 > 65535) || (0.1 <= (float64(1e10) + ((1e10 - 0.1) / 1.5))))
}

func main() {
    var v1 uint64 = 13
    var v2 string = "hello"
    var v3 int = 32768
    v33 := GenS3{F0: 2}
    for fi34 := 0; fi34 < 12; fi34++ {
        v1 = (^(v1))
        v1 += (v1 % ((42) | 1))
        fmt.Println(v2)
        _ = fi34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        switch 4294967295 {
        case 634:
            var v36 bool = (!(false && false))
            _ = v36
        case 42:
            v37 := []int{(fi35 * (-(fi35)))}
            _ = v37
        }
        var v38 bool = (-1 >= 1)
        _ = fi35
        _ = v38
    }
    v2 = ("a" + v2)
    var v39 int16 = 32767
    v39 += v39
    v40 := []int{2147483647, v3}
    v3 -= 752
    var v41 int8 = 100
    if false {
        v39 -= v39
    }
    v1 = v1
    var v42 int16 = v39
    if len(v40) >= 2 {
        v40 = v40[1:len(v40)]
    }
    if len(v40) > 0 {
        v40[0] = (v3 - (v3 - (v3 << ((255) & 63))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
