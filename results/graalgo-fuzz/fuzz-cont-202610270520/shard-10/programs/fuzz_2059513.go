package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint64; F1 int32; F2 bool }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 uint32; F2 bool }
type GenS8 struct { F0 uint16; F1 bool; F2 string }
type GenS9 struct { F0 int; F1 uint32; F2 string }
type GenS10 struct { F0 int16; F1 bool }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint16; F1 int32; F2 int }
type GenS15 struct { F0 uint32; F1 int; F2 uint16 }

func main() {
    var v1 uint = 698
    v2 := make(map[uint32]bool)
    var v3 float32 = 1e10
    var v4 int32 = int32(v1)
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint64 = 0
        var v7 uint32 = 42
        v6 = (v6 ^ 138)
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint8 = 49
    v9 := []int{}
    var v10 uint32 = 65535
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS1{F0: ((true || true) || (!false)), F1: fi11}
        switch fi11 {
        case 2147483647:
            if len(v9) > 0 {
                v9[0] = fi11
            }
        default:
            for fi13 := 0; fi13 < 4; fi13++ {
                fmt.Println(fi11, v4)
                var v14 uint8 = v8
                switch fi11 {
                case 9223372036854775807:
                    if true {
                        v15 := GenS2{F0: 2, F1: (v4 << ((v1) & 63)), F2: true}
                        v1 = v1
                        if false {
                            v16 := GenS3{F0: 21, F1: (-(fi13))}
                            switch "hello" {
                            case "a":
                                v1 = (v1 | uint(v14))
                            case "foo bar":
                                v1 -= v1
                            }
                            if false {
                                v17 := GenS1{F0: false, F1: (fi13 % ((-2147483648) | 1))}
                                fmt.Println(v3)
                                var v18 int = fi11
                                switch (fi13 ^ v18) {
                                case 1:
                                    v10 -= 2
                                case 311:
                                    switch int(v1) {
                                    case 451:
                                        var v19 float32 = float32(1e10)
                                        _ = v19
                                    case -962:
                                        var v20 int = (fi13 - (^(fi13)))
                                        _ = v20
                                    }
                                default:
                                    _ = len(v2)
                                }
                                _ = v17
                                _ = v18
                            }
                            var v21 string = "hello"
                            _ = len(v9)
                            v1 -= (v1 / ((v1) | 1))
                            _ = v16
                            _ = v21
                        } else {
                            _ = len(v9)
                            v22 := GenS1{F0: true, F1: (^(fi13))}
                            v23 := GenS4{F0: ("hello" + "foo bar"), F1: 73}
                            _ = v22
                            _ = v23
                        }
                        _ = v15
                    }
                }
                _ = fi13
                _ = v14
            }
        }
        _ = fi11
        _ = v12
    }
    var v24 int16 = 0
    for fi25 := 0; fi25 < 11; fi25++ {
        v4 -= (v4 << (((v1 | v1)) & 63))
        switch (fi25 | (fi25 >> ((609) & 63))) {
        case 42:
            if (!false) {
                if true {
                    v26 := GenS5{F0: "a", F1: ""}
                    if (100 < 0) {
                        switch fi25 {
                        case 2147483647:
                            var v27 uint8 = (^(v8))
                            _ = v27
                        case 128:
                            if (-25 == 9223372036854775807) {
                                if (float32(1e10) != v3) {
                                    _ = len(v9)
                                    fmt.Println(v3)
                                    v3 -= 0.1
                                }
                            }
                        }
                    }
                    _ = v26
                } else {
                    v10 = (v10 + v10)
                }
            } else {
                if false {
                    v28 := GenS6{F0: 65535}
                    if true {
                        v1 = (v1 ^ uint(v10))
                        var v29 float32 = (((v3 * ((((v3 - float32(-2.25)) / 0.1) / float32(1e10)) / float32(3.14159))) / v3) / 1.0)
                        _ = v29
                    }
                    v30 := GenS2{F0: 787, F1: int32(v24), F2: true}
                    _ = v28
                    _ = v30
                }
            }
        case 1:
            v24 -= 0
        case 703:
            switch "" {
            case "":
                switch int(v4) {
                case 32767:
                    v1 += v1
                case 1:
                    if (v3 >= (3.14159 * (v3 + v3))) {
                        switch (fi25 % ((32768) | 1)) {
                        case 128:
                            switch "" {
                            case "foo bar":
                                var v31 uint16 = 32767
                                _ = v31
                            }
                        }
                        v8 = (v8 * 1)
                        switch fi25 {
                        case 32767:
                            if ((826 >= 9223372036854775807) && (false || (!false))) {
                                v32 := GenS7{F0: -835, F1: (v10 >> ((v1) & 63)), F2: true}
                                delete(v2, v10)
                                v4 = 2147483647
                                _ = v32
                            }
                        default:
                            var v33 string = (("" + "\n\t") + "")
                            _ = v33
                        }
                    } else {
                        for fi34 := 0; fi34 < 3; fi34++ {
                            if true {
                                v3 = float32((0.1 - 0.1))
                                v24 = 883
                                v35 := v2[v10]
                                v9 = append(v9, -364)
                                v36 := GenS6{F0: 736}
                                v24 -= (-(v24))
                                _ = v35
                                _ = v36
                            }
                            v37 := GenS8{F0: 0, F1: (v10 == 65535), F2: ""}
                            _ = fi34
                            _ = v37
                        }
                        v8 = v8
                        if len(v9) > 0 {
                            v9[0] = (-(fi25))
                        }
                    }
                case -128:
                    fmt.Println(v10, v24)
                default:
                    delete(v2, (^(v10)))
                }
            case "foo bar":
                v38 := GenS1{F0: true, F1: (-(fi25))}
                _ = v38
            default:
                fmt.Println(fi25, v1, v4)
            }
        default:
            v10 = (v10 >> ((v1) & 63))
        }
        v39 := v2[v10]
        _ = fi25
        _ = v39
    }
    v40 := GenS9{F0: -1, F1: (v10 + v10), F2: "a"}
    v41 := GenS10{F0: 1, F1: false}
    v4 -= (v4 & v4)
    switch "foo bar" {
    case "a":
        for fi42 := 0; fi42 < 4; fi42++ {
            switch "foo bar" {
            case "":
                if true {
                    v10 = 65535
                    if true {
                        fmt.Println(v10, v1)
                        fmt.Println(v24, v1)
                    } else {
                        fmt.Println(v3, v8, v1)
                        for fi43 := 0; fi43 < 4; fi43++ {
                            if true {
                                if false {
                                    v44 := GenS11{F0: 949}
                                    _ = v44
                                } else {
                                    if false {
                                        v45 := GenS9{F0: (fi43 << ((v1) & 63)), F1: v10, F2: "hello"}
                                        switch "" {
                                        case "\n\t":
                                            var v46 uint32 = 4294967295
                                            _ = v46
                                        }
                                        _ = v45
                                    } else {
                                        v1 = (v1 >> ((v1) & 63))
                                    }
                                }
                                v47 := GenS12{F0: (v8 % ((v8) | 1))}
                                _ = v47
                            } else {
                                if false {
                                    v1 -= v1
                                    v48 := GenS1{F0: true, F1: fi42}
                                    v8 = v8
                                    _ = v48
                                }
                                v24 += v24
                            }
                            switch (fi43 % (((fi43 & (fi42 * 42))) | 1)) {
                            case 9223372036854775807:
                                var v49 int16 = int16(v4)
                                _ = v49
                            }
                            if false {
                                var v50 uint64 = 2
                                if (2147483648 != v10) {
                                    var v51 uint = v1
                                    if false {
                                        v4 += 914
                                        var v52 int64 = 42
                                        v53 := GenS7{F0: -627, F1: v10, F2: true}
                                        _ = v52
                                        _ = v53
                                    }
                                    if (true && true) {
                                        var v54 uint = 18446744073709551615
                                        var v55 int = (fi42 * 32767)
                                        var v56 float64 = -2.25
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    } else {
                                        v51 = 47
                                        fmt.Println(v24, v1, fi43)
                                    }
                                    _ = v51
                                } else {
                                    fmt.Println(v3)
                                    fmt.Println(v24, fi42)
                                }
                                var v57 uint8 = 203
                                _ = v50
                                _ = v57
                            }
                            _ = fi43
                        }
                    }
                    if (v4 <= 42) {
                        v58 := GenS6{F0: 1}
                        _ = v58
                    } else {
                        v24 += 32767
                        v59 := GenS13{F0: v4}
                        v24 -= -32768
                        _ = v59
                    }
                }
            case "foo bar":
                fmt.Println(v4, fi42, fi42)
            default:
                v60 := GenS9{F0: fi42, F1: v10, F2: (("" + "\n\t") + "")}
                _ = v60
            }
            switch fi42 {
            case 0:
                v3 += 1e10
            case 739:
                v4 = (v4 % (((v4 ^ v4)) | 1))
            case 128:
                v61 := GenS14{F0: 605, F1: v4, F2: 32767}
                _ = v61
            }
            v8 = 128
            _ = fi42
        }
    case "foo bar":
        v62 := GenS15{F0: (v10 % (((v10 << ((65535) & 63))) | 1)), F1: 1, F2: uint16(v1)}
        _ = v62
    }
    fmt.Println(v3, v1)
    v8 -= (v8 & v8)
    if true {
        v2[v10] = (0.1 < 0.0)
    }
    v8 = (v8 | (v8 ^ (v8 ^ (v8 % (((v8 % ((127) | 1))) | 1)))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v8
    _ = v9
}
