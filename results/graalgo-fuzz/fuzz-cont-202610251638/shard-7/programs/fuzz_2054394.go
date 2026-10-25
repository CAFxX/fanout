package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 uint16; F1 int; F2 int16 }
func genFn0(p0 uint64, p1 int8) int16 {
    v6 := []int{4294967295, 523, 9223372036854775807, 32767}
    p0 = (^(p0))
    if len(v6) > 0 {
        v6[0] = 0
    }
    switch 1 {
    case 42:
        var v7 int16 = 0
        _ = v7
    }
    var v8 uint64 = (p0 ^ (p0 - 18446744073709551615))
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 553
}
func genFn1(p0 int64) int8 {
    var v9 uint16 = 678
    v9 = (^(v9))
    if false {
        v10 := []int{int(v9), 128}
        _ = len(v10)
        switch ("" + (((("" + "a") + ("" + ("" + "hello"))) + "a") + "\n\t")) {
        case "hello":
            v11 := GenS1{F0: 332, F1: (p0 << ((65535) & 63))}
            _ = v11
        case "foo bar":
            for fi12 := 0; fi12 < 8; fi12++ {
                p0 += p0
                p0 -= 9223372036854775807
                var v13 uint8 = 83
                _ = fi12
                _ = v13
            }
        case "":
            if (0 != 0) {
                fmt.Println(p0)
            } else {
                v14 := GenS2{F0: v9, F1: 32767}
                switch (("a" + (("\n\t" + "hello") + ((("\n\t" + "hello") + "\n\t") + ("foo bar" + "\n\t")))) + ("foo bar" + "hello")) {
                case "\n\t":
                    v15 := GenS1{F0: 2, F1: p0}
                    _ = v15
                case "hello":
                    var v16 uint = 0
                    _ = v16
                case "foo bar":
                    var v17 int8 = 127
                    _ = v17
                }
                _ = v14
            }
        }
        _ = v10
    }
    p0 -= (p0 | (p0 & p0))
    var v18 bool = (824 >= uint64(v9))
    _ = p0
    _ = v18
    _ = v9
    return 26
}

func main() {
    v1 := []string{"", ("foo bar" + "hello"), "\n\t", ""}
    v2 := make(map[string]uint32)
    var v3 float64 = 0.0
    var v4 uint = 2
    var v5 int8 = 1
    var v19 uint = (v4 >> ((669) & 63))
    if (!false) {
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 uint64 = 133
            v22 := GenS3{F0: (v21 + v21), F1: 1000}
            _ = len(v1)
            _ = fi20
            _ = v21
            _ = v22
        }
        v5 += 42
        v3 = float64(0.0)
    }
    fmt.Println(v5)
    v23 := GenS1{F0: 42, F1: -1}
    var v24 int32 = 0
    for fi25 := 0; fi25 < 7; fi25++ {
        v19 -= 431
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 float64 = ((1e10 + 0.1) * float64(((0.1 / 1.5) - (0.1 / ((((3.14159 + 1e10) - 1e10) * 1.5) / (3.14159 + 3.14159))))))
            fmt.Println(fi26, v3, v3)
            switch "\n\t" {
            case "\n\t":
                var v28 float32 = 1e10
                _ = v28
            case "foo bar":
                var v29 int8 = (v5 ^ v5)
                _ = v29
            }
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    if (false && true) {
        switch 32767 {
        case -1:
            v1 = append(v1, (("hello" + "hello") + "\n\t"))
        }
    }
    if (("" + (("hello" + "a") + ("" + ("\n\t" + "a")))) < "hello") {
        var v30 int32 = v24
        var v31 uint32 = 4294967295
        v32 := GenS3{F0: 1, F1: 1}
        _ = v30
        _ = v31
        _ = v32
    }
    v33 := GenS1{F0: 42, F1: 9223372036854775807}
    var v34 float32 = 1e10
    var v35 int64 = 814
    v1 = append(v1, ("a" + ""))
    v4 = v4
    v36 := GenS4{F0: v5}
    v37 := GenS5{F0: 993, F1: 2147483648}
    for fi38 := 0; fi38 < 15; fi38++ {
        for fi39 := 0; fi39 < 18; fi39++ {
            v40 := GenS6{F0: 294, F1: 587, F2: 0}
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
