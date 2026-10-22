package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int16; F1 uint16 }

func main() {
    var v1 uint32 = 2147483648
    v2 := []string{("" + ("\n\t" + "foo bar")), "foo bar", "\n\t", "foo bar"}
    var v3 int = 521
    v4 := make(map[uint32]bool)
    var v5 uint16 = 1
    switch (("foo bar" + ("" + "\n\t")) + "foo bar") {
    case "a":
        if (1 < 2) {
            v3 -= 9223372036854775807
        } else {
            v4[(v1 >> ((4294967295) & 63))] = (-1 > -1)
        }
    case "hello":
        switch "hello" {
        case "a":
            var v6 uint8 = 127
            _ = v6
        case "hello":
            v1 = 2147483648
        case "\n\t":
            v7 := GenS1{F0: 4294967295}
            _ = v7
        }
    case "foo bar":
        v8 := GenS1{F0: 4294967295}
        _ = v8
    default:
        if (uint(v5) == 4294967295) {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            var v9 uint32 = 4294967295
            var v10 int32 = 821
            _ = v10
            _ = v9
        } else {
            switch (v3 | (v3 | 9223372036854775807)) {
            case 1:
                v11 := GenS2{F0: 1, F1: -42}
                _ = v11
            case 0:
                if (false && (42 > 642)) {
                    v2 = append(v2, "")
                } else {
                    v5 = (v5 | (v5 & v5))
                    switch "hello" {
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    v3 = (^(v3))
                }
            }
            switch "hello" {
            case "":
                for fi12 := 0; fi12 < 10; fi12++ {
                    v5 = (^(v5))
                    if true {
                        if false {
                            _ = len(v2)
                            if (42 != 42) {
                                switch v3 {
                                case 42:
                                    var v13 string = ("a" + "")
                                    _ = v13
                                }
                                var v14 uint32 = (^(v1))
                                _ = v14
                            } else {
                                v3 -= v3
                                v2 = append(v2, (("a" + ("hello" + ("" + "\n\t"))) + ("\n\t" + "hello")))
                            }
                        }
                    } else {
                        v5 = v5
                    }
                    for fi15 := 0; fi15 < 6; fi15++ {
                        var v16 uint = 42
                        var v17 float32 = 1.5
                        var v18 int32 = 0
                        fmt.Println(v1, fi15, v5)
                        var v19 float64 = 3.14159
                        v19 -= ((0.0 / -2.25) - v19)
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = fi12
                }
            case "foo bar":
                _ = len(v2)
            }
            var v20 float32 = 0.0
            _ = v20
        }
    }
    var v21 uint32 = (v1 % (((v1 * (v1 | (v1 & v1)))) | 1))
    _ = len(v4)
    switch (v3 & (-(v3))) {
    case 9223372036854775807:
        v21 -= 69
    case 381:
        switch v3 {
        case 9223372036854775807:
            var v22 uint16 = v5
            _ = v22
        case 0:
            var v23 int = v3
            _ = v23
        default:
            if (159 <= 2) {
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 string = "\n\t"
                    fmt.Println(v1)
                    if true {
                        var v26 int64 = int64(v1)
                        var v27 int64 = v26
                        var v28 bool = (-2147483648 != 42)
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    _ = fi24
                    _ = v25
                }
            } else {
                fmt.Println(v21, v3)
                for fi29 := 0; fi29 < 19; fi29++ {
                    v30 := GenS2{F0: 255, F1: -128}
                    var v31 int64 = -84
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
        }
    }
    var v32 uint16 = 255
    fmt.Println(v21, v3, v1)
    v32 += 0
    switch "" {
    case "\n\t":
        v21 = v21
    }
    v32 = 42
    v33 := GenS3{F0: 42, F1: (v5 / (((v32 >> ((42) & 63))) | 1))}
    v3 = v3
    var v34 float64 = 0.0
    v32 = v32
    var v35 uint16 = v32
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v36 int8 = 127
    for fi37 := 0; fi37 < 16; fi37++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi37
    }
    if (!((!((v36 | (v36 / (((v36 >> ((1) & 63))) | 1))) > v36)) && (v5 <= (v5 + 0)))) {
        v2 = append(v2, ("hello" + "a"))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
