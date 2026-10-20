package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int16; F1 int16; F2 uint16 }
func genFn0(p0 bool) bool {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint8 = 75
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 1}
    v10 := []int{32768, -2147483648, 791}
    if ((-27 <= 574) && ((0 >= 170) && (p0 || (970 != 32767)))) {
        v11 := GenS1{F0: -487}
        _ = v11
    } else {
        p0 = (!p0)
        p0 = p0
        if (0 > 2) {
            p0 = (117 > 44)
        } else {
            p0 = (0 < 505)
            for fi12 := 0; fi12 < 18; fi12++ {
                if (!true) {
                    fmt.Println(fi12)
                }
                _ = fi12
            }
            for fi13 := 0; fi13 < 19; fi13++ {
                p0 = (((((("hello" + "") + ("\n\t" + "\n\t")) + "hello") + "") + "") >= "\n\t")
                switch (fi13 / ((4294967295) | 1)) {
                case 4294967295:
                    switch ("hello" + "\n\t") {
                    case "foo bar":
                        if ("" == "foo bar") {
                            p0 = (42 == -128)
                            for fi14 := 0; fi14 < 7; fi14++ {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                p0 = true
                                _ = fi14
                            }
                            v15 := GenS1{F0: fi13}
                            _ = v15
                        }
                    case "a":
                        var v16 bool = p0
                        _ = v16
                    }
                case 9223372036854775807:
                    var v17 uint32 = 65535
                    _ = v17
                }
                _ = fi13
            }
        }
    }
    var v18 int64 = 42
    _ = p0
    _ = v10
    _ = v18
    _ = v9
    return p0
}

func main() {
    var v1 float64 = 3.14159
    v2 := []string{("\n\t" + "foo bar"), "hello"}
    var v3 int = -9223372036854775808
    var v4 int32 = 2147483647
    var v5 bool = true
    var v6 string = "hello"
    v19 := GenS1{F0: v3}
    var v20 uint16 = 0
    var v21 uint16 = v20
    for fi22 := 0; fi22 < 5; fi22++ {
        v21 -= 984
        _ = fi22
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v23 := GenS2{F0: v5, F1: 4294967295}
    v2 = append(v2, v6)
    switch v6 {
    case "hello":
        var v24 int32 = v4
        _ = v24
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v25 float64 = float64(-2.25)
    fmt.Println(v4)
    if true {
        v26 := GenS3{F0: int16(v20), F1: -32768, F2: v20}
        for fi27 := 0; fi27 < 14; fi27++ {
            if v5 {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v28 float32 = 1e10
                v1 -= 1.5
                _ = v28
            } else {
                var v29 float32 = (0.0 * float32(float64(3.14159)))
                var v30 uint32 = 4294967295
                var v31 uint8 = 128
                _ = v29
                _ = v30
                _ = v31
            }
            _ = fi27
        }
        fmt.Println(v21, v5)
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
