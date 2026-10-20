package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint }
func genFn0(p0 int64, p1 uint8, p2 int64) uint32 {
    v6 := []int{42, 4294967295, 32767}
    p0 = (p2 + (p0 - (^(p0))))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    v1 := []string{("hello" + "a"), "foo bar", "a", "foo bar"}
    var v2 float64 = float64(1.5)
    v3 := make(map[string]bool)
    var v4 int8 = 115
    v5 := []uint32{749}
    fmt.Println(v2, v2)
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            fmt.Println(fi7)
            v3[("foo bar" + "hello")] = (!false)
            _ = fi8
        }
        v4 = (v4 - v4)
        _ = fi7
    }
    v9 := GenS1{F0: -970, F1: 72}
    var v10 uint = 18446744073709551615
    _ = len(v3)
    v11 := GenS1{F0: -274, F1: 42}
    switch -1 {
    case 9223372036854775807:
        fmt.Println(v4)
    case 128:
        fmt.Println(v4, v2)
    }
    fmt.Println(v4, v10)
    var v12 int = 127
    var v13 string = ((("a" + ("a" + "foo bar")) + "") + "a")
    delete(v3, "a")
    if (642 != v10) {
        delete(v3, v13)
        v14 := GenS2{F0: v4}
        _ = v14
    } else {
        var v15 uint64 = 18446744073709551615
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 int8 = (-(v4))
            var v18 int8 = v4
            var v19 uint = v10
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        if true {
            var v20 float32 = 1e10
            if len(v5) > 0 {
                v5[0] = 0
            }
            _ = v20
        }
        _ = v15
    }
    switch v13 {
    case "foo bar":
        var v21 float64 = v2
        _ = v21
    case "hello":
        v22 := GenS1{F0: 42, F1: 128}
        _ = v22
    case "":
        v13 = (v13 + v13)
    }
    switch ((("" + "foo bar") + "foo bar") + v13) {
    case "hello":
        fmt.Println(v10, v4)
    case "foo bar":
        if false {
            for fi23 := 0; fi23 < 16; fi23++ {
                v13 = "a"
                if true {
                    v4 += -76
                    var v24 float64 = (v2 - 0.1)
                    var v25 float64 = (0.1 / (((v2 / 1.5) / v2) - -2.25))
                    _ = v24
                    _ = v25
                }
                fmt.Println(v10, v12)
                _ = fi23
            }
            var v26 uint16 = 0
            v27 := GenS3{F0: (v13 + ""), F1: v10}
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 10; fi28++ {
                switch ((v13 + "foo bar") + "\n\t") {
                case "":
                    var v29 uint = v10
                    _ = v29
                case "a":
                    fmt.Println(v13, fi28)
                case "hello":
                    delete(v3, v13)
                default:
                    for fi30 := 0; fi30 < 20; fi30++ {
                        v2 += 1e10
                        var v31 float32 = (1.5 * 1e10)
                        _ = fi30
                        _ = v31
                    }
                }
                _ = fi28
            }
        }
    case "a":
        if ((6 == 0) || true) {
            v4 -= v4
        }
    default:
        if false {
            fmt.Println(v4)
            var v32 float32 = 0.0
            for fi33 := 0; fi33 < 14; fi33++ {
                for fi34 := 0; fi34 < 6; fi34++ {
                    if (42 > 1) {
                        var v35 uint32 = 1
                        var v36 uint16 = 255
                        v12 += (fi33 * (fi34 * fi33))
                        _ = v35
                        _ = v36
                    }
                    v37 := GenS3{F0: v13, F1: (v10 - 1)}
                    var v38 int16 = 833
                    v10 = (v10 - 255)
                    _ = fi34
                    _ = v37
                    _ = v38
                }
                v32 += v32
                v39 := GenS3{F0: ("a" + "\n\t"), F1: 4294967295}
                _ = fi33
                _ = v39
            }
            _ = v32
        }
    }
    v12 = -248
    var v40 int16 = -252
    v10 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v9
}
