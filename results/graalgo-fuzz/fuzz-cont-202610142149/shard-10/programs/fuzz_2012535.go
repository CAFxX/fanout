package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint16 }
type GenS2 struct { F0 string; F1 int32; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint; F2 int }
type GenS4 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 int16; F2 uint16 }
func genFn0() int {
    switch 0 {
    case -37:
        for fi7 := 0; fi7 < 20; fi7++ {
            fmt.Println(fi7)
            var v8 int = fi7
            if false {
                fmt.Println(fi7, v8)
                v8 = (v8 * 32767)
                v9 := GenS1{F0: 0, F1: -128, F2: 255}
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    case -2147483648:
        var v10 uint32 = 2147483648
        _ = v10
    default:
        v11 := GenS2{F0: "a", F1: 2147483647, F2: 2147483647}
        _ = v11
    }
    v12 := GenS3{F0: (1e10 > 3.14159), F1: 309, F2: 128}
    if true {
        var v13 bool = (("foo bar" != (("\n\t" + "foo bar") + ((("a" + ((("" + ("\n\t" + "\n\t")) + (("foo bar" + "\n\t") + ("\n\t" + "foo bar"))) + ("" + "foo bar"))) + "a") + "a"))) || false)
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 float64 = (float64((-2.25 / (0.0 + 3.14159))) / -2.25)
            v15 = 3.14159
            fmt.Println(v15)
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 int16 = 175
            switch (fi16 ^ (-(fi16))) {
            case 32768:
                for fi18 := 0; fi18 < 3; fi18++ {
                    v19 := GenS3{F0: false, F1: 65535, F2: 9223372036854775807}
                    var v20 int8 = 0
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            case 42:
                v17 -= v17
            default:
                v21 := GenS4{F0: 255, F1: 2, F2: 191}
                _ = v21
            }
            _ = fi16
            _ = v17
        }
    }
    v22 := GenS5{F0: 98, F1: -960, F2: 272}
    _ = v12
    _ = v22
    return 183
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint = 572
    var v3 float32 = -2.25
    v4 := []string{"foo bar", "foo bar", "\n\t"}
    var v5 uint32 = 1
    var v6 int8 = -100
    v3 = v3
    for fi23 := 0; fi23 < 3; fi23++ {
        for fi24 := 0; fi24 < 19; fi24++ {
            v4 = append(v4, "\n\t")
            var v25 uint32 = v5
            _ = fi24
            _ = v25
        }
        v1 += (-(fi23))
        _ = fi23
    }
    var v26 int16 = 42
    if true {
        v1 = genFn0()
        var v27 uint = v2
        _ = v27
    }
    var v28 uint16 = 120
    v29 := GenS1{F0: v26, F1: int8(v28), F2: v28}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v30 uint = v2
    var v31 int = 796
    for fi32 := 0; fi32 < 11; fi32++ {
        if (fi32 < (v1 / (((^(v1))) | 1))) {
            v33 := GenS1{F0: 1000, F1: v6, F2: (v28 + v28)}
            switch "hello" {
            case "a":
                v34 := GenS6{F0: genFn0()}
                _ = v34
            case "":
                var v35 string = "\n\t"
                _ = v35
            case "\n\t":
                var v36 float64 = 0.1
                _ = v36
            }
            v4 = append(v4, ("a" + "hello"))
            _ = v33
        }
        _ = fi32
    }
    v37 := GenS7{F0: 722, F1: v26, F2: v28}
    v5 = v5
    v2 -= (v2 * 4294967295)
    var v38 int32 = 42
    var v39 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
