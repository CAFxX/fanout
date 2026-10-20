package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint16 }

func main() {
    var v1 uint16 = 944
    var v2 int32 = -1
    v3 := make(map[string]bool)
    var v4 int64 = 1
    v5 := []uint32{42, 65535}
    v6 := []uint32{2147483647, 812, 4294967295}
    if false {
        for fi7 := 0; fi7 < 5; fi7++ {
            v4 = v4
            _ = fi7
        }
        var v8 int8 = -128
        _ = v8
    }
    fmt.Println(v1, v1)
    if false {
        fmt.Println(v1)
        switch "foo bar" {
        case "":
            if false {
                v9 := GenS1{F0: -63, F1: 42}
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v4)
                }
                _ = v9
            } else {
                if (v4 < 428) {
                    v10 := GenS2{F0: false}
                    v1 = 32767
                    var v11 int64 = 479
                    _ = v10
                    _ = v11
                } else {
                    v3["foo bar"] = (!(1 != v1))
                    v1 -= 321
                }
            }
        case "foo bar":
            _ = len(v3)
        case "hello":
            v12 := GenS3{F0: 42}
            _ = v12
        default:
            delete(v3, ("hello" + "a"))
        }
        v13 := GenS4{F0: 153, F1: (v4 | (v4 * 848)), F2: (v1 - v1)}
        _ = v13
    }
    if ((-266 <= v2) || false) {
        v14 := v3["hello"]
        delete(v3, (((("hello" + "\n\t") + "\n\t") + "\n\t") + ("hello" + "a")))
        _ = v14
    } else {
        if (0.0 >= 3.14159) {
            if (false && true) {
                if false {
                    v5 = append(v5, 0)
                    v1 = 543
                    var v15 int8 = 93
                    _ = v15
                }
                if len(v6) > 0 {
                    v6[0] = 862
                }
                fmt.Println(v2, v4)
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v4 = (-(v4))
            }
            var v16 uint = 18446744073709551615
            _ = v16
        } else {
            var v17 uint32 = 802
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 uint = 2
                var v20 int64 = v4
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v17
        }
    }
    v21 := GenS2{F0: true}
    var v22 bool = (-473 != -990)
    var v23 uint = 0
    for fi24 := 0; fi24 < 19; fi24++ {
        fmt.Println(v1)
        v3["hello"] = (2 >= 42)
        _ = fi24
    }
    if v22 {
        if v22 {
            fmt.Println(v4, v4, v2)
        }
        switch 42 {
        case 0:
            for fi25 := 0; fi25 < 14; fi25++ {
                v23 -= v23
                v22 = v22
                var v26 string = "foo bar"
                _ = fi25
                _ = v26
            }
        case 616:
            v4 -= (v4 / ((9223372036854775807) | 1))
        default:
            var v27 float32 = 0.1
            _ = v27
        }
        v2 = v2
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        v22 = (false || true)
        _ = fi28
    }
    var v29 float64 = 0.0
    v4 = 2
    switch 32767 {
    case -1:
        v4 -= (-(v4))
    case 42:
        if ("hello" <= (("a" + ("" + "foo bar")) + "hello")) {
            var v30 uint32 = 2147483648
            _ = v30
        }
    case 4294967295:
        if (float64((-2.25 + -2.25)) >= 1e10) {
            if (-128 == -128) {
                switch 127 {
                case -1:
                    for fi31 := 0; fi31 < 15; fi31++ {
                        if (1 == 42) {
                            v32 := GenS2{F0: (v1 != 1)}
                            fmt.Println(v2, fi31, v29)
                            fmt.Println(v1, fi31)
                            _ = v32
                        } else {
                            fmt.Println(v1, v4, v1)
                        }
                        _ = fi31
                    }
                case 1:
                    var v33 int = int(v23)
                    _ = v33
                }
                var v34 int32 = v2
                _ = v34
            }
            delete(v3, ("hello" + ("\n\t" + "")))
            var v35 uint64 = 646
            _ = v35
        } else {
            if (-1 <= -117) {
                v22 = (773 <= (v23 & uint(v1)))
            }
        }
    default:
        v4 = (v4 / ((v4) | 1))
    }
    v29 -= 1e10
    v2 = v2
    var v36 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
