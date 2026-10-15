package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint; F1 int; F2 int }
type GenS5 struct { F0 int8; F1 bool; F2 string }
func genFn0(p0 int64) uint {
    var v7 uint8 = 215
    switch 9223372036854775807 {
    case 2147483647:
        for fi8 := 0; fi8 < 4; fi8++ {
            v7 = (^(v7))
            _ = fi8
        }
    }
    switch int(v7) {
    case 2147483647:
        var v9 string = "a"
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 int8 = 1
        v12 := GenS1{F0: 949, F1: -661, F2: (v7 | (v7 / ((v7) | 1)))}
        fmt.Println(p0)
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: "a"}
    _ = p0
    _ = v13
    _ = v7
    return 65535
}
func genFn1(p0 uint32, p1 uint64, p2 string) int64 {
    if (((!false) || ((317 == 418) || (p0 < 0))) || true) {
        v14 := GenS2{F0: (p2 + "\n\t")}
        var v15 uint64 = (^(p1))
        _ = v14
        _ = v15
    } else {
        switch p2 {
        case "":
            v16 := []int{2147483647}
            _ = v16
        }
    }
    var v17 int64 = 9223372036854775807
    var v18 uint16 = 32767
    p0 += 414
    if false {
        var v19 uint64 = p1
        v20 := GenS1{F0: 2147483647, F1: 1, F2: 42}
        var v21 uint32 = (p0 >> ((1) & 63))
        _ = v19
        _ = v20
        _ = v21
    } else {
        v22 := GenS3{F0: "foo bar", F1: (^(p0)), F2: int8(p0)}
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return v17
}

func main() {
    var v1 int16 = 32767
    var v2 float32 = float32((-2.25 + (1e10 + (0.1 - 0.0))))
    var v3 uint16 = 849
    var v4 int8 = -128
    v5 := []int{0, -729, 127}
    var v6 float64 = 3.14159
    for fi23 := 0; fi23 < 10; fi23++ {
        v4 = v4
        if true {
            v1 -= (v1 * 32767)
        } else {
            var v24 uint32 = 4294967295
            _ = v24
        }
        for fi25 := 0; fi25 < 6; fi25++ {
            if false {
                if len(v5) > 0 {
                    v5[0] = (fi25 * 836)
                }
            }
            _ = fi25
        }
        _ = fi23
    }
    var v26 uint8 = 118
    if false {
        var v27 uint32 = 2
        switch ("\n\t" + "a") {
        case "a":
            _ = len(v5)
        }
        v28 := GenS2{F0: ""}
        _ = v27
        _ = v28
    }
    v2 = float32(v6)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        for fi30 := 0; fi30 < 12; fi30++ {
            var v31 int32 = 124
            var v32 uint16 = v3
            _ = fi30
            _ = v31
            _ = v32
        }
        fmt.Println(v3, v3, v6)
        _ = fi29
    }
    fmt.Println(v3)
    var v33 int16 = 1
    switch ((("\n\t" + "a") + "a") + "hello") {
    case "":
        switch -676 {
        case 127:
            v4 = v4
        case -539:
            var v34 int16 = -1
            _ = v34
        }
    case "\n\t":
        fmt.Println(v3, v3)
    default:
        fmt.Println(v3, v2)
    }
    _ = len(v5)
    v33 = (v1 * -79)
    for fi35 := 0; fi35 < 13; fi35++ {
        switch fi35 {
        case 332:
            for fi36 := 0; fi36 < 7; fi36++ {
                v4 = 0
                _ = fi36
            }
        default:
            var v37 int32 = int32(fi35)
            _ = v37
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi35
    }
    switch "\n\t" {
    case "":
        for fi38 := 0; fi38 < 4; fi38++ {
            if true {
                var v39 uint32 = 235
                switch "\n\t" {
                case "foo bar":
                    for fi40 := 0; fi40 < 3; fi40++ {
                        if false {
                            var v41 uint16 = (v3 % ((882) | 1))
                            v3 = (v41 / ((v41) | 1))
                            _ = v41
                        } else {
                            v42 := GenS2{F0: "\n\t"}
                            v43 := GenS4{F0: 151, F1: fi38, F2: fi38}
                            v33 = (v1 & -1)
                            _ = v42
                            _ = v43
                        }
                        _ = fi40
                    }
                case "a":
                    var v44 uint16 = 65535
                    _ = v44
                case "":
                    if false {
                        v45 := GenS2{F0: ("foo bar" + (("\n\t" + ("hello" + "\n\t")) + "a"))}
                        fmt.Println(fi38, v6, v2)
                        for fi46 := 0; fi46 < 4; fi46++ {
                            var v47 int8 = (-(v4))
                            v48 := GenS2{F0: "a"}
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                        _ = v45
                    } else {
                        var v49 float32 = v2
                        v3 = (^(v3))
                        _ = v49
                    }
                default:
                    for fi50 := 0; fi50 < 8; fi50++ {
                        switch "" {
                        case "a":
                            fmt.Println(v39)
                        case "foo bar":
                            if false {
                                var v51 uint16 = (v3 << ((858) & 63))
                                v4 = v4
                                _ = len(v5)
                                var v52 bool = false
                                var v53 int = 0
                                var v54 uint32 = (v39 + v39)
                                v5 = append(v5, v53)
                                _ = v51
                                _ = v52
                                _ = v53
                                _ = v54
                            }
                        }
                        switch (fi50 + (fi38 & (fi38 & 127))) {
                        case 128:
                            v55 := GenS5{F0: 105, F1: (true && false), F2: ""}
                            _ = v55
                        case 1:
                            if len(v5) > 0 {
                                v5[0] = 346
                            }
                        }
                        _ = fi50
                    }
                }
                v6 = v6
                _ = v39
            }
            for fi56 := 0; fi56 < 18; fi56++ {
                var v57 int16 = (-(v1))
                v26 -= (v26 / ((2) | 1))
                v3 = (v3 | v3)
                _ = fi56
                _ = v57
            }
            _ = fi38
        }
    }
    switch "hello" {
    case "":
        var v58 float32 = v2
        _ = v58
    case "foo bar":
        fmt.Println(v26, v3, v26)
    default:
        v26 = (v26 * 255)
    }
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
