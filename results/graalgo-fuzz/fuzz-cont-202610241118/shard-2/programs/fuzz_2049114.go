package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int64; F2 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 string; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS7 struct { F0 int64; F1 int8 }
func genFn0(p0 uint, p1 uint) string {
    var v4 uint64 = 4294967295
    switch 42 {
    case -1:
        v5 := GenS1{F0: 105, F1: 42, F2: 65535}
        _ = v5
    default:
        v6 := GenS2{F0: 32767}
        _ = v6
    }
    var v7 int64 = 42
    if ((("hello" + "") + "a") <= "hello") {
        switch -9223372036854775808 {
        case 4294967295:
            for fi8 := 0; fi8 < 11; fi8++ {
                v4 = v4
                for fi9 := 0; fi9 < 14; fi9++ {
                    switch (fi9 + (fi8 - 128)) {
                    case -1:
                        if (fi9 < (-(fi8))) {
                            v7 = (v7 % ((-1) | 1))
                            var v10 uint8 = 128
                            v11 := []int{-2147483648, fi8}
                            _ = v10
                            _ = v11
                        }
                    case 2147483647:
                        var v12 string = (("hello" + "foo bar") + "\n\t")
                        _ = v12
                    case -9223372036854775808:
                        switch "foo bar" {
                        case "\n\t":
                            p0 = (p0 / ((p1) | 1))
                        default:
                            var v13 float32 = 1.5
                            _ = v13
                        }
                    }
                    var v14 uint32 = 42
                    var v15 bool = (-32768 >= 1)
                    fmt.Println(fi8, v14)
                    var v16 uint64 = 0
                    _ = fi9
                    _ = v14
                    _ = v15
                    _ = v16
                }
                _ = fi8
            }
        case 128:
            v17 := GenS3{F0: 255, F1: (v7 | v7), F2: (v7 ^ 42)}
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        p0 += (^(p1))
        v19 := []int{(fi18 - fi18)}
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return "hello"
}
func genFn1(p0 string) bool {
    v20 := GenS2{F0: 42}
    v21 := []int{32767, 42}
    v22 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 42}
    for fi23 := 0; fi23 < 5; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            v25 := GenS4{F0: 65535}
            p0 = (p0 + p0)
            _ = fi24
            _ = v25
        }
        var v26 float32 = 1e10
        fmt.Println(p0, p0, v26)
        _ = fi23
        _ = v26
    }
    _ = p0
    _ = v20
    _ = v21
    _ = v22
    return false
}

func main() {
    var v1 int16 = -1
    var v2 int32 = 0
    var v3 int16 = v1
    if false {
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 uint8 = 2
            _ = fi27
            _ = v28
        }
        if (false && (0.1 == (0.1 / 3.14159))) {
            switch -2147483648 {
            case 457:
                v29 := GenS3{F0: 255, F1: 9223372036854775807, F2: -677}
                _ = v29
            case -2147483648:
                v2 -= v2
            }
            v30 := GenS2{F0: 32768}
            _ = v30
        }
        var v31 uint64 = 2
        _ = v31
    }
    fmt.Println(v2, v1)
    v32 := []int{}
    _ = len(v32)
    for fi33 := 0; fi33 < 3; fi33++ {
        v3 = v1
        switch 32767 {
        case -2147483648:
            switch fi33 {
            case 627:
                v34 := GenS5{F0: 0, F1: "", F2: 885}
                _ = v34
            }
        case -9223372036854775808:
            if true {
                var v35 int16 = v3
                fmt.Println(v2)
                _ = v35
            } else {
                for fi36 := 0; fi36 < 15; fi36++ {
                    v37 := GenS2{F0: (fi33 % (((-(fi36))) | 1))}
                    v38 := GenS3{F0: 772, F1: 9223372036854775807, F2: 265}
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                v1 = v3
            }
        default:
            v39 := GenS6{F0: 32767, F1: 284, F2: v2}
            _ = v39
        }
        _ = fi33
    }
    var v40 int32 = (v2 * v2)
    var v41 int32 = (v40 << ((65535) & 63))
    for fi42 := 0; fi42 < 13; fi42++ {
        for fi43 := 0; fi43 < 19; fi43++ {
            v41 -= int32(v1)
            if (-1 < 25) {
                v32 = append(v32, 4294967295)
            }
            _ = fi43
        }
        _ = fi42
    }
    if false {
        var v44 uint64 = 4294967295
        v45 := GenS3{F0: 0, F1: -858, F2: -9223372036854775808}
        _ = v44
        _ = v45
    } else {
        var v46 int = 2147483647
        for fi47 := 0; fi47 < 17; fi47++ {
            var v48 uint = 1
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    var v49 string = "a"
    switch (((v49 + v49) + "hello") + v49) {
    case "\n\t":
        v49 = (v49 + v49)
    case "hello":
        for fi50 := 0; fi50 < 7; fi50++ {
            var v51 uint64 = 42
            _ = fi50
            _ = v51
        }
    default:
        v52 := GenS1{F0: 758, F1: 9223372036854775807, F2: 946}
        _ = v52
    }
    for fi53 := 0; fi53 < 1; fi53++ {
        var v54 int16 = (v3 * (^(v1)))
        v3 -= v54
        if (false && (-1 != (v40 / ((-1) | 1)))) {
            var v55 float32 = 1e10
            v32 = append(v32, fi53)
            _ = v55
        } else {
            if len(v32) > 0 {
                v32[0] = fi53
            }
        }
        _ = fi53
        _ = v54
    }
    for fi56 := 0; fi56 < 12; fi56++ {
        v57 := GenS4{F0: 719}
        if true {
            v1 = -32768
            v1 += (v1 << ((1) & 63))
        } else {
            v58 := GenS7{F0: 9223372036854775807, F1: 1}
            _ = v58
        }
        fmt.Println(v3, v3)
        _ = fi56
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v40
    _ = v41
    _ = v49
}
