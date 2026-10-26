package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint8; F1 uint }
func genFn0(p0 int8) int8 {
    switch 32767 {
    case 9223372036854775807:
        v4 := []int{2147483647}
        _ = v4
    case -672:
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 int64 = 0
            _ = fi5
            _ = v6
        }
    default:
        v7 := GenS1{F0: 65535, F1: 1}
        _ = v7
    }
    switch 10 {
    case 32768:
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 int = (fi8 % ((fi8) | 1))
            _ = fi8
            _ = v9
        }
    case -279:
        v10 := []int{}
        _ = v10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int32 = 0
        _ = fi11
        _ = v12
    }
    _ = p0
    return 0
}

func main() {
    var v1 bool = (4294967295 <= 65535)
    var v2 uint8 = 255
    var v3 float64 = (1e10 * float64((0.1 + (0.0 + (-2.25 - 0.0)))))
    v3 -= (v3 - v3)
    var v13 uint32 = 945
    v3 = float64(0.0)
    fmt.Println(v3, v13, v13)
    for fi14 := 0; fi14 < 7; fi14++ {
        if false {
            var v15 int32 = 2147483647
            _ = v15
        } else {
            switch fi14 {
            case 0:
                var v16 uint16 = 836
                _ = v16
            case 32768:
                for fi17 := 0; fi17 < 15; fi17++ {
                    v18 := []int{(fi14 << ((838) & 63)), fi14, (fi17 % (((^(fi17))) | 1)), -242}
                    var v19 uint = 42
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case 4294967295:
                v20 := GenS1{F0: 32767, F1: v2}
                _ = v20
            default:
                v3 += v3
            }
            if v1 {
                v2 += v2
            }
        }
        _ = fi14
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        v13 = (^(v13))
        switch ("hello" + ("\n\t" + ("\n\t" + ("" + "a")))) {
        case "\n\t":
            var v22 int8 = 100
            _ = v22
        case "":
            var v23 uint64 = 508
            _ = v23
        case "a":
            v24 := GenS2{F0: fi21, F1: 18446744073709551615}
            _ = v24
        default:
            var v25 int32 = -109
            _ = v25
        }
        _ = fi21
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        if (18446744073709551615 <= 259) {
            v2 = 127
            var v27 uint8 = (v2 | v2)
            fmt.Println(v3)
            _ = v27
        }
        var v28 int16 = -957
        _ = fi26
        _ = v28
    }
    if false {
        fmt.Println(v2, v2)
    } else {
        v29 := []int{-2147483648, 42, -9223372036854775808}
        v30 := GenS3{F0: 633, F1: v1}
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        v1 = (18446744073709551615 > 42)
        v32 := []int{fi31, fi31}
        v33 := GenS2{F0: (^(fi31)), F1: 0}
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println(v1)
    v34 := []int{32767}
    v1 = v1
    v3 = v3
    if (2147483648 > (v13 << ((0) & 63))) {
        v1 = ((genFn0(genFn0(98)) != 0) || (v13 == v13))
    } else {
        if v1 {
            for fi35 := 0; fi35 < 19; fi35++ {
                var v36 bool = (42 == 1)
                v37 := GenS1{F0: 616, F1: (v2 % (((v2 + v2)) | 1))}
                _ = fi35
                _ = v36
                _ = v37
            }
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
            switch "a" {
            case "\n\t":
                if (uint64(v13) > 215) {
                    if v1 {
                        var v38 bool = false
                        if true {
                            fmt.Println(v3, v38, v13)
                        } else {
                            var v39 bool = ((809 < int(v13)) || false)
                            v13 -= (^(v13))
                            _ = v39
                        }
                        _ = v38
                    }
                    v40 := GenS3{F0: -1, F1: (42 < 255)}
                    _ = v40
                }
            }
        }
    }
    v41 := GenS4{F0: v2, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v41
}
