package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint64, p2 uint8) int16 {
    fmt.Println(p1, p0, p2)
    var v5 int16 = 1000
    var v6 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 int32, p1 uint16, p2 uint16) string {
    p1 = (^(p1))
    switch 4294967295 {
    case 1:
        if true {
            switch 2147483647 {
            case -2147483648:
                v7 := []int{}
                _ = v7
            case 0:
                for fi8 := 0; fi8 < 13; fi8++ {
                    if true {
                        var v9 bool = true
                        var v10 int8 = -1
                        _ = v10
                        _ = v9
                    }
                    p0 = 1
                    v11 := GenS1{F0: fi8, F1: 0}
                    _ = fi8
                    _ = v11
                }
            }
        } else {
            var v12 int64 = 9223372036854775807
            p1 = 1
            for fi13 := 0; fi13 < 7; fi13++ {
                v14 := GenS2{F0: 65535, F1: ("foo bar" + "")}
                _ = fi13
                _ = v14
            }
            _ = v12
        }
    case 42:
        if true {
            fmt.Println(p2, p2, p0)
            v15 := []int{752, 9223372036854775807}
            _ = len(v15)
            _ = v15
        }
    case -9223372036854775808:
        var v16 int64 = 0
        _ = v16
    }
    var v17 float64 = (-2.25 * (1.5 + 0.1))
    p1 -= p2
    if (1 > 2147483647) {
        var v18 uint8 = 128
        if (-850 <= 1000) {
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 string = "a"
                if (!false) {
                    v17 -= (((0.1 / (float64((0.1 - (1.5 / 1.0))) * (-2.25 - 1.5))) / (v17 * float64(1e10))) + v17)
                    switch 42 {
                    case 170:
                        for fi21 := 0; fi21 < 16; fi21++ {
                            if false {
                                var v22 string = "foo bar"
                                var v23 uint = 0
                                _ = v22
                                _ = v23
                            }
                            fmt.Println(p0)
                            var v24 int32 = p0
                            _ = fi21
                            _ = v24
                        }
                    }
                }
                _ = fi19
                _ = v20
            }
            p0 -= p0
            fmt.Println(p1, p1)
        }
        _ = v18
    } else {
        v25 := []int{-9223372036854775808}
        if true {
            v26 := GenS2{F0: 310, F1: "a"}
            v27 := GenS3{F0: -1, F1: ("" <= "hello")}
            var v28 uint = 288
            _ = v26
            _ = v27
            _ = v28
        }
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return "\n\t"
}

func main() {
    var v1 float32 = 0.0
    var v2 uint16 = 32767
    var v3 float32 = v1
    v4 := []uint32{0, 1, 2147483647}
    v1 = v1
    if false {
        v1 = (((1e10 + 1.5) / -2.25) - (v1 / v3))
    } else {
        v29 := GenS2{F0: 835, F1: ("hello" + "\n\t")}
        fmt.Println(v2)
        fmt.Println(v2)
        _ = v29
    }
    fmt.Println(v1)
    v3 += v3
    var v30 bool = (true && (true && false))
    v31 := GenS3{F0: 0, F1: true}
    _ = len(v4)
    if true {
        switch "\n\t" {
        case "a":
            v32 := GenS4{F0: 100}
            _ = v32
        case "foo bar":
            v3 = (v3 + v3)
        }
        switch "a" {
        case "foo bar":
            var v33 uint = 18446744073709551615
            _ = v33
        }
    } else {
        v2 = 0
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        v3 -= v3
        for fi35 := 0; fi35 < 18; fi35++ {
            fmt.Println(v30, v2, v30)
            _ = fi35
        }
        _ = fi34
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        for fi36 := 0; fi36 < 6; fi36++ {
            var v37 int8 = -1
            v3 += v1
            _ = fi36
            _ = v37
        }
    default:
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
