package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int64; F1 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 357
    var v3 int8 = 127
    v4 := GenS1{F0: 175, F1: v2, F2: (v3 ^ (v3 << ((1000) & 63)))}
    for fi5 := 0; fi5 < 8; fi5++ {
        if false {
            fmt.Println(v3)
            _ = len(v1)
        }
        if true {
            var v6 int64 = 9223372036854775807
            _ = v6
        }
        switch (("\n\t" + "") + ("\n\t" + "hello")) {
        case "foo bar":
            var v7 uint32 = 42
            _ = v7
        default:
            fmt.Println(fi5, fi5)
        }
        _ = fi5
    }
    switch "" {
    case "a":
        var v8 int = 32767
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 6; fi9++ {
            v3 = (-(v3))
            _ = fi9
        }
    }
    fmt.Println(v2)
    fmt.Println(v3, v2)
    var v10 uint64 = (^(v2))
    var v11 int32 = -2147483648
    switch ("a" + "a") {
    case "foo bar":
        if (681 <= -1) {
            if ("\n\t" < "\n\t") {
                if (709 > 255) {
                    if false {
                        v2 -= 1
                        v12 := []int{1, 128, -9223372036854775808}
                        for fi13 := 0; fi13 < 17; fi13++ {
                            v14 := GenS2{F0: false}
                            var v15 int = fi13
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        _ = v12
                    }
                    if (2 != 65535) {
                        var v16 uint16 = 2
                        v17 := GenS3{F0: 42, F1: 360, F2: 255}
                        var v18 string = (("\n\t" + "hello") + "a")
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    var v19 uint16 = 65535
                    _ = v19
                }
                v10 -= (^(v2))
                v20 := []int{-441}
                _ = v20
            } else {
                v21 := GenS2{F0: ((-32768 <= 42) && false)}
                fmt.Println(v2, v3)
                _ = v21
            }
            var v22 uint = 65535
            _ = v22
        }
    case "":
        if false {
            delete(v1, 1)
            var v23 uint32 = 402
            _ = v23
        } else {
            delete(v1, -9223372036854775808)
            var v24 int = 1
            v25 := GenS1{F0: 42, F1: (v10 >> ((88) & 63)), F2: v3}
            _ = v24
            _ = v25
        }
    case "\n\t":
        fmt.Println(v11)
    }
    if (v3 == v3) {
        for fi26 := 0; fi26 < 3; fi26++ {
            var v27 int32 = (v11 * (v11 << ((42) & 63)))
            fmt.Println(v3, v10, fi26)
            _ = fi26
            _ = v27
        }
        var v28 int64 = 36
        _ = v28
    } else {
        var v29 int8 = (v3 ^ (-(v3)))
        var v30 string = "hello"
        fmt.Println(v29)
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        for fi32 := 0; fi32 < 13; fi32++ {
            v10 = v10
            v11 += (v11 & (v11 * 2147483647))
            _ = fi32
        }
        fmt.Println(fi31, v2, v11)
        _ = fi31
    }
    v3 += -20
    switch "a" {
    case "":
        v33 := []int{127}
        _ = v33
    case "foo bar":
        if ((v3 | 1) <= 1) {
            v34 := GenS2{F0: (!(-1 < 1000))}
            switch "" {
            case "foo bar":
                v3 += v3
            }
            _ = v34
        } else {
            for fi35 := 0; fi35 < 19; fi35++ {
                v36 := GenS2{F0: false}
                for fi37 := 0; fi37 < 19; fi37++ {
                    if false {
                        _ = len(v1)
                        v3 = -128
                        fmt.Println(fi37, v2)
                    }
                    _ = fi37
                }
                if false {
                    for fi38 := 0; fi38 < 4; fi38++ {
                        v39 := []int{fi38, fi38, 1, (fi35 << ((42) & 63))}
                        v10 = (^(v2))
                        v40 := GenS4{F0: 1}
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    _ = len(v1)
                    var v41 uint32 = 2147483647
                    _ = v41
                } else {
                    v2 = (v10 % ((v10) | 1))
                    var v42 string = ""
                    _ = v42
                }
                _ = fi35
                _ = v36
            }
        }
    }
    switch "" {
    case "":
        for fi43 := 0; fi43 < 6; fi43++ {
            if (((v2 - v2) <= 110) || true) {
                v11 = -1
            }
            fmt.Println(v11)
            for fi44 := 0; fi44 < 15; fi44++ {
                if false {
                    if (0 > 42) {
                        v45 := GenS4{F0: (v3 % ((v3) | 1))}
                        fmt.Println(v11, v11, v2)
                        _ = v45
                    } else {
                        v46 := []int{fi43, (fi44 + fi44)}
                        switch (fi43 ^ fi43) {
                        case 32767:
                            v47 := GenS5{F0: 2147483647, F1: 0}
                            _ = v47
                        case -677:
                            v46 = append(v46, (fi44 ^ fi43))
                        case 9223372036854775807:
                            v48 := GenS6{F0: 9223372036854775807, F1: 42}
                            _ = v48
                        default:
                            var v49 int32 = v11
                            _ = v49
                        }
                        _ = v46
                    }
                    var v50 int16 = -32768
                    fmt.Println(v10, v3)
                    fmt.Println(fi43, v50)
                    switch ((("\n\t" + "foo bar") + "hello") + "") {
                    case "foo bar":
                        delete(v1, -237)
                    case "":
                        var v51 uint16 = 65535
                        _ = v51
                    default:
                        v3 = -32
                    }
                    _ = v50
                }
                v11 -= (-(v11))
                _ = fi44
            }
            _ = fi43
        }
    case "\n\t":
        fmt.Println(v10, v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
