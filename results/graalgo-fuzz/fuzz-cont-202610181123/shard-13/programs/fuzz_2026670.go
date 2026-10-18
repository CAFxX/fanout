package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 0
    var v3 uint16 = 65535
    var v4 int16 = 32767
    var v5 float32 = -2.25
    v6 := []int{-9223372036854775808, 127, 521}
    var v7 float64 = 1e10
    if ((v2 + v2) > v2) {
        var v8 uint16 = v3
        for fi9 := 0; fi9 < 15; fi9++ {
            fmt.Println(v2, v8, fi9)
            _ = fi9
        }
        var v10 string = ""
        _ = v10
        _ = v8
    } else {
        var v11 uint16 = v3
        var v12 uint8 = 128
        switch 127 {
        case -281:
            _ = len(v6)
        case 1:
            v13 := GenS1{F0: 960, F1: v4}
            _ = v13
        case -2147483648:
            for fi14 := 0; fi14 < 10; fi14++ {
                if true {
                    fmt.Println(v5, v5, v3)
                } else {
                    v15 := GenS2{F0: 1, F1: "\n\t", F2: (^(v12))}
                    v7 -= (v7 + (0.0 - v7))
                    _ = v15
                }
                _ = fi14
            }
        default:
            var v16 float64 = v7
            _ = v16
        }
        _ = v11
        _ = v12
    }
    switch "hello" {
    case "a":
        if false {
            if true {
                fmt.Println(v4, v2)
                v5 -= 3.14159
            } else {
                if false {
                    switch ("foo bar" + "") {
                    case "\n\t":
                        _ = len(v6)
                    case "a":
                        for fi17 := 0; fi17 < 7; fi17++ {
                            v3 += (v3 & (v3 % ((255) | 1)))
                            var v18 uint8 = 255
                            _ = fi17
                            _ = v18
                        }
                    case "hello":
                        for fi19 := 0; fi19 < 5; fi19++ {
                            switch (fi19 + (fi19 >> ((0) & 63))) {
                            case -2147483648:
                                var v20 bool = false
                                _ = v20
                            }
                            if (!(false || true)) {
                                switch "hello" {
                                case "foo bar":
                                    v2 -= (v2 ^ 0)
                                }
                                var v21 int8 = 1
                                fmt.Println(v5)
                                _ = v21
                            } else {
                                fmt.Println(v2, fi19)
                                v7 = v7
                                for fi22 := 0; fi22 < 2; fi22++ {
                                    v23 := GenS3{F0: true}
                                    _ = fi22
                                    _ = v23
                                }
                            }
                            _ = fi19
                        }
                    }
                }
            }
            for fi24 := 0; fi24 < 2; fi24++ {
                v4 = 789
                var v25 int32 = 499
                _ = fi24
                _ = v25
            }
        } else {
            var v26 uint32 = 42
            var v27 bool = false
            v28 := GenS3{F0: v27}
            _ = v26
            _ = v27
            _ = v28
        }
    case "":
        v29 := GenS1{F0: 0, F1: v4}
        _ = v29
    case "\n\t":
        v30 := GenS4{F0: 42}
        _ = v30
    default:
        if (!(32768 <= 1)) {
            if (!(!false)) {
                var v31 uint16 = (v3 - (v3 ^ v3))
                _ = v31
            } else {
                if (false || true) {
                    var v32 int = 127
                    var v33 uint32 = 65535
                    _ = v32
                    _ = v33
                }
                fmt.Println(v4)
            }
            if true {
                v2 -= (^(v2))
                switch ("a" + "") {
                case "\n\t":
                    v3 = v3
                }
                var v34 uint16 = 322
                _ = v34
            } else {
                v3 = (v3 + 42)
                var v35 uint = 591
                _ = v35
            }
        }
    }
    v2 = v2
    fmt.Println(v2)
    var v36 float64 = v7
    v37 := GenS2{F0: (v3 >> ((uint(v2)) & 63)), F1: "\n\t", F2: 1}
    for fi38 := 0; fi38 < 19; fi38++ {
        switch "hello" {
        case "hello":
            fmt.Println(v2)
        case "\n\t":
            for fi39 := 0; fi39 < 14; fi39++ {
                var v40 string = ("a" + "hello")
                if (18 <= -1) {
                    if false {
                        fmt.Println(v40, v7)
                        var v41 int = fi38
                        _ = v41
                    }
                }
                _ = fi39
                _ = v40
            }
        }
        _ = fi38
    }
    v36 -= v7
    for fi42 := 0; fi42 < 4; fi42++ {
        switch (fi42 & fi42) {
        case 42:
            v2 = (v2 % ((uint64(v3)) | 1))
        case -1:
            var v43 float64 = (0.1 - 0.0)
            _ = v43
        default:
            v44 := GenS2{F0: v3, F1: "foo bar", F2: 49}
            _ = v44
        }
        _ = fi42
    }
    switch 32767 {
    case -2147483648:
        var v45 float32 = 0.0
        _ = v45
    }
    for fi46 := 0; fi46 < 2; fi46++ {
        var v47 int = (fi46 + (fi46 - (^(fi46))))
        v5 -= (v5 - v5)
        v5 = (0.0 + (float32(3.14159) / v5))
        _ = fi46
        _ = v47
    }
    fmt.Println(v3, v7)
    fmt.Println(v7, v4, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
