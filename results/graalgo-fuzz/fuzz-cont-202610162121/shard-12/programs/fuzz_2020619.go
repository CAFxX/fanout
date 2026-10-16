package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 int; F2 uint }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64; F1 uint }
func genFn0() int8 {
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 float32 = 0.0
        v7 += v7
        v8 := GenS1{F0: -114, F1: 32767}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 string = (("foo bar" + "a") + ("hello" + "hello"))
    fmt.Println(v9)
    for fi10 := 0; fi10 < 11; fi10++ {
        v9 = v9
        var v11 int = fi10
        _ = fi10
        _ = v11
    }
    _ = v9
    return -1
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]uint32)
    var v3 bool = false
    var v4 int64 = -1
    var v5 int64 = (^(v4))
    var v12 int = 42
    switch (-(v12)) {
    case 127:
        v13 := GenS2{F0: v3}
        _ = v13
    case 1:
        var v14 float32 = 1e10
        _ = v14
    case -90:
        var v15 uint64 = 4294967295
        _ = v15
    }
    var v16 uint16 = 965
    switch v12 {
    case 4294967295:
        v17 := []int{0, v12}
        _ = v17
    case 32767:
        for fi18 := 0; fi18 < 4; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                switch ("" + "foo bar") {
                case "foo bar":
                    v4 = (v4 << ((2) & 63))
                case "\n\t":
                    v5 = (v4 - (^(v4)))
                    switch "" {
                    case "\n\t":
                        v1 += -58
                    case "hello":
                        if (!v3) {
                            if (float64(0.1) <= (3.14159 - float64(1.5))) {
                                fmt.Println(v1, v16, fi19)
                                if false {
                                    if v3 {
                                        v20 := GenS1{F0: 100, F1: (-(v1))}
                                        switch ("" + "\n\t") {
                                        case "hello":
                                            switch "hello" {
                                            case "":
                                                v21 := GenS3{F0: 729, F1: "a"}
                                                var v22 uint = 1
                                                if v3 {
                                                    var v23 int = v12
                                                    if (32767 < v1) {
                                                        var v24 uint16 = v16
                                                        v25 := GenS3{F0: -2147483648, F1: ""}
                                                        v26 := GenS4{F0: 914}
                                                        _ = v24
                                                        _ = v25
                                                        _ = v26
                                                    } else {
                                                        v2[65535] = 1
                                                    }
                                                    v12 -= (v23 | (fi18 / ((v12) | 1)))
                                                    _ = v23
                                                } else {
                                                    var v27 int16 = (v1 ^ v1)
                                                    _ = v27
                                                }
                                                _ = v21
                                                _ = v22
                                            case "foo bar":
                                                v5 += v4
                                            case "\n\t":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v28 int32 = -1
                                                    _ = v28
                                                case "hello":
                                                    var v29 float64 = 3.14159
                                                    _ = v29
                                                }
                                            default:
                                                var v30 int32 = 0
                                                _ = v30
                                            }
                                        case "foo bar":
                                            fmt.Println(v16)
                                        }
                                        if v3 {
                                            v2[1] = 861
                                            fmt.Println(v1, v3)
                                            var v31 float32 = -2.25
                                            if v3 {
                                                v32 := []int{-278}
                                                _ = v32
                                            } else {
                                                delete(v2, 797)
                                                if v3 {
                                                    if (((42 <= 167) || (1 < (v1 % (((^(v1))) | 1)))) || (!v3)) {
                                                        if true {
                                                            v33 := GenS2{F0: v3}
                                                            v34 := []int{fi18, fi19, v12, v12}
                                                            _ = v33
                                                            _ = v34
                                                        } else {
                                                            var v35 int16 = (v1 & v1)
                                                            v12 = 0
                                                            _ = v35
                                                        }
                                                        v36 := GenS1{F0: genFn0(), F1: v1}
                                                        var v37 uint32 = 2147483648
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    var v38 int = int(v1)
                                                    _ = v38
                                                } else {
                                                    v16 = (v16 >> ((0) & 63))
                                                    v39 := []int{fi19}
                                                    v40 := GenS5{F0: 65535, F1: (fi19 / ((int(v16)) | 1)), F2: 1000}
                                                    _ = v39
                                                    _ = v40
                                                }
                                            }
                                            v41 := GenS6{F0: 42, F1: -1}
                                            _ = v31
                                            _ = v41
                                        }
                                        _ = v20
                                    } else {
                                        v42 := GenS7{F0: (fi19 << ((18446744073709551615) & 63))}
                                        _ = v42
                                    }
                                    v5 = 1
                                    v12 -= (fi19 ^ fi18)
                                } else {
                                    switch "" {
                                    case "foo bar":
                                        if true {
                                            switch "\n\t" {
                                            case "foo bar":
                                                v4 -= v5
                                            case "hello":
                                                var v43 uint8 = 1
                                                _ = v43
                                            case "a":
                                                if true {
                                                    v3 = v3
                                                    v44 := []int{fi19, 9223372036854775807, 128}
                                                    if (false && v3) {
                                                        if false {
                                                            v45 := GenS1{F0: 0, F1: v1}
                                                            _ = v45
                                                        }
                                                        var v46 uint16 = v16
                                                        v47 := GenS8{F0: v4, F1: 6}
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        if (2 != 2147483647) {
                                                            if len(v44) > 0 {
                                                                v44[0] = (v12 + 2147483647)
                                                            }
                                                        }
                                                    }
                                                    _ = v44
                                                } else {
                                                    var v48 int = (fi19 * 42)
                                                    v49 := GenS4{F0: v5}
                                                    var v50 int8 = int8(v16)
                                                    _ = v48
                                                    _ = v49
                                                    _ = v50
                                                }
                                            }
                                            v51 := []int{}
                                            var v52 float32 = -2.25
                                            if true {
                                                var v53 uint = 42
                                                var v54 uint = (v53 - (v53 + 954))
                                                _ = v53
                                                _ = v54
                                            }
                                            _ = v51
                                            _ = v52
                                        }
                                    case "":
                                        v4 -= (v5 << ((42) & 63))
                                    case "\n\t":
                                        fmt.Println(fi19, v16, v12)
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        fmt.Println(v5, v5, fi19)
                                    case "hello":
                                        fmt.Println(v5)
                                    case "a":
                                        var v55 string = ""
                                        _ = v55
                                    }
                                    var v56 uint = 0
                                    _ = v56
                                }
                            }
                            var v57 uint8 = 100
                            var v58 float32 = ((float32(1e10) * ((0.0 * 0.0) * -2.25)) - (1e10 * (1e10 - 0.0)))
                            v59 := GenS2{F0: true}
                            _ = v57
                            _ = v58
                            _ = v59
                        } else {
                            var v60 int8 = 1
                            _ = v60
                        }
                    default:
                        var v61 uint8 = 127
                        _ = v61
                    }
                    fmt.Println(v4, v1, v3)
                }
                var v62 int = fi18
                v4 = v4
                _ = fi19
                _ = v62
            }
            var v63 float64 = -2.25
            _ = fi18
            _ = v63
        }
    }
    v5 -= v4
    var v64 uint16 = (v16 / ((42) | 1))
    v3 = (3.14159 >= -2.25)
    fmt.Println(v64, v12, v4)
    if ((v3 || v3) || ("a" >= "hello")) {
        v65 := GenS6{F0: 1, F1: -124}
        v1 = v1
        _ = v65
    } else {
        v3 = (!false)
        switch (-(v12)) {
        case 1:
            fmt.Println(v4, v4, v5)
        case 42:
            fmt.Println(v1, v16, v1)
        }
        var v66 int32 = 900
        _ = v66
    }
    v67 := []int{-9223372036854775808, v12, v12, -9223372036854775808}
    var v68 string = ""
    switch v68 {
    case "":
        if ((v5 + (v4 ^ v5)) == (v4 >> ((886) & 63))) {
            var v69 int32 = -1
            var v70 int8 = -128
            var v71 uint64 = 994
            _ = v69
            _ = v70
            _ = v71
        } else {
            if v3 {
                var v72 int32 = 0
                _ = v72
            } else {
                fmt.Println(v5, v64)
                var v73 uint8 = 144
                _ = v73
            }
        }
    default:
        var v74 uint32 = 65535
        _ = v74
    }
    v75 := GenS6{F0: 42, F1: 127}
    v2[1] = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v67
    _ = v68
    _ = v75
}
