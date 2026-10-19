package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
func genFn0(p0 uint16, p1 int8, p2 int64) bool {
    v4 := []int{}
    if ((p2 << ((581) & 63)) != 15) {
        p1 += p1
        v5 := GenS1{F0: "\n\t", F1: 164}
        v4 = append(v4, 2147483647)
        _ = v5
    } else {
        p0 = p0
        var v6 uint = 720
        switch -2147483648 {
        case -1:
            for fi7 := 0; fi7 < 17; fi7++ {
                for fi8 := 0; fi8 < 16; fi8++ {
                    fmt.Println(p2)
                    v6 += v6
                    if true {
                        fmt.Println(p0)
                    } else {
                        var v9 uint16 = 0
                        _ = v9
                    }
                    _ = fi8
                }
                _ = fi7
            }
        case 128:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = v6
    }
    p2 = 398
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ((!((p1 >> ((18446744073709551615) & 63)) < 0)) && false)
}
func genFn1() string {
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS1{F0: ((("hello" + "\n\t") + "a") + ""), F1: 42}
        var v12 uint = 255
        _ = fi10
        _ = v11
        _ = v12
    }
    switch 4294967295 {
    case 2147483647:
        var v13 float64 = (0.0 + (float64(0.0) + ((3.14159 * 1.5) / 3.14159)))
        _ = v13
    case 9223372036854775807:
        v14 := []int{1}
        _ = v14
    case 32767:
        var v15 int32 = -493
        _ = v15
    }
    var v16 bool = (!true)
    fmt.Println(v16)
    _ = v16
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 411
    v3 := []int{}
    var v17 int8 = 127
    fmt.Println(v17, v17)
    for fi18 := 0; fi18 < 20; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            fmt.Println(fi18)
            v2 += 2
            switch "foo bar" {
            case "\n\t":
                var v20 uint32 = (v2 >> ((0) & 63))
                _ = v20
            case "hello":
                v21 := GenS1{F0: "a", F1: v2}
                _ = v21
            case "foo bar":
                v17 += -42
            }
            _ = fi19
        }
        v2 = v2
        _ = fi18
    }
    var v22 bool = false
    v23 := GenS1{F0: "a", F1: v2}
    var v24 uint16 = 612
    v25 := GenS1{F0: "\n\t", F1: 42}
    var v26 int8 = v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
