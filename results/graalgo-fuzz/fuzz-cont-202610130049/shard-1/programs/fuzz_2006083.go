package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 bool; F2 string }
func genFn0(p0 uint16, p1 uint32) uint {
    var v4 int = 42
    v5 := GenS1{F0: 1, F1: 981}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 255
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{541, 2147483647}
    var v3 int8 = 100
    switch -9223372036854775808 {
    case 4294967295:
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 int = (fi6 + (fi6 << ((4294967295) & 63)))
            _ = fi6
            _ = v7
        }
    case -2147483648:
        v3 = v3
    }
    v8 := GenS2{F0: 255, F1: 0}
    v9 := GenS3{F0: 2}
    if ("foo bar" < ("a" + ("hello" + ""))) {
        v3 -= (v3 ^ v3)
    } else {
        v3 = (^(v3))
        v3 -= (v3 * (-(v3)))
        if false {
            var v10 uint32 = 4294967295
            switch ((("" + ((("\n\t" + ("" + "\n\t")) + "a") + "hello")) + ("foo bar" + "")) + "\n\t") {
            case "foo bar":
                v10 = (^(v10))
            }
            _ = v10
        }
    }
    if false {
        var v11 int32 = 489
        v3 = (v3 % (((v3 * v3)) | 1))
        if (genFn0(42, 2) < 60) {
            v11 = v11
            var v12 float32 = 3.14159
            if ((0.1 / (1e10 + (1e10 / 0.1))) < (3.14159 * -2.25)) {
                var v13 uint8 = 1
                if (false || (786 >= 42)) {
                    var v14 int32 = (v11 + (v11 << ((224) & 63)))
                    var v15 string = "hello"
                    var v16 bool = ((false || true) && true)
                    _ = v14
                    _ = v15
                    _ = v16
                }
                _ = v13
            }
            _ = v12
        }
        _ = v11
    } else {
        v3 = (-(v3))
        for fi17 := 0; fi17 < 12; fi17++ {
            switch "\n\t" {
            case "":
                var v18 int32 = 2147483647
                _ = v18
            case "hello":
                v3 -= v3
            default:
                var v19 uint64 = 2
                _ = v19
            }
            _ = fi17
        }
    }
    _ = len(v2)
    v3 -= -128
    var v20 uint64 = 18446744073709551615
    if false {
        v21 := GenS2{F0: 68, F1: 261}
        if true {
            var v22 int64 = 42
            _ = v22
        } else {
            v3 = 42
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 uint64 = (v20 - (^(v20)))
                if false {
                    for fi25 := 0; fi25 < 7; fi25++ {
                        switch (("a" + "foo bar") + "\n\t") {
                        case "":
                            var v26 uint32 = 42
                            _ = v26
                        case "\n\t":
                            switch "" {
                            case "hello":
                                v27 := GenS4{F0: -9223372036854775808, F1: false, F2: ("" + "foo bar")}
                                _ = v27
                            default:
                                if (118 <= -808) {
                                    _ = len(v1)
                                }
                                var v28 int16 = int16(fi25)
                                _ = v28
                            }
                        case "foo bar":
                            v24 -= v24
                        default:
                            var v29 bool = (279 <= 9223372036854775807)
                            _ = v29
                        }
                        fmt.Println(v24)
                        v24 = (v20 & (^(v20)))
                        var v30 uint = 2
                        _ = fi25
                        _ = v30
                    }
                } else {
                    v31 := v1[-9]
                    v3 = v3
                    _ = v31
                }
                _ = fi23
                _ = v24
            }
        }
        _ = v21
    } else {
        v32 := GenS2{F0: 0, F1: 2}
        _ = v32
    }
    fmt.Println(v20, v3)
    if false {
        v3 = (v3 + v3)
        if false {
            var v33 string = "a"
            if (!(0 == 1)) {
                if false {
                    var v34 uint16 = 255
                    _ = v34
                } else {
                    var v35 int8 = v3
                    _ = len(v1)
                    _ = v35
                }
                var v36 int64 = -1
                _ = v36
            }
            var v37 int8 = v3
            _ = v33
            _ = v37
        }
    } else {
        for fi38 := 0; fi38 < 12; fi38++ {
            if (18446744073709551615 < 4294967295) {
                if false {
                    var v39 string = "\n\t"
                    for fi40 := 0; fi40 < 4; fi40++ {
                        _ = len(v1)
                        _ = fi40
                    }
                    _ = v39
                }
            } else {
                fmt.Println(fi38)
                for fi41 := 0; fi41 < 20; fi41++ {
                    v3 += 1
                    v1[9223372036854775807] = 559
                    v3 -= (v3 & v3)
                    v20 = 862
                    v42 := GenS3{F0: 339}
                    _ = fi41
                    _ = v42
                }
            }
            v3 -= v3
            if (-1 == 1) {
                v43 := GenS3{F0: 2}
                _ = v43
            }
            _ = fi38
        }
        v20 -= (v20 / ((v20) | 1))
        v44 := GenS1{F0: 12, F1: 0}
        _ = v44
    }
    switch 1 {
    case -621:
        v3 += -29
    }
    if true {
        var v45 uint16 = 556
        _ = v45
    }
    for fi46 := 0; fi46 < 11; fi46++ {
        for fi47 := 0; fi47 < 5; fi47++ {
            v3 -= 122
            _ = fi47
        }
        v3 = (v3 + v3)
        _ = fi46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
