package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 string }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int8) int32 {
    var v6 bool = ("a" >= ("" + "hello"))
    switch ("\n\t" + ("a" + "foo bar")) {
    case "foo bar":
        var v7 float64 = 3.14159
        _ = v7
    case "":
        if (199 > 1) {
            if (!true) {
                p0 = (p0 << ((42) & 63))
            }
            v8 := []int{4294967295}
            _ = v8
        } else {
            if v6 {
                var v9 string = ""
                v10 := GenS1{F0: "foo bar", F1: 22, F2: v9}
                _ = v10
                _ = v9
            }
            if true {
                if (3.14159 == 1e10) {
                    if false {
                        v11 := []int{}
                        v12 := GenS1{F0: "a", F1: 0, F2: "hello"}
                        _ = v11
                        _ = v12
                    }
                    switch ("foo bar" + ("" + "")) {
                    case "a":
                        v6 = true
                    case "foo bar":
                        var v13 float32 = float32(0.1)
                        _ = v13
                    }
                    v14 := GenS2{F0: 2, F1: 18446744073709551615}
                    _ = v14
                } else {
                    if ((p0 ^ (p0 << ((0) & 63))) == 100) {
                        switch 132 {
                        case 128:
                            var v15 float32 = 0.0
                            _ = v15
                        }
                    } else {
                        p0 = (p0 | p0)
                        switch -9223372036854775808 {
                        case 458:
                            fmt.Println(v6, p0)
                        }
                    }
                }
                var v16 float64 = 3.14159
                _ = v16
            } else {
                var v17 uint64 = 1
                for fi18 := 0; fi18 < 11; fi18++ {
                    v19 := []int{fi18, (fi18 >> ((4294967295) & 63)), -299, -285}
                    var v20 int8 = p0
                    var v21 string = "hello"
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = v17
            }
        }
    }
    p0 = (p0 ^ p0)
    fmt.Println(v6)
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 uint8 = 255
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = v6
    return -1
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]int)
    var v3 float64 = 3.14159
    v4 := []string{"", "\n\t"}
    var v5 int8 = -71
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v24 float32 = ((0.1 + v1) - v1)
    v5 = 0
    switch 32768 {
    case 536:
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 uint32 = 140
            _ = fi25
            _ = v26
        }
    case 9223372036854775807:
        switch 0 {
        case -1:
            v24 -= (0.0 - 0.1)
        default:
            v3 = float64(-2.25)
        }
    }
    var v27 uint16 = 727
    _ = len(v2)
    if true {
        fmt.Println(v1)
        var v28 float32 = (v24 / v24)
        for fi29 := 0; fi29 < 19; fi29++ {
            var v30 int16 = 0
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    if (((-(v5)) >= v5) || true) {
        v24 = (v24 * v1)
        v31 := GenS1{F0: "hello", F1: v27, F2: ("foo bar" + "foo bar")}
        var v32 int = -9223372036854775808
        _ = v31
        _ = v32
    }
    var v33 uint64 = 149
    v3 += float64(v24)
    _ = len(v4)
    for fi34 := 0; fi34 < 2; fi34++ {
        switch "" {
        case "hello":
            var v35 int16 = -732
            _ = v35
        }
        _ = fi34
    }
    v33 -= (^(v33))
    for fi36 := 0; fi36 < 2; fi36++ {
        var v37 string = ("" + "")
        v38 := GenS1{F0: "hello", F1: uint16(v33), F2: "hello"}
        v33 += (v33 & (v33 << ((4294967295) & 63)))
        _ = fi36
        _ = v37
        _ = v38
    }
    switch 4294967295 {
    case -792:
        fmt.Println(v5, v1)
    case 4294967295:
        var v39 int8 = v5
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
