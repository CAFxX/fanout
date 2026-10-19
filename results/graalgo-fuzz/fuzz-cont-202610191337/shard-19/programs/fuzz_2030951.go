package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
func genFn0(p0 bool) uint8 {
    v4 := GenS1{F0: 1}
    var v5 int = -261
    if (!("" >= "hello")) {
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 bool = ("hello" <= "hello")
            switch 1 {
            case 127:
                var v8 int16 = 0
                _ = v8
            }
            v9 := GenS1{F0: 2}
            _ = fi6
            _ = v7
            _ = v9
        }
    }
    v10 := []int{}
    _ = p0
    _ = v10
    _ = v4
    _ = v5
    return 1
}
func genFn1() uint8 {
    var v11 int64 = 0
    v12 := GenS1{F0: 711}
    if false {
        var v13 int = 1
        var v14 uint = 2
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 bool = (42 > 539)
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 2; fi18++ {
                if false {
                    switch ("" + ("" + "\n\t")) {
                    case "foo bar":
                        v19 := []int{9223372036854775807, -339, fi18}
                        _ = v19
                    case "hello":
                        var v20 uint64 = 4294967295
                        _ = v20
                    case "a":
                        var v21 int = (^(v13))
                        _ = v21
                    default:
                        var v22 int64 = v11
                        _ = v22
                    }
                } else {
                    var v23 int = (^(v13))
                    _ = v23
                }
                v24 := GenS1{F0: 1}
                _ = fi18
                _ = v24
            }
            var v25 int8 = 1
            _ = fi15
            _ = v25
        }
        _ = v13
        _ = v14
    }
    _ = v11
    _ = v12
    return 128
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    var v3 uint = 2
    v26 := []int{2147483647, 42, -34}
    if len(v26) > 0 {
        v26[0] = 32767
    }
    switch "a" {
    case "\n\t":
        v3 = (v3 >> ((716) & 63))
    }
    fmt.Println(v3)
    var v27 int8 = 100
    v28 := GenS1{F0: 890}
    if len(v26) > 0 {
        v26[0] = 2147483647
    }
    if (42 <= 65535) {
        switch 4294967295 {
        case 4294967295:
            v29 := GenS1{F0: 1}
            _ = v29
        case -573:
            v27 = v27
        case -444:
            fmt.Println(v27)
        }
        v30 := GenS1{F0: 4294967295}
        _ = v30
    } else {
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 float32 = (0.0 / (-2.25 + -2.25))
            v33 := GenS1{F0: 42}
            _ = fi31
            _ = v32
            _ = v33
        }
        v34 := GenS2{F0: 383}
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
