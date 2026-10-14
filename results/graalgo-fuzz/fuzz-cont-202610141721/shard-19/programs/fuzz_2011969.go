package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int8 }
type GenS2 struct { F0 int64; F1 uint64; F2 int64 }

func main() {
    v1 := []uint32{1, 2, 42}
    v2 := make(map[string]int)
    var v3 int16 = -1
    var v4 uint64 = 2
    v5 := GenS1{F0: 992, F1: "foo bar", F2: 39}
    v4 -= (v4 >> ((1) & 63))
    v4 = v4
    fmt.Println(v3, v4)
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int16 = v3
        var v8 uint8 = 204
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint32 = 374
    for fi10 := 0; fi10 < 12; fi10++ {
        v4 = 4294967295
        _ = fi10
    }
    if ("hello" < "\n\t") {
        fmt.Println(v4, v3)
    } else {
        if (v3 >= (^(v3))) {
            var v11 int = 9223372036854775807
            if false {
                switch ("\n\t" + ((("a" + "foo bar") + "foo bar") + "a")) {
                case "\n\t":
                    for fi12 := 0; fi12 < 8; fi12++ {
                        var v13 float64 = (float64((((1.5 + 0.0) / 1.0) * 3.14159)) - 1e10)
                        if true {
                            for fi14 := 0; fi14 < 10; fi14++ {
                                if (("hello" == "foo bar") && true) {
                                    v3 = (v3 + (v3 * (-(v3))))
                                }
                                _ = fi14
                            }
                            var v15 bool = false
                            fmt.Println(v13)
                            _ = v15
                        } else {
                            v16 := GenS1{F0: (v11 - -9223372036854775808), F1: "foo bar", F2: -128}
                            v13 = v13
                            if ((255 > 128) && true) {
                                if ((uint32(v4) < (^(v9))) && (0 < -1)) {
                                    for fi17 := 0; fi17 < 6; fi17++ {
                                        v13 -= 1.5
                                        _ = fi17
                                    }
                                }
                            }
                            _ = v16
                        }
                        var v18 string = "\n\t"
                        _ = fi12
                        _ = v13
                        _ = v18
                    }
                }
                var v19 uint64 = v4
                _ = v19
            }
            var v20 int64 = -9223372036854775808
            _ = v11
            _ = v20
        }
    }
    v9 = 74
    fmt.Println(v4)
    if false {
        var v21 uint8 = 127
        for fi22 := 0; fi22 < 14; fi22++ {
            v23 := GenS1{F0: fi22, F1: ("foo bar" + ""), F2: -128}
            _ = fi22
            _ = v23
        }
        v4 += v4
        _ = v21
    }
    v24 := v2[("" + "\n\t")]
    if (v3 < v3) {
        fmt.Println(v4)
    } else {
        v4 -= v4
    }
    v3 = v3
    for fi25 := 0; fi25 < 20; fi25++ {
        if (true && true) {
            var v26 int = (-(v24))
            v26 = int(v9)
            var v27 uint8 = 155
            _ = v26
            _ = v27
        } else {
            v28 := GenS2{F0: 0, F1: (v4 + (v4 & (v4 / ((v4) | 1)))), F2: 0}
            _ = v28
        }
        switch fi25 {
        case 127:
            var v29 uint64 = (v4 >> ((18446744073709551615) & 63))
            _ = v29
        }
        var v30 string = "a"
        _ = fi25
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
