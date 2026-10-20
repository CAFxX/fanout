package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 int; F2 int64 }

func main() {
    v1 := []uint32{793, 574}
    var v2 int16 = -700
    var v3 uint16 = 224
    var v4 int32 = 0
    if false {
        v4 -= (v4 & 2147483647)
        v5 := GenS1{F0: 42}
        _ = v5
    }
    var v6 int = 42
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!(!true)) {
        switch ((("" + "a") + "hello") + "a") {
        case "":
            for fi7 := 0; fi7 < 18; fi7++ {
                for fi8 := 0; fi8 < 13; fi8++ {
                    v1 = append(v1, 65535)
                    if false {
                        var v9 uint8 = 128
                        _ = v9
                    }
                    _ = fi8
                }
                _ = fi7
            }
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "\n\t":
            var v10 string = ""
            _ = v10
        default:
            if true {
                for fi11 := 0; fi11 < 16; fi11++ {
                    for fi12 := 0; fi12 < 17; fi12++ {
                        v13 := GenS1{F0: 0}
                        v4 -= v4
                        _ = fi12
                        _ = v13
                    }
                    _ = fi11
                }
            }
        }
    } else {
        var v14 uint8 = 2
        v2 -= (v2 % ((v2) | 1))
        _ = v14
    }
    v3 = (v3 | v3)
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS1{F0: v4}
        v2 = -1
        if (((4 <= 0) || true) && true) {
            v1 = append(v1, 2147483648)
        }
        _ = fi15
        _ = v16
    }
    v3 = v3
    v1 = append(v1, 2147483648)
    if ((v4 >= -1) || false) {
        v3 -= v3
        if (42 < 0) {
            v3 = 2
        } else {
            var v17 uint = 18446744073709551615
            v3 -= v3
            _ = v17
        }
        var v18 uint16 = v3
        _ = v18
    } else {
        var v19 uint = 2
        var v20 int32 = v4
        _ = v19
        _ = v20
    }
    var v21 string = "hello"
    for fi22 := 0; fi22 < 7; fi22++ {
        v6 = 590
        switch v6 {
        case 1:
            v23 := GenS2{F0: (v4 ^ (^(v4)))}
            _ = v23
        }
        fmt.Println(v4, v21)
        _ = fi22
    }
    v3 = 0
    v4 = (v4 << ((1) & 63))
    for fi24 := 0; fi24 < 15; fi24++ {
        for fi25 := 0; fi25 < 5; fi25++ {
            v3 = 255
            var v26 int8 = 87
            switch v21 {
            case "\n\t":
                var v27 int16 = 0
                _ = v27
            default:
                var v28 int64 = -1
                _ = v28
            }
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    if true {
        var v29 int16 = (v2 ^ (v2 & -277))
        v3 += v3
        _ = v29
    } else {
        switch (^(v6)) {
        case 0:
            if true {
                fmt.Println(v2, v4)
                if (!true) {
                    switch (v21 + v21) {
                    case "hello":
                        fmt.Println(v2)
                    case "foo bar":
                        v4 = (v4 << ((212) & 63))
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v2 = int16(v3)
                    for fi30 := 0; fi30 < 11; fi30++ {
                        v1 = append(v1, 153)
                        _ = fi30
                    }
                }
                var v31 uint32 = 1
                _ = v31
            } else {
                for fi32 := 0; fi32 < 14; fi32++ {
                    switch fi32 {
                    case -734:
                        fmt.Println(v4, fi32, v6)
                    default:
                        v6 = (fi32 % ((fi32) | 1))
                    }
                    _ = fi32
                }
                var v33 int64 = 42
                _ = v33
            }
        case 1:
            v34 := GenS3{F0: 42}
            _ = v34
        default:
            switch (v6 / (((v6 | (v6 + 9223372036854775807))) | 1)) {
            case 499:
                v35 := GenS3{F0: 42}
                _ = v35
            case 1:
                v36 := GenS4{F0: 30, F1: (v6 ^ 0), F2: -751}
                _ = v36
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v6
}
