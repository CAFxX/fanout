package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := []int{-1, 127}
    var v2 int16 = 0
    var v3 uint8 = 61
    v4 := make(map[string]bool)
    v5 := make(map[int]bool)
    var v6 int64 = 1
    v1 = append(v1, 32768)
    var v7 string = "foo bar"
    var v8 string = v7
    v2 += v2
    v7 = (((v8 + (v7 + (v8 + (v7 + v8)))) + v8) + "a")
    var v9 float64 = 0.0
    v7 = v8
    if false {
        fmt.Println(v2, v2)
        if (false || true) {
            fmt.Println(v3, v6)
            for fi10 := 0; fi10 < 18; fi10++ {
                if false {
                    switch (v8 + (((v7 + ((v7 + v7) + v8)) + v7) + v8)) {
                    case "\n\t":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        var v11 uint = 255
                        _ = v11
                    }
                    v3 += (v3 ^ (v3 % ((v3) | 1)))
                } else {
                    var v12 float64 = v9
                    for fi13 := 0; fi13 < 7; fi13++ {
                        fmt.Println(fi10, fi13, v2)
                        v14 := GenS1{F0: (float32(v9) < 0.1), F1: 291, F2: -617}
                        _ = fi13
                        _ = v14
                    }
                    _ = v12
                }
                _ = fi10
            }
            v2 -= v2
        }
    } else {
        var v15 uint8 = v3
        v9 += v9
        v6 -= (v6 >> ((171) & 63))
        _ = v15
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch (v7 + "a") {
    case "foo bar":
        switch (v7 + "\n\t") {
        case "a":
            v16 := GenS1{F0: (-2147483648 == -1), F1: 687, F2: 56}
            _ = v16
        }
    case "a":
        fmt.Println(v9, v2)
    default:
        if true {
            var v17 float64 = 0.1
            v18 := GenS1{F0: (1.5 <= 1e10), F1: 1, F2: 1}
            _ = v17
            _ = v18
        } else {
            if ((^(v6)) == (v6 | v6)) {
                v7 = v7
                if true {
                    v7 = (("a" + v7) + (v7 + v7))
                } else {
                    var v19 int = 127
                    _ = v19
                }
                for fi20 := 0; fi20 < 15; fi20++ {
                    v21 := GenS1{F0: ((^(v2)) == v2), F1: 2, F2: 1}
                    switch (fi20 / ((63) | 1)) {
                    case -2147483648:
                        v1 = append(v1, fi20)
                    default:
                        v7 = "a"
                    }
                    v22 := GenS1{F0: false, F1: 1, F2: -1}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            } else {
                var v23 uint = 0
                var v24 uint64 = 698
                _ = v23
                _ = v24
            }
            v25 := GenS1{F0: true, F1: 981, F2: -1}
            for fi26 := 0; fi26 < 12; fi26++ {
                if true {
                    var v27 int = fi26
                    _ = v27
                }
                _ = fi26
            }
            _ = v25
        }
    }
    if ((v2 ^ 1) > -32768) {
        if (-2147483648 >= -2147483648) {
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 int32 = -378
                if false {
                    var v30 string = (v7 + (v7 + v8))
                    v31 := GenS2{F0: 42, F1: 2}
                    _ = v30
                    _ = v31
                } else {
                    v32 := GenS2{F0: 339, F1: 780}
                    var v33 string = ("\n\t" + v8)
                    _ = v32
                    _ = v33
                }
                switch v8 {
                case "":
                    var v34 int32 = v29
                    _ = v34
                }
                _ = fi28
                _ = v29
            }
            v35 := GenS2{F0: 2147483648, F1: 18446744073709551615}
            v5[2147483647] = true
            _ = v35
        } else {
            var v36 uint32 = 0
            var v37 float64 = 0.1
            _ = v36
            _ = v37
        }
        fmt.Println(v6)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
