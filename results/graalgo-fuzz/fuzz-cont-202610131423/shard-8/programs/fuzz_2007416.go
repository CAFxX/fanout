package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
func genFn0() uint {
    v7 := GenS1{F0: 0, F1: 32767}
    var v8 uint = 42
    fmt.Println(v8)
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 int64 = -1
        _ = fi9
        _ = v10
    }
    var v11 int16 = -32768
    _ = v11
    _ = v7
    _ = v8
    return v8
}
func genFn1(p0 bool) uint8 {
    var v12 uint16 = 3
    if (true || false) {
        p0 = (!("a" <= "foo bar"))
        v12 -= v12
    } else {
        v13 := []int{}
        fmt.Println(p0, p0)
        if (!((((("a" + ("" + (("\n\t" + (("" + ("\n\t" + "foo bar")) + "a")) + ("a" + "")))) + "\n\t") + "foo bar") + "\n\t") == "foo bar")) {
            fmt.Println(p0, p0)
            p0 = false
        }
        _ = v13
    }
    _ = p0
    _ = v12
    return uint8(v12)
}

func main() {
    var v1 uint16 = 51
    v2 := make(map[int]uint32)
    var v3 uint8 = 42
    var v4 uint32 = 2147483647
    var v5 bool = ((-860 >= -1) && false)
    var v6 float32 = (0.0 * (0.0 + -2.25))
    fmt.Println(v6)
    v14 := GenS2{F0: "hello"}
    if v5 {
        switch "a" {
        case "foo bar":
            v5 = ((true || (!v5)) || v5)
        case "hello":
            v1 += (^(v1))
        default:
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 int64 = 9223372036854775807
                if (v3 == uint8(v4)) {
                    var v17 uint = genFn0()
                    _ = v17
                } else {
                    fmt.Println(v16, v6, v1)
                }
                _ = fi15
                _ = v16
            }
        }
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 uint32 = (v4 + 0)
            var v20 int16 = 0
            _ = fi18
            _ = v19
            _ = v20
        }
        v4 -= (v4 & v4)
    } else {
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 uint16 = (v1 % (((v1 - 902)) | 1))
            v5 = (true && ((v3 ^ (v3 & (^(v3)))) > v3))
            v23 := GenS2{F0: "\n\t"}
            _ = fi21
            _ = v22
            _ = v23
        }
        v6 = v6
    }
    var v24 uint64 = 0
    v24 = v24
    v4 += (^(v4))
    var v25 uint16 = (v1 % ((v1) | 1))
    var v26 int32 = 859
    fmt.Println(v4, v6, v1)
    if v5 {
        var v27 int16 = 0
        _ = len(v2)
        var v28 float32 = 0.1
        _ = v27
        _ = v28
    } else {
        if (uint8(v24) != genFn1((v5 || true))) {
            v6 = v6
        }
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        v24 += uint64(v4)
        v5 = true
        switch (fi29 >> ((uint(v24)) & 63)) {
        case 1:
            v6 += v6
        default:
            v6 += (1.5 - 1.5)
        }
        _ = fi29
    }
    var v30 uint16 = v25
    fmt.Println(v3, v5, v6)
    v31 := GenS1{F0: 748, F1: -2147483648}
    v32 := []int{0, -91}
    delete(v2, -461)
    v33 := GenS1{F0: -32768, F1: 2147483647}
    switch 695 {
    case 0:
        for fi34 := 0; fi34 < 7; fi34++ {
            if v5 {
                for fi35 := 0; fi35 < 17; fi35++ {
                    if ((v25 ^ v1) != (v1 ^ v30)) {
                        v3 = 152
                        v36 := GenS2{F0: "foo bar"}
                        fmt.Println(v30, fi34, v30)
                        _ = v36
                    } else {
                        v25 = v1
                        if len(v32) > 0 {
                            v32[0] = fi34
                        }
                    }
                    _ = fi35
                }
            } else {
                v5 = false
            }
            v37 := GenS3{F0: 2}
            _ = fi34
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
