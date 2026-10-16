package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 bool; F2 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 int8; F2 uint64 }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 int32 }

func main() {
    var v1 string = ("foo bar" + "hello")
    var v2 float32 = 0.1
    v3 := []string{"a", (v1 + "a")}
    var v4 float64 = 1e10
    if false {
        v2 -= float32(float64((v2 + -2.25)))
        v4 = -2.25
    } else {
        v1 = "foo bar"
        for fi5 := 0; fi5 < 3; fi5++ {
            v4 = v4
            if (!(false && false)) {
                v2 -= -2.25
            }
            for fi6 := 0; fi6 < 17; fi6++ {
                var v7 uint = 255
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
        switch 42 {
        case 128:
            if false {
                var v8 float32 = v2
                v1 = (("hello" + "\n\t") + ("foo bar" + v1))
                _ = v8
            } else {
                var v9 int = 2147483647
                v10 := GenS1{F0: 42, F1: -1}
                v9 = v9
                _ = v10
                _ = v9
            }
        case -1:
            for fi11 := 0; fi11 < 14; fi11++ {
                fmt.Println(v4, fi11, v1)
                _ = fi11
            }
        default:
            var v12 float32 = 1.5
            _ = v12
        }
    }
    var v13 int = 9223372036854775807
    fmt.Println(v2)
    for fi14 := 0; fi14 < 6; fi14++ {
        v1 = "foo bar"
        _ = fi14
    }
    var v15 string = v1
    for fi16 := 0; fi16 < 3; fi16++ {
        _ = len(v3)
        _ = fi16
    }
    v17 := GenS1{F0: -2147483648, F1: -384}
    if true {
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(v4, v2, v15)
            v19 := GenS2{F0: 1}
            _ = fi18
            _ = v19
        }
        v4 = v4
    } else {
        switch v13 {
        case -735:
            if false {
                v20 := GenS3{F0: 42}
                var v21 int64 = 9223372036854775807
                _ = v20
                _ = v21
            } else {
                if (545 > 42) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    v22 := GenS2{F0: 42}
                    if (!true) {
                        v23 := GenS3{F0: 0}
                        v2 -= v2
                        _ = v23
                    } else {
                        switch "" {
                        case "a":
                            v24 := GenS4{F0: 0, F1: ((207 != 255) && (true || true)), F2: 42}
                            _ = v24
                        default:
                            fmt.Println(v2, v4)
                        }
                        fmt.Println(v4, v15, v1)
                    }
                    _ = v22
                }
                for fi25 := 0; fi25 < 16; fi25++ {
                    if false {
                        var v26 bool = (false && (!true))
                        var v27 uint16 = 1
                        var v28 float32 = float32(v4)
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    switch (-(fi25)) {
                    case -9223372036854775808:
                        fmt.Println(v2, v1, v4)
                    case 4294967295:
                        fmt.Println(v1)
                    case 32768:
                        fmt.Println(v2)
                    }
                    for fi29 := 0; fi29 < 19; fi29++ {
                        v30 := GenS5{F0: 47}
                        var v31 int64 = -1
                        var v32 int64 = (^(v31))
                        _ = fi29
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    _ = fi25
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        case -9223372036854775808:
            for fi33 := 0; fi33 < 19; fi33++ {
                v4 = v4
                _ = fi33
            }
        }
        fmt.Println(v15)
    }
    v34 := GenS6{F0: ("" + "hello"), F1: 0, F2: 42}
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 int = (fi35 ^ (fi35 * fi35))
        fmt.Println(v1, v13, v4)
        var v37 int16 = 781
        _ = fi35
        _ = v36
        _ = v37
    }
    v13 -= 0
    if ((42 != 42) || true) {
        for fi38 := 0; fi38 < 15; fi38++ {
            var v39 uint = 443
            if (float64(3.14159) == v4) {
                var v40 uint8 = 255
                if true {
                    var v41 uint32 = 0
                    _ = v41
                }
                v1 = (v15 + v1)
                _ = v40
            } else {
                v4 = v4
                switch 0 {
                case -2147483648:
                    var v42 int64 = 1
                    _ = v42
                default:
                    if true {
                        fmt.Println(v13)
                        fmt.Println(fi38, v1)
                    } else {
                        v2 += 0.1
                        v15 = v15
                        for fi43 := 0; fi43 < 14; fi43++ {
                            v44 := GenS7{F0: 2147483647, F1: 207}
                            var v45 float64 = 3.14159
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                    }
                }
            }
            _ = len(v3)
            _ = fi38
            _ = v39
        }
    }
    for fi46 := 0; fi46 < 7; fi46++ {
        if false {
            if true {
                var v47 int = (v13 - 2147483647)
                var v48 string = (("\n\t" + v1) + "a")
                _ = v47
                _ = v48
            }
            fmt.Println(v4, v15, v15)
        }
        v49 := GenS8{F0: 42}
        _ = fi46
        _ = v49
    }
    var v50 int = (-(v13))
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v13
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v50
}
