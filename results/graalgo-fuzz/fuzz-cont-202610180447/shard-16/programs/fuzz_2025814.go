package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0() uint16 {
    var v7 uint = 325
    switch "" {
    case "":
        switch "a" {
        case "foo bar":
            var v8 float64 = 1e10
            _ = v8
        case "hello":
            v9 := GenS1{F0: (v7 << ((v7) & 63)), F1: -32768, F2: true}
            _ = v9
        case "":
            var v10 float32 = 1.5
            _ = v10
        }
    case "foo bar":
        v7 = v7
    }
    v11 := GenS1{F0: v7, F1: -354, F2: true}
    if false {
        var v12 uint64 = 665
        _ = v12
    }
    v7 = (^(v7))
    _ = v11
    _ = v7
    return 0
}
func genFn1(p0 int64) int16 {
    var v13 float64 = 0.1
    fmt.Println(p0)
    p0 = (p0 - p0)
    _ = p0
    _ = v13
    return -248
}

func main() {
    var v1 int8 = -128
    var v2 uint = 18446744073709551615
    var v3 uint8 = 128
    v4 := []string{}
    var v5 uint32 = 65535
    var v6 uint32 = 0
    v3 = 103
    var v14 uint = 18446744073709551615
    var v15 int = -9223372036854775808
    v5 -= (^(v6))
    fmt.Println(v15, v5, v14)
    if false {
        v5 -= v6
        v16 := GenS2{F0: 2}
        _ = v16
    } else {
        if (2147483647 >= int32(v14)) {
            var v17 bool = false
            v1 = v1
            fmt.Println(v5, v17, v5)
            _ = v17
        } else {
            v3 += 127
            if true {
                for fi18 := 0; fi18 < 3; fi18++ {
                    if (false || true) {
                        var v19 bool = (v1 > v1)
                        _ = v19
                    } else {
                        fmt.Println(v2, v14, fi18)
                        if true {
                            var v20 string = ""
                            _ = v20
                        }
                    }
                    _ = fi18
                }
                for fi21 := 0; fi21 < 1; fi21++ {
                    var v22 uint = 42
                    v23 := GenS2{F0: 2}
                    if (v3 < (v3 % ((v3) | 1))) {
                        var v24 int = fi21
                        var v25 int16 = 793
                        if true {
                            v26 := GenS1{F0: v14, F1: (^(v25)), F2: true}
                            switch fi21 {
                            case -1000:
                                var v27 uint8 = (v3 - 2)
                                _ = v27
                            case 127:
                                var v28 uint64 = 42
                                _ = v28
                            default:
                                var v29 float64 = (1e10 - 1e10)
                                _ = v29
                            }
                            for fi30 := 0; fi30 < 17; fi30++ {
                                v31 := GenS3{F0: (v3 ^ 2)}
                                v4 = append(v4, "\n\t")
                                if (-914 != (v25 >> (((v2 / ((v14) | 1))) & 63))) {
                                    var v32 string = "hello"
                                    var v33 bool = ((!((0 >= 32767) && (-442 < -410))) && true)
                                    _ = v32
                                    _ = v33
                                } else {
                                    v25 += v25
                                    var v34 int = fi30
                                    _ = len(v4)
                                    _ = v34
                                }
                                _ = fi30
                                _ = v31
                            }
                            _ = v26
                        }
                        _ = v24
                        _ = v25
                    } else {
                        v6 = v5
                        v35 := GenS1{F0: v2, F1: -32768, F2: (v2 != uint(v3))}
                        switch (^(fi21)) {
                        case -2147483648:
                            v36 := GenS1{F0: v2, F1: int16(v5), F2: false}
                            _ = v36
                        case 9223372036854775807:
                            v5 = 2147483647
                        case 127:
                            var v37 int8 = (v1 & v1)
                            _ = v37
                        }
                        _ = v35
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
            var v38 uint64 = 2
            _ = v38
        }
        v39 := GenS2{F0: uint64(v2)}
        v14 += (v14 - 4294967295)
        _ = v39
    }
    v4 = append(v4, ("a" + "foo bar"))
    if (((0.0 - ((((0.1 - float64(0.0)) * 3.14159) * ((-2.25 / -2.25) - float64(0.0))) / 0.1)) - 1e10) <= (-2.25 + 0.1)) {
        switch 0 {
        case 4294967295:
            if len(v4) > 0 {
                v4[0] = ("foo bar" + "")
            }
        case 872:
            for fi40 := 0; fi40 < 14; fi40++ {
                for fi41 := 0; fi41 < 20; fi41++ {
                    if false {
                        if (2 == uint64(v3)) {
                            v42 := GenS1{F0: v14, F1: 32767, F2: (730 <= 1)}
                            var v43 int16 = genFn1(-823)
                            var v44 bool = false
                            _ = v42
                            _ = v43
                            _ = v44
                        } else {
                            v2 += (v14 | (v14 << (((v2 >> (((v2 ^ (^(v2)))) & 63))) & 63)))
                            fmt.Println(v15, v6, v1)
                            var v45 int8 = (v1 % (((-(v1))) | 1))
                            _ = v45
                        }
                        if false {
                            var v46 string = "\n\t"
                            switch (v46 + "hello") {
                            case "hello":
                                fmt.Println(v46, v5)
                            case "":
                                fmt.Println(fi41)
                            case "\n\t":
                                _ = len(v4)
                            }
                            fmt.Println(v3)
                            _ = v46
                        }
                        var v47 uint64 = 995
                        if true {
                            if len(v4) > 0 {
                                v4[0] = (("hello" + "a") + "")
                            }
                            v47 = v47
                        } else {
                            var v48 string = "foo bar"
                            _ = v48
                        }
                        fmt.Println(v3)
                        var v49 uint = (^(v2))
                        var v50 int64 = 0
                        _ = v47
                        _ = v49
                        _ = v50
                    } else {
                        fmt.Println(fi40, v3, v15)
                        var v51 float32 = 0.1
                        var v52 string = "hello"
                        var v53 uint64 = 0
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    _ = fi41
                }
                if (v3 > v3) {
                    var v54 int = (-(v15))
                    _ = v54
                } else {
                    _ = len(v4)
                    var v55 bool = (694 != v14)
                    _ = v55
                }
                if ((float32((1e10 / 1.0)) / 1.0) != ((float32(0.0) - -2.25) - 1e10)) {
                    fmt.Println(v14, fi40, v5)
                    if ((v2 * (v2 ^ v2)) < v2) {
                        v56 := GenS2{F0: 1}
                        var v57 uint8 = v3
                        for fi58 := 0; fi58 < 10; fi58++ {
                            fmt.Println(v14)
                            if (uint(v57) == (v14 * (v2 & v14))) {
                                if ((true || ((!((107 < (v1 ^ v1)) || true)) && (true || ((1e10 * -2.25) >= ((0.0 / 1.5) + ((0.1 * float64((3.14159 - ((-2.25 * float32(float64(3.14159))) / 1.5)))) + (3.14159 / 3.14159))))))) || false) {
                                    v4 = append(v4, "foo bar")
                                    _ = len(v4)
                                    var v59 int32 = -2147483648
                                    _ = v59
                                }
                                fmt.Println(v14, v14)
                                fmt.Println(fi58)
                            }
                            _ = fi58
                        }
                        _ = v56
                        _ = v57
                    }
                    var v60 uint64 = 42
                    _ = v60
                }
                _ = fi40
            }
        case 9223372036854775807:
            fmt.Println(v15, v6, v6)
        default:
            var v61 int64 = 1
            _ = v61
        }
    } else {
        if true {
            var v62 bool = true
            var v63 int8 = -91
            v1 = v63
            _ = v62
            _ = v63
        }
        var v64 string = ("hello" + ("" + "a"))
        _ = v64
    }
    var v65 uint8 = (v3 + v3)
    var v66 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
}
