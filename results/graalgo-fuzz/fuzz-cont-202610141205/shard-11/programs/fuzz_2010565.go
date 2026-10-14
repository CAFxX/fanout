package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint16 }

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]string)
    var v3 uint = 883
    var v4 float32 = 0.0
    var v5 string = ((("" + "foo bar") + ("a" + "")) + ((("foo bar" + "hello") + ("" + ("" + ("hello" + "\n\t")))) + (("" + (("" + "hello") + "")) + ("foo bar" + ((("foo bar" + "") + "a") + ("\n\t" + "hello"))))))
    for fi6 := 0; fi6 < 4; fi6++ {
        v2[2147483648] = ((v5 + "\n\t") + (v5 + v5))
        var v7 uint16 = 449
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 1000, F1: true, F2: 0}
    var v9 string = "foo bar"
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            fmt.Println(v5)
            var v12 uint32 = 292
            v13 := []int{(fi10 % (((-(fi10))) | 1)), 32767, fi10}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    if (true || (false && false)) {
        var v14 float64 = float64(1.5)
        v5 = (v9 + v5)
        _ = v14
    } else {
        var v15 bool = true
        _ = v15
    }
    var v16 string = v9
    var v17 uint16 = 0
    var v18 float64 = 1e10
    switch "" {
    case "hello":
        var v19 string = v5
        _ = v19
    case "foo bar":
        var v20 uint64 = (v1 % ((v1) | 1))
        _ = v20
    case "\n\t":
        fmt.Println(v9, v9, v16)
    }
    var v21 uint32 = 2147483647
    var v22 int8 = 33
    var v23 uint32 = v21
    if false {
        v24 := []int{4294967295, 42, 127}
        if ((v5 + (v16 + "\n\t")) <= (v16 + "")) {
            var v25 int64 = 42
            _ = v25
        } else {
            switch "" {
            case "":
                switch 32768 {
                case 42:
                    if false {
                        if (v17 <= (v17 ^ v17)) {
                            var v26 float32 = float32(v18)
                            _ = v26
                        }
                    }
                default:
                    fmt.Println(v22)
                }
            case "\n\t":
                v27 := v2[v21]
                _ = v27
            default:
                v5 = (v16 + "\n\t")
            }
        }
        _ = v24
    }
    var v28 float32 = ((((float32(v18) / v4) - 3.14159) - float32(float64(1e10))) - (((v4 - (v4 / v4)) / ((0.1 / v4) - (3.14159 - ((1e10 / (-2.25 - (3.14159 - v4))) + 0.0)))) / v4))
    fmt.Println(v16)
    v3 = 4294967295
    v29 := GenS1{F0: 42, F1: true, F2: (v17 ^ 2)}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
