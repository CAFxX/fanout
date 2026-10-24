package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int16; F2 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64; F1 int32; F2 int }
type GenS7 struct { F0 int8; F1 uint64; F2 uint64 }
func genFn0(p0 uint16, p1 int64) uint64 {
    p1 -= int64(p0)
    switch "a" {
    case "hello":
        if (1 < p0) {
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 uint64 = 42
                v8 = v8
                v8 += v8
                _ = fi7
                _ = v8
            }
            p1 -= p1
            p1 = (-(p1))
        } else {
            if true {
                v9 := []int{800, 907}
                _ = v9
            }
            switch "foo bar" {
            case "foo bar":
                p0 = (^(p0))
            case "":
                fmt.Println(p0, p1)
            default:
                v10 := []int{9223372036854775807, -182, 1}
                _ = v10
            }
            var v11 string = "hello"
            _ = v11
        }
    default:
        if (false || (0.1 < -2.25)) {
            p0 = (p0 % ((1) | 1))
            if true {
                var v12 string = ""
                fmt.Println(v12, p0, v12)
                v13 := []int{}
                _ = v12
                _ = v13
            } else {
                if (0.1 > (3.14159 / (1e10 - (0.0 / 3.14159)))) {
                    for fi14 := 0; fi14 < 16; fi14++ {
                        for fi15 := 0; fi15 < 5; fi15++ {
                            switch (fi15 - int(p0)) {
                            case 42:
                                var v16 float64 = float64((-2.25 / (((1e10 / ((3.14159 * (((0.1 / float32((3.14159 - (1.5 - 0.1)))) + 1e10) + (-2.25 * ((0.0 / 1.0) / 1.0)))) * 3.14159)) + 3.14159) * 1e10)))
                                _ = v16
                            case -664:
                                switch ("hello" + "foo bar") {
                                case "":
                                    p1 -= (p1 - p1)
                                case "hello":
                                    var v17 uint8 = 2
                                    _ = v17
                                }
                            default:
                                p1 -= p1
                            }
                            var v18 uint16 = (p0 | 582)
                            _ = fi15
                            _ = v18
                        }
                        _ = fi14
                    }
                    p0 -= (^(p0))
                } else {
                    p1 = (p1 - 0)
                }
            }
        }
    }
    p0 += 0
    _ = p0
    _ = p1
    return 4294967295
}
func genFn1() bool {
    var v19 uint8 = 2
    v19 = v19
    if true {
        v19 -= v19
    }
    v20 := GenS1{F0: 0}
    v21 := []int{0, 127, 1, 0}
    _ = v19
    _ = v20
    _ = v21
    return (((1.5 - (1.5 + 0.1)) < 3.14159) && ((0.0 * (-2.25 * -2.25)) != 0.0))
}

func main() {
    v1 := []int{42, 589, 9223372036854775807}
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    var v4 uint8 = 1
    var v5 uint = 804
    var v6 int = 2147483647
    if false {
        var v22 float64 = float64((1e10 / float32(((0.1 / float64(-2.25)) / (-2.25 + float64(3.14159))))))
        v23 := GenS2{F0: 1, F1: (39 <= 1)}
        fmt.Println(v5, v22, v5)
        _ = v22
        _ = v23
    }
    v24 := GenS3{F0: -13}
    if (!((false || (1 != 127)) && (v4 != 127))) {
        v1 = append(v1, (v6 / ((v6) | 1)))
    }
    v6 -= 2147483647
    switch "a" {
    case "hello":
        v5 = (^(v5))
    case "a":
        switch ((((("\n\t" + ("hello" + "")) + "\n\t") + "foo bar") + ("\n\t" + "a")) + ("" + "")) {
        case "foo bar":
            v25 := GenS4{F0: -1, F1: 32767, F2: -2147483648}
            _ = v25
        case "":
            for fi26 := 0; fi26 < 13; fi26++ {
                switch (-(v6)) {
                case -1:
                    v27 := GenS1{F0: (^(v4))}
                    _ = v27
                case 2147483647:
                    v28 := GenS3{F0: 0}
                    _ = v28
                case 127:
                    var v29 int64 = 0
                    _ = v29
                }
                _ = len(v1)
                var v30 string = "foo bar"
                _ = fi26
                _ = v30
            }
        case "hello":
            var v31 uint64 = genFn0(0, 373)
            _ = v31
        }
    case "foo bar":
        v32 := GenS5{F0: -1}
        _ = v32
    default:
        for fi33 := 0; fi33 < 16; fi33++ {
            for fi34 := 0; fi34 < 20; fi34++ {
                fmt.Println(v4)
                v35 := GenS6{F0: 1, F1: -1, F2: (fi34 & (-(v6)))}
                v6 -= (v6 * (-(fi33)))
                _ = fi34
                _ = v35
            }
            if false {
                fmt.Println(fi33)
                v6 = (v6 % ((-172) | 1))
                v5 -= v5
            }
            _ = fi33
        }
    }
    var v36 uint64 = 398
    fmt.Println(v6, v36)
    if false {
        fmt.Println(v36)
    }
    fmt.Println(v6)
    for fi37 := 0; fi37 < 3; fi37++ {
        v38 := v2[v6]
        var v39 bool = (v38 < v38)
        _ = fi37
        _ = v38
        _ = v39
    }
    v40 := GenS4{F0: -94, F1: 949, F2: 0}
    var v41 bool = (false && ((-2.25 == 3.14159) && true))
    if false {
        if (v41 || (1.5 > float64((-2.25 * -2.25)))) {
            v41 = (!((v36 >> ((v5) & 63)) > genFn0(842, 0)))
            var v42 uint32 = 42
            switch 128 {
            case 128:
                v43 := v3[uint32(v5)]
                _ = v43
            case 127:
                if (!true) {
                    var v44 int8 = 2
                    v45 := GenS1{F0: v4}
                    var v46 int = v6
                    _ = v44
                    _ = v45
                    _ = v46
                } else {
                    fmt.Println(v41, v4, v5)
                    var v47 int16 = 42
                    for fi48 := 0; fi48 < 19; fi48++ {
                        var v49 int16 = v47
                        _ = fi48
                        _ = v49
                    }
                    _ = v47
                }
            }
            _ = v42
        }
        v1 = append(v1, v6)
        for fi50 := 0; fi50 < 10; fi50++ {
            for fi51 := 0; fi51 < 6; fi51++ {
                var v52 int16 = -512
                var v53 string = "a"
                _ = fi51
                _ = v52
                _ = v53
            }
            if v41 {
                fmt.Println(fi50, v41)
            } else {
                var v54 uint = v5
                _ = v54
            }
            _ = fi50
        }
    } else {
        if (!v41) {
            if (!(852 < 1000)) {
                fmt.Println(v4, v36)
                v41 = (-123 >= -33)
            }
            v55 := GenS4{F0: 0, F1: -1, F2: int32(v4)}
            _ = v55
        }
        var v56 int32 = -1
        _ = v56
    }
    if v41 {
        var v57 int8 = 100
        var v58 bool = ((v5 == (v5 ^ 562)) || v41)
        _ = v57
        _ = v58
    }
    if true {
        switch 375 {
        case -982:
            var v59 int64 = -9223372036854775808
            _ = v59
        case 32767:
            v41 = false
        }
        var v60 int8 = 100
        _ = v60
    } else {
        v61 := GenS7{F0: 78, F1: (^(v36)), F2: genFn0(255, 577)}
        for fi62 := 0; fi62 < 16; fi62++ {
            v1 = append(v1, (v6 - v6))
            var v63 uint64 = (v36 % (((v36 + v36)) | 1))
            _ = fi62
            _ = v63
        }
        _ = v61
    }
    switch v6 {
    case 2147483647:
        v6 += (v6 >> (((v5 % ((v5) | 1))) & 63))
    }
    fmt.Println(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
