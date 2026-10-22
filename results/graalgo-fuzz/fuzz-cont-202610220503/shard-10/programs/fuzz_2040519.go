package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 string }
type GenS2 struct { F0 bool; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 int16; F2 bool }

func main() {
    v1 := []string{"foo bar", (("" + "a") + ("a" + ((("hello" + "foo bar") + "hello") + ((("hello" + "") + "\n\t") + "hello"))))}
    var v2 float32 = 0.0
    var v3 uint = 0
    var v4 uint8 = 255
    var v5 uint = (^(v3))
    var v6 float64 = float64(1e10)
    var v7 bool = (226 != 4294967295)
    v8 := GenS1{F0: ("" + "\n\t"), F1: 0, F2: "foo bar"}
    var v9 float32 = (v2 / 1.0)
    if v7 {
        fmt.Println(v2)
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS1{F0: "hello", F1: 0, F2: ("foo bar" + "a")}
            v12 := GenS1{F0: "hello", F1: 127, F2: ""}
            _ = fi10
            _ = v11
            _ = v12
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v13 := GenS2{F0: ((!(v7 && v7)) && (1 > 715)), F1: 4294967295, F2: -78}
    switch 32767 {
    case -59:
        v7 = (4294967295 != 623)
    case -891:
        for fi14 := 0; fi14 < 20; fi14++ {
            for fi15 := 0; fi15 < 4; fi15++ {
                fmt.Println(v6, v5, fi15)
                v4 -= v4
                v16 := GenS3{F0: uint8(v3), F1: -1, F2: (v3 * (^(v3)))}
                if ("\n\t" >= "\n\t") {
                    switch fi14 {
                    case 683:
                        _ = len(v1)
                    case 32767:
                        var v17 uint8 = 18
                        _ = v17
                    case -296:
                        switch (fi14 | 31) {
                        case 1:
                            v18 := GenS2{F0: (uint(v4) >= v5), F1: 2, F2: -736}
                            _ = v18
                        }
                    }
                }
                var v19 int8 = -1
                v1 = append(v1, (("\n\t" + "\n\t") + "a"))
                _ = fi15
                _ = v16
                _ = v19
            }
            _ = fi14
        }
    case 9223372036854775807:
        if (float32(3.14159) < ((3.14159 - (1.5 / v9)) * v9)) {
            var v20 bool = ((v4 % ((v4) | 1)) < (v4 - 254))
            _ = v20
        }
    default:
        if false {
            v3 = v3
            var v21 uint64 = 2
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v21
        }
    }
    _ = len(v1)
    if v7 {
        if (42 >= 0) {
            if len(v1) > 0 {
                v1[0] = (("hello" + "a") + ("\n\t" + "hello"))
            }
            switch -645 {
            case 0:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case -795:
                _ = len(v1)
            case -394:
                v22 := GenS4{F0: -1}
                _ = v22
            default:
                if len(v1) > 0 {
                    v1[0] = ""
                }
            }
        } else {
            v2 = 1e10
            if ("" <= ("" + (((("hello" + ("" + "foo bar")) + "\n\t") + "\n\t") + (("a" + "a") + "\n\t")))) {
                var v23 int32 = 565
                v24 := GenS5{F0: ((v7 && (true && (float32((v6 - v6)) <= (v2 / (0.1 / 1e10))))) || v7)}
                v5 += (^(v5))
                _ = v23
                _ = v24
            } else {
                v9 = (v9 - (v2 + v9))
                var v25 float64 = ((-2.25 - ((v6 * v6) + v6)) / (-2.25 - v6))
                _ = v25
            }
            var v26 uint = v5
            _ = v26
        }
    } else {
        v6 = float64(-2.25)
        fmt.Println(v5, v5)
        var v27 int32 = 42
        _ = v27
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        v29 := GenS6{F0: 32767, F1: 1000, F2: v7}
        _ = fi28
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
