package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint64; F1 uint64; F2 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint64, p2 int) int64 {
    v6 := GenS1{F0: 32768, F1: 349}
    v7 := GenS1{F0: 9223372036854775807, F1: 42}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 968
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{""}
    var v3 int8 = 42
    var v4 int16 = 32767
    var v5 int8 = (v3 ^ v3)
    var v8 int64 = 220
    v9 := GenS1{F0: -1, F1: 2}
    if (genFn0(42, 2, -1) < (^(v8))) {
        var v10 int64 = v8
        switch -2147483648 {
        case -328:
            v8 = (v10 - (^(v8)))
        case 127:
            v8 = v8
        default:
            v8 = 0
        }
        var v11 float64 = 0.1
        _ = v10
        _ = v11
    }
    v5 += (v5 >> ((1) & 63))
    var v12 float64 = 1.5
    if (-2.25 < -2.25) {
        var v13 float64 = (float64(0.0) - 0.1)
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 uint32 = 65535
            switch (fi14 + -9223372036854775808) {
            case -603:
                if false {
                    var v16 uint = 255
                    v3 = v3
                    v17 := GenS1{F0: fi14, F1: (v16 + (v16 >> (((^(v16))) & 63)))}
                    _ = v16
                    _ = v17
                } else {
                    if (("a" != ("\n\t" + "\n\t")) || true) {
                        var v18 string = ("hello" + "")
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 18; fi19++ {
                            v13 = v13
                            _ = fi19
                        }
                    }
                    if (false || false) {
                        var v20 uint64 = 0
                        _ = v20
                    }
                }
            default:
                if true {
                    if false {
                        fmt.Println(v13, v12)
                        fmt.Println(v8, v3, v5)
                    } else {
                        v21 := GenS1{F0: (fi14 * (fi14 * (fi14 << ((255) & 63)))), F1: 1}
                        fmt.Println(v13, v12)
                        _ = v21
                    }
                }
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        v22 := GenS2{F0: 459, F1: 276, F2: "hello"}
        _ = len(v2)
        if (!(-363 == -463)) {
            switch "hello" {
            case "a":
                for fi23 := 0; fi23 < 7; fi23++ {
                    v3 = v3
                    v24 := GenS2{F0: 66, F1: 1, F2: "hello"}
                    if true {
                        for fi25 := 0; fi25 < 10; fi25++ {
                            v3 -= (v5 << ((42) & 63))
                            _ = fi25
                        }
                    } else {
                        for fi26 := 0; fi26 < 8; fi26++ {
                            if (!(280 == -2147483648)) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(fi26, v4, v8)
                                v27 := GenS1{F0: (fi23 * (fi23 % ((fi23) | 1))), F1: 4294967295}
                                _ = v27
                            } else {
                                v2 = append(v2, "hello")
                            }
                            _ = fi26
                        }
                    }
                    _ = fi23
                    _ = v24
                }
            case "\n\t":
                v12 = v12
            case "":
                v28 := GenS3{F0: 171}
                _ = v28
            default:
                var v29 int = -2147483648
                _ = v29
            }
        } else {
            switch (("a" + "hello") + ((("foo bar" + "hello") + ("\n\t" + "hello")) + "hello")) {
            case "hello":
                var v30 uint32 = 42
                _ = v30
            case "\n\t":
                if (-529 != 1) {
                    var v31 bool = true
                    fmt.Println(v5, v3)
                    _ = v31
                } else {
                    for fi32 := 0; fi32 < 12; fi32++ {
                        for fi33 := 0; fi33 < 12; fi33++ {
                            v34 := GenS2{F0: 61, F1: 2, F2: "a"}
                            var v35 int = 32768
                            v12 = 1.5
                            v36 := GenS4{F0: 768}
                            var v37 int64 = (v8 * (v8 % ((-599) | 1)))
                            _ = fi33
                            _ = v34
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        if (!true) {
                            switch (fi32 - fi32) {
                            case 9223372036854775807:
                                v38 := GenS2{F0: 4294967295, F1: 2, F2: "foo bar"}
                                _ = v38
                            default:
                                v8 -= v8
                            }
                        } else {
                            if ((v3 >= v3) && (1e10 == -2.25)) {
                                v4 -= (v4 + (v4 >> ((2) & 63)))
                            } else {
                                delete(v1, 4294967295)
                                var v39 float32 = 1.5
                                _ = v39
                            }
                            var v40 string = "\n\t"
                            _ = v40
                        }
                        _ = fi32
                    }
                }
            case "a":
                switch "" {
                case "\n\t":
                    v8 = -1
                case "":
                    _ = len(v2)
                default:
                    v5 += (v3 << ((255) & 63))
                }
            }
        }
        _ = v22
    }
    fmt.Println(v5, v12)
    if (!false) {
        for fi41 := 0; fi41 < 14; fi41++ {
            fmt.Println(v4)
            fmt.Println(v8, v3, v3)
            if ((!(2 == 2147483647)) && (2 < 663)) {
                var v42 int16 = (v4 - int16(v5))
                _ = v42
            } else {
                var v43 int64 = genFn0(0, 42, fi41)
                v5 += (v5 & (v3 + v3))
                _ = v43
            }
            _ = fi41
        }
        delete(v1, 2)
    } else {
        fmt.Println(v5, v4)
    }
    var v44 int32 = -970
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v8
    _ = v9
}
