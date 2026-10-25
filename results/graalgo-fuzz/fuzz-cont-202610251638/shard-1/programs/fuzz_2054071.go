package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }

func main() {
    var v1 bool = false
    var v2 uint64 = 251
    var v3 bool = (((v2 & 4294967295) > v2) && v1)
    var v4 int32 = -2147483648
    for fi5 := 0; fi5 < 14; fi5++ {
        fmt.Println(v2, v3, fi5)
        fmt.Println(v2, v4, v3)
        if ((v4 <= (v4 ^ v4)) && v3) {
            if true {
                var v6 uint16 = 65535
                _ = v6
            }
        }
        _ = fi5
    }
    var v7 string = "foo bar"
    var v8 float32 = 1.5
    fmt.Println(v7, v7)
    var v9 float64 = (3.14159 / 1.0)
    var v10 int = 560
    var v11 uint = 0
    v4 = 2147483647
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 bool = ((v3 || (65535 != 1)) || v1)
        for fi14 := 0; fi14 < 12; fi14++ {
            if v3 {
                v15 := []int{(v10 >> ((2) & 63)), -923, (fi14 | (v10 >> ((v11) & 63)))}
                fmt.Println(v13, v10, fi14)
                var v16 int8 = -128
                _ = v15
                _ = v16
            } else {
                v10 = (fi12 - fi14)
                v17 := GenS1{F0: 18446744073709551615}
                if v1 {
                    v18 := GenS1{F0: v2}
                    if (fi14 != 127) {
                        if v3 {
                            switch fi14 {
                            case -1:
                                if v13 {
                                    v19 := []int{(^(fi14)), fi12, fi12}
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    _ = v19
                                }
                            case 0:
                                var v20 float32 = (v8 - v8)
                                _ = v20
                            case 127:
                                fmt.Println(fi12, fi12, v13)
                            }
                        }
                        var v21 float32 = 1.5
                        switch ((v7 + v7) + "foo bar") {
                        case "foo bar":
                            v22 := GenS1{F0: (^(v2))}
                            _ = v22
                        case "a":
                            fmt.Println(v2)
                        case "hello":
                            var v23 int8 = -108
                            _ = v23
                        }
                        _ = v21
                    }
                    v24 := []int{}
                    if v3 {
                        fmt.Println(v8)
                    } else {
                        switch "a" {
                        case "hello":
                            fmt.Println(v3)
                        case "foo bar":
                            switch (((v7 + "a") + v7) + (v7 + v7)) {
                            case "foo bar":
                                if len(v24) > 0 {
                                    v24[0] = v10
                                }
                            default:
                                fmt.Println(fi12, v9)
                            }
                        default:
                            var v25 uint = (^(v11))
                            _ = v25
                        }
                    }
                    var v26 string = "\n\t"
                    _ = v18
                    _ = v24
                    _ = v26
                } else {
                    v27 := []int{v10, (fi12 * (fi12 % ((fi12) | 1)))}
                    _ = v27
                }
                _ = v17
            }
            v11 += (^(v11))
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    var v28 uint64 = 0
    v29 := []int{v10}
    if ((551 > 0) && (0 >= (-(v4)))) {
        v30 := GenS2{F0: v3}
        _ = v30
    }
    var v31 bool = (128 <= 0)
    if v3 {
        v32 := GenS3{F0: "a"}
        var v33 uint32 = 42
        for fi34 := 0; fi34 < 14; fi34++ {
            fmt.Println(v1)
            v8 += v8
            _ = fi34
        }
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
