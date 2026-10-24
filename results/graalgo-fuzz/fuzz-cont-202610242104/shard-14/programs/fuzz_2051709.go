package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 bool; F1 int8; F2 int16 }
type GenS6 struct { F0 bool }
func genFn0(p0 bool, p1 int64, p2 string) uint32 {
    var v4 uint32 = 579
    var v5 uint32 = (^(v4))
    if p0 {
        v6 := GenS1{F0: 0}
        _ = v6
    }
    switch -229 {
    case -1:
        var v7 int = 128
        _ = v7
    case 32768:
        if ((65535 == 4294967295) && false) {
            v8 := GenS1{F0: 0}
            _ = v8
        } else {
            var v9 float32 = 0.0
            _ = v9
        }
    default:
        v10 := []int{32767, 42}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (v4 >> ((184) & 63))
}
func genFn1(p0 uint64) string {
    var v11 int8 = 100
    if false {
        if (!true) {
            v12 := GenS2{F0: v11, F1: 589}
            p0 = (p0 | 632)
            v13 := GenS1{F0: v11}
            _ = v12
            _ = v13
        } else {
            v11 = (v11 >> ((uint(p0)) & 63))
            p0 -= 1
            switch 128 {
            case 9223372036854775807:
                v14 := GenS3{F0: false, F1: int(p0)}
                _ = v14
            case 127:
                v11 += -42
            }
        }
        var v15 int64 = 0
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 int16 = 42
        v11 -= (-(v11))
        p0 = (p0 >> ((915) & 63))
        _ = fi16
        _ = v17
    }
    switch 4294967295 {
    case -279:
        p0 = 1
    default:
        p0 = p0
    }
    _ = p0
    _ = v11
    return "a"
}

func main() {
    var v1 bool = (21 < 0)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    if v1 {
        fmt.Println(v1)
        v18 := []int{2147483647, -952, 453, 350}
        for fi19 := 0; fi19 < 18; fi19++ {
            for fi20 := 0; fi20 < 17; fi20++ {
                v1 = false
                _ = len(v18)
                _ = fi20
            }
            v1 = ((0.1 > float32((0.1 - -2.25))) || ((^(fi19)) == 32767))
            _ = fi19
        }
        _ = v18
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 int8 = 100
        _ = fi21
        _ = v22
    }
    v23 := GenS4{F0: 42, F1: genFn0((!false), -9223372036854775808, "hello")}
    var v24 int64 = -9223372036854775808
    if v1 {
        var v25 int32 = 0
        for fi26 := 0; fi26 < 8; fi26++ {
            v27 := GenS1{F0: -1}
            var v28 int32 = (v25 >> ((4294967295) & 63))
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = v25
    } else {
        v29 := GenS5{F0: (-1 >= 30), F1: 100, F2: 103}
        v1 = true
        _ = v29
    }
    if false {
        var v30 uint32 = genFn0((!false), v24, "a")
        for fi31 := 0; fi31 < 3; fi31++ {
            if v1 {
                if (-32768 < 1) {
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        var v32 float64 = 3.14159
                        _ = v32
                    case "\n\t":
                        delete(v2, (v30 ^ (v30 % (((^(v30))) | 1))))
                    case "hello":
                        v33 := GenS3{F0: (2 == 42), F1: 2147483647}
                        _ = v33
                    default:
                        var v34 float64 = float64(1e10)
                        _ = v34
                    }
                    v35 := GenS2{F0: -43, F1: 0}
                    switch ("\n\t" + "a") {
                    case "\n\t":
                        v24 += -516
                    }
                    _ = v35
                } else {
                    v36 := GenS1{F0: 127}
                    var v37 int64 = (v24 - v24)
                    _ = v36
                    _ = v37
                }
                v38 := GenS3{F0: (461 != 32767), F1: fi31}
                _ = v38
            }
            v2[v30] = 2
            _ = fi31
        }
        _ = v30
    } else {
        v39 := []int{}
        _ = v39
    }
    for fi40 := 0; fi40 < 9; fi40++ {
        var v41 uint = 0
        switch fi40 {
        case -262:
            v24 = v24
        case 32768:
            if v1 {
                v24 = -254
                if (fi40 < 32768) {
                    fmt.Println(fi40, v24)
                    var v42 int8 = 1
                    var v43 float64 = 1.5
                    _ = v42
                    _ = v43
                } else {
                    v24 = 1
                    var v44 float64 = (1e10 + 0.0)
                    var v45 uint32 = 2147483648
                    _ = v44
                    _ = v45
                }
                var v46 int16 = -1
                _ = v46
            } else {
                v2[923] = genFn0((float32(3.14159) == -2.25), v24, "")
            }
        }
        v41 = v41
        _ = fi40
        _ = v41
    }
    switch "hello" {
    case "":
        for fi47 := 0; fi47 < 11; fi47++ {
            v48 := GenS1{F0: 100}
            switch ("foo bar" + "a") {
            case "":
                if (0.1 != (3.14159 * ((0.1 / 3.14159) - 1e10))) {
                    var v49 uint64 = 87
                    switch (("\n\t" + "a") + "foo bar") {
                    case "\n\t":
                        fmt.Println(v1, v1, v1)
                    }
                    v50 := []int{(-(fi47)), 42, fi47, 9223372036854775807}
                    _ = v49
                    _ = v50
                }
            case "hello":
                v24 -= v24
            default:
                for fi51 := 0; fi51 < 19; fi51++ {
                    v52 := []int{(fi47 ^ fi47), fi47, 2147483647, (fi47 ^ fi47)}
                    _ = fi51
                    _ = v52
                }
            }
            switch fi47 {
            case 0:
                var v53 uint16 = 2
                _ = v53
            }
            _ = fi47
            _ = v48
        }
    case "hello":
        v1 = (true && v1)
    default:
        if (v1 && (true && v1)) {
            var v54 int8 = 0
            v1 = ((0.1 * 1.5) > (0.1 - 1e10))
            _ = v54
        } else {
            var v55 uint8 = 0
            var v56 float64 = (1e10 / -2.25)
            switch 13 {
            case 127:
                var v57 bool = false
                _ = v57
            case 0:
                if v1 {
                    switch "" {
                    case "foo bar":
                        var v58 float64 = 0.1
                        _ = v58
                    default:
                        _ = len(v3)
                    }
                } else {
                    v59 := GenS2{F0: 42, F1: 1}
                    _ = v59
                }
            default:
                if false {
                    fmt.Println(v1, v55)
                }
            }
            _ = v55
            _ = v56
        }
    }
    switch "foo bar" {
    case "hello":
        if v1 {
            if (42 >= 128) {
                v24 = 976
            } else {
                v60 := GenS1{F0: -87}
                v61 := v2[1]
                _ = v60
                _ = v61
            }
            for fi62 := 0; fi62 < 1; fi62++ {
                for fi63 := 0; fi63 < 20; fi63++ {
                    v24 = (v24 >> ((1) & 63))
                    v64 := v3[2147483648]
                    _ = fi63
                    _ = v64
                }
                _ = fi62
            }
        }
    }
    for fi65 := 0; fi65 < 14; fi65++ {
        var v66 uint32 = 4294967295
        if (true || (42 != 127)) {
            v66 = v66
            if false {
                var v67 int64 = v24
                var v68 int64 = -9223372036854775808
                _ = v67
                _ = v68
            }
        } else {
            v69 := GenS6{F0: v1}
            var v70 float32 = float32(0.0)
            var v71 uint64 = 48
            _ = v69
            _ = v70
            _ = v71
        }
        if ((-2.25 > 3.14159) && (0.0 < float32(float64((((0.0 * ((0.1 / 1e10) + 1.5)) / (1e10 - float32(3.14159))) * ((3.14159 / float32((1e10 / (float64(3.14159) * (3.14159 / (float64(float32((1.5 + 0.0))) + 1e10)))))) - (1e10 * 1.5))))))) {
            for fi72 := 0; fi72 < 12; fi72++ {
                v24 = (v24 & 42)
                v73 := []int{32768, (-(fi65)), fi72, fi65}
                _ = fi72
                _ = v73
            }
        }
        _ = fi65
        _ = v66
    }
    switch 9223372036854775807 {
    case 0:
        switch 821 {
        case -263:
            for fi74 := 0; fi74 < 20; fi74++ {
                v24 -= v24
                _ = fi74
            }
        case -2147483648:
            fmt.Println(v24, v24)
        }
    }
    var v75 uint = 1
    v76 := []int{600}
    for fi77 := 0; fi77 < 1; fi77++ {
        _ = len(v2)
        if len(v76) > 0 {
            v76[0] = fi77
        }
        _ = fi77
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v75:", v75)
    fmt.Println("slicelen v76:", len(v76), v76)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v75
    _ = v76
}
