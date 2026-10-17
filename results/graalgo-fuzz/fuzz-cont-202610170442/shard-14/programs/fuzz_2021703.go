package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int8 }
type GenS2 struct { F0 int16; F1 uint32; F2 bool }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 bool; F1 int32 }

func main() {
    v1 := []string{}
    var v2 uint8 = 161
    v3 := make(map[int]bool)
    var v4 uint32 = 2147483647
    if false {
        if len(v1) > 0 {
            v1[0] = ("a" + "\n\t")
        }
        if false {
            v4 += 2
            for fi5 := 0; fi5 < 8; fi5++ {
                fmt.Println(v4, v2, v4)
                _ = fi5
            }
            switch -120 {
            case 9223372036854775807:
                var v6 uint = 255
                _ = v6
            }
        }
    }
    v4 += v4
    switch "a" {
    case "a":
        for fi7 := 0; fi7 < 1; fi7++ {
            _ = len(v1)
            switch 127 {
            case -2147483648:
                v4 += v4
            case -9223372036854775808:
                fmt.Println(v4, fi7, v4)
            case 42:
                if true {
                    v4 = v4
                    _ = len(v1)
                }
            default:
                v8 := GenS1{F0: "foo bar", F1: "foo bar", F2: 42}
                _ = v8
            }
            _ = fi7
        }
    case "foo bar":
        v4 = (v4 >> ((206) & 63))
    default:
        for fi9 := 0; fi9 < 16; fi9++ {
            v4 = (v4 - 65535)
            v10 := GenS2{F0: -1, F1: (v4 * 48), F2: (true || ((-1 == 1000) || (4294967295 < 0)))}
            _ = fi9
            _ = v10
        }
    }
    fmt.Println(v2, v4)
    switch "foo bar" {
    case "foo bar":
        v11 := GenS1{F0: ((("\n\t" + "") + "foo bar") + "a"), F1: (("hello" + ("\n\t" + ("" + ""))) + ("foo bar" + ((("\n\t" + "hello") + "a") + "foo bar"))), F2: 10}
        _ = v11
    case "hello":
        var v12 uint8 = 126
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 7; fi13++ {
            v14 := GenS1{F0: "\n\t", F1: (((("" + "hello") + "a") + "foo bar") + ("foo bar" + "\n\t")), F2: 42}
            v4 = 2147483647
            v2 -= (^(v2))
            _ = fi13
            _ = v14
        }
        v15 := GenS3{F0: 625, F1: ""}
        var v16 int = 1
        _ = v15
        _ = v16
    } else {
        switch -401 {
        case 290:
            var v17 float64 = 0.1
            _ = v17
        case 7:
            if true {
                v2 += (v2 << ((0) & 63))
            }
        case 1:
            for fi18 := 0; fi18 < 4; fi18++ {
                v4 = v4
                fmt.Println(v2)
                _ = fi18
            }
        }
        v4 += v4
        for fi19 := 0; fi19 < 7; fi19++ {
            if (2 < 655) {
                for fi20 := 0; fi20 < 13; fi20++ {
                    v4 -= 4294967295
                    _ = fi20
                }
            }
            var v21 uint = 1
            var v22 uint8 = (v2 << ((uint(v2)) & 63))
            _ = fi19
            _ = v21
            _ = v22
        }
    }
    fmt.Println(v2, v4)
    var v23 uint = 4294967295
    if true {
        if false {
            switch (("a" + "\n\t") + "a") {
            case "hello":
                var v24 float32 = (-2.25 + 1e10)
                _ = v24
            case "foo bar":
                v25 := GenS4{F0: true, F1: int32(v23)}
                _ = v25
            }
            var v26 uint32 = 4294967295
            _ = v26
        } else {
            switch "hello" {
            case "a":
                v27 := GenS4{F0: (v2 > v2), F1: 788}
                _ = v27
            case "":
                if len(v1) > 0 {
                    v1[0] = ("hello" + "")
                }
            default:
                var v28 float32 = (1e10 + 1.5)
                _ = v28
            }
        }
        v23 += 4294967295
    } else {
        var v29 bool = true
        _ = v29
    }
    var v30 int8 = 100
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v4
}
