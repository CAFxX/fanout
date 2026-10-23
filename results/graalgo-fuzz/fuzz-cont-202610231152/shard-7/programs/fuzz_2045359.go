package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint32; F2 int16 }

func main() {
    var v1 float32 = 1.5
    var v2 float32 = 0.0
    var v3 float64 = 1.5
    var v4 uint32 = 42
    v5 := GenS1{F0: -1}
    v4 += v4
    v1 = 0.0
    v4 = v4
    v6 := []int{9223372036854775807, 32768}
    var v7 uint16 = 1
    var v8 uint8 = uint8(v7)
    if true {
        v8 += v8
        v1 += 1e10
        switch ("" + ("a" + "\n\t")) {
        case "a":
            if (false || (!false)) {
                for fi9 := 0; fi9 < 18; fi9++ {
                    if len(v6) > 0 {
                        v6[0] = fi9
                    }
                    v4 += (v4 - v4)
                    switch "" {
                    case "":
                        v10 := GenS2{F0: (v8 * v8)}
                        _ = v10
                    default:
                        for fi11 := 0; fi11 < 10; fi11++ {
                            v12 := GenS1{F0: 1}
                            _ = fi11
                            _ = v12
                        }
                    }
                    _ = fi9
                }
                v3 += v3
            }
        }
    } else {
        switch -2147483648 {
        case 1:
            var v13 float64 = (3.14159 * (-2.25 / float64((v2 - v2))))
            _ = v13
        }
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int32 = -2147483648
            v7 -= 255
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS3{F0: (!true), F1: 42, F2: -171}
    var v17 uint32 = (v4 ^ (v4 * (v4 >> ((849) & 63))))
    v6 = append(v6, int(v4))
    var v18 bool = false
    for fi19 := 0; fi19 < 11; fi19++ {
        if len(v6) > 0 {
            v6[0] = fi19
        }
        _ = fi19
    }
    v4 = 291
    v20 := GenS1{F0: -397}
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
