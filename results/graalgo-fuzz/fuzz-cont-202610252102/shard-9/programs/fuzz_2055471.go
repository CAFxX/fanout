package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int; F2 string }
type GenS3 struct { F0 uint8; F1 uint16; F2 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 string; F2 uint }

func main() {
    var v1 int = 607
    var v2 int32 = -1
    var v3 bool = (18446744073709551615 < 902)
    var v4 float64 = 1e10
    var v5 int8 = -1
    v6 := []int{(v1 ^ (v1 - (v1 + (^(v1)))))}
    var v7 int16 = -1
    switch (v1 / ((-9223372036854775808) | 1)) {
    case 734:
        if (!((v7 + (^(v7))) > (v7 & v7))) {
            var v8 string = "hello"
            _ = v8
        } else {
            fmt.Println(v7)
        }
    }
    if len(v6) > 0 {
        v6[0] = 2147483647
    }
    fmt.Println(v5)
    switch ("hello" + "hello") {
    case "hello":
        v9 := GenS1{F0: 270}
        _ = v9
    case "a":
        var v10 float64 = v4
        _ = v10
    default:
        if v3 {
            for fi11 := 0; fi11 < 15; fi11++ {
                switch (v1 >> ((1) & 63)) {
                case 32768:
                    if (!v3) {
                        fmt.Println(v4, v7, v4)
                    } else {
                        for fi12 := 0; fi12 < 18; fi12++ {
                            var v13 float64 = 1e10
                            fmt.Println(v1)
                            _ = fi12
                            _ = v13
                        }
                        v4 -= v4
                        var v14 int64 = 0
                        _ = v14
                    }
                case 4294967295:
                    v15 := GenS1{F0: 0}
                    _ = v15
                case -1:
                    v6 = append(v6, (^(fi11)))
                }
                v7 += v7
                _ = fi11
            }
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 bool = v3
                _ = fi16
                _ = v17
            }
            _ = len(v6)
        }
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v2, v3)
    case "foo bar":
        if true {
            var v18 int32 = (v2 & -1)
            for fi19 := 0; fi19 < 1; fi19++ {
                fmt.Println(v18)
                _ = fi19
            }
            v20 := GenS2{F0: (v3 && true), F1: (v1 >> ((42) & 63)), F2: ""}
            _ = v18
            _ = v20
        } else {
            var v21 uint64 = 885
            switch "foo bar" {
            case "\n\t":
                switch (^(v1)) {
                case -273:
                    for fi22 := 0; fi22 < 8; fi22++ {
                        var v23 uint64 = 4294967295
                        _ = fi22
                        _ = v23
                    }
                }
            case "":
                v24 := GenS2{F0: (((0.1 / float64(float32(float64((-2.25 * 3.14159))))) - v4) >= float64(1e10)), F1: 32767, F2: "hello"}
                _ = v24
            case "a":
                if (false && (false || false)) {
                    var v25 uint64 = v21
                    v26 := GenS3{F0: 128, F1: 268, F2: (v1 & 42)}
                    var v27 uint8 = 2
                    _ = v25
                    _ = v26
                    _ = v27
                }
            }
            for fi28 := 0; fi28 < 11; fi28++ {
                if false {
                    v6 = append(v6, (fi28 >> ((42) & 63)))
                } else {
                    switch "hello" {
                    case "hello":
                        var v29 bool = ("a" < "foo bar")
                        _ = v29
                    case "a":
                        if false {
                            v30 := GenS1{F0: 32767}
                            for fi31 := 0; fi31 < 13; fi31++ {
                                v4 = float64(0.0)
                                _ = fi31
                            }
                            _ = v30
                        }
                    default:
                        v32 := GenS4{F0: 825}
                        _ = v32
                    }
                    fmt.Println(v1, v2, v7)
                }
                v33 := GenS5{F0: 9223372036854775807, F1: ("foo bar" + "a"), F2: 1}
                switch (fi28 % ((-9223372036854775808) | 1)) {
                case -9223372036854775808:
                    var v34 uint64 = 18446744073709551615
                    _ = v34
                default:
                    var v35 uint32 = 65535
                    _ = v35
                }
                _ = fi28
                _ = v33
            }
            _ = v21
        }
    case "\n\t":
        v2 = (v2 + (v2 ^ v2))
    }
    _ = len(v6)
    if len(v6) > 0 {
        v6[0] = v1
    }
    var v36 uint8 = 0
    var v37 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
