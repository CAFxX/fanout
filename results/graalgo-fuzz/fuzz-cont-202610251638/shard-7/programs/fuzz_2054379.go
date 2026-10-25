package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 662
    v3 := make(map[uint32]string)
    var v4 uint16 = 2
    var v5 int8 = int8(v4)
    var v6 uint32 = 65535
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int64 = -9223372036854775808
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: (v2 - v2), F1: -1, F2: 65535}
    v10 := GenS1{F0: (v2 * v2), F1: 9223372036854775807, F2: 4294967295}
    v11 := []int{-9223372036854775808}
    switch ("hello" + "") {
    case "\n\t":
        v12 := GenS2{F0: 42}
        _ = v12
    case "hello":
        if true {
            fmt.Println(v6, v6)
            var v13 float64 = 1.5
            _ = v13
        } else {
            switch ("a" + "") {
            case "a":
                v6 = 2
            default:
                _ = len(v11)
            }
        }
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        v4 -= 0
        if len(v11) > 0 {
            v11[0] = (fi14 << ((46) & 63))
        }
        switch (fi14 ^ 127) {
        case 9223372036854775807:
            if false {
                var v15 uint64 = 42
                var v16 int = 128
                _ = v15
                _ = v16
            } else {
                var v17 string = "foo bar"
                if true {
                    v18 := GenS3{F0: 9223372036854775807}
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v6 -= 2147483647
                        var v20 int32 = -2147483648
                        _ = fi19
                        _ = v20
                    }
                    _ = v18
                }
                _ = v17
            }
        }
        _ = fi14
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        _ = len(v3)
        _ = fi21
    }
    v22 := GenS3{F0: 42}
    v23 := GenS3{F0: -981}
    if true {
        fmt.Println(v6, v5)
    } else {
        var v24 uint8 = 255
        for fi25 := 0; fi25 < 14; fi25++ {
            for fi26 := 0; fi26 < 9; fi26++ {
                v4 += 433
                fmt.Println(fi26, v4, fi25)
                var v27 uint32 = (v6 | v6)
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
        _ = v24
    }
    v4 += (v4 + 255)
    switch ("hello" + "foo bar") {
    case "a":
        var v28 int = -734
        _ = v28
    case "foo bar":
        var v29 uint = 0
        _ = v29
    case "\n\t":
        var v30 int16 = 0
        _ = v30
    }
    fmt.Println(v6, v2, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
