package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int; F2 uint64 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint16; F1 uint; F2 uint8 }
func genFn0(p0 int8) int8 {
    switch "hello" {
    case "a":
        if false {
            var v4 uint = 18446744073709551615
            var v5 int8 = -75
            _ = v4
            _ = v5
        } else {
            fmt.Println(p0)
            p0 += (p0 << ((65535) & 63))
        }
    case "foo bar":
        for fi6 := 0; fi6 < 6; fi6++ {
            fmt.Println(p0, fi6)
            _ = fi6
        }
    case "\n\t":
        var v7 int = -9223372036854775808
        _ = v7
    default:
        p0 = (p0 & p0)
    }
    fmt.Println(p0)
    v8 := []int{32767, 754}
    _ = p0
    _ = v8
    return 42
}

func main() {
    v1 := []string{"\n\t", ("\n\t" + ("a" + ("" + "foo bar"))), ("hello" + "")}
    v2 := []string{}
    v3 := []int{42, 4294967295, 127, 0}
    var v9 int16 = 0
    v9 = (v9 | 1000)
    v1 = append(v1, "\n\t")
    v10 := GenS1{F0: 42}
    if ("a" <= "foo bar") {
        v11 := GenS2{F0: 778, F1: 4294967295, F2: 411}
        _ = v11
    } else {
        v9 = -867
    }
    if true {
        var v12 float32 = ((-2.25 - 1.5) / 1.0)
        _ = v12
    }
    if (1 >= 851) {
        if (((0.1 + (1.5 * 3.14159)) / 3.14159) == 1.5) {
            v13 := GenS1{F0: -1}
            switch 32768 {
            case 128:
                for fi14 := 0; fi14 < 6; fi14++ {
                    switch 295 {
                    case -2147483648:
                        _ = len(v2)
                    case 4294967295:
                        v9 = v9
                    case 1:
                        if true {
                            v9 = (v9 + 1)
                            var v15 float32 = ((3.14159 - 3.14159) - 3.14159)
                            _ = v15
                        }
                    }
                    for fi16 := 0; fi16 < 1; fi16++ {
                        v17 := GenS1{F0: 2147483647}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v18 int8 = -1
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    fmt.Println(fi14, v9)
                    _ = fi14
                }
            case 636:
                if (-9223372036854775808 < 742) {
                    var v19 int64 = -260
                    var v20 uint = 255
                    _ = v19
                    _ = v20
                }
            case 9223372036854775807:
                switch "foo bar" {
                case "a":
                    for fi21 := 0; fi21 < 13; fi21++ {
                        var v22 uint16 = 1
                        var v23 int32 = -1
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                case "":
                    v24 := GenS2{F0: 65535, F1: 1, F2: 1}
                    _ = v24
                case "hello":
                    v9 = -575
                default:
                    for fi25 := 0; fi25 < 2; fi25++ {
                        if (128 == 2) {
                            for fi26 := 0; fi26 < 11; fi26++ {
                                v9 = v9
                                if (759 >= 1) {
                                    fmt.Println(fi25, fi25, fi26)
                                }
                                v3 = append(v3, (-(fi26)))
                                _ = fi26
                            }
                            if (true || false) {
                                v9 -= -32768
                                var v27 int64 = -101
                                _ = v27
                            } else {
                                v28 := GenS3{F0: 42, F1: false}
                                _ = v28
                            }
                        }
                        var v29 int16 = 42
                        _ = fi25
                        _ = v29
                    }
                }
            }
            switch 127 {
            case -1:
                var v30 float32 = (((0.1 - ((0.1 * -2.25) + 1.5)) * -2.25) / -2.25)
                _ = v30
            case 768:
                v9 -= v9
            }
            _ = v13
        } else {
            _ = len(v1)
            if len(v3) > 0 {
                v3[0] = -1
            }
        }
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        var v32 float32 = 1e10
        fmt.Println(v32, v9, v9)
        _ = fi31
        _ = v32
    }
    switch "foo bar" {
    case "a":
        if true {
            for fi33 := 0; fi33 < 6; fi33++ {
                var v34 int32 = -1
                v9 += (v9 - (v9 / ((int16(v34)) | 1)))
                _ = fi33
                _ = v34
            }
            v9 -= 936
        } else {
            v9 += 142
            v3 = append(v3, 4294967295)
            v35 := GenS1{F0: 42}
            _ = v35
        }
    case "hello":
        fmt.Println(v9)
    case "":
        switch 128 {
        case -1:
            var v36 uint64 = 0
            _ = v36
        case 272:
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 int32 = 2147483647
                v9 = v9
                _ = fi37
                _ = v38
            }
        case 24:
            for fi39 := 0; fi39 < 15; fi39++ {
                v40 := GenS4{F0: 42, F1: 329, F2: 0}
                if (255 <= 72) {
                    var v41 int8 = 42
                    var v42 bool = (((((((1.5 * 0.1) + 1e10) / (((-2.25 * 3.14159) / ((1.5 / 3.14159) / float64((-2.25 / 1.0)))) + (-2.25 - 3.14159))) / 1.5) / 3.14159) != 1e10) || ((true || true) || (fi39 == -99)))
                    _ = v41
                    _ = v42
                }
                for fi43 := 0; fi43 < 17; fi43++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v9 += -1
                    if (!true) {
                        var v44 uint8 = 255
                        _ = v44
                    } else {
                        _ = len(v1)
                        fmt.Println(v9, fi43, fi43)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = fi43
                }
                _ = fi39
                _ = v40
            }
        }
    }
    for fi45 := 0; fi45 < 14; fi45++ {
        for fi46 := 0; fi46 < 4; fi46++ {
            v47 := GenS4{F0: 32767, F1: 255, F2: 128}
            _ = fi46
            _ = v47
        }
        fmt.Println(fi45)
        _ = fi45
    }
    var v48 float32 = float32(3.14159)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v48:", v48)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v48
    _ = v9
}
