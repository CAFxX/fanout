package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int16 }

func main() {
    var v1 int8 = 127
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    var v4 int32 = 0
    v5 := make(map[string]int)
    var v6 int8 = (v1 * (v1 % ((127) | 1)))
    v1 = -11
    if true {
        switch 127 {
        case -9223372036854775808:
            switch 804 {
            case 4294967295:
                fmt.Println(v4, v4, v6)
            case -9223372036854775808:
                switch 2147483647 {
                case -2147483648:
                    v3[(((("foo bar" + "a") + "") + "") + ("a" + "hello"))] = true
                case 0:
                    switch 42 {
                    case -671:
                        v1 = (v6 << ((18446744073709551615) & 63))
                    case 1:
                        switch 4294967295 {
                        case 734:
                            v7 := []int{}
                            _ = v7
                        case 204:
                            v6 -= (v6 & (v6 << ((550) & 63)))
                        }
                    default:
                        if (false || ((v4 ^ 0) == (v4 | -499))) {
                            for fi8 := 0; fi8 < 17; fi8++ {
                                v6 += v6
                                if false {
                                    for fi9 := 0; fi9 < 19; fi9++ {
                                        delete(v3, (("hello" + "foo bar") + "a"))
                                        fmt.Println(fi9, v4, v4)
                                        _ = fi9
                                    }
                                    fmt.Println(v4, v1)
                                    if (!false) {
                                        switch fi8 {
                                        case -652:
                                            for fi10 := 0; fi10 < 16; fi10++ {
                                                v11 := GenS1{F0: (v6 + (-(v6))), F1: (916 < 2), F2: fi8}
                                                _ = fi10
                                                _ = v11
                                            }
                                        }
                                        delete(v5, "foo bar")
                                    }
                                } else {
                                    v12 := GenS2{F0: 0}
                                    _ = v12
                                }
                                _ = fi8
                            }
                        }
                    }
                }
            default:
                if (0 >= 255) {
                    var v13 float32 = 0.0
                    if (("" + "hello") != ((("\n\t" + "") + "a") + "\n\t")) {
                        v13 = v13
                        var v14 uint32 = 2147483647
                        _ = v14
                    }
                    _ = v13
                }
            }
        case -352:
            var v15 int16 = 260
            _ = v15
        case 0:
            var v16 float32 = 0.0
            _ = v16
        default:
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 float64 = ((0.0 * float64(0.1)) - 0.1)
                _ = fi17
                _ = v18
            }
        }
        v19 := v2["foo bar"]
        var v20 string = ((("hello" + "a") + "a") + "a")
        _ = v19
        _ = v20
    } else {
        fmt.Println(v4, v1)
    }
    switch 169 {
    case 0:
        var v21 uint64 = 2
        _ = v21
    case 127:
        for fi22 := 0; fi22 < 10; fi22++ {
            v4 = (v4 | 323)
            var v23 float64 = 3.14159
            for fi24 := 0; fi24 < 12; fi24++ {
                delete(v3, "")
                _ = fi24
            }
            _ = fi22
            _ = v23
        }
    case 85:
        fmt.Println(v1, v6, v4)
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        v26 := GenS2{F0: int64(fi25)}
        fmt.Println(fi25, v6, v6)
        if ((true || true) && (32767 < 32768)) {
            v1 += v1
        } else {
            fmt.Println(v1, v1)
            v27 := v3[(("foo bar" + "\n\t") + "hello")]
            for fi28 := 0; fi28 < 2; fi28++ {
                fmt.Println(fi28, v6, fi28)
                var v29 uint8 = 0
                _ = fi28
                _ = v29
            }
            _ = v27
        }
        _ = fi25
        _ = v26
    }
    v6 += (v1 * v6)
    var v30 uint8 = 2
    v31 := []int{1, -2147483648, 32768, 431}
    v32 := GenS3{F0: 0, F1: 802}
    var v33 int8 = v1
    var v34 uint8 = v30
    var v35 int32 = 0
    fmt.Println(v33, v35, v35)
    v36 := GenS4{F0: -872}
    v34 = 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
