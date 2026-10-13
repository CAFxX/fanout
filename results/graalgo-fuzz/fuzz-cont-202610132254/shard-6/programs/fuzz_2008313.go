package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int32 }
type GenS2 struct { F0 int32; F1 int; F2 string }

func main() {
    var v1 int64 = -76
    var v2 int64 = (v1 + (v1 % ((v1) | 1)))
    var v3 int16 = 0
    v4 := []uint32{2, 139, 320}
    var v5 int8 = -1
    var v6 int32 = 1
    v7 := GenS1{F0: 127, F1: 474, F2: 1}
    switch "\n\t" {
    case "foo bar":
        v2 -= 9223372036854775807
    case "":
        v8 := GenS1{F0: 32767, F1: 310, F2: (v6 / ((v6) | 1))}
        _ = v8
    default:
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(v2)
            switch "\n\t" {
            case "":
                v2 -= v1
            case "\n\t":
                v10 := GenS1{F0: (fi9 & 0), F1: 2, F2: 2147483647}
                _ = v10
            default:
                for fi11 := 0; fi11 < 17; fi11++ {
                    fmt.Println(fi9)
                    var v12 uint = 356
                    v3 -= (v3 - int16(fi11))
                    _ = fi11
                    _ = v12
                }
            }
            var v13 uint16 = 65535
            _ = fi9
            _ = v13
        }
    }
    if (((-(v3)) == 1) || true) {
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 int8 = 41
            var v16 int32 = v6
            var v17 bool = true
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        var v18 uint64 = 1
        v19 := GenS1{F0: 9223372036854775807, F1: 2, F2: v6}
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        v4 = append(v4, 1)
        _ = fi20
    }
    switch -2147483648 {
    case 128:
        var v21 int64 = 0
        _ = v21
    case 127:
        if true {
            v22 := GenS1{F0: 32767, F1: 4294967295, F2: v6}
            var v23 int8 = v5
            v4 = append(v4, 4294967295)
            _ = v22
            _ = v23
        } else {
            for fi24 := 0; fi24 < 18; fi24++ {
                v3 = v3
                for fi25 := 0; fi25 < 17; fi25++ {
                    var v26 int = fi25
                    _ = fi25
                    _ = v26
                }
                if len(v4) > 0 {
                    v4[0] = 42
                }
                _ = fi24
            }
            v6 = (v6 ^ v6)
        }
    default:
        v1 -= (v1 + (v2 + (v1 + v2)))
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        for fi28 := 0; fi28 < 20; fi28++ {
            var v29 uint64 = 0
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    var v30 uint64 = 1
    switch "\n\t" {
    case "\n\t":
        _ = len(v4)
    default:
        for fi31 := 0; fi31 < 18; fi31++ {
            v5 = int8(fi31)
            _ = fi31
        }
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        fmt.Println(v3, fi32, v3)
        for fi33 := 0; fi33 < 1; fi33++ {
            fmt.Println(v2)
            if false {
                v5 = (v5 / ((v5) | 1))
                v4 = append(v4, 5)
                v34 := GenS1{F0: (fi33 * -953), F1: (^(v30)), F2: (-(v6))}
                _ = v34
            }
            _ = fi33
        }
        _ = fi32
    }
    if true {
        if true {
            v1 = 0
        } else {
            if true {
                v6 = (v6 >> ((743) & 63))
                var v35 int32 = (^(v6))
                for fi36 := 0; fi36 < 15; fi36++ {
                    var v37 int64 = (v2 & int64(v5))
                    _ = fi36
                    _ = v37
                }
                _ = v35
            }
            for fi38 := 0; fi38 < 16; fi38++ {
                var v39 int = (fi38 >> ((0) & 63))
                var v40 int64 = v1
                _ = fi38
                _ = v39
                _ = v40
            }
        }
        v41 := GenS2{F0: v6, F1: 0, F2: "\n\t"}
        _ = v41
    } else {
        if true {
            var v42 uint = 206
            _ = v42
        } else {
            v1 = (v1 & (v2 & (v2 / ((-1) | 1))))
        }
        v43 := GenS2{F0: v6, F1: 128, F2: ""}
        fmt.Println(v2, v6, v5)
        _ = v43
    }
    v44 := GenS1{F0: 1, F1: 42, F2: (-(v6))}
    v2 -= (^(v2))
    var v45 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
