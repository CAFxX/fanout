package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS5 struct { F0 int16; F1 int64; F2 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 int32; F1 int32 }

func main() {
    v1 := []string{}
    var v2 int32 = -1
    var v3 uint32 = 403
    var v4 string = "hello"
    v5 := make(map[int]bool)
    switch "a" {
    case "a":
        if false {
            for fi6 := 0; fi6 < 17; fi6++ {
                fmt.Println(v4)
                var v7 uint16 = 32767
                v1 = append(v1, v4)
                _ = fi6
                _ = v7
            }
            v8 := GenS1{F0: 98, F1: "foo bar", F2: 32767}
            _ = v8
        }
    case "hello":
        var v9 int64 = -9223372036854775808
        _ = v9
    case "":
        _ = len(v1)
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v1 = append(v1, "foo bar")
        _ = fi11
    }
    if (65535 != 42) {
        for fi12 := 0; fi12 < 17; fi12++ {
            v4 = ((v4 + v4) + "hello")
            _ = fi12
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (((32767 == 0) || true) || false) {
            v13 := GenS1{F0: 89, F1: v4, F2: 951}
            for fi14 := 0; fi14 < 4; fi14++ {
                for fi15 := 0; fi15 < 8; fi15++ {
                    v3 = 1
                    v2 = 0
                    _ = fi15
                }
                v16 := GenS2{F0: 128}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi14
                _ = v16
            }
            v17 := GenS3{F0: -9223372036854775808, F1: v4}
            _ = v13
            _ = v17
        }
    }
    switch -1 {
    case 495:
        v5[-2147483648] = (v2 <= 42)
    case 141:
        switch 4294967295 {
        case 32767:
            v3 = 2147483648
        case 830:
            if ((32768 > 42) || false) {
                v4 = ("hello" + ((v4 + ((v4 + v4) + v4)) + (v4 + v4)))
            } else {
                v3 += v3
                var v18 uint8 = 95
                if (-9223372036854775808 == 9223372036854775807) {
                    if (819 != 65535) {
                        if true {
                            v3 -= v3
                            fmt.Println(v3)
                        } else {
                            var v19 float32 = 0.0
                            var v20 int64 = 0
                            _ = v19
                            _ = v20
                        }
                    } else {
                        if (1e10 > -2.25) {
                            fmt.Println(v3)
                        }
                        v21 := GenS3{F0: -2147483648, F1: "\n\t"}
                        v22 := GenS4{F0: 0, F1: 2, F2: 42}
                        _ = v21
                        _ = v22
                    }
                    v18 -= v18
                    for fi23 := 0; fi23 < 12; fi23++ {
                        v24 := GenS4{F0: 42, F1: 145, F2: 1}
                        switch (-(fi23)) {
                        case 32768:
                            fmt.Println(fi23, v2)
                        case 1:
                            if true {
                                var v25 string = v4
                                if false {
                                    var v26 int32 = (v2 >> ((1) & 63))
                                    for fi27 := 0; fi27 < 16; fi27++ {
                                        v18 -= 148
                                        var v28 float32 = ((-2.25 - (3.14159 * 1e10)) / ((1.5 / (float32(-2.25) / 1.0)) * (-2.25 + (0.0 * 1.5))))
                                        _ = fi27
                                        _ = v28
                                    }
                                    _ = v26
                                } else {
                                    switch (("foo bar" + ("a" + "hello")) + v25) {
                                    case "hello":
                                        if true {
                                            fmt.Println(v25)
                                            var v29 float32 = (0.0 - 1e10)
                                            fmt.Println(v25)
                                            _ = v29
                                        } else {
                                            var v30 uint32 = v3
                                            _ = v30
                                        }
                                    case "":
                                        if false {
                                            var v31 int16 = -32768
                                            _ = v31
                                        }
                                    case "a":
                                        fmt.Println(v3, v18)
                                    default:
                                        v32 := GenS4{F0: -1, F1: 42, F2: 1}
                                        _ = v32
                                    }
                                    v25 = v25
                                }
                                v33 := GenS2{F0: v18}
                                _ = v25
                                _ = v33
                            }
                        case 2147483647:
                            for fi34 := 0; fi34 < 1; fi34++ {
                                v35 := GenS4{F0: 1, F1: 4294967295, F2: 804}
                                _ = len(v1)
                                _ = fi34
                                _ = v35
                            }
                        }
                        switch (^(fi23)) {
                        case 127:
                            v5[-1] = true
                        case 42:
                            if true {
                                fmt.Println(v3, v18, v3)
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v18 = v18
                                v36 := GenS2{F0: 84}
                                _ = v36
                            }
                        case -392:
                            v37 := GenS1{F0: 45, F1: (v4 + v4), F2: 477}
                            _ = v37
                        }
                        _ = fi23
                        _ = v24
                    }
                } else {
                    if ((42 >= int(v2)) || (42 != 255)) {
                        var v38 int64 = 1
                        _ = v38
                    }
                    v18 = 127
                    v39 := GenS5{F0: 32767, F1: 9223372036854775807, F2: 42}
                    _ = v39
                }
                _ = v18
            }
        default:
            v4 = (("a" + v4) + ("" + ("" + (("" + (((v4 + "foo bar") + "a") + (v4 + "hello"))) + ((v4 + "foo bar") + ((v4 + "hello") + v4))))))
        }
    default:
        _ = len(v5)
    }
    v40 := GenS6{F0: 128}
    v41 := GenS7{F0: 1, F1: 100}
    v4 = "a"
    if (false && (false && true)) {
        v42 := GenS6{F0: 127}
        _ = len(v1)
        if ("" <= v4) {
            var v43 int32 = v2
            var v44 int = -1
            _ = v43
            _ = v44
        }
        _ = v42
    }
    v4 = v4
    for fi45 := 0; fi45 < 14; fi45++ {
        if (2147483647 < fi45) {
            v4 = v4
            v46 := GenS8{F0: v2, F1: (^(v2))}
            v2 -= v2
            _ = v46
        } else {
            for fi47 := 0; fi47 < 18; fi47++ {
                v48 := GenS2{F0: 0}
                v49 := GenS5{F0: -830, F1: 1, F2: 0}
                _ = fi47
                _ = v48
                _ = v49
            }
            switch fi45 {
            case 295:
                v50 := GenS6{F0: 76}
                _ = v50
            case 1:
                fmt.Println(v4, v2, fi45)
            case 127:
                fmt.Println(fi45)
            }
        }
        fmt.Println(fi45)
        for fi51 := 0; fi51 < 16; fi51++ {
            var v52 int16 = 1000
            var v53 int32 = v2
            _ = fi51
            _ = v52
            _ = v53
        }
        _ = fi45
    }
    if (!false) {
        fmt.Println(v3, v4, v2)
        fmt.Println(v3)
        for fi54 := 0; fi54 < 10; fi54++ {
            _ = len(v1)
            _ = fi54
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
