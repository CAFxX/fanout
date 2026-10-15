package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int64; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string }
func genFn0(p0 bool, p1 string) uint {
    var v6 float32 = 1e10
    var v7 uint8 = 192
    v6 -= (v6 - 3.14159)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 uint, p1 uint64) bool {
    v8 := []int{42, 2147483647, 1}
    switch 9223372036854775807 {
    case 900:
        _ = len(v8)
    case 128:
        var v9 float32 = 1.5
        _ = v9
    case 4294967295:
        _ = len(v8)
    default:
        v10 := GenS1{F0: 1, F1: 0}
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if (("foo bar" != ("a" + "hello")) && ((((100 == 100) || false) || true) || true)) {
            var v12 float32 = 1.5
            p0 = genFn0(((fi11 - (fi11 - 0)) < fi11), "\n\t")
            _ = v12
        } else {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            p0 = genFn0(true, "hello")
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := GenS1{F0: 65535, F1: 42}
                v15 := GenS2{F0: 176, F1: (("a" + ("\n\t" + "")) != "a"), F2: 0}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        _ = fi11
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = p1
    _ = v8
    return false
}

func main() {
    var v1 uint16 = 2
    var v2 int16 = -32768
    var v3 uint = 840
    var v4 string = (("hello" + "\n\t") + "a")
    var v5 int64 = 42
    v16 := []int{}
    for fi17 := 0; fi17 < 20; fi17++ {
        v5 -= (v5 ^ 609)
        _ = len(v16)
        var v18 string = "hello"
        _ = fi17
        _ = v18
    }
    v3 = (v3 | v3)
    var v19 int8 = 100
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 int32 = 42
        var v22 uint32 = 4294967295
        _ = fi20
        _ = v21
        _ = v22
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        if (!(0 >= -1)) {
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 int8 = v19
                _ = fi24
                _ = v25
            }
            if false {
                v26 := GenS1{F0: 934, F1: 1}
                _ = v26
            } else {
                var v27 uint32 = 2
                _ = v27
            }
        }
        var v28 uint32 = 2
        v29 := GenS2{F0: v5, F1: (((v2 / (((^(v2))) | 1)) != v2) && true), F2: 173}
        _ = fi23
        _ = v28
        _ = v29
    }
    var v30 uint = genFn0(false, v4)
    var v31 uint16 = (v1 - 42)
    for fi32 := 0; fi32 < 3; fi32++ {
        for fi33 := 0; fi33 < 10; fi33++ {
            if (fi32 == fi32) {
                v34 := GenS2{F0: v5, F1: false, F2: (v31 + 0)}
                _ = v34
            }
            v35 := GenS3{F0: (v31 % (((v31 >> ((v3) & 63))) | 1))}
            var v36 int64 = v5
            v1 = v31
            var v37 string = (((((v4 + v4) + "a") + "") + "a") + v4)
            fmt.Println(v3)
            _ = fi33
            _ = v35
            _ = v36
            _ = v37
        }
        switch -9223372036854775808 {
        case 4294967295:
            v16 = append(v16, fi32)
        }
        v38 := GenS4{F0: v5}
        _ = fi32
        _ = v38
    }
    v5 += (v5 - int64(v3))
    for fi39 := 0; fi39 < 12; fi39++ {
        v40 := GenS2{F0: (v5 & 9223372036854775807), F1: ((-1 <= 701) && ((v5 << (((v30 >> (((v3 + uint(v31))) & 63))) & 63)) < -9223372036854775808)), F2: (^(v31))}
        v19 -= v19
        _ = fi39
        _ = v40
    }
    switch "a" {
    case "hello":
        for fi41 := 0; fi41 < 16; fi41++ {
            if true {
                fmt.Println(v4, v4)
            } else {
                v42 := GenS1{F0: 883, F1: 42}
                fmt.Println(fi41, v3, v2)
                v43 := GenS5{F0: (v4 + "\n\t")}
                _ = v42
                _ = v43
            }
            v16 = append(v16, fi41)
            _ = fi41
        }
    case "":
        switch "hello" {
        case "\n\t":
            var v44 uint64 = 1
            _ = v44
        case "a":
            if (v5 == v5) {
                if true {
                    if ((0.0 + 1.5) >= (3.14159 + float64(-2.25))) {
                        v5 = 65
                    }
                    var v45 string = v4
                    v2 += (v2 % (((-(v2))) | 1))
                    _ = v45
                }
                var v46 bool = false
                v31 = v31
                _ = v46
            } else {
                if true {
                    v47 := GenS1{F0: 2147483648, F1: -1}
                    if len(v16) > 0 {
                        v16[0] = int(v2)
                    }
                    for fi48 := 0; fi48 < 17; fi48++ {
                        for fi49 := 0; fi49 < 18; fi49++ {
                            v16 = append(v16, -876)
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            _ = fi49
                        }
                        fmt.Println(v31)
                        var v50 int64 = (v5 / ((int64(v2)) | 1))
                        _ = fi48
                        _ = v50
                    }
                    _ = v47
                }
            }
        case "hello":
            var v51 uint = v3
            _ = v51
        }
    case "\n\t":
        for fi52 := 0; fi52 < 18; fi52++ {
            for fi53 := 0; fi53 < 17; fi53++ {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = fi53
            }
            v54 := GenS4{F0: (v5 / ((v5) | 1))}
            _ = fi52
            _ = v54
        }
    default:
        v19 = (v19 ^ v19)
    }
    _ = len(v16)
    var v55 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v55
}
