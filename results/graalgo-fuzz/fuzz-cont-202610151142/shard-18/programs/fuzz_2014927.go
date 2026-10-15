package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint; F2 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 string; F1 int8 }
func genFn0(p0 int64) int8 {
    var v5 uint8 = 88
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int32 = 0
        v8 := GenS1{F0: 127}
        v9 := GenS1{F0: 0}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = v5
    return int8(v5)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{("\n\t" + ("\n\t" + "a")), "foo bar", "a"}
    v3 := []string{"", ""}
    v4 := make(map[int]uint32)
    var v10 int16 = -32768
    if false {
        delete(v1, "\n\t")
        delete(v4, -2147483648)
        v10 += (v10 % ((v10) | 1))
    } else {
        delete(v1, ("" + "a"))
        var v11 int32 = 2147483647
        v11 = 1
        _ = v11
    }
    v10 += 42
    v10 = 32767
    var v12 bool = (95 > 4294967295)
    var v13 uint16 = 900
    var v14 int32 = 0
    v12 = false
    var v15 float64 = 3.14159
    fmt.Println(v12, v14, v12)
    if true {
        switch "\n\t" {
        case "foo bar":
            switch -422 {
            case -9223372036854775808:
                _ = len(v1)
            case 2147483647:
                if (!(128 != 255)) {
                    for fi16 := 0; fi16 < 5; fi16++ {
                        switch ("a" + ("a" + "foo bar")) {
                        case "hello":
                            if len(v3) > 0 {
                                v3[0] = ("hello" + (("" + ("a" + "a")) + "hello"))
                            }
                        }
                        _ = fi16
                    }
                } else {
                    switch 413 {
                    case 0:
                        var v17 int8 = -121
                        _ = v17
                    case -868:
                        v18 := GenS2{F0: false, F1: 759, F2: (v14 >> ((0) & 63))}
                        _ = v18
                    case -2147483648:
                        fmt.Println(v14, v10)
                    default:
                        v19 := GenS3{F0: 42, F1: 65535, F2: ""}
                        _ = v19
                    }
                    var v20 uint8 = 1
                    v10 = int16(v20)
                    _ = v20
                }
            default:
                for fi21 := 0; fi21 < 17; fi21++ {
                    if (v12 && false) {
                        v13 -= 932
                        var v22 int32 = (v14 % ((v14) | 1))
                        _ = len(v1)
                        _ = v22
                    }
                    _ = fi21
                }
            }
        case "":
            fmt.Println(v12)
        default:
            for fi23 := 0; fi23 < 19; fi23++ {
                v24 := GenS4{F0: 42}
                v10 = v10
                _ = fi23
                _ = v24
            }
        }
    } else {
        v25 := GenS1{F0: -1}
        var v26 uint64 = 1
        _ = v25
        _ = v26
    }
    var v27 uint16 = (v13 << ((208) & 63))
    var v28 int16 = 1000
    fmt.Println(v14, v10)
    if v12 {
        var v29 int = 1
        _ = v29
    } else {
        var v30 float64 = 0.1
        if v12 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v31 := GenS5{F0: 1, F1: 65535}
            v32 := GenS6{F0: "", F1: 42}
            switch "" {
            case "a":
                for fi33 := 0; fi33 < 18; fi33++ {
                    var v34 int16 = -90
                    v15 -= 3.14159
                    _ = fi33
                    _ = v34
                }
            case "\n\t":
                v30 -= (v30 * (v30 / v15))
            }
            _ = v31
            _ = v32
        }
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
