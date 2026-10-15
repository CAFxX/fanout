package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint32 }
func genFn0(p0 int) bool {
    var v4 float32 = 3.14159
    if false {
        for fi5 := 0; fi5 < 8; fi5++ {
            v6 := []int{(p0 | p0), -1, fi5, 0}
            p0 -= fi5
            for fi7 := 0; fi7 < 18; fi7++ {
                var v8 int8 = -128
                v8 = v8
                fmt.Println(v8, fi7)
                _ = fi7
                _ = v8
            }
            _ = fi5
            _ = v6
        }
        v9 := []int{9223372036854775807, (^(p0)), (^(p0)), (p0 * p0)}
        var v10 float32 = v4
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 3; fi11++ {
            if false {
                v12 := GenS1{F0: (("" + "hello") + ("" + ((("" + "\n\t") + "foo bar") + ""))), F1: (p0 ^ 2147483647), F2: 2147483647}
                v13 := GenS1{F0: "foo bar", F1: fi11, F2: 4294967295}
                _ = v12
                _ = v13
            } else {
                var v14 int = -766
                v14 += fi11
                var v15 int16 = 903
                _ = v14
                _ = v15
            }
            var v16 float32 = -2.25
            _ = fi11
            _ = v16
        }
    }
    v17 := []int{(-(p0))}
    v18 := GenS1{F0: "a", F1: -9223372036854775808, F2: 42}
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    _ = p0
    _ = v17
    _ = v18
    _ = v4
    return false
}

func main() {
    var v1 uint8 = 243
    var v2 string = "foo bar"
    v3 := make(map[uint32]string)
    var v19 uint16 = 32767
    var v20 int = 1
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 uint16 = v19
        if (!((1 >= 32767) || false)) {
            v23 := GenS1{F0: v2, F1: (fi21 & v20), F2: 2147483647}
            if false {
                fmt.Println(v20, v2, v19)
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 int = (v20 % ((4294967295) | 1))
                    fmt.Println(v1, v1)
                    var v26 int = -1
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v2 = "foo bar"
            } else {
                v2 = (v2 + v2)
            }
            v20 += (^(fi21))
            _ = v23
        }
        v19 -= (v22 | v19)
        _ = fi21
        _ = v22
    }
    v19 = 65535
    v2 = ((((v2 + (v2 + ((("hello" + "") + v2) + v2))) + "a") + v2) + v2)
    switch v20 {
    case -9223372036854775808:
        var v27 string = (v2 + ((v2 + v2) + v2))
        _ = v27
    default:
        var v28 float64 = float64(float32(3.14159))
        _ = v28
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        v30 := []int{}
        fmt.Println(v19, v19)
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 float32 = 0.1
            v33 := GenS1{F0: "foo bar", F1: fi31, F2: 42}
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            v32 -= v32
            v32 += ((v32 / 1.5) - v32)
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi29
        _ = v30
    }
    var v34 uint32 = 279
    v1 -= (v1 | v1)
    var v35 int32 = 0
    if ((v1 - v1) >= (v1 >> ((42) & 63))) {
        fmt.Println(v2)
        if true {
            if true {
                v36 := GenS1{F0: v2, F1: (v20 / ((0) | 1)), F2: 2147483648}
                _ = v36
            }
            var v37 int16 = 1
            v35 = (v35 << ((4294967295) & 63))
            _ = v37
        }
    }
    var v38 uint8 = v1
    var v39 string = "a"
    v38 = (^(v38))
    if (!(false && false)) {
        v40 := GenS2{F0: v20, F1: (^(v34))}
        _ = v40
    } else {
        v41 := GenS3{F0: 42, F1: (v34 & 4294967295)}
        _ = v41
    }
    v34 = v34
    for fi42 := 0; fi42 < 12; fi42++ {
        if true {
            var v43 int64 = 0
            v44 := []int{}
            _ = v43
            _ = v44
        }
        switch "foo bar" {
        case "hello":
            v35 = -234
        case "\n\t":
            var v45 uint64 = 4294967295
            _ = v45
        case "":
            v46 := GenS1{F0: ("hello" + v2), F1: fi42, F2: 42}
            _ = v46
        }
        _ = fi42
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v39
}
