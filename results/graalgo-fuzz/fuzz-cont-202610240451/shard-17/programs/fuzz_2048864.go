package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint64 = 885
    var v2 int = 42
    v3 := []int{v2}
    var v4 float32 = (1e10 + 3.14159)
    var v5 uint32 = 0
    var v6 float64 = float64((v4 + v4))
    if true {
        var v7 int16 = -895
        if true {
            if ((!true) || (255 < 127)) {
                v7 += -32768
                var v8 float64 = 3.14159
                _ = v8
            } else {
                var v9 int64 = 42
                _ = v9
            }
            if true {
                _ = len(v3)
            }
            fmt.Println(v5)
        } else {
            v10 := GenS1{F0: -128}
            _ = v10
        }
        switch ("foo bar" + "hello") {
        case "foo bar":
            var v11 uint16 = 0
            _ = v11
        case "\n\t":
            v6 -= v6
        default:
            v1 += v1
        }
        _ = v7
    }
    switch -181 {
    case 42:
        var v12 int64 = 1
        _ = v12
    case 127:
        if (!true) {
            v2 = v2
            if ((v1 ^ (v1 - v1)) == (v1 >> ((2) & 63))) {
                for fi13 := 0; fi13 < 16; fi13++ {
                    v3 = append(v3, (fi13 / ((v2) | 1)))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v14 := GenS1{F0: 40}
                    _ = fi13
                    _ = v14
                }
                for fi15 := 0; fi15 < 1; fi15++ {
                    switch "hello" {
                    case "a":
                        v2 += 32768
                    default:
                        for fi16 := 0; fi16 < 11; fi16++ {
                            v5 = 2
                            _ = fi16
                        }
                    }
                    fmt.Println(fi15, v2, v1)
                    _ = fi15
                }
                _ = len(v3)
            }
        }
    default:
        if ((^(v5)) > v5) {
            var v17 float64 = 0.1
            v1 += (v1 >> ((65535) & 63))
            _ = v17
        } else {
            v6 = 0.0
            fmt.Println(v6)
        }
    }
    v18 := GenS2{F0: -1, F1: 52}
    var v19 uint64 = (v1 - (^(v1)))
    v4 = v4
    v20 := GenS2{F0: -1, F1: 100}
    switch ("hello" + "a") {
    case "\n\t":
        for fi21 := 0; fi21 < 3; fi21++ {
            for fi22 := 0; fi22 < 9; fi22++ {
                if (-699 == 284) {
                    v5 += (v5 + 42)
                    v3 = append(v3, 32767)
                }
                _ = fi22
            }
            _ = fi21
        }
    case "":
        fmt.Println(v1, v19)
    case "a":
        fmt.Println(v19, v19)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
