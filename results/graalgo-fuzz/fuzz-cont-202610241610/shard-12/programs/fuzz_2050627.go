package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }
type GenS2 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    var v3 int64 = -1
    v4 := make(map[int]bool)
    v5 := make(map[int]int)
    var v6 uint32 = 0
    var v7 int16 = -535
    switch "" {
    case "":
        for fi8 := 0; fi8 < 15; fi8++ {
            v5[fi8] = (^(fi8))
            _ = fi8
        }
    case "foo bar":
        fmt.Println(v3, v3, v7)
    }
    v9 := GenS1{F0: -775, F1: 505, F2: 1}
    switch ("hello" + "hello") {
    case "":
        var v10 string = ("\n\t" + ("foo bar" + "\n\t"))
        _ = v10
    case "a":
        v7 -= (v7 % ((1) | 1))
    case "foo bar":
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS1{F0: (fi11 % ((-2147483648) | 1)), F1: 2, F2: -1}
            _ = fi11
            _ = v12
        }
    }
    var v13 uint32 = 4294967295
    if (!(3.14159 > (3.14159 - 1.5))) {
        if false {
            v6 = 720
            var v14 uint = 0
            v15 := []int{32768, 127, -1}
            _ = v14
            _ = v15
        } else {
            v13 = v13
            for fi16 := 0; fi16 < 12; fi16++ {
                var v17 bool = true
                _ = fi16
                _ = v17
            }
        }
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        if (2 < 255) {
            fmt.Println(v3)
            var v19 uint16 = 0
            v20 := GenS1{F0: (fi18 | -9223372036854775808), F1: 42, F2: 127}
            _ = v19
            _ = v20
        } else {
            v21 := GenS2{F0: -1, F1: 42}
            var v22 int32 = -260
            _ = v21
            _ = v22
        }
        fmt.Println(fi18, v6, fi18)
        _ = fi18
    }
    if ((2 <= 18446744073709551615) && (((!true) || false) || true)) {
        for fi23 := 0; fi23 < 18; fi23++ {
            v24 := []int{-2147483648}
            _ = fi23
            _ = v24
        }
        for fi25 := 0; fi25 < 7; fi25++ {
            var v26 uint = 42
            fmt.Println(fi25)
            _ = fi25
            _ = v26
        }
    } else {
        switch -1 {
        case 351:
            var v27 int64 = v3
            _ = v27
        }
        v28 := []int{-2147483648}
        _ = v28
    }
    if false {
        v6 = v13
        var v29 bool = false
        _ = v29
    } else {
        var v30 float64 = (3.14159 + (((0.0 * 0.0) - 0.1) - -2.25))
        _ = len(v4)
        fmt.Println(v30)
        _ = v30
    }
    switch "a" {
    case "foo bar":
        var v31 float32 = (3.14159 - 1.5)
        _ = v31
    default:
        var v32 int8 = 127
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
