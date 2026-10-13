package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint8; F1 uint64 }
func genFn0() string {
    v4 := GenS1{F0: 42, F1: 0}
    v5 := GenS1{F0: 645, F1: 0}
    v6 := GenS1{F0: -2147483648, F1: 42}
    if (!true) {
        switch ("a" + "\n\t") {
        case "a":
            var v7 uint32 = 65535
            _ = v7
        case "":
            v8 := []int{-904}
            _ = v8
        case "hello":
            for fi9 := 0; fi9 < 10; fi9++ {
                for fi10 := 0; fi10 < 20; fi10++ {
                    fmt.Println(fi10, fi10)
                    _ = fi10
                }
                _ = fi9
            }
        }
        var v11 float32 = 1.5
        for fi12 := 0; fi12 < 14; fi12++ {
            v11 = float32(1.5)
            fmt.Println(v11, fi12)
            var v13 uint8 = 255
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    var v14 float32 = 0.1
    _ = v14
    _ = v4
    _ = v5
    _ = v6
    return (((((("hello" + ((("a" + "a") + "foo bar") + "hello")) + "\n\t") + ("a" + ("hello" + "a"))) + "foo bar") + "foo bar") + "a")
}

func main() {
    v1 := []int{9223372036854775807}
    v2 := make(map[string]int)
    var v3 float32 = 0.1
    v3 = 1.5
    var v15 int8 = 6
    for fi16 := 0; fi16 < 20; fi16++ {
        v15 = v15
        for fi17 := 0; fi17 < 11; fi17++ {
            delete(v2, ("" + ""))
            var v18 int8 = v15
            _ = fi17
            _ = v18
        }
        _ = len(v2)
        _ = fi16
    }
    v15 -= v15
    v3 = 0.0
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(v3, v15)
        for fi20 := 0; fi20 < 14; fi20++ {
            _ = len(v1)
            _ = fi20
        }
        var v21 uint = 18446744073709551615
        _ = fi19
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        v15 += 25
        v23 := GenS1{F0: 42, F1: 2}
        v3 += float32(3.14159)
        _ = fi22
        _ = v23
    }
    if ((1e10 - float64(v3)) >= (((((3.14159 / 1.0) * 0.0) / -2.25) + 1e10) - 1.5)) {
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
        var v24 int32 = 2147483647
        _ = v24
    }
    fmt.Println(v3, v3)
    if true {
        for fi25 := 0; fi25 < 20; fi25++ {
            v15 = v15
            _ = fi25
        }
        v26 := GenS1{F0: -2147483648, F1: 0}
        for fi27 := 0; fi27 < 16; fi27++ {
            for fi28 := 0; fi28 < 20; fi28++ {
                v15 = -128
                var v29 int16 = 445
                _ = fi28
                _ = v29
            }
            var v30 int16 = -887
            if ((0 == 2) || true) {
                var v31 uint16 = 42
                _ = v31
            } else {
                for fi32 := 0; fi32 < 14; fi32++ {
                    var v33 int32 = -2147483648
                    if (2 < 128) {
                        var v34 bool = true
                        v3 -= v3
                        _ = v34
                    } else {
                        v35 := GenS1{F0: v33, F1: 612}
                        fmt.Println(fi32, v33)
                        _ = v35
                    }
                    switch "a" {
                    case "\n\t":
                        fmt.Println(fi27, v30)
                    case "hello":
                        _ = len(v2)
                    case "":
                        v30 = v30
                    default:
                        if (((v3 - (v3 - float32(-2.25))) + (1e10 + (v3 / 1.0))) == (3.14159 + 0.0)) {
                            if false {
                                var v36 int32 = (v33 % ((42) | 1))
                                _ = v36
                            } else {
                                v15 += (v15 >> ((1) & 63))
                                var v37 float32 = v3
                                var v38 int64 = 664
                                _ = v37
                                _ = v38
                            }
                            var v39 string = "\n\t"
                            var v40 uint16 = 883
                            if true {
                                v30 += (v30 << ((255) & 63))
                                if false {
                                    fmt.Println(v40, v15)
                                } else {
                                    _ = len(v1)
                                }
                            } else {
                                if ("" > "a") {
                                    fmt.Println(fi27, v3, v15)
                                } else {
                                    v40 -= (v40 << ((65535) & 63))
                                    var v41 uint = 287
                                    _ = v41
                                }
                                fmt.Println(v33, v15)
                            }
                            switch fi32 {
                            case -9223372036854775808:
                                v3 -= 0.0
                            case 127:
                                v2[(v39 + "")] = fi32
                            default:
                                if (1e10 < v3) {
                                    switch fi32 {
                                    case 2147483647:
                                        if false {
                                            if (((-2.25 * 0.0) / 1.0) != 3.14159) {
                                                v42 := GenS2{F0: 42, F1: 255}
                                                v1 = append(v1, fi32)
                                                v43 := v2[(v39 + v39)]
                                                v40 = v40
                                                v1 = append(v1, (fi32 % (((-(fi27))) | 1)))
                                                _ = v42
                                                _ = v43
                                            }
                                            v3 += v3
                                            v40 += v40
                                        }
                                    default:
                                        var v44 uint64 = 1
                                        _ = v44
                                    }
                                    switch (fi27 - fi32) {
                                    case 128:
                                        if (976 >= 1) {
                                            switch (v39 + "foo bar") {
                                            case "foo bar":
                                                var v45 int16 = (v30 - (-(v30)))
                                                _ = v45
                                            default:
                                                v40 += (v40 >> ((18446744073709551615) & 63))
                                            }
                                        }
                                        var v46 uint8 = 185
                                        _ = v46
                                    case 2147483647:
                                        _ = len(v2)
                                    case 32767:
                                        switch ("hello" + ((v39 + "") + "a")) {
                                        case "\n\t":
                                            v47 := GenS3{F0: (fi32 ^ 42)}
                                            _ = v47
                                        }
                                    }
                                    v30 -= (^(v30))
                                } else {
                                    if (true || true) {
                                        var v48 bool = (498 > 18446744073709551615)
                                        if (42 >= 124) {
                                            v48 = (!false)
                                            var v49 bool = (v48 || v48)
                                            _ = v49
                                        }
                                        v33 += (-(v33))
                                        _ = v48
                                    } else {
                                        var v50 int = fi32
                                        var v51 int = (^(fi27))
                                        v30 -= 42
                                        v3 -= v3
                                        _ = v50
                                        _ = v51
                                    }
                                }
                            }
                            var v52 bool = false
                            _ = v39
                            _ = v40
                            _ = v52
                        } else {
                            if (v33 != (v33 << ((0) & 63))) {
                                switch "hello" {
                                case "a":
                                    v53 := GenS1{F0: v33, F1: 18446744073709551615}
                                    _ = v53
                                case "":
                                    v33 = v33
                                }
                                switch int(v33) {
                                case 4294967295:
                                    if false {
                                        if true {
                                            delete(v2, ("" + "hello"))
                                            if false {
                                                fmt.Println(fi27, v15)
                                                _ = len(v2)
                                                delete(v2, ("a" + "hello"))
                                                var v54 int64 = -1
                                                _ = v54
                                            } else {
                                                fmt.Println(v33, fi27, v30)
                                                v55 := GenS2{F0: 139, F1: 229}
                                                fmt.Println(v30, v33)
                                                _ = v55
                                            }
                                            switch (("hello" + ((("" + "hello") + "hello") + "")) + "") {
                                            case "\n\t":
                                                var v56 string = ""
                                                _ = v56
                                            case "foo bar":
                                                v57 := GenS4{F0: 1000}
                                                _ = v57
                                            case "hello":
                                                var v58 string = "\n\t"
                                                _ = v58
                                            default:
                                                v59 := GenS5{F0: 916, F1: ("\n\t" + "a")}
                                                _ = v59
                                            }
                                        } else {
                                            var v60 uint16 = 918
                                            _ = v60
                                        }
                                        var v61 uint32 = 2147483648
                                        _ = v61
                                    }
                                case 105:
                                    var v62 uint16 = 0
                                    _ = v62
                                case 1:
                                    var v63 float64 = 0.0
                                    _ = v63
                                default:
                                    v64 := GenS6{F0: 59, F1: 2}
                                    _ = v64
                                }
                            }
                        }
                    }
                    var v65 uint = 1
                    _ = fi32
                    _ = v33
                    _ = v65
                }
                for fi66 := 0; fi66 < 13; fi66++ {
                    switch "hello" {
                    case "a":
                        v3 = v3
                    case "":
                        v15 += 0
                    default:
                        if (1 <= 42) {
                            switch ("hello" + "") {
                            case "foo bar":
                                fmt.Println(fi27)
                            default:
                                var v67 uint16 = 65535
                                _ = v67
                            }
                        } else {
                            var v68 bool = (1 > -2147483648)
                            _ = v68
                        }
                        var v69 uint64 = 18446744073709551615
                        var v70 uint64 = v69
                        v69 = v70
                        v30 += (v30 + (v30 >> ((0) & 63)))
                        var v71 uint16 = 1
                        _ = v69
                        _ = v70
                        _ = v71
                    }
                    _ = fi66
                }
            }
            _ = fi27
            _ = v30
        }
        _ = v26
    } else {
        switch "foo bar" {
        case "":
            v72 := v2["hello"]
            _ = v72
        case "hello":
            fmt.Println(v15)
        default:
            v15 = 100
        }
        v15 -= 1
        var v73 int64 = -1
        _ = v73
    }
    var v74 int = 42
    for fi75 := 0; fi75 < 10; fi75++ {
        var v76 uint = 555
        _ = fi75
        _ = v76
    }
    fmt.Println(v3)
    v15 = v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v74
}
