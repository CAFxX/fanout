package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 string; F2 int }
func genFn0(p0 uint32, p1 bool) int32 {
    var v5 uint64 = 405
    for fi6 := 0; fi6 < 2; fi6++ {
        switch (((("foo bar" + "foo bar") + ("a" + "")) + "foo bar") + "hello") {
        case "foo bar":
            var v7 int8 = -5
            _ = v7
        case "hello":
            v8 := GenS1{F0: 42, F1: v5, F2: -1}
            _ = v8
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v5
    return -2147483648
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    v4 := []string{"foo bar", "", ("hello" + "\n\t"), "\n\t"}
    v9 := GenS1{F0: 305, F1: 18446744073709551615, F2: 9223372036854775807}
    v2[0] = 128
    var v10 float32 = (0.0 - 1.5)
    if len(v4) > 0 {
        v4[0] = (((("foo bar" + "foo bar") + "hello") + ("a" + "a")) + "a")
    }
    if (!false) {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 float32 = (v10 / 0.1)
            delete(v3, (fi11 & (fi11 / ((-41) | 1))))
            if false {
                v10 += 0.0
                var v13 uint64 = 2
                _ = v13
            }
            _ = fi11
            _ = v12
        }
    }
    v14 := GenS1{F0: 1, F1: 135, F2: 222}
    var v15 int8 = 100
    fmt.Println(v10)
    v16 := GenS1{F0: -2147483648, F1: 18446744073709551615, F2: 0}
    switch 0 {
    case -706:
        switch 0 {
        case 4294967295:
            if (3.14159 == 1e10) {
                v17 := GenS1{F0: 143, F1: 18446744073709551615, F2: 9223372036854775807}
                if (652 <= 42) {
                    if (false && (!(false && (276 != 4294967295)))) {
                        var v18 uint64 = 617
                        if true {
                            if len(v4) > 0 {
                                v4[0] = ("\n\t" + (("hello" + "") + ("" + "\n\t")))
                            }
                        } else {
                            var v19 int16 = 0
                            switch 128 {
                            case 1:
                                if true {
                                    if ((v15 % ((v15) | 1)) <= 42) {
                                        v20 := v3[814]
                                        _ = v20
                                    } else {
                                        var v21 int8 = v15
                                        _ = v21
                                    }
                                } else {
                                    v22 := GenS1{F0: 132, F1: (v18 % (((v18 ^ 18446744073709551615)) | 1)), F2: -1}
                                    v23 := v3[-887]
                                    var v24 int64 = -9223372036854775808
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                            default:
                                var v25 uint = 473
                                _ = v25
                            }
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                            _ = v19
                        }
                        _ = v18
                    }
                }
                for fi26 := 0; fi26 < 16; fi26++ {
                    fmt.Println(v10)
                    v27 := GenS2{F0: 1}
                    _ = fi26
                    _ = v27
                }
                _ = v17
            }
        case 42:
            v28 := GenS2{F0: 42}
            _ = v28
        }
    case 4294967295:
        v10 += v10
    case -2147483648:
        if false {
            v29 := GenS3{F0: -1, F1: ("\n\t" + ""), F2: 0}
            _ = v29
        } else {
            v15 -= -1
            if false {
                v30 := GenS3{F0: 1, F1: ("hello" + ""), F2: -412}
                _ = v30
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
