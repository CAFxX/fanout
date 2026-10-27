package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32; F1 string }
func genFn0(p0 int, p1 int64) bool {
    fmt.Println(p0, p0)
    v7 := GenS1{F0: 0, F1: 0, F2: 2147483647}
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS2{F0: 45}
            var v11 string = ""
            switch (^(fi9)) {
            case 1:
                var v12 float32 = float32(-2.25)
                _ = v12
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    var v13 int = p0
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return (1e10 >= 0.0)
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    var v4 int32 = 840
    v5 := []string{"hello"}
    v6 := []string{"hello", "\n\t"}
    v1 += 1
    if (65535 <= uint32(v1)) {
        if true {
            v4 = -2147483648
        }
        fmt.Println(v1, v1)
        if true {
            var v14 uint32 = 2147483647
            fmt.Println(v14, v1)
            v15 := GenS2{F0: 444}
            _ = v14
            _ = v15
        }
    } else {
        switch "a" {
        case "\n\t":
            var v16 int = 735
            _ = v16
        case "":
            for fi17 := 0; fi17 < 9; fi17++ {
                v1 = (v1 | v1)
                var v18 int32 = (v4 * (-(v4)))
                v19 := GenS3{F0: 127, F1: 2147483648, F2: 32767}
                _ = fi17
                _ = v18
                _ = v19
            }
        case "foo bar":
            var v20 float64 = (0.0 - 1e10)
            _ = v20
        }
        var v21 uint16 = 393
        _ = v21
    }
    v4 = (v4 / ((-2147483648) | 1))
    var v22 int8 = 1
    if (int(v4) > -762) {
        v22 -= (v22 >> ((65535) & 63))
        v23 := GenS4{F0: true, F1: 0}
        _ = v23
    } else {
        v4 = v4
        var v24 bool = (3.14159 <= 0.0)
        v25 := GenS5{F0: v24}
        _ = v24
        _ = v25
    }
    switch 766 {
    case 32768:
        var v26 float32 = (((0.0 / 1e10) / 1.0) / 1e10)
        _ = v26
    case 0:
        v1 += (v1 % ((v1) | 1))
    default:
        fmt.Println(v4)
    }
    if false {
        for fi27 := 0; fi27 < 11; fi27++ {
            var v28 int = (fi27 | fi27)
            fmt.Println(v4, fi27)
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            if (!false) {
                var v30 float32 = 0.1
                _ = v30
            }
            for fi31 := 0; fi31 < 15; fi31++ {
                switch "\n\t" {
                case "\n\t":
                    if (false || true) {
                        v6 = append(v6, "\n\t")
                        fmt.Println(v4, v22, v1)
                    }
                    switch fi29 {
                    case 522:
                        v1 = v1
                    case -761:
                        v22 = v22
                        _ = len(v6)
                    }
                    var v32 float64 = 0.0
                    if true {
                        switch "a" {
                        case "\n\t":
                            delete(v2, -570)
                        default:
                            var v33 int8 = v22
                            _ = v33
                        }
                        fmt.Println(v4)
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(fi29, v32)
                        case "a":
                            if true {
                                var v34 float32 = 0.1
                                v22 -= v22
                                _ = v34
                            } else {
                                var v35 uint64 = 689
                                _ = v35
                            }
                        default:
                            fmt.Println(v32, fi31, fi31)
                        }
                        v1 += v1
                        if true {
                            v36 := GenS3{F0: (^(v1)), F1: 1, F2: 255}
                            _ = v36
                        } else {
                            switch fi29 {
                            case 863:
                                v37 := GenS3{F0: v1, F1: 2147483647, F2: 42}
                                _ = v37
                            }
                        }
                    }
                    v1 += v1
                    _ = v32
                case "hello":
                    v38 := GenS6{F0: (v4 + (-(v4))), F1: ("hello" + "\n\t")}
                    _ = v38
                }
                if (1.5 == (0.1 + (1e10 - (0.0 / ((0.1 * (0.1 * ((1.5 + 1.5) * 0.1))) / ((1.5 / 1e10) * float64(((0.0 + float32((float64(0.0) - (1e10 - ((float64((float32(0.0) / ((0.0 + 1e10) + 1.5))) + (-2.25 * 3.14159)) * 3.14159))))) * (0.1 - 0.0))))))))) {
                    delete(v3, 1)
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v39 := GenS6{F0: v4, F1: ("a" + "")}
                    _ = v39
                }
                _ = fi31
            }
            delete(v2, fi29)
            _ = fi29
        }
    }
    v40 := GenS3{F0: 189, F1: 821, F2: 42}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        v1 = (v1 & v1)
        _ = fi41
    }
    switch 127 {
    case 128:
        v4 = int32(v1)
    }
    fmt.Println(v4)
    fmt.Println(v22, v22, v22)
    var v42 uint = 65535
    v43 := v3[65535]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
