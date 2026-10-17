package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 string; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
func genFn0() uint {
    v5 := []int{-62}
    switch "foo bar" {
    case "hello":
        var v6 int8 = -19
        _ = v6
    }
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = true
    var v3 float32 = (((1e10 + 0.1) + (1e10 - (3.14159 - 1e10))) * 0.1)
    v4 := []uint32{}
    v3 -= v3
    if (!v2) {
        switch "" {
        case "hello":
            var v7 uint16 = 255
            _ = v7
        case "a":
            var v8 float32 = (v3 / -2.25)
            _ = v8
        case "\n\t":
            if (v2 && ((v2 && (true && v2)) || (42 == 42))) {
                var v9 int32 = 42
                v9 -= (v9 >> ((620) & 63))
                _ = v9
            }
        }
        v3 = v3
        if (true && (v2 && ("" != "hello"))) {
            v2 = v2
            for fi10 := 0; fi10 < 2; fi10++ {
                v4 = append(v4, 959)
                _ = fi10
            }
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS1{F0: -1}
                fmt.Println(fi11)
                v4 = append(v4, 4294967295)
                _ = fi11
                _ = v12
            }
        }
    } else {
        v1["a"] = true
    }
    fmt.Println(v3)
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS2{F0: "hello", F1: "hello", F2: 42}
        _ = fi13
        _ = v14
    }
    v2 = v2
    v3 -= (v3 - (v3 + 0.1))
    if v2 {
        for fi15 := 0; fi15 < 10; fi15++ {
            if (v2 && true) {
                v3 -= float32((3.14159 / 1e10))
            } else {
                var v16 float64 = 1e10
                v17 := GenS1{F0: -1}
                v18 := GenS1{F0: -263}
                _ = v16
                _ = v17
                _ = v18
            }
            switch "foo bar" {
            case "a":
                delete(v1, "hello")
            case "\n\t":
                var v19 bool = false
                _ = v19
            }
            _ = fi15
        }
        fmt.Println(v3, v3)
        v3 = v3
    } else {
        v3 = (v3 - -2.25)
        var v20 float32 = v3
        if len(v4) > 0 {
            v4[0] = 65535
        }
        _ = v20
    }
    v21 := GenS3{F0: -1}
    var v22 float64 = -2.25
    for fi23 := 0; fi23 < 10; fi23++ {
        for fi24 := 0; fi24 < 3; fi24++ {
            v22 += v22
            v22 -= (float64(1.5) - 3.14159)
            _ = fi24
        }
        _ = fi23
    }
    v25 := v1["hello"]
    v26 := GenS4{F0: 190}
    v1[""] = ((465 > 2147483647) || ((!(42 != 685)) || (42 <= 0)))
    var v27 int64 = 0
    for fi28 := 0; fi28 < 15; fi28++ {
        v29 := v1[(("\n\t" + "foo bar") + ("foo bar" + (("\n\t" + "") + "hello")))]
        _ = fi28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
