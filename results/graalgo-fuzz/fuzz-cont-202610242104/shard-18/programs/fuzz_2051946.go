package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int64 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint64) string {
    v7 := GenS1{F0: 1, F1: -1, F2: -468}
    var v8 int8 = 127
    v9 := GenS1{F0: -790, F1: v8, F2: -769}
    v8 = 0
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return "a"
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]bool)
    var v3 string = ""
    v4 := make(map[string]string)
    v5 := []string{}
    v6 := make(map[int]bool)
    switch v3 {
    case "a":
        var v10 bool = true
        _ = v10
    case "":
        v11 := GenS1{F0: 92, F1: 127, F2: 0}
        _ = v11
    default:
        for fi12 := 0; fi12 < 13; fi12++ {
            if ((1.5 >= 0.1) || false) {
                var v13 int16 = 42
                switch (fi12 % ((fi12) | 1)) {
                case 2147483647:
                    v2[(v1 >> ((433) & 63))] = true
                default:
                    v13 = (^(v13))
                }
                var v14 uint64 = 71
                _ = v13
                _ = v14
            } else {
                v3 = ("hello" + "hello")
                for fi15 := 0; fi15 < 15; fi15++ {
                    if true {
                        var v16 int8 = 119
                        _ = len(v4)
                        v17 := v2[v1]
                        _ = v16
                        _ = v17
                    }
                    v1 += 642
                    v3 = (("\n\t" + "foo bar") + ("" + (("" + ("hello" + ("a" + (((v3 + v3) + "a") + (v3 + v3))))) + "hello")))
                    var v18 uint16 = 0
                    _ = fi15
                    _ = v18
                }
            }
            _ = fi12
        }
    }
    v3 = ((v3 + "foo bar") + "")
    for fi19 := 0; fi19 < 13; fi19++ {
        if ((1.5 != 0.1) && ((255 >= 0) || (2 != 469))) {
            if (!(63 <= 100)) {
                v20 := GenS2{F0: 1, F1: fi19}
                if (1.5 <= ((0.1 * 0.0) + 1.5)) {
                    for fi21 := 0; fi21 < 15; fi21++ {
                        fmt.Println(fi21, v3)
                        _ = len(v5)
                        var v22 uint = 4294967295
                        var v23 uint8 = 128
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    v1 = v1
                    for fi24 := 0; fi24 < 18; fi24++ {
                        if true {
                            var v25 uint64 = 42
                            v26 := GenS3{F0: 9223372036854775807, F1: (v25 << ((42) & 63))}
                            _ = len(v4)
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS4{F0: (!((true || true) && (2147483647 == 1)))}
                        v3 = "hello"
                        _ = fi24
                        _ = v27
                    }
                }
                _ = v20
            }
        }
        _ = fi19
    }
    if true {
        var v28 int = 42
        _ = v28
    } else {
        if true {
            v1 = v1
        }
    }
    var v29 string = v3
    v29 = v29
    v1 = 4294967295
    for fi30 := 0; fi30 < 1; fi30++ {
        fmt.Println(v1)
        var v31 int8 = 100
        _ = fi30
        _ = v31
    }
    if true {
        switch v29 {
        case "":
            v32 := GenS5{F0: 65, F1: (^(v1))}
            _ = v32
        case "hello":
            var v33 int16 = 1
            _ = v33
        }
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 string = ""
            _ = fi34
            _ = v35
        }
    } else {
        switch 32767 {
        case 4294967295:
            if true {
                v29 = v29
                var v36 int32 = 270
                v29 = ""
                _ = v36
            } else {
                v37 := GenS1{F0: -9223372036854775808, F1: -72, F2: 1}
                _ = v37
            }
        case -1:
            if (!((((3.14159 / 1e10) >= float64((0.0 - ((0.0 / -2.25) * (-2.25 - 1e10))))) || false) && true)) {
                v38 := GenS6{F0: 42}
                for fi39 := 0; fi39 < 20; fi39++ {
                    switch (fi39 * (-(fi39))) {
                    case 701:
                        fmt.Println(v1, v3)
                    default:
                        for fi40 := 0; fi40 < 13; fi40++ {
                            switch v3 {
                            case "hello":
                                v41 := v4["a"]
                                _ = v41
                            }
                            v42 := GenS6{F0: -1}
                            v43 := GenS1{F0: 42, F1: 100, F2: 0}
                            _ = fi40
                            _ = v42
                            _ = v43
                        }
                    }
                    _ = fi39
                }
                var v44 uint64 = 489
                _ = v38
                _ = v44
            }
        case -943:
            v29 = ""
        default:
            if (float32(1.5) > ((3.14159 * (0.1 - 0.0)) / -2.25)) {
                _ = len(v5)
                v1 = (v1 >> ((42) & 63))
                v1 += (v1 / ((65535) | 1))
            }
        }
        v3 = ((v3 + v3) + "\n\t")
        for fi45 := 0; fi45 < 4; fi45++ {
            v1 -= v1
            v1 += (v1 | (v1 - (v1 >> ((353) & 63))))
            _ = fi45
        }
    }
    v46 := GenS4{F0: false}
    v1 = (v1 % (((v1 % (((^(v1))) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
