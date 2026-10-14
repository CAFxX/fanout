package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 string; F2 uint32 }
type GenS4 struct { F0 int8; F1 bool }

func main() {
    var v1 uint16 = 505
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    if false {
        v4 := GenS1{F0: 0, F1: ("a" + "\n\t")}
        var v5 int8 = -128
        _ = v4
        _ = v5
    } else {
        var v6 uint = 255
        fmt.Println(v1, v1)
        if false {
            v7 := GenS2{F0: 65535}
            fmt.Println(v6, v6)
            if ((4294967295 >= 0) || (!(2 <= 249))) {
                fmt.Println(v6, v1)
                if (-1 >= int16(v1)) {
                    for fi8 := 0; fi8 < 4; fi8++ {
                        v9 := GenS1{F0: -9223372036854775808, F1: "\n\t"}
                        _ = fi8
                        _ = v9
                    }
                }
                switch "" {
                case "":
                    var v10 int8 = 60
                    _ = v10
                case "foo bar":
                    for fi11 := 0; fi11 < 11; fi11++ {
                        if (uint(v1) < 945) {
                            var v12 uint8 = 2
                            _ = v12
                        }
                        switch (-(fi11)) {
                        case -2147483648:
                            v6 = v6
                        case 127:
                            v13 := GenS3{F0: (v6 / (((v6 >> (((v6 | (v6 >> ((v6) & 63)))) & 63))) | 1)), F1: "\n\t", F2: 329}
                            _ = v13
                        }
                        _ = fi11
                    }
                case "\n\t":
                    v14 := []int{-70, 466}
                    _ = v14
                }
            } else {
                v1 -= v1
                fmt.Println(v1, v6)
            }
            _ = v7
        }
        _ = v6
    }
    v1 -= v1
    switch ((("\n\t" + "hello") + ("\n\t" + "hello")) + ("\n\t" + "hello")) {
    case "hello":
        v1 -= v1
    case "a":
        fmt.Println(v1)
    default:
        v1 = v1
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS1{F0: 1, F1: "a"}
        fmt.Println(fi15)
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := []int{32767, 127, (fi17 * -241), (^(fi15))}
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    v19 := GenS4{F0: 100, F1: true}
    v1 -= (^(v1))
    switch 794 {
    case -900:
        v20 := []int{641}
        _ = v20
    default:
        for fi21 := 0; fi21 < 13; fi21++ {
            v22 := GenS2{F0: 42}
            v23 := GenS3{F0: 65535, F1: "\n\t", F2: 2147483648}
            if ((("" + ("\n\t" + "foo bar")) == "foo bar") || (!(false && (!((0.0 / ((1e10 + (-2.25 / (float64(float32(0.0)) + float64((1e10 * 1e10))))) / -2.25)) != 1e10))))) {
                v24 := []int{fi21, fi21, fi21, (fi21 + fi21)}
                v1 -= v1
                _ = v24
            } else {
                v1 = v1
                var v25 uint = 42
                _ = v25
            }
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    v1 += (v1 - 604)
    v1 += (^(v1))
    v26 := GenS4{F0: 1, F1: false}
    fmt.Println(v1)
    if (!false) {
        var v27 float64 = 0.1
        var v28 float64 = v27
        v29 := GenS4{F0: 0, F1: (19 == 982)}
        _ = v27
        _ = v28
        _ = v29
    } else {
        v30 := GenS1{F0: 59, F1: ((("" + ((("\n\t" + "hello") + "a") + "foo bar")) + "foo bar") + "")}
        v31 := []int{-1, 1, 32767}
        _ = v30
        _ = v31
    }
    v1 = v1
    v1 = v1
    var v32 uint32 = 2
    for fi33 := 0; fi33 < 7; fi33++ {
        var v34 uint32 = v32
        _ = fi33
        _ = v34
    }
    v1 = 345
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v32
}
