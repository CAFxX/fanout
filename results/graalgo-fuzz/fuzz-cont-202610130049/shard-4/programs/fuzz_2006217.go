package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int16 }

func main() {
    var v1 int32 = 0
    v2 := []uint32{4294967295, 107}
    var v3 int64 = 580
    var v4 float64 = (1.5 + 0.1)
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 bool = true
        v7 := GenS1{F0: 183, F1: "hello", F2: 602}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        v2 = append(v2, 1)
        v1 = v1
        var v9 uint32 = 0
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v1 += (v1 - 0)
        _ = fi10
    }
    v3 -= v3
    fmt.Println(v3)
    switch -551 {
    case 42:
        if (float64(-2.25) > (v4 + 0.0)) {
            v11 := GenS2{F0: (v1 / ((v1) | 1))}
            v3 -= (v3 * -9223372036854775808)
            _ = v11
        }
    case -796:
        var v12 int8 = -1
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        switch "foo bar" {
        case "":
            var v14 uint16 = 255
            _ = v14
        case "a":
            for fi15 := 0; fi15 < 3; fi15++ {
                switch fi13 {
                case 128:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    var v16 uint16 = 1
                    _ = v16
                }
                _ = fi15
            }
        default:
            var v17 bool = ((1e10 / 1.5) == float32(v4))
            _ = v17
        }
        v3 += (v3 >> ((1) & 63))
        _ = fi13
    }
    switch "foo bar" {
    case "hello":
        var v18 uint8 = 182
        _ = v18
    case "a":
        if len(v2) > 0 {
            v2[0] = 2
        }
    case "foo bar":
        if false {
            fmt.Println(v3, v3, v1)
            switch ((("foo bar" + ("a" + "foo bar")) + "foo bar") + "foo bar") {
            case "foo bar":
                switch "" {
                case "\n\t":
                    if (true && false) {
                        v3 = (v3 | v3)
                        v19 := GenS3{F0: 116, F1: 2147483647, F2: -114}
                        _ = v19
                    }
                case "a":
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                default:
                    fmt.Println(v4, v4, v4)
                }
            }
        } else {
            var v20 bool = true
            _ = v20
        }
    }
    switch "\n\t" {
    case "\n\t":
        if true {
            v1 -= v1
            for fi21 := 0; fi21 < 2; fi21++ {
                for fi22 := 0; fi22 < 16; fi22++ {
                    v3 = v3
                    var v23 uint64 = 4294967295
                    fmt.Println(v3, fi22)
                    _ = fi22
                    _ = v23
                }
                switch "" {
                case "":
                    fmt.Println(v4, v3, fi21)
                case "a":
                    if true {
                        var v24 uint32 = 2
                        v25 := GenS1{F0: 177, F1: "foo bar", F2: 163}
                        _ = v24
                        _ = v25
                    }
                }
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 uint8 = 127
                    v3 = v3
                    _ = fi26
                    _ = v27
                }
                _ = fi21
            }
            if true {
                v28 := GenS3{F0: 171, F1: 541, F2: 0}
                v3 += v3
                if (!false) {
                    fmt.Println(v1, v4, v1)
                    var v29 uint = 255
                    _ = v29
                }
                _ = v28
            } else {
                var v30 int16 = 746
                _ = v30
            }
        } else {
            v3 = 0
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
