package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint }
type GenS2 struct { F0 uint8; F1 int; F2 int64 }

func main() {
    var v1 uint = 606
    var v2 uint64 = 1
    var v3 float64 = 1.5
    v4 := make(map[uint32]string)
    for fi5 := 0; fi5 < 8; fi5++ {
        v2 = v2
        var v6 int8 = -7
        _ = fi5
        _ = v6
    }
    switch "a" {
    case "foo bar":
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 uint32 = 2147483648
            var v9 string = "hello"
            var v10 string = v9
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    case "hello":
        v1 = v1
    }
    v11 := []int{2147483647, 127, 127, 127}
    v4[1] = ("" + "a")
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 int32 = -466
        _ = fi12
        _ = v13
    }
    var v14 uint64 = (v2 * v2)
    v14 = (v14 * v2)
    if (false && true) {
        var v15 string = "\n\t"
        v2 += v2
        _ = v15
    }
    var v16 int16 = 1000
    v14 -= (v2 + 0)
    v14 = v2
    if true {
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := GenS1{F0: "a", F1: -1, F2: v1}
            v3 = v3
            v19 := GenS2{F0: 127, F1: fi17, F2: 9223372036854775807}
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    if (9223372036854775807 > 1) {
        v3 -= ((-2.25 / 1e10) - -2.25)
        delete(v4, 2147483647)
    }
    var v20 uint16 = 32767
    if false {
        var v21 float32 = 0.1
        switch 9223372036854775807 {
        case -2147483648:
            fmt.Println(v1, v21)
        case 330:
            v4[2] = ("hello" + "\n\t")
        default:
            v11 = append(v11, 32768)
        }
        _ = v21
    } else {
        if true {
            v14 = v2
        } else {
            if (v2 == v2) {
                switch ("\n\t" + ("" + "")) {
                case "a":
                    var v22 uint64 = v14
                    _ = v22
                case "foo bar":
                    var v23 int8 = 42
                    _ = v23
                }
            }
            if true {
                fmt.Println(v20, v16, v1)
                var v24 string = "hello"
                if (v20 >= v20) {
                    fmt.Println(v16)
                } else {
                    for fi25 := 0; fi25 < 15; fi25++ {
                        for fi26 := 0; fi26 < 5; fi26++ {
                            v20 = (v20 / ((v20) | 1))
                            _ = fi26
                        }
                        _ = fi25
                    }
                }
                _ = v24
            }
        }
        v11 = append(v11, 416)
        fmt.Println(v20, v2)
    }
    v20 -= v20
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v14
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
