package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
func genFn0(p0 int32) int32 {
    v5 := []int{32768}
    _ = len(v5)
    fmt.Println(p0)
    if ((-9223372036854775808 <= 1) && false) {
        p0 = p0
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
    } else {
        p0 += 1
        var v6 int8 = 0
        _ = len(v5)
        _ = v6
    }
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 uint = 255
    v2 := []string{("a" + (("" + "hello") + "")), ("\n\t" + ("\n\t" + "a")), (("" + "") + ""), ""}
    v3 := []int{0, 2147483647, 0, 128}
    v4 := []string{"\n\t"}
    var v7 uint16 = 255
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint32 = uint32(v7)
        var v10 bool = (100 <= 100)
        var v11 string = ""
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    if (!(9223372036854775807 != 1)) {
        var v12 int32 = 0
        for fi13 := 0; fi13 < 2; fi13++ {
            if true {
                for fi14 := 0; fi14 < 2; fi14++ {
                    var v15 int64 = -9223372036854775808
                    var v16 int8 = -11
                    var v17 int64 = v15
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
                v7 = 42
                v7 = (v7 * v7)
            } else {
                var v18 bool = true
                _ = v18
            }
            _ = fi13
        }
        _ = v12
    } else {
        var v19 float64 = -2.25
        v19 = (v19 - v19)
        v20 := GenS1{F0: 542}
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: true}
    v7 -= (v7 & (v7 % ((764) | 1)))
    switch ("hello" + "hello") {
    case "hello":
        v7 = 65535
    case "":
        switch int(v7) {
        case 9223372036854775807:
            if false {
                var v22 uint16 = (v7 * (v7 & v7))
                v2 = append(v2, "foo bar")
                _ = v22
            }
        case 4294967295:
            for fi23 := 0; fi23 < 5; fi23++ {
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
                var v24 string = "a"
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi23
                _ = v24
            }
        }
    case "foo bar":
        fmt.Println(v7, v7)
    default:
        if (42 <= 0) {
            v7 -= v7
            var v25 int16 = 1
            v1 -= 42
            _ = v25
        }
    }
    var v26 int16 = 1
    fmt.Println(v7, v1)
    v1 += (^(v1))
    var v27 int = -9223372036854775808
    fmt.Println(v1, v7)
    for fi28 := 0; fi28 < 20; fi28++ {
        var v29 int64 = -1
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if ("hello" != "hello") {
            var v30 uint32 = 1
            _ = v30
        }
        _ = fi28
        _ = v29
    }
    v1 = v1
    switch v27 {
    case -1:
        v7 = v7
    case 1:
        v31 := GenS2{F0: false}
        _ = v31
    case 209:
        v27 -= v27
    }
    var v32 int64 = -214
    var v33 bool = false
    var v34 uint64 = 53
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v7
}
