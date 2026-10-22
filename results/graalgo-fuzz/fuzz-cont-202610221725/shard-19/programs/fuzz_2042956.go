package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 string }
type GenS2 struct { F0 int32; F1 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 495
    v3 := []int{127, 2147483647, 128, 9223372036854775807}
    var v4 int16 = -830
    v5 := GenS1{F0: 0, F1: 884, F2: ("foo bar" + ("a" + (("\n\t" + "\n\t") + "")))}
    switch 4294967295 {
    case -9223372036854775808:
        var v6 uint64 = 2
        _ = v6
    default:
        v1[uint32(v2)] = 32767
    }
    v4 -= (v4 | v4)
    v4 = -823
    v4 = 0
    v4 += v4
    if true {
        v2 = v2
        fmt.Println(v4)
        v4 = v4
    } else {
        for fi7 := 0; fi7 < 1; fi7++ {
            if false {
                if (246 != 127) {
                    v4 += v4
                    fmt.Println(v4)
                    var v8 uint8 = 255
                    _ = v8
                }
            } else {
                if false {
                    if (!true) {
                        for fi9 := 0; fi9 < 1; fi9++ {
                            v10 := GenS1{F0: 864, F1: 4294967295, F2: "hello"}
                            _ = fi9
                            _ = v10
                        }
                        v11 := v1[0]
                        _ = v11
                    } else {
                        if (-1 < -2147483648) {
                            v3 = append(v3, fi7)
                        } else {
                            for fi12 := 0; fi12 < 4; fi12++ {
                                var v13 int = fi12
                                _ = fi12
                                _ = v13
                            }
                            var v14 bool = true
                            _ = v14
                        }
                        switch "a" {
                        case "foo bar":
                            if (v4 <= (-(v4))) {
                                switch (fi7 + fi7) {
                                case -690:
                                    if true {
                                        fmt.Println(fi7, v4, v2)
                                        v2 = v2
                                        for fi15 := 0; fi15 < 12; fi15++ {
                                            _ = len(v1)
                                            switch (fi15 << ((941) & 63)) {
                                            case 4294967295:
                                                var v16 uint8 = 1
                                                _ = v16
                                            default:
                                                var v17 float32 = 1e10
                                                _ = v17
                                            }
                                            v2 += 0
                                            _ = fi15
                                        }
                                    }
                                default:
                                    for fi18 := 0; fi18 < 11; fi18++ {
                                        switch "\n\t" {
                                        case "":
                                            if (!(!true)) {
                                                v4 -= v4
                                                var v19 float64 = (float64(0.1) - ((-2.25 - 3.14159) - 0.1))
                                                v3 = append(v3, (fi18 << ((65535) & 63)))
                                                _ = v19
                                            }
                                        }
                                        v1[4294967295] = fi18
                                        var v20 string = "hello"
                                        _ = fi18
                                        _ = v20
                                    }
                                }
                                var v21 float32 = (float32(3.14159) / 1.0)
                                for fi22 := 0; fi22 < 8; fi22++ {
                                    _ = len(v3)
                                    v2 -= 0
                                    _ = fi22
                                }
                                _ = v21
                            } else {
                                var v23 int8 = 100
                                v4 = (-(v4))
                                var v24 uint64 = 4294967295
                                _ = v23
                                _ = v24
                            }
                        case "\n\t":
                            v25 := GenS1{F0: 1, F1: 255, F2: "hello"}
                            _ = v25
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    v26 := GenS1{F0: -9223372036854775808, F1: 2, F2: ""}
                    for fi27 := 0; fi27 < 10; fi27++ {
                        v1[65535] = fi7
                        _ = fi27
                    }
                    _ = v26
                } else {
                    if false {
                        var v28 int8 = -1
                        _ = v28
                    } else {
                        fmt.Println(v4)
                    }
                    var v29 uint = 255
                    v30 := GenS1{F0: 0, F1: 42, F2: ""}
                    _ = v29
                    _ = v30
                }
            }
            v31 := GenS1{F0: int64(v4), F1: 4294967295, F2: "a"}
            _ = fi7
            _ = v31
        }
        var v32 int = 4294967295
        _ = v32
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        _ = len(v1)
        _ = fi33
    }
    v2 += (v2 % ((v2) | 1))
    v2 = v2
    for fi34 := 0; fi34 < 3; fi34++ {
        v2 = 42
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi34
    }
    v4 -= v4
    v2 = (v2 + v2)
    if true {
        v2 = v2
        v35 := GenS1{F0: 42, F1: uint(v2), F2: "\n\t"}
        _ = v35
    } else {
        for fi36 := 0; fi36 < 18; fi36++ {
            for fi37 := 0; fi37 < 6; fi37++ {
                fmt.Println(fi36, fi36, fi36)
                _ = fi37
            }
            var v38 float64 = 0.1
            var v39 bool = (1 <= 26)
            _ = fi36
            _ = v38
            _ = v39
        }
        v40 := GenS2{F0: 2147483647, F1: v4}
        _ = v40
    }
    var v41 uint = 500
    var v42 uint32 = 50
    v43 := GenS2{F0: 925, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
