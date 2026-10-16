package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint; F1 uint8; F2 string }
func genFn0(p0 int8) int32 {
    v6 := GenS1{F0: -1, F1: -132}
    if (!(((("hello" + "foo bar") + "foo bar") != "a") || false)) {
        fmt.Println(p0)
        p0 = p0
        if false {
            var v7 int64 = 0
            _ = v7
        }
    }
    var v8 uint32 = 1
    var v9 int32 = -2147483648
    switch ("hello" + "foo bar") {
    case "foo bar":
        v10 := []int{int(v8)}
        _ = v10
    case "a":
        var v11 bool = (!("hello" != "a"))
        _ = v11
    default:
        v12 := []int{}
        _ = v12
    }
    _ = p0
    _ = v6
    _ = v8
    _ = v9
    return (v9 & (v9 >> ((207) & 63)))
}
func genFn1() uint64 {
    var v13 bool = false
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS1{F0: -2147483648, F1: fi14}
        v16 := GenS1{F0: genFn0(127), F1: 2147483647}
        fmt.Println(v13, fi14)
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS1{F0: -1, F1: (fi17 | (fi17 * fi18))}
            fmt.Println(fi18, v13)
            switch -567 {
            case -843:
                fmt.Println(fi18, fi18, v13)
            }
            _ = fi18
            _ = v19
        }
        var v20 int64 = -9223372036854775808
        _ = fi17
        _ = v20
    }
    _ = v13
    return 2
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 199
    v3 := make(map[string]uint32)
    var v4 uint64 = v2
    var v5 float32 = float32(0.0)
    var v21 uint64 = (v2 << ((953) & 63))
    if true {
        switch 357 {
        case 128:
            v2 -= v4
        case 711:
            switch -9223372036854775808 {
            case 251:
                _ = len(v1)
            case 479:
                var v22 uint16 = 65535
                _ = v22
            default:
                for fi23 := 0; fi23 < 2; fi23++ {
                    v24 := v1["hello"]
                    v2 = 4294967295
                    v25 := GenS1{F0: genFn0(-100), F1: v24}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        }
    }
    if false {
        delete(v3, "hello")
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        v4 = 18446744073709551615
        _ = fi26
    }
    var v27 int = -157
    for fi28 := 0; fi28 < 3; fi28++ {
        if (!(!true)) {
            delete(v3, "foo bar")
            var v29 string = ""
            _ = v29
        }
        _ = fi28
    }
    var v30 bool = false
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 float32 = 1.5
        switch v27 {
        case -9223372036854775808:
            if v30 {
                v5 = v5
                fmt.Println(v21)
            } else {
                v27 = (-(fi31))
            }
        case 128:
            v33 := GenS2{F0: 1, F1: 2, F2: "foo bar"}
            _ = v33
        }
        _ = fi31
        _ = v32
    }
    v4 -= 128
    var v34 uint = 864
    v35 := GenS2{F0: v34, F1: uint8(v21), F2: (("\n\t" + "a") + ("" + ("\n\t" + "")))}
    var v36 int64 = 1
    v21 = (v21 * (v4 & v2))
    delete(v3, ("\n\t" + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
