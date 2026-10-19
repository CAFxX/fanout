package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 string }
func genFn0(p0 int8, p1 int8, p2 int8) uint16 {
    p1 = p1
    var v5 bool = false
    p1 -= 42
    v6 := []int{1, 32768}
    p0 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 51
}

func main() {
    v1 := []string{"foo bar", ((("" + (("" + "foo bar") + ("" + ("foo bar" + ("foo bar" + "hello"))))) + (("a" + "") + "\n\t")) + "")}
    v2 := make(map[string]int)
    var v3 int32 = -2147483648
    var v4 uint16 = 0
    var v7 uint8 = 128
    var v8 int8 = 1
    for fi9 := 0; fi9 < 11; fi9++ {
        switch "hello" {
        case "":
            switch (-(fi9)) {
            case 32767:
                for fi10 := 0; fi10 < 19; fi10++ {
                    v3 = v3
                    _ = fi10
                }
            case 149:
                for fi11 := 0; fi11 < 12; fi11++ {
                    v3 += v3
                    v4 += (v4 * v4)
                    if (((481 == 490) && (!false)) || (v7 > v7)) {
                        if (!(((true && (fi9 >= fi9)) && (!false)) && (!false))) {
                            v8 = 100
                            v12 := GenS1{F0: 49, F1: "\n\t"}
                            v13 := v2[(("" + "hello") + "")]
                            _ = v12
                            _ = v13
                        }
                        switch "hello" {
                        case "hello":
                            v8 += -128
                        default:
                            v2["hello"] = fi9
                        }
                        fmt.Println(v3, v7)
                    }
                    _ = fi11
                }
            case 42:
                var v14 int = (fi9 << ((666) & 63))
                _ = v14
            }
        case "hello":
            fmt.Println(v4, v8, v7)
        case "a":
            v7 = v7
        }
        v15 := GenS2{F0: uint16(v7)}
        _ = fi9
        _ = v15
    }
    if true {
        v7 = v7
        v3 += (v3 ^ v3)
        var v16 int32 = v3
        _ = v16
    } else {
        v17 := GenS1{F0: (v8 + v8), F1: "foo bar"}
        _ = v17
    }
    var v18 bool = false
    _ = len(v2)
    v7 += 1
    v19 := GenS3{F0: v8, F1: "\n\t"}
    v18 = v18
    fmt.Println(v8, v18)
    var v20 int16 = 1000
    var v21 int = -634
    switch v21 {
    case 4294967295:
        for fi22 := 0; fi22 < 8; fi22++ {
            v7 += v7
            v20 = (v20 & v20)
            var v23 int32 = 42
            _ = fi22
            _ = v23
        }
    case -2147483648:
        fmt.Println(v20)
    case 266:
        v3 -= (v3 * v3)
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        fmt.Println(fi24)
        _ = fi24
    }
    var v25 int = v21
    var v26 uint8 = v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
