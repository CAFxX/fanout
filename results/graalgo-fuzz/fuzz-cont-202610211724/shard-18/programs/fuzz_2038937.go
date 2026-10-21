package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int64; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint16; F2 string }

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]bool)
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]uint32)
    var v5 int32 = -2147483648
    var v6 bool = (true || (!(((!(!false)) && ((v5 | v5) >= v5)) && (!false))))
    switch -544 {
    case 32768:
        v5 -= 42
    case 42:
        v7 := []int{939, 42}
        _ = v7
    case -444:
        v1 = (v1 | 128)
    }
    v8 := []int{2147483647}
    for fi9 := 0; fi9 < 17; fi9++ {
        v1 = v1
        switch "foo bar" {
        case "\n\t":
            v5 = (v5 & (^(v5)))
        case "a":
            v4[1] = 706
        default:
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 int32 = 42
                var v12 uint64 = 18446744073709551615
                switch "foo bar" {
                case "\n\t":
                    var v13 bool = ((!v6) || (0 > 613))
                    _ = v13
                case "":
                    var v14 uint32 = 784
                    _ = v14
                case "a":
                    var v15 uint8 = 255
                    _ = v15
                default:
                    switch (("hello" + "a") + "\n\t") {
                    case "hello":
                        var v16 uint16 = uint16(v3)
                        _ = v16
                    case "foo bar":
                        var v17 float64 = 0.1
                        _ = v17
                    case "a":
                        v5 = 2147483647
                    }
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        switch (^(fi9)) {
        case 127:
            v18 := GenS1{F0: 128}
            _ = v18
        }
        _ = fi9
    }
    switch 128 {
    case 127:
        v19 := GenS2{F0: 0, F1: "hello"}
        _ = v19
    case 572:
        for fi20 := 0; fi20 < 16; fi20++ {
            if ("\n\t" > "a") {
                var v21 int8 = 1
                _ = v21
            } else {
                v22 := GenS1{F0: (v1 / ((v1) | 1))}
                _ = v22
            }
            _ = fi20
        }
    case 128:
        v5 = -748
    }
    v3 = 42
    v2[""] = v6
    var v23 int32 = v5
    v6 = v6
    if false {
        if (0 == 42) {
            v3 -= v3
        } else {
            if (v6 || true) {
                var v24 int32 = (-(v5))
                v8 = append(v8, 281)
                _ = v24
            }
            v25 := GenS3{F0: 9223372036854775807, F1: 102, F2: 2}
            v26 := GenS4{F0: (^(v1)), F1: 65535, F2: "foo bar"}
            _ = v25
            _ = v26
        }
    }
    v27 := v4[0]
    if (!((v3 >> ((4294967295) & 63)) != (^(v3)))) {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    } else {
        v5 += v5
        if v6 {
            var v28 uint16 = 42
            _ = v28
        } else {
            v1 = 1
        }
        v5 += 89
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
