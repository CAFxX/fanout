package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int; F1 int8; F2 uint16 }

func main() {
    v1 := []string{"\n\t", "", "foo bar"}
    var v2 uint = 950
    var v3 uint = v2
    var v4 string = "foo bar"
    var v5 int32 = -1
    var v6 float64 = (-2.25 / 1.5)
    switch 660 {
    case -480:
        switch int(v5) {
        case 2147483647:
            var v7 uint8 = 102
            _ = v7
        case 838:
            switch 611 {
            case -722:
                switch 0 {
                case 285:
                    _ = len(v1)
                case 4294967295:
                    if (127 < -58) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v8 uint64 = 42
                        for fi9 := 0; fi9 < 17; fi9++ {
                            v8 = 1
                            _ = fi9
                        }
                        _ = v8
                    }
                case -2147483648:
                    var v10 bool = (42 > -1)
                    _ = v10
                }
            case 42:
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 bool = (((false && (1.5 >= -2.25)) && (0 < 18446744073709551615)) || ((-32768 < 42) || (((("a" + (v4 + "")) + (("hello" + v4) + v4)) + "hello") == v4)))
                    v4 = ""
                    _ = fi11
                    _ = v12
                }
            }
        }
    case 609:
        var v13 uint32 = 177
        _ = v13
    }
    var v14 int32 = v5
    if (((v2 * v2) < v2) && true) {
        if (245 <= 2) {
            fmt.Println(v14, v5)
            v5 += 42
        }
        fmt.Println(v6, v5, v5)
    } else {
        v15 := GenS1{F0: 42, F1: 127}
        _ = v15
    }
    fmt.Println(v5, v14)
    switch "foo bar" {
    case "":
        switch -958 {
        case 2147483647:
            v16 := GenS2{F0: 4294967295, F1: -1, F2: 2}
            _ = v16
        default:
            if false {
                v5 += v5
                var v17 uint16 = 858
                _ = v17
            }
        }
    }
    v5 += v14
    if len(v1) > 0 {
        v1[0] = v4
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            fmt.Println(v14)
            _ = fi19
        }
        v2 -= (v3 + (v3 / ((v2) | 1)))
        _ = fi18
    }
    v20 := GenS2{F0: -1, F1: -1, F2: 65535}
    for fi21 := 0; fi21 < 11; fi21++ {
        v1 = append(v1, v4)
        var v22 int8 = -93
        _ = fi21
        _ = v22
    }
    fmt.Println(v5, v3, v4)
    v4 = (("a" + "a") + ((v4 + "a") + (v4 + v4)))
    fmt.Println(v6, v14, v5)
    var v23 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
