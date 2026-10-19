package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 string; F2 uint }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int; F1 int64; F2 uint16 }
type GenS5 struct { F0 int64; F1 int32; F2 int16 }
type GenS6 struct { F0 uint }

func main() {
    var v1 uint32 = 2
    var v2 string = (("" + (("\n\t" + ("hello" + "hello")) + (("\n\t" + "") + ("foo bar" + "")))) + "\n\t")
    var v3 int64 = 1
    var v4 uint16 = 1
    v5 := make(map[string]string)
    switch v2 {
    case "":
        var v6 uint = 65535
        _ = v6
    case "\n\t":
        v3 -= v3
    default:
        var v7 int32 = 1
        _ = v7
    }
    v8 := GenS1{F0: 573}
    var v9 uint64 = 206
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := []int{(^(fi10))}
        if (!false) {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            v1 += (v1 & 992)
            for fi12 := 0; fi12 < 18; fi12++ {
                v9 -= v9
                if (!false) {
                    var v13 bool = (((fi10 + -2147483648) < fi12) || (1 >= 18446744073709551615))
                    if (0.0 <= 0.0) {
                        v4 = uint16(v9)
                        var v14 bool = (fi10 < fi10)
                        switch v2 {
                        case "\n\t":
                            v15 := GenS2{F0: v3, F1: "hello", F2: 649}
                            _ = v15
                        case "foo bar":
                            v16 := GenS3{F0: 128, F1: v14}
                            _ = v16
                        default:
                            delete(v5, v2)
                        }
                        if (fi12 == fi10) {
                            if (v1 != v1) {
                                var v17 int32 = -2147483648
                                _ = v17
                            }
                        }
                        _ = v14
                    } else {
                        var v18 uint8 = 2
                        v18 = v18
                        switch (fi10 & fi10) {
                        case 42:
                            v4 += (^(v4))
                        default:
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        }
                        _ = v18
                    }
                    _ = v13
                } else {
                    v19 := GenS1{F0: 0}
                    if len(v11) > 0 {
                        v11[0] = 32768
                    }
                    _ = v19
                }
                _ = fi12
            }
        } else {
            if true {
                for fi20 := 0; fi20 < 16; fi20++ {
                    var v21 uint64 = v9
                    _ = fi20
                    _ = v21
                }
                switch 811 {
                case 4294967295:
                    v22 := GenS2{F0: v3, F1: v2, F2: 42}
                    _ = v22
                case -487:
                    var v23 float64 = -2.25
                    _ = v23
                }
            } else {
                fmt.Println(v3, fi10, v3)
            }
            v5[v2] = "foo bar"
            var v24 int8 = 56
            _ = v24
        }
        _ = fi10
        _ = v11
    }
    var v25 float64 = float64(0.0)
    _ = len(v5)
    if true {
        v2 = ""
        var v26 uint32 = 42
        v27 := v5[(("hello" + "a") + v2)]
        _ = v26
        _ = v27
    }
    v28 := GenS1{F0: 42}
    v29 := []int{0, -691, 32767, 4294967295}
    for fi30 := 0; fi30 < 4; fi30++ {
        if (42 >= 1) {
            var v31 uint = uint(v1)
            switch (fi30 & (fi30 ^ fi30)) {
            case -789:
                for fi32 := 0; fi32 < 7; fi32++ {
                    var v33 uint16 = (v4 + 488)
                    var v34 int = fi32
                    v35 := GenS4{F0: 4294967295, F1: -1, F2: 296}
                    _ = fi32
                    _ = v33
                    _ = v34
                    _ = v35
                }
            case 9223372036854775807:
                v5[v2] = (v2 + v2)
            default:
                switch 975 {
                case -526:
                    v31 = (^(v31))
                }
            }
            if false {
                v36 := v5[""]
                var v37 uint16 = 0
                _ = v36
                _ = v37
            } else {
                v9 -= (v9 % (((v9 ^ v9)) | 1))
                v38 := GenS5{F0: 0, F1: -1, F2: -32768}
                v39 := v5["\n\t"]
                _ = v38
                _ = v39
            }
            _ = v31
        }
        _ = fi30
    }
    v2 = v2
    v9 += 1
    for fi40 := 0; fi40 < 8; fi40++ {
        v3 = (v3 + -481)
        v41 := GenS3{F0: 255, F1: (float64((0.0 / -2.25)) <= v25)}
        _ = fi40
        _ = v41
    }
    if len(v29) > 0 {
        v29[0] = 464
    }
    fmt.Println(v2, v2, v1)
    fmt.Println(v3)
    v42 := GenS6{F0: 18446744073709551615}
    var v43 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v8
    _ = v9
}
