package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint64 }
func genFn0(p0 int16, p1 uint16) uint16 {
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := []int{127, (-(fi5)), -1, (fi5 | -1)}
        _ = fi5
        _ = v6
    }
    fmt.Println(p0)
    var v7 bool = false
    _ = p0
    _ = p1
    _ = v7
    return p1
}

func main() {
    var v1 string = "a"
    var v2 int16 = 42
    var v3 int64 = 0
    v4 := make(map[int]string)
    fmt.Println(v2, v2, v3)
    if true {
        for fi8 := 0; fi8 < 8; fi8++ {
            switch fi8 {
            case 127:
                var v9 bool = false
                _ = v9
            case 403:
                v10 := v4[(fi8 >> ((42) & 63))]
                _ = v10
            default:
                v11 := GenS1{F0: -128}
                _ = v11
            }
            _ = fi8
        }
        v12 := []int{}
        var v13 uint8 = 42
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 12; fi14++ {
            if false {
                fmt.Println(fi14, fi14, fi14)
            }
            _ = fi14
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 uint16 = 65535
                _ = fi16
                _ = v17
            }
            if (!(67 <= 190)) {
                v2 = -218
                fmt.Println(v3, fi15)
                v1 = "a"
            } else {
                switch (fi15 % (((^(fi15))) | 1)) {
                case 2147483647:
                    switch (v1 + v1) {
                    case "\n\t":
                        v2 -= v2
                    }
                case -9223372036854775808:
                    v2 += (v2 + v2)
                }
            }
            _ = fi15
        }
        if (!false) {
            delete(v4, -825)
        }
    }
    if true {
        for fi18 := 0; fi18 < 14; fi18++ {
            var v19 uint8 = 189
            _ = fi18
            _ = v19
        }
        var v20 int64 = (v3 | v3)
        _ = v20
    }
    var v21 uint8 = 184
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 int8 = 127
        if false {
            v3 = 483
            if true {
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 int64 = (v3 / ((763) | 1))
                    var v26 float64 = float64(0.0)
                    fmt.Println(v2)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            } else {
                v27 := GenS1{F0: v23}
                _ = v27
            }
        }
        var v28 string = ("hello" + v1)
        _ = fi22
        _ = v23
        _ = v28
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        if false {
            _ = len(v4)
        }
        var v30 int = 32768
        _ = fi29
        _ = v30
    }
    v4[128] = "hello"
    v31 := GenS2{F0: (v21 ^ v21), F1: 783, F2: 464}
    v2 += (v2 % (((-(v2))) | 1))
    v32 := GenS1{F0: 81}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
