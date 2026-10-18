package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int64; F1 uint; F2 int }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 100
    var v3 uint64 = 304
    v4 := make(map[uint32]int)
    var v5 uint16 = 2
    v2 -= (^(v2))
    v6 := []int{}
    switch -41 {
    case -9223372036854775808:
        v6 = append(v6, 32768)
    case -810:
        for fi7 := 0; fi7 < 8; fi7++ {
            v2 -= v2
            _ = fi7
        }
    }
    v8 := GenS1{F0: v3, F1: "a"}
    switch ("a" + "") {
    case "a":
        if false {
            switch 4294967295 {
            case 2147483647:
                var v9 float32 = 1.5
                _ = v9
            case -26:
                v10 := GenS2{F0: (v5 != v5), F1: (9223372036854775807 > 251)}
                _ = v10
            default:
                switch 42 {
                case -385:
                    for fi11 := 0; fi11 < 9; fi11++ {
                        var v12 int16 = 1
                        for fi13 := 0; fi13 < 20; fi13++ {
                            var v14 string = "\n\t"
                            delete(v4, 0)
                            fmt.Println(v2, v3)
                            switch (fi11 & 99) {
                            case 127:
                                fmt.Println(fi13, fi13, v12)
                            case 456:
                                switch (("a" + v14) + (v14 + ("" + ((("" + "hello") + "foo bar") + ("" + v14))))) {
                                case "a":
                                    var v15 float64 = float64(float32(float64(1e10)))
                                    _ = v15
                                default:
                                    v12 -= (v12 ^ 1)
                                }
                            }
                            fmt.Println(v2, v2, v14)
                            _ = fi13
                            _ = v14
                        }
                        _ = fi11
                        _ = v12
                    }
                case 42:
                    v16 := GenS3{F0: 1, F1: 1, F2: 128}
                    _ = v16
                case -262:
                    v17 := GenS3{F0: 0, F1: 18446744073709551615, F2: -328}
                    _ = v17
                }
            }
            var v18 int32 = 0
            _ = v18
        }
    default:
        var v19 string = ("\n\t" + "foo bar")
        _ = v19
    }
    var v20 uint16 = (v5 ^ v5)
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 uint32 = 747
        _ = fi21
        _ = v22
    }
    if (42 == 987) {
        v5 += v20
        if false {
            v1[4294967295] = ("hello" + "foo bar")
            if true {
                var v23 int16 = 32767
                v24 := GenS3{F0: 1, F1: 2, F2: -467}
                _ = v23
                _ = v24
            } else {
                for fi25 := 0; fi25 < 14; fi25++ {
                    v26 := GenS3{F0: -166, F1: 2, F2: (fi25 + 981)}
                    if true {
                        var v27 string = "a"
                        v6 = append(v6, -137)
                        _ = v27
                    }
                    _ = fi25
                    _ = v26
                }
            }
            var v28 float32 = (3.14159 + (0.0 + 3.14159))
            _ = v28
        } else {
            for fi29 := 0; fi29 < 19; fi29++ {
                for fi30 := 0; fi30 < 6; fi30++ {
                    var v31 string = "foo bar"
                    v2 = (v2 + v2)
                    var v32 int8 = (v2 * 1)
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                for fi33 := 0; fi33 < 20; fi33++ {
                    var v34 int32 = -2147483648
                    if len(v6) > 0 {
                        v6[0] = fi33
                    }
                    _ = fi33
                    _ = v34
                }
                var v35 bool = true
                _ = fi29
                _ = v35
            }
            for fi36 := 0; fi36 < 17; fi36++ {
                for fi37 := 0; fi37 < 15; fi37++ {
                    switch fi37 {
                    case -2147483648:
                        var v38 float32 = -2.25
                        _ = v38
                    case -522:
                        switch (fi37 << ((65535) & 63)) {
                        case 32768:
                            v20 += (v5 >> ((42) & 63))
                        case 0:
                            var v39 bool = (812 != 4294967295)
                            _ = v39
                        case -2147483648:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    }
                    if false {
                        _ = len(v1)
                        v40 := GenS4{F0: 215, F1: v5, F2: 255}
                        _ = v40
                    } else {
                        v3 = v3
                        switch 619 {
                        case 32768:
                            v2 = (v2 & -109)
                        case 484:
                            fmt.Println(v5)
                        case -543:
                            var v41 uint16 = v20
                            _ = v41
                        default:
                            var v42 float64 = (float64(float32(-2.25)) * 0.0)
                            _ = v42
                        }
                        var v43 int32 = 42
                        _ = v43
                    }
                    var v44 float32 = 1.5
                    _ = fi37
                    _ = v44
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi36
            }
        }
    }
    v45 := GenS4{F0: 0, F1: 926, F2: 2}
    fmt.Println(v20, v3, v20)
    switch 32768 {
    case 127:
        fmt.Println(v3, v3, v3)
    case 42:
        v3 = 0
    case 0:
        for fi46 := 0; fi46 < 2; fi46++ {
            if len(v6) > 0 {
                v6[0] = (fi46 | 0)
            }
            _ = fi46
        }
    default:
        var v47 uint32 = 323
        _ = v47
    }
    v5 = v5
    for fi48 := 0; fi48 < 19; fi48++ {
        var v49 int = fi48
        _ = fi48
        _ = v49
    }
    fmt.Println(v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v8
}
