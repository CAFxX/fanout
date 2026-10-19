package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int; F1 uint; F2 int }
type GenS3 struct { F0 bool; F1 uint64; F2 uint }
type GenS4 struct { F0 int32; F1 string }

func main() {
    v1 := make(map[string]string)
    v2 := []string{"a", "\n\t", "foo bar", ("a" + ("hello" + ("" + "a")))}
    var v3 float32 = 1e10
    v4 := []string{"hello"}
    switch "foo bar" {
    case "a":
        if false {
            for fi5 := 0; fi5 < 17; fi5++ {
                v6 := GenS1{F0: (!true), F1: 1}
                v3 = v3
                _ = fi5
                _ = v6
            }
            if (168 <= 134) {
                v7 := GenS2{F0: 42, F1: 415, F2: 2147483647}
                _ = v7
            }
            var v8 uint = 350
            _ = v8
        } else {
            v3 = -2.25
        }
    case "":
        var v9 float32 = v3
        _ = v9
    case "\n\t":
        if true {
            var v10 int = 127
            _ = v10
        } else {
            if (42 < -1) {
                switch 32767 {
                case 32767:
                    switch 32768 {
                    case 9223372036854775807:
                        var v11 int8 = -74
                        _ = v11
                    case 127:
                        v12 := GenS2{F0: 498, F1: 859, F2: 32768}
                        _ = v12
                    default:
                        fmt.Println(v3)
                    }
                case 1:
                    v3 -= 0.1
                default:
                    var v13 int8 = 0
                    _ = v13
                }
            } else {
                var v14 uint16 = 818
                v15 := GenS1{F0: true, F1: -1}
                _ = v14
                _ = v15
            }
            var v16 float64 = float64(float32((0.1 - (float64(float32(((1.5 / 1.5) / 3.14159))) + 0.0))))
            var v17 bool = true
            _ = v16
            _ = v17
        }
    default:
        v3 += -2.25
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int32 = 959
        v20 := GenS2{F0: (fi18 / ((fi18) | 1)), F1: 65535, F2: (fi18 % (((fi18 + fi18)) | 1))}
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int64 = 1
    var v22 int = 128
    switch (v22 | (-(v22))) {
    case 32767:
        if (18446744073709551615 <= 1) {
            var v23 uint32 = 1
            var v24 int16 = 42
            _ = v23
            _ = v24
        } else {
            v21 = (v21 / ((-564) | 1))
        }
    case 2147483647:
        var v25 int16 = -1
        _ = v25
    case -945:
        if (!false) {
            switch v22 {
            case -643:
                fmt.Println(v3)
            default:
                var v26 int16 = 1
                _ = v26
            }
        }
    default:
        v21 = 1
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 uint8 = 128
        var v29 int16 = 32767
        _ = fi27
        _ = v28
        _ = v29
    }
    v3 = 1.5
    v22 = 2147483647
    if len(v2) > 0 {
        v2[0] = ("a" + "a")
    }
    v30 := GenS2{F0: 398, F1: 1, F2: (v22 + v22)}
    for fi31 := 0; fi31 < 12; fi31++ {
        switch ((("" + "") + "hello") + "") {
        case "hello":
            v1["a"] = "\n\t"
        case "\n\t":
            if ((!(false && ((true && ((!true) && (0 != 134))) && false))) || true) {
                fmt.Println(v3, fi31, v21)
                for fi32 := 0; fi32 < 20; fi32++ {
                    _ = len(v1)
                    _ = fi32
                }
                for fi33 := 0; fi33 < 2; fi33++ {
                    v1["hello"] = "hello"
                    v34 := GenS2{F0: (fi31 << ((0) & 63)), F1: 821, F2: (fi31 & (v22 / ((1) | 1)))}
                    if (false && false) {
                        v35 := GenS2{F0: v22, F1: 1, F2: (^(fi33))}
                        fmt.Println(fi33, v22, fi33)
                        var v36 bool = false
                        _ = v35
                        _ = v36
                    } else {
                        if (v21 > (v21 & (v21 ^ 1))) {
                            v37 := GenS3{F0: (false && false), F1: 2, F2: 1}
                            _ = v37
                        } else {
                            delete(v1, ("\n\t" + "hello"))
                        }
                        if false {
                            var v38 bool = true
                            _ = v38
                        }
                        if false {
                            v39 := GenS4{F0: 2147483647, F1: ("" + ("a" + ("hello" + "a")))}
                            _ = v39
                        }
                    }
                    var v40 uint16 = 42
                    var v41 uint16 = v40
                    _ = fi33
                    _ = v34
                    _ = v40
                    _ = v41
                }
            }
        case "a":
            fmt.Println(v21, v3)
        }
        var v42 int8 = 100
        var v43 int16 = -32768
        _ = fi31
        _ = v42
        _ = v43
    }
    var v44 int16 = -901
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v44
}
