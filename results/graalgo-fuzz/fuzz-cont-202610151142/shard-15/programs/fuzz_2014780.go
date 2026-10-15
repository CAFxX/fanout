package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint16 = 255
    var v2 float32 = ((0.1 + 0.0) + 0.1)
    var v3 string = "a"
    var v4 int64 = -1
    switch 2147483647 {
    case 32767:
        v5 := []int{128}
        _ = v5
    default:
        var v6 int8 = -128
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            v1 -= v1
            var v9 uint8 = 0
            _ = fi8
            _ = v9
        }
        var v10 int32 = 0
        fmt.Println(v4, fi7, v10)
        _ = fi7
        _ = v10
    }
    var v11 uint = 219
    v4 -= v4
    for fi12 := 0; fi12 < 19; fi12++ {
        if false {
            var v13 bool = (1 < v4)
            _ = v13
        }
        v14 := GenS1{F0: -854}
        if true {
            v3 = v3
            var v15 bool = (1 < 2147483647)
            if false {
                fmt.Println(v15, v15, v3)
                if ((v15 || (!true)) || v15) {
                    v2 += ((v2 + v2) / ((0.0 + v2) - (v2 - (v2 * 1e10))))
                }
                for fi16 := 0; fi16 < 16; fi16++ {
                    switch "foo bar" {
                    case "a":
                        fmt.Println(fi12, v11)
                    case "hello":
                        var v17 uint16 = (^(v1))
                        _ = v17
                    case "\n\t":
                        if v15 {
                            fmt.Println(v11)
                            if false {
                                if (v2 >= ((-2.25 + 1.5) + ((float32(0.0) + v2) * v2))) {
                                    v4 -= v4
                                    var v18 uint8 = 0
                                    _ = v18
                                } else {
                                    v19 := GenS1{F0: 2147483647}
                                    v20 := GenS1{F0: 972}
                                    _ = v19
                                    _ = v20
                                }
                            } else {
                                var v21 float32 = 3.14159
                                _ = v21
                            }
                            if v15 {
                                if true {
                                    var v22 uint64 = 0
                                    if false {
                                        v22 = (^(v22))
                                        v15 = (2 < 128)
                                        v23 := GenS2{F0: 255, F1: 849}
                                        _ = v23
                                    } else {
                                        var v24 string = "hello"
                                        if true {
                                            fmt.Println(v2, v24, v11)
                                            var v25 int16 = 32767
                                            _ = v25
                                        } else {
                                            if v15 {
                                                var v26 int64 = v4
                                                _ = v26
                                            }
                                            var v27 int8 = 0
                                            if (521 <= 2147483648) {
                                                fmt.Println(v24, v15)
                                            } else {
                                                var v28 int = fi16
                                                var v29 int32 = 626
                                                v30 := GenS2{F0: (v11 / ((945) | 1)), F1: (v22 ^ (v22 % ((42) | 1)))}
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                            }
                                            _ = v27
                                        }
                                        _ = v24
                                    }
                                    v2 -= (float32(3.14159) / ((float32((3.14159 - (3.14159 * 0.0))) - (v2 - v2)) * v2))
                                    _ = v22
                                } else {
                                    var v31 uint64 = 18446744073709551615
                                    _ = v31
                                }
                                if (907 >= 177) {
                                    v32 := []int{}
                                    var v33 int16 = 231
                                    v11 = (v11 & (v11 ^ (^(v11))))
                                    _ = v32
                                    _ = v33
                                }
                            }
                        }
                    }
                    _ = fi16
                }
            } else {
                if v15 {
                    if true {
                        v15 = v15
                        for fi34 := 0; fi34 < 3; fi34++ {
                            v35 := GenS2{F0: v11, F1: uint64(v1)}
                            if false {
                                var v36 uint = (v11 % (((v11 >> ((v11) & 63))) | 1))
                                v36 -= (v11 * 1)
                                if true {
                                    var v37 uint8 = 237
                                    v38 := GenS2{F0: (v36 | (v11 * (v11 / ((v36) | 1)))), F1: 735}
                                    v1 -= v1
                                    var v39 int16 = 1
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                } else {
                                    v40 := GenS2{F0: (v36 & v36), F1: 2}
                                    _ = v40
                                }
                                _ = v36
                            }
                            _ = fi34
                            _ = v35
                        }
                    }
                    v11 = v11
                } else {
                    for fi41 := 0; fi41 < 3; fi41++ {
                        fmt.Println(v3, v4)
                        _ = fi41
                    }
                    v11 = (v11 & 307)
                    v3 = v3
                }
                v42 := GenS3{F0: -32768}
                _ = v42
            }
            _ = v15
        }
        _ = fi12
        _ = v14
    }
    v1 += 32767
    if (!true) {
        if true {
            if (!(false && false)) {
                v43 := []int{}
                var v44 uint8 = 1
                switch ("a" + v3) {
                case "foo bar":
                    var v45 int8 = -128
                    _ = v45
                case "a":
                    if len(v43) >= 2 {
                        v43 = v43[1:len(v43)]
                    }
                default:
                    var v46 int8 = 100
                    _ = v46
                }
                _ = v43
                _ = v44
            } else {
                v3 = ""
            }
        } else {
            var v47 int16 = -526
            v4 = (v4 * 9223372036854775807)
            _ = v47
        }
        v48 := []int{-399}
        _ = v48
    }
    if ((true || true) && true) {
        v49 := GenS3{F0: 1000}
        _ = v49
    }
    v50 := []int{}
    var v51 int64 = (^(v4))
    var v52 uint16 = (v1 + (v1 + 0))
    fmt.Println(v52, v52)
    if (!(65535 > 2)) {
        fmt.Println(v1, v51, v1)
    } else {
        var v53 int = 32768
        if len(v50) > 0 {
            v50[0] = (-(v53))
        }
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v52
}
