package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 bool; F2 bool }
func genFn0(p0 uint32, p1 int16) int32 {
    p1 -= 42
    var v4 int16 = 42
    p0 = (p0 + 142)
    v5 := GenS1{F0: 18446744073709551615, F1: (!false)}
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 float64 = 1e10
        var v8 int = fi6
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 42
    var v3 int16 = 0
    v9 := []int{32768}
    v2 = (v2 / ((v2) | 1))
    switch 9223372036854775807 {
    case 0:
        v2 += v2
    case 127:
        switch "" {
        case "\n\t":
            v10 := GenS1{F0: 908, F1: (1 < 42)}
            _ = v10
        case "hello":
            v2 = (v2 + -90)
        }
    case 1:
        v11 := GenS1{F0: 42, F1: (true && (657 <= (v3 << ((1) & 63))))}
        _ = v11
    }
    delete(v1, ("hello" + "hello"))
    var v12 int64 = 9223372036854775807
    _ = len(v1)
    switch "a" {
    case "a":
        var v13 int8 = -128
        _ = v13
    case "":
        var v14 int8 = 6
        _ = v14
    case "hello":
        for fi15 := 0; fi15 < 15; fi15++ {
            if ((fi15 % (((fi15 + (fi15 | 1))) | 1)) != 0) {
                var v16 uint16 = 255
                if (1 < 1) {
                    if false {
                        for fi17 := 0; fi17 < 4; fi17++ {
                            v18 := GenS1{F0: 0, F1: true}
                            _ = fi17
                            _ = v18
                        }
                        v9 = append(v9, -370)
                        delete(v1, "")
                    } else {
                        var v19 uint32 = 1
                        _ = v19
                    }
                    var v20 int32 = genFn0(150, (v3 | v3))
                    _ = v20
                }
                fmt.Println(v16)
                _ = v16
            }
            _ = fi15
        }
    default:
        var v21 bool = (32767 >= 32767)
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        if len(v9) > 0 {
            v9[0] = fi22
        }
        _ = fi22
    }
    var v23 int32 = genFn0(2147483647, 42)
    v3 = 294
    if true {
        if true {
            switch "hello" {
            case "":
                for fi24 := 0; fi24 < 14; fi24++ {
                    v25 := GenS1{F0: 292, F1: (!false)}
                    if false {
                        v3 -= (v3 >> ((758) & 63))
                        _ = len(v1)
                        fmt.Println(v12, v23)
                    }
                    v1["foo bar"] = (fi24 * fi24)
                    _ = fi24
                    _ = v25
                }
            }
            v3 = v3
            v12 += (v12 / ((-685) | 1))
        }
        var v26 int8 = 0
        _ = v26
    } else {
        for fi27 := 0; fi27 < 9; fi27++ {
            v28 := GenS1{F0: 994, F1: (("foo bar" + "hello") > (("" + "hello") + "\n\t"))}
            for fi29 := 0; fi29 < 12; fi29++ {
                fmt.Println(fi29, fi29)
                _ = fi29
            }
            switch fi27 {
            case 4294967295:
                v3 = -32768
            }
            _ = fi27
            _ = v28
        }
        switch "foo bar" {
        case "foo bar":
            var v30 float64 = 1e10
            _ = v30
        }
    }
    if (0.1 >= 0.1) {
        switch ("foo bar" + "") {
        case "hello":
            for fi31 := 0; fi31 < 7; fi31++ {
                if (0 <= 2147483648) {
                    var v32 float32 = (1e10 + ((1e10 * ((0.1 - -2.25) - 1.5)) * float32((-2.25 - 1e10))))
                    _ = v32
                }
                for fi33 := 0; fi33 < 16; fi33++ {
                    if (1 != 469) {
                        v34 := GenS1{F0: 822, F1: (65535 > 61)}
                        v35 := GenS2{F0: (("" + "") + "\n\t")}
                        _ = v34
                        _ = v35
                    }
                    _ = fi33
                }
                if (v3 < (v3 | (v3 ^ 1))) {
                    var v36 int16 = 560
                    v37 := v1[("a" + "\n\t")]
                    v37 += -699
                    _ = v36
                    _ = v37
                } else {
                    switch (fi31 << ((4294967295) & 63)) {
                    case 9223372036854775807:
                        fmt.Println(v3)
                    default:
                        for fi38 := 0; fi38 < 14; fi38++ {
                            v12 += (v12 | 42)
                            _ = fi38
                        }
                    }
                    v39 := GenS2{F0: (("foo bar" + "foo bar") + "foo bar")}
                    _ = v39
                }
                _ = fi31
            }
        }
        fmt.Println(v23, v23, v3)
        var v40 float64 = (1.5 / 1.0)
        _ = v40
    }
    v41 := GenS3{F0: 4294967295, F1: true, F2: ((65535 == 2147483648) || true)}
    v2 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v41
    _ = v9
}
