package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32, p1 uint) uint {
    var v7 bool = false
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint16 = 933
            _ = fi9
            _ = v10
        }
        fmt.Println(fi8, p0)
        v11 := GenS1{F0: p1, F1: "a"}
        _ = fi8
        _ = v11
    }
    fmt.Println(p0, p1, p1)
    _ = p0
    _ = p1
    _ = v7
    return (^(p1))
}
func genFn1(p0 uint64) uint64 {
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := []int{fi12, (fi12 % ((127) | 1))}
        if false {
            var v14 string = ((("foo bar" + "") + "\n\t") + "")
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint64 = (p0 & p0)
        _ = fi15
        _ = v16
    }
    _ = p0
    return (p0 + 222)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    v3 := make(map[uint32]bool)
    var v4 uint16 = 988
    var v5 float32 = (float32((-2.25 / (1e10 + (3.14159 + (3.14159 + 3.14159))))) - -2.25)
    v6 := []string{"hello", ("" + "\n\t"), "foo bar", "\n\t"}
    if false {
        if len(v6) > 0 {
            v6[0] = ""
        }
        delete(v1, 2147483647)
    }
    if (!true) {
        v1[4294967295] = 2147483647
        v5 = v5
        switch (("\n\t" + "hello") + "") {
        case "foo bar":
            v5 = v5
        default:
            fmt.Println(v4, v5)
        }
    }
    var v17 float32 = 3.14159
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS1{F0: 55, F1: ""}
        for fi20 := 0; fi20 < 1; fi20++ {
            if len(v6) > 0 {
                v6[0] = "foo bar"
            }
            _ = fi20
        }
        fmt.Println(fi18, fi18)
        _ = fi18
        _ = v19
    }
    if true {
        if false {
            if false {
                fmt.Println(v5, v4, v5)
                v21 := v1[557]
                fmt.Println(v21)
                _ = v21
            } else {
                var v22 string = (("" + "a") + "hello")
                v23 := v3[2147483648]
                v24 := GenS2{F0: v4}
                _ = v22
                _ = v23
                _ = v24
            }
        } else {
            v5 = (0.1 + 0.0)
        }
        v17 -= 1e10
    }
    _ = len(v2)
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 float32 = v17
        var v27 bool = true
        _ = fi25
        _ = v26
        _ = v27
    }
    if true {
        v2 = append(v2, "")
        fmt.Println(v17)
        var v28 uint = 0
        _ = v28
    } else {
        var v29 int64 = int64(v4)
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        for fi31 := 0; fi31 < 2; fi31++ {
            v32 := v3[2147483648]
            _ = fi31
            _ = v32
        }
        fmt.Println(v17)
        _ = fi30
    }
    v33 := GenS1{F0: 4294967295, F1: ("a" + ("foo bar" + "a"))}
    var v34 uint = 148
    var v35 int16 = -1
    v34 = genFn0(0, v34)
    for fi36 := 0; fi36 < 15; fi36++ {
        if false {
            var v37 uint64 = 2
            _ = v37
        }
        _ = fi36
    }
    v1[2] = 571
    v38 := GenS2{F0: v4}
    fmt.Println(v5, v4, v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
