package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint8; F2 int8 }

func main() {
    var v1 uint16 = 1
    var v2 uint16 = 65535
    var v3 uint32 = 1
    v4 := []string{"hello", "a"}
    var v5 uint = 1
    var v6 int16 = 0
    var v7 uint = (v5 | v5)
    v1 -= (^(v1))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = 4294967295
    v8 := GenS1{F0: v6, F1: 65535, F2: 18446744073709551615}
    var v9 int16 = 1
    var v10 int = -938
    for fi11 := 0; fi11 < 20; fi11++ {
        _ = len(v4)
        if (171 != 2) {
            if true {
                fmt.Println(v5, v2)
                var v12 uint = (^(v7))
                _ = v12
            } else {
                if false {
                    if true {
                        fmt.Println(v6)
                    } else {
                        if len(v4) > 0 {
                            v4[0] = "a"
                        }
                    }
                    var v13 int8 = 42
                    _ = v13
                }
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := GenS2{F0: true, F1: 255, F2: -128}
                v16 := GenS2{F0: ((v7 | v7) <= v7), F1: 2, F2: -121}
                _ = fi14
                _ = v15
                _ = v16
            }
            v17 := GenS1{F0: 332, F1: v2, F2: 637}
            _ = v17
        }
        var v18 bool = (-105 == 28)
        _ = fi11
        _ = v18
    }
    var v19 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
