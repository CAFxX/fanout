package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 int64; F2 int16 }
type GenS6 struct { F0 int64; F1 string; F2 int32 }
type GenS7 struct { F0 int32; F1 int32; F2 uint64 }
type GenS8 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS9 struct { F0 int8 }
func genFn0(p0 uint16, p1 uint64) int16 {
    v6 := GenS1{F0: 1000, F1: p0}
    for fi7 := 0; fi7 < 1; fi7++ {
        p1 = 272
        _ = fi7
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    var v4 string = "a"
    var v5 int8 = -128
    v5 = (v5 - 76)
    v8 := GenS2{F0: (23 == 1)}
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int32 = 42
        _ = fi9
        _ = v10
    }
    switch -392 {
    case 466:
        v5 += -1
    case 32767:
        for fi11 := 0; fi11 < 19; fi11++ {
            v12 := GenS3{F0: 127, F1: 255, F2: fi11}
            _ = fi11
            _ = v12
        }
    case 2147483647:
        v4 = (v4 + ("hello" + ((v4 + v4) + v4)))
    }
    v4 = v4
    delete(v1, v4)
    v13 := GenS2{F0: (1 == 4294967295)}
    if (v5 != (^(v5))) {
        if (362 > 401) {
            var v14 int = -9223372036854775808
            v15 := GenS4{F0: 42}
            _ = len(v2)
            _ = v14
            _ = v15
        }
    }
    switch v4 {
    case "a":
        if true {
            if (!(9223372036854775807 > 1)) {
                for fi16 := 0; fi16 < 19; fi16++ {
                    var v17 uint32 = 440
                    v17 += (^(v17))
                    var v18 int64 = 268
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                v19 := GenS5{F0: 2147483647, F1: -9223372036854775808, F2: 376}
                var v20 uint32 = 2
                _ = v19
                _ = v20
            }
            v21 := GenS6{F0: 1, F1: "foo bar", F2: -1}
            v22 := []int{}
            _ = v21
            _ = v22
        }
    case "\n\t":
        switch v4 {
        case "hello":
            v23 := []int{-2147483648, 1, 127, -1}
            _ = v23
        case "foo bar":
            if ((500 != 42) || true) {
                _ = len(v1)
                v5 = 1
            } else {
                if false {
                    if (953 != 42) {
                        var v24 int32 = 1
                        delete(v1, (v4 + v4))
                        switch ("foo bar" + v4) {
                        case "foo bar":
                            switch 9223372036854775807 {
                            case 127:
                                v5 = v5
                            case 32767:
                                for fi25 := 0; fi25 < 18; fi25++ {
                                    for fi26 := 0; fi26 < 11; fi26++ {
                                        var v27 float32 = 1e10
                                        _ = fi26
                                        _ = v27
                                    }
                                    delete(v3, fi25)
                                    delete(v2, v4)
                                    _ = fi25
                                }
                            case 0:
                                switch 1 {
                                case 2147483647:
                                    v24 += 0
                                case 42:
                                    v28 := GenS7{F0: v24, F1: v24, F2: 2}
                                    _ = v28
                                }
                            }
                        default:
                            var v29 float32 = (3.14159 * 1e10)
                            _ = v29
                        }
                        _ = v24
                    }
                    var v30 uint32 = 2147483647
                    if (255 >= 319) {
                        v31 := []int{127, -211, 908, 32768}
                        switch -988 {
                        case 127:
                            for fi32 := 0; fi32 < 17; fi32++ {
                                v3[fi32] = v4
                                v2[v4] = v4
                                _ = fi32
                            }
                        default:
                            fmt.Println(v5, v30, v5)
                        }
                        _ = v31
                    } else {
                        delete(v1, "\n\t")
                        var v33 float32 = -2.25
                        if false {
                            switch v4 {
                            case "":
                                if true {
                                    v30 -= 2147483647
                                }
                            case "\n\t":
                                fmt.Println(v33, v33, v33)
                            case "foo bar":
                                fmt.Println(v30, v5)
                            }
                            for fi34 := 0; fi34 < 17; fi34++ {
                                fmt.Println(v33)
                                v5 -= v5
                                _ = fi34
                            }
                            switch 4294967295 {
                            case -1:
                                fmt.Println(v5, v30)
                            case -962:
                                v1[v4] = 42
                            default:
                                v5 = v5
                            }
                        } else {
                            var v35 bool = true
                            switch ("" + ("a" + ((v4 + (v4 + v4)) + ("\n\t" + v4)))) {
                            case "foo bar":
                                var v36 uint64 = 1
                                _ = v36
                            case "":
                                var v37 uint8 = 1
                                _ = v37
                            case "a":
                                if (((!v35) && v35) || true) {
                                    var v38 uint64 = 4294967295
                                    v39 := v3[215]
                                    var v40 int64 = 0
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                }
                            default:
                                switch -2147483648 {
                                case 4294967295:
                                    if v35 {
                                        v41 := GenS2{F0: (1 < 42)}
                                        v42 := []int{2147483647, 4294967295}
                                        _ = v41
                                        _ = v42
                                    } else {
                                        if v35 {
                                            var v43 uint64 = 392
                                            _ = v43
                                        } else {
                                            fmt.Println(v35)
                                            v35 = true
                                            fmt.Println(v4, v33)
                                        }
                                        switch 4294967295 {
                                        case -9223372036854775808:
                                            v33 = (v33 * v33)
                                        case -350:
                                            v4 = "a"
                                        case 528:
                                            for fi44 := 0; fi44 < 5; fi44++ {
                                                v30 += v30
                                                _ = fi44
                                            }
                                        default:
                                            if (v4 > v4) {
                                                var v45 float64 = (0.1 * -2.25)
                                                v35 = v35
                                                v45 -= float64(float32(float64((v33 + ((v33 * (v33 + float32(3.14159))) - v33)))))
                                                _ = v45
                                            } else {
                                                var v46 int64 = -813
                                                v47 := GenS8{F0: 392, F1: 18446744073709551615, F2: 0}
                                                v48 := GenS9{F0: 42}
                                                _ = v46
                                                _ = v47
                                                _ = v48
                                            }
                                        }
                                    }
                                case 127:
                                    var v49 float32 = v33
                                    _ = v49
                                default:
                                    var v50 uint32 = (v30 + v30)
                                    _ = v50
                                }
                            }
                            var v51 int8 = (^(v5))
                            _ = v35
                            _ = v51
                        }
                        _ = v33
                    }
                    _ = v30
                } else {
                    v4 = v4
                }
            }
        case "a":
            v1["hello"] = 4294967295
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
