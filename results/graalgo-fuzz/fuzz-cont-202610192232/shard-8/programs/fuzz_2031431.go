package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 string }

func main() {
    var v1 int32 = 291
    var v2 uint32 = 2147483647
    var v3 string = ""
    v4 := make(map[uint32]bool)
    var v5 uint32 = (^(v2))
    v6 := GenS1{F0: ((!false) || ((1.5 - (0.1 / ((1.5 - 0.0) - 0.0))) > 0.0)), F1: 791, F2: v3}
    for fi7 := 0; fi7 < 13; fi7++ {
        v2 -= 580
        switch (fi7 * (fi7 / ((fi7) | 1))) {
        case -337:
            switch (fi7 / (((-(fi7))) | 1)) {
            case 712:
                switch v3 {
                case "\n\t":
                    v8 := []int{(fi7 | (fi7 ^ 112)), 2147483647, fi7}
                    _ = v8
                }
            }
        case 1:
            if false {
                if false {
                    switch ("foo bar" + v3) {
                    case "":
                        v9 := GenS1{F0: false, F1: -87, F2: v3}
                        _ = v9
                    case "\n\t":
                        for fi10 := 0; fi10 < 5; fi10++ {
                            v3 = v3
                            v4[v2] = ((v1 / (((^(v1))) | 1)) >= v1)
                            _ = fi10
                        }
                    default:
                        v1 += (v1 - v1)
                    }
                    v5 += v2
                }
                if (false || true) {
                    var v11 uint64 = 4294967295
                    for fi12 := 0; fi12 < 7; fi12++ {
                        var v13 int = (^(fi12))
                        var v14 uint = 18446744073709551615
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = v11
                }
            }
        default:
            var v15 int = (-(fi7))
            _ = v15
        }
        _ = fi7
    }
    fmt.Println(v1, v3, v3)
    if true {
        v2 -= (v5 - v5)
        if true {
            var v16 int8 = -90
            var v17 int8 = (v16 * (v16 ^ v16))
            _ = v16
            _ = v17
        }
    }
    var v18 bool = true
    var v19 int64 = 9223372036854775807
    v1 -= v1
    if (73 < 42) {
        var v20 float64 = 3.14159
        _ = v20
    } else {
        for fi21 := 0; fi21 < 18; fi21++ {
            delete(v4, v5)
            v19 -= v19
            _ = fi21
        }
        fmt.Println(v19, v18, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
