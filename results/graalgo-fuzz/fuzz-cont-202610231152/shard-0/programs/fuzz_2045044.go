package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int8 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint }
func genFn0(p0 int) uint32 {
    var v7 uint64 = 18446744073709551615
    var v8 uint64 = v7
    _ = p0
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 string) uint64 {
    v9 := GenS1{F0: 75, F1: -182, F2: -13}
    fmt.Println(p0)
    var v10 int16 = 32767
    var v11 float32 = ((1e10 + 0.1) - 0.0)
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 42
}

func main() {
    var v1 int16 = 32767
    var v2 uint16 = 42
    var v3 int64 = 1
    var v4 int = 4294967295
    var v5 float32 = -2.25
    var v6 int64 = (v3 / (((v3 >> ((4294967295) & 63))) | 1))
    switch 295 {
    case -621:
        switch (("\n\t" + "") + "\n\t") {
        case "hello":
            if true {
                v12 := GenS1{F0: (v4 ^ (v4 & v4)), F1: 42, F2: 127}
                fmt.Println(v4, v5, v2)
                v13 := []int{v4}
                _ = v12
                _ = v13
            }
        case "foo bar":
            v6 += v3
        default:
            for fi14 := 0; fi14 < 17; fi14++ {
                if (false || false) {
                    for fi15 := 0; fi15 < 8; fi15++ {
                        fmt.Println(v2, v4)
                        switch "hello" {
                        case "foo bar":
                            if (-128 < 49) {
                                v4 = (fi15 + (v4 ^ v4))
                                v16 := GenS2{F0: (((-(v1)) < v1) && true), F1: 42}
                                var v17 bool = false
                                _ = v16
                                _ = v17
                            }
                        default:
                            var v18 uint32 = 2
                            _ = v18
                        }
                        switch (v4 >> ((519) & 63)) {
                        case 127:
                            if true {
                                v19 := GenS3{F0: 2147483647}
                                var v20 uint32 = genFn0(fi14)
                                switch "hello" {
                                case "\n\t":
                                    v6 -= (v3 ^ (-(v6)))
                                case "":
                                    if false {
                                        v4 = (fi14 >> ((65535) & 63))
                                        if false {
                                            var v21 uint64 = genFn1("\n\t")
                                            var v22 uint16 = (v2 & v2)
                                            fmt.Println(v5, v5)
                                            _ = v21
                                            _ = v22
                                        } else {
                                            var v23 int16 = 1
                                            _ = v23
                                        }
                                        var v24 uint16 = (^(v2))
                                        _ = v24
                                    } else {
                                        v25 := []int{128}
                                        switch (v4 & 605) {
                                        case 2147483647:
                                            v26 := GenS4{F0: (^(fi14)), F1: 453}
                                            v1 -= -246
                                            var v27 bool = (!(0.0 > 0.0))
                                            var v28 uint64 = genFn1("")
                                            v29 := GenS2{F0: true, F1: 0}
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        case 73:
                                            v30 := GenS3{F0: (^(v20))}
                                            _ = v30
                                        default:
                                            if ((v6 * v6) <= (v3 ^ v3)) {
                                                fmt.Println(v5)
                                                v31 := GenS3{F0: (v20 / ((v20) | 1))}
                                                v32 := GenS2{F0: (((v20 >> ((42) & 63)) <= (v20 + v20)) && false), F1: 404}
                                                _ = v31
                                                _ = v32
                                            }
                                        }
                                        var v33 float32 = v5
                                        _ = v25
                                        _ = v33
                                    }
                                    v34 := []int{v4}
                                    _ = v34
                                }
                                _ = v19
                                _ = v20
                            } else {
                                v6 -= v3
                                var v35 int64 = -342
                                _ = v35
                            }
                        case 0:
                            v3 -= -9223372036854775808
                        default:
                            var v36 int16 = v1
                            _ = v36
                        }
                        v37 := GenS2{F0: true, F1: 2}
                        v38 := GenS1{F0: (fi14 - 32767), F1: -2147483648, F2: 100}
                        _ = fi15
                        _ = v37
                        _ = v38
                    }
                }
                _ = fi14
            }
        }
    case 32767:
        v39 := []int{(v4 - v4), int(v3)}
        _ = v39
    case 9223372036854775807:
        switch v4 {
        case 32768:
            var v40 uint32 = 4294967295
            _ = v40
        }
    }
    if (false || true) {
        var v41 int8 = -128
        for fi42 := 0; fi42 < 8; fi42++ {
            v6 -= (v6 << ((4294967295) & 63))
            for fi43 := 0; fi43 < 4; fi43++ {
                if (!((("a" + "a") + "\n\t") >= "")) {
                    v3 = 0
                    v44 := []int{-522, fi42}
                    fmt.Println(v41)
                    v2 = v2
                    _ = v44
                } else {
                    var v45 bool = true
                    _ = v45
                }
                v46 := GenS3{F0: 1}
                var v47 bool = (false && true)
                v48 := GenS2{F0: true, F1: 2147483648}
                _ = fi43
                _ = v46
                _ = v47
                _ = v48
            }
            _ = fi42
        }
        _ = v41
    } else {
        var v49 string = ((("" + "foo bar") + "foo bar") + "")
        _ = v49
    }
    v50 := []int{}
    var v51 int = (v4 >> ((18446744073709551615) & 63))
    v3 = v3
    v52 := GenS4{F0: 4294967295, F1: 1}
    var v53 string = ("" + "\n\t")
    switch (v53 + v53) {
    case "hello":
        var v54 int8 = -1
        _ = v54
    }
    if (v3 >= v6) {
        fmt.Println(v51, v53)
        if false {
            if false {
                var v55 float32 = (v5 * 1e10)
                var v56 float32 = 1e10
                _ = v55
                _ = v56
            } else {
                fmt.Println(v1, v51, v51)
                v57 := GenS3{F0: 493}
                _ = v57
            }
        }
    } else {
        v58 := GenS3{F0: 2147483648}
        _ = v58
    }
    for fi59 := 0; fi59 < 17; fi59++ {
        fmt.Println(v6, v3)
        for fi60 := 0; fi60 < 2; fi60++ {
            v6 = v6
            v61 := GenS4{F0: (^(v4)), F1: 222}
            _ = fi60
            _ = v61
        }
        var v62 uint64 = 4294967295
        _ = fi59
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v6
}
