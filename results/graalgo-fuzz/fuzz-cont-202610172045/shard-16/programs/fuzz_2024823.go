package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int }

func main() {
    var v1 bool = false
    var v2 uint8 = 118
    var v3 uint32 = 672
    var v4 uint8 = v2
    var v5 float64 = 0.1
    if v1 {
        if v1 {
            fmt.Println(v5, v3, v3)
            v1 = (v1 && (true || true))
            v6 := []int{int(v2), 42}
            _ = v6
        } else {
            v3 = v3
        }
    }
    v4 = 255
    v7 := []int{2147483647, -256, 127}
    v8 := GenS1{F0: 4294967295, F1: 1}
    var v9 int16 = int16(v2)
    for fi10 := 0; fi10 < 8; fi10++ {
        if len(v7) > 0 {
            v7[0] = (fi10 * (fi10 ^ (fi10 >> ((4294967295) & 63))))
        }
        _ = fi10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 string = (((((("a" + ("foo bar" + "")) + ((("a" + "a") + ("a" + "foo bar")) + ((("" + "foo bar") + "\n\t") + "\n\t"))) + "\n\t") + "foo bar") + "a") + ("foo bar" + ("hello" + ("" + "\n\t"))))
            var v14 int8 = -72
            v7 = append(v7, fi12)
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    fmt.Println(v5, v1, v4)
    v1 = true
    v9 = (v9 % (((v9 % ((v9) | 1))) | 1))
    if (false && ("foo bar" == ("\n\t" + "foo bar"))) {
        fmt.Println(v2)
    }
    fmt.Println(v3)
    v1 = v1
    v15 := GenS2{F0: 4294967295, F1: 4294967295}
    v5 -= (v5 * v5)
    var v16 int16 = (v9 / ((v9) | 1))
    for fi17 := 0; fi17 < 11; fi17++ {
        if v1 {
            switch fi17 {
            case -2147483648:
                _ = len(v7)
            case -606:
                v1 = true
            case -9223372036854775808:
                _ = len(v7)
            default:
                for fi18 := 0; fi18 < 6; fi18++ {
                    v3 -= v3
                    var v19 uint8 = 128
                    _ = fi18
                    _ = v19
                }
            }
            v9 -= (^(v16))
        }
        switch "a" {
        case "\n\t":
            v1 = (1 == v2)
        case "a":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case "":
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := GenS1{F0: 32768, F1: 2}
                _ = fi20
                _ = v21
            }
        }
        v7 = append(v7, fi17)
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
