package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint8 }
func genFn0(p0 int8, p1 uint8, p2 string) int16 {
    v4 := GenS1{F0: p1, F1: p2, F2: p1}
    var v5 float32 = (3.14159 / 1.0)
    for fi6 := 0; fi6 < 7; fi6++ {
        p0 = (p0 / ((p0) | 1))
        var v7 uint8 = 128
        _ = fi6
        _ = v7
    }
    var v8 int16 = -292
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v8
    return (v8 | 1)
}

func main() {
    var v1 uint32 = 65535
    var v2 uint8 = 48
    v3 := make(map[uint32]int)
    if (218 == 2) {
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 int = (fi9 >> ((255) & 63))
            if (572 <= fi9) {
                v1 = 54
                for fi11 := 0; fi11 < 7; fi11++ {
                    var v12 uint8 = (v2 & v2)
                    _ = fi11
                    _ = v12
                }
                var v13 int8 = -43
                _ = v13
            } else {
                for fi14 := 0; fi14 < 19; fi14++ {
                    v2 += v2
                    var v15 int64 = -62
                    if false {
                        var v16 uint = 65535
                        _ = v16
                    }
                    _ = fi14
                    _ = v15
                }
            }
            switch ("\n\t" + ((("" + "\n\t") + (("" + ("a" + "hello")) + "hello")) + ("" + ""))) {
            case "foo bar":
                var v17 float32 = 0.0
                _ = v17
            }
            _ = fi9
            _ = v10
        }
    }
    var v18 uint64 = 42
    for fi19 := 0; fi19 < 15; fi19++ {
        if true {
            v2 += v2
            delete(v3, 2)
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 int = (-(fi19))
                var v22 int16 = -32768
                v23 := GenS1{F0: (v2 >> ((1) & 63)), F1: "foo bar", F2: v2}
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        }
        _ = fi19
    }
    delete(v3, v1)
    v24 := []int{}
    for fi25 := 0; fi25 < 8; fi25++ {
        if (true && true) {
            _ = len(v3)
            if true {
                var v26 int16 = 1000
                var v27 uint32 = v1
                v28 := GenS1{F0: (v2 << ((4294967295) & 63)), F1: ("" + ("\n\t" + ("" + "\n\t"))), F2: v2}
                _ = v26
                _ = v27
                _ = v28
            } else {
                v29 := GenS1{F0: v2, F1: "foo bar", F2: v2}
                _ = v29
            }
        }
        _ = fi25
    }
    if (672 < 1) {
        fmt.Println(v1, v1, v1)
    } else {
        var v30 int32 = 2147483647
        v3[v1] = 1
        var v31 uint = 255
        _ = v30
        _ = v31
    }
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    if (0 != -9223372036854775808) {
        var v32 string = ("a" + "")
        if (v2 < v2) {
            fmt.Println(v1, v32, v1)
        }
        v33 := GenS1{F0: v2, F1: "\n\t", F2: uint8(v1)}
        _ = v32
        _ = v33
    }
    var v34 uint8 = 255
    var v35 uint = 253
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    var v36 float64 = 1e10
    var v37 uint64 = (v18 >> ((uint(v18)) & 63))
    switch -9223372036854775808 {
    case 9223372036854775807:
        v1 = 728
    case 1:
        v38 := GenS1{F0: v2, F1: "foo bar", F2: v34}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
