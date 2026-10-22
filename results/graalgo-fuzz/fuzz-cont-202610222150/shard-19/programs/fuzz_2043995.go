package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int64, p1 uint32) int16 {
    if (!true) {
        var v5 int = 0
        p1 = (^(p1))
        var v6 uint8 = 2
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: 839, F1: 0}
        fmt.Println(p0)
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 673, F1: 2}
    var v11 int32 = -916
    var v12 uint8 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 32767
}
func genFn1(p0 uint16, p1 int16) int16 {
    v13 := []int{32767, 127}
    fmt.Println(p1, p0)
    for fi14 := 0; fi14 < 14; fi14++ {
        switch fi14 {
        case -2147483648:
            v15 := GenS1{F0: 4294967295, F1: 605}
            _ = v15
        }
        p1 = (p1 & p1)
        _ = fi14
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v13
    return (p1 << ((974) & 63))
}

func main() {
    var v1 int8 = -128
    var v2 int64 = -1
    var v3 int16 = 32767
    v4 := []string{"hello", ("a" + (("" + "") + "")), "foo bar", "a"}
    var v16 int = 0
    v17 := GenS1{F0: 817, F1: 42}
    switch 618 {
    case -226:
        var v18 int64 = (v2 + -1)
        _ = v18
    case 127:
        fmt.Println(v16, v2, v2)
    default:
        var v19 int32 = 0
        _ = v19
    }
    var v20 float32 = 1.5
    var v21 int64 = (v2 >> ((4294967295) & 63))
    fmt.Println(v3, v2, v2)
    v4 = append(v4, "a")
    var v22 int16 = v3
    v4 = append(v4, "")
    switch "" {
    case "hello":
        var v23 float32 = v20
        _ = v23
    case "foo bar":
        v24 := GenS1{F0: 4294967295, F1: 18446744073709551615}
        _ = v24
    default:
        fmt.Println(v20)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
