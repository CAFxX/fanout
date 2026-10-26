package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 string; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint64 }

func main() {
    var v1 float32 = (float32((0.1 / float64(1e10))) + (float32((0.0 + 0.1)) * 0.1))
    var v2 uint8 = 0
    v3 := []uint32{65535}
    var v4 uint = 42
    var v5 bool = (!true)
    var v6 uint8 = v2
    v7 := GenS1{F0: 9223372036854775807}
    fmt.Println(v6, v2)
    v8 := GenS1{F0: 0}
    var v9 uint = v4
    var v10 int64 = 42
    var v11 string = "\n\t"
    if false {
        var v12 string = v11
        var v13 uint8 = v2
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 14; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 uint8 = v6
                v10 += v10
                _ = fi15
                _ = v16
            }
            var v17 string = (v11 + (("a" + v11) + "\n\t"))
            _ = fi14
            _ = v17
        }
    }
    if v5 {
        v5 = (v11 >= ("\n\t" + v11))
        fmt.Println(v4, v1, v1)
    }
    v18 := GenS1{F0: int64(v6)}
    switch v11 {
    case "foo bar":
        v3 = append(v3, 1)
    case "\n\t":
        v1 -= float32((-2.25 - ((float64(v1) - -2.25) / (float64(1e10) / 1.0))))
    default:
        for fi19 := 0; fi19 < 14; fi19++ {
            switch (^(fi19)) {
            case 32767:
                v20 := GenS1{F0: 134}
                _ = v20
            case 4294967295:
                switch 32767 {
                case -1:
                    v21 := GenS2{F0: 42, F1: "hello", F2: 1}
                    _ = v21
                case -9223372036854775808:
                    var v22 int64 = v10
                    _ = v22
                default:
                    v23 := GenS2{F0: -128, F1: ((v11 + v11) + v11), F2: -128}
                    _ = v23
                }
            }
            _ = fi19
        }
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        switch v11 {
        case "":
            v4 = (^(v9))
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            v11 = v11
            if v5 {
                switch 9223372036854775807 {
                case 381:
                    switch fi24 {
                    case 37:
                        v26 := GenS2{F0: 100, F1: "hello", F2: 100}
                        _ = v26
                    case 708:
                        fmt.Println(v10, fi24)
                    case -392:
                        _ = len(v3)
                    default:
                        switch (v11 + ("" + v11)) {
                        case "a":
                            if (4294967295 > 938) {
                                if v5 {
                                    var v27 string = "\n\t"
                                    if len(v3) > 0 {
                                        v3[0] = 4294967295
                                    }
                                    _ = v27
                                }
                                v1 -= (float32(0.1) * v1)
                            }
                        default:
                            var v28 bool = (32767 <= 42)
                            _ = v28
                        }
                    }
                case -22:
                    v5 = ((!v5) || (!v5))
                }
                v4 = (v9 >> ((v4) & 63))
            }
            if len(v3) > 0 {
                v3[0] = 1
            }
            fmt.Println(v6)
            _ = fi25
        }
        _ = fi24
    }
    switch -450 {
    case 128:
        for fi29 := 0; fi29 < 10; fi29++ {
            if v5 {
                v9 -= (v9 >> (((v9 >> ((243) & 63))) & 63))
                v30 := GenS2{F0: -128, F1: "a", F2: 1}
                _ = v30
            }
            for fi31 := 0; fi31 < 1; fi31++ {
                v3 = append(v3, 4294967295)
                var v32 bool = v5
                fmt.Println(v2, v4)
                _ = fi31
                _ = v32
            }
            _ = fi29
        }
    case -475:
        if true {
            if v5 {
                v33 := GenS2{F0: 1, F1: ("foo bar" + ("\n\t" + "a")), F2: -123}
                _ = v33
            } else {
                var v34 uint = 65535
                if ((2147483648 == 2147483648) || true) {
                    for fi35 := 0; fi35 < 9; fi35++ {
                        var v36 uint8 = (^(v2))
                        for fi37 := 0; fi37 < 5; fi37++ {
                            if len(v3) > 0 {
                                v3[0] = 574
                            }
                            v3 = append(v3, 0)
                            v6 -= v2
                            _ = fi37
                        }
                        v34 = 4294967295
                        _ = fi35
                        _ = v36
                    }
                    var v38 uint = 58
                    if (v5 && (65535 != 255)) {
                        var v39 int16 = 306
                        var v40 uint64 = 4294967295
                        _ = v39
                        _ = v40
                    }
                    _ = v38
                } else {
                    for fi41 := 0; fi41 < 8; fi41++ {
                        v42 := GenS3{F0: v2, F1: 65535, F2: 0}
                        v34 -= v9
                        _ = fi41
                        _ = v42
                    }
                }
                _ = v34
            }
            v11 = v11
        } else {
            v9 = v4
            var v43 int16 = 1000
            var v44 int = 32768
            _ = v43
            _ = v44
        }
    }
    v4 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
