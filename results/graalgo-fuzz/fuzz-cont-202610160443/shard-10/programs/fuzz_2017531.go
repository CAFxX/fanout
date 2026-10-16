package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 bool }
func genFn0(p0 uint16, p1 uint16) bool {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := []int{(-(fi6)), (^(fi6)), fi6}
        fmt.Println(p0, fi6)
        _ = fi6
        _ = v7
    }
    switch 42 {
    case 32768:
        if false {
            fmt.Println(p1, p1)
        }
    case 128:
        v8 := []int{-2147483648}
        _ = v8
    case 32767:
        var v9 int64 = -9223372036854775808
        _ = v9
    default:
        for fi10 := 0; fi10 < 16; fi10++ {
            if (uint32(p1) < 65535) {
                p1 = (p1 ^ (^(p0)))
                p1 -= p1
                p1 -= 607
            }
            v11 := GenS1{F0: 0, F1: (-2.25 < (1.5 / -2.25)), F2: 42}
            _ = fi10
            _ = v11
        }
    }
    if ("" > ("\n\t" + "hello")) {
        var v12 int64 = 1
        if true {
            if false {
                var v13 int32 = 521
                _ = v13
            } else {
                if false {
                    v14 := GenS1{F0: 2, F1: false, F2: 937}
                    v15 := []int{1}
                    p1 = (p1 + p1)
                    _ = v14
                    _ = v15
                } else {
                    fmt.Println(p1)
                    for fi16 := 0; fi16 < 14; fi16++ {
                        var v17 int = -1
                        v18 := []int{}
                        v18 = append(v18, (fi16 / ((-213) | 1)))
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    v19 := []int{}
                    _ = v19
                }
                fmt.Println(v12)
            }
            var v20 uint = 65535
            v12 -= (v12 / ((v12) | 1))
            _ = v20
        } else {
            v12 = 0
            v21 := []int{32767, -9223372036854775808}
            _ = v21
        }
        for fi22 := 0; fi22 < 10; fi22++ {
            v23 := []int{fi22}
            _ = fi22
            _ = v23
        }
        _ = v12
    }
    var v24 int16 = -368
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v24
    return (828 != 928)
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 201
    var v3 uint16 = 1
    var v4 int32 = 42
    v5 := []uint32{0, 1, 318}
    switch v2 {
    case 128:
        var v25 int64 = -9223372036854775808
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        v27 := v1[-744]
        _ = fi26
        _ = v27
    }
    var v28 uint = 416
    v3 = 32767
    for fi29 := 0; fi29 < 4; fi29++ {
        v30 := GenS2{F0: (!true)}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        switch ("hello" + ("hello" + "a")) {
        case "a":
            var v31 int64 = 42
            _ = v31
        case "\n\t":
            var v32 uint64 = 42
            _ = v32
        default:
            var v33 int64 = 9223372036854775807
            _ = v33
        }
        _ = fi29
        _ = v30
    }
    v34 := GenS3{F0: 369}
    switch ("a" + (("a" + ("" + ("a" + ("\n\t" + "\n\t")))) + "hello")) {
    case "\n\t":
        switch ("hello" + ("a" + "")) {
        case "foo bar":
            v35 := GenS1{F0: 66, F1: ((v2 >> ((v28) & 63)) > 9223372036854775807), F2: v28}
            _ = v35
        }
    }
    v4 = 660
    if (!(v2 >= -1)) {
        fmt.Println(v4, v2, v2)
        if (false || (false && false)) {
            if (-414 >= v2) {
                v4 += (v4 | v4)
            } else {
                for fi36 := 0; fi36 < 19; fi36++ {
                    v28 = (v28 | (v28 >> (((v28 ^ 43)) & 63)))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = len(v5)
                    _ = fi36
                }
                var v37 bool = false
                _ = v37
            }
            var v38 bool = false
            for fi39 := 0; fi39 < 13; fi39++ {
                v40 := GenS2{F0: false}
                for fi41 := 0; fi41 < 17; fi41++ {
                    var v42 uint16 = uint16(v28)
                    delete(v1, (^(v2)))
                    if ((1 < 32767) || v38) {
                        v4 += (v4 % ((v4) | 1))
                        fmt.Println(fi39, fi41, fi39)
                        if v38 {
                            v43 := GenS3{F0: 42}
                            var v44 uint8 = 139
                            fmt.Println(v4)
                            _ = v43
                            _ = v44
                        } else {
                            var v45 int16 = 1000
                            _ = v45
                        }
                    } else {
                        v46 := GenS4{F0: true, F1: (1 != 128)}
                        fmt.Println(v4)
                        _ = v46
                    }
                    _ = fi41
                    _ = v42
                }
                _ = fi39
                _ = v40
            }
            _ = v38
        }
        var v47 bool = (1.5 <= ((-2.25 * (((((-2.25 + (1.5 * float32(3.14159))) / 0.1) / 3.14159) + 1e10) + 3.14159)) + 0.1))
        _ = v47
    }
    v2 -= v2
    var v48 uint8 = 127
    var v49 uint8 = (v48 + v48)
    if true {
        delete(v1, v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v48
    _ = v49
    _ = v5
}
