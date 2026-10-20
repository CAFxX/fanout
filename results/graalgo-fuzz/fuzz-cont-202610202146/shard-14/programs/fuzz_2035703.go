package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int64; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int; F2 string }
func genFn0(p0 int64) int32 {
    var v7 int32 = 0
    v8 := GenS1{F0: 127, F1: (^(p0))}
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := []int{(fi9 % (((fi9 & fi9)) | 1)), (fi9 << ((980) & 63)), -289, (-(fi9))}
        if true {
            if (p0 == 547) {
                var v11 int = (fi9 ^ fi9)
                _ = v11
            } else {
                v12 := GenS1{F0: 255, F1: (p0 * (p0 + p0))}
                _ = len(v10)
                _ = v12
            }
            v7 = (v7 * v7)
            v10 = append(v10, (fi9 | int(p0)))
        } else {
            v13 := GenS1{F0: 255, F1: p0}
            if false {
                for fi14 := 0; fi14 < 17; fi14++ {
                    var v15 int64 = 9223372036854775807
                    _ = fi14
                    _ = v15
                }
                for fi16 := 0; fi16 < 4; fi16++ {
                    v10 = append(v10, fi16)
                    _ = fi16
                }
                for fi17 := 0; fi17 < 7; fi17++ {
                    if (true || true) {
                        var v18 uint = 1
                        _ = v18
                    } else {
                        v7 = 2147483647
                        if len(v10) > 0 {
                            v10[0] = 9223372036854775807
                        }
                        if false {
                            if (p0 == (-(p0))) {
                                switch ("" + "foo bar") {
                                case "\n\t":
                                    v7 = (v7 % (((v7 >> ((407) & 63))) | 1))
                                case "a":
                                    v7 = (^(v7))
                                default:
                                    var v19 int8 = -57
                                    _ = v19
                                }
                                var v20 int64 = 0
                                if true {
                                    var v21 bool = false
                                    _ = v21
                                }
                                _ = v20
                            }
                            if ((2 < 814) && true) {
                                switch ("" + "\n\t") {
                                case "hello":
                                    v22 := GenS1{F0: 42, F1: (p0 + (p0 * p0))}
                                    _ = v22
                                }
                                v23 := GenS2{F0: -1, F1: 44, F2: 42}
                                if false {
                                    switch (("a" + ("a" + "\n\t")) + "hello") {
                                    case "a":
                                        v24 := GenS1{F0: 13, F1: (p0 % (((p0 + (p0 & p0))) | 1))}
                                        _ = v24
                                    case "":
                                        var v25 int8 = 100
                                        _ = v25
                                    case "foo bar":
                                        switch ("hello" + "") {
                                        case "":
                                            fmt.Println(p0, v7)
                                        case "hello":
                                            v10 = append(v10, fi17)
                                        default:
                                            var v26 int = fi17
                                            _ = v26
                                        }
                                    }
                                    p0 += (^(p0))
                                    var v27 int32 = v7
                                    _ = v27
                                }
                                _ = v23
                            }
                        } else {
                            switch (fi9 % ((fi9) | 1)) {
                            case 19:
                                var v28 float32 = -2.25
                                _ = v28
                            case -882:
                                switch ("hello" + "hello") {
                                case "hello":
                                    v7 = v7
                                default:
                                    p0 += (p0 * p0)
                                }
                            default:
                                var v29 uint16 = 42
                                _ = v29
                            }
                            var v30 int64 = 0
                            v30 += p0
                            if (2 != 1) {
                                v10 = append(v10, (fi9 << ((179) & 63)))
                            } else {
                                if (!(p0 != (p0 & (p0 >> ((330) & 63))))) {
                                    v30 = v30
                                    v7 -= v7
                                } else {
                                    var v31 float32 = (1.5 / 3.14159)
                                    v30 = (p0 >> ((18446744073709551615) & 63))
                                    v32 := GenS3{F0: 114, F1: fi9, F2: "\n\t"}
                                    fmt.Println(fi17)
                                    _ = v31
                                    _ = v32
                                }
                            }
                            _ = v30
                        }
                        v7 += (v7 / (((v7 % ((v7) | 1))) | 1))
                    }
                    var v33 uint = 65535
                    v34 := GenS2{F0: 1, F1: 100, F2: 244}
                    _ = fi17
                    _ = v33
                    _ = v34
                }
            } else {
                for fi35 := 0; fi35 < 3; fi35++ {
                    var v36 int32 = v7
                    fmt.Println(fi9, fi35)
                    _ = fi35
                    _ = v36
                }
                if (42 < 0) {
                    v7 = v7
                    for fi37 := 0; fi37 < 14; fi37++ {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        fmt.Println(fi37)
                        _ = fi37
                    }
                }
            }
            fmt.Println(v7, v7)
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    p0 = (p0 - (p0 / ((p0) | 1)))
    var v38 int32 = v7
    _ = p0
    _ = v38
    _ = v7
    _ = v8
    return (-(v7))
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]bool)
    var v3 bool = false
    v4 := make(map[string]uint32)
    var v5 bool = (1 <= 2147483647)
    var v6 int64 = 9223372036854775807
    if false {
        var v39 uint64 = 4294967295
        var v40 int32 = -1
        v3 = v3
        _ = v39
        _ = v40
    }
    switch 128 {
    case -730:
        _ = len(v4)
    default:
        var v41 float32 = 0.1
        _ = v41
    }
    switch 4294967295 {
    case 32767:
        fmt.Println(v3)
    }
    v42 := GenS2{F0: (v6 << ((26) & 63)), F1: 0, F2: v1}
    delete(v4, (("foo bar" + ((((("a" + ("foo bar" + "hello")) + "a") + "a") + "a") + ("a" + ""))) + "\n\t"))
    var v43 uint8 = v1
    for fi44 := 0; fi44 < 15; fi44++ {
        for fi45 := 0; fi45 < 7; fi45++ {
            v46 := []int{fi45}
            _ = fi45
            _ = v46
        }
        var v47 int8 = -128
        _ = fi44
        _ = v47
    }
    var v48 uint8 = v1
    var v49 uint32 = 42
    switch 32767 {
    case -884:
        v50 := []int{-1, -9223372036854775808, 32767, -2147483648}
        _ = v50
    default:
        var v51 int8 = 42
        _ = v51
    }
    if (!false) {
        v52 := []int{979, 128}
        _ = v52
    } else {
        v6 += v6
        fmt.Println(v48, v43)
        var v53 float32 = (-2.25 * -2.25)
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
