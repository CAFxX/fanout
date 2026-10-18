package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int32 }
type GenS2 struct { F0 int; F1 string; F2 int64 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int64; F1 uint64; F2 int64 }
func genFn0() int8 {
    var v4 uint = 1
    v5 := []int{127}
    if len(v5) > 0 {
        v5[0] = -711
    }
    var v6 bool = true
    if v6 {
        fmt.Println(v6, v4)
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(v6, fi7)
            _ = fi7
        }
        var v8 string = (("hello" + "\n\t") + "\n\t")
        _ = v8
    } else {
        v9 := GenS1{F0: 0, F1: 255, F2: 42}
        fmt.Println(v4, v4)
        if v6 {
            var v10 float64 = 1e10
            if v6 {
                v11 := GenS2{F0: -79, F1: (("a" + "foo bar") + ((((("\n\t" + "\n\t") + "foo bar") + "foo bar") + "hello") + "")), F2: 9223372036854775807}
                if v6 {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        var v13 int16 = 0
                        if v6 {
                            switch "a" {
                            case "hello":
                                v5 = append(v5, (fi12 << ((v4) & 63)))
                            case "a":
                                v13 -= (v13 * v13)
                            case "":
                                var v14 uint16 = 973
                                _ = v14
                            default:
                                switch "" {
                                case "":
                                    for fi15 := 0; fi15 < 19; fi15++ {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v6 = v6
                                        v4 -= v4
                                        _ = fi15
                                    }
                                }
                            }
                            for fi16 := 0; fi16 < 2; fi16++ {
                                v10 = 1.5
                                var v17 uint64 = 1
                                switch (fi16 * fi16) {
                                case 1:
                                    if len(v5) > 0 {
                                        v5[0] = fi16
                                    }
                                }
                                _ = fi16
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 14; fi18++ {
                                v19 := GenS3{F0: 552, F1: 42}
                                switch (fi12 % ((fi18) | 1)) {
                                case 0:
                                    var v20 int16 = (v13 * v13)
                                    _ = v20
                                case -922:
                                    fmt.Println(v13, fi18)
                                case 9223372036854775807:
                                    if (-68 >= 1) {
                                        fmt.Println(v13, fi12, v6)
                                    } else {
                                        fmt.Println(fi12, v4, v13)
                                    }
                                }
                                if len(v5) > 0 {
                                    v5[0] = fi12
                                }
                                _ = fi18
                                _ = v19
                            }
                        }
                        for fi21 := 0; fi21 < 18; fi21++ {
                            var v22 int8 = 1
                            v6 = (2147483647 != 0)
                            fmt.Println(v6, fi21)
                            _ = fi21
                            _ = v22
                        }
                        _ = fi12
                        _ = v13
                    }
                    switch 328 {
                    case 42:
                        var v23 string = "a"
                        _ = v23
                    }
                }
                _ = v11
            }
            _ = v10
        } else {
            var v24 int = -1
            var v25 int16 = -759
            _ = v24
            _ = v25
        }
        _ = v9
    }
    _ = v4
    _ = v5
    _ = v6
    return -42
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 32767
    var v3 string = ""
    v1[-9223372036854775808] = -9223372036854775808
    v26 := GenS4{F0: 9223372036854775807, F1: 565, F2: 1}
    v3 = "foo bar"
    var v27 string = (v3 + "")
    var v28 int64 = 42
    switch v3 {
    case "a":
        if ((v28 <= v28) || false) {
            switch 42 {
            case 0:
                delete(v1, -151)
            case -112:
                switch 4294967295 {
                case 128:
                    var v29 float64 = 3.14159
                    _ = v29
                case 32767:
                    var v30 int16 = v2
                    _ = v30
                case 42:
                    v27 = "foo bar"
                default:
                    for fi31 := 0; fi31 < 20; fi31++ {
                        var v32 uint16 = 42
                        _ = fi31
                        _ = v32
                    }
                }
            }
            switch 127 {
            case -9223372036854775808:
                fmt.Println(v27)
            case 0:
                var v33 uint32 = 0
                _ = v33
            }
            var v34 uint8 = 128
            _ = v34
        }
    default:
        v35 := []int{715}
        _ = v35
    }
    v36 := GenS2{F0: 1, F1: v3, F2: (v28 / (((v28 >> ((18446744073709551615) & 63))) | 1))}
    var v37 int64 = v28
    fmt.Println(v3, v3)
    v38 := GenS3{F0: 4294967295, F1: -9223372036854775808}
    var v39 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
