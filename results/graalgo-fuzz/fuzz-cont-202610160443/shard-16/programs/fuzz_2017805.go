package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS6 struct { F0 uint; F1 uint32; F2 int32 }
func genFn0(p0 string) uint32 {
    p0 = (p0 + "\n\t")
    var v5 bool = (false && false)
    v6 := []int{-9223372036854775808, 910}
    _ = p0
    _ = v5
    _ = v6
    return 2147483648
}
func genFn1() int {
    var v7 uint64 = 18446744073709551615
    if false {
        var v8 int64 = -9223372036854775808
        v9 := GenS1{F0: "hello", F1: v8}
        _ = v8
        _ = v9
    } else {
        v10 := GenS2{F0: 1, F1: (128 <= 32768), F2: v7}
        var v11 bool = true
        _ = v10
        _ = v11
    }
    if true {
        for fi12 := 0; fi12 < 10; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 int16 = 1
                v15 := []int{(fi12 + int(v7)), 32767, fi12, fi12}
                v16 := GenS2{F0: 1, F1: false, F2: v7}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi12
        }
        v7 -= (v7 << ((2) & 63))
    } else {
        v7 = (v7 + (v7 - v7))
    }
    switch -230 {
    case -1:
        v7 = v7
    case 167:
        var v17 uint = 1
        _ = v17
    case 32767:
        if false {
            var v18 float32 = (((0.1 * float32(3.14159)) + 3.14159) + (1.5 / float32((3.14159 - 1e10))))
            _ = v18
        } else {
            v19 := []int{32768, -1, -920, -292}
            var v20 uint = 431
            _ = v19
            _ = v20
        }
    }
    v7 = (v7 + (^(v7)))
    _ = v7
    return -335
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []uint32{}
    var v3 int8 = -49
    var v4 uint = 940
    var v21 int = 128
    if (float64(1e10) < 3.14159) {
        if len(v2) > 0 {
            v2[0] = 327
        }
    } else {
        fmt.Println(v21)
    }
    var v22 uint = 4294967295
    v21 = -2147483648
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 uint8 = 128
        var v25 float64 = (1.5 * float64(1.5))
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 int = (v21 * v21)
    v2 = append(v2, 4294967295)
    var v27 uint32 = 2
    v28 := GenS3{F0: 0, F1: "hello"}
    var v29 int16 = int16(v27)
    for fi30 := 0; fi30 < 15; fi30++ {
        v31 := GenS4{F0: 42}
        for fi32 := 0; fi32 < 11; fi32++ {
            if true {
                if (v4 < (v4 / (((v4 % ((v22) | 1))) | 1))) {
                    v1 = (v1 - v1)
                    v21 = (v21 + (fi30 * 32767))
                } else {
                    v29 = (v29 + int16(v1))
                }
            }
            fmt.Println(v21, fi30)
            fmt.Println(v4)
            _ = fi32
        }
        for fi33 := 0; fi33 < 12; fi33++ {
            v34 := GenS5{F0: 0, F1: 301, F2: genFn0("foo bar")}
            _ = fi33
            _ = v34
        }
        _ = fi30
        _ = v31
    }
    switch (v21 / ((int(v22)) | 1)) {
    case 596:
        for fi35 := 0; fi35 < 7; fi35++ {
            v4 = (^(v4))
            switch (("" + "") + ("\n\t" + "a")) {
            case "hello":
                switch "foo bar" {
                case "\n\t":
                    v36 := GenS3{F0: 2, F1: ""}
                    _ = v36
                default:
                    var v37 uint32 = v27
                    _ = v37
                }
            case "a":
                var v38 uint8 = 2
                _ = v38
            default:
                v39 := GenS3{F0: 444, F1: "a"}
                _ = v39
            }
            v40 := GenS3{F0: 2, F1: "foo bar"}
            _ = fi35
            _ = v40
        }
    case -2147483648:
        if ((v26 % ((-2147483648) | 1)) != v21) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v41 := GenS6{F0: (v4 ^ 1), F1: uint32(v22), F2: 0}
            _ = v41
        }
    default:
        for fi42 := 0; fi42 < 3; fi42++ {
            if (true && (!false)) {
                var v43 int16 = v29
                v26 = (v26 & 32768)
                _ = v43
            } else {
                for fi44 := 0; fi44 < 8; fi44++ {
                    if (!(-1 != 42)) {
                        v26 = int(v1)
                        if false {
                            var v45 uint = 0
                            _ = v45
                        } else {
                            var v46 uint = v4
                            v22 = v46
                            _ = v46
                        }
                        fmt.Println(v4, v22, v27)
                    } else {
                        var v47 uint64 = 42
                        v4 -= (v4 - v22)
                        _ = v47
                    }
                    _ = fi44
                }
                var v48 int = v21
                _ = v48
            }
            if true {
                fmt.Println(v27, v27)
                for fi49 := 0; fi49 < 8; fi49++ {
                    switch (("a" + "\n\t") + (("foo bar" + ("foo bar" + "")) + ("" + "hello"))) {
                    case "hello":
                        v1 = 0
                    case "":
                        fmt.Println(v1, v3, fi49)
                    default:
                        var v50 int16 = v29
                        _ = v50
                    }
                    _ = fi49
                }
            }
            var v51 float32 = 0.0
            _ = fi42
            _ = v51
        }
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v26, v26, v1)
    case "foo bar":
        _ = len(v2)
    default:
        if (226 <= 128) {
            var v52 uint8 = 154
            _ = v52
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
