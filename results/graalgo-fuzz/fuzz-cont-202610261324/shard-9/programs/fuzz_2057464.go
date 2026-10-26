package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint16 }

func main() {
    var v1 uint8 = 175
    var v2 int64 = 0
    var v3 uint16 = 968
    var v4 bool = false
    var v5 int = 32767
    v6 := make(map[int]bool)
    if false {
        var v7 string = "\n\t"
        _ = v7
    } else {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint8 = v1
            v10 := []int{v5, (v5 + (fi8 + v5))}
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            switch ("a" + "foo bar") {
            case "hello":
                switch "a" {
                case "":
                    v2 = -258
                case "hello":
                    var v12 uint64 = 2
                    _ = v12
                default:
                    if (0.0 == float32(((1e10 + (0.0 / -2.25)) - (1e10 - 3.14159)))) {
                        for fi13 := 0; fi13 < 9; fi13++ {
                            var v14 uint16 = v3
                            var v15 uint32 = 223
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    }
                }
            default:
                if false {
                    var v16 float64 = 0.0
                    var v17 int64 = v2
                    v18 := GenS1{F0: (v3 ^ v3), F1: 366}
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
            v2 = v2
            _ = fi11
        }
    }
    switch (v5 + -2147483648) {
    case 1:
        var v19 bool = (!v4)
        _ = v19
    }
    if (!v4) {
        v4 = true
        switch 4294967295 {
        case 306:
            v20 := []int{}
            _ = v20
        case 32768:
            v21 := GenS1{F0: v3, F1: 4294967295}
            _ = v21
        case 32767:
            v2 -= int64(v3)
        }
    }
    v1 -= (v1 % ((v1) | 1))
    var v22 int8 = -106
    v5 -= v5
    v23 := []int{(v5 & -2147483648), (v5 + 42), v5}
    if false {
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := GenS2{F0: (255 < 112)}
            v26 := GenS1{F0: uint16(v1), F1: 713}
            var v27 uint8 = v1
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 4; fi28++ {
            if true {
                for fi29 := 0; fi29 < 20; fi29++ {
                    var v30 bool = (int64(v1) < v2)
                    v22 = 1
                    var v31 uint32 = 4294967295
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                for fi32 := 0; fi32 < 15; fi32++ {
                    switch "hello" {
                    case "hello":
                        v5 = -164
                    case "a":
                        if v4 {
                            v23 = append(v23, (fi32 << ((871) & 63)))
                            var v33 string = "a"
                            _ = v33
                        }
                    }
                    var v34 float64 = -2.25
                    _ = fi32
                    _ = v34
                }
            }
            for fi35 := 0; fi35 < 15; fi35++ {
                v3 = uint16(v1)
                v36 := GenS3{F0: v1, F1: 2147483648}
                fmt.Println(v1, fi35)
                _ = fi35
                _ = v36
            }
            _ = fi28
        }
        for fi37 := 0; fi37 < 18; fi37++ {
            for fi38 := 0; fi38 < 2; fi38++ {
                switch v5 {
                case -726:
                    fmt.Println(v5, v5)
                case 4294967295:
                    v4 = (18446744073709551615 > 2)
                case 128:
                    v2 = v2
                default:
                    v3 = (v3 & v3)
                }
                _ = fi38
            }
            if (fi37 >= (fi37 << ((4294967295) & 63))) {
                v2 = -893
            }
            _ = fi37
        }
    } else {
        var v39 float64 = -2.25
        var v40 uint64 = 2
        _ = v39
        _ = v40
    }
    v5 += v5
    fmt.Println(v2, v22)
    if (-2147483648 >= 0) {
        switch "" {
        case "":
            v41 := GenS4{F0: 18446744073709551615, F1: (^(v3))}
            _ = v41
        default:
            fmt.Println(v2, v5, v22)
        }
    } else {
        v42 := GenS3{F0: (v1 & (v1 << ((18446744073709551615) & 63))), F1: 65535}
        for fi43 := 0; fi43 < 11; fi43++ {
            v4 = (!v4)
            fmt.Println(v1, v4)
            _ = fi43
        }
        var v44 bool = ((!(3.14159 >= 3.14159)) || (1 >= 2147483647))
        _ = v42
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
