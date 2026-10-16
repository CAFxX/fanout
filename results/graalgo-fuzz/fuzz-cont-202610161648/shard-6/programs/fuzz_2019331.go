package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint32; F2 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 uint32; F2 int32 }

func main() {
    var v1 int8 = 100
    var v2 uint64 = 1
    var v3 int = -1
    v4 := []int{-722}
    for fi5 := 0; fi5 < 16; fi5++ {
        if true {
            var v6 float32 = 1e10
            var v7 string = "foo bar"
            switch (v7 + v7) {
            case "foo bar":
                v8 := GenS1{F0: 723, F1: "a", F2: 42}
                _ = v8
            case "\n\t":
                if true {
                    if len(v4) > 0 {
                        v4[0] = 775
                    }
                    var v9 int16 = -460
                    v6 = 1.5
                    _ = v9
                } else {
                    fmt.Println(v2, v3)
                    if (((-2.25 / (-2.25 - (v6 + 3.14159))) * v6) == (v6 + v6)) {
                        if len(v4) > 0 {
                            v4[0] = (v3 ^ fi5)
                        }
                        v4 = append(v4, (^(v3)))
                    } else {
                        var v10 int = v3
                        _ = len(v4)
                        var v11 float64 = float64(v6)
                        _ = v10
                        _ = v11
                    }
                }
            case "hello":
                for fi12 := 0; fi12 < 10; fi12++ {
                    fmt.Println(v3)
                    _ = fi12
                }
            }
            _ = v6
            _ = v7
        } else {
            v13 := GenS1{F0: -32768, F1: ((("" + "foo bar") + ("foo bar" + (("" + (("\n\t" + "foo bar") + "\n\t")) + "foo bar"))) + (("hello" + "") + "")), F2: -9223372036854775808}
            _ = v13
        }
        var v14 int8 = v1
        _ = fi5
        _ = v14
    }
    var v15 uint64 = 4294967295
    v16 := GenS2{F0: (false || false), F1: 42}
    v3 -= 127
    if len(v4) > 0 {
        v4[0] = (-(v3))
    }
    if false {
        var v17 float32 = 0.1
        _ = v17
    } else {
        var v18 uint16 = 2
        v19 := GenS3{F0: (v2 - v2), F1: 619, F2: "foo bar"}
        for fi20 := 0; fi20 < 6; fi20++ {
            v4 = append(v4, 4294967295)
            var v21 float64 = (3.14159 / 0.1)
            v22 := GenS4{F0: (v2 & (v15 & (^(v15))))}
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = v18
        _ = v19
    }
    v23 := GenS5{F0: 295, F1: 1, F2: -823}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
