package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 string }
func genFn0() int16 {
    var v4 uint32 = 42
    var v5 string = "\n\t"
    switch 32768 {
    case 42:
        v6 := GenS1{F0: -824}
        _ = v6
    case -47:
        v7 := GenS2{F0: 4294967295}
        _ = v7
    case 851:
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint8 = 128
            switch "" {
            case "":
                if false {
                    v10 := GenS2{F0: 4294967295}
                    var v11 uint16 = 32767
                    v12 := GenS3{F0: false, F1: 9223372036854775807}
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    var v13 float64 = -2.25
                    _ = v13
                }
            default:
                if (false && false) {
                    v9 = v9
                    for fi14 := 0; fi14 < 12; fi14++ {
                        v15 := GenS3{F0: (49 > 127), F1: -568}
                        fmt.Println(v9, v4)
                        _ = fi14
                        _ = v15
                    }
                    var v16 uint = 4294967295
                    _ = v16
                }
            }
            v17 := []int{(fi8 * 0)}
            _ = fi8
            _ = v17
            _ = v9
        }
    }
    fmt.Println(v4)
    var v18 uint8 = 0
    _ = v18
    _ = v4
    _ = v5
    return -820
}

func main() {
    var v1 int16 = 42
    var v2 float32 = -2.25
    v3 := make(map[string]bool)
    if (-227 > 0) {
        v1 += v1
    } else {
        if (((-(v1)) <= (v1 / ((v1) | 1))) && (false && true)) {
            v19 := GenS2{F0: 548}
            var v20 float64 = 1.5
            _ = v19
            _ = v20
        }
    }
    switch ("\n\t" + "hello") {
    case "a":
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 uint64 = 1
            v23 := GenS2{F0: v22}
            v24 := GenS3{F0: ((!(208 >= 42)) && false), F1: 130}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
    case "hello":
        if (1.5 == v2) {
            if (64 <= 42) {
                var v25 string = ("" + "")
                switch "\n\t" {
                case "hello":
                    v2 += -2.25
                case "a":
                    v26 := GenS2{F0: 878}
                    _ = v26
                case "foo bar":
                    var v27 int32 = -293
                    _ = v27
                }
                delete(v3, v25)
                _ = v25
            } else {
                if true {
                    var v28 uint8 = 128
                    _ = v28
                }
                if false {
                    v1 -= -32768
                    v1 -= v1
                    v2 = float32(1.5)
                } else {
                    for fi29 := 0; fi29 < 12; fi29++ {
                        v1 = (v1 >> ((65535) & 63))
                        _ = fi29
                    }
                }
                v2 = v2
            }
        } else {
            for fi30 := 0; fi30 < 1; fi30++ {
                v2 -= (v2 - (float32(-2.25) / v2))
                v31 := v3["foo bar"]
                _ = fi30
                _ = v31
            }
            v32 := []int{334, 32768}
            _ = v32
        }
    case "\n\t":
        v3[("\n\t" + "hello")] = (v1 == v1)
    }
    v33 := GenS4{F0: -1, F1: genFn0()}
    for fi34 := 0; fi34 < 9; fi34++ {
        var v35 int = 1
        var v36 int64 = 9223372036854775807
        if (1 > 65535) {
            if (65535 == 143) {
                var v37 uint16 = 551
                _ = v37
            }
        } else {
            for fi38 := 0; fi38 < 17; fi38++ {
                var v39 int32 = int32(fi34)
                if false {
                    var v40 bool = true
                    var v41 uint8 = 0
                    v41 = (^(v41))
                    _ = v40
                    _ = v41
                } else {
                    var v42 int16 = v1
                    var v43 uint16 = 42
                    v44 := v3[("a" + "foo bar")]
                    _ = v42
                    _ = v43
                    _ = v44
                }
                fmt.Println(v2, fi34)
                v45 := v3[("" + ("hello" + "foo bar"))]
                switch 4294967295 {
                case 128:
                    var v46 string = "\n\t"
                    _ = v46
                case -2147483648:
                    var v47 string = "a"
                    _ = v47
                case 9223372036854775807:
                    fmt.Println(fi34)
                default:
                    v48 := GenS5{F0: 1}
                    var v49 uint = 18446744073709551615
                    fmt.Println(v1)
                    _ = v48
                    _ = v49
                }
                fmt.Println(v35, v2)
                v50 := GenS6{F0: 18446744073709551615, F1: ""}
                _ = fi38
                _ = v39
                _ = v45
                _ = v50
            }
            var v51 uint = 1
            _ = v51
        }
        _ = fi34
        _ = v35
        _ = v36
    }
    v2 = ((0.0 * v2) + v2)
    var v52 uint16 = 0
    var v53 uint = 424
    v54 := GenS3{F0: (390 < 1), F1: 42}
    if true {
        fmt.Println(v2, v52)
        var v55 int8 = 55
        _ = v55
    }
    for fi56 := 0; fi56 < 20; fi56++ {
        v57 := []int{}
        _ = fi56
        _ = v57
    }
    v58 := GenS6{F0: 2, F1: "a"}
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v52
    _ = v53
    _ = v54
    _ = v58
}
