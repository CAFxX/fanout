package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 int16 = 1000
    var v2 int16 = v1
    var v3 int32 = -1
    var v4 int = 9223372036854775807
    var v5 float64 = float64(-2.25)
    v6 := make(map[string]bool)
    v7 := []int{v4, (v4 ^ v4)}
    var v8 float64 = 3.14159
    switch v4 {
    case 4294967295:
        v9 := GenS1{F0: 2, F1: 42}
        _ = v9
    case 0:
        if true {
            v10 := GenS1{F0: 1, F1: 91}
            fmt.Println(v4, v2, v4)
            var v11 int = v4
            _ = v10
            _ = v11
        }
    case -658:
        for fi12 := 0; fi12 < 20; fi12++ {
            switch (fi12 ^ 9223372036854775807) {
            case 42:
                var v13 string = "foo bar"
                _ = v13
            case 487:
                for fi14 := 0; fi14 < 9; fi14++ {
                    if true {
                        fmt.Println(v4, fi14)
                    }
                    v2 -= v1
                    switch v4 {
                    case 227:
                        v1 += v2
                    case -391:
                        var v15 string = ""
                        _ = v15
                    default:
                        v2 -= -658
                    }
                    _ = fi14
                }
            }
            var v16 float64 = 3.14159
            _ = fi12
            _ = v16
        }
    }
    fmt.Println(v2, v3, v3)
    switch v4 {
    case -244:
        v1 = v2
    case -2147483648:
        var v17 int64 = -1
        _ = v17
    case 1:
        v3 -= (v3 & (v3 / ((v3) | 1)))
    }
    fmt.Println(v2, v8)
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := GenS1{F0: 255, F1: 42}
            var v21 int16 = (v2 / ((538) | 1))
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi19
            _ = v20
            _ = v21
        }
        v22 := GenS1{F0: 192, F1: 2}
        _ = fi18
        _ = v22
    }
    v4 = v4
    v4 += 4294967295
    v1 -= 44
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
