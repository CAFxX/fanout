package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint }
func genFn0() uint8 {
    v4 := []int{-456, 42, 1, 9223372036854775807}
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int8 = -128
        v7 := GenS1{F0: 0}
        switch (fi5 * (fi5 - fi5)) {
        case -2147483648:
            v8 := GenS1{F0: 429}
            _ = v8
        case -176:
            if (1.5 <= 0.0) {
                var v9 uint64 = 4294967295
                _ = v9
            } else {
                v6 -= v6
                for fi10 := 0; fi10 < 7; fi10++ {
                    v6 = (-(v6))
                    _ = fi10
                }
            }
        case 9223372036854775807:
            v6 -= 42
        default:
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := GenS1{F0: 1}
                _ = fi11
                _ = v12
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v13 := GenS1{F0: 18446744073709551615}
    for fi14 := 0; fi14 < 14; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            fmt.Println(fi15, fi15)
            var v16 int32 = 733
            v16 = (-(v16))
            _ = fi15
            _ = v16
        }
        var v17 uint16 = 2
        v17 -= (v17 + v17)
        _ = fi14
        _ = v17
    }
    _ = v13
    _ = v4
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -469
    v3 := make(map[uint32]int)
    var v18 float32 = 1.5
    v18 = (1.5 + v18)
    v18 = (v18 - (v18 - 0.1))
    var v19 string = "hello"
    v18 = float32((1e10 / 1.0))
    var v20 bool = ((v2 << ((255) & 63)) == -9223372036854775808)
    switch v19 {
    case "\n\t":
        for fi21 := 0; fi21 < 8; fi21++ {
            if (391 <= -32768) {
                var v22 bool = v20
                switch v19 {
                case "foo bar":
                    switch (("hello" + (v19 + ("\n\t" + (v19 + v19)))) + "\n\t") {
                    case "foo bar":
                        fmt.Println(v2, v18, fi21)
                    case "":
                        switch v19 {
                        case "":
                            var v23 bool = true
                            _ = v23
                        case "foo bar":
                            v24 := GenS2{F0: -86, F1: 1}
                            _ = v24
                        case "\n\t":
                            if v20 {
                                var v25 uint16 = 917
                                var v26 int = (fi21 - 9223372036854775807)
                                _ = v25
                                _ = v26
                            } else {
                                v27 := GenS1{F0: 1}
                                v28 := GenS1{F0: 4294967295}
                                fmt.Println(v18, v20)
                                _ = v27
                                _ = v28
                            }
                        }
                    case "hello":
                        v29 := GenS2{F0: -75, F1: 4294967295}
                        _ = v29
                    default:
                        v30 := []int{fi21, 399}
                        _ = v30
                    }
                }
                _ = v22
            }
            var v31 int = 127
            v32 := v1[0]
            _ = fi21
            _ = v31
            _ = v32
        }
    }
    fmt.Println(v20, v19, v19)
    for fi33 := 0; fi33 < 11; fi33++ {
        v34 := GenS1{F0: 42}
        v20 = (501 >= 192)
        _ = fi33
        _ = v34
    }
    var v35 int16 = 821
    var v36 uint = 812
    if (140 != 0) {
        var v37 string = v19
        v38 := []int{4294967295, 32768, 495}
        v19 = v19
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 15; fi39++ {
        for fi40 := 0; fi40 < 15; fi40++ {
            fmt.Println(v20)
            v35 = (v35 % ((v35) | 1))
            fmt.Println(v19)
            _ = fi40
        }
        _ = fi39
    }
    if (!(!v20)) {
        v35 -= v35
    }
    if v20 {
        var v41 uint16 = 65535
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
}
