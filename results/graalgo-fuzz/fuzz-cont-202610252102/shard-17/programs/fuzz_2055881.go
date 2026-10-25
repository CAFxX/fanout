package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint }
func genFn0(p0 bool) uint32 {
    switch (("" + "\n\t") + ("foo bar" + (("a" + "\n\t") + ("" + "hello")))) {
    case "":
        if true {
            var v6 int16 = -1
            if true {
                if false {
                    v6 -= (^(v6))
                    for fi7 := 0; fi7 < 14; fi7++ {
                        p0 = (p0 || p0)
                        _ = fi7
                    }
                } else {
                    var v8 uint8 = 2
                    v6 += (v6 / (((v6 * (-(v6)))) | 1))
                    v9 := GenS1{F0: (!p0)}
                    _ = v8
                    _ = v9
                }
            } else {
                var v10 uint32 = 2
                switch ("\n\t" + (("" + "\n\t") + "")) {
                case "hello":
                    v11 := GenS2{F0: "hello", F1: 306}
                    _ = v11
                case "foo bar":
                    if (22 >= -1) {
                        v12 := []int{127}
                        fmt.Println(p0, p0)
                        if p0 {
                            if (-2.25 == -2.25) {
                                fmt.Println(v10, v10, v10)
                            }
                        } else {
                            var v13 uint32 = 65535
                            _ = len(v12)
                            _ = v13
                        }
                        _ = v12
                    }
                case "\n\t":
                    var v14 string = ""
                    _ = v14
                default:
                    for fi15 := 0; fi15 < 8; fi15++ {
                        var v16 int8 = -93
                        p0 = (9223372036854775807 == -651)
                        p0 = (!true)
                        _ = fi15
                        _ = v16
                    }
                }
                _ = v10
            }
            _ = v6
        }
    }
    if (-29 >= 1) {
        for fi17 := 0; fi17 < 7; fi17++ {
            switch (fi17 / ((fi17) | 1)) {
            case -820:
                switch "hello" {
                case "":
                    if false {
                        fmt.Println(fi17, fi17)
                    }
                }
            }
            _ = fi17
        }
        p0 = (255 == 510)
    } else {
        v18 := []int{1}
        var v19 int32 = -1
        _ = v18
        _ = v19
    }
    _ = p0
    return 60
}
func genFn1() uint16 {
    v20 := GenS1{F0: (!((false || (0 < 1000)) && false))}
    v21 := GenS1{F0: true}
    _ = v20
    _ = v21
    return 1
}

func main() {
    var v1 string = ""
    v2 := make(map[string]bool)
    var v3 uint8 = 185
    v4 := make(map[int]bool)
    v5 := []int{9223372036854775807, 529, 621, int(v3)}
    if false {
        v1 = v1
    }
    v1 = ""
    if true {
        _ = len(v4)
        if (1 == 0) {
            switch 4294967295 {
            case 127:
                v1 = (v1 + ("hello" + ((v1 + "\n\t") + ("foo bar" + (v1 + v1)))))
            }
            var v22 uint64 = 4294967295
            v22 += 1
            _ = v22
        }
    }
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v3, v1)
    var v23 uint64 = 0
    fmt.Println(v23, v1, v1)
    fmt.Println(v23, v3)
    v23 = v23
    v23 = (v23 >> ((4294967295) & 63))
    v3 = (v3 >> ((818) & 63))
    v23 = (v23 & v23)
    if (0.1 == (-2.25 + 0.0)) {
        v23 += 2
        var v24 uint16 = 32767
        v24 = (v24 << ((920) & 63))
        _ = v24
    }
    fmt.Println(v1, v1)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
