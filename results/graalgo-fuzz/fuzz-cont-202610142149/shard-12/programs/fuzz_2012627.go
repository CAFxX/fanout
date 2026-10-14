package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 float64 = ((0.0 * 3.14159) - 0.0)
    var v2 int8 = 100
    v3 := []string{("a" + ("hello" + ""))}
    var v4 int8 = v2
    v2 -= 42
    switch "foo bar" {
    case "a":
        _ = len(v3)
    case "hello":
        v5 := GenS1{F0: 372, F1: (v4 + v2), F2: true}
        _ = v5
    default:
        var v6 uint64 = 4294967295
        _ = v6
    }
    var v7 int16 = -1
    var v8 int8 = v2
    v2 += v4
    if (255 == 2) {
        var v9 uint32 = 932
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            v3 = append(v3, "foo bar")
        case "a":
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        }
        _ = v9
    } else {
        var v10 int64 = 9223372036854775807
        _ = v10
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: v7, F1: v4, F2: (fi11 > 2147483647)}
        switch (^(fi11)) {
        case 32768:
            if true {
                fmt.Println(v7, v2)
                v13 := GenS2{F0: fi11}
                _ = v13
            }
        default:
            var v14 int = fi11
            _ = v14
        }
        if false {
            switch "a" {
            case "":
                v1 = v1
            }
            v15 := GenS3{F0: (false || (2147483647 < 2147483647))}
            var v16 int = (fi11 << ((1) & 63))
            _ = v15
            _ = v16
        } else {
            fmt.Println(v4, v2)
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 int32 = -2147483648
                if (-128 <= v2) {
                    if len(v3) > 0 {
                        v3[0] = (("hello" + "") + "\n\t")
                    }
                    var v19 int32 = 2147483647
                    _ = v19
                } else {
                    var v20 int = 1
                    _ = v20
                }
                _ = fi17
                _ = v18
            }
        }
        _ = fi11
        _ = v12
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        if (-1 != 9223372036854775807) {
            var v22 uint = 0
            _ = v22
        } else {
            var v23 int32 = -2147483648
            _ = v23
        }
        _ = fi21
    }
    v7 = 536
    _ = len(v3)
    var v24 bool = (0 < 0)
    fmt.Println(v2)
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := GenS4{F0: 32767}
        fmt.Println(v4, v1, v8)
        v1 -= (v1 - (v1 + 1.5))
        _ = fi25
        _ = v26
    }
    if v24 {
        if false {
            v27 := GenS1{F0: (v7 % (((v7 << ((404) & 63))) | 1)), F1: (v8 + (v2 + (v2 >> ((255) & 63)))), F2: (0 < 0)}
            for fi28 := 0; fi28 < 3; fi28++ {
                fmt.Println(v2, v2)
                _ = fi28
            }
            for fi29 := 0; fi29 < 18; fi29++ {
                v7 = v7
                _ = len(v3)
                v30 := GenS2{F0: (fi29 | fi29)}
                _ = fi29
                _ = v30
            }
            _ = v27
        }
        switch -9223372036854775808 {
        case 4294967295:
            if v24 {
                _ = len(v3)
                fmt.Println(v8)
                fmt.Println(v2, v1, v8)
            } else {
                switch "foo bar" {
                case "a":
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                default:
                    var v31 uint8 = 1
                    _ = v31
                }
                var v32 uint32 = 42
                _ = v32
            }
        case 127:
            for fi33 := 0; fi33 < 7; fi33++ {
                fmt.Println(v2)
                _ = fi33
            }
        }
        v34 := GenS1{F0: (v7 | v7), F1: (v8 & v8), F2: (921 == 0)}
        _ = v34
    }
    _ = len(v3)
    fmt.Println(v7, v8)
    if v24 {
        v1 -= 0.1
    } else {
        if v24 {
            v4 -= (-(v2))
            for fi35 := 0; fi35 < 19; fi35++ {
                v36 := GenS3{F0: ((v24 || true) && (int64(fi35) > 1))}
                var v37 int8 = -82
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        v3 = append(v3, ("hello" + "\n\t"))
        for fi38 := 0; fi38 < 11; fi38++ {
            var v39 string = ""
            _ = fi38
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
