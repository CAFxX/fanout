package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int64; F2 uint }
type GenS4 struct { F0 string; F1 string; F2 uint64 }
func genFn0(p0 int64, p1 int) int16 {
    v4 := []int{}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi5 := 0; fi5 < 4; fi5++ {
        if true {
            var v6 int32 = 2147483647
            var v7 uint16 = 0
            p1 -= fi5
            _ = v6
            _ = v7
        }
        if (!false) {
            var v8 int64 = p0
            v9 := GenS1{F0: 2147483647, F1: 4294967295}
            if (true && false) {
                var v10 bool = true
                v11 := GenS1{F0: -2147483648, F1: 229}
                _ = v10
                _ = v11
            } else {
                if false {
                    fmt.Println(v8, fi5)
                    v12 := GenS2{F0: p0}
                    fmt.Println(p1, v8)
                    _ = v12
                } else {
                    var v13 int8 = 36
                    _ = v13
                }
                fmt.Println(p0, v8)
                for fi14 := 0; fi14 < 13; fi14++ {
                    fmt.Println(p0, fi14)
                    p0 -= (-(v8))
                    p1 += fi14
                    _ = fi14
                }
            }
            _ = v8
            _ = v9
        } else {
            v15 := GenS3{F0: (p0 | (p0 * -9223372036854775808)), F1: p0, F2: 0}
            switch (fi5 | (fi5 % ((int(p0)) | 1))) {
            case 966:
                switch p1 {
                case 94:
                    var v16 bool = false
                    _ = v16
                case 392:
                    p0 = (p0 - (-(p0)))
                default:
                    var v17 uint8 = 1
                    _ = v17
                }
            case 327:
                var v18 int = (fi5 * (fi5 + 302))
                _ = v18
            default:
                p1 -= (fi5 >> ((2) & 63))
            }
            if len(v4) > 0 {
                v4[0] = 9223372036854775807
            }
            _ = v15
        }
        _ = fi5
    }
    p0 -= 0
    var v19 uint64 = 2
    _ = p0
    _ = p1
    _ = v19
    _ = v4
    return 32767
}
func genFn1(p0 uint32) int {
    if ("\n\t" != (((("foo bar" + ("hello" + "a")) + "foo bar") + ("a" + "")) + ("\n\t" + ""))) {
        v20 := []int{70, -1, 693}
        v20 = append(v20, 471)
        _ = v20
    }
    var v21 int64 = 1
    var v22 uint8 = 255
    var v23 uint32 = (p0 - p0)
    _ = p0
    _ = v21
    _ = v22
    _ = v23
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 float32 = 0.0
    v3 -= v3
    switch -2147483648 {
    case 9223372036854775807:
        var v24 uint32 = 2
        _ = v24
    case 127:
        var v25 float64 = 0.1
        _ = v25
    }
    switch "foo bar" {
    case "foo bar":
        v26 := []int{129, genFn1(2147483648), 0}
        _ = v26
    default:
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 uint = 42
            v28 += (v28 | (v28 | (v28 ^ (v28 % ((v28) | 1)))))
            _ = fi27
            _ = v28
        }
    }
    v3 += v3
    var v29 int8 = -1
    for fi30 := 0; fi30 < 5; fi30++ {
        v3 += v3
        _ = fi30
    }
    v31 := GenS2{F0: 42}
    var v32 int32 = 466
    v2[42] = ("\n\t" + ("hello" + "hello"))
    for fi33 := 0; fi33 < 2; fi33++ {
        fmt.Println(fi33, v32, v32)
        for fi34 := 0; fi34 < 13; fi34++ {
            if (true || (!true)) {
                var v35 int8 = (-(v29))
                _ = v35
            }
            var v36 int64 = 42
            if (!(!((0.0 * v3) == -2.25))) {
                if (!(32767 <= 2)) {
                    var v37 int32 = (-(v32))
                    if true {
                        v3 = 1e10
                        switch fi34 {
                        case 1:
                            switch (^(fi34)) {
                            case 9223372036854775807:
                                v37 += v37
                            default:
                                v38 := GenS4{F0: (("" + ("a" + "a")) + "foo bar"), F1: "", F2: 519}
                                _ = v38
                            }
                        default:
                            v39 := []int{(-(fi34)), 715, (^(fi33))}
                            _ = v39
                        }
                        delete(v2, -539)
                    } else {
                        var v40 uint16 = 860
                        if false {
                            v37 = 835
                            v41 := []int{}
                            _ = v41
                        }
                        _ = v40
                    }
                    v37 = v37
                    v1["\n\t"] = ("" + (("foo bar" + "\n\t") + ("hello" + "a")))
                    var v42 int8 = (v29 * 1)
                    _ = v37
                    _ = v42
                }
            }
            _ = fi34
            _ = v36
        }
        _ = fi33
    }
    v29 = 127
    v43 := []int{-629, genFn1(2)}
    for fi44 := 0; fi44 < 9; fi44++ {
        fmt.Println(v32, fi44, v3)
        v3 = 1.5
        _ = fi44
    }
    fmt.Println(v32)
    delete(v2, 32767)
    for fi45 := 0; fi45 < 14; fi45++ {
        if true {
            _ = len(v43)
        } else {
            for fi46 := 0; fi46 < 20; fi46++ {
                switch genFn1(57) {
                case 213:
                    v29 += (v29 / ((v29) | 1))
                }
                _ = fi46
            }
        }
        _ = fi45
    }
    v32 = v32
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v43
}
