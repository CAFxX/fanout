package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint) uint64 {
    var v4 uint16 = 42
    var v5 int64 = 1
    _ = p0
    _ = v4
    _ = v5
    return 915
}
func genFn1(p0 uint) int32 {
    var v6 int8 = 100
    for fi7 := 0; fi7 < 20; fi7++ {
        if (32767 != 673) {
            var v8 uint8 = 153
            fmt.Println(v8, p0, v6)
            v9 := GenS1{F0: 0}
            _ = v8
            _ = v9
        } else {
            v10 := GenS2{F0: 295, F1: 2}
            v11 := GenS2{F0: 1, F1: p0}
            _ = v10
            _ = v11
        }
        switch "\n\t" {
        case "\n\t":
            var v12 int32 = 1
            _ = v12
        case "hello":
            var v13 float32 = -2.25
            _ = v13
        case "foo bar":
            var v14 int16 = -32768
            _ = v14
        default:
            fmt.Println(v6, v6)
        }
        p0 = p0
        _ = fi7
    }
    if true {
        if (0 <= 149) {
            p0 -= 4294967295
        } else {
            switch 32767 {
            case 42:
                var v15 int = 42
                _ = v15
            case -125:
                fmt.Println(v6, p0)
            case -621:
                if true {
                    switch 4294967295 {
                    case 9223372036854775807:
                        switch -720 {
                        case -2147483648:
                            for fi16 := 0; fi16 < 13; fi16++ {
                                if (!(0.0 < float32(((3.14159 - 0.0) + float64(1.5))))) {
                                    v17 := GenS1{F0: -625}
                                    var v18 int64 = 9223372036854775807
                                    _ = v17
                                    _ = v18
                                }
                                _ = fi16
                            }
                        case -1:
                            var v19 bool = (167 >= 188)
                            _ = v19
                        }
                    default:
                        var v20 string = "foo bar"
                        _ = v20
                    }
                    switch -2147483648 {
                    case 543:
                        switch -2147483648 {
                        case 248:
                            switch "a" {
                            case "foo bar":
                                v21 := []int{-2147483648}
                                _ = v21
                            case "a":
                                p0 = p0
                            default:
                                for fi22 := 0; fi22 < 20; fi22++ {
                                    for fi23 := 0; fi23 < 17; fi23++ {
                                        v24 := GenS2{F0: 0, F1: (p0 % (((p0 >> ((p0) & 63))) | 1))}
                                        _ = fi23
                                        _ = v24
                                    }
                                    for fi25 := 0; fi25 < 12; fi25++ {
                                        v26 := GenS1{F0: -9223372036854775808}
                                        _ = fi25
                                        _ = v26
                                    }
                                    var v27 float64 = 1e10
                                    _ = fi22
                                    _ = v27
                                }
                            }
                        default:
                            p0 += p0
                        }
                    }
                }
            default:
                for fi28 := 0; fi28 < 15; fi28++ {
                    if false {
                        p0 -= (p0 & p0)
                    } else {
                        var v29 int16 = 1000
                        _ = v29
                    }
                    _ = fi28
                }
            }
        }
    }
    v30 := GenS3{F0: 1, F1: 9223372036854775807}
    switch ("foo bar" + "a") {
    case "a":
        v6 += -128
    case "":
        var v31 uint64 = 754
        _ = v31
    }
    _ = p0
    _ = v30
    _ = v6
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{1}
    v3 := make(map[string]string)
    _ = len(v2)
    if true {
        for fi32 := 0; fi32 < 14; fi32++ {
            switch (fi32 | fi32) {
            case 4294967295:
                switch (fi32 + fi32) {
                case -148:
                    var v33 uint = 18446744073709551615
                    _ = v33
                case -9223372036854775808:
                    var v34 string = "\n\t"
                    _ = v34
                case 34:
                    for fi35 := 0; fi35 < 8; fi35++ {
                        v1[2147483647] = false
                        var v36 int64 = -259
                        fmt.Println(fi32, fi32)
                        var v37 float32 = -2.25
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                }
            case 177:
                v38 := v3[""]
                _ = v38
            case -444:
                var v39 int8 = 42
                _ = v39
            }
            var v40 int = fi32
            if (877 == -986) {
                v40 = (v40 << ((264) & 63))
            }
            _ = fi32
            _ = v40
        }
    } else {
        var v41 int32 = -893
        v41 = genFn1(65535)
        v42 := GenS1{F0: -640}
        _ = v41
        _ = v42
    }
    var v43 uint8 = 57
    var v44 int = 0
    v45 := GenS1{F0: 42}
    var v46 int = (v44 << ((18446744073709551615) & 63))
    v47 := GenS2{F0: 42, F1: 42}
    v1[65535] = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
