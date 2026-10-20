package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int16; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int32; F1 uint8 }

func main() {
    var v1 uint16 = 42
    v2 := make(map[string]int)
    var v3 uint = 798
    v4 := make(map[uint32]int)
    switch int(v3) {
    case 267:
        if (!(false || ((!false) || true))) {
            var v5 int32 = 320
            for fi6 := 0; fi6 < 15; fi6++ {
                v7 := GenS1{F0: 1, F1: "\n\t"}
                v8 := []int{fi6, (fi6 + (fi6 >> ((4294967295) & 63))), fi6}
                v9 := GenS2{F0: v1}
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            if (-2.25 >= (0.1 * ((-2.25 / 1.5) + -2.25))) {
                var v10 int16 = 849
                v1 += v1
                _ = v10
            }
            _ = v5
        } else {
            for fi11 := 0; fi11 < 13; fi11++ {
                v1 = (v1 / (((v1 & 32767)) | 1))
                _ = fi11
            }
        }
    case -516:
        var v12 int16 = 0
        _ = v12
    case -1:
        var v13 int = 2147483647
        _ = v13
    default:
        var v14 uint32 = 65535
        _ = v14
    }
    if (true || (false || false)) {
        for fi15 := 0; fi15 < 5; fi15++ {
            if false {
                for fi16 := 0; fi16 < 20; fi16++ {
                    if false {
                        v17 := GenS1{F0: 127, F1: (("\n\t" + "foo bar") + "")}
                        switch ("hello" + ("" + "hello")) {
                        case "\n\t":
                            var v18 float32 = 0.1
                            _ = v18
                        default:
                            fmt.Println(fi15, v3)
                        }
                        _ = v17
                    }
                    v19 := []int{32768}
                    v1 = 1
                    _ = fi16
                    _ = v19
                }
                for fi20 := 0; fi20 < 19; fi20++ {
                    v1 += v1
                    if ((v3 != 2) && (0 > 547)) {
                        var v21 int32 = 42
                        v22 := []int{}
                        _ = v21
                        _ = v22
                    }
                    if true {
                        v23 := []int{(fi20 >> ((v3) & 63))}
                        fmt.Println(v1)
                        _ = v23
                    }
                    fmt.Println(v3, v3, v1)
                    v1 -= (v1 + (v1 | v1))
                    if true {
                        switch (fi15 - fi15) {
                        case 1:
                            var v24 int = fi15
                            _ = v24
                        default:
                            v1 -= v1
                        }
                    } else {
                        v25 := v4[2147483647]
                        _ = v25
                    }
                    if (false && true) {
                        _ = len(v2)
                        v3 = v3
                    }
                    _ = fi20
                }
            } else {
                if true {
                    var v26 uint = v3
                    v3 = (v26 >> (((v26 - (v3 % (((v3 & (v26 | (v3 / (((v26 * (v3 % ((v26) | 1)))) | 1))))) | 1)))) & 63))
                    var v27 bool = (-1 == -690)
                    _ = v26
                    _ = v27
                } else {
                    switch fi15 {
                    case 0:
                        var v28 int = fi15
                        _ = v28
                    }
                    if (0.1 > (0.0 + (1e10 * 1e10))) {
                        var v29 uint = (^(v3))
                        for fi30 := 0; fi30 < 15; fi30++ {
                            v1 += v1
                            v31 := GenS3{F0: 0, F1: 42, F2: 42}
                            _ = fi30
                            _ = v31
                        }
                        _ = v29
                    } else {
                        v32 := []int{(fi15 & -9223372036854775808), (fi15 & fi15)}
                        v1 -= (^(v1))
                        v3 -= uint(v1)
                        _ = v32
                    }
                }
            }
            v3 += (v3 - (v3 * 0))
            fmt.Println(fi15, v3, v1)
            _ = fi15
        }
        v3 = (v3 >> ((2) & 63))
        for fi33 := 0; fi33 < 10; fi33++ {
            var v34 uint16 = v1
            if true {
                if (2 > 2) {
                    if ("hello" > "") {
                        for fi35 := 0; fi35 < 17; fi35++ {
                            v3 = v3
                            _ = fi35
                        }
                        for fi36 := 0; fi36 < 12; fi36++ {
                            fmt.Println(fi33)
                            _ = fi36
                        }
                        if (((!((false && true) && true)) && (!((4294967295 >= 2) || false))) && false) {
                            var v37 string = ("\n\t" + ("a" + "a"))
                            v38 := GenS4{F0: 42, F1: (v3 >> ((42) & 63))}
                            delete(v2, v37)
                            _ = v37
                            _ = v38
                        }
                    } else {
                        v39 := GenS3{F0: -128, F1: -32768, F2: 42}
                        _ = v39
                    }
                    var v40 uint = (v3 | v3)
                    v40 -= 65535
                    _ = v40
                }
                v41 := []int{fi33, fi33, 42, 42}
                _ = v41
            } else {
                v42 := GenS1{F0: 37, F1: ("a" + ("foo bar" + "a"))}
                var v43 int16 = -1
                _ = v42
                _ = v43
            }
            _ = fi33
            _ = v34
        }
    }
    v44 := GenS2{F0: v1}
    var v45 string = ""
    var v46 int64 = -563
    var v47 uint = (v3 << ((v3) & 63))
    v47 += 795
    v46 += -797
    if false {
        v3 = 624
        v48 := GenS5{F0: 1, F1: uint8(v1)}
        _ = v48
    } else {
        v45 = ""
        v47 = (v3 + v47)
        for fi49 := 0; fi49 < 17; fi49++ {
            var v50 uint8 = 1
            v51 := GenS2{F0: (v1 >> (((v47 ^ v47)) & 63))}
            var v52 int32 = 1
            _ = fi49
            _ = v50
            _ = v51
            _ = v52
        }
    }
    fmt.Println(v1)
    var v53 string = v45
    v54 := GenS5{F0: -1, F1: 196}
    var v55 int = -818
    fmt.Println(v47, v47, v3)
    if true {
        var v56 uint16 = 625
        _ = v56
    } else {
        if (1 != 116) {
            var v57 float32 = (1e10 / 1e10)
            v58 := []int{(v55 - (v55 % (((v55 ^ 854)) | 1))), (v55 - v55), (v55 * v55), (v55 / (((v55 % (((v55 << ((391) & 63))) | 1))) | 1))}
            fmt.Println(v57)
            _ = v57
            _ = v58
        }
        var v59 uint16 = (v1 >> ((v47) & 63))
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v53
    _ = v54
    _ = v55
}
