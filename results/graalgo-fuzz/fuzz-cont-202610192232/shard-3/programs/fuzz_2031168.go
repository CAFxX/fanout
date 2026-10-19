package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int }
type GenS2 struct { F0 uint16; F1 int8 }
func genFn0() uint {
    var v4 int64 = -1
    switch 9223372036854775807 {
    case 242:
        v5 := []int{-1, -337}
        _ = v5
    }
    _ = v4
    return 4294967295
}

func main() {
    var v1 float32 = (-2.25 + 1.5)
    v2 := []int{}
    var v3 float64 = ((0.1 + ((float64((v1 / 1.5)) / 0.1) - ((float64(1.5) * 0.1) - 1e10))) / -2.25)
    var v6 int32 = 206
    v3 -= (v3 + 3.14159)
    switch 237 {
    case -28:
        switch 2147483647 {
        case 32768:
            if true {
                v6 = (v6 ^ (v6 | 0))
                v1 += 3.14159
            } else {
                fmt.Println(v3)
                switch "hello" {
                case "foo bar":
                    v3 -= v3
                case "\n\t":
                    var v7 uint32 = 2147483648
                    _ = v7
                default:
                    var v8 uint64 = 1
                    _ = v8
                }
            }
        case 9223372036854775807:
            v1 += 1.5
        case 128:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case 128:
        for fi9 := 0; fi9 < 13; fi9++ {
            v3 -= v3
            _ = fi9
        }
    }
    var v10 int8 = 0
    if false {
        v10 -= v10
        if false {
            var v11 float64 = (v3 * v3)
            var v12 string = "hello"
            _ = v11
            _ = v12
        }
    }
    _ = len(v2)
    v13 := GenS1{F0: 255, F1: -1, F2: -9223372036854775808}
    var v14 int32 = v6
    var v15 bool = false
    switch ("\n\t" + ((("a" + "a") + "") + "\n\t")) {
    case "\n\t":
        v15 = v15
    }
    v16 := GenS2{F0: 77, F1: (v10 << ((4294967295) & 63))}
    switch ("" + "") {
    case "":
        fmt.Println(v3, v6)
    case "a":
        v17 := GenS1{F0: 255, F1: -32768, F2: 1}
        _ = v17
    case "hello":
        if false {
            v15 = ("hello" <= ("foo bar" + ("a" + "hello")))
            var v18 uint16 = 255
            _ = len(v2)
            _ = v18
        }
    }
    var v19 float64 = ((v3 - ((v3 * v3) - 0.0)) * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v6
}
