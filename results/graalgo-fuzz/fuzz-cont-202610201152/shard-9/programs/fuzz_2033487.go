package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 bool) uint64 {
    var v5 bool = (p0 && (692 < 632))
    v6 := []int{}
    v7 := GenS1{F0: "foo bar"}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 715
}
func genFn1() uint {
    v8 := GenS1{F0: ((("a" + ("" + "foo bar")) + "hello") + "\n\t")}
    v9 := []int{}
    if (4294967295 != 4294967295) {
        var v10 uint64 = 42
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint64 = 2
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 uint = 65535
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v15 int8 = 1
    _ = v15
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint8 = 127
    var v2 int8 = -128
    v3 := []uint32{2147483648}
    var v4 int = 0
    v16 := GenS1{F0: ""}
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 int32 = -2147483648
        var v19 int64 = -9223372036854775808
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v4)
    v4 = (v4 << ((0) & 63))
    switch "hello" {
    case "foo bar":
        v20 := GenS1{F0: "foo bar"}
        _ = v20
    }
    switch ("hello" + "") {
    case "":
        fmt.Println(v2)
    case "foo bar":
        _ = len(v3)
    case "a":
        for fi21 := 0; fi21 < 3; fi21++ {
            v4 -= v4
            var v22 uint8 = (^(v1))
            _ = fi21
            _ = v22
        }
    default:
        var v23 uint32 = 41
        _ = v23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        for fi25 := 0; fi25 < 9; fi25++ {
            if true {
                var v26 uint = 909
                _ = v26
            }
            var v27 uint16 = 73
            v28 := GenS1{F0: "a"}
            v1 += v1
            _ = fi25
            _ = v27
            _ = v28
        }
        v2 = (v2 | (v2 - (-(v2))))
        var v29 int16 = int16(fi24)
        _ = fi24
        _ = v29
    }
    var v30 int64 = -9223372036854775808
    var v31 float64 = 0.0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
