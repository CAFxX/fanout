package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
func genFn0() uint64 {
    switch ("a" + "\n\t") {
    case "a":
        if (true || false) {
            var v6 uint16 = 964
            _ = v6
        } else {
            for fi7 := 0; fi7 < 8; fi7++ {
                if true {
                    if (!(1 < 998)) {
                        var v8 bool = true
                        v8 = false
                        var v9 string = "\n\t"
                        _ = v8
                        _ = v9
                    }
                }
                var v10 uint = 846
                _ = fi7
                _ = v10
            }
        }
    case "foo bar":
        if true {
            var v11 float32 = ((-2.25 / 3.14159) + float32((1.5 - (0.0 * 1.5))))
            _ = v11
        }
    }
    v12 := GenS1{F0: "foo bar", F1: (-90 != 100)}
    _ = v12
    return 4294967295
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"\n\t"}
    v3 := make(map[int]string)
    var v4 uint16 = 1
    v5 := []uint32{}
    var v13 uint16 = v4
    switch "" {
    case "hello":
        switch ("a" + "hello") {
        case "\n\t":
            for fi14 := 0; fi14 < 8; fi14++ {
                v4 += 65535
                _ = fi14
            }
        case "":
            for fi15 := 0; fi15 < 1; fi15++ {
                v16 := GenS1{F0: ("foo bar" + "hello"), F1: true}
                _ = fi15
                _ = v16
            }
        default:
            var v17 int64 = 1
            _ = v17
        }
    default:
        fmt.Println(v13, v4)
    }
    fmt.Println(v13, v13)
    v13 -= v13
    v18 := GenS1{F0: "\n\t", F1: (0 >= 241)}
    var v19 uint8 = 0
    if false {
        var v20 bool = (1.5 >= -2.25)
        if (-1 > int32(v13)) {
            if (v20 || ((531 > 201) || ((v20 || ((819 > 1) || (2147483647 != 1))) || (((556 != 756) && v20) && v20)))) {
                switch ("hello" + "") {
                case "a":
                    v21 := GenS1{F0: "hello", F1: false}
                    _ = v21
                case "foo bar":
                    v4 += (v4 + 936)
                default:
                    v13 = v13
                }
            } else {
                switch 128 {
                case 4294967295:
                    var v22 uint64 = 2
                    _ = v22
                }
                fmt.Println(v19)
            }
            fmt.Println(v4, v13, v20)
        }
        _ = v20
    } else {
        switch 1 {
        case 2147483647:
            v13 += v4
        }
    }
    var v23 int8 = -49
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
