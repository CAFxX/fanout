package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 bool }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint = 255
    var v2 int64 = 9223372036854775807
    v3 := []uint32{162}
    v4 := make(map[int]int)
    v5 := GenS1{F0: "foo bar", F1: -111, F2: false}
    v6 := GenS2{F0: 42}
    var v7 string = "\n\t"
    var v8 float64 = (-2.25 / float64(-2.25))
    switch -622 {
    case 919:
        if false {
            for fi9 := 0; fi9 < 4; fi9++ {
                switch "hello" {
                case "a":
                    fmt.Println(v2, v1, v8)
                case "hello":
                    switch fi9 {
                    case 42:
                        var v10 float32 = 1e10
                        _ = v10
                    case -9223372036854775808:
                        fmt.Println(v1, v2, v1)
                    case 787:
                        var v11 uint64 = 2
                        _ = v11
                    }
                case "foo bar":
                    if false {
                        if (!true) {
                            var v12 uint32 = 1
                            _ = v12
                        } else {
                            v4[(-(fi9))] = -1
                            if true {
                                if (1 != 32767) {
                                    if false {
                                        v13 := v4[fi9]
                                        _ = v13
                                    }
                                    var v14 float32 = -2.25
                                    _ = v14
                                } else {
                                    for fi15 := 0; fi15 < 4; fi15++ {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v16 int32 = int32(fi9)
                                        v17 := GenS2{F0: -32768}
                                        var v18 uint16 = 65535
                                        _ = fi15
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                    }
                                    fmt.Println(fi9)
                                }
                            } else {
                                v8 = float64(3.14159)
                            }
                            var v19 float64 = -2.25
                            _ = v19
                        }
                        if (fi9 <= int(v1)) {
                            var v20 float32 = 1e10
                            switch "hello" {
                            case "hello":
                                for fi21 := 0; fi21 < 12; fi21++ {
                                    var v22 float64 = v8
                                    switch 0 {
                                    case -527:
                                        switch 2147483647 {
                                        case 127:
                                            var v23 bool = true
                                            _ = v23
                                        case 42:
                                            var v24 uint16 = 756
                                            _ = v24
                                        case -254:
                                            if true {
                                                _ = len(v4)
                                                v25 := GenS2{F0: 1000}
                                                var v26 int = fi9
                                                v3 = append(v3, 854)
                                                _ = v25
                                                _ = v26
                                            }
                                        }
                                    }
                                    _ = fi21
                                    _ = v22
                                }
                            default:
                                if (0 != 32767) {
                                    for fi27 := 0; fi27 < 6; fi27++ {
                                        v8 = v8
                                        _ = fi27
                                    }
                                    for fi28 := 0; fi28 < 19; fi28++ {
                                        if (65535 == 2) {
                                            v20 -= ((v20 + (0.1 - 1e10)) * v20)
                                        } else {
                                            v29 := GenS2{F0: 0}
                                            var v30 uint16 = 0
                                            v7 = v7
                                            _ = v29
                                            _ = v30
                                        }
                                        v31 := GenS2{F0: -32768}
                                        if (!false) {
                                            var v32 bool = (3.14159 >= v20)
                                            var v33 bool = v32
                                            if (fi28 > (fi9 % ((fi28) | 1))) {
                                                var v34 uint = 980
                                                _ = v34
                                            }
                                            _ = v32
                                            _ = v33
                                        }
                                        _ = fi28
                                        _ = v31
                                    }
                                }
                            }
                            v2 = v2
                            _ = v20
                        } else {
                            v2 += 1
                            for fi35 := 0; fi35 < 12; fi35++ {
                                v8 -= 0.1
                                _ = fi35
                            }
                        }
                        for fi36 := 0; fi36 < 10; fi36++ {
                            var v37 float64 = -2.25
                            v4[(fi36 * (-(fi36)))] = fi9
                            _ = fi36
                            _ = v37
                        }
                    } else {
                        var v38 int64 = (v2 << (((v1 | v1)) & 63))
                        _ = v38
                    }
                }
                v4[(-(fi9))] = fi9
                _ = fi9
            }
            fmt.Println(v2, v7, v7)
        } else {
            if (false || (v2 < (v2 << (((v1 << (((v1 ^ 4)) & 63))) & 63)))) {
                v7 = v7
            } else {
                var v39 int64 = 1
                _ = v39
            }
            v3 = append(v3, 562)
        }
    }
    var v40 string = ("a" + "foo bar")
    v2 = v2
    v8 -= v8
    v1 -= 8
    var v41 uint = v1
    v41 = 0
    var v42 uint64 = 0
    v43 := v4[-175]
    v40 = v7
    switch v43 {
    case 4294967295:
        var v44 uint64 = v42
        _ = v44
    case 42:
        var v45 string = "\n\t"
        _ = v45
    }
    if true {
        v2 = 0
    } else {
        for fi46 := 0; fi46 < 1; fi46++ {
            v40 = (v40 + ("\n\t" + v7))
            v41 = v1
            _ = fi46
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
