package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint64; F1 int; F2 uint }
func genFn0(p0 string, p1 int16) int32 {
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(p1, fi6)
        _ = fi6
    }
    switch 2147483647 {
    case 32767:
        for fi7 := 0; fi7 < 10; fi7++ {
            p0 = p0
            fmt.Println(fi7)
            _ = fi7
        }
    case -9223372036854775808:
        p1 = p1
    case 127:
        var v8 uint8 = 12
        _ = v8
    default:
        var v9 uint64 = 863
        _ = v9
    }
    p1 -= p1
    _ = p0
    _ = p1
    return 126
}
func genFn1(p0 bool, p1 int) uint16 {
    v10 := []int{p1, p1}
    for fi11 := 0; fi11 < 19; fi11++ {
        if p0 {
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 uint32 = 147
                _ = fi12
                _ = v13
            }
            fmt.Println(fi11, p0, p0)
            var v14 int32 = 0
            _ = v14
        } else {
            switch (("" + "") + "") {
            case "a":
                v15 := GenS1{F0: 9223372036854775807}
                _ = v15
            case "\n\t":
                var v16 uint32 = 284
                _ = v16
            case "":
                p0 = false
            }
        }
        p1 -= (fi11 << ((2) & 63))
        _ = fi11
    }
    v17 := GenS2{F0: -9223372036854775808, F1: 255}
    for fi18 := 0; fi18 < 9; fi18++ {
        if true {
            var v19 uint16 = 0
            var v20 int16 = 1000
            _ = v19
            _ = v20
        } else {
            p1 -= -1
            var v21 string = "\n\t"
            switch "hello" {
            case "foo bar":
                p0 = false
            case "\n\t":
                for fi22 := 0; fi22 < 10; fi22++ {
                    switch -442 {
                    case 32767:
                        var v23 uint32 = 76
                        _ = v23
                    }
                    v24 := GenS1{F0: p1}
                    var v25 uint16 = 1
                    _ = fi22
                    _ = v24
                    _ = v25
                }
            case "":
                v21 = v21
            }
            _ = v21
        }
        if p0 {
            switch (fi18 / ((42) | 1)) {
            case 9223372036854775807:
                v10 = append(v10, (fi18 << ((1) & 63)))
            default:
                _ = len(v10)
            }
            if ((0.1 * 1.5) >= 1e10) {
                if (157 == genFn0("hello", 1000)) {
                    fmt.Println(p1)
                }
                p1 = p1
            } else {
                fmt.Println(p0, p0)
            }
            var v26 uint8 = 127
            _ = v26
        } else {
            var v27 float64 = ((3.14159 + 0.1) / 0.1)
            v27 -= 3.14159
            _ = v27
        }
        if (!(false && p0)) {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            p0 = (1e10 == (float64(3.14159) / float64(1.5)))
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v17
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 float64 = 0.1
    v4 := []uint32{0, 2, 2147483647, 2147483648}
    var v5 bool = (32767 > 42)
    v3 -= v3
    if len(v4) > 0 {
        v4[0] = 2147483648
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        var v29 int16 = 0
        for fi30 := 0; fi30 < 13; fi30++ {
            v3 = 0.0
            var v31 uint32 = 2147483647
            _ = fi30
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    v5 = v5
    delete(v1, 32767)
    if (!(42 != 0)) {
        var v32 int16 = 32767
        var v33 float32 = 1.5
        _ = v32
        _ = v33
    } else {
        if v5 {
            for fi34 := 0; fi34 < 16; fi34++ {
                if v5 {
                    if (!((^(fi34)) <= fi34)) {
                        if v5 {
                            v35 := GenS1{F0: fi34}
                            v5 = (!(!(v5 && v5)))
                            _ = v35
                        } else {
                            switch fi34 {
                            case 128:
                                if (128 >= 127) {
                                    delete(v2, (("hello" + "\n\t") + "foo bar"))
                                    v3 = v3
                                } else {
                                    v3 = 1.5
                                    for fi36 := 0; fi36 < 9; fi36++ {
                                        if (!v5) {
                                            _ = len(v4)
                                            if false {
                                                fmt.Println(fi34, v5)
                                                v37 := GenS1{F0: (fi34 + fi36)}
                                                _ = v37
                                            } else {
                                                if (!v5) {
                                                    v5 = false
                                                }
                                                var v38 string = (("foo bar" + "hello") + "a")
                                                _ = v38
                                            }
                                        } else {
                                            v39 := GenS3{F0: 2, F1: fi34, F2: 0}
                                            _ = v39
                                        }
                                        var v40 float32 = 1.5
                                        _ = fi36
                                        _ = v40
                                    }
                                }
                            case 127:
                                if (58 > 1) {
                                    v5 = (1e10 == float32(1e10))
                                    if false {
                                        v5 = false
                                        v2["a"] = -9223372036854775808
                                    }
                                }
                            default:
                                v2[("foo bar" + "a")] = (fi34 >> ((677) & 63))
                            }
                            v4 = append(v4, 4294967295)
                        }
                    } else {
                        fmt.Println(v3, v5, v5)
                        if (255 > 255) {
                            switch (^(fi34)) {
                            case 42:
                                if v5 {
                                    for fi41 := 0; fi41 < 15; fi41++ {
                                        delete(v1, (fi34 + (fi34 & -455)))
                                        var v42 uint8 = 42
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = fi41
                                        _ = v42
                                    }
                                    v3 = v3
                                }
                            case 778:
                                switch "foo bar" {
                                case "foo bar":
                                    v43 := GenS3{F0: 1, F1: (fi34 + 32768), F2: 929}
                                    _ = v43
                                case "\n\t":
                                    for fi44 := 0; fi44 < 9; fi44++ {
                                        v45 := GenS3{F0: 4294967295, F1: 0, F2: 1}
                                        _ = fi44
                                        _ = v45
                                    }
                                default:
                                    v46 := GenS2{F0: 644, F1: 255}
                                    _ = v46
                                }
                            default:
                                for fi47 := 0; fi47 < 12; fi47++ {
                                    switch (fi47 + fi47) {
                                    case 4294967295:
                                        var v48 uint8 = 255
                                        _ = v48
                                    case -986:
                                        var v49 float64 = v3
                                        _ = v49
                                    }
                                    fmt.Println(fi34, v3, v3)
                                    _ = len(v4)
                                    var v50 int64 = 42
                                    _ = fi47
                                    _ = v50
                                }
                            }
                            var v51 string = "a"
                            switch v51 {
                            case "a":
                                v51 = v51
                            default:
                                var v52 int64 = -897
                                _ = v52
                            }
                            _ = v51
                        }
                    }
                } else {
                    var v53 int = fi34
                    _ = len(v2)
                    _ = v53
                }
                v5 = false
                var v54 int64 = -1
                _ = fi34
                _ = v54
            }
            var v55 int = 42
            _ = v55
        } else {
            fmt.Println(v3, v3)
        }
        v3 -= v3
        v5 = (-2147483648 < 0)
    }
    var v56 string = "foo bar"
    v3 += (v3 - v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
}
