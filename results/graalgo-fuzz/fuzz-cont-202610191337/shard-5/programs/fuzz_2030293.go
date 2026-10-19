package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 uint64; F1 int32; F2 uint32 }
func genFn0() string {
    var v4 float64 = 3.14159
    var v5 float32 = (0.1 / (1.5 + 3.14159))
    v6 := []int{4294967295}
    _ = len(v6)
    v5 += v5
    _ = v4
    _ = v5
    _ = v6
    return "hello"
}
func genFn1(p0 string, p1 string) int8 {
    p0 = p0
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: 0}
        p0 = p1
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            p0 = p1
            p0 = p0
            var v11 int = (-(fi10))
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v12 := GenS2{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v12
    return 127
}

func main() {
    var v1 uint8 = 2
    var v2 uint = 1
    var v3 int = 124
    v13 := GenS3{F0: "\n\t", F1: 18446744073709551615}
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := []int{}
        if len(v15) > 0 {
            v15[0] = v3
        }
        _ = fi14
        _ = v15
    }
    v3 = 127
    var v16 int16 = -885
    v2 += (v2 + v2)
    v17 := GenS4{F0: 255, F1: v16}
    v1 += v1
    v1 = (v1 >> ((v2) & 63))
    if (!false) {
        var v18 int32 = -1
        v18 = (v18 << ((v2) & 63))
        _ = v18
    } else {
        if true {
            v16 = int16(v1)
            v19 := GenS5{F0: 4294967295, F1: 0}
            v20 := GenS3{F0: ((("\n\t" + "foo bar") + "foo bar") + "hello"), F1: 993}
            _ = v19
            _ = v20
        }
        var v21 string = "a"
        v22 := []int{v3, 942}
        _ = v21
        _ = v22
    }
    switch ("hello" + "foo bar") {
    case "a":
        fmt.Println(v1)
    case "":
        switch v3 {
        case -1:
            switch "a" {
            case "":
                var v23 bool = false
                _ = v23
            default:
                if false {
                    var v24 uint64 = 2
                    v25 := GenS6{F0: (v24 + 791), F1: 49, F2: 1}
                    v24 += (v24 * 18446744073709551615)
                    _ = v24
                    _ = v25
                }
            }
        case -777:
            v2 += (^(v2))
        case -9223372036854775808:
            v1 = 2
        }
    case "hello":
        if true {
            v26 := []int{(v3 >> ((v2) & 63)), (v3 + v3), v3, v3}
            var v27 float32 = 1e10
            _ = v26
            _ = v27
        }
    default:
        var v28 uint64 = 986
        _ = v28
    }
    v29 := GenS4{F0: 2, F1: 0}
    fmt.Println(v2)
    switch "\n\t" {
    case "hello":
        var v30 int16 = 568
        _ = v30
    }
    var v31 int16 = (^(v16))
    var v32 uint = v2
    v33 := []int{v3, v3, (v3 % ((int(v32)) | 1))}
    if false {
        switch v3 {
        case 9223372036854775807:
            fmt.Println(v16, v2, v16)
        case -2147483648:
            v31 -= v31
        default:
            v3 = v3
        }
        var v34 uint16 = 2
        _ = v34
    } else {
        for fi35 := 0; fi35 < 2; fi35++ {
            if (2 >= 2) {
                if (100 >= -1) {
                    fmt.Println(v16, fi35)
                }
                var v36 bool = true
                _ = v36
            }
            _ = fi35
        }
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        for fi37 := 0; fi37 < 13; fi37++ {
            var v38 int = v3
            for fi39 := 0; fi39 < 10; fi39++ {
                var v40 int = -9223372036854775808
                _ = fi39
                _ = v40
            }
            _ = fi37
            _ = v38
        }
    }
    v31 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
