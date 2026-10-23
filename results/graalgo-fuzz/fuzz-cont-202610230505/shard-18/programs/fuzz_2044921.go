package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int; F1 int8; F2 bool }
type GenS6 struct { F0 uint32; F1 uint32 }
func genFn0(p0 bool, p1 int64) int8 {
    switch 42 {
    case 707:
        p1 += (p1 << ((42) & 63))
    case -9223372036854775808:
        if false {
            if ((1 >= 1) || p0) {
                var v7 int8 = -128
                p0 = p0
                if (!((3.14159 * 1e10) >= 0.0)) {
                    v8 := []int{4294967295}
                    _ = v8
                } else {
                    var v9 uint = 42
                    p1 -= p1
                    _ = v9
                }
                _ = v7
            } else {
                switch ("hello" + "hello") {
                case "\n\t":
                    p1 += (p1 - (^(p1)))
                case "":
                    v10 := GenS1{F0: 56}
                    _ = v10
                }
                for fi11 := 0; fi11 < 7; fi11++ {
                    v12 := GenS2{F0: (0.0 <= ((0.1 / (1.5 + 1.5)) / 3.14159)), F1: p1}
                    if ("hello" <= "\n\t") {
                        var v13 uint = 42
                        v14 := []int{2147483647, fi11, fi11, (-(fi11))}
                        _ = v13
                        _ = v14
                    }
                    _ = fi11
                    _ = v12
                }
                for fi15 := 0; fi15 < 1; fi15++ {
                    var v16 int16 = -664
                    if true {
                        var v17 string = "hello"
                        _ = v17
                    } else {
                        var v18 string = ("hello" + "a")
                        p0 = p0
                        _ = v18
                    }
                    p1 += (-(p1))
                    _ = fi15
                    _ = v16
                }
            }
            var v19 uint8 = 1
            _ = v19
        }
    }
    switch ("a" + ("foo bar" + ("" + "hello"))) {
    case "foo bar":
        for fi20 := 0; fi20 < 5; fi20++ {
            v21 := GenS3{F0: (507 == 255)}
            _ = fi20
            _ = v21
        }
    case "\n\t":
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 uint32 = 28
            v24 := []int{fi22, fi22}
            _ = fi22
            _ = v23
            _ = v24
        }
    case "a":
        var v25 int16 = -1
        _ = v25
    default:
        var v26 float32 = (3.14159 + 1.5)
        _ = v26
    }
    v27 := []int{4294967295, -1, -1}
    if len(v27) > 0 {
        v27[0] = 32768
    }
    _ = p0
    _ = p1
    _ = v27
    return int8(p1)
}
func genFn1(p0 int8, p1 int64, p2 uint64) uint8 {
    var v28 uint8 = 75
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 uint32 = 1
        var v31 float32 = (((-2.25 + (1e10 + 1e10)) * (0.1 * (-2.25 - 1e10))) + ((3.14159 + 1e10) * 0.1))
        for fi32 := 0; fi32 < 12; fi32++ {
            v33 := []int{fi29, fi29}
            _ = fi32
            _ = v33
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    if false {
        for fi34 := 0; fi34 < 1; fi34++ {
            for fi35 := 0; fi35 < 11; fi35++ {
                v36 := []int{(fi34 << ((0) & 63)), (fi34 | 1)}
                var v37 uint16 = 384
                _ = fi35
                _ = v36
                _ = v37
            }
            fmt.Println(p2, p0, v28)
            v38 := GenS1{F0: v28}
            _ = fi34
            _ = v38
        }
        var v39 uint16 = 18
        _ = v39
    } else {
        v40 := []int{-881, 540, 4294967295}
        p2 = p2
        _ = v40
    }
    for fi41 := 0; fi41 < 9; fi41++ {
        v42 := GenS4{F0: (p1 % ((p1) | 1)), F1: "\n\t"}
        v43 := []int{fi41, fi41, fi41}
        _ = fi41
        _ = v42
        _ = v43
    }
    v44 := GenS5{F0: 463, F1: p0, F2: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v44
    return (^(v28))
}

func main() {
    var v1 uint8 = 128
    var v2 float32 = 3.14159
    v3 := make(map[string]bool)
    var v4 int8 = -128
    var v5 int16 = 32767
    v6 := make(map[int]string)
    var v45 uint32 = 2
    v46 := GenS6{F0: v45, F1: v45}
    if (v2 < (0.0 / v2)) {
        var v47 uint64 = 4294967295
        var v48 int32 = 42
        _ = v47
        _ = v48
    } else {
        var v49 uint = 255
        _ = v49
    }
    var v50 int64 = -9223372036854775808
    v51 := GenS3{F0: false}
    v45 = v45
    switch 0 {
    case 128:
        if false {
            var v52 uint8 = (v1 + genFn1((v4 ^ v4), -9223372036854775808, 18446744073709551615))
            _ = v52
        }
    default:
        delete(v6, 876)
    }
    v53 := GenS1{F0: v1}
    fmt.Println(v1, v4, v50)
    v54 := GenS4{F0: (v50 % ((0) | 1)), F1: ""}
    v55 := v3[""]
    v45 = v45
    v50 = v50
    var v56 int32 = 976
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v50
    _ = v51
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v6
}
