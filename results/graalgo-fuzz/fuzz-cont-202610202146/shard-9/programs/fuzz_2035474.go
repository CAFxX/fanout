package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
func genFn0(p0 int, p1 uint16) bool {
    if false {
        var v4 uint = 0
        _ = v4
    }
    switch ("\n\t" + "\n\t") {
    case "hello":
        fmt.Println(p0, p0)
    default:
        var v5 int64 = -716
        _ = v5
    }
    var v6 uint8 = 165
    for fi7 := 0; fi7 < 8; fi7++ {
        if false {
            for fi8 := 0; fi8 < 16; fi8++ {
                v9 := GenS1{F0: 42, F1: p1}
                v10 := []int{(p0 % (((-(fi7))) | 1)), fi8}
                _ = fi8
                _ = v10
                _ = v9
            }
            if (65535 < 726) {
                v6 -= 25
                v11 := GenS1{F0: 304, F1: 32767}
                _ = v11
            } else {
                var v12 int8 = -1
                _ = v12
            }
            if true {
                fmt.Println(fi7)
                v13 := []int{(^(p0)), (fi7 - fi7), (fi7 & fi7)}
                for fi14 := 0; fi14 < 8; fi14++ {
                    v15 := GenS1{F0: 264, F1: (^(p1))}
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
        } else {
            var v16 string = ""
            var v17 float32 = 3.14159
            _ = v16
            _ = v17
        }
        v18 := GenS1{F0: -595, F1: 0}
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := GenS1{F0: int16(p0), F1: p1}
            p0 -= (p0 >> ((14) & 63))
            _ = fi19
            _ = v20
        }
        _ = fi7
        _ = v18
    }
    var v21 int16 = -1
    _ = p0
    _ = p1
    _ = v21
    _ = v6
    return ((v21 - 0) > v21)
}

func main() {
    v1 := []string{}
    var v2 uint64 = 490
    v3 := []uint32{}
    var v22 int64 = -9223372036854775808
    if ((!((19 < 42) || (1000 < 32767))) && (!(!(("hello" + "") == ("hello" + ""))))) {
        var v23 int = 32767
        for fi24 := 0; fi24 < 17; fi24++ {
            var v25 uint16 = 1
            switch (^(fi24)) {
            case 312:
                v26 := GenS1{F0: 0, F1: v25}
                _ = v26
            }
            if (!(!(int32(v22) < -1))) {
                var v27 string = "foo bar"
                _ = v27
            }
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    v22 = (v22 & v22)
    if len(v1) > 0 {
        v1[0] = ("\n\t" + "")
    }
    fmt.Println(v22, v22)
    v2 = 476
    var v28 int8 = 1
    var v29 int32 = -1
    v30 := GenS1{F0: 197, F1: 682}
    var v31 int64 = 91
    v32 := GenS1{F0: 42, F1: 255}
    v3 = append(v3, 4294967295)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
