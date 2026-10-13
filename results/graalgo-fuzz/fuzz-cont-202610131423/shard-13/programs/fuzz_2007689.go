package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int32; F1 uint8; F2 int16 }
type GenS5 struct { F0 bool; F1 bool }

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]uint32)
    v3 := []string{"", (("" + "") + ""), ((("\n\t" + ("" + "foo bar")) + "a") + "hello"), "foo bar"}
    var v4 uint16 = 636
    var v5 float64 = -2.25
    for fi6 := 0; fi6 < 19; fi6++ {
        delete(v2, 2)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v7 uint16 = v4
        _ = fi6
        _ = v7
    }
    _ = len(v2)
    fmt.Println(v1)
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 -= v1
    }
    v4 -= v4
    switch 32767 {
    case 2147483647:
        switch "" {
        case "hello":
            v5 += 3.14159
        default:
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 uint32 = 2147483647
                switch fi8 {
                case -2147483648:
                    v10 := GenS1{F0: -67}
                    _ = v10
                default:
                    fmt.Println(fi8, fi8, v9)
                }
                var v11 uint16 = v4
                _ = fi8
                _ = v11
                _ = v9
            }
        }
    default:
        _ = len(v3)
    }
    v4 = (v4 / ((309) | 1))
    if ("a" > "hello") {
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 6; fi13++ {
                if ((fi13 % ((fi13) | 1)) == fi12) {
                    v14 := GenS2{F0: 706, F1: 169}
                    if (!false) {
                        v1 -= (v1 - (-(v1)))
                    }
                    var v15 uint16 = 898
                    _ = v14
                    _ = v15
                } else {
                    v16 := GenS2{F0: 0, F1: 42}
                    var v17 int16 = 0
                    fmt.Println(fi13, v5, v17)
                    var v18 uint32 = 400
                    v5 = ((0.1 + v5) + v5)
                    v19 := GenS1{F0: 32767}
                    _ = v16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = fi13
            }
            switch (fi12 >> ((801) & 63)) {
            case 0:
                v1 = 9223372036854775807
            }
            v1 = (^(v1))
            _ = fi12
        }
        switch -1 {
        case -9223372036854775808:
            switch (("a" + ("foo bar" + (("\n\t" + "foo bar") + (((("" + "") + "\n\t") + "foo bar") + ("foo bar" + ""))))) + (((("foo bar" + "\n\t") + "hello") + ("a" + "foo bar")) + "a")) {
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    } else {
        v5 = v5
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 string = "\n\t"
            _ = fi20
            _ = v21
        }
    }
    v5 = 0.1
    v4 = (v4 ^ (v4 - (v4 << ((1) & 63))))
    if (0 > -1) {
        v22 := GenS1{F0: 0}
        if (true && (true || (-2.25 > float32((((v5 - float64(1.5)) / 1.5) / (v5 + v5)))))) {
            if true {
                var v23 uint64 = 399
                for fi24 := 0; fi24 < 15; fi24++ {
                    v25 := GenS3{F0: v1, F1: 60}
                    _ = fi24
                    _ = v25
                }
                v26 := v2[2147483647]
                _ = v23
                _ = v26
            }
            v27 := GenS4{F0: -1, F1: 10, F2: 1000}
            _ = v27
        } else {
            if (!(!(118 >= 98))) {
                v28 := GenS5{F0: (false && (true && (65535 > 2147483647))), F1: ((v4 >> ((uint(v4)) & 63)) != (^(v4)))}
                _ = v28
            }
            v4 = v4
        }
        _ = v22
    }
    v4 -= v4
    v4 += v4
    _ = len(v2)
    var v29 uint64 = 0
    var v30 bool = false
    var v31 bool = (-1 >= 0)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
