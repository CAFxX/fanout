package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 bool }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int16; F1 uint32; F2 int }
type GenS4 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 int32 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    var v4 uint8 = 42
    fmt.Println(v1, v1)
    v4 = (v4 - (v4 * (v4 % ((v4) | 1))))
    if ("a" <= "hello") {
        var v5 int64 = (v1 >> ((65535) & 63))
        var v6 int8 = 42
        v7 := GenS1{F0: ((v5 < v1) && false), F1: 4294967295, F2: true}
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 int16 = -188
    for fi9 := 0; fi9 < 20; fi9++ {
        v8 -= 32767
        var v10 float64 = 1e10
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 0, F1: true}
    v8 -= 1
    switch 127 {
    case 4294967295:
        fmt.Println(v8, v4)
    case 42:
        var v12 int32 = -1
        _ = v12
    }
    if true {
        for fi13 := 0; fi13 < 7; fi13++ {
            if false {
                if (true && true) {
                    v14 := GenS3{F0: (v8 & v8), F1: 2, F2: (fi13 / (((^(fi13))) | 1))}
                    _ = v14
                }
                _ = len(v3)
                v1 = -9223372036854775808
            } else {
                for fi15 := 0; fi15 < 4; fi15++ {
                    v16 := []int{-2147483648, int(v1)}
                    switch "" {
                    case "foo bar":
                        if len(v16) > 0 {
                            v16[0] = fi13
                        }
                    default:
                        var v17 uint32 = 2147483647
                        _ = v17
                    }
                    fmt.Println(v4)
                    _ = fi15
                    _ = v16
                }
                if (0.1 <= (-2.25 + 0.0)) {
                    var v18 uint8 = 128
                    var v19 uint32 = 4294967295
                    fmt.Println(v18, v18)
                    _ = v18
                    _ = v19
                }
            }
            v3[("foo bar" + "")] = "\n\t"
            v20 := GenS2{F0: 366, F1: true}
            _ = fi13
            _ = v20
        }
        var v21 string = ("hello" + ("a" + "a"))
        _ = v21
    }
    switch 32768 {
    case 42:
        v8 = 1
    default:
        v4 = (v4 & v4)
    }
    var v22 int64 = v1
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 uint64 = 2
        _ = fi23
        _ = v24
    }
    fmt.Println(v1, v22)
    switch 128 {
    case -1:
        if true {
            _ = len(v2)
            if ((v8 * 0) < (v8 & v8)) {
                for fi25 := 0; fi25 < 19; fi25++ {
                    var v26 int = (fi25 << ((42) & 63))
                    for fi27 := 0; fi27 < 4; fi27++ {
                        v22 = -980
                        v4 = (v4 >> ((0) & 63))
                        _ = fi27
                    }
                    var v28 float64 = 0.0
                    _ = fi25
                    _ = v26
                    _ = v28
                }
                if (!(v1 < -1)) {
                    var v29 uint64 = 42
                    v30 := v2["foo bar"]
                    if (false || true) {
                        v29 = (^(v29))
                        v29 += 387
                        var v31 int = -692
                        _ = v31
                    } else {
                        switch "hello" {
                        case "":
                            v32 := GenS4{F0: (v30 % ((v30) | 1)), F1: 42, F2: (v8 - (^(v8)))}
                            _ = v32
                        }
                    }
                    _ = v29
                    _ = v30
                }
            }
            for fi33 := 0; fi33 < 2; fi33++ {
                v34 := []int{fi33, (fi33 & (fi33 & (^(fi33)))), fi33}
                v35 := GenS5{F0: ("foo bar" + "foo bar"), F1: 32767}
                _ = fi33
                _ = v34
                _ = v35
            }
        }
    case 42:
        for fi36 := 0; fi36 < 10; fi36++ {
            if true {
                var v37 uint = 1
                _ = len(v3)
                _ = v37
            }
            var v38 uint = 2
            _ = fi36
            _ = v38
        }
    case 1:
        fmt.Println(v1, v22, v4)
    }
    if false {
        if true {
            v39 := GenS6{F0: 1}
            _ = v39
        } else {
            var v40 bool = false
            _ = v40
        }
        v41 := GenS3{F0: 1, F1: 121, F2: -275}
        _ = v41
    }
    var v42 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v8
}
