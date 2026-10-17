package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint; F2 string }
type GenS3 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS4 struct { F0 string; F1 uint16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{-2147483648}
    var v3 uint = 65535
    var v4 uint = v3
    v5 := make(map[string]uint32)
    var v6 int16 = -32768
    v7 := v1[42]
    var v8 float64 = 3.14159
    v4 = 473
    var v9 int64 = -1
    if len(v2) > 0 {
        v2[0] = 718
    }
    switch "a" {
    case "hello":
        var v10 int = v7
        _ = v10
    }
    var v11 uint64 = 42
    var v12 float64 = ((float64(0.1) / v8) / float64(((1.5 * 3.14159) * -2.25)))
    var v13 uint16 = 822
    var v14 int = (^(v7))
    fmt.Println(v12)
    for fi15 := 0; fi15 < 17; fi15++ {
        v4 = (v4 % ((42) | 1))
        _ = fi15
    }
    v4 = 4294967295
    v4 -= (v3 + v3)
    for fi16 := 0; fi16 < 9; fi16++ {
        if (v6 > v6) {
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 bool = ((!(((v11 >= uint64(v13)) || true) && ("hello" == ""))) || true)
                _ = fi17
                _ = v18
            }
            v19 := GenS1{F0: v11}
            v20 := GenS2{F0: 36, F1: v4, F2: ("hello" + ("a" + ("hello" + ("foo bar" + ("" + ("" + "\n\t"))))))}
            _ = v19
            _ = v20
        } else {
            if true {
                var v21 float64 = (v8 * v8)
                v14 -= (fi16 << (((v4 >> ((v4) & 63))) & 63))
                _ = v21
            } else {
                for fi22 := 0; fi22 < 13; fi22++ {
                    var v23 uint8 = 128
                    if (true || ((v13 * v13) >= (v13 & (v13 ^ (v13 >> ((v4) & 63)))))) {
                        if (false || ((float64(1.5) - (v8 + ((v12 - v8) * ((v12 / (v8 - (v12 * 0.1))) - v12)))) < v12)) {
                            v13 = (v13 * 255)
                            var v24 uint = (v3 >> (((v3 / ((503) | 1))) & 63))
                            if false {
                                v2 = append(v2, v7)
                                switch (fi22 | int(v11)) {
                                case 9223372036854775807:
                                    var v25 uint = 363
                                    _ = v25
                                case 32768:
                                    if false {
                                        v5["a"] = 1
                                    }
                                }
                            }
                            var v26 uint16 = v13
                            v14 = -418
                            _ = v24
                            _ = v26
                        }
                        var v27 string = ("foo bar" + "hello")
                        if true {
                            v9 = v9
                            fmt.Println(fi22, v7, v8)
                        }
                        _ = v27
                    } else {
                        delete(v1, 2147483648)
                        if true {
                            var v28 uint32 = 776
                            v6 += v6
                            _ = v28
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v29 string = "a"
                            fmt.Println(v23)
                            _ = v29
                        }
                        v30 := GenS3{F0: v3, F1: (v23 & v23), F2: 53}
                        _ = v30
                    }
                    _ = fi22
                    _ = v23
                }
            }
            v31 := GenS4{F0: "a", F1: 1}
            _ = v31
        }
        _ = fi16
    }
    v6 = (v6 << (((v3 | (v3 ^ 2))) & 63))
    for fi32 := 0; fi32 < 11; fi32++ {
        _ = len(v2)
        var v33 int16 = v6
        _ = fi32
        _ = v33
    }
    if (!(127 >= 39)) {
        fmt.Println(v14)
    } else {
        v34 := GenS4{F0: "foo bar", F1: 689}
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
