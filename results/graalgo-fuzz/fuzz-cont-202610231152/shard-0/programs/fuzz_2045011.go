package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int16 }

func main() {
    var v1 int8 = 1
    var v2 uint32 = 544
    var v3 uint16 = 816
    var v4 uint64 = 571
    var v5 string = ("a" + "hello")
    var v6 bool = false
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(fi7, v2, v6)
        _ = fi7
    }
    var v8 int = 0
    v9 := GenS1{F0: (v4 ^ v4), F1: "", F2: v1}
    switch (v8 - (^(v8))) {
    case -1:
        if ((v3 - v3) < v3) {
            fmt.Println(v5)
        }
    case 42:
        var v10 uint8 = 2
        _ = v10
    }
    switch (v5 + v5) {
    case "hello":
        var v11 int32 = -2147483648
        _ = v11
    case "foo bar":
        var v12 int16 = 1
        _ = v12
    case "a":
        v13 := GenS1{F0: (^(v4)), F1: v5, F2: 74}
        _ = v13
    }
    v6 = (127 > 2)
    var v14 int = (v8 ^ (-(v8)))
    for fi15 := 0; fi15 < 12; fi15++ {
        if (v5 <= v5) {
            fmt.Println(v8)
            v16 := []int{v14, v14, (fi15 - (v14 >> ((962) & 63)))}
            _ = v16
        } else {
            v4 -= v4
            fmt.Println(v4, fi15)
            var v17 string = ("\n\t" + "foo bar")
            _ = v17
        }
        for fi18 := 0; fi18 < 11; fi18++ {
            v19 := []int{(v8 ^ 9223372036854775807)}
            if ((!((0.1 == 1e10) || ((1.5 / 1e10) == float32((float64(1e10) + 1.5))))) && v6) {
                v2 = 42
                switch -933 {
                case -1:
                    var v20 int8 = (v1 & (v1 | v1))
                    _ = v20
                case -2147483648:
                    fmt.Println(v8, v14)
                }
                if v6 {
                    v14 = v14
                    var v21 int64 = -9223372036854775808
                    var v22 bool = (v6 && (v6 && (!false)))
                    fmt.Println(v21)
                    fmt.Println(fi18, v4)
                    var v23 bool = v22
                    _ = v21
                    _ = v22
                    _ = v23
                } else {
                    switch "" {
                    case "hello":
                        if true {
                            v19 = append(v19, (v14 * fi18))
                            _ = len(v19)
                            var v24 uint8 = 1
                            _ = v24
                        } else {
                            if v6 {
                                fmt.Println(fi18, v4, v4)
                            }
                            var v25 int32 = -1
                            var v26 int32 = v25
                            _ = v25
                            _ = v26
                        }
                    default:
                        var v27 int = fi15
                        _ = v27
                    }
                    var v28 int32 = -2147483648
                    _ = v28
                }
                var v29 int64 = -802
                _ = v29
            } else {
                if (!false) {
                    var v30 uint8 = uint8(v3)
                    _ = v30
                } else {
                    var v31 int = (v14 * (^(fi18)))
                    switch 4294967295 {
                    case -1:
                        var v32 uint64 = (v4 & (v4 >> ((18446744073709551615) & 63)))
                        _ = v32
                    case 529:
                        v33 := GenS1{F0: uint64(v3), F1: v5, F2: 1}
                        if (0 >= (v2 | v2)) {
                            _ = len(v19)
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                        } else {
                            v1 = (-(v1))
                            var v34 bool = false
                            _ = v34
                        }
                        v2 = (v2 ^ (v2 % ((v2) | 1)))
                        _ = v33
                    default:
                        switch v5 {
                        case "\n\t":
                            v35 := GenS2{F0: -2147483648}
                            _ = v35
                        }
                    }
                    _ = v31
                }
            }
            v36 := GenS3{F0: v5, F1: 1}
            _ = fi18
            _ = v19
            _ = v36
        }
        var v37 float32 = ((float32((1e10 / float64(1.5))) * (0.0 / -2.25)) + 3.14159)
        _ = fi15
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
