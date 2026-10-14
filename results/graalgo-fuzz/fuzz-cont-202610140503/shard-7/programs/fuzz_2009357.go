package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint64 }

func main() {
    var v1 bool = (float32(0.1) <= (0.0 + float32(1.5)))
    v2 := []int{37, 32767, 791, 42}
    v3 := []string{"a"}
    var v4 int32 = 2147483647
    var v5 uint8 = 127
    switch 2147483647 {
    case 9223372036854775807:
        v6 := GenS1{F0: (((("\n\t" + "") + "hello") + ((("" + "\n\t") + "") + "")) + "foo bar"), F1: 42}
        _ = v6
    case 128:
        switch ("a" + ("foo bar" + ("" + "foo bar"))) {
        case "\n\t":
            fmt.Println(v4, v4)
        case "foo bar":
            if v1 {
                v4 -= (^(v4))
            } else {
                fmt.Println(v5, v4)
                v7 := GenS2{F0: 42, F1: 1000}
                v8 := GenS2{F0: 2, F1: 637}
                _ = v7
                _ = v8
            }
        case "a":
            var v9 uint16 = 1
            _ = v9
        }
    default:
        v10 := GenS1{F0: "a", F1: 1}
        _ = v10
    }
    if len(v2) > 0 {
        v2[0] = int(v5)
    }
    switch 1 {
    case 42:
        v1 = false
    default:
        v1 = ((((100 >= 100) || ((0 > 42) || false)) && v1) && true)
    }
    switch 127 {
    case -1:
        var v11 uint64 = 4294967295
        _ = v11
    }
    switch "a" {
    case "foo bar":
        var v12 uint64 = 18446744073709551615
        _ = v12
    default:
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 uint32 = 65535
            _ = fi13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        if len(v2) > 0 {
            v2[0] = fi15
        }
        _ = fi15
    }
    var v16 int = -9223372036854775808
    v4 -= v4
    var v17 int16 = 0
    v17 = 1000
    var v18 int8 = -1
    if v1 {
        fmt.Println(v16)
        var v19 int32 = (-(v4))
        _ = v19
    } else {
        if v1 {
            var v20 float64 = 1.5
            v3 = append(v3, ("a" + ("" + ("a" + "a"))))
            _ = v20
        }
        v21 := GenS3{F0: 128, F1: -1, F2: 18446744073709551615}
        _ = v21
    }
    v22 := GenS1{F0: "a", F1: 352}
    v2 = append(v2, 0)
    if (v16 == (-(v16))) {
        var v23 uint = 2
        switch v16 {
        case 4294967295:
            switch v16 {
            case 128:
                v18 -= v18
            case 9223372036854775807:
                if ((v23 / ((v23) | 1)) > (v23 << (((v23 + v23)) & 63))) {
                    if ((1 < 42) || (2 <= 2147483647)) {
                        v18 -= -1
                    } else {
                        if v1 {
                            var v24 int64 = 9223372036854775807
                            var v25 int16 = (v17 >> (((^(v23))) & 63))
                            v26 := GenS3{F0: 2, F1: 0, F2: 42}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v16 = v16
                    }
                    var v27 int16 = v17
                    _ = v27
                } else {
                    fmt.Println(v23, v1, v23)
                    var v28 bool = (true && v1)
                    _ = v28
                }
            }
        default:
            if false {
                fmt.Println(v17)
                var v29 uint8 = (v5 / (((v5 & (v5 & (v5 | 117)))) | 1))
                _ = v29
            } else {
                if true {
                    for fi30 := 0; fi30 < 7; fi30++ {
                        var v31 uint = uint(v5)
                        _ = fi30
                        _ = v31
                    }
                    var v32 uint8 = (v5 & v5)
                    _ = v32
                } else {
                    var v33 uint64 = 2
                    _ = v33
                }
            }
        }
        _ = v23
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v5)
    }
    var v34 bool = false
    var v35 int = v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
