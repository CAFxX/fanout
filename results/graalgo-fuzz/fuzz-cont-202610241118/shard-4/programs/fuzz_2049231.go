package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint8; F2 int64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 uint8 = 244
    v4 := make(map[uint32]uint32)
    var v5 string = "foo bar"
    var v6 uint32 = 333
    v5 = "hello"
    fmt.Println(v6)
    v7 := GenS1{F0: ((0.0 / 1.5) >= 3.14159), F1: 1, F2: 1}
    if (false && true) {
        v3 += uint8(v6)
        v6 -= v6
    }
    v5 = v5
    for fi8 := 0; fi8 < 18; fi8++ {
        _ = len(v2)
        _ = fi8
    }
    var v9 float32 = (-2.25 - (1.5 + float32(((0.0 / 3.14159) + (1e10 - (1e10 * 0.1))))))
    v10 := GenS2{F0: (!false)}
    delete(v4, v6)
    for fi11 := 0; fi11 < 10; fi11++ {
        if false {
            v12 := GenS2{F0: (0 <= 740)}
            for fi13 := 0; fi13 < 6; fi13++ {
                v3 -= 1
                v14 := GenS3{F0: 1, F1: (v3 & (^(v3))), F2: -9223372036854775808}
                _ = fi13
                _ = v14
            }
            delete(v1, (v5 + "a"))
            _ = v12
        } else {
            for fi15 := 0; fi15 < 4; fi15++ {
                switch ("\n\t" + (v5 + ("foo bar" + ("\n\t" + v5)))) {
                case "a":
                    v16 := []int{(fi11 | fi15), -2147483648}
                    _ = v16
                case "\n\t":
                    v17 := GenS1{F0: true, F1: 1, F2: uint16(v6)}
                    _ = v17
                }
                switch v5 {
                case "a":
                    var v18 int8 = 1
                    _ = v18
                case "foo bar":
                    fmt.Println(v3, v3, v9)
                    v3 = 79
                }
                var v19 int32 = 1
                if true {
                    fmt.Println(v9, fi15, fi15)
                } else {
                    var v20 bool = (759 <= 95)
                    if v20 {
                        v21 := GenS1{F0: true, F1: -128, F2: 335}
                        fmt.Println(v5, fi11, v3)
                        _ = v21
                    }
                    _ = v20
                }
                _ = fi15
                _ = v19
            }
        }
        fmt.Println(v3, v9)
        _ = fi11
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 float64 = 0.1
        var v24 int64 = -1
        if (0 > 42) {
            if false {
                v5 = "foo bar"
                v24 = v24
                if (!false) {
                    _ = len(v2)
                    for fi25 := 0; fi25 < 6; fi25++ {
                        var v26 int = (fi22 ^ -2147483648)
                        _ = fi25
                        _ = v26
                    }
                }
            } else {
                v9 -= v9
                delete(v2, (-(fi22)))
            }
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    switch ("foo bar" + "\n\t") {
    case "a":
        switch 506 {
        case 8:
            v27 := []int{}
            _ = v27
        case 32767:
            v3 -= 247
        case -1000:
            fmt.Println(v6, v9)
        }
    case "\n\t":
        fmt.Println(v5, v9, v5)
    case "":
        fmt.Println(v3)
    default:
        v28 := []int{32768, 2147483647}
        _ = v28
    }
    if false {
        v3 += v3
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 uint = 2
        v9 += (v9 / 1.0)
        for fi31 := 0; fi31 < 2; fi31++ {
            v3 = 42
            v32 := []int{fi29}
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    if true {
        v33 := GenS2{F0: (v3 <= 0)}
        if ((!true) && (963 > 191)) {
            var v34 uint64 = uint64(v3)
            _ = v34
        } else {
            var v35 uint32 = 42
            _ = v35
        }
        fmt.Println(v3, v6, v3)
        _ = v33
    } else {
        v5 = "a"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
