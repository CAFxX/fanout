package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int; F1 int8 }

func main() {
    var v1 uint16 = 32767
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]string)
    var v4 uint8 = 0
    var v5 int = 32767
    var v6 int64 = -1
    v6 = (v6 - v6)
    v7 := []int{(-(v5)), (-(v5)), v5}
    v2 -= v2
    var v8 uint32 = (v2 >> ((2) & 63))
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            v4 = v4
            _ = fi10
        }
        _ = fi9
    }
    fmt.Println(v2)
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: -1}
        var v13 int = -400
        switch ("" + ("hello" + "a")) {
        case "\n\t":
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 uint64 = 4294967295
                _ = fi14
                _ = v15
            }
        default:
            v16 := GenS1{F0: -85}
            _ = v16
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    var v17 int64 = (v6 - v6)
    if (v4 <= v4) {
        v5 -= v5
        if ((^(v1)) == v1) {
            if true {
                var v18 uint16 = 65535
                v7 = append(v7, (-(v5)))
                _ = v18
            }
            var v19 int32 = -850
            switch (v5 + (v5 << ((323) & 63))) {
            case 617:
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 string = ("foo bar" + "hello")
                    var v22 uint16 = v1
                    var v23 uint32 = 65535
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
            case -1:
                var v24 int8 = -14
                _ = v24
            default:
                v25 := GenS2{F0: v19, F1: false, F2: v8}
                _ = v25
            }
            _ = v19
        } else {
            fmt.Println(v6, v17)
        }
        if (v5 <= v5) {
            if false {
                v8 += 407
                if (!false) {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        v2 = (v8 * (v8 / (((v8 ^ (v2 * (v8 + 2147483647)))) | 1)))
                        _ = fi26
                    }
                } else {
                    v17 -= v6
                }
                for fi27 := 0; fi27 < 19; fi27++ {
                    fmt.Println(v17)
                    _ = fi27
                }
            }
        }
    } else {
        _ = len(v7)
        var v28 int16 = -32768
        _ = v28
    }
    var v29 int64 = v6
    fmt.Println(v6, v2)
    if (uint64(v1) == 839) {
        var v30 uint64 = 0
        var v31 bool = (true || (3.14159 < -2.25))
        v29 = (v6 << ((1) & 63))
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        v2 = (v8 + v2)
        for fi33 := 0; fi33 < 18; fi33++ {
            v34 := GenS3{F0: (v4 >> ((1) & 63)), F1: 0}
            v7 = append(v7, v5)
            var v35 uint64 = 989
            _ = fi33
            _ = v34
            _ = v35
        }
        switch "foo bar" {
        case "":
            var v36 string = ((("a" + "foo bar") + ("" + "foo bar")) + "hello")
            _ = v36
        case "a":
            if (false && ((!true) && false)) {
                v29 = (v29 - int64(v2))
                v37 := GenS3{F0: 2, F1: 27}
                _ = v37
            } else {
                v29 -= v6
            }
        case "\n\t":
            fmt.Println(v4)
        default:
            var v38 uint8 = v4
            _ = v38
        }
        _ = fi32
    }
    var v39 uint64 = 1
    v40 := GenS4{F0: (v5 * (v5 / ((v5) | 1))), F1: 106}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
