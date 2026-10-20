package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint32; F1 string }
func genFn0() uint8 {
    switch 2147483647 {
    case 32768:
        if (0 >= 2) {
            v4 := GenS1{F0: 4294967295}
            for fi5 := 0; fi5 < 9; fi5++ {
                v6 := []int{}
                _ = fi5
                _ = v6
            }
            _ = v4
        } else {
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 uint32 = 0
                _ = fi7
                _ = v8
            }
            var v9 uint64 = 1
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 uint32 = 65535
            v12 := GenS2{F0: 1, F1: 2}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 uint = 1
    switch "hello" {
    case "foo bar":
        v14 := []int{2147483647, 724}
        _ = v14
    }
    fmt.Println(v13)
    for fi15 := 0; fi15 < 3; fi15++ {
        v13 -= v13
        _ = fi15
    }
    _ = v13
    return 42
}
func genFn1(p0 bool, p1 uint64) uint8 {
    p1 -= (p1 >> ((65535) & 63))
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := []int{(fi16 ^ fi16), fi16, fi16}
        for fi18 := 0; fi18 < 14; fi18++ {
            if p0 {
                if len(v17) > 0 {
                    v17[0] = 42
                }
                var v19 int8 = 127
                v19 += (v19 % (((v19 / (((v19 >> ((42) & 63))) | 1))) | 1))
                _ = v19
            }
            var v20 uint8 = genFn0()
            v21 := GenS1{F0: (fi16 * (fi18 - (fi16 | fi18)))}
            _ = fi18
            _ = v20
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    v22 := GenS2{F0: 4294967295, F1: 1}
    var v23 uint8 = 41
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return (v23 & 138)
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -890
    var v3 string = ("\n\t" + "a")
    var v24 int8 = 42
    for fi25 := 0; fi25 < 15; fi25++ {
        if false {
            if (32767 == 474) {
                if false {
                    fmt.Println(v2, v2, v2)
                    v1[("hello" + "\n\t")] = (false && true)
                    var v26 string = v3
                    _ = v26
                }
                if (!false) {
                    v27 := GenS3{F0: 0}
                    v2 += -1
                    if true {
                        if (false && false) {
                            for fi28 := 0; fi28 < 6; fi28++ {
                                switch ("a" + (v3 + v3)) {
                                case "a":
                                    v29 := GenS4{F0: 2, F1: 0, F2: 255}
                                    _ = v29
                                }
                                _ = fi28
                            }
                            v2 = v2
                            for fi30 := 0; fi30 < 1; fi30++ {
                                var v31 int = -2147483648
                                if (!false) {
                                    switch -9223372036854775808 {
                                    case -2147483648:
                                        fmt.Println(v31, v24, v3)
                                    case -856:
                                        v32 := GenS3{F0: 0}
                                        _ = v32
                                    case 9223372036854775807:
                                        v31 -= fi30
                                    }
                                    var v33 float32 = 0.1
                                    v24 = (v24 & (v24 << ((2) & 63)))
                                    _ = v33
                                } else {
                                    var v34 uint16 = 95
                                    var v35 float32 = 1e10
                                    v36 := []int{fi30}
                                    v2 += (v2 * (v2 - v2))
                                    v37 := GenS3{F0: 2147483647}
                                    if false {
                                        v35 = ((-2.25 / v35) + (v35 / v35))
                                    }
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                                _ = fi30
                                _ = v31
                            }
                        }
                    } else {
                        if false {
                            if true {
                                fmt.Println(v24)
                                v38 := GenS5{F0: v2, F1: v2}
                                _ = v38
                            }
                        } else {
                            v39 := []int{}
                            v40 := GenS3{F0: 2147483648}
                            _ = v39
                            _ = v40
                        }
                        v24 += 127
                    }
                    _ = v27
                } else {
                    fmt.Println(v3, v24)
                    var v41 float32 = -2.25
                    for fi42 := 0; fi42 < 7; fi42++ {
                        fmt.Println(v24, v24)
                        v43 := GenS3{F0: 394}
                        v44 := GenS6{F0: 2147483647, F1: v3}
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    _ = v41
                }
                if (fi25 == -108) {
                    v45 := GenS6{F0: 4294967295, F1: v3}
                    v46 := GenS3{F0: 42}
                    _ = v45
                    _ = v46
                } else {
                    for fi47 := 0; fi47 < 9; fi47++ {
                        switch fi25 {
                        case -174:
                            var v48 int8 = v24
                            _ = v48
                        }
                        _ = fi47
                    }
                    for fi49 := 0; fi49 < 15; fi49++ {
                        switch fi49 {
                        case 42:
                            switch v3 {
                            case "\n\t":
                                var v50 string = v3
                                _ = v50
                            }
                        case 32767:
                            if ((320 >= (v2 >> ((18446744073709551615) & 63))) || true) {
                                var v51 uint16 = 42
                                switch (v3 + ("hello" + v3)) {
                                case "":
                                    if (!((v51 > v51) && true)) {
                                        v52 := []int{(fi25 / ((782) | 1)), fi25, (fi49 / ((int(v51)) | 1))}
                                        v3 = v3
                                        _ = v52
                                    }
                                }
                                fmt.Println(v2, v51)
                                _ = v51
                            } else {
                                var v53 float64 = 3.14159
                                switch 2147483647 {
                                case -404:
                                    var v54 float64 = (v53 - (v53 + 1e10))
                                    _ = v54
                                case 1:
                                    if false {
                                        fmt.Println(fi25, fi25, v53)
                                    }
                                case -1:
                                    var v55 int32 = 600
                                    _ = v55
                                }
                                _ = v53
                            }
                        default:
                            var v56 uint16 = 65535
                            _ = v56
                        }
                        _ = fi49
                    }
                }
            }
            v2 = 1
            v3 = v3
        } else {
            if (1 < 135) {
                v24 = 0
            }
            v3 = "foo bar"
        }
        _ = fi25
    }
    var v57 uint32 = 936
    if (false && (!(false || (9223372036854775807 > -594)))) {
        v58 := GenS3{F0: v57}
        v59 := []int{-9223372036854775808, -1}
        _ = v58
        _ = v59
    } else {
        fmt.Println(v24, v24, v3)
        v57 += v57
    }
    switch (v3 + v3) {
    case "hello":
        fmt.Println(v57, v2, v3)
    case "a":
        v60 := GenS1{F0: 2147483647}
        _ = v60
    default:
        var v61 bool = false
        _ = v61
    }
    var v62 int16 = 302
    v63 := []int{128}
    var v64 uint8 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v57:", v57)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v63:", len(v63), v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v57
    _ = v62
    _ = v63
    _ = v64
}
