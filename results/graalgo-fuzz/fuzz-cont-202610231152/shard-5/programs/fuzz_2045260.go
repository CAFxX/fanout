package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0() uint8 {
    switch 127 {
    case 0:
        if false {
            for fi6 := 0; fi6 < 7; fi6++ {
                v7 := []int{-9223372036854775808, (^(fi6)), (^(fi6)), (fi6 | fi6)}
                _ = fi6
                _ = v7
            }
            var v8 int = 127
            fmt.Println(v8)
            _ = v8
        } else {
            var v9 int32 = 42
            _ = v9
        }
    case 42:
        switch ((("a" + "foo bar") + "foo bar") + "hello") {
        case "a":
            v10 := GenS1{F0: -1, F1: 2147483647}
            _ = v10
        }
    case 9223372036854775807:
        switch "foo bar" {
        case "a":
            var v11 int = 32768
            _ = v11
        case "":
            var v12 int16 = 0
            _ = v12
        }
    }
    var v13 int = 1
    v13 += (^(v13))
    _ = v13
    return 199
}

func main() {
    var v1 int16 = 1
    var v2 int32 = -991
    var v3 string = ""
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]string)
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := v5[65535]
        var v16 uint32 = 2147483648
        _ = fi14
        _ = v15
        _ = v16
    }
    v2 = v2
    v17 := GenS1{F0: 42, F1: 847}
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := GenS1{F0: -284, F1: 915}
        v20 := v4[(^(fi18))]
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int32 = v2
    var v22 uint64 = 169
    if false {
        var v23 string = v3
        _ = v23
    } else {
        v3 = v3
    }
    var v24 uint16 = 65535
    v1 = (v1 * 42)
    v1 -= v1
    var v25 int32 = 37
    v26 := v5[4294967295]
    var v27 int64 = -990
    switch 1 {
    case 4294967295:
        for fi28 := 0; fi28 < 4; fi28++ {
            v3 = (("a" + v3) + v26)
            v29 := GenS1{F0: (-(v27)), F1: 2}
            _ = fi28
            _ = v29
        }
    default:
        v4[-358] = 126
    }
    fmt.Println(v22)
    if ((!(505 != 0)) || (false || ((v2 << ((344) & 63)) > -32))) {
        if (!false) {
            v30 := GenS1{F0: v27, F1: 943}
            _ = v30
        } else {
            for fi31 := 0; fi31 < 14; fi31++ {
                for fi32 := 0; fi32 < 1; fi32++ {
                    v22 = (v22 + 648)
                    _ = fi32
                }
                _ = len(v4)
                v33 := []int{}
                _ = fi31
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 7; fi34++ {
            switch fi34 {
            case 128:
                for fi35 := 0; fi35 < 17; fi35++ {
                    v36 := GenS2{F0: 42}
                    _ = fi35
                    _ = v36
                }
            }
            var v37 int32 = (v2 * (-(v21)))
            v1 += v1
            _ = fi34
            _ = v37
        }
    } else {
        var v38 string = v26
        v3 = v26
        _ = v38
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        v24 += (v24 % ((v24) | 1))
    case -9223372036854775808:
        fmt.Println(v27, v2)
    case 1:
        if (2 > 65535) {
            var v39 float32 = -2.25
            var v40 int16 = (v1 >> ((1) & 63))
            _ = v39
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
