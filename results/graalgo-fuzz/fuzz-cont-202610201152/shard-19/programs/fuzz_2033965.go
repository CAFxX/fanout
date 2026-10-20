package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int16 }
type GenS2 struct { F0 bool; F1 string; F2 int32 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int16 = -32768
    var v2 uint32 = 0
    var v3 uint16 = 2
    var v4 float64 = 3.14159
    var v5 int8 = 35
    v6 := []int{402, 127, 1, 42}
    v7 := GenS1{F0: 386, F1: (v3 * (^(v3))), F2: v1}
    v8 := GenS2{F0: (-599 == 1), F1: ("foo bar" + ("hello" + "a")), F2: -643}
    v3 = (v3 % ((v3) | 1))
    v4 = 3.14159
    fmt.Println(v2, v4)
    fmt.Println(v5, v4)
    var v9 float32 = -2.25
    var v10 int16 = (v1 - 227)
    var v11 float64 = v4
    fmt.Println(v5, v9, v11)
    var v12 uint16 = v3
    for fi13 := 0; fi13 < 19; fi13++ {
        switch (("a" + "a") + ("a" + ("hello" + ("a" + "a")))) {
        case "\n\t":
            if false {
                fmt.Println(v3)
                v9 -= float32((1e10 / (((v11 + float64(v9)) / ((v4 * (float64((float32((3.14159 * ((3.14159 - -2.25) + (float64((v9 + float32(float64((3.14159 + v9))))) + 1e10)))) - (1.5 - v9))) * v4)) / (0.0 + float64(((v9 - v9) * -2.25))))) * 0.0)))
            }
        case "hello":
            switch "\n\t" {
            case "a":
                switch ("a" + "\n\t") {
                case "":
                    fmt.Println(v3, v5, fi13)
                default:
                    v6 = append(v6, 629)
                }
            default:
                fmt.Println(v4, v9, v4)
            }
        case "":
            for fi14 := 0; fi14 < 18; fi14++ {
                v4 = 1.5
                v15 := GenS3{F0: v10}
                switch ("foo bar" + ("" + ("hello" + ((((((("a" + "a") + "hello") + "a") + "hello") + "\n\t") + "foo bar") + "hello")))) {
                case "":
                    v6 = append(v6, int(v10))
                default:
                    var v16 int16 = (v10 ^ (v10 | (v1 % ((v1) | 1))))
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
        default:
            for fi17 := 0; fi17 < 9; fi17++ {
                fmt.Println(v2, v10, v10)
                var v18 int = fi13
                _ = fi17
                _ = v18
            }
        }
        _ = fi13
    }
    var v19 float32 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
