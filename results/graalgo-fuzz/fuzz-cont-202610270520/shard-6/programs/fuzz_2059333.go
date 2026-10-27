package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 1000
    v3 := []int{int(v2), -110, -666}
    if false {
        var v4 float64 = 1e10
        var v5 int16 = (v2 ^ (v2 << ((18446744073709551615) & 63)))
        for fi6 := 0; fi6 < 7; fi6++ {
            fmt.Println(fi6, v5)
            _ = fi6
        }
        _ = v4
        _ = v5
    }
    if (false || (938 != 467)) {
        switch (("foo bar" + (((("\n\t" + "foo bar") + ("hello" + "")) + "\n\t") + ("" + "hello"))) + ("a" + (("" + "a") + "hello"))) {
        case "\n\t":
            if false {
                if len(v3) > 0 {
                    v3[0] = 1
                }
                v7 := GenS1{F0: -9223372036854775808, F1: false}
                var v8 uint = 4294967295
                _ = v7
                _ = v8
            }
        case "a":
            var v9 float64 = -2.25
            _ = v9
        }
    } else {
        if true {
            v2 -= (v2 >> ((255) & 63))
            var v10 uint32 = 4294967295
            fmt.Println(v10, v2)
            _ = v10
        }
    }
    var v11 uint16 = 1
    switch -2147483648 {
    case 4294967295:
        for fi12 := 0; fi12 < 1; fi12++ {
            v2 = -1
            v2 = (v2 - v2)
            if false {
                for fi13 := 0; fi13 < 11; fi13++ {
                    v2 = (v2 & v2)
                    if false {
                        v11 -= v11
                        fmt.Println(v11, v11, fi13)
                    }
                    _ = fi13
                }
                switch "hello" {
                case "hello":
                    v11 = 1
                case "foo bar":
                    v14 := GenS1{F0: 1, F1: (((4294967295 == 1) && (127 != -69)) && (false || (((!true) || false) && false)))}
                    _ = v14
                default:
                    v2 -= (-(v2))
                }
            }
            _ = fi12
        }
    case 388:
        fmt.Println(v11, v11)
    }
    var v15 bool = true
    v16 := GenS2{F0: 105, F1: (v11 << ((42) & 63)), F2: int64(v11)}
    var v17 int32 = 110
    v11 += v11
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 int8 = 86
        fmt.Println(v17)
        switch fi18 {
        case 2147483647:
            v20 := GenS3{F0: (fi18 != fi18), F1: ("foo bar" + "hello")}
            _ = v20
        case 0:
            var v21 string = ""
            _ = v21
        default:
            v15 = (!("" > "hello"))
        }
        _ = fi18
        _ = v19
    }
    var v22 bool = (126 < -1)
    v23 := GenS2{F0: 42, F1: (v11 % ((v11) | 1)), F2: -741}
    if false {
        if len(v3) > 0 {
            v3[0] = -9223372036854775808
        }
    }
    v24 := GenS4{F0: 0, F1: 18446744073709551615, F2: 1}
    v3 = append(v3, -9223372036854775808)
    v25 := GenS5{F0: 0}
    if true {
        fmt.Println(v17, v22)
        for fi26 := 0; fi26 < 11; fi26++ {
            v17 = v17
            v27 := GenS2{F0: 909, F1: v11, F2: -389}
            _ = fi26
            _ = v27
        }
    } else {
        v28 := GenS6{F0: 36, F1: 65535}
        fmt.Println(v11)
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := GenS4{F0: 42, F1: 42, F2: 100}
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        v32 := GenS5{F0: 231}
        _ = fi31
        _ = v32
    }
    fmt.Println(v22, v17, v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
