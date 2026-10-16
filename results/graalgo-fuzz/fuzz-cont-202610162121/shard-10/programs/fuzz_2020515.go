package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = -128
    var v3 int64 = 0
    var v4 int64 = (-(v3))
    var v5 float32 = 0.1
    var v6 string = "\n\t"
    v7 := GenS1{F0: v3, F1: 285}
    v6 = v6
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            if (v5 != float32(1e10)) {
                switch (v6 + v6) {
                case "a":
                    v10 := GenS1{F0: (-(v4)), F1: -32768}
                    _ = v10
                case "":
                    v6 = v6
                }
                var v11 uint16 = 0
                _ = v11
            }
            _ = fi9
        }
        if false {
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 string = ("" + v6)
                v3 += -114
                _ = fi12
                _ = v13
            }
        } else {
            var v14 int8 = v2
            fmt.Println(v4, v5, fi8)
            var v15 int64 = v4
            _ = v14
            _ = v15
        }
        _ = fi8
    }
    v2 = (v2 ^ (v2 ^ (v2 << ((356) & 63))))
    v6 = (("a" + v6) + v6)
    var v16 uint16 = 255
    if ((v2 >> ((373) & 63)) == 0) {
        v3 = v4
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS1{F0: (v4 << ((255) & 63)), F1: int16(v4)}
        _ = fi17
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
