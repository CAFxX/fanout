package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int64; F1 int8; F2 uint32 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }

func main() {
    var v1 int16 = 42
    var v2 int = 2147483647
    v3 := make(map[string]bool)
    v4 := []string{"foo bar", ("foo bar" + "hello"), "a", "a"}
    v5 := make(map[uint32]string)
    var v6 string = "a"
    switch -240 {
    case 32768:
        var v7 int32 = 0
        _ = v7
    default:
        switch v2 {
        case 1:
            for fi8 := 0; fi8 < 6; fi8++ {
                if ((100 > 127) || false) {
                    var v9 uint64 = 4294967295
                    v2 += (fi8 ^ (fi8 << ((18446744073709551615) & 63)))
                    delete(v5, 2)
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 5; fi10++ {
                        var v11 float64 = 0.1
                        switch fi10 {
                        case 1:
                            if (231 <= -593) {
                                _ = len(v5)
                                v1 += 847
                            }
                        case 375:
                            v12 := GenS1{F0: 255, F1: 14}
                            _ = v12
                        case -998:
                            switch (v6 + v6) {
                            case "":
                                var v13 float32 = (0.1 / 1e10)
                                _ = v13
                            }
                        }
                        _ = fi10
                        _ = v11
                    }
                    var v14 int = 2147483647
                    if len(v4) > 0 {
                        v4[0] = v6
                    }
                    _ = v14
                }
                v5[65535] = ("" + v6)
                _ = fi8
            }
        }
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        if (true || (86 > 255)) {
            v16 := GenS2{F0: -1, F1: 65, F2: 1}
            fmt.Println(v1, v1)
            switch 127 {
            case -9223372036854775808:
                var v17 uint64 = 1
                _ = v17
            case 9223372036854775807:
                if (!(0.1 == float32(1.5))) {
                    v18 := GenS1{F0: 149, F1: 42}
                    fmt.Println(v6)
                    _ = v18
                } else {
                    var v19 uint16 = 32767
                    _ = v19
                }
            default:
                v20 := GenS1{F0: 42, F1: 40}
                _ = v20
            }
            _ = v16
        }
        v2 = v2
        _ = fi15
    }
    switch v6 {
    case "":
        var v21 float64 = 1.5
        _ = v21
    }
    if (false && ((v6 < v6) || (482 == 0))) {
        v22 := GenS3{F0: 70, F1: -32768}
        _ = v22
    }
    var v23 uint8 = 67
    if (true || true) {
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 int = fi24
            if true {
                for fi26 := 0; fi26 < 11; fi26++ {
                    var v27 uint16 = 42
                    var v28 int64 = 9223372036854775807
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = len(v3)
                if true {
                    if (42 > 217) {
                        fmt.Println(v2, fi24, v6)
                        delete(v5, 868)
                        if true {
                            v29 := v3["a"]
                            fmt.Println(v25)
                            _ = v29
                        }
                    } else {
                        for fi30 := 0; fi30 < 18; fi30++ {
                            v31 := GenS4{F0: (fi30 | v2)}
                            var v32 float32 = (float32((0.0 - 0.1)) * 1.5)
                            var v33 uint64 = 18446744073709551615
                            _ = fi30
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                    }
                    v34 := GenS3{F0: -9223372036854775808, F1: v1}
                    v35 := GenS5{F0: 9223372036854775807}
                    _ = v34
                    _ = v35
                }
            }
            _ = fi24
            _ = v25
        }
        var v36 int = v2
        switch -843 {
        case -1:
            v1 = v1
        }
        _ = v36
    } else {
        for fi37 := 0; fi37 < 17; fi37++ {
            v23 += v23
            fmt.Println(fi37)
            _ = fi37
        }
        _ = len(v4)
    }
    v38 := GenS4{F0: (^(v2))}
    fmt.Println(v23)
    for fi39 := 0; fi39 < 1; fi39++ {
        v40 := GenS2{F0: 1, F1: 93, F2: 65535}
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
