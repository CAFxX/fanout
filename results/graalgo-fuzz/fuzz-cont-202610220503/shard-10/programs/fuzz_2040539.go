package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 int8 }

func main() {
    var v1 bool = false
    var v2 float64 = 3.14159
    var v3 int64 = 9223372036854775807
    var v4 string = (("" + "a") + ("foo bar" + ("" + "foo bar")))
    var v5 float32 = 1e10
    fmt.Println(v5, v4, v3)
    v6 := GenS1{F0: 2147483648, F1: v1, F2: v3}
    v3 += 420
    v7 := []int{127, 4294967295}
    v4 = ("hello" + (("hello" + v4) + v4))
    v2 = (v2 / v2)
    v8 := GenS2{F0: 0, F1: 127}
    fmt.Println(v3, v5)
    v9 := GenS3{F0: (v3 << ((65535) & 63))}
    var v10 uint32 = 65535
    fmt.Println(v2)
    switch "" {
    case "a":
        for fi11 := 0; fi11 < 8; fi11++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v12 float64 = v2
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 string = (v4 + ((v4 + ((v4 + (((v4 + "hello") + "") + v4)) + ("" + v4))) + ""))
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    case "\n\t":
        fmt.Println(v2)
    case "hello":
        v7 = append(v7, 808)
    default:
        switch v4 {
        case "a":
            switch 127 {
            case -534:
                for fi15 := 0; fi15 < 1; fi15++ {
                    v3 += (v3 >> ((4294967295) & 63))
                    fmt.Println(v3, fi15)
                    v16 := GenS4{F0: 42, F1: 0}
                    _ = fi15
                    _ = v16
                }
            case -1:
                v4 = v4
            case 617:
                if len(v7) > 0 {
                    v7[0] = 127
                }
            default:
                if (false && (v1 || (v3 != v3))) {
                    var v17 bool = true
                    _ = v17
                } else {
                    v5 = v5
                    v1 = v1
                    v18 := GenS5{F0: 1}
                    _ = v18
                }
            }
        }
    }
    switch (v4 + ("\n\t" + v4)) {
    case "\n\t":
        for fi19 := 0; fi19 < 11; fi19++ {
            v1 = ((v1 || v1) || true)
            _ = fi19
        }
    case "a":
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 int64 = v3
            v22 := GenS1{F0: 927, F1: (674 >= 1), F2: v21}
            if true {
                if (!(604 >= 42)) {
                    for fi23 := 0; fi23 < 11; fi23++ {
                        _ = len(v7)
                        _ = fi23
                    }
                    fmt.Println(v1, v21, v2)
                }
                if ((v10 - (v10 - 1)) > 655) {
                    for fi24 := 0; fi24 < 11; fi24++ {
                        if len(v7) > 0 {
                            v7[0] = 2147483647
                        }
                        fmt.Println(v10)
                        _ = fi24
                    }
                    v4 = (v4 + v4)
                }
                for fi25 := 0; fi25 < 8; fi25++ {
                    var v26 uint8 = 226
                    v10 = (v10 - v10)
                    v2 -= v2
                    _ = fi25
                    _ = v26
                }
            }
            _ = fi20
            _ = v21
            _ = v22
        }
    case "foo bar":
        v27 := GenS6{F0: (36 == 1), F1: 1}
        _ = v27
    default:
        v3 -= (v3 * (v3 | v3))
    }
    v4 = "a"
    v28 := GenS5{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
