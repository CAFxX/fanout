package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint32 }

func main() {
    var v1 bool = true
    var v2 int64 = 9223372036854775807
    var v3 float64 = (1.5 + 0.1)
    var v4 float32 = 1.5
    var v5 uint16 = 224
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{(^(fi6)), -1}
        var v8 uint = 255
        _ = fi6
        _ = v7
        _ = v8
    }
    if v1 {
        for fi9 := 0; fi9 < 7; fi9++ {
            v2 += (v2 | v2)
            _ = fi9
        }
        switch ("foo bar" + "\n\t") {
        case "":
            var v10 uint = 18446744073709551615
            _ = v10
        default:
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 string = "foo bar"
                var v13 int = (fi11 & (-(fi11)))
                for fi14 := 0; fi14 < 5; fi14++ {
                    v3 = (v3 - (((v3 - ((v3 / -2.25) / 1.5)) - (1.5 / 1.5)) * 1.5))
                    v2 = -256
                    _ = fi14
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    }
    fmt.Println(v1)
    v3 = v3
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint = 65535
        v17 := []int{42}
        v4 = v4
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        v19 := GenS1{F0: 231, F1: 2147483648, F2: -1}
        _ = fi18
        _ = v19
    }
    if true {
        if (!false) {
            v20 := []int{32767, 0}
            v3 = 1.5
            _ = v20
        }
        v21 := []int{54}
        v21 = append(v21, 695)
        _ = v21
    } else {
        if v1 {
            v2 = (v2 + 1)
            for fi22 := 0; fi22 < 3; fi22++ {
                v1 = v1
                _ = fi22
            }
            v23 := []int{32768, -9223372036854775808, 551, 4294967295}
            _ = v23
        }
        if v1 {
            fmt.Println(v3)
            v2 += v2
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 uint8 = 2
                fmt.Println(v25, v2)
                _ = fi24
                _ = v25
            }
        }
        fmt.Println(v1)
    }
    v26 := GenS1{F0: 0, F1: 4294967295, F2: -1}
    v27 := []int{-729, 32767, 998, 896}
    if len(v27) > 0 {
        v27[0] = 128
    }
    v4 += float32(-2.25)
    if ((49 != 42) && ((false && (v1 || v1)) && (v1 && v1))) {
        var v28 uint64 = 0
        v29 := GenS1{F0: 2147483647, F1: 2147483648, F2: -670}
        _ = v28
        _ = v29
    }
    switch 891 {
    case -688:
        for fi30 := 0; fi30 < 20; fi30++ {
            switch "hello" {
            case "foo bar":
                switch (fi30 << ((904) & 63)) {
                case 0:
                    for fi31 := 0; fi31 < 3; fi31++ {
                        var v32 float32 = (0.0 - v4)
                        v1 = (127 != 255)
                        _ = fi31
                        _ = v32
                    }
                case 42:
                    fmt.Println(v2, v1, v2)
                }
            }
            fmt.Println(v2, v3, v4)
            v27 = append(v27, (fi30 & (fi30 / ((fi30) | 1))))
            _ = fi30
        }
    }
    if v1 {
        v2 -= v2
        v33 := GenS2{F0: -1, F1: 421}
        var v34 int8 = -128
        _ = v33
        _ = v34
    } else {
        v35 := GenS3{F0: 100, F1: 42, F2: 1}
        v3 = v3
        _ = v35
    }
    var v36 uint32 = 750
    v2 = 9223372036854775807
    v5 += 333
    v1 = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
