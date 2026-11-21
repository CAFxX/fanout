package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 774
    v3 := []int{4294967295, 127, -573}
    var v4 string = "\n\t"
    for fi5 := 0; fi5 < 15; fi5++ {
        _ = len(v3)
        if len(v3) > 0 {
            v3[0] = (fi5 >> ((0) & 63))
        }
        v2 += 42
        _ = fi5
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        for fi6 := 0; fi6 < 16; fi6++ {
            for fi7 := 0; fi7 < 7; fi7++ {
                v2 += (v2 / (((v2 << ((255) & 63))) | 1))
                fmt.Println(v4, fi7, fi7)
                v2 = v2
                if (!(2 == 391)) {
                    if false {
                        v8 := GenS1{F0: 42, F1: -862}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v8
                    }
                    fmt.Println(fi7, v4)
                } else {
                    var v9 int = fi7
                    _ = v9
                }
                delete(v1, 2)
                _ = fi7
            }
            _ = fi6
        }
    }
    var v10 bool = false
    v10 = (0 == 42)
    fmt.Println(v10)
    var v11 float32 = 1e10
    for fi12 := 0; fi12 < 15; fi12++ {
        switch v4 {
        case "":
            switch v4 {
            case "":
                if v10 {
                    v13 := v1[2147483647]
                    _ = v13
                } else {
                    var v14 int16 = 879
                    var v15 int16 = (v2 ^ v14)
                    var v16 uint8 = 42
                    _ = v14
                    _ = v15
                    _ = v16
                }
            case "foo bar":
                switch fi12 {
                case 42:
                    v17 := GenS2{F0: 4294967295}
                    _ = v17
                case 4294967295:
                    if len(v3) > 0 {
                        v3[0] = (^(fi12))
                    }
                default:
                    v18 := GenS3{F0: 42, F1: v10, F2: 255}
                    _ = v18
                }
            }
        case "foo bar":
            if (80 != 1) {
                v10 = (!(!(493 >= 4294967295)))
                for fi19 := 0; fi19 < 20; fi19++ {
                    if ((!(v10 && false)) || false) {
                        switch 437 {
                        case 42:
                            _ = len(v1)
                        case -9223372036854775808:
                            var v20 int16 = v2
                            _ = v20
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        v21 := GenS4{F0: 175}
                        switch "\n\t" {
                        case "foo bar":
                            v22 := GenS5{F0: 811, F1: 2147483647}
                            _ = v22
                        case "a":
                            if v10 {
                                var v23 string = (v4 + v4)
                                if (fi12 == (fi19 & 127)) {
                                    var v24 uint16 = 32767
                                    v25 := GenS6{F0: 1}
                                    _ = v24
                                    _ = v25
                                } else {
                                    fmt.Println(v23)
                                    v11 = 1e10
                                }
                                v10 = v10
                                v1[0] = 325
                                _ = v23
                            }
                        }
                        _ = v21
                    }
                    _ = fi19
                }
            }
        }
        v4 = "a"
        _ = fi12
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 string = (v4 + "hello")
        var v28 uint64 = 0
        var v29 int16 = (v2 - 32767)
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    v30 := GenS5{F0: 4294967295, F1: 858}
    if (158 != 255) {
        fmt.Println(v11, v10)
        for fi31 := 0; fi31 < 5; fi31++ {
            if v10 {
                var v32 uint = 1
                var v33 uint16 = 0
                for fi34 := 0; fi34 < 17; fi34++ {
                    v33 += v33
                    _ = fi34
                }
                _ = v32
                _ = v33
            }
            v35 := GenS3{F0: 309, F1: true, F2: 0}
            v36 := GenS1{F0: 111, F1: -9223372036854775808}
            _ = fi31
            _ = v35
            _ = v36
        }
    } else {
        var v37 uint8 = 78
        v11 -= v11
        for fi38 := 0; fi38 < 15; fi38++ {
            for fi39 := 0; fi39 < 5; fi39++ {
                var v40 float32 = 3.14159
                var v41 uint = 940
                var v42 uint8 = (v37 ^ (v37 | 127))
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
            }
            _ = fi38
        }
        _ = v37
    }
    v43 := v1[42]
    var v44 uint32 = (v43 >> ((0) & 63))
    var v45 float32 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
