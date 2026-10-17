package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint32, p1 int64) int64 {
    fmt.Println(p0, p1)
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint64 = 42
        var v9 uint8 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(p0, p1)
    p0 += (^(p0))
    var v10 int32 = 548
    _ = p0
    _ = p1
    _ = v10
    return (p1 - p1)
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int32 = 42
    var v3 int32 = v2
    var v4 float64 = -2.25
    var v5 int16 = 1
    v6 := make(map[string]string)
    v11 := GenS1{F0: 42, F1: 2}
    if false {
        v12 := []int{9223372036854775807, -223, 276}
        _ = v12
    } else {
        switch (("\n\t" + "a") + ("foo bar" + "hello")) {
        case "hello":
            fmt.Println(v4)
        }
        var v13 int = -1
        _ = v13
    }
    v14 := GenS1{F0: -1, F1: 2}
    fmt.Println(v4, v1)
    fmt.Println(v3, v5, v2)
    v2 += v3
    if (v1 >= 0) {
        v15 := GenS1{F0: 120, F1: 1}
        if false {
            if true {
                v4 = v4
            }
            switch "a" {
            case "hello":
                var v16 bool = (false || true)
                _ = v16
            case "":
                v17 := GenS1{F0: 1, F1: 207}
                _ = v17
            }
            switch ((((((("a" + "foo bar") + "foo bar") + "a") + "\n\t") + ("\n\t" + "a")) + "foo bar") + "a") {
            case "\n\t":
                for fi18 := 0; fi18 < 20; fi18++ {
                    fmt.Println(v1, v4, v5)
                    v5 = v5
                    _ = fi18
                }
            default:
                if (v3 > -524) {
                    switch "\n\t" {
                    case "a":
                        switch ("\n\t" + ("foo bar" + ("foo bar" + (("a" + "\n\t") + ("hello" + "a"))))) {
                        case "foo bar":
                            if false {
                                if (349 <= 385) {
                                    v19 := v6["foo bar"]
                                    var v20 uint32 = 2
                                    _ = v19
                                    _ = v20
                                } else {
                                    if (992 != 18446744073709551615) {
                                        for fi21 := 0; fi21 < 7; fi21++ {
                                            var v22 int = 115
                                            var v23 uint16 = 680
                                            _ = fi21
                                            _ = v22
                                            _ = v23
                                        }
                                        var v24 float32 = (0.1 * 1e10)
                                        _ = v24
                                    } else {
                                        for fi25 := 0; fi25 < 3; fi25++ {
                                            var v26 uint64 = 1
                                            var v27 int32 = v2
                                            _ = fi25
                                            _ = v26
                                            _ = v27
                                        }
                                    }
                                    v28 := GenS2{F0: 42}
                                    fmt.Println(v2, v5, v5)
                                    _ = v28
                                }
                            }
                        case "\n\t":
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v4, v5)
                            }
                        }
                    case "hello":
                        _ = len(v6)
                    default:
                        v29 := v6[("" + "\n\t")]
                        _ = v29
                    }
                } else {
                    v30 := GenS3{F0: (v3 / ((v3) | 1))}
                    var v31 uint32 = 381
                    v3 = 0
                    _ = v30
                    _ = v31
                }
            }
        } else {
            if (!true) {
                fmt.Println(v5, v3, v1)
            }
        }
        var v32 uint32 = 2147483648
        _ = v15
        _ = v32
    } else {
        fmt.Println(v3, v2, v5)
        switch "\n\t" {
        case "a":
            var v33 bool = true
            _ = v33
        }
        var v34 int16 = v5
        _ = v34
    }
    if (false || (-63 >= -128)) {
        v3 -= v2
        for fi35 := 0; fi35 < 20; fi35++ {
            v36 := GenS3{F0: (v3 + v3)}
            v37 := GenS3{F0: -1}
            var v38 int32 = (^(v2))
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
