package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint; F1 uint64 }

func main() {
    v1 := []string{"foo bar", "hello", ("foo bar" + ("\n\t" + ""))}
    v2 := make(map[int]string)
    var v3 uint64 = 2
    var v4 int64 = 43
    if true {
        if (!(((uint16(v3) >= 1) && false) && (!false))) {
            v5 := GenS1{F0: false, F1: int64(v3)}
            switch -459 {
            case -9223372036854775808:
                var v6 int = -952
                _ = v6
            }
            _ = len(v1)
            _ = v5
        }
        v3 = (v3 >> ((4294967295) & 63))
        switch 32768 {
        case -9223372036854775808:
            for fi7 := 0; fi7 < 1; fi7++ {
                var v8 int32 = 839
                for fi9 := 0; fi9 < 8; fi9++ {
                    var v10 string = "foo bar"
                    _ = fi9
                    _ = v10
                }
                _ = fi7
                _ = v8
            }
        case -2147483648:
            for fi11 := 0; fi11 < 6; fi11++ {
                v1 = append(v1, ("" + "hello"))
                _ = fi11
            }
        case 32767:
            v12 := GenS1{F0: false, F1: v4}
            _ = v12
        default:
            v13 := GenS1{F0: false, F1: (^(v4))}
            _ = v13
        }
    } else {
        v3 -= (v3 / ((42) | 1))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v14 float32 = 1e10
    for fi15 := 0; fi15 < 12; fi15++ {
        switch ("hello" + "foo bar") {
        case "a":
            v14 -= 0.0
        case "":
            var v16 float64 = 3.14159
            _ = v16
        case "\n\t":
            var v17 uint = 18446744073709551615
            _ = v17
        }
        v18 := GenS1{F0: (!true), F1: 42}
        _ = fi15
        _ = v18
    }
    v19 := GenS1{F0: (42 < 1), F1: (v4 / ((v4) | 1))}
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 int = 1
        v22 := GenS2{F0: 1, F1: 4294967295}
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 uint16 = 42
        _ = fi23
        _ = v24
    }
    v14 += v14
    var v25 int = -779
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
