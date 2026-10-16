package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int8; F1 int8; F2 uint8 }
func genFn0(p0 uint, p1 int8, p2 uint32) uint8 {
    var v4 bool = (4294967295 > p2)
    v5 := []int{0, -829, 2147483647, 167}
    var v6 int16 = -32768
    v6 -= (v6 + v6)
    var v7 int16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 uint) uint {
    fmt.Println(p0)
    if (!true) {
        p0 -= p0
    }
    fmt.Println(p0)
    _ = p0
    return p0
}

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]string)
    var v3 uint = v1
    fmt.Println(v1, v3)
    fmt.Println(v3)
    if (65535 <= 2147483648) {
        fmt.Println(v1, v1)
        if true {
            v8 := GenS1{F0: ("" + "foo bar"), F1: 127, F2: "a"}
            _ = v8
        }
        for fi9 := 0; fi9 < 5; fi9++ {
            if false {
                if true {
                    v10 := v2[2147483648]
                    delete(v2, 66)
                    v10 = ("\n\t" + "a")
                    _ = v10
                }
                v1 = (^(v3))
                var v11 bool = (false && false)
                _ = v11
            }
            switch "\n\t" {
            case "hello":
                v3 = genFn1(v1)
            case "a":
                v1 -= (v1 / (((v3 | v1)) | 1))
            case "\n\t":
                switch fi9 {
                case 32768:
                    v12 := GenS2{F0: (fi9 / ((fi9) | 1))}
                    _ = v12
                default:
                    if true {
                        var v13 float32 = (3.14159 + 0.0)
                        v14 := GenS1{F0: "\n\t", F1: 196, F2: "foo bar"}
                        _ = v13
                        _ = v14
                    } else {
                        var v15 uint8 = 37
                        if true {
                            v16 := GenS2{F0: 32767}
                            var v17 string = "foo bar"
                            if (1 < -301) {
                                delete(v2, 0)
                            } else {
                                v18 := GenS1{F0: ("" + (((v17 + v17) + v17) + "foo bar")), F1: (v15 ^ genFn0(v1, 100, 42)), F2: ""}
                                _ = v18
                            }
                            _ = v16
                            _ = v17
                        }
                        _ = v15
                    }
                }
            }
            switch "hello" {
            case "":
                v19 := GenS1{F0: "a", F1: 128, F2: ("" + "\n\t")}
                _ = v19
            case "a":
                if true {
                    v20 := v2[2147483647]
                    v3 += (v3 | 782)
                    _ = v20
                }
            }
            _ = fi9
        }
    } else {
        switch 127 {
        case 127:
            switch -777 {
            case -213:
                v2[2147483648] = "a"
            case 885:
                v21 := GenS3{F0: 65535}
                _ = v21
            case -1:
                v22 := GenS4{F0: 32767, F1: 1}
                _ = v22
            default:
                v1 = v3
            }
        case 2147483647:
            for fi23 := 0; fi23 < 3; fi23++ {
                fmt.Println(fi23, fi23, v3)
                v1 = (v1 >> ((18446744073709551615) & 63))
                for fi24 := 0; fi24 < 5; fi24++ {
                    v1 += (v1 | v3)
                    fmt.Println(fi24, v1)
                    fmt.Println(fi23, v1)
                    _ = fi24
                }
                _ = fi23
            }
        case -619:
            var v25 string = ("a" + "a")
            _ = v25
        default:
            var v26 uint64 = 749
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        v28 := GenS2{F0: -2147483648}
        _ = fi27
        _ = v28
    }
    v29 := []int{32767, 32767}
    var v30 float64 = (1.5 - (3.14159 * float64(0.1)))
    for fi31 := 0; fi31 < 2; fi31++ {
        fmt.Println(v30, v3, v30)
        v32 := GenS5{F0: 100, F1: 1, F2: 2}
        v33 := v2[2147483648]
        _ = fi31
        _ = v32
        _ = v33
    }
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    v29 = append(v29, -924)
    switch -2147483648 {
    case 128:
        v1 = (v3 / ((1) | 1))
    default:
        for fi34 := 0; fi34 < 1; fi34++ {
            _ = len(v2)
            _ = fi34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
