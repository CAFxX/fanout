package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint8, p2 uint32) uint32 {
    for fi7 := 0; fi7 < 10; fi7++ {
        switch (("a" + "") + ("\n\t" + "foo bar")) {
        case "\n\t":
            var v8 int16 = 0
            _ = v8
        default:
            p2 = (p2 - p2)
        }
        v9 := GenS1{F0: -428, F1: (!(!(!true))), F2: -2147483648}
        var v10 uint = 2
        _ = fi7
        _ = v10
        _ = v9
    }
    var v11 int8 = -128
    fmt.Println(v11, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 4294967295
}

func main() {
    var v1 int32 = 0
    v2 := []string{"", "", "\n\t", ((("hello" + ("foo bar" + "a")) + ("a" + "hello")) + "foo bar")}
    var v3 int32 = 1
    var v4 int64 = -183
    var v5 int = 9223372036854775807
    var v6 int32 = (v1 >> ((65535) & 63))
    v12 := GenS1{F0: (^(v1)), F1: true, F2: (v5 / ((4294967295) | 1))}
    v13 := GenS1{F0: v6, F1: ((v1 ^ -2147483648) == v6), F2: 2147483647}
    _ = len(v2)
    _ = len(v2)
    var v14 int32 = (v1 & 1)
    if true {
        for fi15 := 0; fi15 < 5; fi15++ {
            v14 += (v14 % (((^(v6))) | 1))
            _ = fi15
        }
        var v16 uint64 = 699
        _ = v16
    } else {
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 float64 = (0.1 + -2.25)
            fmt.Println(v6, v6)
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            _ = fi17
            _ = v18
        }
        v5 -= v5
        var v19 int64 = v4
        _ = v19
    }
    v5 = v5
    switch "" {
    case "\n\t":
        for fi20 := 0; fi20 < 20; fi20++ {
            fmt.Println(fi20, v14)
            switch (^(v5)) {
            case 9223372036854775807:
                for fi21 := 0; fi21 < 11; fi21++ {
                    if false {
                        if false {
                            if false {
                                var v22 uint = 525
                                v2 = append(v2, ((("a" + ("hello" + ("hello" + "hello"))) + "foo bar") + "foo bar"))
                                var v23 float64 = 1.5
                                _ = v22
                                _ = v23
                            }
                        }
                        fmt.Println(v14)
                        v1 += (v6 >> ((0) & 63))
                    } else {
                        v1 = (^(v6))
                        fmt.Println(fi21, fi21, fi21)
                    }
                    v24 := GenS2{F0: int32(fi21)}
                    _ = fi21
                    _ = v24
                }
            case 813:
                v4 += 9223372036854775807
            }
            if (-11 > -113) {
                fmt.Println(v3)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi20
        }
    case "a":
        fmt.Println(v3, v3)
    case "":
        fmt.Println(v5, v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
