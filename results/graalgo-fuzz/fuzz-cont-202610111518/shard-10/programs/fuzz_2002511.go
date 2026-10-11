package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 string }

func main() {
    v1 := []int{42, -9223372036854775808, -730, -2147483648}
    v2 := []string{"a", "\n\t", "hello"}
    var v3 bool = (-2147483648 != 910)
    var v4 uint8 = 96
    switch ("a" + "hello") {
    case "hello":
        v5 := GenS1{F0: 48, F1: 127, F2: "foo bar"}
        _ = v5
    case "\n\t":
        for fi6 := 0; fi6 < 4; fi6++ {
            fmt.Println(fi6, v3, fi6)
            if (fi6 == 32768) {
                v3 = v3
                v3 = v3
                var v7 int32 = 2147483647
                _ = v7
            } else {
                if (true || (46 <= -96)) {
                    fmt.Println(v4, fi6)
                }
            }
            _ = fi6
        }
    default:
        if (((1e10 + 1e10) * 1.5) <= 1e10) {
            var v8 uint8 = v4
            _ = v8
        }
    }
    v3 = (((false && v3) || v3) && true)
    fmt.Println(v4)
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint64 = 2
        _ = fi9
        _ = v10
    }
    var v11 string = "foo bar"
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 float32 = 1.5
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        fmt.Println(v11, v4, v3)
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 2, F1: 24, F2: "foo bar"}
    v15 := GenS1{F0: 2, F1: -1, F2: (v11 + v11)}
    switch -2147483648 {
    case 0:
        var v16 uint = 2
        _ = v16
    case 9223372036854775807:
        v17 := GenS1{F0: 578, F1: 103, F2: ("foo bar" + (v11 + "a"))}
        _ = v17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 int = (fi18 << ((255) & 63))
        if v3 {
            v11 = v11
        }
        if v3 {
            v4 += v4
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch (v19 & 2147483647) {
                case -1:
                    var v20 uint16 = 255
                    _ = v20
                case 127:
                    if (v11 == (v11 + v11)) {
                        var v21 int16 = 1
                        v11 = v11
                        var v22 uint = 18446744073709551615
                        _ = v21
                        _ = v22
                    }
                case -393:
                    if v3 {
                        fmt.Println(v19, v4, v11)
                        v19 = fi18
                    }
                }
                v3 = v3
            }
            for fi23 := 0; fi23 < 6; fi23++ {
                if (!false) {
                    v11 = ("\n\t" + ("foo bar" + (((("foo bar" + v11) + v11) + "foo bar") + "hello")))
                    fmt.Println(v11, v19)
                    v24 := GenS1{F0: uint16(v4), F1: -128, F2: ((v11 + v11) + ("foo bar" + v11))}
                    _ = v24
                } else {
                    if (784 <= -1) {
                        v25 := GenS1{F0: 255, F1: 42, F2: (v11 + (v11 + ("\n\t" + v11)))}
                        v4 = (v4 + 255)
                        v19 = int(v4)
                        v19 -= 2147483647
                        _ = v25
                    }
                }
                var v26 int16 = 0
                _ = fi23
                _ = v26
            }
        }
        _ = fi18
        _ = v19
    }
    var v27 int = -9223372036854775808
    _ = len(v2)
    var v28 int8 = 100
    if v3 {
        var v29 uint32 = 786
        var v30 int64 = -343
        v29 += (^(v29))
        _ = v29
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
