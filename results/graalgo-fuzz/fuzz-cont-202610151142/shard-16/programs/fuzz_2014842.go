package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint64 }
func genFn0() int16 {
    if false {
        v5 := GenS1{F0: -1}
        switch "" {
        case "hello":
            var v6 uint = 65535
            _ = v6
        default:
            if false {
                var v7 float64 = ((-2.25 - (0.0 / (1e10 / 0.1))) * 3.14159)
                switch ("" + "\n\t") {
                case "hello":
                    fmt.Println(v7)
                default:
                    v8 := GenS1{F0: 42}
                    _ = v8
                }
                if true {
                    for fi9 := 0; fi9 < 9; fi9++ {
                        v10 := GenS1{F0: 42}
                        _ = fi9
                        _ = v10
                    }
                    v7 = v7
                    v11 := GenS1{F0: 9223372036854775807}
                    _ = v11
                } else {
                    switch "hello" {
                    case "":
                        if (0 < 433) {
                            if (true || true) {
                                v12 := GenS1{F0: 0}
                                v7 = (v7 + 3.14159)
                                _ = v12
                            }
                            var v13 uint8 = 235
                            _ = v13
                        }
                    case "foo bar":
                        v14 := GenS1{F0: 0}
                        _ = v14
                    default:
                        fmt.Println(v7)
                    }
                }
                _ = v7
            }
        }
        _ = v5
    } else {
        for fi15 := 0; fi15 < 4; fi15++ {
            fmt.Println(fi15)
            var v16 int16 = 42
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 uint8 = 127
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            switch (fi19 & (fi19 % ((32768) | 1))) {
            case 42:
                fmt.Println(fi19)
            case -432:
                var v20 int32 = -1
                _ = v20
            case -159:
                for fi21 := 0; fi21 < 12; fi21++ {
                    v22 := []int{(^(fi21)), (^(fi21)), fi21, (-(fi21))}
                    _ = fi21
                    _ = v22
                }
            default:
                if false {
                    var v23 bool = true
                    _ = v23
                } else {
                    if (!true) {
                        v24 := []int{}
                        _ = v24
                    }
                }
            }
            var v25 int64 = 1
            fmt.Println(v25, v25)
            _ = fi19
            _ = v25
        }
    }
    v26 := []int{32767, 32768}
    _ = v26
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 252
    var v3 int = 128
    var v4 uint16 = 244
    v27 := GenS1{F0: 42}
    v4 += 65535
    if (v3 != v3) {
        for fi28 := 0; fi28 < 9; fi28++ {
            var v29 string = ""
            _ = fi28
            _ = v29
        }
        _ = len(v1)
        v4 += (v4 ^ v4)
    }
    var v30 uint16 = v4
    if (((!false) && (255 > 471)) || (753 >= 2147483647)) {
        var v31 uint = 939
        _ = v31
    } else {
        fmt.Println(v4, v4)
    }
    v30 = v4
    switch "a" {
    case "hello":
        var v32 int16 = 32767
        _ = v32
    default:
        fmt.Println(v3, v3, v2)
    }
    var v33 int64 = -9223372036854775808
    v1[4294967295] = v3
    v34 := GenS1{F0: (v33 + 0)}
    if true {
        for fi35 := 0; fi35 < 11; fi35++ {
            v30 = (^(v4))
            _ = fi35
        }
        v36 := GenS2{F0: (v4 | 1), F1: 255, F2: 982}
        for fi37 := 0; fi37 < 12; fi37++ {
            v38 := []int{fi37, int(v4), fi37}
            _ = fi37
            _ = v38
        }
        _ = v36
    }
    v33 = (v33 / ((1) | 1))
    v4 -= 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
}
