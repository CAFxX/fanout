package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int8; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint64; F1 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -2147483648
    var v3 int16 = -786
    v4 := []string{"a", "foo bar", ""}
    var v5 bool = (1.5 <= 1.5)
    for fi6 := 0; fi6 < 5; fi6++ {
        v5 = v5
        switch "a" {
        case "\n\t":
            if v5 {
                fmt.Println(v2)
                fmt.Println(v2)
                v3 = (v3 | v3)
            }
        case "":
            var v7 int8 = 1
            _ = v7
        case "a":
            var v8 uint8 = 0
            _ = v8
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int32 = v2
        v2 += 42
        _ = fi9
        _ = v10
    }
    switch -1 {
    case 0:
        v1["foo bar"] = false
    case 2147483647:
        if len(v4) > 0 {
            v4[0] = "a"
        }
    case 1:
        var v11 int = 127
        _ = v11
    default:
        switch 4294967295 {
        case 4294967295:
            for fi12 := 0; fi12 < 9; fi12++ {
                if v5 {
                    fmt.Println(fi12, fi12)
                    v13 := GenS1{F0: 0, F1: -851}
                    if (1 > 2) {
                        v14 := GenS1{F0: -1, F1: 9223372036854775807}
                        delete(v1, "\n\t")
                        var v15 int64 = -379
                        _ = v14
                        _ = v15
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v16 float32 = 1.5
                        _ = v16
                    }
                    _ = v13
                }
                for fi17 := 0; fi17 < 16; fi17++ {
                    v2 -= (v2 % ((-1) | 1))
                    _ = fi17
                }
                _ = fi12
            }
        case -2147483648:
            fmt.Println(v2)
        default:
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 uint64 = 2
                for fi20 := 0; fi20 < 9; fi20++ {
                    v21 := GenS1{F0: 100, F1: 1}
                    fmt.Println(v19, v5, v3)
                    fmt.Println(v19)
                    v22 := GenS1{F0: 127, F1: 752}
                    var v23 int16 = (v3 >> ((654) & 63))
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                var v24 int16 = (v3 + (-(v3)))
                _ = fi18
                _ = v19
                _ = v24
            }
        }
    }
    if (-134 < -2147483648) {
        if v5 {
            var v25 int64 = -9223372036854775808
            for fi26 := 0; fi26 < 2; fi26++ {
                if v5 {
                    _ = len(v4)
                    v2 -= int32(v3)
                } else {
                    var v27 string = "a"
                    _ = v27
                }
                if v5 {
                    var v28 string = "hello"
                    _ = v28
                } else {
                    fmt.Println(v25, v3, v3)
                    fmt.Println(v25, fi26)
                    v2 += v2
                }
                _ = fi26
            }
            var v29 uint64 = 0
            _ = v25
            _ = v29
        } else {
            for fi30 := 0; fi30 < 4; fi30++ {
                v31 := GenS1{F0: int8(v2), F1: -1}
                var v32 float64 = -2.25
                switch fi30 {
                case -602:
                    switch "hello" {
                    case "hello":
                        switch (-(fi30)) {
                        case 42:
                            delete(v1, "hello")
                        default:
                            v3 += (v3 >> ((65535) & 63))
                        }
                    }
                case 595:
                    switch fi30 {
                    case -9223372036854775808:
                        switch "a" {
                        case "a":
                            for fi33 := 0; fi33 < 17; fi33++ {
                                v2 -= v2
                                _ = fi33
                            }
                        case "hello":
                            var v34 uint64 = 2
                            _ = v34
                        case "":
                            if v5 {
                                v5 = false
                                for fi35 := 0; fi35 < 15; fi35++ {
                                    if v5 {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v36 := GenS1{F0: -1, F1: 42}
                                        v3 = 42
                                        _ = v36
                                    }
                                    v37 := GenS2{F0: 42, F1: (!(!v5)), F2: 0}
                                    v2 -= -1
                                    _ = fi35
                                    _ = v37
                                }
                                for fi38 := 0; fi38 < 3; fi38++ {
                                    fmt.Println(fi30, v32)
                                    fmt.Println(v3, fi38)
                                    _ = fi38
                                }
                            } else {
                                fmt.Println(v2)
                            }
                        default:
                            var v39 float32 = (3.14159 / 0.1)
                            _ = v39
                        }
                    case 1:
                        v2 -= (v2 >> ((4294967295) & 63))
                    default:
                        var v40 int = (fi30 & (fi30 % (((^(fi30))) | 1)))
                        _ = v40
                    }
                }
                _ = fi30
                _ = v31
                _ = v32
            }
            var v41 int = 0
            var v42 int16 = v3
            _ = v41
            _ = v42
        }
        var v43 uint64 = 1
        v44 := GenS1{F0: 0, F1: -200}
        _ = v43
        _ = v44
    }
    var v45 int16 = -239
    var v46 uint64 = 544
    v5 = (0 >= 42)
    v2 += (-(v2))
    for fi47 := 0; fi47 < 12; fi47++ {
        v48 := GenS3{F0: 402, F1: v5}
        _ = fi47
        _ = v48
    }
    v5 = ((1000 <= v45) || v5)
    var v49 int32 = (-(v2))
    delete(v1, "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v49
    _ = v5
}
