package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int; F1 int64; F2 uint32 }
type GenS11 struct { F0 bool; F1 uint8; F2 uint32 }
func genFn0(p0 bool, p1 uint16, p2 uint) uint16 {
    var v4 uint16 = p1
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS1{F0: 204, F1: 2147483648, F2: -1}
        var v7 uint8 = 255
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return 32767
}
func genFn1() uint32 {
    if true {
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 uint8 = 127
            _ = fi9
            _ = v10
        }
        switch 4294967295 {
        case 1:
            if (95 < -745) {
                v11 := GenS2{F0: 539}
                _ = v11
            } else {
                var v12 int8 = -128
                _ = v12
            }
        case 513:
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := []int{}
                _ = fi13
                _ = v14
            }
        case 42:
            var v15 uint32 = 2
            _ = v15
        default:
            if (127 >= 128) {
                var v16 int8 = -128
                _ = v16
            }
        }
    } else {
        var v17 int64 = 1
        _ = v17
    }
    v18 := []int{2147483647, 0, 32767}
    switch ("" + ("foo bar" + "a")) {
    case "\n\t":
        for fi19 := 0; fi19 < 20; fi19++ {
            fmt.Println(fi19)
            _ = fi19
        }
    case "a":
        var v20 int8 = 1
        _ = v20
    }
    _ = v18
    return 4294967295
}

func main() {
    v1 := []string{}
    var v2 uint64 = 4294967295
    var v3 uint8 = 1
    v3 += v3
    var v21 int64 = -392
    for fi22 := 0; fi22 < 5; fi22++ {
        if true {
            v3 += 255
        }
        v23 := GenS3{F0: -145}
        for fi24 := 0; fi24 < 1; fi24++ {
            if len(v1) > 0 {
                v1[0] = ""
            }
            switch "\n\t" {
            case "":
                v2 += v2
            case "\n\t":
                if ((v3 & v3) == v3) {
                    v25 := GenS4{F0: 100}
                    if ((v2 * 2) == 2) {
                        switch "a" {
                        case "foo bar":
                            if ("\n\t" == "a") {
                                v26 := GenS5{F0: 878}
                                var v27 uint16 = uint16(v2)
                                var v28 uint8 = 140
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                        }
                    } else {
                        if true {
                            if true {
                                fmt.Println(v21)
                                v29 := GenS1{F0: 127, F1: 65535, F2: v21}
                                _ = v29
                            }
                            v1 = append(v1, "")
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                            if (!(false || false)) {
                                if (0.1 == (1.5 - 3.14159)) {
                                    v2 -= 4294967295
                                } else {
                                    v21 += (v21 & v21)
                                    v30 := GenS6{F0: fi24, F1: 32767}
                                    v2 += 42
                                    fmt.Println(v2)
                                    var v31 int8 = -128
                                    _ = v30
                                    _ = v31
                                }
                                var v32 uint = 2
                                var v33 float32 = float32((1e10 / (float64(1.5) / (1.5 - ((1e10 + 1e10) + (3.14159 * 1e10))))))
                                _ = v32
                                _ = v33
                            }
                        } else {
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            var v34 int8 = -1
                            _ = v34
                        }
                    }
                    _ = v25
                } else {
                    v35 := GenS7{F0: 123, F1: 42}
                    if false {
                        fmt.Println(v21, v21)
                    }
                    var v36 int64 = (v21 << ((293) & 63))
                    _ = v35
                    _ = v36
                }
            default:
                if true {
                    if true {
                        v37 := GenS8{F0: v2}
                        fmt.Println(v3, v2)
                        if true {
                            v3 -= v3
                        }
                        if ((1 <= -83) && (!true)) {
                            var v38 int32 = int32(v21)
                            _ = v38
                        }
                        _ = v37
                    } else {
                        if (!true) {
                            var v39 bool = true
                            _ = v39
                        } else {
                            _ = len(v1)
                        }
                    }
                    if (0 == 0) {
                        if len(v1) > 0 {
                            v1[0] = ((("\n\t" + "") + "") + "")
                        }
                        v40 := GenS8{F0: (v2 >> ((541) & 63))}
                        fmt.Println(fi22, fi24, v21)
                        _ = v40
                    } else {
                        var v41 int64 = -881
                        _ = v41
                    }
                    v2 = 42
                }
            }
            _ = fi24
        }
        _ = fi22
        _ = v23
    }
    if false {
        if false {
            if (((float32((1e10 / -2.25)) / 1.0) - (0.0 / -2.25)) != 1e10) {
                v21 = v21
                v3 = (v3 >> ((1) & 63))
            }
            var v42 string = "foo bar"
            if (true || true) {
                if false {
                    var v43 int16 = -738
                    v44 := GenS9{F0: 42}
                    _ = v43
                    _ = v44
                }
                switch v42 {
                case "\n\t":
                    if false {
                        if (42 <= (v21 % ((v21) | 1))) {
                            v3 -= 211
                            fmt.Println(v2, v42, v3)
                            switch v42 {
                            case "a":
                                var v45 string = ""
                                _ = v45
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = v42
                        }
                    } else {
                        if true {
                            var v46 float32 = (1.5 + float32(0.0))
                            _ = v46
                        }
                        v42 = v42
                    }
                }
            }
            _ = v42
        } else {
            v21 -= v21
        }
        if true {
            for fi47 := 0; fi47 < 7; fi47++ {
                fmt.Println(fi47)
                switch "hello" {
                case "a":
                    v2 = 2
                case "\n\t":
                    for fi48 := 0; fi48 < 9; fi48++ {
                        v3 += (v3 + (v3 + v3))
                        _ = fi48
                    }
                }
                v49 := GenS10{F0: -1, F1: v21, F2: 4294967295}
                _ = fi47
                _ = v49
            }
            v50 := GenS11{F0: true, F1: (v3 << ((65535) & 63)), F2: genFn1()}
            v3 = v3
            _ = v50
        }
    }
    v1 = append(v1, ("hello" + ""))
    v51 := GenS3{F0: 42}
    for fi52 := 0; fi52 < 11; fi52++ {
        var v53 int64 = v21
        v2 += (v2 - (^(v2)))
        v3 += v3
        _ = fi52
        _ = v53
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v51
}
