package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 string; F2 int32 }
func genFn0(p0 uint32) uint32 {
    switch "a" {
    case "hello":
        var v4 int64 = 155
        _ = v4
    default:
        for fi5 := 0; fi5 < 20; fi5++ {
            v6 := GenS1{F0: 0, F1: 2}
            var v7 int = fi5
            var v8 int32 = 0
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    }
    v9 := GenS1{F0: 25, F1: 255}
    var v10 int8 = 127
    _ = p0
    _ = v10
    _ = v9
    return (p0 << ((2) & 63))
}

func main() {
    var v1 uint16 = 32767
    v2 := []int{-542, 32767, -300, 9223372036854775807}
    var v3 float64 = 1e10
    fmt.Println(v3)
    var v11 bool = false
    _ = len(v2)
    fmt.Println(v1, v3)
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS1{F0: 931, F1: 128}
            if ((v1 >> ((524) & 63)) == v1) {
                var v15 int16 = 1
                _ = v15
            } else {
                var v16 int16 = 32767
                _ = v16
            }
            v17 := GenS2{F0: -1, F1: 65535}
            _ = fi13
            _ = v14
            _ = v17
        }
        fmt.Println(v11, fi12, v1)
        switch (fi12 >> ((2) & 63)) {
        case 2147483647:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 127:
            fmt.Println(v3, v1)
        default:
            switch fi12 {
            case 42:
                var v18 string = "\n\t"
                _ = v18
            case -9223372036854775808:
                switch "\n\t" {
                case "a":
                    var v19 int = fi12
                    _ = v19
                case "":
                    var v20 uint16 = 2
                    _ = v20
                default:
                    v21 := GenS2{F0: 42, F1: 2}
                    _ = v21
                }
            }
        }
        _ = fi12
    }
    var v22 uint8 = 42
    v23 := GenS3{F0: (v1 / ((v1) | 1)), F1: 2147483647, F2: -622}
    switch -1 {
    case 42:
        fmt.Println(v1, v11, v3)
    }
    v24 := GenS1{F0: -309, F1: (^(v22))}
    var v25 int32 = 2147483647
    var v26 int = -280
    v27 := GenS4{F0: "foo bar"}
    if ((!(!(int8(v25) <= -1))) && (((2 != 1) && (539 <= 18446744073709551615)) && true)) {
        for fi28 := 0; fi28 < 10; fi28++ {
            v11 = false
            _ = fi28
        }
        v29 := GenS5{F0: 2147483647, F1: "\n\t", F2: 1}
        var v30 string = ("hello" + "hello")
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        if (!false) {
            var v32 uint = 4294967295
            _ = v32
        } else {
            var v33 float64 = (v3 / v3)
            for fi34 := 0; fi34 < 3; fi34++ {
                if v11 {
                    var v35 uint = 18446744073709551615
                    var v36 int16 = 32767
                    _ = v35
                    _ = v36
                }
                _ = fi34
            }
            _ = len(v2)
            _ = v33
        }
        var v37 float32 = (1.5 * 0.0)
        v26 -= -731
        _ = fi31
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
