package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int; F1 string; F2 uint8 }
func genFn0(p0 uint8, p1 int64, p2 int64) int64 {
    fmt.Println(p1, p0, p2)
    p0 = (p0 - (p0 % ((147) | 1)))
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := []int{}
        var v6 uint16 = 123
        if len(v5) > 0 {
            v5[0] = 862
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(p0, p2)
    var v7 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    v1 := []string{("foo bar" + "foo bar"), "", "hello", "foo bar"}
    var v2 int8 = -117
    v3 := []string{"foo bar", (("" + "\n\t") + ("hello" + "")), ("\n\t" + "\n\t")}
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: 0, F1: -117}
        v2 -= (v2 & v2)
        fmt.Println(v2)
        _ = fi8
        _ = v9
    }
    v2 = v2
    v3 = append(v3, ("\n\t" + "\n\t"))
    var v10 bool = true
    for fi11 := 0; fi11 < 11; fi11++ {
        if true {
            fmt.Println(v10, v10)
            v12 := GenS2{F0: (v2 & -128), F1: -1}
            v2 -= 127
            _ = v12
        } else {
            v2 += (v2 << ((651) & 63))
            var v13 int32 = -49
            _ = v13
        }
        v10 = false
        _ = fi11
    }
    var v14 uint64 = 939
    v10 = (65535 < uint16(v14))
    var v15 uint16 = 65535
    v10 = true
    v16 := GenS3{F0: 9223372036854775807, F1: "hello", F2: 255}
    switch ("" + "a") {
    case "\n\t":
        v1 = append(v1, "")
    case "a":
        v17 := GenS1{F0: -1, F1: 1}
        _ = v17
    }
    switch 1 {
    case 128:
        v10 = (42 != 9223372036854775807)
    case 42:
        if v10 {
            v10 = (!true)
            v14 -= v14
        }
    }
    var v18 uint8 = 127
    v14 += (v14 + (v14 | 2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
}
