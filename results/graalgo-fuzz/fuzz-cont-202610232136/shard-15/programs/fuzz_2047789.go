package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint) bool {
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint16 = 932
        v6 = v6
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: 623, F1: 1}
            var v9 int16 = -645
            if false {
                v6 = (v6 + v6)
                v9 = (v9 + v9)
            } else {
                v6 = 939
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    v10 := []int{32767}
    switch (("\n\t" + "hello") + "hello") {
    case "foo bar":
        p0 -= (p0 & (p0 - 1))
    case "\n\t":
        v11 := GenS2{F0: false}
        _ = v11
    default:
        var v12 float64 = 0.0
        _ = v12
    }
    if true {
        if false {
            fmt.Println(p0)
        }
    }
    _ = p0
    _ = v10
    return (!(true && (!false)))
}
func genFn1(p0 uint) uint64 {
    v13 := []int{-9223372036854775808, 1}
    var v14 uint32 = 2147483648
    v14 = v14
    switch (("foo bar" + "a") + ((("" + ("a" + "\n\t")) + "hello") + ("hello" + ("foo bar" + "\n\t")))) {
    case "":
        var v15 int32 = -2147483648
        _ = v15
    case "foo bar":
        v14 -= v14
    case "a":
        var v16 int = 4294967295
        _ = v16
    }
    _ = p0
    _ = v13
    _ = v14
    return 18446744073709551615
}

func main() {
    var v1 string = ""
    var v2 int8 = 127
    var v3 float32 = 0.0
    var v4 uint16 = 837
    v2 += -37
    v2 += (v2 & v2)
    v2 = (v2 << ((42) & 63))
    v17 := GenS1{F0: 1, F1: -215}
    switch 0 {
    case 1:
        v3 += 0.0
    }
    v2 = (^(v2))
    switch -9223372036854775808 {
    case 328:
        fmt.Println(v4, v3)
    case 824:
        var v18 uint8 = 127
        _ = v18
    case -1:
        v2 = (v2 ^ (v2 & 42))
    }
    v2 = (v2 << ((545) & 63))
    var v19 int8 = v2
    v20 := []int{556}
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 1; fi22++ {
            fmt.Println(fi21, fi22)
            _ = fi22
        }
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
