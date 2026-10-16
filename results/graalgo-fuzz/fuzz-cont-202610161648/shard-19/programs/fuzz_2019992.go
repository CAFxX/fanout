package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int16; F2 string }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint64 = 42
    var v2 uint64 = (v1 >> ((42) & 63))
    var v3 float64 = 0.0
    var v4 uint = 0
    var v5 uint = v4
    var v6 bool = false
    if ((v2 == (v2 % ((18446744073709551615) | 1))) && (57 > 2)) {
        v2 = v1
        v7 := GenS1{F0: "a"}
        _ = v7
    } else {
        switch -198 {
        case 32767:
            var v8 uint8 = 127
            _ = v8
        }
        fmt.Println(v3)
        v9 := []int{-274, -746, 128}
        _ = v9
    }
    var v10 uint8 = 127
    v4 += (v4 & (v5 >> (((^(v5))) & 63)))
    v11 := GenS2{F0: v1, F1: -32768, F2: "\n\t"}
    if v6 {
        var v12 float32 = -2.25
        fmt.Println(v12)
        _ = v12
    }
    v13 := GenS1{F0: ("" + "hello")}
    v14 := GenS3{F0: v1}
    var v15 uint = uint(v10)
    var v16 uint8 = v10
    switch 32768 {
    case 2147483647:
        var v17 int8 = 0
        _ = v17
    case 32767:
        fmt.Println(v4, v2)
    case -534:
        for fi18 := 0; fi18 < 9; fi18++ {
            v4 = 0
            v19 := []int{(fi18 >> (((v5 ^ v15)) & 63)), (-(fi18)), fi18, (fi18 * fi18)}
            _ = fi18
            _ = v19
        }
    }
    if ((-92 > -274) && v6) {
        var v20 uint = (v5 % ((v5) | 1))
        v5 += 255
        _ = v20
    }
    v15 += (v4 >> ((v5) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
