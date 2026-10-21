package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint32 }

func main() {
    v1 := []uint32{65535, 65535, 1}
    var v2 int32 = 2147483647
    var v3 int8 = -1
    v4 := make(map[string]int)
    v5 := make(map[uint32]string)
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 255, F1: "a", F2: 111}
        for fi9 := 0; fi9 < 5; fi9++ {
            delete(v4, "a")
            v3 = v3
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    var v10 int64 = 42
    _ = len(v4)
    v2 = v2
    var v11 int16 = 118
    v12 := GenS1{F0: 127, F1: "\n\t", F2: 16}
    if (true && (!(false && (((false || (true && true)) && true) || (156 < 2147483648))))) {
        if false {
            if false {
                v3 = v3
                var v13 float64 = (3.14159 * ((0.0 * 0.0) + ((1.5 * 1.5) + 0.0)))
                var v14 uint32 = 1
                _ = v13
                _ = v14
            } else {
                _ = len(v4)
            }
            if false {
                v11 += (v11 >> ((2) & 63))
                v15 := v4["hello"]
                switch "a" {
                case "a":
                    v5[65535] = ""
                }
                _ = v15
            }
            v10 += -611
        } else {
            if false {
                var v16 float64 = (-2.25 * 0.0)
                _ = v16
            }
            v1 = append(v1, 0)
            v17 := GenS2{F0: 42}
            _ = v17
        }
        var v18 string = ""
        switch -614 {
        case -9223372036854775808:
            for fi19 := 0; fi19 < 15; fi19++ {
                if ((^(v3)) > 127) {
                    delete(v4, v18)
                }
                var v20 int = (fi19 - fi19)
                fmt.Println(v10, v11)
                _ = fi19
                _ = v20
            }
        case 9223372036854775807:
            switch 237 {
            case -752:
                fmt.Println(v10, v2, v3)
            case -621:
                v21 := GenS1{F0: 127, F1: v18, F2: 110}
                _ = v21
            default:
                v22 := GenS2{F0: 18446744073709551615}
                _ = v22
            }
        case 1:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            var v23 uint = 2
            _ = v23
        }
        _ = v18
    }
    var v24 uint16 = 232
    for fi25 := 0; fi25 < 4; fi25++ {
        if (42 <= 0) {
            _ = len(v6)
            v3 += v3
        } else {
            fmt.Println(v24)
            v24 = (v24 << ((454) & 63))
            v5[4294967295] = ("foo bar" + "")
        }
        for fi26 := 0; fi26 < 10; fi26++ {
            if len(v1) > 0 {
                v1[0] = 42
            }
            if true {
                if true {
                    delete(v5, 2147483647)
                    v24 = (v24 + 0)
                    var v27 int64 = -195
                    _ = v27
                } else {
                    if true {
                        v1 = append(v1, 65535)
                        switch fi26 {
                        case -9223372036854775808:
                            v24 = v24
                        default:
                            if (((315 > 1) || true) || true) {
                                v10 -= (v10 + 42)
                                if (((v2 >> ((1) & 63)) != 42) || false) {
                                    fmt.Println(fi25, fi26, v11)
                                    var v28 uint8 = 42
                                    _ = v28
                                } else {
                                    v29 := v5[0]
                                    _ = v29
                                }
                                v30 := GenS2{F0: 4294967295}
                                _ = v30
                            }
                        }
                        var v31 bool = true
                        _ = v31
                    }
                    v24 -= v24
                    var v32 int = (fi25 << ((0) & 63))
                    _ = v32
                }
                v3 -= v3
            }
            var v33 uint64 = 2
            v34 := GenS3{F0: 42, F1: 675}
            var v35 float64 = 0.0
            v36 := GenS2{F0: 1}
            var v37 float32 = (1.5 / (3.14159 + 0.1))
            _ = fi26
            _ = v33
            _ = v34
            _ = v35
            _ = v36
            _ = v37
        }
        _ = fi25
    }
    v38 := GenS2{F0: 2}
    if ((v10 << ((65535) & 63)) > 9223372036854775807) {
        var v39 int = -9223372036854775808
        v2 = (v2 * (^(v2)))
        _ = v39
    } else {
        if len(v1) > 0 {
            v1[0] = 65535
        }
    }
    fmt.Println(v11, v24)
    v10 = 474
    v10 += v10
    if len(v1) > 0 {
        v1[0] = 1
    }
    if ((679 != 2147483648) || (!true)) {
        fmt.Println(v10)
        v24 = v24
    }
    switch -9 {
    case -9223372036854775808:
        var v40 string = ((("hello" + "hello") + "a") + "\n\t")
        _ = v40
    case 10:
        v3 = v3
    case 42:
        var v41 int32 = (^(v2))
        _ = v41
    default:
        var v42 float64 = -2.25
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
