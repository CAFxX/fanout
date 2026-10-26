package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint64; F2 bool }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint32, p1 uint) uint32 {
    switch "hello" {
    case "\n\t":
        if (42 >= 1) {
            fmt.Println(p1, p1)
            var v4 int16 = -1
            p0 += p0
            _ = v4
        }
    case "":
        var v5 uint8 = 127
        _ = v5
    case "a":
        var v6 int16 = 1
        _ = v6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{(fi7 << ((p1) & 63)), fi7, 577}
        v9 := GenS1{F0: p1}
        if true {
            var v10 float32 = float32(1e10)
            v8 = append(v8, 128)
            _ = v10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        v11 := GenS2{F0: "hello", F1: 42, F2: (2 >= 2)}
        v12 := GenS2{F0: "\n\t", F1: 304, F2: true}
        _ = v11
        _ = v12
    } else {
        switch 0 {
        case 0:
            p0 = p0
        case 32768:
            var v13 uint8 = 2
            _ = v13
        case -1:
            var v14 int = 42
            _ = v14
        default:
            for fi15 := 0; fi15 < 11; fi15++ {
                p1 = p1
                switch "\n\t" {
                case "a":
                    if ((669 < 255) && (true && true)) {
                        if true {
                            fmt.Println(p1)
                        } else {
                            if (85 >= 128) {
                                fmt.Println(p1)
                            } else {
                                var v16 int64 = -469
                                v16 -= -527
                                _ = v16
                            }
                            p0 -= (p0 | (p0 & p0))
                            v17 := GenS3{F0: 425}
                            _ = v17
                        }
                    }
                case "foo bar":
                    v18 := GenS3{F0: 42}
                    _ = v18
                case "\n\t":
                    var v19 uint64 = 0
                    _ = v19
                default:
                    switch (fi15 + 4294967295) {
                    case 4294967295:
                        var v20 uint32 = (^(p0))
                        _ = v20
                    default:
                        var v21 int16 = 1
                        _ = v21
                    }
                }
                _ = fi15
            }
        }
    }
    switch 609 {
    case 1:
        fmt.Println(p1, p0)
    case 42:
        p0 -= p0
    case 128:
        v22 := []int{42, 32768}
        _ = v22
    default:
        v23 := []int{9223372036854775807}
        _ = v23
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 327
}

func main() {
    var v1 uint = 0
    var v2 bool = (2147483647 <= 1)
    v3 := make(map[string]uint32)
    for fi24 := 0; fi24 < 1; fi24++ {
        v2 = ((1000 <= 4) && (86 <= 42))
        var v25 uint8 = 127
        fmt.Println(v2, fi24, v1)
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        if v2 {
            var v27 string = (("foo bar" + ("\n\t" + "a")) + "")
            _ = v27
        }
        _ = fi26
    }
    switch 32767 {
    case 32768:
        for fi28 := 0; fi28 < 20; fi28++ {
            switch ("foo bar" + "foo bar") {
            case "":
                v29 := GenS2{F0: "\n\t", F1: 2, F2: (-1 < -487)}
                _ = v29
            default:
                fmt.Println(v2)
            }
            fmt.Println(v2)
            var v30 int8 = 65
            _ = fi28
            _ = v30
        }
    case 487:
        var v31 float32 = 0.0
        _ = v31
    case 128:
        v2 = (255 < 1)
    }
    var v32 uint64 = 18446744073709551615
    var v33 uint16 = 32767
    var v34 uint64 = (^(v32))
    var v35 uint32 = 697
    var v36 int = -1
    var v37 int = -1
    v2 = v2
    _ = len(v3)
    v34 = (^(v32))
    if (!(1 != 100)) {
        var v38 uint = uint(v35)
        var v39 uint = v1
        _ = v38
        _ = v39
    } else {
        v40 := v3["foo bar"]
        v33 += (v33 % ((65535) | 1))
        _ = v40
    }
    v41 := []int{}
    v42 := GenS2{F0: ("" + "foo bar"), F1: 18446744073709551615, F2: v2}
    if (0.0 != 0.0) {
        v43 := GenS2{F0: ("a" + (((("hello" + "hello") + "\n\t") + ("" + "hello")) + (("foo bar" + ("foo bar" + "foo bar")) + ("foo bar" + "foo bar")))), F1: (v32 - v32), F2: (0 == -9223372036854775808)}
        var v44 int64 = 42
        fmt.Println(v34, v44, v37)
        _ = v43
        _ = v44
    }
    var v45 int8 = 127
    fmt.Println(v33)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v41
    _ = v42
    _ = v45
}
