package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 string }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 2
    var v3 string = ("" + ((("foo bar" + "") + "\n\t") + "\n\t"))
    var v4 uint = 2
    var v5 int32 = 841
    v6 := GenS1{F0: (((v3 + (((("hello" + "hello") + v3) + ((v3 + "foo bar") + v3)) + (("foo bar" + "hello") + v3))) + "hello") + ((v3 + v3) + v3))}
    v7 := GenS1{F0: ""}
    v3 = ("a" + v3)
    switch v3 {
    case "hello":
        switch v3 {
        case "hello":
            var v8 int = 0
            _ = v8
        case "\n\t":
            v4 += v4
        case "":
            v9 := GenS1{F0: "foo bar"}
            _ = v9
        default:
            switch -415 {
            case -9223372036854775808:
                v5 -= v5
            }
        }
    default:
        v3 = "\n\t"
    }
    v4 += (^(v4))
    var v10 int64 = 656
    fmt.Println(v2)
    var v11 int32 = v5
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS2{F0: 32767, F1: v3}
        if (!false) {
            if true {
                fmt.Println(v3, fi12, v3)
                switch "a" {
                case "\n\t":
                    var v14 float64 = 1.5
                    _ = v14
                case "foo bar":
                    v5 = (v5 - v11)
                case "a":
                    for fi15 := 0; fi15 < 7; fi15++ {
                        v2 = (v2 + v2)
                        _ = fi15
                    }
                default:
                    v3 = "hello"
                }
                v4 = (v4 | (v4 | 65535))
            } else {
                for fi16 := 0; fi16 < 11; fi16++ {
                    v4 -= v4
                    _ = fi16
                }
            }
            v11 = int32(v10)
            var v17 float32 = 3.14159
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v4, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
