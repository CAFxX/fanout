package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int8; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint64; F2 int16 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 int16, p1 uint64, p2 int8) bool {
    v7 := []int{}
    var v8 int8 = p2
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            _ = len(v7)
            var v11 uint16 = 32767
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: 42}
        switch 0 {
        case -9223372036854775808:
            if false {
                var v13 float32 = float32(1.5)
                _ = v13
            }
        case 42:
            var v14 float64 = 3.14159
            _ = v14
        case 393:
            v8 = (p2 >> ((319) & 63))
        default:
            if (!false) {
                p1 = (p1 * (^(p1)))
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                if ((v8 >> ((998) & 63)) >= (p2 / (((-(p2))) | 1))) {
                    switch fi9 {
                    case 2147483647:
                        if len(v7) > 0 {
                            v7[0] = fi9
                        }
                    default:
                        if (65535 != 2) {
                            var v15 int16 = p0
                            _ = v15
                        }
                    }
                    var v16 float32 = 1e10
                    _ = v16
                }
            }
        }
        _ = fi9
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (1 > 65535)
}
func genFn1(p0 uint32, p1 uint32, p2 int32) uint64 {
    switch ("hello" + ("" + "a")) {
    case "\n\t":
        switch 42 {
        case 32768:
            var v17 int8 = 6
            _ = v17
        }
    case "hello":
        if true {
            v18 := GenS1{F0: 42}
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 uint64 = 0
                _ = fi19
                _ = v20
            }
            _ = v18
        }
    case "":
        fmt.Println(p2, p1)
    }
    p0 += p1
    _ = p0
    _ = p1
    _ = p2
    return 456
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 255
    v3 := make(map[uint32]uint32)
    var v4 float32 = 1e10
    var v5 int64 = 42
    var v6 float64 = ((0.1 * 0.0) * -2.25)
    var v21 uint32 = 2147483647
    v1[(v21 & v21)] = true
    for fi22 := 0; fi22 < 13; fi22++ {
        switch "a" {
        case "hello":
            fmt.Println(v4, v21)
        case "foo bar":
            switch (fi22 + -266) {
            case -385:
                fmt.Println(fi22, v4, fi22)
            case -1:
                v23 := GenS2{F0: 1, F1: -128, F2: 47}
                _ = v23
            case 4294967295:
                switch (^(fi22)) {
                case -9223372036854775808:
                    v24 := GenS1{F0: 193}
                    _ = v24
                case 112:
                    switch "\n\t" {
                    case "":
                        switch "a" {
                        case "":
                            var v25 uint16 = 753
                            _ = v25
                        default:
                            _ = len(v3)
                        }
                    case "foo bar":
                        var v26 int16 = 1
                        _ = v26
                    default:
                        v27 := GenS2{F0: genFn1(v21, v21, -2147483648), F1: 104, F2: -1}
                        _ = v27
                    }
                default:
                    v28 := GenS3{F0: (v5 / (((-(v5))) | 1))}
                    _ = v28
                }
            default:
                v29 := GenS2{F0: 42, F1: int8(v21), F2: 0}
                _ = v29
            }
        case "\n\t":
            if false {
                v2 = (^(v2))
                v30 := GenS4{F0: (v5 | v5), F1: 57, F2: -1}
                _ = v30
            }
        }
        switch "a" {
        case "foo bar":
            if false {
                var v31 uint16 = v2
                _ = v31
            } else {
                var v32 int8 = -1
                v33 := GenS4{F0: v5, F1: 523, F2: 1}
                _ = v32
                _ = v33
            }
        }
        switch "\n\t" {
        case "hello":
            var v34 bool = true
            _ = v34
        default:
            for fi35 := 0; fi35 < 1; fi35++ {
                v21 += v21
                _ = fi35
            }
        }
        _ = fi22
    }
    fmt.Println(v2, v5)
    var v36 float32 = (0.1 - 1e10)
    if false {
        if true {
            var v37 int64 = v5
            v5 += 9223372036854775807
            switch "hello" {
            case "":
                var v38 uint16 = 625
                _ = v38
            case "foo bar":
                for fi39 := 0; fi39 < 19; fi39++ {
                    v2 += v2
                    _ = fi39
                }
            case "hello":
                var v40 int = 32768
                _ = v40
            default:
                var v41 uint32 = v21
                _ = v41
            }
            _ = v37
        } else {
            var v42 uint64 = 0
            fmt.Println(v21, v2, v21)
            _ = v42
        }
        var v43 bool = (2 == 55)
        _ = v43
    } else {
        var v44 float64 = v6
        _ = v44
    }
    for fi45 := 0; fi45 < 10; fi45++ {
        v21 += 42
        if (0 == uint64(v21)) {
            v46 := GenS5{F0: -1, F1: 0}
            _ = v46
        } else {
            v47 := v3[2]
            for fi48 := 0; fi48 < 9; fi48++ {
                var v49 int = -73
                var v50 uint64 = uint64(v2)
                var v51 bool = true
                v52 := GenS5{F0: 0, F1: 128}
                _ = fi48
                _ = v49
                _ = v50
                _ = v51
                _ = v52
            }
            v53 := GenS6{F0: v2}
            _ = v47
            _ = v53
        }
        _ = fi45
    }
    var v54 uint32 = v21
    if false {
        fmt.Println(v5, v36, v5)
        switch 2147483647 {
        case 305:
            v36 = v4
        }
    } else {
        fmt.Println(v6)
        switch -9223372036854775808 {
        case 9223372036854775807:
            var v55 uint32 = 42
            _ = v55
        case 128:
            var v56 int16 = -1
            _ = v56
        }
    }
    _ = len(v3)
    if false {
        if false {
            v57 := []int{}
            _ = v57
        } else {
            var v58 int = int(v21)
            if true {
                if true {
                    v21 = v54
                    v59 := GenS4{F0: v5, F1: 18446744073709551615, F2: -32768}
                    _ = v59
                }
                fmt.Println(v21)
                v58 = 127
            }
            v54 = (^(v54))
            _ = v58
        }
        if (4294967295 != 308) {
            fmt.Println(v5)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
