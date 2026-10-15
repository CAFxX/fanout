package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint }

func main() {
    var v1 uint8 = 127
    var v2 uint16 = 0
    var v3 float64 = float64((0.0 * 1.5))
    var v4 bool = (!true)
    v5 := GenS1{F0: v1, F1: -423, F2: 231}
    var v6 bool = ((2 < 449) || (!("a" >= (("\n\t" + (("" + "a") + "foo bar")) + (("" + "hello") + "hello")))))
    if (39 < 279) {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{}
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: v1, F1: 1, F2: 0}
        _ = v9
    }
    if v4 {
        var v10 uint16 = (v2 >> ((0) & 63))
        var v11 uint64 = 243
        _ = v10
        _ = v11
    } else {
        var v12 uint16 = v2
        _ = v12
    }
    v13 := []int{-1, -473}
    var v14 bool = v6
    if false {
        for fi15 := 0; fi15 < 14; fi15++ {
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 uint64 = 42
                if v14 {
                    fmt.Println(v3, v14, v2)
                    v18 := GenS2{F0: (v2 + (v2 * (v2 >> ((893) & 63)))), F1: int32(v1), F2: 65535}
                    fmt.Println(fi16, fi16)
                    _ = v18
                }
                fmt.Println(fi15, fi15, fi16)
                v19 := GenS1{F0: v1, F1: 271, F2: (v1 & 1)}
                var v20 float64 = (v3 * float64(((0.0 - 3.14159) * ((float32((v3 * v3)) - (1e10 / 1.0)) / -2.25))))
                _ = fi16
                _ = v17
                _ = v19
                _ = v20
            }
            _ = fi15
        }
    } else {
        switch "a" {
        case "hello":
            var v21 int64 = 361
            _ = v21
        case "foo bar":
            var v22 uint64 = 756
            _ = v22
        case "":
            var v23 int32 = -2147483648
            _ = v23
        }
        v13 = append(v13, 1)
        v13 = append(v13, 4294967295)
    }
    v24 := GenS1{F0: 116, F1: 1000, F2: v1}
    v6 = (-1 <= 0)
    switch 889 {
    case -959:
        var v25 int64 = -1
        _ = v25
    case 4294967295:
        if v6 {
            var v26 string = ""
            v3 = (v3 / 1e10)
            _ = v26
        }
    case 1:
        v2 += v2
    default:
        fmt.Println(v6)
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        v14 = true
        _ = fi27
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        v2 -= v2
        for fi29 := 0; fi29 < 10; fi29++ {
            switch (fi28 << ((568) & 63)) {
            case 50:
                var v30 uint32 = 2147483647
                _ = v30
            }
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = fi29
        }
        _ = fi28
    }
    if false {
        for fi31 := 0; fi31 < 4; fi31++ {
            var v32 float64 = v3
            _ = fi31
            _ = v32
        }
    } else {
        var v33 float32 = 1e10
        _ = v33
    }
    v2 += (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
