package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }

func main() {
    v1 := []string{(("" + "hello") + "\n\t"), "hello", "a"}
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    v4 := []uint32{92, 2, 4294967295, 42}
    v5 := []string{"foo bar", "hello", ("\n\t" + "a"), ("foo bar" + ("" + ("hello" + "")))}
    v6 := []string{}
    var v7 uint64 = 2
    v8 := v3[442]
    var v9 bool = true
    var v10 int8 = -45
    if (255 >= 127) {
        v8 -= v8
        switch "hello" {
        case "\n\t":
            var v11 int64 = -448
            _ = v11
        case "hello":
            var v12 float64 = -2.25
            _ = v12
        case "a":
            fmt.Println(v7, v7)
        default:
            var v13 int16 = 789
            _ = v13
        }
    } else {
        v7 = (v7 & 2)
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        if (0 >= 1) {
            switch (v8 | v8) {
            case 32767:
                var v15 int64 = 361
                _ = v15
            }
        }
        _ = fi14
    }
    _ = len(v2)
    v8 = v8
    if ((653 != 0) && (2147483648 > 2147483648)) {
        v16 := GenS1{F0: (v7 + (^(v7)))}
        v3[42] = (v8 >> ((65535) & 63))
        _ = v16
    } else {
        for fi17 := 0; fi17 < 18; fi17++ {
            fmt.Println(v10, v8)
            v8 = -9223372036854775808
            var v18 int8 = v10
            _ = fi17
            _ = v18
        }
        if (1 > 2) {
            var v19 int64 = -477
            fmt.Println(v19, v8, v19)
            _ = v19
        }
    }
    v8 = 32768
    v10 -= -128
    var v20 string = "hello"
    if v9 {
        v21 := GenS2{F0: v8}
        if ((v9 && v9) && (v20 == (v20 + v20))) {
            if (!(true || (116 >= int64(v7)))) {
                var v22 int64 = 1
                _ = v22
            }
        }
        _ = v21
    } else {
        var v23 bool = (42 == -37)
        _ = v23
    }
    var v24 uint64 = (v7 | 2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
