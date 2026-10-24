package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 0
    var v3 float64 = float64((float32((1.5 + 0.1)) - 3.14159))
    for fi4 := 0; fi4 < 10; fi4++ {
        if (float32((-2.25 / v3)) <= 0.1) {
            var v5 bool = true
            if v5 {
                fmt.Println(fi4)
            }
            _ = v5
        }
        if false {
            v6 := []int{(fi4 >> ((1) & 63))}
            _ = v6
        } else {
            if (fi4 >= fi4) {
                fmt.Println(fi4)
                v1[""] = 2147483647
            }
        }
        _ = fi4
    }
    v7 := GenS1{F0: 482}
    v8 := []int{42, 4294967295, 531}
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v3, v3)
    default:
        v2 += (v2 * 128)
    }
    var v9 uint32 = 42
    var v10 float64 = v3
    var v11 int64 = 42
    delete(v1, "foo bar")
    fmt.Println(v9)
    fmt.Println(v2)
    v10 = float64(0.1)
    fmt.Println(v2)
    v2 = 182
    var v12 float64 = 1e10
    for fi13 := 0; fi13 < 13; fi13++ {
        switch "" {
        case "foo bar":
            if true {
                if true {
                    var v14 uint = 0
                    _ = v14
                }
                v15 := GenS1{F0: v9}
                _ = v15
            }
        }
        _ = fi13
    }
    var v16 uint8 = v2
    var v17 int16 = 32767
    switch "" {
    case "":
        if (true || (!true)) {
            v2 -= (v2 >> ((0) & 63))
            var v18 float32 = ((1e10 / 1e10) / 1e10)
            var v19 uint = 0
            _ = v18
            _ = v19
        }
    case "foo bar":
        v20 := GenS1{F0: uint32(v16)}
        _ = v20
    case "a":
        v1[((("\n\t" + "hello") + "a") + ((("a" + "\n\t") + "") + "foo bar"))] = (^(v9))
    default:
        if true {
            var v21 int = 551
            _ = v21
        } else {
            var v22 int64 = 0
            for fi23 := 0; fi23 < 8; fi23++ {
                v9 = v9
                _ = fi23
            }
            switch "\n\t" {
            case "a":
                for fi24 := 0; fi24 < 8; fi24++ {
                    for fi25 := 0; fi25 < 20; fi25++ {
                        var v26 float64 = ((1e10 * v3) - float64(0.1))
                        _ = fi25
                        _ = v26
                    }
                    delete(v1, "foo bar")
                    if (true && (v3 > v12)) {
                        var v27 uint64 = 761
                        fmt.Println(v17, v2)
                        _ = v27
                    } else {
                        _ = len(v8)
                    }
                    _ = fi24
                }
            }
            _ = v22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
