package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    var v1 int8 = 45
    var v2 uint = 4294967295
    v3 := make(map[string]int)
    v4 := make(map[string]int)
    v5 := []string{"hello", "hello", "\n\t"}
    for fi6 := 0; fi6 < 14; fi6++ {
        v5 = append(v5, "\n\t")
        _ = fi6
    }
    var v7 string = "hello"
    v8 := GenS1{F0: 926, F1: (1 <= 255)}
    v7 = ("foo bar" + v7)
    v9 := GenS1{F0: 42, F1: (3.14159 >= 1.5)}
    v1 += v1
    v2 += v2
    switch 1 {
    case -875:
        v2 = 255
    default:
        v10 := GenS1{F0: -643, F1: (true || false)}
        _ = v10
    }
    v11 := v4[v7]
    v1 += (v1 << ((v2) & 63))
    if ((977 == 1) && true) {
        if ((v1 * (v1 + v1)) == v1) {
            fmt.Println(v2, v1)
            v12 := GenS1{F0: 1, F1: (false && false)}
            switch 32768 {
            case 32767:
                v13 := GenS1{F0: -9223372036854775808, F1: (1e10 == ((float64(-2.25) * ((1.5 - ((0.1 + 0.0) + (1.5 - (0.0 / 1.5)))) / 1.0)) / 1e10))}
                _ = v13
            case 42:
                var v14 bool = (false || false)
                _ = v14
            }
            _ = v12
        }
        v5 = append(v5, v7)
        v3[v7] = v11
    }
    var v15 int8 = 127
    var v16 uint16 = 705
    if false {
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 uint = v2
            if true {
                var v19 uint8 = 100
                _ = v19
            } else {
                if (true && false) {
                    if true {
                        var v20 bool = false
                        _ = v20
                    } else {
                        if false {
                            v11 = (v11 + (fi17 - v11))
                        } else {
                            v5 = append(v5, ((v7 + ((v7 + v7) + "\n\t")) + "hello"))
                        }
                        var v21 int64 = 9223372036854775807
                        _ = v21
                    }
                    switch v11 {
                    case 604:
                        v22 := GenS1{F0: 412, F1: ((!(true && ((v7 >= "\n\t") || false))) || (128 > 42))}
                        _ = v22
                    case 703:
                        var v23 uint = v18
                        _ = v23
                    default:
                        v15 -= v1
                    }
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v7 = (v7 + ((v7 + v7) + ""))
                    switch fi17 {
                    case -1:
                        switch fi17 {
                        case 2147483647:
                            fmt.Println(v1)
                        case 32767:
                            v24 := GenS1{F0: 1, F1: false}
                            _ = v24
                        case 42:
                            var v25 int64 = 9223372036854775807
                            _ = v25
                        }
                    case 2147483647:
                        var v26 string = "hello"
                        _ = v26
                    }
                }
            }
            _ = fi17
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
