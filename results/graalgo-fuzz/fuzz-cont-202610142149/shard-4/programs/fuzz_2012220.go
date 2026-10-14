package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint }
type GenS2 struct { F0 uint32; F1 int16; F2 int16 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int64; F1 string; F2 int64 }
type GenS5 struct { F0 int64; F1 int8; F2 int8 }
type GenS6 struct { F0 bool; F1 int16 }
func genFn0(p0 int16, p1 int64) int16 {
    p0 -= 1
    var v5 uint16 = 282
    p0 -= p0
    v6 := []int{9223372036854775807, -912}
    v7 := GenS1{F0: 468, F1: 173, F2: 915}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 int8, p1 uint) int32 {
    v8 := GenS2{F0: 780, F1: int16(p0), F2: 784}
    var v9 float32 = 0.0
    var v10 int16 = -906
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return -736
}

func main() {
    var v1 int64 = 0
    var v2 uint32 = 2147483648
    var v3 int64 = (-(v1))
    var v4 int16 = 1
    v3 -= 0
    var v11 uint16 = 28
    switch 593 {
    case -637:
        v12 := GenS1{F0: v2, F1: 127, F2: 0}
        _ = v12
    case 9223372036854775807:
        v13 := GenS3{F0: -2147483648, F1: 2}
        _ = v13
    case -2147483648:
        v14 := GenS3{F0: -628, F1: (v11 ^ (v11 >> ((65535) & 63)))}
        _ = v14
    default:
        if false {
            fmt.Println(v1, v11, v1)
            for fi15 := 0; fi15 < 7; fi15++ {
                for fi16 := 0; fi16 < 19; fi16++ {
                    if true {
                        fmt.Println(v2, v4, v2)
                        var v17 uint16 = 32767
                        if true {
                            fmt.Println(v3, v4, v3)
                        } else {
                            switch "" {
                            case "\n\t":
                                fmt.Println(fi15, v3, v11)
                            }
                            var v18 uint32 = v2
                            var v19 uint64 = 18446744073709551615
                            _ = v18
                            _ = v19
                        }
                        _ = v17
                    } else {
                        var v20 uint16 = uint16(v2)
                        if false {
                            var v21 uint = 4294967295
                            if true {
                                v22 := []int{fi16, (-(fi15)), 390, fi15}
                                var v23 uint16 = 0
                                _ = v22
                                _ = v23
                            }
                            _ = v21
                        }
                        _ = v20
                    }
                    var v24 float64 = (3.14159 * (1e10 * 0.0))
                    if (127 >= 128) {
                        var v25 uint = 92
                        v26 := GenS2{F0: v2, F1: (v4 >> ((uint(v11)) & 63)), F2: 1000}
                        _ = v25
                        _ = v26
                    }
                    _ = fi16
                    _ = v24
                }
                v27 := GenS3{F0: (fi15 & (fi15 >> ((65535) & 63))), F1: v11}
                var v28 bool = (!false)
                _ = fi15
                _ = v27
                _ = v28
            }
            v3 = (v1 % (((v3 % ((v1) | 1))) | 1))
        }
    }
    if ("" > "\n\t") {
        fmt.Println(v11)
        v4 -= 418
        for fi29 := 0; fi29 < 14; fi29++ {
            var v30 uint = 65535
            v31 := GenS4{F0: v1, F1: ("\n\t" + "a"), F2: 0}
            _ = fi29
            _ = v30
            _ = v31
        }
    } else {
        fmt.Println(v3, v2)
        switch (("" + ((("" + ((("" + "hello") + (("" + "a") + ("foo bar" + "foo bar"))) + ("a" + "\n\t"))) + "hello") + "hello")) + "hello") {
        case "a":
            v32 := []int{0, -161, 42, 127}
            _ = v32
        case "\n\t":
            for fi33 := 0; fi33 < 14; fi33++ {
                if (575 < genFn1(0, 2)) {
                    for fi34 := 0; fi34 < 12; fi34++ {
                        if true {
                            v35 := []int{(fi33 | 0)}
                            v36 := GenS1{F0: (v2 % (((v2 ^ v2)) | 1)), F1: 0, F2: 2}
                            _ = v35
                            _ = v36
                        }
                        v37 := []int{(fi34 | fi34)}
                        v37 = append(v37, int(v11))
                        _ = fi34
                        _ = v37
                    }
                    v38 := GenS5{F0: (v3 ^ (^(v1))), F1: 96, F2: -1}
                    _ = v38
                }
                _ = fi33
            }
        }
        v1 += v1
    }
    v3 = 157
    var v39 uint8 = 2
    if false {
        switch "a" {
        case "a":
            v40 := GenS6{F0: (-2.25 < 3.14159), F1: v4}
            _ = v40
        }
        fmt.Println(v1)
        v41 := []int{2147483647}
        _ = v41
    } else {
        v42 := []int{32767, 4294967295}
        switch 42 {
        case 479:
            v42 = append(v42, 32767)
        case 42:
            fmt.Println(v11, v4, v3)
        }
        _ = v42
    }
    v43 := GenS1{F0: 116, F1: (v39 * (^(v39))), F2: 712}
    v2 = (^(v2))
    var v44 int32 = -848
    v45 := []int{}
    v46 := GenS3{F0: -1, F1: v11}
    if (!false) {
        if (!((("a" > "\n\t") || (100 != 1)) && (581 < 9223372036854775807))) {
            if true {
                fmt.Println(v3, v39, v3)
                v11 = (v11 * v11)
                var v47 int = int(v4)
                _ = v47
            }
        }
        if (!(-329 == 4294967295)) {
            var v48 int64 = -9223372036854775808
            v48 += v48
            var v49 bool = false
            _ = v48
            _ = v49
        } else {
            v1 += (v1 % ((0) | 1))
            switch "foo bar" {
            case "foo bar":
                v50 := GenS4{F0: 42, F1: "a", F2: v1}
                _ = v50
            case "hello":
                var v51 uint16 = v11
                _ = v51
            case "a":
                var v52 bool = false
                _ = v52
            }
        }
    }
    var v53 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v53
}
