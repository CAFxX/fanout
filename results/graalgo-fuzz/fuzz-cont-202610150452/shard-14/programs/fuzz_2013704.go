package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 int }
func genFn0() bool {
    var v4 uint16 = 42
    var v5 int32 = 42
    for fi6 := 0; fi6 < 5; fi6++ {
        if (((!false) || (true && (false || (1 == 9223372036854775807)))) && ((("" + ("a" + "a")) == "") && (219 == 9223372036854775807))) {
            v7 := GenS1{F0: true, F1: false}
            var v8 int = (fi6 >> ((18446744073709551615) & 63))
            _ = v7
            _ = v8
        } else {
            switch "hello" {
            case "foo bar":
                if ("foo bar" > "hello") {
                    for fi9 := 0; fi9 < 13; fi9++ {
                        switch "a" {
                        case "foo bar":
                            v4 = v4
                        case "\n\t":
                            v10 := []int{-1, 56, fi9}
                            _ = v10
                        case "":
                            v4 += (v4 + v4)
                        default:
                            v4 -= v4
                        }
                        _ = fi9
                    }
                }
            case "hello":
                var v11 bool = false
                _ = v11
            }
            v5 -= v5
        }
        var v12 float64 = 3.14159
        var v13 bool = (v12 > v12)
        _ = fi6
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS1{F0: true, F1: (!(42 <= 0))}
        fmt.Println(v5, v5, v4)
        var v16 int = (fi14 % ((fi14) | 1))
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: 0}
    _ = v17
    _ = v4
    _ = v5
    return (false && (!true))
}
func genFn1(p0 uint16) int32 {
    if false {
        v18 := []int{-1, 0, 42, 127}
        var v19 uint8 = 1
        for fi20 := 0; fi20 < 3; fi20++ {
            if ("hello" > ("a" + (((("foo bar" + "a") + "\n\t") + "\n\t") + ("" + (("a" + "") + ("a" + "foo bar")))))) {
                if false {
                    var v21 uint16 = p0
                    v19 = 215
                    _ = v21
                }
                for fi22 := 0; fi22 < 20; fi22++ {
                    var v23 int16 = 32767
                    v24 := GenS2{F0: -9223372036854775808}
                    if false {
                        var v25 uint8 = (^(v19))
                        var v26 uint64 = 0
                        fmt.Println(v25)
                        _ = len(v18)
                        v27 := GenS2{F0: 1}
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        var v28 int16 = v23
                        _ = v28
                    }
                    v23 = 275
                    var v29 string = ""
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v29
                }
            } else {
                for fi30 := 0; fi30 < 13; fi30++ {
                    var v31 string = "hello"
                    _ = fi30
                    _ = v31
                }
                fmt.Println(v19, fi20, v19)
                v32 := GenS3{F0: 1}
                _ = v32
            }
            _ = fi20
        }
        _ = v18
        _ = v19
    } else {
        fmt.Println(p0)
        for fi33 := 0; fi33 < 4; fi33++ {
            fmt.Println(p0, p0)
            var v34 uint8 = 76
            v34 = 42
            _ = fi33
            _ = v34
        }
        for fi35 := 0; fi35 < 1; fi35++ {
            switch fi35 {
            case 0:
                p0 -= 32767
            case 9223372036854775807:
                for fi36 := 0; fi36 < 7; fi36++ {
                    v37 := []int{fi35}
                    _ = fi36
                    _ = v37
                }
            default:
                v38 := []int{1, (fi35 | fi35), 0, (fi35 + (fi35 & 32768))}
                _ = v38
            }
            p0 = (^(p0))
            _ = fi35
        }
    }
    var v39 float64 = 0.1
    switch 4294967295 {
    case 9223372036854775807:
        v40 := GenS4{F0: 254}
        _ = v40
    case 135:
        v41 := GenS2{F0: 1}
        _ = v41
    case 1:
        fmt.Println(v39)
    }
    _ = p0
    _ = v39
    return 1
}

func main() {
    var v1 uint = 994
    v2 := []string{("hello" + "foo bar"), "hello"}
    v3 := make(map[int]uint32)
    for fi42 := 0; fi42 < 1; fi42++ {
        fmt.Println(v1, v1)
        v43 := GenS5{F0: 32767, F1: 0}
        for fi44 := 0; fi44 < 20; fi44++ {
            var v45 uint16 = 42
            _ = len(v2)
            v2 = append(v2, ("" + "a"))
            if (!((58 < 697) && true)) {
                var v46 int16 = 0
                _ = v46
            } else {
                var v47 float32 = (1e10 * 1e10)
                _ = v47
            }
            _ = fi44
            _ = v45
        }
        _ = fi42
        _ = v43
    }
    if (false && true) {
        fmt.Println(v1)
        for fi48 := 0; fi48 < 2; fi48++ {
            v1 += 18446744073709551615
            _ = fi48
        }
    } else {
        v49 := GenS1{F0: false, F1: false}
        if (0 <= int8(v1)) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v1)
            delete(v3, -9223372036854775808)
        } else {
            v1 += v1
            for fi50 := 0; fi50 < 4; fi50++ {
                v1 -= (v1 & 0)
                _ = fi50
            }
        }
        v1 = v1
        _ = v49
    }
    v51 := GenS1{F0: (-9223372036854775808 < 9223372036854775807), F1: false}
    for fi52 := 0; fi52 < 10; fi52++ {
        for fi53 := 0; fi53 < 15; fi53++ {
            _ = len(v2)
            _ = fi53
        }
        _ = fi52
    }
    if true {
        v54 := GenS4{F0: 1}
        _ = v54
    }
    delete(v3, 9223372036854775807)
    v1 += v1
    if false {
        v1 = (v1 + (v1 >> ((v1) & 63)))
        v1 -= (^(v1))
    }
    if (!true) {
        var v55 int16 = -32768
        if (true || (!false)) {
            v56 := v3[128]
            _ = len(v2)
            _ = v56
        }
        if len(v2) > 0 {
            v2[0] = "a"
        }
        _ = v55
    }
    v1 = 238
    if (((float64(float32(1.5)) >= 0.1) || false) || ((false || (false && (true || (42 == 4294967295)))) || false)) {
        fmt.Println(v1)
        var v57 float32 = 1e10
        var v58 uint32 = 2
        _ = v57
        _ = v58
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v51
}
