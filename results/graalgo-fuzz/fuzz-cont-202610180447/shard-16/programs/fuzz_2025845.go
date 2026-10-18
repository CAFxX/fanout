package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []string{"\n\t", ""}
    v2 := make(map[string]uint32)
    var v3 string = (("a" + ((("foo bar" + "hello") + "foo bar") + "foo bar")) + "foo bar")
    v4 := []uint32{208, 340}
    v3 = v3
    if true {
        v3 = v3
        v5 := GenS1{F0: -9223372036854775808, F1: "foo bar", F2: 2}
        var v6 float32 = 0.0
        _ = v5
        _ = v6
    }
    var v7 int8 = 0
    switch -1 {
    case 2147483647:
        _ = len(v1)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (42 > -1) {
        var v8 string = "\n\t"
        var v9 bool = (32768 == 2147483647)
        _ = len(v2)
        _ = v8
        _ = v9
    } else {
        if false {
            fmt.Println(v3)
            v3 = "a"
        } else {
            var v10 uint32 = 812
            _ = v10
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 int64 = -813
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 string = v3
                _ = fi13
                _ = v14
            }
            var v15 int64 = v12
            _ = fi11
            _ = v12
            _ = v15
        }
        var v16 int32 = -1
        _ = v16
    }
    switch v3 {
    case "a":
        v3 = v3
    default:
        var v17 int64 = 1
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 uint64 = 118
        for fi20 := 0; fi20 < 13; fi20++ {
            if false {
                v3 = ("a" + v3)
                delete(v2, "")
                if false {
                    var v21 uint = 118
                    v22 := GenS2{F0: (v7 % (((v7 / ((int8(v19)) | 1))) | 1))}
                    _ = v21
                    _ = v22
                } else {
                    var v23 string = ("foo bar" + ("foo bar" + v3))
                    v19 += 18446744073709551615
                    v24 := GenS1{F0: 813, F1: v3, F2: 991}
                    _ = v23
                    _ = v24
                }
            } else {
                switch v3 {
                case "foo bar":
                    if (!false) {
                        var v25 uint64 = (^(v19))
                        fmt.Println(fi18, v19, v3)
                        fmt.Println(fi20)
                        _ = v25
                    }
                default:
                    v26 := GenS1{F0: 1, F1: v3, F2: 539}
                    _ = v26
                }
                v7 -= v7
                if (!(true || false)) {
                    var v27 uint64 = 2
                    v19 -= (v27 + (^(v19)))
                    var v28 uint = uint(v27)
                    _ = v27
                    _ = v28
                } else {
                    if len(v4) > 0 {
                        v4[0] = 2147483648
                    }
                    v7 = int8(v19)
                    _ = len(v2)
                }
            }
            _ = fi20
        }
        switch "" {
        case "":
            var v29 int16 = 17
            _ = v29
        case "\n\t":
            v30 := GenS3{F0: 42}
            _ = v30
        }
        _ = fi18
        _ = v19
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch v3 {
    case "foo bar":
        for fi31 := 0; fi31 < 7; fi31++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v32 string = "a"
            var v33 int = (fi31 << ((358) & 63))
            _ = fi31
            _ = v32
            _ = v33
        }
    default:
        if len(v4) > 0 {
            v4[0] = 1
        }
    }
    fmt.Println(v3, v3)
    v3 = v3
    for fi34 := 0; fi34 < 11; fi34++ {
        v4 = append(v4, 65535)
        var v35 string = ("foo bar" + "hello")
        _ = fi34
        _ = v35
    }
    v36 := GenS2{F0: 0}
    switch -2147483648 {
    case 925:
        if (283 > 255) {
            if true {
                v7 = (^(v7))
                switch 176 {
                case 803:
                    v7 = -128
                default:
                    if false {
                        var v37 uint8 = 127
                        var v38 uint = 4294967295
                        _ = v37
                        _ = v38
                    }
                }
            } else {
                if (2147483647 < -2147483648) {
                    if len(v1) > 0 {
                        v1[0] = v3
                    }
                    fmt.Println(v3)
                    if false {
                        for fi39 := 0; fi39 < 16; fi39++ {
                            var v40 float32 = 0.1
                            _ = fi39
                            _ = v40
                        }
                        v41 := v2[v3]
                        if (1 == 4294967295) {
                            var v42 int32 = -2147483648
                            _ = v42
                        }
                        _ = v41
                    } else {
                        v7 += v7
                    }
                } else {
                    _ = len(v2)
                }
                var v43 uint64 = 0
                fmt.Println(v7)
                _ = v43
            }
            fmt.Println(v3, v7)
            var v44 uint8 = 128
            _ = v44
        }
    case 879:
        v7 = (-(v7))
    }
    if len(v1) > 0 {
        v1[0] = ("a" + v3)
    }
    v7 = (-(v7))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v7
}
