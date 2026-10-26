package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 string; F2 int64 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 uint8 = 255
    var v2 int32 = 0
    v3 := make(map[uint32]uint32)
    var v4 int = int(v2)
    var v5 float32 = 0.0
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: v2}
        var v9 int = (fi7 + -371)
        fmt.Println(v1)
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        if (0 > (^(v1))) {
            for fi11 := 0; fi11 < 2; fi11++ {
                var v12 float64 = 1e10
                v13 := []int{}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        v1 = v1
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 uint8 = v1
            _ = fi14
            _ = v15
        }
        _ = fi10
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        fmt.Println(v5)
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int8 = 100
            v4 -= fi16
            v18 = v18
            v18 = v18
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    switch "" {
    case "\n\t":
        if (!true) {
            fmt.Println(v4, v5, v2)
        }
    case "foo bar":
        for fi19 := 0; fi19 < 16; fi19++ {
            v4 += (^(v4))
            var v20 float64 = (-2.25 * (((-2.25 * -2.25) * 0.0) / 3.14159))
            _ = fi19
            _ = v20
        }
    default:
        v5 -= v5
    }
    v21 := GenS2{F0: 2147483648}
    v4 = (v4 % (((v4 >> ((255) & 63))) | 1))
    var v22 uint = 255
    var v23 uint = 2
    var v24 uint = (v23 | v22)
    v25 := GenS3{F0: ((1.5 - 1.5) < float64(v5)), F1: (("foo bar" + (("" + "\n\t") + ("\n\t" + ("\n\t" + ("" + ("" + ((("hello" + (("a" + ("" + "")) + ("foo bar" + "a"))) + "hello") + "hello"))))))) + (("foo bar" + (("\n\t" + "\n\t") + "")) + "a")), F2: -1}
    v5 -= (v5 + v5)
    v26 := GenS2{F0: 2147483648}
    v27 := GenS4{F0: 32767}
    for fi28 := 0; fi28 < 20; fi28++ {
        v29 := GenS4{F0: 426}
        v22 += (v24 | 42)
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
