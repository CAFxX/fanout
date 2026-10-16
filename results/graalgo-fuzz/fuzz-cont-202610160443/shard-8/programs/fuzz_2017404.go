package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() uint32 {
    v5 := []int{}
    var v6 int = 42
    _ = v5
    _ = v6
    return 568
}

func main() {
    var v1 int32 = 42
    v2 := []int{32768, int(v1), -146}
    var v3 int32 = (v1 * (v1 * v1))
    v4 := []string{"foo bar", (("\n\t" + ("foo bar" + ("hello" + ("foo bar" + "\n\t")))) + "foo bar")}
    v3 = (^(v1))
    var v7 int = 970
    if (!(!true)) {
        fmt.Println(v1, v3, v3)
    } else {
        var v8 float64 = ((((-2.25 + 0.1) - 1.5) - (1.5 - (-2.25 + ((0.1 - 3.14159) + 1e10)))) / 1.0)
        _ = v8
    }
    v9 := GenS1{F0: 2}
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch "foo bar" {
        case "hello":
            _ = len(v4)
        }
        fmt.Println(v7, v1, v1)
    } else {
        if ("" <= ("hello" + "hello")) {
            switch v7 {
            case -335:
                v2 = append(v2, 9223372036854775807)
            default:
                switch "a" {
                case "foo bar":
                    var v10 uint8 = 2
                    _ = v10
                }
            }
        } else {
            v11 := GenS1{F0: 250}
            fmt.Println(v1, v3, v7)
            _ = v11
        }
        for fi12 := 0; fi12 < 2; fi12++ {
            switch "hello" {
            case "hello":
                var v13 uint8 = 127
                _ = v13
            case "a":
                v3 -= (-(v3))
            }
            var v14 float64 = 1e10
            if len(v2) > 0 {
                v2[0] = 32767
            }
            _ = fi12
            _ = v14
        }
    }
    fmt.Println(v3, v3, v3)
    fmt.Println(v1, v1, v7)
    var v15 string = "foo bar"
    if (42 > 59) {
        for fi16 := 0; fi16 < 20; fi16++ {
            v15 = v15
            if false {
                v15 = ("" + v15)
                for fi17 := 0; fi17 < 17; fi17++ {
                    switch v15 {
                    case "foo bar":
                        var v18 uint8 = 42
                        _ = v18
                    case "":
                        _ = len(v2)
                    default:
                        v1 = v3
                    }
                    v19 := GenS1{F0: 0}
                    if true {
                        var v20 uint64 = 0
                        v2 = append(v2, (-(v7)))
                        _ = v20
                    } else {
                        switch v7 {
                        case 568:
                            v3 += (v3 ^ v3)
                        default:
                            v15 = ("\n\t" + v15)
                        }
                        var v21 uint8 = 0
                        _ = v21
                    }
                    var v22 int = (fi17 % ((fi17) | 1))
                    _ = fi17
                    _ = v19
                    _ = v22
                }
                var v23 bool = true
                _ = v23
            }
            if (((1 < 1) || false) && false) {
                fmt.Println(v3, v15)
                var v24 int16 = 1000
                _ = v24
            }
            _ = fi16
        }
        v15 = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
