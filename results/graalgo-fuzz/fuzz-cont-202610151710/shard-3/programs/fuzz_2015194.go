package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int; F1 uint; F2 int64 }
func genFn0() int64 {
    v4 := GenS1{F0: -812, F1: -1}
    var v5 uint8 = 138
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 int16 = 340
    var v2 int64 = -594
    var v3 int32 = -2147483648
    v3 += (v3 << ((4294967295) & 63))
    v2 += genFn0()
    var v6 bool = (255 > 1)
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: v3, F1: v1}
        _ = fi7
        _ = v8
    }
    if (!(145 >= 787)) {
        v3 = (v3 + 1)
        var v9 string = ("a" + "")
        _ = v9
    }
    fmt.Println(v6)
    var v10 bool = (float32((1.5 / 1e10)) < (((1e10 - 1e10) * ((1.5 + (1e10 * 0.0)) + (0.1 * 1.5))) + 3.14159))
    if true {
        var v11 uint32 = 900
        if (v1 != v1) {
            v2 -= v2
            var v12 float32 = 0.1
            _ = v12
        }
        switch 1 {
        case -2147483648:
            var v13 uint = 295
            _ = v13
        }
        _ = v11
    } else {
        for fi14 := 0; fi14 < 1; fi14++ {
            if (!((!(!(!((fi14 ^ fi14) < fi14)))) && v6)) {
                v1 = (-(v1))
            } else {
                for fi15 := 0; fi15 < 9; fi15++ {
                    var v16 uint32 = 876
                    _ = fi15
                    _ = v16
                }
            }
            fmt.Println(v1)
            v17 := []int{(fi14 & fi14)}
            _ = fi14
            _ = v17
        }
        fmt.Println(v1)
    }
    switch "a" {
    case "":
        switch 32767 {
        case 2147483647:
            v18 := GenS1{F0: (v3 * 796), F1: 42}
            _ = v18
        case 32768:
            v19 := []int{32768, 565, -950, 202}
            _ = v19
        default:
            if v10 {
                v20 := GenS1{F0: (v3 & 833), F1: (v1 * (v1 + -201))}
                var v21 uint16 = 218
                _ = v20
                _ = v21
            }
        }
    case "foo bar":
        for fi22 := 0; fi22 < 13; fi22++ {
            if v6 {
                if (v10 && (147 < 42)) {
                    if v6 {
                        v23 := []int{}
                        _ = v23
                    } else {
                        v24 := GenS1{F0: v3, F1: v1}
                        _ = v24
                    }
                    switch fi22 {
                    case 367:
                        var v25 uint32 = 2147483648
                        _ = v25
                    default:
                        var v26 uint8 = 42
                        _ = v26
                    }
                } else {
                    v27 := GenS2{F0: int8(v3), F1: (v2 & (v2 & int64(v1)))}
                    _ = v27
                }
            } else {
                v1 = -32768
            }
            _ = fi22
        }
    case "\n\t":
        v28 := []int{42, 32768}
        _ = v28
    }
    fmt.Println(v3, v6, v6)
    for fi29 := 0; fi29 < 9; fi29++ {
        if (0 != (-(v1))) {
            v30 := GenS2{F0: 124, F1: int64(v3)}
            fmt.Println(v10)
            _ = v30
        } else {
            v1 = v1
        }
        for fi31 := 0; fi31 < 16; fi31++ {
            switch (("foo bar" + ((("hello" + "foo bar") + "foo bar") + "foo bar")) + ("" + "foo bar")) {
            case "a":
                v32 := []int{fi31, (fi31 / ((fi31) | 1)), (-(fi31)), 32767}
                _ = v32
            case "\n\t":
                v33 := GenS3{F0: (^(fi29)), F1: 18446744073709551615, F2: v2}
                _ = v33
            }
            if ("foo bar" != "") {
                v34 := GenS3{F0: int(v2), F1: 255, F2: (v2 * (v2 / (((v2 - (v2 >> ((2) & 63)))) | 1)))}
                _ = v34
            }
            v6 = true
            _ = fi31
        }
        _ = fi29
    }
    fmt.Println(v2)
    var v35 string = "foo bar"
    var v36 int16 = (v1 + 0)
    v37 := GenS1{F0: (v3 ^ 299), F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v6
}
