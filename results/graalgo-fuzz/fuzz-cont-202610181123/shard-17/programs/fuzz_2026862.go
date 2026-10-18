package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int32 }
func genFn0() int {
    var v7 float32 = 0.1
    v8 := []int{32767}
    if (false && (42 == 1)) {
        var v9 uint64 = 2
        if false {
            var v10 int32 = -2147483648
            _ = v10
        }
        v7 -= (v7 * 3.14159)
        _ = v9
    } else {
        v7 = 0.1
        var v11 uint = 4294967295
        _ = v11
    }
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 uint16, p1 int64) int8 {
    var v12 float32 = 3.14159
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := []int{genFn0(), (fi13 + (fi13 + (fi13 + (fi13 % (((fi13 >> ((uint(p0)) & 63))) | 1))))), fi13, fi13}
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: int8(p1), F1: false, F2: 700}
    var v16 int8 = 100
    _ = p0
    _ = p1
    _ = v12
    _ = v15
    _ = v16
    return (v16 * v16)
}

func main() {
    v1 := []int{843, 1, 127, 529}
    v2 := make(map[int]int)
    v3 := []string{("foo bar" + (("foo bar" + "a") + "\n\t")), "foo bar", ("foo bar" + ("" + ""))}
    v4 := make(map[uint32]uint32)
    v5 := []int{127, 128, -1, -9223372036854775808}
    var v6 float32 = ((1.5 / ((1e10 * float32((0.0 - -2.25))) - 1.5)) / -2.25)
    switch -257 {
    case 715:
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int8 = 42
            var v19 uint8 = 255
            fmt.Println(v6)
            _ = fi17
            _ = v18
            _ = v19
        }
    case 0:
        var v20 int64 = 42
        _ = v20
    case -1:
        switch "\n\t" {
        case "hello":
            v6 = float32((1.5 / 0.1))
        case "foo bar":
            delete(v4, 4294967295)
        default:
            v3 = append(v3, "hello")
        }
    }
    if false {
        switch -1 {
        case 9223372036854775807:
            fmt.Println(v6)
        case -1:
            for fi21 := 0; fi21 < 17; fi21++ {
                v6 = float32(0.1)
                v2[fi21] = fi21
                _ = fi21
            }
        }
        if (((1e10 / 1e10) - 1e10) <= 3.14159) {
            fmt.Println(v6)
            var v22 float64 = 1.5
            _ = v22
        }
        v23 := GenS1{F0: -1, F1: (v6 != 0.0), F2: 42}
        _ = v23
    } else {
        switch "" {
        case "hello":
            var v24 uint64 = 2
            _ = v24
        case "a":
            if true {
                switch "\n\t" {
                case "\n\t":
                    v25 := GenS1{F0: -1, F1: false, F2: 0}
                    _ = v25
                case "hello":
                    var v26 uint64 = 1
                    _ = v26
                case "a":
                    var v27 bool = false
                    _ = v27
                default:
                    switch "hello" {
                    case "\n\t":
                        v6 = v6
                    }
                }
            }
        case "foo bar":
            v6 -= -2.25
        default:
            v6 = 3.14159
        }
        switch "a" {
        case "foo bar":
            if (1 > -1) {
                for fi28 := 0; fi28 < 14; fi28++ {
                    if len(v3) > 0 {
                        v3[0] = ("\n\t" + "hello")
                    }
                    switch ("a" + "hello") {
                    case "foo bar":
                        v6 -= 1e10
                    case "hello":
                        v6 = (v6 / 1e10)
                    }
                    _ = fi28
                }
                var v29 uint32 = 4294967295
                var v30 int32 = 122
                _ = v29
                _ = v30
            } else {
                if false {
                    v6 = v6
                } else {
                    if (882 == 9223372036854775807) {
                        for fi31 := 0; fi31 < 13; fi31++ {
                            var v32 int64 = 9223372036854775807
                            fmt.Println(fi31, v6, fi31)
                            _ = fi31
                            _ = v32
                        }
                        v33 := GenS1{F0: -1, F1: false, F2: 741}
                        for fi34 := 0; fi34 < 17; fi34++ {
                            for fi35 := 0; fi35 < 15; fi35++ {
                                var v36 int64 = -1
                                if true {
                                    v37 := GenS1{F0: -128, F1: (v6 < (v6 / v6)), F2: 460}
                                    _ = v37
                                }
                                if ((fi34 - fi34) >= (fi34 % (((fi34 + -959)) | 1))) {
                                    v6 += 3.14159
                                }
                                _ = fi35
                                _ = v36
                            }
                            var v38 uint = 2
                            v39 := GenS2{F0: -248, F1: "a"}
                            _ = fi34
                            _ = v38
                            _ = v39
                        }
                        _ = v33
                    } else {
                        var v40 uint16 = 2
                        v6 = (v6 * -2.25)
                        _ = v40
                    }
                    var v41 float64 = ((1e10 - (((0.1 * 1.5) - 0.1) + 3.14159)) - 3.14159)
                    _ = v41
                }
            }
        case "hello":
            _ = len(v3)
        }
        v42 := GenS3{F0: 0}
        _ = v42
    }
    var v43 uint8 = 128
    var v44 float64 = 3.14159
    var v45 uint16 = 133
    var v46 float32 = v6
    var v47 int32 = int32(v45)
    v43 = v43
    var v48 int64 = 1
    fmt.Println(v47)
    switch ("a" + "\n\t") {
    case "":
        var v49 uint32 = 659
        _ = v49
    case "\n\t":
        switch "\n\t" {
        case "a":
            v47 -= 7
        case "":
            var v50 int = -715
            _ = v50
        case "hello":
            fmt.Println(v45, v47)
        }
    }
    for fi51 := 0; fi51 < 18; fi51++ {
        v52 := GenS2{F0: genFn0(), F1: ""}
        _ = fi51
        _ = v52
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 -= (0.1 - v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
