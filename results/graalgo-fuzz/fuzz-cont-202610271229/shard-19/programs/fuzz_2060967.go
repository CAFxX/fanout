package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 bool; F2 int8 }

func main() {
    var v1 int64 = -850
    var v2 uint8 = 0
    var v3 int32 = 2147483647
    v4 := make(map[string]string)
    var v5 uint16 = 42
    v6 := []int{1}
    switch "hello" {
    case "":
        v4["foo bar"] = ""
    case "\n\t":
        switch ((("foo bar" + ("a" + "\n\t")) + "foo bar") + ("\n\t" + "hello")) {
        case "foo bar":
            v3 = 42
        case "hello":
            v5 += 1
        default:
            var v7 int32 = v3
            _ = v7
        }
    default:
        v5 -= 42
    }
    switch 0 {
    case -1:
        for fi8 := 0; fi8 < 5; fi8++ {
            v4[("foo bar" + "foo bar")] = "\n\t"
            _ = fi8
        }
    case 4294967295:
        v1 = (-(v1))
    default:
        switch -462 {
        case 42:
            if ((v5 / (((^(v5))) | 1)) >= (^(v5))) {
                v5 = 42
                var v9 int16 = 1
                _ = v9
            } else {
                var v10 uint16 = 65535
                _ = v10
            }
        case -1:
            for fi11 := 0; fi11 < 7; fi11++ {
                v12 := GenS1{F0: (v5 ^ v5)}
                _ = fi11
                _ = v12
            }
        default:
            v13 := GenS2{F0: 57, F1: ((v3 < (v3 + (v3 | -192))) || true), F2: 100}
            _ = v13
        }
    }
    var v14 uint64 = 18446744073709551615
    delete(v4, "foo bar")
    var v15 uint16 = (^(v5))
    delete(v4, "hello")
    v1 += (-(v1))
    var v16 int64 = (v1 ^ (v1 - 0))
    switch "foo bar" {
    case "foo bar":
        v6 = append(v6, -501)
    case "\n\t":
        var v17 string = "foo bar"
        _ = v17
    case "a":
        for fi18 := 0; fi18 < 7; fi18++ {
            fmt.Println(v15, v2)
            v1 = v1
            v19 := GenS1{F0: uint16(v14)}
            _ = fi18
            _ = v19
        }
    default:
        var v20 int32 = 2147483647
        _ = v20
    }
    if true {
        var v21 int16 = 42
        fmt.Println(v2, v21)
        _ = v21
    } else {
        switch int(v1) {
        case -2147483648:
            v16 -= v16
        case 9223372036854775807:
            delete(v4, ("" + "hello"))
        default:
            fmt.Println(v16, v16)
        }
    }
    if len(v6) > 0 {
        v6[0] = 2147483647
    }
    switch (("hello" + "foo bar") + ("\n\t" + "a")) {
    case "foo bar":
        switch int(v16) {
        case -9223372036854775808:
            for fi22 := 0; fi22 < 17; fi22++ {
                for fi23 := 0; fi23 < 13; fi23++ {
                    v24 := GenS2{F0: 0, F1: true, F2: 100}
                    if (((^(v15)) > (^(v5))) && true) {
                        _ = len(v4)
                        var v25 uint = 65535
                        _ = len(v4)
                        _ = v25
                    }
                    _ = fi23
                    _ = v24
                }
                _ = fi22
            }
        case 714:
            var v26 float64 = -2.25
            _ = v26
        case 32768:
            _ = len(v4)
        }
    default:
        v16 = v16
    }
    fmt.Println(v3, v16)
    if false {
        var v27 float64 = (1.5 + 0.0)
        _ = v27
    } else {
        v1 += v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
