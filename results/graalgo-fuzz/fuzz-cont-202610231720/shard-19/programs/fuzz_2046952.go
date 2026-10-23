package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint = 737
    var v2 bool = ((4294967295 != 50) && true)
    v3 := make(map[string]uint32)
    var v4 int = -989
    var v5 uint32 = 65535
    var v6 string = "a"
    var v7 int64 = -505
    var v8 uint64 = 2
    v9 := []int{(v4 / ((int(v7)) | 1)), v4}
    var v10 uint8 = 0
    if true {
        if len(v9) > 0 {
            v9[0] = 2147483647
        }
        v11 := GenS1{F0: uint32(v10)}
        _ = v11
    } else {
        var v12 float64 = (float64(((((1.5 / 1e10) + 1.5) + float32((0.0 * 3.14159))) * float32((-2.25 + 3.14159)))) - (((0.0 * (float64(0.0) - 3.14159)) / 1e10) - (3.14159 / -2.25)))
        fmt.Println(v6)
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        v9 = append(v9, -108)
        if ((fi13 >= (v4 | (fi13 ^ 349))) && true) {
            var v14 int64 = 1
            _ = v14
        }
        _ = fi13
    }
    v9 = append(v9, -1)
    fmt.Println(v6, v7)
    v15 := GenS1{F0: v5}
    v3[("foo bar" + "a")] = 1
    v10 -= v10
    if v2 {
        v9 = append(v9, v4)
        for fi16 := 0; fi16 < 4; fi16++ {
            if (!(v6 < v6)) {
                if v2 {
                    if (int64(v1) >= v7) {
                        var v17 int8 = -128
                        for fi18 := 0; fi18 < 10; fi18++ {
                            if (1 >= 0) {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                v8 -= v8
                                fmt.Println(v7, v7)
                                v19 := GenS1{F0: (v5 ^ 226)}
                                if false {
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(v2)
                                    case "a":
                                        var v20 uint8 = v10
                                        v7 = v7
                                        var v21 uint16 = 1
                                        _ = v20
                                        _ = v21
                                    case "\n\t":
                                        fmt.Println(v10, v10)
                                    }
                                    v17 = int8(v7)
                                } else {
                                    if true {
                                        switch 127 {
                                        case -513:
                                            var v22 int = fi16
                                            _ = v22
                                        case -931:
                                            _ = len(v3)
                                        default:
                                            if true {
                                                v17 -= 31
                                                if (54 <= v10) {
                                                    var v23 string = ((((v6 + "\n\t") + v6) + (v6 + v6)) + v6)
                                                    _ = v23
                                                }
                                            }
                                        }
                                        if (!(v2 && ((v1 / ((4294967295) | 1)) != 962))) {
                                            v6 = v6
                                            v24 := GenS1{F0: v5}
                                            v7 = v7
                                            _ = v24
                                        }
                                        if v2 {
                                            switch v4 {
                                            case -9223372036854775808:
                                                fmt.Println(v5, v7, fi16)
                                            case 9223372036854775807:
                                                v2 = v2
                                            }
                                        }
                                    }
                                    fmt.Println(fi16)
                                }
                                _ = v19
                            } else {
                                v6 = v6
                                var v25 float64 = 1e10
                                _ = v25
                            }
                            v26 := GenS1{F0: 93}
                            fmt.Println(v10, v4, v6)
                            var v27 int8 = (v17 - v17)
                            if ((v5 | v5) == (^(v5))) {
                                var v28 uint16 = 1
                                var v29 int8 = v27
                                _ = len(v9)
                                _ = v28
                                _ = v29
                            } else {
                                var v30 uint32 = (v5 / ((uint32(v8)) | 1))
                                if len(v9) > 0 {
                                    v9[0] = (fi16 ^ fi18)
                                }
                                _ = v30
                            }
                            _ = fi18
                            _ = v26
                            _ = v27
                        }
                        v9 = append(v9, -1)
                        _ = v17
                    }
                    for fi31 := 0; fi31 < 14; fi31++ {
                        v1 = 282
                        _ = fi31
                    }
                }
            } else {
                v10 += v10
                fmt.Println(v5)
                var v32 int8 = -128
                _ = v32
            }
            _ = fi16
        }
    } else {
        var v33 int64 = int64(v1)
        _ = v33
    }
    v10 -= uint8(v8)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        var v35 float64 = 0.0
        fmt.Println(v6, v7)
        _ = fi34
        _ = v35
    }
    v9 = append(v9, v4)
    fmt.Println(v1, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
