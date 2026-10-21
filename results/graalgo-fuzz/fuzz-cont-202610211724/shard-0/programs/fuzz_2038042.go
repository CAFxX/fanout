package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int16 }
type GenS2 struct { F0 int32; F1 uint; F2 uint }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 uint; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{9223372036854775807, 2147483647, -292}
    var v3 uint64 = 0
    v4 := make(map[string]string)
    v5 := make(map[uint32]bool)
    v6 := make(map[string]int)
    _ = len(v4)
    if true {
        switch int(v3) {
        case 0:
            v3 -= (v3 >> ((255) & 63))
        default:
            if true {
                if true {
                    v3 += v3
                }
                v2 = append(v2, 127)
            } else {
                var v7 int = 0
                var v8 float64 = 0.1
                var v9 float64 = 3.14159
                _ = v7
                _ = v8
                _ = v9
            }
        }
        var v10 float64 = ((-2.25 - (0.0 + (1e10 / (1e10 + 0.0)))) + 3.14159)
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        if false {
            switch "foo bar" {
            case "a":
                for fi12 := 0; fi12 < 13; fi12++ {
                    var v13 uint8 = 42
                    fmt.Println(v13, fi12)
                    _ = fi12
                    _ = v13
                }
            case "\n\t":
                var v14 int8 = 1
                _ = v14
            case "":
                for fi15 := 0; fi15 < 20; fi15++ {
                    if false {
                        var v16 int8 = 0
                        switch (fi11 & (-(fi15))) {
                        case 642:
                            switch fi15 {
                            case 32768:
                                v3 += 42
                            default:
                                v17 := GenS1{F0: 1, F1: false, F2: 86}
                                _ = v17
                            }
                        }
                        _ = v16
                    } else {
                        v2 = append(v2, 128)
                        fmt.Println(v3)
                        switch 325 {
                        case -1:
                            var v18 string = "hello"
                            _ = v18
                        case 1:
                            if (((0 > 303) && (false || (42 < 0))) || (1 < 0)) {
                                v3 = v3
                            } else {
                                var v19 uint8 = 220
                                if len(v2) > 0 {
                                    v2[0] = fi15
                                }
                                v20 := GenS2{F0: 2147483647, F1: 1, F2: 17}
                                v3 -= v3
                                _ = v19
                                _ = v20
                            }
                        case -557:
                            v3 = v3
                        }
                        v3 = v3
                    }
                    switch fi11 {
                    case -443:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = len(v2)
                        if (1 < 0) {
                            v3 = 42
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                        } else {
                            fmt.Println(v3, fi15, v3)
                        }
                        var v21 int8 = 127
                        var v22 uint = 65535
                        _ = v21
                        _ = v22
                    default:
                        switch "a" {
                        case "foo bar":
                            if (!(!true)) {
                                var v23 int32 = 42
                                _ = v23
                            } else {
                                var v24 uint32 = 2147483647
                                if true {
                                    var v25 int = fi11
                                    var v26 float32 = ((0.1 + (0.0 + ((0.1 + (1e10 - (float32(0.1) / 1.5))) / 1.0))) * 1e10)
                                    _ = v25
                                    _ = v26
                                }
                                var v27 uint8 = 127
                                _ = v24
                                _ = v27
                            }
                        case "":
                            if false {
                                switch "a" {
                                case "foo bar":
                                    var v28 string = ""
                                    _ = v28
                                case "hello":
                                    fmt.Println(v3, fi15)
                                }
                                v29 := GenS3{F0: 0, F1: 32768}
                                v3 -= 42
                                _ = v29
                            }
                        }
                    }
                    if (true || (801 >= 1000)) {
                        v3 += 1
                    }
                    fmt.Println(fi15, v3, fi11)
                    v30 := GenS4{F0: (fi11 & (fi11 + fi11))}
                    var v31 uint64 = (v3 >> ((1) & 63))
                    _ = fi15
                    _ = v30
                    _ = v31
                }
            }
            if true {
                for fi32 := 0; fi32 < 6; fi32++ {
                    _ = len(v1)
                    v3 -= (^(v3))
                    v3 -= v3
                    _ = fi32
                }
                v3 = (v3 | (v3 + 1))
            }
        } else {
            if true {
                var v33 uint32 = 2
                _ = len(v6)
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 10; fi34++ {
            v35 := GenS5{F0: 2}
            v3 = 4294967295
            _ = fi34
            _ = v35
        }
        _ = fi11
    }
    var v36 float64 = (((1.5 + 0.0) / (1e10 + 1e10)) - float64(float32(0.0)))
    if true {
        if len(v2) > 0 {
            v2[0] = 40
        }
    }
    fmt.Println(v36, v3)
    v3 = v3
    switch "foo bar" {
    case "\n\t":
        switch "hello" {
        case "a":
            v37 := v4["foo bar"]
            _ = v37
        case "foo bar":
            v38 := GenS6{F0: 1, F1: 728, F2: 255}
            _ = v38
        case "hello":
            if false {
                switch "a" {
                case "a":
                    switch 978 {
                    case -9223372036854775808:
                        v3 += 4294967295
                    case 0:
                        _ = len(v2)
                    case 471:
                        v3 -= v3
                    }
                case "":
                    var v39 uint64 = (^(v3))
                    _ = v39
                }
                var v40 uint8 = 239
                _ = v40
            } else {
                if false {
                    var v41 float64 = -2.25
                    _ = v41
                } else {
                    var v42 int16 = -32768
                    for fi43 := 0; fi43 < 2; fi43++ {
                        var v44 string = ((("hello" + "a") + "") + "\n\t")
                        var v45 uint16 = 0
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    fmt.Println(v3, v42, v3)
                    _ = v42
                }
                delete(v6, "foo bar")
            }
        default:
            var v46 int16 = 361
            _ = v46
        }
    case "a":
        v36 += 1.5
    case "hello":
        var v47 float64 = (0.1 - float64(float32((0.1 - (v36 / v36)))))
        _ = v47
    }
    v48 := GenS5{F0: 620}
    var v49 int32 = 990
    switch 2147483647 {
    case -9223372036854775808:
        if (!(32767 < 1)) {
            fmt.Println(v36)
            fmt.Println(v49, v36)
            v50 := GenS2{F0: v49, F1: 504, F2: 2}
            _ = v50
        }
    case -128:
        for fi51 := 0; fi51 < 1; fi51++ {
            v36 += 3.14159
            v52 := GenS2{F0: 0, F1: 2, F2: 2}
            for fi53 := 0; fi53 < 11; fi53++ {
                var v54 uint32 = 2147483647
                var v55 string = "foo bar"
                _ = fi53
                _ = v54
                _ = v55
            }
            _ = fi51
            _ = v52
        }
    }
    _ = len(v2)
    fmt.Println(v3, v49)
    for fi56 := 0; fi56 < 13; fi56++ {
        v49 += 2147483647
        _ = fi56
    }
    var v57 uint64 = (^(v3))
    if len(v2) > 0 {
        v2[0] = int(v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v57
    _ = v6
}
