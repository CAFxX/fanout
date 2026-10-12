package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 uint32 = 4294967295
    v2 := []string{"hello", "a"}
    v3 := make(map[int]int)
    var v4 uint16 = 1
    fmt.Println(v4)
    v4 -= v4
    v5 := GenS1{F0: 127}
    v4 -= (^(v4))
    v6 := GenS1{F0: 0}
    var v7 uint64 = 4294967295
    var v8 uint64 = v7
    var v9 uint8 = 2
    v9 = v9
    if ((0.0 + 0.1) == 1e10) {
        var v10 uint32 = (v1 << ((uint(v9)) & 63))
        var v11 uint = 42
        fmt.Println(v11, v1)
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        if ((false && true) || true) {
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := GenS2{F0: -9223372036854775808, F1: 2, F2: (v7 & 42)}
                v15 := GenS1{F0: 1}
                _ = fi13
                _ = v14
                _ = v15
            }
        } else {
            v7 -= v8
            fmt.Println(v1, v1, v7)
        }
        if true {
            v16 := GenS2{F0: -9223372036854775808, F1: uint(v4), F2: (^(v7))}
            _ = v16
        }
        _ = fi12
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        if ("hello" >= "foo bar") {
            if true {
                var v18 int32 = 850
                var v19 uint8 = 2
                switch fi17 {
                case -9223372036854775808:
                    for fi20 := 0; fi20 < 14; fi20++ {
                        var v21 int64 = 9223372036854775807
                        v2 = append(v2, "foo bar")
                        if true {
                            if (!false) {
                                switch "" {
                                case "":
                                    var v22 uint32 = 2147483648
                                    _ = v22
                                case "hello":
                                    var v23 uint32 = v1
                                    _ = v23
                                case "foo bar":
                                    if false {
                                        var v24 int32 = v18
                                        v9 -= v19
                                        _ = v24
                                    } else {
                                        var v25 uint = 2
                                        _ = v25
                                    }
                                default:
                                    var v26 bool = false
                                    _ = v26
                                }
                            } else {
                                var v27 bool = false
                                _ = v27
                            }
                            fmt.Println(v1, v21, v9)
                            switch (^(fi20)) {
                            case -510:
                                switch ("" + ("a" + "")) {
                                case "a":
                                    if true {
                                        var v28 uint = 197
                                        _ = v28
                                    }
                                case "foo bar":
                                    v8 += (^(v8))
                                case "":
                                    fmt.Println(fi17, v7, v19)
                                default:
                                    fmt.Println(v19, v21, fi17)
                                }
                            case -632:
                                fmt.Println(v8, v19, fi20)
                            }
                        } else {
                            delete(v3, fi20)
                            var v29 uint = 912
                            _ = v29
                        }
                        fmt.Println(v7, v1)
                        v30 := GenS3{F0: 1, F1: v4, F2: 0}
                        _ = fi20
                        _ = v21
                        _ = v30
                    }
                case 819:
                    fmt.Println(fi17, v1)
                }
                _ = v18
                _ = v19
            }
            if len(v2) > 0 {
                v2[0] = (("hello" + "foo bar") + "\n\t")
            }
            _ = len(v3)
        } else {
            var v31 int = fi17
            _ = v31
        }
        _ = fi17
    }
    if false {
        v4 += (v4 << ((181) & 63))
        if (!false) {
            switch "foo bar" {
            case "foo bar":
                var v32 uint16 = 77
                _ = v32
            }
            if true {
                v33 := GenS1{F0: 100}
                for fi34 := 0; fi34 < 4; fi34++ {
                    for fi35 := 0; fi35 < 8; fi35++ {
                        switch fi34 {
                        case 127:
                            fmt.Println(v7, v8)
                            fmt.Println(v9, v8, v8)
                        }
                        _ = fi35
                    }
                    for fi36 := 0; fi36 < 14; fi36++ {
                        if false {
                            var v37 uint32 = (^(v1))
                            v38 := GenS4{F0: (^(fi36))}
                            v1 = 2147483648
                            _ = v37
                            _ = v38
                        }
                        var v39 bool = ((9223372036854775807 > -9223372036854775808) && (2 >= 54))
                        if (!(!(432 >= 18446744073709551615))) {
                            var v40 int = 1
                            _ = v40
                        } else {
                            var v41 int = fi34
                            _ = v41
                        }
                        _ = fi36
                        _ = v39
                    }
                    fmt.Println(fi34, v8, v4)
                    _ = fi34
                }
                var v42 uint = 0
                _ = v33
                _ = v42
            }
            if (9223372036854775807 <= -9223372036854775808) {
                v4 = 32767
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v1, v7)
            if (v1 <= (v1 | v1)) {
                v1 += 1
                if len(v2) > 0 {
                    v2[0] = ("foo bar" + "\n\t")
                }
                switch "" {
                case "\n\t":
                    var v43 uint8 = 0
                    _ = v43
                case "":
                    for fi44 := 0; fi44 < 17; fi44++ {
                        v45 := GenS5{F0: -1, F1: v8, F2: 42}
                        _ = fi44
                        _ = v45
                    }
                case "foo bar":
                    v9 += (v9 ^ v9)
                default:
                    v4 = 2
                }
            } else {
                var v46 uint32 = 65535
                _ = v46
            }
        }
    } else {
        v47 := v3[-437]
        var v48 int16 = -390
        v49 := GenS6{F0: (v7 << ((892) & 63))}
        _ = v47
        _ = v48
        _ = v49
    }
    var v50 float64 = ((float64((-2.25 / 1.0)) + 0.0) - float64(((1.5 + (float32(0.0) / float32(3.14159))) / 1.5)))
    var v51 uint64 = uint64(v9)
    switch "hello" {
    case "a":
        _ = len(v3)
    case "hello":
        fmt.Println(v9)
    default:
        for fi52 := 0; fi52 < 20; fi52++ {
            switch fi52 {
            case 127:
                switch (fi52 / ((32767) | 1)) {
                case 127:
                    if true {
                        for fi53 := 0; fi53 < 11; fi53++ {
                            v8 = 4294967295
                            _ = len(v2)
                            v51 += v51
                            _ = fi53
                        }
                    }
                case 58:
                    fmt.Println(v50, fi52)
                case 635:
                    fmt.Println(fi52, v4)
                default:
                    var v54 uint = 2
                    _ = v54
                }
            case 4294967295:
                for fi55 := 0; fi55 < 15; fi55++ {
                    v8 = uint64(v1)
                    _ = fi55
                }
            case 32767:
                v4 -= (v4 ^ (v4 >> ((481) & 63)))
            }
            _ = fi52
        }
    }
    if false {
        var v56 float64 = -2.25
        fmt.Println(v9, v7, v8)
        v57 := GenS2{F0: int64(v4), F1: 201, F2: uint64(v9)}
        _ = v56
        _ = v57
    } else {
        v9 -= (v9 ^ v9)
    }
    v50 -= ((0.1 * 3.14159) / float64(3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
