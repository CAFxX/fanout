package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int16; F1 string; F2 uint64 }
type GenS4 struct { F0 int16; F1 int }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 int64 = -1
    switch -1 {
    case -1:
        for fi4 := 0; fi4 < 8; fi4++ {
            var v5 int64 = v3
            fmt.Println(v5, v5, fi4)
            _ = fi4
            _ = v5
        }
    case 9223372036854775807:
        v3 = (v3 + (v3 << ((286) & 63)))
    case -359:
        var v6 string = ("a" + ("a" + "a"))
        _ = v6
    }
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint64 = 586
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 int8 = 15
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(v3)
    if false {
        if (-188 < -2147483648) {
            v12 := GenS1{F0: 100}
            _ = v12
        }
        if ((false || ("" > (("\n\t" + "\n\t") + ("hello" + ("foo bar" + "foo bar"))))) && (31 <= 18446744073709551615)) {
            var v13 int64 = v3
            v13 -= v3
            _ = v13
        } else {
            v14 := []int{125, -556}
            if (!false) {
                switch -9223372036854775808 {
                case 672:
                    v3 = v3
                }
            }
            _ = v14
        }
    }
    v15 := GenS2{F0: "hello", F1: 32767}
    switch ("" + "") {
    case "a":
        for fi16 := 0; fi16 < 3; fi16++ {
            v3 = -1
            _ = fi16
        }
    case "hello":
        var v17 uint = 4294967295
        _ = v17
    case "":
        fmt.Println(v3)
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 uint64 = 42
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 5; fi21++ {
            v3 = (^(v3))
            _ = fi21
        }
        var v22 uint16 = 312
        _ = fi18
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 string = ""
        v25 := GenS2{F0: (v24 + (v24 + "foo bar")), F1: 0}
        _ = fi23
        _ = v24
        _ = v25
    }
    v26 := []int{128, 665}
    v3 = v3
    var v27 float32 = (0.0 - 1e10)
    var v28 bool = true
    v28 = v28
    if v28 {
        fmt.Println(v27, v3)
        switch -353 {
        case 259:
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
        case -1:
            for fi29 := 0; fi29 < 8; fi29++ {
                if (189 != -449) {
                    if (v28 && v28) {
                        v3 = (v3 << ((4294967295) & 63))
                        var v30 uint8 = 42
                        v31 := GenS3{F0: 409, F1: "", F2: 4294967295}
                        _ = v30
                        _ = v31
                    } else {
                        switch ("" + ("foo bar" + "a")) {
                        case "hello":
                            var v32 int = fi29
                            _ = v32
                        case "foo bar":
                            for fi33 := 0; fi33 < 16; fi33++ {
                                v28 = v28
                                v27 = (((0.0 / ((v27 / 1e10) / v27)) * ((v27 * ((-2.25 * 0.0) / 1e10)) - float32((3.14159 / 1.0)))) * v27)
                                _ = fi33
                            }
                        default:
                            v3 = v3
                        }
                        v34 := GenS4{F0: 0, F1: (-(fi29))}
                        fmt.Println(fi29)
                        _ = v34
                    }
                }
                _ = fi29
            }
        default:
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
        }
    } else {
        fmt.Println(v28)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v7
}
