package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int32; F2 uint64 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int; F1 int32; F2 uint8 }
func genFn0(p0 string) uint16 {
    v4 := GenS1{F0: 2, F1: 42}
    v5 := []int{42, 1, -106}
    if true {
        _ = len(v5)
        fmt.Println(p0)
    }
    if false {
        p0 = (p0 + (((p0 + p0) + ((p0 + "a") + "foo bar")) + (p0 + ((p0 + p0) + p0))))
        if ((!true) || false) {
            var v6 int32 = -1
            var v7 string = p0
            _ = v6
            _ = v7
        } else {
            switch "\n\t" {
            case "a":
                var v8 float64 = -2.25
                _ = v8
            case "hello":
                var v9 int32 = -1
                _ = v9
            case "":
                if (2147483647 == 68) {
                    _ = len(v5)
                    if (128 == 645) {
                        p0 = (("\n\t" + (p0 + "\n\t")) + "\n\t")
                        if (!false) {
                            if len(v5) > 0 {
                                v5[0] = 127
                            }
                        }
                        for fi10 := 0; fi10 < 1; fi10++ {
                            v11 := GenS1{F0: 255, F1: 18446744073709551615}
                            if (true || (!false)) {
                                p0 = (p0 + p0)
                                fmt.Println(fi10, fi10)
                            } else {
                                v5 = append(v5, (fi10 / (((fi10 + fi10)) | 1)))
                                var v12 int = fi10
                                p0 = p0
                                _ = v12
                            }
                            v13 := GenS2{F0: 0}
                            _ = fi10
                            _ = v11
                            _ = v13
                        }
                    }
                    for fi14 := 0; fi14 < 13; fi14++ {
                        p0 = ((p0 + p0) + (("" + p0) + ("" + ("\n\t" + "a"))))
                        p0 = p0
                        _ = fi14
                    }
                } else {
                    for fi15 := 0; fi15 < 14; fi15++ {
                        p0 = (("foo bar" + p0) + p0)
                        if (4294967295 >= 4294967295) {
                            fmt.Println(p0, fi15)
                            switch p0 {
                            case "hello":
                                v16 := GenS1{F0: 4294967295, F1: 40}
                                _ = v16
                            case "foo bar":
                                _ = len(v5)
                            case "":
                                var v17 uint32 = 49
                                _ = v17
                            }
                        }
                        var v18 uint16 = 65535
                        _ = fi15
                        _ = v18
                    }
                }
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1() uint64 {
    var v19 uint16 = 276
    for fi20 := 0; fi20 < 14; fi20++ {
        v21 := []int{fi20, int(v19)}
        switch fi20 {
        case -282:
            v19 = 255
        }
        v22 := GenS3{F0: fi20, F1: -1, F2: 1}
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        v19 -= (v19 - 0)
        _ = fi23
    }
    _ = v19
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{""}
    v3 := []int{}
    var v24 uint8 = 42
    var v25 uint32 = 2147483647
    var v26 int = 564
    for fi27 := 0; fi27 < 17; fi27++ {
        _ = len(v1)
        _ = fi27
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        if false {
            var v29 float32 = 1.5
            v3 = append(v3, fi28)
            _ = v29
        } else {
            var v30 int32 = -251
            if true {
                for fi31 := 0; fi31 < 16; fi31++ {
                    delete(v1, (v25 - v25))
                    _ = fi31
                }
                if (!false) {
                    v26 = 594
                } else {
                    if true {
                        fmt.Println(v26, v30, v25)
                        var v32 uint64 = 1
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 9; fi33++ {
                        v25 -= (v25 >> ((42) & 63))
                        var v34 int = fi33
                        v26 = 112
                        _ = fi33
                        _ = v34
                    }
                    if false {
                        fmt.Println(v30, v24, v24)
                    }
                }
                for fi35 := 0; fi35 < 7; fi35++ {
                    if (("\n\t" + "hello") < "\n\t") {
                        v3 = append(v3, (v26 / (((fi28 - 32768)) | 1)))
                        v36 := GenS4{F0: 127, F1: 0}
                        var v37 string = ("hello" + (("foo bar" + "hello") + "hello"))
                        switch v37 {
                        case "foo bar":
                            if (!false) {
                                v38 := GenS5{F0: 42}
                                _ = v38
                            }
                        }
                        _ = v36
                        _ = v37
                    }
                    _ = fi35
                }
            }
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            _ = v30
        }
        _ = fi28
    }
    v39 := GenS2{F0: 0}
    v26 = v26
    v25 = 65535
    var v40 uint64 = 317
    delete(v1, v25)
    var v41 int16 = -32768
    v25 -= (v25 | v25)
    var v42 int32 = 42
    fmt.Println(v40)
    for fi43 := 0; fi43 < 15; fi43++ {
        v42 += v42
        switch "\n\t" {
        case "a":
            for fi44 := 0; fi44 < 7; fi44++ {
                v45 := GenS5{F0: (v40 | 42)}
                if true {
                    if true {
                        fmt.Println(fi43, fi43)
                        v46 := GenS6{F0: v42, F1: true}
                        if false {
                            var v47 uint32 = v25
                            _ = v47
                        }
                        delete(v1, 4294967295)
                        v40 -= (^(v40))
                        var v48 uint32 = 2147483647
                        v49 := GenS1{F0: 4294967295, F1: 2}
                        var v50 int64 = 42
                        if true {
                            switch "\n\t" {
                            case "":
                                v41 += v41
                                fmt.Println(v48, v40, v26)
                            case "hello":
                                v42 -= (-(v42))
                            }
                            switch ("a" + "hello") {
                            case "\n\t":
                                if true {
                                    v51 := GenS5{F0: (v40 >> ((uint(v40)) & 63))}
                                    fmt.Println(fi44)
                                    _ = v51
                                }
                            case "hello":
                                v50 += -1
                            }
                            var v52 float64 = -2.25
                            _ = v52
                        }
                        var v53 uint32 = (v48 & v48)
                        _ = v46
                        _ = v48
                        _ = v49
                        _ = v50
                        _ = v53
                    } else {
                        fmt.Println(v42)
                        _ = len(v2)
                    }
                } else {
                    v40 -= v40
                    v54 := v1[(v25 >> ((1) & 63))]
                    _ = v54
                }
                switch "" {
                case "foo bar":
                    if (false || true) {
                        v55 := GenS7{F0: uint32(v24)}
                        v41 -= -95
                        var v56 int32 = v42
                        _ = v55
                        _ = v56
                    }
                    var v57 uint = 497
                    _ = v57
                case "a":
                    v58 := GenS8{F0: fi44, F1: v42, F2: (v24 | v24)}
                    _ = v58
                }
                _ = fi44
                _ = v45
            }
        case "":
            for fi59 := 0; fi59 < 18; fi59++ {
                v41 += v41
                var v60 uint8 = (^(v24))
                v40 = v40
                v61 := GenS7{F0: (v25 & v25)}
                _ = len(v1)
                _ = fi59
                _ = v60
                _ = v61
            }
        default:
            fmt.Println(v41, v25)
        }
        var v62 int32 = v42
        _ = fi43
        _ = v62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
