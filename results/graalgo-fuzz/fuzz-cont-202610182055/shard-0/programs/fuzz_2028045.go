package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }

func main() {
    var v1 bool = (false || (true || (float32(float64(1.5)) <= 1.5)))
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    v4 := []uint32{65535, 833}
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]int)
    switch -2147483648 {
    case 9223372036854775807:
        var v7 int8 = 1
        _ = v7
    case 357:
        for fi8 := 0; fi8 < 19; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                v1 = (42 <= 115)
                v10 := v2[("" + "")]
                _ = fi9
                _ = v10
            }
            switch "a" {
            case "foo bar":
                if (138 >= 1) {
                    for fi11 := 0; fi11 < 12; fi11++ {
                        v1 = (100 == -128)
                        if (127 > 0) {
                            v1 = true
                        } else {
                            v1 = v1
                            var v12 uint16 = 281
                            _ = v12
                        }
                        switch "foo bar" {
                        case "hello":
                            if len(v4) > 0 {
                                v4[0] = 4294967295
                            }
                        case "a":
                            v13 := GenS1{F0: 100}
                            _ = v13
                        default:
                            v1 = (!(0 >= 0))
                        }
                        _ = fi11
                    }
                }
            }
            _ = fi8
        }
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        switch (fi14 / ((fi14) | 1)) {
        case 32767:
            var v15 uint = 279
            _ = v15
        case -690:
            for fi16 := 0; fi16 < 5; fi16++ {
                fmt.Println(v1)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if true {
                    if (("" + ("hello" + "")) != "a") {
                        v17 := GenS1{F0: 1}
                        _ = v17
                    }
                }
                _ = fi16
            }
        case 385:
            var v18 uint8 = 133
            _ = v18
        }
        switch ("\n\t" + "") {
        case "hello":
            fmt.Println(fi14, v1)
        default:
            var v19 string = "a"
            _ = v19
        }
        v1 = (0 >= -1)
        _ = fi14
    }
    var v20 uint16 = 0
    v4 = append(v4, 4294967295)
    if ("foo bar" >= "hello") {
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := GenS1{F0: -89}
            _ = fi21
            _ = v22
        }
        v23 := GenS2{F0: 9223372036854775807, F1: 0}
        _ = v23
    }
    v24 := GenS2{F0: 778, F1: -2147483648}
    v25 := GenS1{F0: 0}
    for fi26 := 0; fi26 < 8; fi26++ {
        if v1 {
            v1 = v1
        }
        var v27 int16 = 1000
        var v28 bool = (0.1 >= -2.25)
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        if v1 {
            if false {
                v30 := GenS1{F0: 1}
                _ = v30
            }
        }
        v31 := GenS3{F0: 2}
        if (0 == 42) {
            var v32 uint16 = 65535
            v33 := GenS4{F0: (!((-128 <= -103) || v1))}
            _ = v32
            _ = v33
        } else {
            v4 = append(v4, 4294967295)
            v20 -= (^(v20))
        }
        _ = fi29
        _ = v31
    }
    var v34 int = 128
    if v1 {
        var v35 float32 = 1e10
        for fi36 := 0; fi36 < 19; fi36++ {
            v37 := GenS2{F0: 1, F1: -257}
            var v38 uint16 = (v20 + v20)
            _ = len(v2)
            _ = fi36
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 11; fi39++ {
            var v40 uint64 = 2
            _ = fi39
            _ = v40
        }
        _ = v35
    } else {
        v1 = true
        v6[65535] = -250
    }
    if (2147483647 != 806) {
        v4 = append(v4, 2147483648)
        var v41 float32 = float32(1e10)
        switch "foo bar" {
        case "\n\t":
            if false {
                _ = len(v4)
                for fi42 := 0; fi42 < 11; fi42++ {
                    if len(v4) > 0 {
                        v4[0] = 535
                    }
                    var v43 int = v34
                    _ = fi42
                    _ = v43
                }
            }
        case "a":
            if (v1 && ((!v1) && (42 >= 127))) {
                v34 = v34
                v41 += -2.25
            }
        default:
            v20 += v20
        }
        _ = v41
    }
    fmt.Println(v34, v34, v1)
    v4 = append(v4, 1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
