package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
func genFn0(p0 uint8) string {
    p0 = (p0 >> ((768) & 63))
    p0 = 117
    p0 += p0
    if ((65535 < 317) && false) {
        v6 := GenS1{F0: "\n\t", F1: 42}
        if (true || ((101 > 127) && (true && false))) {
            p0 += (p0 >> ((2) & 63))
            fmt.Println(p0)
        }
        _ = v6
    } else {
        p0 -= p0
        if true {
            var v7 int8 = -15
            var v8 string = "hello"
            v8 = ("hello" + v8)
            _ = v7
            _ = v8
        }
        v9 := []int{127, 9223372036854775807}
        _ = v9
    }
    v10 := []int{682, -9223372036854775808}
    _ = p0
    _ = v10
    return "foo bar"
}
func genFn1(p0 int, p1 int16) int8 {
    for fi11 := 0; fi11 < 19; fi11++ {
        p0 -= p0
        _ = fi11
    }
    v12 := []int{}
    _ = p0
    _ = p1
    _ = v12
    return -126
}

func main() {
    var v1 int8 = -1
    var v2 uint8 = 1
    var v3 bool = false
    var v4 uint64 = 2
    var v5 bool = (!false)
    var v13 bool = (!(2 > 18446744073709551615))
    v14 := GenS1{F0: "\n\t", F1: 1}
    v13 = v13
    v15 := []int{138, -1, -662}
    _ = len(v15)
    var v16 uint8 = v2
    if true {
        if len(v15) > 0 {
            v15[0] = -2147483648
        }
    } else {
        v5 = v3
        _ = len(v15)
    }
    if (1 == -1) {
        fmt.Println(v16, v16)
    } else {
        var v17 int8 = v1
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = v17
    }
    var v18 uint32 = uint32(v2)
    var v19 int8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
