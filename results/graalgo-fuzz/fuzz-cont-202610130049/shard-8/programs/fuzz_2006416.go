package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 bool }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int }

func main() {
    var v1 int8 = 57
    var v2 int = 32767
    var v3 float32 = 0.0
    var v4 int64 = 42
    v5 := []int{(v2 | (v2 >> ((0) & 63)))}
    v6 := make(map[string]string)
    v7 := GenS1{F0: v2, F1: (629 != 42), F2: (!true)}
    v4 = v4
    for fi8 := 0; fi8 < 8; fi8++ {
        v1 += v1
        if (true || false) {
            fmt.Println(v3)
        } else {
            v6[("\n\t" + ("a" + "\n\t"))] = ((("" + "foo bar") + "hello") + "foo bar")
        }
        v3 = v3
        _ = fi8
    }
    v3 = (((1e10 + v3) - (3.14159 * (v3 / ((v3 - (1.5 / (v3 / (1e10 * v3)))) + v3)))) * (-2.25 * (v3 / 1.5)))
    var v9 int64 = (^(v4))
    v10 := GenS2{F0: 2, F1: (false && ((-649 >= -281) && true))}
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 6; fi12++ {
            if len(v5) > 0 {
                v5[0] = fi12
            }
            _ = fi12
        }
        v13 := GenS2{F0: 2, F1: true}
        var v14 float64 = -2.25
        _ = fi11
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        switch "\n\t" {
        case "\n\t":
            switch -1 {
            case 127:
                v9 = (v4 | v4)
            case -1:
                if len(v5) > 0 {
                    v5[0] = v2
                }
            case 9223372036854775807:
                fmt.Println(fi15, v1)
            default:
                for fi16 := 0; fi16 < 7; fi16++ {
                    if len(v5) > 0 {
                        v5[0] = (v2 ^ -422)
                    }
                    v6[("" + ((("foo bar" + "hello") + "") + ""))] = "a"
                    _ = fi16
                }
            }
        case "foo bar":
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 uint8 = 127
                var v19 int16 = 965
                switch fi15 {
                case 9223372036854775807:
                    v19 = v19
                case -976:
                    if true {
                        v19 += (v19 * (^(v19)))
                    } else {
                        v2 = (v2 << ((255) & 63))
                        v4 = v4
                    }
                case 680:
                    if false {
                        var v20 bool = true
                        _ = v20
                    }
                }
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        fmt.Println(v1, v9, v9)
        _ = len(v6)
        _ = fi15
    }
    v1 -= -1
    v21 := GenS3{F0: (v2 | (^(v2))), F1: ""}
    var v22 uint = 18446744073709551615
    if true {
        var v23 uint64 = 2
        v24 := v6["a"]
        var v25 uint = (v22 + 255)
        _ = v23
        _ = v24
        _ = v25
    }
    v2 += v2
    fmt.Println(v9, v2, v1)
    v26 := GenS4{F0: -780}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
