package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int64; F2 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }

func main() {
    v1 := []int{441, 128, -1, 42}
    v2 := make(map[string]uint32)
    var v3 uint32 = 42
    var v4 uint = 4294967295
    var v5 int32 = -2147483648
    switch 128 {
    case 774:
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 float32 = 1e10
            v8 := v2[("foo bar" + ("\n\t" + "foo bar"))]
            if ("" > ("\n\t" + ((((((("\n\t" + "") + "a") + "") + "hello") + "\n\t") + "foo bar") + ""))) {
                v9 := GenS1{F0: 78}
                _ = v9
            } else {
                v3 -= 349
                fmt.Println(v5)
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    case -65:
        for fi10 := 0; fi10 < 1; fi10++ {
            v5 = v5
            var v11 uint = 18446744073709551615
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 int = 94
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
    }
    v3 = 65535
    v1 = append(v1, 1)
    var v14 int = 1
    var v15 int64 = -1
    for fi16 := 0; fi16 < 17; fi16++ {
        switch (fi16 / ((v14) | 1)) {
        case 2147483647:
            v2[""] = (^(v3))
        default:
            for fi17 := 0; fi17 < 20; fi17++ {
                v1 = append(v1, fi17)
                var v18 int16 = -829
                v15 = v15
                _ = fi17
                _ = v18
            }
        }
        if true {
            var v19 int16 = -1
            _ = v19
        }
        switch "hello" {
        case "foo bar":
            v20 := GenS1{F0: 801}
            _ = v20
        case "":
            var v21 uint = (v4 << ((v4) & 63))
            _ = v21
        default:
            var v22 string = ("\n\t" + "foo bar")
            _ = v22
        }
        _ = fi16
    }
    switch (v14 - 2147483647) {
    case 1:
        v2["hello"] = (v3 - v3)
    case 32767:
        v23 := GenS2{F0: int(v4), F1: (^(v15)), F2: 1000}
        _ = v23
    case 42:
        v14 = int(v15)
    default:
        if len(v1) > 0 {
            v1[0] = v14
        }
    }
    if false {
        v5 -= 1
        v3 -= (v3 - (v3 >> ((v4) & 63)))
    }
    v15 = (v15 | int64(v5))
    switch ("hello" + "a") {
    case "hello":
        fmt.Println(v14, v3)
    case "foo bar":
        v5 -= (v5 & (v5 << ((v4) & 63)))
    case "\n\t":
        v24 := GenS2{F0: v14, F1: 828, F2: -100}
        _ = v24
    }
    v25 := GenS3{F0: 60}
    switch (-(v14)) {
    case 123:
        fmt.Println(v4, v4, v4)
    case 127:
        for fi26 := 0; fi26 < 5; fi26++ {
            switch "" {
            case "\n\t":
                fmt.Println(v4)
            case "a":
                for fi27 := 0; fi27 < 20; fi27++ {
                    var v28 uint32 = v3
                    _ = fi27
                    _ = v28
                }
            default:
                switch "hello" {
                case "\n\t":
                    for fi29 := 0; fi29 < 11; fi29++ {
                        v4 -= v4
                        var v30 uint = (v4 | (v4 | (v4 | (v4 & 528))))
                        v15 = -1
                        switch "hello" {
                        case "\n\t":
                            if ((false || true) || (300 > 65535)) {
                                var v31 int32 = v5
                                switch v14 {
                                case 32768:
                                    v31 = v5
                                case -26:
                                    v3 -= v3
                                default:
                                    v32 := GenS1{F0: 1}
                                    _ = v32
                                }
                                v14 = (-(v14))
                                _ = v31
                            }
                        case "hello":
                            var v33 int32 = (v5 ^ (v5 % (((v5 - (v5 - (v5 % ((42) | 1))))) | 1)))
                            _ = v33
                        case "foo bar":
                            switch fi29 {
                            case 127:
                                switch fi29 {
                                case 128:
                                    var v34 int8 = 0
                                    _ = v34
                                default:
                                    var v35 uint8 = 128
                                    if (uint16(v4) <= 680) {
                                        switch "foo bar" {
                                        case "a":
                                            v2[("foo bar" + "a")] = v3
                                        case "hello":
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v36 bool = false
                                                _ = v36
                                            case "\n\t":
                                                v37 := GenS2{F0: v14, F1: (v15 >> ((v4) & 63)), F2: -32768}
                                                _ = v37
                                            case "a":
                                                var v38 uint16 = 2
                                                _ = v38
                                            }
                                        }
                                        var v39 float64 = 0.1
                                        _ = v39
                                    } else {
                                        _ = len(v2)
                                    }
                                    var v40 uint32 = v3
                                    _ = v35
                                    _ = v40
                                }
                            case 949:
                                v41 := GenS4{F0: ("" + "foo bar")}
                                _ = v41
                            }
                        default:
                            v42 := GenS3{F0: 1}
                            _ = v42
                        }
                        _ = fi29
                        _ = v30
                    }
                default:
                    v4 = v4
                }
            }
            v43 := GenS2{F0: v14, F1: (v15 ^ (v15 - v15)), F2: 1000}
            _ = fi26
            _ = v43
        }
    default:
        fmt.Println(v4, v4, v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
