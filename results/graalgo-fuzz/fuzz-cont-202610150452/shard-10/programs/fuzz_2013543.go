package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := []int{}
    var v2 bool = (!true)
    v3 := []string{"\n\t", "a", "hello", "\n\t"}
    var v4 float32 = 0.0
    v5 := make(map[uint32]int)
    var v6 uint8 = 237
    var v7 uint8 = (^(v6))
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint8 = v7
            switch (-(fi8)) {
            case -9223372036854775808:
                v6 = (^(v10))
            default:
                v11 := GenS1{F0: -1, F1: 588, F2: 42}
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            fmt.Println(v2)
            fmt.Println(v2)
            fmt.Println(v4, v7, fi12)
            v4 += 0.0
            v13 := v5[4294967295]
            _ = fi12
            _ = v13
        }
        _ = fi8
    }
    v4 = v4
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 bool = (2147483647 != -2147483648)
        _ = fi14
        _ = v15
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v16 := GenS2{F0: 726}
    v3 = append(v3, "a")
    _ = len(v5)
    switch 0 {
    case 1:
        v2 = (-1 >= -1)
    case -535:
        switch -9223372036854775808 {
        case 0:
            var v17 uint8 = 12
            _ = v17
        default:
            var v18 float32 = 1e10
            _ = v18
        }
    default:
        v19 := GenS2{F0: -191}
        _ = v19
    }
    fmt.Println(v7, v7)
    _ = len(v3)
    v20 := GenS1{F0: 450, F1: 18446744073709551615, F2: 4294967295}
    v4 = 1.5
    switch -821 {
    case 128:
        if (-659 < 128) {
            var v21 bool = (!v2)
            _ = v21
        } else {
            var v22 bool = (!true)
            fmt.Println(v7)
            switch ((("a" + "") + "a") + ("foo bar" + ("foo bar" + ""))) {
            case "":
                v23 := GenS2{F0: 0}
                _ = v23
            case "foo bar":
                for fi24 := 0; fi24 < 3; fi24++ {
                    for fi25 := 0; fi25 < 18; fi25++ {
                        v7 = v7
                        if v22 {
                            var v26 uint16 = 238
                            _ = v26
                        } else {
                            var v27 uint32 = 42
                            switch "\n\t" {
                            case "foo bar":
                                switch (((("foo bar" + ("hello" + "a")) + "foo bar") + "a") + "\n\t") {
                                case "foo bar":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                            var v28 float64 = 0.0
                            _ = v27
                            _ = v28
                        }
                        _ = fi25
                    }
                    v4 -= (3.14159 + v4)
                    _ = fi24
                }
            case "\n\t":
                var v29 int16 = 42
                _ = v29
            default:
                if ((!((true && true) && (124 == 100))) && ((v7 >> ((42) & 63)) >= v7)) {
                    delete(v5, 51)
                    v7 += v6
                }
            }
            _ = v22
        }
    case 4294967295:
        v30 := GenS3{F0: v6, F1: 1}
        _ = v30
    case -1:
        var v31 int32 = 95
        _ = v31
    }
    v4 += (float32(0.0) * 1e10)
    v3 = append(v3, "\n\t")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
