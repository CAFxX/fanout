package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 bool; F2 int16 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint64 }
func genFn0() int64 {
    v4 := GenS1{F0: -1, F1: 127, F2: -392}
    var v5 int32 = -108
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[uint32]int)
    v3 := []int{0, 32768, 1, 2147483647}
    switch -1 {
    case 4294967295:
        v6 := GenS2{F0: 570, F1: 2147483648}
        _ = v6
    }
    if false {
        fmt.Println(v1)
    } else {
        var v7 bool = false
        _ = v7
    }
    if ((-800 != -956) || (944 != v1)) {
        var v8 int64 = -149
        _ = v8
    } else {
        var v9 uint16 = (v1 << ((0) & 63))
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        switch "" {
        case "hello":
            var v11 uint32 = 59
            _ = v11
        case "a":
            var v12 uint = 42
            _ = v12
        case "foo bar":
            var v13 uint = 369
            _ = v13
        default:
            for fi14 := 0; fi14 < 5; fi14++ {
                _ = len(v3)
                var v15 int64 = 9223372036854775807
                fmt.Println(v15, fi14)
                _ = fi14
                _ = v15
            }
        }
        v16 := GenS2{F0: 42, F1: 65535}
        v17 := GenS3{F0: 0}
        _ = fi10
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS3{F0: 0}
        v1 += v1
        v1 = (v1 % ((0) | 1))
        _ = fi18
        _ = v19
    }
    var v20 float64 = 3.14159
    if false {
        if true {
            var v21 int64 = 0
            if (true || (("a" + (((("hello" + ("" + "hello")) + "\n\t") + "a") + "\n\t")) == ("foo bar" + ("" + "hello")))) {
                _ = len(v2)
            }
            _ = v21
        } else {
            fmt.Println(v1, v20)
            v20 -= v20
            if false {
                for fi22 := 0; fi22 < 9; fi22++ {
                    var v23 uint = 506
                    switch 127 {
                    case 9223372036854775807:
                        if true {
                            for fi24 := 0; fi24 < 3; fi24++ {
                                var v25 uint32 = 42
                                var v26 string = ("a" + (("hello" + "\n\t") + "foo bar"))
                                v27 := GenS4{F0: v26, F1: (v25 != v25), F2: 687}
                                _ = fi24
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                            switch fi22 {
                            case 616:
                                switch "\n\t" {
                                case "foo bar":
                                    v28 := v2[65535]
                                    _ = v28
                                case "hello":
                                    if false {
                                        for fi29 := 0; fi29 < 1; fi29++ {
                                            v30 := GenS5{F0: -1, F1: 255}
                                            if (!false) {
                                                var v31 uint = v23
                                                _ = v31
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v32 := GenS2{F0: 683, F1: 815}
                                                _ = v32
                                            }
                                            v23 += 65535
                                            _ = fi29
                                            _ = v30
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                }
                            case -282:
                                v33 := GenS6{F0: 0}
                                _ = v33
                            default:
                                var v34 int32 = 161
                                _ = v34
                            }
                        }
                    case -701:
                        if ((127 == 139) || true) {
                            var v35 uint16 = 1
                            var v36 int16 = -235
                            if ((^(fi22)) <= fi22) {
                                if ((^(v35)) < (v1 << ((255) & 63))) {
                                    var v37 uint8 = 2
                                    _ = v37
                                }
                            }
                            _ = v35
                            _ = v36
                        } else {
                            v1 += (v1 * (v1 - (v1 + 39)))
                            v1 = v1
                        }
                    case 32768:
                        for fi38 := 0; fi38 < 15; fi38++ {
                            var v39 int16 = -357
                            _ = fi38
                            _ = v39
                        }
                    }
                    var v40 int64 = int64(v23)
                    _ = fi22
                    _ = v23
                    _ = v40
                }
                fmt.Println(v20, v1)
            }
        }
    }
    v41 := v2[276]
    var v42 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v41
    _ = v42
}
