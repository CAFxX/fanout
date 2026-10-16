package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 string = "hello"
    v3 := []int{-9223372036854775808}
    var v4 float64 = float64(0.1)
    var v5 int = 128
    v6 := []uint32{42}
    var v7 uint16 = 255
    if (!true) {
        v8 := GenS1{F0: 1}
        v2 = ""
        _ = v8
    } else {
        switch ("" + (v2 + "foo bar")) {
        case "a":
            var v9 bool = true
            _ = v9
        case "":
            v5 = v5
        case "hello":
            v7 = 220
        default:
            v7 += v7
        }
    }
    v4 = v4
    var v10 uint16 = 255
    for fi11 := 0; fi11 < 18; fi11++ {
        if true {
            v4 -= v4
            var v12 float32 = float32(v4)
            v13 := GenS1{F0: 4294967295}
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    switch "\n\t" {
    case "a":
        switch v2 {
        case "hello":
            v14 := GenS1{F0: 0}
            _ = v14
        case "":
            if ("\n\t" >= "hello") {
                if (1 < 42) {
                    if (true || false) {
                        var v15 float64 = v4
                        v5 += v5
                        for fi16 := 0; fi16 < 11; fi16++ {
                            fmt.Println(v2, v15)
                            v10 = v7
                            fmt.Println(fi16)
                            _ = fi16
                        }
                        _ = v15
                    } else {
                        if (965 > 1) {
                            v2 = "hello"
                        }
                        var v17 uint32 = 2
                        switch v5 {
                        case 644:
                            v5 = v5
                        case 42:
                            v18 := GenS1{F0: 18446744073709551615}
                            _ = v18
                        case -1:
                            if false {
                                for fi19 := 0; fi19 < 6; fi19++ {
                                    if true {
                                        if false {
                                            v7 += v10
                                            fmt.Println(v17, v5)
                                        } else {
                                            if ((-(fi19)) == -1) {
                                                fmt.Println(v4, fi19)
                                            } else {
                                                v10 -= v10
                                            }
                                        }
                                        v20 := GenS1{F0: 42}
                                        var v21 float64 = 1.5
                                        _ = v20
                                        _ = v21
                                    } else {
                                        var v22 uint32 = v17
                                        _ = v22
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        for fi23 := 0; fi23 < 16; fi23++ {
                                            switch (v2 + (v2 + "")) {
                                            case "hello":
                                                _ = len(v1)
                                            case "":
                                                var v24 string = v2
                                                _ = v24
                                            case "\n\t":
                                                var v25 float64 = v4
                                                _ = v25
                                            }
                                            v3 = append(v3, 128)
                                            _ = fi23
                                        }
                                    case "\n\t":
                                        v10 = 392
                                    }
                                    _ = fi19
                                }
                                var v26 uint8 = 255
                                _ = len(v1)
                                _ = v26
                            }
                        default:
                            var v27 int8 = 100
                            _ = v27
                        }
                        _ = v17
                    }
                    v3 = append(v3, v5)
                    var v28 int = v5
                    _ = v28
                }
                v29 := GenS1{F0: 286}
                _ = v29
            }
        case "\n\t":
            v10 += v10
        default:
            var v30 int = v5
            _ = v30
        }
    }
    delete(v1, v5)
    for fi31 := 0; fi31 < 5; fi31++ {
        if (1 <= 1) {
            var v32 float32 = (0.0 + 0.1)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            for fi33 := 0; fi33 < 17; fi33++ {
                switch v2 {
                case "a":
                    var v34 float32 = v32
                    _ = v34
                case "hello":
                    if ((v5 & v5) > (-(fi33))) {
                        var v35 uint = 514
                        _ = v35
                    }
                case "\n\t":
                    var v36 int32 = -748
                    _ = v36
                default:
                    v32 = v32
                }
                switch v2 {
                case "a":
                    v10 += v7
                default:
                    v2 = ("foo bar" + "")
                }
                v5 = fi31
                _ = fi33
            }
            _ = v32
        }
        if (-1 >= 1) {
            var v37 uint16 = (v10 | 42)
            _ = v37
        }
        if ((float64(3.14159) + v4) <= (v4 / v4)) {
            if ((fi31 * 9223372036854775807) >= (v5 % ((fi31) | 1))) {
                v38 := GenS2{F0: 42, F1: v10}
                _ = v38
            } else {
                for fi39 := 0; fi39 < 3; fi39++ {
                    var v40 int8 = 42
                    v40 = (-(v40))
                    _ = fi39
                    _ = v40
                }
                for fi41 := 0; fi41 < 3; fi41++ {
                    v42 := GenS3{F0: -2147483648}
                    var v43 int64 = -1
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            }
        }
        _ = fi31
    }
    var v44 string = v2
    v5 = v5
    v1[(v5 >> ((2) & 63))] = v5
    if ("a" <= v44) {
        _ = len(v1)
        switch 4294967295 {
        case -559:
            var v45 bool = true
            _ = v45
        case -9223372036854775808:
            v10 -= v7
        case 1:
            v10 = v7
        }
    } else {
        for fi46 := 0; fi46 < 13; fi46++ {
            v5 = (v5 >> ((805) & 63))
            _ = fi46
        }
        switch (v5 << ((920) & 63)) {
        case -653:
            var v47 string = "foo bar"
            _ = v47
        case 1:
            v48 := GenS1{F0: 4294967295}
            _ = v48
        default:
            v49 := GenS4{F0: 765, F1: 1}
            _ = v49
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
