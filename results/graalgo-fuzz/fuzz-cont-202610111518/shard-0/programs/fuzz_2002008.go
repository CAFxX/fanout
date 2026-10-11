package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint; F1 int16; F2 int32 }
type GenS3 struct { F0 uint; F1 int; F2 int8 }

func main() {
    var v1 int64 = -753
    var v2 uint16 = 65535
    var v3 string = (((("hello" + ("a" + ("" + ("a" + "foo bar")))) + "hello") + ("" + "a")) + "a")
    var v4 string = v3
    v3 = v4
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int64 = v1
        if (!false) {
            v3 = "hello"
            switch fi5 {
            case 128:
                fmt.Println(v6, v3)
            case -9223372036854775808:
                fmt.Println(v1)
            default:
                v7 := GenS1{F0: -32768, F1: 2}
                _ = v7
            }
        } else {
            var v8 float32 = float32(0.1)
            var v9 int = fi5
            _ = v8
            _ = v9
        }
        var v10 uint8 = 101
        _ = fi5
        _ = v10
        _ = v6
    }
    v2 -= 723
    fmt.Println(v2)
    switch 41 {
    case -421:
        var v11 uint64 = 0
        _ = v11
    case -279:
        v12 := []int{128, 4294967295, 0, 454}
        _ = v12
    default:
        v13 := []int{42, 32768, 9223372036854775807, 2147483647}
        _ = v13
    }
    if (true || false) {
        var v14 int32 = 42
        v15 := GenS1{F0: int16(v2), F1: 395}
        _ = v14
        _ = v15
    } else {
        v16 := []int{32768, -9223372036854775808}
        switch v4 {
        case "\n\t":
            fmt.Println(v4)
        case "":
            v4 = ((v4 + ("hello" + v4)) + v3)
        case "foo bar":
            fmt.Println(v2, v3)
        default:
            if (2 <= 18446744073709551615) {
                for fi17 := 0; fi17 < 10; fi17++ {
                    fmt.Println(v1, v2, v1)
                    var v18 bool = (2147483647 <= 0)
                    _ = fi17
                    _ = v18
                }
            }
        }
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 float64 = (1.5 * (0.0 + -2.25))
            var v21 uint32 = 675
            v22 := GenS2{F0: 18446744073709551615, F1: 1000, F2: 42}
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = v16
    }
    var v23 float32 = float32(((float64(float32(1e10)) / (3.14159 + 0.1)) - (0.1 - 1e10)))
    for fi24 := 0; fi24 < 16; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := GenS3{F0: 255, F1: (fi24 * fi24), F2: 100}
            _ = fi25
            _ = v26
        }
        v4 = v3
        if (!true) {
            v27 := GenS2{F0: 0, F1: -80, F2: -857}
            for fi28 := 0; fi28 < 16; fi28++ {
                v29 := GenS1{F0: 42, F1: 4294967295}
                var v30 int = fi28
                _ = fi28
                _ = v29
                _ = v30
            }
            switch "foo bar" {
            case "":
                fmt.Println(v3)
            case "a":
                switch fi24 {
                case 0:
                    var v31 bool = (1 >= 42)
                    _ = v31
                case 32768:
                    if (127 == 127) {
                        if false {
                            switch fi24 {
                            case 587:
                                v23 = ((float32(0.0) + (((v23 / (v23 + v23)) - 1.5) + v23)) * float32(0.0))
                            case 128:
                                var v32 uint8 = 65
                                _ = v32
                            case 1:
                                if ((v4 + (("a" + ("hello" + v3)) + v3)) != "\n\t") {
                                    fmt.Println(v23, v23)
                                    switch 0 {
                                    case 42:
                                        fmt.Println(v2, v4, v1)
                                    case 150:
                                        for fi33 := 0; fi33 < 19; fi33++ {
                                            var v34 int8 = 100
                                            v1 -= (^(v1))
                                            _ = fi33
                                            _ = v34
                                        }
                                    }
                                }
                            default:
                                fmt.Println(v1, fi24, v3)
                            }
                        } else {
                            for fi35 := 0; fi35 < 20; fi35++ {
                                v23 = (1.5 * v23)
                                var v36 float64 = -2.25
                                _ = fi35
                                _ = v36
                            }
                        }
                        for fi37 := 0; fi37 < 1; fi37++ {
                            v23 += float32(3.14159)
                            _ = fi37
                        }
                    }
                }
            default:
                v1 = 470
            }
            _ = v27
        } else {
            fmt.Println(v2, v2)
            var v38 int32 = 1
            _ = v38
        }
        _ = fi24
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        v40 := []int{fi39}
        var v41 float32 = v23
        _ = fi39
        _ = v40
        _ = v41
    }
    if (!(2147483647 == 1)) {
        var v42 string = v4
        _ = v42
    }
    v2 = v2
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
