package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 bool = false
    v2 := []string{"hello", (("foo bar" + (("" + "hello") + ("foo bar" + ((("\n\t" + "") + "\n\t") + "foo bar")))) + ((("hello" + "hello") + "\n\t") + "\n\t"))}
    var v3 uint64 = 166
    var v4 float64 = (1.5 * (1e10 * 0.1))
    v3 += (v3 / (((v3 ^ v3)) | 1))
    v4 = float64((3.14159 - 0.1))
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := GenS1{F0: 65535, F1: 0}
        var v7 int16 = 1000
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint64 = (v3 ^ v3)
    switch 32768 {
    case -139:
        for fi9 := 0; fi9 < 12; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                if v1 {
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                }
                v3 = v8
                _ = fi10
            }
            v1 = v1
            v3 += (v3 % (((v8 >> ((1) & 63))) | 1))
            _ = fi9
        }
    case -9223372036854775808:
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                v2 = append(v2, ("\n\t" + ("foo bar" + ("hello" + "hello"))))
                _ = fi12
            }
            var v13 int32 = 0
            _ = fi11
            _ = v13
        }
    case 128:
        v4 = v4
    default:
        v3 = 2
    }
    v3 = v8
    v14 := GenS1{F0: 241, F1: 255}
    v2 = append(v2, ("" + "a"))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
