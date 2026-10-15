package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 int8; F1 int8 }
func genFn0() uint8 {
    switch 14 {
    case 454:
        var v5 uint = 1
        _ = v5
    case 1:
        var v6 int8 = 77
        _ = v6
    case -1:
        var v7 uint32 = 65535
        _ = v7
    default:
        v8 := []int{}
        _ = v8
    }
    if ("" > "\n\t") {
        switch -2147483648 {
        case 184:
            var v9 uint = 255
            _ = v9
        default:
            switch "foo bar" {
            case "\n\t":
                var v10 bool = false
                _ = v10
            }
        }
        v11 := GenS1{F0: false}
        var v12 uint64 = 315
        _ = v11
        _ = v12
    }
    var v13 int32 = -1
    for fi14 := 0; fi14 < 20; fi14++ {
        v13 -= v13
        var v15 int = 2147483647
        var v16 int32 = (v13 | v13)
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := []int{0}
    _ = v13
    _ = v17
    return 42
}
func genFn1() int {
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 int32 = 227
        if false {
            var v20 bool = (!(0 <= 42))
            switch (fi18 % ((fi18) | 1)) {
            case 0:
                v21 := GenS2{F0: ("hello" + ("a" + (("a" + "") + "foo bar")))}
                _ = v21
            case 9223372036854775807:
                v19 += v19
            case -1:
                var v22 float64 = 0.0
                _ = v22
            }
            _ = v20
        } else {
            var v23 uint8 = genFn0()
            v23 = v23
            v24 := GenS3{F0: "foo bar", F1: v23, F2: (v23 - v23)}
            _ = v23
            _ = v24
        }
        _ = fi18
        _ = v19
    }
    v25 := GenS3{F0: "", F1: 0, F2: 189}
    var v26 uint = 101
    _ = v25
    _ = v26
    return -9223372036854775808
}

func main() {
    v1 := []uint32{}
    var v2 bool = (false || false)
    v3 := []int{1, -2147483648, -1, 4294967295}
    v4 := make(map[uint32]bool)
    switch "\n\t" {
    case "":
        _ = len(v3)
    }
    var v27 uint64 = 63
    if v2 {
        v2 = ("hello" < "\n\t")
    } else {
        if true {
            var v28 int = 128
            if v2 {
                for fi29 := 0; fi29 < 8; fi29++ {
                    var v30 uint = 4294967295
                    _ = fi29
                    _ = v30
                }
                v4[503] = true
            } else {
                var v31 int = -9223372036854775808
                if (0 != 127) {
                    if v2 {
                        for fi32 := 0; fi32 < 6; fi32++ {
                            if v2 {
                                var v33 uint = 258
                                v34 := v4[1]
                                for fi35 := 0; fi35 < 7; fi35++ {
                                    if v2 {
                                        switch fi35 {
                                        case -315:
                                            var v36 bool = (1000 >= 1)
                                            _ = v36
                                        case -9223372036854775808:
                                            v37 := GenS3{F0: "hello", F1: 156, F2: 128}
                                            _ = v37
                                        case 864:
                                            switch -2147483648 {
                                            case -257:
                                                v28 += v31
                                            case 32767:
                                                fmt.Println(v2, fi32)
                                            }
                                        }
                                    }
                                    var v38 int32 = 2147483647
                                    if false {
                                        var v39 uint8 = 0
                                        fmt.Println(v28, v27, v28)
                                        _ = v39
                                    }
                                    _ = fi35
                                    _ = v38
                                }
                                _ = v33
                                _ = v34
                            } else {
                                if false {
                                    v40 := GenS2{F0: "hello"}
                                    _ = v40
                                } else {
                                    for fi41 := 0; fi41 < 20; fi41++ {
                                        if (!v2) {
                                            v42 := GenS3{F0: "a", F1: 2, F2: 128}
                                            var v43 int16 = 1000
                                            _ = v42
                                            _ = v43
                                        }
                                        var v44 int64 = 544
                                        var v45 int64 = v44
                                        _ = fi41
                                        _ = v44
                                        _ = v45
                                    }
                                }
                                v2 = ((929 >= 2) && (false && (!(1e10 > 1e10))))
                            }
                            _ = fi32
                        }
                    } else {
                        if v2 {
                            for fi46 := 0; fi46 < 6; fi46++ {
                                for fi47 := 0; fi47 < 8; fi47++ {
                                    v2 = v2
                                    var v48 int32 = 0
                                    var v49 float32 = 1e10
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                                if (1 < 127) {
                                    v2 = (1 == 160)
                                } else {
                                    fmt.Println(v2, v27, v2)
                                    v27 += v27
                                }
                                _ = fi46
                            }
                            if (v27 < (v27 + v27)) {
                                fmt.Println(v28, v31, v27)
                                for fi50 := 0; fi50 < 1; fi50++ {
                                    v27 = v27
                                    var v51 string = ""
                                    var v52 int = (fi50 * (v31 / ((v31) | 1)))
                                    _ = fi50
                                    _ = v51
                                    _ = v52
                                }
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = 65535
                                }
                                if (true || v2) {
                                    v53 := GenS3{F0: "foo bar", F1: 254, F2: 255}
                                    _ = v53
                                }
                                v54 := GenS4{F0: v28}
                                _ = v54
                            }
                        } else {
                            for fi55 := 0; fi55 < 9; fi55++ {
                                for fi56 := 0; fi56 < 7; fi56++ {
                                    var v57 string = ("a" + "foo bar")
                                    _ = fi56
                                    _ = v57
                                }
                                _ = fi55
                            }
                            switch (("hello" + ("" + "\n\t")) + "") {
                            case "a":
                                var v58 uint = 4294967295
                                _ = v58
                            case "foo bar":
                                v28 = (v28 & genFn1())
                            case "\n\t":
                                var v59 float64 = 1e10
                                _ = v59
                            }
                        }
                        v27 = (v27 / (((v27 * v27)) | 1))
                        _ = len(v1)
                    }
                }
                if true {
                    var v60 string = "hello"
                    v2 = (950 <= 65535)
                    var v61 bool = (1000 > 1)
                    _ = v60
                    _ = v61
                } else {
                    v4[2] = (!(v27 <= 4294967295))
                    var v62 float64 = (float64(1.5) * 3.14159)
                    _ = v62
                }
                _ = v31
            }
            v27 = 1
            _ = v28
        }
    }
    if (1 == 1) {
        switch 2147483647 {
        case 32767:
            v63 := GenS1{F0: (!(2 >= 0))}
            _ = v63
        case 32768:
            v27 += 0
        case 9223372036854775807:
            fmt.Println(v2)
        default:
            var v64 uint16 = 168
            _ = v64
        }
    }
    fmt.Println(v2, v27)
    var v65 int16 = 137
    for fi66 := 0; fi66 < 1; fi66++ {
        v27 -= v27
        v67 := GenS4{F0: fi66}
        v68 := GenS5{F0: 44, F1: 1}
        _ = fi66
        _ = v67
        _ = v68
    }
    v69 := GenS6{F0: 42, F1: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v65
    _ = v69
}
