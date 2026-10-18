package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint32 = 149
    var v2 float32 = 0.1
    var v3 uint64 = 1
    var v4 string = "foo bar"
    v1 += 655
    fmt.Println(v2)
    for fi5 := 0; fi5 < 6; fi5++ {
        switch "foo bar" {
        case "":
            fmt.Println(v3, v1, v1)
        case "\n\t":
            v3 = 1
        case "a":
            v4 = ("a" + v4)
        default:
            v6 := GenS1{F0: 9223372036854775807, F1: 0, F2: (fi5 % ((fi5) | 1))}
            _ = v6
        }
        _ = fi5
    }
    if ((0.1 < (0.0 + float64((float32(1e10) + v2)))) && (-18 > 0)) {
        v3 = (^(v3))
    } else {
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := GenS1{F0: 9223372036854775807, F1: -128, F2: fi7}
            v9 := GenS1{F0: -1, F1: -1, F2: (fi7 + 2147483647)}
            var v10 bool = false
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        if false {
            for fi11 := 0; fi11 < 9; fi11++ {
                for fi12 := 0; fi12 < 9; fi12++ {
                    var v13 uint16 = 42
                    v14 := []int{2147483647, (fi12 >> ((42) & 63)), -1, -295}
                    switch fi12 {
                    case 320:
                        _ = len(v14)
                    default:
                        v15 := GenS2{F0: -1, F1: 9223372036854775807}
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                v16 := []int{fi11, (fi11 | (fi11 + 128)), -973}
                _ = fi11
                _ = v16
            }
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 uint32 = v1
                _ = fi17
                _ = v18
            }
        }
        if ((^(v1)) > v1) {
            v19 := GenS3{F0: (^(v3))}
            if ((^(v3)) < 0) {
                v20 := []int{127}
                _ = v20
            }
            _ = v19
        } else {
            var v21 string = "\n\t"
            v1 -= v1
            _ = v21
        }
    }
    v22 := []int{2147483647, -96, 127, -1}
    v23 := GenS2{F0: -1, F1: 9223372036854775807}
    v2 = v2
    fmt.Println(v3, v1)
    switch v4 {
    case "foo bar":
        fmt.Println(v3)
    case "a":
        for fi24 := 0; fi24 < 6; fi24++ {
            switch -2147483648 {
            case 127:
                for fi25 := 0; fi25 < 2; fi25++ {
                    if false {
                        fmt.Println(v4, v4, fi25)
                        v3 += v3
                        if ((v1 >> ((65535) & 63)) <= v1) {
                            v2 -= v2
                            var v26 string = "hello"
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                            _ = v26
                        }
                    }
                    switch fi25 {
                    case -1:
                        if ((834 == 42) || (706 > -1)) {
                            var v27 int16 = -32768
                            var v28 int64 = 0
                            _ = v27
                            _ = v28
                        } else {
                            fmt.Println(v1)
                        }
                    case 9223372036854775807:
                        v4 = (v4 + v4)
                    default:
                        if (v2 >= v2) {
                            v1 += v1
                        }
                    }
                    if false {
                        v29 := GenS1{F0: -822, F1: 127, F2: (fi25 * fi25)}
                        _ = v29
                    }
                    _ = fi25
                }
            default:
                v30 := GenS2{F0: -128, F1: 42}
                _ = v30
            }
            for fi31 := 0; fi31 < 6; fi31++ {
                v22 = append(v22, fi31)
                _ = fi31
            }
            _ = fi24
        }
    default:
        for fi32 := 0; fi32 < 18; fi32++ {
            fmt.Println(v1)
            fmt.Println(v3, fi32, fi32)
            _ = fi32
        }
    }
    _ = len(v22)
    for fi33 := 0; fi33 < 4; fi33++ {
        var v34 uint = 65535
        _ = fi33
        _ = v34
    }
    v35 := GenS2{F0: 104, F1: 970}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v35
    _ = v4
}
