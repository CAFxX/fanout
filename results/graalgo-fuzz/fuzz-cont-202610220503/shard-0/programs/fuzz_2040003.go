package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 uint64 }
func genFn0() int {
    var v4 int16 = -945
    fmt.Println(v4)
    _ = v4
    return -318
}

func main() {
    var v1 int8 = 100
    var v2 int = 277
    v3 := []int{}
    v2 += 32768
    var v5 uint32 = 4294967295
    fmt.Println(v2, v2, v2)
    v6 := GenS1{F0: 127, F1: genFn0(), F2: 42}
    switch 1 {
    case 127:
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS2{F0: 1, F1: 42}
            if (!false) {
                var v9 float32 = ((0.1 / 1.0) / (3.14159 + 1.5))
                if false {
                    v1 = v1
                    var v10 uint = 42
                    _ = v10
                } else {
                    var v11 string = ("" + "")
                    _ = v11
                }
                for fi12 := 0; fi12 < 13; fi12++ {
                    var v13 int64 = int64(v5)
                    _ = fi12
                    _ = v13
                }
                _ = v9
            }
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 uint16 = 2
                v2 -= -1
                v16 := GenS2{F0: 42, F1: 2}
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi7
            _ = v8
        }
    default:
        switch "a" {
        case "a":
            switch "foo bar" {
            case "\n\t":
                v5 += v5
            case "a":
                if true {
                    v17 := GenS1{F0: 148, F1: (v2 << ((255) & 63)), F2: 1}
                    _ = v17
                } else {
                    v2 -= (^(v2))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi18 := 0; fi18 < 19; fi18++ {
                        v1 -= v1
                        v1 -= v1
                        var v19 uint64 = 42
                        _ = fi18
                        _ = v19
                    }
                }
            }
        }
    }
    if false {
        switch genFn0() {
        case 1:
            var v20 string = "hello"
            _ = v20
        case 32768:
            v21 := GenS3{F0: v5}
            _ = v21
        case -686:
            v22 := GenS4{F0: -9223372036854775808, F1: 470}
            _ = v22
        default:
            _ = len(v3)
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            if (1 == 605) {
                var v24 uint = 18446744073709551615
                var v25 uint64 = 1
                for fi26 := 0; fi26 < 8; fi26++ {
                    _ = len(v3)
                    v3 = append(v3, (v2 | genFn0()))
                    v24 = (v24 % ((uint(v25)) | 1))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi26
                }
                _ = v24
                _ = v25
            } else {
                for fi27 := 0; fi27 < 8; fi27++ {
                    v5 += v5
                    _ = fi27
                }
                var v28 bool = (v5 != (^(v5)))
                v5 += v5
                _ = v28
            }
            if true {
                var v29 string = ("hello" + "\n\t")
                _ = v29
            } else {
                var v30 uint32 = 2147483648
                _ = v30
            }
            if (42 == 128) {
                fmt.Println(v5, v2, fi23)
                if true {
                    _ = len(v3)
                    fmt.Println(v1, v1)
                }
                v1 = v1
            } else {
                fmt.Println(v1, v1)
                v5 += (v5 - 2)
            }
            _ = fi23
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v31 int8 = 1
    _ = len(v3)
    var v32 uint16 = 2
    fmt.Println(v31, v31, v32)
    var v33 uint32 = (v5 * (^(v5)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v5
    _ = v6
}
