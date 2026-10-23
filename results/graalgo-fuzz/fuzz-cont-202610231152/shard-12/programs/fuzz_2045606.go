package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int; F2 uint }
type GenS4 struct { F0 uint8; F1 int8; F2 int64 }

func main() {
    var v1 bool = (!true)
    var v2 float32 = (0.1 / 1.5)
    var v3 float32 = 3.14159
    v4 := []uint32{2, 42, 687, 2147483647}
    var v5 float32 = float32(1.5)
    var v6 int16 = 645
    v5 = v2
    v1 = v1
    var v7 uint8 = 127
    v5 = (v2 / v3)
    var v8 uint64 = 173
    if v1 {
        var v9 int = 127
        var v10 uint64 = 506
        for fi11 := 0; fi11 < 14; fi11++ {
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
            _ = fi11
        }
        _ = v10
        _ = v9
    }
    v5 = v5
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            if true {
                v5 = v3
                v14 := GenS1{F0: true, F1: "a", F2: 0}
                switch "\n\t" {
                case "hello":
                    var v15 uint16 = 827
                    _ = v15
                case "a":
                    fmt.Println(v3, v6, v7)
                case "foo bar":
                    v7 += 255
                }
                _ = v14
            } else {
                var v16 bool = false
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (v1 || true) {
                    var v17 uint = 53
                    _ = v17
                } else {
                    switch ("a" + "foo bar") {
                    case "":
                        v18 := GenS1{F0: (v8 > (v8 * (v8 << ((uint(v8)) & 63)))), F1: "a", F2: 65535}
                        _ = v18
                    case "foo bar":
                        var v19 int64 = 623
                        var v20 int32 = -2147483648
                        if true {
                            if v16 {
                                switch (((("hello" + ("foo bar" + "foo bar")) + "hello") + "a") + "foo bar") {
                                case "hello":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                v8 = (^(v8))
                            }
                            if true {
                                v2 -= (v2 / ((((((0.0 + 0.1) * ((1e10 - v2) + 3.14159)) / 1e10) / v5) * v5) + v2))
                                if len(v4) > 0 {
                                    v4[0] = 4294967295
                                }
                                v4 = append(v4, 816)
                                var v21 int16 = 32767
                                var v22 string = "a"
                                _ = v21
                                _ = v22
                            } else {
                                var v23 string = (("hello" + "") + ((("foo bar" + ("" + (("" + "") + ("foo bar" + "foo bar")))) + ("foo bar" + "a")) + "foo bar"))
                                _ = v23
                            }
                            v8 = (v8 % ((v8) | 1))
                        } else {
                            var v24 int8 = 1
                            _ = v24
                        }
                        v4 = append(v4, 157)
                        if (255 != 2) {
                            v2 -= v5
                            var v25 float32 = (3.14159 - v5)
                            _ = v25
                        }
                        _ = v19
                        _ = v20
                    }
                }
                var v26 uint64 = 4294967295
                _ = v16
                _ = v26
            }
            _ = fi13
        }
        _ = len(v4)
        var v27 uint32 = 4294967295
        _ = fi12
        _ = v27
    }
    if v1 {
        var v28 uint8 = v7
        if len(v4) > 0 {
            v4[0] = 2147483648
        }
        for fi29 := 0; fi29 < 11; fi29++ {
            var v30 uint32 = 535
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    var v31 int64 = 0
    if v1 {
        v32 := GenS1{F0: ((1 == 127) && (9223372036854775807 < -875)), F1: "\n\t", F2: 778}
        switch -1 {
        case -402:
            v4 = append(v4, 81)
        default:
            v6 = v6
        }
        _ = v32
    } else {
        for fi33 := 0; fi33 < 8; fi33++ {
            var v34 uint16 = 255
            v8 += (^(v8))
            if (1e10 == 1.5) {
                fmt.Println(fi33, v31, v34)
                switch "" {
                case "foo bar":
                    v3 = (float32((float64((v5 - v2)) + ((-2.25 - 3.14159) * 0.0))) + (float32(((float64(v5) - 1.5) / 3.14159)) * v3))
                case "hello":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "a":
                    v31 = v31
                }
                for fi35 := 0; fi35 < 19; fi35++ {
                    fmt.Println(v7)
                    _ = fi35
                }
            }
            _ = fi33
            _ = v34
        }
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        _ = len(v4)
        switch ("\n\t" + (("" + "") + "\n\t")) {
        case "hello":
            if (true && v1) {
                v6 += v6
            } else {
                fmt.Println(v5, v3)
            }
        case "":
            v37 := GenS2{F0: (v7 | 127)}
            _ = v37
        }
        for fi38 := 0; fi38 < 10; fi38++ {
            if (83 <= 1) {
                v5 = v5
                v39 := GenS1{F0: (303 == 18446744073709551615), F1: ((("a" + "hello") + "") + "hello"), F2: 42}
                _ = v39
            } else {
                v40 := GenS3{F0: (v6 & (^(v6))), F1: 1, F2: 442}
                _ = v40
            }
            _ = len(v4)
            _ = fi38
        }
        _ = fi36
    }
    _ = len(v4)
    switch ("foo bar" + "hello") {
    case "hello":
        v41 := GenS3{F0: v6, F1: 488, F2: 448}
        _ = v41
    case "a":
        v42 := GenS4{F0: 186, F1: 100, F2: (v31 ^ v31)}
        _ = v42
    }
    if (!(!v1)) {
        var v43 uint64 = 2
        _ = v43
    } else {
        for fi44 := 0; fi44 < 14; fi44++ {
            for fi45 := 0; fi45 < 2; fi45++ {
                v6 += v6
                v31 += -1
                _ = fi45
            }
            for fi46 := 0; fi46 < 1; fi46++ {
                if len(v4) > 0 {
                    v4[0] = 65535
                }
                _ = fi46
            }
            _ = fi44
        }
        fmt.Println(v3, v31)
    }
    v47 := GenS1{F0: (-1 != 100), F1: "foo bar", F2: 427}
    var v48 float32 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
