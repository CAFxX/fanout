package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int64; F1 uint32; F2 int32 }

func main() {
    var v1 uint = 682
    var v2 uint = (^(v1))
    var v3 float64 = 1e10
    var v4 int32 = -2147483648
    var v5 string = "\n\t"
    switch 808 {
    case -505:
        v5 = ("foo bar" + v5)
    }
    switch "hello" {
    case "":
        var v6 float64 = (v3 + -2.25)
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 float64 = v3
            var v9 float64 = float64(-2.25)
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v10 := []int{0, 42}
    var v11 uint = v2
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 uint16 = 65535
        if (!(42 < 155)) {
            v13 = 481
            switch "\n\t" {
            case "a":
                v14 := GenS1{F0: 4294967295}
                _ = v14
            case "":
                switch v5 {
                case "a":
                    switch v5 {
                    case "foo bar":
                        _ = len(v10)
                    case "a":
                        if false {
                            var v15 int16 = 32767
                            _ = v15
                        }
                    case "\n\t":
                        v16 := GenS2{F0: (^(v13)), F1: 42, F2: 42}
                        _ = v16
                    default:
                        v2 -= 0
                    }
                case "":
                    var v17 uint32 = 1
                    _ = v17
                case "foo bar":
                    switch (("\n\t" + v5) + v5) {
                    case "foo bar":
                        var v18 uint64 = 42
                        _ = v18
                    case "hello":
                        switch v5 {
                        case "foo bar":
                            fmt.Println(v1)
                        case "a":
                            fmt.Println(v11)
                        case "":
                            if (v1 != (^(v2))) {
                                var v19 int = fi12
                                v11 -= (v2 * v11)
                                _ = v19
                            }
                        }
                    case "a":
                        switch v5 {
                        case "a":
                            for fi20 := 0; fi20 < 9; fi20++ {
                                v21 := GenS1{F0: 317}
                                var v22 int64 = 42
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        case "foo bar":
                            v10 = append(v10, fi12)
                        case "hello":
                            if ((fi12 & fi12) != fi12) {
                                v23 := GenS3{F0: (v4 | -869), F1: -98}
                                for fi24 := 0; fi24 < 2; fi24++ {
                                    var v25 bool = false
                                    switch v5 {
                                    case "":
                                        var v26 uint64 = 4294967295
                                        _ = v26
                                    case "\n\t":
                                        fmt.Println(v3, v2, v1)
                                    case "a":
                                        var v27 int = fi24
                                        _ = v27
                                    }
                                    _ = fi24
                                    _ = v25
                                }
                                _ = v23
                            }
                        default:
                            if (false && true) {
                                if true {
                                    fmt.Println(fi12)
                                }
                            }
                        }
                    default:
                        var v28 bool = (int64(fi12) != 9223372036854775807)
                        _ = v28
                    }
                }
            }
        }
        _ = fi12
        _ = v13
    }
    switch 0 {
    case 4294967295:
        v29 := GenS4{F0: 561, F1: v2}
        _ = v29
    case 799:
        if ((255 == 1) || true) {
            v2 = (v1 >> ((v11) & 63))
            fmt.Println(v3, v4)
        }
    default:
        if (!false) {
            v4 = v4
        } else {
            var v30 uint64 = 965
            var v31 bool = (v5 > v5)
            fmt.Println(v2, v2)
            _ = v30
            _ = v31
        }
    }
    var v32 uint8 = 128
    fmt.Println(v3, v3)
    v4 += int32(v32)
    for fi33 := 0; fi33 < 1; fi33++ {
        var v34 uint = v2
        v35 := GenS5{F0: 9223372036854775807, F1: 42, F2: (v4 + 835)}
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
