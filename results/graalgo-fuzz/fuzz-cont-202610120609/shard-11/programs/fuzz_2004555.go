package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int64; F1 int16 }
func genFn0(p0 int16, p1 uint32, p2 int) string {
    var v5 int8 = -51
    v6 := []int{127, p2, 9223372036854775807, p2}
    for fi7 := 0; fi7 < 20; fi7++ {
        p2 -= -899
        _ = len(v6)
        _ = fi7
    }
    v8 := GenS1{F0: (-(v5)), F1: 175}
    v9 := GenS1{F0: v5, F1: 192}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v8
    _ = v9
    return (("hello" + "foo bar") + ("a" + "a"))
}

func main() {
    var v1 int32 = 894
    var v2 uint = 2
    var v3 float64 = (1e10 * (-2.25 / 1e10))
    v4 := make(map[uint32]string)
    if (uint16(v2) >= 2) {
        switch 128 {
        case 127:
            fmt.Println(v2, v1)
        case 0:
            fmt.Println(v3, v2, v3)
        }
        var v10 bool = false
        if v10 {
            v11 := GenS2{F0: -1, F1: -1}
            _ = v11
        }
        _ = v10
    }
    var v12 uint8 = 255
    var v13 uint16 = 1
    v13 -= v13
    v14 := GenS2{F0: 9223372036854775807, F1: -24}
    var v15 float64 = (3.14159 - ((float64(0.0) / (float64(0.1) + v3)) - v3))
    v1 -= -497
    if (-128 > 42) {
        for fi16 := 0; fi16 < 4; fi16++ {
            v12 = v12
            if (475 != v13) {
                var v17 bool = (42 <= 1)
                for fi18 := 0; fi18 < 3; fi18++ {
                    if (v17 && ((v1 & v1) == v1)) {
                        if v17 {
                            var v19 uint = v2
                            _ = v19
                        }
                        switch ("\n\t" + "foo bar") {
                        case "hello":
                            v20 := []int{(fi18 & fi18), int(v12), 32767}
                            _ = v20
                        case "foo bar":
                            v21 := GenS1{F0: 127, F1: v12}
                            _ = v21
                        default:
                            if v17 {
                                var v22 int8 = -1
                                if (v17 && (fi16 <= (fi16 | (fi18 ^ fi18)))) {
                                    v12 = 77
                                    var v23 int32 = v1
                                    _ = v23
                                }
                                _ = v22
                            } else {
                                if (v13 == v13) {
                                    var v24 bool = (42 > 42)
                                    _ = v24
                                } else {
                                    v1 += 42
                                    v25 := GenS3{F0: "foo bar", F1: -578}
                                    var v26 int16 = 32767
                                    v12 -= v12
                                    _ = v25
                                    _ = v26
                                }
                                switch "" {
                                case "foo bar":
                                    fmt.Println(v3, v17)
                                }
                                var v27 float64 = v3
                                v12 = v12
                                if (v17 || false) {
                                    switch (fi18 << ((42) & 63)) {
                                    case 281:
                                        switch 32768 {
                                        case 465:
                                            switch 0 {
                                            case 483:
                                                if true {
                                                    fmt.Println(v13, v12)
                                                    v2 -= v2
                                                } else {
                                                    v27 += v15
                                                    v1 -= (v1 >> ((v2) & 63))
                                                    if ((^(v12)) != v12) {
                                                        var v28 int64 = 1
                                                        _ = v28
                                                    } else {
                                                        fmt.Println(v1, fi18)
                                                        v1 = v1
                                                        var v29 uint64 = 658
                                                        fmt.Println(v3, v27)
                                                        _ = v29
                                                    }
                                                }
                                            case 127:
                                                v12 -= 42
                                            case -633:
                                                v2 -= 227
                                            }
                                        }
                                    }
                                    v30 := GenS4{F0: -9223372036854775808, F1: -765}
                                    v17 = (-128 > 127)
                                    var v31 int16 = -32768
                                    v13 = 32767
                                    var v32 int = fi16
                                    v33 := []int{9223372036854775807}
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                                _ = v27
                            }
                        }
                        if ((9223372036854775807 > -104) && true) {
                            v34 := v4[2147483648]
                            _ = v34
                        }
                        var v35 string = ""
                        var v36 uint64 = 1
                        v35 = ""
                        var v37 uint64 = v36
                        var v38 int32 = -1
                        _ = v35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    var v39 uint8 = 2
                    var v40 int = (fi18 & (^(fi18)))
                    _ = fi18
                    _ = v39
                    _ = v40
                }
                _ = v17
            }
            var v41 float64 = (v3 / 1.0)
            _ = fi16
            _ = v41
        }
        for fi42 := 0; fi42 < 12; fi42++ {
            var v43 uint16 = 322
            _ = fi42
            _ = v43
        }
    }
    var v44 int64 = 0
    v45 := GenS3{F0: "hello", F1: 772}
    var v46 string = "\n\t"
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
}
