package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int64 }
func genFn0(p0 uint32) int32 {
    if true {
        var v4 int16 = 32767
        fmt.Println(p0, p0)
        _ = v4
    } else {
        p0 = p0
        if (!(9223372036854775807 > 397)) {
            p0 -= 65535
            p0 -= p0
        } else {
            p0 = p0
        }
    }
    var v5 uint16 = 1
    var v6 uint = 255
    v7 := []int{-699}
    for fi8 := 0; fi8 < 17; fi8++ {
        v5 = (v5 / ((475) | 1))
        v9 := GenS1{F0: 1, F1: 336, F2: 0}
        if true {
            var v10 string = ("foo bar" + ("foo bar" + "a"))
            _ = v10
        } else {
            var v11 int64 = -416
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 733
}
func genFn1(p0 uint16, p1 int32, p2 uint) string {
    fmt.Println(p0, p2)
    var v12 int = -528
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 float32 = 1.5
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return ("\n\t" + "\n\t")
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = (v1 << ((1) & 63))
    v3 := make(map[int]uint32)
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 int16 = 1
        switch "foo bar" {
        case "foo bar":
            var v17 int8 = 0
            _ = v17
        }
        v18 := GenS1{F0: 0, F1: (v2 / ((v1) | 1)), F2: 1}
        _ = fi15
        _ = v16
        _ = v18
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        v2 = v1
        v20 := GenS1{F0: 128, F1: v1, F2: 0}
        var v21 int = fi19
        _ = fi19
        _ = v20
        _ = v21
    }
    v2 = 375
    v1 -= (v1 % ((v2) | 1))
    for fi22 := 0; fi22 < 6; fi22++ {
        delete(v3, fi22)
        fmt.Println(v2, fi22)
        var v23 int32 = genFn0(457)
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        for fi25 := 0; fi25 < 16; fi25++ {
            v1 += 605
            _ = fi25
        }
        var v26 float64 = 1e10
        _ = fi24
        _ = v26
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        var v28 uint32 = 47
        v28 = 2
        _ = fi27
        _ = v28
    }
    switch 32768 {
    case 32767:
        fmt.Println(v2, v2)
    case 128:
        v2 -= v1
    case 9223372036854775807:
        for fi29 := 0; fi29 < 12; fi29++ {
            v1 = 527
            var v30 bool = (!false)
            switch fi29 {
            case 0:
                v30 = (v1 > 368)
            case 2147483647:
                v30 = ((fi29 >> ((1) & 63)) > fi29)
            case 127:
                for fi31 := 0; fi31 < 3; fi31++ {
                    v2 -= (v2 * 42)
                    _ = fi31
                }
            default:
                delete(v3, 9223372036854775807)
            }
            _ = fi29
            _ = v30
        }
    }
    v1 = v2
    if false {
        v32 := GenS1{F0: 2, F1: v1, F2: 9223372036854775807}
        var v33 int8 = 100
        var v34 int64 = 1
        _ = v32
        _ = v33
        _ = v34
    } else {
        _ = len(v3)
        switch "hello" {
        case "hello":
            if true {
                v2 = (v1 * 2)
            } else {
                if (691 <= uint16(v1)) {
                    var v35 uint = 513
                    _ = v35
                }
            }
        }
    }
    v36 := []int{330, -9223372036854775808, -164}
    var v37 int = 32768
    switch (("a" + ("a" + "")) + "a") {
    case "\n\t":
        fmt.Println(v2)
    case "a":
        v36 = append(v36, int(v1))
    default:
        var v38 int64 = 584
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}
