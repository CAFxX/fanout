package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int64 }
func genFn0(p0 uint8, p1 uint32, p2 uint) string {
    for fi4 := 0; fi4 < 2; fi4++ {
        if true {
            p1 += (p1 ^ 0)
            v5 := GenS1{F0: p1}
            p2 += (p2 << (((p2 + p2)) & 63))
            _ = v5
        } else {
            var v6 bool = (!true)
            _ = v6
        }
        _ = fi4
    }
    if true {
        v7 := GenS2{F0: (p2 / (((^(p2))) | 1)), F1: 0}
        switch 4294967295 {
        case 820:
            var v8 int64 = 1
            _ = v8
        case -1:
            v9 := GenS2{F0: 65535, F1: -1}
            _ = v9
        }
        v10 := GenS1{F0: (^(p1))}
        _ = v10
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return (("a" + "\n\t") + "hello")
}
func genFn1(p0 bool, p1 int64, p2 uint8) int8 {
    switch ("hello" + "foo bar") {
    case "a":
        p0 = p0
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := []int{fi11}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 59
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []string{}
    var v3 bool = true
    if len(v2) > 0 {
        v2[0] = "a"
    }
    var v13 int8 = -128
    for fi14 := 0; fi14 < 11; fi14++ {
        if v3 {
            v13 -= (v13 >> ((259) & 63))
        }
        _ = fi14
    }
    v15 := GenS2{F0: 1, F1: -512}
    var v16 uint64 = v1
    fmt.Println(v13, v1, v3)
    var v17 uint64 = (v16 >> ((0) & 63))
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int32 = -1
        _ = fi18
        _ = v19
    }
    fmt.Println(v13)
    v3 = true
    v1 += (v17 / (((v16 ^ v16)) | 1))
    var v20 int64 = -1
    if (!("hello" <= "a")) {
        switch ("hello" + ("\n\t" + "a")) {
        case "":
            fmt.Println(v1, v17, v1)
        case "foo bar":
            v21 := GenS1{F0: 849}
            _ = v21
        case "hello":
            v16 -= 0
        }
    } else {
        var v22 uint = 899
        var v23 float64 = (1e10 + 0.0)
        var v24 int64 = (-(v20))
        _ = v22
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: 1}
    v17 = 737
    v2 = append(v2, "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v25
    _ = v3
}
