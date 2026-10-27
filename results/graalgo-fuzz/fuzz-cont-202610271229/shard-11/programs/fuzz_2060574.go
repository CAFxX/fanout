package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int16; F1 int; F2 int16 }
type GenS5 struct { F0 bool; F1 uint16 }

func main() {
    var v1 int16 = 1000
    var v2 uint16 = 205
    var v3 int64 = -9223372036854775808
    v4 := []string{("\n\t" + "a"), "foo bar", ("\n\t" + "\n\t"), "a"}
    v5 := make(map[string]string)
    var v6 uint64 = 1
    var v7 int64 = (v3 * (v3 / ((42) | 1)))
    if true {
        v8 := GenS1{F0: 1, F1: 4294967295, F2: "a"}
        v1 = 42
        _ = v8
    }
    var v9 float32 = 3.14159
    var v10 float64 = (float64(v9) / 1.5)
    var v11 float32 = ((v9 + (0.0 / v9)) / (-2.25 * ((v9 - v9) - (v9 + v9))))
    var v12 float64 = 1e10
    if (!true) {
        var v13 int16 = (v1 & (-(v1)))
        _ = v13
    } else {
        v6 -= v6
        if false {
            v14 := GenS2{F0: 42}
            var v15 float32 = v9
            v16 := GenS2{F0: -1}
            _ = v14
            _ = v15
            _ = v16
        }
        fmt.Println(v9)
    }
    v17 := GenS2{F0: -145}
    switch -388 {
    case 9223372036854775807:
        var v18 int32 = 2147483647
        _ = v18
    case 127:
        if true {
            for fi19 := 0; fi19 < 8; fi19++ {
                for fi20 := 0; fi20 < 7; fi20++ {
                    if (-2.25 < v10) {
                        v12 = (v10 * (v12 - v10))
                        _ = len(v5)
                        v21 := GenS3{F0: 42, F1: ""}
                        if false {
                            _ = len(v4)
                        }
                        _ = v21
                    } else {
                        var v22 int8 = 0
                        if true {
                            var v23 float32 = v9
                            delete(v5, ("" + "foo bar"))
                            _ = v23
                        } else {
                            fmt.Println(v1)
                        }
                        v24 := GenS4{F0: (v1 << ((42) & 63)), F1: -9223372036854775808, F2: 42}
                        var v25 uint64 = 0
                        if false {
                            fmt.Println(v1, v6, v12)
                            fmt.Println(v1)
                        } else {
                            v25 -= v6
                            v22 = 109
                            v26 := GenS1{F0: v22, F1: 65535, F2: ("a" + "a")}
                            _ = v26
                        }
                        _ = v22
                        _ = v24
                        _ = v25
                    }
                    _ = fi20
                }
                if (fi19 > fi19) {
                    delete(v5, "foo bar")
                }
                _ = fi19
            }
            v27 := GenS4{F0: v1, F1: 127, F2: -1}
            _ = v27
        } else {
            v9 = (v9 * v9)
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 int32 = 42
                if true {
                    for fi30 := 0; fi30 < 7; fi30++ {
                        if (((^(fi28)) != int(v2)) && ((!(true || true)) || (-29 == int8(v7)))) {
                            v31 := GenS5{F0: ("hello" < ("\n\t" + "a")), F1: v2}
                            v32 := v5["\n\t"]
                            _ = v31
                            _ = v32
                        } else {
                            v11 += (((v9 * 1e10) * v9) * 3.14159)
                            var v33 uint = 255
                            _ = v33
                        }
                        var v34 bool = false
                        _ = fi30
                        _ = v34
                    }
                    v35 := GenS1{F0: 42, F1: 4294967295, F2: "foo bar"}
                    _ = v35
                } else {
                    _ = len(v5)
                    switch "hello" {
                    case "hello":
                        fmt.Println(v9)
                    case "foo bar":
                        v36 := v5[("hello" + "foo bar")]
                        _ = v36
                    case "a":
                        v37 := v5[""]
                        _ = v37
                    default:
                        var v38 int16 = int16(v2)
                        _ = v38
                    }
                    if ((true && false) && ((v3 >> ((2) & 63)) == v3)) {
                        switch (fi28 * fi28) {
                        case 32768:
                            v7 += (v3 % ((v3) | 1))
                        case 0:
                            v2 = (^(v2))
                        case 4294967295:
                            v2 = 32767
                        }
                        fmt.Println(v29)
                    } else {
                        _ = len(v4)
                    }
                }
                _ = fi28
                _ = v29
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
