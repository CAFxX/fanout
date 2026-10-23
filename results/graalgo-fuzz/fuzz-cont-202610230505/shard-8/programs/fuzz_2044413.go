package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8 }
func genFn0() int16 {
    var v4 uint8 = 128
    fmt.Println(v4)
    v5 := GenS1{F0: 33}
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]string)
    var v3 uint16 = 32767
    if (((v1 & (^(v1))) != (v1 >> (((v1 / ((18446744073709551615) | 1))) & 63))) || (!(true || (v3 != 893)))) {
        switch -9223372036854775808 {
        case 131:
            v6 := GenS2{F0: -1, F1: (v3 % (((v3 * 255)) | 1))}
            _ = v6
        case 2147483647:
            if (-1 >= 9223372036854775807) {
                if true {
                    var v7 uint32 = 2
                    if false {
                        v8 := []int{32768, -1}
                        for fi9 := 0; fi9 < 2; fi9++ {
                            v3 = v3
                            var v10 float64 = 0.1
                            _ = fi9
                            _ = v10
                        }
                        _ = v8
                    } else {
                        v11 := GenS3{F0: (((((("" + "\n\t") + "") + "foo bar") + (("a" + "") + "hello")) + ((("" + "hello") + "") + "foo bar")) + ("hello" + ("\n\t" + "")))}
                        switch (("" + "hello") + (("\n\t" + "") + ("a" + "foo bar"))) {
                        case "\n\t":
                            v12 := GenS2{F0: 1, F1: v3}
                            _ = v12
                        default:
                            var v13 int32 = 2147483647
                            _ = v13
                        }
                        v14 := []int{-2147483648, -1, 127, 8}
                        _ = v11
                        _ = v14
                    }
                    _ = v7
                }
                fmt.Println(v3)
            }
        case -187:
            fmt.Println(v1, v1)
        }
    } else {
        var v15 int8 = 100
        _ = v15
    }
    v1 = 757
    var v16 int = 0
    if ((^(v16)) >= v16) {
        v16 -= 32767
        var v17 int32 = 42
        v17 += (-(v17))
        _ = v17
    }
    switch "" {
    case "a":
        for fi18 := 0; fi18 < 8; fi18++ {
            if (1 != -1) {
                var v19 uint64 = 519
                if (-81 == -105) {
                    var v20 uint64 = 18446744073709551615
                    v1 += v1
                    v21 := GenS3{F0: "\n\t"}
                    _ = v20
                    _ = v21
                }
                _ = v19
            }
            switch 42 {
            case -9223372036854775808:
                for fi22 := 0; fi22 < 18; fi22++ {
                    switch fi22 {
                    case 805:
                        var v23 int16 = -487
                        _ = v23
                    case -1:
                        v24 := []int{-121, 9223372036854775807, v16, v16}
                        _ = v24
                    default:
                        if true {
                            v25 := []int{fi18, 128}
                            _ = v25
                        }
                        v26 := GenS2{F0: -1, F1: (^(v3))}
                        _ = v26
                    }
                    var v27 float32 = (3.14159 + 0.1)
                    _ = fi22
                    _ = v27
                }
            }
            _ = fi18
        }
    case "foo bar":
        for fi28 := 0; fi28 < 10; fi28++ {
            fmt.Println(fi28)
            v3 -= v3
            _ = fi28
        }
    case "\n\t":
        v29 := GenS3{F0: ((("foo bar" + "foo bar") + ("a" + "\n\t")) + "a")}
        _ = v29
    default:
        switch ("\n\t" + "hello") {
        case "foo bar":
            v30 := []int{9223372036854775807, v16, 9223372036854775807, v16}
            _ = v30
        case "a":
            v1 = 255
        case "":
            for fi31 := 0; fi31 < 19; fi31++ {
                delete(v2, 639)
                _ = fi31
            }
        }
    }
    fmt.Println(v16, v3, v16)
    v3 = (^(v3))
    if true {
        fmt.Println(v3, v3, v16)
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 int16 = -32768
            fmt.Println(v3)
            v34 := GenS4{F0: v3}
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    if (!true) {
        var v35 uint8 = 0
        _ = v35
    }
    v36 := GenS5{F0: uint8(v1)}
    var v37 int64 = 358
    v38 := []int{}
    for fi39 := 0; fi39 < 11; fi39++ {
        var v40 int64 = v37
        _ = fi39
        _ = v40
    }
    v16 -= (v16 >> ((v1) & 63))
    v3 = 470
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}
