package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int16; F1 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -1
    var v3 string = "foo bar"
    if true {
        fmt.Println(v2, v3)
        if ((false && (702 < -1)) || (((!(42 >= -397)) && false) && (("foo bar" + ("\n\t" + v3)) == v3))) {
            v3 = "foo bar"
        }
        var v4 int = -1
        _ = v4
    }
    if (false && false) {
        v2 -= v2
        for fi5 := 0; fi5 < 18; fi5++ {
            var v6 int = fi5
            var v7 int64 = 1
            _ = fi5
            _ = v6
            _ = v7
        }
    } else {
        v2 -= (v2 & -1)
    }
    switch -499 {
    case 1:
        switch "foo bar" {
        case "hello":
            switch 0 {
            case 42:
                var v8 int8 = 127
                _ = v8
            default:
                v2 = (v2 << ((65535) & 63))
            }
        case "":
            var v9 int64 = -9223372036854775808
            _ = v9
        }
    case 9223372036854775807:
        var v10 int8 = -1
        _ = v10
    case 4294967295:
        if (!false) {
            _ = len(v1)
            v11 := GenS1{F0: 425, F1: 42}
            v12 := GenS1{F0: 0, F1: 1}
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 uint = 18446744073709551615
        v15 := GenS1{F0: -408, F1: 436}
        v16 := []int{fi13}
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    if true {
        v17 := GenS1{F0: -2147483648, F1: 42}
        switch (v3 + "\n\t") {
        case "foo bar":
            fmt.Println(v3, v2)
        case "":
            v18 := GenS2{F0: -2147483648, F1: 2147483647}
            _ = v18
        case "hello":
            _ = len(v1)
        }
        _ = v17
    } else {
        var v19 int16 = (v2 - -619)
        v2 += v2
        for fi20 := 0; fi20 < 20; fi20++ {
            v21 := GenS3{F0: 868, F1: v2}
            fmt.Println(v19)
            fmt.Println(v3, v19, v2)
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    switch v3 {
    case "foo bar":
        for fi22 := 0; fi22 < 7; fi22++ {
            if (0.0 < 0.1) {
                if false {
                    var v23 bool = false
                    v24 := GenS1{F0: 2147483647, F1: 9223372036854775807}
                    v25 := []int{32767, 536, (fi22 >> ((18446744073709551615) & 63))}
                    _ = v23
                    _ = v24
                    _ = v25
                }
                v26 := GenS1{F0: -23, F1: -817}
                var v27 uint = 65535
                _ = v26
                _ = v27
            }
            _ = fi22
        }
    case "a":
        v28 := GenS4{F0: 425, F1: 32767}
        _ = v28
    case "\n\t":
        delete(v1, 953)
    }
    v3 = v3
    v3 = v3
    v29 := GenS1{F0: 1, F1: -9223372036854775808}
    var v30 int = 1
    var v31 int16 = (v2 >> ((0) & 63))
    fmt.Println(v31, v31, v31)
    if (127 <= 127) {
        if true {
            v32 := GenS4{F0: 1, F1: -1}
            for fi33 := 0; fi33 < 7; fi33++ {
                v3 = v3
                v34 := GenS5{F0: (v2 >> ((18446744073709551615) & 63)), F1: v30}
                _ = fi33
                _ = v34
            }
            _ = v32
        }
    }
    switch v3 {
    case "foo bar":
        v1[-9223372036854775808] = 42
    case "\n\t":
        v3 = v3
    case "":
        v2 = 1
    default:
        v35 := []int{v30, (v30 >> ((0) & 63))}
        _ = v35
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
