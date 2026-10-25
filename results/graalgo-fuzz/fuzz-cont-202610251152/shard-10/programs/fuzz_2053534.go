package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint8 }
type GenS2 struct { F0 string; F1 uint8; F2 uint8 }

func main() {
    var v1 int64 = -709
    var v2 uint32 = 0
    v3 := []string{}
    var v4 uint8 = 255
    switch -567 {
    case -578:
        for fi5 := 0; fi5 < 3; fi5++ {
            if false {
                switch fi5 {
                case 32767:
                    var v6 int = (fi5 ^ fi5)
                    _ = v6
                default:
                    if true {
                        var v7 uint16 = 2
                        if (true || true) {
                            var v8 int = -1
                            v2 += v2
                            _ = v8
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v7
                    }
                }
                v3 = append(v3, "")
                if len(v3) > 0 {
                    v3[0] = ("\n\t" + (((("a" + "hello") + "\n\t") + "a") + "hello"))
                }
            } else {
                if (true && false) {
                    v2 = 2147483647
                } else {
                    fmt.Println(v4, fi5)
                    v9 := GenS1{F0: fi5, F1: ("a" + (((("\n\t" + "foo bar") + "") + "hello") + "\n\t")), F2: (v4 | uint8(v2))}
                    _ = v9
                }
                var v10 float64 = float64(1.5)
                _ = v10
            }
            var v11 uint16 = 602
            _ = fi5
            _ = v11
        }
    case -1:
        v12 := GenS1{F0: -852, F1: "hello", F2: 220}
        _ = v12
    }
    v4 = 42
    switch 0 {
    case 0:
        v1 = v1
    case 36:
        v13 := GenS1{F0: -161, F1: "", F2: v4}
        _ = v13
    }
    v14 := GenS2{F0: "foo bar", F1: (v4 << ((65535) & 63)), F2: v4}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v15 uint16 = 337
    v15 += (v15 / ((v15) | 1))
    v1 += -9223372036854775808
    v1 += 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
