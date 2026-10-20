package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 int64 = 863
    v2 := []uint32{2147483647, 4294967295, 42}
    v3 := make(map[uint32]bool)
    for fi4 := 0; fi4 < 13; fi4++ {
        if true {
            var v5 uint16 = 0
            _ = v5
        }
        _ = fi4
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        switch "hello" {
        case "a":
            if ((1.5 - (3.14159 / 1.0)) != 3.14159) {
                var v7 uint8 = 234
                v8 := GenS1{F0: (1.5 != -2.25), F1: "a", F2: fi6}
                _ = v7
                _ = v8
            }
        }
        v3[74] = true
        fmt.Println(v1, fi6)
        _ = fi6
    }
    var v9 int8 = -127
    if (false && (("\n\t" + "foo bar") <= ("" + ("a" + "a")))) {
        var v10 int32 = -644
        fmt.Println(v9, v10, v10)
        v10 += v10
        _ = v10
    }
    switch -2147483648 {
    case 128:
        var v11 uint = 153
        _ = v11
    default:
        fmt.Println(v9)
    }
    var v12 uint64 = 1
    for fi13 := 0; fi13 < 4; fi13++ {
        switch (fi13 & (fi13 << ((42) & 63))) {
        case -9223372036854775808:
            v9 -= (-(v9))
        default:
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 int = 384
                fmt.Println(v9, v12, v9)
                switch 0 {
                case 128:
                    if (((195 == 941) || (false && ((v1 / ((-850) | 1)) >= (v1 | (v1 >> ((799) & 63)))))) && false) {
                        v9 -= (-(v9))
                        switch ("foo bar" + ("\n\t" + "\n\t")) {
                        case "a":
                            var v16 uint16 = 255
                            _ = v16
                        case "\n\t":
                            if (2 >= 1) {
                                switch "" {
                                case "hello":
                                    v17 := GenS1{F0: false, F1: "\n\t", F2: fi13}
                                    fmt.Println(v9)
                                    var v18 uint64 = v12
                                    _ = v17
                                    _ = v18
                                default:
                                    switch "a" {
                                    case "a":
                                        var v19 int8 = 127
                                        _ = v19
                                    default:
                                        if false {
                                            v20 := GenS1{F0: false, F1: "foo bar", F2: fi13}
                                            var v21 int16 = -392
                                            v1 -= v1
                                            _ = v20
                                            _ = v21
                                        } else {
                                            v9 = 127
                                            var v22 int64 = (-(v1))
                                            if (false && false) {
                                                if true {
                                                    fmt.Println(v9)
                                                    var v23 int8 = v9
                                                    _ = v23
                                                } else {
                                                    var v24 uint8 = 255
                                                    _ = v24
                                                }
                                                fmt.Println(fi14, v15)
                                            }
                                            _ = v22
                                        }
                                    }
                                }
                                var v25 int64 = v1
                                _ = v25
                            }
                        case "hello":
                            v12 = v12
                        default:
                            var v26 int64 = (v1 % (((v1 * v1)) | 1))
                            _ = v26
                        }
                    }
                case -9223372036854775808:
                    switch (fi13 | fi14) {
                    case 41:
                        switch int(v9) {
                        case 127:
                            fmt.Println(fi13, v12, v1)
                        case 2147483647:
                            if (v9 > int8(v12)) {
                                v27 := GenS2{F0: 65535}
                                v15 += (^(fi13))
                                _ = v27
                            }
                        }
                    case 42:
                        v9 = v9
                    }
                }
                _ = fi14
                _ = v15
            }
        }
        _ = fi13
    }
    if true {
        v28 := GenS1{F0: ((v1 ^ (v1 >> ((18446744073709551615) & 63))) > v1), F1: (("foo bar" + ("hello" + "a")) + ""), F2: 32767}
        var v29 uint = 2
        for fi30 := 0; fi30 < 11; fi30++ {
            switch (-(fi30)) {
            case 1:
                var v31 int64 = (v1 ^ -1)
                _ = v31
            case 9223372036854775807:
                for fi32 := 0; fi32 < 6; fi32++ {
                    v29 = 0
                    v33 := GenS3{F0: (fi32 << ((v29) & 63)), F1: 88, F2: 2}
                    v34 := GenS4{F0: -860}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
            for fi35 := 0; fi35 < 6; fi35++ {
                fmt.Println(fi35, fi30, v12)
                _ = fi35
            }
            _ = fi30
        }
        _ = v28
        _ = v29
    }
    switch 2147483647 {
    case 42:
        var v36 int = 127
        _ = v36
    case 546:
        switch "\n\t" {
        case "hello":
            var v37 int16 = 0
            _ = v37
        default:
            var v38 uint16 = 893
            _ = v38
        }
    default:
        fmt.Println(v1, v12)
    }
    if false {
        for fi39 := 0; fi39 < 19; fi39++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v12 += v12
            var v40 int16 = 1
            _ = fi39
            _ = v40
        }
    } else {
        if (0 < 2) {
            v3[1] = true
        }
        var v41 int8 = 1
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
