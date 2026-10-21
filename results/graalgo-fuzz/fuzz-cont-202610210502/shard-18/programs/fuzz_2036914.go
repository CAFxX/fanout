package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 string }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 string = ("a" + "\n\t")
    var v2 uint64 = 42
    v3 := []string{v1, v1, "hello", (v1 + ("hello" + v1))}
    var v4 float64 = float64((-2.25 / -2.25))
    var v5 uint16 = 322
    var v6 uint64 = 18446744073709551615
    var v7 bool = true
    if (1e10 <= 1e10) {
        if (v7 && v7) {
            var v8 string = ((v1 + "") + v1)
            if len(v3) > 0 {
                v3[0] = v1
            }
            v7 = (42 <= 1000)
            _ = v8
        }
    } else {
        var v9 uint32 = 0
        _ = v9
    }
    v5 -= (v5 / (((^(v5))) | 1))
    v10 := GenS1{F0: (v6 % ((v6) | 1)), F1: 1, F2: v1}
    v11 := GenS1{F0: (^(v2)), F1: 289, F2: (v1 + ("\n\t" + (((v1 + v1) + "hello") + "foo bar")))}
    switch v1 {
    case "hello":
        if (!v7) {
            for fi12 := 0; fi12 < 19; fi12++ {
                for fi13 := 0; fi13 < 7; fi13++ {
                    v3 = append(v3, "a")
                    _ = fi13
                }
                _ = fi12
            }
            fmt.Println(v1, v5)
        } else {
            v2 = v2
        }
    case "a":
        var v14 uint16 = v5
        _ = v14
    default:
        var v15 uint16 = (v5 / (((^(v5))) | 1))
        _ = v15
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 int16 = 42
        v3 = append(v3, "a")
        v2 = v6
        _ = fi16
        _ = v17
    }
    switch 32767 {
    case -2147483648:
        _ = len(v3)
    case 9223372036854775807:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 32768:
        switch 174 {
        case 0:
            var v18 int64 = 0
            _ = v18
        case -9223372036854775808:
            var v19 uint64 = v2
            _ = v19
        case 2147483647:
            for fi20 := 0; fi20 < 7; fi20++ {
                switch v1 {
                case "foo bar":
                    v21 := GenS2{F0: 2}
                    _ = v21
                }
                if v7 {
                    v1 = ((v1 + "foo bar") + (((("hello" + (v1 + (v1 + v1))) + ("" + ((("foo bar" + v1) + ((v1 + ((v1 + "a") + v1)) + v1)) + v1))) + "") + v1))
                    var v22 uint16 = 42
                    _ = v22
                }
                _ = fi20
            }
        }
    }
    switch ("\n\t" + (v1 + "")) {
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = (("a" + "") + v1)
        }
    case "hello":
        v6 += v2
    }
    var v23 uint8 = 232
    var v24 string = ""
    fmt.Println(v7, v4, v24)
    var v25 float32 = (0.1 * ((0.0 + -2.25) + 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
