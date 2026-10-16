package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint64; F2 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int64; F2 int16 }

func main() {
    var v1 uint32 = 1
    v2 := make(map[string]int)
    var v3 uint16 = 42
    var v4 int16 = -1
    var v5 int16 = v4
    v6 := v2[("\n\t" + (("a" + "") + ("hello" + "")))]
    v7 := GenS1{F0: -1}
    v8 := []int{v6, (-(v6)), v6, 32767}
    if false {
        fmt.Println(v6)
        var v9 uint32 = (v1 / ((2147483648) | 1))
        var v10 uint64 = uint64(v1)
        _ = v10
        _ = v9
    }
    if true {
        if true {
            v11 := GenS2{F0: (("" + "foo bar") + "a"), F1: 337, F2: (^(v6))}
            switch ((("\n\t" + "") + "foo bar") + ("" + "foo bar")) {
            case "\n\t":
                v3 -= v3
            }
            fmt.Println(v4, v1)
            _ = v11
        } else {
            var v12 uint = 81
            var v13 int8 = 1
            for fi14 := 0; fi14 < 20; fi14++ {
                for fi15 := 0; fi15 < 4; fi15++ {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    v16 := GenS2{F0: (((("foo bar" + "foo bar") + (("" + "foo bar") + "foo bar")) + "foo bar") + ("hello" + "\n\t")), F1: 4294967295, F2: fi15}
                    _ = fi15
                    _ = v16
                }
                switch "hello" {
                case "\n\t":
                    var v17 int8 = 42
                    _ = v17
                case "hello":
                    v18 := GenS1{F0: -9223372036854775808}
                    _ = v18
                default:
                    v19 := GenS2{F0: (("foo bar" + "hello") + "a"), F1: 42, F2: 9223372036854775807}
                    _ = v19
                }
                v3 += v3
                _ = fi14
            }
            _ = v12
            _ = v13
        }
        v20 := GenS2{F0: "", F1: 1, F2: (v6 ^ v6)}
        _ = v20
    } else {
        v21 := GenS2{F0: ("hello" + "hello"), F1: 699, F2: (v6 - v6)}
        _ = v21
    }
    v5 = v5
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 int = 42
        _ = fi22
        _ = v23
    }
    var v24 bool = (!false)
    for fi25 := 0; fi25 < 10; fi25++ {
        for fi26 := 0; fi26 < 3; fi26++ {
            var v27 float64 = (0.0 / (((((1.5 + (0.1 - ((-2.25 * 1e10) / (float64(float32(1e10)) / -2.25)))) / 1.0) * (0.1 - 0.0)) * -2.25) + float64(0.0)))
            var v28 bool = (v24 || v24)
            v28 = (175 != 210)
            _ = fi26
            _ = v27
            _ = v28
        }
        switch "a" {
        case "a":
            v29 := GenS3{F0: 787}
            _ = v29
        case "hello":
            for fi30 := 0; fi30 < 18; fi30++ {
                switch (fi25 & (fi30 & fi25)) {
                case 32767:
                    v4 = -32768
                }
                if ((1 != 1) && (0.0 <= 0.0)) {
                    var v31 string = "hello"
                    var v32 int32 = -1
                    v33 := GenS4{F0: int(v4), F1: 459, F2: v5}
                    _ = v31
                    _ = v32
                    _ = v33
                } else {
                    fmt.Println(v3, fi30)
                    if (v24 || (!v24)) {
                        v5 -= (v4 / ((-32768) | 1))
                    }
                }
                _ = fi30
            }
        case "foo bar":
            v34 := v2["a"]
            _ = v34
        }
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 uint8 = 0
            v24 = v24
            var v37 int16 = (v4 & (v4 * -1))
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi25
    }
    var v38 uint16 = 1
    var v39 bool = (("foo bar" + "foo bar") > "a")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
