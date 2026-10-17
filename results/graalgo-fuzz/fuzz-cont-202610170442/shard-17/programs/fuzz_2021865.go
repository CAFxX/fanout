package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint; F1 uint16; F2 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 int64; F2 bool }

func main() {
    v1 := []string{"", ("foo bar" + "foo bar"), "hello"}
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    var v4 int32 = 738
    if true {
        v4 += v4
        v4 += 0
    }
    switch (("hello" + ("a" + "\n\t")) + "a") {
    case "\n\t":
        v5 := GenS1{F0: (v4 % ((v4) | 1))}
        _ = v5
    case "a":
        switch 2147483647 {
        case -1:
            switch ("a" + "foo bar") {
            case "foo bar":
                v4 = v4
            case "a":
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
            default:
                fmt.Println(v4)
            }
        case 2147483647:
            for fi6 := 0; fi6 < 11; fi6++ {
                v7 := GenS1{F0: 655}
                for fi8 := 0; fi8 < 7; fi8++ {
                    v9 := GenS2{F0: 92, F1: 42}
                    _ = fi8
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
        default:
            var v10 bool = (255 != 1)
            _ = v10
        }
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 uint16 = 255
        v4 = (v4 * (v4 | (v4 | (-(v4)))))
        _ = fi11
        _ = v12
    }
    v4 -= v4
    v4 = (v4 & v4)
    var v13 float64 = 0.1
    switch 127 {
    case -813:
        var v14 uint = 40
        _ = v14
    case 127:
        _ = len(v3)
    case 32768:
        v15 := GenS1{F0: (v4 * -1)}
        _ = v15
    default:
        v16 := GenS1{F0: 42}
        _ = v16
    }
    v1 = append(v1, "a")
    if (false && (v13 != (v13 * ((v13 - v13) / v13)))) {
        delete(v2, 127)
        for fi17 := 0; fi17 < 2; fi17++ {
            if true {
                v18 := v2[(fi17 | (fi17 >> ((784) & 63)))]
                _ = v18
            } else {
                v3[fi17] = (473 != 0)
            }
            var v19 int = fi17
            _ = fi17
            _ = v19
        }
        v20 := GenS3{F0: ((((0.1 - 1e10) / 1.0) <= (float32(v13) * float32(0.0))) && false), F1: (v4 & (v4 % ((1) | 1)))}
        _ = v20
    }
    _ = len(v3)
    v21 := GenS4{F0: 0, F1: 309, F2: -741}
    v4 = v4
    switch 1 {
    case 127:
        v22 := GenS4{F0: 4294967295, F1: 32767, F2: 126}
        _ = v22
    default:
        switch "hello" {
        case "\n\t":
            if (!(!((!true) || (!(!(("a" + "\n\t") >= "")))))) {
                for fi23 := 0; fi23 < 8; fi23++ {
                    for fi24 := 0; fi24 < 12; fi24++ {
                        v25 := GenS3{F0: (42 > 1), F1: v4}
                        switch fi24 {
                        case 4294967295:
                            var v26 uint32 = 0
                            _ = v26
                        default:
                            if (2147483647 <= 1) {
                                if (false || true) {
                                    if (!(!(669 == 684))) {
                                        _ = len(v2)
                                        v4 -= (^(v4))
                                    }
                                    v27 := GenS5{F0: v4}
                                    _ = v27
                                }
                                switch (-(fi23)) {
                                case 0:
                                    var v28 uint = 0
                                    _ = v28
                                }
                                fmt.Println(v4)
                            }
                        }
                        fmt.Println(fi23, fi23)
                        _ = fi24
                        _ = v25
                    }
                    v29 := GenS6{F0: -128, F1: -9223372036854775808, F2: (127 == 42)}
                    _ = fi23
                    _ = v29
                }
            } else {
                var v30 int = 127
                if false {
                    v4 = 97
                    for fi31 := 0; fi31 < 12; fi31++ {
                        v32 := GenS1{F0: (v4 % ((267) | 1))}
                        _ = fi31
                        _ = v32
                    }
                    var v33 uint32 = 2147483648
                    _ = v33
                } else {
                    v13 = (v13 + v13)
                    var v34 uint64 = 4294967295
                    _ = v34
                }
                _ = v30
            }
        case "foo bar":
            if true {
                var v35 int64 = 42
                _ = v35
            }
        case "a":
            var v36 uint8 = 127
            _ = v36
        default:
            if (97 != 255) {
                var v37 float32 = 1.5
                fmt.Println(v13, v37)
                v37 = v37
                _ = v37
            } else {
                v13 += float64((1e10 + (1e10 * (1e10 + float32(((v13 * (0.0 + v13)) - 1e10))))))
            }
        }
    }
    var v38 float32 = (float32(3.14159) - 1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v38
    _ = v4
}
