package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 uint64, p1 bool) int8 {
    var v7 float64 = 1e10
    v8 := []int{}
    p0 = (p0 >> ((587) & 63))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -39
}

func main() {
    v1 := []string{"", (("foo bar" + "foo bar") + "\n\t"), "foo bar"}
    v2 := make(map[uint32]uint32)
    v3 := []int{-359}
    v4 := []string{"", ((("a" + (("\n\t" + ("hello" + ((("foo bar" + (("\n\t" + ("\n\t" + ("" + "hello"))) + ("hello" + ("a" + ("" + (("\n\t" + ("foo bar" + ("\n\t" + ""))) + "foo bar")))))) + "\n\t") + ("a" + "foo bar")))) + "hello")) + "") + "foo bar")}
    var v5 uint64 = 2
    v6 := make(map[string]uint32)
    _ = len(v3)
    switch "" {
    case "a":
        delete(v6, "a")
    case "":
        var v9 uint = 65535
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            if true {
                var v12 uint64 = (v5 - (v5 & (v5 + (v5 % ((18446744073709551615) | 1)))))
                switch "" {
                case "a":
                    var v13 string = ((("a" + "foo bar") + "\n\t") + "hello")
                    fmt.Println(v13, fi10, fi11)
                    _ = v13
                case "":
                    v12 -= v12
                default:
                    if (!(((v5 % (((v5 & (v12 >> ((4294967295) & 63)))) | 1)) == v5) || (true || ((^(fi11)) >= (-(fi10)))))) {
                        if ((((float32(3.14159) / 1e10) / (0.1 + float32(-2.25))) <= 3.14159) || (0 <= 128)) {
                            if ((v12 ^ (v12 / ((v5) | 1))) == v12) {
                                _ = len(v1)
                            }
                            var v14 string = "hello"
                            var v15 int16 = -32768
                            _ = v14
                            _ = v15
                        } else {
                            var v16 uint16 = 32767
                            switch fi11 {
                            case -2147483648:
                                v12 = 2
                            case 1:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case 0:
                                v17 := GenS1{F0: 123, F1: 2147483647, F2: 2}
                                _ = v17
                            default:
                                fmt.Println(fi11, fi11, fi10)
                            }
                            _ = v16
                        }
                        if (32767 <= 124) {
                            v18 := v2[1]
                            _ = v18
                        }
                        v5 = v5
                        delete(v2, 0)
                    }
                }
                v12 = (v12 ^ (v5 / (((v12 * 0)) | 1)))
                v12 -= 42
                if (-2.25 < (1.5 - 0.1)) {
                    fmt.Println(v5, fi11)
                }
                switch (("\n\t" + ("\n\t" + ("" + "\n\t"))) + "a") {
                case "hello":
                    v19 := GenS1{F0: 128, F1: 2147483647, F2: 435}
                    _ = v19
                }
                var v20 int = fi10
                if true {
                    v5 = v12
                    v12 += 401
                }
                switch (fi10 | (fi10 + (v20 / ((fi11) | 1)))) {
                case -422:
                    if true {
                        fmt.Println(v20, v5, fi11)
                        v12 += 860
                    } else {
                        var v21 bool = (!true)
                        _ = v21
                    }
                case -2147483648:
                    var v22 uint32 = 908
                    switch "\n\t" {
                    case "a":
                        switch "hello" {
                        case "":
                            var v23 int64 = -1
                            _ = v23
                        case "a":
                            v24 := GenS1{F0: 203, F1: 42, F2: v22}
                            _ = v24
                        case "\n\t":
                            v5 = v12
                            switch fi11 {
                            case 1:
                                v12 -= 1
                            case 167:
                                v25 := GenS1{F0: 1, F1: int32(v5), F2: v22}
                                _ = v25
                            case 128:
                                v26 := GenS1{F0: 0, F1: -1, F2: v22}
                                if len(v3) > 0 {
                                    v3[0] = int(v22)
                                }
                                var v27 uint16 = 255
                                _ = v26
                                _ = v27
                            default:
                                v5 += v12
                            }
                            v20 += 419
                        default:
                            switch fi11 {
                            case 127:
                                if (!true) {
                                    var v28 uint8 = 2
                                    v5 -= 4294967295
                                    _ = v28
                                } else {
                                    v20 -= 128
                                }
                            case 833:
                                var v29 uint8 = 1
                                _ = v29
                            case 256:
                                var v30 uint = 2
                                _ = v30
                            }
                        }
                    default:
                        if true {
                            var v31 int = v20
                            v22 = v22
                            var v32 float32 = 1e10
                            fmt.Println(v5, v12)
                            v32 -= float32(-2.25)
                            _ = v31
                            _ = v32
                        } else {
                            v12 = v5
                            v6["foo bar"] = v22
                            if (("hello" + "hello") > "") {
                                v33 := v6["a"]
                                _ = v33
                            } else {
                                _ = len(v3)
                                fmt.Println(fi10, fi11)
                                delete(v2, v22)
                                v12 = v12
                            }
                        }
                    }
                    v34 := GenS2{F0: 255, F1: 1}
                    _ = v22
                    _ = v34
                default:
                    v35 := GenS3{F0: 0, F1: false}
                    _ = v35
                }
                _ = v12
                _ = v20
            } else {
                switch fi10 {
                case 9223372036854775807:
                    v5 = (v5 ^ 2)
                }
                var v36 string = ("foo bar" + "a")
                if (((1e10 * 3.14159) * (float64(1.5) + -2.25)) >= float64(0.1)) {
                    switch fi11 {
                    case 0:
                        if false {
                            if false {
                                v5 = 204
                                if true {
                                    v3 = append(v3, (fi10 / (((fi10 + 0)) | 1)))
                                    v5 += v5
                                }
                                delete(v6, v36)
                                var v37 uint8 = 2
                                _ = v37
                            }
                            if false {
                                v36 = (v36 + v36)
                                v38 := GenS3{F0: 0, F1: true}
                                _ = v38
                            }
                            fmt.Println(v36, v5)
                        }
                    }
                    v5 += v5
                    if (true && (0 <= genFn0((v5 - (v5 * (v5 + (v5 % ((v5) | 1))))), (0.0 > float64((1.5 - 1e10)))))) {
                        v5 -= 948
                    } else {
                        _ = len(v1)
                        if (false || false) {
                            v39 := GenS2{F0: 42, F1: (v5 << ((65535) & 63))}
                            v5 = v5
                            fmt.Println(fi10, fi10, v36)
                            fmt.Println(fi10, fi11, fi11)
                            _ = v39
                        }
                    }
                    v40 := GenS4{F0: 9223372036854775807, F1: v5, F2: 255}
                    fmt.Println(fi10)
                    v4 = append(v4, (v36 + v36))
                    v41 := GenS3{F0: 0, F1: (127 < 2)}
                    if (9223372036854775807 <= 0) {
                        if (-2.25 > (-2.25 * 3.14159)) {
                            v36 = (v36 + ("hello" + "foo bar"))
                            if (((v5 ^ (v5 % (((v5 >> ((175) & 63))) | 1))) == 51) && true) {
                                v5 -= (v5 | v5)
                                v36 = v36
                                v5 = (v5 + (v5 / ((42) | 1)))
                            }
                        } else {
                            v36 = (v36 + "a")
                        }
                        v42 := GenS5{F0: 50, F1: 0, F2: 88}
                        _ = v42
                    }
                    if false {
                        v43 := GenS6{F0: v36}
                        v44 := GenS7{F0: 35, F1: 100}
                        _ = v43
                        _ = v44
                    } else {
                        if (375 >= 2) {
                            var v45 int64 = -668
                            if ((!(59 > 32767)) && true) {
                                var v46 int8 = 1
                                if false {
                                    if len(v4) > 0 {
                                        v4[0] = v36
                                    }
                                    v47 := v6[v36]
                                    _ = v47
                                }
                                _ = v46
                            } else {
                                var v48 float64 = 1e10
                                _ = v48
                            }
                            _ = v45
                        }
                        fmt.Println(v36, fi11)
                    }
                    _ = v40
                    _ = v41
                } else {
                    v5 = (v5 + (^(v5)))
                }
                _ = v36
            }
            _ = fi11
        }
        v49 := GenS8{F0: 255}
        var v50 uint64 = (^(v5))
        _ = fi10
        _ = v49
        _ = v50
    }
    var v51 float32 = 0.1
    v5 = v5
    fmt.Println(v51)
    var v52 int8 = -128
    _ = len(v2)
    for fi53 := 0; fi53 < 8; fi53++ {
        var v54 string = ""
        v55 := GenS4{F0: 0, F1: 742, F2: 1}
        _ = fi53
        _ = v54
        _ = v55
    }
    v56 := GenS8{F0: 160}
    switch -641 {
    case 42:
        var v57 float64 = -2.25
        _ = v57
    case 260:
        switch 32768 {
        case 42:
            var v58 bool = false
            _ = v58
        }
    case 32768:
        var v59 string = ("a" + ("foo bar" + "\n\t"))
        _ = v59
    default:
        var v60 int8 = v52
        _ = v60
    }
    v52 = v52
    switch ("\n\t" + "a") {
    case "\n\t":
        _ = len(v4)
    default:
        var v61 int = -962
        _ = v61
    }
    switch "foo bar" {
    case "foo bar":
        var v62 int = -1
        _ = v62
    case "a":
        var v63 int16 = 0
        _ = v63
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v56
    _ = v6
}
