package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int; F1 int64; F2 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint16; F2 int32 }
func genFn0() uint16 {
    switch 32768 {
    case 457:
        var v7 uint16 = 138
        _ = v7
    case -775:
        var v8 uint8 = 1
        _ = v8
    case -97:
        v9 := GenS1{F0: 1}
        _ = v9
    }
    v10 := []int{-9223372036854775808}
    v10 = append(v10, -1)
    if (!false) {
        _ = len(v10)
        v11 := GenS2{F0: 0, F1: 42}
        var v12 string = ("" + "foo bar")
        _ = v11
        _ = v12
    }
    _ = v10
    return 1
}

func main() {
    var v1 int8 = 1
    var v2 string = "foo bar"
    v3 := make(map[int]bool)
    var v4 float64 = 0.0
    var v5 uint8 = 200
    v6 := make(map[int]string)
    if true {
        fmt.Println(v2, v5, v1)
        if (!(282 < 42)) {
            var v13 float64 = v4
            _ = v13
        }
    }
    v14 := []int{}
    fmt.Println(v4, v4)
    fmt.Println(v1)
    v4 = v4
    if false {
        v15 := GenS3{F0: 2147483647, F1: 42, F2: (v1 * v1)}
        _ = v15
    } else {
        v5 -= v5
    }
    v4 -= (3.14159 - v4)
    v16 := GenS2{F0: -1, F1: 255}
    if len(v14) > 0 {
        v14[0] = 1
    }
    if true {
        fmt.Println(v1, v1, v1)
        var v17 int32 = 42
        v17 = (v17 >> ((255) & 63))
        _ = v17
    }
    if (v2 > (v2 + (v2 + ""))) {
        v2 = v2
        v4 += v4
    } else {
        var v18 uint32 = 2147483647
        v19 := GenS4{F0: 146}
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: 0}
    v21 := v3[-599]
    if (v2 == ((v2 + (v2 + "")) + v2)) {
        var v22 int16 = 126
        if (2147483647 < 0) {
            if false {
                if (!(!((18446744073709551615 != 62) || (!(1 < 2147483647))))) {
                    switch -898 {
                    case 128:
                        fmt.Println(v4)
                    default:
                        v22 = v22
                    }
                }
                v14 = append(v14, 0)
                fmt.Println(v1, v4, v2)
            } else {
                switch 2147483647 {
                case -472:
                    v2 = v2
                case 2147483647:
                    switch 0 {
                    case 42:
                        switch -125 {
                        case 1:
                            switch 32768 {
                            case 127:
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                            case -210:
                                fmt.Println(v4, v5, v2)
                            default:
                                v4 = v4
                            }
                        }
                    case -775:
                        _ = len(v6)
                    default:
                        v2 = "hello"
                    }
                case 9223372036854775807:
                    v1 = (v1 << ((255) & 63))
                default:
                    v23 := GenS4{F0: 128}
                    _ = v23
                }
            }
            switch ("" + v2) {
            case "\n\t":
                delete(v6, int(v1))
            }
        }
        if (("" + v2) <= v2) {
            var v24 uint32 = 42
            switch 1 {
            case 42:
                v25 := GenS5{F0: v22, F1: 886, F2: int32(v1)}
                _ = v25
            case 128:
                var v26 float64 = ((v4 * v4) / float64(3.14159))
                _ = v26
            }
            _ = v24
        }
        _ = v22
    } else {
        for fi27 := 0; fi27 < 6; fi27++ {
            v28 := GenS3{F0: 920, F1: 0, F2: v1}
            switch 32767 {
            case 4294967295:
                var v29 bool = ((((((v4 / 3.14159) * float64(3.14159)) / v4) == ((0.0 + 1e10) / v4)) && (("" >= v2) || true)) && v21)
                _ = v29
            case -2147483648:
                v14 = append(v14, (-(fi27)))
            case 928:
                fmt.Println(v5)
            default:
                if v21 {
                    v5 -= (v5 << ((42) & 63))
                }
            }
            _ = fi27
            _ = v28
        }
    }
    if (v21 && v21) {
        fmt.Println(v1)
    }
    if v21 {
        v5 = (v5 ^ v5)
        for fi30 := 0; fi30 < 7; fi30++ {
            if (v1 == v1) {
                switch "" {
                case "\n\t":
                    var v31 uint16 = 2
                    _ = v31
                case "hello":
                    v32 := GenS5{F0: 32767, F1: 0, F2: 2147483647}
                    _ = v32
                case "a":
                    v4 = v4
                default:
                    v1 = v1
                }
            } else {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            }
            fmt.Println(v2, v2)
            if v21 {
                var v33 int64 = 1
                _ = v33
            }
            _ = fi30
        }
        for fi34 := 0; fi34 < 2; fi34++ {
            var v35 float32 = 1.5
            _ = fi34
            _ = v35
        }
    }
    switch v2 {
    case "":
        var v36 int64 = -401
        _ = v36
    case "foo bar":
        v5 = v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
