package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int; F1 string }
func genFn0(p0 uint) int64 {
    var v5 float64 = 3.14159
    if (true || true) {
        fmt.Println(p0, v5)
        v5 -= (0.0 * 1.5)
        for fi6 := 0; fi6 < 17; fi6++ {
            p0 -= p0
            v7 := GenS1{F0: (("foo bar" + "hello") + "\n\t"), F1: 1}
            p0 += 18446744073709551615
            _ = fi6
            _ = v7
        }
    } else {
        p0 += (p0 % ((p0) | 1))
        var v8 uint = (p0 | 65535)
        _ = v8
    }
    if true {
        var v9 uint16 = 174
        switch 752 {
        case -870:
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 bool = (-1 <= int8(v9))
                fmt.Println(p0)
                var v12 int64 = -9223372036854775808
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        _ = v9
    }
    var v13 string = (("" + ("a" + ("" + ""))) + "foo bar")
    _ = p0
    _ = v13
    _ = v5
    return 0
}

func main() {
    var v1 int = 127
    var v2 string = "hello"
    var v3 int16 = -1
    v4 := []string{}
    for fi14 := 0; fi14 < 11; fi14++ {
        if (!((!(!(false && (1 > 127)))) || false)) {
            if true {
                _ = len(v4)
                var v15 string = "a"
                _ = v15
            }
            var v16 int16 = (v3 / ((v3) | 1))
            if (true || (228 < -43)) {
                v2 = (v2 + (v2 + "hello"))
                var v17 string = (("foo bar" + "\n\t") + ((v2 + (v2 + v2)) + v2))
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v17
            } else {
                if false {
                    if (false || false) {
                        v18 := GenS2{F0: 32767, F1: genFn0(207)}
                        _ = v18
                    }
                } else {
                    switch "foo bar" {
                    case "a":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case "":
                        var v19 int64 = 0
                        _ = v19
                    case "foo bar":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        var v20 float64 = (0.1 - (3.14159 / 1.5))
                        _ = v20
                    }
                    v4 = append(v4, "\n\t")
                }
                v16 -= -1
            }
            _ = v16
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi14
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = v2
    }
    switch v1 {
    case 0:
        var v21 uint64 = 18446744073709551615
        _ = v21
    case 128:
        var v22 uint16 = 65535
        _ = v22
    case 32767:
        v1 = v1
    }
    var v23 float32 = 3.14159
    var v24 int8 = 127
    for fi25 := 0; fi25 < 16; fi25++ {
        if (!false) {
            if (false || (false && (4294967295 >= 65535))) {
                fmt.Println(v2, fi25, v23)
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = len(v4)
        }
        v26 := GenS3{F0: fi25, F1: (v2 + v2)}
        switch v2 {
        case "\n\t":
            v2 = (v2 + ("hello" + (("" + ("a" + (v2 + "a"))) + v2)))
        case "foo bar":
            v24 += v24
        default:
            v3 = v3
        }
        _ = fi25
        _ = v26
    }
    var v27 int16 = 42
    v28 := GenS1{F0: v2, F1: 8}
    for fi29 := 0; fi29 < 9; fi29++ {
        fmt.Println(fi29, fi29)
        _ = fi29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        v2 = v2
        for fi31 := 0; fi31 < 2; fi31++ {
            switch v1 {
            case -373:
                var v32 int8 = -88
                _ = v32
            case 127:
                fmt.Println(v24, fi31, v1)
            case 0:
                v3 = -313
            }
            v1 = (fi30 & v1)
            v33 := GenS3{F0: (v1 % ((v1) | 1)), F1: v2}
            _ = fi31
            _ = v33
        }
        _ = fi30
    }
    if (!false) {
        v34 := GenS3{F0: v1, F1: "foo bar"}
        switch (v1 + 208) {
        case -704:
            v3 += 32767
        }
        v3 += (^(v27))
        _ = v34
    }
    fmt.Println(v23, v23, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
