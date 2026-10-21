package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int8 {
    v4 := []int{32767, -890, 487}
    v5 := GenS1{F0: 157}
    var v6 float64 = 3.14159
    var v7 string = "hello"
    var v8 int = 95
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint8) int {
    var v9 int = 637
    switch "\n\t" {
    case "a":
        if false {
            for fi10 := 0; fi10 < 9; fi10++ {
                switch "a" {
                case "hello":
                    if true {
                        v9 = v9
                        var v11 uint32 = 2147483647
                        _ = v11
                    }
                }
                _ = fi10
            }
        }
    case "\n\t":
        v12 := GenS1{F0: p0}
        _ = v12
    }
    var v13 int64 = -1
    _ = p0
    _ = v13
    _ = v9
    return 819
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{("" + ("a" + ("foo bar" + "\n\t"))), "", "hello", ""}
    var v3 uint16 = 2
    if (!true) {
        var v14 int8 = genFn0()
        _ = v14
    } else {
        if false {
            if (1 >= 2147483648) {
                if (255 > 2) {
                    var v15 uint8 = 248
                    _ = v15
                } else {
                    v3 += v3
                }
                var v16 uint = 340
                _ = v16
            } else {
                v3 += 32767
                fmt.Println(v3)
            }
            var v17 float64 = (3.14159 * float64(0.0))
            var v18 int64 = 0
            _ = v17
            _ = v18
        }
        var v19 uint32 = 0
        _ = v19
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v3)
    case "\n\t":
        v1["hello"] = 42
    case "foo bar":
        fmt.Println(v3)
    default:
        fmt.Println(v3)
    }
    v20 := GenS1{F0: 42}
    for fi21 := 0; fi21 < 6; fi21++ {
        v3 += v3
        v3 -= (v3 + 804)
        if false {
            v22 := GenS1{F0: 255}
            v3 = 42
            _ = v22
        }
        _ = fi21
    }
    switch "a" {
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "a":
        for fi23 := 0; fi23 < 2; fi23++ {
            for fi24 := 0; fi24 < 15; fi24++ {
                v25 := GenS1{F0: 1}
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
    case "":
        var v26 int16 = 42
        _ = v26
    }
    var v27 int8 = genFn0()
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        delete(v1, "a")
        _ = fi28
    }
    if false {
        fmt.Println(v27, v3)
        var v29 uint64 = 18446744073709551615
        _ = v29
    }
    var v30 int16 = -230
    v2 = append(v2, ("" + (((("hello" + ("" + ("" + "hello"))) + "foo bar") + "a") + "hello")))
    v31 := GenS1{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v30
    _ = v31
}
