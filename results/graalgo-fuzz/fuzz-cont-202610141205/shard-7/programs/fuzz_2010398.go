package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    v3 := []int{-2147483648, 9223372036854775807, 32768, -647}
    _ = len(v1)
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 int32 = -2147483648
        v5 -= v5
        _ = fi4
        _ = v5
    }
    var v6 uint = 4294967295
    v6 = (v6 % ((1) | 1))
    switch 42 {
    case 2147483647:
        if false {
            if (!(1 != 0)) {
                if true {
                    var v7 int16 = 175
                    for fi8 := 0; fi8 < 6; fi8++ {
                        var v9 string = ""
                        _ = fi8
                        _ = v9
                    }
                    _ = v7
                }
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := GenS1{F0: 0}
                    _ = fi10
                    _ = v11
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                if (uint32(v6) <= 147) {
                    fmt.Println(v6)
                    v12 := GenS1{F0: 42}
                    var v13 string = ("a" + "foo bar")
                    _ = v12
                    _ = v13
                }
                for fi14 := 0; fi14 < 6; fi14++ {
                    for fi15 := 0; fi15 < 6; fi15++ {
                        _ = len(v3)
                        var v16 uint16 = 0
                        _ = fi15
                        _ = v16
                    }
                    v17 := GenS1{F0: 1}
                    var v18 int64 = -589
                    _ = fi14
                    _ = v17
                    _ = v18
                }
            }
            var v19 bool = (2147483647 > -99)
            _ = v19
        }
    case -445:
        v20 := GenS1{F0: -40}
        _ = v20
    case -943:
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 uint64 = 1
            fmt.Println(v22, fi21)
            _ = fi21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        v1[2147483648] = fi23
        fmt.Println(fi23)
        if true {
            delete(v1, 1)
        }
        _ = fi23
    }
    v6 = 1
    if (!(1 > 0)) {
        v24 := GenS1{F0: 0}
        _ = v24
    } else {
        var v25 uint8 = uint8(v6)
        v26 := GenS1{F0: 0}
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
}
