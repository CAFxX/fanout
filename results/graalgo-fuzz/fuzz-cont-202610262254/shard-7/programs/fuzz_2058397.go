package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{559, 326}
    v3 := make(map[int]string)
    var v4 uint8 = 116
    var v5 uint = 0
    var v6 float64 = (0.0 * 1e10)
    if false {
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS1{F0: 42}
            if true {
                var v9 float64 = 0.1
                _ = v9
            }
            _ = fi7
            _ = v8
        }
        var v10 int = 32768
        _ = v10
    } else {
        fmt.Println(v5)
        if (((-1 == -1) || (3.14159 != 3.14159)) || (!(false || (0 <= 393)))) {
            switch "foo bar" {
            case "a":
                var v11 int64 = -9223372036854775808
                _ = v11
            }
            var v12 uint8 = (v4 % ((255) | 1))
            _ = v12
        } else {
            v2 = append(v2, 4294967295)
            v13 := GenS1{F0: 810}
            _ = v13
        }
        var v14 int64 = -9223372036854775808
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 uint16 = 836
            if true {
                var v18 int8 = 127
                _ = v18
            } else {
                v4 = (v4 >> ((18446744073709551615) & 63))
                var v19 uint32 = 847
                _ = v19
            }
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    var v20 uint8 = (v4 - v4)
    var v21 uint = (v5 * 931)
    if true {
        switch (("" + ("\n\t" + "a")) + "foo bar") {
        case "\n\t":
            var v22 uint8 = (v20 - v4)
            _ = v22
        }
        fmt.Println(v20)
    } else {
        var v23 uint64 = 42
        _ = v23
    }
    var v24 uint8 = (v20 / ((v4) | 1))
    v24 = (v20 - v24)
    fmt.Println(v20, v5)
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 uint16 = 1
        v27 := GenS1{F0: 42}
        var v28 uint16 = (^(v26))
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    v29 := GenS1{F0: 275}
    var v30 uint64 = uint64(v4)
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
