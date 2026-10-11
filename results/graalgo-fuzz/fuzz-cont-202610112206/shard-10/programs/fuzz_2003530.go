package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 467
    v3 := []uint32{2, 65535, 2147483647}
    var v4 string = ""
    var v5 string = "a"
    for fi6 := 0; fi6 < 10; fi6++ {
        switch fi6 {
        case -2147483648:
            _ = len(v1)
        }
        var v7 int64 = 1
        for fi8 := 0; fi8 < 17; fi8++ {
            _ = len(v3)
            v5 = v4
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 250, F1: -9223372036854775808, F2: -84}
        if true {
            for fi11 := 0; fi11 < 18; fi11++ {
                v4 = v5
                _ = fi11
            }
            for fi12 := 0; fi12 < 15; fi12++ {
                v2 = v2
                _ = fi12
            }
            var v13 uint = 18446744073709551615
            _ = v13
        }
        v14 := GenS1{F0: 65535, F1: 42, F2: -32768}
        _ = fi9
        _ = v10
        _ = v14
    }
    switch 123 {
    case 32767:
        v15 := GenS1{F0: 0, F1: -698, F2: 1}
        _ = v15
    case 4294967295:
        switch "hello" {
        case "hello":
            v16 := GenS1{F0: 1, F1: -1, F2: 0}
            _ = v16
        }
    default:
        for fi17 := 0; fi17 < 9; fi17++ {
            fmt.Println(v5, v2, v5)
            _ = fi17
        }
    }
    var v18 int8 = -122
    v2 -= v2
    var v19 float32 = 3.14159
    var v20 uint16 = (v2 * v2)
    var v21 uint64 = 2
    for fi22 := 0; fi22 < 8; fi22++ {
        v23 := GenS1{F0: 65535, F1: 205, F2: 1}
        v4 = (((v4 + v4) + v5) + (("foo bar" + v4) + v4))
        _ = fi22
        _ = v23
    }
    switch 0 {
    case 32768:
        var v24 float32 = (1e10 - float32(1e10))
        _ = v24
    default:
        var v25 int64 = 226
        _ = v25
    }
    if true {
        var v26 bool = (992 <= -353)
        _ = v26
    } else {
        v19 = float32(3.14159)
        var v27 float32 = v19
        var v28 int64 = 9223372036854775807
        _ = v27
        _ = v28
    }
    v29 := GenS2{F0: 1}
    if false {
        var v30 int64 = 130
        if false {
            for fi31 := 0; fi31 < 11; fi31++ {
                v32 := GenS1{F0: 0, F1: (v30 + (v30 >> ((18446744073709551615) & 63))), F2: -1}
                var v33 float32 = (v19 * float32((1.5 / 1.5)))
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        _ = v30
    }
    v2 = (v2 | 972)
    if ((^(v21)) <= 400) {
        var v34 int32 = -1
        for fi35 := 0; fi35 < 1; fi35++ {
            switch "" {
            case "\n\t":
                for fi36 := 0; fi36 < 20; fi36++ {
                    switch v4 {
                    case "foo bar":
                        v2 += 42
                    case "":
                        v4 = "foo bar"
                    default:
                        var v37 uint8 = 0
                        _ = v37
                    }
                    if false {
                        v2 -= v2
                        _ = len(v1)
                    } else {
                        switch (fi36 ^ fi36) {
                        case -1:
                            var v38 int8 = v18
                            _ = v38
                        }
                    }
                    v18 -= v18
                    _ = fi36
                }
            case "hello":
                for fi39 := 0; fi39 < 2; fi39++ {
                    v1[fi39] = (fi39 << ((18446744073709551615) & 63))
                    v40 := GenS3{F0: (^(v18)), F1: (v18 - 42)}
                    _ = fi39
                    _ = v40
                }
            default:
                for fi41 := 0; fi41 < 17; fi41++ {
                    v18 = v18
                    _ = fi41
                }
            }
            if len(v3) > 0 {
                v3[0] = 65535
            }
            _ = fi35
        }
        _ = v34
    } else {
        if (((v19 - v19) * 3.14159) < v19) {
            var v42 int16 = 299
            v42 += 1000
            for fi43 := 0; fi43 < 10; fi43++ {
                if len(v3) > 0 {
                    v3[0] = uint32(v20)
                }
                v44 := GenS4{F0: false, F1: 2147483647}
                _ = fi43
                _ = v44
            }
            _ = v42
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
