package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 uint32; F2 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32; F1 int64; F2 uint16 }
type GenS8 struct { F0 uint8; F1 int8; F2 bool }
func genFn0() bool {
    var v4 float32 = 0.1
    v5 := GenS1{F0: 0, F1: 42}
    var v6 uint8 = 128
    if (2 < 42) {
        v7 := GenS2{F0: 88}
        v8 := GenS3{F0: 0}
        _ = v7
        _ = v8
    }
    _ = v4
    _ = v5
    _ = v6
    return false
}

func main() {
    var v1 int8 = 0
    var v2 int64 = 593
    var v3 uint16 = 521
    var v9 float32 = 1e10
    var v10 int64 = (v2 % (((^(v2))) | 1))
    var v11 float64 = 0.1
    if false {
        fmt.Println(v2)
        for fi12 := 0; fi12 < 8; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 int8 = v1
                var v15 int16 = 359
                var v16 uint = 577
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            v17 := []int{(-(fi12))}
            for fi18 := 0; fi18 < 9; fi18++ {
                switch (^(fi12)) {
                case -458:
                    var v19 float32 = 0.0
                    _ = v19
                case 42:
                    v20 := GenS4{F0: 32767, F1: 2147483648, F2: ""}
                    _ = v20
                }
                fmt.Println(v1, fi12, fi12)
                _ = fi18
            }
            _ = fi12
            _ = v17
        }
        v9 += 3.14159
    } else {
        var v21 uint16 = v3
        fmt.Println(v11, v21)
        if (false || true) {
            if (65535 != 1) {
                for fi22 := 0; fi22 < 2; fi22++ {
                    if false {
                        v23 := []int{fi22, (fi22 * fi22), fi22, fi22}
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                        _ = v23
                    }
                    _ = fi22
                }
                v11 = -2.25
                fmt.Println(v11, v10, v2)
            }
        }
        _ = v21
    }
    var v24 uint64 = 42
    if true {
        switch "\n\t" {
        case "\n\t":
            var v25 int8 = v1
            _ = v25
        case "a":
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 uint16 = v3
                v28 := []int{fi26, (fi26 % ((fi26) | 1)), 0, fi26}
                _ = fi26
                _ = v27
                _ = v28
            }
        }
        v9 = v9
    } else {
        v29 := GenS5{F0: 4294967295}
        v1 -= v1
        _ = v29
    }
    v30 := []int{-9223372036854775808, 128, -45}
    v30 = append(v30, 9223372036854775807)
    fmt.Println(v1, v1)
    if false {
        var v31 float64 = v11
        var v32 string = ("" + "\n\t")
        var v33 string = "hello"
        _ = v31
        _ = v32
        _ = v33
    } else {
        switch ("foo bar" + (("a" + ("hello" + "\n\t")) + (("foo bar" + (("a" + "\n\t") + ("a" + ("" + ("foo bar" + "hello"))))) + ("foo bar" + "hello")))) {
        case "\n\t":
            if true {
                switch "hello" {
                case "":
                    var v34 int16 = -56
                    _ = v34
                }
                var v35 uint = 163
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 5; fi36++ {
            var v37 bool = false
            var v38 int32 = -1
            _ = fi36
            _ = v37
            _ = v38
        }
        v39 := GenS6{F0: (v1 + (v1 / ((v1) | 1)))}
        _ = v39
    }
    v40 := GenS4{F0: 1, F1: 798, F2: (("\n\t" + "") + ("hello" + "a"))}
    v9 += 1e10
    v1 = 22
    var v41 int64 = -785
    var v42 int64 = v2
    fmt.Println(v41, v2)
    for fi43 := 0; fi43 < 11; fi43++ {
        if true {
            var v44 uint64 = (v24 + v24)
            for fi45 := 0; fi45 < 4; fi45++ {
                if (-1 > (^(v1))) {
                    var v46 uint = 0
                    var v47 int32 = 1
                    v48 := GenS4{F0: -1, F1: uint32(v44), F2: "\n\t"}
                    v49 := GenS6{F0: v1}
                    var v50 int8 = v1
                    _ = v46
                    _ = v47
                    _ = v48
                    _ = v49
                    _ = v50
                }
                if true {
                    switch "hello" {
                    case "foo bar":
                        _ = len(v30)
                        var v51 float32 = v9
                        v52 := GenS7{F0: 42, F1: v2, F2: (v3 ^ 1)}
                        _ = v51
                        _ = v52
                    case "a":
                        fmt.Println(v3)
                    default:
                        v42 = v10
                    }
                    switch 42 {
                    case -9223372036854775808:
                        if true {
                            v3 -= 325
                            var v53 float64 = -2.25
                            fmt.Println(fi43)
                            _ = v53
                        } else {
                            var v54 int8 = int8(v3)
                            _ = v54
                        }
                    case -135:
                        var v55 bool = false
                        _ = v55
                    default:
                        if (!(-349 < 14)) {
                            fmt.Println(v1, v1)
                        }
                    }
                } else {
                    fmt.Println(v44)
                }
                if len(v30) > 0 {
                    v30[0] = (fi45 | 0)
                }
                _ = fi45
            }
            _ = v44
        }
        v10 -= (v41 % ((v2) | 1))
        v1 = 1
        _ = fi43
    }
    if (v10 > 0) {
        var v56 int64 = v42
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            v57 := GenS7{F0: 1, F1: v56, F2: v3}
            _ = v57
        }
        if (-1 >= 2147483647) {
            var v58 int = 1
            _ = v58
        } else {
            if false {
                var v59 uint32 = 4294967295
                _ = v59
            } else {
                fmt.Println(v3, v2, v56)
                switch -18 {
                case 127:
                    var v60 uint32 = 42
                    _ = v60
                case 128:
                    switch 32767 {
                    case 32767:
                        v1 = -128
                    case -48:
                        switch "foo bar" {
                        case "":
                            for fi61 := 0; fi61 < 5; fi61++ {
                                var v62 int8 = v1
                                var v63 int32 = -396
                                fmt.Println(v11, v41)
                                _ = fi61
                                _ = v62
                                _ = v63
                            }
                        case "foo bar":
                            v30 = append(v30, 42)
                        }
                    case 128:
                        if false {
                            switch (("" + "") + ("foo bar" + "hello")) {
                            case "\n\t":
                                v64 := GenS7{F0: 605, F1: v10, F2: 65535}
                                _ = v64
                            }
                            v65 := GenS6{F0: v1}
                            _ = v65
                        }
                    default:
                        v66 := GenS3{F0: 968}
                        _ = v66
                    }
                default:
                    v30 = append(v30, 4294967295)
                }
                v67 := GenS8{F0: 2, F1: -1, F2: (0 < 1)}
                _ = v67
            }
            var v68 uint32 = 0
            _ = v68
        }
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v40
    _ = v41
    _ = v42
    _ = v9
}
