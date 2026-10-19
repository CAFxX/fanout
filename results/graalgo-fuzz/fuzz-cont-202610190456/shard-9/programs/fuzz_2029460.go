package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int16; F1 uint8; F2 int16 }
func genFn0() int16 {
    for fi4 := 0; fi4 < 12; fi4++ {
        switch "a" {
        case "a":
            if (46 < 193) {
                var v5 int32 = 1
                _ = v5
            }
        case "\n\t":
            var v6 float32 = 1e10
            _ = v6
        case "foo bar":
            v7 := GenS1{F0: 127, F1: fi4, F2: 42}
            _ = v7
        default:
            if (false && true) {
                switch ("" + "foo bar") {
                case "":
                    switch (^(fi4)) {
                    case -2147483648:
                        var v8 int = (fi4 % ((fi4) | 1))
                        _ = v8
                    case -484:
                        for fi9 := 0; fi9 < 20; fi9++ {
                            var v10 uint16 = 2
                            v11 := GenS1{F0: 127, F1: 489, F2: 526}
                            var v12 uint16 = v10
                            _ = fi9
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                    case 4294967295:
                        for fi13 := 0; fi13 < 20; fi13++ {
                            var v14 uint = 4294967295
                            fmt.Println(fi13, v14)
                            var v15 float32 = 0.0
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    }
                case "\n\t":
                    v16 := []int{}
                    _ = v16
                case "foo bar":
                    var v17 uint = 243
                    _ = v17
                default:
                    var v18 int8 = -1
                    _ = v18
                }
                var v19 int64 = 400
                switch (fi4 * fi4) {
                case 1:
                    v20 := []int{fi4, (fi4 % ((fi4) | 1)), fi4}
                    _ = v20
                case 2147483647:
                    for fi21 := 0; fi21 < 12; fi21++ {
                        var v22 uint16 = 487
                        v19 += 337
                        _ = fi21
                        _ = v22
                    }
                }
                _ = v19
            } else {
                for fi23 := 0; fi23 < 16; fi23++ {
                    v24 := []int{fi4}
                    switch (fi4 << ((821) & 63)) {
                    case 349:
                        var v25 uint8 = 2
                        _ = v25
                    }
                    v26 := GenS1{F0: int8(fi23), F1: fi4, F2: 2}
                    _ = fi23
                    _ = v24
                    _ = v26
                }
                var v27 string = "foo bar"
                var v28 bool = true
                _ = v27
                _ = v28
            }
        }
        _ = fi4
    }
    var v29 uint8 = 127
    for fi30 := 0; fi30 < 6; fi30++ {
        v31 := []int{-9223372036854775808, fi30}
        var v32 bool = (!false)
        v32 = true
        _ = fi30
        _ = v31
        _ = v32
    }
    _ = v29
    return 32767
}
func genFn1(p0 string, p1 uint64) int64 {
    var v33 bool = (("" + (p0 + p0)) > p0)
    switch "" {
    case "hello":
        v33 = (1.5 >= 1.5)
    case "foo bar":
        var v34 int16 = -163
        _ = v34
    case "":
        var v35 uint8 = 255
        _ = v35
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v33
    return -361
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{9223372036854775807, 47}
    var v3 uint8 = 1
    v36 := GenS1{F0: -1, F1: -117, F2: 2}
    v37 := GenS2{F0: 0, F1: -128}
    switch 9223372036854775807 {
    case -9223372036854775808:
        fmt.Println(v3)
    }
    v38 := GenS2{F0: 137, F1: -1}
    switch 4294967295 {
    case -106:
        v3 = v3
    case 0:
        for fi39 := 0; fi39 < 16; fi39++ {
            v3 -= v3
            fmt.Println(fi39, fi39)
            _ = fi39
        }
    case 883:
        if true {
            for fi40 := 0; fi40 < 20; fi40++ {
                fmt.Println(v3, fi40)
                var v41 float32 = 3.14159
                _ = fi40
                _ = v41
            }
        }
    }
    v3 = v3
    var v42 uint64 = 4294967295
    if true {
        if (642 == 4294967295) {
            var v43 uint = 2
            _ = v43
        } else {
            v44 := GenS3{F0: genFn0(), F1: v3, F2: 0}
            _ = v44
        }
        if (-2.25 != 1e10) {
            v3 += (^(v3))
        }
        v45 := GenS3{F0: 0, F1: v3, F2: -242}
        _ = v45
    }
    v3 = 128
    for fi46 := 0; fi46 < 1; fi46++ {
        v47 := GenS1{F0: 100, F1: (fi46 & fi46), F2: 182}
        for fi48 := 0; fi48 < 17; fi48++ {
            if (false && ((v3 ^ 128) < 2)) {
                var v49 uint32 = 4294967295
                _ = v49
            } else {
                v42 -= (v42 ^ (v42 % (((v42 >> ((1) & 63))) | 1)))
                fmt.Println(fi48, v3)
                v3 = (v3 & (^(v3)))
            }
            if len(v2) > 0 {
                v2[0] = -2147483648
            }
            if false {
                fmt.Println(v3, fi46)
                var v50 uint = 19
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v50
            }
            _ = fi48
        }
        _ = fi46
        _ = v47
    }
    var v51 int = 32767
    v51 -= 819
    v52 := GenS2{F0: v42, F1: -1}
    var v53 int16 = -1
    v53 += v53
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v42
    _ = v51
    _ = v52
    _ = v53
}
