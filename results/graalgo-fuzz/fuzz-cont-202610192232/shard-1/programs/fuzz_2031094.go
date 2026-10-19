package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 uint16; F2 string }
type GenS5 struct { F0 int8 }

func main() {
    v1 := []string{}
    var v2 float64 = 1e10
    var v3 int8 = 100
    var v4 uint32 = 65535
    v2 = v2
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int16 = 329
        if true {
            for fi7 := 0; fi7 < 6; fi7++ {
                if true {
                    switch (("foo bar" + ("" + ("" + "hello"))) + "hello") {
                    case "a":
                        if (127 != 2147483647) {
                            var v8 uint = 445
                            _ = v8
                        }
                    case "foo bar":
                        var v9 int = (fi5 | fi5)
                        _ = v9
                    case "\n\t":
                        var v10 uint16 = 1
                        _ = v10
                    }
                    _ = len(v1)
                    if false {
                        var v11 uint64 = 777
                        _ = v11
                    }
                } else {
                    _ = len(v1)
                    if false {
                        var v12 uint64 = 18446744073709551615
                        _ = v12
                    }
                    var v13 int32 = 2147483647
                    _ = v13
                }
                v14 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615}
                v4 += v4
                _ = fi7
                _ = v14
            }
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS1{F0: 1, F1: 1}
                _ = fi15
                _ = v16
            }
            switch fi5 {
            case 32767:
                var v17 bool = false
                _ = v17
            }
        }
        v4 += v4
        _ = fi5
        _ = v6
    }
    var v18 bool = (42 >= 42)
    for fi19 := 0; fi19 < 3; fi19++ {
        if (4294967295 <= (fi19 * (fi19 & (^(fi19))))) {
            for fi20 := 0; fi20 < 13; fi20++ {
                if ((v3 + (v3 % ((-1) | 1))) == v3) {
                    var v21 int8 = v3
                    v1 = append(v1, "foo bar")
                    var v22 int8 = (v3 | 0)
                    _ = v21
                    _ = v22
                }
                _ = fi20
            }
            v4 += (^(v4))
        }
        fmt.Println(v4)
        _ = fi19
    }
    if (!true) {
        if (-32768 == 0) {
            v3 = (^(v3))
            if (v18 && (1 <= 495)) {
                fmt.Println(v3)
            } else {
                var v23 uint64 = 18446744073709551615
                _ = v23
            }
            switch "" {
            case "a":
                fmt.Println(v4, v2)
            case "foo bar":
                v24 := GenS1{F0: 885, F1: 925}
                _ = v24
            case "\n\t":
                v25 := GenS2{F0: 32767}
                _ = v25
            }
        } else {
            switch "a" {
            case "":
                v2 -= (float64(-2.25) + float64(((0.1 - (1.5 * 1e10)) * (-2.25 - ((-2.25 + 3.14159) / (float32(-2.25) / -2.25))))))
            case "\n\t":
                v2 -= 0.0
            }
            for fi26 := 0; fi26 < 4; fi26++ {
                fmt.Println(v4, v4, v2)
                fmt.Println(v18, fi26, v3)
                _ = fi26
            }
            switch 685 {
            case 32768:
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 int8 = (^(v3))
                    for fi29 := 0; fi29 < 20; fi29++ {
                        v30 := GenS2{F0: 0}
                        v18 = v18
                        v2 = (float64(-2.25) - v2)
                        _ = fi29
                        _ = v30
                    }
                    var v31 int32 = 2147483647
                    _ = fi27
                    _ = v28
                    _ = v31
                }
            case 32767:
                if v18 {
                    for fi32 := 0; fi32 < 20; fi32++ {
                        v2 -= 0.0
                        for fi33 := 0; fi33 < 1; fi33++ {
                            switch -264 {
                            case -640:
                                var v34 int8 = (-(v3))
                                _ = v34
                            default:
                                v3 = v3
                            }
                            switch "a" {
                            case "foo bar":
                                v4 -= v4
                            case "":
                                var v35 float32 = 0.1
                                v35 = (float32(v2) + float32(v2))
                                _ = v35
                            default:
                                v18 = (v4 <= v4)
                            }
                            var v36 uint32 = (v4 & v4)
                            _ = fi33
                            _ = v36
                        }
                        var v37 int64 = 0
                        _ = fi32
                        _ = v37
                    }
                    if ((v3 & v3) > (v3 / (((-(v3))) | 1))) {
                        if v18 {
                            var v38 uint = uint(v4)
                            v2 = 0.1
                            if (v38 != 1) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v39 float32 = 3.14159
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                                _ = v39
                            }
                            _ = v38
                        } else {
                            fmt.Println(v4, v2)
                            v40 := GenS2{F0: 42}
                            _ = v40
                        }
                        v41 := GenS3{F0: 685}
                        _ = v41
                    }
                }
            case -1:
                v42 := GenS4{F0: v18, F1: 712, F2: "hello"}
                _ = v42
            }
        }
        fmt.Println(v4)
    } else {
        var v43 int = -9223372036854775808
        _ = v43
    }
    fmt.Println(v3, v18, v2)
    var v44 string = ""
    var v45 uint = 42
    if v18 {
        var v46 uint32 = v4
        if ((0 != -2147483648) || false) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v47 int64 = 1
        _ = v46
        _ = v47
    } else {
        var v48 int32 = 2147483647
        v4 += v4
        for fi49 := 0; fi49 < 11; fi49++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if true {
                switch int(v3) {
                case 1:
                    _ = len(v1)
                case 32768:
                    v3 -= v3
                case -966:
                    var v50 bool = false
                    _ = v50
                default:
                    v51 := GenS3{F0: 1}
                    _ = v51
                }
                v48 -= v48
            }
            switch fi49 {
            case 42:
                fmt.Println(v44)
            default:
                if ((0 == int16(v45)) || false) {
                    var v52 string = v44
                    v4 = (^(v4))
                    switch (fi49 ^ (-(fi49))) {
                    case -441:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 127:
                        var v53 bool = (!v18)
                        _ = v53
                    }
                    _ = v52
                }
            }
            _ = fi49
        }
        _ = v48
    }
    var v54 string = "foo bar"
    for fi55 := 0; fi55 < 9; fi55++ {
        for fi56 := 0; fi56 < 15; fi56++ {
            fmt.Println(v18, v2)
            _ = fi56
        }
        v45 = 65535
        var v57 float64 = v2
        _ = fi55
        _ = v57
    }
    var v58 int = -260
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v59 := GenS5{F0: (v3 - v3)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v54:", v54)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v54
    _ = v58
    _ = v59
}
