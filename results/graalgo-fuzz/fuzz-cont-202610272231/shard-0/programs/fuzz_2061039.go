package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int16 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 bool }
func genFn0(p0 int, p1 uint, p2 string) int16 {
    if (!((((0 > 2147483647) || ((!(42 < 42)) || (0 >= 0))) && true) && (true || (true && (!true))))) {
        for fi6 := 0; fi6 < 14; fi6++ {
            switch fi6 {
            case 32768:
                p1 -= p1
            }
            p1 -= (^(p1))
            if false {
                p2 = p2
                var v7 bool = (2 <= 0)
                p2 = p2
                _ = v7
            }
            _ = fi6
        }
    } else {
        if false {
            v8 := GenS1{F0: 0, F1: 1, F2: 116}
            var v9 uint64 = 4294967295
            _ = v8
            _ = v9
        } else {
            fmt.Println(p1, p1)
        }
        v10 := GenS1{F0: 0, F1: -1, F2: 0}
        var v11 uint32 = 2147483648
        _ = v10
        _ = v11
    }
    if (false || (p0 >= (^(p0)))) {
        v12 := GenS2{F0: 65535, F1: -890}
        v13 := []int{(p0 % (((p0 >> ((0) & 63))) | 1)), p0, (^(p0))}
        v14 := GenS3{F0: 2}
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 int8 = -1
    v16 := []int{p0, p0, -9223372036854775808, 1}
    if true {
        var v17 int = p0
        fmt.Println(p2, p0, p0)
        _ = v17
    } else {
        p0 -= (^(p0))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 749
}
func genFn1(p0 uint) bool {
    v18 := GenS4{F0: (p0 ^ p0), F1: (^(p0))}
    fmt.Println(p0)
    for fi19 := 0; fi19 < 13; fi19++ {
        switch ("foo bar" + "\n\t") {
        case "a":
            v20 := GenS3{F0: 1}
            _ = v20
        default:
            var v21 int8 = 0
            _ = v21
        }
        _ = fi19
    }
    _ = p0
    _ = v18
    return ((42 >= 1) || (false || (genFn0(-630, (p0 ^ 18446744073709551615), "foo bar") < genFn0(-2147483648, 18446744073709551615, ("\n\t" + ("a" + "\n\t"))))))
}

func main() {
    var v1 int32 = 1
    v2 := []int{374, 128, 42, 42}
    var v3 string = ""
    var v4 uint = 18446744073709551615
    var v5 int = 2147483647
    fmt.Println(v1, v5)
    v1 -= (v1 - 313)
    if true {
        var v22 int8 = -36
        if false {
            fmt.Println(v3)
            if len(v2) > 0 {
                v2[0] = v5
            }
        }
        v1 = -2147483648
        _ = v22
    }
    v3 = v3
    switch (v5 >> (((^(v4))) & 63)) {
    case 622:
        if true {
            switch ((v3 + (v3 + v3)) + v3) {
            case "foo bar":
                v23 := GenS5{F0: (127 >= 64)}
                _ = v23
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v5 = v5
            switch (v5 / ((v5) | 1)) {
            case 9223372036854775807:
                v1 = v1
            case 202:
                var v24 uint = v4
                _ = v24
            default:
                v25 := GenS5{F0: (255 <= 951)}
                _ = v25
            }
        }
    case -320:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 42:
        switch "hello" {
        case "hello":
            for fi26 := 0; fi26 < 3; fi26++ {
                switch v5 {
                case -1:
                    fmt.Println(v5, v1, v4)
                }
                if len(v2) > 0 {
                    v2[0] = (v5 >> ((v4) & 63))
                }
                v3 = (v3 + ("hello" + v3))
                _ = fi26
            }
        }
    }
    switch ("\n\t" + v3) {
    case "hello":
        for fi27 := 0; fi27 < 4; fi27++ {
            v3 = "\n\t"
            v28 := GenS6{F0: -9223372036854775808, F1: 1}
            _ = fi27
            _ = v28
        }
    default:
        v29 := GenS2{F0: 4294967295, F1: (v1 >> ((42) & 63))}
        _ = v29
    }
    v1 = 521
    var v30 float64 = (((-2.25 * (3.14159 / 1.0)) - (-2.25 + (3.14159 - ((3.14159 + (0.1 * -2.25)) - 0.0)))) - 0.0)
    v31 := GenS7{F0: (128 == 218)}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    v30 += v30
    v3 = ("a" + v3)
    var v32 uint = (^(v4))
    _ = len(v2)
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
