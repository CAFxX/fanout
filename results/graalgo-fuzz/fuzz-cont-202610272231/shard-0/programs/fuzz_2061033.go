package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 uint32; F1 uint8; F2 string }
func genFn0(p0 uint8) uint64 {
    var v5 float64 = 1.5
    switch -9223372036854775808 {
    case -553:
        var v6 uint8 = (p0 | p0)
        _ = v6
    default:
        for fi7 := 0; fi7 < 4; fi7++ {
            v5 = (((0.1 + float64(0.1)) * float64(1e10)) + float64(0.0))
            switch (-(fi7)) {
            case 32768:
                v5 = 0.0
            case 127:
                var v8 int32 = -2147483648
                _ = v8
            case -466:
                var v9 string = "foo bar"
                _ = v9
            }
            p0 = p0
            _ = fi7
        }
    }
    v10 := []int{0}
    if false {
        v11 := GenS1{F0: (136 > 2)}
        _ = v11
    }
    p0 = p0
    _ = p0
    _ = v10
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 int16, p1 uint16) int64 {
    if false {
        v12 := []int{}
        _ = v12
    }
    switch ("" + "\n\t") {
    case "\n\t":
        var v13 float32 = ((-2.25 * 3.14159) * -2.25)
        _ = v13
    case "a":
        var v14 int8 = 127
        _ = v14
    }
    p0 = p0
    switch "\n\t" {
    case "hello":
        if true {
            if true {
                p1 += 65535
                if false {
                    var v15 uint16 = p1
                    v16 := []int{0}
                    if true {
                        var v17 float64 = 0.0
                        _ = v17
                    } else {
                        if (-2147483648 == -1) {
                            _ = len(v16)
                        }
                    }
                    _ = v15
                    _ = v16
                } else {
                    p1 = 42
                    var v18 uint64 = 4294967295
                    if (-2.25 > float32(3.14159)) {
                        v19 := GenS2{F0: -792, F1: (p0 * (-(p0)))}
                        _ = v19
                    }
                    _ = v18
                }
                for fi20 := 0; fi20 < 4; fi20++ {
                    fmt.Println(p0)
                    for fi21 := 0; fi21 < 4; fi21++ {
                        var v22 float32 = float32(float64((3.14159 + (1e10 / (0.1 / ((3.14159 * -2.25) * 1e10))))))
                        _ = fi21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 15; fi23++ {
                        v24 := GenS3{F0: 298, F1: 2147483648}
                        var v25 float64 = 0.1
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = fi20
                }
            }
        } else {
            if ((p1 < (p1 * (p1 + 32767))) && ((!(227 > 255)) && true)) {
                if true {
                    p0 += p0
                    v26 := GenS1{F0: (34 != 127)}
                    var v27 int64 = -622
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(p0)
                    p1 += 42
                    var v28 uint8 = 128
                    _ = v28
                }
            } else {
                var v29 uint64 = 1
                p0 -= (p0 | p0)
                _ = v29
            }
            var v30 uint = 447
            p1 = (p1 >> ((18446744073709551615) & 63))
            _ = v30
        }
    case "":
        if false {
            v31 := []int{}
            for fi32 := 0; fi32 < 11; fi32++ {
                fmt.Println(fi32, fi32, p1)
                if len(v31) >= 2 {
                    v31 = v31[1:len(v31)]
                }
                fmt.Println(p1, p1, fi32)
                _ = fi32
            }
            v33 := GenS2{F0: 48, F1: p0}
            _ = v31
            _ = v33
        }
    case "\n\t":
        var v34 uint = 2
        _ = v34
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 string = "hello"
    v2 := []uint32{2147483647, 1, 2, 4294967295}
    var v3 uint32 = 2147483647
    var v4 string = v1
    v4 = "a"
    v35 := GenS3{F0: 0, F1: (v3 & (^(v3)))}
    var v36 float64 = 1e10
    v1 = v1
    v1 = ((v1 + v4) + v1)
    if (!(1 <= 2147483647)) {
        v1 = v1
        var v37 int8 = 127
        for fi38 := 0; fi38 < 6; fi38++ {
            var v39 uint = 1
            _ = fi38
            _ = v39
        }
        _ = v37
    }
    v40 := GenS4{F0: v4, F1: 2147483647}
    if true {
        v41 := GenS2{F0: 0, F1: 1000}
        _ = v41
    } else {
        switch -2147483648 {
        case -1:
            fmt.Println(v4)
        case 674:
            fmt.Println(v4)
        }
        _ = len(v2)
    }
    var v42 int8 = -128
    v2 = append(v2, (v3 * v3))
    var v43 uint16 = 2
    if len(v2) > 0 {
        v2[0] = v3
    }
    v44 := GenS5{F0: 4294967295, F1: 2, F2: v4}
    var v45 uint16 = (v43 ^ (^(v43)))
    v1 = ("foo bar" + ("\n\t" + v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
