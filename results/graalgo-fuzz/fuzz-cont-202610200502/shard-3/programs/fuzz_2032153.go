package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint32; F2 string }
func genFn0(p0 uint64, p1 bool, p2 int16) int64 {
    switch -15 {
    case 127:
        v4 := []int{}
        _ = v4
    default:
        v5 := GenS1{F0: 1, F1: p0}
        _ = v5
    }
    switch "\n\t" {
    case "\n\t":
        p1 = (-912 >= 1)
    case "foo bar":
        if (490 <= 2147483647) {
            var v6 float32 = (1e10 - 0.1)
            switch 128 {
            case 0:
                p1 = (286 <= 893)
            case 9223372036854775807:
                var v7 uint16 = 0
                _ = v7
            }
            _ = v6
        } else {
            switch 32767 {
            case 4294967295:
                var v8 string = ("" + ((("\n\t" + "") + "foo bar") + "hello"))
                _ = v8
            case 9223372036854775807:
                var v9 int64 = -569
                _ = v9
            }
            if true {
                for fi10 := 0; fi10 < 19; fi10++ {
                    v11 := GenS1{F0: -795, F1: 0}
                    fmt.Println(p1)
                    _ = fi10
                    _ = v11
                }
            } else {
                v12 := GenS1{F0: -807, F1: p0}
                _ = v12
            }
            p0 = 4294967295
        }
    default:
        p0 = 137
    }
    v13 := []int{2147483647, -1, -2147483648}
    if p1 {
        if len(v13) > 0 {
            v13[0] = -2147483648
        }
        var v14 int64 = 9223372036854775807
        switch 392 {
        case 989:
            switch ("hello" + "a") {
            case "foo bar":
                for fi15 := 0; fi15 < 5; fi15++ {
                    switch fi15 {
                    case 2147483647:
                        if p1 {
                            if false {
                                fmt.Println(p1)
                                var v16 int64 = (-(v14))
                                fmt.Println(p2, v16, p1)
                                _ = v16
                            }
                            v13 = append(v13, -1)
                            _ = len(v13)
                        }
                    }
                    _ = fi15
                }
            case "":
                _ = len(v13)
            }
        }
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    var v1 float64 = 1.5
    var v2 float64 = 0.1
    v3 := make(map[int]string)
    v2 = 0.1
    if (0 != -76) {
        v17 := GenS2{F0: false, F1: 593, F2: ("foo bar" + ((("a" + ("" + "foo bar")) + "") + "foo bar"))}
        _ = v17
    } else {
        v18 := GenS2{F0: (3.14159 >= (v2 - v2)), F1: 1, F2: ((("hello" + "hello") + ("foo bar" + ("" + ("" + "a")))) + ((((("a" + "hello") + "") + "") + "a") + "foo bar"))}
        _ = v18
    }
    var v19 bool = ((false || (-111 == -1)) && (!(false && (32767 == 65535))))
    var v20 uint16 = 32767
    if true {
        var v21 bool = ((v19 && v19) && v19)
        if (((0.0 / (1e10 + 0.0)) - 3.14159) != 0.0) {
            var v22 int16 = 1000
            _ = v22
        }
        _ = v21
    } else {
        v19 = v19
        for fi23 := 0; fi23 < 8; fi23++ {
            v24 := []int{-9223372036854775808, fi23, (fi23 * 32768)}
            if len(v24) > 0 {
                v24[0] = (fi23 + fi23)
            }
            _ = fi23
            _ = v24
        }
    }
    var v25 string = "foo bar"
    if (!(v1 > (3.14159 / (v1 - v1)))) {
        fmt.Println(v1, v1)
        fmt.Println(v2)
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 float64 = (float64((float32(3.14159) - (float32(v1) * 1e10))) + float64((((1.5 / 3.14159) - 3.14159) + float32(((v1 + v2) - v1)))))
        v25 = (v25 + (("foo bar" + "") + (v25 + ("\n\t" + v25))))
        var v28 int = (fi26 / (((fi26 ^ (-(fi26)))) | 1))
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v25, v20, v19)
    v25 = (v25 + v25)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
}
