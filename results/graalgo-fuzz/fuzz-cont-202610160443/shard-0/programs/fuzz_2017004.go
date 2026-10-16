package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint; F1 int64; F2 string }
type GenS3 struct { F0 bool; F1 int16; F2 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }
func genFn0(p0 int16) int64 {
    v5 := GenS1{F0: -2147483648, F1: 0}
    v6 := GenS1{F0: -693, F1: 779}
    _ = p0
    _ = v5
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 uint64) int64 {
    if true {
        p0 = 0
        switch 0 {
        case -2147483648:
            var v7 int = 42
            _ = v7
        case 42:
            v8 := GenS2{F0: 42, F1: int64(p0), F2: "\n\t"}
            _ = v8
        }
        var v9 uint16 = 380
        _ = v9
    }
    switch "foo bar" {
    case "hello":
        var v10 int = -2147483648
        _ = v10
    case "foo bar":
        p0 -= p0
    case "\n\t":
        v11 := GenS3{F0: true, F1: 0, F2: -337}
        _ = v11
    default:
        p0 -= p0
    }
    fmt.Println(p0)
    _ = p0
    return -1
}

func main() {
    var v1 int32 = -1
    var v2 float32 = 0.0
    var v3 string = ""
    var v4 int = -2147483648
    if false {
        v12 := []int{}
        if (194 != genFn0(-323)) {
            if true {
                v13 := GenS4{F0: (v3 + v3)}
                v1 = v1
                _ = v13
            } else {
                for fi14 := 0; fi14 < 11; fi14++ {
                    fmt.Println(v3)
                    v3 = (v3 + (v3 + ""))
                    v3 = "hello"
                    _ = fi14
                }
                var v15 int8 = 0
                for fi16 := 0; fi16 < 6; fi16++ {
                    v2 = (v2 * ((3.14159 - 0.1) * v2))
                    _ = fi16
                }
                _ = v15
            }
        } else {
            var v17 int = -9223372036854775808
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 uint64 = 2
                var v20 float32 = (0.1 / v2)
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v17
        }
        _ = v12
    }
    var v21 uint64 = 114
    var v22 int8 = 0
    fmt.Println(v2, v2, v2)
    var v23 int8 = v22
    fmt.Println(v21, v1)
    fmt.Println(v22, v1)
    var v24 float32 = (0.0 + v2)
    for fi25 := 0; fi25 < 6; fi25++ {
        for fi26 := 0; fi26 < 16; fi26++ {
            if false {
                if true {
                    switch 953 {
                    case 887:
                        var v27 uint = 4294967295
                        _ = v27
                    }
                    v28 := GenS1{F0: 1, F1: 2}
                    _ = v28
                } else {
                    fmt.Println(v22, v1, v22)
                    v29 := GenS5{F0: genFn0(0)}
                    v22 = v22
                    _ = v29
                }
            } else {
                if false {
                    v23 -= 100
                    switch (v4 << ((255) & 63)) {
                    case 47:
                        v23 -= (v23 >> ((918) & 63))
                    }
                    fmt.Println(v1, v4, v22)
                } else {
                    v30 := GenS4{F0: "hello"}
                    _ = v30
                }
            }
            if (false || false) {
                var v31 int = (fi26 % (((^(fi25))) | 1))
                v23 -= v22
                var v32 int8 = v23
                fmt.Println(v24, v4, v32)
                var v33 int32 = (v1 ^ (v1 << ((4294967295) & 63)))
                var v34 uint64 = v21
                _ = v31
                _ = v32
                _ = v33
                _ = v34
            } else {
                var v35 uint8 = 255
                _ = v35
            }
            fmt.Println(v22)
            _ = fi26
        }
        for fi36 := 0; fi36 < 3; fi36++ {
            if (1e10 < 0.1) {
                if false {
                    v37 := []int{(fi25 << ((255) & 63)), fi25, (fi36 | v4), 9223372036854775807}
                    fmt.Println(v1, v21)
                    var v38 int64 = 546
                    _ = v37
                    _ = v38
                }
            }
            v39 := GenS5{F0: -1}
            _ = fi36
            _ = v39
        }
        v3 = "hello"
        _ = fi25
    }
    var v40 uint = 628
    if (false && true) {
        if true {
            for fi41 := 0; fi41 < 9; fi41++ {
                var v42 float32 = v2
                v2 -= v42
                var v43 uint32 = 162
                _ = fi41
                _ = v42
                _ = v43
            }
            v44 := []int{v4}
            _ = v44
        } else {
            var v45 string = ((v3 + (("" + v3) + "")) + "")
            _ = v45
        }
        var v46 uint64 = (^(v21))
        var v47 uint64 = 18446744073709551615
        _ = v46
        _ = v47
    } else {
        v4 = (v4 ^ 941)
    }
    v48 := GenS6{F0: -631}
    var v49 uint16 = 255
    if true {
        switch v4 {
        case -575:
            var v50 uint = 303
            _ = v50
        case 127:
            if ((v40 ^ 65535) != v40) {
                v51 := GenS4{F0: (v3 + (v3 + "a"))}
                _ = v51
            } else {
                v40 = 18446744073709551615
            }
        case -584:
            switch v3 {
            case "a":
                v52 := GenS6{F0: (v4 >> ((391) & 63))}
                _ = v52
            case "":
                fmt.Println(v40, v3, v1)
            default:
                var v53 uint64 = (v21 * v21)
                _ = v53
            }
        }
    } else {
        var v54 int8 = (v22 | (v23 - v22))
        _ = v54
    }
    v55 := GenS5{F0: -532}
    if true {
        v56 := []int{}
        fmt.Println(v24, v49, v4)
        _ = v56
    }
    for fi57 := 0; fi57 < 2; fi57++ {
        v58 := GenS2{F0: (v40 ^ (v40 ^ v40)), F1: 42, F2: ((v3 + "foo bar") + "hello")}
        if (v21 >= v21) {
            if true {
                v24 -= v2
            }
            var v59 int = -578
            _ = v59
        } else {
            var v60 int16 = 0
            v4 += (^(fi57))
            for fi61 := 0; fi61 < 15; fi61++ {
                v4 = -625
                v40 += v40
                if (2147483647 < v1) {
                    v3 = "hello"
                } else {
                    if false {
                        v62 := GenS5{F0: -411}
                        fmt.Println(v49, fi61, v4)
                        v63 := []int{fi57, (v4 >> (((v40 + v40)) & 63)), (-(fi61)), (-(v4))}
                        _ = v62
                        _ = v63
                    }
                    var v64 uint16 = 1
                    if (!(true || (128 >= 2))) {
                        var v65 float64 = -2.25
                        _ = v65
                    } else {
                        var v66 uint64 = uint64(v64)
                        _ = v66
                    }
                    _ = v64
                }
                v67 := []int{(fi57 + 32768), (v4 * (fi61 % (((fi57 ^ (^(fi61)))) | 1)))}
                v23 += (v22 / (((v22 >> ((v40) & 63))) | 1))
                _ = fi61
                _ = v67
            }
            _ = v60
        }
        var v68 string = "a"
        _ = fi57
        _ = v58
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v48
    _ = v49
    _ = v55
}
