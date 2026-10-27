package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16) int16 {
    v5 := []int{157}
    var v6 string = ("foo bar" + ("hello" + "hello"))
    var v7 uint32 = 233
    if false {
        v7 = 192
        v5 = append(v5, 9223372036854775807)
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int32 = 2147483647
    var v2 int16 = 1
    v3 := make(map[string]int)
    var v4 uint16 = 620
    var v8 bool = (1.5 < 3.14159)
    var v9 float64 = -2.25
    var v10 float32 = 0.1
    var v11 float32 = (0.0 / v10)
    var v12 int32 = v1
    switch "foo bar" {
    case "a":
        if (!(0 == 764)) {
            v13 := GenS1{F0: 151, F1: (v2 << ((696) & 63))}
            var v14 float64 = (1e10 - (-2.25 / -2.25))
            if v8 {
                v11 += ((1.5 * v10) + 0.0)
                v12 = (v1 ^ (v12 + (v12 + v12)))
            }
            _ = v13
            _ = v14
        } else {
            fmt.Println(v4)
            for fi15 := 0; fi15 < 7; fi15++ {
                switch "hello" {
                case "":
                    v16 := GenS1{F0: (fi15 >> ((887) & 63)), F1: genFn0(1)}
                    _ = v16
                case "\n\t":
                    var v17 uint32 = 2147483648
                    _ = v17
                }
                v1 -= (^(v1))
                _ = fi15
            }
        }
    case "foo bar":
        var v18 string = "\n\t"
        _ = v18
    case "\n\t":
        if false {
            switch ("\n\t" + "\n\t") {
            case "":
                v19 := GenS1{F0: 4294967295, F1: v2}
                _ = v19
            default:
                v20 := GenS2{F0: genFn0(v4)}
                _ = v20
            }
        }
    }
    if (v1 <= v12) {
        v9 = ((-2.25 + float64(0.0)) + float64(v10))
    }
    v3[("" + "")] = 32767
    if v8 {
        _ = len(v3)
        v21 := GenS1{F0: 127, F1: (v2 >> ((18446744073709551615) & 63))}
        _ = v21
    }
    var v22 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
