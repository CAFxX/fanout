package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 bool; F2 int32 }

func main() {
    v1 := []string{"foo bar", "hello", "\n\t", ("a" + "hello")}
    v2 := []int{-2147483648}
    v3 := make(map[string]uint32)
    var v4 float32 = float32(0.1)
    var v5 float32 = float32((0.0 - -2.25))
    v6 := make(map[string]bool)
    var v7 int64 = 9223372036854775807
    if false {
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 18; fi9++ {
                switch (fi8 * 4294967295) {
                case 42:
                    var v10 int = (-(fi9))
                    _ = v10
                case 672:
                    delete(v6, ("hello" + "a"))
                case 4294967295:
                    fmt.Println(v7, fi9, fi8)
                default:
                    v4 = (0.0 / float32(0.1))
                }
                v4 = v5
                v7 -= (v7 >> ((1) & 63))
                v11 := GenS1{F0: v7, F1: -32768}
                _ = fi9
                _ = v11
            }
            v12 := GenS1{F0: v7, F1: 1}
            var v13 int16 = 935
            _ = fi8
            _ = v12
            _ = v13
        }
        fmt.Println(v7, v4)
        var v14 int32 = 2147483647
        _ = v14
    }
    fmt.Println(v4)
    var v15 uint32 = 0
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 uint32 = v15
        var v18 string = (("a" + ("a" + "a")) + "\n\t")
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        if (!(127 >= -34)) {
            v2 = append(v2, fi19)
            var v20 bool = (((v7 % ((v7) | 1)) <= v7) || true)
            switch "a" {
            case "\n\t":
                for fi21 := 0; fi21 < 2; fi21++ {
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    var v22 int64 = (-(v7))
                    var v23 int16 = 1
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
            _ = v20
        }
        v5 = v5
        _ = fi19
    }
    v24 := GenS1{F0: v7, F1: 1000}
    _ = len(v2)
    if ((v7 - v7) > (v7 + v7)) {
        for fi25 := 0; fi25 < 18; fi25++ {
            v26 := GenS1{F0: (^(v7)), F1: -32768}
            var v27 uint = 65535
            _ = fi25
            _ = v26
            _ = v27
        }
        var v28 int64 = (v7 << ((65535) & 63))
        if false {
            var v29 string = "hello"
            v30 := GenS2{F0: v15}
            v31 := GenS1{F0: 42, F1: 32767}
            _ = v29
            _ = v30
            _ = v31
        } else {
            v7 = v7
            if true {
                v5 += 1e10
                if false {
                    switch -1 {
                    case 2147483647:
                        switch (("hello" + ("a" + "")) + "foo bar") {
                        case "a":
                            if false {
                                v4 += v5
                                for fi32 := 0; fi32 < 5; fi32++ {
                                    v5 = (0.0 * 3.14159)
                                    _ = fi32
                                }
                                var v33 bool = (v5 < (((v4 * v5) / -2.25) / 1.5))
                                _ = v33
                            } else {
                                _ = len(v6)
                            }
                        case "hello":
                            fmt.Println(v7)
                        default:
                            if ((^(v15)) < v15) {
                                v34 := GenS1{F0: (v7 | v7), F1: 201}
                                _ = v34
                            } else {
                                v6[""] = (true || true)
                            }
                        }
                    }
                }
                fmt.Println(v5, v28)
            } else {
                v35 := GenS2{F0: v15}
                var v36 uint8 = 128
                _ = v35
                _ = v36
            }
        }
        _ = v28
    } else {
        v15 = 1
    }
    v2 = append(v2, 128)
    fmt.Println(v5)
    if true {
        switch 633 {
        case 23:
            var v37 int = 565
            _ = v37
        case 9223372036854775807:
            v38 := v3["a"]
            _ = v38
        case 967:
            var v39 uint8 = 202
            _ = v39
        default:
            fmt.Println(v15, v4)
        }
        if true {
            v4 -= 1.5
            if true {
                v40 := GenS3{F0: 108}
                v41 := GenS1{F0: -678, F1: 1}
                _ = v40
                _ = v41
            }
            v42 := GenS2{F0: v15}
            _ = v42
        } else {
            v15 = v15
        }
        var v43 uint8 = 101
        _ = v43
    } else {
        v4 += v5
        var v44 uint32 = v15
        _ = v44
    }
    var v45 uint32 = v15
    if true {
        v46 := GenS3{F0: 1}
        _ = v46
    }
    v1 = append(v1, "a")
    v47 := GenS2{F0: v15}
    v48 := GenS4{F0: -1, F1: true, F2: 68}
    if false {
        var v49 uint8 = 45
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v45
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v7
}
