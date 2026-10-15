package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint) uint32 {
    v4 := []int{int(p0), 127}
    if len(v4) > 0 {
        v4[0] = 858
    }
    _ = len(v4)
    p1 += (p1 ^ (p1 | p1))
    _ = p0
    _ = p1
    _ = v4
    return 206
}
func genFn1(p0 uint64) string {
    v5 := GenS1{F0: 714, F1: 2147483647}
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint64 = p0
            v9 := GenS1{F0: 2, F1: 0}
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 uint32 = 2147483647
        _ = fi6
        _ = v10
    }
    var v11 int16 = -774
    v12 := []int{942, -1}
    _ = p0
    _ = v11
    _ = v12
    _ = v5
    return "foo bar"
}

func main() {
    var v1 uint = 1
    var v2 uint64 = 2
    var v3 bool = true
    switch "" {
    case "foo bar":
        var v13 uint32 = 42
        _ = v13
    default:
        fmt.Println(v2)
    }
    var v14 uint = (v1 % (((v1 / ((4294967295) | 1))) | 1))
    var v15 int64 = 42
    if v3 {
        v14 = v1
    } else {
        v2 += (^(v2))
        if (1 == 27) {
            if (42 == 1) {
                var v16 int64 = v15
                _ = v16
            }
            v1 -= 4294967295
            var v17 uint64 = v2
            _ = v17
        } else {
            v18 := []int{-9223372036854775808, -614}
            if (884 > 345) {
                for fi19 := 0; fi19 < 5; fi19++ {
                    switch ((("\n\t" + ("\n\t" + "foo bar")) + "a") + "") {
                    case "a":
                        var v20 float64 = 1e10
                        _ = v20
                    default:
                        if ((1e10 / 0.1) < float64(0.1)) {
                            fmt.Println(v15, v15, fi19)
                            if (2147483648 > 2) {
                                fmt.Println(v1, v1, fi19)
                            }
                        }
                    }
                    v14 -= (v1 ^ (^(v14)))
                    v2 -= 714
                    _ = fi19
                }
            } else {
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
            }
            _ = v18
        }
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        if (42 < 2) {
            v22 := GenS2{F0: (("foo bar" + ("" + ("\n\t" + "foo bar"))) + "foo bar")}
            switch ("foo bar" + ((((("hello" + "foo bar") + ("\n\t" + (("hello" + (((("\n\t" + "\n\t") + "") + "foo bar") + "foo bar")) + ("hello" + (("a" + "\n\t") + "foo bar"))))) + "a") + ("" + "a")) + "a")) {
            case "":
                switch "a" {
                case "hello":
                    var v23 bool = (255 > 65535)
                    _ = v23
                case "a":
                    var v24 int32 = -1
                    _ = v24
                }
            }
            _ = v22
        }
        v25 := GenS3{F0: 129}
        v14 += v1
        _ = fi21
        _ = v25
    }
    if (((-(v15)) != (^(v15))) && v3) {
        fmt.Println(v14)
        v1 -= (v14 * v1)
    } else {
        v2 = (v2 % (((v2 - (v2 + 1))) | 1))
        var v26 uint32 = 42
        _ = v26
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        v2 = v2
        _ = fi27
    }
    if (-1 > -128) {
        if (!(42 <= 42)) {
            switch 1 {
            case 2147483647:
                if v3 {
                    v14 -= (v1 & v1)
                    var v28 bool = v3
                    var v29 uint8 = 2
                    _ = v28
                    _ = v29
                } else {
                    fmt.Println(v3)
                    var v30 int16 = 1000
                    _ = v30
                }
            }
            v3 = v3
            fmt.Println(v15)
        } else {
            v3 = (v3 || (!true))
        }
        v31 := GenS1{F0: 32767, F1: -922}
        _ = v31
    } else {
        for fi32 := 0; fi32 < 8; fi32++ {
            v14 += (^(v14))
            _ = fi32
        }
        switch -2147483648 {
        case 716:
            for fi33 := 0; fi33 < 6; fi33++ {
                v34 := []int{(fi33 << ((v1) & 63)), (fi33 & fi33), fi33, 9223372036854775807}
                _ = fi33
                _ = v34
            }
        case -2147483648:
            var v35 string = "hello"
            _ = v35
        case -9223372036854775808:
            fmt.Println(v2, v1)
        default:
            v3 = (!(v3 || (32767 == -58)))
        }
        fmt.Println(v2, v2, v2)
    }
    var v36 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
}
