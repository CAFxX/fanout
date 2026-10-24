package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 bool }
type GenS2 struct { F0 uint; F1 bool; F2 string }
func genFn0(p0 bool, p1 int) uint {
    switch 32767 {
    case -9223372036854775808:
        v6 := GenS1{F0: 0, F1: -128, F2: ((("foo bar" + "foo bar") + "\n\t") != "hello")}
        _ = v6
    default:
        p1 = p1
    }
    p1 -= p1
    v7 := []int{}
    if (!p0) {
        p1 = (^(p1))
    } else {
        switch "a" {
        case "":
            if p0 {
                var v8 int = p1
                v8 -= p1
                _ = v8
            }
        case "hello":
            var v9 int16 = -1
            _ = v9
        default:
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 string = ""
                v12 := GenS1{F0: 0, F1: -105, F2: (32767 <= 75)}
                var v13 int64 = 9223372036854775807
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
        }
    }
    p0 = (((-2.25 + ((3.14159 - 0.1) - -2.25)) == 1.5) && false)
    _ = p0
    _ = p1
    _ = v7
    return 255
}

func main() {
    var v1 int32 = -280
    var v2 int32 = 0
    v3 := make(map[uint32]bool)
    var v4 int64 = -881
    v5 := make(map[uint32]uint32)
    fmt.Println(v4, v4, v4)
    switch "" {
    case "foo bar":
        var v14 uint64 = 42
        _ = v14
    default:
        v15 := GenS2{F0: 255, F1: true, F2: (((("" + (("\n\t" + ("a" + "")) + ("foo bar" + ("foo bar" + "a")))) + "\n\t") + (("\n\t" + ("a" + "\n\t")) + ("hello" + "hello"))) + ("hello" + ""))}
        _ = v15
    }
    if ((!true) || false) {
        v1 = 933
        v1 = v2
    }
    var v16 uint32 = 65535
    var v17 int16 = -1
    switch 4294967295 {
    case 32768:
        for fi18 := 0; fi18 < 17; fi18++ {
            fmt.Println(fi18, v4, v2)
            _ = fi18
        }
    case 1:
        v4 = int64(v16)
    default:
        v2 += (v1 << ((1) & 63))
    }
    v3[v16] = false
    _ = len(v5)
    v17 = (v17 | 32767)
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := []int{fi19}
        _ = fi19
        _ = v20
    }
    v21 := []int{1, 2147483647}
    fmt.Println(v16, v2)
    var v22 uint64 = 697
    if (!true) {
        delete(v3, 327)
        v16 = (v16 << ((4294967295) & 63))
    } else {
        if false {
            v23 := GenS2{F0: 255, F1: ((v16 << ((uint(v22)) & 63)) == 2), F2: ""}
            v4 = v4
            _ = v23
        } else {
            v24 := GenS2{F0: 415, F1: (!true), F2: "\n\t"}
            _ = v24
        }
        v17 = v17
    }
    _ = len(v21)
    var v25 string = ""
    var v26 bool = (!((-(v17)) != v17))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
