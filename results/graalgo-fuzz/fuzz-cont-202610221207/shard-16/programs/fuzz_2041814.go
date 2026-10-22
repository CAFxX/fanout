package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 string; F1 bool; F2 uint32 }
func genFn0(p0 int8) int64 {
    v4 := GenS1{F0: (false && (-900 == 0)), F1: 65535}
    v5 := GenS2{F0: 1}
    if false {
        var v6 string = "hello"
        _ = v6
    }
    _ = p0
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 int16 = 32767
    var v2 int16 = v1
    v3 := make(map[int]int)
    switch 129 {
    case 9223372036854775807:
        if (-355 > 2147483647) {
            var v7 int16 = 32767
            if false {
                var v8 string = "foo bar"
                v1 -= v2
                switch -9223372036854775808 {
                case 42:
                    var v9 int16 = (v7 + 32767)
                    _ = v9
                case 805:
                    var v10 uint32 = 42
                    _ = v10
                case 607:
                    switch -9223372036854775808 {
                    case 718:
                        var v11 bool = true
                        _ = v11
                    case -1000:
                        v1 -= (v1 & 42)
                    default:
                        fmt.Println(v7)
                    }
                default:
                    var v12 int32 = -628
                    _ = v12
                }
                _ = v8
            }
            _ = v7
        } else {
            for fi13 := 0; fi13 < 18; fi13++ {
                fmt.Println(v1, v2, v1)
                _ = fi13
            }
            v14 := GenS2{F0: 32767}
            fmt.Println(v1)
            _ = v14
        }
    }
    switch 839 {
    case 4294967295:
        if true {
            switch -2147483648 {
            case 2147483647:
                var v15 float32 = float32(3.14159)
                _ = v15
            }
        } else {
            fmt.Println(v1, v2)
            v2 += v1
            fmt.Println(v1, v1)
        }
    }
    v16 := v3[-1]
    v16 = v16
    switch "a" {
    case "\n\t":
        var v17 int32 = 2147483647
        _ = v17
    case "hello":
        if true {
            var v18 uint = 356
            v2 -= 1000
            _ = v18
        } else {
            for fi19 := 0; fi19 < 2; fi19++ {
                v16 = (v16 & (fi19 - (v16 ^ (v16 * fi19))))
                switch ("a" + "a") {
                case "\n\t":
                    switch v16 {
                    case 32768:
                        var v20 float64 = (0.0 + 1e10)
                        _ = v20
                    case 9223372036854775807:
                        v16 = (-(v16))
                    case -9223372036854775808:
                        if true {
                            var v21 int8 = -21
                            v22 := GenS3{F0: 490, F1: 0, F2: 1}
                            _ = v21
                            _ = v22
                        }
                    }
                case "":
                    fmt.Println(v1, v2, v1)
                default:
                    fmt.Println(v1, v16)
                }
                for fi23 := 0; fi23 < 10; fi23++ {
                    v2 += v1
                    if false {
                        var v24 float32 = ((3.14159 / 0.1) + (0.0 / 3.14159))
                        v24 += ((v24 * 0.0) / 0.1)
                        fmt.Println(v16, fi23, v24)
                        _ = v24
                    } else {
                        if false {
                            v2 += (v2 * (v1 % (((v1 * (v2 / (((v1 | v1)) | 1)))) | 1)))
                            fmt.Println(v16, v2, v2)
                        } else {
                            v25 := GenS4{F0: 42, F1: -1}
                            v16 = 9223372036854775807
                            if ((^(v16)) >= (v16 ^ fi19)) {
                                v2 -= v1
                                v26 := GenS5{F0: (fi19 ^ v16), F1: genFn0(38)}
                                v2 = (v1 * v1)
                                _ = v26
                            }
                            _ = v25
                        }
                        v2 = v1
                    }
                    _ = fi23
                }
                _ = fi19
            }
        }
    default:
        v16 = v16
    }
    switch (("" + "a") + "foo bar") {
    case "":
        fmt.Println(v16)
    default:
        if false {
            v27 := GenS6{F0: "hello", F1: ((2147483647 != 532) || false), F2: 901}
            if (1e10 != 1e10) {
                var v28 int16 = (^(v2))
                _ = v28
            } else {
                for fi29 := 0; fi29 < 4; fi29++ {
                    if (!true) {
                        var v30 uint = 298
                        _ = v30
                    }
                    if false {
                        if (!(42 < 127)) {
                            var v31 float64 = 0.1
                            for fi32 := 0; fi32 < 15; fi32++ {
                                fmt.Println(v31)
                                _ = fi32
                            }
                            _ = v31
                        }
                        fmt.Println(v16)
                        v33 := GenS2{F0: 861}
                        _ = v33
                    }
                    v2 = 32767
                    _ = fi29
                }
                var v34 int8 = 127
                _ = v34
            }
            _ = v27
        } else {
            v2 += (v2 & (v1 % ((v2) | 1)))
        }
    }
    v2 = (v2 | 1)
    fmt.Println(v16, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
}
