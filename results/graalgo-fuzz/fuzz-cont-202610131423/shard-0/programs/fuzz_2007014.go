package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int8 }

func main() {
    var v1 int32 = 42
    var v2 int16 = 42
    v3 := []int{0, 127, 32767, -995}
    var v4 int8 = 0
    v5 := GenS1{F0: 1}
    var v6 int8 = (v4 << ((78) & 63))
    if true {
        v7 := GenS1{F0: 201}
        if true {
            for fi8 := 0; fi8 < 8; fi8++ {
                switch "a" {
                case "foo bar":
                    fmt.Println(v6, v1, fi8)
                }
                _ = len(v3)
                _ = len(v3)
                _ = fi8
            }
            var v9 int8 = -84
            _ = v9
        }
        _ = v7
    }
    var v10 int8 = v4
    v10 += 0
    fmt.Println(v6, v1, v4)
    v11 := GenS1{F0: 2}
    v1 += v1
    if true {
        if true {
            var v12 int = 9223372036854775807
            _ = v12
        } else {
            if false {
                if (848 < 42) {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        switch "" {
                        case "hello":
                            var v14 float64 = -2.25
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 5; fi15++ {
                            v16 := GenS2{F0: 2, F1: v4}
                            _ = fi15
                            _ = v16
                        }
                        var v17 uint = 255
                        _ = fi13
                        _ = v17
                    }
                } else {
                    var v18 uint16 = 255
                    _ = v18
                }
                v19 := GenS2{F0: 1, F1: v10}
                _ = v19
            } else {
                var v20 int32 = (v1 % (((-(v1))) | 1))
                var v21 uint32 = 579
                var v22 uint32 = 42
                _ = v20
                _ = v21
                _ = v22
            }
        }
        switch ("\n\t" + "") {
        case "a":
            var v23 uint8 = 73
            _ = v23
        case "\n\t":
            var v24 int = 413
            _ = v24
        case "":
            if (1.5 <= 3.14159) {
                _ = len(v3)
            } else {
                var v25 uint32 = 2147483648
                v1 = v1
                fmt.Println(v4, v25, v2)
                _ = v25
            }
        default:
            v26 := GenS1{F0: 1}
            _ = v26
        }
        _ = len(v3)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v27 uint8 = 2
    var v28 int16 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
