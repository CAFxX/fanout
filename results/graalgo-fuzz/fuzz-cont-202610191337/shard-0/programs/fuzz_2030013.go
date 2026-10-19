package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int16; F1 int; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint64) bool {
    var v4 string = (("a" + "hello") + "\n\t")
    for fi5 := 0; fi5 < 15; fi5++ {
        fmt.Println(v4, v4, fi5)
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := []int{(fi5 + fi5)}
            _ = fi6
            _ = v7
        }
        v4 = ""
        _ = fi5
    }
    _ = p0
    _ = v4
    return true
}

func main() {
    var v1 int16 = 1
    var v2 uint32 = 316
    var v3 float64 = float64(0.1)
    if ((1.5 != 1e10) || (false && false)) {
        v8 := GenS1{F0: (v1 << ((4294967295) & 63)), F1: (^(v1))}
        v9 := []int{32767, 1, 32768, 32767}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 uint16 = 2
        if (208 != 128) {
            var v12 float32 = 0.0
            var v13 float32 = 1e10
            v14 := GenS2{F0: true, F1: (v2 >> ((42) & 63))}
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    var v15 int = -629
    if (1 > 127) {
        var v16 float32 = 1e10
        v3 = v3
        var v17 bool = false
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 7; fi18++ {
            v19 := GenS3{F0: -1, F1: (v15 + (fi18 % ((-996) | 1))), F2: (v2 >> ((331) & 63))}
            _ = fi18
            _ = v19
        }
        switch 476 {
        case 408:
            switch 0 {
            case -992:
                for fi20 := 0; fi20 < 8; fi20++ {
                    v21 := []int{}
                    var v22 float64 = float64(1e10)
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            default:
                v3 = (((v3 * v3) / v3) / 1e10)
            }
        case 189:
            v23 := GenS4{F0: 0, F1: 127}
            _ = v23
        }
        if (false && (!true)) {
            var v24 uint8 = 2
            var v25 string = ("a" + ("\n\t" + ("hello" + (("a" + ("a" + "a")) + "\n\t"))))
            _ = v24
            _ = v25
        }
    }
    var v26 int = 128
    v27 := []int{(^(v15))}
    for fi28 := 0; fi28 < 1; fi28++ {
        var v29 uint = 65535
        _ = fi28
        _ = v29
    }
    switch v26 {
    case 0:
        for fi30 := 0; fi30 < 16; fi30++ {
            switch "foo bar" {
            case "":
                v2 = 1
            case "hello":
                fmt.Println(v1, v3, v3)
            case "a":
                if len(v27) > 0 {
                    v27[0] = 32767
                }
            }
            _ = fi30
        }
    case 42:
        for fi31 := 0; fi31 < 3; fi31++ {
            var v32 uint64 = 1
            _ = fi31
            _ = v32
        }
    case 534:
        v33 := GenS4{F0: 128, F1: 66}
        _ = v33
    default:
        if len(v27) > 0 {
            v27[0] = (^(v26))
        }
    }
    fmt.Println(v2)
    v3 = float64(0.1)
    for fi34 := 0; fi34 < 2; fi34++ {
        for fi35 := 0; fi35 < 8; fi35++ {
            fmt.Println(v2, v3, v1)
            v2 += v2
            _ = fi35
        }
        v1 += v1
        if (!((fi34 & v26) >= 2147483647)) {
            var v36 string = ("\n\t" + "")
            fmt.Println(v15)
            v1 -= (v1 << ((413) & 63))
            _ = v36
        } else {
            if true {
                fmt.Println(v1, v26, v15)
                v27 = append(v27, fi34)
            } else {
                var v37 uint = 2
                v38 := GenS1{F0: -222, F1: 32767}
                switch 4294967295 {
                case 898:
                    if len(v27) > 0 {
                        v27[0] = (fi34 + v26)
                    }
                }
                _ = v37
                _ = v38
            }
            var v39 int32 = 42
            var v40 int16 = (v1 / (((v1 << ((18446744073709551615) & 63))) | 1))
            _ = v39
            _ = v40
        }
        _ = fi34
    }
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    var v41 int32 = 0
    v27 = append(v27, (v26 + (v26 * 42)))
    v41 = 1
    v42 := GenS4{F0: 255, F1: 93}
    v26 -= (v15 ^ 42)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v41
    _ = v42
}
