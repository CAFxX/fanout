package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint16 }

func main() {
    var v1 bool = (0 >= 32767)
    var v2 uint64 = 2
    v3 := make(map[uint32]bool)
    v3[198] = true
    v4 := v3[2147483648]
    for fi5 := 0; fi5 < 7; fi5++ {
        delete(v3, 0)
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := []int{32768, (fi5 ^ fi6), (fi5 | fi6)}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v1 = v1
    var v8 uint32 = 1
    fmt.Println(v4, v2, v1)
    var v9 int = -964
    for fi10 := 0; fi10 < 4; fi10++ {
        v1 = v1
        v4 = true
        _ = fi10
    }
    switch "hello" {
    case "":
        if v4 {
            var v11 int8 = -1
            v4 = (v1 && v4)
            v3[v8] = (-1 >= 42)
            _ = v11
        }
    }
    _ = len(v3)
    switch "" {
    case "a":
        switch 997 {
        case -59:
            v4 = v1
        case -862:
            v12 := GenS1{F0: 0, F1: "a", F2: 2}
            _ = v12
        case -767:
            v8 += (v8 % ((v8) | 1))
        }
    case "hello":
        v9 += int(v2)
    case "":
        v4 = ((^(v8)) <= (v8 << ((360) & 63)))
    default:
        var v13 uint32 = (^(v8))
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        switch "a" {
        case "":
            fmt.Println(fi14, v1)
        case "hello":
            if false {
                var v15 int8 = -128
                _ = v15
            } else {
                v3[(v8 ^ 0)] = (!true)
                fmt.Println(v8, v1)
            }
        case "\n\t":
            delete(v3, v8)
        }
        if (v4 || (619 != 1)) {
            var v16 uint64 = (^(v2))
            for fi17 := 0; fi17 < 14; fi17++ {
                if true {
                    v18 := GenS1{F0: (fi17 << ((65535) & 63)), F1: ((("" + "hello") + ((("hello" + "foo bar") + "hello") + ("foo bar" + "hello"))) + ""), F2: 65535}
                    _ = v18
                }
                v19 := []int{32768, (fi17 + 32767)}
                switch int(v8) {
                case 32767:
                    v4 = v4
                case 42:
                    v20 := GenS1{F0: (fi14 & (fi17 & (fi17 >> ((1) & 63)))), F1: ("" + ("" + "foo bar")), F2: 911}
                    _ = v20
                }
                _ = fi17
                _ = v19
            }
            var v21 uint16 = 185
            _ = v16
            _ = v21
        } else {
            v1 = v4
            v1 = ((v8 | v8) < (v8 * (v8 - v8)))
        }
        _ = fi14
    }
    v4 = v4
    if v1 {
        v22 := GenS1{F0: (v9 % (((v9 / ((v9) | 1))) | 1)), F1: ("" + "foo bar"), F2: 42}
        var v23 int32 = 2147483647
        _ = v22
        _ = v23
    }
    var v24 uint8 = 2
    var v25 bool = v1
    var v26 int8 = -95
    v27 := []int{v9}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
