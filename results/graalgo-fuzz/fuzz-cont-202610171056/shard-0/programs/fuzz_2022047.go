package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() bool {
    v4 := []int{-1, 4294967295, 4294967295, 32768}
    var v5 uint8 = 2
    _ = v4
    _ = v5
    return (false && false)
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []string{""}
    var v3 int32 = 2147483647
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v6 int32 = 0
    v3 = v6
    var v7 uint64 = (v1 ^ v1)
    v7 = v1
    if true {
        if true {
            v3 += (v3 + v6)
            v1 = 42
        } else {
            var v8 string = "hello"
            fmt.Println(v6)
            v9 := GenS1{F0: 952}
            _ = v8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v3 -= 1
        _ = fi10
    }
    fmt.Println(v6, v7, v1)
    switch "hello" {
    case "a":
        switch 128 {
        case -9223372036854775808:
            var v11 uint8 = 149
            _ = v11
        case 2147483647:
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                fmt.Println(v7, v3)
            }
        case 0:
            var v12 uint = 2
            _ = v12
        }
    case "hello":
        var v13 int = 261
        _ = v13
    case "\n\t":
        for fi14 := 0; fi14 < 4; fi14++ {
            for fi15 := 0; fi15 < 2; fi15++ {
                v2 = append(v2, "a")
                if (42 >= 0) {
                    var v16 uint32 = 4294967295
                    var v17 float64 = ((float64(((((1e10 - 0.0) - float32(((((0.1 * 1.5) + (((3.14159 + ((1.5 - (0.0 - 1e10)) - 1.5)) * (float64(1.5) + float64(1.5))) + ((0.0 * ((1.5 / ((0.1 - ((1e10 / ((1.5 + (float64(1.5) / float64((3.14159 - (3.14159 - -2.25))))) + float64(-2.25))) + (float64(3.14159) + -2.25))) * float64(3.14159))) + 1e10)) * float64(float32((0.1 * 0.0)))))) - ((3.14159 - 0.0) * (1e10 / 1.5))) * 0.0))) - -2.25) + float32(1.5))) + 1.5) + 1e10)
                    var v18 float64 = (((v17 - ((((v17 / (v17 - ((v17 + ((v17 + (1.5 - v17)) + 0.0)) * 1e10))) + (v17 / ((0.1 / (1.5 * (1e10 - v17))) / 3.14159))) + v17) / (float64(((-2.25 / (((0.1 + (0.0 - 1e10)) + 0.1) - 0.0)) - 0.0)) - v17))) / ((1e10 / 1.5) * v17)) * v17)
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    var v19 bool = false
                    _ = v19
                }
                if (42 < 1000) {
                    _ = len(v2)
                    v20 := GenS1{F0: 65535}
                    _ = v20
                } else {
                    v3 = (-(v6))
                    var v21 int16 = 983
                    _ = v21
                }
                _ = fi15
            }
            v3 -= 42
            _ = fi14
        }
    }
    var v22 uint = 377
    var v23 int8 = 74
    v1 -= (v7 << (((v22 >> (((v22 + 18446744073709551615)) & 63))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v6
    _ = v7
}
