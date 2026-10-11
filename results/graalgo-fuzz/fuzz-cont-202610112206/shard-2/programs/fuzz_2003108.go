package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int16, p1 int8) string {
    var v5 uint8 = 2
    var v6 uint32 = 351
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (((("a" + "a") + "hello") + "foo bar") + "a")
}
func genFn1(p0 int8, p1 int, p2 int16) bool {
    if false {
        p0 -= (p0 * (p0 * (p0 >> ((0) & 63))))
    } else {
        for fi7 := 0; fi7 < 3; fi7++ {
            if true {
                if false {
                    fmt.Println(p2, p1, p0)
                }
                if (4294967295 >= 1) {
                    v8 := []int{}
                    fmt.Println(p0, fi7, p1)
                    v9 := GenS1{F0: (p1 * fi7)}
                    _ = v8
                    _ = v9
                }
                switch ("a" + "a") {
                case "foo bar":
                    for fi10 := 0; fi10 < 6; fi10++ {
                        p2 += (p2 >> ((42) & 63))
                        _ = fi10
                    }
                default:
                    var v11 uint32 = 42
                    _ = v11
                }
            }
            p2 += 1000
            _ = fi7
        }
        var v12 int8 = p0
        _ = v12
    }
    v13 := []int{(p1 % ((361) | 1)), -1, p1, -664}
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return (!((-2.25 > -2.25) || (!true)))
}

func main() {
    var v1 string = "a"
    var v2 int64 = 711
    v3 := make(map[uint32]bool)
    v4 := []uint32{163, 927, 2147483647, 4294967295}
    var v14 uint8 = 232
    if (false || ("a" <= "\n\t")) {
        fmt.Println(v14, v2, v1)
        v14 -= (^(v14))
    }
    fmt.Println(v2)
    v1 = "foo bar"
    v1 = ("" + v1)
    var v15 int8 = 42
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 int16 = 20
            _ = fi16
            _ = v17
        }
        var v18 int8 = -128
        if (0 >= -1) {
            switch 4294967295 {
            case 259:
                if false {
                    v1 = (v1 + "hello")
                    var v19 int32 = -2147483648
                    v20 := GenS1{F0: 340}
                    _ = v19
                    _ = v20
                }
            case -9223372036854775808:
                var v21 uint8 = 42
                _ = v21
            case -2147483648:
                var v22 float64 = (1e10 / 1.5)
                _ = v22
            }
        }
        _ = v18
    }
    v2 += v2
    var v23 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
