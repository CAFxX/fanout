package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 bool, p1 string, p2 int16) int {
    switch (p1 + "foo bar") {
    case "a":
        p2 = p2
    case "":
        for fi5 := 0; fi5 < 5; fi5++ {
            p1 = "foo bar"
            switch (-(fi5)) {
            case 32768:
                fmt.Println(fi5, p2)
            case 42:
                v6 := GenS1{F0: (636 < 746), F1: 334, F2: 9223372036854775807}
                _ = v6
            case -1:
                p2 -= 0
            default:
                v7 := GenS1{F0: p0, F1: 2147483647, F2: 75}
                _ = v7
            }
            v8 := []int{fi5, (fi5 | 127), (fi5 ^ 596)}
            _ = fi5
            _ = v8
        }
    case "foo bar":
        var v9 uint64 = 1
        _ = v9
    default:
        v10 := []int{}
        _ = v10
    }
    fmt.Println(p0)
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 float32 = 1e10
        if false {
            p1 = p1
            var v13 int64 = 102
            _ = v13
        } else {
            fmt.Println(p1, p1)
        }
        p0 = p0
        _ = fi11
        _ = v12
    }
    v14 := GenS1{F0: (625 <= 42), F1: 1, F2: 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 32767
}
func genFn1() bool {
    v15 := GenS2{F0: 1}
    v16 := GenS1{F0: ((false || true) && (!false)), F1: 65535, F2: 265}
    v17 := GenS1{F0: (!false), F1: 819, F2: 42}
    v18 := GenS1{F0: ((876 <= 42) && false), F1: 0, F2: 1}
    v19 := []int{513, 42}
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    return (42 != -32768)
}

func main() {
    var v1 float32 = 1e10
    var v2 float64 = 0.1
    var v3 int = -1
    var v4 int8 = 42
    var v20 float64 = (float64(v1) * v2)
    var v21 uint8 = 99
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 float32 = (((v1 + v1) * v1) - ((v1 / v1) + ((((v1 * v1) - ((((float32(3.14159) - (1e10 / v1)) / v1) - v1) / v1)) * v1) * v1)))
        _ = fi22
        _ = v23
    }
    v24 := []int{}
    if ((65535 > 4294967295) && (32767 < (-(v3)))) {
        v25 := GenS1{F0: (1 > 258), F1: 42, F2: 9223372036854775807}
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        v3 -= (^(fi26))
        _ = fi26
    }
    var v27 string = ((("" + "") + ("hello" + "foo bar")) + "")
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 int32 = 1
        v27 = ((v27 + (v27 + (("foo bar" + "hello") + ""))) + "hello")
        fmt.Println(v1, v4)
        _ = fi28
        _ = v29
    }
    var v30 uint = 0
    var v31 uint8 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
