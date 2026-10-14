package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint32 }
type GenS2 struct { F0 uint32; F1 string; F2 int8 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int32; F1 bool }

func main() {
    var v1 int = 128
    var v2 int = v1
    v3 := make(map[int]uint32)
    v4 := GenS1{F0: 255, F1: ("hello" + "hello"), F2: 2147483648}
    v5 := GenS1{F0: 934, F1: "", F2: 2}
    var v6 float64 = 0.0
    switch (v2 + 155) {
    case 32768:
        var v7 string = "\n\t"
        _ = v7
    default:
        v8 := GenS2{F0: 715, F1: "a", F2: 127}
        _ = v8
    }
    v2 = v2
    v6 = v6
    fmt.Println(v6, v6, v2)
    switch "" {
    case "foo bar":
        var v9 int64 = 522
        _ = v9
    }
    var v10 uint32 = 42
    v3[(^(v2))] = v10
    var v11 int32 = 2147483647
    var v12 int64 = -867
    v2 = (-(v1))
    v13 := GenS3{F0: 1, F1: v12}
    var v14 bool = (!false)
    var v15 uint16 = 42
    if false {
        v16 := GenS4{F0: 0, F1: (!false)}
        v11 = (-(v11))
        var v17 uint32 = 680
        _ = v16
        _ = v17
    } else {
        v18 := GenS4{F0: v11, F1: v14}
        var v19 uint32 = v10
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 int16 = 1000
            switch "" {
            case "foo bar":
                for fi22 := 0; fi22 < 12; fi22++ {
                    if v14 {
                        var v23 uint32 = (v19 / (((^(v19))) | 1))
                        if v14 {
                            if (!(((v1 < (fi22 >> ((0) & 63))) || v14) || true)) {
                                delete(v3, fi20)
                                v14 = v14
                                v24 := GenS3{F0: v21, F1: (v12 + (^(v12)))}
                                _ = v24
                            }
                            var v25 uint32 = v10
                            var v26 int64 = (v12 % (((-(v12))) | 1))
                            _ = v25
                            _ = v26
                        }
                        _ = v23
                    } else {
                        var v27 float64 = (-2.25 + (v6 / v6))
                        switch fi22 {
                        case 115:
                            v28 := GenS3{F0: -1, F1: (v12 & (v12 | (v12 << ((1) & 63))))}
                            _ = v28
                        default:
                            fmt.Println(fi22)
                        }
                        var v29 uint = 42
                        _ = v27
                        _ = v29
                    }
                    fmt.Println(v12)
                    var v30 uint16 = (v15 & 255)
                    _ = fi22
                    _ = v30
                }
            default:
                for fi31 := 0; fi31 < 1; fi31++ {
                    switch (v2 << ((65535) & 63)) {
                    case 421:
                        v14 = ((2 >= 0) || (false || (true && v14)))
                    default:
                        switch (v2 / ((656) | 1)) {
                        case 2147483647:
                            v32 := GenS1{F0: 32767, F1: "hello", F2: v10}
                            _ = v32
                        }
                    }
                    fmt.Println(v10, fi20)
                    _ = fi31
                }
            }
            _ = fi20
            _ = v21
        }
        _ = v18
        _ = v19
    }
    var v33 int = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
