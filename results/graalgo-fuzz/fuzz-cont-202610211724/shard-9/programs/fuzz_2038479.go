package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint; F1 int8 }
func genFn0(p0 bool, p1 string) int8 {
    v5 := []int{2147483647}
    switch "a" {
    case "foo bar":
        fmt.Println(p0, p0)
    case "hello":
        switch 2147483647 {
        case 1:
            v6 := GenS1{F0: false}
            _ = v6
        case 368:
            p0 = true
        }
    case "\n\t":
        p0 = p0
    default:
        switch p1 {
        case "hello":
            if p0 {
                for fi7 := 0; fi7 < 13; fi7++ {
                    _ = len(v5)
                    p1 = p1
                    _ = fi7
                }
                fmt.Println(p0, p1)
                _ = len(v5)
            } else {
                var v8 int8 = 36
                if len(v5) > 0 {
                    v5[0] = 32768
                }
                v9 := GenS2{F0: 634, F1: -32768}
                _ = v8
                _ = v9
            }
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return -1
}

func main() {
    var v1 string = "\n\t"
    var v2 int = 1
    v3 := []string{v1, v1, ""}
    v4 := make(map[uint32]int)
    fmt.Println(v2, v2)
    if len(v3) > 0 {
        v3[0] = ("a" + v1)
    }
    v10 := GenS1{F0: (!false)}
    var v11 int8 = -128
    switch v1 {
    case "a":
        v12 := GenS3{F0: 783, F1: v1}
        _ = v12
    case "hello":
        v11 += (v11 >> ((0) & 63))
    default:
        var v13 int = 1
        _ = v13
    }
    if false {
        if true {
            if (true || false) {
                for fi14 := 0; fi14 < 18; fi14++ {
                    if (false || (0 == 548)) {
                        for fi15 := 0; fi15 < 7; fi15++ {
                            switch (fi14 & 128) {
                            case 9223372036854775807:
                                switch "" {
                                case "a":
                                    v11 = 57
                                }
                            default:
                                delete(v4, 176)
                            }
                            v1 = (v1 + "")
                            if false {
                                var v16 bool = true
                                if v16 {
                                    v17 := v4[2147483647]
                                    _ = v17
                                } else {
                                    var v18 string = "a"
                                    v4[843] = (fi15 % (((v2 - fi15)) | 1))
                                    if (144 > 0) {
                                        v4[636] = fi15
                                    } else {
                                        v19 := GenS2{F0: 4294967295, F1: -941}
                                        var v20 int32 = -368
                                        v1 = v1
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = v18
                                }
                                var v21 float32 = (3.14159 - 0.1)
                                _ = v16
                                _ = v21
                            }
                            _ = fi15
                        }
                        fmt.Println(v11, v11)
                        v3 = append(v3, v1)
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if false {
                            v3 = append(v3, "\n\t")
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                    for fi22 := 0; fi22 < 5; fi22++ {
                        v23 := GenS3{F0: 67, F1: (v1 + "foo bar")}
                        var v24 uint64 = 2
                        var v25 float64 = (-2.25 * 1e10)
                        fmt.Println(v24, v2, v2)
                        switch v2 {
                        case 32767:
                            var v26 uint32 = 699
                            _ = v26
                        case -390:
                            v2 += fi22
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    _ = fi14
                }
                v2 = (v2 | -2147483648)
                v3 = append(v3, "\n\t")
            } else {
                v4[598] = 711
                v4[4294967295] = 32767
                v11 = genFn0((false || false), "")
            }
            if (v11 < (v11 | -63)) {
                var v27 float32 = 0.0
                for fi28 := 0; fi28 < 17; fi28++ {
                    fmt.Println(fi28)
                    _ = fi28
                }
                _ = v27
            } else {
                v29 := GenS1{F0: (42 > v11)}
                v1 = (v1 + (v1 + v1))
                _ = v29
            }
        }
    } else {
        fmt.Println(v11)
    }
    v3 = append(v3, ("hello" + (v1 + v1)))
    fmt.Println(v1, v11, v11)
    fmt.Println(v2, v1, v1)
    for fi30 := 0; fi30 < 13; fi30++ {
        v31 := GenS4{F0: 18446744073709551615, F1: -128}
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        v4[0] = -1
        _ = fi32
    }
    delete(v4, 237)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
