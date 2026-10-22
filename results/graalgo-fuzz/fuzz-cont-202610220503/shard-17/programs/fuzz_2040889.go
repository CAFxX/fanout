package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int; F2 bool }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 string }
func genFn0(p0 int16, p1 uint16) int8 {
    v7 := GenS1{F0: 92}
    v8 := []int{-787, -2147483648, -1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 93
}
func genFn1(p0 int16, p1 uint) uint8 {
    switch ("a" + "hello") {
    case "hello":
        v9 := []int{-1, -483}
        _ = v9
    case "\n\t":
        switch ("" + ("foo bar" + "a")) {
        case "foo bar":
            p0 += p0
        case "hello":
            fmt.Println(p0)
        case "a":
            v10 := GenS1{F0: 255}
            _ = v10
        default:
            for fi11 := 0; fi11 < 5; fi11++ {
                p1 -= p1
                _ = fi11
            }
        }
    case "":
        v12 := GenS1{F0: 127}
        _ = v12
    default:
        v13 := []int{32768, 32767}
        _ = v13
    }
    if true {
        var v14 int8 = 1
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 uint8 = 116
        var v17 int16 = (p0 + (^(p0)))
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 236}
    _ = p0
    _ = p1
    _ = v18
    return 128
}

func main() {
    var v1 int = -9223372036854775808
    var v2 float32 = 0.1
    var v3 uint = 18446744073709551615
    var v4 int16 = 0
    v5 := []int{0}
    var v6 uint8 = 99
    if ((!true) || ((v1 >= v1) && true)) {
        switch -1 {
        case -922:
            if false {
                var v19 uint = 255
                _ = v19
            } else {
                v5 = append(v5, (v1 / (((-(v1))) | 1)))
                _ = len(v5)
                var v20 float32 = (v2 * 3.14159)
                _ = v20
            }
        case 9223372036854775807:
            if ((v4 / (((^(v4))) | 1)) < (v4 ^ -1)) {
                fmt.Println(v2)
                for fi21 := 0; fi21 < 11; fi21++ {
                    for fi22 := 0; fi22 < 10; fi22++ {
                        v6 = 1
                        _ = fi22
                    }
                    var v23 string = ""
                    if (("a" + "hello") == v23) {
                        v2 = (1e10 / v2)
                        fmt.Println(fi21, v6)
                    } else {
                        var v24 int16 = (v4 % (((v4 % (((v4 * v4)) | 1))) | 1))
                        v3 -= (v3 << ((916) & 63))
                        _ = v24
                    }
                    _ = fi21
                    _ = v23
                }
            }
        case -1:
            if (v4 != (v4 << (((v3 * v3)) & 63))) {
                var v25 uint64 = 18446744073709551615
                _ = v25
            }
        default:
            fmt.Println(v4, v1)
        }
        if (759 == 65535) {
            var v26 float64 = ((0.1 + 0.0) / float64(1.5))
            switch v1 {
            case -590:
                var v27 bool = false
                _ = v27
            case 801:
                v4 = (-(v4))
            }
            _ = v26
        } else {
            var v28 uint32 = 0
            _ = v28
        }
        v2 = (v2 / v2)
    } else {
        var v29 int8 = -91
        if len(v5) > 0 {
            v5[0] = 287
        }
        var v30 int16 = (v4 << ((v3) & 63))
        _ = v29
        _ = v30
    }
    if false {
        for fi31 := 0; fi31 < 6; fi31++ {
            v32 := GenS2{F0: (v3 | 697), F1: 0}
            _ = fi31
            _ = v32
        }
    } else {
        v4 += (^(v4))
        var v33 bool = false
        _ = v33
    }
    fmt.Println(v6)
    if true {
        if true {
            switch "a" {
            case "":
                var v34 string = (("foo bar" + "\n\t") + ("" + ("hello" + ("" + ""))))
                _ = v34
            case "\n\t":
                v35 := GenS1{F0: (v6 | 8)}
                _ = v35
            case "hello":
                var v36 uint = v3
                _ = v36
            }
            for fi37 := 0; fi37 < 11; fi37++ {
                var v38 int32 = -1
                _ = fi37
                _ = v38
            }
            for fi39 := 0; fi39 < 3; fi39++ {
                var v40 uint64 = 2
                v41 := GenS3{F0: 255, F1: 938, F2: v40}
                _ = len(v5)
                _ = fi39
                _ = v40
                _ = v41
            }
        } else {
            if len(v5) > 0 {
                v5[0] = (v1 / (((^(v1))) | 1))
            }
            var v42 float64 = 1.5
            _ = v42
        }
        if (1e10 < 1e10) {
            v43 := GenS4{F0: genFn1((^(v4)), v3), F1: (v1 << ((568) & 63)), F2: (((v1 * (v1 + (v1 + v1))) > 0) || true)}
            _ = v43
        } else {
            v44 := GenS5{F0: uint(v6), F1: 600}
            if len(v5) > 0 {
                v5[0] = v1
            }
            var v45 int32 = 1
            _ = v44
            _ = v45
        }
    } else {
        switch (v1 ^ (-(v1))) {
        case 0:
            fmt.Println(v1)
        case 9223372036854775807:
            var v46 int8 = -1
            _ = v46
        case -720:
            v2 -= v2
        }
    }
    if false {
        v47 := GenS6{F0: uint64(v3), F1: v1}
        v4 -= 42
        _ = v47
    } else {
        if len(v5) > 0 {
            v5[0] = 42
        }
        var v48 bool = true
        fmt.Println(v6, v48, v4)
        _ = v48
    }
    v49 := GenS7{F0: "foo bar"}
    for fi50 := 0; fi50 < 3; fi50++ {
        v4 = v4
        v6 -= (v6 >> (((v3 - (v3 ^ 0))) & 63))
        _ = fi50
    }
    var v51 uint8 = v6
    v51 = 1
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v51
    _ = v6
}
