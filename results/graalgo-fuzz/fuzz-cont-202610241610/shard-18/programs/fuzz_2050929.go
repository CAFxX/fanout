package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint }

func main() {
    var v1 string = (("foo bar" + ("\n\t" + "hello")) + "\n\t")
    var v2 uint16 = 2
    v3 := []string{"", v1, "\n\t", v1}
    for fi4 := 0; fi4 < 5; fi4++ {
        for fi5 := 0; fi5 < 3; fi5++ {
            if len(v3) > 0 {
                v3[0] = (v1 + "")
            }
            v2 = (v2 & (v2 | (v2 << ((18446744073709551615) & 63))))
            if (!true) {
                if false {
                    v1 = v1
                }
                fmt.Println(fi4)
                var v6 uint8 = 2
                if true {
                    v7 := GenS1{F0: 0, F1: 120, F2: 18}
                    v8 := GenS1{F0: v6, F1: v6, F2: 65535}
                    _ = v7
                    _ = v8
                }
                if (0 > 0) {
                    var v9 uint8 = (^(v6))
                    _ = v9
                } else {
                    var v10 float64 = float64(0.0)
                    _ = v10
                }
                var v11 string = "foo bar"
                _ = v11
                _ = v6
            }
            _ = fi5
        }
        _ = fi4
    }
    var v12 int8 = -1
    v2 = (v2 ^ v2)
    _ = len(v3)
    for fi13 := 0; fi13 < 16; fi13++ {
        _ = len(v3)
        v14 := GenS1{F0: 2, F1: 40, F2: 18446744073709551615}
        _ = fi13
        _ = v14
    }
    v12 = (v12 * -50)
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 bool = (-462 < 0)
        for fi17 := 0; fi17 < 4; fi17++ {
            v2 = (v2 ^ 42)
            v18 := GenS1{F0: 1, F1: 236, F2: 18446744073709551615}
            _ = fi17
            _ = v18
        }
        var v19 float32 = 0.0
        _ = fi15
        _ = v16
        _ = v19
    }
    var v20 uint8 = 24
    if true {
        v2 -= v2
    }
    v1 = (v1 + v1)
    v12 = v12
    v12 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v3
}
