package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int16 }
type GenS2 struct { F0 int; F1 int32 }

func main() {
    var v1 string = "\n\t"
    var v2 uint32 = 1
    var v3 float32 = (3.14159 - 0.1)
    v4 := []int{4294967295}
    var v5 int32 = 1
    var v6 uint16 = 1
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 float64 = 0.1
            _ = fi8
            _ = v9
        }
        var v10 int8 = 100
        for fi11 := 0; fi11 < 17; fi11++ {
            fmt.Println(fi11, v2)
            switch (fi11 + -1) {
            case -9223372036854775808:
                v6 = v6
            default:
                var v12 uint64 = 1
                _ = v12
            }
            _ = fi11
        }
        _ = fi7
        _ = v10
    }
    var v13 bool = (!false)
    for fi14 := 0; fi14 < 4; fi14++ {
        if true {
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                v2 = (v2 * v2)
                v1 = (v1 + v1)
            }
        }
        _ = fi14
    }
    if (32767 > int16(v2)) {
        if false {
            var v15 int = 867
            var v16 uint16 = 65535
            switch v1 {
            case "foo bar":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                fmt.Println(v15, v6, v5)
            }
            _ = v15
            _ = v16
        } else {
            var v17 uint = 2
            _ = v17
        }
        var v18 bool = true
        _ = v18
    }
    var v19 bool = (v13 && (-775 <= -450))
    v13 = true
    switch 42 {
    case -9223372036854775808:
        var v20 uint32 = (^(v2))
        _ = v20
    case -481:
        switch -245 {
        case 127:
            v2 -= 65535
        case -80:
            var v21 float64 = (((-2.25 / (1e10 + 0.1)) + (float64(v3) + 3.14159)) / ((0.1 * float64((v3 + 1.5))) - 1e10))
            _ = v21
        case 858:
            if ("hello" != "a") {
                v22 := GenS1{F0: 32767, F1: v1, F2: 32767}
                if v19 {
                    fmt.Println(v6)
                    v13 = (1 <= 32767)
                    v19 = true
                } else {
                    var v23 bool = v19
                    fmt.Println(v6, v13, v19)
                    var v24 int = 32767
                    _ = v23
                    _ = v24
                }
                for fi25 := 0; fi25 < 19; fi25++ {
                    for fi26 := 0; fi26 < 11; fi26++ {
                        _ = len(v4)
                        var v27 int64 = 789
                        v28 := GenS2{F0: (^(fi25)), F1: v5}
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    v13 = (true || (255 < 216))
                    _ = fi25
                }
                _ = v22
            }
        }
    }
    var v29 uint32 = v2
    v30 := GenS2{F0: 9223372036854775807, F1: v5}
    fmt.Println(v19)
    for fi31 := 0; fi31 < 3; fi31++ {
        for fi32 := 0; fi32 < 4; fi32++ {
            var v33 float64 = (0.0 * 1e10)
            v6 = v6
            if (v29 >= 2147483647) {
                if len(v4) > 0 {
                    v4[0] = (fi31 % ((fi32) | 1))
                }
            }
            _ = fi32
            _ = v33
        }
        var v34 uint16 = (v6 / ((v6) | 1))
        if len(v4) > 0 {
            v4[0] = 1
        }
        _ = fi31
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
