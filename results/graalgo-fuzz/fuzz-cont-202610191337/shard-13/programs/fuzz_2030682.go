package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int32; F2 string }
type GenS3 struct { F0 uint16; F1 int; F2 uint8 }
func genFn0(p0 int32, p1 uint32, p2 int) bool {
    p0 = p0
    fmt.Println(p2, p2)
    var v4 uint = 255
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 int = (p2 + fi5)
        if (2 != 18446744073709551615) {
            fmt.Println(p0)
            var v7 uint = v4
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint64 = 94
        fmt.Println(v4, p1, p2)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return false
}
func genFn1() int {
    var v10 int32 = -832
    var v11 int8 = 27
    var v12 uint16 = 2
    fmt.Println(v10)
    _ = v10
    _ = v11
    _ = v12
    return 2147483647
}

func main() {
    var v1 uint32 = 0
    var v2 int32 = 42
    v3 := make(map[uint32]uint32)
    v13 := GenS1{F0: (!false), F1: ((("\n\t" + ((("hello" + "a") + "hello") + "\n\t")) + "\n\t") + ""), F2: v1}
    var v14 uint32 = v1
    v15 := v3[v14]
    fmt.Println(v15, v15, v15)
    switch 32767 {
    case 42:
        var v16 int = -2147483648
        _ = v16
    case 1:
        if (2147483647 == 22) {
            v2 = 42
        } else {
            v2 += v2
            var v17 int64 = -248
            _ = v17
        }
    default:
        if false {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 uint64 = 0
                v20 := v3[0]
                fmt.Println(v2)
                _ = fi18
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 7; fi21++ {
                v14 -= 838
                fmt.Println(v2)
                var v22 string = "a"
                _ = fi21
                _ = v22
            }
            switch ((("" + ("hello" + "\n\t")) + "foo bar") + "hello") {
            case "hello":
                var v23 int16 = 412
                _ = v23
            default:
                v14 = (v14 + (v15 | v15))
            }
        } else {
            v24 := GenS2{F0: 946, F1: v2, F2: ("foo bar" + "foo bar")}
            v25 := v3[2147483647]
            _ = v24
            _ = v25
        }
    }
    fmt.Println(v2, v14, v14)
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 uint16 = 32767
        v27 = v27
        _ = fi26
        _ = v27
    }
    if (!(!(false || (!false)))) {
        fmt.Println(v14, v1, v1)
        v28 := GenS3{F0: 1, F1: -9223372036854775808, F2: 1}
        _ = v28
    }
    var v29 uint8 = 42
    fmt.Println(v1, v1)
    v1 -= (v14 + (v15 * v1))
    v30 := []int{42, 32768}
    if false {
        var v31 int32 = 981
        _ = v31
    }
    v32 := GenS3{F0: 22, F1: 1, F2: 1}
    var v33 int16 = int16(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
}
