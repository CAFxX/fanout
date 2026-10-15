package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
type GenS2 struct { F0 int8; F1 int64; F2 int64 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    var v3 int16 = -1
    var v4 uint8 = 1
    var v5 uint32 = 65535
    v3 = v3
    v4 += v4
    if true {
        v3 = (v3 % ((-207) | 1))
        switch 9223372036854775807 {
        case -2147483648:
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v4, v5)
            case "":
                switch 761 {
                case -1:
                    var v6 int = 83
                    _ = v6
                }
            case "foo bar":
                v7 := GenS1{F0: v4, F1: 2, F2: 1}
                _ = v7
            default:
                var v8 string = ("" + "hello")
                _ = v8
            }
        case -454:
            v9 := GenS1{F0: 0, F1: 562, F2: -206}
            _ = v9
        }
        var v10 uint16 = 406
        _ = v10
    } else {
        v4 += v4
        v11 := []int{2147483647, -967, -24}
        v5 -= 1
        _ = v11
    }
    var v12 int16 = (v3 >> ((97) & 63))
    if false {
        v3 = (-(v3))
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            v5 -= (v5 ^ 2147483647)
            switch (("hello" + "\n\t") + "hello") {
            case "foo bar":
                v5 = v5
            }
            v3 = v3
            _ = fi13
        }
        switch "a" {
        case "foo bar":
            v14 := []int{144}
            _ = v14
        case "a":
            v15 := GenS2{F0: 42, F1: 772, F2: -801}
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 bool = true
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 uint16 = 42
        v20 := GenS2{F0: -1, F1: 1, F2: 9223372036854775807}
        var v21 string = "foo bar"
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    var v22 uint64 = 42
    var v23 uint16 = 523
    for fi24 := 0; fi24 < 4; fi24++ {
        v1["hello"] = uint32(v23)
        _ = fi24
    }
    v25 := []int{}
    if false {
        if (((("foo bar" + "\n\t") + ("\n\t" + "a")) == "\n\t") || (!((v23 + (v23 << ((42) & 63))) < (^(v23))))) {
            fmt.Println(v4, v3)
        }
        fmt.Println(v3)
    }
    if (1.5 < 0.0) {
        v26 := GenS2{F0: 0, F1: -495, F2: -512}
        var v27 bool = true
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        switch (("foo bar" + "a") + (("foo bar" + "a") + "foo bar")) {
        case "foo bar":
            v23 -= (v23 / ((1) | 1))
        case "a":
            fmt.Println(v4, v22, v23)
        default:
            for fi29 := 0; fi29 < 3; fi29++ {
                v12 += (v12 + (v3 & 1000))
                v3 += 1
                _ = fi29
            }
        }
        v4 = (v4 | v4)
        v23 = uint16(v4)
        _ = fi28
    }
    fmt.Println(v5)
    if (32767 <= 1) {
        if false {
            fmt.Println(v5, v3)
        }
    }
    if false {
        if false {
            var v30 int8 = -124
            _ = v30
        } else {
            v12 = 42
            if (!((0 > int64(v23)) || true)) {
                v12 -= int16(v23)
                var v31 uint8 = 1
                _ = v31
            }
        }
        switch 127 {
        case 1:
            v5 = (v5 << ((uint(v23)) & 63))
        case 4294967295:
            fmt.Println(v22, v5)
        case 32767:
            for fi32 := 0; fi32 < 4; fi32++ {
                switch "hello" {
                case "foo bar":
                    v4 -= v4
                case "a":
                    _ = len(v25)
                }
                var v33 uint64 = v22
                _ = fi32
                _ = v33
            }
        default:
            v34 := GenS3{F0: -9223372036854775808, F1: v23, F2: 42}
            _ = v34
        }
        v12 = -786
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
