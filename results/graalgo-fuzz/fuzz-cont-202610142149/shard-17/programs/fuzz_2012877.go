package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint64; F2 bool }
type GenS5 struct { F0 int64; F1 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 143
    var v4 string = "hello"
    for fi5 := 0; fi5 < 16; fi5++ {
        switch (fi5 * 1) {
        case -1:
            v6 := GenS1{F0: 0, F1: 2, F2: 127}
            _ = v6
        }
        v7 := GenS1{F0: -972, F1: 2, F2: 42}
        v8 := GenS2{F0: 235, F1: 255, F2: (v2 * (v2 >> ((4294967295) & 63)))}
        _ = fi5
        _ = v7
        _ = v8
    }
    v4 = v4
    v4 = ("foo bar" + v4)
    if (109 <= 65535) {
        fmt.Println(v4, v2)
        switch v4 {
        case "hello":
            switch v4 {
            case "":
                switch -264 {
                case 9223372036854775807:
                    v1[2147483647] = 166
                default:
                    v3 -= (v3 & v3)
                }
            default:
                v1[1] = 0
            }
        case "foo bar":
            v9 := GenS3{F0: 0}
            _ = v9
        case "a":
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 float32 = -2.25
                v12 := GenS2{F0: 42, F1: 255, F2: v2}
                v4 = "hello"
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    }
    v13 := GenS3{F0: 731}
    v2 += v2
    for fi14 := 0; fi14 < 1; fi14++ {
        fmt.Println(fi14, v2)
        switch "hello" {
        case "foo bar":
            v15 := GenS4{F0: (v4 + ((((v4 + "a") + ("" + v4)) + v4) + (v4 + v4))), F1: (v3 ^ (v3 ^ v3)), F2: (false || false)}
            _ = v15
        }
        if true {
            if false {
                v16 := GenS1{F0: 1, F1: 886, F2: 1}
                _ = v16
            } else {
                if false {
                    switch 128 {
                    case -1:
                        if ((v4 > v4) && false) {
                            v4 = (("" + (((v4 + v4) + "foo bar") + v4)) + "hello")
                        } else {
                            v3 += (v3 & v3)
                        }
                    case -878:
                        var v17 bool = ((v3 ^ 460) >= (v3 << ((uint(v3)) & 63)))
                        _ = v17
                    case 128:
                        var v18 string = v4
                        _ = v18
                    default:
                        v19 := GenS5{F0: -230, F1: 0}
                        _ = v19
                    }
                    v3 -= (^(v3))
                    var v20 int = fi14
                    _ = v20
                }
            }
        }
        _ = fi14
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        v1[0] = 2147483647
        _ = len(v1)
        for fi22 := 0; fi22 < 4; fi22++ {
            if (!true) {
                v2 = 1
                var v23 uint8 = 255
                _ = v23
            }
            var v24 string = v4
            var v25 uint = 4294967295
            fmt.Println(v3, fi21, v3)
            _ = fi22
            _ = v24
            _ = v25
        }
        _ = fi21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
