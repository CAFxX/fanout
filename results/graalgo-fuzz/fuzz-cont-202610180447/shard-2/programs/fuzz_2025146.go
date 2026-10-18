package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int8 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int, p1 bool, p2 uint16) uint16 {
    v5 := []int{32768, 156, -2147483648, p0}
    var v6 uint = 572
    v6 = (^(v6))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 277
}
func genFn1(p0 int8) int {
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(fi7, p0)
        p0 += (^(p0))
        _ = fi7
    }
    v8 := GenS1{F0: -32768, F1: 9223372036854775807, F2: 20}
    _ = p0
    _ = v8
    return 9223372036854775807
}

func main() {
    v1 := []string{"hello", "\n\t"}
    v2 := []uint32{42, 2147483648, 876}
    var v3 int = 42
    var v4 bool = true
    if v4 {
        var v9 uint32 = 0
        _ = v9
    }
    v3 -= (v3 % ((genFn1(-74)) | 1))
    if ("" == "hello") {
        switch ((((("" + "foo bar") + (("foo bar" + ("hello" + (("foo bar" + "a") + ("foo bar" + "\n\t")))) + "foo bar")) + (("" + "\n\t") + "\n\t")) + "a") + "foo bar") {
        case "a":
            v4 = ((112 > 2) && v4)
        case "\n\t":
            var v10 float32 = 1.5
            _ = v10
        case "foo bar":
            var v11 uint = 255
            _ = v11
        default:
            v3 = v3
        }
        v3 += (v3 >> ((0) & 63))
        switch (v3 + v3) {
        case -788:
            v12 := GenS1{F0: 0, F1: v3, F2: 36}
            _ = v12
        case 2147483647:
            fmt.Println(v4)
        case -9223372036854775808:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    } else {
        _ = len(v2)
        v13 := GenS1{F0: -500, F1: v3, F2: 100}
        _ = v13
    }
    fmt.Println(v4, v4)
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 int64 = 0
        _ = len(v1)
        _ = fi14
        _ = v15
    }
    switch (v3 ^ 4294967295) {
    case 1:
        v3 += v3
    default:
        fmt.Println(v3, v3)
    }
    if v4 {
        v3 += (v3 - genFn1(-39))
    }
    if false {
        v1 = append(v1, ("hello" + "hello"))
        for fi16 := 0; fi16 < 17; fi16++ {
            if v4 {
                fmt.Println(fi16)
            }
            for fi17 := 0; fi17 < 16; fi17++ {
                switch "a" {
                case "a":
                    var v18 uint64 = 2
                    _ = v18
                }
                if (17 > 9223372036854775807) {
                    var v19 int = fi16
                    if false {
                        if v4 {
                            if (53 <= 235) {
                                v4 = (100 == -128)
                                var v20 int32 = -1
                                fmt.Println(fi16, v4)
                                _ = v20
                            }
                            switch ("hello" + "") {
                            case "hello":
                                v21 := GenS1{F0: -806, F1: (v19 & v19), F2: -128}
                                _ = v21
                            case "a":
                                v19 += 454
                            case "":
                                if v4 {
                                    var v22 int8 = 0
                                    _ = v22
                                } else {
                                    v23 := GenS1{F0: -652, F1: fi16, F2: 127}
                                    _ = v23
                                }
                            default:
                                _ = len(v1)
                                var v24 uint = 944
                                _ = v24
                            }
                        }
                        var v25 float64 = ((float64(0.0) + (((((-2.25 - 3.14159) - float64(0.1)) / ((3.14159 + (1.5 - float64((0.1 - ((0.1 + 0.0) + (((1.5 - 0.0) / 1.0) * 1e10)))))) - 0.1)) * 0.1) / -2.25)) - 1.5)
                        _ = len(v2)
                        if (!v4) {
                            if ((!("hello" >= "foo bar")) || true) {
                                var v26 float64 = v25
                                var v27 int32 = 993
                                var v28 int = (v3 << ((1) & 63))
                                _ = v26
                                _ = v27
                                _ = v28
                            } else {
                                v19 = (fi16 + -207)
                                v29 := GenS1{F0: int16(fi17), F1: (-(v19)), F2: -128}
                                v30 := GenS2{F0: 32767, F1: 128}
                                _ = v29
                                _ = v30
                            }
                            if (false && (2147483647 == 2147483647)) {
                                v31 := GenS1{F0: 42, F1: fi17, F2: 42}
                                v25 += float64(((3.14159 * (1e10 * 1e10)) * (0.1 / 1.0)))
                                _ = v31
                            } else {
                                fmt.Println(v3)
                                if false {
                                    var v32 int64 = int64(fi16)
                                    if v4 {
                                        v4 = ((("hello" + "foo bar") + "\n\t") == "")
                                        v32 += 0
                                        v4 = v4
                                    }
                                    fmt.Println(v4, v19, v3)
                                    _ = v32
                                }
                            }
                        }
                        fmt.Println(v25)
                        var v33 bool = true
                        v3 = v3
                        _ = v25
                        _ = v33
                    }
                    var v34 uint32 = 898
                    _ = v19
                    _ = v34
                } else {
                    v35 := GenS3{F0: 2}
                    if (fi17 > fi17) {
                        var v36 uint16 = 255
                        _ = v36
                    } else {
                        v2 = append(v2, 42)
                        switch "foo bar" {
                        case "\n\t":
                            var v37 int = fi17
                            _ = v37
                        case "":
                            var v38 uint = 0
                            _ = v38
                        default:
                            switch "a" {
                            case "a":
                                v4 = (-867 != -539)
                            default:
                                var v39 float32 = -2.25
                                _ = v39
                            }
                        }
                    }
                    _ = v35
                }
                _ = fi17
            }
            _ = fi16
        }
        var v40 bool = ("\n\t" > ("hello" + "\n\t"))
        _ = v40
    } else {
        v3 = (v3 >> ((874) & 63))
        for fi41 := 0; fi41 < 19; fi41++ {
            for fi42 := 0; fi42 < 1; fi42++ {
                fmt.Println(fi41, v4)
                var v43 uint64 = 286
                var v44 uint = 456
                _ = fi42
                _ = v43
                _ = v44
            }
            if (1 == 65) {
                fmt.Println(fi41, v3)
            }
            v4 = (4294967295 < 255)
            _ = fi41
        }
    }
    var v45 uint64 = 42
    switch (("" + ("\n\t" + ((("foo bar" + "\n\t") + "\n\t") + ((("foo bar" + "") + ("a" + (("a" + "foo bar") + "a"))) + ("\n\t" + "\n\t"))))) + ("hello" + "")) {
    case "a":
        v45 += (v45 >> ((0) & 63))
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
}
