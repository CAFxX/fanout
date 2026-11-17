package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int32 }
func genFn0(p0 uint32, p1 int8) uint {
    var v4 int32 = -540
    if (255 == 265) {
        v5 := []int{}
        fmt.Println(p0)
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return 913
}
func genFn1() uint16 {
    var v6 float64 = 1.5
    v6 -= 0.0
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: ("foo bar" + ((("foo bar" + "hello") + "a") + ("foo bar" + (("hello" + "") + "hello"))))}
        switch "" {
        case "a":
            fmt.Println(v6, v6)
        case "hello":
            v9 := []int{}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    _ = v6
    return 997
}

func main() {
    var v1 bool = (2 != 42)
    v2 := make(map[int]uint32)
    var v3 uint16 = 65535
    fmt.Println(v3, v3)
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(v3)
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := GenS1{F0: (("\n\t" + (((("\n\t" + ((((("" + "foo bar") + "foo bar") + "") + (("foo bar" + (("" + "\n\t") + ("\n\t" + ""))) + "\n\t")) + ("a" + ""))) + "foo bar") + "foo bar") + "a")) + "a")}
            v2[fi10] = 633
            v3 += (v3 >> ((42) & 63))
            _ = fi11
            _ = v12
        }
        var v13 int64 = 0
        _ = fi10
        _ = v13
    }
    var v14 uint32 = 2147483648
    var v15 uint16 = (v3 % (((v3 & (v3 % ((v3) | 1)))) | 1))
    v1 = v1
    switch -23 {
    case 32768:
        var v16 float64 = 0.0
        _ = v16
    case -1:
        v3 += v15
    case -2147483648:
        v15 -= v15
    default:
        for fi17 := 0; fi17 < 15; fi17++ {
            if v1 {
                var v18 uint16 = 107
                var v19 bool = v1
                fmt.Println(v3, fi17)
                _ = v18
                _ = v19
            } else {
                v14 -= (v14 ^ v14)
                v14 += (v14 - 240)
                v20 := GenS2{F0: -2147483648, F1: -2147483648}
                _ = v20
            }
            v15 -= v3
            _ = fi17
        }
    }
    var v21 uint64 = 1
    var v22 float32 = (1e10 + 1e10)
    var v23 uint16 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
