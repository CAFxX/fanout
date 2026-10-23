package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint16; F2 int64 }
type GenS5 struct { F0 int; F1 int8; F2 int16 }
type GenS6 struct { F0 uint8; F1 int; F2 int16 }
type GenS7 struct { F0 int64 }
func genFn0(p0 int32, p1 uint8, p2 uint64) int8 {
    v4 := []int{2147483647, 765, 4294967295, 4294967295}
    var v5 uint64 = (p2 * (p2 << ((42) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 int8, p1 int16, p2 uint64) bool {
    p0 = (^(p0))
    if false {
        if (!(("hello" + "") <= "a")) {
            v6 := []int{0}
            _ = v6
        } else {
            var v7 uint16 = 2
            v8 := GenS1{F0: -1, F1: (p1 * (p1 >> ((930) & 63))), F2: 2147483647}
            _ = v7
            _ = v8
        }
    } else {
        v9 := []int{}
        p2 += (p2 / ((p2) | 1))
        v9 = append(v9, 958)
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return (-2.25 != -2.25)
}

func main() {
    var v1 int64 = 42
    v2 := []uint32{2, 65535, 170, 0}
    v3 := make(map[int]int)
    v10 := GenS1{F0: 42, F1: -777, F2: 2147483648}
    var v11 float32 = 0.0
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 uint32 = 2
        var v14 uint16 = 0
        v14 -= v14
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS2{F0: 2}
        _ = fi15
        _ = v16
    }
    if ((v11 * v11) > v11) {
        fmt.Println(v1)
        if (-32768 == 0) {
            for fi17 := 0; fi17 < 4; fi17++ {
                var v18 int16 = 42
                fmt.Println(v1)
                switch ((("foo bar" + "foo bar") + "\n\t") + ("hello" + ("a" + "\n\t"))) {
                case "":
                    switch fi17 {
                    case 32768:
                        if false {
                            v3[747] = fi17
                            for fi19 := 0; fi19 < 6; fi19++ {
                                if false {
                                    v20 := v3[fi19]
                                    _ = v20
                                } else {
                                    if true {
                                        v21 := GenS1{F0: 2147483647, F1: (v18 + (v18 << ((1) & 63))), F2: 2147483647}
                                        if (!((^(v18)) > v18)) {
                                            fmt.Println(fi19)
                                            var v22 int16 = -387
                                            _ = v22
                                        }
                                        _ = v21
                                    }
                                    v1 += (v1 << ((2) & 63))
                                }
                                var v23 int8 = 100
                                switch "" {
                                case "foo bar":
                                    v23 = v23
                                case "a":
                                    if (false && (v18 == 1)) {
                                        if ((v18 << ((1) & 63)) >= v18) {
                                            var v24 int64 = (v1 << ((65535) & 63))
                                            var v25 int64 = -9223372036854775808
                                            var v26 uint16 = 975
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        }
                                        switch fi17 {
                                        case -1:
                                            v18 = v18
                                        case 908:
                                            v27 := GenS3{F0: 65535, F1: 750, F2: 1}
                                            _ = v27
                                        case 32768:
                                            switch (fi19 & fi17) {
                                            case 641:
                                                var v28 int16 = v18
                                                _ = v28
                                            }
                                        default:
                                            var v29 uint64 = 4294967295
                                            _ = v29
                                        }
                                    }
                                }
                                _ = fi19
                                _ = v23
                            }
                            v1 += (v1 * v1)
                        }
                    case -79:
                        if (!false) {
                            v1 = (v1 + v1)
                            for fi30 := 0; fi30 < 14; fi30++ {
                                switch (fi30 * (fi17 * (fi17 + (fi30 << ((1) & 63))))) {
                                case 749:
                                    if true {
                                        v11 += float32(-2.25)
                                        var v31 uint32 = 4294967295
                                        _ = v31
                                    }
                                default:
                                    _ = len(v2)
                                }
                                var v32 int = 128
                                v11 = float32((((-2.25 + (0.0 + 0.0)) * (float64(((1e10 * 0.1) - (v11 / ((v11 - ((v11 + (1.5 / ((1.5 + v11) * 0.0))) * v11)) * 0.0)))) / -2.25)) * float64((v11 + (0.0 + v11)))))
                                _ = fi30
                                _ = v32
                            }
                            for fi33 := 0; fi33 < 7; fi33++ {
                                if (528 < 1) {
                                    v18 = (v18 ^ -1)
                                } else {
                                    var v34 int = (fi33 >> ((0) & 63))
                                    v11 -= v11
                                    v35 := GenS4{F0: 445, F1: 1, F2: (v1 + (v1 / (((-(v1))) | 1)))}
                                    _ = v34
                                    _ = v35
                                }
                                var v36 int64 = v1
                                if ((^(fi17)) <= fi17) {
                                    if len(v2) > 0 {
                                        v2[0] = 65535
                                    }
                                    fmt.Println(v1)
                                } else {
                                    _ = len(v2)
                                }
                                _ = fi33
                                _ = v36
                            }
                        }
                    case 87:
                        var v37 int32 = 42
                        _ = v37
                    }
                case "hello":
                    v2 = append(v2, 529)
                case "\n\t":
                    v38 := GenS1{F0: -2147483648, F1: v18, F2: 42}
                    _ = v38
                default:
                    v39 := GenS5{F0: fi17, F1: -1, F2: v18}
                    _ = v39
                }
                _ = fi17
                _ = v18
            }
        }
        v11 = float32(1e10)
    } else {
        for fi40 := 0; fi40 < 2; fi40++ {
            for fi41 := 0; fi41 < 12; fi41++ {
                v42 := GenS5{F0: (fi40 + (-(fi41))), F1: 0, F2: 344}
                var v43 int16 = 0
                _ = fi41
                _ = v42
                _ = v43
            }
            _ = fi40
        }
    }
    for fi44 := 0; fi44 < 12; fi44++ {
        v45 := GenS6{F0: 128, F1: fi44, F2: -718}
        _ = fi44
        _ = v45
    }
    fmt.Println(v1)
    var v46 uint8 = 127
    v47 := GenS6{F0: v46, F1: -2147483648, F2: -32768}
    v48 := GenS7{F0: v1}
    fmt.Println(v1)
    v49 := GenS2{F0: 42}
    if true {
        v11 += (float32(1.5) / (v11 / -2.25))
    } else {
        v2 = append(v2, 1)
    }
    v11 = (1.5 + ((0.1 * 1.5) - v11))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
