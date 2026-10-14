package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 uint16 }

func main() {
    var v1 float64 = (3.14159 + float64(3.14159))
    var v2 uint = 0
    v3 := make(map[string]string)
    var v4 uint64 = 129
    v1 = float64((0.1 - (3.14159 / (1.5 - ((0.1 * 1e10) - (-2.25 + 0.1))))))
    if (!(!false)) {
        var v5 uint32 = 2147483647
        _ = v5
    } else {
        var v6 int8 = -1
        var v7 int8 = v6
        fmt.Println(v1, v4)
        _ = v6
        _ = v7
    }
    v8 := []int{}
    for fi9 := 0; fi9 < 20; fi9++ {
        delete(v3, "\n\t")
        var v10 string = "foo bar"
        v11 := GenS1{F0: "", F1: 65535, F2: 619}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint = (v2 / ((v2) | 1))
    v1 -= 1e10
    var v13 float32 = 3.14159
    v14 := GenS2{F0: (v4 + (v4 % ((2) | 1))), F1: 65535}
    v15 := GenS3{F0: -2147483648}
    if false {
        var v16 int64 = -600
        v17 := GenS4{F0: -45}
        for fi18 := 0; fi18 < 17; fi18++ {
            if (v4 > 18446744073709551615) {
                v4 -= (v4 * (v4 & v4))
            }
            v13 -= -2.25
            _ = fi18
        }
        _ = v16
        _ = v17
    }
    if false {
        v4 = v4
        switch 32768 {
        case 42:
            fmt.Println(v4, v12)
        }
        if false {
            v13 = (v13 - (v13 - v13))
            var v19 uint16 = 2
            v20 := GenS5{F0: v19}
            _ = v19
            _ = v20
        }
    }
    switch ((("" + "foo bar") + "\n\t") + "") {
    case "":
        switch 32767 {
        case 127:
            fmt.Println(v2, v1, v2)
        case 106:
            for fi21 := 0; fi21 < 13; fi21++ {
                for fi22 := 0; fi22 < 14; fi22++ {
                    v1 = float64(0.1)
                    switch "a" {
                    case "\n\t":
                        var v23 int = (fi22 & (fi21 >> (((v12 / ((333) | 1))) & 63)))
                        _ = v23
                    case "foo bar":
                        v1 = (v1 * ((v1 * v1) - v1))
                    }
                    _ = fi22
                }
                fmt.Println(v2, v13, v12)
                _ = fi21
            }
        case 42:
            fmt.Println(v4, v4)
        }
    default:
        v13 = float32(v1)
    }
    switch "\n\t" {
    case "a":
        v24 := v3["foo bar"]
        _ = v24
    case "hello":
        for fi25 := 0; fi25 < 11; fi25++ {
            v2 -= 255
            var v26 int16 = 1000
            v13 -= (v13 + 0.1)
            _ = fi25
            _ = v26
        }
    case "":
        v27 := GenS6{F0: (v12 / ((v12) | 1)), F1: 4294967295}
        _ = v27
    default:
        _ = len(v3)
    }
    v28 := v3["\n\t"]
    for fi29 := 0; fi29 < 1; fi29++ {
        if true {
            var v30 bool = (!(!false))
            v12 = (v12 / ((v2) | 1))
            var v31 int32 = -665
            _ = v30
            _ = v31
        } else {
            fmt.Println(fi29)
            v2 += (^(v12))
            fmt.Println(v1, v28, v12)
        }
        v13 = (v13 * 3.14159)
        _ = fi29
    }
    switch 0 {
    case 0:
        v32 := GenS4{F0: -1}
        _ = v32
    case -1:
        v33 := GenS2{F0: v4, F1: 812}
        _ = v33
    case 20:
        if (!(239 <= 2)) {
            var v34 uint16 = 32767
            var v35 float32 = (((float32(0.1) - (v13 + v13)) + v13) * 0.0)
            _ = v34
            _ = v35
        } else {
            var v36 float64 = (v1 / (v1 / (v1 + (v1 + (3.14159 * 1e10)))))
            var v37 float32 = float32(v36)
            if true {
                var v38 uint64 = (v4 & 808)
                var v39 float32 = v13
                _ = v38
                _ = v39
            } else {
                if false {
                    var v40 int16 = 42
                    _ = v40
                } else {
                    var v41 float32 = ((v37 / v37) + (v37 - v13))
                    _ = v41
                }
            }
            _ = v36
            _ = v37
        }
    }
    switch -774 {
    case -1:
        v28 = v28
    case 1:
        for fi42 := 0; fi42 < 5; fi42++ {
            var v43 uint = (^(v12))
            if (-1 < -1) {
                v44 := GenS2{F0: (v4 / (((v4 << ((v2) & 63))) | 1)), F1: 42}
                _ = v44
            }
            v1 += v1
            _ = fi42
            _ = v43
        }
    case 32767:
        v45 := GenS3{F0: 398}
        _ = v45
    default:
        if true {
            var v46 int64 = -9223372036854775808
            _ = len(v8)
            _ = v46
        } else {
            var v47 int16 = 581
            v48 := GenS7{F0: 393}
            _ = v47
            _ = v48
        }
    }
    var v49 uint64 = (v4 >> (((v12 - v12)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v49
    _ = v8
}
