package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int }
func genFn0(p0 uint8, p1 uint64, p2 uint) uint16 {
    var v4 float32 = ((-2.25 / 0.1) * 0.0)
    p1 += uint64(p2)
    var v5 int16 = 1
    v4 = float32(-2.25)
    fmt.Println(p0, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 float64 = (0.1 * 1e10)
    var v2 float64 = 3.14159
    var v3 string = (("hello" + ("a" + "hello")) + "hello")
    switch -2147483648 {
    case 128:
        var v6 int8 = 100
        _ = v6
    case 2147483647:
        var v7 int = -9223372036854775808
        _ = v7
    }
    v2 += v2
    v8 := GenS1{F0: -32768, F1: 127}
    var v9 float32 = (0.0 * ((1.5 * 3.14159) / (((3.14159 + 0.0) * (1e10 / float32((v2 / v2)))) / 1.5)))
    if false {
        switch 1 {
        case -1:
            v10 := GenS2{F0: 128}
            _ = v10
        }
        v1 = (float64(3.14159) * (3.14159 + (0.1 * ((v1 / 0.1) + ((v2 * v1) - (v2 / v1))))))
        var v11 int = 9223372036854775807
        _ = v11
    } else {
        if false {
            var v12 int = 32768
            v9 = -2.25
            var v13 float32 = (v9 + v9)
            _ = v12
            _ = v13
        } else {
            fmt.Println(v2)
            var v14 uint16 = 316
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint8 = 127
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: -601}
    var v18 int8 = 78
    switch (v3 + "\n\t") {
    case "a":
        var v19 uint16 = 255
        _ = v19
    case "\n\t":
        v20 := []int{248, 0, 42}
        _ = v20
    default:
        v9 -= v9
    }
    var v21 bool = (3.14159 == (((0.0 / v2) - float64((v9 - 1.5))) + ((1e10 * 0.0) * float64(v9))))
    fmt.Println(v18, v9, v3)
    var v22 int = 427
    if v21 {
        var v23 string = v3
        _ = v23
    } else {
        for fi24 := 0; fi24 < 11; fi24++ {
            fmt.Println(v3, v3)
            if (0 >= -1) {
                for fi25 := 0; fi25 < 5; fi25++ {
                    var v26 uint8 = 127
                    v18 = -1
                    v27 := []int{fi24, fi24, (fi25 - -2)}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            } else {
                var v28 string = (("" + "hello") + v3)
                _ = v28
            }
            _ = fi24
        }
        var v29 string = ("" + v3)
        v18 += (-(v18))
        _ = v29
    }
    fmt.Println(v9, v22, v3)
    v9 += float32((3.14159 / (((float64(v9) / (0.1 - v2)) * 0.1) - v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v8
    _ = v9
}
