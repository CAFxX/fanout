package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint32 = 42
    var v2 int32 = 801
    v3 := []string{("a" + "")}
    if false {
        v4 := GenS1{F0: 0, F1: 1}
        _ = v4
    }
    var v5 uint16 = 1
    v1 += (v1 ^ 2147483648)
    fmt.Println(v2)
    v6 := GenS2{F0: 2147483648, F1: v2, F2: v1}
    for fi7 := 0; fi7 < 6; fi7++ {
        _ = len(v3)
        _ = fi7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        switch "foo bar" {
        case "foo bar":
            for fi9 := 0; fi9 < 17; fi9++ {
                switch "a" {
                case "foo bar":
                    v5 += (v5 / ((v5) | 1))
                case "a":
                    var v10 int8 = int8(v1)
                    _ = v10
                default:
                    if ((!true) || false) {
                        v11 := GenS1{F0: 527, F1: 0}
                        _ = v11
                    } else {
                        v12 := GenS3{F0: (v2 ^ v2)}
                        v3 = append(v3, "foo bar")
                        _ = v12
                    }
                }
                switch ("" + ("hello" + "a")) {
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "":
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        if true {
                            v13 := GenS1{F0: 4294967295, F1: 128}
                            _ = v13
                        } else {
                            var v14 float32 = 1.5
                            var v15 int32 = v2
                            v14 += v14
                            _ = v14
                            _ = v15
                        }
                        v3 = append(v3, "\n\t")
                    }
                }
                _ = fi9
            }
        case "":
            var v16 uint16 = v5
            _ = v16
        }
        _ = fi8
    }
    var v17 int32 = -2147483648
    var v18 int32 = (v2 >> ((1) & 63))
    var v19 float64 = float64(float32((((0.0 - (1.5 + 0.1)) / float64(1.5)) / (((0.1 - 1e10) / 1e10) + (-2.25 / ((3.14159 / 1.5) * 1.5))))))
    v17 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
