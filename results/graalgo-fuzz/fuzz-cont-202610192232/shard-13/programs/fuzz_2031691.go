package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint32, p1 uint32, p2 string) int64 {
    p1 = 0
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint8 = 255
        var v7 int16 = 32767
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 int64 = -1
    var v2 uint = 0
    var v3 int8 = -1
    var v4 bool = (-246 >= -2147483648)
    fmt.Println(v3, v4, v4)
    switch (((((("a" + (("a" + "a") + (("foo bar" + "foo bar") + "a"))) + "hello") + "hello") + "\n\t") + "a") + "a") {
    case "":
        var v8 uint32 = 65535
        _ = v8
    case "hello":
        switch ("foo bar" + ("a" + "hello")) {
        case "":
            if v4 {
                v9 := []int{9223372036854775807, 0}
                _ = v9
            }
        case "foo bar":
            switch ("a" + "") {
            case "a":
                var v10 uint64 = 4294967295
                _ = v10
            case "":
                var v11 uint = v2
                _ = v11
            default:
                var v12 string = ("hello" + ("foo bar" + "hello"))
                _ = v12
            }
        default:
            for fi13 := 0; fi13 < 7; fi13++ {
                switch 1 {
                case -1:
                    for fi14 := 0; fi14 < 13; fi14++ {
                        switch 127 {
                        case -1:
                            v15 := []int{fi13, int(v2), (^(fi13))}
                            _ = v15
                        case 2147483647:
                            if (1000 <= 32767) {
                                v16 := []int{(fi13 >> (((v2 - (v2 | v2))) & 63))}
                                v2 = (v2 + v2)
                                var v17 uint32 = 65535
                                _ = v16
                                _ = v17
                            } else {
                                fmt.Println(v4)
                                var v18 uint64 = 1
                                _ = v18
                            }
                        case 4294967295:
                            var v19 uint8 = 136
                            _ = v19
                        default:
                            var v20 uint16 = 170
                            _ = v20
                        }
                        _ = fi14
                    }
                case 9223372036854775807:
                    var v21 uint8 = 255
                    _ = v21
                case -505:
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v23 := GenS1{F0: v1}
                        v3 = (v3 * (v3 >> ((0) & 63)))
                        v24 := GenS1{F0: (v1 * (v1 & v1))}
                        v4 = v4
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
                _ = fi13
            }
        }
    case "\n\t":
        v3 = v3
    default:
        v25 := GenS2{F0: -1, F1: v3}
        _ = v25
    }
    var v26 int8 = (v3 ^ (v3 / ((v3) | 1)))
    switch -891 {
    case 32768:
        if v4 {
            fmt.Println(v26, v1, v26)
            if true {
                if true {
                    for fi27 := 0; fi27 < 14; fi27++ {
                        fmt.Println(v26, fi27, v4)
                        _ = fi27
                    }
                    var v28 uint64 = 162
                    _ = v28
                }
                v1 += (v1 % ((419) | 1))
            } else {
                var v29 int8 = 127
                if v4 {
                    v30 := GenS1{F0: (v1 * (v1 ^ (v1 * 241)))}
                    switch ("" + (("\n\t" + "a") + "\n\t")) {
                    case "a":
                        fmt.Println(v2, v2)
                    case "\n\t":
                        var v31 int16 = -202
                        _ = v31
                    }
                    _ = v30
                }
                fmt.Println(v29)
                _ = v29
            }
        }
    case 660:
        var v32 int64 = (v1 * 0)
        _ = v32
    case 9223372036854775807:
        if false {
            v2 -= v2
            v1 += v1
        }
    default:
        v33 := GenS3{F0: -2147483648}
        _ = v33
    }
    if false {
        v34 := []int{-1, 12, -167}
        var v35 bool = v4
        switch "foo bar" {
        case "foo bar":
            switch "hello" {
            case "a":
                v26 += v26
            }
        case "hello":
            var v36 int16 = -704
            _ = v36
        }
        _ = v34
        _ = v35
    }
    switch "\n\t" {
    case "a":
        if ((v1 + 1) <= v1) {
            v4 = (626 <= 4294967295)
            v37 := []int{-1, 128}
            _ = v37
        }
    }
    v26 = (v26 + v3)
    v1 = genFn0(2147483647, 617, "foo bar")
    for fi38 := 0; fi38 < 18; fi38++ {
        var v39 uint8 = 128
        _ = fi38
        _ = v39
    }
    v40 := GenS2{F0: -32768, F1: (v3 >> ((255) & 63))}
    v41 := []int{32767, 32768}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
