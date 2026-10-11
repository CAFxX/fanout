package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 string = ((("\n\t" + ((("foo bar" + ("\n\t" + ("a" + "\n\t"))) + "a") + ("hello" + "foo bar"))) + (("hello" + ("foo bar" + "")) + "")) + ("hello" + "a"))
    v4 := []uint32{1, 860, 173}
    var v5 int16 = -1
    var v6 uint8 = 1
    v7 := GenS1{F0: 1}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int = fi8
        if (-32768 < 1) {
            v10 := GenS1{F0: 522}
            _ = v10
        } else {
            var v11 string = v3
            for fi12 := 0; fi12 < 6; fi12++ {
                fmt.Println(v9, fi12)
                var v13 float64 = 1e10
                v4 = append(v4, 2)
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        switch v3 {
        case "hello":
            fmt.Println(v6, v9)
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(v6, v6, v3)
    v6 -= (v6 - 42)
    if false {
        if true {
            v14 := GenS1{F0: 18446744073709551615}
            _ = v14
        } else {
            var v15 int = 200
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 uint32 = 625
                _ = fi16
                _ = v17
            }
            _ = v15
        }
        var v18 string = (("" + v3) + v3)
        if false {
            for fi19 := 0; fi19 < 7; fi19++ {
                v20 := GenS2{F0: -2147483648}
                _ = fi19
                _ = v20
            }
        }
        _ = v18
    }
    fmt.Println(v3, v3, v3)
    v21 := GenS1{F0: 1}
    var v22 uint32 = 0
    v3 = (v3 + "foo bar")
    v3 = ("hello" + ((v3 + v3) + v3))
    switch v3 {
    case "foo bar":
        v5 = v5
    default:
        v3 = ((v3 + (v3 + "")) + v3)
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        if false {
            v24 := v2[v22]
            _ = v24
        } else {
            fmt.Println(v22, v6)
            v5 = 42
        }
        _ = fi23
    }
    if (-1 < 27) {
        for fi25 := 0; fi25 < 11; fi25++ {
            var v26 uint8 = (v6 & (v6 ^ (v6 % ((145) | 1))))
            _ = fi25
            _ = v26
        }
        v27 := GenS2{F0: -2147483648}
        _ = v27
    }
    v5 -= (v5 - 1)
    v6 = v6
    var v28 int64 = -902
    if true {
        var v29 uint32 = 712
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
