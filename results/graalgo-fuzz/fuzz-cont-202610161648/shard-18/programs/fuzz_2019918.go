package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 uint32; F1 int16; F2 int8 }
type GenS6 struct { F0 uint }

func main() {
    var v1 int = 0
    var v2 uint16 = 65535
    var v3 float32 = ((-2.25 / 1.0) - 3.14159)
    var v4 uint64 = 18446744073709551615
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 float64 = (1e10 / float64((0.1 - (v3 * (v3 + v3)))))
        fmt.Println(v6)
        if true {
            if ((float64(v3) >= float64((1e10 - 0.0))) || true) {
                var v7 int16 = 1000
                _ = v7
            }
        } else {
            switch ((("hello" + "") + "hello") + ((("hello" + ("" + ("hello" + "\n\t"))) + "foo bar") + "foo bar")) {
            case "\n\t":
                if ((808 <= 42) || (!(-1 < 42))) {
                    v6 = float64(1e10)
                    v2 = (v2 / ((65535) | 1))
                }
            case "hello":
                switch (v1 / ((fi5) | 1)) {
                case 32767:
                    v1 = 875
                default:
                    v8 := []int{v1, fi5, v1, fi5}
                    _ = v8
                }
            }
            for fi9 := 0; fi9 < 17; fi9++ {
                if (1 != 2) {
                    var v10 int8 = 0
                    _ = v10
                } else {
                    v11 := []int{2147483647, (fi9 % ((2147483647) | 1))}
                    _ = v11
                }
                v12 := []int{(^(fi5)), (-(fi5))}
                _ = fi9
                _ = v12
            }
        }
        _ = fi5
        _ = v6
    }
    v13 := GenS1{F0: 1}
    v1 += v1
    v14 := GenS2{F0: v4}
    fmt.Println(v1)
    if false {
        v15 := []int{(v1 - v1), (v1 ^ (-(v1)))}
        _ = v15
    } else {
        fmt.Println(v2)
        v16 := []int{(-(v1)), 0, (v1 | (v1 >> ((0) & 63)))}
        v1 -= (^(v1))
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 uint64 = 939
        for fi19 := 0; fi19 < 18; fi19++ {
            if (65535 >= 2147483648) {
                var v20 int32 = 42
                switch "foo bar" {
                case "a":
                    v21 := GenS1{F0: v20}
                    _ = v21
                case "":
                    var v22 int32 = (v20 & v20)
                    _ = v22
                }
                _ = v20
            } else {
                v23 := GenS2{F0: (v4 >> ((0) & 63))}
                _ = v23
            }
            v24 := []int{(^(fi19)), (fi17 * (fi17 ^ 4294967295))}
            _ = fi19
            _ = v24
        }
        _ = fi17
        _ = v18
    }
    var v25 int = v1
    for fi26 := 0; fi26 < 17; fi26++ {
        for fi27 := 0; fi27 < 2; fi27++ {
            var v28 float32 = v3
            if false {
                v29 := GenS3{F0: 104, F1: 42}
                v30 := []int{128, fi26}
                _ = v29
                _ = v30
            } else {
                v28 = ((((1.5 * (v3 - -2.25)) + ((1e10 + v3) * 1.5)) * v28) / v28)
            }
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    v31 := []int{917, 32768}
    var v32 uint64 = (^(v4))
    var v33 int16 = -1
    switch "\n\t" {
    case "\n\t":
        switch "hello" {
        case "\n\t":
            v25 -= v1
        case "":
            var v34 uint32 = 2147483648
            _ = v34
        default:
            fmt.Println(v25)
        }
    case "":
        var v35 int64 = -58
        _ = v35
    default:
        if true {
            var v36 uint32 = 2147483648
            _ = v36
        } else {
            if false {
                fmt.Println(v2, v1)
                switch "foo bar" {
                case "foo bar":
                    switch ("foo bar" + "foo bar") {
                    case "":
                        v3 += (v3 / -2.25)
                    }
                case "a":
                    for fi37 := 0; fi37 < 15; fi37++ {
                        v38 := GenS4{F0: 255, F1: 128}
                        switch ("\n\t" + "") {
                        case "hello":
                            for fi39 := 0; fi39 < 4; fi39++ {
                                v31 = append(v31, (^(fi39)))
                                fmt.Println(fi37)
                                var v40 uint64 = (^(v4))
                                _ = fi39
                                _ = v40
                            }
                        case "":
                            if true {
                                if false {
                                    var v41 uint = 42
                                    _ = v41
                                }
                                for fi42 := 0; fi42 < 16; fi42++ {
                                    if len(v31) > 0 {
                                        v31[0] = fi37
                                    }
                                    _ = fi42
                                }
                            }
                        default:
                            for fi43 := 0; fi43 < 4; fi43++ {
                                var v44 float32 = 0.0
                                var v45 float64 = 1.5
                                fmt.Println(fi43)
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                        }
                        switch (fi37 % ((fi37) | 1)) {
                        case 0:
                            var v46 int64 = 0
                            _ = v46
                        case -525:
                            var v47 int = fi37
                            _ = v47
                        default:
                            v48 := GenS5{F0: 449, F1: v33, F2: int8(v1)}
                            _ = v48
                        }
                        _ = fi37
                        _ = v38
                    }
                default:
                    if (!(v32 <= 2)) {
                        v4 = v4
                    }
                }
            } else {
                if false {
                    v33 = (v33 * v33)
                } else {
                    fmt.Println(v32)
                }
                switch (("a" + "\n\t") + ("hello" + "hello")) {
                case "\n\t":
                    v31 = append(v31, (v1 & (^(v25))))
                case "a":
                    var v49 bool = (uint16(v32) <= 1)
                    _ = v49
                case "hello":
                    v50 := GenS2{F0: 680}
                    _ = v50
                }
                if (!true) {
                    v51 := GenS4{F0: 836, F1: 42}
                    v52 := GenS6{F0: 18446744073709551615}
                    for fi53 := 0; fi53 < 12; fi53++ {
                        v32 = (v4 >> ((0) & 63))
                        var v54 int8 = -20
                        _ = fi53
                        _ = v54
                    }
                    _ = v51
                    _ = v52
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
