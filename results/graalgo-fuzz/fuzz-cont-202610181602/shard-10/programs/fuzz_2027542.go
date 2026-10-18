package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint32) int {
    var v6 string = "\n\t"
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 634
}
func genFn1(p0 uint) int64 {
    var v7 int = 365
    v8 := GenS1{F0: (^(p0))}
    var v9 float64 = float64((1.5 - float32(0.1)))
    var v10 int16 = -32768
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = (true || false)
    var v3 int8 = 42
    var v4 int = 32767
    var v5 float32 = 3.14159
    v3 -= (-(v3))
    v4 -= (^(v4))
    var v11 int8 = v3
    v5 -= v5
    var v12 bool = false
    v13 := GenS1{F0: 18446744073709551615}
    var v14 uint8 = 255
    switch (v4 & (^(v4))) {
    case -619:
        fmt.Println(v4)
    default:
        var v15 uint = 42
        _ = v15
    }
    var v16 uint32 = 2
    var v17 int64 = 42
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 int16 = -869
        switch "hello" {
        case "foo bar":
            switch "\n\t" {
            case "hello":
                var v20 int64 = genFn1(0)
                _ = v20
            case "foo bar":
                var v21 uint8 = (v14 + 190)
                _ = v21
            case "\n\t":
                v19 -= v19
            default:
                switch ((((("a" + "foo bar") + ("a" + (("" + "") + "hello"))) + ("foo bar" + "a")) + ("\n\t" + "")) + "a") {
                case "":
                    var v22 uint64 = 0
                    _ = v22
                }
            }
        case "hello":
            switch "\n\t" {
            case "hello":
                var v23 string = ("\n\t" + "")
                _ = v23
            case "foo bar":
                v4 = int(v16)
            }
        case "\n\t":
            v3 = 42
        default:
            var v24 uint32 = uint32(v14)
            _ = v24
        }
        _ = fi18
        _ = v19
    }
    v4 -= int(v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
