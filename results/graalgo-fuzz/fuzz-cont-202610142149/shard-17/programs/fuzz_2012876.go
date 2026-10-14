package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 uint; F2 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 int32; F1 uint16; F2 int32 }
type GenS7 struct { F0 uint8; F1 int64; F2 int32 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 bool, p1 int) int32 {
    v7 := GenS1{F0: 2147483647, F1: p1, F2: "\n\t"}
    fmt.Println(p0, p1)
    fmt.Println(p1, p1)
    p0 = (((!true) || p0) || ("hello" >= "a"))
    v8 := GenS1{F0: 369, F1: p1, F2: ("foo bar" + ("a" + "\n\t"))}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 312
    var v3 uint16 = 122
    var v4 int32 = v2
    v5 := make(map[int]string)
    var v6 uint64 = 0
    var v9 bool = (true && (754 < 2))
    if v9 {
        if false {
            v6 -= (v6 / ((v6) | 1))
        } else {
            for fi10 := 0; fi10 < 12; fi10++ {
                delete(v5, (fi10 / ((-2147483648) | 1)))
                _ = fi10
            }
        }
        switch (((((((("" + (("a" + "") + "hello")) + "\n\t") + "foo bar") + "a") + "hello") + "") + ("" + "foo bar")) + "foo bar") {
        case "a":
            v9 = false
        case "foo bar":
            if (!(42 > 2147483647)) {
                if false {
                    v11 := []int{2147483647, 9223372036854775807, 128}
                    v11 = append(v11, 42)
                    _ = v11
                }
                v12 := GenS2{F0: ""}
                _ = v12
            }
        case "hello":
            var v13 float64 = 0.1
            _ = v13
        default:
            switch -9223372036854775808 {
            case 4294967295:
                v2 += (v2 + 1)
            case 9223372036854775807:
                delete(v1, 1)
            default:
                for fi14 := 0; fi14 < 6; fi14++ {
                    v9 = (!((254 >= 191) || (-503 >= fi14)))
                    _ = fi14
                }
            }
        }
    }
    v15 := GenS3{F0: 440, F1: 42, F2: 42}
    v6 -= 42
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 float32 = float32(((((float64(1.5) - 0.1) + 0.0) - ((1.5 + 1e10) / 0.1)) / 3.14159))
        _ = fi16
        _ = v17
    }
    if v9 {
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := v5[(fi18 & fi18)]
            _ = fi18
            _ = v19
        }
        v2 -= -388
        switch "hello" {
        case "foo bar":
            if v9 {
                switch 777 {
                case 579:
                    v20 := GenS4{F0: v4}
                    _ = v20
                case 390:
                    var v21 uint8 = 183
                    _ = v21
                }
                v22 := GenS5{F0: (v3 % ((65535) | 1)), F1: 1}
                _ = v22
            }
        default:
            if (!v9) {
                var v23 uint64 = (v6 + (v6 - (v6 % ((v6) | 1))))
                var v24 int64 = -9223372036854775808
                for fi25 := 0; fi25 < 3; fi25++ {
                    if (true || true) {
                        var v26 float64 = (float64(0.0) + ((((3.14159 + (1e10 - ((float64(float32(1e10)) - ((1.5 + ((1.5 + (float64((1e10 / 1.0)) * 0.1)) + 1e10)) - (1.5 - 0.0))) + 1e10))) * 0.0) + float64((1.5 - 1.5))) / -2.25))
                        var v27 string = ("a" + "")
                        var v28 int32 = 42
                        _ = v26
                        _ = v27
                        _ = v28
                    } else {
                        v29 := GenS5{F0: (v3 + 989), F1: 616}
                        v24 += 1
                        v30 := []int{fi25}
                        _ = v29
                        _ = v30
                    }
                    var v31 int64 = (v24 & v24)
                    switch "hello" {
                    case "a":
                        v32 := GenS1{F0: (-(v4)), F1: (fi25 >> ((433) & 63)), F2: "foo bar"}
                        _ = v32
                    case "foo bar":
                        v4 = genFn0((583 > 287), (^(fi25)))
                    default:
                        if (v9 && v9) {
                            switch (fi25 % (((^(fi25))) | 1)) {
                            case -913:
                                fmt.Println(v31)
                            case -74:
                                if v9 {
                                    v4 = (v4 % (((v4 & 2147483647)) | 1))
                                } else {
                                    var v33 uint64 = (v6 ^ 1)
                                    _ = v33
                                }
                            default:
                                switch "foo bar" {
                                case "hello":
                                    v34 := GenS6{F0: v2, F1: (v3 + (v3 ^ 42)), F2: (v2 + v4)}
                                    _ = v34
                                }
                            }
                        }
                    }
                    _ = fi25
                    _ = v31
                }
                _ = v23
                _ = v24
            } else {
                v35 := []int{9223372036854775807, 0}
                _ = v35
            }
        }
    }
    if ((false && true) || false) {
        delete(v1, 2147483647)
    } else {
        var v36 int64 = 1
        for fi37 := 0; fi37 < 11; fi37++ {
            var v38 bool = v9
            _ = fi37
            _ = v38
        }
        fmt.Println(v6, v9)
        _ = v36
    }
    var v39 uint16 = 1
    var v40 int16 = -308
    switch 128 {
    case -9223372036854775808:
        for fi41 := 0; fi41 < 18; fi41++ {
            var v42 uint64 = 163
            _ = fi41
            _ = v42
        }
    case -203:
        if true {
            v4 = genFn0(true, 127)
            var v43 uint = 2
            _ = v43
        } else {
            if (v9 && v9) {
                fmt.Println(v3)
                var v44 uint = 894
                _ = v44
            } else {
                var v45 int32 = (-(v4))
                _ = v45
            }
            v5[292] = ("a" + "a")
        }
    case 9223372036854775807:
        fmt.Println(v4)
    default:
        var v46 uint32 = 2147483648
        _ = v46
    }
    for fi47 := 0; fi47 < 11; fi47++ {
        for fi48 := 0; fi48 < 17; fi48++ {
            v9 = (0.1 == 1e10)
            v49 := GenS7{F0: 127, F1: 9223372036854775807, F2: v4}
            if true {
                v50 := GenS8{F0: (^(v3))}
                fmt.Println(v9, v4)
                v51 := []int{(fi48 ^ (-(fi48))), fi47, int(v39)}
                _ = v50
                _ = v51
            }
            _ = fi48
            _ = v49
        }
        switch fi47 {
        case 906:
            for fi52 := 0; fi52 < 6; fi52++ {
                switch "" {
                case "hello":
                    v53 := GenS3{F0: (fi47 ^ fi47), F1: 2, F2: v40}
                    _ = v53
                case "a":
                    var v54 int64 = -1
                    _ = v54
                case "":
                    v39 += v3
                }
                var v55 bool = (v40 > -632)
                fmt.Println(v40, v4, v55)
                _ = fi52
                _ = v55
            }
        case 1:
            fmt.Println(v6, v2, v9)
        }
        _ = fi47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
