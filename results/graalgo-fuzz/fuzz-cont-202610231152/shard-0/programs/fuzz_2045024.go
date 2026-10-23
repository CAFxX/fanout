package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint8, p1 uint8, p2 int64) uint32 {
    p1 = p1
    for fi4 := 0; fi4 < 12; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            p2 = p2
            v6 := GenS1{F0: (1 == uint16(p1))}
            v7 := []int{(fi4 * fi4), fi5}
            _ = fi5
            _ = v6
            _ = v7
        }
        for fi8 := 0; fi8 < 15; fi8++ {
            if false {
                if false {
                    v9 := []int{330, fi4, (fi4 * 32767)}
                    _ = v9
                }
                if (!(1 <= 2147483647)) {
                    v10 := GenS1{F0: ((!(!(100 < 14))) && false)}
                    p1 += p0
                    fmt.Println(fi8, p0, fi8)
                    _ = v10
                } else {
                    var v11 float64 = (0.0 * 3.14159)
                    v12 := GenS1{F0: false}
                    p0 -= 128
                    _ = v11
                    _ = v12
                }
            }
            if true {
                v13 := GenS1{F0: false}
                _ = v13
            }
            _ = fi8
        }
        _ = fi4
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint16 = 390
    var v2 int16 = 1
    var v3 float64 = 1e10
    v1 -= (v1 * (^(v1)))
    v14 := []int{}
    v2 += (v2 & v2)
    switch -604 {
    case -9223372036854775808:
        for fi15 := 0; fi15 < 13; fi15++ {
            fmt.Println(v1)
            v2 = 0
            var v16 int8 = 58
            _ = fi15
            _ = v16
        }
    }
    v17 := GenS1{F0: false}
    var v18 int64 = -9223372036854775808
    if false {
        v3 += float64(1.5)
    } else {
        switch (("foo bar" + ("hello" + "foo bar")) + ("" + "")) {
        case "hello":
            v18 += -235
        }
        v2 -= v2
        fmt.Println(v18, v2)
    }
    v14 = append(v14, -281)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
