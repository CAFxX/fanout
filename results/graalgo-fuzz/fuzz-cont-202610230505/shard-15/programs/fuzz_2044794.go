package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint }

func main() {
    var v1 bool = (2 == 4294967295)
    v2 := []string{("foo bar" + ("a" + "foo bar"))}
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    v5 := []int{}
    var v6 uint = 255
    fmt.Println(v6, v6)
    v6 -= (v6 % ((0) | 1))
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(v6, v1)
        for fi8 := 0; fi8 < 7; fi8++ {
            if (1 >= -79) {
                if len(v5) > 0 {
                    v5[0] = fi8
                }
                fmt.Println(fi7)
            }
            var v9 int64 = 1
            _ = fi8
            _ = v9
        }
        switch ("" + "foo bar") {
        case "hello":
            var v10 bool = false
            _ = v10
        case "":
            v6 += (v6 * v6)
        case "\n\t":
            if ("foo bar" != ("a" + "\n\t")) {
                v11 := GenS1{F0: 2, F1: int32(v6), F2: v6}
                _ = v11
            }
        }
        _ = fi7
    }
    v6 -= v6
    switch -911 {
    case 0:
        v12 := GenS1{F0: 2, F1: -150, F2: 35}
        _ = v12
    case -690:
        v1 = (464 >= 4294967295)
    default:
        if v1 {
            if (0 != 961) {
                fmt.Println(v6)
            }
            switch 554 {
            case 4294967295:
                v6 = v6
            }
        } else {
            var v13 bool = (0 < 861)
            _ = v13
        }
    }
    var v14 bool = v1
    v14 = v1
    var v15 int32 = -37
    v16 := GenS1{F0: 18446744073709551615, F1: (v15 ^ (v15 + (v15 >> ((4294967295) & 63)))), F2: (v6 | (^(v6)))}
    fmt.Println(v1, v6, v6)
    if true {
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 1; fi18++ {
                switch 127 {
                case -600:
                    fmt.Println(fi17)
                case -394:
                    if (false && (18446744073709551615 <= v6)) {
                        v1 = false
                    } else {
                        v19 := GenS1{F0: 464, F1: (v15 + 0), F2: (v6 << ((2) & 63))}
                        _ = v19
                    }
                case -9223372036854775808:
                    var v20 uint = v6
                    fmt.Println(v20, v15)
                    if (3.14159 <= (1.5 + float32(1e10))) {
                        var v21 uint = (^(v6))
                        _ = v21
                    } else {
                        v15 -= v15
                    }
                    _ = v20
                }
                switch "\n\t" {
                case "\n\t":
                    v6 += v6
                }
                switch (("" + "hello") + "") {
                case "\n\t":
                    var v22 uint32 = 1
                    _ = v22
                default:
                    v6 -= (v6 | v6)
                    fmt.Println(v15, v1)
                }
                _ = fi18
            }
            _ = len(v3)
            _ = fi17
        }
        delete(v3, 127)
    } else {
        v1 = (42 != 198)
        var v23 uint64 = 695
        _ = v23
    }
    v15 = -63
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
