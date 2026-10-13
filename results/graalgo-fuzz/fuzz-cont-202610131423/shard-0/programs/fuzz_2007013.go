package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint16; F1 string; F2 uint }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 int32; F1 uint64; F2 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint; F1 int16; F2 uint16 }
func genFn0() uint64 {
    var v7 int = -853
    v8 := []int{(v7 << ((65535) & 63)), v7, v7}
    var v9 int64 = -308
    if (!false) {
        var v10 uint64 = 42
        var v11 float64 = (float64((float32(3.14159) + float32(0.1))) - 1.5)
        _ = v10
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int8 = 127
    v2 := make(map[string]uint32)
    var v3 int64 = 1
    var v4 float64 = 1e10
    var v5 uint8 = 2
    v6 := make(map[string]int)
    v12 := v2[("\n\t" + "")]
    _ = len(v6)
    v13 := GenS1{F0: v12}
    v14 := GenS1{F0: v12}
    var v15 uint16 = 490
    switch 42 {
    case 4294967295:
        var v16 int32 = -153
        _ = v16
    case 9223372036854775807:
        var v17 int8 = (v1 ^ 100)
        _ = v17
    case 0:
        var v18 float32 = 3.14159
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := GenS1{F0: (v12 | (v12 >> ((2) & 63)))}
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: "a", F1: (0.0 != 0.0)}
    if true {
        if (-261 <= int(v12)) {
            var v22 int64 = v3
            _ = v22
        } else {
            var v23 float32 = 1.5
            v1 = 42
            switch 809 {
            case 42:
                v23 -= ((v23 - (float32(v4) / float32(v4))) + (v23 * (0.1 + (v23 + v23))))
            case -254:
                if false {
                    fmt.Println(v23, v4, v12)
                } else {
                    if true {
                        if true {
                            var v24 int64 = (-(v3))
                            _ = v24
                        }
                        if false {
                            if true {
                                if false {
                                    v25 := GenS3{F0: (3.14159 >= float32((3.14159 * 1e10))), F1: (v15 >= v15)}
                                    fmt.Println(v12, v1, v15)
                                    _ = v25
                                }
                                v5 -= v5
                                fmt.Println(v3, v4, v23)
                            } else {
                                v26 := GenS4{F0: (v15 >> ((18446744073709551615) & 63)), F1: "a", F2: 0}
                                _ = v26
                            }
                            switch "\n\t" {
                            case "\n\t":
                                for fi27 := 0; fi27 < 15; fi27++ {
                                    if (2 < 944) {
                                        var v28 uint = 4294967295
                                        _ = v28
                                    }
                                    var v29 float64 = v4
                                    _ = fi27
                                    _ = v29
                                }
                            case "":
                                v1 += (v1 % (((v1 << ((42) & 63))) | 1))
                            case "foo bar":
                                v30 := []int{-2147483648}
                                _ = v30
                            default:
                                for fi31 := 0; fi31 < 3; fi31++ {
                                    var v32 int16 = 291
                                    _ = fi31
                                    _ = v32
                                }
                            }
                        } else {
                            v33 := GenS1{F0: v12}
                            v23 = float32(-2.25)
                            if ((809 < 753) || (!true)) {
                                var v34 int8 = (v1 << ((11) & 63))
                                fmt.Println(v34, v5)
                                var v35 int = 4294967295
                                _ = v34
                                _ = v35
                            } else {
                                v36 := GenS1{F0: v12}
                                var v37 uint = 2
                                v38 := []int{42}
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                            _ = v33
                        }
                        v39 := GenS4{F0: v15, F1: "foo bar", F2: 342}
                        _ = v39
                    } else {
                        v40 := GenS5{F0: (^(v15)), F1: (v1 << ((65535) & 63))}
                        for fi41 := 0; fi41 < 7; fi41++ {
                            v3 = 1
                            _ = fi41
                        }
                        _ = v40
                    }
                    var v42 int = 127
                    if (true || false) {
                        var v43 float64 = (v4 - float64(((float32(0.1) + 1e10) * (v23 / 1e10))))
                        var v44 uint8 = v5
                        _ = v43
                        _ = v44
                    }
                    _ = v42
                }
            case 4294967295:
                v45 := GenS5{F0: v15, F1: 126}
                _ = v45
            }
            _ = v23
        }
        switch int(v5) {
        case 13:
            var v46 int = 128
            _ = v46
        case 32768:
            var v47 bool = (1 <= v12)
            _ = v47
        case 0:
            for fi48 := 0; fi48 < 14; fi48++ {
                v49 := GenS6{F0: int32(fi48), F1: 42, F2: ("" + "hello")}
                _ = fi48
                _ = v49
            }
        default:
            if true {
                v50 := GenS7{F0: 0}
                if false {
                    if (true || true) {
                        if false {
                            var v51 int8 = (v1 - (v1 + v1))
                            var v52 uint64 = 4294967295
                            _ = v51
                            _ = v52
                        } else {
                            var v53 uint = 1
                            var v54 int64 = v3
                            _ = v53
                            _ = v54
                        }
                        var v55 uint16 = v15
                        v12 = uint32(v5)
                        _ = v55
                    } else {
                        var v56 uint = 249
                        v57 := []int{-9223372036854775808, 2147483647}
                        v56 = (v56 << ((65535) & 63))
                        _ = v56
                        _ = v57
                    }
                    v58 := GenS8{F0: 2, F1: -336, F2: 0}
                    _ = v58
                }
                if false {
                    var v59 string = ("foo bar" + "\n\t")
                    if (1 != 1) {
                        var v60 float64 = (v4 + 1e10)
                        _ = v60
                    }
                    _ = v59
                } else {
                    switch ("foo bar" + ("foo bar" + "")) {
                    case "a":
                        for fi61 := 0; fi61 < 2; fi61++ {
                            if false {
                                fmt.Println(v15)
                                v4 = v4
                                v1 -= (v1 * v1)
                            } else {
                                v62 := GenS3{F0: true, F1: false}
                                v3 = (v3 >> ((737) & 63))
                                _ = v62
                            }
                            _ = fi61
                        }
                    default:
                        var v63 int16 = 42
                        _ = v63
                    }
                    v15 -= (v15 >> ((65535) & 63))
                }
                _ = v50
            }
        }
    }
    var v64 int = 32767
    if (!(!(!true))) {
        v4 -= (1.5 * ((float64(0.0) + v4) - v4))
    } else {
        fmt.Println(v12)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
