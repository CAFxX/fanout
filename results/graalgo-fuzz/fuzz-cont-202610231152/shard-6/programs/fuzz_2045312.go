package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint16; F1 int8 }
func genFn0(p0 int64, p1 int, p2 string) int8 {
    p2 = p2
    var v7 float64 = 0.0
    var v8 int64 = p0
    switch (("" + p2) + ((p2 + "hello") + p2)) {
    case "foo bar":
        if true {
            switch (p1 >> ((255) & 63)) {
            case 127:
                var v9 float32 = (((1e10 / float32(-2.25)) / -2.25) - (-2.25 + 1.5))
                _ = v9
            case 2147483647:
                var v10 uint = 42
                _ = v10
            }
            v11 := []int{p1, p1, (p1 | (p1 << ((136) & 63)))}
            for fi12 := 0; fi12 < 14; fi12++ {
                switch p1 {
                case -9223372036854775808:
                    v13 := GenS1{F0: 0, F1: 1}
                    _ = v13
                case 4294967295:
                    v14 := GenS1{F0: -1, F1: -1}
                    _ = v14
                }
                v8 = (p0 << ((724) & 63))
                _ = fi12
            }
            _ = v11
        }
    case "a":
        switch ("" + (p2 + "\n\t")) {
        case "foo bar":
            p0 -= 107
        }
    }
    var v15 bool = (-275 != int32(v8))
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    _ = v8
    return -86
}
func genFn1(p0 uint8, p1 int8, p2 uint16) bool {
    if false {
        p0 = (p0 * uint8(p2))
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := []int{(fi16 | (fi16 >> ((1) & 63))), (-(fi16)), 4294967295, fi16}
            _ = fi16
            _ = v17
        }
        p1 = p1
    }
    var v18 uint8 = p0
    var v19 string = (("hello" + "hello") + "hello")
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return true
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    var v5 string = "a"
    v6 := make(map[int]int)
    v20 := GenS1{F0: -663, F1: 42}
    var v21 uint16 = 32767
    v1 -= (v1 * v1)
    switch v5 {
    case "foo bar":
        v22 := GenS1{F0: 1000, F1: 0}
        _ = v22
    case "hello":
        switch "\n\t" {
        case "\n\t":
            var v23 int32 = 42
            _ = v23
        case "foo bar":
            if true {
                fmt.Println(v5)
                v24 := GenS1{F0: 869, F1: -2147483648}
                _ = v24
            } else {
                if (false || (true || (uint64(v1) != 1))) {
                    switch 32768 {
                    case 32:
                        v5 = "\n\t"
                    case 0:
                        v4[(v1 / (((v1 | v1)) | 1))] = 2147483647
                    }
                    if true {
                        for fi25 := 0; fi25 < 2; fi25++ {
                            fmt.Println(v1)
                            _ = fi25
                        }
                        if true {
                            var v26 int32 = -1
                            var v27 float64 = (3.14159 / 1.5)
                            _ = v26
                            _ = v27
                        } else {
                            v28 := GenS1{F0: 42, F1: -906}
                            if (!true) {
                                var v29 uint16 = (^(v21))
                                var v30 int16 = -32768
                                v30 = v30
                                _ = v29
                                _ = v30
                            }
                            v31 := GenS1{F0: 32767, F1: 42}
                            _ = v28
                            _ = v31
                        }
                    } else {
                        var v32 float32 = ((((float32(float64(1.5)) + 0.0) - (1e10 + (0.0 + (float32(1e10) - 1.5)))) + 1e10) + (3.14159 * (0.0 - float32(-2.25))))
                        for fi33 := 0; fi33 < 7; fi33++ {
                            var v34 uint64 = 4294967295
                            for fi35 := 0; fi35 < 19; fi35++ {
                                if true {
                                    if ("foo bar" == "\n\t") {
                                        fmt.Println(v1, v34, fi35)
                                        var v36 int32 = -2147483648
                                        _ = v36
                                    }
                                    v37 := GenS1{F0: 337, F1: 0}
                                    var v38 uint = 255
                                    if (!false) {
                                        v5 = v5
                                    }
                                    var v39 uint64 = (v34 << (((v38 / (((v38 | v38)) | 1))) & 63))
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                                fmt.Println(v1, v1)
                                var v40 int = (fi35 << ((243) & 63))
                                _ = fi35
                                _ = v40
                            }
                            for fi41 := 0; fi41 < 13; fi41++ {
                                var v42 uint32 = 464
                                _ = fi41
                                _ = v42
                            }
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    }
                    switch v5 {
                    case "foo bar":
                        v21 -= 138
                    }
                }
                var v43 int8 = genFn0(9223372036854775807, -9223372036854775808, v5)
                _ = v43
            }
        }
    case "":
        fmt.Println(v1, v5)
    default:
        v44 := []int{597}
        _ = v44
    }
    fmt.Println(v1)
    v45 := []int{0, -263, 32768}
    if len(v45) > 0 {
        v45[0] = 32767
    }
    v46 := GenS1{F0: -174, F1: -2147483648}
    switch ((v5 + ((v5 + "hello") + "hello")) + (v5 + "a")) {
    case "\n\t":
        v21 += 0
    case "":
        if true {
            var v47 uint32 = (v1 >> ((1) & 63))
            _ = v47
        }
    case "hello":
        if (!true) {
            if (!(!(100 < -128))) {
                v48 := GenS2{F0: 602, F1: -1}
                _ = v48
            } else {
                if ((-882 >= 1) && false) {
                    _ = len(v45)
                    var v49 uint64 = 42
                    if false {
                        if false {
                            fmt.Println(v1, v1, v5)
                            v5 = "a"
                            var v50 float32 = (0.1 * 0.0)
                            _ = v50
                        }
                        _ = len(v4)
                        v5 = ((v5 + "") + v5)
                    } else {
                        var v51 string = ((v5 + (("hello" + v5) + "hello")) + "")
                        var v52 uint32 = v1
                        _ = v51
                        _ = v52
                    }
                    _ = v49
                } else {
                    for fi53 := 0; fi53 < 1; fi53++ {
                        for fi54 := 0; fi54 < 14; fi54++ {
                            var v55 uint64 = 0
                            var v56 int = fi53
                            _ = fi54
                            _ = v55
                            _ = v56
                        }
                        _ = fi53
                    }
                    var v57 uint = 517
                    fmt.Println(v21, v21, v1)
                    _ = v57
                }
                var v58 int32 = -2147483648
                _ = v58
            }
            fmt.Println(v1, v21, v5)
        } else {
            var v59 uint16 = (v21 + (v21 << ((329) & 63)))
            v5 = v5
            if len(v45) > 0 {
                v45[0] = -2147483648
            }
            _ = v59
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
