package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int64 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []uint32{}
    v2 := make(map[int]bool)
    var v3 int32 = -1
    var v4 int32 = 1
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            if false {
                var v7 float64 = 1.5
                _ = len(v2)
                if (4294967295 <= 42) {
                    var v8 int16 = 42
                    v4 -= (v3 << ((65535) & 63))
                    _ = v8
                }
                _ = v7
            }
            v3 = 121
            if (((0.1 / ((1e10 - 0.1) + 1.5)) <= (1.5 * 0.1)) || ((fi6 ^ 503) != fi5)) {
                v3 = v3
                v3 = v4
            }
            _ = fi6
        }
        for fi9 := 0; fi9 < 16; fi9++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v10 := GenS1{F0: (v3 >> ((255) & 63)), F1: 38, F2: -1}
            switch ("foo bar" + "\n\t") {
            case "":
                var v11 int8 = 127
                _ = v11
            case "\n\t":
                var v12 uint = 821
                _ = v12
            case "hello":
                if true {
                    var v13 int64 = 1
                    v4 -= (v4 ^ 907)
                    if ((-128 != 78) && ((^(v4)) < 2147483647)) {
                        var v14 float32 = 3.14159
                        _ = v14
                    } else {
                        var v15 uint = 255
                        v4 += 247
                        _ = v15
                    }
                    fmt.Println(fi5, v4, v3)
                    var v16 string = "foo bar"
                    _ = v13
                    _ = v16
                }
            }
            _ = fi9
            _ = v10
        }
        _ = fi5
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            switch "\n\t" {
            case "a":
                v3 += (-(v4))
            case "foo bar":
                fmt.Println(v3, fi17, v4)
            }
            _ = fi18
        }
        var v19 int16 = 1
        _ = fi17
        _ = v19
    }
    v20 := GenS1{F0: v3, F1: -1, F2: -1}
    delete(v2, -9223372036854775808)
    if false {
        var v21 uint16 = 31
        fmt.Println(v21, v3)
        for fi22 := 0; fi22 < 19; fi22++ {
            v3 += 42
            _ = fi22
        }
        _ = v21
    } else {
        v3 = -1
    }
    if (!(-9223372036854775808 != -9223372036854775808)) {
        var v23 int32 = (v3 + (-(v4)))
        v24 := GenS2{F0: 2}
        _ = v23
        _ = v24
    } else {
        v4 += v4
        for fi25 := 0; fi25 < 6; fi25++ {
            v26 := GenS2{F0: 2}
            v3 = v4
            _ = fi25
            _ = v26
        }
    }
    if (-32768 == 42) {
        switch "" {
        case "":
            var v27 uint64 = 239
            _ = v27
        case "\n\t":
            for fi28 := 0; fi28 < 12; fi28++ {
                if (255 >= 42) {
                    var v29 int8 = 42
                    v3 = v3
                    _ = v29
                }
                _ = fi28
            }
        default:
            var v30 int32 = v3
            _ = v30
        }
        _ = len(v2)
        if len(v1) > 0 {
            v1[0] = 754
        }
    } else {
        for fi31 := 0; fi31 < 13; fi31++ {
            v3 = v4
            for fi32 := 0; fi32 < 5; fi32++ {
                v3 += -1
                fmt.Println(fi32)
                _ = fi32
            }
            _ = fi31
        }
    }
    v4 += (v4 & -2147483648)
    var v33 int16 = -765
    var v34 uint = 42
    var v35 int32 = v3
    var v36 int16 = v33
    fmt.Println(v36, v34, v36)
    v36 = 1
    for fi37 := 0; fi37 < 3; fi37++ {
        v2[fi37] = (!(((1.5 * float64(1e10)) >= ((3.14159 - (1.5 / 3.14159)) / -2.25)) || false))
        v35 -= v4
        var v38 int16 = v36
        _ = fi37
        _ = v38
    }
    if true {
        var v39 uint8 = 255
        v34 = 392
        var v40 uint16 = 2
        _ = v39
        _ = v40
    }
    var v41 int32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
}
