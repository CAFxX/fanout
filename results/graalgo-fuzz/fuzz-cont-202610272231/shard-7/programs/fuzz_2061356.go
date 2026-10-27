package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint }

func main() {
    var v1 uint64 = 4294967295
    v2 := []uint32{4294967295, 2147483647, 2}
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 1
    v5 := []string{"foo bar", (("\n\t" + "a") + "a"), ""}
    v1 = 0
    v4 -= v1
    switch 127 {
    case 9223372036854775807:
        fmt.Println(v1, v1)
    case 717:
        var v6 int8 = 1
        _ = v6
    case -9223372036854775808:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
    }
    v7 := GenS1{F0: 1}
    var v8 bool = true
    v9 := GenS2{F0: 42, F1: 914, F2: 65535}
    for fi10 := 0; fi10 < 5; fi10++ {
        switch fi10 {
        case 9223372036854775807:
            if (v8 || false) {
                fmt.Println(v4, v8, v8)
            } else {
                if false {
                    var v11 int64 = 263
                    var v12 int32 = 393
                    var v13 float64 = 0.1
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    var v14 int16 = 1
                    _ = v14
                }
                v15 := GenS2{F0: 42, F1: 0, F2: 163}
                _ = v15
            }
        case 32767:
            if (v8 && (((-779 == -32768) && (0 > v4)) || v8)) {
                switch ("foo bar" + "hello") {
                case "a":
                    var v16 float64 = 1e10
                    _ = v16
                default:
                    for fi17 := 0; fi17 < 2; fi17++ {
                        if true {
                            v18 := GenS2{F0: 1, F1: 18446744073709551615, F2: 2}
                            _ = v18
                        } else {
                            v19 := GenS3{F0: 127, F1: 535}
                            fmt.Println(v8, v8)
                            var v20 int64 = 152
                            _ = v19
                            _ = v20
                        }
                        _ = fi17
                    }
                }
                switch ((("\n\t" + "") + "a") + ("" + "foo bar")) {
                case "hello":
                    for fi21 := 0; fi21 < 5; fi21++ {
                        _ = len(v3)
                        var v22 int16 = 393
                        var v23 int32 = 364
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                }
                var v24 uint32 = 2
                _ = v24
            }
        }
        v8 = v8
        v4 = v4
        _ = fi10
    }
    if (true || true) {
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 int = fi25
            var v27 int16 = 94
            _ = fi25
            _ = v26
            _ = v27
        }
        v1 -= (v4 * 0)
    }
    v1 += (v4 << ((4294967295) & 63))
    var v28 int16 = 1
    var v29 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
