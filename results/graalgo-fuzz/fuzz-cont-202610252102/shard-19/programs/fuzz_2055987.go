package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint16) uint {
    p0 = p0
    var v7 string = (("foo bar" + ("a" + "hello")) + ("a" + ("foo bar" + "\n\t")))
    v8 := GenS1{F0: 42, F1: 0}
    fmt.Println(p0, v7)
    _ = p0
    _ = v7
    _ = v8
    return 100
}

func main() {
    var v1 int = -1
    v2 := []string{"\n\t", "a", ("a" + "a")}
    var v3 int16 = 32767
    var v4 int32 = 42
    var v5 uint16 = 275
    var v6 int64 = -333
    var v9 float32 = 0.0
    v6 += (v6 - 0)
    for fi10 := 0; fi10 < 17; fi10++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi10
    }
    v11 := GenS1{F0: 396, F1: (v1 & v1)}
    if ((-108 != 0) || (814 < 42)) {
        var v12 int8 = 1
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        if len(v2) > 0 {
            v2[0] = ("\n\t" + "")
        }
        _ = v12
    }
    switch 353 {
    case -498:
        var v13 uint8 = 170
        _ = v13
    case 32768:
        switch -1 {
        case -556:
            v3 = (v3 | (v3 | (^(v3))))
        default:
            if len(v2) > 0 {
                v2[0] = ""
            }
        }
    case -523:
        v2 = append(v2, "a")
    default:
        var v14 int16 = v3
        _ = v14
    }
    var v15 int8 = -70
    var v16 int32 = v4
    v5 -= (v5 * 970)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((18446744073709551615 < 1) || false) {
        v17 := GenS2{F0: 4294967295}
        _ = v17
    } else {
        var v18 uint = 921
        if true {
            for fi19 := 0; fi19 < 20; fi19++ {
                v15 = (v15 / (((v15 >> ((v18) & 63))) | 1))
                v5 -= v5
                _ = fi19
            }
        }
        var v20 int = (v1 ^ 42)
        _ = v18
        _ = v20
    }
    var v21 uint8 = 1
    v22 := GenS1{F0: 42, F1: v1}
    v2 = append(v2, "a")
    v21 += 2
    if false {
        for fi23 := 0; fi23 < 2; fi23++ {
            for fi24 := 0; fi24 < 12; fi24++ {
                if len(v2) > 0 {
                    v2[0] = ""
                }
                v6 = (^(v6))
                v5 = 32767
                switch (fi24 * int(v15)) {
                case 1:
                    switch "" {
                    case "hello":
                        fmt.Println(v6)
                    case "a":
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                    }
                default:
                    var v25 uint64 = 266
                    _ = v25
                }
                _ = fi24
            }
            v2 = append(v2, "")
            _ = fi23
        }
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := GenS1{F0: uint64(v5), F1: 2147483647}
            if true {
                _ = len(v2)
                var v28 float64 = 0.1
                if true {
                    for fi29 := 0; fi29 < 8; fi29++ {
                        var v30 float64 = float64(1e10)
                        _ = fi29
                        _ = v30
                    }
                }
                _ = v28
            } else {
                if (v6 < (v6 - v6)) {
                    var v31 string = "hello"
                    var v32 int = (v1 ^ (fi26 & (fi26 ^ fi26)))
                    _ = v31
                    _ = v32
                } else {
                    if true {
                        v5 = 0
                        v9 = v9
                    } else {
                        var v33 float64 = -2.25
                        v34 := GenS3{F0: (v15 << ((1) & 63))}
                        _ = v33
                        _ = v34
                    }
                }
            }
            fmt.Println(v16, v3)
            _ = fi26
            _ = v27
        }
        fmt.Println(v6, v1, v9)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
