package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }

func main() {
    var v1 float32 = (3.14159 * 0.0)
    var v2 string = "\n\t"
    var v3 uint8 = 0
    fmt.Println(v2, v1, v2)
    fmt.Println(v3, v1)
    for fi4 := 0; fi4 < 15; fi4++ {
        if false {
            if true {
                v5 := []int{}
                var v6 float64 = (-2.25 + 0.0)
                _ = v5
                _ = v6
            }
        }
        _ = fi4
    }
    var v7 uint32 = 981
    var v8 int8 = 1
    v3 = v3
    v8 += v8
    v9 := GenS1{F0: 127, F1: 4294967295, F2: v2}
    var v10 uint = 17
    switch 229 {
    case 815:
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 float32 = (1.5 * float32(0.1))
            switch fi11 {
            case 0:
                switch fi11 {
                case 196:
                    var v13 bool = (!(0 >= -2147483648))
                    _ = v13
                }
            default:
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := GenS1{F0: (fi11 * (^(fi14))), F1: (v10 & v10), F2: (v2 + v2)}
                    _ = fi14
                    _ = v15
                }
            }
            if (float64(float32((0.0 / 3.14159))) > (1e10 * 1.5)) {
                v16 := GenS1{F0: fi11, F1: 0, F2: (v2 + "\n\t")}
                fmt.Println(v8, fi11, v8)
                switch fi11 {
                case -9223372036854775808:
                    var v17 int64 = -9223372036854775808
                    _ = v17
                default:
                    v2 = v2
                }
                _ = v16
            }
            _ = fi11
            _ = v12
        }
    }
    var v18 uint8 = v3
    v2 = "a"
    switch (v2 + v2) {
    case "\n\t":
        if (("\n\t" + v2) != "foo bar") {
            v10 = 255
            var v19 uint = 18446744073709551615
            _ = v19
        }
    case "foo bar":
        v8 += (v8 ^ (v8 * (-(v8))))
    case "":
        if true {
            switch 0 {
            case 42:
                v1 = float32((1.5 - (0.1 / 0.1)))
            case 862:
                v18 = uint8(v7)
            case 127:
                var v20 int8 = 127
                _ = v20
            default:
                var v21 uint = (v10 * 65535)
                _ = v21
            }
            v22 := GenS1{F0: 128, F1: (^(v10)), F2: "a"}
            _ = v22
        }
    }
    var v23 int = 2147483647
    v23 = v23
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
