package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 int8; F2 int32 }
func genFn0(p0 int32, p1 int8) uint32 {
    for fi7 := 0; fi7 < 19; fi7++ {
        p0 -= (p0 << ((42) & 63))
        var v8 int = (^(fi7))
        var v9 uint32 = 4294967295
        _ = fi7
        _ = v8
        _ = v9
    }
    p1 = (p1 | (-(p1)))
    var v10 int64 = -9223372036854775808
    switch 1 {
    case 1:
        var v11 int32 = (p0 << ((2) & 63))
        _ = v11
    case -1:
        for fi12 := 0; fi12 < 15; fi12++ {
            switch "\n\t" {
            case "":
                p1 = (p1 / (((^(p1))) | 1))
            }
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 int8 = p1
                _ = fi13
                _ = v14
            }
            fmt.Println(p0, v10, p1)
            _ = fi12
        }
    case 128:
        p1 = 100
    }
    _ = p0
    _ = p1
    _ = v10
    return 0
}
func genFn1(p0 int32, p1 string) uint {
    p0 -= 1
    v15 := []int{}
    v16 := GenS1{F0: (!(708 > 18446744073709551615))}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 42
}

func main() {
    v1 := []uint32{114, 102, 449, 317}
    v2 := []string{("foo bar" + ("hello" + "hello")), (("hello" + "foo bar") + ""), "\n\t"}
    var v3 uint32 = 1
    v4 := make(map[uint32]uint32)
    var v5 float32 = -2.25
    v6 := make(map[string]int)
    if (false && (!(196 == 1))) {
        var v17 uint16 = 42
        _ = v17
    }
    fmt.Println(v5)
    v18 := GenS2{F0: 9223372036854775807, F1: 2, F2: 882}
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS1{F0: (!(!true))}
        if (-41 > 127) {
            v21 := GenS3{F0: 256}
            for fi22 := 0; fi22 < 13; fi22++ {
                v3 = v3
                v5 = float32(((0.0 / -2.25) * ((-2.25 + (0.0 * ((1e10 / 0.1) - float64(v5)))) + 0.0)))
                v5 += float32(-2.25)
                delete(v6, "hello")
                _ = fi22
            }
            v5 = v5
            _ = v21
        } else {
            var v23 uint16 = 255
            _ = v23
        }
        _ = fi19
        _ = v20
    }
    v24 := GenS4{F0: (v3 << ((genFn1(-1, ("a" + "hello"))) & 63)), F1: 191}
    var v25 float32 = 1.5
    var v26 bool = (127 == 1)
    var v27 string = ""
    if len(v1) > 0 {
        v1[0] = v3
    }
    switch -270 {
    case 1:
        switch 253 {
        case 4294967295:
            if len(v1) > 0 {
                v1[0] = 643
            }
        case -602:
            fmt.Println(v25, v27, v3)
        }
    case -9223372036854775808:
        v5 -= ((((-2.25 * v25) / v25) / -2.25) + v5)
    case 127:
        for fi28 := 0; fi28 < 20; fi28++ {
            fmt.Println(v26, fi28, fi28)
            v26 = ("\n\t" < "foo bar")
            _ = fi28
        }
    }
    v29 := GenS5{F0: 2}
    var v30 int32 = -280
    v31 := GenS5{F0: 4294967295}
    switch 0 {
    case -9223372036854775808:
        fmt.Println(v5, v30)
    case -738:
        v32 := GenS6{F0: 18446744073709551615, F1: 1, F2: (v30 ^ 1)}
        _ = v32
    case 127:
        for fi33 := 0; fi33 < 2; fi33++ {
            v6["\n\t"] = fi33
            _ = fi33
        }
    }
    if ("hello" < "foo bar") {
        var v34 uint = 0
        _ = v34
    } else {
        v26 = (v26 || true)
    }
    var v35 uint16 = 473
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
