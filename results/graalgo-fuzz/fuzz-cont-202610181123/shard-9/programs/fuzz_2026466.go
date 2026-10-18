package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int8; F1 int32; F2 int }
func genFn0() uint {
    v6 := []int{-277, 42, 127, 32768}
    var v7 float32 = (0.0 + -2.25)
    for fi8 := 0; fi8 < 1; fi8++ {
        v7 = (float32((1.5 * ((-2.25 / 1e10) + 1.5))) - (-2.25 * 0.0))
        _ = fi8
    }
    if (false || (-501 != 0)) {
        v7 -= v7
        var v9 int = 536
        _ = v9
    } else {
        var v10 uint = 701
        v7 = v7
        _ = v10
    }
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 uint8, p1 int8, p2 string) string {
    fmt.Println(p0, p2)
    if true {
        v11 := GenS1{F0: 1}
        p1 = p1
        _ = v11
    }
    p2 = (p2 + "")
    _ = p0
    _ = p1
    _ = p2
    return (p2 + p2)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{(("foo bar" + "\n\t") + ""), "hello"}
    v3 := make(map[uint32]bool)
    var v4 float32 = 0.0
    var v5 uint = 65535
    v4 = -2.25
    for fi12 := 0; fi12 < 16; fi12++ {
        v5 = v5
        _ = fi12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            fmt.Println(fi14, fi13)
            if (!(fi13 >= fi13)) {
                switch ("foo bar" + (("hello" + ("foo bar" + "a")) + (("a" + "\n\t") + "hello"))) {
                case "":
                    fmt.Println(fi13, v4, fi13)
                }
                v2 = append(v2, "")
            } else {
                var v15 bool = true
                if ((1e10 / 3.14159) < 0.1) {
                    if ((v4 > v4) || v15) {
                        if v15 {
                            var v16 int64 = -9223372036854775808
                            if (-1 != 0) {
                                var v17 bool = (!((("\n\t" + "foo bar") + ("a" + "")) < ("foo bar" + "")))
                                _ = v17
                            } else {
                                fmt.Println(fi13, v4)
                            }
                            _ = v16
                        } else {
                            var v18 uint64 = 0
                            fmt.Println(v4)
                            _ = v18
                        }
                    }
                }
                _ = v15
            }
            v19 := GenS1{F0: 91}
            _ = fi14
            _ = v19
        }
        switch "hello" {
        case "":
            v20 := GenS1{F0: 42}
            _ = v20
        default:
            fmt.Println(v4, v4, v4)
        }
        _ = fi13
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        if false {
            switch "foo bar" {
            case "":
                switch "foo bar" {
                case "":
                    var v22 string = ("hello" + (("\n\t" + "foo bar") + ""))
                    _ = v22
                default:
                    if (false || true) {
                        var v23 float32 = v4
                        _ = v23
                    } else {
                        _ = len(v2)
                        v24 := GenS2{F0: -1, F1: 1}
                        _ = v24
                    }
                }
            case "foo bar":
                for fi25 := 0; fi25 < 9; fi25++ {
                    v26 := GenS1{F0: 100}
                    var v27 uint64 = 315
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            case "a":
                v4 += float32(-2.25)
            default:
                for fi28 := 0; fi28 < 9; fi28++ {
                    v5 = v5
                    _ = fi28
                }
            }
            var v29 float64 = (1e10 * (0.0 + (0.1 / float64(float32((((1e10 * float64(v4)) + float64(0.1)) - -2.25))))))
            if (true && true) {
                var v30 float32 = v4
                v31 := GenS1{F0: -1}
                _ = v30
                _ = v31
            }
            _ = v29
        } else {
            if len(v2) > 0 {
                v2[0] = ""
            }
            var v32 string = (("" + "") + "foo bar")
            _ = v32
        }
        _ = fi21
    }
    if (false && (!(true || false))) {
        var v33 uint64 = 494
        v5 = v5
        if (int(v33) < -9223372036854775808) {
            fmt.Println(v5, v5)
            if false {
                v4 = float32(0.1)
                for fi34 := 0; fi34 < 12; fi34++ {
                    fmt.Println(v4)
                    if (0 == 0) {
                        var v35 uint64 = v33
                        _ = v35
                    }
                    _ = fi34
                }
                v4 = ((float32(1.5) - 3.14159) - (v4 - 3.14159))
            } else {
                v5 = (v5 / ((genFn0()) | 1))
                v5 -= (v5 - v5)
            }
            v36 := GenS3{F0: 100, F1: 299, F2: 4294967295}
            _ = v36
        }
        _ = v33
    } else {
        v37 := GenS1{F0: 1}
        fmt.Println(v4, v5)
        _ = v37
    }
    v1[((("" + "a") + "\n\t") + "a")] = 2147483647
    fmt.Println(v4, v5)
    var v38 float32 = v4
    var v39 uint32 = 42
    var v40 int = 32768
    for fi41 := 0; fi41 < 2; fi41++ {
        v39 = 2147483647
        v40 = (v40 / (((v40 ^ v40)) | 1))
        _ = fi41
    }
    v40 = v40
    var v42 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v5
}
