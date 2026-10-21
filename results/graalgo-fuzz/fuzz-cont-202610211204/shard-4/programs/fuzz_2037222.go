package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int16; F2 int }

func main() {
    var v1 uint16 = 32767
    var v2 int64 = 1
    var v3 float64 = 0.0
    var v4 string = "foo bar"
    var v5 uint64 = 18446744073709551615
    var v6 int = 70
    var v7 uint64 = (v5 ^ v5)
    if false {
        for fi8 := 0; fi8 < 2; fi8++ {
            switch 1 {
            case -598:
                for fi9 := 0; fi9 < 9; fi9++ {
                    v7 = v5
                    v10 := []int{fi8, (v6 ^ 237), (v6 | 32767), (fi9 >> ((18446744073709551615) & 63))}
                    fmt.Println(v1, fi8)
                    _ = fi9
                    _ = v10
                }
            case 594:
                v11 := []int{v6, (v6 * (^(v6)))}
                _ = v11
            }
            v12 := GenS1{F0: v4, F1: 568}
            fmt.Println(v5)
            _ = fi8
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        if false {
            v14 := []int{fi13, v6}
            var v15 int = (^(v6))
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 bool = true
                var v18 int16 = 398
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v14
            _ = v15
        } else {
            for fi19 := 0; fi19 < 14; fi19++ {
                if false {
                    if (0 != -541) {
                        switch (("hello" + v4) + v4) {
                        case "":
                            v20 := GenS2{F0: "foo bar", F1: v7}
                            _ = v20
                        case "foo bar":
                            if (!false) {
                                if (v2 < (^(v2))) {
                                    v4 = "\n\t"
                                    switch 2147483647 {
                                    case -114:
                                        if (198 >= 2147483648) {
                                            v21 := []int{(fi13 % ((127) | 1)), -2147483648, fi13, (fi19 << ((230) & 63))}
                                            _ = v21
                                        } else {
                                            v22 := GenS1{F0: v4, F1: 851}
                                            fmt.Println(fi19)
                                            v4 = (v4 + "")
                                            _ = v22
                                        }
                                    case 0:
                                        fmt.Println(v6)
                                    }
                                    v23 := GenS3{F0: 42}
                                    _ = v23
                                } else {
                                    v2 -= -1
                                    v24 := []int{}
                                    v6 -= v6
                                    _ = v24
                                }
                            }
                        case "hello":
                            var v25 int32 = 1
                            _ = v25
                        }
                        fmt.Println(v4, v1, v4)
                    }
                    if (!((v4 + v4) <= (v4 + v4))) {
                        v4 = ("a" + v4)
                    } else {
                        switch v4 {
                        case "foo bar":
                            if true {
                                if false {
                                    var v26 int32 = -969
                                    _ = v26
                                }
                                v1 += v1
                            }
                        case "\n\t":
                            var v27 uint = 0
                            _ = v27
                        default:
                            if false {
                                v4 = v4
                                var v28 string = v4
                                v29 := []int{(-(v6)), (^(v6)), 957}
                                _ = v28
                                _ = v29
                            }
                        }
                    }
                } else {
                    var v30 uint16 = 167
                    v31 := GenS4{F0: 65535, F1: -1, F2: int(v5)}
                    _ = v30
                    _ = v31
                }
                v3 = (v3 + (float64(1e10) / ((0.0 / 1.0) / (((0.1 + (v3 - ((((v3 + 1e10) * v3) + v3) + (v3 * v3)))) / v3) / (v3 * (v3 / (1.5 / (1.5 / v3))))))))
                var v32 float64 = v3
                _ = fi19
                _ = v32
            }
            v33 := []int{}
            _ = v33
        }
        _ = fi13
    }
    v34 := GenS3{F0: 197}
    v4 = (v4 + v4)
    v35 := []int{v6, v6}
    fmt.Println(v6)
    v36 := GenS2{F0: "hello", F1: 2}
    fmt.Println(v4)
    v7 += 33
    if (!false) {
        for fi37 := 0; fi37 < 16; fi37++ {
            v2 = (v2 | (v2 % ((v2) | 1)))
            _ = fi37
        }
    } else {
        if (true && true) {
            var v38 int32 = 42
            v6 -= (^(v6))
            _ = v38
        } else {
            fmt.Println(v3, v1, v2)
            var v39 int32 = 2147483647
            _ = v39
        }
        v2 -= -1
    }
    fmt.Println(v2, v1, v6)
    v7 = (v7 / ((v7) | 1))
    var v40 int = (v6 % (((v6 - -1)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
