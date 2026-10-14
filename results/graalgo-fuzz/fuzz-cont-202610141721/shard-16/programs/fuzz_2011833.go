package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint }
type GenS2 struct { F0 int64; F1 int8 }

func main() {
    var v1 int = -2147483648
    var v2 uint = 148
    var v3 uint16 = 32767
    v4 := make(map[string]uint32)
    var v5 uint = 4294967295
    v6 := make(map[string]uint32)
    if true {
        v5 += (v2 - (v2 | 255))
        v3 += (v3 ^ (^(v3)))
        var v7 string = "\n\t"
        _ = v7
    } else {
        for fi8 := 0; fi8 < 5; fi8++ {
            delete(v4, "a")
            _ = fi8
        }
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        switch 32768 {
        case -9223372036854775808:
            if true {
                for fi10 := 0; fi10 < 8; fi10++ {
                    v1 = -930
                    var v11 float32 = 0.0
                    _ = fi10
                    _ = v11
                }
            }
        case 32767:
            var v12 float32 = 1.5
            _ = v12
        case 9223372036854775807:
            v5 = (v5 + v2)
        }
        v1 -= v1
        v5 += (v2 ^ (v2 >> (((^(v2))) & 63)))
        _ = fi9
    }
    switch v1 {
    case 0:
        v5 = 759
    case 32768:
        var v13 float32 = float32(-2.25)
        _ = v13
    default:
        var v14 uint = v5
        _ = v14
    }
    var v15 bool = true
    v2 -= 694
    var v16 uint16 = v3
    if (!(!(!v15))) {
        if (!(v16 <= 2)) {
            v17 := GenS1{F0: "hello", F1: 42, F2: (v5 - v5)}
            switch "\n\t" {
            case "hello":
                var v18 int = (v1 / ((v1) | 1))
                _ = v18
            case "":
                for fi19 := 0; fi19 < 16; fi19++ {
                    if v15 {
                        v20 := GenS1{F0: ("hello" + "foo bar"), F1: 52, F2: (v2 ^ v2)}
                        _ = v20
                    } else {
                        var v21 int64 = -1
                        v1 = 42
                        _ = v21
                    }
                    if v15 {
                        v16 -= (v3 ^ (v3 / ((42) | 1)))
                        v22 := GenS2{F0: -9223372036854775808, F1: 127}
                        v23 := GenS2{F0: 426, F1: 1}
                        _ = v22
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 10; fi24++ {
                        v25 := []int{}
                        v3 = (v3 & v16)
                        _ = fi24
                        _ = v25
                    }
                    _ = fi19
                }
            }
            _ = v17
        } else {
            v26 := GenS2{F0: 881, F1: -45}
            _ = v26
        }
        fmt.Println(v5, v16)
        v2 = uint(v16)
    } else {
        var v27 int64 = -9223372036854775808
        _ = v27
    }
    v2 += 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
