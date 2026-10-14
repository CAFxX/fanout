package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint16; F1 bool; F2 int32 }
type GenS4 struct { F0 int16; F1 string }

func main() {
    var v1 uint16 = 2
    v2 := []int{730}
    v3 := make(map[uint32]string)
    v4 := make(map[int]bool)
    switch -2147483648 {
    case -1:
        v1 = 32767
    case -802:
        v1 += (^(v1))
    case 32767:
        v1 += (v1 * v1)
    }
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := v4[(fi5 * fi5)]
        if false {
            for fi7 := 0; fi7 < 10; fi7++ {
                if true {
                    v1 -= (v1 * (v1 >> ((255) & 63)))
                }
                if false {
                    var v8 float64 = float64((3.14159 * (-2.25 / 0.1)))
                    var v9 int64 = 42
                    v10 := GenS1{F0: 1}
                    _ = v10
                    _ = v8
                    _ = v9
                } else {
                    fmt.Println(v1, v6)
                    var v11 string = ""
                    v12 := v4[0]
                    _ = v11
                    _ = v12
                }
                v13 := GenS2{F0: 28, F1: true}
                _ = fi7
                _ = v13
            }
            switch fi5 {
            case 267:
                if v6 {
                    v6 = (true && false)
                    fmt.Println(v1, fi5)
                }
            }
            v4[(fi5 % ((128) | 1))] = (v6 || (-32768 >= 87))
        }
        var v14 int32 = -2147483648
        _ = fi5
        _ = v14
        _ = v6
    }
    var v15 int64 = 1
    switch 127 {
    case 238:
        var v16 uint = 65535
        _ = v16
    case 1:
        fmt.Println(v15, v1)
    case -163:
        v17 := GenS1{F0: 2147483648}
        _ = v17
    default:
        v18 := GenS2{F0: 65535, F1: false}
        _ = v18
    }
    switch "foo bar" {
    case "foo bar":
        var v19 int16 = 1
        _ = v19
    }
    if false {
        var v20 bool = ((42 != 32767) && false)
        _ = v20
    } else {
        if (1e10 >= 1.5) {
            switch 127 {
            case 814:
                if (false && (3.14159 > 0.0)) {
                    fmt.Println(v15)
                    if (282 >= 951) {
                        var v21 float32 = 0.0
                        _ = v21
                    }
                } else {
                    if true {
                        switch -998 {
                        case -530:
                            v2 = append(v2, 0)
                        case 306:
                            var v22 uint8 = 56
                            _ = v22
                        case -1:
                            v23 := GenS3{F0: (^(v1)), F1: true, F2: 117}
                            _ = v23
                        default:
                            if true {
                                _ = len(v4)
                            }
                        }
                        v2 = append(v2, 127)
                        v1 -= (v1 ^ 1)
                    } else {
                        v24 := GenS1{F0: 916}
                        v25 := GenS1{F0: 932}
                        _ = v24
                        _ = v25
                    }
                    var v26 string = ""
                    _ = v26
                }
            case -2147483648:
                for fi27 := 0; fi27 < 3; fi27++ {
                    v15 += 0
                    _ = fi27
                }
            }
            v28 := v3[2]
            var v29 uint16 = v1
            _ = v28
            _ = v29
        }
        v30 := v4[388]
        _ = v30
    }
    if (true || (18446744073709551615 != 2)) {
        v31 := GenS4{F0: 32767, F1: ""}
        var v32 uint64 = 0
        v33 := GenS2{F0: 388, F1: false}
        _ = v31
        _ = v32
        _ = v33
    } else {
        v1 = (v1 % ((2) | 1))
    }
    if (!true) {
        var v34 bool = (v1 < (v1 - v1))
        _ = v34
    }
    v1 += (^(v1))
    var v35 uint = 1
    var v36 int8 = 1
    v37 := GenS2{F0: uint32(v35), F1: (false || true)}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
