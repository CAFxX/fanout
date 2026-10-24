package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int32; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 string }
func genFn0(p0 uint32, p1 int64, p2 uint16) uint8 {
    var v7 uint8 = 127
    var v8 int16 = 0
    v7 -= (v7 | (v7 % (((v7 << ((0) & 63))) | 1)))
    if (uint64(p2) >= 4294967295) {
        var v9 float32 = 0.1
        var v10 string = ("a" + "")
        var v11 string = ("hello" + "\n\t")
        _ = v10
        _ = v11
        _ = v9
    } else {
        var v12 int32 = -2147483648
        v8 = v8
        _ = v12
    }
    if false {
        v8 = (-(v8))
        var v13 uint16 = 32767
        if true {
            for fi14 := 0; fi14 < 16; fi14++ {
                for fi15 := 0; fi15 < 3; fi15++ {
                    v16 := GenS1{F0: "foo bar"}
                    if (-2.25 != float64(-2.25)) {
                        v17 := GenS1{F0: "a"}
                        _ = v17
                    }
                    var v18 uint = uint(p0)
                    if true {
                        if false {
                            var v19 uint64 = 963
                            _ = v19
                        } else {
                            var v20 uint16 = v13
                            if (!false) {
                                var v21 int16 = v8
                                _ = v21
                            }
                            var v22 int64 = (p1 + 9223372036854775807)
                            _ = v20
                            _ = v22
                        }
                        var v23 bool = ((^(p0)) > (p0 * (p0 % ((p0) | 1))))
                        var v24 int16 = v8
                        _ = v23
                        _ = v24
                    }
                    if false {
                        var v25 int64 = p1
                        p1 = (p1 ^ v25)
                        _ = v25
                    } else {
                        v26 := []int{}
                        switch (-(fi15)) {
                        case 1:
                            var v27 uint16 = (v13 + p2)
                            _ = v27
                        case 0:
                            if true {
                                var v28 float32 = 3.14159
                                _ = v28
                            } else {
                                p0 = p0
                                p1 += int64(v7)
                                switch ("a" + "a") {
                                case "":
                                    if (1.5 != ((float64(3.14159) / -2.25) * 0.1)) {
                                        v29 := GenS2{F0: "a", F1: -918, F2: -49}
                                        v18 += uint(v7)
                                        switch (fi15 % ((fi15) | 1)) {
                                        case 42:
                                            var v30 int32 = 937
                                            _ = v30
                                        case 9223372036854775807:
                                            if true {
                                                v31 := GenS2{F0: "foo bar", F1: -395, F2: 100}
                                                switch "hello" {
                                                case "\n\t":
                                                    p0 = 42
                                                }
                                                if false {
                                                    _ = len(v26)
                                                }
                                                _ = v31
                                            } else {
                                                if (3.14159 >= 0.1) {
                                                    if (!true) {
                                                        if len(v26) > 0 {
                                                            v26[0] = fi14
                                                        }
                                                        v8 += (^(v8))
                                                        if len(v26) >= 2 {
                                                            v26 = v26[1:len(v26)]
                                                        }
                                                    }
                                                } else {
                                                    var v32 int32 = -745
                                                    _ = v32
                                                }
                                                v13 += 16
                                                v26 = append(v26, fi15)
                                            }
                                        case 4294967295:
                                            fmt.Println(p0, v18, v18)
                                        }
                                        _ = v29
                                    } else {
                                        v33 := GenS1{F0: ("\n\t" + "a")}
                                        v7 += 3
                                        switch (^(fi14)) {
                                        case 128:
                                            if true {
                                                if true {
                                                    if (p2 < p2) {
                                                        var v34 uint64 = 18446744073709551615
                                                        p0 += p0
                                                        v35 := GenS2{F0: ("foo bar" + ("foo bar" + "a")), F1: 1, F2: 127}
                                                        _ = v34
                                                        _ = v35
                                                    } else {
                                                        var v36 int8 = -127
                                                        _ = v36
                                                    }
                                                    var v37 int64 = p1
                                                    _ = v37
                                                } else {
                                                    if false {
                                                        v38 := GenS2{F0: "\n\t", F1: 0, F2: 0}
                                                        v39 := GenS1{F0: ""}
                                                        if len(v26) > 0 {
                                                            v26[0] = -395
                                                        }
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                }
                                                var v40 int16 = 32767
                                                fmt.Println(p2)
                                                _ = v40
                                            }
                                        case 4294967295:
                                            switch "hello" {
                                            case "a":
                                                fmt.Println(v18)
                                            case "hello":
                                                if true {
                                                    p1 = -9223372036854775808
                                                    switch "\n\t" {
                                                    case "hello":
                                                        fmt.Println(p0, v18)
                                                    case "foo bar":
                                                        if (p1 >= (p1 << ((v18) & 63))) {
                                                            if true {
                                                                if (!(false && false)) {
                                                                    v41 := GenS3{F0: (^(v18))}
                                                                    var v42 int8 = 1
                                                                    _ = v41
                                                                    _ = v42
                                                                }
                                                                if len(v26) > 0 {
                                                                    v26[0] = fi14
                                                                }
                                                                p2 += (p2 + p2)
                                                            } else {
                                                                var v43 float64 = 0.0
                                                                _ = v43
                                                            }
                                                            v44 := GenS2{F0: "a", F1: 0, F2: 1}
                                                            if true {
                                                                p0 += p0
                                                            } else {
                                                                var v45 uint8 = v7
                                                                _ = v45
                                                            }
                                                            _ = v44
                                                        }
                                                    }
                                                    fmt.Println(v18, fi15, p1)
                                                }
                                            }
                                        case 9223372036854775807:
                                            if (((1e10 - 0.1) + -2.25) < (-2.25 - (1e10 * 1e10))) {
                                                if len(v26) >= 2 {
                                                    v26 = v26[1:len(v26)]
                                                }
                                                if false {
                                                    v18 = v18
                                                    v26 = append(v26, (fi15 % (((fi14 | 32767)) | 1)))
                                                    if false {
                                                        fmt.Println(v7)
                                                    } else {
                                                        if false {
                                                            var v46 bool = false
                                                            _ = v46
                                                        } else {
                                                            _ = len(v26)
                                                        }
                                                        if (!(false || (0 == v7))) {
                                                            if ((v8 % ((-1) | 1)) != (v8 - v8)) {
                                                                v47 := GenS3{F0: (^(v18))}
                                                                _ = len(v26)
                                                                var v48 float32 = 1.5
                                                                _ = v47
                                                                _ = v48
                                                            } else {
                                                                var v49 int64 = (p1 - p1)
                                                                _ = v49
                                                            }
                                                            var v50 uint32 = p0
                                                            fmt.Println(v50)
                                                            var v51 int8 = -68
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                        p0 += 4294967295
                                                    }
                                                    var v52 int16 = v8
                                                    _ = v52
                                                } else {
                                                    fmt.Println(p1, p2)
                                                    v18 = v18
                                                }
                                            }
                                        default:
                                            if (!(true && (!false))) {
                                                if (true && false) {
                                                    p1 -= -1
                                                    p1 -= (p1 + (-(p1)))
                                                }
                                            } else {
                                                fmt.Println(p0, v7)
                                                v7 = (v7 | (v7 % (((v7 * (v7 << (((v18 / (((v18 | v18)) | 1))) & 63)))) | 1)))
                                            }
                                        }
                                        _ = v33
                                    }
                                case "a":
                                    p0 -= (p0 * (p0 & (^(p0))))
                                }
                                p2 -= p2
                                if len(v26) > 0 {
                                    v26[0] = fi14
                                }
                            }
                            var v53 uint64 = 1
                            _ = v53
                        }
                        switch "\n\t" {
                        case "foo bar":
                            var v54 int16 = (v8 / ((-32768) | 1))
                            _ = v54
                        case "a":
                            p1 += (p1 - p1)
                        case "\n\t":
                            var v55 uint64 = 18446744073709551615
                            _ = v55
                        }
                        if (!false) {
                            var v56 uint32 = 625
                            var v57 float32 = 0.0
                            p0 -= (v56 / (((v56 - p0)) | 1))
                            _ = v56
                            _ = v57
                        } else {
                            v58 := GenS4{F0: 18446744073709551615, F1: v8}
                            if true {
                                var v59 string = "foo bar"
                                _ = v59
                            }
                            _ = v58
                        }
                        _ = v26
                    }
                    _ = fi15
                    _ = v16
                    _ = v18
                }
                v60 := []int{(fi14 % ((-525) | 1)), fi14}
                if len(v60) >= 2 {
                    v60 = v60[1:len(v60)]
                }
                _ = fi14
                _ = v60
            }
            v61 := []int{0, -2147483648}
            p1 = 42
            _ = v61
        } else {
            v62 := GenS2{F0: "hello", F1: 2147483647, F2: 127}
            _ = v62
        }
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]string)
    var v3 int = 9223372036854775807
    var v4 int = v3
    var v5 bool = false
    var v6 uint8 = 2
    for fi63 := 0; fi63 < 4; fi63++ {
        fmt.Println(v4)
        v4 -= (^(v4))
        _ = fi63
    }
    v64 := []int{v3}
    fmt.Println(v4, v6)
    v64 = append(v64, 0)
    var v65 uint8 = 255
    v66 := GenS4{F0: 255, F1: 1}
    fmt.Println(v3, v6, v6)
    fmt.Println(v4, v3)
    v67 := GenS3{F0: 266}
    if v5 {
        if (false || (830 != (v1 & v1))) {
            if (v5 && ((v4 / ((42) | 1)) < v4)) {
                _ = len(v2)
            }
        }
        var v68 int16 = int16(v1)
        var v69 string = ""
        _ = v68
        _ = v69
    }
    v70 := GenS5{F0: "a"}
    v65 = v65
    switch "foo bar" {
    case "hello":
        switch ("\n\t" + ("a" + "")) {
        case "":
            fmt.Println(v5, v1)
        case "foo bar":
            fmt.Println(v4, v3, v4)
        }
    }
    v64 = append(v64, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v64:", len(v64), v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v70
}
