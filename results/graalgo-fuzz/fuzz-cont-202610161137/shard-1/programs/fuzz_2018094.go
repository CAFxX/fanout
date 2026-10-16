package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 string; F1 uint8; F2 uint32 }

func main() {
    v1 := []uint32{65535, 4294967295, 1}
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    var v4 bool = false
    fmt.Println(v4)
    fmt.Println(v4)
    var v5 bool = v4
    switch 9223372036854775807 {
    case 32768:
        v5 = (!v5)
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int8 = 127
        v7 = v7
        if (!(65535 > 0)) {
            var v8 float64 = 0.0
            v9 := GenS1{F0: 80, F1: 764}
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 3; fi10++ {
                v4 = v5
                switch "" {
                case "":
                    var v11 bool = (v5 || (v4 || (-2.25 >= -2.25)))
                    _ = v11
                case "foo bar":
                    if false {
                        if v5 {
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                            var v12 int = fi10
                            _ = v12
                        }
                        var v13 bool = (!v5)
                        var v14 float64 = 0.1
                        _ = v13
                        _ = v14
                    }
                case "a":
                    var v15 uint64 = 4294967295
                    _ = v15
                }
                var v16 int32 = 1
                fmt.Println(fi6)
                _ = fi10
                _ = v16
            }
            var v17 int32 = 0
            _ = v17
        }
        _ = fi6
        _ = v7
    }
    v18 := GenS2{F0: -797, F1: 1}
    var v19 int32 = 0
    if (128 <= 255) {
        for fi20 := 0; fi20 < 18; fi20++ {
            if len(v1) > 0 {
                v1[0] = 0
            }
            switch "" {
            case "":
                v21 := v3[128]
                _ = v21
            case "hello":
                for fi22 := 0; fi22 < 19; fi22++ {
                    if v5 {
                        var v23 bool = false
                        v5 = true
                        if (2147483647 > 65535) {
                            var v24 string = "\n\t"
                            _ = v24
                        } else {
                            var v25 int32 = 0
                            v1 = append(v1, 656)
                            _ = v25
                        }
                        _ = v23
                    }
                    v26 := GenS3{F0: -774, F1: 788}
                    _ = fi22
                    _ = v26
                }
            case "foo bar":
                switch (-(fi20)) {
                case 9223372036854775807:
                    var v27 int32 = (v19 & v19)
                    _ = v27
                case -2147483648:
                    v4 = false
                case 32767:
                    v4 = false
                default:
                    fmt.Println(fi20)
                }
            default:
                v5 = (false && (!(((((float64(((3.14159 / -2.25) / 1e10)) - (-2.25 / (0.1 + 1e10))) + 0.1) * -2.25) != ((float64(1.5) / -2.25) / (0.1 + ((0.1 + 1e10) + 1e10)))) && true)))
            }
            v28 := v2[65]
            _ = fi20
            _ = v28
        }
        var v29 uint8 = 217
        v29 = v29
        _ = v29
    } else {
        var v30 int = 127
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi31 := 0; fi31 < 10; fi31++ {
            var v32 float64 = 0.1
            switch (v30 ^ fi31) {
            case -489:
                v4 = (2 == 4294967295)
            case 1:
                var v33 int32 = v19
                _ = v33
            }
            _ = fi31
            _ = v32
        }
        _ = v30
    }
    var v34 uint8 = 246
    if false {
        var v35 float32 = 3.14159
        v36 := GenS1{F0: 42, F1: 9223372036854775807}
        fmt.Println(v35)
        _ = v35
        _ = v36
    }
    switch "hello" {
    case "\n\t":
        v37 := GenS4{F0: 936, F1: 9223372036854775807}
        _ = v37
    case "":
        var v38 uint16 = 42
        _ = v38
    case "a":
        for fi39 := 0; fi39 < 11; fi39++ {
            fmt.Println(fi39, v19, v5)
            v4 = false
            v19 -= v19
            _ = fi39
        }
    }
    for fi40 := 0; fi40 < 11; fi40++ {
        var v41 float32 = (-2.25 - 0.0)
        v5 = false
        _ = fi40
        _ = v41
    }
    v3[42] = 65535
    v42 := GenS5{F0: "a", F1: v34, F2: 2147483648}
    for fi43 := 0; fi43 < 1; fi43++ {
        v19 = v19
        var v44 float32 = 1.5
        _ = fi43
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v5
}
