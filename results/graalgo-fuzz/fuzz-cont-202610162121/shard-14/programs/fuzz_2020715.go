package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint; F1 string; F2 uint64 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 bool, p1 int16) int32 {
    for fi4 := 0; fi4 < 19; fi4++ {
        for fi5 := 0; fi5 < 5; fi5++ {
            fmt.Println(p0, fi5)
            _ = fi5
        }
        _ = fi4
    }
    switch 9223372036854775807 {
    case 128:
        fmt.Println(p0)
    case -2147483648:
        v6 := GenS1{F0: 9223372036854775807, F1: 4294967295, F2: 225}
        _ = v6
    case -1:
        p0 = (133 != 2)
    default:
        var v7 int8 = 100
        _ = v7
    }
    fmt.Println(p0, p1)
    switch (("foo bar" + "foo bar") + "hello") {
    case "\n\t":
        switch -364 {
        case 705:
            for fi8 := 0; fi8 < 11; fi8++ {
                for fi9 := 0; fi9 < 12; fi9++ {
                    p1 = p1
                    if (p0 && ((p0 || p0) || true)) {
                        var v10 int = (fi9 >> ((900) & 63))
                        if (!p0) {
                            fmt.Println(p0)
                            v11 := GenS2{F0: 2, F1: "foo bar", F2: 951}
                            var v12 uint32 = 0
                            _ = v11
                            _ = v12
                        }
                        if (((("" + "foo bar") + ("\n\t" + "foo bar")) > "foo bar") || true) {
                            v13 := GenS3{F0: (p0 && p0), F1: false}
                            fmt.Println(p0)
                            _ = v13
                        } else {
                            v14 := GenS4{F0: 141, F1: 42}
                            var v15 uint8 = 42
                            if p0 {
                                fmt.Println(v15, fi9, v10)
                                var v16 uint = 2
                                _ = v16
                            }
                            _ = v14
                            _ = v15
                        }
                        _ = v10
                    } else {
                        v17 := GenS3{F0: true, F1: (p0 && (!(p0 && p0)))}
                        _ = v17
                    }
                    _ = fi9
                }
                _ = fi8
            }
        }
    case "hello":
        var v18 int64 = -1
        _ = v18
    case "":
        var v19 string = ""
        _ = v19
    }
    _ = p0
    _ = p1
    return 35
}
func genFn1() int16 {
    var v20 uint64 = 360
    v21 := GenS1{F0: 253, F1: 18446744073709551615, F2: 65535}
    _ = v20
    _ = v21
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{}
    v3 := []int{-1}
    v22 := v1[32768]
    v22 -= v22
    v22 = v22
    v23 := GenS5{F0: 1, F1: genFn0(("a" == "\n\t"), 1)}
    v24 := GenS2{F0: 4294967295, F1: "foo bar", F2: 610}
    v25 := GenS6{F0: 65535}
    fmt.Println(v22)
    fmt.Println(v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
