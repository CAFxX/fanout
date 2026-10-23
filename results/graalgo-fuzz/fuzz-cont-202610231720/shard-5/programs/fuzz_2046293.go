package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint; F1 int8; F2 int32 }

func main() {
    var v1 int16 = 42
    v2 := make(map[int]bool)
    var v3 uint = 18446744073709551615
    switch -2147483648 {
    case 32768:
        v4 := GenS1{F0: -1, F1: 4294967295, F2: 797}
        _ = v4
    case 2147483647:
        v5 := GenS1{F0: 286, F1: (v3 + (v3 - 653)), F2: -476}
        _ = v5
    case 0:
        v1 -= v1
    default:
        var v6 uint8 = 198
        _ = v6
    }
    if (v3 < v3) {
        v7 := []int{}
        var v8 int8 = 42
        var v9 int64 = 952
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 float32 = 0.0
    fmt.Println(v10, v3)
    var v11 uint16 = 355
    switch 0 {
    case 21:
        v3 += 0
    default:
        v12 := GenS1{F0: 1, F1: v3, F2: 42}
        _ = v12
    }
    fmt.Println(v1, v11, v10)
    v13 := []int{32768, 527, -9223372036854775808, 770}
    switch ("hello" + "\n\t") {
    case "hello":
        v10 = v10
    case "a":
        v1 -= v1
    case "foo bar":
        v10 -= 1e10
    }
    fmt.Println(v11, v1)
    for fi14 := 0; fi14 < 12; fi14++ {
        switch (-(fi14)) {
        case 395:
            for fi15 := 0; fi15 < 10; fi15++ {
                fmt.Println(v10, v10, fi14)
                var v16 uint16 = (^(v11))
                if true {
                    var v17 float32 = float32(1.5)
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
        case 32767:
            fmt.Println(fi14, v10, v3)
        }
        for fi18 := 0; fi18 < 13; fi18++ {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = len(v13)
            fmt.Println(v11, v1, fi14)
            var v19 int32 = 0
            v19 += (v19 - (-(v19)))
            _ = fi18
            _ = v19
        }
        if false {
            var v20 int64 = 65
            if true {
                var v21 string = ("foo bar" + "foo bar")
                _ = v21
            }
            switch ("a" + "foo bar") {
            case "foo bar":
                var v22 int8 = 0
                _ = v22
            case "\n\t":
                if true {
                    for fi23 := 0; fi23 < 6; fi23++ {
                        switch fi14 {
                        case 569:
                            switch fi14 {
                            case -229:
                                v3 += v3
                            case 1:
                                var v24 bool = true
                                _ = v24
                            }
                        case -250:
                            if (!true) {
                                var v25 int64 = -1
                                _ = v25
                            }
                        case -486:
                            v26 := GenS1{F0: 9223372036854775807, F1: v3, F2: 811}
                            _ = v26
                        default:
                            if (((v20 / (((v20 / ((v20) | 1))) | 1)) >= v20) && (100 >= 127)) {
                                if (true && false) {
                                    v3 -= (^(v3))
                                }
                                var v27 float32 = float32(0.0)
                                var v28 int16 = -21
                                _ = v27
                                _ = v28
                            }
                        }
                        var v29 float32 = v10
                        v30 := v2[int(v3)]
                        _ = fi23
                        _ = v29
                        _ = v30
                    }
                    if (false || false) {
                        var v31 int16 = v1
                        var v32 string = "\n\t"
                        _ = v31
                        _ = v32
                    } else {
                        delete(v2, (^(fi14)))
                    }
                } else {
                    v33 := GenS2{F0: 18446744073709551615, F1: (v3 / ((2) | 1))}
                    var v34 string = "foo bar"
                    v35 := GenS1{F0: (fi14 / ((-561) | 1)), F1: 1, F2: -1}
                    _ = v33
                    _ = v34
                    _ = v35
                }
            case "hello":
                v11 = (^(v11))
            default:
                v20 += v20
            }
            _ = v20
        } else {
            for fi36 := 0; fi36 < 7; fi36++ {
                switch 745 {
                case 2147483647:
                    var v37 string = "\n\t"
                    _ = v37
                case 11:
                    var v38 uint64 = 964
                    _ = v38
                }
                _ = fi36
            }
            v39 := GenS3{F0: 42, F1: 193, F2: v11}
            v10 = float32(1e10)
            _ = v39
        }
        _ = fi14
    }
    var v40 uint = 18446744073709551615
    v40 += v40
    v41 := GenS4{F0: v40, F1: 0, F2: -697}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v40
    _ = v41
}
