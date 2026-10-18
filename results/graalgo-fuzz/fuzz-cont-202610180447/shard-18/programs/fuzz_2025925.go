package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int64; F1 int8 }

func main() {
    var v1 uint32 = 650
    v2 := make(map[int]bool)
    v3 := []int{42, 42}
    var v4 int32 = -1
    var v5 int16 = 192
    v4 = (v4 * v4)
    v2[4294967295] = true
    for fi6 := 0; fi6 < 8; fi6++ {
        v5 = (^(v5))
        if true {
            for fi7 := 0; fi7 < 19; fi7++ {
                v8 := GenS1{F0: 127}
                _ = fi7
                _ = v8
            }
            v5 = -371
            var v9 int64 = 45
            _ = v9
        }
        _ = fi6
    }
    v1 = v1
    if false {
        var v10 int64 = -9223372036854775808
        if true {
            switch int(v5) {
            case 42:
                if true {
                    var v11 int32 = (v4 % (((v4 >> ((18446744073709551615) & 63))) | 1))
                    _ = v11
                }
            case 199:
                switch 578 {
                case 42:
                    for fi12 := 0; fi12 < 3; fi12++ {
                        fmt.Println(v5, v10)
                        _ = fi12
                    }
                case 1:
                    fmt.Println(v5, v5, v4)
                case -289:
                    for fi13 := 0; fi13 < 18; fi13++ {
                        for fi14 := 0; fi14 < 1; fi14++ {
                            fmt.Println(v5)
                            fmt.Println(fi14, fi13)
                            _ = fi14
                        }
                        switch (("foo bar" + "") + ("hello" + "")) {
                        case "hello":
                            if ((-(v10)) != (v10 - (v10 << ((18446744073709551615) & 63)))) {
                                var v15 uint = 487
                                var v16 uint16 = 2
                                _ = v15
                                _ = v16
                            }
                        case "foo bar":
                            v5 -= v5
                        default:
                            var v17 int16 = 343
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 3; fi18++ {
                            var v19 uint32 = 278
                            switch fi13 {
                            case 32767:
                                _ = len(v3)
                            case 32768:
                                var v20 uint8 = 240
                                _ = v20
                            case 0:
                                var v21 int = fi18
                                _ = v21
                            default:
                                switch fi18 {
                                case 1:
                                    v22 := GenS2{F0: v5, F1: -740}
                                    _ = v22
                                }
                            }
                            _ = fi18
                            _ = v19
                        }
                        _ = fi13
                    }
                default:
                    v4 = v4
                }
            }
        }
        v5 = v5
        _ = v10
    } else {
        var v23 float64 = 0.0
        fmt.Println(v23)
        _ = v23
    }
    fmt.Println(v5, v1)
    var v24 bool = ((v4 & v4) >= v4)
    _ = len(v2)
    if true {
        var v25 float32 = -2.25
        switch (("foo bar" + ("\n\t" + "hello")) + "foo bar") {
        case "":
            var v26 float64 = -2.25
            _ = v26
        case "hello":
            switch "" {
            case "a":
                var v27 int16 = v5
                _ = v27
            case "\n\t":
                v24 = (v5 < -32768)
            default:
                for fi28 := 0; fi28 < 2; fi28++ {
                    fmt.Println(fi28, v4)
                    v29 := v2[(fi28 << ((488) & 63))]
                    v4 = v4
                    _ = fi28
                    _ = v29
                }
            }
        case "a":
            v1 = v1
        }
        _ = len(v3)
        _ = v25
    }
    v5 = (v5 + v5)
    v24 = (!v24)
    v30 := GenS3{F0: -487, F1: 0}
    if v24 {
        for fi31 := 0; fi31 < 10; fi31++ {
            if ((v5 * (v5 / (((v5 / (((^(v5))) | 1))) | 1))) >= (^(v5))) {
                var v32 string = "hello"
                for fi33 := 0; fi33 < 9; fi33++ {
                    fmt.Println(fi33)
                    v34 := GenS3{F0: 9223372036854775807, F1: 0}
                    _ = fi33
                    _ = v34
                }
                v35 := GenS2{F0: 717, F1: (v4 | v4)}
                _ = v32
                _ = v35
            }
            _ = fi31
        }
    }
    v4 = (v4 ^ v4)
    if (((("" + ("hello" + "")) + ("foo bar" + (("" + "") + ("foo bar" + "")))) + "hello") != (("a" + ("a" + ("" + "\n\t"))) + "foo bar")) {
        for fi36 := 0; fi36 < 15; fi36++ {
            v5 += (v5 / ((32767) | 1))
            var v37 bool = ((0 < 32767) || v24)
            _ = fi36
            _ = v37
        }
    } else {
        var v38 int = -1
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
