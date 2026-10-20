package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int32; F1 uint16; F2 int }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{}
    var v3 float64 = 1e10
    var v4 float64 = v3
    var v5 int = -2147483648
    var v6 bool = (1 < 1)
    var v7 int16 = 0
    var v8 uint8 = 128
    var v9 int32 = -2147483648
    v4 = 1e10
    if v6 {
        if v6 {
            var v10 bool = (654 < -1)
            var v11 int64 = 874
            v4 = v4
            _ = v10
            _ = v11
        } else {
            v2 = append(v2, ("hello" + ("foo bar" + ("" + "hello"))))
        }
        v3 += 1e10
        switch 128 {
        case 836:
            v8 = v8
        case 128:
            var v12 int8 = -10
            _ = v12
        default:
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 uint = 4294967295
                _ = fi13
                _ = v14
            }
        }
    } else {
        var v15 uint = 255
        if v6 {
            var v16 float32 = (float32(float64((1.5 - 3.14159))) - (3.14159 * -2.25))
            for fi17 := 0; fi17 < 4; fi17++ {
                var v18 uint64 = 18446744073709551615
                if (v6 || v6) {
                    v19 := v1[2147483648]
                    switch (v19 + ("" + v19)) {
                    case "a":
                        var v20 uint = (v15 + uint(v8))
                        _ = v20
                    }
                    _ = v19
                } else {
                    var v21 uint16 = 253
                    _ = v21
                }
                v7 = 1
                _ = fi17
                _ = v18
            }
            fmt.Println(v6, v4, v8)
            _ = v16
        }
        _ = v15
    }
    fmt.Println(v3)
    for fi22 := 0; fi22 < 11; fi22++ {
        v6 = ("\n\t" >= "")
        if v6 {
            for fi23 := 0; fi23 < 15; fi23++ {
                v7 = (v7 * v7)
                _ = fi23
            }
            for fi24 := 0; fi24 < 8; fi24++ {
                if v6 {
                    v25 := v1[1]
                    if ((v25 + "hello") > ((v25 + ("a" + (v25 + v25))) + "a")) {
                        if (v9 >= (v9 % ((v9) | 1))) {
                            switch v25 {
                            case "\n\t":
                                v9 += (v9 + (v9 | (v9 - (-(v9)))))
                            }
                            if len(v2) > 0 {
                                v2[0] = ((v25 + v25) + (v25 + v25))
                            }
                            var v26 int8 = -1
                            _ = v26
                        }
                        switch ("a" + v25) {
                        case "\n\t":
                            v27 := GenS1{F0: -123, F1: 0}
                            _ = v27
                        }
                    } else {
                        v28 := GenS1{F0: 42, F1: -561}
                        _ = v28
                    }
                    _ = v25
                } else {
                    v5 = (fi24 >> ((18446744073709551615) & 63))
                    _ = len(v2)
                    v4 += float64(0.1)
                }
                var v29 uint = 2
                _ = fi24
                _ = v29
            }
        } else {
            var v30 uint16 = 255
            for fi31 := 0; fi31 < 9; fi31++ {
                v7 -= (v7 % (((v7 * (v7 << ((4294967295) & 63)))) | 1))
                switch "hello" {
                case "hello":
                    switch "hello" {
                    case "\n\t":
                        if (v30 >= 1) {
                            v9 -= (v9 / ((461) | 1))
                            v32 := GenS2{F0: int32(v7), F1: (v30 ^ (v30 & (^(v30)))), F2: fi31}
                            _ = v32
                        }
                    case "":
                        v33 := GenS1{F0: -7, F1: -9223372036854775808}
                        if v6 {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if v6 {
                                v8 = (v8 << ((432) & 63))
                                v34 := GenS2{F0: (v9 % (((v9 << ((18446744073709551615) & 63))) | 1)), F1: 32767, F2: 563}
                                var v35 uint16 = 32767
                                _ = v34
                                _ = v35
                            } else {
                                if (0 != 1) {
                                    v36 := GenS3{F0: 279}
                                    var v37 uint16 = v30
                                    _ = v36
                                    _ = v37
                                } else {
                                    var v38 int16 = (v7 - (v7 >> ((255) & 63)))
                                    _ = len(v2)
                                    _ = v38
                                }
                                v1[4294967295] = "foo bar"
                                v8 = 255
                            }
                            var v39 int8 = -12
                            _ = v39
                        } else {
                            var v40 int32 = 475
                            var v41 int16 = (v7 >> ((42) & 63))
                            fmt.Println(v30, v5, v3)
                            _ = v40
                            _ = v41
                        }
                        _ = v33
                    case "foo bar":
                        v8 -= v8
                    default:
                        var v42 uint64 = 1
                        _ = v42
                    }
                default:
                    v5 -= v5
                    fmt.Println(v8, v9)
                }
                switch (fi22 % ((fi31) | 1)) {
                case 32767:
                    fmt.Println(v30, fi22)
                case 1:
                    fmt.Println(v6, v3)
                default:
                    var v43 float32 = 0.1
                    _ = v43
                }
                _ = fi31
            }
            v6 = (v30 <= v30)
            _ = v30
        }
        v4 -= float64(((-2.25 * 0.0) - float32((v3 + ((3.14159 * float64(0.0)) * v4)))))
        _ = fi22
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        v5 += 42
        delete(v1, 2)
        _ = fi44
    }
    if v6 {
        fmt.Println(v7, v3)
        delete(v1, 125)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
