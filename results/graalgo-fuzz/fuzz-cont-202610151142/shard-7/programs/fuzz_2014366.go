package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int64; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -532
    var v3 int = v2
    if ((3.14159 * 0.0) < float32((1.5 / 3.14159))) {
        v4 := GenS1{F0: 717, F1: true}
        switch "a" {
        case "a":
            if ((^(v2)) < (v3 ^ v2)) {
                var v5 uint8 = 127
                for fi6 := 0; fi6 < 8; fi6++ {
                    v7 := []int{}
                    v8 := GenS1{F0: 42, F1: true}
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                if true {
                    v3 += (v2 | (v3 * 998))
                } else {
                    for fi9 := 0; fi9 < 6; fi9++ {
                        switch fi9 {
                        case 4294967295:
                            v10 := GenS1{F0: 4294967295, F1: false}
                            _ = v10
                        }
                        var v11 int64 = 9223372036854775807
                        var v12 bool = (!(1.5 >= 0.1))
                        _ = fi9
                        _ = v11
                        _ = v12
                    }
                    switch (v3 % ((9223372036854775807) | 1)) {
                    case 9223372036854775807:
                        v5 = (v5 * (v5 >> ((980) & 63)))
                    default:
                        v13 := []int{v2}
                        _ = v13
                    }
                    v2 = v3
                }
                _ = v5
            }
        }
        v2 = v2
        _ = v4
    } else {
        v14 := GenS1{F0: 18446744073709551615, F1: false}
        _ = v14
    }
    v15 := GenS2{F0: 255, F1: "\n\t"}
    fmt.Println(v3, v3)
    delete(v1, ("\n\t" + "a"))
    var v16 int8 = -97
    v3 += 42
    if false {
        v17 := GenS3{F0: 131, F1: v3}
        _ = v17
    }
    v18 := v1["foo bar"]
    delete(v1, "a")
    v18 = (v18 >> ((400) & 63))
    for fi19 := 0; fi19 < 2; fi19++ {
        if (1e10 < 0.1) {
            v20 := GenS4{F0: -9223372036854775808, F1: 42}
            switch "hello" {
            case "":
                v21 := GenS1{F0: 42, F1: true}
                _ = v21
            case "hello":
                switch (fi19 >> ((42) & 63)) {
                case 32767:
                    var v22 int32 = 511
                    _ = v22
                case 4294967295:
                    v16 = v16
                }
            case "a":
                for fi23 := 0; fi23 < 5; fi23++ {
                    var v24 int16 = -32768
                    fmt.Println(v24, v3, v24)
                    var v25 uint32 = 54
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            default:
                v26 := GenS4{F0: 649, F1: 2147483647}
                _ = v26
            }
            _ = v20
        } else {
            v16 -= (v16 / (((v16 | (v16 + (v16 * v16)))) | 1))
            if false {
                switch "foo bar" {
                case "":
                    var v27 string = "foo bar"
                    _ = v27
                case "foo bar":
                    var v28 bool = (!false)
                    _ = v28
                default:
                    var v29 uint = 85
                    _ = v29
                }
                v2 -= -1
                v30 := GenS4{F0: 1, F1: 42}
                _ = v30
            } else {
                v31 := []int{9223372036854775807}
                v16 -= (v16 << ((255) & 63))
                _ = v31
            }
        }
        v16 = (-(v16))
        _ = fi19
    }
    v16 = v16
    var v32 int64 = 42
    fmt.Println(v18, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v32
}
