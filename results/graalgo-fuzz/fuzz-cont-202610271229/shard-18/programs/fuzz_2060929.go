package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int64 }

func main() {
    v1 := []string{"\n\t", "a", "hello", "\n\t"}
    v2 := []string{""}
    v3 := []uint32{0, 42, 2147483647}
    var v4 string = (("a" + "hello") + "")
    v5 := GenS1{F0: 18446744073709551615}
    v4 = (("a" + "a") + v4)
    var v6 uint32 = 2
    if (!true) {
        for fi7 := 0; fi7 < 19; fi7++ {
            if len(v1) > 0 {
                v1[0] = ""
            }
            var v8 int32 = 976
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS1{F0: 0}
    var v10 int64 = 9223372036854775807
    v10 = v10
    _ = len(v2)
    switch 118 {
    case 0:
        v6 -= 1
    default:
        var v11 uint32 = v6
        _ = v11
    }
    v12 := GenS1{F0: 1}
    var v13 float32 = 0.0
    v4 = ((v4 + (((v4 + (v4 + (v4 + "foo bar"))) + v4) + v4)) + ("\n\t" + v4))
    if ((42 >= 940) && false) {
        if true {
            if false {
                v14 := GenS2{F0: 360, F1: (-(v10))}
                _ = v14
            } else {
                for fi15 := 0; fi15 < 19; fi15++ {
                    v6 -= 2
                    v4 = v4
                    _ = fi15
                }
            }
        } else {
            fmt.Println(v4, v10, v10)
        }
        var v16 int16 = 0
        _ = v16
    } else {
        v4 = v4
    }
    v4 = ((v4 + v4) + (v4 + v4))
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
