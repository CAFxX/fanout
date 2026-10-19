package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS6 struct { F0 uint8; F1 int8; F2 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32; F1 string }
func genFn0(p0 int, p1 string) uint64 {
    p0 = p0
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(p0, fi6, p1)
        if false {
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := GenS1{F0: ((p1 + (p1 + ((p1 + "hello") + ((p1 + p1) + "foo bar")))) + p1), F1: fi6}
                var v9 string = "a"
                _ = fi7
                _ = v8
                _ = v9
            }
            v10 := GenS2{F0: 2147483647}
            p0 = p0
            _ = v10
        }
        _ = fi6
    }
    v11 := GenS3{F0: 32767}
    switch (p0 | -883) {
    case 127:
        fmt.Println(p0, p0)
    default:
        v12 := GenS4{F0: 532, F1: 0, F2: 0}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v11
    return 335
}
func genFn1() int32 {
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 uint32 = 472
        _ = fi13
        _ = v14
    }
    var v15 string = "hello"
    v15 = ("hello" + v15)
    _ = v15
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 int = -961
    var v3 int32 = 0
    var v4 float32 = float32((1.5 * 0.0))
    var v5 uint64 = 314
    if true {
        var v16 int32 = v3
        fmt.Println(v5, v16)
        _ = v16
    }
    v1 = append(v1, 2)
    v17 := GenS1{F0: ("" + (("foo bar" + "hello") + ("foo bar" + "hello"))), F1: v2}
    switch (v2 ^ v2) {
    case -249:
        v18 := GenS3{F0: 656}
        _ = v18
    case 9223372036854775807:
        var v19 int32 = v3
        _ = v19
    case 377:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v5 = (^(v5))
    for fi20 := 0; fi20 < 3; fi20++ {
        _ = len(v1)
        if false {
            v4 += 1e10
            v21 := GenS2{F0: 2147483648}
            _ = v21
        }
        _ = fi20
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS5{F0: 389, F1: 0, F2: 116}
        if false {
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v3 += 471
                for fi24 := 0; fi24 < 12; fi24++ {
                    var v25 int32 = (v3 << ((0) & 63))
                    fmt.Println(v3, v2, v4)
                    var v26 string = "hello"
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            var v27 string = ("foo bar" + ("" + "hello"))
            var v28 int32 = 42
            _ = v27
            _ = v28
        } else {
            fmt.Println(v3, fi22)
            v5 = v5
            if (!false) {
                var v29 int8 = 100
                switch v2 {
                case 127:
                    if (127 < (-(v29))) {
                        var v30 uint = uint(v5)
                        _ = v30
                    } else {
                        v31 := GenS6{F0: 0, F1: v29, F2: 1}
                        var v32 uint64 = v5
                        for fi33 := 0; fi33 < 2; fi33++ {
                            v1 = append(v1, 0)
                            if true {
                                v34 := GenS4{F0: (v3 * 42), F1: 789, F2: 128}
                                var v35 int32 = v3
                                _ = v34
                                _ = v35
                            }
                            v29 = 127
                            _ = fi33
                        }
                        _ = v31
                        _ = v32
                    }
                case 32768:
                    var v36 uint64 = (^(v5))
                    _ = v36
                default:
                    v37 := GenS5{F0: 497, F1: 2, F2: -208}
                    _ = v37
                }
                var v38 int16 = -805
                _ = v29
                _ = v38
            }
        }
        v5 = 42
        _ = fi22
        _ = v23
    }
    v4 = (v4 * v4)
    switch (v2 >> ((79) & 63)) {
    case -2147483648:
        switch v2 {
        case 4294967295:
            fmt.Println(v4)
        default:
            for fi39 := 0; fi39 < 17; fi39++ {
                var v40 string = "hello"
                var v41 uint = 180
                _ = fi39
                _ = v40
                _ = v41
            }
        }
    }
    var v42 uint16 = 255
    if true {
        switch ("\n\t" + (("a" + ("a" + "")) + ("hello" + "foo bar"))) {
        case "hello":
            v1 = append(v1, 2)
        case "a":
            v2 += v2
        case "foo bar":
            v4 += float32(-2.25)
        default:
            for fi43 := 0; fi43 < 1; fi43++ {
                v42 = v42
                var v44 bool = ("foo bar" > "")
                v3 = (v3 - genFn1())
                _ = fi43
                _ = v44
            }
        }
    }
    for fi45 := 0; fi45 < 16; fi45++ {
        v42 = 255
        fmt.Println(v2)
        if false {
            switch "a" {
            case "foo bar":
                v4 = v4
            case "hello":
                v42 = v42
            case "\n\t":
                for fi46 := 0; fi46 < 20; fi46++ {
                    v47 := GenS7{F0: 0}
                    v3 += genFn1()
                    _ = fi46
                    _ = v47
                }
            default:
                if (v3 <= (v3 << ((4294967295) & 63))) {
                    if ((false && false) && ((175 != 42) || (4294967295 == 255))) {
                        var v48 string = ("hello" + "a")
                        switch -933 {
                        case 0:
                            fmt.Println(v2, v2)
                        case 32768:
                            if (int8(v3) != 1) {
                                var v49 int8 = 0
                                _ = v49
                            } else {
                                _ = len(v1)
                                v48 = v48
                                v1 = append(v1, 1)
                            }
                        case 128:
                            v3 += int32(v5)
                        default:
                            v50 := GenS4{F0: int32(fi45), F1: 275, F2: 255}
                            _ = v50
                        }
                        _ = v48
                    } else {
                        v5 -= genFn0((-(v2)), ("" + "foo bar"))
                    }
                }
            }
        } else {
            v51 := GenS8{F0: 1, F1: ""}
            if true {
                _ = len(v1)
                v5 += (v5 - v5)
            }
            _ = v51
        }
        _ = fi45
    }
    if false {
        var v52 string = "hello"
        _ = v52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
