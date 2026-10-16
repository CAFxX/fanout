package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 bool; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 0.0
    var v3 uint16 = 2
    var v4 int16 = -1
    fmt.Println(v4, v2)
    v3 = v3
    var v5 int = 42
    var v6 uint64 = 0
    switch "foo bar" {
    case "\n\t":
        switch "\n\t" {
        case "hello":
            if (100 != 0) {
                if (true || (1 > 251)) {
                    for fi7 := 0; fi7 < 4; fi7++ {
                        for fi8 := 0; fi8 < 16; fi8++ {
                            v5 += v5
                            if (42 < -128) {
                                var v9 int32 = 0
                                _ = v9
                            }
                            v1[(v5 - (fi7 ^ (v5 | fi8)))] = ("hello" + "")
                            _ = fi8
                        }
                        v10 := []int{}
                        v1[-601] = "\n\t"
                        _ = fi7
                        _ = v10
                    }
                    if (false && false) {
                        var v11 int32 = 116
                        fmt.Println(v2, v4, v11)
                        _ = v11
                    }
                    v12 := GenS1{F0: false}
                    _ = v12
                }
                v13 := GenS1{F0: true}
                var v14 bool = true
                _ = v13
                _ = v14
            }
        case "foo bar":
            if true {
                var v15 uint64 = (v6 / (((^(v6))) | 1))
                v2 += v2
                v6 += (v15 ^ (v6 % (((^(v6))) | 1)))
                _ = v15
            }
        default:
            v16 := GenS2{F0: (v6 & (v6 | 429)), F1: (2 >= 4294967295), F2: -161}
            _ = v16
        }
    case "a":
        delete(v1, v5)
    case "":
        v3 = v3
    default:
        fmt.Println(v5)
    }
    switch (("\n\t" + "a") + "\n\t") {
    case "":
        if true {
            v4 += (-(v4))
        } else {
            v3 = 255
        }
    case "foo bar":
        for fi17 := 0; fi17 < 9; fi17++ {
            v5 -= fi17
            _ = fi17
        }
    }
    v6 -= 0
    v18 := v1[int(v6)]
    v3 = v3
    v18 = ("hello" + (v18 + ""))
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
