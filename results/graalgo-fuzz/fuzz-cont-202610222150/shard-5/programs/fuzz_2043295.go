package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 bool; F1 string; F2 string }
type GenS3 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 float32 = ((-2.25 + float32((float64(float32(-2.25)) - 1.5))) / -2.25)
    v2 := make(map[int]int)
    var v3 int64 = 1
    var v4 int64 = v3
    var v5 uint8 = 174
    v6 := GenS1{F0: v5, F1: true}
    delete(v2, 1)
    fmt.Println(v5, v1)
    if false {
        fmt.Println(v1, v5, v3)
        v7 := []int{32768, 2147483647, 238}
        _ = v7
    } else {
        var v8 int = -384
        if true {
            v9 := []int{1, (-(v8))}
            if false {
                if len(v9) > 0 {
                    v9[0] = 2147483647
                }
            }
            v4 += v3
            _ = v9
        } else {
            var v10 uint = 0
            fmt.Println(v3)
            _ = v10
        }
        if false {
            v11 := v2[(v8 * v8)]
            _ = v11
        } else {
            v12 := v2[(-(v8))]
            _ = v12
        }
        _ = v8
    }
    v4 += (v3 << ((803) & 63))
    v3 = (v3 % (((-(v3))) | 1))
    v4 = (v3 + 42)
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(fi13)
        v5 -= (v5 / ((v5) | 1))
        _ = fi13
    }
    if (!false) {
        if (v1 >= 3.14159) {
            fmt.Println(v1)
            var v14 uint16 = 1
            switch "" {
            case "":
                v14 -= v14
            case "\n\t":
                var v15 uint = 4294967295
                _ = v15
            case "hello":
                for fi16 := 0; fi16 < 5; fi16++ {
                    switch -1 {
                    case 344:
                        var v17 float64 = (1.5 * -2.25)
                        _ = v17
                    case 2147483647:
                        v5 -= 1
                    case 0:
                        v2[(^(fi16))] = fi16
                    default:
                        v18 := GenS2{F0: (true || (!(true && true))), F1: "\n\t", F2: ("a" + "hello")}
                        _ = v18
                    }
                    var v19 int8 = 0
                    var v20 uint16 = 1
                    _ = fi16
                    _ = v19
                    _ = v20
                }
            default:
                if ((-2147483648 == 0) || ((3.14159 - 1e10) != 0.1)) {
                    var v21 uint16 = v14
                    _ = v21
                }
            }
            _ = v14
        }
        v3 = 9223372036854775807
        switch "\n\t" {
        case "":
            v5 = v5
        case "a":
            v22 := GenS1{F0: (v5 >> ((556) & 63)), F1: false}
            _ = v22
        }
    } else {
        var v23 string = ((("foo bar" + "hello") + ("hello" + (("a" + "a") + "hello"))) + "a")
        for fi24 := 0; fi24 < 15; fi24++ {
            v25 := GenS3{F0: (^(v5)), F1: (v5 ^ v5)}
            _ = fi24
            _ = v25
        }
        _ = len(v2)
        _ = v23
    }
    v1 += (float32((float64(1e10) + 0.1)) / (v1 + 3.14159))
    switch -2147483648 {
    case 211:
        for fi26 := 0; fi26 < 14; fi26++ {
            if true {
                if true {
                    for fi27 := 0; fi27 < 11; fi27++ {
                        var v28 float64 = 0.0
                        var v29 int16 = 42
                        v29 = 1000
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
                fmt.Println(fi26, v4, v5)
            } else {
                var v30 float32 = v1
                v31 := []int{(fi26 ^ (fi26 + (fi26 ^ -9223372036854775808))), (fi26 ^ (fi26 / ((fi26) | 1))), fi26, (fi26 & fi26)}
                v5 -= v5
                _ = v30
                _ = v31
            }
            v32 := GenS2{F0: (((!true) && (1 > 330)) || false), F1: ("hello" + "hello"), F2: "hello"}
            if true {
                v33 := []int{}
                v5 -= v5
                _ = v33
            } else {
                v3 = (v3 | v3)
                var v34 uint8 = (v5 & v5)
                _ = v34
            }
            _ = fi26
            _ = v32
        }
    case 2147483647:
        v35 := GenS3{F0: 128, F1: (^(v5))}
        _ = v35
    default:
        fmt.Println(v5, v1)
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        v4 -= v3
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
