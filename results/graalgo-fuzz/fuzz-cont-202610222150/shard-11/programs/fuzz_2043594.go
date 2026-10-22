package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int16 = -355
    var v2 uint64 = 18446744073709551615
    v3 := []string{(("" + "foo bar") + ("a" + ((("a" + ("hello" + ("" + ""))) + ("" + (((("hello" + "a") + ("hello" + "foo bar")) + "\n\t") + "\n\t"))) + ""))), ""}
    var v4 int64 = -337
    var v5 int8 = -74
    fmt.Println(v1)
    if (255 == 42) {
        var v6 uint8 = 53
        _ = v6
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 int16 = v1
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 2, F1: "a", F2: 1}
        for fi10 := 0; fi10 < 2; fi10++ {
            v3 = append(v3, ("\n\t" + ("foo bar" + "\n\t")))
            for fi11 := 0; fi11 < 16; fi11++ {
                v5 = 100
                v12 := GenS1{F0: 1, F1: ((("hello" + ("a" + "a")) + "\n\t") + "hello"), F2: 1}
                v4 = v4
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
        _ = v9
    }
    switch "foo bar" {
    case "foo bar":
        if true {
            if (2 <= v2) {
                v1 = v1
                v13 := GenS1{F0: 42, F1: "\n\t", F2: 139}
                _ = v13
            }
            switch 128 {
            case 410:
                v14 := GenS1{F0: 255, F1: ("foo bar" + ("a" + "a")), F2: 32}
                _ = v14
            case -796:
                fmt.Println(v4)
            case -9223372036854775808:
                v15 := GenS1{F0: 120, F1: "hello", F2: 0}
                _ = v15
            }
        }
    case "a":
        v16 := GenS1{F0: uint16(v2), F1: "", F2: 2}
        _ = v16
    case "\n\t":
        var v17 int = 127
        _ = v17
    default:
        v18 := GenS1{F0: uint16(v2), F1: "a", F2: 91}
        _ = v18
    }
    if false {
        v2 = (^(v2))
    }
    switch 1 {
    case 128:
        if (!(1.5 >= (1.5 + (-2.25 * 1e10)))) {
            switch 32768 {
            case 42:
                fmt.Println(v4)
            default:
                v19 := GenS2{F0: 2147483647}
                _ = v19
            }
            if len(v3) > 0 {
                v3[0] = "a"
            }
        }
    case 2147483647:
        var v20 bool = true
        _ = v20
    case 32767:
        var v21 string = "foo bar"
        _ = v21
    default:
        v1 += (v1 + -32768)
    }
    var v22 float64 = 1.5
    v22 -= (((v22 * 3.14159) * float64((3.14159 * -2.25))) * (v22 + ((v22 - v22) + float64(1e10))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
