package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint64 = 0
    var v2 int16 = 1
    v3 := make(map[string]bool)
    v4 := []uint32{586, 385, 357, 674}
    var v5 float32 = float32(float64(0.1))
    switch "foo bar" {
    case "\n\t":
        switch -557 {
        case 127:
            for fi6 := 0; fi6 < 12; fi6++ {
                var v7 int8 = 42
                if false {
                    if (0 == 2) {
                        v2 += (v2 & v2)
                        var v8 bool = true
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                        _ = v8
                    }
                    if false {
                        var v9 uint8 = 255
                        var v10 bool = (v5 != v5)
                        _ = v10
                        _ = v9
                    }
                    v5 -= v5
                }
                _ = fi6
                _ = v7
            }
        case 60:
            var v11 uint32 = 2147483648
            _ = v11
        }
    case "a":
        for fi12 := 0; fi12 < 12; fi12++ {
            v5 -= v5
            if false {
                v1 = 18446744073709551615
                v4 = append(v4, 4294967295)
            } else {
                switch fi12 {
                case 114:
                    for fi13 := 0; fi13 < 8; fi13++ {
                        var v14 float32 = 1.5
                        _ = len(v3)
                        _ = fi13
                        _ = v14
                    }
                case 0:
                    v5 -= v5
                default:
                    v5 = v5
                }
                var v15 float64 = -2.25
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 uint32 = 64
                    fmt.Println(v5, fi12, fi12)
                    v15 = v15
                    v2 = (-(v2))
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
            _ = fi12
        }
    case "foo bar":
        var v18 uint = 0
        _ = v18
    }
    var v19 uint32 = 4294967295
    if false {
        switch "" {
        case "":
            if len(v4) > 0 {
                v4[0] = 567
            }
        default:
            v20 := GenS1{F0: 127}
            _ = v20
        }
        v19 += v19
        var v21 uint8 = 2
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 float64 = float64(v5)
            _ = fi23
            _ = v24
        }
        for fi25 := 0; fi25 < 17; fi25++ {
            fmt.Println(v5, fi22, fi25)
            if false {
                var v26 int32 = -699
                var v27 int64 = -1
                _ = v26
                _ = v27
            }
            _ = fi25
        }
        _ = fi22
    }
    if true {
        v28 := GenS2{F0: -55}
        if len(v4) > 0 {
            v4[0] = (v19 >> ((4294967295) & 63))
        }
        _ = v28
    } else {
        var v29 int = 695
        v29 -= int(v2)
        switch (v29 + -303) {
        case 127:
            for fi30 := 0; fi30 < 12; fi30++ {
                switch (v29 / (((v29 | (fi30 >> ((1) & 63)))) | 1)) {
                case 68:
                    delete(v3, "foo bar")
                case -1:
                    for fi31 := 0; fi31 < 15; fi31++ {
                        v2 -= v2
                        _ = fi31
                    }
                case 74:
                    fmt.Println(v5)
                }
                _ = fi30
            }
        default:
            var v32 int = v29
            _ = v32
        }
        _ = v29
    }
    if len(v4) > 0 {
        v4[0] = (v19 ^ v19)
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        switch (fi33 - fi33) {
        case 9223372036854775807:
            v4 = append(v4, (v19 / ((v19) | 1)))
        case 1:
            fmt.Println(v5)
        case 128:
            for fi34 := 0; fi34 < 7; fi34++ {
                v1 = v1
                var v35 uint8 = 92
                _ = fi34
                _ = v35
            }
        default:
            fmt.Println(v5, fi33)
        }
        _ = len(v4)
        v19 += (v19 % ((v19) | 1))
        _ = fi33
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    delete(v3, ("\n\t" + (("hello" + "\n\t") + "hello")))
    fmt.Println(v5, v2, v5)
    v2 += 42
    var v36 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
