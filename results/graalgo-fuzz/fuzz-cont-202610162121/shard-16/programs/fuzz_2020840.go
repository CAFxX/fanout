package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint) int8 {
    var v7 bool = (924 > 121)
    var v8 int = -1
    var v9 uint8 = 71
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -103
}

func main() {
    var v1 uint = 2
    v2 := make(map[int]bool)
    v3 := []uint32{1}
    var v4 float64 = 1e10
    var v5 uint = v1
    var v6 int32 = 42
    v10 := GenS1{F0: v1, F1: 128}
    _ = len(v2)
    fmt.Println(v6, v4, v4)
    var v11 int32 = -2147483648
    if false {
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(v1, v4)
            v4 += (v4 + float64(-2.25))
            _ = fi12
        }
        v13 := GenS1{F0: v1, F1: 42}
        var v14 int64 = -9223372036854775808
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 16; fi15++ {
            v1 -= 4294967295
            _ = fi15
        }
        var v16 string = ((("a" + "\n\t") + "hello") + "\n\t")
        v6 = (v6 + (v6 << (((^(v5))) & 63)))
        _ = v16
    }
    if true {
        switch "hello" {
        case "foo bar":
            if false {
                fmt.Println(v1, v5)
            }
        case "":
            v3 = append(v3, 552)
        default:
            var v17 uint = (v1 ^ (^(v5)))
            _ = v17
        }
        if true {
            if false {
                v1 -= (v1 ^ v1)
                var v18 int = 4294967295
                _ = len(v2)
                _ = v18
            }
            var v19 uint16 = 613
            for fi20 := 0; fi20 < 14; fi20++ {
                v11 += 941
                var v21 int32 = v11
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 string = ("hello" + "a")
        v11 += -431
        _ = fi22
        _ = v23
    }
    v1 += 255
    var v24 uint16 = 32767
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        v26 := GenS1{F0: uint(v24), F1: 128}
        switch (fi25 + (fi25 - (fi25 ^ 127))) {
        case 127:
            for fi27 := 0; fi27 < 5; fi27++ {
                var v28 string = "\n\t"
                _ = fi27
                _ = v28
            }
        case 333:
            var v29 uint = (v1 / ((255) | 1))
            _ = v29
        case -1:
            _ = len(v3)
        }
        var v30 int8 = 0
        _ = fi25
        _ = v26
        _ = v30
    }
    v31 := GenS1{F0: v1, F1: 42}
    if false {
        var v32 uint = v1
        v33 := GenS2{F0: 42, F1: uint32(v1)}
        if (60 < 435) {
            _ = len(v2)
        } else {
            v1 -= (v5 << ((0) & 63))
            fmt.Println(v24, v4)
        }
        _ = v32
        _ = v33
    }
    if (int(v24) >= 0) {
        if false {
            v34 := v2[-1]
            fmt.Println(v34)
            _ = v34
        }
        var v35 int64 = 1
        _ = v35
    } else {
        v36 := GenS1{F0: (^(v1)), F1: 172}
        switch 4294967295 {
        case -302:
            switch 127 {
            case 716:
                var v37 uint8 = uint8(v24)
                _ = v37
            default:
                v38 := GenS2{F0: 4294967295, F1: 37}
                _ = v38
            }
        case -626:
            for fi39 := 0; fi39 < 5; fi39++ {
                v40 := GenS3{F0: 42}
                _ = fi39
                _ = v40
            }
        case 4294967295:
            fmt.Println(v4, v11, v4)
        }
        _ = v36
    }
    var v41 int16 = 586
    v11 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
