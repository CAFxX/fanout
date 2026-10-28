package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int }
type GenS5 struct { F0 int; F1 int8; F2 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 int64; F1 int }
func genFn0() int32 {
    var v7 string = ("" + "foo bar")
    v8 := []int{}
    for fi9 := 0; fi9 < 6; fi9++ {
        v7 = v7
        var v10 int8 = 0
        v11 := GenS1{F0: 42, F1: 32767, F2: 100}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v7
    _ = v8
    return 252
}

func main() {
    var v1 int16 = 860
    var v2 bool = false
    var v3 uint16 = 65535
    var v4 int32 = 42
    var v5 uint64 = 1
    var v6 uint64 = v5
    v12 := GenS2{F0: 4294967295, F1: 2, F2: (v1 % (((v1 + (v1 % ((1000) | 1)))) | 1))}
    for fi13 := 0; fi13 < 5; fi13++ {
        if v2 {
            var v14 uint32 = 42
            v15 := GenS1{F0: 535, F1: v1, F2: 9}
            var v16 int64 = 42
            _ = v14
            _ = v15
            _ = v16
        } else {
            v17 := GenS2{F0: 107, F1: 0, F2: (v1 - v1)}
            switch (("hello" + "a") + "hello") {
            case "hello":
                v18 := []int{-1}
                _ = v18
            default:
                var v19 int32 = 297
                _ = v19
            }
            _ = v17
        }
        v5 -= (v5 / ((v5) | 1))
        v1 = (v1 ^ v1)
        _ = fi13
    }
    var v20 int = int(v6)
    var v21 uint = 2
    v22 := []int{42}
    var v23 int64 = 242
    fmt.Println(v2, v5, v1)
    switch (("foo bar" + "\n\t") + ("a" + "hello")) {
    case "":
        if (900 < v23) {
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 uint = 65535
                if true {
                    var v26 int = (fi24 - (fi24 / (((v20 ^ 15)) | 1)))
                    var v27 uint64 = (v5 / ((v5) | 1))
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(v2, fi24)
                    v28 := GenS1{F0: (^(v25)), F1: v1, F2: -1}
                    if ((v20 - v20) > -514) {
                        if (340 >= (v1 | (v1 + (v1 - (v1 * (v1 >> ((v25) & 63))))))) {
                            if v2 {
                                for fi29 := 0; fi29 < 11; fi29++ {
                                    if (v6 >= (v5 % ((v6) | 1))) {
                                        var v30 float32 = 3.14159
                                        _ = v30
                                    } else {
                                        var v31 bool = (v23 >= (^(v23)))
                                        if (!v2) {
                                            v32 := GenS1{F0: 18446744073709551615, F1: v1, F2: int8(fi29)}
                                            v33 := GenS3{F0: (-(v23))}
                                            if len(v22) > 0 {
                                                v22[0] = -1
                                            }
                                            v31 = ((v4 * v4) == (v4 - v4))
                                            _ = v32
                                            _ = v33
                                        }
                                        v6 = 42
                                        _ = v31
                                    }
                                    _ = fi29
                                }
                                var v34 float64 = -2.25
                                _ = v34
                            } else {
                                switch (v20 ^ int(v23)) {
                                case 32767:
                                    v25 += (v25 & 4294967295)
                                default:
                                    var v35 float32 = (0.1 + 3.14159)
                                    _ = v35
                                }
                            }
                            for fi36 := 0; fi36 < 11; fi36++ {
                                switch (fi24 >> ((v21) & 63)) {
                                case -983:
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                default:
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                }
                                var v37 uint = v25
                                v38 := GenS2{F0: 416, F1: 42, F2: int16(v37)}
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                            fmt.Println(v1, v23)
                        } else {
                            var v39 bool = (!(true || false))
                            _ = v39
                        }
                        v4 += (v4 % ((v4) | 1))
                    }
                    _ = v28
                }
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                _ = fi24
                _ = v25
            }
        } else {
            fmt.Println(v20, v21)
            v40 := GenS4{F0: 18446744073709551615, F1: v3, F2: (v20 << (((^(v21))) & 63))}
            _ = v40
        }
    case "foo bar":
        fmt.Println(v20, v23)
    case "hello":
        switch ("a" + "") {
        case "\n\t":
            var v41 bool = false
            _ = v41
        }
    }
    var v42 int64 = (v23 * (v23 & (v23 + v23)))
    var v43 int64 = v42
    if true {
        v4 = int32(v5)
        v6 -= (v6 * v5)
    } else {
        if (v2 || v2) {
            v44 := GenS5{F0: v20, F1: -1, F2: true}
            _ = v44
        } else {
            switch 128 {
            case 84:
                var v45 int32 = v4
                _ = v45
            case -2147483648:
                v46 := GenS5{F0: (^(v20)), F1: 127, F2: v2}
                _ = v46
            case -647:
                var v47 float64 = (-2.25 * (0.0 / ((1e10 * 1.5) / 1.0)))
                _ = v47
            default:
                if len(v22) > 0 {
                    v22[0] = v20
                }
            }
            var v48 uint16 = v3
            if (1 >= 0) {
                for fi49 := 0; fi49 < 8; fi49++ {
                    v50 := GenS4{F0: uint64(v48), F1: v48, F2: (fi49 + (fi49 / ((32768) | 1)))}
                    _ = fi49
                    _ = v50
                }
                v5 = v5
            } else {
                v51 := GenS6{F0: "hello"}
                if ((v42 & -1) != (v42 | (v23 % (((v23 % ((v43) | 1))) | 1)))) {
                    v52 := GenS5{F0: v20, F1: 127, F2: v2}
                    _ = v52
                } else {
                    for fi53 := 0; fi53 < 2; fi53++ {
                        var v54 string = "foo bar"
                        var v55 uint8 = 1
                        switch (v54 + "a") {
                        case "\n\t":
                            switch ("\n\t" + (v54 + (v54 + v54))) {
                            case "a":
                                switch "hello" {
                                case "\n\t":
                                    var v56 uint = v21
                                    _ = v56
                                default:
                                    var v57 int16 = (v1 << (((v21 * v21)) & 63))
                                    _ = v57
                                }
                            }
                        case "":
                            fmt.Println(v4)
                        case "a":
                            v55 += v55
                        }
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                    fmt.Println(v48)
                }
                _ = v51
            }
            _ = v48
        }
        fmt.Println(v1)
        v58 := GenS7{F0: 0}
        _ = v58
    }
    for fi59 := 0; fi59 < 9; fi59++ {
        v60 := GenS8{F0: 0, F1: v6}
        _ = fi59
        _ = v60
    }
    var v61 uint32 = 42
    v43 = (^(v23))
    if (2 > 2) {
        fmt.Println(v5)
    } else {
        v62 := GenS9{F0: v42, F1: (v20 * v20)}
        var v63 uint = (^(v21))
        _ = v62
        _ = v63
    }
    for fi64 := 0; fi64 < 19; fi64++ {
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        var v65 bool = v2
        _ = fi64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v61
}
