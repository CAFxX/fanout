package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 uint16 }
func genFn0(p0 int32, p1 int16, p2 int8) int {
    for fi4 := 0; fi4 < 19; fi4++ {
        p0 = (p0 % ((p0) | 1))
        _ = fi4
    }
    v5 := []int{-424}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 float32 = ((((0.1 - (0.0 / float32((3.14159 / 1.0)))) - 0.0) / ((float32((((((-2.25 + (-2.25 / 0.1)) / 1.5) * 0.0) - (0.1 - 3.14159)) * -2.25)) / (3.14159 + 3.14159)) - 0.1)) / 3.14159)
    var v2 string = "a"
    var v3 bool = (0 >= -1)
    var v6 string = "\n\t"
    var v7 bool = v3
    if v7 {
        v6 = v6
        if v3 {
            v3 = v7
            if true {
                v3 = true
                v2 = ("\n\t" + v2)
                v2 = ("\n\t" + "foo bar")
            }
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 uint32 = 42
                v10 := GenS1{F0: 42, F1: (v9 ^ (v9 >> ((42) & 63)))}
                v7 = (!(82 == 158))
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        v1 = (v1 * v1)
    } else {
        if (v3 && (-2147483648 != -2147483648)) {
            v11 := GenS2{F0: 4294967295, F1: 75}
            var v12 bool = v7
            v3 = (86 > 789)
            _ = v11
            _ = v12
        } else {
            switch genFn0(-751, -32768, -1) {
            case -894:
                var v13 int = 128
                _ = v13
            case 32768:
                for fi14 := 0; fi14 < 5; fi14++ {
                    switch -2147483648 {
                    case -907:
                        switch "foo bar" {
                        case "a":
                            for fi15 := 0; fi15 < 19; fi15++ {
                                v16 := []int{fi14, (fi14 / ((4294967295) | 1)), (fi15 & (fi15 - (fi15 + (-(fi15))))), fi14}
                                _ = fi15
                                _ = v16
                            }
                        }
                    case 162:
                        switch (v2 + v2) {
                        case "a":
                            for fi17 := 0; fi17 < 12; fi17++ {
                                v18 := []int{(-(fi17)), fi17, (fi17 - (fi17 / (((^(fi17))) | 1))), fi14}
                                _ = fi17
                                _ = v18
                            }
                        default:
                            for fi19 := 0; fi19 < 8; fi19++ {
                                v6 = v6
                                fmt.Println(v2)
                                _ = fi19
                            }
                        }
                    }
                    v20 := GenS2{F0: 829, F1: 42}
                    _ = fi14
                    _ = v20
                }
            case 175:
                var v21 float64 = ((1.5 - (0.0 + float64(float32(3.14159)))) * 1e10)
                _ = v21
            }
            v3 = v7
            for fi22 := 0; fi22 < 9; fi22++ {
                v23 := GenS3{F0: fi22}
                _ = fi22
                _ = v23
            }
        }
        v2 = (v2 + ("foo bar" + (v6 + ("foo bar" + ((v2 + (v2 + (((("hello" + v2) + (("a" + "a") + "a")) + "foo bar") + "\n\t"))) + v6)))))
    }
    var v24 int8 = 1
    for fi25 := 0; fi25 < 11; fi25++ {
        if v7 {
            var v26 int32 = 42
            _ = v26
        } else {
            for fi27 := 0; fi27 < 13; fi27++ {
                v28 := GenS4{F0: (!(-1 != 0))}
                if v3 {
                    v29 := GenS5{F0: (v6 + v2), F1: 2}
                    v30 := []int{(fi25 | fi25), (fi27 + (fi27 / (((fi25 % ((fi27) | 1))) | 1))), (fi27 - (^(fi27)))}
                    if (v1 > (0.1 + -2.25)) {
                        var v31 int16 = 32767
                        fmt.Println(v6)
                        _ = v31
                    }
                    _ = v29
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
        }
        var v32 uint8 = 0
        v3 = ((v32 >> ((1) & 63)) >= (v32 & v32))
        _ = fi25
        _ = v32
    }
    fmt.Println(v7, v24, v6)
    var v33 uint64 = 718
    if v7 {
        if v3 {
            fmt.Println(v33, v1)
            var v34 float32 = (v1 - (v1 / 1.0))
            _ = v34
        } else {
            v35 := []int{-9223372036854775808}
            _ = v35
        }
        v2 = "hello"
    }
    var v36 uint = 1
    var v37 int16 = 431
    for fi38 := 0; fi38 < 8; fi38++ {
        var v39 float32 = v1
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v6
    _ = v7
}
