package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS3 struct { F0 string; F1 bool }
func genFn0() int16 {
    var v6 int32 = -605
    var v7 float64 = -2.25
    v8 := []int{-9223372036854775808, -9223372036854775808, int(v6), -9223372036854775808}
    switch 32767 {
    case 1:
        if false {
            v8 = append(v8, -9223372036854775808)
            switch "\n\t" {
            case "\n\t":
                v6 = (-(v6))
            }
        } else {
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 int64 = 1
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: v6}
            _ = v11
        }
    }
    _ = v6
    _ = v7
    _ = v8
    return 690
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 255
    v3 := make(map[int]int)
    v4 := []uint32{249}
    var v5 uint = 319
    var v12 uint64 = 18446744073709551615
    v13 := GenS1{F0: 392}
    var v14 int16 = 0
    if (!(v5 == (^(v5)))) {
        v15 := v3[32768]
        _ = v15
    }
    var v16 bool = false
    if (v16 || (((!((v1 + (((v1 + v1) + v1) + "\n\t")) > (((v1 + ((v1 + v1) + "")) + ("foo bar" + "")) + v1))) && v16) || (-28 < 127))) {
        v2 = 255
        v12 = v12
        v17 := GenS1{F0: -1}
        _ = v17
    }
    v18 := GenS2{F0: v12, F1: 4294967295, F2: v2}
    fmt.Println(v1)
    switch (("" + v1) + "foo bar") {
    case "hello":
        var v19 int16 = (v14 % (((-(v14))) | 1))
        _ = v19
    }
    delete(v3, -806)
    v20 := GenS3{F0: (v1 + ""), F1: (float64((3.14159 - 1.5)) == 0.0)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
