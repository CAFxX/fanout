package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 bool }
func genFn0(p0 uint32, p1 int) int {
    p1 = (p1 << ((310) & 63))
    p1 = (p1 ^ (-(p1)))
    for fi4 := 0; fi4 < 19; fi4++ {
        p1 = (^(fi4))
        fmt.Println(fi4)
        var v5 uint8 = 255
        _ = fi4
        _ = v5
    }
    switch (p1 * p1) {
    case -614:
        var v6 int8 = 42
        _ = v6
    case 127:
        if ("a" == ("hello" + "foo bar")) {
            var v7 int32 = -1
            if false {
                switch (p1 & p1) {
                case 4294967295:
                    v8 := []int{int(p0)}
                    _ = v8
                default:
                    v9 := GenS1{F0: 2, F1: ("a" + ("\n\t" + (("" + ("foo bar" + "hello")) + "foo bar")))}
                    _ = v9
                }
            } else {
                if false {
                    switch (("" + "\n\t") + ("\n\t" + "foo bar")) {
                    case "\n\t":
                        v10 := []int{(-(p1))}
                        _ = v10
                    }
                    var v11 uint16 = 32767
                    _ = v11
                }
                p0 += p0
                var v12 uint16 = 209
                _ = v12
            }
            var v13 uint32 = p0
            _ = v13
            _ = v7
        } else {
            if false {
                for fi14 := 0; fi14 < 20; fi14++ {
                    var v15 int = -1
                    _ = fi14
                    _ = v15
                }
            }
        }
    case -9223372036854775808:
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS1{F0: 97, F1: "hello"}
            fmt.Println(p1)
            _ = fi16
            _ = v17
        }
    }
    if true {
        v18 := GenS1{F0: 65, F1: ("a" + "a")}
        _ = v18
    } else {
        var v19 uint16 = 65535
        _ = v19
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 int8 = 100
    var v2 string = "foo bar"
    var v3 uint64 = 1
    v3 = 2
    v20 := []int{779, 32767, 1}
    v1 = (v1 - 1)
    v3 += v3
    for fi21 := 0; fi21 < 7; fi21++ {
        _ = len(v20)
        var v22 int32 = -1
        switch "" {
        case "\n\t":
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
        case "a":
            var v23 uint32 = 2147483647
            _ = v23
        case "hello":
            v24 := GenS2{F0: (42 < 0)}
            _ = v24
        default:
            var v25 uint64 = 2
            _ = v25
        }
        _ = fi21
        _ = v22
    }
    switch "a" {
    case "a":
        fmt.Println(v2)
    case "":
        var v26 uint8 = 231
        _ = v26
    case "hello":
        fmt.Println(v1, v3, v1)
    default:
        var v27 bool = true
        _ = v27
    }
    _ = len(v20)
    if (4294967295 <= 432) {
        v2 = "\n\t"
        v1 = v1
    } else {
        if true {
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 int8 = -128
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 19; fi30++ {
                if len(v20) > 0 {
                    v20[0] = fi30
                }
                _ = fi30
            }
            fmt.Println(v1, v1)
        } else {
            v2 = v2
        }
        var v31 uint8 = 1
        _ = v31
    }
    var v32 string = v2
    v2 = (v32 + "a")
    for fi33 := 0; fi33 < 11; fi33++ {
        var v34 uint8 = 29
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
}
