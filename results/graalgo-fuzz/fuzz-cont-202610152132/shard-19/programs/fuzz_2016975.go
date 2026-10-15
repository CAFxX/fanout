package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS5 struct { F0 int; F1 int64; F2 int64 }
type GenS6 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 uint8; F1 int64; F2 int32 }

func main() {
    var v1 uint32 = 787
    v2 := []string{"hello"}
    var v3 uint64 = 2
    v4 := make(map[int]bool)
    for fi5 := 0; fi5 < 16; fi5++ {
        if true {
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 bool = (!true)
                var v8 bool = (65535 <= 602)
                if v8 {
                    v9 := GenS1{F0: 9223372036854775807}
                    if (v3 < v3) {
                        v8 = v8
                        var v10 int = fi6
                        _ = v10
                    }
                    _ = v9
                }
                _ = fi6
                _ = v7
                _ = v8
            }
            for fi11 := 0; fi11 < 5; fi11++ {
                v12 := GenS1{F0: 0}
                v13 := GenS2{F0: uint16(v1), F1: 1}
                switch (-(fi11)) {
                case 811:
                    v14 := GenS3{F0: v1}
                    v15 := v4[(-(fi11))]
                    if v15 {
                        var v16 uint64 = uint64(v1)
                        var v17 string = "foo bar"
                        fmt.Println(fi5)
                        v4[(fi5 * 889)] = v15
                        _ = v16
                        _ = v17
                    }
                    _ = v14
                    _ = v15
                case -2147483648:
                    v1 = (v1 & (^(v1)))
                case 42:
                    delete(v4, (fi11 | fi11))
                default:
                    v18 := GenS4{F0: 51, F1: 65535, F2: 151}
                    _ = v18
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        } else {
            v19 := GenS5{F0: (fi5 & int(v1)), F1: 1, F2: 261}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v20 int32 = 315
            _ = v19
            _ = v20
        }
        _ = fi5
    }
    var v21 float64 = 0.0
    var v22 int8 = 127
    switch "" {
    case "foo bar":
        var v23 int16 = 42
        _ = v23
    case "":
        v22 += (v22 % ((-128) | 1))
    case "hello":
        switch 9223372036854775807 {
        case 36:
            v22 += v22
        default:
            for fi24 := 0; fi24 < 8; fi24++ {
                switch fi24 {
                case 416:
                    v4[(fi24 ^ 42)] = (!true)
                case 4294967295:
                    var v25 int8 = 0
                    _ = v25
                }
                if false {
                    v4[(fi24 % ((fi24) | 1))] = false
                    v26 := GenS5{F0: fi24, F1: 781, F2: 948}
                    _ = v26
                }
                v27 := GenS2{F0: 1, F1: -901}
                _ = fi24
                _ = v27
            }
        }
    }
    var v28 int64 = 588
    v29 := GenS4{F0: 0, F1: 169, F2: 0}
    v28 = 0
    var v30 uint64 = v3
    fmt.Println(v21, v1)
    v21 = 1.5
    fmt.Println(v30, v28)
    switch "foo bar" {
    case "a":
        v2 = append(v2, ("" + ("a" + "")))
    case "foo bar":
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 int64 = 9223372036854775807
            v33 := GenS6{F0: -1, F1: 62, F2: v30}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi31
            _ = v32
            _ = v33
        }
    case "":
        if false {
            var v34 float32 = 0.1
            v35 := GenS7{F0: true, F1: 1}
            _ = v34
            _ = v35
        } else {
            var v36 bool = false
            for fi37 := 0; fi37 < 3; fi37++ {
                fmt.Println(v21, v21, v22)
                if true {
                    for fi38 := 0; fi38 < 15; fi38++ {
                        var v39 float32 = (1e10 * 0.1)
                        v40 := GenS4{F0: 42, F1: 65535, F2: 255}
                        fmt.Println(fi38, v1)
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    v41 := GenS4{F0: 42, F1: 0, F2: 127}
                    _ = v41
                } else {
                    var v42 int64 = v28
                    _ = v42
                }
                _ = fi37
            }
            v43 := GenS8{F0: 127, F1: (v28 ^ (v28 - (v28 / (((v28 - int64(v22))) | 1)))), F2: 42}
            _ = v36
            _ = v43
        }
    }
    delete(v4, int(v28))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
