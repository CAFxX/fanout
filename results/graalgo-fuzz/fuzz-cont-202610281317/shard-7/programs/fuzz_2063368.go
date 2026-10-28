package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
func genFn0(p0 int16, p1 bool, p2 uint16) uint32 {
    p2 += (^(p2))
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := []int{fi4, -980, 849, 127}
        var v6 float64 = (3.14159 - (float64(1e10) * -2.25))
        v7 := GenS1{F0: false, F1: 2, F2: 164}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483648
}

func main() {
    v1 := []uint32{4294967295, 42}
    v2 := make(map[int]string)
    var v3 uint16 = 0
    for fi8 := 0; fi8 < 18; fi8++ {
        v3 = 1
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := GenS1{F0: (!false), F1: 42, F2: -244}
            var v11 uint64 = 742
            v12 := GenS1{F0: (18446744073709551615 == 42), F1: uint(v11), F2: -606}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    v1 = append(v1, genFn0(42, (0 > uint8(v3)), 32767))
    for fi13 := 0; fi13 < 4; fi13++ {
        v3 += v3
        v3 = v3
        var v14 string = ""
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint = 4294967295
        fmt.Println(v3, v3)
        _ = fi15
        _ = v16
    }
    if false {
        v17 := GenS1{F0: false, F1: 937, F2: -2147483648}
        var v18 float32 = (1.5 * 3.14159)
        _ = v17
        _ = v18
    } else {
        v3 = (^(v3))
        var v19 float32 = (-2.25 / 1.5)
        _ = v19
    }
    delete(v2, 2147483647)
    switch "\n\t" {
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "foo bar":
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 int16 = -1
            delete(v2, fi20)
            _ = len(v2)
            _ = fi20
            _ = v21
        }
    default:
        v3 -= (v3 >> ((56) & 63))
    }
    _ = len(v1)
    _ = len(v1)
    switch 0 {
    case 32767:
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := GenS1{F0: true, F1: 0, F2: 566}
            fmt.Println(v3)
            var v24 float32 = 0.0
            _ = fi22
            _ = v23
            _ = v24
        }
    case 712:
        v25 := GenS1{F0: (42 <= 1), F1: 255, F2: 1}
        _ = v25
    default:
        v3 -= (v3 | v3)
    }
    var v26 string = ("hello" + "\n\t")
    v26 = (v26 + "a")
    for fi27 := 0; fi27 < 20; fi27++ {
        switch 958 {
        case 421:
            _ = len(v2)
        case 32768:
            var v28 uint64 = uint64(v3)
            _ = v28
        case -9223372036854775808:
            v29 := GenS1{F0: false, F1: 1, F2: -188}
            _ = v29
        }
        _ = fi27
    }
    var v30 uint = 18446744073709551615
    if true {
        var v31 uint16 = v3
        _ = v31
    } else {
        var v32 bool = false
        _ = v32
    }
    var v33 int = 32767
    var v34 int = 42
    for fi35 := 0; fi35 < 17; fi35++ {
        fmt.Println(v3, v34, v34)
        fmt.Println(fi35)
        _ = fi35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v33
    _ = v34
}
