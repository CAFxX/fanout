package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 string; F1 int16; F2 uint }
type GenS3 struct { F0 int8; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]string)
    var v3 int32 = -606
    v4 := []int{-397, 4294967295}
    if true {
        for fi5 := 0; fi5 < 14; fi5++ {
            switch ("foo bar" + ("" + ("a" + "foo bar"))) {
            case "a":
                if false {
                    if ((fi5 ^ fi5) < fi5) {
                        v3 -= (-(v3))
                        var v6 int16 = 753
                        _ = v6
                    }
                    var v7 int32 = (v3 | 1)
                    if true {
                        v1 += 0.0
                        for fi8 := 0; fi8 < 5; fi8++ {
                            if true {
                                v2[""] = ("" + "\n\t")
                                _ = len(v2)
                                v9 := GenS1{F0: false, F1: -1}
                                v4 = append(v4, fi8)
                                _ = v9
                            } else {
                                v2["foo bar"] = "foo bar"
                            }
                            var v10 int32 = 561
                            _ = fi8
                            _ = v10
                        }
                    } else {
                        v1 = 0.0
                        delete(v2, "hello")
                        var v11 string = "\n\t"
                        _ = v11
                    }
                    _ = v7
                }
            case "":
                for fi12 := 0; fi12 < 2; fi12++ {
                    if (!(-875 < 42)) {
                        v13 := GenS2{F0: "\n\t", F1: 0, F2: 42}
                        var v14 bool = true
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v13
                        _ = v14
                    }
                    var v15 int64 = 42
                    _ = fi12
                    _ = v15
                }
            default:
                _ = len(v4)
            }
            fmt.Println(fi5)
            if (984 < 65535) {
                v16 := GenS2{F0: "\n\t", F1: 32767, F2: 255}
                for fi17 := 0; fi17 < 6; fi17++ {
                    v18 := GenS1{F0: ((!true) && ((true || false) || true)), F1: (v3 - (v3 + v3))}
                    v3 += v3
                    if ((float32(float64(float32((1.5 - 1e10)))) + v1) != v1) {
                        fmt.Println(fi5, fi5, fi17)
                        if true {
                            var v19 int = fi17
                            if false {
                                fmt.Println(fi5)
                            }
                            switch "a" {
                            case "hello":
                                fmt.Println(v19, v1, fi17)
                            }
                            _ = v19
                        } else {
                            v3 = v3
                        }
                        fmt.Println(fi17)
                    }
                    v3 = v3
                    v1 = (v1 * v1)
                    _ = fi17
                    _ = v18
                }
                v20 := GenS2{F0: ("a" + ("" + "")), F1: 415, F2: 47}
                _ = v16
                _ = v20
            }
            _ = fi5
        }
        _ = len(v2)
        if (!false) {
            var v21 uint32 = 0
            v21 += v21
            if false {
                var v22 int = -1
                v22 += -956
                switch (v22 - (-(v22))) {
                case -608:
                    v23 := GenS3{F0: 106, F1: true, F2: 18446744073709551615}
                    _ = v23
                default:
                    var v24 bool = true
                    _ = v24
                }
                _ = v22
            } else {
                var v25 uint64 = 42
                _ = v25
            }
            _ = v21
        }
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        if true {
            fmt.Println(v1, v3, fi26)
            var v27 uint16 = 32767
            if (!(4294967295 >= 65535)) {
                if (false || true) {
                    v4 = append(v4, (fi26 & -880))
                } else {
                    var v28 uint32 = 2147483647
                    v27 += (^(v27))
                    _ = v28
                }
            } else {
                v27 -= v27
            }
            _ = v27
        }
        v3 -= v3
        _ = fi26
    }
    v4 = append(v4, -615)
    v1 += float32(((1e10 / float64(v1)) + float64(v1)))
    if (!(169 == 127)) {
        fmt.Println(v1, v1)
        var v29 int32 = (v3 & -2147483648)
        v3 -= v3
        _ = v29
    } else {
        var v30 bool = false
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 int8 = 127
            v2[(("foo bar" + "foo bar") + "\n\t")] = "hello"
            switch (("\n\t" + "\n\t") + ("\n\t" + "hello")) {
            case "":
                v32 -= v32
            case "foo bar":
                for fi33 := 0; fi33 < 10; fi33++ {
                    v30 = (false && (true || (4294967295 <= 1)))
                    fmt.Println(v3, v32, fi33)
                    _ = fi33
                }
            }
            _ = fi31
            _ = v32
        }
        var v34 string = ""
        _ = v30
        _ = v34
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = 42
    }
    switch 9223372036854775807 {
    case 0:
        if (-962 != v3) {
            v3 = (v3 >> ((103) & 63))
            var v35 uint = 327
            _ = v35
        }
    }
    v3 -= (^(v3))
    var v36 int = 924
    v3 = v3
    v37 := GenS4{F0: 42, F1: 642, F2: 306}
    fmt.Println(v3, v3)
    var v38 int32 = v3
    v36 = 32768
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
