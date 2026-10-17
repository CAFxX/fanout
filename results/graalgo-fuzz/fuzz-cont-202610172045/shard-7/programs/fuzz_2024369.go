package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int8, p1 int, p2 uint) uint {
    p2 += 680
    var v5 int8 = p0
    var v6 string = ""
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            switch fi8 {
            case 455:
                v9 := GenS1{F0: 42, F1: p1, F2: -999}
                _ = v9
            }
            var v10 uint16 = 0
            p2 = p2
            _ = fi8
            _ = v10
        }
        _ = fi7
    }
    v11 := GenS2{F0: -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v5
    _ = v6
    return 255
}

func main() {
    var v1 int64 = -1
    var v2 int8 = 127
    var v3 int64 = (-(v1))
    var v4 uint32 = 2147483648
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS2{F0: 2147483647}
        _ = fi12
        _ = v13
    }
    v14 := []int{}
    switch (("a" + "") + ("\n\t" + "foo bar")) {
    case "\n\t":
        fmt.Println(v2)
    default:
        var v15 string = ("hello" + "\n\t")
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = fi16
    }
    _ = len(v14)
    var v17 uint = 500
    var v18 uint64 = 454
    if false {
        for fi19 := 0; fi19 < 17; fi19++ {
            switch fi19 {
            case 32768:
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 string = "hello"
                    _ = len(v14)
                    _ = fi20
                    _ = v21
                }
            case -9223372036854775808:
                v4 += 342
            case 32767:
                for fi22 := 0; fi22 < 18; fi22++ {
                    var v23 bool = ((true && (32767 == 1000)) && (0.1 <= (1e10 - 1.5)))
                    _ = fi22
                    _ = v23
                }
            default:
                for fi24 := 0; fi24 < 11; fi24++ {
                    v1 -= v3
                    var v25 string = "foo bar"
                    _ = fi24
                    _ = v25
                }
            }
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = fi19
        }
    }
    v26 := GenS2{F0: 1}
    fmt.Println(v2, v2, v17)
    v3 += v1
    var v27 int32 = 0
    var v28 int16 = -765
    v3 -= -1
    var v29 string = "hello"
    if (!true) {
        fmt.Println(v4, v3, v29)
        var v30 string = ""
        _ = v30
    } else {
        if ((127 <= uint8(v17)) && true) {
            for fi31 := 0; fi31 < 18; fi31++ {
                v32 := GenS2{F0: v27}
                _ = fi31
                _ = v32
            }
            for fi33 := 0; fi33 < 13; fi33++ {
                fmt.Println(v4)
                v34 := GenS2{F0: (v27 << ((uint(v4)) & 63))}
                fmt.Println(v27, v1)
                _ = fi33
                _ = v34
            }
        }
        v35 := GenS2{F0: v27}
        _ = v35
    }
    v36 := GenS2{F0: (-(v27))}
    v2 = 73
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
}
