package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = -783
    var v3 int8 = 0
    var v4 float64 = (3.14159 - 1.5)
    v3 += 127
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: 2, F1: 2147483648, F2: 1}
        fmt.Println(v2)
        _ = fi5
        _ = v6
    }
    var v7 int8 = (v3 & v3)
    var v8 uint16 = 2
    var v9 uint = 172
    var v10 uint64 = 2
    if (!true) {
        v11 := []int{128, 0, 9223372036854775807}
        switch "hello" {
        case "\n\t":
            var v12 int = 407
            _ = v12
        case "a":
            switch 1 {
            case -9223372036854775808:
                switch 4294967295 {
                case 0:
                    for fi13 := 0; fi13 < 5; fi13++ {
                        v9 += v9
                        if true {
                            v7 -= v7
                        } else {
                            for fi14 := 0; fi14 < 16; fi14++ {
                                switch fi14 {
                                case -9223372036854775808:
                                    if false {
                                        v10 += v10
                                    }
                                }
                                var v15 float32 = 1.5
                                _ = fi14
                                _ = v15
                            }
                        }
                        var v16 uint64 = v10
                        _ = fi13
                        _ = v16
                    }
                case 1:
                    switch ("\n\t" + ("" + "\n\t")) {
                    case "":
                        _ = len(v11)
                    case "foo bar":
                        var v17 uint32 = uint32(v9)
                        _ = v17
                    default:
                        v2 -= (v2 - v2)
                    }
                case -1:
                    if true {
                        v7 -= 0
                        v8 -= uint16(v9)
                    } else {
                        fmt.Println(v10)
                    }
                }
            case 0:
                switch "foo bar" {
                case "hello":
                    var v18 uint8 = 127
                    _ = v18
                case "\n\t":
                    if (true && (!(-746 < int64(v7)))) {
                        _ = len(v11)
                        if (1.5 < 1e10) {
                            var v19 bool = (1e10 == float32(v4))
                            var v20 int16 = v2
                            var v21 uint64 = 42
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    }
                case "a":
                    v7 -= (v7 / ((100) | 1))
                default:
                    v8 = 255
                }
            case 32768:
                var v22 uint32 = 4294967295
                _ = v22
            }
        default:
            for fi23 := 0; fi23 < 3; fi23++ {
                for fi24 := 0; fi24 < 11; fi24++ {
                    var v25 float32 = (1e10 - -2.25)
                    if (("a" < "foo bar") || false) {
                        fmt.Println(v3)
                    } else {
                        fmt.Println(v8, v3, v2)
                        var v26 uint16 = (v8 & v8)
                        fmt.Println(v26, v25)
                        switch ("\n\t" + ((("" + ("hello" + "foo bar")) + "") + "\n\t")) {
                        case "\n\t":
                            var v27 uint32 = 4294967295
                            _ = v27
                        case "a":
                            v28 := GenS1{F0: 0, F1: 722, F2: (v9 + (v9 * (v9 - (v9 + (v9 & (v9 % (((v9 * 65535)) | 1)))))))}
                            _ = v28
                        case "hello":
                            v25 -= v25
                            var v29 int = fi24
                            v30 := GenS1{F0: 2, F1: 0, F2: v9}
                            _ = v29
                            _ = v30
                        default:
                            var v31 int16 = int16(v10)
                            _ = v31
                        }
                        v7 -= 100
                        _ = v26
                    }
                    _ = fi24
                    _ = v25
                }
                _ = fi23
            }
        }
        _ = v11
    }
    var v32 int32 = -1
    var v33 uint16 = 42
    var v34 uint32 = 65535
    if false {
        _ = len(v1)
        if (!false) {
            var v35 uint8 = 196
            switch "hello" {
            case "":
                var v36 int64 = 1
                _ = v36
            case "hello":
                var v37 int32 = v32
                _ = v37
            default:
                v35 = (v35 ^ (v35 << ((v9) & 63)))
            }
            _ = v35
        } else {
            v34 = (^(v34))
        }
        for fi38 := 0; fi38 < 12; fi38++ {
            fmt.Println(fi38, v4)
            var v39 int32 = (v32 << (((v9 + (v9 / ((1) | 1)))) & 63))
            fmt.Println(fi38, v39, v9)
            _ = fi38
            _ = v39
        }
    } else {
        for fi40 := 0; fi40 < 15; fi40++ {
            v10 += uint64(v9)
            fmt.Println(v2)
            _ = fi40
        }
        v41 := v1[864]
        var v42 string = ("\n\t" + "a")
        _ = v41
        _ = v42
    }
    var v43 int32 = (v32 * int32(v34))
    v44 := GenS1{F0: (v34 + v34), F1: v34, F2: v9}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v43
    _ = v44
    _ = v7
    _ = v8
    _ = v9
}
