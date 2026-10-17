package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64; F1 uint64; F2 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int = -870
    var v2 uint16 = 2
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    v5 := []uint32{306, 4294967295, uint32(v2)}
    v6 := make(map[uint32]int)
    if (!((958 >= 18446744073709551615) && false)) {
        for fi7 := 0; fi7 < 2; fi7++ {
            for fi8 := 0; fi8 < 10; fi8++ {
                switch (fi7 | v1) {
                case -2147483648:
                    if (227 >= 1) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v6[65535] = fi8
                    } else {
                        var v9 int16 = -1
                        if len(v5) > 0 {
                            v5[0] = 1
                        }
                        if len(v5) > 0 {
                            v5[0] = 2147483648
                        }
                        _ = v9
                    }
                case -9223372036854775808:
                    var v10 uint64 = 42
                    _ = v10
                }
                _ = fi8
            }
            _ = fi7
        }
    } else {
        v11 := GenS1{F0: 461, F1: -663}
        switch "\n\t" {
        case "a":
            v2 += v2
        case "foo bar":
            var v12 bool = ((!true) || true)
            _ = v12
        }
        var v13 int16 = 1000
        _ = v11
        _ = v13
    }
    var v14 int64 = -1
    for fi15 := 0; fi15 < 19; fi15++ {
        if (!(("\n\t" + "hello") < ("" + "\n\t"))) {
            v1 += 32768
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 float64 = (1e10 - ((float64(0.0) - 1e10) / 0.1))
                switch "foo bar" {
                case "\n\t":
                    if (0.0 > ((v17 + v17) * v17)) {
                        v18 := GenS1{F0: (fi16 / ((-1) | 1)), F1: int32(v2)}
                        _ = len(v5)
                        _ = v18
                    }
                default:
                    if (492 >= 2) {
                        v19 := GenS1{F0: (fi16 % (((^(fi16))) | 1)), F1: 0}
                        _ = v19
                    }
                }
                v1 += 481
                _ = fi16
                _ = v17
            }
            if ((0.0 - (3.14159 - 1.5)) != 3.14159) {
                v2 -= (v2 >> ((4294967295) & 63))
            } else {
                var v20 uint = 255
                v14 -= 9223372036854775807
                delete(v3, fi15)
                _ = v20
            }
        } else {
            v1 = 203
            for fi21 := 0; fi21 < 6; fi21++ {
                v1 = (v1 & fi15)
                _ = fi21
            }
        }
        if ((140 <= 0) && (!(!(("hello" + "\n\t") == "a")))) {
            var v22 uint = 209
            _ = v22
        } else {
            v23 := GenS2{F0: 0, F1: -1}
            _ = v23
        }
        var v24 float64 = float64(1.5)
        _ = fi15
        _ = v24
    }
    fmt.Println(v2, v14)
    for fi25 := 0; fi25 < 5; fi25++ {
        v14 = (v14 % ((v14) | 1))
        _ = fi25
    }
    if (!(("a" + (("foo bar" + ("a" + "hello")) + "foo bar")) > ("\n\t" + (("\n\t" + (("hello" + ("hello" + "\n\t")) + ((("a" + "\n\t") + "\n\t") + "\n\t"))) + ("" + ""))))) {
        v14 = v14
        v14 = 9223372036854775807
    }
    v14 = (^(v14))
    switch v1 {
    case 4294967295:
        for fi26 := 0; fi26 < 20; fi26++ {
            v27 := GenS3{F0: 9223372036854775807, F1: 42, F2: true}
            if false {
                v28 := GenS4{F0: 0}
                _ = v28
            } else {
                for fi29 := 0; fi29 < 14; fi29++ {
                    v30 := GenS5{F0: 127}
                    v31 := GenS3{F0: v14, F1: 291, F2: (true && (494 > 4294967295))}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                v2 = 174
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi26
            _ = v27
        }
    }
    v2 += 2
    if false {
        if false {
            if (-2.25 == 1.5) {
                switch (v1 << ((uint(v2)) & 63)) {
                case 226:
                    for fi32 := 0; fi32 < 17; fi32++ {
                        var v33 int32 = -1
                        v1 = 0
                        _ = fi32
                        _ = v33
                    }
                default:
                    var v34 float64 = 3.14159
                    _ = v34
                }
                v2 += v2
                if (!(("a" + "a") != ((("a" + "foo bar") + (("foo bar" + "\n\t") + "hello")) + "a"))) {
                    var v35 float32 = -2.25
                    for fi36 := 0; fi36 < 9; fi36++ {
                        v37 := GenS3{F0: 42, F1: 697, F2: (!false)}
                        _ = len(v5)
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                }
            }
        }
        for fi38 := 0; fi38 < 20; fi38++ {
            _ = len(v6)
            _ = fi38
        }
        if (((0.0 - 1e10) / (1.5 * -2.25)) != ((-2.25 + 1.5) / 1.0)) {
            v39 := GenS1{F0: 9223372036854775807, F1: -222}
            v14 = v14
            delete(v3, v1)
            _ = v39
        }
    } else {
        v6[175] = (-(v1))
        var v40 int = (-(v1))
        for fi41 := 0; fi41 < 17; fi41++ {
            v42 := GenS1{F0: (-(v1)), F1: -2147483648}
            _ = fi41
            _ = v42
        }
        _ = v40
    }
    var v43 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
