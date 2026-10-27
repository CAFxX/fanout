package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int64 = 42
    var v2 int32 = -1
    var v3 float64 = -2.25
    var v4 int64 = v1
    var v5 int = 0
    var v6 uint8 = 228
    v2 -= (v2 * 1)
    var v7 float64 = ((1.5 - float64(1.5)) / 1.0)
    v8 := GenS1{F0: "a"}
    switch (v5 & (v5 / (((v5 ^ 924)) | 1))) {
    case 42:
        for fi9 := 0; fi9 < 4; fi9++ {
            v6 = v6
            _ = fi9
        }
    case 4294967295:
        for fi10 := 0; fi10 < 19; fi10++ {
            fmt.Println(v6, v6)
            _ = fi10
        }
    default:
        switch 9223372036854775807 {
        case 682:
            v11 := []int{}
            _ = v11
        case -1:
            fmt.Println(v4, v7, v2)
        case 2147483647:
            v12 := []int{(-(v5))}
            _ = v12
        default:
            fmt.Println(v1, v1)
        }
    }
    v13 := []int{v5}
    if true {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 int64 = 42
            var v16 int = (v5 * (v5 & fi14))
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    var v17 float64 = (0.0 * (v7 + (((0.0 + v3) - v7) + ((-2.25 * (((0.1 * (v3 * (v3 - float64((0.1 - 0.1))))) / (v3 - (float64(((0.0 * ((3.14159 + ((0.1 - (0.1 - -2.25)) * 0.0)) * float32(v7))) / 1e10)) + (v7 - ((v3 * (3.14159 / float64(3.14159))) + v7))))) + (float64(-2.25) + (0.0 / (-2.25 * float64(3.14159)))))) + (v7 - float64(0.0))))))
    if false {
        v13 = append(v13, -9223372036854775808)
        for fi18 := 0; fi18 < 10; fi18++ {
            v6 -= (^(v6))
            _ = fi18
        }
    } else {
        v19 := GenS1{F0: ((("" + ("a" + "foo bar")) + (((("" + (("a" + "hello") + ("" + "a"))) + ("\n\t" + "foo bar")) + "foo bar") + ((("a" + ("a" + "foo bar")) + "hello") + "a"))) + "foo bar")}
        v7 += float64(((0.0 - (0.0 - (1.5 + 3.14159))) * (-2.25 * -2.25)))
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
