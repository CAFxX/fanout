package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint32; F1 uint8; F2 int8 }
func genFn0(p0 uint8) uint {
    v5 := []int{4294967295}
    for fi6 := 0; fi6 < 18; fi6++ {
        if false {
            p0 = (p0 >> ((uint(p0)) & 63))
            var v7 uint = 65535
            switch ("" + "\n\t") {
            case "hello":
                v5 = append(v5, (fi6 << ((v7) & 63)))
            case "":
                v7 += 0
            default:
                var v8 uint16 = 65535
                _ = v8
            }
            _ = v7
        }
        v9 := GenS1{F0: 1, F1: 74}
        fmt.Println(p0, fi6)
        _ = fi6
        _ = v9
    }
    var v10 int64 = -9223372036854775808
    _ = p0
    _ = v10
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint64 = 887
    var v2 uint64 = (^(v1))
    var v3 int = 986
    v4 := []int{}
    switch ("" + "foo bar") {
    case "\n\t":
        if (((713 == 0) || true) || true) {
            v11 := GenS2{F0: ("" + "a"), F1: true}
            _ = v11
        } else {
            switch -1 {
            case -731:
                v1 = 898
            case 584:
                v2 = (^(v2))
            }
        }
    case "foo bar":
        if (!(("a" + "\n\t") < ("\n\t" + "a"))) {
            for fi12 := 0; fi12 < 11; fi12++ {
                v1 += 42
                _ = fi12
            }
            if len(v4) > 0 {
                v4[0] = (v3 | v3)
            }
        }
    default:
        var v13 int32 = -1
        _ = v13
    }
    if false {
        if (42 < 0) {
            var v14 uint32 = 646
            _ = v14
        }
    }
    v3 = v3
    var v15 int16 = -1
    for fi16 := 0; fi16 < 3; fi16++ {
        v2 = v1
        v15 = v15
        _ = fi16
    }
    if (int32(v3) >= 2147483647) {
        var v17 uint = 2
        v2 = 1
        var v18 float64 = 3.14159
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 18; fi19++ {
            v15 = 1000
            v20 := GenS1{F0: -1, F1: -43}
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 11; fi21++ {
            v22 := GenS3{F0: 65535, F1: 128, F2: 42}
            var v23 int16 = v15
            v24 := GenS1{F0: -809, F1: -1}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
    }
    v1 = (v1 - (v2 << ((genFn0(179)) & 63)))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
