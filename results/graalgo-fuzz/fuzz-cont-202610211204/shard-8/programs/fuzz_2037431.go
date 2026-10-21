package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint8 }
func genFn0(p0 string) uint8 {
    fmt.Println(p0)
    var v5 float64 = 0.0
    _ = p0
    _ = v5
    return 63
}
func genFn1() bool {
    if false {
        if ((0 <= 42) && true) {
            switch "a" {
            case "":
                if true {
                    if false {
                        v6 := GenS1{F0: 28, F1: 1000, F2: 85}
                        var v7 float64 = 1.5
                        _ = v6
                        _ = v7
                    }
                    v8 := []int{}
                    var v9 int16 = -1
                    _ = v8
                    _ = v9
                }
            }
        } else {
            var v10 float64 = (float64(0.0) * (3.14159 + (0.0 / 3.14159)))
            _ = v10
        }
        if false {
            v11 := GenS1{F0: 584, F1: 1, F2: 255}
            var v12 int8 = 0
            _ = v11
            _ = v12
        } else {
            var v13 int = 32767
            fmt.Println(v13)
            _ = v13
        }
        var v14 uint16 = 845
        _ = v14
    }
    v15 := []int{}
    _ = v15
    return (9223372036854775807 != 42)
}

func main() {
    v1 := []string{"hello", ("\n\t" + "hello")}
    var v2 bool = true
    var v3 int8 = 100
    v4 := []uint32{604}
    var v16 string = "hello"
    if (v2 || (true && v2)) {
        v2 = true
    } else {
        var v17 float32 = 3.14159
        var v18 int64 = -774
        fmt.Println(v3, v17)
        _ = v17
        _ = v18
    }
    fmt.Println(v2, v16, v3)
    switch v16 {
    case "foo bar":
        fmt.Println(v16, v2)
    case "\n\t":
        var v19 bool = v2
        _ = v19
    case "hello":
        _ = len(v4)
    }
    var v20 uint16 = 2
    for fi21 := 0; fi21 < 19; fi21++ {
        if false {
            v20 += v20
        } else {
            var v22 bool = (-2147483648 >= 128)
            var v23 uint8 = genFn0(v16)
            _ = v22
            _ = v23
        }
        v24 := GenS1{F0: v20, F1: 32767, F2: 21}
        _ = fi21
        _ = v24
    }
    v25 := GenS1{F0: (v20 << ((4294967295) & 63)), F1: 1000, F2: 16}
    var v26 uint32 = 1
    v1 = append(v1, "a")
    if false {
        var v27 float32 = (3.14159 / 1.0)
        v2 = true
        fmt.Println(v2, v3, v27)
        _ = v27
    }
    v1 = append(v1, v16)
    for fi28 := 0; fi28 < 8; fi28++ {
        fmt.Println(v20, v2, v20)
        _ = fi28
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        fmt.Println(v26, v26, v20)
        _ = len(v4)
        _ = fi29
    }
    var v30 float64 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
}
