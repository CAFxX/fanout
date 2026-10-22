package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int; F1 uint64; F2 int32 }
func genFn0() int8 {
    v5 := GenS1{F0: 32767}
    if false {
        switch 555 {
        case 32767:
            v6 := GenS2{F0: ((("" + "foo bar") + "") + ""), F1: 65535}
            _ = v6
        case 0:
            var v7 float32 = 1.5
            _ = v7
        case 128:
            var v8 int32 = 0
            _ = v8
        }
        v9 := []int{9223372036854775807}
        _ = v9
    } else {
        v10 := GenS2{F0: (("foo bar" + "foo bar") + ("hello" + "\n\t")), F1: 32767}
        _ = v10
    }
    switch "foo bar" {
    case "a":
        var v11 float64 = float64(0.0)
        _ = v11
    case "\n\t":
        if false {
            var v12 int8 = 109
            _ = v12
        }
    default:
        var v13 int64 = -2
        _ = v13
    }
    if false {
        var v14 uint64 = 0
        _ = v14
    }
    _ = v5
    return 15
}
func genFn1(p0 uint32, p1 uint64) uint {
    for fi15 := 0; fi15 < 1; fi15++ {
        fmt.Println(fi15, p1, fi15)
        _ = fi15
    }
    p1 -= p1
    v16 := GenS2{F0: "a", F1: 301}
    if (!(!(826 != p1))) {
        var v17 int = -796
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v16
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 74
    var v3 int8 = -1
    v4 := []uint32{998}
    delete(v1, ("a" + ("" + "foo bar")))
    var v18 uint16 = 756
    var v19 int8 = (v3 & v3)
    if true {
        var v20 float64 = ((3.14159 * 1e10) / 3.14159)
        v20 -= (float64((1e10 - 0.1)) - (v20 * v20))
        if len(v4) > 0 {
            v4[0] = 379
        }
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 18; fi22++ {
            if (0 == -168) {
                switch fi21 {
                case 124:
                    if true {
                        var v23 bool = false
                        if (!(32767 <= -11)) {
                            v2 += 2
                        } else {
                            var v24 uint64 = v2
                            delete(v1, "foo bar")
                            if true {
                                var v25 uint8 = 128
                                var v26 uint = genFn1(317, v24)
                                _ = v25
                                _ = v26
                            } else {
                                var v27 int8 = v19
                                var v28 int64 = -689
                                _ = v27
                                _ = v28
                            }
                            _ = v24
                        }
                        fmt.Println(v18, v18, v18)
                        _ = v23
                    }
                case 32768:
                    _ = len(v1)
                    v29 := GenS1{F0: 2}
                    if true {
                        var v30 bool = ((v18 - (v18 * (^(v18)))) >= v18)
                        v18 += (^(v18))
                        if (v30 && v30) {
                            v2 += v2
                            _ = len(v1)
                        }
                        _ = v30
                    }
                    v3 = v3
                    if len(v4) > 0 {
                        v4[0] = 2
                    }
                    v3 += 127
                    _ = v29
                }
            } else {
                v18 = v18
                _ = len(v4)
                fmt.Println(fi22, v19)
            }
            if true {
                v19 -= (v3 - 1)
                var v31 float32 = -2.25
                _ = v31
            } else {
                if (!(151 > 42)) {
                    switch (fi22 ^ (fi21 - (fi21 | fi22))) {
                    case 127:
                        v32 := GenS3{F0: (fi21 % (((fi22 + 42)) | 1)), F1: v2, F2: 1}
                        _ = v32
                    default:
                        v33 := v1["hello"]
                        _ = v33
                    }
                }
            }
            _ = fi22
        }
        var v34 float32 = float32(3.14159)
        v18 = (v18 | 32767)
        _ = fi21
        _ = v34
    }
    var v35 float32 = ((1e10 - (1e10 - 0.1)) / (-2.25 + (1e10 + (1e10 * 3.14159))))
    v35 += (0.1 + (v35 - 0.1))
    for fi36 := 0; fi36 < 10; fi36++ {
        var v37 int = fi36
        v1[("foo bar" + "a")] = true
        v38 := GenS2{F0: ("" + "\n\t"), F1: v18}
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println(v35)
    v19 = v3
    switch -1 {
    case 128:
        v1["foo bar"] = true
    case 2147483647:
        var v39 uint64 = v2
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
