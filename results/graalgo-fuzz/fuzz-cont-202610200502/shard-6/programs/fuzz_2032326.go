package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int; F1 int64; F2 int16 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint64, p1 bool) uint {
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(p0, p1)
        _ = fi7
    }
    switch ("hello" + ("" + "")) {
    case "hello":
        p0 = 18446744073709551615
    }
    v8 := []int{42, -1, -52, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v8
    return 579
}

func main() {
    var v1 int8 = 0
    var v2 int32 = 0
    v3 := []int{42, -1}
    v4 := []int{32767, 32767, 114, -9223372036854775808}
    var v5 string = "\n\t"
    v6 := []uint32{}
    var v9 uint = genFn0(1, false)
    if false {
        for fi10 := 0; fi10 < 6; fi10++ {
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 float64 = 0.1
                fmt.Println(fi10, fi11)
                _ = fi11
                _ = v12
            }
            var v13 uint = v9
            var v14 int64 = -9223372036854775808
            _ = fi10
            _ = v13
            _ = v14
        }
    }
    switch (("a" + v5) + (v5 + ("foo bar" + (v5 + (("hello" + v5) + ("hello" + ((((v5 + v5) + (("a" + (v5 + v5)) + "hello")) + v5) + ((v5 + v5) + v5)))))))) {
    case "hello":
        v15 := GenS1{F0: 0, F1: v9}
        _ = v15
    default:
        var v16 uint = 0
        _ = v16
    }
    var v17 uint32 = 40
    for fi18 := 0; fi18 < 1; fi18++ {
        v19 := GenS2{F0: uint8(v9), F1: 127}
        v4 = append(v4, (^(fi18)))
        _ = fi18
        _ = v19
    }
    var v20 uint64 = 18446744073709551615
    if ((true && ((v1 >> ((genFn0(v20, false)) & 63)) > 3)) && false) {
        var v21 int8 = 0
        v9 = (v9 & (^(v9)))
        _ = v21
    }
    fmt.Println(v2)
    v17 -= (v17 - v17)
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 int32 = (-(v2))
        v17 += (^(v17))
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        v17 += 0
        var v25 uint8 = 0
        fmt.Println(v5)
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            if false {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            } else {
                if true {
                    v28 := GenS1{F0: 280, F1: v9}
                    switch (v5 + v5) {
                    case "\n\t":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    case "foo bar":
                        var v29 string = v5
                        _ = v29
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = len(v3)
                    _ = v28
                }
                var v30 int32 = v2
                if (-2.25 <= float64((0.1 * (0.1 - (0.0 / (3.14159 * (1e10 + 1.5))))))) {
                    v31 := GenS3{F0: -680, F1: int64(v9), F2: -32768}
                    var v32 float32 = -2.25
                    var v33 float32 = v32
                    v30 = v30
                    v2 -= v30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = v30
            }
            _ = fi27
        }
        v34 := GenS3{F0: (^(fi26)), F1: 765, F2: 1000}
        _ = fi26
        _ = v34
    }
    _ = len(v3)
    switch "\n\t" {
    case "\n\t":
        v17 = v17
    case "hello":
        fmt.Println(v1, v5, v17)
    case "a":
        var v35 int32 = (v2 % (((v2 & (v2 ^ v2))) | 1))
        _ = v35
    }
    switch v5 {
    case "foo bar":
        fmt.Println(v5, v1)
    case "hello":
        v36 := GenS1{F0: 32767, F1: v9}
        _ = v36
    default:
        if (2 > 334) {
            switch v5 {
            case "a":
                var v37 float32 = 3.14159
                _ = v37
            case "\n\t":
                if (v5 <= "a") {
                    var v38 string = ("\n\t" + "\n\t")
                    _ = v38
                } else {
                    if (v17 > (v17 & 950)) {
                        var v39 int = 32768
                        _ = v39
                    }
                }
            case "hello":
                v1 = v1
            }
            fmt.Println(v2)
            v40 := GenS4{F0: 128, F1: 888}
            _ = v40
        } else {
            switch v5 {
            case "hello":
                if ((genFn0((v20 % (((v20 + 4294967295)) | 1)), true) == (v9 / (((^(v9))) | 1))) && (v2 != (v2 | 42))) {
                    switch -2147483648 {
                    case 127:
                        v17 -= (v17 | (v17 & (v17 * v17)))
                    case 32768:
                        v41 := GenS3{F0: 490, F1: -277, F2: 1000}
                        _ = v41
                    case 128:
                        v2 = (v2 / (((v2 * (v2 / ((v2) | 1)))) | 1))
                    }
                    var v42 int32 = v2
                    var v43 int = 4294967295
                    _ = v42
                    _ = v43
                } else {
                    var v44 uint64 = v20
                    var v45 float32 = -2.25
                    _ = v44
                    _ = v45
                }
            case "":
                v2 = -168
            case "a":
                v46 := GenS2{F0: 0, F1: 105}
                _ = v46
            }
            v47 := GenS5{F0: true}
            if len(v6) > 0 {
                v6[0] = 42
            }
            _ = v47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
