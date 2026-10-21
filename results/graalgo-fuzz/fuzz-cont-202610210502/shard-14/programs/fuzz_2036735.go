package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 string; F1 uint64; F2 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8 }
func genFn0() uint8 {
    switch 32767 {
    case 128:
        var v5 uint16 = 1
        _ = v5
    case -1:
        var v6 bool = (false && true)
        _ = v6
    default:
        var v7 int8 = 102
        _ = v7
    }
    var v8 int32 = 1
    if true {
        if true {
            switch 424 {
            case 42:
                if (1e10 >= float64(1.5)) {
                    if (42 < 1000) {
                        var v9 int32 = (v8 % ((v8) | 1))
                        var v10 string = "foo bar"
                        _ = v10
                        _ = v9
                    }
                    var v11 int64 = 844
                    v8 = -1
                    _ = v11
                }
            case 2147483647:
                if (-9223372036854775808 <= -419) {
                    var v12 float32 = 1.5
                    var v13 int8 = 1
                    _ = v12
                    _ = v13
                } else {
                    v14 := []int{755, 716, 0}
                    switch ("foo bar" + "foo bar") {
                    case "":
                        var v15 int32 = (v8 + v8)
                        _ = v15
                    case "foo bar":
                        var v16 uint32 = 1
                        _ = v16
                    case "hello":
                        fmt.Println(v8)
                    default:
                        var v17 int64 = 750
                        _ = v17
                    }
                    _ = v14
                }
            case 188:
                fmt.Println(v8)
            }
            if (!(!false)) {
                if false {
                    v8 += v8
                    fmt.Println(v8)
                    var v18 uint16 = 32767
                    _ = v18
                }
            }
        } else {
            v19 := []int{0}
            v8 -= (^(v8))
            _ = v19
        }
        fmt.Println(v8)
        switch 529 {
        case 451:
            if true {
                v20 := []int{32767}
                _ = v20
            } else {
                v21 := []int{127, 42}
                _ = v21
            }
        case 32767:
            for fi22 := 0; fi22 < 16; fi22++ {
                switch (("\n\t" + "hello") + "") {
                case "hello":
                    v8 -= (^(v8))
                case "":
                    v8 = v8
                case "a":
                    for fi23 := 0; fi23 < 9; fi23++ {
                        v24 := GenS1{F0: 42}
                        _ = fi23
                        _ = v24
                    }
                default:
                    v8 += 2147483647
                }
                var v25 int16 = 765
                for fi26 := 0; fi26 < 9; fi26++ {
                    v27 := GenS1{F0: 0}
                    var v28 uint16 = 0
                    switch "" {
                    case "":
                        if ((("a" + "hello") + "foo bar") == "a") {
                            fmt.Println(v25, v28, v8)
                            v29 := []int{fi26, (fi26 - (fi26 * (fi22 / (((fi26 % ((42) | 1))) | 1)))), fi22, fi26}
                            _ = v29
                        } else {
                            v30 := []int{fi22}
                            _ = v30
                        }
                        v31 := []int{}
                        switch ("a" + "foo bar") {
                        case "a":
                            var v32 uint32 = 558
                            var v33 float64 = 0.0
                            _ = v32
                            _ = v33
                        case "hello":
                            var v34 string = ""
                            _ = v34
                        case "foo bar":
                            if (((0.0 / (1e10 * float64(-2.25))) + 1e10) >= (-2.25 * 0.0)) {
                                v25 -= v25
                                if len(v31) > 0 {
                                    v31[0] = fi26
                                }
                            } else {
                                if len(v31) >= 2 {
                                    v31 = v31[1:len(v31)]
                                }
                                v8 -= 42
                            }
                        default:
                            v25 = (v25 / (((v25 >> ((1) & 63))) | 1))
                        }
                        _ = v31
                    case "a":
                        if (32767 <= v28) {
                            var v35 string = "hello"
                            var v36 int32 = (v8 ^ (v8 << ((0) & 63)))
                            _ = v35
                            _ = v36
                        } else {
                            v37 := GenS2{F0: (^(v25)), F1: 197}
                            fmt.Println(fi26, v8)
                            v38 := GenS1{F0: 78}
                            _ = v37
                            _ = v38
                        }
                    case "foo bar":
                        fmt.Println(fi22)
                        switch (^(fi26)) {
                        case 1:
                            var v39 int8 = -1
                            _ = v39
                        }
                        var v40 int64 = -9223372036854775808
                        _ = v40
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            var v41 uint = 301
                            _ = v41
                        case "\n\t":
                            var v42 int = 0
                            _ = v42
                        case "a":
                            fmt.Println(v28, fi26, v25)
                        }
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = fi22
                _ = v25
            }
        case 389:
            var v43 uint64 = 2
            _ = v43
        }
    }
    v8 -= v8
    if false {
        fmt.Println(v8)
        var v44 bool = (!false)
        _ = v44
    } else {
        if false {
            for fi45 := 0; fi45 < 9; fi45++ {
                v46 := []int{(fi45 | fi45)}
                _ = fi45
                _ = v46
            }
        }
        var v47 float64 = 1.5
        v48 := GenS3{F0: 2, F1: -394}
        _ = v47
        _ = v48
    }
    _ = v8
    return 255
}

func main() {
    var v1 int64 = 1
    var v2 float32 = 0.1
    v3 := []uint32{0, 2147483647, 2147483648}
    var v4 int = 128
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = 128
    var v49 uint32 = 65535
    v3 = append(v3, v49)
    v50 := GenS4{F0: ("hello" + ("a" + "foo bar")), F1: 976, F2: 778}
    var v51 int16 = -1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, (v49 & (v49 | v49)))
    if (((v1 ^ v1) >= (v1 << ((42) & 63))) || false) {
        if (42 <= 2) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) > 0 {
                v3[0] = v49
            }
        }
    } else {
        for fi52 := 0; fi52 < 4; fi52++ {
            if true {
                v53 := GenS5{F0: 0}
                _ = v53
            } else {
                v51 -= (v51 & 42)
            }
            _ = fi52
        }
    }
    v49 = v49
    v54 := GenS6{F0: 2}
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v55 int32 = 700
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v51:", v51)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v51
    _ = v54
    _ = v55
}
