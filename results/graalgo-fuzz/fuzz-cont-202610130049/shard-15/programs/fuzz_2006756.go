package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int8 }
func genFn0(p0 string, p1 int16, p2 int16) uint8 {
    v6 := []int{2147483647, 127, 9223372036854775807}
    fmt.Println(p0, p1, p0)
    v7 := GenS1{F0: 110, F1: 2, F2: 32767}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 143
}

func main() {
    var v1 string = "foo bar"
    var v2 uint = 0
    var v3 uint8 = 127
    v4 := []string{v1}
    v5 := []string{v1, v1, ("" + ("hello" + v1))}
    if len(v4) > 0 {
        v4[0] = "a"
    }
    fmt.Println(v1, v3)
    var v8 int32 = -2147483648
    var v9 float32 = ((0.1 / 0.1) + ((3.14159 * 1e10) - (1e10 * -2.25)))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v10 := GenS1{F0: 42, F1: 477, F2: 2}
    if (!true) {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint = v2
            _ = fi11
            _ = v12
        }
    }
    v13 := GenS2{F0: 0, F1: 2147483647}
    if (v8 != -146) {
        if len(v5) > 0 {
            v5[0] = ("a" + v1)
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            switch "foo bar" {
            case "hello":
                if (!true) {
                    v4 = append(v4, "a")
                    switch (fi14 % (((^(fi14))) | 1)) {
                    case 0:
                        v15 := GenS3{F0: 1, F1: 2}
                        _ = v15
                    case -2147483648:
                        v9 = ((v9 * float32((1e10 * (0.0 * (-2.25 - 1e10))))) / (v9 - ((v9 * v9) + v9)))
                    }
                    v5 = append(v5, v1)
                }
            case "":
                if ((-(v8)) < (v8 >> ((v2) & 63))) {
                    v8 = v8
                    v9 = (v9 / (-2.25 + float32(((1e10 + (1e10 * 1.5)) + 1.5))))
                    v8 = v8
                } else {
                    v1 = ""
                }
            case "a":
                if true {
                    var v16 uint8 = (v3 * genFn0("\n\t", -870, -1))
                    var v17 int32 = 93
                    v3 -= (v16 + (^(v3)))
                    _ = v16
                    _ = v17
                } else {
                    var v18 int = (fi14 * 9223372036854775807)
                    _ = v18
                }
            default:
                var v19 uint64 = 829
                _ = v19
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := GenS2{F0: 4, F1: 555}
                v22 := GenS4{F0: 42}
                _ = fi20
                _ = v21
                _ = v22
            }
            v3 += (v3 << ((65535) & 63))
            _ = fi14
        }
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        v1 = "a"
        var v24 uint8 = v3
        fmt.Println(v3)
        _ = fi23
        _ = v24
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v25 uint16 = 952
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
