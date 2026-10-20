package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 bool; F1 uint8; F2 bool }
type GenS6 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS7 struct { F0 int8 }
func genFn0(p0 uint16, p1 uint8) int32 {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: p0, F1: 2147483647}
        v8 := GenS1{F0: p0, F1: -735}
        _ = fi6
        _ = v7
        _ = v8
    }
    if (2147483647 == -9223372036854775808) {
        v9 := []int{32767, -836, 42}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = len(v9)
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        if (!true) {
            p1 += (p1 * (p1 >> ((320) & 63)))
            for fi11 := 0; fi11 < 8; fi11++ {
                p0 += (p0 << ((4294967295) & 63))
                _ = fi11
            }
        }
        for fi12 := 0; fi12 < 14; fi12++ {
            fmt.Println(fi10, p1, fi12)
            _ = fi12
        }
        var v13 uint64 = 0
        _ = fi10
        _ = v13
    }
    if false {
        if true {
            if true {
                var v14 string = ("\n\t" + "a")
                _ = v14
            }
            v15 := GenS2{F0: "hello"}
            if false {
                v16 := GenS3{F0: p0, F1: 46, F2: 0}
                _ = v16
            } else {
                fmt.Println(p1, p1)
                if (false || true) {
                    v17 := []int{2147483647, 127}
                    if len(v17) > 0 {
                        v17[0] = -9223372036854775808
                    }
                    _ = v17
                } else {
                    v18 := []int{32768}
                    if (true || true) {
                        for fi19 := 0; fi19 < 7; fi19++ {
                            fmt.Println(p0, p0, fi19)
                            _ = fi19
                        }
                        var v20 uint64 = 391
                        var v21 float32 = 1.5
                        _ = v20
                        _ = v21
                    } else {
                        if (42 == (p1 >> ((923) & 63))) {
                            if (-74 >= 9223372036854775807) {
                                if false {
                                    var v22 bool = ((p1 / ((p1) | 1)) != (p1 * 128))
                                    _ = v22
                                }
                                fmt.Println(p1, p0)
                            } else {
                                var v23 bool = (127 <= 128)
                                switch 32768 {
                                case -594:
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                case 128:
                                    var v24 string = ((("a" + (("foo bar" + "foo bar") + "\n\t")) + ((("" + "foo bar") + "a") + "hello")) + ("\n\t" + "\n\t"))
                                    _ = v24
                                }
                                p0 -= (p0 + (p0 << ((65535) & 63)))
                                _ = v23
                            }
                            p1 += (^(p1))
                        } else {
                            fmt.Println(p1)
                        }
                        switch 32767 {
                        case 32768:
                            if false {
                                fmt.Println(p1, p1)
                            } else {
                                fmt.Println(p1, p0)
                                for fi25 := 0; fi25 < 2; fi25++ {
                                    var v26 uint = 39
                                    _ = fi25
                                    _ = v26
                                }
                                v27 := GenS3{F0: (p0 + p0), F1: 1, F2: 42}
                                _ = v27
                            }
                        default:
                            fmt.Println(p1, p0)
                        }
                    }
                    _ = v18
                }
            }
            _ = v15
        }
        v28 := GenS4{F0: "\n\t", F1: (p1 - p1)}
        _ = v28
    }
    p0 += (p0 - 183)
    _ = p0
    _ = p1
    return -1
}

func main() {
    v1 := []int{42}
    v2 := []int{-112, 128, 9223372036854775807, 0}
    var v3 uint32 = 994
    var v4 uint8 = 0
    v5 := make(map[string]bool)
    v29 := GenS5{F0: (!(-1 == 0)), F1: (v4 / ((v4) | 1)), F2: ((!(((-2.25 * float64(1.5)) + 0.0) > 3.14159)) || true)}
    switch "foo bar" {
    case "foo bar":
        for fi30 := 0; fi30 < 11; fi30++ {
            for fi31 := 0; fi31 < 16; fi31++ {
                v2 = append(v2, (fi31 & fi30))
                v4 -= v4
                _ = fi31
            }
            _ = fi30
        }
    case "hello":
        if (int64(v4) != 0) {
            _ = len(v5)
            fmt.Println(v3, v3)
        }
    case "a":
        switch "" {
        case "hello":
            switch "" {
            case "":
                v32 := GenS6{F0: 0, F1: 244, F2: uint8(v3)}
                _ = v32
            }
        case "a":
            _ = len(v5)
        }
    default:
        v33 := GenS7{F0: 100}
        _ = v33
    }
    delete(v5, "foo bar")
    fmt.Println(v4, v3)
    if ("\n\t" > "") {
        if false {
            if len(v2) > 0 {
                v2[0] = 128
            }
            fmt.Println(v3)
        }
        if false {
            _ = len(v5)
        } else {
            if false {
                switch -2147483648 {
                case -2147483648:
                    for fi34 := 0; fi34 < 2; fi34++ {
                        switch "" {
                        case "":
                            var v35 uint32 = 745
                            _ = v35
                        }
                        _ = fi34
                    }
                case 42:
                    v3 += (v3 & (v3 + v3))
                }
                v36 := GenS5{F0: false, F1: v4, F2: true}
                _ = v36
            }
        }
    }
    delete(v5, "foo bar")
    switch 127 {
    case 848:
        switch "hello" {
        case "":
            var v37 int64 = -398
            _ = v37
        case "hello":
            v3 = 2147483648
        default:
            fmt.Println(v3, v4)
        }
    default:
        switch "foo bar" {
        case "hello":
            v4 -= v4
        case "":
            v4 += 2
        default:
            var v38 int32 = genFn0(2, v4)
            _ = v38
        }
    }
    var v39 bool = false
    _ = len(v1)
    for fi40 := 0; fi40 < 1; fi40++ {
        v3 -= (v3 << ((113) & 63))
        _ = fi40
    }
    var v41 uint = 0
    v42 := v5["foo bar"]
    if len(v1) > 0 {
        v1[0] = 32768
    }
    fmt.Println(v39)
    delete(v5, "a")
    v3 += 2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
