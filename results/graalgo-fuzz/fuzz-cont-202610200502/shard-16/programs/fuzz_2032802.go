package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int32; F2 int }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int64 }
func genFn0(p0 int8, p1 int16, p2 uint8) uint16 {
    if (1e10 != 1e10) {
        p0 += (-(p0))
    }
    v4 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 32767
}

func main() {
    var v1 uint = 0
    v2 := []string{("" + ""), ((("\n\t" + (("\n\t" + "foo bar") + ("\n\t" + "foo bar"))) + "a") + ("\n\t" + ""))}
    var v3 uint64 = 42
    switch ((("" + "a") + "") + "\n\t") {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "":
        var v5 int64 = -965
        _ = v5
    }
    switch 822 {
    case 32767:
        v6 := GenS1{F0: 2147483647, F1: 4294967295, F2: 2}
        _ = v6
    case 127:
        v1 = v1
    }
    if (!false) {
        v7 := GenS1{F0: 1, F1: 423, F2: 42}
        v8 := GenS2{F0: uint32(v1), F1: -1, F2: 4294967295}
        fmt.Println(v3, v1)
        _ = v7
        _ = v8
    }
    var v9 int8 = -1
    if true {
        v10 := GenS2{F0: 2, F1: 604, F2: 32767}
        _ = v10
    } else {
        var v11 int32 = 42
        fmt.Println(v3, v3)
        _ = v11
    }
    var v12 int64 = -173
    switch -9223372036854775808 {
    case 32767:
        if len(v2) > 0 {
            v2[0] = ""
        }
    case 528:
        for fi13 := 0; fi13 < 19; fi13++ {
            if false {
                v14 := GenS1{F0: fi13, F1: fi13, F2: 0}
                v12 = (-(v12))
                v3 = v3
                _ = v14
            }
            _ = fi13
        }
    default:
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 bool = false
            _ = fi15
            _ = v16
        }
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    var v17 uint = (v1 >> ((4294967295) & 63))
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 uint64 = uint64(v17)
        var v20 int64 = v12
        v3 -= (v19 + 18446744073709551615)
        _ = fi18
        _ = v19
        _ = v20
    }
    switch 32767 {
    case -772:
        var v21 uint16 = genFn0(int8(v12), -343, 81)
        _ = v21
    case 32767:
        switch int(v12) {
        case 1:
            var v22 int16 = 1000
            _ = v22
        default:
            fmt.Println(v1, v1)
        }
    case 2147483647:
        if len(v2) > 0 {
            v2[0] = ("\n\t" + "a")
        }
    default:
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 uint32 = 0
            for fi25 := 0; fi25 < 14; fi25++ {
                v26 := GenS3{F0: 32767, F1: (v12 * -9223372036854775808)}
                if false {
                    if true {
                        var v27 uint32 = (v24 - (^(v24)))
                        var v28 uint16 = 32767
                        _ = v27
                        _ = v28
                    } else {
                        var v29 uint = v1
                        _ = v29
                    }
                } else {
                    var v30 uint = uint(v24)
                    v9 = 42
                    var v31 int32 = 42
                    _ = len(v2)
                    fmt.Println(v24, fi23, v12)
                    _ = v30
                    _ = v31
                }
                _ = fi25
                _ = v26
            }
            v32 := GenS4{F0: v12}
            _ = fi23
            _ = v24
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
