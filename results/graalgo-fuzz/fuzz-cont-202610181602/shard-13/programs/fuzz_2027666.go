package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 string }

func main() {
    var v1 int32 = 0
    var v2 float64 = 1.5
    var v3 uint = 65535
    var v4 int32 = (v1 >> ((v3) & 63))
    var v5 bool = true
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 15; fi7++ {
        v3 += v3
        switch (("\n\t" + "hello") + "hello") {
        case "\n\t":
            var v8 float64 = ((v2 * float64((-2.25 + float32(-2.25)))) + (v2 / v2))
            _ = v8
        case "hello":
            v6["foo bar"] = 540
        case "foo bar":
            if true {
                if (true && v5) {
                    for fi9 := 0; fi9 < 13; fi9++ {
                        fmt.Println(v5, v1, v3)
                        _ = fi9
                    }
                } else {
                    v4 -= -202
                }
            } else {
                v2 = ((v2 * (float64(3.14159) / (v2 - v2))) * v2)
            }
        default:
            switch fi7 {
            case 42:
                var v10 uint16 = 65535
                _ = v10
            case 128:
                v3 += v3
            }
        }
        _ = fi7
    }
    var v11 uint = v3
    v4 -= -1
    if (2147483648 < uint32(v3)) {
        v12 := GenS1{F0: -9223372036854775808, F1: 42}
        switch 127 {
        case 462:
            v3 -= (v11 + v11)
        default:
            var v13 uint16 = 42
            _ = v13
        }
        _ = v12
    }
    v2 = 1e10
    var v14 uint32 = 2
    for fi15 := 0; fi15 < 10; fi15++ {
        if (!(v14 != 1)) {
            v14 = 680
        }
        if ((^(v11)) < v11) {
            v2 = (v2 + 0.0)
            switch (fi15 / ((fi15) | 1)) {
            case -9223372036854775808:
                if (2 != (v14 ^ (v14 | v14))) {
                    v5 = (true && (!(1 >= 1)))
                    var v16 uint8 = 1
                    _ = v16
                } else {
                    var v17 uint16 = 42
                    _ = v17
                }
            }
            var v18 float64 = (v2 / v2)
            _ = v18
        }
        fmt.Println(v1, v2, v4)
        _ = fi15
    }
    fmt.Println(v11, v4, v5)
    v6[""] = (v14 / (((^(v14))) | 1))
    v19 := GenS2{F0: 42}
    if (!true) {
        if (!v5) {
            v20 := GenS3{F0: 1, F1: 716}
            _ = v20
        } else {
            if (!v5) {
                var v21 bool = (36 < -1)
                _ = v21
            }
            var v22 int32 = -269
            _ = v22
        }
    } else {
        var v23 int = -1
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 float32 = 0.1
            v3 = (v11 << (((^(v3))) & 63))
            v26 := GenS1{F0: -1, F1: v14}
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 17; fi27++ {
            v14 = 653
            var v28 uint = v3
            _ = fi27
            _ = v28
        }
        _ = v23
    }
    fmt.Println(v5, v1, v14)
    v29 := GenS4{F0: "foo bar"}
    var v30 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v14
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
