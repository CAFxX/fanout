package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 int64; F2 int64 }

func main() {
    var v1 float32 = (((1.5 - ((((3.14159 + 3.14159) - float32((((-2.25 - 0.1) * 1e10) / (3.14159 / -2.25)))) + -2.25) * float32(0.0))) / 0.1) / 0.1)
    var v2 int = 32767
    var v3 float32 = (v1 * (((v1 + float32(3.14159)) / 1e10) + v1))
    fmt.Println(v3, v3)
    v4 := []int{v2}
    var v5 uint64 = 42
    v5 = v5
    var v6 int8 = 105
    v7 := GenS1{F0: (v6 ^ (v6 + v6))}
    v8 := GenS1{F0: 0}
    if false {
        if (2 != (v5 % (((v5 | 2)) | 1))) {
            _ = len(v4)
            switch v2 {
            case -155:
                fmt.Println(v1)
            case 711:
                v9 := GenS2{F0: (!true)}
                _ = v9
            }
        }
        var v10 int16 = -1
        _ = v10
    } else {
        for fi11 := 0; fi11 < 20; fi11++ {
            fmt.Println(v2, v6, v3)
            if (-2147483648 == v2) {
                v5 -= 616
                if false {
                    if true {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v12 := GenS3{F0: 185}
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        if false {
                            v14 := GenS4{F0: 2147483647, F1: (-1 != 1000)}
                            if (!(42 == 4294967295)) {
                                var v15 string = "a"
                                var v16 bool = true
                                _ = v15
                                _ = v16
                            }
                            _ = v14
                        }
                        if ((189 < 0) || false) {
                            if len(v4) > 0 {
                                v4[0] = 0
                            }
                        } else {
                            if ("\n\t" >= "\n\t") {
                                switch 9223372036854775807 {
                                case -9223372036854775808:
                                    var v17 bool = true
                                    _ = v17
                                case 0:
                                    if (false && true) {
                                        v3 = v1
                                    } else {
                                        _ = len(v4)
                                    }
                                default:
                                    var v18 int8 = (v6 - (v6 * 127))
                                    _ = v18
                                }
                                if (200 == -1) {
                                    _ = len(v4)
                                } else {
                                    if (32767 <= -518) {
                                        v19 := GenS3{F0: 0}
                                        _ = v19
                                    } else {
                                        fmt.Println(v5, v1, v6)
                                        var v20 bool = (-2147483648 > -584)
                                        _ = v20
                                    }
                                    fmt.Println(fi13, v2, v1)
                                }
                            }
                            v2 = (v2 ^ int(v6))
                        }
                        if true {
                            v21 := GenS3{F0: 437}
                            _ = v21
                        }
                        _ = len(v4)
                        _ = fi13
                    }
                    if (false && false) {
                        var v22 float64 = (1e10 + 0.0)
                        var v23 float32 = 1.5
                        _ = v22
                        _ = v23
                    } else {
                        var v24 int64 = -1
                        if false {
                            v2 = -1
                            v3 = (((v1 + float32(3.14159)) / v3) * v1)
                            fmt.Println(v5, v1)
                        } else {
                            v25 := GenS5{F0: "a"}
                            _ = v25
                        }
                        _ = v24
                    }
                    v26 := GenS4{F0: 42, F1: (2147483647 < 2)}
                    _ = v26
                }
                var v27 uint = 18446744073709551615
                _ = v27
            }
            _ = fi11
        }
        var v28 int32 = 42
        _ = v28
    }
    if (((float32(0.0) - v1) + float32(3.14159)) >= ((3.14159 - 0.1) / v3)) {
        var v29 int64 = -1
        if (42 > (^(v5))) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch (v2 % (((v2 ^ v2)) | 1)) {
            case 113:
                var v30 float32 = v1
                _ = v30
            default:
                var v31 int = 128
                _ = v31
            }
            for fi32 := 0; fi32 < 15; fi32++ {
                v33 := GenS6{F0: -111, F1: (v29 >> ((1) & 63)), F2: (^(v29))}
                _ = fi32
                _ = v33
            }
        } else {
            var v34 float32 = v3
            _ = v34
        }
        var v35 uint16 = 255
        _ = v29
        _ = v35
    }
    fmt.Println(v1, v1, v1)
    v2 -= (v2 ^ (v2 - v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
