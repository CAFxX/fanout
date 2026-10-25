package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int }
type GenS2 struct { F0 uint16; F1 int; F2 string }

func main() {
    v1 := []int{746, -1, 1, 128}
    v2 := make(map[string]uint32)
    var v3 float64 = -2.25
    for fi4 := 0; fi4 < 12; fi4++ {
        switch (("\n\t" + "") + (("" + "a") + "a")) {
        case "hello":
            v5 := GenS1{F0: 0, F1: 1, F2: (fi4 << ((255) & 63))}
            _ = v5
        }
        _ = fi4
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        _ = len(v2)
        _ = fi6
    }
    var v7 int64 = 1
    for fi8 := 0; fi8 < 12; fi8++ {
        switch (("" + (("hello" + "foo bar") + "a")) + "hello") {
        case "foo bar":
            switch (-(fi8)) {
            case -889:
                fmt.Println(fi8)
            case 128:
                v1 = append(v1, fi8)
            case 774:
                var v9 uint32 = 42
                _ = v9
            default:
                v10 := v2["foo bar"]
                _ = v10
            }
        case "hello":
            for fi11 := 0; fi11 < 1; fi11++ {
                if false {
                    v3 = (1.5 * 0.1)
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        v2["hello"] = 1
                        v12 := v2[("\n\t" + ("" + "\n\t"))]
                        _ = v12
                    case "a":
                        fmt.Println(v3, fi11)
                    case "hello":
                        fmt.Println(v7)
                    default:
                        v13 := GenS1{F0: 100, F1: -1, F2: fi11}
                        _ = v13
                    }
                    var v14 int64 = (v7 * (v7 / ((v7) | 1)))
                    _ = v14
                }
                var v15 int16 = 42
                var v16 uint16 = 105
                var v17 float32 = -2.25
                v18 := v2["hello"]
                _ = fi11
                _ = v15
                _ = v16
                _ = v17
                _ = v18
            }
        }
        _ = fi8
    }
    if true {
        if ((!false) || (!(false && (42 != -185)))) {
            for fi19 := 0; fi19 < 15; fi19++ {
                if (!(942 != 2)) {
                    fmt.Println(v7, v7)
                }
                v3 += v3
                v3 = v3
                _ = fi19
            }
            var v20 bool = false
            _ = v20
        }
    }
    v3 -= (v3 * v3)
    var v21 uint64 = 695
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := v2["\n\t"]
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi22
        _ = v23
    }
    var v24 float32 = 1.5
    for fi25 := 0; fi25 < 12; fi25++ {
        switch "foo bar" {
        case "hello":
            switch "a" {
            case "hello":
                if ((9223372036854775807 != (^(fi25))) || false) {
                    switch -1 {
                    case 9223372036854775807:
                        _ = len(v2)
                    case 42:
                        v3 += v3
                    default:
                        var v26 uint = 18446744073709551615
                        _ = v26
                    }
                    var v27 int64 = (v7 / ((647) | 1))
                    var v28 bool = true
                    _ = v27
                    _ = v28
                }
            case "foo bar":
                v1 = append(v1, (fi25 >> ((1) & 63)))
            case "a":
                v29 := GenS2{F0: 786, F1: 9223372036854775807, F2: "hello"}
                _ = v29
            }
        case "foo bar":
            v7 -= (v7 & (v7 + v7))
        }
        v1 = append(v1, 127)
        if len(v1) > 0 {
            v1[0] = (fi25 >> ((26) & 63))
        }
        _ = fi25
    }
    if (-1 <= 42) {
        delete(v2, "\n\t")
        switch 1 {
        case 920:
            v30 := GenS2{F0: 276, F1: 0, F2: "foo bar"}
            _ = v30
        default:
            v21 = (v21 - (v21 | v21))
        }
    }
    v7 = v7
    v21 = 18446744073709551615
    v3 = v3
    v31 := GenS1{F0: 42, F1: 119, F2: -710}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v31
    _ = v7
}
