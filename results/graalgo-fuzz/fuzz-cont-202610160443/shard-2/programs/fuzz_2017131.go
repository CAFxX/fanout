package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 int8, p1 int16) uint16 {
    var v7 uint64 = 177
    v8 := GenS1{F0: (p1 | (p1 / ((461) | 1)))}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint32) string {
    var v9 int16 = 21
    if false {
        for fi10 := 0; fi10 < 15; fi10++ {
            fmt.Println(p0, p0)
            _ = fi10
        }
        if true {
            var v11 uint = 471
            v9 = (v9 + (^(v9)))
            if false {
                v9 = v9
                for fi12 := 0; fi12 < 6; fi12++ {
                    v13 := GenS2{F0: "a", F1: 479}
                    _ = fi12
                    _ = v13
                }
                p0 = p0
            }
            _ = v11
        }
        var v14 int = -9223372036854775808
        _ = v14
    } else {
        v15 := []int{768, 879, 1}
        switch -372 {
        case 653:
            if len(v15) > 0 {
                v15[0] = 32767
            }
        }
        _ = v15
    }
    if (false || (false || (623 < 9223372036854775807))) {
        v16 := []int{}
        for fi17 := 0; fi17 < 3; fi17++ {
            if true {
                var v18 float32 = ((0.0 / 3.14159) + 1.5)
                v9 += (v9 << ((42) & 63))
                _ = v18
            }
            v9 = v9
            _ = fi17
        }
        v16 = append(v16, -9223372036854775808)
        _ = v16
    } else {
        if ((-2147483648 >= 2147483647) || (15 <= 2)) {
            var v19 int64 = -9223372036854775808
            var v20 bool = true
            _ = v19
            _ = v20
        } else {
            v21 := []int{9223372036854775807, -85, 263, -1}
            _ = v21
        }
        var v22 string = (("\n\t" + ("foo bar" + "")) + "foo bar")
        v22 = v22
        _ = v22
    }
    var v23 bool = false
    _ = p0
    _ = v23
    _ = v9
    return "foo bar"
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 634
    var v3 float64 = (float64((1e10 - ((3.14159 * 1e10) + -2.25))) + -2.25)
    v4 := make(map[uint32]int)
    var v5 uint16 = 255
    var v6 float32 = ((0.1 / (((3.14159 + (3.14159 * float32(0.1))) / 1e10) - -2.25)) * (0.0 - 0.1))
    var v24 int8 = -55
    for fi25 := 0; fi25 < 10; fi25++ {
        v6 -= (0.0 - v6)
        var v26 uint32 = v2
        v5 = v5
        _ = fi25
        _ = v26
    }
    var v27 int64 = -782
    v24 = 58
    if (0 > 18446744073709551615) {
        _ = len(v1)
        var v28 uint = 691
        fmt.Println(v28)
        _ = v28
    } else {
        v27 -= 42
    }
    delete(v4, v2)
    var v29 uint8 = 128
    var v30 uint64 = 1
    var v31 uint16 = (v5 ^ v5)
    v32 := GenS2{F0: ("foo bar" + "foo bar"), F1: int64(v30)}
    fmt.Println(v31, v31)
    for fi33 := 0; fi33 < 6; fi33++ {
        var v34 uint = 42
        v3 = ((v3 + 0.0) / v3)
        _ = fi33
        _ = v34
    }
    switch 32767 {
    case 32767:
        var v35 string = ("foo bar" + ("foo bar" + ("hello" + "foo bar")))
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
