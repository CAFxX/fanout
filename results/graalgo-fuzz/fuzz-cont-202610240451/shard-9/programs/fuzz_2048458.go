package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint32, p1 int64, p2 int) int32 {
    fmt.Println(p1)
    v5 := GenS1{F0: 1, F1: -1, F2: 0}
    if (false || true) {
        var v6 int8 = 100
        p0 = p0
        var v7 float64 = (((float64((1.5 * 0.0)) * (0.0 / 1e10)) - ((3.14159 - 0.1) * 0.0)) + 0.1)
        _ = v6
        _ = v7
    } else {
        var v8 uint32 = 2
        switch ((("\n\t" + ("hello" + ("foo bar" + ""))) + "foo bar") + "") {
        case "\n\t":
            var v9 int32 = 2147483647
            _ = v9
        }
        p1 = 547
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    var v1 float32 = 0.1
    v2 := []uint32{1, 0}
    var v3 string = "hello"
    v4 := make(map[uint32]string)
    v1 += v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        switch 32767 {
        case 128:
            var v10 float32 = float32(-2.25)
            _ = v10
        }
        fmt.Println(v1, v3)
        fmt.Println(v1, v1)
    }
    delete(v4, 42)
    for fi11 := 0; fi11 < 19; fi11++ {
        v1 = 0.1
        switch 2147483647 {
        case 1:
            fmt.Println(fi11)
        case 977:
            v12 := GenS1{F0: 225, F1: -17, F2: 603}
            _ = v12
        }
        if (true && (true || false)) {
            var v13 bool = false
            _ = v13
        }
        _ = fi11
    }
    fmt.Println(v1, v3)
    for fi14 := 0; fi14 < 4; fi14++ {
        switch v3 {
        case "a":
            switch "foo bar" {
            case "hello":
                v15 := GenS1{F0: 255, F1: -12, F2: 42}
                _ = v15
            case "foo bar":
                for fi16 := 0; fi16 < 10; fi16++ {
                    delete(v4, 42)
                    _ = fi16
                }
            case "a":
                var v17 uint16 = 65535
                _ = v17
            }
        case "":
            if true {
                v18 := GenS1{F0: 42, F1: -1, F2: 42}
                if (!false) {
                    v19 := GenS2{F0: 271, F1: 255}
                    v1 += float32(1e10)
                    _ = v19
                }
                switch fi14 {
                case -803:
                    for fi20 := 0; fi20 < 18; fi20++ {
                        _ = len(v4)
                        if len(v2) > 0 {
                            v2[0] = 276
                        }
                        _ = fi20
                    }
                default:
                    v21 := v4[2]
                    _ = v21
                }
                _ = v18
            }
        default:
            for fi22 := 0; fi22 < 17; fi22++ {
                if (2147483647 == 21) {
                    switch fi14 {
                    case 673:
                        var v23 bool = true
                        _ = v23
                    case 2147483647:
                        v1 -= float32(-2.25)
                        var v24 int8 = 42
                        if (813 != 2) {
                            fmt.Println(fi14)
                            switch fi22 {
                            case -867:
                                v25 := GenS2{F0: 0, F1: 127}
                                _ = v25
                            case 9223372036854775807:
                                if (!true) {
                                    v3 = "a"
                                }
                            }
                            fmt.Println(fi22)
                        }
                        _ = v24
                    case -1:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        _ = len(v4)
                        v26 := GenS1{F0: 255, F1: -59, F2: 2}
                        _ = v26
                    }
                }
                var v27 uint16 = 1
                var v28 int8 = 119
                _ = len(v2)
                _ = fi22
                _ = v27
                _ = v28
            }
        }
        _ = fi14
    }
    var v29 string = v3
    var v30 int16 = 237
    v30 = (-(v30))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
