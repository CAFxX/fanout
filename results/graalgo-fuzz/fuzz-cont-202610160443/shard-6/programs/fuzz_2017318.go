package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint16; F1 int64 }
func genFn0(p0 int64, p1 bool, p2 uint16) uint32 {
    v6 := GenS1{F0: 984, F1: (p2 - (p2 / (((p2 * p2)) | 1))), F2: -2147483648}
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            if p1 {
                if p1 {
                    switch "foo bar" {
                    case "foo bar":
                        v9 := GenS2{F0: 739, F1: 151}
                        _ = v9
                    case "\n\t":
                        v10 := GenS2{F0: 255, F1: 65535}
                        _ = v10
                    default:
                        v11 := []int{}
                        _ = v11
                    }
                    v12 := GenS1{F0: (fi7 << ((4294967295) & 63)), F1: 42, F2: 0}
                    _ = v12
                } else {
                    v13 := []int{(fi7 << ((18446744073709551615) & 63)), -2147483648, fi8}
                    _ = v13
                }
                p1 = p1
                fmt.Println(p2, p1)
            }
            var v14 float64 = float64(float32(-2.25))
            _ = fi8
            _ = v14
        }
        var v15 int32 = 42
        _ = fi7
        _ = v15
    }
    v16 := GenS3{F0: ""}
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(p2, p2)
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v6
    return 42
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]bool)
    var v3 bool = true
    var v4 float64 = (-2.25 + (((0.1 * 1e10) / (((1.5 - 1e10) + 0.0) * -2.25)) + 1e10))
    var v5 int64 = 792
    v3 = (v3 || v3)
    v4 = ((1.5 * 0.1) + float64(1.5))
    var v18 uint64 = 18446744073709551615
    fmt.Println(v3)
    v18 = v18
    v1 = append(v1, -2147483648)
    var v19 int = 2147483647
    if (v5 >= (v5 / ((-464) | 1))) {
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 string = ""
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 int64 = v5
                fmt.Println(fi22)
                v21 = (("\n\t" + (("a" + "") + v21)) + "hello")
                v19 += (fi22 << ((1) & 63))
                _ = fi22
                _ = v23
            }
            switch v21 {
            case "foo bar":
                v5 -= (v5 & v5)
            }
            _ = fi20
            _ = v21
        }
        if (v19 <= (v19 | (^(v19)))) {
            v1 = append(v1, v19)
            v1 = append(v1, (v19 & v19))
        } else {
            v24 := GenS1{F0: 42, F1: 1, F2: int32(v18)}
            for fi25 := 0; fi25 < 20; fi25++ {
                for fi26 := 0; fi26 < 19; fi26++ {
                    v27 := GenS1{F0: fi26, F1: 255, F2: -2147483648}
                    v1 = append(v1, 307)
                    var v28 string = "a"
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                if (1 > 0) {
                    fmt.Println(v18, v19)
                    fmt.Println(fi25)
                }
                _ = fi25
            }
            _ = v24
        }
        switch v19 {
        case 9223372036854775807:
            v29 := GenS3{F0: (("" + "a") + "foo bar")}
            _ = v29
        case 1:
            var v30 int = 9223372036854775807
            _ = v30
        case 42:
            switch 32768 {
            case 32768:
                switch ("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))) {
                case "foo bar":
                    delete(v2, 24)
                case "hello":
                    v31 := v2[2147483648]
                    _ = v31
                case "a":
                    v4 -= float64(((float32(v4) - 0.1) - -2.25))
                default:
                    if (((("foo bar" + "foo bar") + "") + "") <= (("a" + "hello") + "")) {
                        v32 := GenS1{F0: (v19 + v19), F1: 65535, F2: 1}
                        if v3 {
                            switch (("a" + "\n\t") + "foo bar") {
                            case "hello":
                                var v33 int8 = 0
                                _ = v33
                            case "a":
                                for fi34 := 0; fi34 < 1; fi34++ {
                                    var v35 string = "\n\t"
                                    v36 := v2[genFn0((v5 ^ v5), (((v5 / (((v5 & v5)) | 1)) >= v5) || true), 743)]
                                    v18 -= 0
                                    _ = fi34
                                    _ = v35
                                    _ = v36
                                }
                            }
                            _ = len(v1)
                        } else {
                            switch ("hello" + (("\n\t" + "") + "\n\t")) {
                            case "foo bar":
                                var v37 float32 = 1e10
                                _ = v37
                            case "":
                                if false {
                                    v18 = (v18 % (((v18 + 0)) | 1))
                                    v1 = append(v1, v19)
                                }
                            default:
                                v19 += -1
                            }
                            v38 := GenS4{F0: (v18 & v18), F1: uint(v18)}
                            _ = v38
                        }
                        _ = v32
                    }
                }
            case -60:
                v39 := GenS5{F0: 65535, F1: (v5 * 1)}
                _ = v39
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
