package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 string; F1 int }

func main() {
    v1 := []uint32{}
    var v2 int = 127
    var v3 float64 = ((0.0 + -2.25) - 1.5)
    var v4 uint8 = 127
    if false {
        if true {
            v5 := GenS1{F0: 42, F1: 890}
            v6 := GenS2{F0: "foo bar", F1: (v2 >> ((334) & 63))}
            for fi7 := 0; fi7 < 4; fi7++ {
                v3 += -2.25
                fmt.Println(fi7, fi7)
                var v8 uint = 2
                _ = fi7
                _ = v8
            }
            _ = v5
            _ = v6
        } else {
            var v9 int16 = -32768
            var v10 int = 32767
            if len(v1) > 0 {
                v1[0] = 2
            }
            _ = v10
            _ = v9
        }
    }
    v1 = append(v1, 0)
    switch (v2 - (v2 % ((42) | 1))) {
    case 32768:
        fmt.Println(v2, v3, v4)
    default:
        v1 = append(v1, 2147483647)
    }
    v11 := GenS2{F0: ("a" + (((((("\n\t" + "hello") + ("hello" + ("foo bar" + ("foo bar" + "hello")))) + (("hello" + ("a" + "hello")) + ("\n\t" + ("a" + ("" + "\n\t"))))) + "foo bar") + ("" + "foo bar")) + ("foo bar" + "\n\t"))), F1: v2}
    v12 := GenS1{F0: -32768, F1: 859}
    for fi13 := 0; fi13 < 7; fi13++ {
        v2 = (fi13 % (((v2 * v2)) | 1))
        var v14 uint16 = 614
        _ = fi13
        _ = v14
    }
    v2 = (v2 << ((521) & 63))
    fmt.Println(v3, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
