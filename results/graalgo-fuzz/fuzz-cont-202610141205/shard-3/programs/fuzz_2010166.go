package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint32, p1 uint8, p2 int64) int16 {
    fmt.Println(p2, p2, p0)
    if true {
        var v5 string = ("hello" + ("a" + (("" + "\n\t") + ("a" + ""))))
        if false {
            fmt.Println(v5, v5)
            var v6 bool = true
            _ = v6
        }
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := []int{2147483647, -9223372036854775808}
    var v2 int16 = 1000
    v3 := []int{-341, -958, 1, 9223372036854775807}
    var v4 uint64 = 302
    if false {
        v3 = append(v3, 1)
    } else {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 bool = (2147483648 == 2147483648)
            switch (("a" + "") + (("hello" + "a") + "foo bar")) {
            case "a":
                v2 = -302
            default:
                _ = len(v1)
            }
            switch "foo bar" {
            case "foo bar":
                v8 = ((-(v2)) > (v2 ^ (-(v2))))
            }
            _ = fi7
            _ = v8
        }
        switch 193 {
        case -194:
            v2 = 42
        case 42:
            v4 = (v4 >> ((65535) & 63))
        default:
            v3 = append(v3, 128)
        }
        var v9 float32 = ((1.5 - (3.14159 + 0.0)) - (3.14159 * float32(0.1)))
        _ = v9
    }
    if false {
        var v10 float64 = (0.1 * float64((float32(0.0) - ((0.0 + (-2.25 / 1e10)) * 0.0))))
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            if false {
                fmt.Println(v2)
                v1 = append(v1, (fi12 | (^(fi12))))
                fmt.Println(v4, v4)
            }
            v3 = append(v3, fi12)
            _ = fi12
        }
        var v13 uint64 = (v4 * v4)
        var v14 float64 = 1e10
        _ = fi11
        _ = v13
        _ = v14
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 = (v4 & 18446744073709551615)
    var v15 uint16 = 943
    v2 = 0
    fmt.Println(v2, v2)
    if (255 > 511) {
        for fi16 := 0; fi16 < 19; fi16++ {
            if len(v3) > 0 {
                v3[0] = fi16
            }
            _ = fi16
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            fmt.Println(v4)
            _ = fi17
        }
    } else {
        if true {
            v3 = append(v3, -357)
            v15 = (v15 % ((v15) | 1))
        } else {
            switch ("a" + "hello") {
            case "\n\t":
                v2 -= (v2 | genFn0(673, 2, 9223372036854775807))
            default:
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 int32 = 2147483647
                    _ = fi18
                    _ = v19
                }
            }
            fmt.Println(v4, v15, v15)
        }
        if false {
            v20 := GenS1{F0: 127, F1: genFn0(65535, 175, int64(v15)), F2: -9223372036854775808}
            if len(v1) > 0 {
                v1[0] = 9223372036854775807
            }
            _ = v20
        }
        for fi21 := 0; fi21 < 12; fi21++ {
            v2 += v2
            v15 -= 835
            v4 += v4
            _ = fi21
        }
    }
    v22 := GenS2{F0: 2147483647}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
