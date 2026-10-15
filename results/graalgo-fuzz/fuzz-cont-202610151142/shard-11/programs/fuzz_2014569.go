package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int32; F2 int }

func main() {
    var v1 int64 = 0
    var v2 uint16 = 301
    v3 := make(map[uint32]bool)
    var v4 int64 = v1
    v5 := make(map[string]bool)
    var v6 int = -1
    var v7 int64 = 0
    var v8 float32 = (1e10 + (1.5 - 1.5))
    v9 := GenS1{F0: 2}
    v10 := GenS1{F0: 255}
    for fi11 := 0; fi11 < 18; fi11++ {
        switch fi11 {
        case -2147483648:
            fmt.Println(v4)
        case 128:
            for fi12 := 0; fi12 < 13; fi12++ {
                v6 = fi11
                _ = fi12
            }
        }
        switch ("" + ("foo bar" + "foo bar")) {
        case "\n\t":
            for fi13 := 0; fi13 < 13; fi13++ {
                fmt.Println(v7)
                _ = fi13
            }
        case "":
            v4 -= 366
        case "foo bar":
            switch "" {
            case "hello":
                v5[("\n\t" + "a")] = true
            case "\n\t":
                v7 -= v1
            }
        }
        _ = fi11
    }
    var v14 float32 = v8
    var v15 bool = false
    v7 = -9223372036854775808
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS1{F0: 246}
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v19 := GenS2{F0: (v7 + v7), F1: -2147483648, F2: (v6 % ((v6) | 1))}
    if (3.14159 >= ((float64(float32((3.14159 * 3.14159))) * 0.1) / 0.1)) {
        if (!(!false)) {
            delete(v3, 2147483648)
            v20 := GenS1{F0: 42}
            _ = v20
        } else {
            var v21 string = ("hello" + ((((("foo bar" + ("" + "")) + "a") + ("a" + "")) + ("hello" + "")) + "a"))
            _ = v21
        }
        v22 := []int{v6, v6}
        var v23 uint8 = 128
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: 127}
    var v25 string = "foo bar"
    switch 483 {
    case -2147483648:
        v2 -= 32767
    case 127:
        v26 := []int{(v6 & (v6 + v6)), v6, v6}
        _ = v26
    default:
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 int32 = int32(v2)
            _ = fi27
            _ = v28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
