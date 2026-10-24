package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 uint16; F1 string; F2 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 uint }
func genFn0(p0 int8) uint8 {
    var v5 int32 = 1
    if (v5 == -199) {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: (("" + "foo bar") + "a"), F1: 201, F2: (v5 - v5)}
            v8 := GenS1{F0: "a", F1: 968, F2: (v5 & v5)}
            v9 := GenS1{F0: "hello", F1: -639, F2: v5}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        if (((127 != 255) && (!((663 <= 4294967295) && (("a" + ("a" + "\n\t")) != "\n\t")))) || false) {
            v10 := GenS2{F0: (42 <= 2)}
            p0 = p0
            _ = v10
        } else {
            var v11 int64 = 1
            var v12 float32 = (0.1 - 1.5)
            _ = v11
            _ = v12
        }
    }
    var v13 string = "\n\t"
    v5 -= v5
    var v14 string = "hello"
    _ = p0
    _ = v13
    _ = v14
    _ = v5
    return 2
}
func genFn1(p0 uint16) bool {
    p0 = p0
    switch 763 {
    case -424:
        if false {
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 string = ("hello" + "foo bar")
                fmt.Println(v16, p0, fi15)
                _ = fi15
                _ = v16
            }
        } else {
            p0 = 327
            var v17 int8 = 0
            _ = v17
        }
    case -9223372036854775808:
        v18 := []int{464, -497, 996}
        _ = v18
    default:
        var v19 bool = true
        _ = v19
    }
    p0 = (p0 | 65535)
    _ = p0
    return (-2.25 < 3.14159)
}

func main() {
    v1 := []uint32{105, 2}
    v2 := make(map[string]bool)
    var v3 float64 = float64(3.14159)
    v4 := []uint32{2}
    var v20 uint8 = 42
    switch "a" {
    case "a":
        if true {
            if (!(42 < 0)) {
                switch ("a" + ("foo bar" + "hello")) {
                case "hello":
                    if false {
                        for fi21 := 0; fi21 < 7; fi21++ {
                            v22 := GenS3{F0: 42, F1: 0}
                            _ = fi21
                            _ = v22
                        }
                        var v23 int16 = 0
                        fmt.Println(v3, v20)
                        _ = v23
                    }
                default:
                    if false {
                        var v24 uint8 = 2
                        for fi25 := 0; fi25 < 10; fi25++ {
                            switch (fi25 << ((42) & 63)) {
                            case 0:
                                v3 = (v3 - float64(float32(0.1)))
                            case -901:
                                for fi26 := 0; fi26 < 2; fi26++ {
                                    v20 = 1
                                    var v27 uint64 = 555
                                    _ = fi26
                                    _ = v27
                                }
                            }
                            _ = fi25
                        }
                        if false {
                            var v28 int16 = 0
                            fmt.Println(v24)
                            _ = v28
                        } else {
                            v3 = v3
                            if len(v4) > 0 {
                                v4[0] = 2147483648
                            }
                            if false {
                                v29 := GenS4{F0: false, F1: false}
                                for fi30 := 0; fi30 < 12; fi30++ {
                                    v3 = (v3 * float64((3.14159 + float32(v3))))
                                    _ = fi30
                                }
                                v20 += (^(v24))
                                _ = v29
                            }
                        }
                        _ = v24
                    } else {
                        if (42 >= -88) {
                            v3 += (v3 * v3)
                        } else {
                            v31 := GenS2{F0: (9223372036854775807 == 42)}
                            v32 := GenS5{F0: (v20 - 0), F1: -1}
                            _ = v31
                            _ = v32
                        }
                        var v33 int8 = 42
                        _ = v33
                    }
                }
                for fi34 := 0; fi34 < 14; fi34++ {
                    _ = len(v2)
                    v35 := GenS3{F0: -1, F1: 4294967295}
                    var v36 uint64 = 4294967295
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                var v37 int8 = 0
                _ = v37
            }
        } else {
            fmt.Println(v20, v20)
        }
    case "foo bar":
        fmt.Println(v20, v20)
    default:
        for fi38 := 0; fi38 < 5; fi38++ {
            switch "foo bar" {
            case "a":
                v3 = v3
            case "hello":
                delete(v2, "")
            case "\n\t":
                if false {
                    var v39 uint = 0
                    _ = v39
                } else {
                    if false {
                        v40 := GenS4{F0: ((float64(1e10) == (v3 - 0.0)) || true), F1: false}
                        _ = v40
                    } else {
                        var v41 uint64 = 0
                        v20 = genFn0(42)
                        if (1 < 1) {
                            var v42 uint32 = 494
                            v43 := GenS3{F0: 0, F1: 65535}
                            _ = v42
                            _ = v43
                        }
                        _ = v41
                    }
                }
            default:
                v44 := GenS6{F0: 255, F1: "hello", F2: 42}
                _ = v44
            }
            _ = fi38
        }
    }
    if (!false) {
        fmt.Println(v20, v20)
        fmt.Println(v20)
        _ = len(v4)
    }
    fmt.Println(v3)
    fmt.Println(v20)
    for fi45 := 0; fi45 < 19; fi45++ {
        var v46 int32 = 42
        var v47 uint8 = (v20 ^ (^(v20)))
        v2[("a" + "")] = true
        _ = fi45
        _ = v46
        _ = v47
    }
    var v48 int64 = 1
    if false {
        v20 = 246
    }
    v49 := GenS7{F0: 42}
    v50 := GenS8{F0: -1, F1: 974}
    delete(v2, ("foo bar" + "a"))
    for fi51 := 0; fi51 < 15; fi51++ {
        var v52 uint16 = 65535
        _ = fi51
        _ = v52
    }
    v53 := GenS8{F0: 2147483647, F1: 18446744073709551615}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v50
    _ = v53
}
