package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS3 struct { F0 bool; F1 bool; F2 uint8 }

func main() {
    var v1 uint = 1
    var v2 uint = (v1 >> ((0) & 63))
    var v3 float32 = 0.0
    var v4 uint16 = 65535
    v5 := make(map[uint32]uint32)
    v6 := []int{-2147483648}
    var v7 int64 = 9223372036854775807
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(v7)
        if (-2.25 <= 0.0) {
            var v9 int16 = -252
            fmt.Println(v2)
            _ = v9
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            if true {
                var v11 uint8 = 0
                v3 = (((float32(3.14159) + v3) + v3) * (v3 * ((1e10 + (((v3 + (v3 - (v3 - float32(((1e10 / float64(3.14159)) * 1e10))))) / float32(float64(((float32(1e10) + v3) * 1.5)))) + 1e10)) * v3)))
                _ = len(v5)
                _ = v11
            } else {
                fmt.Println(v7, fi8, v2)
                switch ((("a" + ("foo bar" + "\n\t")) + "\n\t") + "a") {
                case "foo bar":
                    v12 := GenS1{F0: v4}
                    _ = v12
                default:
                    if (uint64(v2) > 882) {
                        var v13 int = fi8
                        var v14 float32 = v3
                        v15 := v5[31]
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        var v16 bool = (2 == 150)
                        v17 := GenS1{F0: 65535}
                        _ = v16
                        _ = v17
                    }
                }
            }
            v3 = v3
            fmt.Println(fi8, v7, fi10)
            v18 := GenS2{F0: 449, F1: 22, F2: 42}
            var v19 uint32 = 2147483648
            var v20 uint16 = v4
            fmt.Println(fi10)
            _ = fi10
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi8
    }
    switch "" {
    case "hello":
        v21 := GenS2{F0: v7, F1: 249, F2: 42}
        _ = v21
    case "foo bar":
        v22 := GenS2{F0: 9223372036854775807, F1: 1, F2: 886}
        _ = v22
    }
    v2 = v2
    var v23 float32 = -2.25
    for fi24 := 0; fi24 < 6; fi24++ {
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 int32 = 1
            _ = fi25
            _ = v26
        }
        fmt.Println(fi24, v4)
        if (18446744073709551615 != 0) {
            v2 += v2
        } else {
            var v27 uint = 937
            _ = v27
        }
        _ = fi24
    }
    if false {
        if true {
            v1 = v1
        } else {
            v28 := GenS2{F0: 1, F1: 75, F2: 42}
            var v29 uint8 = 162
            v30 := GenS2{F0: -9223372036854775808, F1: (v29 ^ 3), F2: 0}
            _ = v28
            _ = v29
            _ = v30
        }
    }
    delete(v5, 619)
    v7 -= v7
    if (true || true) {
        var v31 uint = v2
        _ = v31
    }
    _ = len(v6)
    v32 := GenS2{F0: v7, F1: 1, F2: 0}
    var v33 int8 = 0
    if true {
        if (!false) {
            v34 := GenS3{F0: (false && (v4 >= 32767)), F1: (float32(float64(((1e10 * 1.5) * float32((0.0 + 3.14159))))) == (1.5 - (v3 / (v3 + float32((0.0 * float64(v3))))))), F2: 57}
            _ = v34
        }
        var v35 int = -2147483648
        _ = v35
    } else {
        if len(v6) > 0 {
            v6[0] = 0
        }
        if (false && (0.1 != -2.25)) {
            v4 = v4
            for fi36 := 0; fi36 < 19; fi36++ {
                for fi37 := 0; fi37 < 11; fi37++ {
                    v38 := GenS3{F0: true, F1: (true && true), F2: 42}
                    var v39 uint8 = 170
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                v4 = 255
                if (false || ((-1 < 0) || (751 == 1))) {
                    var v40 bool = (32767 > 65535)
                    fmt.Println(v1, fi36)
                    _ = v40
                } else {
                    var v41 uint = (v1 << ((v1) & 63))
                    _ = v41
                }
                _ = fi36
            }
            var v42 int8 = v33
            _ = v42
        } else {
            if ("foo bar" == "hello") {
                if len(v6) > 0 {
                    v6[0] = 127
                }
            } else {
                switch (("hello" + "hello") + "hello") {
                case "hello":
                    v7 -= v7
                case "foo bar":
                    switch ("a" + "hello") {
                    case "a":
                        v43 := GenS2{F0: v7, F1: 0, F2: 4294967295}
                        _ = v43
                    case "hello":
                        if true {
                            var v44 uint = v1
                            if false {
                                v2 = 332
                                var v45 int8 = (v33 / (((v33 >> (((v2 * 2)) & 63))) | 1))
                                var v46 uint32 = 930
                                _ = v45
                                _ = v46
                            } else {
                                switch 9223372036854775807 {
                                case -750:
                                    var v47 uint = (^(v1))
                                    _ = v47
                                case 1:
                                    var v48 int16 = 1
                                    _ = v48
                                }
                            }
                            v44 = v1
                            _ = v44
                        }
                    default:
                        _ = len(v6)
                    }
                case "\n\t":
                    var v49 uint = (v1 | (^(v2)))
                    _ = v49
                default:
                    v7 += int64(v4)
                }
            }
        }
        switch -82 {
        case 4294967295:
            v4 -= (v4 >> (((v1 << ((v2) & 63))) & 63))
        case 32768:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
    }
    v50 := GenS1{F0: (v4 >> ((v1) & 63))}
    for fi51 := 0; fi51 < 4; fi51++ {
        var v52 uint64 = 42
        var v53 bool = (("hello" + "") <= "hello")
        var v54 uint = (v1 << ((2) & 63))
        _ = fi51
        _ = v52
        _ = v53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v7
}
