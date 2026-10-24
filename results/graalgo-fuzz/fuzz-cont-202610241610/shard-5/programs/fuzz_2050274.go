package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
func genFn0(p0 int, p1 uint32) int8 {
    p0 += (^(p0))
    p1 += p1
    v4 := []int{p0}
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := GenS1{F0: fi6}
            _ = fi6
            _ = v7
        }
        var v8 bool = (((!true) || (1e10 < 0.0)) && false)
        _ = fi5
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1(p0 int16, p1 uint8) uint64 {
    var v9 int8 = 62
    v9 -= (v9 << ((255) & 63))
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS2{F0: (^(fi10)), F1: 4294967295}
        v12 := []int{(fi10 + (fi10 + -670))}
        var v13 int64 = -1
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    switch 1 {
    case 0:
        switch "\n\t" {
        case "hello":
            var v14 uint64 = 0
            _ = v14
        case "a":
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := []int{(fi15 * (fi15 % (((fi15 / (((fi15 + fi15)) | 1))) | 1))), (fi15 | (fi15 / ((fi15) | 1))), 127}
                p0 = 0
                var v17 int32 = 470
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    case 32767:
        v9 -= (v9 / (((v9 | (v9 << ((18446744073709551615) & 63)))) | 1))
    }
    var v18 int = -2147483648
    _ = p0
    _ = p1
    _ = v18
    _ = v9
    return 2
}

func main() {
    var v1 float32 = -2.25
    var v2 int32 = -2147483648
    v3 := []string{"\n\t", ("hello" + "a"), "\n\t", "foo bar"}
    var v19 int64 = 9223372036854775807
    switch 127 {
    case -992:
        v1 = v1
    case -737:
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 float64 = (-2.25 + -2.25)
            v21 = (v21 / 1e10)
            v19 -= (v19 - -628)
            _ = fi20
            _ = v21
        }
    default:
        var v22 int = -895
        _ = v22
    }
    var v23 int16 = -32768
    var v24 uint32 = 457
    var v25 bool = ((false || (-128 != -29)) && ((true || false) || false))
    var v26 float64 = (0.0 * (-2.25 - -2.25))
    for fi27 := 0; fi27 < 12; fi27++ {
        var v28 float32 = 0.0
        v1 += -2.25
        _ = fi27
        _ = v28
    }
    switch 481 {
    case 268:
        v3 = append(v3, "a")
    default:
        _ = len(v3)
    }
    if v25 {
        v29 := GenS1{F0: 32767}
        switch "hello" {
        case "a":
            v1 = v1
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = "hello"
            }
        }
        v30 := GenS2{F0: -845, F1: (v24 << ((18446744073709551615) & 63))}
        _ = v29
        _ = v30
    } else {
        var v31 int8 = 121
        if len(v3) > 0 {
            v3[0] = "a"
        }
        v26 += float64(v1)
        _ = v31
    }
    if v25 {
        v32 := GenS3{F0: (v24 | 0)}
        v2 = 1
        _ = v32
    } else {
        fmt.Println(v26, v1, v26)
        for fi33 := 0; fi33 < 3; fi33++ {
            fmt.Println(v2)
            if len(v3) > 0 {
                v3[0] = "a"
            }
            _ = fi33
        }
        if true {
            switch -124 {
            case 2147483647:
                v19 = 42
            case 9223372036854775807:
                var v34 int32 = (^(v2))
                _ = v34
            case 32767:
                v2 -= (v2 >> ((924) & 63))
            }
            v35 := GenS4{F0: "\n\t"}
            _ = v35
        } else {
            fmt.Println(v2, v24, v1)
        }
    }
    v24 = (v24 % ((v24) | 1))
    for fi36 := 0; fi36 < 13; fi36++ {
        v25 = (127 > 52)
        switch "hello" {
        case "":
            v25 = v25
        case "\n\t":
            var v37 int16 = 42
            _ = v37
        case "foo bar":
            fmt.Println(fi36)
        }
        _ = fi36
    }
    if (!(v25 || (true && true))) {
        for fi38 := 0; fi38 < 5; fi38++ {
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
            for fi39 := 0; fi39 < 11; fi39++ {
                var v40 uint32 = 2
                _ = fi39
                _ = v40
            }
            for fi41 := 0; fi41 < 17; fi41++ {
                var v42 bool = ((false || (int8(v23) > 1)) && v25)
                v43 := GenS1{F0: fi41}
                fmt.Println(v19)
                _ = fi41
                _ = v42
                _ = v43
            }
            _ = fi38
        }
        if v25 {
            var v44 float32 = 1e10
            _ = v44
        }
        v26 += float64((-2.25 * v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
