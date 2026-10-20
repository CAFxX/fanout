package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 0
    v3 := []int{-1}
    v2 = v2
    v2 += (v2 << ((0) & 63))
    if true {
        for fi4 := 0; fi4 < 13; fi4++ {
            v2 -= v2
            v2 = v2
            _ = fi4
        }
    }
    v3 = append(v3, -1)
    v2 = v2
    var v5 uint64 = 804
    var v6 uint32 = (v2 >> ((2) & 63))
    var v7 bool = (486 >= 0)
    delete(v1, v2)
    v5 = (v5 % (((v5 ^ 1)) | 1))
    v2 = 210
    v5 += (v5 % ((42) | 1))
    _ = len(v1)
    var v8 int = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (((3.14159 + ((1e10 + 3.14159) * ((-2.25 - (-2.25 - -2.25)) + 3.14159))) / 0.1) != (0.1 + 1.5)) {
        v9 := GenS1{F0: -2147483648}
        var v10 int = (v8 ^ 0)
        _ = v10
        _ = v9
    } else {
        if v7 {
            if (!v7) {
                if false {
                    var v11 string = ("foo bar" + "\n\t")
                    fmt.Println(v8, v8)
                    _ = v11
                } else {
                    fmt.Println(v6, v8, v2)
                    switch (v8 & v8) {
                    case -2147483648:
                        v12 := GenS2{F0: "a"}
                        _ = v12
                    case -9223372036854775808:
                        v6 = v6
                    default:
                        v13 := GenS3{F0: v7}
                        _ = v13
                    }
                }
                var v14 float32 = 0.1
                switch (v8 - (v8 & v8)) {
                case 1000:
                    v3 = append(v3, v8)
                case -222:
                    if (v14 == v14) {
                        var v15 uint = 551
                        switch "a" {
                        case "hello":
                            fmt.Println(v5)
                        case "":
                            _ = len(v1)
                        }
                        _ = len(v3)
                        _ = v15
                    }
                }
                _ = v14
            } else {
                for fi16 := 0; fi16 < 18; fi16++ {
                    switch (fi16 | fi16) {
                    case 32767:
                        v17 := GenS2{F0: "a"}
                        _ = v17
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    var v18 uint = 4294967295
                    v8 += fi16
                    _ = fi16
                    _ = v18
                }
                delete(v1, (v6 / ((v2) | 1)))
                var v19 bool = v7
                _ = v19
            }
            fmt.Println(v7, v2, v7)
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 bool = (true || true)
                _ = fi20
                _ = v21
            }
        }
        fmt.Println(v2, v8, v8)
        v2 = v2
    }
    v2 = (v2 * 4294967295)
    if v7 {
        if true {
            v6 -= v6
        } else {
            v2 -= v2
            fmt.Println(v5, v8, v6)
            if v7 {
                var v22 bool = ((v8 > v8) || (3.14159 >= -2.25))
                v23 := GenS4{F0: 2147483647}
                v6 = (v2 + 470)
                _ = v22
                _ = v23
            } else {
                if ((^(v2)) <= 2147483648) {
                    for fi24 := 0; fi24 < 6; fi24++ {
                        var v25 float64 = float64((0.0 + 1.5))
                        v3 = append(v3, fi24)
                        v25 = (v25 - v25)
                        _ = fi24
                        _ = v25
                    }
                    fmt.Println(v6, v7, v6)
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
        }
    } else {
        v5 -= (v5 / (((v5 + 564)) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
