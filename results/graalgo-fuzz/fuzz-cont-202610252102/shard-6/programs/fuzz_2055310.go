package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int16; F1 int64 }

func main() {
    var v1 uint8 = 127
    var v2 float64 = (0.0 + (((-2.25 - 0.0) * 0.1) * -2.25))
    v3 := []int{4294967295}
    var v4 uint64 = 258
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v5 int8 = 54
    v4 -= v4
    if (false || ("a" != ("foo bar" + (("foo bar" + "\n\t") + ("hello" + "a"))))) {
        var v6 int16 = 42
        v7 := GenS1{F0: 951, F1: -1}
        var v8 uint16 = 255
        _ = v6
        _ = v7
        _ = v8
    }
    if len(v3) > 0 {
        v3[0] = 485
    }
    if (false && true) {
        v9 := GenS1{F0: 4294967295, F1: -1}
        _ = v9
    }
    if true {
        var v10 float32 = -2.25
        if len(v3) > 0 {
            v3[0] = int(v5)
        }
        if false {
            v11 := GenS1{F0: 4294967295, F1: -9223372036854775808}
            var v12 uint16 = 32767
            v12 = (v12 + 42)
            _ = v11
            _ = v12
        }
        _ = v10
    } else {
        v13 := GenS2{F0: 4294967295, F1: ((((-2147483648 > 850) || (4294967295 != 2147483647)) || true) || (-921 <= 2147483647))}
        fmt.Println(v5, v4, v2)
        _ = v13
    }
    v1 += 1
    if true {
        v5 += v5
        var v14 int64 = 9223372036854775807
        _ = v14
    }
    if ((((0.1 <= (-2.25 + (0.0 + 1e10))) && (1 == 167)) || (4294967295 == 127)) || ((v5 >= -109) || true)) {
        fmt.Println(v2, v2, v5)
        switch "foo bar" {
        case "a":
            _ = len(v3)
        }
    }
    var v15 int16 = 42
    if true {
        v16 := GenS3{F0: (v15 / (((v15 % ((v15) | 1))) | 1)), F1: -1}
        fmt.Println(v5, v1, v4)
        _ = v16
    } else {
        v15 = v15
        fmt.Println(v1, v15)
    }
    var v17 float64 = 3.14159
    v5 += -51
    for fi18 := 0; fi18 < 16; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 bool = true
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    var v21 float64 = (((((3.14159 / v2) + ((0.0 / 1.5) / (v2 / 1e10))) + (float64(3.14159) - float64((-2.25 + 3.14159)))) + (v2 * (v2 * (0.1 / 1.5)))) + v2)
    if len(v3) > 0 {
        v3[0] = -913
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
