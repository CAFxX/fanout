package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int16; F1 int8; F2 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 uint16; F2 int8 }

func main() {
    var v1 int32 = 42
    var v2 uint = 520
    var v3 int64 = -1
    v4 := make(map[uint32]string)
    var v5 int64 = (^(v3))
    switch "\n\t" {
    case "\n\t":
        v6 := GenS1{F0: (v5 | (v3 % (((v3 | -9223372036854775808)) | 1)))}
        _ = v6
    case "hello":
        var v7 int64 = 9223372036854775807
        _ = v7
    case "":
        switch ("\n\t" + "foo bar") {
        case "":
            v1 += (-(v1))
        case "a":
            switch (("a" + ("\n\t" + "a")) + ("foo bar" + "foo bar")) {
            case "a":
                var v8 float64 = 1.5
                _ = v8
            case "foo bar":
                fmt.Println(v3, v2)
            }
        case "\n\t":
            v9 := GenS2{F0: -193, F1: true}
            _ = v9
        }
    default:
        v2 = (v2 ^ (v2 % ((2) | 1)))
    }
    var v10 float64 = 1e10
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            v2 += v2
            _ = fi12
        }
        _ = fi11
    }
    if true {
        if (v2 != 615) {
            v13 := GenS3{F0: 2147483647, F1: 32767}
            v5 -= v3
            v5 -= int64(v2)
            _ = v13
        } else {
            for fi14 := 0; fi14 < 6; fi14++ {
                if (true || (true && true)) {
                    var v15 string = "\n\t"
                    _ = v15
                } else {
                    var v16 bool = (true && ((!true) && false))
                    if (v16 && true) {
                        var v17 int8 = -120
                        fmt.Println(v5, v16, v17)
                        var v18 uint16 = 42
                        _ = v17
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 8; fi19++ {
                            v20 := []int{(fi14 + (^(fi14))), fi14, (^(fi19)), (fi14 * 4294967295)}
                            _ = fi19
                            _ = v20
                        }
                        v5 = 1
                    }
                    if (v1 != 2147483647) {
                        v21 := GenS3{F0: 9223372036854775807, F1: fi14}
                        for fi22 := 0; fi22 < 3; fi22++ {
                            fmt.Println(v3, fi22)
                            var v23 string = "a"
                            var v24 uint8 = 127
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        _ = v21
                    }
                    _ = v16
                }
                if (255 < 65535) {
                    v25 := GenS4{F0: -32768, F1: 0, F2: v1}
                    _ = v25
                } else {
                    v26 := GenS1{F0: v5}
                    _ = v26
                }
                v27 := v4[194]
                _ = fi14
                _ = v27
            }
            v10 = (v10 / v10)
            for fi28 := 0; fi28 < 17; fi28++ {
                v1 -= (v1 * v1)
                _ = fi28
            }
        }
        fmt.Println(v3, v2, v3)
        v29 := v4[42]
        _ = v29
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        v31 := GenS4{F0: 1, F1: -1, F2: v1}
        if false {
            v32 := GenS2{F0: 128, F1: (false || (v10 == 1.5))}
            fmt.Println(v1)
            _ = v32
        }
        v33 := GenS2{F0: 2147483647, F1: true}
        _ = fi30
        _ = v31
        _ = v33
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        fmt.Println(v5, v10)
        for fi35 := 0; fi35 < 19; fi35++ {
            v36 := []int{(fi34 + (fi34 | 804)), fi34, 1, fi34}
            v2 = (v2 << ((v2) & 63))
            var v37 uint64 = 362
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi34
    }
    switch "\n\t" {
    case "":
        switch ("" + ("a" + "hello")) {
        case "\n\t":
            fmt.Println(v3)
        case "":
            v38 := GenS5{F0: true}
            _ = v38
        default:
            for fi39 := 0; fi39 < 3; fi39++ {
                v10 -= ((-2.25 + ((1e10 * -2.25) + (((1.5 - v10) - (v10 + v10)) * (float64((3.14159 * 1e10)) + 1e10)))) * v10)
                _ = fi39
            }
        }
    default:
        v40 := []int{}
        _ = v40
    }
    switch 9223372036854775807 {
    case 228:
        switch 143 {
        case -9223372036854775808:
            v10 -= v10
        case 127:
            var v41 int = 4294967295
            _ = v41
        }
    case 0:
        v42 := []int{9223372036854775807}
        _ = v42
    case 9223372036854775807:
        var v43 int64 = (v5 % ((1) | 1))
        _ = v43
    }
    v44 := []int{-510, -9223372036854775808}
    for fi45 := 0; fi45 < 13; fi45++ {
        if true {
            v46 := GenS6{F0: (fi45 & fi45), F1: 65535, F2: -61}
            fmt.Println(v3, v2, v1)
            _ = v46
        } else {
            v3 = v5
            v47 := GenS6{F0: fi45, F1: 42, F2: 42}
            _ = v47
        }
        _ = fi45
    }
    var v48 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v48
    _ = v5
}
