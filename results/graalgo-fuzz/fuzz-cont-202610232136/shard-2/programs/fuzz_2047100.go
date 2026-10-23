package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint16; F2 bool }
type GenS4 struct { F0 int; F1 uint8; F2 string }
type GenS5 struct { F0 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = float64((-2.25 * 0.0))
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]bool)
    v2 -= ((v2 / ((v2 + v2) / 1.0)) + v2)
    if (2 > 1) {
        switch 1 {
        case 32768:
            v5 := GenS1{F0: 174, F1: 1}
            _ = v5
        case -1:
            v2 = 1.5
        case 257:
            for fi6 := 0; fi6 < 18; fi6++ {
                if true {
                    v7 := GenS1{F0: 78, F1: 219}
                    v8 := GenS1{F0: 128, F1: 105}
                    fmt.Println(fi6, v2)
                    _ = v7
                    _ = v8
                }
                _ = fi6
            }
        default:
            var v9 string = ("foo bar" + ("foo bar" + ("\n\t" + "foo bar")))
            _ = v9
        }
        var v10 uint64 = 18446744073709551615
        _ = v10
    }
    if (true && true) {
        v11 := v1["a"]
        var v12 float64 = 0.0
        _ = v11
        _ = v12
    } else {
        if true {
            var v13 int = -43
            _ = v13
        }
    }
    var v14 uint16 = 900
    v14 -= (v14 ^ v14)
    v2 = v2
    v15 := GenS1{F0: 137, F1: 42}
    switch ("foo bar" + "foo bar") {
    case "hello":
        if false {
            switch 1 {
            case 9223372036854775807:
                v14 = v14
            case 127:
                fmt.Println(v14, v14)
            default:
                switch -221 {
                case 42:
                    for fi16 := 0; fi16 < 10; fi16++ {
                        fmt.Println(v2)
                        _ = fi16
                    }
                default:
                    v17 := GenS2{F0: 2}
                    _ = v17
                }
            }
            var v18 uint64 = 42
            _ = v18
        }
    case "a":
        v19 := GenS1{F0: 201, F1: 42}
        _ = v19
    case "":
        var v20 uint16 = 618
        _ = v20
    default:
        var v21 uint = 65535
        _ = v21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        v4[846] = (((481 < -1) && (-116 != 91)) || (18446744073709551615 < 0))
        v23 := GenS3{F0: 4294967295, F1: (v14 >> ((4294967295) & 63)), F2: false}
        _ = fi22
        _ = v23
    }
    v24 := GenS4{F0: -1, F1: uint8(v14), F2: ("hello" + "hello")}
    v14 = (v14 + v14)
    var v25 float32 = (float32(3.14159) / 0.1)
    delete(v4, 2147483647)
    v26 := v1["foo bar"]
    if (1 >= -1) {
        v27 := GenS4{F0: int(v14), F1: 2, F2: "foo bar"}
        for fi28 := 0; fi28 < 10; fi28++ {
            if v26 {
                v29 := GenS4{F0: fi28, F1: 250, F2: ""}
                if false {
                    var v30 int = (-(fi28))
                    v2 -= (v2 - -2.25)
                    _ = v30
                }
                _ = v29
            } else {
                switch "" {
                case "":
                    fmt.Println(v14, v14)
                case "foo bar":
                    v14 += v14
                }
            }
            switch 4294967295 {
            case -2147483648:
                fmt.Println(v2)
            default:
                var v31 int64 = 42
                _ = v31
            }
            v25 -= (((float32(float64(((1.5 + (v25 / v25)) - float32(float64(v25))))) / (v25 * 0.0)) / v25) * v25)
            _ = fi28
        }
        _ = v27
    } else {
        v32 := []int{}
        _ = v32
    }
    v14 = (v14 & 158)
    for fi33 := 0; fi33 < 19; fi33++ {
        for fi34 := 0; fi34 < 1; fi34++ {
            v35 := []int{}
            if false {
                v36 := GenS2{F0: 128}
                v37 := GenS2{F0: 184}
                v38 := GenS5{F0: 9223372036854775807}
                _ = v36
                _ = v37
                _ = v38
            }
            _ = fi34
            _ = v35
        }
        var v39 int8 = 51
        _ = fi33
        _ = v39
    }
    fmt.Println(v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
