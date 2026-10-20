package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint32 }
func genFn0(p0 int, p1 int) int8 {
    var v7 uint16 = 2
    p0 = p1
    if true {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint8 = 143
            _ = fi8
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS1{F0: 214, F1: 2}
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return -128
}

func main() {
    var v1 uint = 2
    var v2 int64 = -1
    v3 := []int{128, 32767}
    var v4 bool = (203 > 1)
    var v5 float64 = 1e10
    var v6 int = -154
    fmt.Println(v2, v6)
    if ((v4 && v4) && ("foo bar" != ((((("hello" + "\n\t") + (("\n\t" + "") + ("\n\t" + ("hello" + "\n\t")))) + "\n\t") + (("\n\t" + "a") + ((((((((("" + "") + "foo bar") + "") + "a") + ("\n\t" + ("foo bar" + "a"))) + "hello") + ("\n\t" + "")) + "a") + "\n\t"))) + "foo bar"))) {
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 int32 = 301
            var v14 int64 = v2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    switch -2147483648 {
    case 32767:
        var v15 uint8 = 1
        _ = v15
    }
    var v16 int64 = v2
    switch ("" + "hello") {
    case "\n\t":
        v17 := GenS2{F0: 2147483647, F1: 2}
        _ = v17
    case "hello":
        v6 = v6
    case "":
        switch 775 {
        case -1:
            v16 -= 9223372036854775807
        case 87:
            v18 := GenS2{F0: 2147483648, F1: 0}
            _ = v18
        case 4294967295:
            var v19 string = ""
            _ = v19
        default:
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 string = "hello"
                if (v5 < 1e10) {
                    v6 += fi20
                } else {
                    for fi22 := 0; fi22 < 1; fi22++ {
                        v23 := GenS3{F0: 0, F1: -128}
                        if v4 {
                            _ = len(v3)
                            _ = len(v3)
                            var v24 int = v6
                            _ = v24
                        }
                        _ = fi22
                        _ = v23
                    }
                }
                for fi25 := 0; fi25 < 17; fi25++ {
                    v26 := GenS4{F0: -1}
                    v3 = append(v3, fi20)
                    fmt.Println(v4, v21, v1)
                    v21 = v21
                    var v27 float64 = (v5 * ((v5 + ((0.0 / ((v5 * (((v5 * v5) - v5) * v5)) - 1e10)) + 0.1)) / 3.14159))
                    v1 = (v1 << ((v1) & 63))
                    if true {
                        var v28 int16 = int16(v16)
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = fi20
                _ = v21
            }
        }
    default:
        fmt.Println(v16, v4)
    }
    var v29 float64 = float64(3.14159)
    var v30 int16 = 1000
    v31 := GenS4{F0: (v30 | v30)}
    v1 += v1
    fmt.Println(v2, v1, v4)
    var v32 uint32 = 2
    v33 := GenS5{F0: 4294967295, F1: int64(v6), F2: v32}
    v34 := GenS2{F0: v32, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
