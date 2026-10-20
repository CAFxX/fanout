package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int; F2 int }
type GenS3 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 int32 }

func main() {
    var v1 int8 = -128
    var v2 uint32 = 65535
    v3 := make(map[uint32]bool)
    v4 := []string{"foo bar"}
    var v5 string = ("\n\t" + "a")
    var v6 uint = 260
    var v7 uint = (^(v6))
    delete(v3, (^(v2)))
    v8 := v3[v2]
    switch "hello" {
    case "":
        var v9 string = v5
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        _ = len(v3)
        _ = fi10
    }
    v11 := GenS1{F0: 957}
    if (v8 && v8) {
        v12 := GenS2{F0: (-(v1)), F1: 127, F2: -9223372036854775808}
        for fi13 := 0; fi13 < 10; fi13++ {
            if (19 == 0) {
                if len(v4) > 0 {
                    v4[0] = ("\n\t" + v5)
                }
                if v8 {
                    var v14 uint32 = 1
                    var v15 int32 = -458
                    for fi16 := 0; fi16 < 11; fi16++ {
                        fmt.Println(fi16)
                        if false {
                            v4 = append(v4, v5)
                            v17 := v3[v14]
                            fmt.Println(v14)
                            var v18 int64 = -1
                            _ = v17
                            _ = v18
                        }
                        var v19 string = v5
                        _ = fi16
                        _ = v19
                    }
                    _ = v14
                    _ = v15
                } else {
                    var v20 float64 = float64(1.5)
                    _ = v20
                }
                var v21 int16 = int16(fi13)
                _ = v21
            }
            if len(v4) > 0 {
                v4[0] = ("hello" + v5)
            }
            _ = fi13
        }
        switch (((((v5 + v5) + "foo bar") + v5) + v5) + ((v5 + "a") + (v5 + v5))) {
        case "\n\t":
            var v22 bool = ((!((42 != 147) && false)) || v8)
            _ = v22
        default:
            v23 := GenS3{F0: v1, F1: v2, F2: 2}
            _ = v23
        }
        _ = v12
    } else {
        for fi24 := 0; fi24 < 13; fi24++ {
            for fi25 := 0; fi25 < 5; fi25++ {
                v6 = uint(v2)
                var v26 int8 = -1
                var v27 float64 = 3.14159
                var v28 float32 = 0.0
                if ((v1 - v1) == (v1 - (v26 % ((v1) | 1)))) {
                    v1 = (v26 % (((v1 - v26)) | 1))
                } else {
                    v1 += v1
                    var v29 int64 = -9223372036854775808
                    var v30 string = ("a" + "hello")
                    _ = v29
                    _ = v30
                }
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
            _ = fi24
        }
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        v7 = (v6 / (((v7 >> ((18446744073709551615) & 63))) | 1))
        v32 := v3[4294967295]
        _ = fi31
        _ = v32
    }
    v33 := GenS4{F0: 1, F1: ((v2 % (((^(v2))) | 1)) > (v2 * v2))}
    if ((0.1 != 1.5) && v8) {
        if (v5 < v5) {
            for fi34 := 0; fi34 < 2; fi34++ {
                var v35 float32 = (float32((1e10 * float64((1e10 + 0.1)))) * 3.14159)
                switch v5 {
                case "hello":
                    v1 = int8(fi34)
                case "foo bar":
                    v3[v2] = v8
                default:
                    switch v5 {
                    case "hello":
                        v1 = v1
                    case "\n\t":
                        if (!false) {
                            if (0 >= -489) {
                                var v36 int64 = int64(v7)
                                _ = v36
                            } else {
                                for fi37 := 0; fi37 < 15; fi37++ {
                                    v38 := v3[(v2 >> ((v7) & 63))]
                                    v39 := GenS2{F0: (v1 & 0), F1: fi37, F2: (-(fi37))}
                                    _ = fi37
                                    _ = v38
                                    _ = v39
                                }
                            }
                            for fi40 := 0; fi40 < 7; fi40++ {
                                v6 = v7
                                _ = fi40
                            }
                        }
                    case "a":
                        v41 := GenS5{F0: 0, F1: -835}
                        _ = v41
                    }
                }
                for fi42 := 0; fi42 < 1; fi42++ {
                    if (42 != 42) {
                        switch (v5 + ((v5 + ((v5 + ((v5 + v5) + v5)) + "foo bar")) + v5)) {
                        case "hello":
                            var v43 int32 = -128
                            var v44 int64 = -9223372036854775808
                            v45 := v3[v2]
                            delete(v3, (v2 ^ (^(v2))))
                            _ = v43
                            _ = v44
                            _ = v45
                        case "":
                            switch "foo bar" {
                            case "":
                                v35 = 1e10
                            }
                        case "\n\t":
                            v6 -= 18446744073709551615
                        }
                    }
                    _ = fi42
                }
                _ = fi34
                _ = v35
            }
        }
        if (!false) {
            v6 = 255
        } else {
            v1 -= v1
            v46 := GenS6{F0: 0, F1: (v1 - (v1 - -28))}
            _ = v46
        }
        if true {
            v47 := GenS1{F0: 4294967295}
            v48 := GenS4{F0: -647, F1: (v8 && (0 < (v1 % (((v1 ^ (v1 + 1))) | 1))))}
            _ = v47
            _ = v48
        }
    } else {
        if v8 {
            var v49 string = ((v5 + v5) + v5)
            _ = v49
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch (v5 + "") {
        case "\n\t":
            v8 = v8
        case "a":
            if (!(1 == -2147483648)) {
                v2 += 506
            }
        case "foo bar":
            var v50 float64 = (float64(((-2.25 / 1.5) / 1e10)) + 0.1)
            _ = v50
        default:
            var v51 int8 = (^(v1))
            _ = v51
        }
    }
    _ = len(v4)
    var v52 float32 = (1e10 / 1e10)
    var v53 bool = (128 >= 42)
    switch (v5 + v5) {
    case "a":
        var v54 uint16 = 65535
        _ = v54
    case "foo bar":
        if v8 {
            v55 := GenS7{F0: -243}
            var v56 int8 = (-(v1))
            for fi57 := 0; fi57 < 1; fi57++ {
                for fi58 := 0; fi58 < 12; fi58++ {
                    var v59 uint = 17
                    var v60 uint32 = (v2 | (v2 % ((1) | 1)))
                    _ = fi58
                    _ = v59
                    _ = v60
                }
                v7 = (v7 + (v7 + v7))
                _ = fi57
            }
            _ = v55
            _ = v56
        }
    case "":
        if v53 {
            var v61 bool = v8
            _ = v61
        } else {
            if v8 {
                v53 = false
                v62 := GenS4{F0: 0, F1: false}
                _ = v62
            }
            var v63 uint8 = 0
            fmt.Println(v52)
            _ = v63
        }
    default:
        v8 = v8
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v7
    _ = v8
}
