package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int32; F1 int8; F2 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int64; F2 bool }
type GenS5 struct { F0 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 1.5
    var v3 uint = 4294967295
    var v4 bool = (!(v3 >= 4294967295))
    var v5 float32 = 0.0
    v6 := make(map[string]int)
    if (255 != 2) {
        if (148 < 136) {
            for fi7 := 0; fi7 < 11; fi7++ {
                v5 = 0.0
                if ((-2.25 + 3.14159) >= (1.5 / float32(3.14159))) {
                    var v8 bool = ((fi7 > 977) && (true && v4))
                    _ = v8
                } else {
                    switch fi7 {
                    case 69:
                        v9 := GenS1{F0: 789, F1: -265}
                        _ = v9
                    case 2147483647:
                        if false {
                            v10 := GenS2{F0: 0, F1: 1, F2: -1}
                            _ = v10
                        }
                    default:
                        fmt.Println(fi7, v2)
                    }
                    for fi11 := 0; fi11 < 9; fi11++ {
                        switch (-(fi7)) {
                        case -725:
                            if (false || (v3 >= (v3 % ((4294967295) | 1)))) {
                                v1[2] = 565
                                fmt.Println(v4)
                            }
                            switch fi7 {
                            case 2147483647:
                                fmt.Println(v2, v3)
                            }
                        }
                        var v12 int8 = -1
                        var v13 int64 = -876
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    var v14 uint64 = 18446744073709551615
                    _ = v14
                }
                switch "foo bar" {
                case "hello":
                    var v15 uint16 = 1
                    _ = v15
                case "\n\t":
                    var v16 bool = (!false)
                    _ = v16
                }
                _ = fi7
            }
        } else {
            switch "\n\t" {
            case "\n\t":
                if (v4 && v4) {
                    var v17 uint32 = 65535
                    if v4 {
                        for fi18 := 0; fi18 < 3; fi18++ {
                            var v19 int32 = -2147483648
                            _ = fi18
                            _ = v19
                        }
                    } else {
                        v20 := GenS1{F0: 1000, F1: 9223372036854775807}
                        _ = v20
                    }
                    _ = v17
                }
            case "":
                var v21 uint8 = 0
                _ = v21
            case "hello":
                v22 := GenS3{F0: 42}
                _ = v22
            }
        }
        v23 := GenS4{F0: 889, F1: 691, F2: (!false)}
        fmt.Println(v4)
        _ = v23
    }
    v3 = (v3 + v3)
    var v24 uint8 = 2
    var v25 int16 = 1
    var v26 int = 32768
    v27 := GenS5{F0: (^(v3))}
    v5 = float32(float64(-2.25))
    var v28 string = ("" + "")
    fmt.Println(v2)
    var v29 string = ((v28 + v28) + v28)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
