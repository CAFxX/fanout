package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 string; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint }

func main() {
    var v1 bool = true
    var v2 int32 = 163
    var v3 int64 = -1
    var v4 int32 = (v2 << ((18446744073709551615) & 63))
    var v5 uint = 18446744073709551615
    if (int(v2) > -1) {
        var v6 uint32 = 0
        _ = v6
    } else {
        if v1 {
            v5 += 525
        }
        var v7 uint16 = 0
        v8 := GenS1{F0: true}
        _ = v7
        _ = v8
    }
    v9 := []int{}
    var v10 uint = (v5 & (v5 - v5))
    if v1 {
        var v11 uint16 = 32767
        _ = v11
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        v9 = append(v9, fi12)
        v4 = v2
        _ = fi12
    }
    if ((("" + ("a" + "\n\t")) + "foo bar") <= ((("" + "foo bar") + "a") + "foo bar")) {
        if (128 >= 1) {
            var v13 int8 = -1
            var v14 int8 = v13
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: (-663 >= -600)}
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 uint32 = 555
            for fi18 := 0; fi18 < 3; fi18++ {
                if ((v3 & -9223372036854775808) > v3) {
                    v3 += (v3 | (-(v3)))
                    v5 += 652
                    if v1 {
                        v2 += (v4 | v4)
                        v5 += v10
                        if false {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        }
                    } else {
                        v19 := GenS2{F0: ("\n\t" + ((("foo bar" + "hello") + "") + ("" + "foo bar"))), F1: ("a" + "hello"), F2: -32768}
                        var v20 int8 = -128
                        v3 += (v3 | v3)
                        _ = v19
                        _ = v20
                    }
                    if (!("" != ("hello" + ("" + ("\n\t" + ("\n\t" + "\n\t")))))) {
                        var v21 bool = ((fi16 | 1) != (-(fi18)))
                        _ = v21
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            if v1 {
                                if len(v9) > 0 {
                                    v9[0] = (fi16 - fi16)
                                }
                                var v22 float64 = 3.14159
                                _ = v22
                            } else {
                                fmt.Println(v17, v17)
                                fmt.Println(v10, fi16, v3)
                                if v1 {
                                    if (!(!(!(false && v1)))) {
                                        var v23 int16 = -32768
                                        _ = v23
                                    } else {
                                        v24 := GenS3{F0: 18446744073709551615, F1: 42}
                                        var v25 uint64 = 928
                                        v25 = (v25 >> (((^(v10))) & 63))
                                        _ = v24
                                        _ = v25
                                    }
                                    fmt.Println(v2, fi18)
                                }
                            }
                        case "hello":
                            v26 := GenS3{F0: 481, F1: 74}
                            _ = v26
                        default:
                            var v27 int = 4294967295
                            _ = v27
                        }
                        switch ("hello" + "") {
                        case "\n\t":
                            switch ("" + ((("foo bar" + "hello") + "foo bar") + "foo bar")) {
                            case "a":
                                v1 = (48 == 62)
                            case "foo bar":
                                var v28 bool = (true && (v17 < 2147483647))
                                fmt.Println(v4, v3, v3)
                                _ = v28
                            }
                        }
                    }
                }
                v29 := GenS4{F0: v17, F1: 1, F2: (v5 ^ v5)}
                v30 := GenS1{F0: (!(fi18 <= (fi18 + fi18)))}
                _ = fi18
                _ = v29
                _ = v30
            }
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        var v31 int64 = (v3 & 73)
        _ = v31
    }
    var v32 uint32 = 65535
    fmt.Println(v2, v3)
    var v33 bool = v1
    v5 -= v5
    if ((v10 * v10) == (v5 >> ((479) & 63))) {
        v2 += (^(v4))
    } else {
        v9 = append(v9, 0)
    }
    v10 = 1
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v9
}
