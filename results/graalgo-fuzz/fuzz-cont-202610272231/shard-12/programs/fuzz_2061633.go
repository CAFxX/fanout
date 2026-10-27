package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 int8; F2 int8 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 int; F2 uint64 }
func genFn0(p0 bool, p1 uint64) int16 {
    v7 := []int{-354}
    var v8 uint8 = 127
    switch ("foo bar" + ("hello" + ("\n\t" + "\n\t"))) {
    case "hello":
        for fi9 := 0; fi9 < 1; fi9++ {
            if (p0 || p0) {
                if false {
                    var v10 string = (("hello" + "a") + "a")
                    v11 := GenS1{F0: v8, F1: -193}
                    p1 = p1
                    _ = v10
                    _ = v11
                } else {
                    var v12 uint16 = uint16(p1)
                    _ = v12
                }
            } else {
                p1 += (p1 % ((p1) | 1))
                if (p0 || (!false)) {
                    for fi13 := 0; fi13 < 19; fi13++ {
                        switch (fi9 * (^(fi9))) {
                        case 128:
                            p1 -= (p1 ^ (p1 + 1))
                        case -1:
                            v8 = v8
                        default:
                            var v14 int = fi13
                            _ = v14
                        }
                        v15 := GenS1{F0: (v8 | (v8 ^ 116)), F1: 113}
                        var v16 int64 = 0
                        _ = fi13
                        _ = v15
                        _ = v16
                    }
                }
                if (p0 || (p0 && true)) {
                    v17 := GenS1{F0: 175, F1: 1}
                    fmt.Println(p0)
                    _ = v17
                }
            }
            _ = fi9
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []string{"hello", "foo bar"}
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]string)
    var v4 int64 = (v2 >> ((4294967295) & 63))
    v5 := []int{-9223372036854775808, 0, 127, -2147483648}
    var v6 uint = 1
    fmt.Println(v2, v4)
    _ = len(v1)
    if false {
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 uint16 = 0
            var v20 bool = false
            switch (((("hello" + "hello") + "a") + "") + ("hello" + ("" + ((("foo bar" + "hello") + ("a" + ("" + ""))) + ("" + ("hello" + "hello")))))) {
            case "":
                v5 = append(v5, (fi18 % ((fi18) | 1)))
            case "hello":
                v21 := GenS1{F0: 255, F1: 42}
                _ = v21
            default:
                v2 += 42
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        fmt.Println(v6)
        v22 := GenS1{F0: 175, F1: -2147483648}
        _ = v22
    } else {
        if true {
            v23 := GenS1{F0: 72, F1: 42}
            if false {
                v2 += (v4 * v4)
            } else {
                var v24 uint64 = 42
                _ = v24
            }
            for fi25 := 0; fi25 < 5; fi25++ {
                v6 = (v6 + 0)
                _ = fi25
            }
            _ = v23
        }
        if true {
            delete(v3, 128)
        } else {
            var v26 uint64 = 18446744073709551615
            if true {
                v2 += (v4 / ((v2) | 1))
            } else {
                var v27 int16 = 97
                _ = v27
            }
            fmt.Println(v2, v26, v6)
            _ = v26
        }
    }
    v1 = append(v1, "a")
    for fi28 := 0; fi28 < 2; fi28++ {
        if len(v1) > 0 {
            v1[0] = "a"
        }
        v3[42] = "foo bar"
        var v29 uint = 18446744073709551615
        _ = fi28
        _ = v29
    }
    var v30 int32 = 1
    var v31 int8 = -89
    if false {
        for fi32 := 0; fi32 < 18; fi32++ {
            if ((((1e10 / float64(float32(-2.25))) - 1.5) - 1e10) == (1.5 + (0.0 - 3.14159))) {
                var v33 bool = (0.0 > -2.25)
                _ = v33
            } else {
                v5 = append(v5, (fi32 * 9223372036854775807))
                var v34 uint32 = 304
                _ = v34
            }
            _ = fi32
        }
        var v35 int64 = (v4 % (((v4 << ((v6) & 63))) | 1))
        var v36 int64 = 1
        _ = v35
        _ = v36
    }
    var v37 uint32 = 42
    if (42 <= int16(v31)) {
        v38 := GenS2{F0: 18446744073709551615}
        if false {
            v30 -= -562
            switch 1 {
            case 1:
                if len(v1) > 0 {
                    v1[0] = ("\n\t" + ("a" + ("hello" + "a")))
                }
            }
        } else {
            var v39 float64 = (1.5 - (0.0 - float64(-2.25)))
            if true {
                fmt.Println(v37)
            }
            v30 -= v30
            _ = v39
        }
        _ = v38
    }
    v40 := GenS3{F0: 255, F1: -128, F2: v31}
    switch "\n\t" {
    case "hello":
        for fi41 := 0; fi41 < 15; fi41++ {
            v42 := GenS3{F0: 748, F1: 0, F2: (v31 % ((v31) | 1))}
            _ = fi41
            _ = v42
        }
    case "\n\t":
        v30 = int32(v4)
    case "":
        if (!true) {
            if true {
                fmt.Println(v4)
                fmt.Println(v30, v37, v30)
                v43 := GenS3{F0: uint16(v37), F1: v31, F2: v31}
                _ = v43
            } else {
                if len(v5) > 0 {
                    v5[0] = -2147483648
                }
                v44 := GenS4{F0: 10, F1: v30}
                _ = v44
            }
            switch "" {
            case "foo bar":
                v2 = -1
            case "\n\t":
                var v45 float64 = (1e10 / (-2.25 / 0.1))
                _ = v45
            case "":
                v3[32767] = "hello"
            }
            for fi46 := 0; fi46 < 1; fi46++ {
                for fi47 := 0; fi47 < 15; fi47++ {
                    if false {
                        v48 := GenS5{F0: 920}
                        var v49 uint32 = v37
                        v50 := GenS2{F0: 42}
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                    var v51 float64 = (0.1 / (0.1 - (0.1 / -2.25)))
                    if ((!(false && true)) && true) {
                        var v52 int32 = v30
                        fmt.Println(v2)
                        _ = v52
                    }
                    _ = fi47
                    _ = v51
                }
                _ = fi46
            }
        }
    }
    if true {
        var v53 uint8 = 1
        _ = v53
    }
    switch 1 {
    case -9223372036854775808:
        fmt.Println(v6, v6)
    case 2147483647:
        v54 := GenS6{F0: v31, F1: -2147483648, F2: 18446744073709551615}
        _ = v54
    case 4294967295:
        if (("hello" + "") != "hello") {
            var v55 float32 = -2.25
            for fi56 := 0; fi56 < 19; fi56++ {
                switch fi56 {
                case 0:
                    delete(v3, (-(fi56)))
                case 9223372036854775807:
                    v31 -= v31
                }
                _ = fi56
            }
            var v57 uint = 814
            _ = v55
            _ = v57
        }
    }
    var v58 uint8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v58
    _ = v6
}
