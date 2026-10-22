package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int32; F1 uint; F2 int16 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64; F1 int8; F2 uint32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    var v3 int64 = -9223372036854775808
    v4 := make(map[uint32]string)
    v5 := []uint32{}
    var v6 uint16 = 42
    _ = len(v5)
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: -1, F1: 2}
        v6 = v6
        _ = fi7
        _ = v8
    }
    var v9 float64 = -2.25
    var v10 bool = false
    var v11 uint16 = (^(v6))
    v9 += v9
    switch "a" {
    case "hello":
        var v12 uint8 = 130
        _ = v12
    case "a":
        if v10 {
            switch -2147483648 {
            case 0:
                var v13 bool = v10
                _ = v13
            }
        }
    case "\n\t":
        switch "\n\t" {
        case "foo bar":
            if (v9 >= float64(3.14159)) {
                var v14 bool = true
                for fi15 := 0; fi15 < 6; fi15++ {
                    v16 := GenS1{F0: -1, F1: 945}
                    _ = len(v5)
                    if ((2 < 65535) || (v9 >= ((((1.5 + ((v9 * float64(0.1)) + ((3.14159 * v9) / 1.0))) * v9) - ((v9 - 1.5) - 1e10)) - -2.25))) {
                        var v17 int32 = 2147483647
                        if (v10 && v14) {
                            fmt.Println(v10)
                            v18 := GenS2{F0: int32(fi15), F1: 18446744073709551615, F2: -32768}
                            _ = v18
                        } else {
                            v19 := GenS2{F0: (-(v17)), F1: 65535, F2: 0}
                            _ = v19
                        }
                        _ = v17
                    } else {
                        if v10 {
                            for fi20 := 0; fi20 < 7; fi20++ {
                                v21 := GenS2{F0: 42, F1: 1, F2: 1}
                                switch fi15 {
                                case -19:
                                    switch (-(fi20)) {
                                    case 9223372036854775807:
                                        if ((fi15 >> ((961) & 63)) <= -9223372036854775808) {
                                            switch fi20 {
                                            case -591:
                                                v10 = (!v14)
                                            case 4294967295:
                                                v22 := GenS3{F0: 2, F1: -94}
                                                _ = v22
                                            case 0:
                                                v23 := GenS2{F0: 432, F1: 198, F2: 1}
                                                _ = v23
                                            }
                                            fmt.Println(v9, v3)
                                        } else {
                                            switch 32768 {
                                            case -158:
                                                v3 = v3
                                            case 588:
                                                var v24 int = fi20
                                                _ = v24
                                            default:
                                                v14 = v10
                                            }
                                            v25 := GenS1{F0: 122, F1: 0}
                                            _ = v25
                                        }
                                    case -9223372036854775808:
                                        v11 = (^(v6))
                                    }
                                case 128:
                                    delete(v4, uint32(v6))
                                    v26 := GenS4{F0: 42}
                                    _ = len(v2)
                                    v6 -= v11
                                    _ = v26
                                case 806:
                                    switch 1 {
                                    case 32768:
                                        fmt.Println(fi20, fi15, fi20)
                                    case -85:
                                        switch ((("\n\t" + "hello") + "hello") + ("foo bar" + "foo bar")) {
                                        case "\n\t":
                                            v10 = (0 != v11)
                                        case "foo bar":
                                            var v27 uint8 = 255
                                            _ = v27
                                        case "":
                                            v9 -= -2.25
                                        default:
                                            v11 = v11
                                        }
                                    }
                                }
                                _ = fi20
                                _ = v21
                            }
                            v28 := GenS5{F0: -32768, F1: 100}
                            _ = v28
                        } else {
                            var v29 uint16 = (v6 * v6)
                            v30 := GenS6{F0: 42}
                            _ = v29
                            _ = v30
                        }
                    }
                    _ = fi15
                    _ = v16
                }
                _ = v14
            } else {
                if len(v5) > 0 {
                    v5[0] = 2
                }
                var v31 uint8 = 255
                fmt.Println(v10, v9, v10)
                _ = v31
            }
        case "a":
            for fi32 := 0; fi32 < 19; fi32++ {
                for fi33 := 0; fi33 < 9; fi33++ {
                    var v34 int32 = 2147483647
                    if true {
                        switch (fi33 % ((9223372036854775807) | 1)) {
                        case -190:
                            switch ("hello" + "\n\t") {
                            case "hello":
                                _ = len(v5)
                            default:
                                v9 -= float64(0.1)
                            }
                        case 2147483647:
                            v35 := GenS5{F0: 42, F1: -121}
                            _ = v35
                        default:
                            v6 = 593
                        }
                        var v36 uint64 = 18446744073709551615
                        _ = v36
                    }
                    if v10 {
                        var v37 uint8 = 1
                        _ = v37
                    }
                    if ("" < "\n\t") {
                        v38 := GenS5{F0: 0, F1: 54}
                        switch (-(fi32)) {
                        case -9223372036854775808:
                            v39 := GenS3{F0: 42, F1: 0}
                            _ = v39
                        }
                        v10 = v10
                        _ = v38
                    }
                    v40 := GenS2{F0: -1, F1: 255, F2: -1}
                    _ = fi33
                    _ = v34
                    _ = v40
                }
                _ = fi32
            }
        case "":
            switch "a" {
            case "a":
                fmt.Println(v9, v3)
            }
        }
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        v42 := GenS7{F0: v6}
        v43 := GenS1{F0: 1000, F1: 55}
        if false {
            switch ("" + "hello") {
            case "\n\t":
                v44 := v4[42]
                _ = v44
            case "":
                if true {
                    v5 = append(v5, 4294967295)
                }
            default:
                var v45 int8 = -128
                _ = v45
            }
            v46 := GenS8{F0: 190}
            if v10 {
                v3 += (-(v3))
                for fi47 := 0; fi47 < 7; fi47++ {
                    switch ("hello" + "") {
                    case "\n\t":
                        v48 := GenS2{F0: 0, F1: 65535, F2: 58}
                        _ = v48
                    case "foo bar":
                        v9 += v9
                    }
                    _ = fi47
                }
            } else {
                v9 = 0.1
            }
            _ = v46
        }
        _ = fi41
        _ = v42
        _ = v43
    }
    v9 = v9
    v49 := GenS3{F0: 773, F1: int8(v3)}
    v50 := GenS9{F0: v3, F1: 100, F2: 733}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v9
}
