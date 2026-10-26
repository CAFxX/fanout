package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int64 }
type GenS2 struct { F0 uint; F1 uint; F2 int64 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint; F1 uint64; F2 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = -35
    var v3 uint = 995
    var v4 int64 = 1
    var v5 uint = 59
    v6 := []int{}
    v4 -= v4
    var v7 float64 = float64(0.0)
    var v8 int8 = v2
    v9 := GenS1{F0: (v4 & v4), F1: ("\n\t" + "hello"), F2: v4}
    v8 = (v2 * v2)
    v1[32768] = false
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v10 := GenS1{F0: v4, F1: "foo bar", F2: (v4 << ((582) & 63))}
    var v11 float32 = (float32(float64((3.14159 * ((float32(((1e10 * 0.0) - 0.1)) - -2.25) * 0.0)))) - 1.5)
    v12 := GenS2{F0: 18446744073709551615, F1: 131, F2: 42}
    switch ("hello" + "\n\t") {
    case "a":
        var v13 int64 = v4
        _ = v13
    case "hello":
        v7 = v7
    case "\n\t":
        v11 -= 0.0
    default:
        v11 = 1.5
    }
    var v14 int64 = (v4 | v4)
    for fi15 := 0; fi15 < 4; fi15++ {
        fmt.Println(v11)
        v2 += v8
        _ = fi15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        if (11 > (v8 | v8)) {
            var v17 int16 = -670
            switch fi16 {
            case -128:
                v18 := GenS3{F0: "foo bar", F1: 571}
                _ = v18
            case 4294967295:
                v4 -= v4
            }
            switch ((("hello" + "\n\t") + "hello") + "\n\t") {
            case "hello":
                switch "" {
                case "":
                    for fi19 := 0; fi19 < 13; fi19++ {
                        fmt.Println(v3)
                        _ = fi19
                    }
                case "hello":
                    switch "" {
                    case "hello":
                        switch fi16 {
                        case 2147483647:
                            var v20 int8 = v8
                            _ = v20
                        }
                    case "\n\t":
                        v8 += (v2 * (-(v8)))
                    case "foo bar":
                        v11 = 3.14159
                    }
                }
            case "":
                v14 -= -9223372036854775808
            }
            _ = v17
        } else {
            if true {
                if true {
                    var v21 uint64 = 808
                    _ = v21
                }
            }
        }
        _ = fi16
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        if true {
            _ = len(v1)
            v5 = 65535
            switch "hello" {
            case "a":
                fmt.Println(v14)
            case "\n\t":
                for fi23 := 0; fi23 < 5; fi23++ {
                    switch "\n\t" {
                    case "\n\t":
                        if false {
                            var v24 uint16 = 129
                            switch "foo bar" {
                            case "":
                                var v25 uint8 = 0
                                _ = v25
                            case "\n\t":
                                v24 -= (v24 & v24)
                            case "foo bar":
                                v26 := GenS4{F0: 0, F1: -1}
                                _ = v26
                            default:
                                v27 := GenS3{F0: "foo bar", F1: 2}
                                fmt.Println(v5, v5)
                                _ = v27
                            }
                            v7 = -2.25
                            _ = v24
                        }
                    case "hello":
                        v28 := GenS5{F0: 1, F1: -32768}
                        _ = v28
                    default:
                        fmt.Println(v2, v8)
                    }
                    if false {
                        v4 += v4
                        v29 := GenS6{F0: 243}
                        _ = v29
                    } else {
                        v8 -= (v8 % ((42) | 1))
                        if ((!false) && (uint8(v3) != 42)) {
                            v30 := GenS4{F0: v4, F1: 0}
                            _ = v30
                        }
                        v3 -= 250
                    }
                    v31 := GenS1{F0: v14, F1: (("hello" + "\n\t") + "a"), F2: (v4 + (v4 ^ 42))}
                    _ = fi23
                    _ = v31
                }
            default:
                _ = len(v1)
            }
        }
        v32 := GenS7{F0: 1, F1: 0, F2: 42}
        _ = fi22
        _ = v32
    }
    for fi33 := 0; fi33 < 9; fi33++ {
        v6 = append(v6, (fi33 % (((fi33 / ((fi33) | 1))) | 1)))
        _ = fi33
    }
    var v34 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
