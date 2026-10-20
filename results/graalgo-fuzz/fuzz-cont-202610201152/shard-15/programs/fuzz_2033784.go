package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int64; F1 bool }
func genFn0() int16 {
    var v4 uint16 = 768
    v5 := GenS1{F0: 438}
    _ = v4
    _ = v5
    return 782
}
func genFn1(p0 uint32, p1 int8) uint16 {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint16 = 32767
        if (2 >= 1) {
            switch "a" {
            case "\n\t":
                var v8 uint64 = 102
                _ = v8
            case "":
                var v9 float32 = 1.5
                _ = v9
            case "foo bar":
                if (255 <= 1) {
                    v10 := []int{fi6, fi6}
                    v11 := GenS1{F0: 4294967295}
                    _ = v10
                    _ = v11
                }
            default:
                v12 := GenS1{F0: 4294967295}
                _ = v12
            }
            v13 := []int{fi6}
            _ = v13
        }
        _ = fi6
        _ = v7
    }
    p1 -= p1
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 uint32 = 65535
        if true {
            switch "foo bar" {
            case "a":
                if (!(!false)) {
                    if (900 == 2) {
                        v16 := []int{(^(fi14)), (fi14 | (fi14 >> ((213) & 63)))}
                        fmt.Println(p0, v15)
                        _ = len(v16)
                        _ = v16
                    }
                    if (!(687 <= 18446744073709551615)) {
                        var v17 float32 = (-2.25 * 1e10)
                        var v18 int32 = 574
                        var v19 uint32 = p0
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        v20 := GenS2{F0: (552 > -9223372036854775808), F1: 131}
                        v15 += (^(p0))
                        v21 := GenS1{F0: 65535}
                        _ = v20
                        _ = v21
                    }
                } else {
                    v22 := []int{}
                    p1 += (p1 & p1)
                    _ = v22
                }
            }
        }
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    return 636
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{127, 325, 683}
    v3 := make(map[string]string)
    if (1 < genFn0()) {
        switch 1000 {
        case 32768:
            var v23 int64 = -9223372036854775808
            _ = v23
        case 207:
            var v24 string = ("foo bar" + "")
            _ = v24
        default:
            var v25 float64 = 1.5
            _ = v25
        }
    } else {
        var v26 int32 = -790
        if false {
            fmt.Println(v26)
            if true {
                v27 := GenS2{F0: (("hello" + ((("\n\t" + "\n\t") + "") + ("hello" + ("\n\t" + "hello")))) == ("\n\t" + "\n\t")), F1: 0}
                _ = v27
            }
        }
        _ = v26
    }
    v3[""] = ("a" + "\n\t")
    v28 := v1["foo bar"]
    v28 = (2 > 2)
    var v29 uint32 = 44
    var v30 int32 = -2147483648
    if v28 {
        switch 0 {
        case 32768:
            switch 42 {
            case 127:
                var v31 uint64 = 0
                _ = v31
            }
        case -9223372036854775808:
            v32 := GenS1{F0: 65535}
            _ = v32
        default:
            var v33 uint32 = v29
            _ = v33
        }
        v29 = 667
    }
    v34 := GenS2{F0: true, F1: 65535}
    if ((!true) || true) {
        var v35 float32 = 0.0
        var v36 int32 = 42
        _ = v35
        _ = v36
    } else {
        if true {
            if (true || (!(128 == -355))) {
                v3[("" + (("" + ("\n\t" + "hello")) + ("hello" + "\n\t")))] = "hello"
                switch -947 {
                case 1:
                    v2 = append(v2, -221)
                case 42:
                    var v37 uint64 = 1
                    _ = v37
                case -226:
                    fmt.Println(v30)
                }
            } else {
                v30 += (v30 & 0)
                v38 := GenS3{F0: -205, F1: v28}
                _ = v38
            }
        }
        if true {
            var v39 float32 = (1e10 / 1e10)
            switch 9223372036854775807 {
            case 128:
                var v40 bool = v28
                _ = v40
            case 127:
                var v41 int = -368
                _ = v41
            case -2147483648:
                var v42 uint32 = 42
                _ = v42
            }
            _ = v39
        }
    }
    var v43 uint32 = v29
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v43
}
