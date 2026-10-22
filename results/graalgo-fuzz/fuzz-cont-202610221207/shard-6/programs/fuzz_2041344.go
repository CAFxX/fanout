package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []int{}
    v2 := []int{4294967295, 2147483647}
    var v3 string = ""
    var v4 uint64 = 621
    if true {
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 int = int(v4)
            _ = fi5
            _ = v6
        }
        v2 = append(v2, 0)
    } else {
        v3 = v3
        v7 := GenS1{F0: 255}
        _ = v7
    }
    v1 = append(v1, -9223372036854775808)
    var v8 float64 = 1e10
    v4 = 2
    for fi9 := 0; fi9 < 8; fi9++ {
        switch -787 {
        case 32768:
            fmt.Println(v8, v4)
        case -9223372036854775808:
            if len(v1) > 0 {
                v1[0] = fi9
            }
        case -738:
            switch v3 {
            case "hello":
                var v10 int16 = 32767
                _ = v10
            }
        default:
            var v11 uint32 = 1
            _ = v11
        }
        v8 = v8
        _ = fi9
    }
    v12 := GenS1{F0: 42}
    _ = len(v1)
    if (false || (false && true)) {
        switch 790 {
        case 0:
            fmt.Println(v3, v3)
        case 4294967295:
            fmt.Println(v3, v8, v4)
        case 318:
            var v13 bool = false
            _ = v13
        default:
            for fi14 := 0; fi14 < 9; fi14++ {
                switch v3 {
                case "foo bar":
                    v2 = append(v2, fi14)
                case "\n\t":
                    v8 = (float64(((float32(1.5) * (-2.25 + (0.1 + float32((v8 * v8))))) / 3.14159)) - ((v8 / 0.1) / ((v8 / float64(-2.25)) + v8)))
                default:
                    for fi15 := 0; fi15 < 4; fi15++ {
                        var v16 int8 = -99
                        _ = fi15
                        _ = v16
                    }
                }
                var v17 uint8 = 132
                _ = fi14
                _ = v17
            }
        }
        v3 = "a"
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v3 = v3
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v18 uint8 = 128
    v18 = (v18 - v18)
    var v19 int = 993
    v4 += v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
