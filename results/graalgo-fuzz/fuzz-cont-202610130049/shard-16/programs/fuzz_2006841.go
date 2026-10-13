package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int64 }
type GenS2 struct { F0 uint; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 uint }
func genFn0(p0 int) bool {
    fmt.Println(p0)
    p0 = (p0 | (-(p0)))
    _ = p0
    return (1 < 1)
}
func genFn1(p0 int64) bool {
    switch "foo bar" {
    case "\n\t":
        for fi5 := 0; fi5 < 16; fi5++ {
            fmt.Println(p0)
            for fi6 := 0; fi6 < 14; fi6++ {
                p0 = p0
                v7 := []int{fi6, (fi6 | fi5)}
                if (!false) {
                    if true {
                        var v8 uint64 = 1
                        v8 = (v8 / ((v8) | 1))
                        _ = v8
                    } else {
                        fmt.Println(fi6, fi6)
                        var v9 uint64 = 18446744073709551615
                        _ = v9
                    }
                    switch fi5 {
                    case 274:
                        var v10 int = (fi6 >> ((434) & 63))
                        _ = v10
                    case 2147483647:
                        var v11 string = ("" + "hello")
                        _ = v11
                    }
                } else {
                    p0 += p0
                    switch ("a" + "") {
                    case "hello":
                        var v12 float32 = (float32(1e10) + (float32((float64((1.5 * float32(-2.25))) + (3.14159 * 0.0))) - -2.25))
                        _ = v12
                    case "":
                        if (65535 <= 614) {
                            var v13 int16 = 74
                            _ = v13
                        }
                    case "foo bar":
                        p0 += p0
                    }
                    fmt.Println(fi5, p0, fi5)
                    p0 = p0
                    switch fi6 {
                    case -984:
                        var v14 bool = (false && true)
                        _ = v14
                    }
                }
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
    }
    p0 += p0
    v15 := GenS1{F0: 61, F1: 1, F2: p0}
    _ = p0
    _ = v15
    return true
}

func main() {
    v1 := []string{"", ("foo bar" + "hello"), "\n\t"}
    v2 := make(map[uint32]string)
    var v3 int = -673
    var v4 int = -250
    if false {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        switch ((("" + "\n\t") + "\n\t") + "foo bar") {
        case "\n\t":
            if true {
                var v16 uint64 = 18446744073709551615
                var v17 string = "hello"
                _ = v16
                _ = v17
            }
        default:
            v4 = 4294967295
        }
        var v18 bool = false
        _ = v18
    } else {
        var v19 float32 = float32(float64(-2.25))
        if (!true) {
            v20 := GenS1{F0: 2, F1: 1, F2: 9223372036854775807}
            _ = v20
        }
        for fi21 := 0; fi21 < 14; fi21++ {
            if false {
                var v22 uint = 919
                _ = v22
            } else {
                if false {
                    v4 -= v4
                    var v23 uint = 1
                    _ = v23
                }
            }
            fmt.Println(v19, v19)
            _ = fi21
        }
        _ = v19
    }
    v24 := GenS1{F0: 0, F1: 42, F2: -1}
    var v25 uint32 = 2147483647
    for fi26 := 0; fi26 < 8; fi26++ {
        v27 := GenS2{F0: 0, F1: false, F2: 127}
        v3 = (v3 % (((fi26 & 32767)) | 1))
        _ = fi26
        _ = v27
    }
    if false {
        switch "foo bar" {
        case "hello":
            var v28 uint = 138
            _ = v28
        }
        if true {
            v25 = 65535
        }
    } else {
        v3 += v3
        v29 := GenS2{F0: 4294967295, F1: true, F2: 226}
        v30 := GenS1{F0: 51, F1: 100, F2: 9223372036854775807}
        _ = v29
        _ = v30
    }
    v4 = 32767
    v25 = (v25 ^ 4294967295)
    v4 = 127
    for fi31 := 0; fi31 < 3; fi31++ {
        v32 := GenS3{F0: 0}
        switch "foo bar" {
        case "hello":
            var v33 int8 = -128
            _ = v33
        default:
            v34 := GenS3{F0: 479}
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    v4 -= v4
    var v35 float64 = ((3.14159 + 1.5) / 1.0)
    var v36 float32 = (float32(1.5) + ((-2.25 - (-2.25 / 1.0)) * 3.14159))
    v25 -= (v25 % ((42) | 1))
    var v37 float64 = float64(v36)
    for fi38 := 0; fi38 < 5; fi38++ {
        fmt.Println(v37, fi38, v3)
        var v39 int8 = -1
        v4 = fi38
        _ = fi38
        _ = v39
    }
    v25 = 1
    for fi40 := 0; fi40 < 19; fi40++ {
        v4 -= (v3 >> ((883) & 63))
        for fi41 := 0; fi41 < 2; fi41++ {
            fmt.Println(v3, v35, v35)
            v42 := GenS4{F0: -9223372036854775808, F1: 293}
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    if (false && false) {
        for fi43 := 0; fi43 < 14; fi43++ {
            for fi44 := 0; fi44 < 17; fi44++ {
                if (2147483648 >= 4294967295) {
                    var v45 int64 = 1
                    _ = v45
                } else {
                    v4 += 0
                    v35 = -2.25
                    v1 = append(v1, "")
                    v46 := GenS2{F0: 4294967295, F1: (!(127 != 2)), F2: 154}
                    _ = v46
                }
                _ = fi44
            }
            var v47 uint = 4294967295
            _ = fi43
            _ = v47
        }
        fmt.Println(v3, v35, v4)
        v48 := v2[0]
        _ = v48
    } else {
        delete(v2, v25)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v35 -= (1.5 + (v37 * v37))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
