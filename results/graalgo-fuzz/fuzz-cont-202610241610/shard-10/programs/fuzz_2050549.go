package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int32; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 int16; F1 int; F2 uint }
func genFn0(p0 string, p1 int64) string {
    var v4 float32 = -2.25
    switch p0 {
    case "a":
        if (-809 == 4294967295) {
            switch (("a" + "hello") + ((p0 + "hello") + p0)) {
            case "hello":
                v4 = (v4 / v4)
            case "a":
                fmt.Println(p0)
            case "foo bar":
                p1 -= p1
            default:
                var v5 int32 = 637
                _ = v5
            }
        }
    default:
        v6 := GenS1{F0: 17}
        _ = v6
    }
    p1 = (^(p1))
    v7 := GenS1{F0: 2}
    fmt.Println(v4, v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return p0
}

func main() {
    var v1 float32 = 3.14159
    v2 := []uint32{1}
    var v3 uint32 = 2147483647
    var v8 uint = uint(v3)
    v9 := GenS1{F0: 18446744073709551615}
    v3 -= 65535
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: 2}
        v1 += float32(1e10)
        _ = fi10
        _ = v11
    }
    fmt.Println(v3, v3, v8)
    switch ("\n\t" + "") {
    case "a":
        var v12 int = 2147483647
        _ = v12
    default:
        fmt.Println(v3, v1, v3)
    }
    if (!true) {
        var v13 float64 = (0.1 * (0.0 * ((float64(((v1 + (3.14159 + ((v1 - (v1 + (1.5 / (float32(3.14159) / float32(0.1))))) - ((v1 - 1e10) - v1)))) * 0.1)) * (1e10 + float64(float32(1.5)))) / 1.0)))
        if (127 < 127) {
            v3 = (v3 << (((v8 % ((v8) | 1))) & 63))
            v14 := GenS2{F0: 2147483647}
            v15 := GenS3{F0: (v3 ^ 0), F1: -1, F2: -1}
            _ = v14
            _ = v15
        } else {
            switch 32768 {
            case 4294967295:
                if ("" != "foo bar") {
                    if false {
                        switch (("hello" + (("foo bar" + ("foo bar" + "hello")) + ("foo bar" + ("hello" + ("hello" + ("" + "hello")))))) + "\n\t") {
                        case "\n\t":
                            v8 = v8
                        }
                    } else {
                        v13 = v13
                        var v16 int64 = -9223372036854775808
                        for fi17 := 0; fi17 < 15; fi17++ {
                            if true {
                                v16 -= 42
                                v18 := GenS4{F0: fi17}
                                for fi19 := 0; fi19 < 3; fi19++ {
                                    var v20 bool = false
                                    _ = fi19
                                    _ = v20
                                }
                                _ = v18
                            }
                            for fi21 := 0; fi21 < 1; fi21++ {
                                v22 := GenS5{F0: (fi21 | 733), F1: -61}
                                v23 := GenS3{F0: (v3 / ((v3) | 1)), F1: 2147483647, F2: v16}
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                            var v24 int16 = 0
                            _ = fi17
                            _ = v24
                        }
                        _ = v16
                    }
                    for fi25 := 0; fi25 < 19; fi25++ {
                        var v26 bool = false
                        v27 := GenS6{F0: 32767, F1: fi25, F2: v8}
                        var v28 float32 = 0.1
                        _ = fi25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                }
            case 32768:
                var v29 uint = (v8 % ((v8) | 1))
                _ = v29
            }
            var v30 int64 = 700
            _ = v30
        }
        if (false || false) {
            v2 = append(v2, v3)
            for fi31 := 0; fi31 < 15; fi31++ {
                v3 = (^(v3))
                var v32 int32 = 923
                _ = fi31
                _ = v32
            }
        }
        _ = v13
    } else {
        v2 = append(v2, 294)
        if false {
            v2 = append(v2, (v3 - v3))
            v33 := GenS1{F0: 1}
            _ = v33
        } else {
            v34 := GenS4{F0: 871}
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        fmt.Println(v8, v1, v8)
        v36 := GenS3{F0: v3, F1: 247, F2: -504}
        var v37 bool = true
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 int16 = 114
    v8 -= (v8 % ((v8) | 1))
    var v39 uint8 = 2
    for fi40 := 0; fi40 < 5; fi40++ {
        var v41 int = (fi40 >> (((v8 | v8)) & 63))
        if true {
            v3 -= 751
            if false {
                var v42 int8 = -128
                _ = v42
            }
            var v43 uint = uint(v3)
            _ = v43
        } else {
            fmt.Println(v3)
        }
        _ = fi40
        _ = v41
    }
    v1 -= ((v1 / v1) + (v1 / ((1.5 - ((float32((-2.25 * 0.1)) - float32(((0.0 - -2.25) / (3.14159 + 3.14159)))) - (1.5 + v1))) * (float32(1e10) - (v1 * (1.5 / 1.0))))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v8
    _ = v9
}
