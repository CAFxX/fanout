package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint8 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 3.14159
    var v3 int8 = 0
    fmt.Println(v2, v2)
    if false {
        switch "hello" {
        case "a":
            for fi4 := 0; fi4 < 9; fi4++ {
                if false {
                    var v5 int64 = -57
                    v6 := []int{fi4}
                    _ = v5
                    _ = v6
                }
                v2 = v2
                if false {
                    v3 -= v3
                } else {
                    for fi7 := 0; fi7 < 1; fi7++ {
                        v2 -= float64(0.0)
                        _ = fi7
                    }
                    if false {
                        switch (fi4 + -687) {
                        case 888:
                            v2 = float64(0.1)
                        case 32768:
                            switch (fi4 - fi4) {
                            case 1:
                                switch fi4 {
                                case 400:
                                    for fi8 := 0; fi8 < 13; fi8++ {
                                        fmt.Println(fi8, fi4)
                                        _ = fi8
                                    }
                                case 889:
                                    for fi9 := 0; fi9 < 2; fi9++ {
                                        fmt.Println(v3, fi4, fi9)
                                        _ = fi9
                                    }
                                case 32768:
                                    v2 += -2.25
                                default:
                                    fmt.Println(v2, v3, v2)
                                }
                            case 539:
                                v10 := []int{9223372036854775807}
                                _ = v10
                            case 127:
                                for fi11 := 0; fi11 < 8; fi11++ {
                                    fmt.Println(fi11)
                                    v12 := GenS1{F0: 0, F1: 0, F2: 0}
                                    _ = fi11
                                    _ = v12
                                }
                            }
                        case -2147483648:
                            if (false || (false && (-2.25 <= v2))) {
                                switch "a" {
                                case "":
                                    v1[127] = fi4
                                case "hello":
                                    var v13 bool = (222 == 1)
                                    _ = v13
                                }
                            } else {
                                switch (fi4 & (fi4 % ((42) | 1))) {
                                case -1:
                                    v2 += 3.14159
                                case 526:
                                    v14 := GenS1{F0: -713, F1: (fi4 >> ((42) & 63)), F2: 1}
                                    _ = v14
                                default:
                                    fmt.Println(fi4)
                                }
                                v2 -= (-2.25 + v2)
                                v2 = 1.5
                            }
                        default:
                            var v15 int64 = -1
                            _ = v15
                        }
                    }
                }
                _ = fi4
            }
        case "hello":
            var v16 uint8 = 128
            _ = v16
        default:
            v3 = v3
        }
        switch ("hello" + ("\n\t" + ("hello" + ""))) {
        case "hello":
            for fi17 := 0; fi17 < 16; fi17++ {
                v2 = (1.5 - 3.14159)
                _ = fi17
            }
        default:
            v3 += 24
        }
    }
    v18 := GenS2{F0: 1}
    if true {
        if true {
            v2 = (-2.25 * 1.5)
            var v19 int16 = 0
            _ = v19
        }
        var v20 int16 = 938
        _ = v20
    } else {
        var v21 string = ((("\n\t" + (("" + "a") + (("hello" + "a") + "foo bar"))) + "hello") + "a")
        _ = v21
    }
    v22 := GenS1{F0: 2147483647, F1: -9223372036854775808, F2: 67}
    v3 = (-(v3))
    fmt.Println(v3, v2)
    v23 := []int{2147483647, 42, 32767, 9223372036854775807}
    for fi24 := 0; fi24 < 17; fi24++ {
        for fi25 := 0; fi25 < 8; fi25++ {
            fmt.Println(v2)
            if (false || true) {
                var v26 float32 = float32((v2 + (v2 + (v2 + 3.14159))))
                _ = v26
            }
            _ = fi25
        }
        switch "hello" {
        case "":
            var v27 uint32 = 260
            _ = v27
        default:
            for fi28 := 0; fi28 < 13; fi28++ {
                fmt.Println(v2)
                _ = fi28
            }
        }
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        _ = fi24
    }
    if true {
        fmt.Println(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
