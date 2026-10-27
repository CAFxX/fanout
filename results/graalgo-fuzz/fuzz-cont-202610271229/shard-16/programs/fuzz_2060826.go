package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int32; F1 bool; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int; F2 int }
type GenS5 struct { F0 uint8; F1 int64; F2 int32 }
type GenS6 struct { F0 uint32; F1 bool; F2 int32 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 742
    v3 := make(map[int]int)
    var v4 bool = false
    var v5 int16 = 0
    v6 := []string{"\n\t", ("hello" + ("\n\t" + (("" + "foo bar") + "\n\t"))), "foo bar"}
    switch (v2 | -9223372036854775808) {
    case 634:
        fmt.Println(v4, v4, v2)
    }
    var v7 uint16 = 65535
    fmt.Println(v2)
    switch "" {
    case "":
        switch (v2 / ((300) | 1)) {
        case -9223372036854775808:
            fmt.Println(v2)
        }
    case "hello":
        if (2 < 2) {
            fmt.Println(v7)
        } else {
            if (!v4) {
                if (42 > v5) {
                    v8 := GenS1{F0: 1000, F1: v7}
                    if (!v4) {
                        if len(v6) > 0 {
                            v6[0] = "a"
                        }
                    } else {
                        fmt.Println(v4, v7)
                    }
                    _ = v8
                }
            }
        }
    }
    switch ("\n\t" + ("\n\t" + "")) {
    case "\n\t":
        v7 = v7
    case "a":
        if ((v7 * 255) == (^(v7))) {
            if ((1e10 + (float64(((float32(3.14159) / (((3.14159 * 0.1) / (1.5 / 1e10)) * -2.25)) + -2.25)) * (0.1 + 1.5))) > 0.1) {
                v4 = v4
                if ((18446744073709551615 == 4294967295) && true) {
                    v7 = v7
                }
            } else {
                for fi9 := 0; fi9 < 9; fi9++ {
                    switch int(v5) {
                    case 395:
                        var v10 uint8 = 42
                        _ = v10
                    case -2147483648:
                        v4 = ("\n\t" == "")
                    case 0:
                        delete(v3, v2)
                    }
                    _ = fi9
                }
                v11 := GenS2{F0: 42, F1: (!(42 > -128)), F2: 713}
                v12 := GenS1{F0: (v5 >> ((255) & 63)), F1: 32767}
                _ = v11
                _ = v12
            }
            var v13 int8 = 1
            var v14 string = (("\n\t" + "foo bar") + "\n\t")
            _ = v13
            _ = v14
        } else {
            switch v2 {
            case 127:
                for fi15 := 0; fi15 < 10; fi15++ {
                    var v16 uint64 = 47
                    var v17 int16 = (v5 | v5)
                    for fi18 := 0; fi18 < 16; fi18++ {
                        var v19 int64 = -593
                        var v20 int = fi15
                        v20 = 32767
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            case 625:
                v5 += v5
            }
        }
    }
    if true {
        if v4 {
            v5 = (v5 / (((v5 + v5)) | 1))
        } else {
            if v4 {
                fmt.Println(v4, v5)
            }
            for fi21 := 0; fi21 < 20; fi21++ {
                if ((v4 || (127 == 24)) && true) {
                    v22 := v1[(-(v2))]
                    _ = v22
                } else {
                    if ((((((((3.14159 + (float32(float64(-2.25)) - -2.25)) - (1e10 + 1e10)) + 1.5) + 0.1) + 0.0) * 0.0) / 1.0) <= 3.14159) {
                        v23 := GenS3{F0: fi21}
                        _ = v23
                    } else {
                        var v24 uint64 = 951
                        _ = v24
                    }
                    v5 = v5
                    v25 := GenS4{F0: v2, F1: -1, F2: v2}
                    _ = v25
                }
                for fi26 := 0; fi26 < 13; fi26++ {
                    v3[fi21] = 32768
                    if false {
                        v27 := GenS4{F0: 0, F1: 2147483647, F2: int(v7)}
                        v2 += 686
                        _ = v27
                    } else {
                        v28 := GenS4{F0: 393, F1: fi21, F2: (fi21 ^ 9223372036854775807)}
                        _ = v28
                    }
                    _ = fi26
                }
                fmt.Println(v5, v5)
                _ = fi21
            }
        }
        if (((42 <= 176) || v4) || (v4 && v4)) {
            v5 = v5
            v29 := GenS1{F0: (v5 / ((v5) | 1)), F1: v7}
            if v4 {
                v4 = v4
            } else {
                _ = len(v6)
                v2 += v2
            }
            _ = v29
        } else {
            v4 = true
            var v30 int16 = (^(v5))
            _ = v30
        }
        fmt.Println(v7, v7, v7)
    }
    v31 := GenS5{F0: 0, F1: -1, F2: -137}
    var v32 int16 = v5
    v6 = append(v6, ("hello" + ("hello" + "foo bar")))
    if len(v6) > 0 {
        v6[0] = "foo bar"
    }
    v33 := v1[v2]
    if (float64(3.14159) < 0.1) {
        v34 := GenS5{F0: 78, F1: -822, F2: -1}
        _ = len(v6)
        _ = v34
    }
    v35 := GenS6{F0: 0, F1: (v33 || false), F2: 694}
    var v36 uint8 = 128
    fmt.Println(v36, v32, v5)
    v33 = (!v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
