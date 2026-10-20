package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint32 }
func genFn0() int64 {
    var v4 uint64 = 42
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: true}
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 2
    switch "" {
    case "":
        if (true && true) {
            v4 = 2
            v8 := GenS2{F0: 2, F1: 42, F2: 707}
            var v9 int = -9223372036854775808
            _ = v8
            _ = v9
        } else {
            if true {
                for fi10 := 0; fi10 < 13; fi10++ {
                    if true {
                        var v11 uint8 = 128
                        v12 := GenS1{F0: false}
                        var v13 uint = 577
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    v7 -= (^(v4))
                    switch "foo bar" {
                    case "\n\t":
                        var v14 int32 = -592
                        _ = v14
                    default:
                        var v15 uint16 = 2
                        _ = v15
                    }
                    _ = fi10
                }
                var v16 uint8 = 74
                _ = v16
            } else {
                fmt.Println(v7)
                v17 := []int{}
                _ = v17
            }
            v4 += 1
        }
    case "\n\t":
        if false {
            v4 += v7
            var v18 float32 = (1e10 + ((((1.5 / 1.5) - ((-2.25 + 0.1) / (1e10 / 1e10))) + -2.25) * 0.0))
            _ = v18
        } else {
            v19 := GenS2{F0: 412, F1: 34, F2: 4294967295}
            _ = v19
        }
    }
    _ = v4
    _ = v7
    return -84
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int16 = -32768
    var v3 uint8 = uint8(v1)
    v20 := []int{-9223372036854775808, 93, -2147483648}
    v21 := GenS3{F0: 978, F1: 65535, F2: (v1 / (((v1 ^ (^(v1)))) | 1))}
    var v22 uint8 = v3
    var v23 float64 = 0.1
    fmt.Println(v3, v22)
    for fi24 := 0; fi24 < 16; fi24++ {
        v20 = append(v20, (fi24 - (-(fi24))))
        var v25 int16 = v2
        var v26 uint64 = 0
        _ = fi24
        _ = v25
        _ = v26
    }
    if false {
        switch -1 {
        case 32768:
            v2 += v2
        }
    }
    if ((0.1 < -2.25) || true) {
        var v27 uint8 = (^(v22))
        v2 += (v2 - (v2 >> ((442) & 63)))
        _ = v27
    } else {
        fmt.Println(v23, v23)
        if ((!false) && ((false && true) && (!(float32((0.1 * (v23 + (v23 + v23)))) <= 0.1)))) {
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 int16 = v2
                var v30 int64 = -805
                _ = fi28
                _ = v29
                _ = v30
            }
            var v31 uint32 = v1
            for fi32 := 0; fi32 < 19; fi32++ {
                if len(v20) > 0 {
                    v20[0] = (fi32 << ((4294967295) & 63))
                }
                v20 = append(v20, fi32)
                _ = fi32
            }
            _ = v31
        } else {
            var v33 int8 = 100
            _ = v33
        }
        var v34 float64 = (v23 + 0.0)
        _ = v34
    }
    switch "hello" {
    case "foo bar":
        if (("" + "\n\t") > ((("foo bar" + (("a" + (("hello" + "\n\t") + ("hello" + "foo bar"))) + "hello")) + ("a" + "foo bar")) + (("" + ("a" + ("a" + ("\n\t" + "a")))) + "a"))) {
            _ = len(v20)
        } else {
            if len(v20) > 0 {
                v20[0] = 32767
            }
            var v35 uint16 = 2
            var v36 uint = 0
            _ = v35
            _ = v36
        }
    case "hello":
        v37 := GenS1{F0: (!(32767 >= 32767))}
        _ = v37
    }
    fmt.Println(v3)
    for fi38 := 0; fi38 < 6; fi38++ {
        fmt.Println(v1, v23)
        for fi39 := 0; fi39 < 5; fi39++ {
            _ = len(v20)
            fmt.Println(v22, v23)
            v23 += v23
            _ = fi39
        }
        var v40 float64 = 3.14159
        _ = fi38
        _ = v40
    }
    v1 = (v1 >> ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
