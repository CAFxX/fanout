package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 string; F1 int32; F2 int }
type GenS4 struct { F0 string; F1 int32; F2 uint8 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS8 struct { F0 int64; F1 string }
func genFn0(p0 int, p1 bool) uint32 {
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int32 = -1
        var v8 uint8 = 0
        _ = fi6
        _ = v7
        _ = v8
    }
    switch "a" {
    case "a":
        p1 = (1e10 < 1e10)
    case "foo bar":
        p0 = (-(p0))
    case "\n\t":
        var v9 bool = true
        _ = v9
    default:
        for fi10 := 0; fi10 < 1; fi10++ {
            if (("" + "hello") <= "") {
                p1 = p1
                var v11 float64 = 1e10
                v11 += (3.14159 * float64((3.14159 - 0.1)))
                _ = v11
            } else {
                var v12 uint64 = 18446744073709551615
                v12 = v12
                _ = v12
            }
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 int8 = 0
    var v2 float32 = 0.0
    var v3 int64 = -1
    v4 := make(map[uint32]bool)
    var v5 int64 = (-(v3))
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := []int{(fi13 / ((127) | 1)), (fi13 * (fi13 & (fi13 & int(v1)))), fi13}
        var v15 uint = 0
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v5, v5, v5)
    v16 := GenS1{F0: 0, F1: 1}
    v17 := GenS2{F0: 32768, F1: -821}
    if (v2 < ((3.14159 / (v2 + float32(-2.25))) + 0.0)) {
        var v18 uint = 2
        var v19 bool = false
        _ = v18
        _ = v19
    }
    v3 = (v5 ^ (v3 ^ v5))
    if false {
        if (false || false) {
            switch "hello" {
            case "a":
                fmt.Println(v5, v3)
            default:
                fmt.Println(v2)
            }
            switch ("" + (("\n\t" + "a") + "foo bar")) {
            case "":
                if (42 < 207) {
                    v3 += v3
                }
            default:
                if false {
                    _ = len(v4)
                    for fi20 := 0; fi20 < 11; fi20++ {
                        v21 := []int{(fi20 | fi20), int(v5), int(v1), 127}
                        if (false || (1.5 > v2)) {
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            for fi22 := 0; fi22 < 7; fi22++ {
                                v23 := GenS3{F0: "a", F1: 0, F2: fi20}
                                switch fi22 {
                                case 66:
                                    var v24 int = fi20
                                    _ = v24
                                }
                                _ = fi22
                                _ = v23
                            }
                        }
                        _ = fi20
                        _ = v21
                    }
                    v25 := v4[genFn0(-986, ((true && false) && (721 >= 2147483647)))]
                    _ = v25
                }
            }
        }
        v26 := GenS4{F0: ("foo bar" + "hello"), F1: 0, F2: 248}
        _ = v26
    }
    v27 := GenS5{F0: -225, F1: "a"}
    var v28 int32 = -2147483648
    v29 := GenS4{F0: ("\n\t" + "\n\t"), F1: v28, F2: 127}
    v2 -= v2
    for fi30 := 0; fi30 < 9; fi30++ {
        v28 = 2147483647
        _ = fi30
    }
    v2 += v2
    for fi31 := 0; fi31 < 16; fi31++ {
        if ((v1 & v1) == (^(v1))) {
            switch (fi31 << ((1) & 63)) {
            case 128:
                fmt.Println(v28, v1)
            case 42:
                switch fi31 {
                case 42:
                    var v32 float32 = 3.14159
                    _ = v32
                case -2147483648:
                    for fi33 := 0; fi33 < 11; fi33++ {
                        v1 -= v1
                        _ = fi33
                    }
                default:
                    for fi34 := 0; fi34 < 8; fi34++ {
                        var v35 bool = true
                        _ = fi34
                        _ = v35
                    }
                }
            case -9223372036854775808:
                var v36 uint64 = 306
                _ = v36
            }
            if true {
                v37 := []int{32768, (fi31 >> ((42) & 63)), (fi31 % ((fi31) | 1)), int(v1)}
                v38 := GenS6{F0: 1}
                _ = v37
                _ = v38
            }
            v4[248] = (42 == 1)
        } else {
            v39 := GenS7{F0: 343, F1: 65535, F2: 160}
            for fi40 := 0; fi40 < 15; fi40++ {
                v41 := GenS8{F0: (v3 - v5), F1: "\n\t"}
                v42 := []int{128, fi31, fi40, 1}
                delete(v4, 623)
                v2 += float32(1.5)
                if ((65535 > 2) || (v28 >= v28)) {
                    var v43 float32 = v2
                    _ = v43
                }
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = v39
        }
        _ = fi31
    }
    v3 += (v3 * v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
