package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int; F2 uint8 }
type GenS3 struct { F0 int64; F1 int8; F2 uint }

func main() {
    var v1 int16 = -3
    var v2 uint = 42
    var v3 int16 = -409
    var v4 int8 = -5
    var v5 uint32 = 0
    var v6 int = -9223372036854775808
    var v7 uint64 = 68
    v8 := []int{(-(v6))}
    v3 = v3
    if (!(v7 <= v7)) {
        for fi9 := 0; fi9 < 19; fi9++ {
            fmt.Println(v4, v7, v4)
            _ = fi9
        }
    } else {
        if false {
            v10 := GenS1{F0: 1}
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 uint64 = 4294967295
                if true {
                    if (565 <= (v3 << ((65535) & 63))) {
                        for fi13 := 0; fi13 < 6; fi13++ {
                            if true {
                                var v14 float32 = ((3.14159 - 0.1) * 0.0)
                                v8 = append(v8, (v6 ^ (v6 ^ (^(v6)))))
                                v8 = append(v8, fi11)
                                _ = v14
                            } else {
                                fmt.Println(fi13)
                            }
                            if false {
                                v15 := GenS2{F0: 0, F1: fi11, F2: 116}
                                _ = v15
                            } else {
                                var v16 int64 = -9223372036854775808
                                v8 = append(v8, -2147483648)
                                v7 = (v12 << (((^(v2))) & 63))
                                _ = v16
                            }
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            _ = fi13
                        }
                        for fi17 := 0; fi17 < 6; fi17++ {
                            v4 -= (v4 - (v4 * (-(v4))))
                            var v18 bool = false
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            fmt.Println(v4)
                            var v19 uint64 = 951
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        v6 -= fi11
                    } else {
                        var v20 int16 = v1
                        var v21 uint = 956
                        _ = v20
                        _ = v21
                    }
                } else {
                    if (-2.25 == 0.0) {
                        for fi22 := 0; fi22 < 20; fi22++ {
                            var v23 float64 = (0.1 * float64(-2.25))
                            fmt.Println(fi11, v5)
                            var v24 bool = ((float64(1.5) < (v23 * 1e10)) && (true && ((-(v6)) != -2147483648)))
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    } else {
                        for fi25 := 0; fi25 < 20; fi25++ {
                            var v26 string = "foo bar"
                            _ = fi25
                            _ = v26
                        }
                        v8 = append(v8, (v6 & (fi11 % ((32767) | 1))))
                    }
                    if (!true) {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    }
                }
                v3 = v1
                _ = fi11
                _ = v12
            }
            if false {
                var v27 int = -996
                var v28 uint = (v2 | (v2 >> ((207) & 63)))
                for fi29 := 0; fi29 < 13; fi29++ {
                    v30 := GenS1{F0: 2147483647}
                    if true {
                        if false {
                            switch "\n\t" {
                            case "":
                                v6 = -446
                            case "hello":
                                var v31 int32 = -2147483648
                                _ = v31
                            case "a":
                                v1 += v3
                            }
                            v27 = v27
                            v32 := GenS3{F0: 1, F1: v4, F2: v28}
                            _ = v32
                        } else {
                            v4 = 100
                        }
                        var v33 float64 = 1e10
                        _ = v33
                    }
                    _ = fi29
                    _ = v30
                }
                _ = v27
                _ = v28
            } else {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                switch "hello" {
                case "hello":
                    if len(v8) > 0 {
                        v8[0] = (v6 & 4294967295)
                    }
                case "":
                    var v34 int16 = (v3 % ((v1) | 1))
                    _ = v34
                case "foo bar":
                    if (!(-128 > (v4 & int8(v5)))) {
                        v6 = (v6 / ((v6) | 1))
                    } else {
                        v8 = append(v8, (v6 + int(v2)))
                        switch ("" + ("\n\t" + "a")) {
                        case "":
                            v35 := GenS3{F0: -1, F1: (v4 ^ 1), F2: (v2 / ((1) | 1))}
                            _ = v35
                        case "\n\t":
                            v3 = v3
                        case "a":
                            for fi36 := 0; fi36 < 3; fi36++ {
                                v8 = append(v8, (fi36 & v6))
                                var v37 int16 = -32768
                                _ = fi36
                                _ = v37
                            }
                        default:
                            v1 += v3
                        }
                    }
                default:
                    v7 = (v7 & v7)
                }
                var v38 uint64 = 2
                _ = v38
            }
            _ = v10
        }
    }
    v1 = (v3 + v1)
    var v39 int8 = int8(v7)
    if false {
        var v40 bool = true
        fmt.Println(v40, v39, v3)
        _ = v40
    } else {
        fmt.Println(v1)
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if len(v8) > 0 {
        v8[0] = v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
