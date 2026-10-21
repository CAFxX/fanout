package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint = 1
    var v2 uint8 = 0
    var v3 uint16 = 42
    v4 := []string{("foo bar" + ("hello" + "hello")), ("hello" + "foo bar"), ("foo bar" + "foo bar"), ("" + ("\n\t" + ("foo bar" + "hello")))}
    v5 := GenS1{F0: 42, F1: 918}
    v1 += (v1 * (v1 ^ (v1 - 255)))
    fmt.Println(v1)
    v4 = append(v4, (("\n\t" + "a") + ("a" + "\n\t")))
    var v6 uint = (v1 ^ v1)
    if true {
        v7 := GenS2{F0: 1}
        for fi8 := 0; fi8 < 16; fi8++ {
            _ = len(v4)
            _ = fi8
        }
        var v9 float64 = -2.25
        _ = v7
        _ = v9
    }
    v6 -= 854
    v6 = 18446744073709551615
    if (v3 <= (^(v3))) {
        if (-2147483648 < 711) {
            if len(v4) > 0 {
                v4[0] = (("" + ("hello" + "a")) + "a")
            }
            v3 = v3
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS1{F0: 0, F1: 2}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v1 += (^(v6))
                _ = fi10
                _ = v11
            }
        } else {
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 uint8 = (^(v2))
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if ((!true) && (((0.0 * 3.14159) + 0.1) >= (((0.1 * (-2.25 + 1e10)) - 1.5) + 0.0))) {
                    v13 = v13
                    var v14 uint = (v1 >> ((v1) & 63))
                    _ = v14
                }
                _ = fi12
                _ = v13
            }
        }
    }
    fmt.Println(v6, v3, v2)
    fmt.Println(v2, v6)
    var v15 uint8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
