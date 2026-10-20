package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int; F1 uint32; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint }
type GenS5 struct { F0 bool; F1 int64 }
func genFn0(p0 uint16) uint64 {
    v5 := GenS1{F0: 2}
    p0 = (p0 + p0)
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]bool)
    var v3 uint16 = 65535
    v4 := make(map[string]string)
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int8 = -26
        switch "\n\t" {
        case "hello":
            v1 -= v1
        }
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 2
    var v9 uint = 1
    v1 = (v1 + (v1 % (((v1 + 255)) | 1)))
    switch 127 {
    case -18:
        switch ((("" + "\n\t") + "\n\t") + "\n\t") {
        case "hello":
            v10 := GenS1{F0: uint8(v3)}
            _ = v10
        case "a":
            var v11 int32 = 2147483647
            _ = v11
        }
    case 32767:
        switch ("\n\t" + "foo bar") {
        case "":
            var v12 uint64 = 1
            _ = v12
        case "a":
            var v13 string = "foo bar"
            _ = v13
        case "\n\t":
            if false {
                switch -9223372036854775808 {
                case -9223372036854775808:
                    switch "\n\t" {
                    case "":
                        v9 += (v9 % ((v9) | 1))
                    case "a":
                        if ((float64(-2.25) / 1.0) < ((0.0 / 1.0) * 1e10)) {
                            for fi14 := 0; fi14 < 15; fi14++ {
                                switch "a" {
                                case "a":
                                    for fi15 := 0; fi15 < 4; fi15++ {
                                        fmt.Println(v1, v9, v9)
                                        _ = fi15
                                    }
                                default:
                                    if true {
                                        switch (fi14 % (((fi14 + fi14)) | 1)) {
                                        case 4294967295:
                                            if true {
                                                switch fi14 {
                                                case 652:
                                                    var v16 uint = 18446744073709551615
                                                    _ = v16
                                                default:
                                                    fmt.Println(v9, v9)
                                                }
                                                v1 = v1
                                            }
                                        case 610:
                                            var v17 uint = (v9 ^ 1)
                                            _ = v17
                                        case 128:
                                            v18 := GenS2{F0: uint(v3), F1: (^(v8)), F2: v8}
                                            _ = v18
                                        default:
                                            v3 = v3
                                        }
                                    }
                                }
                                _ = fi14
                            }
                        }
                    case "hello":
                        v1 += (v1 / ((v1) | 1))
                    default:
                        switch "hello" {
                        case "":
                            for fi19 := 0; fi19 < 11; fi19++ {
                                if true {
                                    v20 := []int{1, fi19, (fi19 & 32767), fi19}
                                    fmt.Println(v3, v8)
                                    _ = v20
                                } else {
                                    v21 := GenS2{F0: 255, F1: uint32(v3), F2: (v8 ^ v8)}
                                    _ = v21
                                }
                                if (!("hello" > "a")) {
                                    v8 = v8
                                } else {
                                    v22 := GenS3{F0: fi19, F1: (v8 / ((v8) | 1)), F2: 821}
                                    if false {
                                        var v23 uint32 = (v8 + v8)
                                        _ = v23
                                    } else {
                                        if false {
                                            switch "hello" {
                                            case "foo bar":
                                                _ = len(v4)
                                            case "":
                                                fmt.Println(v8)
                                            case "\n\t":
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v24 int32 = int32(v1)
                                                    _ = v24
                                                case "hello":
                                                    var v25 uint64 = 1
                                                    _ = v25
                                                case "foo bar":
                                                    delete(v2, fi19)
                                                default:
                                                    fmt.Println(v9, v8)
                                                }
                                            default:
                                                v26 := GenS4{F0: (!false), F1: (v1 + v1), F2: v9}
                                                _ = v26
                                            }
                                            v4["foo bar"] = "hello"
                                        }
                                        fmt.Println(v3)
                                        fmt.Println(fi19)
                                    }
                                    _ = v22
                                }
                                if (!((v3 & uint16(v8)) > uint16(v9))) {
                                    var v27 uint8 = v1
                                    v8 += v8
                                    _ = v27
                                }
                                _ = fi19
                            }
                        default:
                            if true {
                                fmt.Println(v1, v9, v8)
                                delete(v2, 419)
                            }
                        }
                    }
                default:
                    fmt.Println(v8, v9)
                }
                v1 += (v1 * (v1 | (v1 % ((v1) | 1))))
            } else {
                var v28 int64 = 42
                v9 -= v9
                _ = v28
            }
        default:
            v4[(("foo bar" + "a") + ((("" + "\n\t") + "a") + ("a" + "hello")))] = ""
        }
    case -9223372036854775808:
        var v29 int8 = -48
        _ = v29
    }
    if false {
        var v30 int8 = 100
        for fi31 := 0; fi31 < 6; fi31++ {
            if true {
                v32 := []int{1000, (fi31 * (-(fi31))), fi31}
                _ = v32
            } else {
                var v33 uint = v9
                v34 := v2[(fi31 * (fi31 - (fi31 | (^(fi31)))))]
                _ = v33
                _ = v34
            }
            if false {
                v9 = 42
            }
            _ = fi31
        }
        _ = v30
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        if (false || (3.14159 >= 1.5)) {
            var v36 uint32 = 1
            switch fi35 {
            case 254:
                if ((127 < 127) && (true || true)) {
                    if false {
                        fmt.Println(v36, v3, v3)
                    }
                    _ = len(v4)
                }
            case -9223372036854775808:
                var v37 uint16 = v3
                _ = v37
            case 32768:
                var v38 uint = v9
                _ = v38
            }
            _ = v36
        }
        v39 := GenS4{F0: (true && ((128 > v1) || false)), F1: (v1 % ((v1) | 1)), F2: (v9 << ((v9) & 63))}
        for fi40 := 0; fi40 < 8; fi40++ {
            v41 := GenS1{F0: 255}
            _ = fi40
            _ = v41
        }
        _ = fi35
        _ = v39
    }
    var v42 int16 = 32767
    for fi43 := 0; fi43 < 16; fi43++ {
        var v44 int16 = v42
        v44 = v44
        _ = fi43
        _ = v44
    }
    v9 = v9
    var v45 int64 = 0
    fmt.Println(v8)
    v45 += int64(v1)
    if (719 < 515) {
        v46 := []int{}
        var v47 float32 = ((0.0 * 0.0) - ((0.0 + 1.5) / 3.14159))
        _ = v46
        _ = v47
    }
    var v48 float32 = 0.1
    switch "hello" {
    case "":
        v49 := GenS5{F0: (false && true), F1: -550}
        _ = v49
    case "a":
        v2[66] = ((true || true) && false)
    case "hello":
        v9 += v9
    default:
        var v50 uint8 = v1
        _ = v50
    }
    var v51 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v51:", v51)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v45
    _ = v48
    _ = v51
    _ = v8
    _ = v9
}
