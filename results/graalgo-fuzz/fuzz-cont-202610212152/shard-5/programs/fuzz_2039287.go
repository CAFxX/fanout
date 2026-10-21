package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 string }
type GenS2 struct { F0 int64; F1 bool }
func genFn0(p0 uint16) uint32 {
    if (-150 == 0) {
        p0 = 227
    }
    p0 += p0
    _ = p0
    return 2147483648
}

func main() {
    v1 := []int{-1, 0, 9223372036854775807, 0}
    var v2 uint8 = 1
    var v3 string = ("foo bar" + "a")
    var v4 int = 429
    if true {
        if false {
            for fi5 := 0; fi5 < 19; fi5++ {
                v4 -= v4
                v6 := GenS1{F0: 0, F1: 18446744073709551615, F2: v3}
                var v7 int8 = 127
                _ = fi5
                _ = v6
                _ = v7
            }
        }
        v2 -= v2
        var v8 float32 = ((0.0 - (3.14159 + 1.5)) / 1.5)
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: 2, F2: (v3 + v3)}
    if true {
        for fi10 := 0; fi10 < 14; fi10++ {
            v4 -= v4
            var v11 uint16 = 10
            _ = fi10
            _ = v11
        }
    } else {
        fmt.Println(v2, v4, v3)
    }
    if (974 > 9223372036854775807) {
        switch v4 {
        case -445:
            if (714 > 0) {
                v12 := GenS1{F0: 32767, F1: 478, F2: (("\n\t" + v3) + v3)}
                v3 = v3
                _ = v12
            }
        }
        switch ("" + ("\n\t" + "foo bar")) {
        case "\n\t":
            v3 = v3
        case "foo bar":
            var v13 int8 = -1
            _ = v13
        default:
            switch 127 {
            case 42:
                v4 -= (v4 - 4294967295)
            }
        }
        var v14 uint8 = (v2 % ((235) | 1))
        _ = v14
    }
    switch 580 {
    case 0:
        _ = len(v1)
    case 9223372036854775807:
        fmt.Println(v2)
    case 32768:
        if ((v2 >= (v2 & v2)) || false) {
            var v15 bool = false
            _ = v15
        } else {
            for fi16 := 0; fi16 < 20; fi16++ {
                v3 = (v3 + ("a" + ((v3 + (v3 + "foo bar")) + (v3 + v3))))
                v4 -= 0
                var v17 int = (^(fi16))
                _ = fi16
                _ = v17
            }
        }
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v19 := GenS1{F0: 395, F1: 276, F2: v3}
        _ = fi18
        _ = v19
    }
    var v20 string = (((((("" + (v3 + ((v3 + ((v3 + "a") + "a")) + (v3 + ("a" + ("" + "\n\t")))))) + ((v3 + "hello") + (v3 + ("a" + v3)))) + ("hello" + (v3 + "\n\t"))) + (v3 + v3)) + v3) + "foo bar")
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v21 := GenS2{F0: 9223372036854775807, F1: (!((v4 & v4) <= v4))}
    fmt.Println(v4)
    fmt.Println(v2, v20, v4)
    v2 -= v2
    v20 = "a"
    _ = len(v1)
    var v22 uint8 = 42
    v2 += (v22 - 2)
    var v23 int8 = -128
    for fi24 := 0; fi24 < 19; fi24++ {
        v20 = ""
        v22 = 2
        _ = fi24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
