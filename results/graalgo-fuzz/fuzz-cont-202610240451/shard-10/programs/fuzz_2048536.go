package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint32; F2 string }
type GenS3 struct { F0 int16; F1 int32; F2 int }

func main() {
    v1 := []string{}
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    var v4 float32 = float32(1e10)
    var v5 int64 = -9223372036854775808
    v6 := []int{-1, 532, 1, -1}
    if (!((42 == 255) || true)) {
        if false {
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := GenS1{F0: (1 >= 174)}
                var v9 float32 = v4
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        var v10 float32 = 0.1
        switch (("" + ("\n\t" + "\n\t")) + ((("foo bar" + "") + "a") + "\n\t")) {
        case "hello":
            if false {
                if (false || true) {
                    v11 := GenS2{F0: 1000, F1: 746, F2: "foo bar"}
                    var v12 int8 = 127
                    _ = v11
                    _ = v12
                }
                v10 = v4
            }
        case "":
            v13 := GenS2{F0: 42, F1: 42, F2: "a"}
            _ = v13
        default:
            if false {
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 int64 = v5
                    _ = fi14
                    _ = v15
                }
                switch "" {
                case "\n\t":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                if (2 != 127) {
                    v16 := GenS2{F0: -926, F1: 1, F2: "\n\t"}
                    _ = v16
                } else {
                    var v17 uint32 = 346
                    _ = v17
                }
            }
        }
        _ = v10
    }
    var v18 uint64 = 18446744073709551615
    v19 := GenS1{F0: true}
    v18 += (v18 | (v18 - 18446744073709551615))
    for fi20 := 0; fi20 < 2; fi20++ {
        fmt.Println(v4)
        v4 = v4
        _ = fi20
    }
    fmt.Println(v4)
    v4 = v4
    var v21 uint8 = 42
    for fi22 := 0; fi22 < 1; fi22++ {
        if true {
            var v23 uint16 = 42
            if ((v4 - v4) != (v4 - 0.0)) {
                if ("" <= "") {
                    var v24 float64 = 0.1
                    var v25 uint16 = 42
                    var v26 int64 = -746
                    _ = v24
                    _ = v25
                    _ = v26
                } else {
                    var v27 float64 = float64(0.1)
                    var v28 int64 = (v5 % (((v5 / ((0) | 1))) | 1))
                    _ = v27
                    _ = v28
                }
                v5 = v5
            }
            _ = len(v1)
            _ = v23
        }
        v5 -= (v5 / ((1) | 1))
        _ = fi22
    }
    v1 = append(v1, (("hello" + "\n\t") + "hello"))
    v5 = v5
    switch 42 {
    case 32768:
        if false {
            if false {
                switch -2147483648 {
                case 2147483647:
                    var v29 uint8 = v21
                    _ = v29
                }
            } else {
                v30 := GenS2{F0: 0, F1: 65535, F2: "\n\t"}
                var v31 int32 = 2147483647
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 17; fi32++ {
                if len(v6) > 0 {
                    v6[0] = (fi32 + fi32)
                }
                v33 := v2[("\n\t" + "a")]
                v34 := GenS3{F0: -32768, F1: 42, F2: (-(fi32))}
                _ = fi32
                _ = v33
                _ = v34
            }
            v35 := GenS3{F0: -32768, F1: -2147483648, F2: 89}
            _ = v35
        } else {
            v4 -= v4
        }
    case 979:
        fmt.Println(v21, v21)
    case 1:
        if true {
            switch 1 {
            case -28:
                v36 := GenS2{F0: 0, F1: 0, F2: ("a" + ("" + ("a" + "\n\t")))}
                _ = v36
            case 1:
                v21 = 255
            default:
                var v37 string = ("" + ("\n\t" + (("" + "hello") + "hello")))
                _ = v37
            }
        } else {
            var v38 string = ((("hello" + (((("a" + ("" + ("\n\t" + "a"))) + "foo bar") + "hello") + "hello")) + ("foo bar" + "foo bar")) + "hello")
            _ = v38
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
