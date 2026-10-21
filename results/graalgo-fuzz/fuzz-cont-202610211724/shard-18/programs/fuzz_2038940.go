package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 int8; F2 int8 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS9 struct { F0 uint16; F1 uint32 }
type GenS10 struct { F0 int; F1 int32; F2 int64 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(fi6)
        switch (fi6 - (fi6 * fi6)) {
        case 463:
            fmt.Println(fi6)
        }
        _ = fi6
    }
    var v7 uint16 = 2
    var v8 uint32 = 42
    v8 = v8
    fmt.Println(v7)
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int16, p1 uint, p2 int) uint8 {
    var v9 uint64 = 413
    switch -2147483648 {
    case 127:
        fmt.Println(v9, p2, p2)
    default:
        v10 := GenS1{F0: ((^(p1)) == p1), F1: genFn0()}
        _ = v10
    }
    var v11 float64 = -2.25
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint = p1
            var v15 uint = (v14 * 65535)
            var v16 int8 = 1
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v9
    return 0
}

func main() {
    v1 := []int{289, 1, 32768}
    var v2 int64 = -1
    var v3 uint16 = 0
    v4 := []uint32{0, 5, 2147483648, 65535}
    v5 := make(map[uint32]string)
    v17 := GenS2{F0: 65535}
    v2 = 0
    var v18 int = -19
    v3 = (^(v3))
    switch 127 {
    case 32768:
        var v19 int = (v18 ^ (v18 / ((v18) | 1)))
        _ = v19
    case 294:
        for fi20 := 0; fi20 < 4; fi20++ {
            if (!true) {
                v21 := GenS1{F0: true, F1: 0}
                switch "a" {
                case "foo bar":
                    v3 = 2
                }
                fmt.Println(v18)
                _ = v21
            } else {
                delete(v5, 4294967295)
                for fi22 := 0; fi22 < 2; fi22++ {
                    if false {
                        v23 := GenS1{F0: (!(42 <= 127)), F1: -1}
                        var v24 int64 = -428
                        v25 := GenS3{F0: 42, F1: -9223372036854775808}
                        _ = v23
                        _ = v24
                        _ = v25
                    } else {
                        var v26 int32 = 0
                        v27 := GenS3{F0: 128, F1: -1}
                        _ = v26
                        _ = v27
                    }
                    fmt.Println(v2, fi20)
                    _ = fi22
                }
            }
            _ = fi20
        }
    case -2147483648:
        for fi28 := 0; fi28 < 13; fi28++ {
            v29 := GenS3{F0: 42, F1: 42}
            var v30 float64 = (3.14159 - ((1.5 - ((((((-2.25 * 0.0) / (1.5 + 0.0)) / 1.5) / float64(1e10)) / -2.25) - float64(float32(1.5)))) + (1.5 * (3.14159 * ((1.5 + 1e10) * (1.5 - (1e10 + float64(3.14159))))))))
            switch (fi28 / ((v18) | 1)) {
            case 800:
                if false {
                    v31 := GenS4{F0: 42}
                    _ = v31
                }
            case -9223372036854775808:
                if (!(!(("\n\t" + "a") > "\n\t"))) {
                    v32 := GenS5{F0: true, F1: 1, F2: -1}
                    for fi33 := 0; fi33 < 3; fi33++ {
                        v34 := GenS6{F0: genFn1(1, 0, v18), F1: -128}
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                } else {
                    var v35 string = (("foo bar" + "a") + "a")
                    v36 := GenS7{F0: (v3 >> ((0) & 63))}
                    _ = v35
                    _ = v36
                }
            case 536:
                for fi37 := 0; fi37 < 8; fi37++ {
                    switch 4294967295 {
                    case 2147483647:
                        switch "\n\t" {
                        case "":
                            v5[1] = (("\n\t" + "\n\t") + "")
                        case "\n\t":
                            if true {
                                v18 = -499
                            }
                        case "a":
                            var v38 int = 127
                            _ = v38
                        default:
                            v39 := GenS1{F0: (true && (!false)), F1: genFn0()}
                            _ = v39
                        }
                        if (!(!(3.14159 != 0.0))) {
                            v40 := GenS8{F0: 1, F1: 47, F2: 1}
                            v18 = -481
                            v5[0] = (("foo bar" + "hello") + "a")
                            _ = v40
                        } else {
                            fmt.Println(v18, v3, v18)
                        }
                    }
                    v41 := GenS9{F0: (v3 << ((918) & 63)), F1: 4294967295}
                    _ = fi37
                    _ = v41
                }
            }
            _ = fi28
            _ = v29
            _ = v30
        }
    default:
        fmt.Println(v18, v2)
    }
    var v42 int32 = -1
    v43 := v5[42]
    var v44 uint16 = 2
    fmt.Println(v2)
    v1 = append(v1, (v18 << ((405) & 63)))
    v44 -= v44
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = 957
    if false {
        _ = len(v5)
    }
    for fi45 := 0; fi45 < 17; fi45++ {
        for fi46 := 0; fi46 < 7; fi46++ {
            fmt.Println(v43, v44, v42)
            fmt.Println(v44, v2, v2)
            if ((true && true) && ((v42 ^ v42) <= 797)) {
                var v47 int32 = v42
                switch v18 {
                case 4294967295:
                    v48 := GenS7{F0: (^(v3))}
                    _ = v48
                case 42:
                    var v49 uint32 = 2
                    _ = v49
                case 792:
                    if ((-2.25 + 3.14159) >= 1e10) {
                        fmt.Println(fi45, v47)
                        var v50 float32 = (0.0 + (0.0 - 1e10))
                        _ = v50
                    } else {
                        fmt.Println(v44)
                        v44 = (v44 << ((4294967295) & 63))
                        var v51 int8 = 41
                        _ = v51
                    }
                default:
                    v52 := GenS10{F0: int(v42), F1: v42, F2: 1}
                    _ = v52
                }
                v3 += 706
                _ = v47
            }
            _ = fi46
        }
        v44 = (^(v3))
        _ = fi45
    }
    var v53 int = (v18 * int(v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v53
}
