package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 string; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 string; F1 int16; F2 int64 }

func main() {
    v1 := []string{"hello", "hello", "foo bar"}
    v2 := []uint32{93, 2, 1, 2}
    var v3 int16 = 506
    var v4 int32 = 42
    if (0 < v3) {
        fmt.Println(v4, v4)
        fmt.Println(v4)
    } else {
        var v5 int = 2147483647
        switch 9223372036854775807 {
        case 127:
            for fi6 := 0; fi6 < 11; fi6++ {
                v7 := GenS1{F0: (!(1 > -1)), F1: 2}
                fmt.Println(fi6)
                _ = fi6
                _ = v7
            }
        }
        _ = v5
    }
    var v8 uint16 = 334
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS1{F0: true, F1: 2}
        v8 = (v8 & 967)
        _ = fi9
        _ = v10
    }
    v2 = append(v2, 677)
    for fi11 := 0; fi11 < 6; fi11++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 uint16 = 42
            if (!(127 == 0)) {
                v13 -= v8
                var v14 string = "a"
                v4 += (-(v4))
                _ = v14
            }
            fmt.Println(v4, fi12, fi11)
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    switch (("a" + "foo bar") + "\n\t") {
    case "":
        var v15 int64 = -9223372036854775808
        _ = v15
    case "hello":
        if true {
            if ((v3 / ((0) | 1)) == (v3 | v3)) {
                v8 += 32767
                v4 += (v4 >> ((169) & 63))
            }
        }
    case "foo bar":
        var v16 int16 = v3
        _ = v16
    }
    v8 = (v8 & (v8 & (^(v8))))
    v4 -= v4
    v3 = (v3 >> ((662) & 63))
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 uint = 255
        _ = fi17
        _ = v18
    }
    if true {
        for fi19 := 0; fi19 < 9; fi19++ {
            if false {
                fmt.Println(fi19)
            }
            _ = fi19
        }
        v4 = v4
        fmt.Println(v4)
    }
    if false {
        v20 := GenS2{F0: ((("foo bar" + "hello") + ("hello" + "foo bar")) + "\n\t"), F1: 1, F2: v4}
        if true {
            v21 := GenS3{F0: 65535, F1: 4294967295}
            _ = v21
        }
        var v22 int = 389
        _ = v20
        _ = v22
    } else {
        var v23 int16 = (v3 + v3)
        var v24 int64 = -219
        v23 = 1000
        _ = v23
        _ = v24
    }
    _ = len(v1)
    v3 = v3
    v8 += 748
    for fi25 := 0; fi25 < 3; fi25++ {
        for fi26 := 0; fi26 < 6; fi26++ {
            if (!false) {
                var v27 uint64 = 2
                v27 = v27
                _ = v27
            } else {
                if (!false) {
                    var v28 uint32 = 4294967295
                    v8 = v8
                    v29 := GenS2{F0: "a", F1: 18446744073709551615, F2: v4}
                    var v30 uint16 = v8
                    _ = v28
                    _ = v29
                    _ = v30
                } else {
                    var v31 uint = 2
                    v32 := GenS1{F0: (uint32(v31) >= 143), F1: 18446744073709551615}
                    v33 := GenS2{F0: ("foo bar" + "hello"), F1: 19, F2: 0}
                    _ = v31
                    _ = v32
                    _ = v33
                }
                var v34 int = int(v4)
                v34 += (v34 | (fi26 ^ v34))
                if (18446744073709551615 != 92) {
                    v34 = v34
                    _ = len(v2)
                }
                _ = v34
            }
            if (784 <= -9223372036854775808) {
                v35 := GenS3{F0: 1, F1: fi26}
                switch (^(fi25)) {
                case 0:
                    if (!(1.5 > 1.5)) {
                        v2 = append(v2, 42)
                        var v36 int64 = -443
                        switch ("hello" + "a") {
                        case "hello":
                            v4 -= 1
                        default:
                            v36 = (v36 / (((v36 ^ int64(v3))) | 1))
                        }
                        _ = v36
                    } else {
                        fmt.Println(v4, fi25, v4)
                        v37 := GenS2{F0: "a", F1: uint64(v8), F2: v4}
                        _ = v37
                    }
                case 127:
                    var v38 bool = true
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v38
                default:
                    v2 = append(v2, 4294967295)
                }
                _ = v35
            }
            switch int(v4) {
            case 218:
                v39 := GenS4{F0: "\n\t", F1: v3, F2: -558}
                _ = v39
            }
            v40 := GenS3{F0: 65535, F1: 683}
            _ = fi26
            _ = v40
        }
        _ = fi25
    }
    v41 := GenS4{F0: ("a" + "hello"), F1: (v3 * (v3 + v3)), F2: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v8
}
