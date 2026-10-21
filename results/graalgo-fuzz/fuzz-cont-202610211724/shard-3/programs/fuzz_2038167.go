package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int32; F2 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 int32; F2 uint64 }
type GenS6 struct { F0 int16; F1 uint16; F2 uint16 }

func main() {
    var v1 int16 = 1
    var v2 float32 = 0.1
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    v5 := make(map[int]int)
    v1 -= v1
    fmt.Println(v1, v1)
    fmt.Println(v2)
    v6 := GenS1{F0: 4294967295, F1: 313, F2: 18446744073709551615}
    delete(v3, 42)
    if false {
        var v7 uint8 = 127
        v8 := GenS2{F0: 65535}
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 13; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                fmt.Println(v1, fi10)
                v11 := GenS1{F0: 533, F1: 776, F2: 18446744073709551615}
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    }
    v2 -= 3.14159
    if (0 == 1) {
        fmt.Println(v1, v1)
        if true {
            if (255 < 0) {
                var v12 int16 = v1
                _ = v12
            } else {
                var v13 uint = 208
                _ = v13
            }
        }
        switch 9223372036854775807 {
        case 89:
            var v14 int32 = -2147483648
            _ = v14
        case 42:
            var v15 uint32 = 2
            _ = v15
        default:
            if true {
                for fi16 := 0; fi16 < 20; fi16++ {
                    fmt.Println(v1)
                    v17 := GenS2{F0: 18446744073709551615}
                    _ = fi16
                    _ = v17
                }
                switch -898 {
                case 9223372036854775807:
                    v18 := GenS3{F0: 2, F1: int32(v1), F2: int32(v1)}
                    _ = v18
                case 32768:
                    if true {
                        v2 = v2
                    } else {
                        v19 := GenS1{F0: 781, F1: 0, F2: 65535}
                        v20 := GenS1{F0: 18446744073709551615, F1: 659, F2: 1}
                        var v21 uint = 42
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                case -1:
                    v1 += v1
                }
                switch -997 {
                case -1:
                    v22 := GenS2{F0: 1}
                    _ = v22
                case -2147483648:
                    fmt.Println(v1, v2)
                case -897:
                    v23 := GenS2{F0: 255}
                    _ = v23
                default:
                    v1 = v1
                }
            } else {
                if false {
                    switch -1 {
                    case 32767:
                        if (0 <= -935) {
                            var v24 float32 = (0.0 - v2)
                            _ = v24
                        } else {
                            v1 = v1
                            for fi25 := 0; fi25 < 9; fi25++ {
                                if false {
                                    v26 := v5[-23]
                                    v27 := GenS1{F0: 391, F1: 342, F2: 820}
                                    _ = v26
                                    _ = v27
                                }
                                _ = fi25
                            }
                        }
                    case -2147483648:
                        v28 := GenS2{F0: 2}
                        _ = v28
                    default:
                        delete(v3, 2147483648)
                    }
                    var v29 uint8 = 255
                    _ = v29
                }
            }
        }
    }
    v30 := []int{-9223372036854775808, 32768}
    for fi31 := 0; fi31 < 3; fi31++ {
        if false {
            switch (fi31 - 4294967295) {
            case 4294967295:
                fmt.Println(fi31)
            default:
                v1 += v1
            }
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            switch "a" {
            case "":
                if (3.14159 == (v2 - v2)) {
                    _ = len(v5)
                    for fi32 := 0; fi32 < 12; fi32++ {
                        var v33 int = fi32
                        if (255 > 296) {
                            v2 += v2
                            switch (fi32 * 1) {
                            case 32768:
                                v33 = (fi32 << ((42) & 63))
                            default:
                                v5[(-(fi31))] = (^(v33))
                            }
                        }
                        _ = fi32
                        _ = v33
                    }
                } else {
                    switch "\n\t" {
                    case "a":
                        if true {
                            v2 = ((v2 / ((v2 * (v2 + float32((0.0 * float64(-2.25))))) - -2.25)) + v2)
                        } else {
                            if false {
                                switch "\n\t" {
                                case "a":
                                    v1 -= (^(v1))
                                case "foo bar":
                                    v34 := GenS4{F0: 128}
                                    _ = v34
                                default:
                                    var v35 bool = true
                                    _ = v35
                                }
                                v1 -= v1
                            }
                            fmt.Println(v2)
                        }
                    case "\n\t":
                        var v36 int64 = -1
                        _ = v36
                    case "foo bar":
                        var v37 int16 = (^(v1))
                        _ = v37
                    default:
                        _ = len(v30)
                    }
                    var v38 float64 = 1.5
                    _ = v38
                }
            default:
                v2 += v2
            }
        } else {
            _ = len(v30)
            switch "hello" {
            case "foo bar":
                v1 -= 32767
            case "a":
                v39 := GenS5{F0: "hello", F1: -2147483648, F2: 1}
                _ = v39
            case "":
                delete(v4, fi31)
            }
            switch fi31 {
            case 308:
                var v40 int32 = int32(fi31)
                _ = v40
            case 4294967295:
                var v41 int8 = -105
                _ = v41
            case -385:
                fmt.Println(v2, v1)
            default:
                var v42 uint = 42
                _ = v42
            }
        }
        for fi43 := 0; fi43 < 19; fi43++ {
            if (!false) {
                if true {
                    v1 = v1
                } else {
                    var v44 int = (fi31 + (fi31 / (((^(fi43))) | 1)))
                    v2 += v2
                    v45 := GenS6{F0: v1, F1: 691, F2: 255}
                    _ = v44
                    _ = v45
                }
            } else {
                fmt.Println(v2, fi31)
            }
            if false {
                v46 := GenS2{F0: 1}
                _ = v46
            }
            v1 = -1
            _ = fi43
        }
        v47 := GenS6{F0: (^(v1)), F1: 42, F2: 2}
        _ = fi31
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
