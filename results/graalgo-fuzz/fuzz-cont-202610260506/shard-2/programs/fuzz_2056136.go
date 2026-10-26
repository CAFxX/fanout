package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint64; F1 int }

func main() {
    v1 := []uint32{441, 4294967295, 392}
    v2 := make(map[int]uint32)
    v3 := []int{}
    var v4 uint8 = 2
    for fi5 := 0; fi5 < 19; fi5++ {
        v4 = (v4 % ((v4) | 1))
        _ = fi5
    }
    fmt.Println(v4)
    var v6 uint8 = v4
    v4 = v6
    if true {
        switch 257 {
        case -920:
            v4 += (v6 + v6)
        case 42:
            var v7 int8 = -1
            _ = v7
        }
        v4 = v4
    } else {
        var v8 uint16 = 255
        switch "" {
        case "foo bar":
            v9 := GenS1{F0: 65535, F1: 0}
            _ = v9
        }
        v2[4294967295] = 2147483647
        _ = v8
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        switch (-(fi10)) {
        case 127:
            v4 -= v6
        case 430:
            var v11 float64 = 1.5
            _ = v11
        case 2147483647:
            if false {
                if false {
                    for fi12 := 0; fi12 < 12; fi12++ {
                        _ = len(v3)
                        _ = fi12
                    }
                    for fi13 := 0; fi13 < 18; fi13++ {
                        var v14 int16 = 42
                        v15 := v2[fi10]
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
            }
        }
        var v16 uint32 = 4294967295
        _ = fi10
        _ = v16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 int = 4294967295
        _ = fi17
        _ = v18
    }
    var v19 bool = false
    switch 9223372036854775807 {
    case 0:
        v2[0] = 42
    case 878:
        v6 -= 1
    }
    v19 = (uint(v6) >= 856)
    for fi20 := 0; fi20 < 20; fi20++ {
        switch "" {
        case "a":
            v21 := GenS2{F0: 0, F1: (fi20 / (((fi20 ^ (-(fi20)))) | 1))}
            _ = v21
        default:
            var v22 uint8 = (v4 / ((v6) | 1))
            _ = v22
        }
        var v23 int32 = 0
        fmt.Println(v19, fi20)
        _ = fi20
        _ = v23
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        if true {
            v19 = true
            switch "foo bar" {
            case "":
                var v25 int8 = -34
                _ = v25
            default:
                if len(v3) > 0 {
                    v3[0] = fi24
                }
            }
            for fi26 := 0; fi26 < 10; fi26++ {
                v2[(fi26 | fi24)] = 4294967295
                v4 -= 76
                _ = fi26
            }
        } else {
            v27 := v2[(fi24 >> ((255) & 63))]
            for fi28 := 0; fi28 < 3; fi28++ {
                fmt.Println(fi24, fi28)
                _ = fi28
            }
            _ = v27
        }
        switch (("foo bar" + "") + ("hello" + (("hello" + "a") + "foo bar"))) {
        case "":
            for fi29 := 0; fi29 < 10; fi29++ {
                var v30 uint32 = 0
                _ = fi29
                _ = v30
            }
        case "hello":
            for fi31 := 0; fi31 < 11; fi31++ {
                v6 = 255
                _ = fi31
            }
        case "a":
            if v19 {
                v6 += (v6 * v4)
                var v32 int64 = 1
                var v33 int = (fi24 >> ((0) & 63))
                _ = v32
                _ = v33
            }
        }
        _ = fi24
    }
    v19 = false
    fmt.Println(v6, v19)
    fmt.Println(v4, v4, v4)
    fmt.Println(v6, v4)
    for fi34 := 0; fi34 < 10; fi34++ {
        for fi35 := 0; fi35 < 6; fi35++ {
            switch "a" {
            case "hello":
                v36 := GenS1{F0: 624, F1: 39}
                _ = v36
            }
            if v19 {
                v4 = 58
                switch 933 {
                case -1:
                    if v19 {
                        v4 = v6
                        v37 := GenS1{F0: 2, F1: 26}
                        _ = v37
                    } else {
                        _ = len(v1)
                    }
                }
            }
            _ = fi35
        }
        _ = fi34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
