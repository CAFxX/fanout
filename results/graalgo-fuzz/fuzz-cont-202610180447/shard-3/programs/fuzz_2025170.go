package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 int8 }

func main() {
    var v1 string = ""
    var v2 uint16 = 42
    var v3 uint8 = 127
    v4 := make(map[uint32]int)
    var v5 uint32 = 845
    _ = len(v4)
    if true {
        v4[(v5 << ((255) & 63))] = 32768
        if (1e10 == float64((((0.0 / 0.1) * 3.14159) + (0.1 - -2.25)))) {
            var v6 int = -9223372036854775808
            _ = v6
        } else {
            var v7 int = -2147483648
            v5 = (^(v5))
            for fi8 := 0; fi8 < 15; fi8++ {
                v3 = (v3 | (^(v3)))
                fmt.Println(fi8, v1, v5)
                delete(v4, 2)
                _ = fi8
            }
            _ = v7
        }
    } else {
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: (!(false && (!false))), F1: 827}
            _ = fi9
            _ = v10
        }
    }
    switch 32768 {
    case 4294967295:
        v4[(v5 / (((^(v5))) | 1))] = 32768
    case 128:
        if true {
            if (65535 == 18446744073709551615) {
                if true {
                    v1 = (v1 + "")
                    var v11 uint8 = (v3 * v3)
                    v12 := []int{-809, 4294967295}
                    _ = v11
                    _ = v12
                }
                for fi13 := 0; fi13 < 19; fi13++ {
                    v4[(v5 << ((630) & 63))] = (fi13 * (fi13 + (^(fi13))))
                    var v14 int32 = -1
                    v3 = (v3 + 2)
                    _ = fi13
                    _ = v14
                }
                if false {
                    v3 += 255
                }
            } else {
                switch 2147483647 {
                case 128:
                    v3 -= v3
                default:
                    if ((v3 >> ((1) & 63)) == (v3 ^ v3)) {
                        fmt.Println(v5, v5, v1)
                        _ = len(v4)
                        fmt.Println(v2)
                    }
                }
                for fi15 := 0; fi15 < 3; fi15++ {
                    var v16 uint32 = 272
                    _ = fi15
                    _ = v16
                }
            }
            if false {
                for fi17 := 0; fi17 < 11; fi17++ {
                    fmt.Println(v1, v5, v2)
                    var v18 string = "a"
                    _ = fi17
                    _ = v18
                }
            }
        } else {
            v19 := GenS1{F0: (535 <= 18446744073709551615), F1: uint64(v2)}
            for fi20 := 0; fi20 < 19; fi20++ {
                switch v1 {
                case "\n\t":
                    v21 := GenS2{F0: v2, F1: 0}
                    _ = v21
                }
                var v22 uint32 = v5
                v23 := []int{fi20, (fi20 - 588)}
                _ = fi20
                _ = v22
                _ = v23
            }
            _ = v19
        }
    case -442:
        v1 = ("foo bar" + ((v1 + v1) + v1))
    default:
        fmt.Println(v5, v5, v3)
    }
    if false {
        v3 = (^(v3))
    } else {
        _ = len(v4)
        var v24 bool = (!false)
        _ = v24
    }
    if true {
        v25 := []int{2147483647}
        v26 := v4[v5]
        _ = v25
        _ = v26
    } else {
        v27 := []int{4294967295, 32767, 128, -9223372036854775808}
        _ = v27
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        if (1e10 <= (3.14159 * 0.0)) {
            var v29 uint8 = (v3 | (v3 >> ((42) & 63)))
            fmt.Println(v5, fi28)
            v30 := []int{0, fi28, (fi28 - (fi28 << ((0) & 63))), 42}
            _ = v29
            _ = v30
        }
        _ = fi28
    }
    v31 := GenS1{F0: true, F1: 4294967295}
    v1 = v1
    if (((1e10 / (1e10 - 0.0)) * 3.14159) > -2.25) {
        v32 := GenS2{F0: v2, F1: 94}
        _ = v32
    }
    v3 = (^(v3))
    v33 := GenS3{F0: "hello"}
    v34 := GenS4{F0: -32768, F1: 100}
    for fi35 := 0; fi35 < 1; fi35++ {
        var v36 float32 = 1e10
        if (358 == -246) {
            fmt.Println(fi35, v2, v3)
        }
        _ = fi35
        _ = v36
    }
    v5 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
