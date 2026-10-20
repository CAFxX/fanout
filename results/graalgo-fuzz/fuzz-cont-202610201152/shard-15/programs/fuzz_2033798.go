package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint16 }
func genFn0() string {
    var v6 int64 = 35
    if true {
        v7 := GenS1{F0: 42, F1: 1, F2: 42}
        _ = v7
    } else {
        fmt.Println(v6)
        if false {
            var v8 int64 = v6
            switch "a" {
            case "hello":
                var v9 uint64 = 1
                _ = v9
            case "":
                var v10 string = ("foo bar" + ((((("foo bar" + ("\n\t" + "")) + "\n\t") + "foo bar") + "\n\t") + ""))
                _ = v10
            case "a":
                v8 += -891
            }
            _ = v8
        } else {
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 float64 = 1.5
                var v13 int16 = 487
                v14 := GenS1{F0: 2147483648, F1: 2, F2: 32767}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        }
    }
    if false {
        v6 = v6
        fmt.Println(v6)
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 int32 = 0
        var v17 int64 = (v6 & v6)
        if (3.14159 < 0.0) {
            v16 -= 1
            fmt.Println(fi15)
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = v6
    return "hello"
}
func genFn1(p0 int32) string {
    switch 289 {
    case -2147483648:
        v18 := []int{32767, 614, -1}
        _ = v18
    }
    v19 := []int{127, 4294967295, -920, 0}
    var v20 string = ("hello" + "foo bar")
    _ = p0
    _ = v19
    _ = v20
    return (v20 + (v20 + v20))
}

func main() {
    var v1 float32 = float32(0.0)
    var v2 bool = true
    var v3 bool = v2
    v4 := make(map[int]string)
    v5 := []string{}
    v1 += v1
    var v21 int64 = 0
    var v22 uint8 = 2
    v22 = (v22 >> ((1) & 63))
    v22 += 128
    v23 := GenS1{F0: 4294967295, F1: (v22 + v22), F2: 1}
    v22 -= (v22 << ((0) & 63))
    v1 -= v1
    fmt.Println(v1)
    fmt.Println(v1, v21, v3)
    var v24 float64 = 0.0
    for fi25 := 0; fi25 < 7; fi25++ {
        if false {
            v21 = 1
            if (v3 && (0.1 == v1)) {
                v26 := GenS1{F0: 65535, F1: v22, F2: 0}
                _ = len(v5)
                switch ("hello" + "\n\t") {
                case "\n\t":
                    fmt.Println(v24, v24, v2)
                }
                _ = v26
            }
            v5 = append(v5, "a")
        } else {
            v2 = (42 <= 0)
            if v3 {
                var v27 uint = 1
                _ = v27
            }
            v24 += v24
        }
        v3 = ((v1 / -2.25) >= (v1 * v1))
        if (!v3) {
            v21 = v21
            var v28 int8 = 100
            fmt.Println(v22, v2, v21)
            _ = v28
        }
        _ = fi25
    }
    v3 = (true && false)
    for fi29 := 0; fi29 < 20; fi29++ {
        var v30 float64 = v24
        _ = len(v4)
        _ = fi29
        _ = v30
    }
    v21 = (-(v21))
    v22 = (v22 & v22)
    fmt.Println(v22, v24, v24)
    fmt.Println(v22, v2, v24)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
