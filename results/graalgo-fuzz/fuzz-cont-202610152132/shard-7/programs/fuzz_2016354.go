package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 bool = true
    var v2 int = 2147483647
    var v3 uint16 = 795
    var v4 uint = 0
    for fi5 := 0; fi5 < 14; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            v1 = false
            fmt.Println(v2, v4)
            _ = fi6
        }
        _ = fi5
    }
    fmt.Println(v2, v4, v4)
    fmt.Println(v3, v2, v2)
    v1 = (2147483647 != 1)
    for fi7 := 0; fi7 < 9; fi7++ {
        switch v2 {
        case 613:
            v8 := GenS1{F0: 59}
            _ = v8
        case 32768:
            v9 := []int{int(v4), (fi7 - (v2 / ((v2) | 1)))}
            _ = v9
        }
        v4 += 854
        _ = fi7
    }
    switch "hello" {
    case "hello":
        v4 = (^(v4))
    }
    if v1 {
        v3 = 2
        fmt.Println(v2)
        v2 = (v2 - (v2 + v2))
    }
    var v10 uint32 = 2
    if (true || (0.1 < (0.0 / 1e10))) {
        if v1 {
            for fi11 := 0; fi11 < 12; fi11++ {
                v4 = (v4 - v4)
                var v12 uint64 = 943
                v13 := []int{(fi11 - v2), (v2 - (-(fi11)))}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        var v14 int16 = -522
        var v15 int8 = 1
        _ = v14
        _ = v15
    }
    switch (-(v2)) {
    case 431:
        v10 -= uint32(v4)
    case 597:
        fmt.Println(v3)
    }
    v16 := []int{(-(v2)), v2, (v2 >> (((v4 + v4)) & 63))}
    v1 = v1
    if ((42 < v4) || v1) {
        var v17 int8 = -126
        v10 += uint32(v4)
        _ = v17
    }
    switch v2 {
    case 2147483647:
        for fi18 := 0; fi18 < 12; fi18++ {
            for fi19 := 0; fi19 < 20; fi19++ {
                v16 = append(v16, v2)
                if len(v16) > 0 {
                    v16[0] = -1
                }
                fmt.Println(fi18, v2)
                _ = fi19
            }
            v20 := GenS1{F0: 0}
            _ = fi18
            _ = v20
        }
    }
    v3 = v3
    if true {
        for fi21 := 0; fi21 < 5; fi21++ {
            v10 += uint32(v3)
            if true {
                v1 = (int8(v10) == 42)
            } else {
                var v22 string = "\n\t"
                var v23 uint16 = (v3 + 0)
                _ = v22
                _ = v23
            }
            fmt.Println(fi21)
            _ = fi21
        }
        v1 = v1
        var v24 uint16 = (v3 & v3)
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
