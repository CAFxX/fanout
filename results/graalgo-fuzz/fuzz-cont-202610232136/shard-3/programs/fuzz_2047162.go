package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
func genFn0(p0 bool, p1 int64, p2 int64) uint64 {
    switch -309 {
    case -2147483648:
        if (((2147483647 < 2147483647) || true) || p0) {
            var v5 uint = 255
            fmt.Println(p2, p1)
            v6 := []int{-9223372036854775808}
            _ = v5
            _ = v6
        }
    case 128:
        fmt.Println(p1, p2, p2)
    case 2147483647:
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 bool = true
            p1 = 42
            _ = fi7
            _ = v8
        }
    }
    var v9 int32 = 2147483647
    var v10 int32 = (v9 / ((v9) | 1))
    p2 = p2
    var v11 uint8 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 2
}

func main() {
    var v1 uint32 = 659
    v2 := make(map[uint32]bool)
    v3 := []string{}
    var v4 float32 = -2.25
    v4 += 3.14159
    v4 += (v4 / v4)
    _ = len(v2)
    for fi12 := 0; fi12 < 2; fi12++ {
        switch "foo bar" {
        case "foo bar":
            var v13 int64 = 42
            _ = v13
        case "\n\t":
            delete(v2, (v1 + (v1 << ((65535) & 63))))
        case "":
            var v14 uint16 = 42
            _ = v14
        default:
            var v15 int16 = 42
            _ = v15
        }
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        var v16 uint8 = 19
        _ = fi12
        _ = v16
    }
    if ((v1 - (v1 % ((0) | 1))) >= (^(v1))) {
        v17 := GenS1{F0: (("hello" + "a") + ("foo bar" + "foo bar")), F1: "hello", F2: ("a" + "")}
        _ = v17
    } else {
        switch 32767 {
        case 377:
            v1 += (^(v1))
        case -723:
            v3 = append(v3, "a")
        default:
            v18 := GenS2{F0: 42}
            _ = v18
        }
        v4 += (v4 * 0.0)
        for fi19 := 0; fi19 < 6; fi19++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi19
        }
    }
    v1 -= (v1 + (v1 | 0))
    var v20 int64 = 1
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 int8 = 68
        if (v22 <= (v22 + (v22 << ((255) & 63)))) {
            if true {
                v3 = append(v3, ("a" + ""))
                if true {
                    v23 := v2[0]
                    var v24 uint32 = v1
                    _ = v23
                    _ = v24
                }
            } else {
                v4 = v4
                v4 = v4
            }
        }
        v4 = float32((0.0 - float64((v4 - -2.25))))
        _ = fi21
        _ = v22
    }
    switch 42 {
    case 494:
        v25 := GenS2{F0: 127}
        _ = v25
    case 9223372036854775807:
        if false {
            var v26 int8 = -1
            _ = v26
        } else {
            if true {
                for fi27 := 0; fi27 < 20; fi27++ {
                    v20 -= -9223372036854775808
                    var v28 int16 = 1000
                    _ = fi27
                    _ = v28
                }
                if ((!((!((v1 >> ((143) & 63)) < 0)) && true)) && true) {
                    var v29 uint64 = 553
                    var v30 uint8 = 227
                    v2[v1] = true
                    _ = v29
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 7; fi31++ {
                        for fi32 := 0; fi32 < 4; fi32++ {
                            v1 += v1
                            _ = fi32
                        }
                        _ = fi31
                    }
                }
            }
            var v33 uint16 = 0
            _ = v33
        }
    case 0:
        if (!false) {
            if (581 >= 110) {
                v34 := GenS2{F0: 1}
                _ = v34
            } else {
                v20 = v20
                switch 128 {
                case 127:
                    v35 := v2[v1]
                    _ = v35
                case -760:
                    var v36 int = 2147483647
                    _ = v36
                default:
                    v2[v1] = ((("\n\t" == (("a" + "\n\t") + "hello")) || (2 <= 0)) && true)
                }
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
            }
            for fi37 := 0; fi37 < 3; fi37++ {
                fmt.Println(v4, v20, v1)
                v38 := GenS3{F0: v20}
                _ = fi37
                _ = v38
            }
        }
    }
    var v39 uint32 = 1
    v39 = 2
    var v40 uint8 = 1
    var v41 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
