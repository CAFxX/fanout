package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0() uint32 {
    if (-9223372036854775808 < 0) {
        var v5 uint = 42
        _ = v5
    }
    var v6 uint64 = 1
    v7 := []int{-496, 42}
    if true {
        var v8 int32 = -1
        var v9 bool = true
        if v9 {
            var v10 uint64 = v6
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := GenS1{F0: 318}
                _ = fi11
                _ = v12
            }
            var v13 float32 = 3.14159
            _ = v10
            _ = v13
        }
        _ = v8
        _ = v9
    } else {
        fmt.Println(v6)
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        switch ((("" + ("foo bar" + (("\n\t" + "foo bar") + "\n\t"))) + ("foo bar" + (("hello" + (("\n\t" + "\n\t") + "foo bar")) + "hello"))) + (("a" + (("\n\t" + "\n\t") + ("hello" + "a"))) + ("a" + "\n\t"))) {
        case "foo bar":
            switch 719 {
            case 32767:
                v6 += 0
            case -585:
                var v14 bool = false
                _ = v14
            }
        default:
            v6 = v6
        }
    }
    if (!true) {
        var v15 string = ("foo bar" + "")
        _ = v15
    }
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int32 = -1
    var v2 uint16 = 128
    var v3 float64 = (1.5 / (3.14159 - 1e10))
    v4 := make(map[int]bool)
    for fi16 := 0; fi16 < 1; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            fmt.Println(v3)
            _ = fi17
        }
        _ = fi16
    }
    v3 += v3
    fmt.Println(v2, v3, v3)
    var v18 int8 = -128
    v19 := GenS2{F0: 125}
    fmt.Println(v3, v1)
    var v20 int = 1
    var v21 int = (v20 - v20)
    v18 -= (v18 >> ((0) & 63))
    var v22 int = v21
    var v23 int64 = 42
    var v24 uint8 = 1
    v25 := []int{}
    fmt.Println(v22)
    v26 := GenS1{F0: (v2 | v2)}
    delete(v4, v20)
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 uint64 = 4294967295
        var v29 int16 = 293
        if true {
            if false {
                v30 := GenS1{F0: v2}
                v20 = -494
                v31 := GenS2{F0: v24}
                _ = v30
                _ = v31
            } else {
                var v32 int32 = v1
                _ = v32
            }
            fmt.Println(v2, v23, v21)
            var v33 float64 = ((float64(3.14159) / 1.5) - ((0.1 + (v3 + float64(0.1))) * 0.0))
            _ = v33
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
