package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint64) uint8 {
    if (0 != (p0 & p0)) {
        p0 += p0
    }
    for fi5 := 0; fi5 < 2; fi5++ {
        switch "" {
        case "":
            p0 = p0
        }
        if false {
            var v6 uint8 = 68
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := []int{(-(fi7))}
                v9 := GenS1{F0: 4, F1: 0}
                v8 = append(v8, fi7)
                p0 -= (p0 * p0)
                _ = fi7
                _ = v8
                _ = v9
            }
            var v10 uint64 = p0
            _ = v10
            _ = v6
        }
        _ = fi5
    }
    switch -248 {
    case 32767:
        var v11 uint16 = 255
        _ = v11
    case 32768:
        v12 := []int{-9223372036854775808, -1, 127, -1}
        _ = v12
    case 128:
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 uint16 = 537
            p0 = (p0 + (p0 ^ p0))
            _ = fi13
            _ = v14
        }
    }
    v15 := []int{32767, 0, 42}
    _ = p0
    _ = v15
    return 2
}
func genFn1(p0 int16, p1 uint32, p2 uint64) uint32 {
    fmt.Println(p0)
    v16 := GenS1{F0: 0, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return p1
}

func main() {
    var v1 uint8 = 2
    var v2 string = ""
    v3 := []uint32{824, uint32(v1), 42}
    var v4 uint = 0
    v4 += v4
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 int8 = 115
        var v19 string = v2
        _ = fi17
        _ = v18
        _ = v19
    }
    if len(v3) > 0 {
        v3[0] = 231
    }
    v2 = "hello"
    var v20 uint32 = 4294967295
    v21 := GenS1{F0: 919, F1: 42}
    var v22 int16 = 32767
    v23 := GenS2{F0: -706}
    var v24 int8 = 100
    fmt.Println(v22, v22)
    for fi25 := 0; fi25 < 8; fi25++ {
        v20 = 2147483647
        if false {
            v26 := GenS1{F0: 65535, F1: -994}
            _ = v26
        }
        _ = fi25
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        for fi28 := 0; fi28 < 14; fi28++ {
            if (!((0 >= -190) || (-9223372036854775808 < -1))) {
                _ = len(v3)
                v4 -= 4294967295
            }
            var v29 int32 = -2147483648
            _ = fi28
            _ = v29
        }
        v22 = 476
        _ = fi27
    }
    fmt.Println(v1, v4)
    var v30 int = -1
    fmt.Println(v20)
    var v31 bool = ((((v24 >> ((v4) & 63)) > v24) || (244 < 2)) || false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
