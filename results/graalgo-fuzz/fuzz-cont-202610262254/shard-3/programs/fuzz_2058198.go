package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint16; F1 uint8; F2 int }
type GenS4 struct { F0 uint16; F1 uint32; F2 int32 }
func genFn0(p0 uint, p1 uint32) int32 {
    v5 := GenS1{F0: ""}
    if (1 > 128) {
        if true {
            v6 := []int{int(p1), 1, -969}
            for fi7 := 0; fi7 < 6; fi7++ {
                for fi8 := 0; fi8 < 14; fi8++ {
                    var v9 float32 = 3.14159
                    var v10 int = (fi8 - fi8)
                    var v11 int32 = -2147483648
                    if false {
                        var v12 uint16 = 255
                        switch ("foo bar" + "hello") {
                        case "hello":
                            var v13 uint32 = (p1 / (((p1 * p1)) | 1))
                            _ = v13
                        case "foo bar":
                            var v14 bool = false
                            _ = v14
                        default:
                            v10 = (v10 + (v10 << ((p0) & 63)))
                        }
                        _ = v12
                    } else {
                        var v15 uint32 = (p1 >> ((42) & 63))
                        _ = v15
                    }
                    v16 := GenS2{F0: 42, F1: 780}
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v16
                    _ = v9
                }
                for fi17 := 0; fi17 < 5; fi17++ {
                    var v18 string = ""
                    _ = fi17
                    _ = v18
                }
                _ = fi7
            }
            v19 := GenS1{F0: "\n\t"}
            _ = v19
            _ = v6
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            fmt.Println(fi20)
            var v21 uint8 = 50
            _ = fi20
            _ = v21
        }
    }
    v22 := []int{}
    fmt.Println(p0, p0)
    var v23 bool = false
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    _ = v5
    return 2147483647
}
func genFn1(p0 int32, p1 int32) uint32 {
    for fi24 := 0; fi24 < 14; fi24++ {
        switch "" {
        case "\n\t":
            var v25 float64 = 3.14159
            _ = v25
        }
        if (((("" + "a") + "\n\t") < "") && true) {
            p1 -= 1
            v26 := GenS1{F0: "foo bar"}
            _ = v26
        }
        p0 = (p0 >> ((793) & 63))
        _ = fi24
    }
    switch 128 {
    case -394:
        if (1.5 >= 0.1) {
            fmt.Println(p0, p0)
            var v27 float32 = ((1e10 + -2.25) / (1.5 * 0.1))
            _ = v27
        }
    case 2147483647:
        p1 = p1
    case -106:
        p0 = p0
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint16 = 764
    var v2 uint64 = 42
    var v3 string = "\n\t"
    var v4 uint = 242
    if false {
        for fi28 := 0; fi28 < 18; fi28++ {
            v2 += (v2 << ((v4) & 63))
            _ = fi28
        }
        var v29 uint32 = uint32(v1)
        if false {
            v4 += v4
            v2 += (v2 % (((v2 % ((v2) | 1))) | 1))
        } else {
            fmt.Println(v3, v3)
            v2 = (v2 * uint64(v4))
        }
        _ = v29
    } else {
        v3 = v3
    }
    fmt.Println(v4)
    v30 := []int{-752, 1}
    var v31 uint8 = 0
    v32 := GenS3{F0: v1, F1: (v31 ^ v31), F2: -9223372036854775808}
    for fi33 := 0; fi33 < 19; fi33++ {
        var v34 int16 = int16(fi33)
        for fi35 := 0; fi35 < 3; fi35++ {
            if true {
                var v36 int32 = -1
                v37 := GenS1{F0: ((v3 + (("\n\t" + (v3 + "hello")) + "")) + "a")}
                fmt.Println(v31, v36, v34)
                _ = v36
                _ = v37
            } else {
                switch ((((v3 + v3) + "hello") + v3) + (v3 + v3)) {
                case "":
                    v38 := GenS4{F0: uint16(v31), F1: 0, F2: genFn0(18446744073709551615, 42)}
                    _ = v38
                case "\n\t":
                    _ = len(v30)
                }
                if len(v30) > 0 {
                    v30[0] = 32768
                }
            }
            var v39 uint8 = 42
            v2 -= (^(v2))
            _ = fi35
            _ = v39
        }
        var v40 float32 = 0.0
        _ = fi33
        _ = v34
        _ = v40
    }
    v41 := GenS3{F0: v1, F1: (v31 ^ v31), F2: 128}
    var v42 uint32 = 2
    var v43 int64 = -1
    var v44 float32 = (-2.25 * 3.14159)
    var v45 uint = v4
    fmt.Println(v2)
    switch v3 {
    case "hello":
        v46 := GenS4{F0: v1, F1: v42, F2: 0}
        _ = v46
    default:
        v47 := GenS4{F0: (v1 * v1), F1: (^(v42)), F2: genFn0((^(v45)), v42)}
        _ = v47
    }
    v45 += (^(v45))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
