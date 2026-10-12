package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int, p1 int16, p2 uint) uint16 {
    v6 := []int{p0, p0, p0}
    switch ("a" + ("" + "foo bar")) {
    case "\n\t":
        switch ("foo bar" + "\n\t") {
        case "a":
            if false {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v7 := GenS1{F0: true, F1: 1}
                if (3.14159 >= ((0.0 / 1.0) * 1e10)) {
                    if true {
                        p0 += p0
                    }
                } else {
                    fmt.Println(p2)
                    p0 = p0
                }
                _ = v7
            }
        case "hello":
            var v8 uint8 = 45
            _ = v8
        case "\n\t":
            if (2 == 40) {
                p0 -= p0
                if true {
                    var v9 uint16 = 2
                    _ = v9
                }
            } else {
                var v10 int8 = 42
                _ = v10
            }
        }
    case "hello":
        p0 += (p0 + p0)
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(fi11, p0, fi11)
        var v12 uint8 = 254
        _ = fi11
        _ = v12
    }
    fmt.Println(p1, p1)
    if (((-1 < -128) || true) && true) {
        fmt.Println(p0)
        v6 = append(v6, (p0 >> ((p2) & 63)))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 255
}

func main() {
    v1 := []int{}
    var v2 uint = 193
    var v3 int64 = 0
    var v4 float32 = 3.14159
    v5 := make(map[string]int)
    fmt.Println(v4, v4)
    var v13 int = 0
    fmt.Println(v3, v13, v2)
    fmt.Println(v2)
    _ = len(v1)
    var v14 int32 = 42
    fmt.Println(v3, v3)
    if (1 == 100) {
        fmt.Println(v2, v4)
        var v15 int = (^(v13))
        _ = v15
    } else {
        v3 -= v3
        v16 := GenS2{F0: 4294967295, F1: 2147483647}
        _ = v16
    }
    if (127 > v13) {
        for fi17 := 0; fi17 < 1; fi17++ {
            v13 = fi17
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 bool = (!(true || (475 < 394)))
                var v20 uint = v2
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 float32 = (3.14159 / v4)
            _ = fi17
            _ = v21
        }
        v14 = int32(v13)
        if true {
            v4 = v4
            if false {
                var v22 uint = 65535
                switch 42 {
                case -9223372036854775808:
                    var v23 float64 = (-2.25 / 1.0)
                    _ = v23
                }
                _ = v22
            } else {
                for fi24 := 0; fi24 < 19; fi24++ {
                    v4 = (((v4 / 3.14159) - v4) / v4)
                    v5["a"] = (fi24 << ((51) & 63))
                    _ = fi24
                }
                v25 := GenS2{F0: 0, F1: 2147483648}
                _ = v25
            }
        }
    }
    v2 = (v2 + v2)
    var v26 uint = 265
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
