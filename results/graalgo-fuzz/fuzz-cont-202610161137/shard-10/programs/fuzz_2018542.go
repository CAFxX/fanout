package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint64 = 2
    v2 := []int{499, 9223372036854775807, 42, 32768}
    var v3 uint16 = 2
    var v4 string = "\n\t"
    if (!true) {
        switch 127 {
        case 42:
            _ = len(v2)
        case 127:
            v3 -= 33
        case 1:
            for fi5 := 0; fi5 < 10; fi5++ {
                v4 = "foo bar"
                var v6 int = fi5
                fmt.Println(v4, v3, v4)
                _ = fi5
                _ = v6
            }
        }
    } else {
        for fi7 := 0; fi7 < 11; fi7++ {
            if (0.1 <= 1.5) {
                if len(v2) > 0 {
                    v2[0] = fi7
                }
            } else {
                v8 := GenS1{F0: true, F1: (^(v1)), F2: (^(v1))}
                if true {
                    switch "foo bar" {
                    case "\n\t":
                        _ = len(v2)
                    }
                    v9 := GenS1{F0: false, F1: (v1 ^ 42), F2: v1}
                    v3 = v3
                    _ = v9
                }
                _ = v8
            }
            if len(v2) > 0 {
                v2[0] = (fi7 % (((fi7 + 32767)) | 1))
            }
            var v10 uint32 = 1
            _ = fi7
            _ = v10
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v11 float32 = float32((0.1 * (-2.25 + 0.1)))
        _ = v11
    }
    v2 = append(v2, -2147483648)
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint = 2
        fmt.Println(v1, v13)
        _ = fi12
        _ = v13
    }
    var v14 uint32 = 42
    var v15 int8 = 61
    var v16 uint16 = v3
    var v17 uint16 = v16
    switch 127 {
    case 235:
        for fi18 := 0; fi18 < 5; fi18++ {
            if true {
                v4 = (("foo bar" + v4) + v4)
                v19 := GenS2{F0: true}
                v3 += v16
                _ = v19
            }
            v17 += 1
            _ = fi18
        }
    case -408:
        for fi20 := 0; fi20 < 4; fi20++ {
            v2 = append(v2, -2147483648)
            fmt.Println(v4)
            fmt.Println(v14, v17)
            _ = fi20
        }
    default:
        var v21 uint64 = 157
        _ = v21
    }
    if (((v1 * (^(v1))) < v1) && true) {
        for fi22 := 0; fi22 < 19; fi22++ {
            if (!false) {
                fmt.Println(v15)
                var v23 float64 = (0.0 / 1.5)
                var v24 uint32 = v14
                _ = v23
                _ = v24
            }
            _ = fi22
        }
        if (v16 != 326) {
            v1 = v1
            _ = len(v2)
            v25 := GenS1{F0: ((v1 | 2) == v1), F1: v1, F2: (v1 & (v1 & v1))}
            _ = v25
        }
        _ = len(v2)
    }
    if (127 != 128) {
        var v26 float64 = (1.5 * 1e10)
        _ = v26
    } else {
        if len(v2) > 0 {
            v2[0] = 0
        }
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    fmt.Println(v15, v16, v16)
    if (!(false || (v14 <= v14))) {
        var v27 int = -753
        if true {
            v27 += v27
            v17 = v17
        } else {
            v3 = uint16(v14)
            for fi28 := 0; fi28 < 5; fi28++ {
                if ((v15 + (v15 | v15)) != v15) {
                    for fi29 := 0; fi29 < 7; fi29++ {
                        var v30 uint32 = (v14 % ((0) | 1))
                        _ = fi29
                        _ = v30
                    }
                }
                for fi31 := 0; fi31 < 9; fi31++ {
                    if false {
                        var v32 int8 = (v15 | (v15 >> ((847) & 63)))
                        var v33 int = (v27 - fi31)
                        _ = v32
                        _ = v33
                    }
                    var v34 int = (v27 >> ((1) & 63))
                    _ = fi31
                    _ = v34
                }
                _ = fi28
            }
        }
        _ = v27
    } else {
        var v35 bool = true
        v4 = v4
        var v36 uint32 = (v14 | (v14 ^ 148))
        _ = v35
        _ = v36
    }
    _ = len(v2)
    switch 32767 {
    case 866:
        var v37 string = v4
        _ = v37
    default:
        if (-1 != 2147483647) {
            fmt.Println(v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
