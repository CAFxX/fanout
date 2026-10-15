package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 string; F2 bool }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 uint64; F1 int8 }
func genFn0(p0 int, p1 uint64) bool {
    switch ((("" + "") + "") + ("" + "\n\t")) {
    case "a":
        p0 -= (p0 / ((0) | 1))
    }
    var v4 uint = 0
    _ = p0
    _ = p1
    _ = v4
    return false
}
func genFn1(p0 int64, p1 uint, p2 int32) uint8 {
    switch ("a" + "hello") {
    case "a":
        switch "a" {
        case "foo bar":
            var v5 int = 128
            _ = v5
        case "\n\t":
            v6 := GenS1{F0: 18446744073709551615}
            _ = v6
        default:
            p0 += p0
        }
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS2{F0: 42, F1: "hello", F2: (!((-(fi8)) < fi8))}
            p1 -= (p1 >> ((p1) & 63))
            _ = fi8
            _ = v9
        }
        if false {
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 uint8 = 128
                _ = fi10
                _ = v11
            }
        }
        switch 127 {
        case -340:
            fmt.Println(p2)
        }
        _ = fi7
    }
    var v12 uint = (p1 ^ 0)
    for fi13 := 0; fi13 < 17; fi13++ {
        v12 = (p1 % (((^(p1))) | 1))
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 212
}

func main() {
    v1 := []int{-547}
    v2 := []int{945, 107, 606}
    v3 := make(map[uint32]int)
    var v14 float32 = 1e10
    switch "foo bar" {
    case "":
        v14 = (0.0 - v14)
    case "foo bar":
        fmt.Println(v14)
    }
    if (true && false) {
        v15 := GenS1{F0: 1}
        switch 531 {
        case 42:
            switch "hello" {
            case "foo bar":
                v16 := GenS3{F0: "a", F1: -1}
                _ = v16
            case "\n\t":
                switch "a" {
                case "hello":
                    switch "\n\t" {
                    case "\n\t":
                        _ = len(v3)
                    case "a":
                        v17 := GenS4{F0: 42}
                        _ = v17
                    }
                case "":
                    if ((42 < -32768) || false) {
                        _ = len(v1)
                        v3[333] = 128
                    }
                case "\n\t":
                    switch -9223372036854775808 {
                    case 32768:
                        for fi18 := 0; fi18 < 4; fi18++ {
                            var v19 uint16 = 32767
                            _ = fi18
                            _ = v19
                        }
                    }
                }
            }
        case -733:
            for fi20 := 0; fi20 < 2; fi20++ {
                switch "hello" {
                case "\n\t":
                    v14 -= v14
                case "a":
                    delete(v3, 1)
                case "":
                    var v21 uint16 = 2
                    _ = v21
                }
                var v22 int32 = -2147483648
                for fi23 := 0; fi23 < 6; fi23++ {
                    if true {
                        v22 -= -1
                        fmt.Println(fi23, fi20, v22)
                        if true {
                            var v24 int64 = -538
                            _ = v24
                        } else {
                            fmt.Println(fi23)
                            v14 += 3.14159
                        }
                    }
                    _ = fi23
                }
                _ = fi20
                _ = v22
            }
        }
        fmt.Println(v14)
        _ = v15
    } else {
        var v25 uint64 = 0
        var v26 uint32 = 2147483648
        for fi27 := 0; fi27 < 11; fi27++ {
            v25 -= 471
            _ = fi27
        }
        _ = v25
        _ = v26
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        _ = len(v3)
        var v29 uint16 = 141
        for fi30 := 0; fi30 < 5; fi30++ {
            v31 := GenS5{F0: 589, F1: (fi28 << ((403) & 63))}
            _ = fi30
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    v14 -= (3.14159 * (0.0 / v14))
    for fi32 := 0; fi32 < 17; fi32++ {
        fmt.Println(v14, fi32)
        var v33 uint32 = 0
        for fi34 := 0; fi34 < 18; fi34++ {
            v35 := v3[v33]
            _ = fi34
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        v14 += float32(0.1)
        _ = fi36
    }
    v37 := GenS6{F0: 1, F1: -128}
    v38 := v3[2147483647]
    v39 := GenS4{F0: -104}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
}
