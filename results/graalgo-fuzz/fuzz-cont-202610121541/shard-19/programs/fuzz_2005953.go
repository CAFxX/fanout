package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 bool; F2 int16 }
type GenS5 struct { F0 uint8; F1 string; F2 string }
type GenS6 struct { F0 int32 }

func main() {
    var v1 uint = 718
    v2 := []string{"", "hello", "a"}
    v3 := []string{"foo bar", (("\n\t" + "") + "hello"), ("foo bar" + "hello")}
    var v4 int64 = -689
    v5 := make(map[uint32]bool)
    switch "" {
    case "foo bar":
        switch "a" {
        case "hello":
            switch "a" {
            case "foo bar":
                if (!(10 > -2147483648)) {
                    v1 -= (^(v1))
                } else {
                    if (0 <= 799) {
                        if true {
                            v5[65535] = true
                            fmt.Println(v1)
                            v1 += (^(v1))
                        }
                        for fi6 := 0; fi6 < 8; fi6++ {
                            delete(v5, 162)
                            for fi7 := 0; fi7 < 4; fi7++ {
                                var v8 uint16 = 0
                                switch "\n\t" {
                                case "a":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                _ = len(v2)
                                _ = fi7
                                _ = v8
                            }
                            _ = fi6
                        }
                        var v9 int = 32768
                        _ = v9
                    }
                    var v10 uint64 = 2
                    _ = v10
                }
            case "\n\t":
                v1 += v1
            default:
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 uint64 = 715
                    if (false || (true && false)) {
                        v13 := GenS1{F0: 63}
                        v14 := GenS2{F0: 32767}
                        switch ("" + ("" + "\n\t")) {
                        case "\n\t":
                            v15 := GenS1{F0: 127}
                            _ = v15
                        case "hello":
                            switch "\n\t" {
                            case "\n\t":
                                switch "hello" {
                                case "\n\t":
                                    var v16 float64 = ((1e10 + 0.0) + 1e10)
                                    _ = v16
                                case "foo bar":
                                    var v17 uint8 = 255
                                    _ = v17
                                }
                            case "foo bar":
                                var v18 string = ("foo bar" + (("\n\t" + ("a" + "\n\t")) + ("foo bar" + ("a" + "foo bar"))))
                                _ = v18
                            }
                        case "foo bar":
                            v12 = (v12 | (v12 - v12))
                        default:
                            var v19 string = ("" + "hello")
                            _ = v19
                        }
                        _ = v13
                        _ = v14
                    }
                    var v20 int8 = int8(fi11)
                    _ = fi11
                    _ = v12
                    _ = v20
                }
            }
        }
    case "":
        v21 := GenS1{F0: -128}
        _ = v21
    }
    v4 = (v4 >> ((v1) & 63))
    if true {
        v22 := GenS3{F0: 468}
        var v23 uint64 = 4294967295
        _ = v22
        _ = v23
    } else {
        v4 -= 0
        var v24 string = ("hello" + "foo bar")
        _ = v24
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v1, v4)
    case "a":
        fmt.Println(v1)
    case "":
        v1 -= 2
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        for fi26 := 0; fi26 < 6; fi26++ {
            if (2 <= 2) {
                var v27 bool = (1e10 != 1.5)
                v2 = append(v2, (("foo bar" + "hello") + "hello"))
                _ = v27
            } else {
                var v28 uint = (v1 * v1)
                var v29 int8 = 1
                var v30 float64 = 3.14159
                fmt.Println(v30)
                v28 = 524
                v3 = append(v3, "hello")
                fmt.Println(fi26, v29)
                v1 += v1
                switch 814 {
                case 4294967295:
                    v30 = ((1e10 - v30) - (0.1 * 0.1))
                case 0:
                    v31 := GenS1{F0: v29}
                    _ = v31
                case -2147483648:
                    switch fi25 {
                    case -806:
                        v29 -= (v29 + v29)
                    }
                }
                _ = v28
                _ = v29
                _ = v30
            }
            _ = len(v3)
            _ = fi26
        }
        v32 := GenS4{F0: 743, F1: ((1.5 > 0.1) && false), F2: 0}
        _ = fi25
        _ = v32
    }
    v33 := GenS5{F0: 0, F1: "hello", F2: "\n\t"}
    for fi34 := 0; fi34 < 10; fi34++ {
        v35 := GenS3{F0: 2147483647}
        var v36 float32 = (0.1 * 1.5)
        _ = fi34
        _ = v35
        _ = v36
    }
    delete(v5, 65535)
    v37 := GenS5{F0: 1, F1: "foo bar", F2: "\n\t"}
    var v38 int16 = -234
    var v39 string = "a"
    for fi40 := 0; fi40 < 4; fi40++ {
        var v41 int8 = -1
        for fi42 := 0; fi42 < 3; fi42++ {
            fmt.Println(fi40)
            _ = fi42
        }
        for fi43 := 0; fi43 < 14; fi43++ {
            switch (("\n\t" + "a") + v39) {
            case "a":
                var v44 int64 = (v4 >> ((721) & 63))
                _ = v44
            default:
                var v45 uint16 = 1
                _ = v45
            }
            switch ("hello" + v39) {
            case "\n\t":
                var v46 float64 = 0.0
                v46 = float64(float32(v46))
                _ = v46
            case "":
                v47 := GenS1{F0: (v41 >> (((v1 ^ (v1 << ((v1) & 63)))) & 63))}
                _ = v47
            case "foo bar":
                v48 := GenS6{F0: -4}
                _ = v48
            }
            var v49 string = ("\n\t" + v39)
            _ = fi43
            _ = v49
        }
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
