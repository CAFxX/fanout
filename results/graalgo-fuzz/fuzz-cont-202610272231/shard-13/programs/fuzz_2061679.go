package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }

func main() {
    var v1 float64 = 1.5
    var v2 int8 = 0
    v3 := make(map[int]uint32)
    fmt.Println(v2)
    var v4 uint = 703
    for fi5 := 0; fi5 < 18; fi5++ {
        switch fi5 {
        case 9223372036854775807:
            if false {
                var v6 int16 = 32767
                var v7 uint = v4
                switch ((((("foo bar" + (("foo bar" + "foo bar") + "\n\t")) + (((("" + "a") + ("\n\t" + "foo bar")) + "\n\t") + "hello")) + ("hello" + "hello")) + "hello") + "foo bar") {
                case "":
                    v7 -= v4
                case "foo bar":
                    v2 += (v2 >> ((v7) & 63))
                default:
                    v8 := GenS1{F0: 981, F1: (fi5 >> ((v7) & 63))}
                    _ = v8
                }
                _ = v6
                _ = v7
            } else {
                switch "a" {
                case "hello":
                    for fi9 := 0; fi9 < 19; fi9++ {
                        if (true && ((0.1 + (-2.25 * 3.14159)) == -2.25)) {
                            if (-580 > 0) {
                                fmt.Println(v2, v2)
                                var v10 uint32 = 42
                                _ = v10
                            }
                            v1 -= ((v1 + (v1 + float64((3.14159 + 3.14159)))) * v1)
                            v11 := GenS1{F0: 211, F1: (fi5 % ((32768) | 1))}
                            _ = v11
                        }
                        switch (fi9 << (((v4 + v4)) & 63)) {
                        case 42:
                            v12 := []int{(-(fi9)), (fi9 | fi5)}
                            _ = v12
                        }
                        _ = fi9
                    }
                }
                if true {
                    v1 += float64(3.14159)
                    var v13 int32 = 1
                    var v14 int = fi5
                    _ = v13
                    _ = v14
                }
                v15 := []int{fi5, (fi5 + (fi5 << (((v4 | v4)) & 63))), -545, fi5}
                _ = v15
            }
        default:
            v16 := []int{}
            _ = v16
        }
        var v17 uint8 = 42
        v2 = v2
        _ = fi5
        _ = v17
    }
    fmt.Println(v4, v2)
    v18 := []int{127}
    v19 := GenS2{F0: false}
    switch ("a" + "") {
    case "":
        v18 = append(v18, -563)
    case "foo bar":
        v20 := GenS3{F0: 0}
        _ = v20
    case "hello":
        _ = len(v18)
    }
    var v21 bool = (32767 <= 127)
    if (-564 <= 9223372036854775807) {
        _ = len(v3)
    } else {
        fmt.Println(v2)
        v4 -= (v4 % ((v4) | 1))
    }
    var v22 int32 = 42
    var v23 uint32 = 4294967295
    v21 = (!((!false) || (!v21)))
    var v24 float32 = -2.25
    v24 -= v24
    v4 -= 0
    var v25 uint = (v4 | v4)
    if v21 {
        for fi26 := 0; fi26 < 3; fi26++ {
            var v27 uint8 = 42
            v27 += 128
            v28 := GenS3{F0: 0}
            _ = fi26
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            v2 += v2
            v21 = (1 == uint16(v23))
            for fi30 := 0; fi30 < 14; fi30++ {
                v4 = v4
                var v31 bool = v21
                var v32 int64 = -9223372036854775808
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi29
        }
        v33 := GenS3{F0: -541}
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
