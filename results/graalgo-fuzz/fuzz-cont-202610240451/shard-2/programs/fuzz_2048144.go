package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint16; F1 string; F2 uint64 }

func main() {
    var v1 float64 = 0.0
    var v2 uint8 = 127
    v3 := make(map[int]bool)
    var v4 float64 = (v1 / (0.1 + v1))
    if true {
        var v5 uint64 = 689
        v5 = (v5 ^ 4294967295)
        _ = v5
    }
    v1 -= 3.14159
    var v6 uint32 = 665
    v7 := []int{128}
    v8 := GenS1{F0: -128, F1: 4294967295}
    v3[32768] = (!false)
    switch "a" {
    case "a":
        var v9 uint16 = 709
        _ = v9
    default:
        v1 = v4
    }
    fmt.Println(v1)
    v3[-418] = (255 == 18446744073709551615)
    v4 -= v1
    var v10 int8 = 101
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS1{F0: (v10 - (v10 & (v10 + v10))), F1: 4294967295}
        _ = fi11
        _ = v12
    }
    var v13 uint8 = v2
    if (48 == v10) {
        if len(v7) > 0 {
            v7[0] = 9223372036854775807
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            v13 -= (v13 * 0)
            var v15 bool = ((true || false) && false)
            _ = fi14
            _ = v15
        }
        if (true || true) {
            _ = len(v3)
            switch 127 {
            case 42:
                if false {
                    fmt.Println(v6, v6)
                }
            case -978:
                switch 0 {
                case 32768:
                    v4 = v4
                default:
                    for fi16 := 0; fi16 < 5; fi16++ {
                        var v17 uint8 = (v2 ^ v13)
                        _ = fi16
                        _ = v17
                    }
                }
            }
            v13 -= (v13 + v13)
        }
    }
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        v1 += (v1 - v1)
    case "hello":
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 11; fi19++ {
                switch fi18 {
                case 987:
                    if false {
                        v2 -= (v2 + 255)
                        v20 := GenS1{F0: v10, F1: 809}
                        _ = v20
                    } else {
                        var v21 int32 = 1
                        v22 := GenS2{F0: 360, F1: ("\n\t" + "foo bar"), F2: 1}
                        v23 := GenS1{F0: v10, F1: 2}
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                case -52:
                    var v24 int = fi19
                    _ = v24
                default:
                    var v25 uint8 = 1
                    _ = v25
                }
                v4 -= ((float64(0.1) + float64(0.0)) * 1.5)
                var v26 string = "\n\t"
                _ = fi19
                _ = v26
            }
            _ = fi18
        }
    case "":
        var v27 int8 = v10
        _ = v27
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
