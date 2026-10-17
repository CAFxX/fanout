package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }

func main() {
    v1 := []string{("" + "\n\t")}
    var v2 int16 = 0
    var v3 uint = 65535
    var v4 float64 = ((-2.25 * -2.25) + 1.5)
    v5 := make(map[int]int)
    v6 := []string{}
    if true {
        v7 := GenS1{F0: 42, F1: ("\n\t" + "\n\t")}
        _ = v7
    } else {
        v8 := GenS1{F0: 255, F1: "hello"}
        _ = v8
    }
    var v9 uint64 = 407
    var v10 float32 = 0.0
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            switch fi12 {
            case -1:
                v2 = v2
            case 9223372036854775807:
                var v13 int8 = -126
                _ = v13
            case 127:
                v14 := GenS1{F0: 716, F1: "foo bar"}
                _ = v14
            }
            _ = fi12
        }
        _ = fi11
    }
    v10 = v10
    for fi15 := 0; fi15 < 8; fi15++ {
        fmt.Println(v2, v2)
        v16 := GenS1{F0: 0, F1: ("a" + "")}
        _ = fi15
        _ = v16
    }
    v4 = v4
    fmt.Println(v10, v9, v9)
    v4 = (0.0 * float64(((v10 / v10) / v10)))
    for fi17 := 0; fi17 < 14; fi17++ {
        v3 += 330
        _ = fi17
    }
    v18 := GenS1{F0: 285, F1: "foo bar"}
    for fi19 := 0; fi19 < 15; fi19++ {
        for fi20 := 0; fi20 < 4; fi20++ {
            fmt.Println(v4, v10, v4)
            if (float32(((float64(((v10 - 0.0) * v10)) / ((1e10 - float64(v10)) + v4)) / (v4 + -2.25))) >= v10) {
                var v21 int32 = -299
                _ = v21
            }
            _ = fi20
        }
        v22 := GenS1{F0: 32767, F1: (("" + "foo bar") + ("a" + ""))}
        _ = fi19
        _ = v22
    }
    v1 = append(v1, "")
    var v23 bool = (true || false)
    v24 := GenS2{F0: 1, F1: 851, F2: "a"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
