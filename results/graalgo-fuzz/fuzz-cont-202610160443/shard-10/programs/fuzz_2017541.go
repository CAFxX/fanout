package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint64 = 2
    var v2 uint8 = 155
    var v3 string = (("\n\t" + ("foo bar" + "")) + "\n\t")
    var v4 float64 = 0.1
    v5 := make(map[string]int)
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 int16 = 8
            var v9 int8 = 1
            _ = fi7
            _ = v8
            _ = v9
        }
        delete(v5, v3)
        switch fi6 {
        case 128:
            v3 = (((v3 + v3) + v3) + "foo bar")
        }
        _ = fi6
    }
    if ((true || (((v4 / v4) <= (3.14159 * 0.1)) && false)) || (0 >= -1)) {
        var v10 int8 = 59
        _ = v10
    }
    var v11 int = 2147483647
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := []int{(v11 << ((42) & 63)), 720, fi12}
        _ = fi12
        _ = v13
    }
    v14 := []int{(v11 | -223), (^(v11)), v11}
    switch v3 {
    case "a":
        if false {
            var v15 float64 = (v4 / v4)
            _ = v15
        }
    case "foo bar":
        v1 = (^(v1))
    case "\n\t":
        fmt.Println(v11)
    default:
        v16 := GenS1{F0: 1}
        _ = v16
    }
    v1 = 4294967295
    var v17 float64 = float64(1.5)
    v18 := GenS1{F0: 65535}
    var v19 int16 = 393
    fmt.Println(v1, v3, v4)
    if true {
        v2 = v2
        for fi20 := 0; fi20 < 8; fi20++ {
            v11 += -9223372036854775808
            _ = fi20
        }
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 string = ((v3 + v3) + v3)
            v19 = 1000
            _ = fi21
            _ = v22
        }
    }
    v23 := GenS1{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
