package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 string; F2 uint64 }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 uint; F1 string; F2 int32 }
type GenS9 struct { F0 int16; F1 uint16 }
type GenS10 struct { F0 int32; F1 int8; F2 uint16 }
type GenS11 struct { F0 string; F1 uint32 }
func genFn0() uint8 {
    if true {
        var v7 float64 = (0.1 + -2.25)
        _ = v7
    } else {
        v8 := []int{4294967295, 399, 0}
        _ = v8
    }
    switch 42 {
    case 42:
        switch 32768 {
        case 0:
            var v9 uint32 = 642
            _ = v9
        case 128:
            v10 := GenS1{F0: "", F1: 9223372036854775807, F2: 32767}
            _ = v10
        default:
            var v11 float64 = -2.25
            _ = v11
        }
    case 1:
        var v12 uint32 = 2147483647
        _ = v12
    }
    var v13 uint8 = 42
    var v14 uint32 = 0
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := []int{}
        fmt.Println(fi15)
        if true {
            v17 := GenS2{F0: (v13 - (^(v13)))}
            for fi18 := 0; fi18 < 11; fi18++ {
                v19 := GenS3{F0: 42}
                v20 := GenS3{F0: 18446744073709551615}
                var v21 int16 = 0
                var v22 float32 = 3.14159
                v23 := GenS4{F0: 1}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
                _ = v22
                _ = v23
            }
            fmt.Println(v14, v13, v14)
            _ = v17
        } else {
            switch "foo bar" {
            case "hello":
                v24 := GenS5{F0: 0, F1: (("a" + "foo bar") + "hello"), F2: 988}
                _ = v24
            case "a":
                var v25 int32 = -1
                _ = v25
            case "\n\t":
                for fi26 := 0; fi26 < 7; fi26++ {
                    if true {
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        if (42 >= 42) {
                            v14 += v14
                            v14 -= (v14 >> ((42) & 63))
                            v14 = 67
                        } else {
                            v13 += (^(v13))
                            if false {
                                var v27 int32 = -1
                                var v28 uint64 = 0
                                _ = v27
                                _ = v28
                            }
                        }
                        v13 -= 0
                    }
                    var v29 uint16 = 682
                    _ = fi26
                    _ = v29
                }
            default:
                switch (fi15 >> ((65535) & 63)) {
                case -1:
                    if len(v16) > 0 {
                        v16[0] = -492
                    }
                default:
                    switch (fi15 - 32767) {
                    case 0:
                        v14 = (v14 - (^(v14)))
                    case 2147483647:
                        v30 := GenS6{F0: 1, F1: "\n\t"}
                        _ = v30
                    default:
                        v14 -= v14
                    }
                }
            }
        }
        _ = fi15
        _ = v16
    }
    _ = v13
    _ = v14
    return v13
}

func main() {
    var v1 float64 = -2.25
    var v2 uint32 = 2147483647
    var v3 float64 = (0.1 * (v1 - -2.25))
    v4 := []int{1}
    v5 := []string{}
    v6 := []string{("foo bar" + "hello")}
    v2 = (^(v2))
    if ((float32(v1) != (((1.5 * 0.0) / 3.14159) / float32(v1))) || (861 > -711)) {
        var v31 uint32 = (^(v2))
        var v32 int8 = 0
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        v34 := GenS7{F0: -785, F1: ("hello" + "")}
        switch fi33 {
        case 96:
            v35 := GenS8{F0: 255, F1: ((("\n\t" + "\n\t") + "a") + "\n\t"), F2: 42}
            _ = v35
        case 940:
            fmt.Println(fi33, v3, v1)
        default:
            var v36 float64 = float64(0.0)
            _ = v36
        }
        v5 = append(v5, "\n\t")
        _ = fi33
        _ = v34
    }
    var v37 int64 = 1
    if (v37 <= (^(v37))) {
        if len(v5) > 0 {
            v5[0] = "a"
        }
        switch -9223372036854775808 {
        case -801:
            var v38 uint64 = 0
            _ = v38
        case 903:
            if true {
                v39 := GenS9{F0: -1, F1: 0}
                v40 := GenS6{F0: 4294967295, F1: "hello"}
                _ = v39
                _ = v40
            }
        case 32767:
            fmt.Println(v3)
        default:
            var v41 uint32 = 507
            _ = v41
        }
    }
    v37 = v37
    if true {
        fmt.Println(v37, v37, v37)
        v37 -= v37
        switch "foo bar" {
        case "":
            v3 = ((v1 / v3) / ((v1 * ((v3 / v1) * v3)) - (float64(float32((float64(-2.25) * (v1 * float64((((3.14159 + ((0.0 - ((0.0 / 0.1) * 0.0)) * 0.1)) * float32(v1)) * 1e10)))))) * 1e10)))
        case "hello":
            v42 := GenS9{F0: -32768, F1: 255}
            _ = v42
        default:
            fmt.Println(v3)
        }
    } else {
        v1 = 1.5
    }
    fmt.Println(v37, v37, v2)
    v43 := GenS10{F0: int32(v37), F1: 70, F2: 1}
    fmt.Println(v1, v2)
    v44 := GenS11{F0: "foo bar", F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
