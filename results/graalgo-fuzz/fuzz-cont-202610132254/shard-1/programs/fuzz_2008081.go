package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 string }
type GenS3 struct { F0 int64; F1 int32 }

func main() {
    var v1 uint64 = 601
    var v2 uint32 = 4294967295
    var v3 string = "hello"
    var v4 int64 = 42
    var v5 int16 = -498
    var v6 int32 = 42
    v7 := []int{127, 32767, -1, 558}
    if false {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint8 = 0
            _ = fi8
            _ = v9
        }
    } else {
        v10 := GenS1{F0: 0}
        switch v3 {
        case "foo bar":
            if (false || (false || false)) {
                for fi11 := 0; fi11 < 13; fi11++ {
                    v7 = append(v7, (fi11 >> ((4294967295) & 63)))
                    _ = fi11
                }
                fmt.Println(v2, v3, v5)
                v2 = (v2 * v2)
            } else {
                if true {
                    var v12 bool = false
                    _ = v12
                }
                v13 := GenS2{F0: 0, F1: 18446744073709551615, F2: v3}
                _ = v13
            }
        }
        _ = v10
    }
    if (655 >= -927) {
        for fi14 := 0; fi14 < 14; fi14++ {
            v15 := GenS1{F0: (v2 - (v2 - v2))}
            v1 += v1
            if true {
                if (v4 < v4) {
                    for fi16 := 0; fi16 < 11; fi16++ {
                        switch v3 {
                        case "foo bar":
                            v5 += v5
                        case "a":
                            v5 = (v5 / ((v5) | 1))
                        default:
                            _ = len(v7)
                        }
                        v5 = v5
                        _ = fi16
                    }
                    v4 += v4
                    var v17 bool = false
                    _ = v17
                }
                if (true && true) {
                    switch "" {
                    case "":
                        v18 := GenS2{F0: 226, F1: 42, F2: "\n\t"}
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 1; fi19++ {
                        v2 += v2
                        v7 = append(v7, 1)
                        v2 -= 2147483648
                        _ = fi19
                    }
                    for fi20 := 0; fi20 < 15; fi20++ {
                        fmt.Println(fi14, fi20)
                        var v21 uint16 = 794
                        fmt.Println(v3, fi20, v2)
                        _ = fi20
                        _ = v21
                    }
                }
            }
            _ = fi14
            _ = v15
        }
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 int32 = (v6 | v6)
            v24 := GenS1{F0: 931}
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    var v25 int64 = 1
    var v26 uint64 = (v1 - (v1 << ((1) & 63)))
    var v27 int64 = (^(v25))
    v26 = 575
    fmt.Println(v2, v25, v27)
    v28 := GenS3{F0: (-(v25)), F1: v6}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
