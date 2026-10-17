package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint64, p1 int8) uint64 {
    v6 := GenS1{F0: -389}
    for fi7 := 0; fi7 < 1; fi7++ {
        if ((((1.5 - -2.25) + 0.1) - (1e10 + 1.5)) == 3.14159) {
            fmt.Println(p1, p0, p0)
        }
        var v8 string = "a"
        if (false || false) {
            var v9 string = ("foo bar" + "hello")
            var v10 string = ("foo bar" + (v9 + ""))
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := []int{fi7, (fi11 - 153)}
                _ = fi11
                _ = v12
            }
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p1)
    for fi13 := 0; fi13 < 2; fi13++ {
        switch "\n\t" {
        case "hello":
            var v14 uint32 = 1
            _ = v14
        default:
            var v15 int8 = p1
            _ = v15
        }
        v16 := []int{}
        if (2 >= 4294967295) {
            p1 = p1
            switch ("foo bar" + "hello") {
            case "hello":
                for fi17 := 0; fi17 < 5; fi17++ {
                    p0 += p0
                    _ = fi17
                }
            }
            p1 -= -1
        } else {
            switch ("foo bar" + "hello") {
            case "hello":
                fmt.Println(fi13, fi13, p0)
            }
            fmt.Println(fi13, p0)
        }
        _ = fi13
        _ = v16
    }
    var v18 int16 = -128
    _ = p0
    _ = p1
    _ = v18
    _ = v6
    return 866
}
func genFn1(p0 int8) bool {
    p0 = (p0 % (((p0 | (-(p0)))) | 1))
    v19 := []int{256, 313, 2147483647}
    _ = p0
    _ = v19
    return (488 >= 65535)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = -116
    v3 := []int{2147483647}
    v4 := make(map[int]int)
    var v5 bool = true
    if false {
        v20 := GenS1{F0: 42}
        var v21 uint16 = 255
        _ = len(v4)
        _ = v20
        _ = v21
    } else {
        v2 = (v2 << ((65535) & 63))
        v22 := GenS1{F0: -2147483648}
        v23 := GenS1{F0: -933}
        _ = v22
        _ = v23
    }
    var v24 uint32 = 2
    fmt.Println(v24, v2, v5)
    if (!true) {
        var v25 float32 = (1e10 * 0.1)
        _ = v25
    }
    var v26 uint32 = (v24 - (v24 / ((v24) | 1)))
    for fi27 := 0; fi27 < 13; fi27++ {
        _ = len(v3)
        _ = fi27
    }
    var v28 string = ("foo bar" + ("foo bar" + "foo bar"))
    var v29 uint16 = 2
    v30 := GenS1{F0: 2147483647}
    fmt.Println(v28)
    v31 := GenS1{F0: -505}
    switch v28 {
    case "foo bar":
        v29 = v29
    case "\n\t":
        v26 = 648
    }
    v24 -= (v24 | 490)
    for fi32 := 0; fi32 < 8; fi32++ {
        fmt.Println(v29)
        _ = fi32
    }
    v3 = append(v3, -9223372036854775808)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
