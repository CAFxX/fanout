package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() int32 {
    var v5 string = "hello"
    for fi6 := 0; fi6 < 7; fi6++ {
        v5 = v5
        if false {
            var v7 int = (fi6 + fi6)
            _ = v7
        } else {
            fmt.Println(fi6)
            for fi8 := 0; fi8 < 15; fi8++ {
                v5 = "hello"
                v5 = (v5 + v5)
                var v9 uint32 = 0
                _ = fi8
                _ = v9
            }
            switch -855 {
            case -89:
                fmt.Println(v5, v5)
            }
        }
        _ = fi6
    }
    var v10 uint = 1
    _ = v10
    _ = v5
    return -2147483648
}

func main() {
    var v1 float32 = float32(((1e10 + float64(1e10)) / (float64(0.1) - 1.5)))
    var v2 int16 = -1
    v3 := []string{}
    var v4 float32 = (3.14159 - -2.25)
    if (-1 <= -9223372036854775808) {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        if true {
            fmt.Println(v2)
        }
        fmt.Println(v1)
    } else {
        _ = len(v3)
        v11 := GenS1{F0: 42}
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        switch 529 {
        case 636:
            v2 = (^(v2))
        default:
            v13 := GenS1{F0: int32(v2)}
            _ = v13
        }
        var v14 int8 = 100
        _ = fi12
        _ = v14
    }
    fmt.Println(v4)
    var v15 uint16 = 255
    switch -240 {
    case 128:
        if len(v3) > 0 {
            v3[0] = (("\n\t" + "hello") + ("hello" + ""))
        }
    }
    var v16 uint = 906
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = (v2 & -574)
    var v17 float32 = 3.14159
    _ = len(v3)
    for fi18 := 0; fi18 < 17; fi18++ {
        v2 = (^(v2))
        _ = fi18
    }
    switch "foo bar" {
    case "hello":
        _ = len(v3)
    case "a":
        var v19 uint32 = 373
        _ = v19
    case "foo bar":
        var v20 uint8 = 255
        _ = v20
    default:
        _ = len(v3)
    }
    var v21 float64 = 0.1
    for fi22 := 0; fi22 < 20; fi22++ {
        fmt.Println(v21, v17)
        _ = fi22
    }
    v4 = float32(float64(((v4 * (((v17 + v4) * ((v4 * (3.14159 + 3.14159)) * (((v1 + (v4 + ((v4 / (v17 - float32((v21 + 1e10)))) / v4))) * ((v17 + v17) - -2.25)) * float32((((v21 - (0.0 / ((3.14159 + v21) + -2.25))) + (((v21 * ((-2.25 - ((v21 / v21) / v21)) / (v21 * v21))) + 1e10) + (v21 / v21))) + v21))))) - float32(v21))) + float32(v21))))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
