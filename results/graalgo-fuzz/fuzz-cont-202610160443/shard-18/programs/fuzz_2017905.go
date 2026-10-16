package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 bool }
type GenS2 struct { F0 int32; F1 int; F2 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 string; F2 int32 }
type GenS5 struct { F0 bool; F1 uint32; F2 string }

func main() {
    var v1 int64 = 0
    v2 := []int{32768, 128, 42}
    v3 := make(map[int]string)
    var v4 int32 = -1
    var v5 uint8 = 1
    var v6 bool = (4294967295 <= -9223372036854775808)
    v7 := GenS1{F0: 1, F1: 220, F2: (128 != 989)}
    var v8 int8 = -128
    switch 4294967295 {
    case 2147483647:
        if ("\n\t" >= "\n\t") {
            var v9 float64 = -2.25
            v10 := GenS2{F0: v4, F1: -839, F2: v4}
            v11 := GenS2{F0: (v4 - -120), F1: 938, F2: (v4 - -2147483648)}
            _ = v10
            _ = v11
            _ = v9
        }
    case -9223372036854775808:
        _ = len(v2)
    case 42:
        fmt.Println(v8, v6)
    default:
        fmt.Println(v4, v8, v1)
    }
    v12 := GenS1{F0: 42, F1: 2147483648, F2: v6}
    v5 -= (v5 ^ v5)
    fmt.Println(v5, v4)
    v13 := GenS3{F0: v4}
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS2{F0: (v4 & v4), F1: (fi14 - (fi14 >> ((uint(v5)) & 63))), F2: -2147483648}
        v5 -= 127
        _ = fi14
        _ = v15
    }
    if v6 {
        var v16 bool = v6
        for fi17 := 0; fi17 < 12; fi17++ {
            v6 = v6
            switch ("hello" + (("a" + "") + "\n\t")) {
            case "a":
                for fi18 := 0; fi18 < 1; fi18++ {
                    _ = len(v2)
                    fmt.Println(fi18, v8)
                    _ = fi18
                }
            case "foo bar":
                v19 := v3[(fi17 & (fi17 * (fi17 - fi17)))]
                _ = v19
            }
            for fi20 := 0; fi20 < 15; fi20++ {
                fmt.Println(v8, v5)
                var v21 uint = 117
                var v22 float64 = (1e10 / (((1.5 - 0.0) * 0.0) - float64((1e10 - ((1e10 * -2.25) / 0.1)))))
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi17
        }
        _ = v16
    } else {
        v23 := GenS4{F0: (v1 & (v1 / ((380) | 1))), F1: ("hello" + ""), F2: -1}
        _ = v23
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        fmt.Println(v1)
        if (!true) {
            var v25 uint8 = (v5 & 234)
            v26 := GenS3{F0: v4}
            var v27 int64 = 1
            _ = v25
            _ = v26
            _ = v27
        } else {
            v2 = append(v2, fi24)
        }
        _ = fi24
    }
    if (!(float64(0.1) != float64(-2.25))) {
        if v6 {
            for fi28 := 0; fi28 < 13; fi28++ {
                v29 := GenS3{F0: (v4 >> ((23) & 63))}
                _ = fi28
                _ = v29
            }
            v8 += 127
            var v30 uint16 = 42
            _ = v30
        } else {
            delete(v3, 17)
        }
        var v31 uint64 = 296
        v32 := GenS2{F0: 2147483647, F1: 128, F2: v4}
        _ = v31
        _ = v32
    }
    var v33 int16 = -1
    for fi34 := 0; fi34 < 10; fi34++ {
        if v6 {
            if v6 {
                _ = len(v3)
                for fi35 := 0; fi35 < 5; fi35++ {
                    v1 -= (v1 / (((v1 | v1)) | 1))
                    _ = fi35
                }
            }
            v36 := GenS5{F0: (2 < 242), F1: 486, F2: "a"}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v36
        }
        _ = fi34
    }
    v37 := v3[127]
    var v38 int32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
