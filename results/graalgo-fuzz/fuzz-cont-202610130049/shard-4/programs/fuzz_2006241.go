package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int8; F1 uint32 }
func genFn0(p0 int, p1 int) int {
    switch (p1 % (((p0 & 9223372036854775807)) | 1)) {
    case 127:
        v4 := []int{(p0 - p1), p1, p0, (p1 | 32767)}
        _ = v4
    case 1:
        v5 := []int{0}
        _ = v5
    }
    v6 := []int{32768}
    fmt.Println(p1, p0)
    v7 := GenS1{F0: (p1 & p1), F1: 42}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 << ((1) & 63))
}
func genFn1(p0 int8) int32 {
    v8 := GenS2{F0: 0, F1: 32768}
    fmt.Println(p0)
    p0 += p0
    _ = p0
    _ = v8
    return -971
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]int)
    var v3 int8 = 127
    v1 = (v1 - 0)
    if ("a" == "foo bar") {
        switch 4294967295 {
        case 465:
            fmt.Println(v3, v1)
        case 1:
            if true {
                v9 := GenS3{F0: 2147483647}
                var v10 float64 = 0.1
                switch "" {
                case "":
                    var v11 uint = 489
                    _ = v11
                }
                _ = v10
                _ = v9
            } else {
                v12 := GenS3{F0: 502}
                _ = len(v2)
                v3 = (^(v3))
                _ = v12
            }
        case 445:
            var v13 uint64 = 737
            _ = v13
        default:
            fmt.Println(v1)
        }
    }
    v3 = (^(v3))
    v3 = (v3 << ((255) & 63))
    switch "a" {
    case "\n\t":
        v3 = v3
    case "a":
        v3 += (v3 >> ((65535) & 63))
    case "hello":
        switch (((("hello" + "\n\t") + ("foo bar" + "foo bar")) + (("foo bar" + "hello") + "a")) + "") {
        case "hello":
            v14 := v2[2]
            _ = v14
        }
    }
    if false {
        var v15 uint32 = 65535
        if false {
            var v16 uint8 = 127
            var v17 uint32 = 974
            var v18 float64 = (1e10 / float64((3.14159 + (float32(((1e10 + -2.25) / 1.5)) - 1.5))))
            _ = v16
            _ = v17
            _ = v18
        } else {
            fmt.Println(v15)
        }
        fmt.Println(v1, v3)
        _ = v15
    } else {
        if ((32767 != 42) || false) {
            v19 := GenS4{F0: 1, F1: 32768}
            var v20 uint = 2
            switch 9223372036854775807 {
            case 4294967295:
                if true {
                    v3 -= v3
                    for fi21 := 0; fi21 < 18; fi21++ {
                        for fi22 := 0; fi22 < 2; fi22++ {
                            v20 -= (^(v20))
                            _ = fi22
                        }
                        fmt.Println(fi21)
                        _ = fi21
                    }
                } else {
                    delete(v2, uint32(v20))
                    var v23 uint32 = 758
                    _ = v23
                }
            }
            _ = v19
            _ = v20
        } else {
            v24 := GenS4{F0: 1, F1: -1}
            _ = v24
        }
        fmt.Println(v1)
    }
    delete(v2, 889)
    var v25 uint16 = 497
    for fi26 := 0; fi26 < 1; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            v28 := GenS4{F0: (v25 << ((2) & 63)), F1: (fi26 / ((fi26) | 1))}
            var v29 int64 = (v1 / ((v1) | 1))
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi26
    }
    v30 := GenS5{F0: (v3 * 38), F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
}
