package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int64; F2 int64 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int8; F1 uint64; F2 int }

func main() {
    var v1 int = -2147483648
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    var v4 float32 = float32(float64(float32(0.0)))
    var v5 uint32 = 719
    v6 := make(map[string]uint32)
    v7 := []int{0, 1}
    var v8 uint64 = 18446744073709551615
    fmt.Println(v8, v4)
    for fi9 := 0; fi9 < 1; fi9++ {
        v1 = (v1 ^ fi9)
        delete(v3, "a")
        v4 = float32((3.14159 - -2.25))
        _ = fi9
    }
    if (!true) {
        var v10 int64 = 522
        v11 := GenS1{F0: 127}
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(v5, v5, v1)
            var v13 uint64 = 1
            _ = fi12
            _ = v13
        }
    }
    fmt.Println(v8)
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 int8 = -47
            var v17 int16 = -1
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    v18 := GenS1{F0: 127}
    switch (v1 >> ((0) & 63)) {
    case -1:
        if (false || true) {
            v19 := v3[("a" + "a")]
            switch "foo bar" {
            case "a":
                switch ("\n\t" + "hello") {
                case "":
                    if true {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        var v20 int32 = -2147483648
                        fmt.Println(v8, v5, v4)
                        _ = v20
                    }
                case "foo bar":
                    var v21 int = -302
                    _ = v21
                case "\n\t":
                    v5 = (^(v5))
                }
            case "hello":
                var v22 int8 = 1
                _ = v22
            default:
                if true {
                    for fi23 := 0; fi23 < 16; fi23++ {
                        if (v8 > v8) {
                            v24 := GenS2{F0: v8, F1: -1, F2: 0}
                            for fi25 := 0; fi25 < 6; fi25++ {
                                var v26 uint16 = 42
                                v27 := GenS2{F0: v8, F1: 0, F2: -408}
                                _ = len(v2)
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                            _ = v24
                        } else {
                            v4 = float32(3.14159)
                        }
                        var v28 int64 = 0
                        _ = fi23
                        _ = v28
                    }
                }
            }
            _ = v19
        } else {
            v29 := GenS2{F0: (v8 & (v8 - 1)), F1: -1, F2: int64(v5)}
            _ = v29
        }
    case 42:
        v30 := GenS2{F0: v8, F1: 792, F2: 219}
        _ = v30
    default:
        v4 += v4
    }
    if len(v7) > 0 {
        v7[0] = v1
    }
    v31 := GenS3{F0: -32768, F1: int64(v8)}
    switch "hello" {
    case "hello":
        for fi32 := 0; fi32 < 6; fi32++ {
            var v33 bool = (2 <= (v5 + v5))
            v34 := GenS4{F0: 2, F1: 128}
            _ = fi32
            _ = v33
            _ = v34
        }
    case "\n\t":
        switch (v1 >> ((413) & 63)) {
        case 9223372036854775807:
            for fi35 := 0; fi35 < 12; fi35++ {
                switch 1 {
                case 27:
                    fmt.Println(v1, v8, v5)
                case 1:
                    var v36 uint64 = (v8 & 136)
                    _ = v36
                case 128:
                    v5 -= (v5 ^ 4294967295)
                }
                var v37 uint = 421
                switch (v1 >> (((v37 % (((^(v37))) | 1))) & 63)) {
                case -2147483648:
                    var v38 int = fi35
                    _ = v38
                }
                _ = fi35
                _ = v37
            }
        case 198:
            switch "" {
            case "":
                v4 -= float32(0.0)
            }
        case 127:
            var v39 uint64 = v8
            _ = v39
        default:
            v40 := GenS5{F0: 1, F1: v8, F2: v1}
            _ = v40
        }
    case "a":
        for fi41 := 0; fi41 < 18; fi41++ {
            delete(v6, "")
            _ = fi41
        }
    default:
        if len(v7) > 0 {
            v7[0] = v1
        }
    }
    v5 -= v5
    var v42 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
