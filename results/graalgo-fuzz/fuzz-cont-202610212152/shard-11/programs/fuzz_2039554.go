package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint64; F1 string }
func genFn0(p0 string, p1 uint64, p2 int64) uint64 {
    for fi6 := 0; fi6 < 10; fi6++ {
        p0 = ("\n\t" + (p0 + "foo bar"))
        var v7 int32 = 2147483647
        _ = fi6
        _ = v7
    }
    switch (p0 + p0) {
    case "\n\t":
        v8 := GenS1{F0: 18446744073709551615}
        _ = v8
    case "":
        fmt.Println(p0, p2, p2)
    }
    if (true || false) {
        var v9 uint8 = 42
        p2 = (p2 << ((287) & 63))
        _ = v9
    } else {
        v10 := []int{-466, 912, 4294967295, 1}
        var v11 int32 = -2147483648
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1(p0 string, p1 uint32, p2 uint64) int32 {
    p1 = (p1 << ((2) & 63))
    fmt.Println(p2)
    for fi12 := 0; fi12 < 15; fi12++ {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(p2, fi12, fi12)
        default:
            p1 -= 42
        }
        fmt.Println(p1)
        v13 := []int{fi12}
        _ = fi12
        _ = v13
    }
    if (2147483647 == -824) {
        p0 = (p0 + "a")
        p2 = 475
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        if false {
            v15 := GenS1{F0: (^(p2))}
            fmt.Println(fi14, p2)
            _ = v15
        }
        p1 -= p1
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 1
    v3 := []int{128, 594, 2147483647}
    var v4 uint64 = 42
    v5 := []uint32{2147483648, 1, 1, 771}
    v4 -= (v4 << ((1) & 63))
    _ = len(v5)
    for fi16 := 0; fi16 < 11; fi16++ {
        if (!true) {
            v4 = 42
            v17 := GenS1{F0: genFn0("foo bar", (v4 % ((1) | 1)), -1)}
            var v18 int64 = 42
            _ = v17
            _ = v18
        }
        var v19 int = -9223372036854775808
        _ = fi16
        _ = v19
    }
    var v20 uint64 = genFn0("", uint64(v2), -9223372036854775808)
    fmt.Println(v20, v4)
    if true {
        fmt.Println(v2, v20, v4)
        v21 := GenS1{F0: 2}
        _ = v21
    } else {
        switch 1 {
        case 32768:
            switch "a" {
            case "":
                var v22 uint64 = (v4 | (v20 % (((^(v4))) | 1)))
                _ = v22
            case "hello":
                switch 9223372036854775807 {
                case -298:
                    delete(v1, -1)
                case 72:
                    var v23 string = "hello"
                    _ = v23
                case 1:
                    var v24 int64 = -703
                    _ = v24
                }
            case "a":
                for fi25 := 0; fi25 < 16; fi25++ {
                    for fi26 := 0; fi26 < 16; fi26++ {
                        var v27 uint64 = genFn0((("hello" + ("\n\t" + "hello")) + "\n\t"), (v20 + 0), 778)
                        switch (fi25 + fi25) {
                        case 2147483647:
                            v2 += 1
                        case 9223372036854775807:
                            var v28 float64 = -2.25
                            _ = v28
                        case -1:
                            v2 -= v2
                        }
                        _ = fi26
                        _ = v27
                    }
                    _ = fi25
                }
            }
        case 32767:
            var v29 string = "a"
            _ = v29
        case 898:
            if false {
                switch -829 {
                case 4294967295:
                    v4 += (v4 | v20)
                }
                v30 := GenS1{F0: 2}
                _ = v30
            }
        }
        for fi31 := 0; fi31 < 15; fi31++ {
            if len(v3) > 0 {
                v3[0] = fi31
            }
            _ = fi31
        }
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 uint32 = 2
            _ = fi32
            _ = v33
        }
    }
    var v34 uint16 = 32767
    v2 += (v2 ^ (v2 << ((255) & 63)))
    v35 := GenS2{F0: 72, F1: 2147483647}
    fmt.Println(v4, v4, v4)
    if true {
        if ((3.14159 * float32(((-2.25 - 3.14159) * 3.14159))) != 0.0) {
            for fi36 := 0; fi36 < 4; fi36++ {
                for fi37 := 0; fi37 < 13; fi37++ {
                    v1[(fi37 ^ (-(fi37)))] = ""
                    v20 += v4
                    fmt.Println(v20)
                    _ = fi37
                }
                _ = fi36
            }
        }
        fmt.Println(v4)
        v38 := GenS1{F0: uint64(v2)}
        _ = v38
    }
    if (!false) {
        if true {
            _ = len(v3)
        }
        switch "foo bar" {
        case "":
            v39 := GenS2{F0: v2, F1: 780}
            _ = v39
        case "hello":
            var v40 bool = true
            _ = v40
        }
        v34 = (v34 | (^(v34)))
    } else {
        if false {
            v41 := GenS3{F0: 1000}
            var v42 uint8 = 255
            _ = v41
            _ = v42
        } else {
            v43 := GenS4{F0: 0, F1: (v20 - genFn0(("" + ("foo bar" + ((("" + "a") + ("" + (("\n\t" + "") + ""))) + "a"))), 18446744073709551615, -10)), F2: v4}
            for fi44 := 0; fi44 < 13; fi44++ {
                var v45 uint8 = (v2 << ((42) & 63))
                _ = fi44
                _ = v45
            }
            if (0 == -715) {
                fmt.Println(v2, v2)
                v34 = (^(v34))
                var v46 uint16 = v34
                _ = v46
            } else {
                if false {
                    fmt.Println(v2, v4)
                    if ((v20 >= genFn0(("" + ("foo bar" + (("" + "hello") + ("" + "")))), genFn0(("hello" + (("hello" + "\n\t") + "hello")), (v20 | v4), 0), 42)) && true) {
                        v47 := GenS5{F0: (v20 - genFn0("", v4, 9223372036854775807)), F1: "hello"}
                        _ = v47
                    }
                    var v48 int16 = 42
                    _ = v48
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v49 int64 = 42
                    v20 += v20
                    _ = v49
                }
            }
            _ = v43
        }
        fmt.Println(v2)
        fmt.Println(v34)
    }
    switch "a" {
    case "a":
        var v50 uint8 = v2
        _ = v50
    case "hello":
        var v51 int16 = 32767
        _ = v51
    case "foo bar":
        v34 = v34
    default:
        for fi52 := 0; fi52 < 16; fi52++ {
            _ = len(v5)
            delete(v1, -2147483648)
            v3 = append(v3, -9223372036854775808)
            _ = fi52
        }
    }
    fmt.Println(v34)
    if (-612 <= -2147483648) {
        v20 += (v20 % (((v20 & (v4 - (^(v20))))) | 1))
    }
    var v53 uint8 = v2
    v53 -= 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v53
}
