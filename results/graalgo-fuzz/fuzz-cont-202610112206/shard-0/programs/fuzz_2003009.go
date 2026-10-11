package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int64; F2 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 uint; F2 uint32 }
type GenS8 struct { F0 int32; F1 int8; F2 string }

func main() {
    v1 := make(map[int]int)
    var v2 int = 127
    v3 := make(map[string]string)
    var v4 float32 = 0.0
    fmt.Println(v2, v4)
    v4 -= v4
    v5 := []int{(v2 % (((v2 | (-(v2)))) | 1)), 4294967295}
    if false {
        delete(v3, "foo bar")
        fmt.Println(v4, v2)
        v6 := GenS1{F0: 223, F1: 1, F2: false}
        _ = v6
    } else {
        v2 += (v2 * (v2 >> ((18446744073709551615) & 63)))
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS2{F0: 127}
            v9 := GenS3{F0: (v2 * fi7), F1: (fi7 + fi7)}
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 986, F1: 175, F2: ((!((v2 ^ (v2 | v2)) >= (v2 % ((v2) | 1)))) || (false || false))}
        _ = v10
    }
    var v11 uint16 = 42
    fmt.Println(v4, v4)
    v1[(^(v2))] = (v2 << ((65535) & 63))
    v12 := GenS1{F0: 1, F1: 42, F2: (!(!false))}
    if false {
        var v13 int = (v2 << ((54) & 63))
        fmt.Println(v13, v11, v13)
        _ = v13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        switch fi14 {
        case 32768:
            switch "" {
            case "a":
                for fi15 := 0; fi15 < 2; fi15++ {
                    v16 := GenS4{F0: 2}
                    v17 := GenS5{F0: (fi15 | fi14), F1: -9223372036854775808, F2: 0}
                    v18 := GenS6{F0: "a"}
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
            case "foo bar":
                if true {
                    for fi19 := 0; fi19 < 13; fi19++ {
                        fmt.Println(v4, fi14)
                        _ = fi19
                    }
                } else {
                    v20 := GenS7{F0: -32768, F1: 208, F2: 211}
                    v2 = 635
                    var v21 uint8 = 139
                    _ = v20
                    _ = v21
                }
            case "\n\t":
                v4 -= 1e10
            default:
                v4 += v4
            }
        }
        var v22 uint32 = 958
        v11 -= 42
        _ = fi14
        _ = v22
    }
    if true {
        v2 = v2
        v2 += (v2 - (v2 & (-(v2))))
    } else {
        var v23 int32 = -2147483648
        v2 = int(v11)
        var v24 uint64 = 893
        _ = v23
        _ = v24
    }
    if false {
        var v25 int32 = int32(v2)
        _ = v25
    }
    v26 := GenS8{F0: -1, F1: 1, F2: ((("hello" + "foo bar") + ("\n\t" + "\n\t")) + "hello")}
    for fi27 := 0; fi27 < 6; fi27++ {
        if true {
            fmt.Println(v4)
        } else {
            if len(v5) > 0 {
                v5[0] = (v2 | v2)
            }
        }
        _ = fi27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        var v29 string = (("" + (((("a" + "\n\t") + "\n\t") + ("a" + "")) + "foo bar")) + "a")
        v29 = (v29 + ("a" + "foo bar"))
        _ = fi28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
