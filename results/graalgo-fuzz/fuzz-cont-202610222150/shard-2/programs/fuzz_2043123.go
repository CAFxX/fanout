package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int32 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 int }
func genFn0(p0 int64, p1 uint32, p2 int64) int16 {
    var v4 int = -2147483648
    switch "" {
    case "":
        if false {
            var v5 uint64 = 700
            v6 := GenS1{F0: 41, F1: 0, F2: 1}
            p0 = p0
            _ = v5
            _ = v6
        } else {
            p2 -= p0
            for fi7 := 0; fi7 < 13; fi7++ {
                switch (fi7 & (v4 + v4)) {
                case 9223372036854775807:
                    for fi8 := 0; fi8 < 6; fi8++ {
                        fmt.Println(fi8, v4, p1)
                        switch (fi8 % ((-725) | 1)) {
                        case 128:
                            fmt.Println(p2, fi7, p2)
                            v9 := GenS2{F0: p1, F1: -123}
                            _ = v9
                        case 414:
                            var v10 int = v4
                            _ = v10
                        default:
                            v11 := GenS2{F0: p1, F1: 0}
                            _ = v11
                        }
                        _ = fi8
                    }
                }
                var v12 uint = 126
                v13 := GenS1{F0: 1, F1: 0, F2: 1}
                _ = fi7
                _ = v12
                _ = v13
            }
        }
    default:
        var v14 uint16 = 65535
        _ = v14
    }
    p0 = (p2 / ((0) | 1))
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := GenS1{F0: 0, F1: 1, F2: -1}
        _ = fi15
        _ = v16
    }
    v4 += (v4 - -1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    var v1 int16 = 427
    v2 := []string{}
    v3 := []uint32{42, 4294967295}
    if (2 != 4294967295) {
        v17 := GenS3{F0: "foo bar"}
        _ = v17
    } else {
        v2 = append(v2, "foo bar")
        var v18 int = 32768
        _ = v18
    }
    v1 -= (v1 << ((960) & 63))
    for fi19 := 0; fi19 < 9; fi19++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi19
    }
    if ((0 > 348) || true) {
        var v20 int64 = 0
        for fi21 := 0; fi21 < 19; fi21++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi22 := 0; fi22 < 16; fi22++ {
                _ = len(v2)
                var v23 int = fi21
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
        _ = v20
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v24 int32 = 1
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 int = (fi25 % ((fi25) | 1))
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v28 uint32 = 81
        _ = fi27
        _ = v28
    }
    var v29 float32 = ((3.14159 / 1.5) * (1.5 * 1.5))
    if true {
        var v30 int16 = 32767
        v31 := GenS4{F0: (("a" + "\n\t") + "hello")}
        if (true || (29 >= -1)) {
            if (2147483647 != -2147483648) {
                var v32 uint = 0
                _ = v32
            }
            if true {
                var v33 float32 = (3.14159 * v29)
                v30 = (v1 + v1)
                v29 += v29
                _ = v33
            }
            v1 = (-(v30))
        }
        _ = v30
        _ = v31
    } else {
        if true {
            fmt.Println(v1, v1)
        } else {
            for fi34 := 0; fi34 < 7; fi34++ {
                v1 += (v1 ^ v1)
                fmt.Println(fi34)
                fmt.Println(v29, v1)
                _ = fi34
            }
            fmt.Println(v1, v1)
            v35 := GenS5{F0: 127, F1: 4294967295}
            _ = v35
        }
    }
    var v36 int64 = -1
    var v37 uint16 = 683
    v1 = (v1 << ((170) & 63))
    fmt.Println(v29)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v36
    _ = v37
}
