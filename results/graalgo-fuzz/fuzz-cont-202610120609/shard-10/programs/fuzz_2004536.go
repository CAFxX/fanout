package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int }
func genFn0() int32 {
    var v7 int32 = -157
    var v8 int64 = 1
    _ = v7
    _ = v8
    return (v7 % ((v7) | 1))
}
func genFn1() int {
    var v9 int64 = 0
    v9 -= v9
    _ = v9
    return -912
}

func main() {
    var v1 uint8 = 127
    var v2 uint16 = 32767
    var v3 uint8 = (v1 - 128)
    v4 := []uint32{65535}
    var v5 uint64 = 85
    var v6 uint8 = v3
    switch "\n\t" {
    case "foo bar":
        var v10 uint32 = 939
        _ = v10
    }
    var v11 uint64 = v5
    fmt.Println(v5, v11, v1)
    _ = len(v4)
    v12 := GenS1{F0: true}
    v1 += 42
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        if false {
            var v13 int8 = int8(v6)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v14 float64 = ((float64((3.14159 + -2.25)) / 1.5) - -2.25)
            _ = v13
            _ = v14
        } else {
            if (false && false) {
                var v15 uint32 = 1
                _ = v15
            }
            v16 := GenS1{F0: true}
            if len(v4) > 0 {
                v4[0] = 50
            }
            _ = v16
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            v11 = (v11 / (((v11 + v11)) | 1))
            _ = fi17
        }
    } else {
        var v18 uint64 = (^(v5))
        v19 := GenS2{F0: -1, F1: -651}
        fmt.Println(v6)
        _ = v18
        _ = v19
    }
    switch 9223372036854775807 {
    case -2147483648:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
