package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int16; F1 uint32 }

func main() {
    var v1 uint8 = 2
    v2 := []uint32{2147483647}
    var v3 float64 = 1e10
    var v4 int64 = 42
    var v5 int = 4294967295
    v6 := make(map[uint32]int)
    var v7 string = "\n\t"
    v6[629] = 4294967295
    for fi8 := 0; fi8 < 3; fi8++ {
        v3 = 3.14159
        _ = len(v6)
        _ = fi8
    }
    fmt.Println(v5)
    v9 := GenS1{F0: (v5 - v5), F1: 0, F2: -96}
    if (((0.1 / (3.14159 * -2.25)) * (1e10 - 0.0)) > (0.0 + 0.0)) {
        if false {
            v5 = (^(v5))
            for fi10 := 0; fi10 < 2; fi10++ {
                if (true && true) {
                    var v11 uint8 = v1
                    v3 -= ((-2.25 * v3) - v3)
                    v12 := GenS2{F0: (v4 % ((v4) | 1))}
                    _ = v11
                    _ = v12
                }
                _ = fi10
            }
        }
    } else {
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS3{F0: 2147483647}
            v4 -= v4
            _ = fi13
            _ = v14
        }
    }
    var v15 int32 = 972
    v1 -= v1
    if (v4 < int64(v5)) {
        fmt.Println(v5, v5, v4)
    } else {
        var v16 float32 = 0.1
        if (42 > 1) {
            v17 := GenS3{F0: (v15 / (((v15 - (v15 - (v15 / ((v15) | 1))))) | 1))}
            _ = v17
        }
        _ = v16
    }
    v18 := GenS4{F0: true, F1: v15}
    var v19 int8 = 42
    if (!false) {
        fmt.Println(v1, v15, v1)
        if len(v2) > 0 {
            v2[0] = 2
        }
    }
    v3 = float64((3.14159 + 0.0))
    v7 = ("hello" + v7)
    if (!((v5 & (v5 * (v5 | v5))) == v5)) {
        if (v19 > 127) {
            v20 := GenS3{F0: (v15 * v15)}
            fmt.Println(v4, v4, v19)
            var v21 float64 = v3
            _ = v20
            _ = v21
        }
        switch (v7 + "") {
        case "\n\t":
            v15 += v15
        case "hello":
            if true {
                fmt.Println(v1, v5)
                var v22 uint8 = (v1 + 42)
                _ = v22
            }
        default:
            switch (v5 << ((42) & 63)) {
            case -630:
                var v23 int16 = 0
                _ = v23
            case 921:
                v24 := GenS3{F0: (-(v15))}
                _ = v24
            case 9223372036854775807:
                v1 -= v1
            }
        }
        for fi25 := 0; fi25 < 1; fi25++ {
            if (968 <= 1) {
                if (v3 > v3) {
                    switch fi25 {
                    case 1:
                        v15 = 0
                    case 4294967295:
                        if ((v19 ^ v19) >= v19) {
                            var v26 string = (v7 + (v7 + ("a" + v7)))
                            if (true && (true && false)) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v27 uint16 = 2
                                v6[2] = 32768
                                _ = v27
                            }
                            v6[2147483648] = (fi25 / (((v5 / ((fi25) | 1))) | 1))
                            _ = v26
                        } else {
                            var v28 int16 = int16(v1)
                            var v29 string = "\n\t"
                            v30 := GenS5{F0: (v28 & (v28 - -409)), F1: 2}
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    }
                }
                fmt.Println(v19, fi25, v19)
            } else {
                switch "foo bar" {
                case "\n\t":
                    v19 += (v19 % (((v19 / ((v19) | 1))) | 1))
                case "foo bar":
                    var v31 uint64 = 42
                    _ = v31
                }
                var v32 uint8 = (v1 << ((4294967295) & 63))
                _ = v32
            }
            switch "" {
            case "foo bar":
                v33 := GenS3{F0: v15}
                _ = v33
            case "\n\t":
                _ = len(v6)
            case "hello":
                var v34 bool = false
                _ = v34
            default:
                for fi35 := 0; fi35 < 7; fi35++ {
                    var v36 uint64 = 549
                    fmt.Println(v1, fi25)
                    var v37 string = ""
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            }
            _ = fi25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
