package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 uint8 = 29
    var v2 int32 = 1
    v3 := []uint32{376, 2147483648, 2147483648}
    v4 := make(map[int]uint32)
    v5 := make(map[int]int)
    v6 := GenS1{F0: 550, F1: 42, F2: 167}
    v1 = v1
    v7 := GenS1{F0: 2, F1: 1000, F2: 65535}
    for fi8 := 0; fi8 < 15; fi8++ {
        v2 -= -1
        v1 += (v1 << ((824) & 63))
        _ = fi8
    }
    v9 := GenS1{F0: 32767, F1: -178, F2: 929}
    var v10 int16 = 0
    var v11 int8 = 79
    for fi12 := 0; fi12 < 19; fi12++ {
        if ((!(false && false)) && true) {
            var v13 bool = (false || false)
            for fi14 := 0; fi14 < 18; fi14++ {
                if (v13 && v13) {
                    var v15 int32 = v2
                    switch int(v11) {
                    case 9223372036854775807:
                        fmt.Println(fi12, v11)
                    default:
                        var v16 string = ("a" + "hello")
                        _ = v16
                    }
                    _ = v15
                }
                switch fi12 {
                case -1:
                    switch fi12 {
                    case 42:
                        switch "a" {
                        case "":
                            v13 = v13
                        }
                    case 32767:
                        switch "hello" {
                        case "a":
                            v10 = v10
                        case "\n\t":
                            var v17 int = fi12
                            _ = v17
                        }
                    }
                    v18 := GenS1{F0: 1, F1: -292, F2: 0}
                    v19 := GenS2{F0: 65535, F1: 42}
                    v5[(fi12 << ((2) & 63))] = (fi12 % (((^(fi14))) | 1))
                    _ = v18
                    _ = v19
                case 2147483647:
                    delete(v5, fi12)
                }
                _ = fi14
            }
            _ = v13
        }
        _ = fi12
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            if (!false) {
                v1 = v1
                delete(v5, fi21)
            } else {
                fmt.Println(v11, v10, v11)
                switch fi21 {
                case 83:
                    var v22 string = ("\n\t" + "\n\t")
                    _ = v22
                default:
                    _ = len(v3)
                }
            }
            v1 = 127
            v23 := GenS1{F0: 255, F1: (-(v10)), F2: 32767}
            _ = fi21
            _ = v23
        }
        var v24 bool = false
        if (v24 || v24) {
            v24 = ((615 != v10) || (-9223372036854775808 > 695))
        } else {
            v10 = (v10 % ((-593) | 1))
            delete(v4, -9223372036854775808)
            v1 = (v1 / (((v1 + 175)) | 1))
        }
        _ = fi20
        _ = v24
    }
    var v25 bool = (2 == 1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
