package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int; F2 int64 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64; F1 int64; F2 uint }
type GenS9 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS10 struct { F0 int; F1 int }
type GenS11 struct { F0 int; F1 int64 }
func genFn0() uint16 {
    if (-1 != 31) {
        if ((465 != 2) && false) {
            var v6 int32 = 0
            _ = v6
        }
        var v7 uint8 = 42
        _ = v7
    }
    var v8 float32 = -2.25
    var v9 int = -9223372036854775808
    v9 += -1
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS1{F0: -742}
        v12 := GenS2{F0: 255, F1: fi10, F2: 1}
        v13 := []int{(^(fi10))}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    _ = v8
    _ = v9
    return 65535
}
func genFn1() string {
    switch ("foo bar" + "\n\t") {
    case "":
        if ("" == "") {
            if true {
                var v14 uint = 2
                v15 := GenS3{F0: 2147483647, F1: 42}
                _ = v14
                _ = v15
            } else {
                var v16 uint16 = 42
                if ("foo bar" > "") {
                    v16 += (^(v16))
                } else {
                    if (!true) {
                        var v17 uint32 = 2147483648
                        v18 := []int{-2147483648, 893, -780, -2147483648}
                        _ = v17
                        _ = v18
                    }
                    if true {
                        v19 := []int{32767, -9223372036854775808}
                        v16 = 942
                        v20 := GenS4{F0: 2147483648, F1: 2}
                        _ = v19
                        _ = v20
                    }
                }
                switch "\n\t" {
                case "hello":
                    var v21 int16 = -32768
                    _ = v21
                case "foo bar":
                    var v22 int64 = 1
                    _ = v22
                case "\n\t":
                    if false {
                        v23 := []int{-9223372036854775808, 4294967295, 70, -9223372036854775808}
                        var v24 bool = (1 < -67)
                        v16 += genFn0()
                        _ = v23
                        _ = v24
                    }
                }
                _ = v16
            }
            var v25 uint16 = 1
            v26 := GenS5{F0: 4294967295, F1: -1, F2: 42}
            _ = v25
            _ = v26
        } else {
            v27 := GenS3{F0: 868, F1: -400}
            var v28 uint64 = 0
            if (-1 == 42) {
                fmt.Println(v28)
            } else {
                v28 += (v28 >> ((361) & 63))
                for fi29 := 0; fi29 < 16; fi29++ {
                    var v30 int16 = 413
                    v30 = (v30 % (((^(v30))) | 1))
                    _ = fi29
                    _ = v30
                }
            }
            _ = v27
            _ = v28
        }
    default:
        var v31 uint32 = 0
        _ = v31
    }
    var v32 uint8 = 1
    if (127 > 37) {
        v33 := []int{1, 32767}
        for fi34 := 0; fi34 < 15; fi34++ {
            v32 -= v32
            if ((fi34 | -2147483648) != fi34) {
                for fi35 := 0; fi35 < 5; fi35++ {
                    v32 = v32
                    v32 -= v32
                    _ = fi35
                }
                var v36 uint8 = v32
                fmt.Println(v32, v32, v32)
                _ = v36
            }
            _ = fi34
        }
        var v37 bool = true
        _ = v33
        _ = v37
    } else {
        if false {
            for fi38 := 0; fi38 < 5; fi38++ {
                for fi39 := 0; fi39 < 10; fi39++ {
                    if true {
                        v40 := []int{(fi39 % (((fi38 << ((248) & 63))) | 1)), (fi38 << ((1) & 63)), (fi39 | -2147483648)}
                        v32 = v32
                        _ = v40
                    }
                    v32 -= 42
                    v41 := []int{32767}
                    _ = fi39
                    _ = v41
                }
                _ = fi38
            }
        } else {
            v42 := GenS6{F0: 42}
            _ = v42
        }
        if false {
            if true {
                var v43 uint32 = 408
                v44 := GenS6{F0: 276}
                switch -9223372036854775808 {
                case 1:
                    v43 = v43
                case 265:
                    switch 32768 {
                    case -612:
                        var v45 int = 32768
                        _ = v45
                    case 2147483647:
                        v43 += v43
                    default:
                        fmt.Println(v43)
                    }
                case 4294967295:
                    var v46 int64 = 9223372036854775807
                    _ = v46
                default:
                    v47 := GenS4{F0: (v43 / ((v43) | 1)), F1: 124}
                    _ = v47
                }
                _ = v43
                _ = v44
            }
        }
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v32)
    default:
        v32 = 0
    }
    _ = v32
    return "\n\t"
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = false
    var v3 float64 = float64(1.5)
    v4 := make(map[uint32]int)
    var v5 uint64 = 42
    var v48 int8 = 123
    var v49 uint = 65535
    fmt.Println(v2, v5, v5)
    for fi50 := 0; fi50 < 20; fi50++ {
        v51 := GenS2{F0: 127, F1: fi50, F2: 650}
        switch "hello" {
        case "foo bar":
            switch "\n\t" {
            case "foo bar":
                var v52 int64 = 793
                _ = v52
            case "\n\t":
                if (1 < 32767) {
                    fmt.Println(v48)
                    v2 = (!(32 == 127))
                    var v53 string = ("a" + (("a" + "\n\t") + ("a" + "a")))
                    _ = v53
                } else {
                    v54 := GenS3{F0: 2147483647, F1: 1}
                    _ = v54
                }
            default:
                v55 := GenS6{F0: genFn0()}
                _ = v55
            }
        }
        if (4294967295 <= 42) {
            if v2 {
                for fi56 := 0; fi56 < 8; fi56++ {
                    var v57 float32 = 0.0
                    switch (fi50 + fi50) {
                    case 42:
                        fmt.Println(v49, fi50)
                    case 128:
                        var v58 uint16 = genFn0()
                        _ = v58
                    case 2147483647:
                        switch (-(fi56)) {
                        case 0:
                            if v2 {
                                v59 := GenS5{F0: 4294967295, F1: 0, F2: v5}
                                v2 = v2
                                fmt.Println(v3, v5)
                                _ = v59
                            }
                        }
                        var v60 uint = (v49 / (((v49 ^ 42)) | 1))
                        v61 := []int{fi56, 9223372036854775807}
                        v62 := GenS7{F0: 61}
                        switch (("\n\t" + "a") + "a") {
                        case "":
                            var v63 int = (fi50 % (((^(fi50))) | 1))
                            _ = v63
                        case "foo bar":
                            v64 := v4[0]
                            _ = v64
                        case "hello":
                            v5 += (v5 / ((v5) | 1))
                        }
                        _ = v60
                        _ = v61
                        _ = v62
                    }
                    _ = fi56
                    _ = v57
                }
                var v65 int16 = -1
                _ = v65
            } else {
                var v66 string = "a"
                if (((3.14159 / 1.5) <= 0.1) || (!v2)) {
                    var v67 int8 = v48
                    delete(v4, 65535)
                    var v68 uint16 = genFn0()
                    _ = v67
                    _ = v68
                } else {
                    v49 += v49
                }
                switch "hello" {
                case "":
                    if (578 > 507) {
                        v5 -= 18446744073709551615
                    }
                case "a":
                    v48 = v48
                }
                _ = v66
            }
            fmt.Println(v3, v5, v3)
        } else {
            fmt.Println(v48, fi50, v49)
            fmt.Println(v5)
        }
        _ = fi50
        _ = v51
    }
    for fi69 := 0; fi69 < 13; fi69++ {
        v70 := GenS8{F0: -9223372036854775808, F1: -9223372036854775808, F2: (v49 * v49)}
        fmt.Println(v5, v5)
        v71 := GenS9{F0: uint16(v49), F1: -9223372036854775808, F2: (v5 | v5)}
        _ = fi69
        _ = v70
        _ = v71
    }
    v72 := GenS10{F0: 591, F1: -965}
    switch -1 {
    case 32767:
        v73 := GenS11{F0: 797, F1: -836}
        _ = v73
    case 0:
        var v74 uint16 = 381
        _ = v74
    }
    v5 = (^(v5))
    v75 := []int{-300, 32768}
    var v76 int32 = 1
    v75 = append(v75, -305)
    fmt.Println(v48, v48)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v75:", len(v75), v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v72
    _ = v75
    _ = v76
}
