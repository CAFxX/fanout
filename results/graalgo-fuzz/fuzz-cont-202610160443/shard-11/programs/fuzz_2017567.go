package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 string }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int32; F1 int8 }

func main() {
    var v1 uint64 = 581
    var v2 int = -9223372036854775808
    var v3 bool = false
    v4 := []int{}
    fmt.Println(v1, v2)
    if len(v4) > 0 {
        v4[0] = v2
    }
    v5 := GenS1{F0: 0, F1: -79, F2: ("hello" + ("\n\t" + (("foo bar" + ("foo bar" + "hello")) + ("\n\t" + ("" + ("" + ("hello" + "")))))))}
    if v3 {
        if true {
            v2 = v2
            v6 := GenS1{F0: 94, F1: -1, F2: ("hello" + "hello")}
            v2 -= 127
            _ = v6
        }
    } else {
        v1 += (^(v1))
    }
    var v7 int8 = 1
    for fi8 := 0; fi8 < 18; fi8++ {
        if false {
            v9 := GenS1{F0: 2147483647, F1: -661, F2: ("hello" + "")}
            _ = v9
        }
        if v3 {
            if (("hello" >= "hello") || ((v2 ^ -9223372036854775808) != (fi8 + 32767))) {
                var v10 uint16 = 348
                var v11 uint16 = 65535
                switch fi8 {
                case 0:
                    v3 = (v3 || ((-1 >= -32768) && (19 == 2147483648)))
                default:
                    v12 := GenS1{F0: 0, F1: 1, F2: "hello"}
                    _ = v12
                }
                _ = v10
                _ = v11
            }
            if v3 {
                var v13 float64 = -2.25
                v3 = (v13 >= (float64(0.0) - v13))
                _ = v13
            }
            v1 = v1
        } else {
            if true {
                if len(v4) > 0 {
                    v4[0] = (fi8 ^ (-(v2)))
                }
                v1 = (^(v1))
                v14 := GenS2{F0: fi8, F1: v1}
                _ = v14
            } else {
                for fi15 := 0; fi15 < 15; fi15++ {
                    v3 = (42 != 32767)
                    fmt.Println(v1, v3, v3)
                    _ = fi15
                }
                v16 := GenS1{F0: 144, F1: -1, F2: "a"}
                v2 += 2147483647
                _ = v16
            }
            for fi17 := 0; fi17 < 12; fi17++ {
                v18 := GenS2{F0: -9223372036854775808, F1: (v1 & (^(v1)))}
                var v19 uint8 = 2
                var v20 int32 = -989
                switch "" {
                case "\n\t":
                    var v21 int = 584
                    _ = v21
                case "foo bar":
                    v3 = (0 < 562)
                case "a":
                    v22 := GenS1{F0: 0, F1: (v20 + (v20 << ((1) & 63))), F2: "hello"}
                    _ = v22
                default:
                    switch "" {
                    case "a":
                        var v23 float64 = 0.0
                        _ = v23
                    case "hello":
                        var v24 float64 = (0.1 * -2.25)
                        _ = v24
                    case "\n\t":
                        if (912 <= 1) {
                            var v25 int32 = v20
                            var v26 float32 = 1e10
                            v3 = (2147483647 != 954)
                            _ = v25
                            _ = v26
                        }
                    default:
                        var v27 int = v2
                        _ = v27
                    }
                }
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi8
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        v3 = true
        _ = fi28
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if v3 {
        var v29 bool = v3
        fmt.Println(v7, v7)
        _ = v29
    } else {
        var v30 bool = true
        v4 = append(v4, v2)
        _ = v30
    }
    var v31 bool = false
    if ((0 != 980) || (-1 == -448)) {
        v4 = append(v4, (-(v2)))
    }
    v32 := GenS1{F0: 0, F1: 0, F2: "\n\t"}
    v33 := GenS3{F0: int32(v1), F1: (v7 / ((127) | 1))}
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    var v34 uint = 18446744073709551615
    fmt.Println(v1, v34)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v7
}
