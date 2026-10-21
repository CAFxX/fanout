package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint64 }
type GenS2 struct { F0 string; F1 int32; F2 int8 }
type GenS3 struct { F0 bool; F1 int; F2 int16 }
type GenS4 struct { F0 uint16; F1 int; F2 uint8 }

func main() {
    var v1 uint32 = 65535
    v2 := []string{}
    v3 := make(map[string]uint32)
    var v4 uint8 = 229
    if false {
        v5 := GenS1{F0: 706, F1: -32768, F2: 2}
        for fi6 := 0; fi6 < 13; fi6++ {
            fmt.Println(v4)
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 int8 = 100
                v1 = v1
                _ = fi7
                _ = v8
            }
            v9 := GenS1{F0: -1, F1: 403, F2: 4294967295}
            _ = fi6
            _ = v9
        }
        _ = v5
    } else {
        for fi10 := 0; fi10 < 4; fi10++ {
            _ = len(v2)
            var v11 uint = 65535
            _ = fi10
            _ = v11
        }
        var v12 int16 = -32768
        v13 := GenS1{F0: 106, F1: (v12 << ((169) & 63)), F2: 18446744073709551615}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 int16 = int16(v4)
        _ = fi14
        _ = v15
    }
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS2{F0: "foo bar", F1: 127, F2: 127}
            for fi18 := 0; fi18 < 7; fi18++ {
                var v19 uint32 = v1
                var v20 uint64 = 0
                var v21 string = "foo bar"
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            v2 = append(v2, "")
            _ = fi16
            _ = v17
        }
    case "\n\t":
        var v22 int = 4294967295
        _ = v22
    }
    v3["\n\t"] = (^(v1))
    v23 := GenS1{F0: 32768, F1: 1, F2: 627}
    fmt.Println(v4)
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 int32 = -1
        v26 := GenS3{F0: false, F1: (fi24 * int(v4)), F2: -78}
        if true {
            for fi27 := 0; fi27 < 13; fi27++ {
                var v28 uint64 = 1
                _ = fi27
                _ = v28
            }
            _ = len(v3)
            var v29 int16 = 1
            _ = v29
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    v4 -= 128
    for fi30 := 0; fi30 < 3; fi30++ {
        var v31 int64 = -863
        _ = fi30
        _ = v31
    }
    switch -351 {
    case 32767:
        var v32 int32 = 0
        _ = v32
    default:
        fmt.Println(v1)
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        if (309 >= 42) {
            v34 := GenS4{F0: uint16(v1), F1: (-(fi33)), F2: v4}
            _ = v34
        }
        _ = fi33
    }
    if (42 > -307) {
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
