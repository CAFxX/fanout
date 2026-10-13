package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint64 }
func genFn0() int8 {
    var v7 uint = 65535
    for fi8 := 0; fi8 < 11; fi8++ {
        v7 = 42
        switch 2147483647 {
        case 9223372036854775807:
            var v9 int = fi8
            _ = v9
        case 127:
            v10 := []int{-1}
            _ = v10
        case 945:
            v7 = v7
        default:
            for fi11 := 0; fi11 < 20; fi11++ {
                fmt.Println(v7, v7)
                if ((false && (33 < fi11)) || (fi8 < -9223372036854775808)) {
                    fmt.Println(fi8, fi11, v7)
                } else {
                    v7 = (v7 << (((v7 | v7)) & 63))
                    v7 = 0
                    var v12 int = 823
                    _ = v12
                }
                _ = fi11
            }
        }
        _ = fi8
    }
    var v13 uint64 = 431
    _ = v13
    _ = v7
    return 1
}

func main() {
    var v1 uint8 = 255
    var v2 uint64 = 2
    var v3 int = 45
    var v4 int8 = 0
    var v5 int = v3
    v6 := []string{}
    if (false || true) {
        var v14 uint16 = 42
        _ = v14
    }
    v5 += v5
    if len(v6) > 0 {
        v6[0] = "\n\t"
    }
    if true {
        v5 -= v5
        v15 := GenS1{F0: 65535}
        switch "\n\t" {
        case "\n\t":
            v2 -= uint64(v1)
        default:
            for fi16 := 0; fi16 < 17; fi16++ {
                if (3.14159 <= 0.0) {
                    fmt.Println(v4, v2, v4)
                } else {
                    v3 = (v3 | v3)
                    if ((v5 * v3) <= (v5 / ((0) | 1))) {
                        fmt.Println(v4, v5, v5)
                    }
                }
                for fi17 := 0; fi17 < 19; fi17++ {
                    fmt.Println(v2, fi17)
                    _ = fi17
                }
                if true {
                    if len(v6) > 0 {
                        v6[0] = "foo bar"
                    }
                } else {
                    switch fi16 {
                    case -806:
                        if len(v6) > 0 {
                            v6[0] = "foo bar"
                        }
                    case -505:
                        for fi18 := 0; fi18 < 20; fi18++ {
                            v3 = v5
                            _ = fi18
                        }
                    default:
                        var v19 int8 = -1
                        _ = v19
                    }
                    switch "" {
                    case "a":
                        v20 := GenS1{F0: 773}
                        _ = v20
                    case "hello":
                        v1 += 255
                    }
                }
                _ = fi16
            }
        }
        _ = v15
    } else {
        v6 = append(v6, ("hello" + ("a" + "hello")))
    }
    v3 = v3
    var v21 int32 = -1
    fmt.Println(v4, v21)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v21, v2, v2)
    v3 += (v5 << ((268) & 63))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (!(false && (!false))) {
        v2 -= v2
        fmt.Println(v2)
    } else {
        fmt.Println(v21)
        for fi22 := 0; fi22 < 10; fi22++ {
            v23 := GenS2{F0: 4294967295, F1: (v2 >> ((1) & 63))}
            if false {
                for fi24 := 0; fi24 < 7; fi24++ {
                    var v25 uint16 = 0
                    switch (("hello" + "a") + "foo bar") {
                    case "hello":
                        var v26 int = -1
                        _ = v26
                    case "a":
                        var v27 int16 = -818
                        _ = v27
                    case "foo bar":
                        var v28 float32 = float32(-2.25)
                        _ = v28
                    }
                    _ = fi24
                    _ = v25
                }
                var v29 int32 = v21
                _ = v29
            }
            _ = fi22
            _ = v23
        }
        var v30 uint32 = 1
        _ = v30
    }
    v3 = v5
    var v31 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
