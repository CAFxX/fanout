package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint32, p1 uint8) bool {
    v5 := GenS1{F0: true}
    for fi6 := 0; fi6 < 19; fi6++ {
        if (int(p0) >= fi6) {
            v7 := GenS2{F0: p0}
            switch 2147483647 {
            case -2147483648:
                var v8 float64 = (float64((((3.14159 + 0.0) * 1.5) * 0.0)) + 3.14159)
                _ = v8
            case 9223372036854775807:
                for fi9 := 0; fi9 < 2; fi9++ {
                    v10 := []int{(fi6 << ((2) & 63)), int(p1), (fi9 << ((750) & 63)), (-(fi6))}
                    _ = fi9
                    _ = v10
                }
            default:
                p1 = (^(p1))
            }
            _ = v7
        } else {
            if true {
                p1 = p1
                if true {
                    if (1e10 <= ((3.14159 - -2.25) + -2.25)) {
                        v11 := GenS3{F0: 0, F1: 1, F2: 9223372036854775807}
                        _ = v11
                    }
                    if true {
                        for fi12 := 0; fi12 < 13; fi12++ {
                            switch fi12 {
                            case 621:
                                v13 := GenS3{F0: 0, F1: 811, F2: -1}
                                _ = v13
                            case 1:
                                var v14 int8 = 1
                                v15 := GenS2{F0: p0}
                                p0 = p0
                                _ = v14
                                _ = v15
                            case 4294967295:
                                p1 = (p1 * p1)
                            }
                            _ = fi12
                        }
                        var v16 float32 = ((float32(((float64(0.0) * -2.25) * (-2.25 / 3.14159))) / (0.1 + float32(-2.25))) / float32(0.1))
                        _ = v16
                    } else {
                        v17 := GenS3{F0: 221, F1: 0, F2: 42}
                        switch 1 {
                        case 491:
                            switch "foo bar" {
                            case "a":
                                fmt.Println(fi6)
                            case "hello":
                                v18 := GenS3{F0: 4294967295, F1: 623, F2: 42}
                                _ = v18
                            default:
                                v19 := GenS3{F0: 252, F1: 309, F2: -9223372036854775808}
                                _ = v19
                            }
                        case 1:
                            v20 := GenS1{F0: false}
                            _ = v20
                        }
                        _ = v17
                    }
                }
                switch (fi6 / ((fi6) | 1)) {
                case 341:
                    p1 -= 42
                case 128:
                    var v21 float32 = 1e10
                    _ = v21
                case -2147483648:
                    var v22 uint32 = (^(p0))
                    _ = v22
                default:
                    fmt.Println(fi6, p0)
                }
            } else {
                var v23 int16 = -32768
                for fi24 := 0; fi24 < 3; fi24++ {
                    switch fi24 {
                    case 9223372036854775807:
                        v25 := GenS4{F0: ("\n\t" + "")}
                        _ = v25
                    case -462:
                        var v26 int64 = int64(fi24)
                        _ = v26
                    }
                    v23 = (^(v23))
                    fmt.Println(fi6)
                    _ = fi24
                }
                fmt.Println(p1)
                _ = v23
            }
            var v27 uint8 = 127
            p0 += 2147483648
            _ = v27
        }
        v28 := GenS5{F0: 42, F1: 42}
        _ = fi6
        _ = v28
    }
    _ = p0
    _ = p1
    _ = v5
    return true
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = -507
    var v3 float64 = (float64((float32(1.5) - (0.0 * 1e10))) * 3.14159)
    v4 := []uint32{1, 711, 1, 4294967295}
    var v29 uint = 18446744073709551615
    delete(v1, "foo bar")
    switch 42 {
    case 0:
        v30 := GenS2{F0: 803}
        _ = v30
    case 128:
        for fi31 := 0; fi31 < 13; fi31++ {
            if (0 > 2147483647) {
                v32 := GenS4{F0: ("hello" + "")}
                _ = v32
            } else {
                switch ("hello" + "") {
                case "":
                    if true {
                        v3 = v3
                    } else {
                        _ = len(v4)
                    }
                default:
                    for fi33 := 0; fi33 < 19; fi33++ {
                        if (true && (100 < 3)) {
                            var v34 int64 = -558
                            _ = v34
                        } else {
                            fmt.Println(v2)
                            v35 := GenS6{F0: (v29 ^ 547)}
                            _ = v35
                        }
                        v3 = v3
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi33
                    }
                }
                if false {
                    _ = len(v1)
                    v36 := v1["foo bar"]
                    _ = v36
                }
                var v37 uint = v29
                _ = v37
            }
            _ = fi31
        }
    }
    if (2 <= 776) {
        var v38 float64 = float64(-2.25)
        v39 := GenS4{F0: "hello"}
        _ = v38
        _ = v39
    } else {
        var v40 uint = (v29 & (^(v29)))
        _ = v40
    }
    var v41 float64 = (v3 - (((v3 / (v3 / -2.25)) - (1e10 + v3)) * (1e10 * -2.25)))
    _ = len(v4)
    switch 438 {
    case 2147483647:
        fmt.Println(v2)
    case 128:
        if true {
            var v42 uint16 = 65535
            _ = v42
        } else {
            v2 = -32768
            v43 := GenS5{F0: v2, F1: v2}
            _ = v43
        }
    default:
        var v44 bool = false
        _ = v44
    }
    v45 := GenS7{F0: 42, F1: v2}
    var v46 uint = (v29 << ((v29) & 63))
    v47 := GenS6{F0: 42}
    var v48 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v41
    _ = v45
    _ = v46
    _ = v47
    _ = v48
}
