package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int64 }

func main() {
    var v1 float64 = (1.5 * (float64(float32(3.14159)) - -2.25))
    var v2 float64 = (v1 + v1)
    var v3 uint8 = 0
    var v4 float64 = (v2 * v2)
    switch -880 {
    case 4294967295:
        var v5 uint64 = 0
        _ = v5
    case -285:
        v4 = 0.1
    case -2147483648:
        fmt.Println(v2)
    default:
        if false {
            if ((9223372036854775807 != -1) && ("foo bar" >= ("a" + ""))) {
                v1 = v1
            }
            var v6 uint16 = 0
            v3 = 255
            _ = v6
        } else {
            if false {
                v4 -= v2
            } else {
                if false {
                    if true {
                        v1 = 1.5
                    } else {
                        var v7 uint16 = 4
                        if (0 == 1) {
                            v8 := GenS1{F0: 18446744073709551615}
                            _ = v8
                        }
                        for fi9 := 0; fi9 < 6; fi9++ {
                            v4 -= v1
                            var v10 uint16 = v7
                            _ = fi9
                            _ = v10
                        }
                        _ = v7
                    }
                }
                v3 += v3
            }
            v4 = 0.1
            v3 += 249
        }
    }
    fmt.Println(v3, v1, v1)
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS2{F0: 42, F1: (("foo bar" + "hello") + "foo bar")}
        _ = fi11
        _ = v12
    }
    var v13 float64 = ((float64(((0.0 * float32((-2.25 + float64(float32(0.1))))) - 3.14159)) * -2.25) / (3.14159 - 1.5))
    for fi14 := 0; fi14 < 3; fi14++ {
        switch (fi14 >> ((255) & 63)) {
        case -330:
            v4 = (0.1 / 1.5)
        case 42:
            if false {
                var v15 int32 = -269
                v15 += v15
                _ = v15
            }
        case 32768:
            if ((100 > -82) && false) {
                for fi16 := 0; fi16 < 7; fi16++ {
                    v17 := []int{fi14, fi14, 32767}
                    _ = fi16
                    _ = v17
                }
            } else {
                for fi18 := 0; fi18 < 19; fi18++ {
                    var v19 bool = true
                    switch 2147483647 {
                    case -9223372036854775808:
                        if true {
                            var v20 int8 = -1
                            var v21 uint32 = 0
                            _ = v20
                            _ = v21
                        }
                    case 128:
                        v22 := GenS2{F0: -977, F1: "hello"}
                        _ = v22
                    case 9223372036854775807:
                        switch (fi14 ^ (fi14 & fi14)) {
                        case 2147483647:
                            v23 := GenS1{F0: 255}
                            _ = v23
                        case -391:
                            v24 := GenS3{F0: 42}
                            _ = v24
                        case -1:
                            fmt.Println(v3)
                        }
                    }
                    _ = fi18
                    _ = v19
                }
            }
        }
        _ = fi14
    }
    v25 := []int{948, 0, 42, 9223372036854775807}
    if (0.1 < 3.14159) {
        v2 = ((1e10 / v1) - v1)
        var v26 float64 = -2.25
        _ = v26
    } else {
        v13 = ((v2 - (-2.25 + 1.5)) - v1)
        _ = len(v25)
    }
    v3 = (^(v3))
    if (-1 != -1) {
        var v27 int8 = -115
        if false {
            if false {
                v3 = (v3 - (^(v3)))
                if false {
                    if true {
                        var v28 uint32 = 42
                        switch 0 {
                        case -647:
                            v2 = (v2 + ((v13 / v4) + 0.0))
                        default:
                            var v29 int32 = -1
                            _ = v29
                        }
                        switch 32768 {
                        case -9223372036854775808:
                            switch 498 {
                            case 0:
                                fmt.Println(v2)
                            case 4294967295:
                                v28 = (^(v28))
                            }
                        case 0:
                            fmt.Println(v2)
                        case 2147483647:
                            v30 := GenS2{F0: 42, F1: "a"}
                            _ = v30
                        }
                        _ = v28
                    }
                    fmt.Println(v27, v2)
                    var v31 uint16 = 255
                    _ = v31
                } else {
                    v13 += (0.1 / (3.14159 * v4))
                }
            } else {
                _ = len(v25)
            }
        }
        _ = v27
    } else {
        var v32 float32 = 0.0
        v33 := GenS2{F0: -1, F1: "\n\t"}
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        v13 -= v4
        if (false && false) {
            var v35 uint = 864
            switch "" {
            case "a":
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
            case "foo bar":
                if len(v25) > 0 {
                    v25[0] = (fi34 / (((-(fi34))) | 1))
                }
            case "hello":
                fmt.Println(v13, v2, v2)
            }
            _ = v35
        }
        _ = len(v25)
        _ = fi34
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
        _ = fi36
    }
    var v37 float32 = 0.1
    fmt.Println(v37)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v37
    _ = v4
}
