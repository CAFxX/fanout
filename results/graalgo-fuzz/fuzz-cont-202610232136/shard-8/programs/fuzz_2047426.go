package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int32; F2 uint8 }

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]bool)
    var v3 uint32 = 1
    var v4 uint32 = v3
    var v5 string = "a"
    v6 := []int{2147483647, 847, 68}
    switch -387 {
    case -2147483648:
        if false {
            switch -9223372036854775808 {
            case 9223372036854775807:
                v7 := GenS1{F0: -1, F1: (42 == 9223372036854775807)}
                _ = v7
            case 127:
                var v8 float64 = (1.5 * ((-2.25 + 1.5) * 0.1))
                _ = v8
            case 32768:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            default:
                switch "" {
                case "hello":
                    fmt.Println(v1, v1, v4)
                default:
                    v4 = (v3 ^ (v3 | v3))
                }
            }
            if false {
                var v9 uint = 0
                fmt.Println(v9)
                v2[-9223372036854775808] = false
                _ = v9
            }
            for fi10 := 0; fi10 < 12; fi10++ {
                if ((!(-1 == 249)) && false) {
                    v11 := GenS1{F0: 3, F1: false}
                    switch -662 {
                    case 32767:
                        if true {
                            v6 = append(v6, fi10)
                            var v12 string = v5
                            _ = v12
                        }
                    case -1:
                        if (-32768 >= 1) {
                            switch ("a" + (v5 + ("hello" + (v5 + ((v5 + (((v5 + (("\n\t" + v5) + ((v5 + "") + "hello"))) + ((v5 + v5) + "foo bar")) + v5)) + ("hello" + ("\n\t" + v5))))))) {
                            case "foo bar":
                                fmt.Println(v3, v3, v3)
                            case "\n\t":
                                for fi13 := 0; fi13 < 2; fi13++ {
                                    v14 := GenS1{F0: 42, F1: true}
                                    _ = fi13
                                    _ = v14
                                }
                            case "a":
                                v15 := GenS1{F0: 100, F1: ((true && (false || (v1 != v1))) && true)}
                                _ = v15
                            default:
                                switch v5 {
                                case "foo bar":
                                    v1 -= 2
                                case "\n\t":
                                    for fi16 := 0; fi16 < 8; fi16++ {
                                        fmt.Println(v4)
                                        _ = fi16
                                    }
                                case "":
                                    v17 := GenS1{F0: 42, F1: (true || false)}
                                    _ = v17
                                }
                            }
                            var v18 int = fi10
                            _ = v18
                        }
                    case 729:
                        if (246 != -2147483648) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            for fi19 := 0; fi19 < 13; fi19++ {
                                var v20 uint32 = v3
                                _ = fi19
                                _ = v20
                            }
                            fmt.Println(v3, v1, fi10)
                        } else {
                            for fi21 := 0; fi21 < 19; fi21++ {
                                fmt.Println(fi21)
                                _ = fi21
                            }
                            var v22 uint64 = 18446744073709551615
                            if false {
                                var v23 int8 = 100
                                v24 := GenS1{F0: (v23 % ((42) | 1)), F1: (32767 <= fi10)}
                                for fi25 := 0; fi25 < 20; fi25++ {
                                    if (2147483647 == 2147483647) {
                                        if len(v6) > 0 {
                                            v6[0] = fi10
                                        }
                                    }
                                    _ = fi25
                                }
                                _ = v23
                                _ = v24
                            } else {
                                for fi26 := 0; fi26 < 8; fi26++ {
                                    fmt.Println(v1, v22, fi10)
                                    var v27 int64 = -1
                                    v4 -= v3
                                    _ = fi26
                                    _ = v27
                                }
                            }
                            _ = v22
                        }
                    }
                    _ = v11
                }
                v3 = v3
                _ = fi10
            }
        } else {
            if (((527 > (v3 / ((487) | 1))) && (((32767 <= 153) && true) && true)) || true) {
                v4 = v4
            }
            switch v5 {
            case "foo bar":
                switch -415 {
                case 344:
                    var v28 int = -1
                    _ = v28
                case 4294967295:
                    v5 = v5
                case 921:
                    for fi29 := 0; fi29 < 9; fi29++ {
                        if (0.0 == 1e10) {
                            v5 = "foo bar"
                        }
                        fmt.Println(v1)
                        _ = fi29
                    }
                default:
                    var v30 float64 = 1e10
                    _ = v30
                }
            case "hello":
                var v31 int16 = -1
                _ = v31
            case "\n\t":
                var v32 bool = true
                _ = v32
            }
            if (float32(1e10) < 0.0) {
                for fi33 := 0; fi33 < 3; fi33++ {
                    var v34 uint8 = v1
                    v4 = v3
                    _ = fi33
                    _ = v34
                }
            }
        }
    case 127:
        v35 := GenS2{F0: 1}
        _ = v35
    case 527:
        fmt.Println(v1, v1, v1)
    }
    v6 = append(v6, 347)
    switch "\n\t" {
    case "a":
        v5 = ("foo bar" + v5)
    case "\n\t":
        v36 := GenS2{F0: 816}
        _ = v36
    }
    v1 = (v1 ^ (v1 | (v1 * v1)))
    for fi37 := 0; fi37 < 7; fi37++ {
        if (v4 > v4) {
            v38 := GenS3{F0: v5, F1: 2147483647, F2: v1}
            _ = v38
        }
        v5 = (v5 + v5)
        if len(v6) > 0 {
            v6[0] = 32767
        }
        _ = fi37
    }
    var v39 uint8 = (v1 * (v1 | (^(v1))))
    switch ("" + "foo bar") {
    case "a":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "\n\t":
        fmt.Println(v5, v4, v39)
    default:
        fmt.Println(v1)
    }
    v4 += (^(v3))
    var v40 uint32 = 177
    var v41 uint64 = uint64(v3)
    var v42 int64 = -284
    var v43 int16 = 42
    v3 += v4
    v42 = v42
    _ = len(v6)
    var v44 float64 = 0.1
    v5 = ("a" + v5)
    var v45 string = v5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
