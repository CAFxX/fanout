package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int }
func genFn0(p0 uint, p1 uint64, p2 uint) int {
    for fi7 := 0; fi7 < 17; fi7++ {
        p2 = p2
        var v8 float64 = 0.1
        _ = fi7
        _ = v8
    }
    var v9 int32 = 2147483647
    fmt.Println(p2, p0)
    var v10 int16 = -32768
    v11 := []int{142, -764, -144, -460}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return -338
}
func genFn1() int16 {
    v12 := GenS1{F0: 1}
    if true {
        v13 := GenS1{F0: 828}
        if (!false) {
            v14 := GenS1{F0: 2147483647}
            var v15 int64 = 9223372036854775807
            v15 -= v15
            _ = v14
            _ = v15
        }
        _ = v13
    }
    _ = v12
    return -32768
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 42
    v3 := make(map[int]string)
    var v4 uint16 = 483
    var v5 uint16 = (v4 >> (((v2 - 2)) & 63))
    v6 := []string{""}
    if false {
        var v16 uint8 = 255
        _ = v16
    }
    _ = len(v6)
    var v17 uint = v2
    var v18 string = ""
    for fi19 := 0; fi19 < 13; fi19++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v5 -= (v5 << ((uint(v5)) & 63))
        v17 = v17
        _ = fi19
    }
    v2 += 370
    fmt.Println(v17, v4, v17)
    fmt.Println(v17, v4)
    delete(v1, -391)
    if false {
        _ = len(v3)
    }
    var v20 uint32 = 2
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v5 -= (v5 * v4)
    if (1000 <= 42) {
        v18 = "a"
    }
    fmt.Println(v4)
    if false {
        v18 = v18
    } else {
        v2 -= v2
        if (!(32767 != 4294967295)) {
            if false {
                var v21 uint16 = 42
                _ = v21
            } else {
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 int32 = 0
                    v20 += (v20 >> (((^(v17))) & 63))
                    _ = fi22
                    _ = v23
                }
                _ = len(v1)
            }
        } else {
            switch v18 {
            case "foo bar":
                fmt.Println(v17, v5)
            case "\n\t":
                if true {
                    for fi24 := 0; fi24 < 13; fi24++ {
                        v18 = (v18 + v18)
                        switch (fi24 - (-(fi24))) {
                        case 1:
                            v17 += 639
                        default:
                            var v25 uint16 = (v4 & (v4 / ((2) | 1)))
                            _ = v25
                        }
                        _ = fi24
                    }
                    var v26 int8 = 1
                    if false {
                        v27 := GenS1{F0: 204}
                        if (-393 >= 0) {
                            v28 := GenS1{F0: (v20 << ((v2) & 63))}
                            for fi29 := 0; fi29 < 5; fi29++ {
                                v30 := GenS2{F0: -448}
                                switch ((v18 + "a") + (v18 + ((v18 + v18) + "\n\t"))) {
                                case "":
                                    var v31 string = ""
                                    _ = v31
                                case "\n\t":
                                    var v32 int32 = 1
                                    _ = v32
                                case "foo bar":
                                    v33 := GenS3{F0: v4, F1: 14, F2: -9223372036854775808}
                                    _ = v33
                                }
                                _ = fi29
                                _ = v30
                            }
                            v34 := GenS4{F0: ((0.0 > 0.0) || false), F1: (v2 & (^(v2)))}
                            _ = v28
                            _ = v34
                        } else {
                            var v35 uint32 = v20
                            _ = v35
                        }
                        v2 -= 65535
                        _ = v27
                    }
                    _ = v26
                } else {
                    var v36 uint64 = 17
                    v37 := GenS5{F0: 4294967295}
                    _ = v36
                    _ = v37
                }
            case "a":
                v38 := GenS1{F0: v20}
                _ = v38
            default:
                var v39 uint16 = 515
                _ = v39
            }
        }
        switch ("foo bar" + v18) {
        case "foo bar":
            v40 := GenS1{F0: 626}
            _ = v40
        case "\n\t":
            if len(v6) > 0 {
                v6[0] = "\n\t"
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
