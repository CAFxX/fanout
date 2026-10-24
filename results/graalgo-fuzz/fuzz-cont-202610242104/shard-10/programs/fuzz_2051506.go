package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 bool; F1 int32; F2 uint16 }
type GenS6 struct { F0 uint8; F1 int8 }

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]bool)
    var v3 uint64 = 18446744073709551615
    v4 := []int{}
    fmt.Println(v3, v3)
    if true {
        v3 += (v3 >> ((65535) & 63))
        if (128 <= 128) {
            fmt.Println(v1, v3)
            v5 := GenS1{F0: 1, F1: -1}
            _ = v5
        }
    }
    var v6 int32 = 47
    delete(v2, ("" + ((("hello" + ("a" + "foo bar")) + "") + "\n\t")))
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS2{F0: 941}
        fmt.Println(v6, fi7, v6)
        if false {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 float32 = float32(1e10)
                switch "" {
                case "hello":
                    v1 += (((v10 / v10) / 1.0) / v10)
                case "":
                    var v11 int32 = -137
                    _ = v11
                case "a":
                    switch 4294967295 {
                    case -1:
                        v4 = append(v4, (fi7 & fi9))
                    case 32768:
                        switch (^(fi7)) {
                        case -472:
                            var v12 int = 1
                            _ = v12
                        }
                    case 1:
                        if true {
                            v13 := GenS2{F0: 0}
                            if (!true) {
                                v14 := v2["a"]
                                v15 := GenS1{F0: 42, F1: 42}
                                _ = v14
                                _ = v15
                            }
                            v16 := GenS2{F0: -1}
                            _ = v13
                            _ = v16
                        }
                    }
                }
                v6 -= 2147483647
                _ = fi9
                _ = v10
            }
        } else {
            v17 := GenS3{F0: "foo bar", F1: (v6 >> ((2) & 63))}
            var v18 int64 = 42
            if (1000 <= 32767) {
                v19 := GenS3{F0: "\n\t", F1: 625}
                _ = v19
            } else {
                if false {
                    if true {
                        v18 = (v18 ^ v18)
                        if true {
                            v18 -= v18
                        }
                        v6 = (v6 ^ v6)
                    }
                    var v20 bool = false
                    _ = v20
                } else {
                    v21 := GenS3{F0: "hello", F1: 1}
                    switch fi7 {
                    case 781:
                        v22 := GenS2{F0: v18}
                        _ = v22
                    case 127:
                        switch (fi7 * 32768) {
                        case -2147483648:
                            v3 += (v3 / ((4294967295) | 1))
                        case -9223372036854775808:
                            var v23 uint32 = 42
                            _ = v23
                        }
                    }
                    _ = v21
                }
                var v24 uint64 = (v3 / (((^(v3))) | 1))
                _ = v24
            }
            _ = v17
            _ = v18
        }
        _ = fi7
        _ = v8
    }
    v25 := GenS4{F0: 1, F1: -345}
    v1 += float32(-2.25)
    v6 = (^(v6))
    var v26 int8 = -87
    switch 1 {
    case 42:
        fmt.Println(v26, v6, v26)
    case 128:
        v26 = -113
    }
    if true {
        switch 128 {
        case -9223372036854775808:
            v27 := GenS3{F0: ("" + ("" + "\n\t")), F1: -670}
            _ = v27
        case 539:
            var v28 float64 = 0.0
            _ = v28
        }
        v3 -= v3
        fmt.Println(v6)
    } else {
        var v29 uint16 = 42
        switch 370 {
        case 127:
            var v30 int8 = v26
            _ = v30
        case 128:
            fmt.Println(v26)
        case 506:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            v3 = (v3 | v3)
        }
        if false {
            v26 -= (-(v26))
            fmt.Println(v3, v29)
        }
        _ = v29
    }
    if false {
        v26 = (v26 << ((18446744073709551615) & 63))
        var v31 float64 = 0.1
        _ = v31
    } else {
        if (false || true) {
            for fi32 := 0; fi32 < 8; fi32++ {
                v6 += v6
                fmt.Println(fi32, v26, v3)
                fmt.Println(v3)
                _ = fi32
            }
        }
    }
    fmt.Println(v26, v26, v26)
    fmt.Println(v26, v6)
    switch ((("hello" + "") + "") + "a") {
    case "a":
        switch -1 {
        case 72:
            fmt.Println(v3)
        case 9223372036854775807:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case "hello":
        v33 := GenS4{F0: 255, F1: v6}
        _ = v33
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        for fi35 := 0; fi35 < 10; fi35++ {
            var v36 bool = false
            if v36 {
                var v37 int8 = (v26 ^ (v26 + (v26 + (v26 ^ (^(v26))))))
                v38 := GenS5{F0: false, F1: v6, F2: 331}
                _ = v37
                _ = v38
            } else {
                var v39 float64 = -2.25
                delete(v2, "\n\t")
                fmt.Println(v3, v1, v39)
                _ = v39
            }
            v26 = v26
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    switch 1 {
    case -2147483648:
        delete(v2, ((("\n\t" + ("a" + ("" + "hello"))) + ("foo bar" + "a")) + ("\n\t" + "\n\t")))
    case 4294967295:
        v40 := GenS6{F0: 1, F1: 75}
        _ = v40
    default:
        var v41 int8 = v26
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v6
}
