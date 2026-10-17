package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int32 }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0(p0 uint32, p1 uint8, p2 uint64) uint64 {
    var v6 uint8 = p1
    v7 := GenS1{F0: 0, F1: ("a" + "foo bar"), F2: -27}
    v8 := []int{}
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int64 = -409
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (^(p2))
}
func genFn1(p0 bool) string {
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 int32 = 42
        if (float32(float64(((0.1 * 0.1) + ((((3.14159 * 0.0) - 0.1) + 3.14159) + (3.14159 / 1.5))))) > float32(1e10)) {
            for fi13 := 0; fi13 < 11; fi13++ {
                if false {
                    v14 := GenS1{F0: (v12 << ((419) & 63)), F1: "hello", F2: 744}
                    v15 := []int{1, fi13, fi13, 4294967295}
                    _ = v14
                    _ = v15
                }
                if p0 {
                    fmt.Println(v12, v12)
                }
                _ = fi13
            }
            fmt.Println(v12, v12)
        } else {
            var v16 uint64 = genFn0(359, 0, 354)
            _ = v16
        }
        _ = fi11
        _ = v12
    }
    v17 := []int{42, 42, -475}
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 float32 = 1.5
        fmt.Println(p0, fi18, v19)
        _ = fi18
        _ = v19
    }
    if (!p0) {
        var v20 uint64 = 2
        _ = v20
    } else {
        p0 = false
    }
    _ = p0
    _ = v17
    return "foo bar"
}

func main() {
    var v1 int32 = 0
    v2 := make(map[int]uint32)
    v3 := []string{"hello", "hello"}
    var v4 bool = (-490 != -682)
    v5 := []string{("foo bar" + "\n\t"), "", "hello"}
    var v21 uint16 = 594
    v4 = (-9223372036854775808 >= -9223372036854775808)
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 bool = (!(1 >= -128))
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 uint8 = 61
        switch "" {
        case "hello":
            switch "foo bar" {
            case "hello":
                var v26 int32 = -1
                _ = v26
            case "\n\t":
                if v4 {
                    v25 -= 255
                    var v27 int32 = 0
                    var v28 uint16 = 255
                    _ = v27
                    _ = v28
                } else {
                    v1 -= -1
                    for fi29 := 0; fi29 < 11; fi29++ {
                        v30 := GenS2{F0: (fi29 ^ fi29), F1: v21}
                        fmt.Println(v1)
                        if true {
                            var v31 float32 = (-2.25 / 1.0)
                            _ = v31
                        }
                        var v32 int32 = (v1 % ((v1) | 1))
                        if (!true) {
                            v5 = append(v5, (("a" + ("hello" + "a")) + ""))
                            if (42 != 2) {
                                fmt.Println(v32, v25, v25)
                            } else {
                                v33 := GenS2{F0: (fi24 + fi29), F1: 255}
                                _ = v33
                            }
                            v34 := GenS1{F0: 1, F1: "", F2: v1}
                            var v35 uint16 = (v21 % (((v21 ^ 926)) | 1))
                            var v36 uint64 = 991
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        delete(v2, 173)
                        _ = fi29
                        _ = v30
                        _ = v32
                    }
                }
            }
        default:
            var v37 float32 = 0.1
            _ = v37
        }
        _ = fi24
        _ = v25
    }
    var v38 uint64 = genFn0(2147483647, 1, 44)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2[128] = 665
    if false {
        v39 := GenS1{F0: v1, F1: (("foo bar" + "foo bar") + "foo bar"), F2: (v1 % ((v1) | 1))}
        _ = v39
    }
    if (((v4 || v4) && false) && true) {
        switch 32768 {
        case 1:
            switch -9223372036854775808 {
            case 9223372036854775807:
                v4 = (v4 && v4)
            default:
                v4 = false
            }
        }
    } else {
        fmt.Println(v21, v1)
        v40 := v2[128]
        fmt.Println(v38, v4, v40)
        _ = v40
    }
    v41 := GenS1{F0: v1, F1: "\n\t", F2: v1}
    switch "a" {
    case "a":
        switch -729 {
        case 42:
            fmt.Println(v21, v38, v38)
        case -453:
            v42 := v2[42]
            _ = v42
        }
    case "":
        v21 = v21
    }
    v21 = 65535
    v43 := v2[0]
    fmt.Println(v21, v4, v43)
    if true {
        for fi44 := 0; fi44 < 16; fi44++ {
            if (0 >= (v1 & 707)) {
                if v4 {
                    var v45 uint32 = 4294967295
                    var v46 int = fi44
                    _ = v45
                    _ = v46
                }
                v47 := GenS2{F0: (^(fi44)), F1: (v21 + 0)}
                var v48 float32 = float32((1e10 + 3.14159))
                _ = v47
                _ = v48
            }
            var v49 uint16 = v21
            var v50 bool = true
            _ = fi44
            _ = v49
            _ = v50
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v38
    _ = v4
    _ = v41
    _ = v43
    _ = v5
}
