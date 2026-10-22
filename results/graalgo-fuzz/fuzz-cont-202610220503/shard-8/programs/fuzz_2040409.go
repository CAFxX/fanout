package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint16 }

func main() {
    v1 := []int{}
    var v2 bool = false
    var v3 int16 = 0
    var v4 int64 = -493
    _ = len(v1)
    for fi5 := 0; fi5 < 18; fi5++ {
        if v2 {
            var v6 float64 = 0.1
            for fi7 := 0; fi7 < 17; fi7++ {
                v6 = v6
                v3 -= (v3 >> ((18446744073709551615) & 63))
                _ = fi7
            }
            _ = v6
        } else {
            fmt.Println(fi5)
            switch ("hello" + "\n\t") {
            case "foo bar":
                var v8 uint = 65535
                _ = v8
            case "hello":
                fmt.Println(fi5, v4)
            }
            if (4294967295 == 42) {
                var v9 int = 32767
                switch 127 {
                case 127:
                    fmt.Println(fi5)
                case 356:
                    switch fi5 {
                    case 339:
                        fmt.Println(fi5, v9, fi5)
                    case 2147483647:
                        if len(v1) > 0 {
                            v1[0] = (^(fi5))
                        }
                    default:
                        var v10 int64 = 0
                        _ = v10
                    }
                default:
                    var v11 uint64 = 0
                    _ = v11
                }
                _ = v9
            }
        }
        v2 = (-2147483648 > -394)
        var v12 uint64 = 395
        _ = fi5
        _ = v12
    }
    var v13 float32 = 3.14159
    v13 = v13
    var v14 float64 = -2.25
    if len(v1) > 0 {
        v1[0] = 256
    }
    v3 -= (v3 - -1)
    v15 := GenS1{F0: v3, F1: 809, F2: 201}
    fmt.Println(v14, v13, v14)
    var v16 bool = (false || v2)
    var v17 uint8 = 42
    var v18 bool = true
    v19 := GenS1{F0: (v3 ^ -578), F1: 255, F2: 42}
    v13 = ((0.1 * ((1.5 / 1.5) / v13)) - float32(v14))
    _ = len(v1)
    fmt.Println(v18, v14, v18)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
