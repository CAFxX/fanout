package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64 }
func genFn0() uint64 {
    var v7 int32 = 0
    v8 := GenS1{F0: 2147483647}
    v9 := GenS2{F0: 32767, F1: (v7 >> ((816) & 63))}
    v7 = (^(v7))
    _ = v7
    _ = v8
    _ = v9
    return 537
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 346
    v3 := []string{((("foo bar" + ((("foo bar" + "") + ((((("a" + "hello") + "a") + "") + "") + "hello")) + ("" + "hello"))) + "a") + "a"), "\n\t", (("a" + "a") + "\n\t"), ((("foo bar" + "a") + ("hello" + "\n\t")) + "\n\t")}
    v4 := []string{((((("\n\t" + "foo bar") + "\n\t") + "foo bar") + "a") + "foo bar")}
    v5 := []string{("foo bar" + ("hello" + "")), "hello"}
    var v6 int32 = -377
    if (254 == 255) {
        v10 := GenS2{F0: 255, F1: v6}
        v11 := GenS1{F0: 237}
        if true {
            v12 := v1[-136]
            var v13 float64 = float64((float32(1e10) * -2.25))
            _ = v12
            _ = v13
        } else {
            if ((true && (65535 > (v2 ^ v2))) && false) {
                if false {
                    var v14 string = (((("a" + "hello") + ("foo bar" + "a")) + "a") + "hello")
                    _ = v14
                }
                var v15 float32 = (-2.25 + (1e10 + 3.14159))
                _ = v15
            }
            switch 2147483647 {
            case 32767:
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 string = "foo bar"
                    v17 = ((v17 + (v17 + (("\n\t" + "hello") + v17))) + "a")
                    v17 = (v17 + v17)
                    _ = fi16
                    _ = v17
                }
            case 2147483647:
                fmt.Println(v2, v2)
            case -847:
                v6 = (^(v6))
            }
            _ = len(v5)
        }
        _ = v10
        _ = v11
    }
    v18 := GenS3{F0: 255, F1: 42}
    var v19 bool = (9223372036854775807 <= -1)
    if v19 {
        if (9223372036854775807 <= int(v2)) {
            v20 := GenS1{F0: 9223372036854775807}
            _ = len(v3)
            _ = v20
        }
    } else {
        v19 = v19
    }
    v21 := GenS1{F0: -463}
    var v22 int64 = 0
    for fi23 := 0; fi23 < 9; fi23++ {
        v5 = append(v5, "")
        v6 += v6
        var v24 uint64 = 108
        _ = fi23
        _ = v24
    }
    var v25 bool = v19
    switch -804 {
    case -266:
        if ((^(v2)) != 4294967295) {
            v26 := GenS4{F0: 566}
            var v27 uint8 = 76
            v28 := GenS5{F0: -484}
            _ = v26
            _ = v27
            _ = v28
        } else {
            var v29 uint16 = 1
            _ = v29
        }
    case 1000:
        if v19 {
            _ = len(v5)
            fmt.Println(v22, v22, v19)
        }
    default:
        v25 = (v19 || false)
    }
    switch ("" + "foo bar") {
    case "":
        var v30 int16 = -32768
        _ = v30
    case "hello":
        if v25 {
            v31 := GenS4{F0: 1}
            _ = v31
        }
    default:
        v32 := GenS4{F0: 0}
        _ = v32
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 string = "\n\t"
        if v25 {
            switch v34 {
            case "foo bar":
                for fi35 := 0; fi35 < 20; fi35++ {
                    switch (v34 + v34) {
                    case "\n\t":
                        v36 := GenS3{F0: 695, F1: v22}
                        _ = v36
                    case "foo bar":
                        fmt.Println(fi35, v22)
                    default:
                        v37 := GenS3{F0: 1, F1: v22}
                        _ = v37
                    }
                    v38 := GenS3{F0: 193, F1: 0}
                    _ = fi35
                    _ = v38
                }
            default:
                v34 = "a"
            }
        } else {
            if (-2147483648 == (v6 / ((v6) | 1))) {
                delete(v1, fi33)
                var v39 uint64 = genFn0()
                _ = v39
            } else {
                var v40 float32 = ((1.5 * (0.1 / 0.1)) - (0.0 - 0.1))
                _ = v40
            }
        }
        _ = fi33
        _ = v34
    }
    switch -9223372036854775808 {
    case 32768:
        for fi41 := 0; fi41 < 12; fi41++ {
            var v42 uint16 = 698
            var v43 int = (^(fi41))
            var v44 bool = (v19 && ((true || ((v42 * 561) > (v42 + v42))) && (float32(1.5) < 1e10)))
            _ = fi41
            _ = v42
            _ = v43
            _ = v44
        }
    case 9223372036854775807:
        var v45 uint = 259
        _ = v45
    case -9223372036854775808:
        fmt.Println(v6, v2, v6)
    default:
        v4 = append(v4, "\n\t")
    }
    v46 := GenS2{F0: 42, F1: (-(v6))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
