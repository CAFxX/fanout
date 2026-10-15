package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    var v1 bool = true
    var v2 string = "foo bar"
    var v3 int16 = -32768
    v4 := make(map[string]int)
    var v5 uint = 2
    v6 := []uint32{400}
    v1 = (838 > 1)
    var v7 int16 = v3
    v5 -= (v5 / (((v5 + (v5 / ((v5) | 1)))) | 1))
    for fi8 := 0; fi8 < 19; fi8++ {
        if (0.0 >= 0.0) {
            v2 = ("a" + (v2 + (v2 + "hello")))
            v1 = v1
            v7 = v7
        }
        _ = fi8
    }
    if ((0.1 + 1e10) < 1.5) {
        for fi9 := 0; fi9 < 14; fi9++ {
            fmt.Println(v5, fi9, v1)
            _ = fi9
        }
        var v10 int32 = 0
        _ = v10
    } else {
        for fi11 := 0; fi11 < 19; fi11++ {
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 int16 = (v3 ^ v7)
                v6 = append(v6, 723)
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        delete(v4, ("foo bar" + v2))
        v5 = 2
    }
    var v14 uint64 = uint64(v5)
    v15 := GenS1{F0: (^(v14)), F1: (42 >= int32(v14)), F2: 2147483647}
    fmt.Println(v1, v5, v2)
    var v16 uint16 = 268
    var v17 uint8 = 255
    var v18 float64 = float64((float32(float64(1.5)) - 0.0))
    var v19 uint64 = uint64(v17)
    v14 += (v19 | (v14 % (((^(v19))) | 1)))
    var v20 int8 = -98
    if v1 {
        for fi21 := 0; fi21 < 3; fi21++ {
            v20 = v20
            _ = fi21
        }
        v3 -= 1
    } else {
        if v1 {
            for fi22 := 0; fi22 < 16; fi22++ {
                fmt.Println(v7, v17)
                for fi23 := 0; fi23 < 14; fi23++ {
                    switch (v2 + v2) {
                    case "hello":
                        v24 := GenS1{F0: v14, F1: (-571 < (v7 ^ (v3 + (-(v3))))), F2: 4294967295}
                        _ = v24
                    case "a":
                        var v25 float32 = 1e10
                        _ = v25
                    case "foo bar":
                        if ((0 <= 202) && v1) {
                            v26 := GenS2{F0: v2}
                            var v27 int16 = (^(v3))
                            _ = v26
                            _ = v27
                        }
                    default:
                        v28 := GenS2{F0: "a"}
                        _ = v28
                    }
                    if ((v17 | v17) >= 128) {
                        var v29 int64 = int64(v20)
                        _ = v29
                    }
                    _ = fi23
                }
                _ = fi22
            }
            v30 := GenS2{F0: "hello"}
            _ = v30
        } else {
            for fi31 := 0; fi31 < 18; fi31++ {
                switch v2 {
                case "foo bar":
                    var v32 uint8 = 119
                    _ = v32
                default:
                    switch (fi31 + fi31) {
                    case 1:
                        v33 := GenS1{F0: (^(v14)), F1: (!(!true)), F2: 65535}
                        _ = v33
                    case 32768:
                        fmt.Println(v16, v7, v19)
                    case 128:
                        v34 := GenS3{F0: -9223372036854775808, F1: 2}
                        _ = v34
                    }
                }
                fmt.Println(v17, v18)
                _ = fi31
            }
            v35 := v4[("foo bar" + "\n\t")]
            _ = v35
        }
        _ = len(v6)
    }
    v2 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
