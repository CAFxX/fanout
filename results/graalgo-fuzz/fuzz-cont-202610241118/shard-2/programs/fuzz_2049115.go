package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 bool; F2 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 1
    var v3 uint64 = v2
    var v4 uint64 = (v2 ^ v2)
    for fi5 := 0; fi5 < 15; fi5++ {
        v3 = (v4 * (v4 % (((^(v4))) | 1)))
        var v6 uint = 21
        var v7 uint = (v6 & 387)
        _ = fi5
        _ = v6
        _ = v7
    }
    v1[(("foo bar" + "") + "hello")] = (!false)
    var v8 int8 = 100
    v9 := v1[("foo bar" + "\n\t")]
    v3 = 867
    v10 := GenS1{F0: 2}
    v11 := GenS2{F0: 1}
    if v9 {
        var v12 int16 = 1
        v4 -= 801
        if v9 {
            var v13 uint8 = 1
            v14 := GenS1{F0: 65535}
            if false {
                fmt.Println(v8)
            }
            _ = v13
            _ = v14
        } else {
            v15 := GenS2{F0: 65535}
            v12 = v12
            _ = v15
        }
        _ = v12
    }
    if (v9 && false) {
        var v16 uint32 = 204
        switch -9223372036854775808 {
        case -9223372036854775808:
            var v17 uint = 229
            _ = v17
        case 2147483647:
            var v18 uint = 42
            _ = v18
        }
        _ = v16
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        v9 = v9
        _ = fi19
    }
    if v9 {
        switch -2147483648 {
        case 484:
            v20 := GenS1{F0: 4294967295}
            _ = v20
        case 0:
            if (v8 != 0) {
                var v21 string = "\n\t"
                var v22 float64 = (-2.25 + 1e10)
                _ = v21
                _ = v22
            } else {
                var v23 uint32 = 274
                _ = v23
            }
        }
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        if true {
            fmt.Println(v4, v9)
        } else {
            fmt.Println(v3, v9)
            v25 := GenS1{F0: 239}
            _ = v25
        }
        if false {
            switch "\n\t" {
            case "foo bar":
                v26 := GenS3{F0: (^(fi24)), F1: 32767, F2: -1}
                _ = v26
            }
            v9 = (0 < int64(v8))
            v3 -= (v2 >> ((445) & 63))
        } else {
            for fi27 := 0; fi27 < 10; fi27++ {
                var v28 uint8 = 2
                var v29 uint16 = 0
                v30 := GenS4{F0: 299, F1: 882}
                v3 = (v2 ^ (v3 - 1))
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        }
        v31 := GenS5{F0: 9223372036854775807}
        _ = fi24
        _ = v31
    }
    v32 := GenS6{F0: 1, F1: (!(v9 && true)), F2: 0}
    v33 := GenS3{F0: 911, F1: 255, F2: v8}
    switch "hello" {
    case "foo bar":
        switch "hello" {
        case "foo bar":
            v1["hello"] = ((253 < 0) || (int16(v4) >= 1000))
        case "a":
            v34 := v1[""]
            _ = v34
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v8
    _ = v9
}
