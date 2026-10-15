package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 bool }
type GenS2 struct { F0 bool; F1 bool; F2 string }

func main() {
    var v1 uint16 = 210
    var v2 uint16 = 571
    var v3 float32 = ((3.14159 + -2.25) * 1.5)
    var v4 uint8 = 2
    var v5 float32 = (0.0 / 1.0)
    if false {
        v6 := GenS1{F0: "hello", F1: 42, F2: (2147483647 > 2147483647)}
        switch "a" {
        case "a":
            for fi7 := 0; fi7 < 13; fi7++ {
                if true {
                    fmt.Println(fi7, fi7, v5)
                    v2 -= (^(v2))
                    v2 = (v2 ^ 32767)
                } else {
                    var v8 int32 = -540
                    v1 = v1
                    var v9 string = "hello"
                    _ = v8
                    _ = v9
                }
                _ = fi7
            }
        }
        for fi10 := 0; fi10 < 2; fi10++ {
            for fi11 := 0; fi11 < 9; fi11++ {
                v1 = v2
                v12 := GenS1{F0: "", F1: 255, F2: (!(v5 <= 0.0))}
                v4 -= 183
                _ = fi11
                _ = v12
            }
            v1 -= (v2 + v2)
            _ = fi10
        }
        _ = v6
    } else {
        var v13 uint32 = 2147483648
        v13 = 2147483648
        _ = v13
    }
    fmt.Println(v3)
    switch "" {
    case "":
        v14 := GenS2{F0: false, F1: false, F2: "hello"}
        _ = v14
    case "a":
        switch "hello" {
        case "hello":
            v2 = (v1 * (v1 * 65535))
        default:
            var v15 int32 = 451
            _ = v15
        }
    }
    v3 -= ((v3 - v5) + (v5 + (v5 / 1.0)))
    v1 = v2
    v2 += v1
    v2 = v2
    if false {
        v3 = v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
