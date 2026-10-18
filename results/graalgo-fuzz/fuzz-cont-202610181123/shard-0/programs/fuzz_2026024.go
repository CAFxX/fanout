package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 string; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint; F1 int; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 uint64; F2 uint64 }
func genFn0(p0 bool, p1 int64) int16 {
    p1 = p1
    p0 = (p1 >= p1)
    if ((p1 != (-(p1))) && p0) {
        var v6 uint = 0
        var v7 string = ("a" + "\n\t")
        _ = v6
        _ = v7
    } else {
        p1 = p1
        p1 -= 0
        p0 = true
    }
    if (297 != 575) {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int8 = 42
            v10 := []int{}
            var v11 float32 = -2.25
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        var v12 uint = 42
        p1 = (p1 + 42)
        _ = v12
    } else {
        v13 := []int{-9223372036854775808}
        _ = v13
    }
    var v14 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v14
    return -1
}
func genFn1() uint8 {
    if false {
        var v15 int64 = 9223372036854775807
        _ = v15
    } else {
        v16 := []int{42, -1}
        var v17 uint64 = 1
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 bool = (-508 < 1000)
        var v20 uint16 = 2
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int16 = 0
    _ = v21
    return 128
}

func main() {
    var v1 float64 = 0.1
    var v2 float32 = float32(v1)
    var v3 uint32 = 4294967295
    var v4 uint64 = 715
    v5 := make(map[int]int)
    v4 = v4
    if false {
        var v22 bool = true
        var v23 bool = (!false)
        if (-2147483648 >= 0) {
            v5[1] = 9223372036854775807
            v24 := GenS1{F0: (v3 / ((v3) | 1))}
            var v25 int32 = 2147483647
            _ = v24
            _ = v25
        } else {
            var v26 int = 42
            if v23 {
                var v27 int16 = -32768
                var v28 int16 = (v27 | (v27 ^ v27))
                for fi29 := 0; fi29 < 19; fi29++ {
                    for fi30 := 0; fi30 < 7; fi30++ {
                        switch ("hello" + "a") {
                        case "foo bar":
                            var v31 string = "\n\t"
                            _ = v31
                        case "hello":
                            var v32 int8 = 100
                            switch fi29 {
                            case -1:
                                var v33 uint16 = 2
                                _ = v33
                            }
                            v34 := GenS2{F0: 0, F1: "a", F2: 12}
                            _ = v32
                            _ = v34
                        }
                        _ = fi30
                    }
                    fmt.Println(v2, fi29, v3)
                    v2 = v2
                    _ = fi29
                }
                _ = v27
                _ = v28
            } else {
                v26 = 176
                if (!v22) {
                    if v23 {
                        if (!v23) {
                            for fi35 := 0; fi35 < 12; fi35++ {
                                delete(v5, (fi35 << ((2) & 63)))
                                v26 = (fi35 * fi35)
                                _ = fi35
                            }
                            _ = len(v5)
                            var v36 int32 = -2147483648
                            _ = v36
                        } else {
                            v37 := GenS3{F0: 65535, F1: 255, F2: -9223372036854775808}
                            for fi38 := 0; fi38 < 13; fi38++ {
                                for fi39 := 0; fi39 < 17; fi39++ {
                                    v40 := v5[fi38]
                                    _ = fi39
                                    _ = v40
                                }
                                _ = fi38
                            }
                            _ = v37
                        }
                        switch ("\n\t" + "\n\t") {
                        case "foo bar":
                            for fi41 := 0; fi41 < 7; fi41++ {
                                fmt.Println(v2, v1)
                                v42 := GenS4{F0: 4294967295, F1: 1, F2: v3}
                                switch "hello" {
                                case "":
                                    switch ("a" + "\n\t") {
                                    case "a":
                                        v43 := GenS3{F0: 960, F1: genFn1(), F2: 391}
                                        _ = v43
                                    }
                                case "hello":
                                    switch "hello" {
                                    case "hello":
                                        if true {
                                            v1 = ((float64(v2) - v1) + float64((v2 * 1.5)))
                                            var v44 uint32 = v3
                                            _ = v44
                                        } else {
                                            var v45 uint16 = 65535
                                            var v46 float64 = (0.1 - v1)
                                            _ = v45
                                            _ = v46
                                        }
                                    case "\n\t":
                                        v23 = v23
                                    case "foo bar":
                                        v22 = (("\n\t" <= "hello") || (!true))
                                    default:
                                        if (false || ((v3 + (v3 ^ (v3 % (((^(v3))) | 1)))) != v3)) {
                                            var v47 int32 = -144
                                            v3 = v3
                                            _ = v47
                                        } else {
                                            v2 = ((v2 + (v2 * v2)) / (v2 * (v2 + (3.14159 + float32(float64((((((((float32(v1) / (0.1 - (0.0 * float32(v1)))) * v2) / (v2 / (v2 + (v2 / ((v2 + ((3.14159 - ((v2 * float32(0.1)) + 3.14159)) * 1e10)) - float32(1.5)))))) + (3.14159 + 3.14159)) + (v2 * v2)) + v2) * (v2 + (v2 * ((1.5 / 0.1) / (((1e10 * v2) - v2) + float32((1.5 * 0.0)))))))))))))
                                        }
                                    }
                                case "a":
                                    v1 += (0.1 + v1)
                                }
                                _ = fi41
                                _ = v42
                            }
                        case "hello":
                            var v48 string = ("foo bar" + "")
                            _ = v48
                        default:
                            fmt.Println(v26, v2)
                        }
                    } else {
                        v49 := []int{(v26 * v26), int(v4), v26}
                        v50 := GenS3{F0: uint16(v3), F1: 129, F2: 9223372036854775807}
                        _ = v49
                        _ = v50
                    }
                    v51 := GenS4{F0: 42, F1: 128, F2: v3}
                    switch "\n\t" {
                    case "":
                        if (-1 > -38) {
                            for fi52 := 0; fi52 < 16; fi52++ {
                                if v22 {
                                    v53 := []int{v26, fi52, fi52}
                                    v22 = (42 != -641)
                                    _ = v53
                                }
                                var v54 uint64 = (^(v4))
                                if v23 {
                                    fmt.Println(v22)
                                    var v55 uint16 = 32767
                                    v56 := GenS5{F0: (v26 & v26)}
                                    _ = v55
                                    _ = v56
                                }
                                _ = fi52
                                _ = v54
                            }
                            v57 := []int{(v26 - (v26 | v26))}
                            _ = v57
                        }
                    case "foo bar":
                        fmt.Println(v26, v3, v2)
                    case "hello":
                        if v23 {
                            v23 = v22
                        } else {
                            var v58 int = 9223372036854775807
                            _ = v58
                        }
                    }
                    _ = v51
                }
            }
            if (!(v23 && true)) {
                if v22 {
                    if v22 {
                        switch ("" + "") {
                        case "":
                            v59 := []int{v26, (v26 - 0), v26, v26}
                            _ = v59
                        }
                    }
                    v60 := GenS4{F0: 255, F1: 2147483647, F2: v3}
                    var v61 int32 = 0
                    _ = v60
                    _ = v61
                }
                var v62 int32 = 2147483647
                _ = v62
            } else {
                var v63 int64 = 384
                v64 := v5[-2147483648]
                _ = v63
                _ = v64
            }
            _ = v26
        }
        _ = v22
        _ = v23
    }
    var v65 uint64 = (v4 << ((255) & 63))
    fmt.Println(v4)
    var v66 int16 = -533
    var v67 int8 = -14
    switch 2147483647 {
    case 2147483647:
        if true {
            if false {
                v2 = v2
            }
            v68 := GenS3{F0: 728, F1: 191, F2: 885}
            _ = v68
        }
    case -357:
        var v69 float32 = float32((float64(v2) * v1))
        _ = v69
    }
    v1 += float64(v2)
    if true {
        v70 := GenS1{F0: 42}
        v71 := GenS6{F0: 255}
        _ = v70
        _ = v71
    } else {
        fmt.Println(v1, v4)
        v72 := GenS7{F0: v3, F1: v4, F2: v65}
        _ = v72
    }
    v73 := []int{2147483647, 32768, 128, 2147483647}
    var v74 uint64 = (v4 << ((65535) & 63))
    var v75 float64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v73:", len(v73), v73)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v67
    _ = v73
    _ = v74
    _ = v75
}
