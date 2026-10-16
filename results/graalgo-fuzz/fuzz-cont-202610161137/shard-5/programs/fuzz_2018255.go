package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int }

func main() {
    v1 := []string{"foo bar", "foo bar", "hello", "foo bar"}
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    var v4 uint = 18446744073709551615
    var v5 int64 = -1
    var v6 int = 128
    for fi7 := 0; fi7 < 2; fi7++ {
        switch "hello" {
        case "foo bar":
            for fi8 := 0; fi8 < 13; fi8++ {
                switch (("hello" + "hello") + "foo bar") {
                case "foo bar":
                    var v9 int = (fi7 & -731)
                    _ = v9
                default:
                    var v10 int8 = -128
                    _ = v10
                }
                _ = fi8
            }
        case "hello":
            v5 -= v5
        }
        v11 := GenS1{F0: 1}
        var v12 int32 = 1
        _ = fi7
        _ = v11
        _ = v12
    }
    switch ("" + "a") {
    case "foo bar":
        if true {
            v6 += (v6 | -9223372036854775808)
        } else {
            var v13 uint64 = 42
            var v14 bool = false
            switch "hello" {
            case "foo bar":
                _ = len(v2)
            }
            _ = v13
            _ = v14
        }
    case "":
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 int64 = -584
            var v17 uint32 = 4294967295
            if (((^(v4)) > 625) || false) {
                fmt.Println(v5, v4)
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v6 -= (v6 | (v6 << ((0) & 63)))
    }
    v18 := GenS1{F0: 2}
    for fi19 := 0; fi19 < 3; fi19++ {
        fmt.Println(v5)
        var v20 int8 = 1
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: 2}
    fmt.Println(v6, v4, v4)
    switch "\n\t" {
    case "\n\t":
        v4 -= v4
    }
    if true {
        v22 := GenS1{F0: 0}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 6; fi23++ {
            switch "" {
            case "foo bar":
                v2[2147483648] = fi23
            }
            v24 := GenS2{F0: 1, F1: (v6 | (fi23 >> ((v4) & 63)))}
            _ = fi23
            _ = v24
        }
        v2[1] = v6
    }
    var v25 uint = (^(v4))
    v26 := GenS2{F0: 2147483647, F1: (v6 * 1)}
    v5 += v5
    for fi27 := 0; fi27 < 18; fi27++ {
        fmt.Println(v5, v25, fi27)
        _ = fi27
    }
    if true {
        var v28 float64 = -2.25
        switch "a" {
        case "foo bar":
            if true {
                var v29 int8 = int8(v25)
                if (18446744073709551615 <= 18446744073709551615) {
                    for fi30 := 0; fi30 < 11; fi30++ {
                        var v31 uint16 = 523
                        var v32 bool = false
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    v6 -= 32767
                }
                _ = v29
            }
        case "a":
            var v33 int64 = v5
            _ = v33
        case "\n\t":
            for fi34 := 0; fi34 < 12; fi34++ {
                v35 := v2[2]
                _ = fi34
                _ = v35
            }
        default:
            switch -2147483648 {
            case 4294967295:
                v36 := GenS1{F0: 81}
                _ = v36
            case 591:
                if ((v28 + v28) != (0.0 / -2.25)) {
                    v25 = (^(v4))
                    switch ("\n\t" + "a") {
                    case "hello":
                        var v37 int32 = int32(v6)
                        _ = v37
                    }
                }
            default:
                var v38 float64 = 0.1
                _ = v38
            }
        }
        _ = v28
    }
    fmt.Println(v25, v25, v5)
    if (65535 <= 523) {
        switch (v6 >> ((391) & 63)) {
        case 32768:
            if (!(238 != 224)) {
                for fi39 := 0; fi39 < 5; fi39++ {
                    v5 = 42
                    fmt.Println(v5, v25)
                    _ = len(v1)
                    _ = fi39
                }
                var v40 int = v6
                _ = v40
            }
        case 2147483647:
            v4 += (v25 << (((v4 + 18446744073709551615)) & 63))
        case 1:
            var v41 uint64 = 42
            _ = v41
        default:
            v1 = append(v1, "a")
        }
        if true {
            if (0.1 >= (0.0 * -2.25)) {
                v4 += v4
            }
            fmt.Println(v5, v6, v25)
            var v42 uint32 = 2147483647
            _ = v42
        } else {
            v4 = v4
            v5 -= (-(v5))
            fmt.Println(v4)
        }
    }
    v2[2147483647] = (v6 % ((v6) | 1))
    v1 = append(v1, "")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
