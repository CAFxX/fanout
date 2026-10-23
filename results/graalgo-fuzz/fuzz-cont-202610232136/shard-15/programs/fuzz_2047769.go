package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
func genFn0(p0 int32, p1 uint64) uint16 {
    var v5 int64 = 42
    fmt.Println(p1, v5, p1)
    var v6 bool = false
    v6 = (0 <= -32768)
    p1 = (p1 ^ p1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := []uint32{0, 2, 363, 931}
    var v2 uint64 = 18446744073709551615
    v3 := []string{"", (((("\n\t" + "foo bar") + "foo bar") + "a") + ""), "hello"}
    v4 := []string{"a", "a", (("hello" + "") + ""), "foo bar"}
    var v7 uint64 = (v2 % ((v2) | 1))
    for fi8 := 0; fi8 < 5; fi8++ {
        if ((false && (958 > 18446744073709551615)) || (fi8 >= 316)) {
            var v9 float64 = ((float64(1e10) - 0.1) + 0.0)
            _ = v9
        }
        _ = fi8
    }
    var v10 int = 108
    var v11 bool = true
    if len(v3) > 0 {
        v3[0] = (((("hello" + "foo bar") + "hello") + "foo bar") + ("a" + "\n\t"))
    }
    v12 := GenS1{F0: 42}
    for fi13 := 0; fi13 < 9; fi13++ {
        switch (v10 >> ((2) & 63)) {
        case 127:
            v14 := GenS1{F0: v10}
            _ = v14
        case -18:
            v7 = v7
        case 9223372036854775807:
            var v15 bool = v11
            _ = v15
        }
        if len(v1) > 0 {
            v1[0] = 377
        }
        _ = fi13
    }
    v11 = v11
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        if len(v3) > 0 {
            v3[0] = "a"
        }
        var v17 uint64 = (v2 * (v2 | 4294967295))
        v10 -= (v10 ^ (fi16 ^ 32768))
        _ = fi16
        _ = v17
    }
    switch "hello" {
    case "a":
        var v18 int64 = -9223372036854775808
        _ = v18
    default:
        if len(v1) > 0 {
            v1[0] = 328
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        if v11 {
            _ = len(v4)
        } else {
            var v20 uint16 = 816
            fmt.Println(v10, v11)
            var v21 int = (v10 >> ((2) & 63))
            _ = v20
            _ = v21
        }
        if ((1 <= 1) || (1 >= 2147483647)) {
            v22 := GenS2{F0: (fi19 >> ((357) & 63))}
            _ = v22
        } else {
            switch ("a" + "hello") {
            case "a":
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 int16 = 919
                    _ = fi23
                    _ = v24
                }
            }
            if v11 {
                v2 += (v2 & (v7 - (v2 % ((v7) | 1))))
                v2 += 727
                if false {
                    for fi25 := 0; fi25 < 20; fi25++ {
                        v2 -= v7
                        var v26 int64 = -1
                        _ = fi25
                        _ = v26
                    }
                    v11 = v11
                    switch "hello" {
                    case "a":
                        fmt.Println(fi19, v11, v7)
                    case "":
                        v27 := GenS1{F0: -183}
                        _ = v27
                    }
                }
            } else {
                v28 := GenS1{F0: fi19}
                _ = v28
            }
        }
        var v29 uint64 = v2
        _ = fi19
        _ = v29
    }
    var v30 int32 = -1
    v2 = (v7 | 0)
    var v31 uint16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v7
}
