package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int64; F1 uint16; F2 int }
type GenS3 struct { F0 int }
func genFn0() int32 {
    v5 := []int{0, 9223372036854775807, 362}
    var v6 uint64 = 4294967295
    switch "foo bar" {
    case "hello":
        if (0.1 < (-2.25 - float32((((float64(1e10) - 0.0) + (0.1 * 3.14159)) / float64((0.1 - ((((0.1 / 1e10) * (float32(3.14159) * -2.25)) - -2.25) * float32(1.5)))))))) {
            var v7 int64 = -991
            for fi8 := 0; fi8 < 15; fi8++ {
                v6 = (v6 & 1)
                var v9 int32 = -1
                v6 = (v6 >> ((143) & 63))
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 18; fi10++ {
                for fi11 := 0; fi11 < 4; fi11++ {
                    switch ("hello" + "") {
                    case "\n\t":
                        v7 += -1
                    }
                    v6 = v6
                    v5 = append(v5, (fi11 >> ((4294967295) & 63)))
                    _ = fi11
                }
                for fi12 := 0; fi12 < 16; fi12++ {
                    fmt.Println(v7)
                    v13 := GenS1{F0: 42, F1: v7}
                    v5 = append(v5, (fi12 >> ((180) & 63)))
                    _ = fi12
                    _ = v13
                }
                _ = fi10
            }
            _ = v7
        }
    case "\n\t":
        for fi14 := 0; fi14 < 7; fi14++ {
            _ = len(v5)
            _ = fi14
        }
    case "foo bar":
        v6 = (v6 % ((210) | 1))
    }
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int = 2147483647
    var v2 string = ""
    var v3 uint16 = 755
    var v4 string = ((v2 + v2) + "foo bar")
    if (!(15 != 40)) {
        v3 -= (v3 | 0)
        switch v2 {
        case "":
            var v15 int = (^(v1))
            _ = v15
        case "a":
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 bool = false
                _ = fi16
                _ = v17
            }
        }
        var v18 int8 = -126
        _ = v18
    }
    v19 := GenS1{F0: genFn0(), F1: -1}
    if (-1 <= 0) {
        v20 := GenS2{F0: int64(v1), F1: v3, F2: -2147483648}
        var v21 int64 = 893
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: 211, F1: 9223372036854775807}
    v4 = v2
    v23 := []int{v1, -375}
    var v24 string = "\n\t"
    v2 = v4
    fmt.Println(v2, v4, v2)
    v4 = ((((("a" + (v2 + v4)) + v24) + (((("" + ((v2 + v4) + (((((v24 + "a") + ((v4 + (v2 + v24)) + "a")) + ((v4 + v24) + v4)) + "") + v2))) + (((v2 + (("foo bar" + (v2 + "foo bar")) + "foo bar")) + "") + (("foo bar" + (v2 + "a")) + v4))) + "") + ("hello" + "foo bar"))) + v2) + "foo bar")
    v3 += v3
    v25 := GenS2{F0: 0, F1: v3, F2: (v1 / ((v1) | 1))}
    switch 128 {
    case -820:
        v26 := GenS1{F0: -2147483648, F1: 629}
        _ = v26
    case 516:
        for fi27 := 0; fi27 < 5; fi27++ {
            if len(v23) > 0 {
                v23[0] = 9223372036854775807
            }
            _ = fi27
        }
    case -1:
        var v28 int8 = 21
        _ = v28
    }
    v29 := GenS3{F0: 32767}
    v24 = v24
    fmt.Println(v2, v1)
    var v30 uint = 255
    var v31 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
