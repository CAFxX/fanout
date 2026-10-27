package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int; F1 uint8; F2 uint8 }
func genFn0(p0 uint16, p1 int64) int64 {
    var v7 int = 127
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v7
    return p1
}
func genFn1() uint {
    var v8 uint = 352
    var v9 string = ("foo bar" + "hello")
    v10 := []int{382, 0}
    if (true && true) {
        var v11 float64 = (0.0 * (3.14159 * 1.5))
        fmt.Println(v9)
        var v12 float64 = 3.14159
        _ = v11
        _ = v12
    } else {
        if (-1 >= 42) {
            v8 = 0
            if false {
                fmt.Println(v8, v9)
                switch (v9 + v9) {
                case "a":
                    for fi13 := 0; fi13 < 16; fi13++ {
                        var v14 int = (fi13 * (fi13 - -1))
                        switch v9 {
                        case "":
                            var v15 bool = (977 <= v8)
                            _ = v15
                        default:
                            fmt.Println(v8, v9, v8)
                        }
                        _ = fi13
                        _ = v14
                    }
                case "\n\t":
                    v16 := GenS1{F0: 2147483647, F1: 127}
                    _ = v16
                }
            } else {
                v10 = append(v10, -593)
                var v17 uint16 = 2
                _ = v17
            }
            if true {
                v8 = (^(v8))
            }
        } else {
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 int64 = 9223372036854775807
                if (1.5 > 0.0) {
                    for fi20 := 0; fi20 < 5; fi20++ {
                        v21 := GenS2{F0: 45}
                        if false {
                            if len(v10) > 0 {
                                v10[0] = (fi20 & 32768)
                            }
                            v19 -= v19
                            v19 += v19
                        }
                        switch v9 {
                        case "":
                            v8 = (^(v8))
                        }
                        _ = fi20
                        _ = v21
                    }
                    var v22 uint16 = 1
                    _ = v22
                }
                _ = fi18
                _ = v19
            }
            var v23 string = (v9 + v9)
            _ = v23
        }
        var v24 int16 = 42
        v8 += 173
        _ = v24
    }
    v8 -= v8
    _ = v10
    _ = v8
    _ = v9
    return (v8 >> ((v8) & 63))
}

func main() {
    var v1 uint = 255
    var v2 int = -1
    var v3 float32 = (float32((1e10 + 3.14159)) - ((1e10 + (float32(0.1) - (0.1 - -2.25))) / -2.25))
    var v4 int32 = 122
    var v5 uint = v1
    v6 := make(map[int]string)
    var v25 uint64 = 0
    var v26 uint16 = 32767
    for fi27 := 0; fi27 < 5; fi27++ {
        fmt.Println(v26, v1)
        if true {
            for fi28 := 0; fi28 < 7; fi28++ {
                if false {
                    v29 := GenS2{F0: 0}
                    _ = v29
                }
                v30 := GenS2{F0: -1}
                switch "\n\t" {
                case "\n\t":
                    if true {
                        v25 = (v25 - v25)
                        switch "foo bar" {
                        case "":
                            v31 := []int{(-(fi27))}
                            _ = v31
                        case "\n\t":
                            if (1e10 <= ((-2.25 / ((0.0 - (3.14159 * (-2.25 - (1e10 * 0.0)))) + ((((-2.25 - (1.5 * 0.1)) + 1e10) / 1.0) * 1e10))) / (-2.25 - 0.1))) {
                                if (("foo bar" + "hello") < ((("\n\t" + "\n\t") + "a") + "\n\t")) {
                                    if true {
                                        v32 := GenS1{F0: 0, F1: 19}
                                        _ = v32
                                    } else {
                                        switch -2147483648 {
                                        case 4294967295:
                                            var v33 uint64 = 2
                                            _ = v33
                                        case 829:
                                            v3 += float32(float64(v3))
                                        default:
                                            var v34 uint = v5
                                            _ = v34
                                        }
                                        var v35 int32 = int32(v1)
                                        fmt.Println(fi27, v25, v3)
                                        var v36 int16 = 1
                                        _ = v35
                                        _ = v36
                                    }
                                    v5 += genFn1()
                                } else {
                                    var v37 uint64 = (v25 ^ v25)
                                    _ = v37
                                }
                                v3 = (0.0 + v3)
                            } else {
                                var v38 float32 = -2.25
                                if ("\n\t" > ("a" + "")) {
                                    switch "a" {
                                    case "foo bar":
                                        var v39 int64 = 42
                                        _ = v39
                                    }
                                }
                                if (!true) {
                                    if true {
                                        v1 += (v1 >> ((v1) & 63))
                                        var v40 uint64 = (v25 * v25)
                                        v41 := GenS1{F0: (-(v4)), F1: 1}
                                        _ = v40
                                        _ = v41
                                    }
                                    var v42 int16 = 678
                                    v4 -= (v4 + 248)
                                    _ = v42
                                }
                                _ = v38
                            }
                        case "a":
                            switch (fi27 ^ -9223372036854775808) {
                            case 0:
                                v43 := GenS3{F0: 1}
                                _ = v43
                            case -2147483648:
                                v44 := GenS4{F0: 2147483648, F1: ("foo bar" + (("\n\t" + "a") + ""))}
                                _ = v44
                            case -826:
                                v45 := GenS5{F0: -45}
                                _ = v45
                            }
                        default:
                            v2 = (fi28 & -2147483648)
                        }
                    }
                    v46 := GenS6{F0: v4, F1: 65535, F2: 127}
                    if true {
                        v25 = (v25 & v25)
                        if (0 <= 2) {
                            var v47 string = "hello"
                            _ = v47
                        } else {
                            var v48 int8 = -32
                            var v49 uint32 = 0
                            v49 = v49
                            _ = v48
                            _ = v49
                        }
                        v4 -= v4
                    } else {
                        var v50 bool = false
                        _ = v50
                    }
                    _ = v46
                case "foo bar":
                    switch ("\n\t" + "\n\t") {
                    case "":
                        fmt.Println(v5, v4, v25)
                    case "\n\t":
                        v51 := []int{764, (-(fi28))}
                        v52 := v6[(v2 - fi27)]
                        _ = v51
                        _ = v52
                    case "foo bar":
                        switch v2 {
                        case 128:
                            v3 -= (v3 / v3)
                        case 246:
                            var v53 int32 = -1
                            v54 := GenS7{F0: (fi27 < (fi28 + fi27))}
                            _ = v53
                            _ = v54
                        default:
                            if true {
                                var v55 uint = uint(v25)
                                _ = v55
                            } else {
                                v4 -= v4
                            }
                        }
                        var v56 uint = (v5 * uint(v25))
                        _ = v56
                    default:
                        v1 -= genFn1()
                    }
                case "a":
                    v57 := GenS1{F0: (^(v4)), F1: 118}
                    _ = v57
                }
                v5 = (v1 ^ (v5 >> ((v1) & 63)))
                var v58 uint16 = 831
                var v59 int16 = -370
                _ = fi28
                _ = v30
                _ = v58
                _ = v59
            }
            var v60 uint16 = (v26 % ((v26) | 1))
            if (924 != -1) {
                var v61 int64 = -9223372036854775808
                v62 := GenS8{F0: v2, F1: 149, F2: 127}
                v63 := GenS8{F0: v2, F1: 255, F2: 153}
                _ = v61
                _ = v62
                _ = v63
            }
            _ = v60
        }
        if false {
            switch fi27 {
            case 821:
                v64 := GenS8{F0: fi27, F1: 127, F2: 190}
                _ = v64
            case 32768:
                switch ("a" + ("hello" + "hello")) {
                case "\n\t":
                    var v65 string = "\n\t"
                    _ = v65
                case "foo bar":
                    if (!((true && true) || false)) {
                        if ((v4 * 1) < v4) {
                            switch ("foo bar" + "") {
                            case "":
                                _ = len(v6)
                            case "hello":
                                v6[fi27] = (("foo bar" + "") + ("\n\t" + "\n\t"))
                            }
                            fmt.Println(v3, v4, v25)
                            fmt.Println(v2, fi27)
                        }
                        var v66 uint32 = 65535
                        _ = v66
                    }
                default:
                    v25 -= 18446744073709551615
                }
            }
            for fi67 := 0; fi67 < 14; fi67++ {
                v3 -= v3
                var v68 string = ""
                _ = fi67
                _ = v68
            }
        } else {
            for fi69 := 0; fi69 < 8; fi69++ {
                var v70 uint = genFn1()
                _ = fi69
                _ = v70
            }
        }
        _ = fi27
    }
    var v71 uint = (v1 * v1)
    fmt.Println(v3, v26)
    v72 := GenS6{F0: (v4 + 1), F1: v26, F2: 66}
    switch -199 {
    case 4294967295:
        v73 := []int{(v2 << (((v71 >> ((v5) & 63))) & 63))}
        _ = v73
    case 127:
        v71 = (v5 / ((65535) | 1))
    }
    v4 = v4
    if (1000 >= 42) {
        fmt.Println(v71, v5, v5)
        if true {
            v26 = 2
            v74 := []int{}
            v75 := GenS4{F0: 340, F1: "a"}
            _ = v74
            _ = v75
        } else {
            v71 -= 891
        }
    }
    v76 := []int{(v2 / ((127) | 1))}
    for fi77 := 0; fi77 < 11; fi77++ {
        var v78 uint16 = (v26 ^ v26)
        v26 = v78
        _ = fi77
        _ = v78
    }
    v79 := v6[v2]
    v80 := GenS4{F0: 55, F1: (((v79 + "foo bar") + "\n\t") + "")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v71:", v71)
    fmt.Println("slicelen v76:", len(v76), v76)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v72
    _ = v76
    _ = v79
    _ = v80
}
