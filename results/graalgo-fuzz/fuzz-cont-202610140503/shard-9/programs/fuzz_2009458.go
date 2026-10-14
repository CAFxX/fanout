package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int) int32 {
    var v7 int8 = -128
    var v8 uint32 = 0
    if (!false) {
        var v9 float32 = (3.14159 / float32(1.5))
        var v10 float32 = 1.5
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            p0 = (p0 >> ((732) & 63))
            var v13 int64 = -9223372036854775808
            v14 := []int{p0, -1, (fi12 >> ((0) & 63)), fi12}
            _ = fi12
            _ = v13
            _ = v14
        }
        fmt.Println(p0)
        v15 := GenS1{F0: 1}
        _ = fi11
        _ = v15
    }
    p0 -= -2147483648
    _ = p0
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 int64 = -665
    var v2 int16 = 1
    var v3 float32 = -2.25
    v4 := make(map[string]string)
    v5 := make(map[uint32]int)
    v6 := []uint32{42}
    v16 := GenS1{F0: 533}
    if (102 > -1) {
        if true {
            if false {
                switch ("foo bar" + (("\n\t" + "") + "foo bar")) {
                case "a":
                    v2 += v2
                case "hello":
                    v4["hello"] = "hello"
                }
                v6 = append(v6, 2147483647)
            } else {
                var v17 uint8 = 42
                v6 = append(v6, 41)
                var v18 float64 = 0.0
                _ = v17
                _ = v18
            }
            var v19 uint8 = 2
            v1 = -140
            _ = v19
        }
        if (!("" == "")) {
            for fi20 := 0; fi20 < 19; fi20++ {
                fmt.Println(fi20, v2, v3)
                _ = fi20
            }
            if len(v6) > 0 {
                v6[0] = 2
            }
        } else {
            _ = len(v4)
        }
        var v21 bool = true
        _ = v21
    }
    switch "" {
    case "hello":
        var v22 uint = 4294967295
        _ = v22
    case "a":
        var v23 uint32 = 1
        _ = v23
    case "\n\t":
        var v24 float64 = float64(v3)
        _ = v24
    default:
        v1 = (v1 + v1)
    }
    if false {
        var v25 int64 = v1
        var v26 uint16 = 144
        _ = v25
        _ = v26
    }
    var v27 float64 = -2.25
    _ = len(v6)
    delete(v4, ("" + ("a" + ("" + ((("foo bar" + "foo bar") + ((("\n\t" + "hello") + "a") + ((("" + "") + "a") + "\n\t"))) + ("foo bar" + "\n\t"))))))
    var v28 int64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
