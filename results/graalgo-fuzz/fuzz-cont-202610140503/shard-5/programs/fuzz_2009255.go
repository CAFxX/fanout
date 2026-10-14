package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int8; F1 int8; F2 bool }

func main() {
    var v1 uint = 0
    var v2 int = 128
    var v3 string = ""
    var v4 int8 = 42
    fmt.Println(v4, v1, v3)
    v1 += 42
    for fi5 := 0; fi5 < 17; fi5++ {
        if (980 < 42) {
            for fi6 := 0; fi6 < 17; fi6++ {
                v4 = (v4 % ((v4) | 1))
                var v7 int16 = 32767
                v2 -= 0
                _ = fi6
                _ = v7
            }
        }
        if false {
            fmt.Println(v2, fi5, fi5)
            if false {
                v3 = "foo bar"
                if (v3 == "hello") {
                    v4 = (v4 - v4)
                    v2 -= fi5
                    switch (v3 + v3) {
                    case "hello":
                        v1 -= v1
                    case "\n\t":
                        for fi8 := 0; fi8 < 2; fi8++ {
                            v9 := []int{(v2 - fi5)}
                            _ = fi8
                            _ = v9
                        }
                    default:
                        switch (v2 ^ 709) {
                        case 378:
                            if false {
                                v2 = 1
                                v10 := []int{(fi5 % ((924) | 1)), (fi5 ^ v2), v2, v2}
                                _ = v10
                            } else {
                                v2 = (^(v2))
                                v11 := []int{fi5, v2, 4294967295}
                                switch v2 {
                                case -22:
                                    v3 = "hello"
                                case -9223372036854775808:
                                    for fi12 := 0; fi12 < 3; fi12++ {
                                        v1 = (v1 | (v1 + (^(v1))))
                                        fmt.Println(fi5)
                                        _ = fi12
                                    }
                                }
                                _ = v11
                            }
                        }
                    }
                } else {
                    var v13 float64 = float64(-2.25)
                    v1 = v1
                    _ = v13
                }
                if (int16(fi5) > 42) {
                    if true {
                        v2 += (v2 & 4294967295)
                        v14 := []int{(v2 & fi5), 661, fi5}
                        _ = v14
                    }
                    v3 = "foo bar"
                } else {
                    switch v2 {
                    case -623:
                        v4 -= (v4 ^ (-(v4)))
                    default:
                        v2 = 32768
                    }
                    v2 += (v2 | (fi5 ^ v2))
                    v2 = (v2 & (v2 + (fi5 << ((770) & 63))))
                }
            } else {
                v4 = v4
                for fi15 := 0; fi15 < 9; fi15++ {
                    v16 := []int{}
                    var v17 int = fi15
                    fmt.Println(v4, v3)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            v18 := GenS1{F0: (v1 * v1), F1: 32767}
            _ = v18
        }
        v1 -= (v1 >> (((^(v1))) & 63))
        _ = fi5
    }
    v19 := GenS2{F0: v4, F1: 40, F2: (v2 <= v2)}
    v3 = (v3 + v3)
    fmt.Println(v2, v3)
    v20 := []int{(v2 ^ (v2 % ((2147483647) | 1))), 0, v2}
    var v21 float64 = (0.0 * 0.0)
    if false {
        for fi22 := 0; fi22 < 13; fi22++ {
            v4 -= v4
            _ = fi22
        }
        v3 = v3
    } else {
        switch (^(v2)) {
        case 1:
            var v23 uint32 = 2147483647
            _ = v23
        case -268:
            if (v2 != v2) {
                v24 := GenS1{F0: (v1 - v1), F1: (v2 * v2)}
                v3 = "a"
                _ = v24
            }
        default:
            if len(v20) > 0 {
                v20[0] = (v2 & -1)
            }
        }
        var v25 float64 = v21
        _ = v25
    }
    if true {
        for fi26 := 0; fi26 < 12; fi26++ {
            _ = len(v20)
            switch 42 {
            case 957:
                switch (v2 / ((4294967295) | 1)) {
                case -104:
                    var v27 uint64 = 42
                    _ = v27
                case 0:
                    v20 = append(v20, (fi26 | v2))
                default:
                    if (2 >= 1) {
                        v4 += (v4 | v4)
                        var v28 int8 = v4
                        _ = v28
                    } else {
                        v29 := GenS2{F0: v4, F1: (v4 | (-(v4))), F2: false}
                        _ = v29
                    }
                }
            }
            _ = len(v20)
            _ = fi26
        }
    }
    var v30 uint = 42
    var v31 bool = (3.14159 < 3.14159)
    for fi32 := 0; fi32 < 15; fi32++ {
        switch (fi32 * (v2 - (fi32 - (-(fi32))))) {
        case 4294967295:
            var v33 int = (fi32 / ((fi32) | 1))
            _ = v33
        }
        _ = fi32
    }
    if true {
        v30 -= v30
        switch "" {
        case "":
            v34 := GenS1{F0: (v1 - (v1 - 42)), F1: 32768}
            _ = v34
        case "\n\t":
            var v35 uint = v30
            _ = v35
        }
        var v36 int64 = -9223372036854775808
        _ = v36
    }
    v37 := GenS1{F0: (v30 << ((v30) & 63)), F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
}
