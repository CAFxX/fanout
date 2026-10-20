package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint8; F1 bool }

func main() {
    v1 := []uint32{2147483647, 259, 65535}
    var v2 uint = 133
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    if true {
        v2 = (v2 & 979)
        v2 -= v2
        var v5 int16 = -887
        _ = v5
    }
    switch -337 {
    case 42:
        if false {
            v2 = (^(v2))
        } else {
            v2 = v2
        }
    case 127:
        var v6 int32 = -410
        _ = v6
    default:
        v4[2] = (!false)
    }
    fmt.Println(v2)
    var v7 string = ("hello" + ("" + (("foo bar" + "a") + "foo bar")))
    var v8 uint = (v2 % ((v2) | 1))
    switch 9223372036854775807 {
    case 32767:
        v2 = (v8 * (v8 ^ (v2 | (v2 << ((659) & 63)))))
    case 0:
        var v9 uint16 = 42
        _ = v9
    case 1:
        v10 := GenS1{F0: 793}
        _ = v10
    default:
        if (-32768 > 42) {
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := GenS2{F0: (0.0 == ((1.5 - 3.14159) * 3.14159)), F1: 808}
                if true {
                    var v13 string = v7
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
            if true {
                if false {
                    for fi14 := 0; fi14 < 15; fi14++ {
                        v7 = v7
                        _ = fi14
                    }
                }
                fmt.Println(v8, v7)
                switch (v7 + ("hello" + "foo bar")) {
                case "":
                    v2 = 76
                case "a":
                    var v15 uint32 = 234
                    _ = v15
                }
            } else {
                v16 := GenS3{F0: 1, F1: (76 < 65)}
                if (2147483647 >= 822) {
                    switch 0 {
                    case 725:
                        v17 := v4[1]
                        _ = v17
                    default:
                        var v18 int16 = 42
                        _ = v18
                    }
                }
                _ = v16
            }
        }
    }
    if ((0.0 == 1e10) || true) {
        v7 = "hello"
        v7 = "foo bar"
        var v19 uint32 = 4294967295
        _ = v19
    } else {
        if false {
            switch "foo bar" {
            case "foo bar":
                v7 = v7
            case "":
                if (!true) {
                    v20 := v4[1]
                    for fi21 := 0; fi21 < 2; fi21++ {
                        v2 -= (v8 - (v2 & (v8 * v2)))
                        delete(v4, 2)
                        _ = fi21
                    }
                    v20 = ((1 > 4294967295) && v20)
                    _ = v20
                } else {
                    switch 4294967295 {
                    case 4294967295:
                        var v22 uint16 = 0
                        _ = v22
                    }
                    _ = len(v1)
                    v23 := GenS1{F0: 2}
                    _ = v23
                }
            case "hello":
                var v24 string = (v7 + v7)
                _ = v24
            }
        }
    }
    var v25 float64 = 3.14159
    v25 -= v25
    v2 = (^(v2))
    var v26 int64 = 0
    v26 += v26
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
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
