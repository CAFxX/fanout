package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint16 }

func main() {
    var v1 uint8 = 127
    var v2 int32 = 42
    var v3 uint = 394
    var v4 uint = (v3 | v3)
    var v5 uint16 = 1
    var v6 bool = false
    v6 = ((0.0 / 0.1) <= -2.25)
    v7 := GenS1{F0: (-32768 <= -1)}
    var v8 uint32 = 4294967295
    v5 = v5
    v9 := []int{-9223372036854775808, 2147483647, -322, 127}
    v10 := GenS2{F0: (v8 / ((v8) | 1))}
    v3 = (v3 & (v4 * 42))
    var v11 string = ("" + "")
    if v6 {
        if (0 == 128) {
            for fi12 := 0; fi12 < 9; fi12++ {
                if v6 {
                    switch 2147483647 {
                    case 0:
                        var v13 uint16 = 65535
                        _ = v13
                    default:
                        for fi14 := 0; fi14 < 20; fi14++ {
                            v8 += (v8 * uint32(v5))
                            v2 = (-(v2))
                            switch (fi14 - (fi12 | 42)) {
                            case -711:
                                var v15 int64 = 600
                                _ = v15
                            case 9223372036854775807:
                                var v16 int32 = (^(v2))
                                _ = v16
                            default:
                                var v17 uint64 = 42
                                _ = v17
                            }
                            _ = fi14
                        }
                    }
                    switch v11 {
                    case "\n\t":
                        _ = len(v9)
                    case "hello":
                        v8 -= v8
                    case "":
                        if v6 {
                            for fi18 := 0; fi18 < 17; fi18++ {
                                if (v8 < v8) {
                                    fmt.Println(v2, v5, v5)
                                } else {
                                    if (v6 || (!v6)) {
                                        if false {
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                        }
                                        v4 = 77
                                    } else {
                                        var v19 int8 = -1
                                        _ = v19
                                    }
                                    var v20 int16 = 32767
                                    v11 = ""
                                    _ = v20
                                }
                                var v21 int16 = -32768
                                _ = fi18
                                _ = v21
                            }
                            v5 = (v5 + (v5 - (^(v5))))
                        }
                    default:
                        var v22 uint64 = 18446744073709551615
                        _ = v22
                    }
                    var v23 int = 281
                    _ = v23
                } else {
                    var v24 uint64 = 2
                    var v25 bool = true
                    if ((^(v24)) <= (^(v24))) {
                        fmt.Println(fi12)
                    } else {
                        if (v6 && (1.5 == 0.0)) {
                            var v26 int32 = v2
                            v27 := GenS2{F0: v8}
                            _ = v26
                            _ = v27
                        }
                        fmt.Println(v24)
                        var v28 uint = v3
                        _ = v28
                    }
                    _ = v24
                    _ = v25
                }
                _ = fi12
            }
        }
        var v29 int64 = 772
        _ = v29
    } else {
        fmt.Println(v11, v8, v11)
        v30 := GenS1{F0: (v2 > 42)}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 uint8 = 2
        for fi33 := 0; fi33 < 3; fi33++ {
            v34 := GenS1{F0: ((!v6) || ((!false) || false))}
            var v35 uint32 = uint32(v5)
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi31
        _ = v32
    }
    if v6 {
        v1 += (^(v1))
        if (((-32 > -128) && true) || (((1 == int(v8)) || (false || (int8(v2) != -128))) || (2 == v8))) {
            switch v11 {
            case "\n\t":
                v1 = v1
            case "hello":
                v8 = 4294967295
            case "foo bar":
                v4 -= (v4 * 931)
            default:
                for fi36 := 0; fi36 < 13; fi36++ {
                    v6 = true
                    v37 := GenS1{F0: (1.5 == float32(float64((float32(3.14159) + float32(1e10)))))}
                    for fi38 := 0; fi38 < 14; fi38++ {
                        var v39 int64 = 9223372036854775807
                        _ = fi38
                        _ = v39
                    }
                    _ = fi36
                    _ = v37
                }
            }
            _ = len(v9)
            for fi40 := 0; fi40 < 1; fi40++ {
                var v41 bool = ((v5 ^ v5) >= (v5 % ((121) | 1)))
                var v42 int16 = 32767
                var v43 uint32 = v8
                _ = fi40
                _ = v41
                _ = v42
                _ = v43
            }
        }
        if v6 {
            v4 -= (^(v3))
        } else {
            switch 0 {
            case 128:
                fmt.Println(v3)
            }
            switch v11 {
            case "":
                var v44 int32 = (v2 / (((v2 - v2)) | 1))
                _ = v44
            }
            v45 := GenS1{F0: false}
            _ = v45
        }
    } else {
        if ((v6 && v6) && (!(!(true || (255 > (v1 + (v1 * (v1 / ((255) | 1))))))))) {
            var v46 int64 = 9223372036854775807
            var v47 int16 = -922
            v46 = v46
            _ = v46
            _ = v47
        }
        fmt.Println(v6)
    }
    v11 = "foo bar"
    if v6 {
        fmt.Println(v11, v1, v2)
    } else {
        v48 := GenS3{F0: uint64(v5), F1: (v2 & (v2 - v2)), F2: (v5 % (((v5 + (v5 << ((v3) & 63)))) | 1))}
        _ = v48
    }
    if false {
        v4 += (v4 * (v4 / ((65535) | 1)))
    }
    var v49 uint16 = v5
    var v50 uint = v4
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
