package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint16 }
func genFn0() int64 {
    var v7 int = -534
    for fi8 := 0; fi8 < 13; fi8++ {
        v7 = v7
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: "foo bar", F1: 65535, F2: 0}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v7 -= (v7 * (v7 * 4294967295))
    v11 := GenS2{F0: -32768}
    _ = v11
    _ = v7
    return -754
}

func main() {
    var v1 bool = false
    var v2 uint16 = 32767
    var v3 int8 = -128
    var v4 uint32 = 42
    var v5 float32 = (3.14159 * (float32(-2.25) * 1e10))
    var v6 bool = v1
    if (340 > 904) {
        v6 = (!(!false))
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 uint32 = v4
            _ = fi12
            _ = v13
        }
    }
    switch "" {
    case "\n\t":
        if (203 < 17) {
            v14 := GenS1{F0: "", F1: v2, F2: 2}
            if true {
                var v15 int = 0
                fmt.Println(v1)
                if ((1e10 + -2.25) != 1e10) {
                    for fi16 := 0; fi16 < 15; fi16++ {
                        v5 += v5
                        var v17 string = "foo bar"
                        fmt.Println(v4, v1, v6)
                        _ = fi16
                        _ = v17
                    }
                    v18 := GenS2{F0: -32768}
                    if (-114 > (v3 << ((0) & 63))) {
                        if v1 {
                            v19 := GenS3{F0: 42, F1: 250, F2: v2}
                            _ = v19
                        } else {
                            var v20 uint8 = 127
                            v6 = v1
                            if (("a" + "") <= "") {
                                v20 -= 127
                                v21 := []int{v15}
                                var v22 string = "a"
                                _ = v21
                                _ = v22
                            }
                            _ = v20
                        }
                    }
                    _ = v18
                } else {
                    var v23 float64 = 1e10
                    switch "hello" {
                    case "a":
                        for fi24 := 0; fi24 < 14; fi24++ {
                            var v25 int = 139
                            v23 = v23
                            _ = fi24
                            _ = v25
                        }
                    }
                    _ = v23
                }
                _ = v15
            } else {
                v3 -= (v3 | v3)
            }
            v2 = v2
            _ = v14
        }
    }
    v3 = v3
    var v26 uint64 = 804
    v1 = (621 <= v26)
    fmt.Println(v3)
    fmt.Println(v3, v4)
    v27 := []int{}
    if (v3 == 127) {
        v6 = ("hello" == ("hello" + "a"))
        for fi28 := 0; fi28 < 1; fi28++ {
            switch ((("" + "\n\t") + "") + "foo bar") {
            case "\n\t":
                v27 = append(v27, (fi28 << ((65535) & 63)))
            default:
                fmt.Println(v4)
            }
            v2 -= 65535
            _ = fi28
        }
        v29 := GenS3{F0: (v2 ^ (v2 % ((v2) | 1))), F1: uint8(v2), F2: 42}
        _ = v29
    }
    v27 = append(v27, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
