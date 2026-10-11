package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int }
type GenS2 struct { F0 int; F1 bool; F2 uint32 }

func main() {
    v1 := []uint32{}
    var v2 string = "a"
    var v3 float64 = 1e10
    var v4 uint64 = 18446744073709551615
    v5 := []string{"\n\t"}
    if (!(4294967295 >= 2147483648)) {
        v2 = v2
        var v6 float32 = (0.0 * (0.1 * (1.5 / (1e10 + (3.14159 - 3.14159)))))
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        switch (-(fi7)) {
        case 2147483647:
            switch -1 {
            case 469:
                switch v2 {
                case "":
                    switch "" {
                    case "a":
                        v4 = (v4 ^ 261)
                    case "foo bar":
                        v4 = v4
                    case "":
                        for fi8 := 0; fi8 < 14; fi8++ {
                            v3 = v3
                            v3 = (v3 * float64(((((1e10 + 1e10) + 0.1) * 3.14159) + 1.5)))
                            var v9 int32 = 2147483647
                            v10 := GenS1{F0: "a", F1: 0, F2: (fi8 * (-(fi8)))}
                            _ = fi8
                            _ = v10
                            _ = v9
                        }
                    default:
                        var v11 uint32 = 65535
                        _ = v11
                    }
                case "\n\t":
                    for fi12 := 0; fi12 < 6; fi12++ {
                        if len(v1) > 0 {
                            v1[0] = 42
                        }
                        _ = fi12
                    }
                }
            }
        case -132:
            var v13 uint16 = 1
            _ = v13
        default:
            for fi14 := 0; fi14 < 9; fi14++ {
                fmt.Println(v2, v3)
                if false {
                    v3 += float64((0.1 * 3.14159))
                }
                var v15 float32 = 0.0
                switch (v2 + v2) {
                case "foo bar":
                    switch v2 {
                    case "foo bar":
                        if (false && (127 != fi7)) {
                            v3 += 1e10
                        }
                    default:
                        var v16 int8 = 0
                        _ = v16
                    }
                }
                _ = fi14
                _ = v15
            }
        }
        _ = fi7
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        _ = len(v5)
        _ = fi17
    }
    if true {
        v18 := GenS2{F0: 128, F1: false, F2: 42}
        _ = v18
    }
    fmt.Println(v3, v3, v4)
    var v19 float32 = (float32(((((0.1 + -2.25) / v3) * float64((0.0 * -2.25))) / v3)) * 1.5)
    var v20 string = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v21 uint16 = 2
    v20 = v2
    var v22 bool = false
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v23 uint64 = (v4 ^ (v4 >> ((1) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
