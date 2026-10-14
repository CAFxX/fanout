package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
func genFn0() uint32 {
    if false {
        switch ((("\n\t" + "foo bar") + ("foo bar" + "\n\t")) + ("a" + "\n\t")) {
        case "hello":
            switch "foo bar" {
            case "a":
                var v5 string = "a"
                _ = v5
            case "hello":
                v6 := GenS1{F0: 255, F1: 42}
                _ = v6
            case "":
                var v7 bool = ((1 == 2) && false)
                _ = v7
            default:
                v8 := GenS2{F0: 32767}
                _ = v8
            }
        }
        v9 := []int{-703}
        _ = v9
    } else {
        var v10 float32 = (((0.1 - ((((1.5 / 0.1) * 0.1) - (float32(3.14159) * 1e10)) * ((0.0 + -2.25) - 3.14159))) * 0.1) - (((3.14159 * (1.5 - 1.5)) / (1.5 + (((1.5 - (((((-2.25 * (-2.25 / -2.25)) + ((0.0 + -2.25) * 0.1)) / ((1e10 * float32(((float64(((1e10 / (3.14159 + 1e10)) / float32(3.14159))) / 0.1) - -2.25))) - 1.5)) - 1.5) / float32(1e10))) + float32((0.1 + (0.0 * ((3.14159 * 3.14159) - (-2.25 + ((3.14159 / float64(1e10)) - (3.14159 * 0.0)))))))) + 1e10))) * (-2.25 - 1e10)))
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := GenS1{F0: 255, F1: 65}
            if ((0 >= -1) && true) {
                fmt.Println(fi11)
                fmt.Println(v10, fi11)
            } else {
                fmt.Println(fi11)
                var v13 bool = true
                var v14 float64 = 3.14159
                _ = v13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        v15 := []int{}
        _ = v10
        _ = v15
    }
    var v16 int64 = -9223372036854775808
    v17 := []int{1, 42, 32768, 128}
    _ = v16
    _ = v17
    return 543
}
func genFn1(p0 bool, p1 int32, p2 uint8) uint {
    p0 = false
    p2 = 90
    p2 += (^(p2))
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 float32 = 3.14159
    var v2 int16 = -1
    var v3 int16 = (v2 & -32768)
    var v4 string = ""
    v1 += (0.0 - -2.25)
    v18 := []int{32768, 2147483647, 4294967295}
    var v19 uint8 = 188
    fmt.Println(v4, v2)
    var v20 int = 403
    var v21 int64 = -1
    fmt.Println(v3)
    if false {
        if (32767 > 2) {
            if ((v4 + (("a" + (v4 + ("foo bar" + "\n\t"))) + v4)) < v4) {
                if (false && true) {
                    var v22 int32 = 42
                    var v23 int64 = 42
                    _ = v22
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 3; fi24++ {
                        v3 += (v2 + v3)
                        _ = fi24
                    }
                    fmt.Println(v4, v4, v19)
                    if false {
                        fmt.Println(v4)
                        var v25 int64 = 355
                        var v26 int64 = -9223372036854775808
                        _ = v25
                        _ = v26
                    }
                }
            }
            v27 := GenS1{F0: 941, F1: 42}
            v28 := GenS3{F0: (v19 >> ((1) & 63)), F1: v3}
            _ = v27
            _ = v28
        }
        v29 := GenS4{F0: 0}
        v1 -= v1
        _ = v29
    }
    v30 := GenS4{F0: 127}
    for fi31 := 0; fi31 < 5; fi31++ {
        if (-80 == 0) {
            var v32 int64 = (v21 | 1)
            fmt.Println(v32)
            fmt.Println(v21)
            _ = v32
        }
        v33 := GenS5{F0: (v21 ^ (v21 & (v21 << ((uint(v19)) & 63))))}
        if false {
            _ = len(v18)
        } else {
            if ((^(v21)) <= v21) {
                fmt.Println(v20)
                v34 := GenS3{F0: 199, F1: v3}
                v21 -= 1
                _ = v34
            } else {
                for fi35 := 0; fi35 < 15; fi35++ {
                    v36 := GenS1{F0: 42, F1: -106}
                    var v37 int32 = 113
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            }
        }
        _ = fi31
        _ = v33
    }
    v38 := GenS2{F0: (v20 % ((int(v19)) | 1))}
    if true {
        var v39 int16 = 0
        _ = v39
    }
    for fi40 := 0; fi40 < 20; fi40++ {
        if true {
            v1 = 1e10
        } else {
            v3 = (-(v3))
        }
        fmt.Println(v2, v19)
        _ = fi40
    }
    fmt.Println(v20)
    if (v3 <= v2) {
        for fi41 := 0; fi41 < 5; fi41++ {
            fmt.Println(v4)
            _ = fi41
        }
    }
    switch v20 {
    case 4294967295:
        var v42 uint8 = (v19 | 128)
        _ = v42
    case 127:
        if (!true) {
            v3 = (v2 % ((v2) | 1))
            var v43 uint32 = 42
            v21 += -9223372036854775808
            _ = v43
        } else {
            if false {
                if true {
                    _ = len(v18)
                    for fi44 := 0; fi44 < 16; fi44++ {
                        v19 -= v19
                        _ = fi44
                    }
                } else {
                    var v45 uint = 65535
                    _ = v45
                }
            } else {
                v46 := GenS4{F0: 0}
                _ = v46
            }
            v2 += v3
        }
    }
    for fi47 := 0; fi47 < 5; fi47++ {
        var v48 string = ((((v4 + "\n\t") + v4) + "") + v4)
        v21 = -694
        switch (fi47 << ((1) & 63)) {
        case 476:
            for fi49 := 0; fi49 < 16; fi49++ {
                v3 = -32768
                _ = fi49
            }
        case 215:
            for fi50 := 0; fi50 < 9; fi50++ {
                v51 := GenS2{F0: (^(v20))}
                _ = fi50
                _ = v51
            }
        default:
            v1 += 1e10
        }
        _ = fi47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v38
    _ = v4
}
