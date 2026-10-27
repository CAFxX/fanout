package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 string; F1 int64 }
func genFn0(p0 bool, p1 uint32, p2 int16) uint {
    switch "\n\t" {
    case "a":
        switch "a" {
        case "foo bar":
            v5 := GenS1{F0: (p2 / ((p2) | 1))}
            _ = v5
        }
    case "":
        fmt.Println(p2)
    case "foo bar":
        if (65535 == 0) {
            switch -9223372036854775808 {
            case 281:
                v6 := GenS1{F0: (p2 / (((p2 | (p2 / ((1) | 1)))) | 1))}
                _ = v6
            default:
                switch 551 {
                case -2147483648:
                    v7 := GenS1{F0: (p2 / ((p2) | 1))}
                    _ = v7
                case -1:
                    var v8 uint8 = 145
                    _ = v8
                case -900:
                    p1 = p1
                default:
                    if (0.1 < 0.0) {
                        v9 := GenS2{F0: 125}
                        fmt.Println(p1, p1, p2)
                        _ = v9
                    } else {
                        var v10 float64 = -2.25
                        switch 32767 {
                        case 32767:
                            var v11 uint16 = 2
                            _ = v11
                        case -648:
                            for fi12 := 0; fi12 < 19; fi12++ {
                                var v13 int = fi12
                                var v14 uint64 = 1
                                switch v13 {
                                case 32768:
                                    v15 := []int{-581, fi12, fi12, (fi12 % ((fi12) | 1))}
                                    _ = v15
                                case -328:
                                    var v16 int8 = -33
                                    _ = v16
                                case 2147483647:
                                    v17 := GenS3{F0: 677, F1: (-(fi12))}
                                    _ = v17
                                }
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                        default:
                            var v18 int8 = -1
                            _ = v18
                        }
                        _ = v10
                    }
                }
            }
            var v19 int16 = 0
            fmt.Println(p2, p2, v19)
            _ = v19
        }
    default:
        p1 += 42
    }
    var v20 uint8 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 479
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 38
    v3 := []uint32{4294967295, uint32(v2), 0, 4294967295}
    v4 := make(map[string]string)
    v2 = (v2 - 42)
    v2 = 0
    v2 -= (v2 | 122)
    var v21 int8 = 35
    v22 := GenS1{F0: -32768}
    var v23 uint32 = 0
    var v24 float64 = 3.14159
    switch "foo bar" {
    case "a":
        var v25 uint8 = v2
        _ = v25
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
    case "":
        var v26 uint = 65535
        _ = v26
    default:
        v27 := GenS4{F0: "a", F1: 42}
        _ = v27
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        v24 = v24
        for fi29 := 0; fi29 < 13; fi29++ {
            v30 := v4[((("" + "foo bar") + ("foo bar" + ("\n\t" + (("hello" + ("\n\t" + ("foo bar" + "a"))) + "a")))) + "\n\t")]
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi29
            _ = v30
        }
        switch "foo bar" {
        case "hello":
            var v31 int8 = 1
            _ = v31
        case "a":
            if len(v3) > 0 {
                v3[0] = 2
            }
        case "\n\t":
            for fi32 := 0; fi32 < 4; fi32++ {
                v23 += (v23 << ((65535) & 63))
                _ = fi32
            }
        default:
            var v33 int32 = 1
            _ = v33
        }
        _ = fi28
    }
    v34 := GenS4{F0: "hello", F1: -1}
    if false {
        _ = len(v3)
        v2 -= (v2 + v2)
    }
    var v35 int16 = -840
    var v36 int8 = 1
    var v37 int32 = 0
    if (-2.25 == v24) {
        v24 = v24
    } else {
        var v38 float32 = 1.5
        fmt.Println(v38)
        _ = v38
    }
    if (0 <= 918) {
        var v39 uint8 = (^(v2))
        var v40 float32 = float32(v24)
        v37 = v37
        _ = v39
        _ = v40
    } else {
        if false {
            for fi41 := 0; fi41 < 14; fi41++ {
                var v42 int32 = (v37 - (^(v37)))
                _ = fi41
                _ = v42
            }
            for fi43 := 0; fi43 < 12; fi43++ {
                switch fi43 {
                case 791:
                    var v44 int8 = 100
                    _ = v44
                case 32767:
                    var v45 uint64 = 18446744073709551615
                    _ = v45
                case 42:
                    var v46 int8 = 42
                    _ = v46
                default:
                    var v47 int64 = -55
                    _ = v47
                }
                _ = len(v4)
                _ = fi43
            }
        }
        var v48 uint32 = v23
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
