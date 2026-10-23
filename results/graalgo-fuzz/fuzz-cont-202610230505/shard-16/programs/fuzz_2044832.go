package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int8; F2 uint64 }
func genFn0(p0 uint64) int32 {
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(fi7, fi7)
        _ = fi7
    }
    p0 += 18446744073709551615
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := []int{}
            var v11 uint32 = 931
            var v12 string = "foo bar"
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        var v13 uint32 = 2147483648
        _ = fi8
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint = 4294967295
        _ = fi14
        _ = v15
    }
    var v16 float64 = (-2.25 - ((3.14159 / ((-2.25 * -2.25) + 3.14159)) - (0.0 / 3.14159)))
    _ = p0
    _ = v16
    return -1
}
func genFn1() int16 {
    if ((0 >= 0) && (49 != 49)) {
        var v17 uint8 = 42
        v18 := GenS1{F0: -1}
        for fi19 := 0; fi19 < 8; fi19++ {
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 uint8 = 1
                v22 := GenS2{F0: 644}
                v17 += (^(v21))
                _ = fi20
                _ = v21
                _ = v22
            }
            if ((-(fi19)) > (fi19 >> ((18446744073709551615) & 63))) {
                v23 := []int{(fi19 - fi19), (fi19 % (((fi19 + fi19)) | 1)), 127}
                _ = v23
            } else {
                v17 = v17
            }
            v24 := []int{-517, (fi19 / ((fi19) | 1)), 443}
            _ = fi19
            _ = v24
        }
        _ = v17
        _ = v18
    } else {
        var v25 float32 = (-2.25 * ((1.5 + (1.5 + 1e10)) - (3.14159 - ((0.1 - -2.25) - 1e10))))
        v26 := []int{32767, -1, 1}
        _ = v25
        _ = v26
    }
    switch 128 {
    case -9223372036854775808:
        v27 := []int{}
        _ = v27
    }
    v28 := GenS2{F0: 18446744073709551615}
    _ = v28
    return 283
}

func main() {
    var v1 uint32 = 2
    var v2 int = 0
    v3 := make(map[uint32]int)
    var v4 uint32 = 1
    v5 := make(map[int]uint32)
    v6 := make(map[string]int)
    if ((v2 >> ((65535) & 63)) <= (v2 << ((695) & 63))) {
        v1 = v4
        var v29 int8 = 127
        _ = v29
    }
    var v30 string = ""
    for fi31 := 0; fi31 < 3; fi31++ {
        fmt.Println(v1, v30)
        v32 := GenS1{F0: 2147483647}
        _ = fi31
        _ = v32
    }
    var v33 float32 = (3.14159 + 3.14159)
    v34 := GenS1{F0: -1}
    if true {
        v35 := GenS3{F0: (v2 ^ (^(v2)))}
        _ = v35
    } else {
        var v36 int16 = 1
        v2 -= (v2 << ((42) & 63))
        v37 := []int{}
        _ = v36
        _ = v37
    }
    var v38 int = (v2 / (((v2 << ((4294967295) & 63))) | 1))
    for fi39 := 0; fi39 < 14; fi39++ {
        v4 = v1
        v40 := GenS4{F0: 803, F1: int16(v1), F2: 69}
        if true {
            var v41 int64 = 9223372036854775807
            _ = v41
        }
        _ = fi39
        _ = v40
    }
    if false {
        var v42 string = v30
        v43 := GenS3{F0: v38}
        if false {
            var v44 int32 = 1
            v45 := []int{-9223372036854775808, v38, 734, v2}
            _ = v44
            _ = v45
        } else {
            var v46 bool = true
            _ = v46
        }
        _ = v42
        _ = v43
    }
    if false {
        fmt.Println(v2)
    } else {
        if (true || (("" <= "") && false)) {
            if true {
                var v47 uint8 = 1
                v48 := v5[-9223372036854775808]
                fmt.Println(v1)
                _ = v47
                _ = v48
            }
            v30 = "foo bar"
            for fi49 := 0; fi49 < 15; fi49++ {
                v50 := []int{(fi49 * v38), v2}
                _ = fi49
                _ = v50
            }
        }
        var v51 int8 = -1
        _ = v51
    }
    if true {
        v52 := GenS2{F0: 18446744073709551615}
        _ = v52
    }
    switch v2 {
    case 0:
        v2 -= 9223372036854775807
    case 9223372036854775807:
        var v53 uint = 2
        _ = v53
    default:
        v4 -= 65535
    }
    v54 := GenS5{F0: 2, F1: -127, F2: 818}
    switch (v30 + "") {
    case "a":
        if (503 >= 32767) {
            fmt.Println(v4, v4, v38)
            var v55 uint32 = v4
            _ = v55
        }
    }
    v56 := []int{v38}
    switch "\n\t" {
    case "hello":
        var v57 uint8 = uint8(v1)
        _ = v57
    default:
        if (false && false) {
            v58 := GenS1{F0: -571}
            for fi59 := 0; fi59 < 13; fi59++ {
                var v60 uint64 = 18446744073709551615
                _ = fi59
                _ = v60
            }
            _ = v58
        }
    }
    v61 := GenS2{F0: 42}
    _ = len(v56)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v38
    _ = v4
    _ = v5
    _ = v54
    _ = v56
    _ = v6
    _ = v61
}
