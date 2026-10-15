package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int32; F2 bool }
type GenS3 struct { F0 int; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint64 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint16 = 2
    var v3 uint = 1
    var v4 int32 = -773
    var v5 uint8 = 2
    fmt.Println(v4, v3)
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{fi6, (fi6 | fi6), (fi6 - (fi6 + fi6)), fi6}
        if (0.1 == float64((1.5 * float32((-2.25 + 0.0))))) {
            fmt.Println(v1, v2, fi6)
            switch -9223372036854775808 {
            case -138:
                for fi8 := 0; fi8 < 3; fi8++ {
                    if true {
                        v1 += v1
                        v1 -= (^(v1))
                        if (true || true) {
                            var v9 uint64 = 4294967295
                            _ = v9
                        } else {
                            fmt.Println(v4)
                        }
                    } else {
                        v7 = append(v7, fi8)
                        fmt.Println(v5, v1)
                        fmt.Println(fi8, fi8)
                    }
                    switch fi6 {
                    case 9223372036854775807:
                        var v10 float32 = 0.0
                        _ = v10
                    case 959:
                        var v11 float64 = -2.25
                        _ = v11
                    }
                    var v12 string = ""
                    _ = fi8
                    _ = v12
                }
            default:
                for fi13 := 0; fi13 < 6; fi13++ {
                    if (false || (false || (false || false))) {
                        v1 = (v1 - v1)
                        var v14 uint = 678
                        _ = v14
                    }
                    v15 := GenS1{F0: -32768, F1: 959, F2: (v1 ^ 2)}
                    v5 -= v5
                    _ = fi13
                    _ = v15
                }
            }
            if (((!false) && (((1.5 / 1.5) - (0.0 * (((3.14159 * 3.14159) - (1e10 / 1.0)) / (float64(0.1) + (1e10 - ((float64(-2.25) / 3.14159) + ((float64(3.14159) - (1e10 + 0.0)) * -2.25))))))) != -2.25)) || (true || false)) {
                if false {
                    fmt.Println(v2, v2, v2)
                    var v16 int16 = 1000
                    _ = v16
                }
            }
        } else {
            v7 = append(v7, 42)
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            v1 = uint64(v2)
            _ = fi17
        }
        _ = fi6
        _ = v7
    }
    v18 := GenS1{F0: int16(v4), F1: v1, F2: v1}
    v4 += (v4 ^ v4)
    var v19 uint32 = 624
    var v20 uint16 = v2
    v21 := []int{782, 128}
    v22 := GenS2{F0: v19, F1: v4, F2: false}
    if (42 == (^(v5))) {
        v20 = v20
    } else {
        v23 := GenS3{F0: -245, F1: 42, F2: 4294967295}
        var v24 string = "foo bar"
        _ = v23
        _ = v24
    }
    if (true && ((((0.0 + 0.0) < 1.5) || (true || (-278 >= 2147483647))) || true)) {
        var v25 bool = (0 == 0)
        if v25 {
            switch (("\n\t" + "hello") + "hello") {
            case "hello":
                _ = len(v21)
            case "":
                for fi26 := 0; fi26 < 6; fi26++ {
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    for fi27 := 0; fi27 < 11; fi27++ {
                        v25 = false
                        fmt.Println(v1, v19)
                        fmt.Println(v25, v2)
                        _ = fi27
                    }
                    var v28 int = (^(fi26))
                    _ = fi26
                    _ = v28
                }
            default:
                switch "foo bar" {
                case "a":
                    v29 := GenS4{F0: uint8(v2), F1: -1, F2: v1}
                    _ = v29
                }
            }
        }
        _ = v25
    } else {
        if ((float32(3.14159) < (0.0 + (0.1 * 0.0))) || (v1 >= uint64(v2))) {
            for fi30 := 0; fi30 < 11; fi30++ {
                var v31 int8 = -1
                var v32 uint = 2
                _ = fi30
                _ = v31
                _ = v32
            }
            v3 = v3
            var v33 int8 = 9
            _ = v33
        } else {
            v4 = v4
        }
        var v34 uint64 = 4294967295
        if len(v21) > 0 {
            v21[0] = 128
        }
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
