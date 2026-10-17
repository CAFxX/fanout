package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint16; F2 string }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 1
    var v3 bool = ("foo bar" != "\n\t")
    var v4 uint32 = (v2 << ((4294967295) & 63))
    switch "a" {
    case "foo bar":
        switch -2147483648 {
        case 0:
            if ((1 >= 2) || v3) {
                if true {
                    if (2 >= 128) {
                        var v5 uint64 = 2
                        _ = v5
                    }
                    for fi6 := 0; fi6 < 4; fi6++ {
                        var v7 int8 = 0
                        _ = fi6
                        _ = v7
                    }
                }
                switch (("hello" + "") + "") {
                case "a":
                    var v8 bool = false
                    _ = v8
                case "foo bar":
                    v3 = true
                default:
                    var v9 string = ("\n\t" + "a")
                    _ = v9
                }
            }
        case -2147483648:
            var v10 float32 = 1e10
            _ = v10
        default:
            for fi11 := 0; fi11 < 1; fi11++ {
                fmt.Println(fi11, fi11)
                _ = fi11
            }
        }
    case "\n\t":
        if true {
            var v12 float64 = (-2.25 * 1.5)
            v13 := GenS1{F0: 0}
            _ = v12
            _ = v13
        } else {
            v2 += 195
            fmt.Println(v3, v4)
        }
    case "hello":
        v4 = (v4 / ((v4) | 1))
    }
    v2 -= v2
    if v3 {
        for fi14 := 0; fi14 < 20; fi14++ {
            v3 = (float64(1e10) < 0.0)
            var v15 int8 = 0
            v16 := GenS1{F0: 245}
            _ = fi14
            _ = v15
            _ = v16
        }
        v17 := GenS2{F0: false, F1: 255, F2: (("" + ((("\n\t" + "\n\t") + "foo bar") + "\n\t")) + "foo bar")}
        _ = v17
    }
    switch "foo bar" {
    case "":
        var v18 float32 = -2.25
        _ = v18
    case "a":
        var v19 uint64 = 0
        _ = v19
    case "\n\t":
        v20 := GenS1{F0: -1}
        _ = v20
    default:
        if (!v3) {
            v21 := GenS3{F0: 848, F1: v3}
            switch ("\n\t" + "") {
            case "":
                v2 -= v4
            default:
                v3 = false
            }
            if v3 {
                if v3 {
                    var v22 int64 = -9223372036854775808
                    switch 2147483647 {
                    case -326:
                        var v23 int64 = v22
                        _ = v23
                    default:
                        fmt.Println(v2)
                    }
                    v24 := GenS2{F0: v3, F1: 32767, F2: ((("" + "hello") + (("foo bar" + "foo bar") + "")) + "foo bar")}
                    _ = v22
                    _ = v24
                }
                var v25 int16 = -77
                _ = v25
            }
            _ = v21
        }
    }
    v26 := []int{2147483647, -273, 0, 42}
    var v27 int32 = 882
    v27 = v27
    for fi28 := 0; fi28 < 18; fi28++ {
        var v29 uint32 = v4
        _ = fi28
        _ = v29
    }
    v1[94] = v2
    for fi30 := 0; fi30 < 16; fi30++ {
        v1[fi30] = 813
        var v31 uint32 = v4
        var v32 int16 = 423
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println(v3, v27, v27)
    var v33 int64 = 9223372036854775807
    var v34 uint32 = (v2 << ((uint(v2)) & 63))
    if ((v3 && v3) || (v3 && ("hello" <= ("a" + "")))) {
        v27 -= v27
    } else {
        v35 := GenS4{F0: 4294967295}
        if (!v3) {
            if false {
                switch "hello" {
                case "":
                    var v36 int64 = (-(v33))
                    _ = v36
                case "a":
                    v3 = v3
                default:
                    var v37 uint = 2
                    _ = v37
                }
                fmt.Println(v27, v4)
            }
        } else {
            v33 -= 42
        }
        if (v3 || v3) {
            v27 += (v27 + (v27 - (v27 ^ 42)))
            fmt.Println(v4, v27, v4)
            if v3 {
                v34 = 1
            } else {
                v38 := GenS3{F0: 0, F1: (!v3)}
                if false {
                    for fi39 := 0; fi39 < 3; fi39++ {
                        v40 := v1[(fi39 ^ fi39)]
                        v34 = v2
                        v41 := GenS4{F0: 1}
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                }
                _ = v38
            }
        }
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
