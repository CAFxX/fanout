package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int16; F2 string }
type GenS3 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int64 }
func genFn0(p0 int64, p1 uint8, p2 uint8) uint8 {
    v5 := GenS1{F0: 42}
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 string = ("" + "")
        _ = fi6
        _ = v7
    }
    switch 1 {
    case 32767:
        switch "a" {
        case "foo bar":
            var v8 uint32 = 262
            _ = v8
        case "a":
            if (!(!true)) {
                var v9 int = -2147483648
                if false {
                    switch "" {
                    case "":
                        for fi10 := 0; fi10 < 8; fi10++ {
                            var v11 int = v9
                            v12 := []int{int(p1), int(p1), 32767, 2147483647}
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                    }
                    var v13 bool = (true && false)
                    v9 = v9
                    _ = v13
                } else {
                    v14 := []int{}
                    if true {
                        fmt.Println(v9)
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        v15 := GenS2{F0: 18446744073709551615, F1: 492, F2: "a"}
                        _ = v15
                    }
                    var v16 string = "hello"
                    _ = v14
                    _ = v16
                }
                if (false && (!false)) {
                    switch "hello" {
                    case "\n\t":
                        var v17 bool = false
                        _ = v17
                    }
                    if false {
                        p0 += p0
                    }
                } else {
                    var v18 float32 = float32(1.5)
                    _ = v18
                }
                _ = v9
            } else {
                for fi19 := 0; fi19 < 19; fi19++ {
                    for fi20 := 0; fi20 < 18; fi20++ {
                        p2 += p2
                        v21 := []int{fi20}
                        _ = fi20
                        _ = v21
                    }
                    if (0 < 2) {
                        v22 := []int{}
                        _ = v22
                    }
                    _ = fi19
                }
            }
        case "hello":
            if ((3.14159 / 1.0) >= -2.25) {
                for fi23 := 0; fi23 < 19; fi23++ {
                    var v24 int32 = -2147483648
                    _ = fi23
                    _ = v24
                }
                switch 128 {
                case 128:
                    p1 = (p1 ^ 61)
                case 32768:
                    for fi25 := 0; fi25 < 19; fi25++ {
                        var v26 int16 = 0
                        fmt.Println(p0, p2)
                        p1 -= (^(p1))
                        _ = fi25
                        _ = v26
                    }
                default:
                    var v27 uint = 65535
                    _ = v27
                }
            } else {
                v28 := GenS2{F0: 674, F1: 77, F2: ("" + ((("" + (("" + ("hello" + ("foo bar" + "a"))) + "")) + ("foo bar" + "")) + "\n\t"))}
                _ = v28
            }
        default:
            v29 := []int{}
            _ = v29
        }
    default:
        var v30 float32 = -2.25
        _ = v30
    }
    v31 := GenS1{F0: 376}
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    _ = v5
    return (p2 >> ((2) & 63))
}

func main() {
    v1 := []uint32{546, 188}
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    var v32 uint32 = 4294967295
    v1 = append(v1, v32)
    for fi33 := 0; fi33 < 14; fi33++ {
        v34 := GenS3{F0: 905, F1: 42, F2: (v32 & v32)}
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        var v36 float32 = -2.25
        for fi37 := 0; fi37 < 6; fi37++ {
            var v38 float32 = v36
            switch 1 {
            case 32768:
                if (true && false) {
                    _ = len(v3)
                }
                v39 := GenS4{F0: 354, F1: -193}
                fmt.Println(v36, v36)
                _ = v39
            case 32767:
                if len(v1) > 0 {
                    v1[0] = (v32 >> ((255) & 63))
                }
            default:
                v38 += 3.14159
            }
            if true {
                var v40 uint = 65535
                _ = v40
            } else {
                var v41 float32 = v36
                var v42 float64 = (-2.25 * ((float64(((0.0 / (-2.25 / v36)) + v41)) / -2.25) + 0.0))
                _ = v41
                _ = v42
            }
            var v43 int64 = -9223372036854775808
            if true {
                v32 = (v32 & (v32 & v32))
            }
            _ = fi37
            _ = v38
            _ = v43
        }
        switch "" {
        case "foo bar":
            var v44 uint = 411
            _ = v44
        case "a":
            var v45 uint32 = 285
            _ = v45
        }
        _ = fi35
        _ = v36
    }
    v46 := GenS4{F0: v32, F1: 1}
    for fi47 := 0; fi47 < 4; fi47++ {
        for fi48 := 0; fi48 < 2; fi48++ {
            var v49 uint32 = v32
            var v50 int16 = 1000
            _ = fi48
            _ = v49
            _ = v50
        }
        _ = fi47
    }
    v51 := GenS4{F0: v32, F1: 42}
    _ = len(v2)
    if false {
        var v52 uint64 = 411
        v53 := GenS1{F0: 42}
        v52 += v52
        _ = v52
        _ = v53
    }
    for fi54 := 0; fi54 < 3; fi54++ {
        var v55 int8 = 0
        for fi56 := 0; fi56 < 13; fi56++ {
            v1 = append(v1, (v32 / ((0) | 1)))
            _ = fi56
        }
        _ = fi54
        _ = v55
    }
    v32 -= (v32 + (^(v32)))
    var v57 uint64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v46
    _ = v51
    _ = v57
}
