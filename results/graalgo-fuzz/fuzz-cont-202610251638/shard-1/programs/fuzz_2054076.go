package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 int64; F1 bool; F2 int32 }
type GenS7 struct { F0 int; F1 uint; F2 uint64 }
type GenS8 struct { F0 uint; F1 int32; F2 int8 }
type GenS9 struct { F0 int16; F1 uint16; F2 uint }
func genFn0(p0 string) bool {
    fmt.Println(p0)
    switch -671 {
    case 1:
        p0 = p0
    }
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 uint32 = 1
        var v6 uint8 = 165
        var v7 int16 = 42
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    p0 = "a"
    var v8 int8 = -1
    _ = p0
    _ = v8
    return true
}
func genFn1(p0 int) int32 {
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: 18446744073709551615}
        if (false || true) {
            var v11 int16 = 42
            fmt.Println(p0)
            v12 := []int{p0, -1}
            _ = v11
            _ = v12
        } else {
            fmt.Println(p0, p0)
            if (0 < 1) {
                v13 := GenS1{F0: 4294967295}
                v14 := []int{fi9, p0, p0}
                _ = v13
                _ = v14
            }
            switch "a" {
            case "a":
                fmt.Println(p0, fi9)
            case "hello":
                var v15 float64 = (-2.25 - float64((float32(0.0) * ((0.1 + (0.0 - float32(3.14159))) - (3.14159 - 1e10)))))
                _ = v15
            }
        }
        fmt.Println(p0, fi9)
        _ = fi9
        _ = v10
    }
    switch "hello" {
    case "\n\t":
        switch ((("a" + "") + (("a" + "foo bar") + "\n\t")) + (("a" + "\n\t") + "hello")) {
        case "\n\t":
            v16 := GenS1{F0: 1}
            _ = v16
        case "":
            p0 = 457
        default:
            var v17 uint8 = 128
            _ = v17
        }
    case "a":
        switch "" {
        case "\n\t":
            var v18 string = ("foo bar" + ("" + ("hello" + "hello")))
            _ = v18
        default:
            if (163 == 396) {
                v19 := []int{p0, p0, 9223372036854775807, 32767}
                _ = v19
            } else {
                switch "a" {
                case "\n\t":
                    switch "\n\t" {
                    case "\n\t":
                        var v20 uint = 504
                        _ = v20
                    case "hello":
                        fmt.Println(p0)
                    default:
                        fmt.Println(p0)
                    }
                case "":
                    if (false || true) {
                        switch (p0 + (p0 % ((p0) | 1))) {
                        case 127:
                            if false {
                                p0 += -205
                                var v21 int = (p0 * (p0 % ((p0) | 1)))
                                v22 := []int{1, v21}
                                _ = v21
                                _ = v22
                            } else {
                                var v23 float32 = float32(1.5)
                                _ = v23
                            }
                        default:
                            var v24 float32 = float32(0.0)
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 2; fi25++ {
                            var v26 uint32 = 206
                            _ = fi25
                            _ = v26
                        }
                        var v27 int = (^(p0))
                        _ = v27
                    }
                }
                if true {
                    fmt.Println(p0)
                } else {
                    var v28 int = p0
                    _ = v28
                }
                if true {
                    if (float32(0.0) == 1e10) {
                        v29 := GenS2{F0: 65535}
                        v30 := []int{p0, 127, p0, -849}
                        _ = len(v30)
                        _ = v29
                        _ = v30
                    }
                    v31 := []int{-2147483648, (p0 / (((p0 << ((261) & 63))) | 1))}
                    var v32 int64 = 0
                    _ = v31
                    _ = v32
                }
            }
        }
    case "foo bar":
        var v33 uint16 = 42
        _ = v33
    default:
        v34 := []int{(p0 | 4294967295)}
        _ = v34
    }
    var v35 float32 = 3.14159
    fmt.Println(p0, p0)
    _ = p0
    _ = v35
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ("foo bar" + "\n\t")
    var v3 int32 = 2147483647
    v2 = "foo bar"
    fmt.Println(v3, v3)
    for fi36 := 0; fi36 < 8; fi36++ {
        v2 = "a"
        _ = fi36
    }
    var v37 int64 = 9223372036854775807
    var v38 uint32 = 65535
    var v39 uint = 172
    v39 = v39
    if (1e10 >= float32(((-2.25 / (3.14159 + (-2.25 - 1e10))) - -2.25))) {
        switch ("\n\t" + "\n\t") {
        case "a":
            var v40 int8 = 42
            _ = v40
        case "\n\t":
            if false {
                v41 := []int{-9223372036854775808, 255, 42, 2147483647}
                _ = v41
            }
        case "":
            v39 += v39
        default:
            fmt.Println(v39, v3, v38)
        }
        v42 := GenS3{F0: 2147483648, F1: 592}
        _ = v42
    } else {
        var v43 uint32 = uint32(v39)
        delete(v1, (v43 * 5))
        _ = v43
    }
    v44 := GenS2{F0: 32767}
    v45 := GenS3{F0: 0, F1: 1}
    if true {
        v37 += (v37 - v37)
    } else {
        var v46 uint16 = 65535
        v38 = v38
        _ = v46
    }
    v47 := GenS4{F0: (v39 >> ((v39) & 63))}
    if true {
        switch 127 {
        case 32768:
            if true {
                if true {
                    v37 = 402
                    v3 += -2147483648
                    var v48 int64 = (v37 - (v37 & v37))
                    _ = v48
                }
            }
        case -9223372036854775808:
            v49 := []int{-2147483648, -204}
            _ = v49
        }
        switch ("hello" + (("\n\t" + ((v2 + ((v2 + v2) + (v2 + v2))) + (v2 + "\n\t"))) + (v2 + "hello"))) {
        case "foo bar":
            switch "\n\t" {
            case "hello":
                delete(v1, v38)
            }
        case "":
            switch "hello" {
            case "":
                v50 := GenS5{F0: v39, F1: true}
                _ = v50
            }
        case "\n\t":
            var v51 uint64 = 873
            _ = v51
        default:
            v52 := GenS6{F0: v37, F1: (127 == uint8(v39)), F2: v3}
            _ = v52
        }
        if (127 != 0) {
            v39 = 2
            var v53 float32 = 0.1
            _ = v53
        } else {
            v39 = 702
            switch 317 {
            case 128:
                for fi54 := 0; fi54 < 15; fi54++ {
                    v55 := []int{436, (fi54 * int(v38))}
                    if false {
                        v3 -= int32(fi54)
                        fmt.Println(v3)
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v39, v38, v2)
                        case "hello":
                            v37 = v37
                        case "a":
                            v55 = append(v55, (fi54 - (fi54 - -500)))
                        default:
                            for fi56 := 0; fi56 < 5; fi56++ {
                                switch (^(fi54)) {
                                case -176:
                                    if (0 >= 128) {
                                        fmt.Println(v2, fi54)
                                        v57 := GenS7{F0: (fi56 ^ fi56), F1: 218, F2: 42}
                                        _ = v57
                                    } else {
                                        if true {
                                            v58 := GenS6{F0: (v37 / (((^(v37))) | 1)), F1: (fi56 <= (fi54 / ((fi56) | 1))), F2: 2147483647}
                                            _ = v58
                                        }
                                        var v59 uint64 = 202
                                        var v60 bool = false
                                        _ = v59
                                        _ = v60
                                    }
                                case -698:
                                    v39 += (v39 & v39)
                                case 32768:
                                    v61 := GenS5{F0: (v39 << (((^(v39))) & 63)), F1: (!(0 < -32768))}
                                    _ = v61
                                }
                                var v62 uint8 = 255
                                _ = fi56
                                _ = v62
                            }
                        }
                    }
                    var v63 uint = 65535
                    _ = fi54
                    _ = v55
                    _ = v63
                }
            case 792:
                fmt.Println(v3, v3, v38)
            case -818:
                var v64 uint64 = 18446744073709551615
                _ = v64
            }
            v65 := GenS8{F0: v39, F1: (-(v3)), F2: 42}
            _ = v65
        }
    }
    v66 := GenS9{F0: 42, F1: 65535, F2: 954}
    fmt.Println(v39)
    var v67 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v44
    _ = v45
    _ = v47
    _ = v66
    _ = v67
}
