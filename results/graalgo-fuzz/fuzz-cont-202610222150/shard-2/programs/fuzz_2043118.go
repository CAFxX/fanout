package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    var v4 uint64 = 0
    v1[2147483647] = 42
    switch "\n\t" {
    case "":
        v4 -= v4
    case "foo bar":
        for fi5 := 0; fi5 < 8; fi5++ {
            v6 := GenS1{F0: 0, F1: false}
            var v7 int32 = -1
            _ = fi5
            _ = v6
            _ = v7
        }
    case "a":
        if false {
            v8 := GenS2{F0: 32767, F1: 0}
            _ = v8
        }
    }
    fmt.Println(v4)
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 int32 = 636
        v4 -= 1
        _ = fi9
        _ = v10
    }
    switch -563 {
    case -2147483648:
        if (1 != 304) {
            v4 = (v4 * (v4 | v4))
        } else {
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 string = "hello"
                v1[883] = 42
                var v13 uint64 = v4
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 int8 = 0
            _ = v14
        }
    default:
        var v15 int = -9223372036854775808
        _ = v15
    }
    v16 := GenS2{F0: 0, F1: 255}
    v17 := GenS1{F0: 2147483648, F1: ((!(4294967295 < 0)) && (9223372036854775807 < 42))}
    v4 = 1
    var v18 uint32 = 0
    if true {
        var v19 uint = 27
        _ = v19
    } else {
        var v20 string = "\n\t"
        v21 := GenS1{F0: (v18 ^ 65535), F1: true}
        fmt.Println(v18)
        _ = v20
        _ = v21
    }
    if false {
        var v22 uint64 = v4
        _ = len(v1)
        _ = v22
    } else {
        v23 := []int{2147483647, -9223372036854775808, -990, 9223372036854775807}
        _ = v23
    }
    if true {
        v24 := []int{32767}
        switch 128 {
        case -362:
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 int32 = 42
                v18 += v18
                if (!((fi25 ^ fi25) > (fi25 + -194))) {
                    if (!false) {
                        switch "foo bar" {
                        case "hello":
                            if true {
                                fmt.Println(v26, v26)
                            } else {
                                v27 := GenS2{F0: 154, F1: 42}
                                v4 += 2
                                for fi28 := 0; fi28 < 1; fi28++ {
                                    v29 := GenS3{F0: (v4 ^ v4), F1: ("foo bar" + "a")}
                                    if false {
                                        var v30 uint64 = (v4 << ((255) & 63))
                                        var v31 int16 = 42
                                        _ = v30
                                        _ = v31
                                    } else {
                                        if (0 >= 555) {
                                            switch -931 {
                                            case 32768:
                                                fmt.Println(v26, fi28, v4)
                                            default:
                                                switch ((("hello" + "") + "") + "") {
                                                case "hello":
                                                    v26 -= (v26 ^ v26)
                                                }
                                            }
                                            var v32 uint64 = v4
                                            if false {
                                                v32 = (^(v4))
                                                switch "" {
                                                case "foo bar":
                                                    v26 += 1
                                                case "":
                                                    if ((!false) && (521 < 255)) {
                                                        v18 += (v18 & 2147483647)
                                                    } else {
                                                        v26 += (v26 << ((629) & 63))
                                                        var v33 float32 = -2.25
                                                        _ = v33
                                                    }
                                                case "\n\t":
                                                    v34 := GenS4{F0: 0, F1: 0}
                                                    _ = v34
                                                }
                                            }
                                            _ = v32
                                        } else {
                                            var v35 string = "foo bar"
                                            v4 += 910
                                            v36 := GenS5{F0: (v18 ^ 65535), F1: 127}
                                            _ = v35
                                            _ = v36
                                        }
                                    }
                                    var v37 int8 = -1
                                    _ = fi28
                                    _ = v29
                                    _ = v37
                                }
                                _ = v27
                            }
                        }
                    } else {
                        for fi38 := 0; fi38 < 17; fi38++ {
                            switch "\n\t" {
                            case "\n\t":
                                v3[fi38] = true
                            }
                            _ = fi38
                        }
                        var v39 int16 = 0
                        _ = v39
                    }
                }
                _ = fi25
                _ = v26
            }
        }
        v18 -= (^(v18))
        _ = v24
    }
    v2[0] = ("" == "hello")
    v18 = (v18 & v18)
    var v40 string = ("foo bar" + "foo bar")
    var v41 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
