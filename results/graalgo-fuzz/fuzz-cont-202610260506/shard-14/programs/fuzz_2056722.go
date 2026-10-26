package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 string; F1 uint64 }
func genFn0(p0 uint, p1 uint16, p2 uint16) int32 {
    fmt.Println(p0)
    p0 = (p0 % ((p0) | 1))
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int32 = -1
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 328, F1: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -485
    v3 := []int{127, -214, 32767}
    var v4 float64 = 1.5
    var v5 float64 = v4
    var v6 int8 = -94
    switch "" {
    case "foo bar":
        if true {
            v10 := GenS1{F0: -32768, F1: 32767}
            var v11 string = "hello"
            _ = v10
            _ = v11
        }
    case "a":
        v6 = 0
    default:
        v1[4294967295] = -435
    }
    v12 := GenS1{F0: 1000, F1: -899}
    v6 += v6
    fmt.Println(v5, v4)
    _ = len(v1)
    v13 := v1[2147483648]
    v14 := GenS2{F0: "a", F1: 42}
    var v15 int = (v13 & (^(v13)))
    var v16 int32 = genFn0(1, 820, 255)
    var v17 int = v15
    for fi18 := 0; fi18 < 11; fi18++ {
        if (!(2 == 42)) {
            v13 = (v15 + v15)
        } else {
            switch v17 {
            case 2147483647:
                switch ("a" + ("\n\t" + "")) {
                case "":
                    fmt.Println(v5, v16)
                case "foo bar":
                    var v19 string = "a"
                    _ = v19
                case "\n\t":
                    v16 = (v16 % (((v16 ^ (v16 * (-(v16))))) | 1))
                default:
                    v13 = v17
                }
            case 42:
                for fi20 := 0; fi20 < 13; fi20++ {
                    var v21 int = (fi18 + (fi20 % ((v13) | 1)))
                    switch 512 {
                    case 32767:
                        var v22 string = "hello"
                        _ = v22
                    case 0:
                        v23 := GenS2{F0: "foo bar", F1: 0}
                        _ = v23
                    }
                    _ = fi20
                    _ = v21
                }
            }
            v24 := GenS2{F0: "a", F1: 18446744073709551615}
            var v25 uint = 1
            _ = v24
            _ = v25
        }
        if true {
            _ = len(v3)
        } else {
            var v26 int64 = -525
            for fi27 := 0; fi27 < 5; fi27++ {
                var v28 int8 = (v6 / ((v6) | 1))
                v17 = fi27
                var v29 float64 = v4
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = v26
        }
        if (!(!(-2.25 >= v4))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi18
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        v15 = (v15 + fi30)
        _ = fi30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi31
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
