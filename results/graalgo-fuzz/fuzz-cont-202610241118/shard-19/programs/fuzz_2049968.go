package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int64 }
type GenS2 struct { F0 int32; F1 int32; F2 int32 }
func genFn0() int64 {
    if false {
        var v4 bool = false
        v5 := GenS1{F0: 42, F1: 931, F2: 9223372036854775807}
        var v6 bool = (!v4)
        _ = v4
        _ = v5
        _ = v6
    }
    var v7 int16 = 0
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint16) uint8 {
    for fi8 := 0; fi8 < 5; fi8++ {
        p0 = p0
        v9 := []int{(fi8 * 127), fi8, (fi8 + fi8)}
        switch "" {
        case "\n\t":
            v10 := GenS1{F0: (p0 & (p0 * p0)), F1: (^(fi8)), F2: 42}
            _ = v10
        case "hello":
            fmt.Println(fi8)
        }
        _ = fi8
        _ = v9
    }
    p0 += p0
    _ = p0
    return 1
}

func main() {
    var v1 uint = 4294967295
    var v2 uint64 = 723
    var v3 float64 = (3.14159 / ((3.14159 / 1.0) - 0.0))
    var v11 int64 = 9223372036854775807
    v12 := []int{-90, 127}
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 string = "a"
        switch (^(fi13)) {
        case 127:
            for fi15 := 0; fi15 < 11; fi15++ {
                switch v14 {
                case "\n\t":
                    if false {
                        v14 = v14
                        v16 := GenS1{F0: 42, F1: fi15, F2: 0}
                        v17 := GenS1{F0: 2, F1: (fi13 | (fi15 % (((fi13 % (((fi15 - (fi15 % ((fi15) | 1)))) | 1))) | 1))), F2: genFn0()}
                        _ = v16
                        _ = v17
                    } else {
                        if (true && true) {
                            v11 = (^(v11))
                        } else {
                            var v18 uint8 = genFn1(985)
                            if true {
                                var v19 int = fi13
                                v2 -= v2
                                v11 = 1
                                _ = v19
                            } else {
                                if (!(65535 > 886)) {
                                    _ = len(v12)
                                }
                            }
                            _ = v18
                        }
                        switch fi13 {
                        case 1:
                            fmt.Println(v11, v2)
                        }
                    }
                case "a":
                    v3 = v3
                }
                v3 = 0.1
                _ = fi15
            }
        case 42:
            var v20 uint16 = 65535
            _ = v20
        }
        var v21 string = v14
        _ = fi13
        _ = v14
        _ = v21
    }
    _ = len(v12)
    if false {
        v3 -= v3
        switch ("\n\t" + "a") {
        case "\n\t":
            var v22 uint8 = 1
            _ = v22
        case "":
            switch "\n\t" {
            case "\n\t":
                v12 = append(v12, -458)
            case "hello":
                switch -144 {
                case 4294967295:
                    v23 := GenS2{F0: 1, F1: 2147483647, F2: -1}
                    _ = v23
                case 295:
                    switch (("a" + ("" + "")) + "hello") {
                    case "hello":
                        v24 := GenS1{F0: 323, F1: 32767, F2: (v11 << ((v1) & 63))}
                        _ = v24
                    default:
                        var v25 int8 = 111
                        _ = v25
                    }
                case 42:
                    fmt.Println(v2, v11)
                default:
                    if (false && false) {
                        v1 -= v1
                        if len(v12) > 0 {
                            v12[0] = -2147483648
                        }
                    }
                }
            }
        default:
            if len(v12) > 0 {
                v12[0] = 495
            }
        }
    }
    var v26 uint32 = 1
    if len(v12) > 0 {
        v12[0] = -9223372036854775808
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        var v28 int8 = 0
        v3 += v3
        for fi29 := 0; fi29 < 1; fi29++ {
            v12 = append(v12, (^(fi27)))
            var v30 uint = v1
            v26 -= (v26 - (^(v26)))
            _ = fi29
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    var v31 int = 42
    var v32 int16 = int16(v31)
    var v33 uint64 = v2
    v1 += v1
    var v34 uint = (v1 & 0)
    var v35 uint64 = 1
    if true {
        v2 = 970
    } else {
        v2 = v33
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        if len(v12) > 0 {
            v12[0] = -1
        }
        _ = fi36
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        v2 = 1
        v12 = append(v12, (fi37 / ((int(v26)) | 1)))
        _ = fi37
    }
    fmt.Println(v26, v26, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
