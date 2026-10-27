package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint8 }
func genFn0(p0 uint64, p1 uint16) uint32 {
    for fi7 := 0; fi7 < 16; fi7++ {
        p1 = (p1 << ((18446744073709551615) & 63))
        _ = fi7
    }
    v8 := []int{-9223372036854775808, 0, 4294967295}
    if false {
        var v9 int16 = 32767
        var v10 uint = 0
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: 127, F1: p0, F2: 42}
        _ = v11
    }
    switch "\n\t" {
    case "foo bar":
        p1 -= 918
    case "":
        _ = len(v8)
    case "a":
        v12 := GenS2{F0: -2147483648}
        _ = v12
    default:
        switch "a" {
        case "foo bar":
            p0 += p0
        case "":
            var v13 uint = 234
            _ = v13
        default:
            v14 := GenS2{F0: 4294967295}
            _ = v14
        }
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v8
    return 2147483648
}

func main() {
    var v1 float32 = 0.0
    var v2 int8 = 0
    v3 := []uint32{1, 65535}
    var v4 int64 = 0
    var v5 int8 = v2
    v6 := make(map[string]uint32)
    for fi15 := 0; fi15 < 4; fi15++ {
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        switch (fi15 & (fi15 | (fi15 ^ fi15))) {
        case 880:
            v2 -= (v5 + (v2 % ((v5) | 1)))
        case 1:
            var v16 uint64 = 18446744073709551615
            _ = v16
        }
        var v17 bool = (true && (false || (3.14159 <= 3.14159)))
        _ = fi15
        _ = v17
    }
    switch "\n\t" {
    case "a":
        if (!true) {
            fmt.Println(v5, v5)
        } else {
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 int64 = v4
                _ = fi18
                _ = v19
            }
            v2 -= (v5 | v5)
        }
    case "\n\t":
        var v20 bool = false
        _ = v20
    }
    if (((255 >= 32767) || false) || (false && (!((3.14159 != 0.1) && false)))) {
        if (4294967295 != 626) {
            v21 := GenS2{F0: 42}
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 bool = (!((fi22 * (-(fi22))) > -1))
                delete(v6, "")
                _ = fi22
                _ = v23
            }
            _ = v21
        } else {
            if false {
                v4 += 213
                var v24 string = (("foo bar" + "\n\t") + (("" + "hello") + "foo bar"))
                v5 -= (v5 >> ((42) & 63))
                _ = v24
            }
            delete(v6, "\n\t")
            for fi25 := 0; fi25 < 20; fi25++ {
                v4 -= (v4 % ((v4) | 1))
                _ = fi25
            }
        }
        v26 := GenS2{F0: 128}
        _ = v26
    }
    var v27 int8 = (v5 + 26)
    var v28 bool = (false && (!false))
    var v29 uint = 4294967295
    if false {
        var v30 float32 = 0.1
        delete(v6, "\n\t")
        v2 -= 0
        _ = v30
    }
    fmt.Println(v4, v5, v4)
    if false {
        v2 = (v5 * (v27 + 100))
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 int16 = -926
            var v33 string = "hello"
            _ = fi31
            _ = v32
            _ = v33
        }
        if ((v1 * float32(3.14159)) >= (v1 / (1e10 / v1))) {
            fmt.Println(v1, v1)
            v28 = v28
            fmt.Println(v4, v5)
        } else {
            if false {
                fmt.Println(v5, v29, v27)
                for fi34 := 0; fi34 < 11; fi34++ {
                    v35 := GenS2{F0: (-(fi34))}
                    var v36 int16 = 42
                    if false {
                        v37 := GenS3{F0: (v5 | (v27 - (v5 | v5))), F1: 4294967295, F2: 2}
                        _ = v37
                    } else {
                        v4 = (^(v4))
                    }
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                var v38 float64 = (0.0 / 1.0)
                _ = v38
            }
        }
    } else {
        fmt.Println(v27)
    }
    v39 := v6["\n\t"]
    var v40 float64 = (0.1 + (1e10 / -2.25))
    var v41 uint32 = (v39 ^ (v39 - (v39 + 0)))
    var v42 uint64 = 1
    for fi43 := 0; fi43 < 9; fi43++ {
        switch "" {
        case "":
            var v44 float64 = -2.25
            _ = v44
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = (v41 >> ((42) & 63))
            }
        case "a":
            var v45 uint8 = 128
            _ = v45
        default:
            v46 := GenS4{F0: 255}
            _ = v46
        }
        switch ("a" + "") {
        case "":
            for fi47 := 0; fi47 < 11; fi47++ {
                v5 = v5
                v48 := v6["\n\t"]
                _ = fi47
                _ = v48
            }
        default:
            v2 = (v2 ^ -1)
        }
        _ = fi43
    }
    v49 := GenS3{F0: v2, F1: v42, F2: 128}
    for fi50 := 0; fi50 < 20; fi50++ {
        v51 := GenS3{F0: v5, F1: (^(v42)), F2: 130}
        v52 := GenS5{F0: (v42 | (v42 / ((18446744073709551615) | 1))), F1: 42, F2: 1}
        _ = fi50
        _ = v51
        _ = v52
    }
    var v53 float64 = v40
    var v54 float32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v49
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
