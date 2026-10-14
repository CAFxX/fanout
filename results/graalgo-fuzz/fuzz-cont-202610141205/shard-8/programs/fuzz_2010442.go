package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 int; F1 uint8; F2 bool }
func genFn0(p0 bool) string {
    v7 := GenS1{F0: 2147483648}
    p0 = (100 >= 42)
    p0 = p0
    p0 = false
    _ = p0
    _ = v7
    return "foo bar"
}
func genFn1() uint {
    if false {
        var v8 uint8 = 196
        _ = v8
    }
    switch "a" {
    case "foo bar":
        v9 := []int{}
        _ = v9
    default:
        var v10 int32 = -2147483648
        _ = v10
    }
    v11 := []int{-9223372036854775808, 42}
    v12 := GenS2{F0: 42, F1: 32767, F2: 2}
    var v13 uint8 = 127
    _ = v11
    _ = v12
    _ = v13
    return 2
}

func main() {
    var v1 uint32 = 0
    var v2 uint64 = uint64(v1)
    var v3 uint64 = (v2 >> ((2) & 63))
    v4 := []string{"foo bar", "hello", "a", "\n\t"}
    v5 := make(map[int]string)
    v6 := []int{32768, 531, 1}
    v3 = (^(v2))
    v2 += (v2 << ((1) & 63))
    _ = len(v5)
    if false {
        v3 = v3
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := GenS2{F0: 1, F1: 217, F2: 2}
        _ = fi14
        _ = v15
    }
    if (false && (1 != -905)) {
        switch 2147483647 {
        case 750:
            var v16 int = 1
            _ = v16
        }
        switch "a" {
        case "\n\t":
            v17 := GenS1{F0: 2}
            _ = v17
        case "a":
            var v18 int = 286
            _ = v18
        case "foo bar":
            var v19 int32 = 384
            _ = v19
        default:
            v2 = v3
        }
        v2 = (v3 % (((v3 - v2)) | 1))
    }
    fmt.Println(v1)
    if len(v4) > 0 {
        v4[0] = (("hello" + ("foo bar" + "foo bar")) + ((("hello" + "\n\t") + ("hello" + ("a" + "hello"))) + "\n\t"))
    }
    switch 42 {
    case -1:
        switch -730 {
        case 9223372036854775807:
            if (-128 <= int8(v1)) {
                switch -900 {
                case 0:
                    for fi20 := 0; fi20 < 10; fi20++ {
                        fmt.Println(v1, v3)
                        v21 := GenS1{F0: v1}
                        _ = fi20
                        _ = v21
                    }
                }
                switch 1 {
                case 9223372036854775807:
                    if (v1 >= 0) {
                        v2 = v3
                    }
                case 128:
                    v2 = 2
                case -236:
                    _ = len(v5)
                }
                switch "" {
                case "":
                    v3 = 18446744073709551615
                case "\n\t":
                    for fi22 := 0; fi22 < 11; fi22++ {
                        v3 += v3
                        _ = fi22
                    }
                case "hello":
                    _ = len(v6)
                }
            } else {
                switch 2147483647 {
                case 2147483647:
                    var v23 int = 32768
                    _ = v23
                default:
                    var v24 int8 = 25
                    _ = v24
                }
                for fi25 := 0; fi25 < 18; fi25++ {
                    v3 += v2
                    _ = fi25
                }
            }
        case 240:
            v1 = (v1 + 42)
        case -136:
            v26 := GenS2{F0: -9223372036854775808, F1: 42, F2: 519}
            _ = v26
        }
    case 1:
        var v27 int64 = -1
        _ = v27
    case 0:
        if false {
            for fi28 := 0; fi28 < 13; fi28++ {
                v29 := GenS3{F0: 42, F1: uint8(v3)}
                v3 += v3
                _ = fi28
                _ = v29
            }
            v30 := GenS1{F0: v1}
            fmt.Println(v3)
            _ = v30
        } else {
            v2 = 765
            var v31 int16 = 1000
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        for fi33 := 0; fi33 < 19; fi33++ {
            switch 32767 {
            case 1:
                v34 := GenS3{F0: 651, F1: 23}
                _ = v34
            case -9223372036854775808:
                v6 = append(v6, fi33)
            default:
                var v35 uint8 = 255
                _ = v35
            }
            var v36 int = fi33
            v4 = append(v4, "hello")
            _ = fi33
            _ = v36
        }
        _ = fi32
    }
    v3 -= v2
    v37 := GenS1{F0: uint32(v2)}
    var v38 int16 = -185
    switch 9223372036854775807 {
    case -41:
        fmt.Println(v1, v38)
    case 1:
        v39 := v5[1]
        _ = v39
    default:
        switch "hello" {
        case "a":
            var v40 int = 128
            _ = v40
        case "":
            v41 := GenS4{F0: (v2 / ((v2) | 1)), F1: 1}
            _ = v41
        case "hello":
            v42 := GenS5{F0: 0, F1: 76, F2: true}
            _ = v42
        }
    }
    fmt.Println(v2, v2)
    var v43 float32 = 0.0
    var v44 uint8 = 77
    for fi45 := 0; fi45 < 7; fi45++ {
        if true {
            if true {
                v1 = v1
                var v46 int16 = (v38 & v38)
                _ = v46
            }
        } else {
            if (65535 == 42) {
                if len(v6) > 0 {
                    v6[0] = (fi45 << ((uint(v3)) & 63))
                }
                var v47 uint64 = (v2 ^ 0)
                v47 -= v47
                _ = v47
            } else {
                var v48 uint16 = 42
                _ = v48
            }
        }
        fmt.Println(v44, v3, v44)
        v38 += v38
        _ = fi45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
