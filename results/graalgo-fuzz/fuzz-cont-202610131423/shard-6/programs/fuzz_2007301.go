package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int16; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 bool; F1 uint; F2 uint64 }
func genFn0(p0 int) uint64 {
    v4 := GenS1{F0: "hello", F1: (841 > 1)}
    v5 := GenS2{F0: 323}
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(p0)
        _ = fi6
    }
    _ = p0
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 uint, p1 uint32) uint8 {
    p0 = (^(p0))
    var v7 string = "a"
    fmt.Println(p0, p1, p0)
    var v8 uint8 = 175
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint16 = 642
    v2 := make(map[string]uint32)
    var v3 float32 = 0.0
    if true {
        v1 += v1
        var v9 uint64 = 0
        fmt.Println(v9, v3, v9)
        _ = v9
    }
    v10 := []int{}
    var v11 int = 297
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := v2["a"]
        switch fi12 {
        case -526:
            switch v11 {
            case -743:
                v14 := GenS2{F0: 573}
                _ = v14
            case 42:
                fmt.Println(v3)
            case -9223372036854775808:
                var v15 int32 = 42
                _ = v15
            default:
                v3 -= 0.1
            }
        }
        var v16 int64 = -9223372036854775808
        _ = fi12
        _ = v13
        _ = v16
    }
    var v17 int32 = 1
    for fi18 := 0; fi18 < 11; fi18++ {
        v19 := GenS1{F0: "", F1: false}
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 int8 = 0
            _ = fi20
            _ = v21
        }
        var v22 int16 = -1
        _ = fi18
        _ = v19
        _ = v22
    }
    v1 += (v1 * 363)
    v23 := GenS1{F0: "hello", F1: (true || ((v1 << ((42) & 63)) == v1))}
    switch v11 {
    case 0:
        if false {
            switch v11 {
            case -832:
                switch (v11 / ((v11) | 1)) {
                case 436:
                    v1 -= v1
                }
            case 905:
                fmt.Println(v3, v3, v11)
            }
            v24 := GenS1{F0: "hello", F1: false}
            _ = v24
        }
    case 9223372036854775807:
        var v25 string = "hello"
        _ = v25
    case -748:
        if false {
            if false {
                if (!true) {
                    delete(v2, "foo bar")
                    if (2 > 4294967295) {
                        v3 += v3
                    } else {
                        v2[((("" + "hello") + ("" + "foo bar")) + "a")] = 1
                        v17 = (v17 * -362)
                    }
                } else {
                    for fi26 := 0; fi26 < 18; fi26++ {
                        v27 := GenS3{F0: -760, F1: 1, F2: -529}
                        var v28 int32 = v17
                        var v29 int8 = -1
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    var v30 uint16 = (v1 % (((v1 / ((65535) | 1))) | 1))
                    _ = v30
                }
                var v31 uint64 = 173
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = v31
            } else {
                var v32 uint16 = 504
                v33 := GenS4{F0: 533, F1: 268}
                _ = v32
                _ = v33
            }
            var v34 int32 = v17
            v34 += -2147483648
            _ = v34
        } else {
            for fi35 := 0; fi35 < 20; fi35++ {
                var v36 uint16 = v1
                if true {
                    fmt.Println(v36, v36)
                } else {
                    if true {
                        v36 = v36
                        if false {
                            switch fi35 {
                            case 128:
                                v17 += -1
                            case -195:
                                v36 = 1
                            default:
                                fmt.Println(v17, v3, v1)
                            }
                            v37 := GenS5{F0: v17, F1: -1}
                            v38 := GenS6{F0: (9223372036854775807 != 381), F1: 18446744073709551615, F2: 4294967295}
                            _ = v37
                            _ = v38
                        } else {
                            for fi39 := 0; fi39 < 18; fi39++ {
                                if true {
                                    var v40 float32 = v3
                                    switch "hello" {
                                    case "a":
                                        v41 := GenS4{F0: 0, F1: 1}
                                        _ = v41
                                    case "":
                                        v42 := GenS5{F0: v17, F1: -128}
                                        _ = v42
                                    }
                                    var v43 float32 = -2.25
                                    _ = v40
                                    _ = v43
                                } else {
                                    fmt.Println(v11, v17)
                                    var v44 int8 = 47
                                    var v45 uint16 = (v1 | (v1 ^ (v36 & v1)))
                                    _ = v44
                                    _ = v45
                                }
                                fmt.Println(fi39, fi39, v36)
                                v46 := GenS6{F0: false, F1: 313, F2: 42}
                                _ = fi39
                                _ = v46
                            }
                        }
                    }
                }
                for fi47 := 0; fi47 < 15; fi47++ {
                    v1 = v36
                    v3 -= v3
                    v11 -= (v11 >> ((0) & 63))
                    _ = fi47
                }
                _ = fi35
                _ = v36
            }
        }
    default:
        fmt.Println(v3)
    }
    var v48 int = v11
    v11 = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v48
}
