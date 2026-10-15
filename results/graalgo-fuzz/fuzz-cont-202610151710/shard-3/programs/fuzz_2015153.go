package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 bool }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint; F1 uint8; F2 string }
type GenS4 struct { F0 bool; F1 int32; F2 int32 }
func genFn0(p0 uint8, p1 uint8, p2 uint16) int64 {
    for fi4 := 0; fi4 < 5; fi4++ {
        if true {
            var v5 uint8 = p1
            var v6 int = 127
            _ = v5
            _ = v6
        }
        p1 = 1
        _ = fi4
    }
    v7 := []int{32767, int(p2), 128, -9223372036854775808}
    v8 := GenS1{F0: 9223372036854775807, F1: 0, F2: (-32768 < 99)}
    v7 = append(v7, -2147483648)
    v9 := GenS2{F0: 816, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 827
}

func main() {
    var v1 int8 = 127
    var v2 uint16 = 32767
    var v3 float32 = (3.14159 - 1e10)
    v2 = v2
    v10 := []int{-220, 0, 2147483647, -1}
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(v3, v2, v3)
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi11
    }
    if ((1 != 42) && (!(762 == 627))) {
        var v12 uint8 = 42
        if len(v10) > 0 {
            v10[0] = -9223372036854775808
        }
        switch (("hello" + "\n\t") + "a") {
        case "\n\t":
            switch 48 {
            case 9223372036854775807:
                var v13 float32 = v3
                _ = v13
            }
        case "hello":
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 uint = 1
                var v16 string = ("" + "foo bar")
                if true {
                    var v17 bool = (1.5 >= float64(v3))
                    v18 := GenS1{F0: 64, F1: v15, F2: v17}
                    _ = v17
                    _ = v18
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        default:
            var v19 int8 = (v1 >> ((18446744073709551615) & 63))
            _ = v19
        }
        _ = v12
    } else {
        v10 = append(v10, 291)
        var v20 uint64 = 4294967295
        _ = v20
    }
    v21 := GenS2{F0: -2147483648, F1: -13}
    var v22 int = -9223372036854775808
    if (2147483648 <= 2147483647) {
        v22 += v22
    }
    v3 += float32((-2.25 * 0.1))
    var v23 uint8 = 108
    for fi24 := 0; fi24 < 5; fi24++ {
        v22 = (fi24 + v22)
        _ = fi24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 int8 = v1
        _ = fi25
        _ = v26
    }
    v27 := GenS1{F0: 42, F1: 42, F2: false}
    _ = len(v10)
    fmt.Println(v3, v23)
    switch (v22 / (((v22 | v22)) | 1)) {
    case 32767:
        var v28 float32 = (0.1 + v3)
        _ = v28
    case -9223372036854775808:
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 uint16 = 933
            v23 = (v23 % ((v23) | 1))
            _ = fi29
            _ = v30
        }
    case 1:
        if false {
            if false {
                fmt.Println(v3, v1, v2)
                _ = len(v10)
                switch "hello" {
                case "hello":
                    if true {
                        v1 -= (v1 * (v1 + 1))
                    }
                case "foo bar":
                    var v31 uint16 = 255
                    _ = v31
                default:
                    var v32 uint16 = (v2 - v2)
                    _ = v32
                }
            } else {
                if (v3 > (0.0 - v3)) {
                    if len(v10) > 0 {
                        v10[0] = 32768
                    }
                    switch ("hello" + ("foo bar" + "a")) {
                    case "hello":
                        v1 = -128
                    }
                    v1 = (v1 % ((100) | 1))
                } else {
                    v22 = (v22 + v22)
                    if true {
                        if ("a" <= "a") {
                            var v33 int64 = genFn0(243, v23, (^(v2)))
                            switch v22 {
                            case -202:
                                v2 -= 255
                            case -614:
                                fmt.Println(v1, v22, v23)
                            default:
                                for fi34 := 0; fi34 < 5; fi34++ {
                                    v33 = genFn0((v23 + (v23 ^ v23)), 1, (v2 << ((2) & 63)))
                                    if true {
                                        for fi35 := 0; fi35 < 18; fi35++ {
                                            v22 -= (v22 % ((-325) | 1))
                                            fmt.Println(fi34, v33)
                                            _ = fi35
                                        }
                                        fmt.Println(v23, v2)
                                    } else {
                                        var v36 string = (("hello" + ("\n\t" + "hello")) + "a")
                                        _ = v36
                                    }
                                    if false {
                                        v23 = 42
                                        var v37 uint8 = (v23 & (v23 & (v23 * (v23 >> ((42) & 63)))))
                                        _ = v37
                                    } else {
                                        var v38 uint8 = 157
                                        _ = v38
                                    }
                                    _ = fi34
                                }
                            }
                            v22 += (^(v22))
                            _ = v33
                        } else {
                            for fi39 := 0; fi39 < 10; fi39++ {
                                switch (v22 / ((fi39) | 1)) {
                                case -471:
                                    for fi40 := 0; fi40 < 16; fi40++ {
                                        switch "hello" {
                                        case "a":
                                            switch v22 {
                                            case 1:
                                                v41 := GenS3{F0: 572, F1: (v23 << ((1) & 63)), F2: "a"}
                                                _ = v41
                                            default:
                                                switch 42 {
                                                case 651:
                                                    v42 := GenS4{F0: ((true || false) || false), F1: int32(v23), F2: -121}
                                                    _ = v42
                                                case -1:
                                                    var v43 uint16 = uint16(v23)
                                                    _ = v43
                                                default:
                                                    var v44 uint64 = 0
                                                    _ = v44
                                                }
                                            }
                                        case "":
                                            v45 := GenS1{F0: 42, F1: 694, F2: (472 <= (v2 & 0))}
                                            _ = v45
                                        case "hello":
                                            var v46 uint16 = uint16(v23)
                                            _ = v46
                                        default:
                                            var v47 int16 = 32767
                                            _ = v47
                                        }
                                        _ = fi40
                                    }
                                case 42:
                                    var v48 uint = 1
                                    _ = v48
                                case 1:
                                    if true {
                                        for fi49 := 0; fi49 < 10; fi49++ {
                                            var v50 int8 = int8(v23)
                                            _ = fi49
                                            _ = v50
                                        }
                                    }
                                }
                                _ = fi39
                            }
                            v3 += (0.1 / -2.25)
                            if true {
                                v2 = (v2 >> ((1) & 63))
                            }
                        }
                    } else {
                        var v51 uint64 = 4294967295
                        v1 = (v1 / (((v1 + 127)) | 1))
                        _ = v51
                    }
                    var v52 int64 = 814
                    _ = v52
                }
                var v53 uint64 = 2
                _ = v53
            }
        }
    }
    var v54 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v54
}
