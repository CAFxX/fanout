package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint; F2 uint16 }
type GenS3 struct { F0 string; F1 int32; F2 uint16 }

func main() {
    var v1 int = 127
    var v2 uint32 = 42
    v3 := make(map[uint32]string)
    var v4 int32 = 1
    v5 := make(map[string]string)
    v6 := make(map[uint32]string)
    var v7 int64 = int64(v1)
    if false {
        v8 := []int{-1}
        _ = v8
    } else {
        fmt.Println(v7, v2)
        if true {
            if (0.1 == (1e10 * 1e10)) {
                for fi9 := 0; fi9 < 8; fi9++ {
                    switch fi9 {
                    case 996:
                        v10 := GenS1{F0: (false && true)}
                        _ = v10
                    case 32768:
                        v2 = v2
                    case -9223372036854775808:
                        var v11 float32 = ((0.0 * 0.0) / float32(((0.1 + 0.0) + 1.5)))
                        _ = v11
                    }
                    v2 += v2
                    var v12 bool = false
                    _ = fi9
                    _ = v12
                }
            }
        } else {
            var v13 int16 = -1
            _ = v13
        }
    }
    switch "foo bar" {
    case "":
        v7 = (v7 >> ((1) & 63))
    }
    if ((false || false) || (127 <= int(v7))) {
        var v14 string = "\n\t"
        var v15 uint16 = uint16(v2)
        var v16 uint16 = v15
        _ = v14
        _ = v15
        _ = v16
    } else {
        if (413 < 1) {
            fmt.Println(v4, v4, v7)
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := GenS2{F0: 32767, F1: 360, F2: 862}
                v19 := GenS3{F0: ("\n\t" + (((("foo bar" + ("foo bar" + ("\n\t" + "hello"))) + ("\n\t" + "foo bar")) + "") + "\n\t")), F1: int32(fi17), F2: 65535}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 1; fi20++ {
            _ = len(v5)
            var v21 string = "foo bar"
            var v22 uint64 = 772
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    switch -1 {
    case -1:
        v1 = (^(v1))
    case 2147483647:
        if false {
            var v23 int64 = (v7 % ((int64(v4)) | 1))
            v24 := v5["hello"]
            v1 = int(v4)
            _ = v23
            _ = v24
        } else {
            switch ("hello" + "") {
            case "hello":
                for fi25 := 0; fi25 < 3; fi25++ {
                    if true {
                        var v26 float64 = 1.5
                        _ = v26
                    } else {
                        v7 += v7
                    }
                    _ = fi25
                }
            case "\n\t":
                for fi27 := 0; fi27 < 9; fi27++ {
                    fmt.Println(v4)
                    _ = fi27
                }
            default:
                switch (v1 | v1) {
                case -9223372036854775808:
                    var v28 int32 = (-(v4))
                    _ = v28
                case -741:
                    switch v1 {
                    case 128:
                        if false {
                            v29 := []int{-424}
                            var v30 float32 = 0.0
                            _ = v29
                            _ = v30
                        }
                    default:
                        if false {
                            for fi31 := 0; fi31 < 17; fi31++ {
                                v2 = (^(v2))
                                _ = fi31
                            }
                        }
                    }
                case 0:
                    for fi32 := 0; fi32 < 20; fi32++ {
                        for fi33 := 0; fi33 < 4; fi33++ {
                            v4 = int32(fi32)
                            var v34 int = (v1 ^ v1)
                            var v35 string = ("hello" + "a")
                            var v36 float32 = float32((float64(1e10) / -2.25))
                            _ = fi33
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        _ = fi32
                    }
                default:
                    v4 -= v4
                }
            }
        }
    case 4294967295:
        if false {
            fmt.Println(v1, v2, v4)
            var v37 int64 = 9223372036854775807
            _ = v37
        }
    }
    var v38 int16 = 1
    switch (-(v1)) {
    case 841:
        var v39 uint32 = (v2 >> ((216) & 63))
        _ = v39
    case 127:
        delete(v6, 1)
    default:
        for fi40 := 0; fi40 < 8; fi40++ {
            switch "a" {
            case "foo bar":
                v41 := []int{v1, (^(fi40)), -9223372036854775808, (v1 / ((fi40) | 1))}
                _ = v41
            }
            v1 -= v1
            _ = fi40
        }
    }
    fmt.Println(v4, v38, v38)
    switch ("\n\t" + "hello") {
    case "":
        v1 = (v1 & (v1 / (((-(v1))) | 1)))
    }
    v7 += v7
    delete(v5, "")
    if (!(!(-2.25 <= 0.1))) {
        v1 += v1
        var v42 bool = true
        _ = v42
    } else {
        v43 := v3[1]
        _ = v43
    }
    fmt.Println(v2)
    v7 -= (-(v7))
    v44 := GenS1{F0: true}
    for fi45 := 0; fi45 < 10; fi45++ {
        _ = len(v6)
        v38 = v38
        _ = fi45
    }
    fmt.Println(v2, v1)
    var v46 float32 = (3.14159 / float32(-2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v44
    _ = v46
    _ = v5
    _ = v6
    _ = v7
}
