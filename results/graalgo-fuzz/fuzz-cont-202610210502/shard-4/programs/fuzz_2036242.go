package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 798
    var v3 int16 = -1
    v4 := make(map[string]bool)
    var v5 float32 = (3.14159 - 3.14159)
    for fi6 := 0; fi6 < 15; fi6++ {
        if (v5 == (v5 - ((0.1 - v5) / v5))) {
            v7 := []int{}
            v5 += 1.5
            var v8 uint64 = 42
            _ = v7
            _ = v8
        } else {
            v9 := []int{1, (fi6 / (((-(fi6))) | 1)), 32768, (fi6 << ((691) & 63))}
            fmt.Println(v2, v5)
            if false {
                v10 := GenS1{F0: "foo bar"}
                _ = v10
            } else {
                var v11 int = fi6
                var v12 int = (-(fi6))
                _ = v11
                _ = v12
            }
            _ = v9
        }
        v2 = (v2 / ((v2) | 1))
        var v13 int32 = v2
        _ = fi6
        _ = v13
    }
    if false {
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 float32 = v5
            v3 -= (v3 / ((32767) | 1))
            if false {
                v2 -= v2
                if true {
                    fmt.Println(fi14, v3)
                    for fi16 := 0; fi16 < 12; fi16++ {
                        v17 := GenS1{F0: "\n\t"}
                        var v18 int = fi16
                        fmt.Println(fi14, v3, v15)
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 18; fi19++ {
                        v20 := GenS1{F0: "\n\t"}
                        var v21 int = (fi14 % (((fi14 | (fi19 / ((fi19) | 1)))) | 1))
                        switch "" {
                        case "foo bar":
                            v15 = (0.1 / v5)
                        case "":
                            v3 += 1000
                        case "a":
                            if false {
                                v22 := GenS1{F0: ("a" + "a")}
                                var v23 uint16 = 255
                                _ = v22
                                _ = v23
                            } else {
                                var v24 int8 = 127
                                switch "" {
                                case "":
                                    var v25 int16 = v3
                                    _ = v25
                                default:
                                    v26 := GenS2{F0: (v3 * (v3 & 454))}
                                    _ = v26
                                }
                                var v27 uint32 = 0
                                _ = v24
                                _ = v27
                            }
                            v5 += 0.0
                            v28 := GenS3{F0: int32(fi19)}
                            _ = v28
                        default:
                            v29 := GenS2{F0: (v3 ^ v3)}
                            _ = v29
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                } else {
                    v15 += v15
                    for fi30 := 0; fi30 < 4; fi30++ {
                        v5 = 0.0
                        v3 = 663
                        _ = fi30
                    }
                    var v31 bool = false
                    _ = v31
                }
            } else {
                v5 += 1.5
            }
            _ = fi14
            _ = v15
        }
        if (((!(42 > 18446744073709551615)) && false) || false) {
            for fi32 := 0; fi32 < 8; fi32++ {
                v3 = (v3 + (v3 << ((2) & 63)))
                _ = fi32
            }
        } else {
            if (!((0 == 125) && true)) {
                var v33 bool = (false || false)
                var v34 float64 = 1e10
                v35 := GenS2{F0: (v3 | (v3 | (v3 * -1)))}
                _ = v33
                _ = v34
                _ = v35
            }
        }
        var v36 bool = true
        _ = v36
    }
    for fi37 := 0; fi37 < 2; fi37++ {
        var v38 uint8 = 127
        _ = fi37
        _ = v38
    }
    if true {
        v39 := GenS2{F0: v3}
        if ("foo bar" > "hello") {
            v5 = -2.25
        }
        _ = v39
    } else {
        v5 = (0.0 + ((float32((-2.25 + (0.0 - 0.1))) + v5) + float32((3.14159 / 3.14159))))
        var v40 uint = 375
        v2 += v2
        _ = v40
    }
    if true {
        var v41 uint = 65535
        if (4294967295 >= 1) {
            v3 = v3
            v2 = v2
            v42 := []int{32768, -9223372036854775808, 9223372036854775807, 42}
            _ = v42
        } else {
            v43 := GenS3{F0: 410}
            v44 := GenS2{F0: v3}
            _ = v43
            _ = v44
        }
        v2 = (v2 << ((v41) & 63))
        _ = v41
    }
    var v45 uint16 = 1
    if false {
        var v46 uint16 = v45
        _ = v46
    }
    var v47 uint = 154
    v48 := GenS2{F0: (^(v3))}
    v3 = (v3 >> ((v47) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v47
    _ = v48
    _ = v5
}
