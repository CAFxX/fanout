package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
func genFn0(p0 uint16, p1 uint8, p2 bool) uint32 {
    p2 = p2
    v6 := GenS1{F0: uint64(p0), F1: 133}
    v7 := GenS1{F0: 0, F1: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return uint32(p0)
}
func genFn1(p0 uint32) uint {
    v8 := GenS1{F0: 42, F1: (^(p0))}
    if false {
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        switch (((("\n\t" + "hello") + "a") + (("" + "foo bar") + "foo bar")) + "hello") {
        case "a":
            if (1 >= 255) {
                for fi11 := 0; fi11 < 18; fi11++ {
                    var v12 uint32 = p0
                    _ = fi11
                    _ = v12
                }
            } else {
                if (1 >= 65535) {
                    var v13 string = "a"
                    var v14 int = 127
                    if false {
                        for fi15 := 0; fi15 < 7; fi15++ {
                            v16 := GenS1{F0: 42, F1: p0}
                            _ = fi15
                            _ = v16
                        }
                    }
                    _ = v13
                    _ = v14
                }
                v17 := GenS1{F0: 18446744073709551615, F1: genFn0(2, 50, true)}
                _ = v17
            }
        default:
            if false {
                switch -9223372036854775808 {
                case -273:
                    for fi18 := 0; fi18 < 4; fi18++ {
                        if (-1 <= 42) {
                            p0 = p0
                        }
                        switch ("" + ((("" + "") + "") + "\n\t")) {
                        case "":
                            v19 := []int{}
                            _ = v19
                        }
                        for fi20 := 0; fi20 < 9; fi20++ {
                            switch (fi20 & (fi20 & fi18)) {
                            case -234:
                                v21 := GenS1{F0: 42, F1: 42}
                                _ = v21
                            }
                            var v22 bool = ((false && (4294967295 >= 1)) && (genFn0(42, 0, (1 >= 517)) > genFn0(836, 174, false)))
                            v22 = (false || (1000 >= 32767))
                            v23 := GenS1{F0: 746, F1: genFn0(42, 30, v22)}
                            _ = fi20
                            _ = v22
                            _ = v23
                        }
                        _ = fi18
                    }
                case 989:
                    for fi24 := 0; fi24 < 12; fi24++ {
                        v25 := GenS1{F0: 1, F1: 2147483648}
                        _ = fi24
                        _ = v25
                    }
                default:
                    fmt.Println(p0)
                }
                p0 = p0
            }
        }
    } else {
        switch ((("\n\t" + "\n\t") + ("" + "hello")) + "hello") {
        case "a":
            var v26 string = ("\n\t" + ("a" + ("a" + ("hello" + "\n\t"))))
            _ = v26
        case "hello":
            p0 -= (^(p0))
        case "":
            fmt.Println(p0)
        default:
            p0 = (^(p0))
        }
        for fi27 := 0; fi27 < 15; fi27++ {
            for fi28 := 0; fi28 < 16; fi28++ {
                v29 := []int{}
                p0 += p0
                _ = fi28
                _ = v29
            }
            p0 = (p0 << ((1) & 63))
            _ = fi27
        }
    }
    var v30 int16 = 0
    _ = p0
    _ = v30
    _ = v8
    return 18446744073709551615
}

func main() {
    v1 := []uint32{760, 2147483648, 65535, 2}
    var v2 int8 = 27
    var v3 int16 = -662
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    var v31 int = -626
    v3 = v3
    v3 -= v3
    if (1e10 == 3.14159) {
        var v32 uint = 40
        fmt.Println(v31)
        fmt.Println(v3, v32, v2)
        _ = v32
    } else {
        var v33 float64 = 0.0
        _ = v33
    }
    v31 -= (v31 << ((65535) & 63))
    var v34 uint8 = 42
    _ = len(v1)
    v2 = -128
    v35 := v4[(-(v31))]
    var v36 int16 = 1000
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
