package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int16; F1 int }

func main() {
    var v1 string = "a"
    v2 := make(map[string]bool)
    v3 := []int{0, 1, 1, 128}
    var v4 int16 = -1
    var v5 uint8 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v6 int8 = 42
    for fi7 := 0; fi7 < 15; fi7++ {
        v5 = v5
        _ = fi7
    }
    if ((!false) && true) {
        var v8 int16 = v4
        fmt.Println(v5)
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            v2["hello"] = true
            v11 := GenS1{F0: -411}
            _ = fi10
            _ = v11
        }
        var v12 int32 = 921
        _ = fi9
        _ = v12
    }
    fmt.Println(v1, v5, v4)
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v4)
    case "":
        delete(v2, v1)
    }
    var v13 int64 = -9223372036854775808
    if false {
        fmt.Println(v4)
        v14 := v2[v1]
        var v15 int32 = 0
        _ = v14
        _ = v15
    } else {
        var v16 float32 = 0.0
        v17 := GenS2{F0: 407}
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: 292}
    switch ("foo bar" + v1) {
    case "\n\t":
        for fi19 := 0; fi19 < 19; fi19++ {
            fmt.Println(v4, v5)
            v13 = -771
            _ = fi19
        }
    case "":
        var v20 float64 = float64(3.14159)
        _ = v20
    case "hello":
        v21 := GenS1{F0: 42}
        _ = v21
    default:
        var v22 int32 = -2147483648
        _ = v22
    }
    fmt.Println(v6)
    for fi23 := 0; fi23 < 13; fi23++ {
        if true {
            for fi24 := 0; fi24 < 11; fi24++ {
                v2[v1] = (false || true)
                switch 44 {
                case 128:
                    if false {
                        v25 := GenS2{F0: 0}
                        v26 := GenS3{F0: (fi23 * 343), F1: false}
                        fmt.Println(v6, v6, v1)
                        var v27 int64 = -226
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                case 32767:
                    v28 := GenS1{F0: fi24}
                    _ = v28
                }
                _ = fi24
            }
        } else {
            for fi29 := 0; fi29 < 13; fi29++ {
                v13 -= v13
                if (false || true) {
                    v6 = 42
                    v2[("hello" + v1)] = true
                    fmt.Println(v6, fi29, v13)
                } else {
                    var v30 uint8 = (v5 & v5)
                    v31 := GenS1{F0: fi29}
                    _ = v30
                    _ = v31
                }
                _ = fi29
            }
        }
        if true {
            v32 := GenS1{F0: (fi23 & (-(fi23)))}
            delete(v2, (v1 + v1))
            _ = v32
        }
        if (false && ((!(1 <= 177)) && false)) {
            if false {
                var v33 uint8 = v5
                v34 := GenS1{F0: -2147483648}
                _ = v33
                _ = v34
            } else {
                fmt.Println(fi23, fi23, v1)
                if true {
                    var v35 uint = 2
                    _ = v35
                } else {
                    if false {
                        delete(v2, ((v1 + v1) + "foo bar"))
                        v36 := GenS4{F0: v4, F1: -9223372036854775808}
                        _ = v36
                    } else {
                        fmt.Println(v13, v4, v13)
                    }
                    for fi37 := 0; fi37 < 14; fi37++ {
                        v38 := v2["\n\t"]
                        _ = fi37
                        _ = v38
                    }
                }
                v2[(v1 + v1)] = false
            }
        } else {
            for fi39 := 0; fi39 < 14; fi39++ {
                _ = len(v2)
                var v40 int8 = (v6 & (v6 | v6))
                if false {
                    fmt.Println(v5, v4)
                    _ = len(v2)
                } else {
                    v40 = v40
                    if (2 > 1) {
                        v41 := GenS4{F0: (v4 ^ v4), F1: 0}
                        _ = v41
                    }
                }
                _ = fi39
                _ = v40
            }
            for fi42 := 0; fi42 < 15; fi42++ {
                if true {
                    var v43 float32 = ((3.14159 / float32(-2.25)) - ((0.0 / 1.0) * 1e10))
                    _ = v43
                }
                _ = fi42
            }
            if (1 > 42) {
                for fi44 := 0; fi44 < 17; fi44++ {
                    v1 = v1
                    var v45 int16 = v4
                    _ = fi44
                    _ = v45
                }
            }
        }
        _ = fi23
    }
    v1 = "foo bar"
    var v46 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
