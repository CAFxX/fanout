package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 string }
func genFn0(p0 bool, p1 int) int16 {
    v4 := []int{(p1 & (^(p1)))}
    switch "foo bar" {
    case "hello":
        fmt.Println(p0, p1)
    default:
        for fi5 := 0; fi5 < 4; fi5++ {
            fmt.Println(fi5)
            p1 += (fi5 * (fi5 / (((p1 >> ((247) & 63))) | 1)))
            for fi6 := 0; fi6 < 20; fi6++ {
                if p0 {
                    if p0 {
                        if false {
                            p1 = p1
                            if (p0 && p0) {
                                v4 = append(v4, fi6)
                                var v7 float32 = (((-2.25 + 1.5) - ((0.0 / float32(1e10)) + (-2.25 + (3.14159 + (1e10 / 0.1))))) + 3.14159)
                                var v8 int16 = 1000
                                _ = len(v4)
                                var v9 uint32 = 0
                                _ = v7
                                _ = v8
                                _ = v9
                            } else {
                                var v10 int16 = 454
                                var v11 int8 = -128
                                var v12 uint = 623
                                fmt.Println(p0, v10)
                                _ = v10
                                _ = v11
                                _ = v12
                            }
                            fmt.Println(p0, p0, p0)
                        }
                        var v13 string = (("foo bar" + (((("\n\t" + "hello") + "") + "hello") + ((("\n\t" + "hello") + "\n\t") + (("hello" + "hello") + "foo bar")))) + ("hello" + "hello"))
                        _ = v13
                    } else {
                        v14 := GenS1{F0: 65535}
                        p1 -= fi6
                        if (false && p0) {
                            v15 := GenS1{F0: 42}
                            switch (fi6 >> ((0) & 63)) {
                            case -2147483648:
                                var v16 uint16 = 1
                                _ = v16
                            case -1:
                                switch p1 {
                                case 426:
                                    var v17 uint16 = 32767
                                    _ = v17
                                }
                            case -843:
                                p0 = (p0 && true)
                            default:
                                fmt.Println(fi5, p1)
                            }
                            p0 = true
                            _ = v15
                        }
                        _ = v14
                    }
                }
                _ = fi6
            }
            _ = fi5
        }
    }
    v18 := GenS1{F0: 0}
    var v19 float64 = 0.1
    if len(v4) > 0 {
        v4[0] = (p1 << ((1) & 63))
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v4
    return 32767
}

func main() {
    v1 := []int{128, 936, 358}
    var v2 int16 = 599
    var v3 int16 = v2
    switch "foo bar" {
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = -957
        }
    }
    v20 := GenS1{F0: 4294967295}
    v3 = 32767
    var v21 uint32 = 2147483647
    for fi22 := 0; fi22 < 5; fi22++ {
        v2 = genFn0((false && ((v3 | -32768) >= v3)), -2147483648)
        _ = fi22
    }
    if false {
        v21 -= (v21 * v21)
        _ = len(v1)
        var v23 int32 = 1
        _ = v23
    } else {
        var v24 int32 = 0
        _ = v24
    }
    if ((786 != 4294967295) || true) {
        var v25 int16 = genFn0(false, 32767)
        var v26 uint32 = 65535
        if (232 < 18446744073709551615) {
            var v27 uint16 = 305
            for fi28 := 0; fi28 < 13; fi28++ {
                v29 := GenS2{F0: 987}
                v26 = (v26 + v21)
                switch "" {
                case "hello":
                    fmt.Println(v3, fi28)
                }
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 16; fi30++ {
                var v31 uint32 = (v26 % ((v21) | 1))
                _ = fi30
                _ = v31
            }
            _ = v27
        }
        _ = v25
        _ = v26
    } else {
        for fi32 := 0; fi32 < 4; fi32++ {
            v33 := GenS3{F0: 65, F1: "a"}
            _ = fi32
            _ = v33
        }
        var v34 int64 = 1
        for fi35 := 0; fi35 < 11; fi35++ {
            var v36 int64 = -930
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    if (!false) {
        if (0 < 0) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            for fi37 := 0; fi37 < 12; fi37++ {
                if false {
                    fmt.Println(v3)
                    var v38 int8 = 1
                    v3 = (v3 | v2)
                    _ = v38
                }
                fmt.Println(fi37)
                _ = fi37
            }
        }
        v21 = (v21 | v21)
    }
    if true {
        v21 -= 90
    } else {
        v2 = (v2 * (v3 << ((2) & 63)))
        var v39 string = ("hello" + "foo bar")
        _ = v39
    }
    var v40 float64 = -2.25
    for fi41 := 0; fi41 < 14; fi41++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi41
    }
    var v42 float64 = float64(0.1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v40
    _ = v42
}
