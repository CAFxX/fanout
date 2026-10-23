package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint }
type GenS2 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 string = ("a" + ((("hello" + "") + "a") + ("\n\t" + "\n\t")))
    var v2 int32 = 706
    v3 := make(map[int]string)
    v4 := []uint32{0}
    fmt.Println(v1)
    v1 = (("\n\t" + (v1 + v1)) + "foo bar")
    if (v2 == 358) {
        v1 = "\n\t"
        _ = len(v4)
    }
    for fi5 := 0; fi5 < 14; fi5++ {
        if false {
            if false {
                v6 := GenS1{F0: (fi5 % (((^(fi5))) | 1)), F1: 481, F2: 18446744073709551615}
                _ = v6
            }
            var v7 bool = true
            var v8 float32 = (float32(((-2.25 - 0.0) - 1.5)) - 3.14159)
            _ = v7
            _ = v8
        }
        if len(v4) > 0 {
            v4[0] = 65535
        }
        if true {
            var v9 float32 = ((0.0 / (float32(float64(float32((1.5 / (0.1 - (-2.25 - ((1.5 - -2.25) / 1.5))))))) * (0.1 - (1e10 + 1e10)))) * (((float32(1.5) * ((0.0 - float32(1e10)) / 0.1)) + -2.25) * 1e10))
            _ = v9
        }
        _ = fi5
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        if len(v4) > 0 {
            v4[0] = 39
        }
        var v11 float32 = (0.0 - (0.1 / (0.1 - 1.5)))
        _ = fi10
        _ = v11
    }
    fmt.Println(v2, v1)
    fmt.Println(v2)
    v12 := GenS1{F0: 32768, F1: 124, F2: 18446744073709551615}
    v13 := GenS2{F0: 1, F1: 151}
    if false {
        v1 = v1
        if true {
            switch "foo bar" {
            case "\n\t":
                v2 = v2
            case "foo bar":
                v1 = "foo bar"
            case "":
                for fi14 := 0; fi14 < 1; fi14++ {
                    if true {
                        var v15 uint32 = 4294967295
                        var v16 int8 = -113
                        v2 -= (-(v2))
                        _ = v15
                        _ = v16
                    }
                    if true {
                        fmt.Println(v1, v1)
                    }
                    _ = len(v3)
                    _ = fi14
                }
            }
        } else {
            fmt.Println(v1, v2)
        }
    }
    var v17 uint16 = 42
    v18 := GenS2{F0: 0, F1: v17}
    delete(v3, 4294967295)
    v19 := GenS1{F0: 4294967295, F1: 1, F2: 553}
    fmt.Println(v2)
    delete(v3, -452)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
