package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 bool; F2 int8 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 int16 = -100
    var v2 string = "a"
    v3 := []string{("\n\t" + ("" + v2)), v2}
    v4 := make(map[string]bool)
    v2 = ""
    v1 -= 482
    for fi5 := 0; fi5 < 8; fi5++ {
        switch 32768 {
        case -329:
            fmt.Println(v2)
        case -2147483648:
            v6 := GenS1{F0: false}
            _ = v6
        }
        _ = fi5
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        switch v2 {
        case "a":
            switch v2 {
            case "\n\t":
                var v8 string = ((v2 + v2) + "a")
                _ = v8
            case "a":
                v2 = "foo bar"
            }
        case "\n\t":
            var v9 float64 = 0.0
            _ = v9
        case "foo bar":
            var v10 int = fi7
            _ = v10
        }
        v11 := GenS2{F0: (((("\n\t" + ((v2 + "hello") + "")) + "") + v2) + v2)}
        var v12 int8 = 1
        _ = fi7
        _ = v11
        _ = v12
    }
    v1 = v1
    for fi13 := 0; fi13 < 19; fi13++ {
        if ((false && false) && (18446744073709551615 != 42)) {
            for fi14 := 0; fi14 < 18; fi14++ {
                v1 -= (v1 % ((1) | 1))
                v2 = (v2 + "\n\t")
                _ = fi14
            }
            switch fi13 {
            case 9223372036854775807:
                v15 := GenS3{F0: 18446744073709551615, F1: true, F2: 0}
                _ = v15
            case 727:
                v16 := GenS2{F0: "hello"}
                _ = v16
            case -418:
                var v17 int32 = 2147483647
                _ = v17
            }
            var v18 int = fi13
            _ = v18
        }
        _ = len(v4)
        if (1 <= 127) {
            if false {
                var v19 string = v2
                _ = v19
            }
            if true {
                var v20 uint64 = 18446744073709551615
                var v21 uint32 = 65535
                v1 -= (v1 >> ((0) & 63))
                _ = v20
                _ = v21
            }
        }
        _ = fi13
    }
    if (42 == 127) {
        v22 := GenS4{F0: 4294967295, F1: 955}
        _ = v22
    } else {
        v2 = (((v2 + ("foo bar" + "a")) + ("foo bar" + ("" + v2))) + v2)
        v1 += v1
        var v23 string = (("" + v2) + v2)
        _ = v23
    }
    v2 = (v2 + ("foo bar" + v2))
    if (!false) {
        var v24 uint = 18446744073709551615
        if (((v24 | (v24 ^ (v24 / ((42) | 1)))) == 1) || false) {
            var v25 int16 = v1
            for fi26 := 0; fi26 < 11; fi26++ {
                v27 := GenS5{F0: 702, F1: 708}
                _ = fi26
                _ = v27
            }
            _ = v25
        }
        v2 = v2
        _ = v24
    }
    var v28 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
