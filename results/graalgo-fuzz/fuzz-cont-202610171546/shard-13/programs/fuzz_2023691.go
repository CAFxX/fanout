package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 string; F1 uint64 }
func genFn0(p0 int64) int64 {
    p0 -= p0
    p0 -= (p0 / ((p0) | 1))
    fmt.Println(p0)
    _ = p0
    return 1
}
func genFn1(p0 uint16, p1 uint64, p2 uint32) int32 {
    p2 += (^(p2))
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 int32 = -994
        var v7 int32 = (^(v6))
        var v8 uint8 = 1
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 float64 = (-2.25 - float64(1e10))
    v10 := []int{1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -2147483648
}

func main() {
    v1 := []string{"hello", (("hello" + "a") + ""), (("hello" + "\n\t") + "foo bar")}
    var v2 uint32 = 2
    var v3 bool = true
    v4 := make(map[uint32]int)
    if v3 {
        var v11 string = ("\n\t" + ("" + ("hello" + "hello")))
        switch v11 {
        case "a":
            v12 := GenS1{F0: (v2 ^ v2), F1: v2}
            _ = v12
        case "foo bar":
            v11 = ("a" + ((("hello" + (v11 + ("hello" + v11))) + ((("a" + v11) + v11) + v11)) + ""))
        case "hello":
            fmt.Println(v11, v11)
        default:
            var v13 int64 = 0
            _ = v13
        }
        _ = v11
    }
    v14 := GenS2{F0: "", F1: 2}
    var v15 uint16 = 2
    v15 = 65535
    if (v3 && false) {
        if len(v1) > 0 {
            v1[0] = "a"
        }
    }
    var v16 int32 = 2147483647
    var v17 uint64 = 1
    for fi18 := 0; fi18 < 5; fi18++ {
        v17 -= v17
        _ = fi18
    }
    if (!(false && (!(v3 || false)))) {
        fmt.Println(v15, v16)
    } else {
        v3 = ((9223372036854775807 < genFn0(202)) || v3)
    }
    v15 = v15
    if v3 {
        _ = len(v4)
        v3 = false
        var v19 uint64 = v17
        _ = v19
    } else {
        delete(v4, v2)
    }
    v4[1] = 128
    if false {
        if v3 {
            fmt.Println(v16, v3, v3)
        } else {
            delete(v4, v2)
            v20 := GenS2{F0: "", F1: v17}
            _ = v20
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            for fi22 := 0; fi22 < 14; fi22++ {
                fmt.Println(v16)
                _ = fi22
            }
            _ = fi21
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            v16 = genFn1(v15, (v17 % ((18446744073709551615) | 1)), v2)
            fmt.Println(v3)
            v17 -= 907
            _ = fi23
        }
    }
    switch 0 {
    case 128:
        var v24 uint64 = 445
        _ = v24
    default:
        v25 := GenS1{F0: (^(v2)), F1: (v2 + (v2 + (v2 >> ((65535) & 63))))}
        _ = v25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 uint64 = 2
        v27 = v27
        _ = fi26
        _ = v27
    }
    if ((-34 != -1) && v3) {
        if (127 >= 127) {
            var v28 bool = false
            _ = v28
        } else {
            if ((v17 / ((v17) | 1)) > v17) {
                var v29 uint8 = 42
                _ = v29
            } else {
                var v30 uint32 = (v2 * 1)
                var v31 uint16 = 32767
                _ = v30
                _ = v31
            }
        }
        v32 := GenS2{F0: "", F1: v17}
        _ = v32
    } else {
        for fi33 := 0; fi33 < 6; fi33++ {
            v15 -= v15
            _ = fi33
        }
    }
    v34 := GenS2{F0: "hello", F1: v17}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
