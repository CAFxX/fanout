package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 bool }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint) bool {
    v7 := []int{326, 127, 476}
    p0 += p0
    fmt.Println(p0)
    fmt.Println(p0)
    if ((127 >= 0) && true) {
        _ = len(v7)
    }
    _ = p0
    _ = v7
    return true
}
func genFn1() uint64 {
    if (!false) {
        var v8 int64 = -946
        _ = v8
    }
    var v9 float64 = (1.5 - -2.25)
    var v10 bool = false
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 int32 = 2147483647
    v2 := []string{"a"}
    var v3 uint64 = 2
    var v4 int8 = -128
    var v5 bool = (true && false)
    v6 := make(map[int]bool)
    if v5 {
        fmt.Println(v4)
        fmt.Println(v5, v5, v4)
        fmt.Println(v4, v5, v3)
    } else {
        v11 := GenS1{F0: 1, F1: 42, F2: (32767 == 42)}
        if false {
            var v12 bool = v5
            _ = v12
        }
        _ = v11
    }
    var v13 int = -348
    var v14 uint8 = 255
    switch ("hello" + "foo bar") {
    case "":
        v15 := GenS1{F0: -9223372036854775808, F1: v14, F2: (32767 <= 42)}
        _ = v15
    case "\n\t":
        v4 -= v4
    default:
        switch 127 {
        case 4294967295:
            if (0.1 == 0.0) {
                switch (v13 + -326) {
                case 9223372036854775807:
                    if ((uint64(v14) >= v3) || (!(32767 <= v13))) {
                        v6[(v13 - (v13 << ((4294967295) & 63)))] = false
                        for fi16 := 0; fi16 < 16; fi16++ {
                            var v17 uint32 = 2147483648
                            var v18 int16 = -32768
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 7; fi19++ {
                            if (v4 <= (v4 - (^(v4)))) {
                                var v20 uint = 255
                                _ = v20
                            } else {
                                var v21 string = "hello"
                                _ = v21
                            }
                            v22 := GenS1{F0: -9223372036854775808, F1: (v14 / ((0) | 1)), F2: v5}
                            _ = fi19
                            _ = v22
                        }
                    } else {
                        fmt.Println(v1)
                        for fi23 := 0; fi23 < 5; fi23++ {
                            fmt.Println(fi23, v5, v1)
                            var v24 float32 = ((1.5 + -2.25) / 1.5)
                            switch "a" {
                            case "":
                                var v25 int8 = 127
                                _ = v25
                            case "hello":
                                var v26 int32 = v1
                                _ = v26
                            }
                            _ = fi23
                            _ = v24
                        }
                        if (v1 >= -1) {
                            v13 = v13
                        } else {
                            var v27 float32 = 3.14159
                            v28 := GenS2{F0: 417, F1: -2147483648}
                            _ = v27
                            _ = v28
                        }
                    }
                case 747:
                    fmt.Println(v5)
                }
                var v29 uint32 = 544
                _ = v29
            } else {
                if v5 {
                    var v30 float32 = float32(1e10)
                    _ = v30
                }
            }
        case 931:
            v1 += (v1 >> ((197) & 63))
        default:
            fmt.Println(v3, v1)
        }
    }
    v13 = 32768
    var v31 uint32 = 4294967295
    switch v13 {
    case 143:
        var v32 uint8 = v14
        _ = v32
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        if (v14 <= 1) {
            for fi34 := 0; fi34 < 20; fi34++ {
                v5 = true
                _ = fi34
            }
        } else {
            v1 = (v1 ^ (-(v1)))
        }
        switch ("" + "a") {
        case "":
            v35 := GenS3{F0: -693}
            _ = v35
        case "a":
            switch fi33 {
            case 9223372036854775807:
                if len(v2) > 0 {
                    v2[0] = (("" + "") + "foo bar")
                }
            case 0:
                _ = len(v2)
            }
        default:
            var v36 uint = 0
            _ = v36
        }
        switch (v13 * fi33) {
        case 709:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 2147483647:
            for fi37 := 0; fi37 < 6; fi37++ {
                if v5 {
                    v31 += (v31 | (v31 - 42))
                } else {
                    var v38 uint = 255
                    if true {
                        v39 := GenS2{F0: (v3 | (v3 << ((0) & 63))), F1: v1}
                        var v40 bool = ((^(fi33)) == -2147483648)
                        _ = v39
                        _ = v40
                    } else {
                        var v41 uint = (v38 | 947)
                        fmt.Println(v38, fi33, v4)
                        var v42 uint8 = v14
                        _ = v41
                        _ = v42
                    }
                    var v43 int = -1
                    _ = v38
                    _ = v43
                }
                fmt.Println(v3, v13)
                _ = len(v6)
                _ = fi37
            }
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
