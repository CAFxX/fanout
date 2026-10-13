package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int16; F1 int; F2 int16 }
func genFn0(p0 bool) int8 {
    p0 = true
    p0 = p0
    _ = p0
    return 42
}
func genFn1(p0 int8) string {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 bool = ((true && (true || false)) || (!(1 <= 1)))
            switch "\n\t" {
            case "\n\t":
                var v10 string = "\n\t"
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        p0 -= p0
        p0 -= (-(p0))
        _ = fi7
    }
    p0 += genFn0(false)
    for fi11 := 0; fi11 < 7; fi11++ {
        switch (fi11 + (fi11 * fi11)) {
        case 32767:
            var v12 uint32 = 65535
            _ = v12
        case -1:
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 int8 = p0
                _ = fi13
                _ = v14
            }
        }
        switch (fi11 % (((fi11 | (-(fi11)))) | 1)) {
        case -174:
            v15 := []int{}
            _ = v15
        case 9223372036854775807:
            for fi16 := 0; fi16 < 18; fi16++ {
                if (0 == 1) {
                    if (true || (44 <= 1)) {
                        if (!((("\n\t" + (("\n\t" + "a") + "hello")) != ("" + ("hello" + "hello"))) && (!false))) {
                            p0 = (p0 / ((0) | 1))
                            p0 = (p0 ^ (p0 & (p0 << ((977) & 63))))
                            var v17 float32 = -2.25
                            fmt.Println(fi16)
                            _ = v17
                        } else {
                            p0 -= -108
                            v18 := GenS1{F0: 42, F1: 65535, F2: -32768}
                            _ = v18
                        }
                    } else {
                        p0 += (p0 ^ 127)
                        p0 += p0
                        if (51 <= 1) {
                            v19 := []int{(fi16 << ((0) & 63))}
                            fmt.Println(fi16)
                            p0 = (^(p0))
                            fmt.Println(fi11, p0, p0)
                            _ = v19
                        } else {
                            v20 := GenS2{F0: -1, F1: 206}
                            var v21 bool = ("hello" != "hello")
                            _ = v20
                            _ = v21
                        }
                    }
                    var v22 int32 = 298
                    var v23 float64 = (0.0 * -2.25)
                    _ = v22
                    _ = v23
                } else {
                    if ((90 < 4294967295) || false) {
                        var v24 int8 = (-(p0))
                        var v25 float64 = 0.0
                        _ = v24
                        _ = v25
                    }
                }
                p0 = (p0 | p0)
                v26 := GenS3{F0: 1000, F1: 169, F2: -957}
                var v27 uint = 0
                var v28 float32 = -2.25
                _ = fi16
                _ = v26
                _ = v27
                _ = v28
            }
        }
        _ = fi11
    }
    _ = p0
    return "a"
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{42}
    var v3 float64 = 3.14159
    v4 := []uint32{}
    var v5 int = -2147483648
    var v6 float32 = 0.1
    switch 4294967295 {
    case 180:
        fmt.Println(v5)
    case -9223372036854775808:
        if (1 >= 0) {
            v1[("foo bar" + ("hello" + "a"))] = ""
        } else {
            if (!true) {
                v5 = v5
                var v29 uint16 = 255
                fmt.Println(v6, v6)
                _ = v29
            }
        }
    }
    var v30 int8 = 42
    fmt.Println(v6)
    var v31 uint32 = 2
    switch "a" {
    case "foo bar":
        if (true || false) {
            fmt.Println(v30, v5, v3)
        } else {
            fmt.Println(v6, v5, v30)
            if len(v4) > 0 {
                v4[0] = v31
            }
            switch (("a" + "") + ((("\n\t" + "hello") + "\n\t") + "")) {
            case "\n\t":
                v32 := v1[""]
                _ = v32
            default:
                v33 := GenS3{F0: 1, F1: (v5 * v5), F2: 0}
                _ = v33
            }
        }
    case "":
        v1["foo bar"] = ("" + "\n\t")
    default:
        var v34 int32 = -964
        _ = v34
    }
    v2 = append(v2, (^(v31)))
    for fi35 := 0; fi35 < 16; fi35++ {
        v36 := GenS2{F0: 1, F1: 18446744073709551615}
        fmt.Println(v30, v30, v6)
        _ = fi35
        _ = v36
    }
    var v37 uint = 0
    var v38 uint32 = (v31 / (((v31 - (^(v31)))) | 1))
    var v39 float64 = v3
    var v40 uint16 = 2
    for fi41 := 0; fi41 < 19; fi41++ {
        v5 -= (fi41 + 127)
        v39 = 3.14159
        var v42 uint32 = 2
        _ = fi41
        _ = v42
    }
    fmt.Println(v40, v39, v3)
    var v43 int64 = -747
    if (-1 <= int16(v5)) {
        v44 := GenS1{F0: 0, F1: v40, F2: 439}
        var v45 int32 = -2147483648
        var v46 float64 = (v3 * v3)
        _ = v44
        _ = v45
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v5
    _ = v6
}
