package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int64; F2 string }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64; F1 int32; F2 int8 }
func genFn0(p0 uint8) int64 {
    if false {
        fmt.Println(p0)
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: -5}
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 12; fi8++ {
            p0 += p0
            _ = fi8
        }
    } else {
        p0 -= p0
        var v9 int16 = 1000
        v10 := []int{-507, 2147483647, 2147483647}
        _ = v10
        _ = v9
    }
    p0 += p0
    switch 0 {
    case 128:
        var v11 uint32 = 2147483648
        _ = v11
    case 1:
        v12 := GenS1{F0: 42}
        _ = v12
    default:
        var v13 int64 = -9223372036854775808
        _ = v13
    }
    v14 := GenS2{F0: 605}
    p0 += (p0 / (((p0 / (((p0 % (((^(p0))) | 1))) | 1))) | 1))
    _ = p0
    _ = v14
    return 42
}
func genFn1(p0 uint16) bool {
    v15 := []int{}
    if true {
        for fi16 := 0; fi16 < 4; fi16++ {
            v15 = append(v15, fi16)
            _ = fi16
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            if len(v15) > 0 {
                v15[0] = fi17
            }
            switch (("foo bar" + "a") + (("" + "") + "\n\t")) {
            case "\n\t":
                for fi18 := 0; fi18 < 12; fi18++ {
                    v19 := GenS1{F0: 55}
                    p0 = (p0 + (^(p0)))
                    _ = fi18
                    _ = v19
                }
            case "foo bar":
                p0 = 2
            case "hello":
                for fi20 := 0; fi20 < 3; fi20++ {
                    p0 = (p0 >> ((1) & 63))
                    var v21 int = fi20
                    _ = fi20
                    _ = v21
                }
            default:
                p0 += 255
            }
            fmt.Println(fi17)
            _ = fi17
        }
    }
    v15 = append(v15, 32767)
    p0 -= p0
    var v22 int64 = -1
    _ = p0
    _ = v15
    _ = v22
    return true
}

func main() {
    var v1 float64 = 0.1
    var v2 uint64 = 590
    v3 := make(map[uint32]uint32)
    var v4 uint = 533
    var v5 int32 = 186
    fmt.Println(v2)
    var v23 uint64 = v2
    v24 := GenS3{F0: 836, F1: 1, F2: "\n\t"}
    switch -1 {
    case -326:
        v3[2] = 2147483647
    }
    v25 := v3[528]
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        switch "\n\t" {
        case "":
            var v26 int8 = 42
            _ = v26
        default:
            v27 := GenS4{F0: -1, F1: 440}
            _ = v27
        }
    case "foo bar":
        switch 529 {
        case 32768:
            for fi28 := 0; fi28 < 20; fi28++ {
                fmt.Println(v4, v2)
                for fi29 := 0; fi29 < 20; fi29++ {
                    var v30 int8 = 0
                    if (false && (!true)) {
                        if (0 > 858) {
                            var v31 int16 = 1
                            v5 += -1
                            _ = v31
                        }
                    }
                    switch ("" + (("\n\t" + "") + ("" + "hello"))) {
                    case "":
                        var v32 uint = (^(v4))
                        _ = v32
                    case "hello":
                        var v33 float32 = float32(0.0)
                        _ = v33
                    default:
                        v3[v25] = v25
                    }
                    _ = fi29
                    _ = v30
                }
                _ = fi28
            }
        }
    case "hello":
        v34 := GenS4{F0: 42, F1: -1}
        _ = v34
    }
    if (255 >= 0) {
        if (-32768 < 1000) {
            v4 += (v4 | v4)
            var v35 float64 = v1
            var v36 int = -1
            _ = v35
            _ = v36
        }
        for fi37 := 0; fi37 < 3; fi37++ {
            if false {
                fmt.Println(fi37)
            } else {
                v38 := GenS5{F0: int64(v4)}
                var v39 uint64 = 1
                _ = v38
                _ = v39
            }
            var v40 bool = ((false || (v25 != 485)) || (!true))
            v41 := GenS6{F0: (v23 * v2)}
            _ = fi37
            _ = v40
            _ = v41
        }
        var v42 float64 = v1
        _ = v42
    }
    var v43 int = -97
    v44 := GenS7{F0: -1, F1: (v5 % ((v5) | 1)), F2: 0}
    var v45 uint64 = (v2 % ((4294967295) | 1))
    var v46 int32 = 1
    var v47 uint8 = 107
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
