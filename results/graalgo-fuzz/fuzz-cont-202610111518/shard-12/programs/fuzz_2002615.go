package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint }
type GenS2 struct { F0 string; F1 int64; F2 int8 }
type GenS3 struct { F0 int }
func genFn0() uint64 {
    v4 := GenS1{F0: 86, F1: 60, F2: 820}
    if false {
        v5 := []int{-551}
        v6 := GenS1{F0: 665, F1: 127, F2: 0}
        var v7 uint = 236
        _ = v5
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int32 = 711
            v10 := GenS2{F0: "foo bar", F1: 0, F2: int8(fi8)}
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := []int{32767, 128}
        _ = v11
    }
    var v12 int64 = -1
    switch 2147483647 {
    case 0:
        var v13 uint8 = 127
        _ = v13
    case 402:
        v12 += v12
    case 42:
        for fi14 := 0; fi14 < 3; fi14++ {
            v12 -= (v12 >> ((4294967295) & 63))
            v15 := []int{}
            _ = fi14
            _ = v15
        }
    }
    _ = v12
    _ = v4
    return 1
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]uint32)
    var v3 uint8 = v1
    for fi16 := 0; fi16 < 19; fi16++ {
        if false {
            if true {
                fmt.Println(fi16, fi16, v3)
            }
        }
        v1 += v1
        _ = fi16
    }
    switch 4294967295 {
    case 1:
        var v17 int = 127
        _ = v17
    case 32767:
        v3 = (v3 << ((42) & 63))
    case 127:
        v2["foo bar"] = 1
    default:
        v18 := GenS3{F0: 9223372036854775807}
        _ = v18
    }
    v3 += v3
    v1 = 65
    _ = len(v2)
    v1 -= (v1 & (v3 - (v3 & v3)))
    var v19 int16 = -1
    v20 := GenS3{F0: -9223372036854775808}
    fmt.Println(v3, v1, v3)
    fmt.Println(v19)
    for fi21 := 0; fi21 < 18; fi21++ {
        v1 -= (v1 + v1)
        if (true || (-183 < 9223372036854775807)) {
            v22 := GenS3{F0: fi21}
            _ = v22
        } else {
            switch ("hello" + ("\n\t" + ("foo bar" + (("\n\t" + ("" + "\n\t")) + "hello")))) {
            case "a":
                v3 += v3
            case "foo bar":
                fmt.Println(v3)
            }
            v19 -= 0
            v3 = v3
        }
        _ = fi21
    }
    v23 := v2["hello"]
    v23 = 79
    if false {
        if true {
            if true {
                var v24 uint16 = 211
                _ = len(v2)
                v23 += (v23 | (^(v23)))
                _ = v24
            }
        }
    } else {
        var v25 uint = 354
        v3 = (^(v3))
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
