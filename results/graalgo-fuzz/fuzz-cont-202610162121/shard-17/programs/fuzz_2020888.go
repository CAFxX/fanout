package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int; F1 uint64; F2 string }

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[uint32]int)
    v3 := []string{}
    var v4 int32 = 42
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS1{F0: 42, F1: 0}
        v2[2147483648] = (fi5 ^ fi5)
        v4 += (v4 >> ((0) & 63))
        _ = fi5
        _ = v6
    }
    v7 := v2[65535]
    switch (^(v7)) {
    case -62:
        if ((v7 != 2147483647) || true) {
            v7 -= v7
        } else {
            for fi8 := 0; fi8 < 19; fi8++ {
                if (2 < 1) {
                    if (-1 == (v4 - (-(v4)))) {
                        v4 = (-(v4))
                        fmt.Println(v7, fi8, v7)
                    } else {
                        switch fi8 {
                        case -375:
                            v7 -= fi8
                        }
                        var v9 uint64 = 655
                        fmt.Println(v4, v4, v9)
                        _ = v9
                    }
                    v10 := GenS1{F0: -128, F1: 42}
                    _ = v10
                }
                _ = fi8
            }
        }
    case 0:
        var v11 int16 = 322
        _ = v11
    default:
        var v12 uint16 = 1
        _ = v12
    }
    if true {
        v13 := GenS2{F0: v7, F1: 1, F2: ((("hello" + "foo bar") + "a") + "hello")}
        v14 := GenS1{F0: 42, F1: 0}
        v7 -= v7
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        v4 -= -758
        var v16 int16 = 1000
        v4 -= v4
        _ = fi15
        _ = v16
    }
    var v17 uint8 = 128
    v17 = (v17 % ((v17) | 1))
    for fi18 := 0; fi18 < 2; fi18++ {
        v19 := v2[2]
        _ = fi18
        _ = v19
    }
    fmt.Println(v17)
    for fi20 := 0; fi20 < 5; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 int16 = -873
            var v23 string = ("hello" + ("\n\t" + "hello"))
            var v24 uint16 = 979
            v25 := v2[1]
            fmt.Println(v22, v17)
            v7 = (-(fi20))
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
            _ = v25
        }
        if false {
            v26 := GenS1{F0: -128, F1: v17}
            v4 += 938
            _ = v26
        } else {
            if (!false) {
                for fi27 := 0; fi27 < 5; fi27++ {
                    v28 := GenS2{F0: -396, F1: 42, F2: "\n\t"}
                    var v29 int64 = 42
                    var v30 float64 = 0.1
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                if (v17 >= (^(v17))) {
                    for fi31 := 0; fi31 < 9; fi31++ {
                        var v32 string = ""
                        _ = fi31
                        _ = v32
                    }
                }
                var v33 int32 = -752
                _ = v33
            }
        }
        v7 = (fi20 | (fi20 % ((4294967295) | 1)))
        _ = fi20
    }
    var v34 uint = 2
    v17 = 36
    v7 -= v7
    var v35 int8 = int8(v7)
    var v36 uint16 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v7
}
