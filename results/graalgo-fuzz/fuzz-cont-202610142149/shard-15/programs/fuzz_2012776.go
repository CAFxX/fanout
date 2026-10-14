package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 bool }
type GenS2 struct { F0 uint; F1 int; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint8; F2 int }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 uint; F2 int }
type GenS7 struct { F0 uint32; F1 int16 }
func genFn0() int64 {
    if false {
        var v4 uint32 = 123
        _ = v4
    }
    var v5 float32 = 3.14159
    _ = v5
    return 462
}
func genFn1(p0 uint32, p1 bool) uint8 {
    fmt.Println(p0, p1)
    switch ("a" + "") {
    case "a":
        p0 += p0
    case "":
        var v6 uint64 = 0
        _ = v6
    default:
        var v7 int = 2147483647
        _ = v7
    }
    v8 := []int{-252, 602, -1, -1}
    _ = p0
    _ = p1
    _ = v8
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 42
    var v3 uint32 = 2147483648
    for fi9 := 0; fi9 < 5; fi9++ {
        fmt.Println(fi9, v3)
        switch "a" {
        case "hello":
            v2 -= (v2 * v2)
        case "\n\t":
            v2 -= (v2 | (v2 ^ 42))
        case "foo bar":
            var v10 uint64 = 42
            _ = v10
        default:
            fmt.Println(v3, v2, v3)
        }
        _ = fi9
    }
    v11 := v1[int(v3)]
    fmt.Println(v11, v3)
    v2 += (v2 / ((v2) | 1))
    _ = len(v1)
    if true {
        v11 = v11
        v12 := []int{927, -419}
        if false {
            v13 := GenS1{F0: 1, F1: -52, F2: false}
            _ = v13
        } else {
            if ((v3 >> ((0) & 63)) < 512) {
                var v14 uint64 = 42
                var v15 float64 = ((-2.25 + 3.14159) * -2.25)
                var v16 int64 = 555
                _ = v14
                _ = v15
                _ = v16
            } else {
                var v17 uint16 = 2
                v3 -= (v11 >> ((1) & 63))
                _ = v17
            }
            var v18 uint16 = 65535
            fmt.Println(v2)
            _ = v18
        }
        _ = v12
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            switch "a" {
            case "a":
                var v21 string = ("foo bar" + (((("" + ((("foo bar" + (("a" + ("foo bar" + "hello")) + "a")) + "foo bar") + "foo bar")) + "a") + "\n\t") + ("a" + "")))
                _ = v21
            case "foo bar":
                var v22 uint = 4294967295
                _ = v22
            default:
                v23 := v1[(fi20 >> ((18446744073709551615) & 63))]
                _ = v23
            }
            v24 := GenS1{F0: 100, F1: 0, F2: false}
            _ = fi20
            _ = v24
        }
        v25 := GenS1{F0: 42, F1: -1, F2: false}
        v3 += 2147483647
        _ = fi19
        _ = v25
    }
    switch "hello" {
    case "foo bar":
        switch 32768 {
        case 2147483647:
            v11 = (v3 ^ v3)
        default:
            var v26 int = -2147483648
            _ = v26
        }
    case "\n\t":
        for fi27 := 0; fi27 < 4; fi27++ {
            v3 -= 42
            _ = fi27
        }
    default:
        for fi28 := 0; fi28 < 6; fi28++ {
            v1[fi28] = v3
            if (false || false) {
                v1[fi28] = v3
            } else {
                var v29 int32 = -403
                _ = v29
            }
            switch "" {
            case "a":
                var v30 float64 = (float64(((0.0 / 1e10) / 3.14159)) + 0.1)
                _ = v30
            default:
                for fi31 := 0; fi31 < 8; fi31++ {
                    v32 := GenS2{F0: 2, F1: (fi28 >> ((42) & 63)), F2: 0}
                    _ = fi31
                    _ = v32
                }
            }
            _ = fi28
        }
    }
    v2 = v2
    v2 = 32767
    var v33 int16 = 1000
    v11 += 515
    v34 := GenS3{F0: 1, F1: 2, F2: -1}
    var v35 string = ("hello" + "")
    _ = len(v1)
    switch v35 {
    case "foo bar":
        switch v35 {
        case "a":
            v1[9223372036854775807] = (^(v11))
        case "hello":
            switch ("\n\t" + v35) {
            case "":
                v33 += (v33 & v2)
            case "foo bar":
                v11 -= v11
            case "hello":
                if false {
                    v36 := GenS1{F0: -128, F1: 127, F2: (false || false)}
                    var v37 float32 = 1e10
                    _ = v36
                    _ = v37
                }
            }
        default:
            var v38 uint64 = 1
            _ = v38
        }
    case "a":
        var v39 uint8 = 128
        _ = v39
    case "hello":
        switch 1 {
        case 4294967295:
            if (false && ((v3 - v3) < (v3 | v3))) {
                var v40 string = v35
                if true {
                    v41 := GenS2{F0: 943, F1: 1, F2: 0}
                    _ = v41
                } else {
                    var v42 int16 = v33
                    for fi43 := 0; fi43 < 4; fi43++ {
                        v3 = (v11 - (v11 >> ((534) & 63)))
                        _ = fi43
                    }
                    v44 := []int{42, -209, -2147483648}
                    _ = v42
                    _ = v44
                }
                _ = v40
            }
        case 1:
            for fi45 := 0; fi45 < 6; fi45++ {
                v46 := GenS3{F0: int64(v2), F1: 42, F2: fi45}
                for fi47 := 0; fi47 < 2; fi47++ {
                    fmt.Println(v2, v3)
                    fmt.Println(v35, v2)
                    _ = fi47
                }
                v35 = (v35 + v35)
                _ = fi45
                _ = v46
            }
        case 204:
            if false {
                fmt.Println(v11, v35)
                v48 := GenS1{F0: 70, F1: 100, F2: true}
                _ = v48
            }
        default:
            v49 := GenS4{F0: v33, F1: 1}
            _ = v49
        }
    default:
        v3 = v3
    }
    if false {
        if true {
            v50 := []int{128, 1, -9223372036854775808, 9223372036854775807}
            v35 = (v35 + "hello")
            _ = v50
        }
        if (-2147483648 > 42) {
            if false {
                v51 := []int{157, -2147483648, 9223372036854775807}
                v35 = (v35 + "\n\t")
                _ = v51
            }
        } else {
            v35 = v35
            v35 = v35
            v52 := []int{}
            _ = v52
        }
        var v53 uint16 = 65535
        _ = v53
    }
    switch "\n\t" {
    case "\n\t":
        v54 := GenS5{F0: (-32768 <= v2)}
        _ = v54
    case "foo bar":
        for fi55 := 0; fi55 < 4; fi55++ {
            var v56 int32 = 786
            var v57 uint32 = 2147483648
            var v58 float64 = 1.5
            _ = fi55
            _ = v56
            _ = v57
            _ = v58
        }
    case "":
        for fi59 := 0; fi59 < 17; fi59++ {
            var v60 int = 9223372036854775807
            _ = fi59
            _ = v60
        }
    default:
        for fi61 := 0; fi61 < 6; fi61++ {
            var v62 uint32 = 1
            if (0.0 > 1.5) {
                var v63 uint = 1
                v35 = ""
                _ = v63
            }
            for fi64 := 0; fi64 < 8; fi64++ {
                v65 := GenS6{F0: 65535, F1: 42, F2: (fi64 * fi64)}
                if true {
                    v35 = "\n\t"
                    var v66 uint64 = 18446744073709551615
                    _ = v66
                }
                v67 := []int{}
                v68 := GenS7{F0: v62, F1: int16(v62)}
                _ = fi64
                _ = v65
                _ = v67
                _ = v68
            }
            _ = fi61
            _ = v62
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
