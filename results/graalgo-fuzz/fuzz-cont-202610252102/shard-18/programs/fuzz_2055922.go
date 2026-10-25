package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 bool }
type GenS2 struct { F0 uint32; F1 bool }
func genFn0() int16 {
    if false {
        for fi5 := 0; fi5 < 8; fi5++ {
            var v6 uint32 = 404
            _ = fi5
            _ = v6
        }
        var v7 int = -186
        switch (v7 & v7) {
        case -9223372036854775808:
            fmt.Println(v7)
        case 4294967295:
            v7 = (v7 | (v7 >> ((758) & 63)))
        default:
            for fi8 := 0; fi8 < 5; fi8++ {
                for fi9 := 0; fi9 < 5; fi9++ {
                    var v10 uint32 = 703
                    v10 -= (v10 - (v10 + 42))
                    v10 -= (v10 - v10)
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
        }
        _ = v7
    }
    var v11 uint64 = 4294967295
    if true {
        fmt.Println(v11)
    }
    _ = v11
    return -76
}

func main() {
    var v1 int64 = 1
    var v2 int8 = 1
    var v3 string = "foo bar"
    var v4 uint16 = 0
    v12 := []int{127}
    if false {
        v3 = "foo bar"
        var v13 int32 = -1
        var v14 uint32 = 65535
        _ = v13
        _ = v14
    }
    var v15 uint64 = 42
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int16 = -32768
        _ = fi16
        _ = v17
    }
    if len(v12) > 0 {
        v12[0] = -270
    }
    var v18 int64 = (v1 / ((-197) | 1))
    switch 2147483647 {
    case 1:
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    case -9223372036854775808:
        var v19 bool = false
        _ = v19
    }
    if true {
        for fi20 := 0; fi20 < 6; fi20++ {
            for fi21 := 0; fi21 < 1; fi21++ {
                switch fi20 {
                case 32767:
                    v2 += v2
                case 1:
                    if ((true && (!(902 == 42))) || true) {
                        v22 := GenS1{F0: v2, F1: true, F2: false}
                        if ((2147483647 == 2147483647) && true) {
                            var v23 bool = false
                            switch v3 {
                            case "a":
                                v24 := GenS1{F0: (-(v2)), F1: (false || (!v23)), F2: true}
                                _ = v24
                            case "foo bar":
                                var v25 bool = v23
                                _ = v25
                            case "\n\t":
                                var v26 int32 = 1
                                _ = v26
                            default:
                                var v27 uint32 = 2147483648
                                _ = v27
                            }
                            v1 -= v1
                            if v23 {
                                v12 = append(v12, -187)
                                if true {
                                    v28 := GenS2{F0: 2147483648, F1: (((float32(3.14159) + 3.14159) == 1.5) && (!v23))}
                                    var v29 string = ("a" + "\n\t")
                                    v30 := GenS2{F0: 4294967295, F1: (v23 && ((-933 <= 1) && (332 > 609)))}
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                                var v31 int32 = 0
                                _ = v31
                            } else {
                                var v32 int64 = v18
                                var v33 uint = 773
                                fmt.Println(v23, v15, v3)
                                _ = v32
                                _ = v33
                            }
                            _ = v23
                        }
                        _ = v22
                    }
                case 4:
                    fmt.Println(fi21)
                }
                v18 = (v18 % ((-386) | 1))
                _ = fi21
            }
            _ = fi20
        }
        v34 := GenS1{F0: v2, F1: true, F2: ((true || (0 <= 0)) && true)}
        v3 = v3
        _ = v34
    } else {
        fmt.Println(v18, v4, v4)
    }
    fmt.Println(v2, v15)
    var v35 int16 = -32768
    var v36 int32 = -300
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    var v37 int32 = v36
    var v38 uint8 = 175
    v12 = append(v12, -353)
    fmt.Println(v36)
    var v39 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
