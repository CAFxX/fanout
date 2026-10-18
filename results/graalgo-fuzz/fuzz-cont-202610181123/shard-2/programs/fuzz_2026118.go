package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int; F1 int }
func genFn0(p0 uint8) bool {
    var v5 uint32 = 2147483648
    v6 := GenS1{F0: (("" + (((("foo bar" + ((("a" + "\n\t") + "\n\t") + "\n\t")) + "") + "") + "")) + "")}
    switch -254 {
    case 1:
        if (1.5 > 0.1) {
            if false {
                v7 := GenS2{F0: -105, F1: (("" + ("foo bar" + "hello")) + "")}
                _ = v7
            }
            if false {
                fmt.Println(v5, p0)
                for fi8 := 0; fi8 < 15; fi8++ {
                    var v9 uint64 = 18446744073709551615
                    var v10 uint8 = p0
                    p0 = (p0 - v10)
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                v11 := []int{2147483647, 0}
                _ = v11
            }
        } else {
            if (2 > 0) {
                if true {
                    v5 = 1
                    switch ("a" + ("hello" + ("hello" + ("a" + "a")))) {
                    case "a":
                        var v12 int = 42
                        _ = v12
                    default:
                        fmt.Println(p0)
                    }
                    var v13 uint8 = (p0 | p0)
                    _ = v13
                }
                var v14 int16 = 32767
                fmt.Println(v14, v5, p0)
                _ = v14
            } else {
                for fi15 := 0; fi15 < 10; fi15++ {
                    v16 := GenS1{F0: (("a" + "a") + ((("hello" + "a") + "\n\t") + "a"))}
                    v17 := GenS1{F0: "a"}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                p0 = (p0 - p0)
            }
        }
    case -1:
        v18 := []int{-951, 2147483647}
        _ = v18
    }
    switch "\n\t" {
    case "foo bar":
        var v19 uint16 = 1
        _ = v19
    default:
        v20 := []int{127, 42, 9223372036854775807, 50}
        _ = v20
    }
    v21 := GenS1{F0: ("hello" + (((("hello" + "\n\t") + (("foo bar" + (("\n\t" + ("" + "foo bar")) + "")) + "foo bar")) + "") + ""))}
    _ = p0
    _ = v21
    _ = v5
    _ = v6
    return (!((p0 - 2) < (p0 + p0)))
}

func main() {
    var v1 string = ""
    v2 := make(map[int]uint32)
    v3 := []string{("hello" + v1), v1, v1, "a"}
    var v4 uint8 = 234
    fmt.Println(v1, v4)
    for fi22 := 0; fi22 < 10; fi22++ {
        v2[(fi22 * fi22)] = 281
        v4 = 128
        var v23 string = v1
        _ = fi22
        _ = v23
    }
    delete(v2, 127)
    if false {
        v24 := GenS2{F0: 127, F1: v1}
        if false {
            for fi25 := 0; fi25 < 13; fi25++ {
                switch (v1 + "\n\t") {
                case "hello":
                    fmt.Println(v4, fi25)
                case "\n\t":
                    v4 -= (v4 * v4)
                case "a":
                    var v26 uint64 = 181
                    _ = v26
                }
                var v27 float32 = float32((float64(0.0) - (1.5 + (1.5 / 1e10))))
                _ = fi25
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 7; fi28++ {
            v1 = ""
            _ = fi28
        }
        _ = v24
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        if (!(-323 < 42)) {
            switch (fi29 - (fi29 * 4294967295)) {
            case 127:
                var v30 uint8 = v4
                _ = v30
            }
            v31 := GenS3{F0: 390, F1: (v4 & v4)}
            fmt.Println(v4, v1, v4)
            _ = v31
        }
        _ = fi29
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        v2[42] = 422
        v4 += (v4 >> ((18446744073709551615) & 63))
        fmt.Println(fi32, v4, fi32)
        _ = fi32
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v33 := GenS3{F0: 138, F1: (^(v4))}
    for fi34 := 0; fi34 < 14; fi34++ {
        if false {
            v35 := GenS4{F0: 32767, F1: 844}
            switch v1 {
            case "":
                if ((true || true) && true) {
                    v2[fi34] = 4294967295
                } else {
                    fmt.Println(fi34, v4)
                }
            case "a":
                v4 += v4
            case "foo bar":
                var v36 int64 = 1
                _ = v36
            default:
                var v37 uint = 255
                _ = v37
            }
            _ = v35
        } else {
            if true {
                v38 := GenS3{F0: 595, F1: (^(v4))}
                v4 = (v4 + (v4 ^ 1))
                if false {
                    v39 := v2[(fi34 >> ((2) & 63))]
                    _ = v39
                } else {
                    v1 = v1
                }
                _ = v38
            } else {
                var v40 bool = (24 >= -67)
                switch fi34 {
                case 128:
                    v40 = false
                case -648:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = v40
            }
            var v41 int64 = 42
            var v42 uint16 = 65535
            _ = v41
            _ = v42
        }
        var v43 uint32 = 2147483647
        _ = fi34
        _ = v43
    }
    if (!true) {
        switch -493 {
        case 0:
            fmt.Println(v4, v1)
        case 9223372036854775807:
            switch 32767 {
            case -400:
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v44 := GenS5{F0: 128, F1: 587}
                    v2[-1] = 65535
                    _ = v44
                } else {
                    var v45 uint8 = v4
                    _ = len(v3)
                    v45 += (v4 << ((1) & 63))
                    _ = v45
                }
            default:
                switch (v1 + ((v1 + v1) + v1)) {
                case "hello":
                    v46 := GenS2{F0: -1, F1: ""}
                    _ = v46
                case "foo bar":
                    fmt.Println(v1, v4)
                }
            }
        }
        v4 += 42
    }
    if ((424 > 1) && true) {
        v4 = (^(v4))
        var v47 uint8 = (v4 % ((v4) | 1))
        _ = v47
    } else {
        fmt.Println(v1)
    }
    for fi48 := 0; fi48 < 7; fi48++ {
        v4 -= v4
        v4 = (v4 % ((42) | 1))
        _ = fi48
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
