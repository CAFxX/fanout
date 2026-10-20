package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 string }
func genFn0() string {
    v5 := []int{32768, 9223372036854775807, -745}
    var v6 int = 127
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint64 = 2
        switch (("foo bar" + "") + "") {
        case "hello":
            v9 = (^(v9))
        case "\n\t":
            v10 := GenS1{F0: 67}
            _ = v10
        default:
            var v11 uint16 = 1
            _ = v11
        }
        v5 = append(v5, 0)
        _ = fi8
        _ = v9
    }
    v6 = (v6 / ((1) | 1))
    _ = v5
    _ = v6
    _ = v7
    return "a"
}
func genFn1(p0 int32, p1 uint8) string {
    p0 += (^(p0))
    var v12 uint32 = 1
    if true {
        fmt.Println(p1)
        var v13 uint64 = 0
        _ = v13
    } else {
        var v14 uint64 = 42
        _ = v14
    }
    v15 := GenS2{F0: (-710 != 640)}
    _ = p0
    _ = p1
    _ = v12
    _ = v15
    return "a"
}

func main() {
    v1 := []int{-1}
    var v2 uint64 = 2
    v3 := make(map[uint32]int)
    var v4 uint16 = 651
    fmt.Println(v4, v2)
    _ = len(v3)
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS2{F0: true}
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 uint8 = 255
            _ = fi18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 uint32 = 0
            switch 127 {
            case 42:
                v23 := GenS2{F0: false}
                _ = v23
            case -307:
                var v24 int64 = -1
                _ = v24
            }
            v2 = 42
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    switch -429 {
    case 32767:
        if (127 == 42) {
            for fi25 := 0; fi25 < 4; fi25++ {
                v4 -= v4
                switch "a" {
                case "\n\t":
                    var v26 uint16 = (v4 - 65535)
                    _ = v26
                default:
                    switch 594 {
                    case 1:
                        var v27 int8 = 127
                        _ = v27
                    }
                }
                _ = fi25
            }
        } else {
            v28 := GenS2{F0: false}
            var v29 int8 = 127
            _ = v28
            _ = v29
        }
    case 414:
        for fi30 := 0; fi30 < 2; fi30++ {
            fmt.Println(v4, v2, v4)
            fmt.Println(fi30, v4, fi30)
            _ = fi30
        }
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        switch ("a" + "a") {
        case "hello":
            switch "\n\t" {
            case "foo bar":
                v4 += (^(v4))
            case "":
                var v32 int8 = -1
                _ = v32
            case "hello":
                var v33 uint16 = (v4 - v4)
                _ = v33
            default:
                var v34 uint64 = (v2 + 0)
                _ = v34
            }
        default:
            for fi35 := 0; fi35 < 10; fi35++ {
                var v36 bool = false
                if v36 {
                    if ((fi35 & 2147483647) == (fi35 % ((fi31) | 1))) {
                        var v37 bool = true
                        var v38 int8 = -23
                        var v39 float32 = 0.0
                        _ = v37
                        _ = v38
                        _ = v39
                    } else {
                        var v40 uint8 = 215
                        _ = v40
                    }
                    var v41 int16 = 42
                    _ = v41
                } else {
                    fmt.Println(v4, v4)
                    var v42 uint8 = 128
                    _ = v42
                }
                _ = fi35
                _ = v36
            }
        }
        _ = fi31
    }
    v4 = (v4 << ((42) & 63))
    var v43 int16 = -32768
    switch (("hello" + "hello") + "a") {
    case "a":
        v43 -= v43
    case "\n\t":
        v1 = append(v1, -1)
    case "hello":
        if true {
            v44 := GenS3{F0: 42, F1: ((("hello" + ((("foo bar" + "a") + "foo bar") + "a")) + "foo bar") + "hello")}
            for fi45 := 0; fi45 < 14; fi45++ {
                fmt.Println(v2)
                _ = fi45
            }
            _ = v44
        }
    }
    for fi46 := 0; fi46 < 15; fi46++ {
        v4 = 888
        _ = fi46
    }
    var v47 int32 = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v47
}
