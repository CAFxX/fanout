package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }

func main() {
    var v1 float32 = -2.25
    var v2 int8 = 100
    v3 := []uint32{352, 2}
    var v4 uint16 = 32767
    v2 -= (v2 & (v2 / ((v2) | 1)))
    v1 = (-2.25 + float32(-2.25))
    if true {
        v2 = (v2 % ((v2) | 1))
    }
    var v5 uint64 = 0
    switch "foo bar" {
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "":
        if (0 == -1) {
            for fi6 := 0; fi6 < 3; fi6++ {
                switch ("\n\t" + "") {
                case "\n\t":
                    if true {
                        fmt.Println(v2)
                        switch ("" + ("foo bar" + "foo bar")) {
                        case "hello":
                            fmt.Println(v5, v5, v4)
                        case "a":
                            _ = len(v3)
                        default:
                            var v7 int16 = -920
                            _ = v7
                        }
                    } else {
                        if (false && true) {
                            v5 -= (v5 & (^(v5)))
                        }
                    }
                case "":
                    if ((511 < -1) && true) {
                        v4 = 255
                    }
                }
                v8 := GenS1{F0: 18446744073709551615, F1: 1}
                _ = fi6
                _ = v8
            }
        }
    case "foo bar":
        fmt.Println(v2, v5, v2)
    default:
        fmt.Println(v4)
    }
    var v9 uint8 = 45
    v1 = v1
    v4 = v4
    v10 := GenS1{F0: 65535, F1: 0}
    v2 = (v2 - v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
