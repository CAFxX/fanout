package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int32, p1 int32) int64 {
    v6 := []int{-577, -2147483648, -9223372036854775808, -2147483648}
    if (42 != 127) {
        fmt.Println(p0, p1)
        if (175 >= 797) {
            v6 = append(v6, 0)
            var v7 uint64 = 4294967295
            v7 += (v7 + 2)
            _ = v7
        }
        var v8 uint8 = 255
        _ = v8
    } else {
        switch ("\n\t" + "a") {
        case "foo bar":
            for fi9 := 0; fi9 < 17; fi9++ {
                p0 += (-(p1))
                if false {
                    if false {
                        p1 -= (p0 - -1)
                    } else {
                        var v10 int64 = 0
                        var v11 bool = false
                        _ = v10
                        _ = v11
                    }
                    var v12 uint64 = 0
                    fmt.Println(p0, p1, v12)
                    _ = v12
                }
                _ = fi9
            }
        case "hello":
            p0 += 0
        case "\n\t":
            switch ("a" + "a") {
            case "hello":
                var v13 float32 = (-2.25 / 3.14159)
                _ = v13
            case "":
                if len(v6) > 0 {
                    v6[0] = 204
                }
            default:
                if (2147483648 == 42) {
                    var v14 uint16 = 32767
                    var v15 uint8 = 2
                    var v16 int64 = 42
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v17 int16 = -32768
                    _ = v17
                }
            }
        default:
            var v18 bool = (false || true)
            _ = v18
        }
    }
    fmt.Println(p1, p0)
    var v19 float32 = 0.1
    _ = p0
    _ = p1
    _ = v19
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = true
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    var v5 int64 = 9223372036854775807
    _ = len(v4)
    v20 := []int{9223372036854775807, -481, 1, 1}
    v1[1] = true
    if v2 {
        v5 -= (v5 % ((genFn0(42, 42)) | 1))
        for fi21 := 0; fi21 < 1; fi21++ {
            for fi22 := 0; fi22 < 15; fi22++ {
                v2 = v2
                var v23 float64 = -2.25
                var v24 int = 738
                _ = fi22
                _ = v23
                _ = v24
            }
            if (v2 && v2) {
                v2 = (845 < 18446744073709551615)
                for fi25 := 0; fi25 < 16; fi25++ {
                    v2 = (65535 < 65535)
                    if v2 {
                        var v26 int16 = -879
                        var v27 int8 = 100
                        fmt.Println(v27, v27, v2)
                        v28 := GenS1{F0: (v27 ^ v27)}
                        fmt.Println(fi25, v27)
                        var v29 uint32 = 0
                        _ = v26
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        var v30 uint32 = 42
                        if ((v30 % ((v30) | 1)) == 1) {
                            var v31 int64 = (v5 << ((82) & 63))
                            var v32 bool = false
                            var v33 int16 = 1
                            _ = v31
                            _ = v32
                            _ = v33
                        } else {
                            v34 := GenS1{F0: 66}
                            v35 := GenS1{F0: 42}
                            v30 = (v30 % (((v30 + 2)) | 1))
                            v5 = 0
                            _ = v34
                            _ = v35
                        }
                        v36 := GenS2{F0: 0, F1: 2}
                        _ = v30
                        _ = v36
                    }
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    _ = fi25
                }
                v5 += (-(v5))
            }
            v37 := GenS1{F0: 22}
            _ = fi21
            _ = v37
        }
        v2 = v2
    }
    v5 += v5
    var v38 uint16 = 0
    fmt.Println(v38, v5, v38)
    var v39 bool = v2
    if v39 {
        v2 = ((!false) || (v2 && (((("hello" + ("a" + "")) + "") + (((("" + ("foo bar" + ("a" + ("\n\t" + "\n\t")))) + "") + ("" + "a")) + "\n\t")) > "")))
        for fi40 := 0; fi40 < 19; fi40++ {
            var v41 uint = 449
            if len(v20) > 0 {
                v20[0] = 128
            }
            _ = fi40
            _ = v41
        }
        var v42 string = "foo bar"
        _ = v42
    }
    v5 = genFn0(-2147483648, -2147483648)
    var v43 float32 = (0.1 * 0.0)
    if v2 {
        _ = len(v3)
    }
    v20 = append(v20, 32768)
    for fi44 := 0; fi44 < 11; fi44++ {
        switch "hello" {
        case "\n\t":
            var v45 int64 = genFn0(2147483647, 42)
            _ = v45
        }
        v4["foo bar"] = "a"
        for fi46 := 0; fi46 < 1; fi46++ {
            var v47 int16 = -1
            v48 := GenS3{F0: 127}
            _ = fi46
            _ = v47
            _ = v48
        }
        _ = fi44
    }
    v43 -= (v43 / v43)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v5
}
