package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 int }
type GenS3 struct { F0 int64; F1 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 42
    var v3 bool = false
    var v4 int = 1
    switch "a" {
    case "foo bar":
        if true {
            v5 := GenS1{F0: false, F1: -32768}
            v6 := v1[""]
            if false {
                v7 := GenS1{F0: v3, F1: -32768}
                _ = v7
            }
            _ = v5
            _ = v6
        } else {
            var v8 int64 = 0
            _ = v8
        }
    case "hello":
        v4 = v4
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        v4 -= v4
        var v10 float32 = 3.14159
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        fmt.Println(fi11)
        var v12 uint = 65535
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: v3, F1: 1000}
    if false {
        switch "foo bar" {
        case "a":
            fmt.Println(v2)
        case "foo bar":
            if true {
                v14 := GenS2{F0: 42, F1: 42, F2: v4}
                _ = v14
            }
        }
    } else {
        for fi15 := 0; fi15 < 4; fi15++ {
            if (0 <= 1) {
                v16 := GenS2{F0: 1000, F1: 999, F2: v4}
                delete(v1, ("\n\t" + "\n\t"))
                _ = v16
            } else {
                var v17 int8 = 84
                v18 := v1["\n\t"]
                _ = v17
                _ = v18
            }
            v19 := v1["foo bar"]
            if true {
                v19 = ((((v19 + "") + v19) + v19) + (v19 + "\n\t"))
            }
            _ = fi15
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        switch "a" {
        case "hello":
            _ = len(v1)
        default:
            _ = len(v1)
        }
        if v3 {
            fmt.Println(v3)
            var v21 int8 = 127
            if (true && v3) {
                if (!((1 <= 1) && v3)) {
                    v4 = fi20
                    if v3 {
                        v4 += (v4 >> ((65535) & 63))
                    } else {
                        var v22 bool = v3
                        v22 = (1 <= 42)
                        _ = v22
                    }
                }
                var v23 float32 = 0.0
                v24 := GenS3{F0: 1, F1: -2147483648}
                _ = v23
                _ = v24
            } else {
                var v25 float64 = (0.0 * (float64((float32(-2.25) + ((3.14159 / 3.14159) * 1.5))) * 1.5))
                switch fi20 {
                case 32768:
                    if (-598 < -812) {
                        for fi26 := 0; fi26 < 1; fi26++ {
                            v1["foo bar"] = "hello"
                            v27 := v1["\n\t"]
                            var v28 float64 = v25
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    }
                default:
                    v25 += (float64((1e10 + (-2.25 - 1.5))) + 3.14159)
                }
                for fi29 := 0; fi29 < 19; fi29++ {
                    v4 = (v4 + (-(v4)))
                    v4 -= v4
                    _ = fi29
                }
                _ = v25
            }
            _ = v21
        } else {
            v2 = v2
        }
        v3 = v3
        _ = fi20
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        v31 := []int{}
        _ = fi30
        _ = v31
    }
    var v32 int = (v4 * (v4 ^ -1))
    fmt.Println(v3)
    v3 = ((32767 == 0) || (25 >= 891))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
