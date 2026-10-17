package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 bool; F1 int }
func genFn0(p0 int16, p1 bool, p2 int) uint64 {
    var v6 uint64 = 4294967295
    var v7 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 uint) uint16 {
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int32 = -498
        _ = fi8
        _ = v9
    }
    if (((((0.1 + (1.5 - float64((float32((((1e10 - 1.5) - 0.1) + 3.14159)) / 1.5)))) / 0.1) * 1.5) - 1e10) >= 1e10) {
        v10 := []int{893}
        v10 = append(v10, 32768)
        _ = v10
    }
    _ = p0
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = -966
    var v3 uint32 = 42
    v4 := make(map[uint32]uint32)
    var v5 uint = 485
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS1{F0: 2, F1: 40, F2: 32767}
        fmt.Println(fi11, v2)
        _ = fi11
        _ = v12
    }
    v13 := []int{v2, (v2 & v2), 42, (v2 + (v2 * v2))}
    for fi14 := 0; fi14 < 20; fi14++ {
        if false {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            if false {
                fmt.Println(fi14, fi14)
                fmt.Println(v5, v3)
                for fi15 := 0; fi15 < 14; fi15++ {
                    _ = len(v13)
                    fmt.Println(v5)
                    v16 := GenS1{F0: 899, F1: 0, F2: genFn1((v5 + (v5 + (v5 + (v5 - v5)))))}
                    var v17 uint16 = 1
                    switch (("hello" + "") + "") {
                    case "foo bar":
                        v3 = v3
                    case "a":
                        switch (fi14 >> ((2) & 63)) {
                        case 32767:
                            fmt.Println(fi14, v2, v5)
                        case 32768:
                            var v18 float64 = 0.1
                            _ = v18
                        default:
                            switch (fi15 % ((1) | 1)) {
                            case 974:
                                var v19 bool = (1.5 >= 0.0)
                                _ = v19
                            default:
                                v20 := GenS1{F0: (v17 - (v17 / ((genFn1((v5 >> ((1) & 63)))) | 1))), F1: 1, F2: v17}
                                _ = v20
                            }
                        }
                    }
                    v17 += (v17 << ((v5) & 63))
                    switch int(v3) {
                    case 32767:
                        if false {
                            switch 32768 {
                            case 0:
                                var v21 bool = true
                                _ = v21
                            case 32767:
                                var v22 uint16 = genFn1((v5 & v5))
                                _ = v22
                            case 602:
                                v3 = (v3 + 706)
                            default:
                                switch "\n\t" {
                                case "foo bar":
                                    var v23 uint8 = uint8(v5)
                                    _ = v23
                                case "a":
                                    v5 += v5
                                case "\n\t":
                                    var v24 float64 = 0.0
                                    _ = v24
                                }
                            }
                        } else {
                            v25 := v1[(v3 - v3)]
                            switch ("\n\t" + "hello") {
                            case "a":
                                var v26 uint8 = 0
                                _ = v26
                            }
                            _ = v25
                        }
                    case 4294967295:
                        v3 += (v3 & 4294967295)
                    case -434:
                        v3 += v3
                    }
                    switch -895 {
                    case 127:
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    }
                    v27 := GenS1{F0: 0, F1: 1, F2: genFn1((v5 | (v5 << ((v5) & 63))))}
                    v2 = (fi15 % (((-(fi14))) | 1))
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v27
                }
            } else {
                var v28 int8 = 83
                if (true && false) {
                    v13 = append(v13, fi14)
                    if true {
                        v29 := GenS2{F0: -1}
                        _ = v29
                    }
                }
                _ = v28
            }
            for fi30 := 0; fi30 < 12; fi30++ {
                v2 -= v2
                _ = fi30
            }
        } else {
            var v31 float32 = ((1e10 * 1.5) + -2.25)
            v32 := GenS1{F0: genFn1(v5), F1: 0, F2: 255}
            _ = v31
            _ = v32
        }
        fmt.Println(v2)
        _ = fi14
    }
    v33 := GenS1{F0: 2, F1: 41, F2: 690}
    v34 := GenS3{F0: "hello"}
    v35 := GenS3{F0: ("\n\t" + ("hello" + "hello"))}
    v36 := GenS4{F0: (^(v5)), F1: 127, F2: 12}
    v37 := GenS5{F0: 65535, F1: 1}
    v1[v3] = (!false)
    for fi38 := 0; fi38 < 5; fi38++ {
        v4[(v3 * v3)] = (v3 << ((255) & 63))
        for fi39 := 0; fi39 < 4; fi39++ {
            if (0.0 >= 0.0) {
                var v40 uint8 = 127
                if ((((false || (0 <= 0)) && true) || true) && false) {
                    _ = len(v13)
                }
                _ = v40
            }
            var v41 string = "hello"
            _ = fi39
            _ = v41
        }
        _ = fi38
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        if (!((!false) || (1 != 127))) {
            if (0 < 42) {
                v5 += v5
            } else {
                switch ((((("foo bar" + "a") + "hello") + "a") + ("\n\t" + "a")) + ("" + ("foo bar" + "hello"))) {
                case "\n\t":
                    if (!("" < "\n\t")) {
                        v2 = (-(v2))
                        switch v2 {
                        case 281:
                            v3 = 1
                        case 42:
                            var v43 int16 = 1
                            _ = v43
                        }
                    }
                }
                fmt.Println(fi42)
            }
            fmt.Println(fi42)
            if len(v13) > 0 {
                v13[0] = fi42
            }
        }
        v44 := GenS1{F0: uint16(v5), F1: 1, F2: genFn1(v5)}
        var v45 uint8 = 127
        _ = fi42
        _ = v44
        _ = v45
    }
    v2 = int(v3)
    v5 = 1
    var v46 bool = true
    v46 = ("foo bar" >= "hello")
    for fi47 := 0; fi47 < 18; fi47++ {
        for fi48 := 0; fi48 < 6; fi48++ {
            var v49 int64 = 42
            v50 := GenS4{F0: v5, F1: 255, F2: 2147483647}
            fmt.Println(v46, fi48)
            var v51 uint = v5
            if false {
                v51 = 2
            } else {
                switch "" {
                case "a":
                    var v52 uint = (^(v5))
                    _ = v52
                case "foo bar":
                    v53 := GenS6{F0: v46, F1: (fi48 ^ 32768)}
                    _ = v53
                }
            }
            _ = fi48
            _ = v49
            _ = v50
            _ = v51
        }
        v2 -= (^(fi47))
        if false {
            var v54 uint32 = (^(v3))
            v46 = v46
            _ = v54
        }
        _ = fi47
    }
    v3 = (v3 & v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v46
    _ = v5
}
