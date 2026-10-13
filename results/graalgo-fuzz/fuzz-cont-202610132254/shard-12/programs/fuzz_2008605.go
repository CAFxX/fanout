package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
func genFn0() uint16 {
    var v5 string = "hello"
    if false {
        switch 32768 {
        case -1:
            fmt.Println(v5)
        case 826:
            v6 := []int{9223372036854775807, 620, 128}
            _ = v6
        case 0:
            v5 = ("\n\t" + "foo bar")
        }
        switch "" {
        case "hello":
            switch ("\n\t" + (v5 + (v5 + "foo bar"))) {
            case "\n\t":
                v7 := []int{9223372036854775807, 4294967295, 949}
                _ = v7
            }
        case "foo bar":
            var v8 bool = true
            _ = v8
        case "a":
            v9 := []int{1, -1}
            _ = v9
        default:
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 int = (-(fi10))
                for fi12 := 0; fi12 < 17; fi12++ {
                    fmt.Println(fi10, fi12)
                    _ = fi12
                }
                _ = fi10
                _ = v11
            }
        }
    }
    switch (v5 + (v5 + v5)) {
    case "hello":
        fmt.Println(v5)
    case "foo bar":
        v5 = v5
    }
    var v13 int = 128
    var v14 int8 = 42
    _ = v13
    _ = v14
    _ = v5
    return 42
}

func main() {
    v1 := []int{9223372036854775807}
    var v2 float64 = 1.5
    var v3 int = -235
    var v4 float64 = ((((v2 * (0.1 + 0.1)) * ((v2 * 0.0) + 1.5)) - (v2 / 1.5)) + float64(0.0))
    v4 = v2
    if len(v1) > 0 {
        v1[0] = 0
    }
    fmt.Println(v2, v3)
    switch "a" {
    case "\n\t":
        v3 += v3
    case "a":
        var v15 uint32 = 65535
        _ = v15
    default:
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 int64 = 42
            v3 += fi16
            _ = fi16
            _ = v17
        }
    }
    v4 += ((1.5 * ((0.1 + -2.25) + v2)) * v4)
    if true {
        if (1 == 476) {
            v18 := GenS1{F0: -32768, F1: (2 != genFn0())}
            _ = v18
        }
        var v19 uint32 = 2147483648
        var v20 string = ("hello" + "hello")
        _ = v19
        _ = v20
    }
    v3 += -739
    var v21 int32 = -2147483648
    v22 := GenS1{F0: 42, F1: (false && (2147483647 <= 65535))}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    v2 -= (v4 - (v4 + (float64((((0.0 - -2.25) - (-2.25 + (0.0 + (((0.0 + (0.1 / 3.14159)) - ((1e10 + 1e10) * ((3.14159 - ((0.1 + 3.14159) / (((0.0 + (0.0 / (0.1 - (((0.1 - (1e10 * 0.0)) * 1e10) - float32(v2))))) + float32(1.5)) + 1.5))) * 0.0))) + ((3.14159 + ((3.14159 / 1.0) - 1.5)) / 0.1))))) / ((1e10 * 0.1) - -2.25))) * 0.1)))
    v3 = (v3 ^ 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
