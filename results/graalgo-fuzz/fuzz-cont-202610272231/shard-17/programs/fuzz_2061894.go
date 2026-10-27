package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    v5 := []int{42, 252, 32768, 128}
    fmt.Println(v1)
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: fi6, F1: 988}
        if true {
            v8 := GenS1{F0: fi6, F1: 1000}
            switch fi6 {
            case 128:
                if (738 <= 639) {
                    if (!true) {
                        if false {
                            v1 += (v1 >> ((681) & 63))
                            var v9 uint16 = 65535
                            _ = v9
                        }
                        v10 := GenS1{F0: fi6, F1: 42}
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 4; fi11++ {
                        var v12 uint64 = 0
                        switch fi11 {
                        case 128:
                            if ("\n\t" == "\n\t") {
                                switch (fi6 / ((2147483647) | 1)) {
                                case -9223372036854775808:
                                    v2[fi11] = (("hello" + "hello") + "foo bar")
                                default:
                                    v13 := GenS1{F0: fi6, F1: 32767}
                                    _ = v13
                                }
                                v14 := GenS1{F0: (fi6 % ((fi6) | 1)), F1: -459}
                                _ = v14
                            } else {
                                var v15 int = (fi11 % (((fi6 >> ((99) & 63))) | 1))
                                if len(v5) > 0 {
                                    v5[0] = 4294967295
                                }
                                v12 += 1
                                _ = v15
                            }
                        case 42:
                            if false {
                                var v16 int8 = -127
                                _ = v16
                            }
                        default:
                            v5 = append(v5, (fi11 * -182))
                        }
                        v5 = append(v5, -1)
                        _ = fi11
                        _ = v12
                    }
                    for fi17 := 0; fi17 < 10; fi17++ {
                        fmt.Println(fi17)
                        if len(v5) > 0 {
                            v5[0] = (fi17 % ((128) | 1))
                        }
                        var v18 string = "hello"
                        _ = fi17
                        _ = v18
                    }
                }
            case 9223372036854775807:
                v1 = (v1 / (((v1 << ((1) & 63))) | 1))
            case 937:
                v1 += (v1 >> ((18446744073709551615) & 63))
            default:
                v1 += (v1 ^ v1)
            }
            _ = v8
        }
        fmt.Println(v1, fi6)
        _ = fi6
        _ = v7
    }
    var v19 int8 = 100
    var v20 int32 = 668
    v20 += (v20 << ((272) & 63))
    var v21 int = 9223372036854775807
    var v22 float64 = 1e10
    v22 = (v22 + 0.1)
    fmt.Println(v1)
    fmt.Println(v1, v22, v21)
    var v23 uint64 = 0
    v23 += 1
    var v24 float32 = 0.0
    var v25 int64 = 48
    if len(v5) > 0 {
        v5[0] = v21
    }
    fmt.Println(v19, v19, v19)
    switch v21 {
    case 0:
        fmt.Println(v22)
    case 2147483647:
        fmt.Println(v22, v25)
    case -3:
        v24 -= v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
