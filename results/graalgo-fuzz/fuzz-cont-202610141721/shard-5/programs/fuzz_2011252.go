package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint8; F2 string }

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]uint32)
    var v3 uint64 = 991
    var v4 int16 = (^(v1))
    v5 := make(map[string]int)
    v6 := GenS1{F0: 42, F1: -2147483648, F2: 42}
    var v7 bool = (!false)
    v2["a"] = 966
    var v8 int16 = (v1 | (^(v4)))
    v9 := v2[""]
    if (42 < 267) {
        v10 := []int{321, 618, 32768, -334}
        _ = v10
    } else {
        v11 := GenS2{F0: 42}
        v3 += v3
        fmt.Println(v9, v7, v3)
        _ = v11
    }
    var v12 string = "foo bar"
    fmt.Println(v9)
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := GenS1{F0: 2147483647, F1: -2147483648, F2: (fi13 + fi13)}
        _ = fi13
        _ = v14
    }
    v15 := []int{32767, -903}
    if len(v15) > 0 {
        v15[0] = 360
    }
    v16 := GenS3{F0: 1}
    switch v12 {
    case "foo bar":
        switch "hello" {
        case "hello":
            if len(v15) > 0 {
                v15[0] = 855
            }
        case "\n\t":
            v1 = v8
        default:
            v12 = "a"
        }
    default:
        for fi17 := 0; fi17 < 10; fi17++ {
            switch (fi17 ^ 654) {
            case 9223372036854775807:
                _ = len(v2)
            case 1:
                switch v12 {
                case "\n\t":
                    switch v12 {
                    case "a":
                        var v18 uint32 = v9
                        _ = v18
                    case "hello":
                        switch fi17 {
                        case 4294967295:
                            switch v12 {
                            case "foo bar":
                                for fi19 := 0; fi19 < 12; fi19++ {
                                    var v20 uint8 = 2
                                    var v21 float64 = 1e10
                                    var v22 string = v12
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                }
                            }
                        case -968:
                            var v23 string = "a"
                            _ = v23
                        default:
                            switch fi17 {
                            case -1:
                                fmt.Println(v1, v3, v7)
                            case 32768:
                                v7 = v7
                            default:
                                var v24 int64 = 477
                                _ = v24
                            }
                        }
                    case "foo bar":
                        for fi25 := 0; fi25 < 10; fi25++ {
                            if (!false) {
                                v26 := GenS3{F0: 255}
                                var v27 string = "hello"
                                switch fi17 {
                                case 128:
                                    _ = len(v5)
                                case 288:
                                    var v28 string = "\n\t"
                                    _ = v28
                                }
                                var v29 int16 = v4
                                _ = v26
                                _ = v27
                                _ = v29
                            } else {
                                v15 = append(v15, 0)
                                var v30 uint16 = 448
                                v31 := GenS2{F0: 186}
                                _ = v30
                                _ = v31
                            }
                            var v32 int64 = -647
                            v9 += (v9 ^ (v9 & (v9 % ((4294967295) | 1))))
                            _ = fi25
                            _ = v32
                        }
                    default:
                        switch fi17 {
                        case 128:
                            switch fi17 {
                            case -2147483648:
                                var v33 float64 = (-2.25 / 1.5)
                                _ = v33
                            case 2147483647:
                                fmt.Println(v12, v12)
                            case 127:
                                for fi34 := 0; fi34 < 1; fi34++ {
                                    if true {
                                        if true {
                                            var v35 uint64 = (v3 + 245)
                                            v36 := GenS2{F0: 192}
                                            v3 = 217
                                            _ = v35
                                            _ = v36
                                        }
                                        if v7 {
                                            var v37 uint8 = 127
                                            v9 += v9
                                            _ = v37
                                        }
                                    }
                                    _ = fi34
                                }
                            default:
                                v7 = (v7 || true)
                            }
                        case 32768:
                            var v38 int = 1
                            _ = v38
                        case 2147483647:
                            var v39 int8 = 100
                            _ = v39
                        }
                    }
                }
            default:
                if false {
                    for fi40 := 0; fi40 < 15; fi40++ {
                        var v41 bool = (1 <= 4294967295)
                        fmt.Println(v9)
                        _ = fi40
                        _ = v41
                    }
                    v42 := GenS4{F0: 27, F1: uint8(v3), F2: "foo bar"}
                    _ = v42
                }
            }
            v43 := GenS4{F0: 65535, F1: 1, F2: (v12 + "\n\t")}
            _ = fi17
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
