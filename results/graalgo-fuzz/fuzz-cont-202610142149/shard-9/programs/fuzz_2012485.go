package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 string; F1 uint64; F2 int64 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 int32; F1 uint8; F2 uint8 }

func main() {
    var v1 string = "\n\t"
    v2 := []string{"\n\t", v1, v1, "a"}
    var v3 string = v1
    v4 := GenS1{F0: 4294967295, F1: 130, F2: 201}
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 uint = 305
        _ = fi5
        _ = v6
    }
    var v7 string = v1
    if true {
        switch v7 {
        case "":
            switch -1 {
            case 42:
                for fi8 := 0; fi8 < 10; fi8++ {
                    if true {
                        fmt.Println(v7, fi8, v7)
                        var v9 uint16 = 65535
                        var v10 uint64 = 2
                        _ = v10
                        _ = v9
                    } else {
                        if true {
                            v3 = (v7 + "hello")
                            for fi11 := 0; fi11 < 14; fi11++ {
                                fmt.Println(v7, v3)
                                v7 = v1
                                var v12 bool = false
                                var v13 uint16 = 149
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                        } else {
                            fmt.Println(v1)
                            switch fi8 {
                            case 4294967295:
                                v14 := GenS1{F0: 2, F1: 255, F2: 0}
                                _ = v14
                            default:
                                switch "" {
                                case "\n\t":
                                    v15 := GenS2{F0: 40, F1: 2}
                                    _ = v15
                                default:
                                    var v16 uint8 = 127
                                    _ = v16
                                }
                            }
                        }
                    }
                    v7 = ((v3 + "foo bar") + v3)
                    v7 = (v7 + (("a" + ("a" + (((v7 + "foo bar") + v1) + v7))) + v1))
                    _ = fi8
                }
            case 127:
                fmt.Println(v1, v7, v1)
            case 225:
                switch "" {
                case "foo bar":
                    switch v3 {
                    case "":
                        if len(v2) > 0 {
                            v2[0] = v3
                        }
                    case "a":
                        v2 = append(v2, v7)
                    }
                case "\n\t":
                    fmt.Println(v7, v7, v7)
                case "hello":
                    if len(v2) > 0 {
                        v2[0] = (v1 + v3)
                    }
                default:
                    fmt.Println(v7, v7, v1)
                }
            default:
                v7 = v3
            }
        case "a":
            var v17 int = 1
            _ = v17
        default:
            var v18 int32 = -730
            _ = v18
        }
    }
    v1 = v7
    if (42 < 0) {
        var v19 uint64 = 2
        var v20 int = -391
        _ = v19
        _ = v20
    } else {
        var v21 int8 = 20
        _ = v21
    }
    var v22 float64 = 1.5
    if (4294967295 < 672) {
        var v23 uint8 = 69
        v3 = "a"
        v1 = v3
        _ = v23
    } else {
        if false {
            switch ("" + "") {
            case "a":
                v1 = (("a" + (v3 + (v1 + "foo bar"))) + (v3 + (v7 + (v1 + v1))))
            case "":
                fmt.Println(v22, v7)
            case "foo bar":
                switch (("\n\t" + v7) + "a") {
                case "a":
                    v24 := GenS3{F0: ((("a" + v1) + v3) + "")}
                    _ = v24
                case "":
                    for fi25 := 0; fi25 < 1; fi25++ {
                        v22 += ((1e10 - v22) + (v22 / v22))
                        var v26 int32 = 2147483647
                        v27 := GenS4{F0: 1, F1: 339}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                case "\n\t":
                    v22 -= 3.14159
                default:
                    if (true && (1 > 65535)) {
                        switch ((v3 + v1) + "hello") {
                        case "\n\t":
                            fmt.Println(v7, v3, v22)
                        default:
                            v28 := GenS2{F0: -1, F1: 0}
                            _ = v28
                        }
                        for fi29 := 0; fi29 < 8; fi29++ {
                            var v30 uint8 = 241
                            v3 = v7
                            if (-9223372036854775808 != -9223372036854775808) {
                                for fi31 := 0; fi31 < 19; fi31++ {
                                    v22 = 1.5
                                    fmt.Println(v1, v7, v7)
                                    v30 = (v30 + (v30 << ((348) & 63)))
                                    v32 := GenS5{F0: v1, F1: 202, F2: 0}
                                    _ = fi31
                                    _ = v32
                                }
                            } else {
                                for fi33 := 0; fi33 < 4; fi33++ {
                                    v30 = v30
                                    _ = fi33
                                }
                                if false {
                                    for fi34 := 0; fi34 < 5; fi34++ {
                                        v35 := GenS5{F0: "foo bar", F1: 168, F2: -9223372036854775808}
                                        var v36 int16 = int16(v30)
                                        _ = fi34
                                        _ = v35
                                        _ = v36
                                    }
                                }
                            }
                            _ = fi29
                            _ = v30
                        }
                        switch 2147483647 {
                        case -9223372036854775808:
                            v37 := GenS6{F0: 42, F1: 1}
                            _ = v37
                        case 2147483647:
                            fmt.Println(v22)
                        case 128:
                            v22 += v22
                        }
                    }
                }
            }
        } else {
            if (!false) {
                v3 = ((v7 + ("hello" + ((v7 + (v1 + v7)) + (v7 + v1)))) + (v7 + v1))
                if len(v2) > 0 {
                    v2[0] = (v1 + v3)
                }
            }
        }
    }
    var v38 uint = 807
    switch v1 {
    case "\n\t":
        var v39 int8 = 1
        _ = v39
    case "a":
        if (42 <= -663) {
            for fi40 := 0; fi40 < 19; fi40++ {
                v41 := GenS7{F0: -868, F1: 211, F2: 0}
                for fi42 := 0; fi42 < 12; fi42++ {
                    v38 += 1
                    v38 += v38
                    _ = fi42
                }
                if false {
                    _ = len(v2)
                    v38 = v38
                    v43 := GenS4{F0: -1, F1: 0}
                    _ = v43
                }
                _ = fi40
                _ = v41
            }
            var v44 uint16 = 2
            var v45 int8 = -109
            _ = v44
            _ = v45
        }
    case "":
        v38 += 896
    }
    for fi46 := 0; fi46 < 19; fi46++ {
        v2 = append(v2, v7)
        var v47 float64 = v22
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v38
    _ = v4
    _ = v7
}
