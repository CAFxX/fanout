package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int16 }
func genFn0(p0 int64) uint32 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 18; fi5++ {
        p0 = (^(p0))
        fmt.Println(p0)
        _ = fi5
    }
    _ = p0
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 127
    v3 := make(map[uint32]string)
    var v4 uint16 = 779
    switch 0 {
    case 42:
        v6 := v3[1]
        _ = v6
    case 32767:
        switch ("a" + ("" + ("a" + "foo bar"))) {
        case "hello":
            for fi7 := 0; fi7 < 17; fi7++ {
                if (!false) {
                    var v8 int32 = 42
                    var v9 uint64 = uint64(v4)
                    v2 += v2
                    _ = v8
                    _ = v9
                } else {
                    _ = len(v3)
                    v10 := GenS1{F0: 2, F1: -32768}
                    var v11 float64 = 3.14159
                    _ = v10
                    _ = v11
                }
                if (false || (true && (false || true))) {
                    if ((!(1 > 0)) || (127 != fi7)) {
                        fmt.Println(v4, fi7)
                        delete(v1, 65535)
                        v12 := GenS2{F0: 2147483647, F1: 358, F2: -32768}
                        _ = v12
                    } else {
                        var v13 uint8 = 2
                        switch "a" {
                        case "\n\t":
                            switch (((("a" + "hello") + "foo bar") + "") + "foo bar") {
                            case "foo bar":
                                var v14 bool = (false || (true || (!(433 <= 2))))
                                _ = v14
                            case "a":
                                _ = len(v3)
                            case "\n\t":
                                var v15 uint8 = 196
                                _ = v15
                            }
                        case "hello":
                            var v16 int32 = 42
                            _ = v16
                        default:
                            v13 -= (^(v13))
                        }
                        _ = v13
                    }
                    v2 -= v2
                }
                var v17 uint8 = 127
                _ = fi7
                _ = v17
            }
        }
    case 9223372036854775807:
        v18 := GenS1{F0: 2147483648, F1: 0}
        _ = v18
    }
    v3[65535] = "a"
    if ((v2 / ((v2) | 1)) >= (v2 & (v2 / (((v2 | (v2 - 1))) | 1)))) {
        switch ("\n\t" + "hello") {
        case "foo bar":
            fmt.Println(v2)
        case "":
            var v19 uint8 = 106
            _ = v19
        default:
            fmt.Println(v2, v2)
        }
    }
    var v20 int = 128
    var v21 int = (v20 & (^(v20)))
    if false {
        var v22 uint32 = 1
        var v23 int32 = 243
        switch -9223372036854775808 {
        case 9223372036854775807:
            for fi24 := 0; fi24 < 19; fi24++ {
                switch "\n\t" {
                case "":
                    var v25 uint16 = 42
                    _ = v25
                }
                var v26 uint32 = (v22 * (v22 * 4294967295))
                _ = fi24
                _ = v26
            }
        case 4294967295:
            v27 := GenS2{F0: v22, F1: (^(v22)), F2: 32767}
            _ = v27
        case 98:
            var v28 uint64 = 18446744073709551615
            _ = v28
        default:
            v29 := GenS2{F0: genFn0(-1), F1: genFn0(9223372036854775807), F2: -633}
            _ = v29
        }
        _ = v22
        _ = v23
    }
    v30 := GenS2{F0: 2, F1: 4294967295, F2: 333}
    if (1 == 874) {
        switch 127 {
        case -1:
            var v31 uint32 = 693
            _ = v31
        case -2147483648:
            var v32 uint = 1
            _ = v32
        case -9223372036854775808:
            var v33 float32 = -2.25
            _ = v33
        }
        var v34 int8 = v2
        _ = v34
    }
    delete(v1, 610)
    if true {
        var v35 float64 = float64(((float32(3.14159) + 0.0) - 0.1))
        _ = v35
    }
    v2 += (v2 ^ (v2 - (v2 | (v2 << ((601) & 63)))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v4
}
