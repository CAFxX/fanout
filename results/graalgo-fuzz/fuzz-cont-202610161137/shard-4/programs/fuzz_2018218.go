package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint32 }

func main() {
    var v1 int16 = 42
    var v2 int32 = -856
    var v3 uint64 = 2
    var v4 uint8 = 128
    switch 955 {
    case 9223372036854775807:
        v5 := GenS1{F0: 1000}
        _ = v5
    case 32767:
        v6 := []int{128, -912}
        _ = v6
    }
    fmt.Println(v2, v2, v1)
    var v7 uint64 = (^(v3))
    v8 := []int{32767, 128}
    fmt.Println(v7, v4, v2)
    var v9 int16 = v1
    var v10 float64 = (float64(float32(((float64(3.14159) * (1.5 * float64(0.0))) - 1e10))) + (0.1 * (((1e10 + -2.25) / -2.25) + (1e10 * ((1.5 * 0.0) * (float64(0.0) + 3.14159))))))
    for fi11 := 0; fi11 < 12; fi11++ {
        switch 128 {
        case -342:
            for fi12 := 0; fi12 < 20; fi12++ {
                v13 := GenS1{F0: v9}
                v14 := GenS2{F0: 65535, F1: (v3 + (v7 << ((0) & 63)))}
                _ = fi12
                _ = v13
                _ = v14
            }
        case 4294967295:
            switch ("foo bar" + (("" + "") + "\n\t")) {
            case "a":
                switch (^(fi11)) {
                case 742:
                    switch "foo bar" {
                    case "foo bar":
                        var v15 uint8 = 127
                        _ = v15
                    case "a":
                        v16 := GenS2{F0: 1, F1: (^(v7))}
                        _ = v16
                    }
                case 4294967295:
                    for fi17 := 0; fi17 < 2; fi17++ {
                        _ = len(v8)
                        var v18 bool = false
                        var v19 uint = 42
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                case 436:
                    for fi20 := 0; fi20 < 7; fi20++ {
                        var v21 int16 = 436
                        var v22 uint = 1
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
            default:
                v8 = append(v8, (fi11 ^ fi11))
            }
        case -2147483648:
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        _ = fi11
    }
    _ = len(v8)
    if ((("a" + "foo bar") < (((("foo bar" + "hello") + ((("foo bar" + "foo bar") + ((("" + "hello") + "") + "a")) + "a")) + ("hello" + "foo bar")) + "foo bar")) || (!false)) {
        var v23 uint32 = 76
        v24 := GenS2{F0: 42, F1: (v3 >> ((18446744073709551615) & 63))}
        if len(v8) > 0 {
            v8[0] = 32767
        }
        _ = v23
        _ = v24
    } else {
        var v25 uint32 = 640
        if true {
            var v26 float32 = 0.1
            var v27 uint64 = v7
            _ = v26
            _ = v27
        } else {
            var v28 int16 = 1000
            switch 961 {
            case 32767:
                var v29 int = -776
                _ = v29
            case 2147483647:
                v30 := GenS2{F0: 4294967295, F1: v3}
                _ = v30
            case 1:
                if true {
                    switch 42 {
                    case -726:
                        var v31 uint64 = v7
                        _ = v31
                    }
                    v25 = (v25 ^ 2)
                    switch -1 {
                    case -1:
                        var v32 int16 = 796
                        _ = v32
                    }
                } else {
                    if ((true || true) || (v1 <= v28)) {
                        for fi33 := 0; fi33 < 13; fi33++ {
                            var v34 uint32 = (^(v25))
                            if false {
                                for fi35 := 0; fi35 < 15; fi35++ {
                                    v36 := GenS3{F0: uint32(v3)}
                                    v25 = v25
                                    _ = fi35
                                    _ = v36
                                }
                                v8 = append(v8, -1)
                            }
                            for fi37 := 0; fi37 < 10; fi37++ {
                                if true {
                                    var v38 int = fi37
                                    switch ("" + "") {
                                    case "a":
                                        if ((1e10 * 1.5) >= 0.0) {
                                            v39 := GenS4{F0: 32767}
                                            _ = v39
                                        } else {
                                            var v40 int = fi37
                                            switch (fi37 | -764) {
                                            case -2147483648:
                                                var v41 int32 = v2
                                                _ = v41
                                            case 9223372036854775807:
                                                var v42 float64 = (v10 + 0.0)
                                                _ = v42
                                            }
                                            _ = v40
                                        }
                                    default:
                                        var v43 uint16 = 2
                                        _ = v43
                                    }
                                    _ = v38
                                } else {
                                    v44 := GenS5{F0: -128, F1: v25, F2: v25}
                                    _ = v44
                                }
                                v45 := GenS5{F0: 97, F1: 682, F2: (v34 / ((2147483648) | 1))}
                                var v46 uint = 2
                                _ = fi37
                                _ = v45
                                _ = v46
                            }
                            _ = fi33
                            _ = v34
                        }
                    }
                    fmt.Println(v3, v7, v2)
                    switch "" {
                    case "foo bar":
                        v7 -= 18446744073709551615
                    case "":
                        if false {
                            fmt.Println(v7, v3, v4)
                        } else {
                            fmt.Println(v9, v1, v10)
                        }
                    case "a":
                        var v47 uint16 = 32767
                        _ = v47
                    }
                }
            }
            _ = v28
        }
        v1 += v1
        _ = v25
    }
    if len(v8) > 0 {
        v8[0] = 32768
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
