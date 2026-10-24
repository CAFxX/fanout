package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 uint8; F1 int32 }
func genFn0(p0 int32, p1 uint64, p2 int64) int32 {
    v7 := GenS1{F0: 0}
    var v8 uint16 = 0
    var v9 int = 2147483647
    if false {
        for fi10 := 0; fi10 < 10; fi10++ {
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS1{F0: 659}
                p2 = (-(p2))
                switch fi10 {
                case 928:
                    if (uint32(p1) != 156) {
                        var v13 int16 = -1
                        var v14 int8 = -8
                        v15 := []int{}
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        if false {
                            v8 += (v8 / (((v8 % ((2) | 1))) | 1))
                        }
                        fmt.Println(fi10, p2, p2)
                        v16 := []int{fi10, -796, fi10}
                        _ = v16
                    }
                default:
                    v17 := []int{fi10, fi10}
                    _ = v17
                }
                _ = fi11
                _ = v12
            }
            v18 := []int{}
            v9 = fi10
            _ = fi10
            _ = v18
        }
        var v19 int = 32768
        _ = v19
    } else {
        p1 += 4294967295
        var v20 uint64 = (p1 + p1)
        v20 = (p1 << ((745) & 63))
        _ = v20
    }
    var v21 float32 = (1e10 - ((0.1 - (0.1 + 3.14159)) * -2.25))
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v7
    _ = v8
    _ = v9
    return (-(p0))
}
func genFn1(p0 int8, p1 int32, p2 int16) bool {
    var v22 int64 = 0
    v23 := GenS1{F0: 997}
    var v24 int16 = p2
    var v25 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    return true
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 2147483648
    var v3 int16 = -14
    var v4 uint64 = 42
    var v5 uint16 = 2
    v6 := make(map[string]string)
    for fi26 := 0; fi26 < 3; fi26++ {
        switch ("" + "foo bar") {
        case "hello":
            if true {
                v2 += (v2 + (v2 * v2))
                var v27 uint64 = 435
                v2 += (v2 % ((980) | 1))
                _ = v27
            } else {
                fmt.Println(v3, fi26, v2)
                for fi28 := 0; fi28 < 1; fi28++ {
                    switch "\n\t" {
                    case "hello":
                        if true {
                            switch "foo bar" {
                            case "":
                                var v29 uint = uint(v5)
                                _ = v29
                            case "hello":
                                v5 += (v5 >> ((624) & 63))
                            case "\n\t":
                                fmt.Println(fi26, v5)
                            default:
                                if (false && true) {
                                    v5 = (^(v5))
                                    var v30 uint64 = 18446744073709551615
                                    if false {
                                        switch (fi28 - -9223372036854775808) {
                                        case 42:
                                            v31 := []int{528, (^(fi28)), fi26, (fi26 * fi28)}
                                            _ = v31
                                        case 32768:
                                            var v32 int64 = -9223372036854775808
                                            switch (fi28 << ((65535) & 63)) {
                                            case 32767:
                                                v3 = (v3 & (-(v3)))
                                            }
                                            var v33 float64 = -2.25
                                            _ = v32
                                            _ = v33
                                        case -980:
                                            v5 = (v5 << ((42) & 63))
                                        default:
                                            fmt.Println(fi28)
                                        }
                                        if (true && false) {
                                            fmt.Println(fi28, v3, fi26)
                                            v30 = 18446744073709551615
                                        } else {
                                            var v34 int32 = -1
                                            var v35 float64 = (float64(-2.25) - (3.14159 + 1.5))
                                            _ = v34
                                            _ = v35
                                        }
                                    } else {
                                        v3 -= (v3 / ((v3) | 1))
                                    }
                                    _ = v30
                                }
                            }
                            switch (-(fi26)) {
                            case 127:
                                switch (fi28 + fi26) {
                                case -1:
                                    v36 := GenS2{F0: "foo bar", F1: genFn0(0, (v4 ^ v4), 0), F2: 127}
                                    _ = v36
                                case 32768:
                                    var v37 string = "foo bar"
                                    _ = v37
                                }
                            case -203:
                                fmt.Println(v3)
                                fmt.Println(fi28, v4)
                            case 1:
                                v38 := GenS2{F0: "foo bar", F1: genFn0(genFn0(0, 42, -9223372036854775808), v4, -1), F2: 127}
                                v2 -= v2
                                _ = v38
                            }
                            v39 := []int{(fi26 ^ fi26), fi26, (fi28 - fi26)}
                            _ = v39
                        }
                    case "a":
                        v40 := GenS1{F0: 862}
                        _ = v40
                    case "\n\t":
                        fmt.Println(fi28, v3)
                    }
                    fmt.Println(v3)
                    _ = fi28
                }
                var v41 int32 = -344
                _ = v41
            }
        case "a":
            var v42 string = "hello"
            _ = v42
        default:
            if ((^(v4)) < (v4 ^ v4)) {
                if (false && false) {
                    for fi43 := 0; fi43 < 5; fi43++ {
                        if (!(((-(v3)) < (v3 | (v3 << ((65535) & 63)))) || (1.5 > float32(0.1)))) {
                            switch fi43 {
                            case 371:
                                var v44 float64 = 0.1
                                _ = v44
                            }
                            v4 -= (^(v4))
                            var v45 int16 = (v3 & (v3 << ((1) & 63)))
                            var v46 float32 = (0.0 + 0.1)
                            if (127 <= 0) {
                                var v47 int16 = (-(v45))
                                v48 := GenS2{F0: (("foo bar" + "\n\t") + ("" + "foo bar")), F1: 1, F2: 0}
                                var v49 bool = ((18446744073709551615 < 4294967295) || (!true))
                                _ = v47
                                _ = v48
                                _ = v49
                            } else {
                                var v50 float32 = v46
                                fmt.Println(v5, v2, fi43)
                                fmt.Println(v3, fi26)
                                _ = v50
                            }
                            _ = v45
                            _ = v46
                        } else {
                            v5 = (v5 + 966)
                            var v51 int64 = -1
                            _ = v51
                        }
                        _ = fi43
                    }
                    for fi52 := 0; fi52 < 16; fi52++ {
                        if true {
                            var v53 uint16 = 95
                            _ = v53
                        }
                        if (((-(v3)) == v3) || true) {
                            if false {
                                fmt.Println(v5)
                                v5 += v5
                                v4 -= 1
                            } else {
                                v54 := []int{9223372036854775807, fi26, (fi26 & (^(fi52)))}
                                v55 := GenS3{F0: (v4 % ((4294967295) | 1))}
                                if ((("" + "\n\t") + (("" + "\n\t") + "\n\t")) != "") {
                                    fmt.Println(v3, fi52, v5)
                                    var v56 uint8 = uint8(v5)
                                    var v57 int = (-(fi52))
                                    _ = v56
                                    _ = v57
                                } else {
                                    v3 = int16(v5)
                                }
                                _ = v54
                                _ = v55
                            }
                        } else {
                            v58 := GenS3{F0: v4}
                            _ = v58
                        }
                        _ = fi52
                    }
                    v59 := GenS4{F0: (v2 | 2147483648)}
                    _ = v59
                }
                if ((("foo bar" + "foo bar") > ("a" + "\n\t")) || false) {
                    for fi60 := 0; fi60 < 4; fi60++ {
                        var v61 uint32 = v2
                        v3 -= v3
                        fmt.Println(v4, fi26)
                        _ = fi60
                        _ = v61
                    }
                    v5 += uint16(v4)
                    switch fi26 {
                    case 127:
                        switch "\n\t" {
                        case "a":
                            _ = len(v6)
                        case "":
                            switch ("\n\t" + "a") {
                            case "a":
                                for fi62 := 0; fi62 < 1; fi62++ {
                                    if (fi62 >= fi26) {
                                        v63 := []int{(fi62 - (-(fi26))), (-(fi62)), 9223372036854775807}
                                        v64 := GenS3{F0: v4}
                                        var v65 string = ("a" + "foo bar")
                                        v66 := GenS5{F0: ("foo bar" + "foo bar"), F1: 176}
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    }
                                    var v67 int16 = (v3 - v3)
                                    _ = fi62
                                    _ = v67
                                }
                            case "":
                                var v68 uint16 = (^(v5))
                                _ = v68
                            }
                        default:
                            var v69 float32 = (3.14159 / 0.1)
                            _ = v69
                        }
                    default:
                        v70 := GenS5{F0: "", F1: 1}
                        _ = v70
                    }
                } else {
                    switch fi26 {
                    case 128:
                        delete(v6, ("hello" + "\n\t"))
                    case 9223372036854775807:
                        var v71 int32 = -531
                        _ = v71
                    case 32768:
                        v72 := v6["foo bar"]
                        _ = v72
                    default:
                        if false {
                            if (79 != 320) {
                                var v73 int64 = 0
                                _ = v73
                            }
                        } else {
                            var v74 int = (fi26 ^ int(v5))
                            v75 := GenS1{F0: 4294967295}
                            _ = v74
                            _ = v75
                        }
                    }
                }
                var v76 uint16 = v5
                _ = v76
            } else {
                fmt.Println(v4)
            }
        }
        var v77 int64 = 479
        _ = fi26
        _ = v77
    }
    fmt.Println(v5)
    for fi78 := 0; fi78 < 7; fi78++ {
        v79 := GenS4{F0: uint32(v4)}
        _ = fi78
        _ = v79
    }
    v80 := GenS6{F0: 246, F1: 0}
    switch -938 {
    case 0:
        for fi81 := 0; fi81 < 9; fi81++ {
            var v82 int16 = 0
            fmt.Println(v82)
            _ = fi81
            _ = v82
        }
    case 32768:
        fmt.Println(v3, v4)
    }
    v83 := GenS5{F0: "", F1: uint8(v2)}
    v84 := GenS4{F0: (v2 * v2)}
    for fi85 := 0; fi85 < 6; fi85++ {
        var v86 string = ("" + "foo bar")
        v3 += v3
        _ = fi85
        _ = v86
    }
    v87 := []int{9223372036854775807}
    fmt.Println(v3)
    v88 := GenS3{F0: (v4 % (((v4 ^ v4)) | 1))}
    v89 := GenS3{F0: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v87:", len(v87), v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v80
    _ = v83
    _ = v84
    _ = v87
    _ = v88
    _ = v89
}
