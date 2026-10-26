package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }

func main() {
    var v1 float64 = 1e10
    var v2 float32 = ((3.14159 - (float32(v1) * (((float32(v1) + -2.25) * 1.5) / ((3.14159 * -2.25) * 1e10)))) - (((float32((3.14159 * (float64(0.1) + v1))) + (1.5 / 3.14159)) - 0.1) / 1e10))
    var v3 bool = ((!false) || ((-32768 < 42) && false))
    var v4 uint32 = 973
    var v5 int64 = 245
    var v6 string = "foo bar"
    var v7 int8 = 1
    var v8 string = ""
    v9 := []int{32768, 2147483647}
    v6 = v8
    if true {
        for fi10 := 0; fi10 < 11; fi10++ {
            switch "" {
            case "\n\t":
                v1 = v1
            case "foo bar":
                v4 = 2147483647
            case "a":
                if true {
                    if true {
                        for fi11 := 0; fi11 < 15; fi11++ {
                            v9 = append(v9, 2147483647)
                            _ = fi11
                        }
                    }
                    v9 = append(v9, (^(fi10)))
                    v12 := GenS1{F0: ((v8 + "a") + "\n\t"), F1: v7, F2: 1}
                    _ = v12
                }
            }
            _ = fi10
        }
        var v13 uint16 = 1
        _ = v13
    }
    v2 = (((v2 / v2) - 1e10) - (((v2 - float32(1.5)) / (0.0 - 1e10)) / v2))
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := GenS2{F0: (v5 & (^(v5)))}
        v16 := GenS3{F0: (fi14 >> ((0) & 63))}
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 int32 = -832
            _ = fi17
            _ = v18
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 int8 = (v7 - -1)
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
