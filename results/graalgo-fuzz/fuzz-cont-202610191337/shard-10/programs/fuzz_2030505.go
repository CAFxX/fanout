package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int64; F2 bool }
type GenS3 struct { F0 int16 }
func genFn0() int8 {
    var v5 uint16 = 219
    var v6 bool = (316 > 2147483647)
    v5 -= v5
    _ = v5
    _ = v6
    return 76
}

func main() {
    var v1 float64 = -2.25
    var v2 uint16 = 2
    var v3 int64 = 659
    v4 := []string{"hello", "foo bar", "foo bar"}
    var v7 int32 = 1
    var v8 bool = (!true)
    switch "a" {
    case "foo bar":
        switch 990 {
        case 128:
            v3 = 0
        }
    case "a":
        v3 = (v3 & (v3 * 898))
    }
    fmt.Println(v3)
    switch 9223372036854775807 {
    case 128:
        v4 = append(v4, "foo bar")
    case -165:
        if (!(4294967295 < 0)) {
            var v9 int32 = (v7 << ((57) & 63))
            var v10 int = 625
            if v8 {
                var v11 bool = false
                v4 = append(v4, ("\n\t" + ("foo bar" + "foo bar")))
                _ = v11
            } else {
                fmt.Println(v10, v2, v8)
                v12 := GenS1{F0: 0}
                v10 = v10
                _ = v12
            }
            _ = v10
            _ = v9
        }
    }
    v3 -= (-(v3))
    switch "a" {
    case "a":
        var v13 int16 = 32767
        _ = v13
    case "\n\t":
        var v14 float64 = (v1 + 0.0)
        _ = v14
    default:
        if (false || (!(((0.0 * (1.5 / (((3.14159 * -2.25) / 1e10) / 3.14159))) - 0.1) != 1e10))) {
            for fi15 := 0; fi15 < 4; fi15++ {
                if (!("\n\t" == "\n\t")) {
                    if (v3 > -1) {
                        var v16 int32 = v7
                        fmt.Println(v7, fi15, v2)
                        var v17 int32 = v7
                        _ = v16
                        _ = v17
                    } else {
                        fmt.Println(v7, fi15, v8)
                        var v18 uint64 = uint64(v2)
                        v2 -= 1
                        _ = v18
                    }
                }
                var v19 float32 = (1.5 * 0.1)
                _ = fi15
                _ = v19
            }
            if v8 {
                v3 = (-(v3))
                var v20 bool = (float32(v1) <= float32(1.5))
                _ = v20
            }
        } else {
            v21 := GenS1{F0: 42}
            if v8 {
                v7 = (v7 | (v7 - v7))
                switch "" {
                case "foo bar":
                    if len(v4) > 0 {
                        v4[0] = ("hello" + "")
                    }
                default:
                    var v22 int8 = 42
                    _ = v22
                }
            } else {
                var v23 uint64 = 18446744073709551615
                v24 := GenS1{F0: (v7 ^ 541)}
                _ = v23
                _ = v24
            }
            v1 = (((v1 / v1) - (v1 + 0.0)) * ((v1 / v1) + (0.0 - float64((-2.25 + 0.1)))))
            _ = v21
        }
    }
    v25 := GenS2{F0: 127, F1: -55, F2: (v8 || true)}
    v7 += (-(v7))
    if len(v4) > 0 {
        v4[0] = (("a" + "a") + ("" + ("\n\t" + "foo bar")))
    }
    var v26 uint64 = 42
    v3 += -907
    _ = len(v4)
    for fi27 := 0; fi27 < 17; fi27++ {
        v2 = v2
        for fi28 := 0; fi28 < 11; fi28++ {
            v29 := GenS3{F0: 42}
            _ = fi28
            _ = v29
        }
        switch (-(fi27)) {
        case -9223372036854775808:
            if (!(2147483647 <= 65535)) {
                v30 := GenS1{F0: (v7 % (((v7 ^ -1)) | 1))}
                v31 := GenS3{F0: int16(v26)}
                if ((!v8) && ((("" + ("\n\t" + "\n\t")) + "") == "hello")) {
                    var v32 uint8 = 234
                    fmt.Println(v3, v3, v3)
                    _ = v32
                } else {
                    fmt.Println(fi27, v3, v1)
                }
                _ = v30
                _ = v31
            } else {
                for fi33 := 0; fi33 < 1; fi33++ {
                    switch (("foo bar" + "hello") + "\n\t") {
                    case "a":
                        var v34 float64 = v1
                        _ = v34
                    case "hello":
                        fmt.Println(fi27)
                    }
                    fmt.Println(v3, v26)
                    if v8 {
                        v8 = (42 < 32767)
                    }
                    _ = fi33
                }
            }
        case 673:
            v4 = append(v4, "foo bar")
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
