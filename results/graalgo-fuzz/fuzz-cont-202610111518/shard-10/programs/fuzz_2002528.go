package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 bool; F1 uint64; F2 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 int32; F2 int8 }
func genFn0() uint16 {
    v7 := []int{42, 1}
    _ = len(v7)
    if true {
        var v8 float32 = float32(1.5)
        var v9 int8 = 98
        fmt.Println(v8)
        _ = v8
        _ = v9
    }
    _ = v7
    return 914
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    v4 := []string{}
    var v5 int = 2147483647
    v6 := make(map[string]bool)
    var v10 uint8 = 111
    fmt.Println(v5, v10, v10)
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 string = "foo bar"
        v1 -= 1
        var v13 uint64 = 18446744073709551615
        _ = fi11
        _ = v12
        _ = v13
    }
    if (!(v5 == v5)) {
        v14 := GenS1{F0: 1, F1: -16}
        switch v5 {
        case -1:
            fmt.Println(v1, v5, v1)
        case -683:
            switch "" {
            case "\n\t":
                var v15 int64 = 9223372036854775807
                _ = v15
            case "a":
                v16 := v2[(v5 & v5)]
                _ = v16
            case "hello":
                if len(v4) > 0 {
                    v4[0] = "a"
                }
            default:
                for fi17 := 0; fi17 < 16; fi17++ {
                    _ = len(v2)
                    if (!true) {
                        if ("" >= "") {
                            if false {
                                for fi18 := 0; fi18 < 17; fi18++ {
                                    var v19 int16 = 0
                                    v10 = 253
                                    var v20 string = ((("" + "\n\t") + ((("\n\t" + "a") + (("a" + "hello") + "a")) + "foo bar")) + "\n\t")
                                    _ = fi18
                                    _ = v19
                                    _ = v20
                                }
                            }
                            for fi21 := 0; fi21 < 3; fi21++ {
                                var v22 int64 = 9223372036854775807
                                v1 = v1
                                if len(v4) > 0 {
                                    v4[0] = "foo bar"
                                }
                                var v23 uint16 = 255
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        } else {
                            fmt.Println(v1)
                            _ = len(v2)
                        }
                    }
                    _ = fi17
                }
            }
        case 128:
            switch 1 {
            case -686:
                for fi24 := 0; fi24 < 7; fi24++ {
                    for fi25 := 0; fi25 < 1; fi25++ {
                        var v26 bool = true
                        v10 -= v10
                        _ = fi25
                        _ = v26
                    }
                    fmt.Println(v1, v1, v10)
                    v4 = append(v4, "hello")
                    _ = fi24
                }
            }
        default:
            if true {
                if ((884 > 25) && true) {
                    var v27 int = v5
                    if true {
                        for fi28 := 0; fi28 < 2; fi28++ {
                            for fi29 := 0; fi29 < 7; fi29++ {
                                v30 := GenS1{F0: 9223372036854775807, F1: (v27 / (((v5 ^ (v5 ^ (v5 + (fi29 ^ (v5 % (((v5 >> ((703) & 63))) | 1))))))) | 1))}
                                _ = len(v6)
                                fmt.Println(v10, fi29, v10)
                                _ = fi29
                                _ = v30
                            }
                            v31 := GenS1{F0: 932, F1: v27}
                            v32 := GenS2{F0: false, F1: 4294967295, F2: (v1 - v1)}
                            _ = fi28
                            _ = v31
                            _ = v32
                        }
                        var v33 int64 = 937
                        switch "a" {
                        case "":
                            if true {
                                var v34 int32 = (v1 << ((256) & 63))
                                v27 = v5
                                v33 = 1
                                _ = v34
                            }
                        default:
                            v10 -= 42
                        }
                        _ = v33
                    }
                    switch "" {
                    case "hello":
                        v35 := GenS3{F0: false}
                        _ = v35
                    case "foo bar":
                        var v36 bool = true
                        _ = v36
                    }
                    _ = v27
                }
                v10 -= v10
            } else {
                switch ("\n\t" + "hello") {
                case "a":
                    v10 = 255
                case "hello":
                    var v37 int64 = 1
                    _ = v37
                case "foo bar":
                    fmt.Println(v5)
                }
            }
        }
        var v38 string = "\n\t"
        _ = v14
        _ = v38
    } else {
        var v39 float64 = 0.1
        v39 -= -2.25
        _ = v39
    }
    v5 += 9223372036854775807
    var v40 string = "foo bar"
    if false {
        v10 = (v10 | v10)
        var v41 bool = true
        _ = v41
    }
    v42 := GenS4{F0: v40, F1: -2147483648, F2: 100}
    v10 += 1
    var v43 float64 = 0.0
    var v44 int32 = v1
    var v45 int32 = -505
    v5 += v5
    v40 = (v40 + (((("a" + "a") + ("foo bar" + v40)) + ("" + v40)) + (v40 + v40)))
    v3[128] = v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
