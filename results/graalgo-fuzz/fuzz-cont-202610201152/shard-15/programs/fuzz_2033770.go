package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint }
type GenS3 struct { F0 uint8; F1 int8; F2 int32 }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]bool)
    var v3 float32 = 0.0
    var v4 string = ("foo bar" + ("" + "a"))
    var v5 string = v4
    v6 := make(map[uint32]uint32)
    v7 := GenS1{F0: 1, F1: 1, F2: 1}
    switch (("" + (v4 + v5)) + ("" + (v5 + v4))) {
    case "":
        fmt.Println(v1)
    default:
        if false {
            fmt.Println(v5, v4)
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 string = v4
                v10 := GenS2{F0: v1, F1: 511, F2: 530}
                var v11 bool = false
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            delete(v2, 2147483647)
        } else {
            if false {
                v4 = ""
                v1 = (v1 * v1)
            } else {
                if true {
                    v6[uint32(v1)] = 0
                    var v12 int64 = 1
                    _ = v12
                }
            }
            delete(v6, 2147483647)
            v13 := GenS2{F0: (v1 / (((^(v1))) | 1)), F1: 2147483648, F2: 973}
            _ = v13
        }
    }
    fmt.Println(v5)
    v4 = v5
    v14 := GenS1{F0: 2, F1: 0, F2: -315}
    if ((!true) || (-684 > 0)) {
        var v15 uint32 = 2
        v16 := v6[v15]
        v17 := GenS1{F0: 49, F1: 1, F2: 9223372036854775807}
        _ = v15
        _ = v16
        _ = v17
    }
    switch "" {
    case "foo bar":
        var v18 float32 = float32(1.5)
        _ = v18
    case "\n\t":
        v1 = 18446744073709551615
    default:
        var v19 int = 1
        _ = v19
    }
    var v20 int = 127
    v3 += -2.25
    switch "\n\t" {
    case "hello":
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 float32 = v3
            for fi23 := 0; fi23 < 1; fi23++ {
                v24 := v2[42]
                v25 := []int{fi23}
                v4 = ("" + ((v5 + v5) + ("foo bar" + "\n\t")))
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi21
            _ = v22
        }
    case "\n\t":
        var v26 uint16 = 0
        _ = v26
    case "foo bar":
        if false {
            if true {
                switch (v20 / (((v20 / (((^(v20))) | 1))) | 1)) {
                case 1:
                    v3 += v3
                }
                var v27 uint32 = 800
                if true {
                    v3 -= v3
                } else {
                    v27 = v27
                    v28 := []int{(v20 * -9223372036854775808), (v20 << ((255) & 63)), v20}
                    switch "hello" {
                    case "":
                        var v29 float32 = 1e10
                        _ = v29
                    case "foo bar":
                        v1 -= 2
                    case "\n\t":
                        v3 -= v3
                    default:
                        v30 := GenS3{F0: 255, F1: 1, F2: 42}
                        _ = v30
                    }
                    _ = v28
                }
                _ = v27
            } else {
                var v31 string = "hello"
                v1 += v1
                _ = v31
            }
            v1 += v1
        } else {
            switch (v20 << ((921) & 63)) {
            case 76:
                if (673 > -933) {
                    _ = len(v6)
                    if true {
                        if true {
                            v32 := GenS1{F0: 65535, F1: 42, F2: -847}
                            v33 := GenS4{F0: 255}
                            _ = v32
                            _ = v33
                        } else {
                            switch v5 {
                            case "hello":
                                v1 += 18446744073709551615
                            case "a":
                                var v34 string = "hello"
                                _ = v34
                            default:
                                var v35 uint8 = 255
                                _ = v35
                            }
                        }
                    } else {
                        v3 -= v3
                    }
                    delete(v6, 2147483648)
                } else {
                    var v36 uint32 = 738
                    v5 = ""
                    if (!false) {
                        for fi37 := 0; fi37 < 7; fi37++ {
                            if (true && false) {
                                var v38 uint16 = uint16(v1)
                                v36 += (v36 ^ v36)
                                switch -2147483648 {
                                case 1:
                                    v39 := GenS3{F0: 255, F1: 100, F2: -1}
                                    _ = v39
                                case 32768:
                                    for fi40 := 0; fi40 < 9; fi40++ {
                                        v5 = v5
                                        _ = fi40
                                    }
                                }
                                _ = v38
                            }
                            var v41 int8 = -128
                            _ = fi37
                            _ = v41
                        }
                        var v42 int8 = 11
                        _ = v42
                    }
                    _ = v36
                }
            case -9223372036854775808:
                var v43 uint64 = v1
                _ = v43
            default:
                var v44 uint = 1
                _ = v44
            }
        }
    default:
        if true {
            for fi45 := 0; fi45 < 18; fi45++ {
                v3 -= v3
                v5 = ("hello" + v4)
                if (0.1 > 0.0) {
                    fmt.Println(v4, v3)
                }
                _ = fi45
            }
        } else {
            fmt.Println(v4, v20, v20)
            switch (v20 % ((35) | 1)) {
            case -2147483648:
                var v46 string = "hello"
                _ = v46
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
