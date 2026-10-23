package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint8; F2 string }
type GenS3 struct { F0 uint64; F1 int8 }

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]string)
    var v3 uint64 = 42
    v4 := []int{-2147483648}
    if false {
        var v5 float64 = 0.1
        switch -2147483648 {
        case 219:
            var v6 string = ("a" + ("\n\t" + ("foo bar" + (("foo bar" + "") + (((("" + (("\n\t" + "\n\t") + "hello")) + "a") + "") + (("\n\t" + ("a" + "")) + "a"))))))
            _ = v6
        }
        switch "foo bar" {
        case "":
            var v7 float32 = (((0.1 / 1.0) * 1e10) + (1e10 - 1e10))
            _ = v7
        }
        _ = v5
    }
    var v8 int = 42
    v9 := GenS1{F0: v8}
    if (((((1.5 * -2.25) / -2.25) * ((1e10 / 3.14159) * -2.25)) * (1.5 / 1e10)) <= 0.0) {
        var v10 int = (v8 << ((v1) & 63))
        _ = v10
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            if (175 <= 223) {
                var v12 bool = (!false)
                fmt.Println(v8)
                _ = v12
            } else {
                fmt.Println(v3, v3, v8)
                v13 := v2[1]
                _ = v13
            }
            fmt.Println(v8)
            switch -232 {
            case -9223372036854775808:
                if true {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        v2[65535] = ("\n\t" + "a")
                        v15 := GenS2{F0: 127, F1: 127, F2: "\n\t"}
                        _ = fi14
                        _ = v15
                    }
                    if ((("hello" + "hello") + "a") != "foo bar") {
                        var v16 float64 = (1e10 * (0.1 / float64(0.1)))
                        if (3 < 204) {
                            var v17 uint32 = 18
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 13; fi18++ {
                            var v19 uint64 = v3
                            fmt.Println(v19, v3, v1)
                            _ = fi18
                            _ = v19
                        }
                        _ = v16
                    } else {
                        fmt.Println(v1, v8, v3)
                    }
                }
            default:
                v1 = v1
            }
            _ = fi11
        }
        if (1e10 != 0.1) {
            v20 := GenS3{F0: 4294967295, F1: 78}
            var v21 int64 = -779
            _ = v20
            _ = v21
        } else {
            v1 = 356
        }
    }
    var v22 int = (v8 << ((4294967295) & 63))
    for fi23 := 0; fi23 < 4; fi23++ {
        switch "foo bar" {
        case "foo bar":
            v24 := GenS2{F0: 0, F1: 255, F2: ""}
            _ = v24
        default:
            if ((v8 % (((-(v8))) | 1)) <= v8) {
                var v25 string = ("a" + (((("" + ("\n\t" + ("\n\t" + "a"))) + "hello") + "foo bar") + ("\n\t" + ("\n\t" + "\n\t"))))
                var v26 uint32 = 42
                v22 -= 9223372036854775807
                _ = v25
                _ = v26
            }
        }
        if len(v4) > 0 {
            v4[0] = (fi23 & (fi23 >> ((v1) & 63)))
        }
        var v27 uint64 = 697
        _ = fi23
        _ = v27
    }
    if false {
        if (!(true && false)) {
            fmt.Println(v3, v22)
        } else {
            v3 += 0
            v1 = v1
        }
        fmt.Println(v3, v3)
        var v28 int64 = -799
        _ = v28
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if true {
            fmt.Println(v3, v1)
        }
        switch -9223372036854775808 {
        case 32768:
            if (false && (42 < 952)) {
                v8 -= v8
            } else {
                var v29 uint = v1
                fmt.Println(v1)
                var v30 int32 = 1
                _ = v29
                _ = v30
            }
        case 873:
            v31 := GenS1{F0: (v22 - (v8 | v8))}
            _ = v31
        default:
            var v32 int8 = -1
            _ = v32
        }
    }
    _ = len(v2)
    fmt.Println(v3)
    v33 := v2[4294967295]
    fmt.Println(v33, v33)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v4
    _ = v8
    _ = v9
}
