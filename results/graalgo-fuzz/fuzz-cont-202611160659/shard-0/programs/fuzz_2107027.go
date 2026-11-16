package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint8 }
func genFn0() int {
    v7 := []int{42, 976}
    for fi8 := 0; fi8 < 4; fi8++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
    }
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 int32, p1 bool, p2 uint8) uint8 {
    switch -1 {
    case 0:
        var v9 uint16 = 1
        _ = v9
    case 1:
        v10 := GenS1{F0: 4294967295}
        _ = v10
    case -2147483648:
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int8 = 1
            _ = fi11
            _ = v12
        }
    default:
        if p1 {
            switch "\n\t" {
            case "foo bar":
                var v13 string = "hello"
                _ = v13
            }
        }
    }
    v14 := GenS1{F0: 424}
    var v15 int64 = 603
    if (p1 || p1) {
        v16 := []int{-1, -2147483648, 127}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return (p2 / (((p2 / ((p2) | 1))) | 1))
}

func main() {
    var v1 int16 = -32768
    var v2 float32 = 1.5
    v3 := []uint32{2, 2147483648, 2147483648}
    var v4 int32 = -716
    var v5 int = -661
    var v6 bool = true
    var v17 bool = v6
    v6 = ((255 != 25) && v6)
    var v18 int64 = 0
    fmt.Println(v4, v18)
    if true {
        for fi19 := 0; fi19 < 16; fi19++ {
            fmt.Println(fi19, v1)
            if ((v2 + (v2 + (v2 - ((v2 - (float32(-2.25) + (0.1 / v2))) * 1.5)))) <= v2) {
                v3 = append(v3, 50)
            } else {
                if v6 {
                    switch "foo bar" {
                    case "\n\t":
                        if true {
                            var v20 int8 = 0
                            _ = v20
                        } else {
                            for fi21 := 0; fi21 < 1; fi21++ {
                                switch (("hello" + "a") + "foo bar") {
                                case "a":
                                    v1 += -829
                                }
                                _ = fi21
                            }
                            v22 := GenS1{F0: 4294967295}
                            _ = v22
                        }
                    case "":
                        if len(v3) > 0 {
                            v3[0] = 767
                        }
                    case "hello":
                        v3 = append(v3, 397)
                    default:
                        fmt.Println(v18)
                    }
                    var v23 uint16 = 32767
                    _ = v23
                } else {
                    switch -152 {
                    case 4294967295:
                        v24 := GenS2{F0: (v18 & (v18 / ((v18) | 1))), F1: -9223372036854775808}
                        _ = v24
                    }
                    if ((v18 * (v18 / ((-1) | 1))) > (^(v18))) {
                        for fi25 := 0; fi25 < 11; fi25++ {
                            switch (fi25 * fi19) {
                            case 127:
                                v17 = (v4 < 2147483647)
                            default:
                                var v26 int64 = v18
                                _ = v26
                            }
                            _ = fi25
                        }
                    }
                }
                var v27 int64 = -1
                _ = v27
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi19
        }
    } else {
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 int32 = 0
            var v30 uint8 = 151
            _ = fi28
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 10; fi31++ {
            for fi32 := 0; fi32 < 1; fi32++ {
                var v33 int16 = -756
                if (255 > 18446744073709551615) {
                    v4 -= (v4 % ((v4) | 1))
                    v33 = v1
                    var v34 int16 = v33
                    v35 := GenS3{F0: 126}
                    _ = len(v3)
                    v36 := GenS1{F0: 4294967295}
                    _ = v34
                    _ = v35
                    _ = v36
                }
                var v37 uint16 = 255
                _ = fi32
                _ = v33
                _ = v37
            }
            v6 = v6
            _ = fi31
        }
        for fi38 := 0; fi38 < 15; fi38++ {
            v2 -= v2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v39 uint16 = 162
            _ = fi38
            _ = v39
        }
    }
    var v40 string = "foo bar"
    v2 = ((v2 / 1.0) + v2)
    if true {
        if false {
            if v6 {
                var v41 int16 = v1
                fmt.Println(v18)
                v6 = (146 > 82)
                _ = v41
            }
            fmt.Println(v5, v40)
        }
        var v42 string = ("" + v40)
        _ = v42
    } else {
        var v43 int16 = (v1 % ((v1) | 1))
        var v44 int16 = v1
        _ = v43
        _ = v44
    }
    var v45 float64 = (-2.25 * 3.14159)
    if (false || (v4 <= (v4 ^ (v4 - (v4 / (((-(v4))) | 1)))))) {
        fmt.Println(v45, v1)
        fmt.Println(v5, v40, v40)
    }
    if true {
        var v46 int64 = -1
        v47 := GenS3{F0: 127}
        _ = v46
        _ = v47
    } else {
        var v48 uint64 = 1
        var v49 string = (v40 + ((v40 + "") + ""))
        _ = v48
        _ = v49
    }
    if len(v3) > 0 {
        v3[0] = 2
    }
    var v50 uint32 = 0
    v51 := GenS3{F0: 40}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v45
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
