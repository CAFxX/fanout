package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint64 }

func main() {
    var v1 uint64 = 2
    v2 := []int{-9223372036854775808}
    var v3 float32 = (1e10 * (0.1 - (-2.25 - ((-2.25 / 3.14159) / 1.0))))
    v4 := make(map[uint32]int)
    var v5 uint = 42
    for fi6 := 0; fi6 < 5; fi6++ {
        switch (("hello" + "a") + "") {
        case "hello":
            _ = len(v2)
        case "\n\t":
            v1 = (v1 ^ 1)
        case "":
            switch "" {
            case "":
                v3 -= (v3 * (v3 * v3))
            case "foo bar":
                switch (fi6 + fi6) {
                case 9223372036854775807:
                    v1 -= (v1 * 1)
                case 143:
                    _ = len(v2)
                case 1:
                    var v7 uint8 = 2
                    _ = v7
                }
            case "hello":
                var v8 uint8 = 127
                _ = v8
            }
        }
        v1 = 482
        for fi9 := 0; fi9 < 15; fi9++ {
            v3 += v3
            _ = fi9
        }
        _ = fi6
    }
    v10 := GenS1{F0: 171, F1: 4294967295}
    v5 = (v5 - (v5 >> ((uint(v1)) & 63)))
    var v11 uint64 = v1
    if (true && false) {
        if (true && ("foo bar" <= ("a" + "hello"))) {
            v5 -= (v5 - v5)
            var v12 int64 = 412
            v13 := GenS2{F0: 2, F1: 843, F2: v11}
            _ = v12
            _ = v13
        } else {
            switch 32768 {
            case -312:
                v14 := GenS2{F0: 42, F1: 438, F2: (v11 >> (((v5 * (v5 + 42))) & 63))}
                _ = v14
            }
            if false {
                v11 = (v11 % ((v1) | 1))
            }
        }
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        if (v3 < 0.0) {
            v16 := v4[2147483647]
            _ = v16
        }
        v5 -= (^(v5))
        _ = fi15
    }
    if ((v1 ^ 2) <= (v11 + v1)) {
        switch 812 {
        case 691:
            var v17 float32 = 3.14159
            _ = v17
        default:
            _ = len(v4)
        }
        fmt.Println(v3)
        var v18 float32 = v3
        _ = v18
    } else {
        fmt.Println(v11, v1, v3)
    }
    v1 = 0
    var v19 int32 = 1
    for fi20 := 0; fi20 < 16; fi20++ {
        fmt.Println(fi20, v1)
        var v21 uint8 = 127
        if (fi20 < (fi20 % (((-(fi20))) | 1))) {
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 float64 = 0.0
                var v24 int32 = v19
                _ = fi22
                _ = v23
                _ = v24
            }
            var v25 uint32 = 0
            _ = v25
        }
        _ = fi20
        _ = v21
    }
    if (v19 == 0) {
        v2 = append(v2, 297)
        v11 = (^(v11))
        var v26 string = ("\n\t" + "")
        _ = v26
    } else {
        for fi27 := 0; fi27 < 6; fi27++ {
            v1 -= (v11 << (((^(v5))) & 63))
            v5 -= 110
            _ = fi27
        }
        if true {
            v3 += v3
            for fi28 := 0; fi28 < 11; fi28++ {
                switch 32768 {
                case -745:
                    _ = len(v2)
                }
                var v29 uint64 = v1
                if false {
                    delete(v4, 2147483648)
                    if true {
                        if ((839 >= fi28) && (542 == 0)) {
                            fmt.Println(v19)
                            fmt.Println(v11, v19)
                            switch fi28 {
                            case -173:
                                v1 -= (v11 * (v11 / (((v11 ^ v1)) | 1)))
                            default:
                                for fi30 := 0; fi30 < 18; fi30++ {
                                    v31 := GenS1{F0: 42, F1: (^(fi28))}
                                    _ = fi30
                                    _ = v31
                                }
                            }
                        }
                        if true {
                            _ = len(v2)
                        } else {
                            var v32 float32 = v3
                            v4[4294967295] = (fi28 / ((fi28) | 1))
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v32
                        }
                    } else {
                        v1 -= (v29 | v11)
                        var v33 int16 = 32767
                        _ = v33
                    }
                    v1 -= 42
                }
                _ = fi28
                _ = v29
            }
            _ = len(v4)
        } else {
            v5 = 4294967295
        }
    }
    var v34 uint32 = 1
    fmt.Println(v11, v34)
    if false {
        var v35 uint = v5
        _ = v35
    } else {
        v5 += 2
        v36 := GenS1{F0: 127, F1: 42}
        for fi37 := 0; fi37 < 4; fi37++ {
            v19 = (^(v19))
            _ = fi37
        }
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
