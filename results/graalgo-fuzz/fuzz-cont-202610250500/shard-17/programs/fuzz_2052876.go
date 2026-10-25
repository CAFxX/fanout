package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint; F1 int16; F2 uint }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 uint8; F2 int }
type GenS9 struct { F0 uint }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 0
    var v3 string = ""
    var v4 bool = true
    var v5 uint = 774
    for fi6 := 0; fi6 < 1; fi6++ {
        v5 = 0
        _ = fi6
    }
    v7 := GenS1{F0: true, F1: 0, F2: 42}
    var v8 int32 = -1000
    switch 42 {
    case 42:
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int64 = -9223372036854775808
            _ = fi9
            _ = v10
        }
    case -432:
        var v11 int8 = 121
        _ = v11
    case 127:
        switch ("\n\t" + v3) {
        case "\n\t":
            v4 = (127 == 40)
        case "foo bar":
            switch (("" + v3) + "hello") {
            case "":
                v12 := GenS2{F0: 83, F1: ((v8 < v8) || true)}
                _ = v12
            case "a":
                var v13 uint = v2
                _ = v13
            case "foo bar":
                var v14 uint64 = 576
                _ = v14
            default:
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS3{F0: v2, F1: 42, F2: (v5 << ((v2) & 63))}
                    _ = fi15
                    _ = v16
                }
            }
        }
    default:
        for fi17 := 0; fi17 < 11; fi17++ {
            delete(v1, (fi17 >> (((v5 / ((255) | 1))) & 63)))
            _ = fi17
        }
    }
    if (42 > 140) {
        if v4 {
            v4 = (4294967295 != v5)
            v3 = v3
        }
    } else {
        v3 = v3
        var v18 int64 = 9223372036854775807
        _ = v18
    }
    if (89 <= 4294967295) {
        v8 = v8
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 uint = v5
            v2 += (^(v2))
            var v21 int16 = 42
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        switch 413 {
        case 32767:
            var v22 int64 = -1
            _ = v22
        }
    }
    v2 = v2
    for fi23 := 0; fi23 < 4; fi23++ {
        for fi24 := 0; fi24 < 12; fi24++ {
            delete(v1, -2147483648)
            _ = fi24
        }
        _ = fi23
    }
    if false {
        if v4 {
            var v25 uint32 = 2
            _ = v25
        } else {
            var v26 int64 = 0
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 uint8 = 2
                if (v4 || v4) {
                    if (true && v4) {
                        switch 128 {
                        case 4294967295:
                            fmt.Println(v26, v4)
                        }
                        for fi29 := 0; fi29 < 17; fi29++ {
                            v30 := GenS4{F0: ((v3 + ("foo bar" + v3)) + "foo bar"), F1: 42}
                            v31 := GenS5{F0: (-(v26))}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        v3 = ((v3 + (("\n\t" + "\n\t") + "\n\t")) + v3)
                    }
                    v32 := GenS6{F0: v2, F1: 65535}
                    _ = v32
                } else {
                    if (0.1 < float32(1.5)) {
                        if (v4 && (42 <= 1)) {
                            if v4 {
                                if v4 {
                                    for fi33 := 0; fi33 < 20; fi33++ {
                                        var v34 float32 = (1.5 - (1e10 + 0.0))
                                        v35 := []int{(fi27 * fi27), -2147483648, fi33, (fi27 / ((4294967295) | 1))}
                                        if (-1 == -41) {
                                            _ = len(v35)
                                        } else {
                                            var v36 uint32 = 2147483648
                                            switch fi33 {
                                            case 128:
                                                _ = len(v1)
                                            case 1:
                                                var v37 bool = (127 > 42)
                                                _ = v37
                                            case -1:
                                                v38 := GenS6{F0: v5, F1: 393}
                                                _ = v38
                                            default:
                                                var v39 uint8 = (v28 << ((v2) & 63))
                                                v36 = (v36 ^ v36)
                                                _ = v39
                                            }
                                            _ = v36
                                        }
                                        var v40 uint64 = 2
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                        _ = v40
                                    }
                                    v41 := []int{fi27, fi27}
                                    fmt.Println(v8)
                                    _ = v41
                                }
                            }
                            delete(v1, fi27)
                            if false {
                                if true {
                                    switch 9223372036854775807 {
                                    case -9223372036854775808:
                                        fmt.Println(v8, fi27, v4)
                                    case 0:
                                        v28 = v28
                                    default:
                                        for fi42 := 0; fi42 < 7; fi42++ {
                                            switch (fi42 - fi42) {
                                            case 680:
                                                v43 := GenS7{F0: 42}
                                                _ = v43
                                            default:
                                                fmt.Println(fi42, fi27)
                                            }
                                            _ = fi42
                                        }
                                    }
                                    if (v4 && (v4 && v4)) {
                                        v44 := GenS7{F0: -987}
                                        if v4 {
                                            v45 := []int{fi27, fi27}
                                            _ = v45
                                        } else {
                                            _ = len(v1)
                                        }
                                        _ = v44
                                    } else {
                                        v46 := GenS8{F0: 42, F1: (^(v28)), F2: -1}
                                        delete(v1, (fi27 << ((195) & 63)))
                                        v47 := GenS3{F0: (v5 + 42), F1: -290, F2: v2}
                                        _ = v46
                                        _ = v47
                                    }
                                } else {
                                    for fi48 := 0; fi48 < 9; fi48++ {
                                        switch "" {
                                        case "foo bar":
                                            v49 := GenS9{F0: 994}
                                            _ = v49
                                        default:
                                            fmt.Println(v4)
                                        }
                                        _ = fi48
                                    }
                                    for fi50 := 0; fi50 < 5; fi50++ {
                                        var v51 float32 = 3.14159
                                        _ = fi50
                                        _ = v51
                                    }
                                }
                            }
                        } else {
                            var v52 uint64 = 1
                            _ = v52
                        }
                    } else {
                        v53 := GenS1{F0: (4294967295 > 755), F1: 0, F2: (^(v8))}
                        var v54 uint16 = 751
                        fmt.Println(fi27, v54)
                        _ = v53
                        _ = v54
                    }
                    var v55 int32 = 1
                    _ = v55
                }
                _ = fi27
                _ = v28
            }
            v56 := GenS5{F0: (v26 / ((v26) | 1))}
            _ = v26
            _ = v56
        }
    } else {
        v57 := []int{}
        _ = v57
    }
    v58 := []int{128, -1, 32767, 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v7
    _ = v8
}
