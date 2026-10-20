package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }

func main() {
    var v1 int32 = -1
    var v2 bool = (65535 <= 42)
    v3 := []string{"\n\t", "foo bar", "hello"}
    var v4 int32 = (v1 ^ v1)
    var v5 uint16 = 0
    v6 := GenS1{F0: 1, F1: (v1 - v1), F2: v5}
    switch int(v5) {
    case 32767:
        _ = len(v3)
    }
    if false {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        if false {
            switch 128 {
            case 158:
                for fi7 := 0; fi7 < 2; fi7++ {
                    switch ("" + "hello") {
                    case "hello":
                        var v8 int16 = -191
                        _ = v8
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    v4 -= int32(v5)
                    _ = fi7
                }
            case -236:
                var v9 float32 = -2.25
                _ = v9
            }
            fmt.Println(v2)
        } else {
            if (194 == 12) {
                switch "\n\t" {
                case "hello":
                    var v10 float32 = 0.0
                    _ = v10
                case "foo bar":
                    v2 = v2
                case "":
                    v4 = v4
                default:
                    var v11 int8 = 127
                    _ = v11
                }
            }
        }
        v12 := GenS1{F0: 1000, F1: v4, F2: v5}
        _ = v12
    }
    var v13 uint32 = 42
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 string = "hello"
        fmt.Println(v13)
        _ = len(v3)
        _ = fi14
        _ = v15
    }
    if (v2 && (-2.25 != 1.5)) {
        var v16 uint16 = (^(v5))
        v17 := GenS1{F0: 641, F1: 228, F2: 169}
        _ = v16
        _ = v17
    }
    v5 = v5
    var v18 uint = 1
    var v19 uint32 = v13
    var v20 int16 = -314
    for fi21 := 0; fi21 < 14; fi21++ {
        _ = len(v3)
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 float32 = (0.1 / float32(-2.25))
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
