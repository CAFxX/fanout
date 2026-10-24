package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int; F1 int16; F2 int8 }
type GenS4 struct { F0 int16; F1 int8 }
func genFn0(p0 int64) uint32 {
    p0 = (p0 % ((9223372036854775807) | 1))
    v5 := []int{362}
    var v6 float64 = (float64(0.0) - 1e10)
    for fi7 := 0; fi7 < 17; fi7++ {
        switch (fi7 << ((0) & 63)) {
        case 32767:
            for fi8 := 0; fi8 < 6; fi8++ {
                _ = len(v5)
                _ = fi8
            }
        case -718:
            if (false && true) {
                var v9 string = "\n\t"
                _ = v9
            }
        default:
            v6 = v6
        }
        if true {
            if len(v5) > 0 {
                v5[0] = 0
            }
        }
        if false {
            var v10 float32 = 1e10
            _ = v10
        } else {
            var v11 float32 = (float32(float64(3.14159)) / 1.0)
            fmt.Println(v11)
            _ = v11
        }
        _ = fi7
    }
    v12 := GenS1{F0: 128}
    _ = p0
    _ = v12
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 uint) uint {
    for fi13 := 0; fi13 < 7; fi13++ {
        switch -48 {
        case 9223372036854775807:
            switch -896 {
            case 297:
                var v14 uint16 = 425
                _ = v14
            default:
                var v15 uint = p0
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 1; fi16++ {
            v17 := GenS2{F0: ((!false) && false), F1: ("foo bar" + "\n\t")}
            _ = fi16
            _ = v17
        }
        _ = fi13
    }
    var v18 uint = (p0 + p0)
    _ = p0
    _ = v18
    return 192
}

func main() {
    var v1 uint8 = 0
    var v2 string = "foo bar"
    var v3 string = ("hello" + (v2 + v2))
    v4 := make(map[int]string)
    var v19 uint32 = 1
    v19 = (v19 & (v19 * (v19 / ((v19) | 1))))
    var v20 bool = (genFn1(796) != 4294967295)
    var v21 int = 0
    v22 := GenS3{F0: (v21 >> ((18446744073709551615) & 63)), F1: -231, F2: 1}
    switch ("foo bar" + "") {
    case "hello":
        fmt.Println(v21, v21, v19)
    case "\n\t":
        var v23 float64 = ((float64(float32((1.5 * 0.0))) - 0.0) + 1e10)
        _ = v23
    case "a":
        if v20 {
            var v24 uint8 = (v1 << ((18446744073709551615) & 63))
            if v20 {
                var v25 int32 = int32(v19)
                v1 = v24
                if (v24 > v24) {
                    fmt.Println(v25, v2, v1)
                }
                _ = v25
            } else {
                var v26 uint16 = 1
                _ = v26
            }
            _ = v24
        }
    default:
        v27 := GenS4{F0: -32768, F1: -103}
        _ = v27
    }
    if v20 {
        v28 := []int{v21, v21}
        _ = v28
    }
    v29 := []int{v21, v21}
    var v30 uint = genFn1(611)
    for fi31 := 0; fi31 < 3; fi31++ {
        v20 = (v20 && true)
        fmt.Println(fi31)
        _ = fi31
    }
    var v32 string = "foo bar"
    fmt.Println(v20, v30, v20)
    for fi33 := 0; fi33 < 17; fi33++ {
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 uint16 = 255
            fmt.Println(fi33)
            switch (v3 + ((v32 + v32) + (("foo bar" + ("foo bar" + v32)) + ("foo bar" + "hello")))) {
            case "hello":
                v36 := GenS2{F0: (!v20), F1: (v2 + ("" + "a"))}
                _ = v36
            default:
                v1 = v1
            }
            _ = fi34
            _ = v35
        }
        switch -721 {
        case 32768:
            for fi37 := 0; fi37 < 4; fi37++ {
                var v38 int = (fi33 / ((fi37) | 1))
                _ = fi37
                _ = v38
            }
        case 0:
            v39 := GenS4{F0: -944, F1: -104}
            _ = v39
        case 949:
            v2 = (v3 + "hello")
        }
        _ = fi33
    }
    var v40 uint32 = uint32(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v40
}
