package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 string }
func genFn0(p0 uint16, p1 uint, p2 int16) uint8 {
    var v5 uint64 = 975
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 int = (-(fi6))
        p2 += 689
        _ = fi6
        _ = v7
    }
    var v8 uint64 = 364
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 0
    var v3 uint = 871
    v4 := []string{("\n\t" + "a"), "a", ("a" + ("" + "foo bar")), ((("" + "hello") + "hello") + ("foo bar" + ("a" + "foo bar")))}
    var v9 int32 = -705
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS1{F0: 0}
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS2{F0: "hello", F1: 18446744073709551615}
            _ = fi12
            _ = v13
        }
        var v14 int8 = 125
        _ = fi10
        _ = v11
        _ = v14
    }
    v9 = 949
    v15 := GenS2{F0: ("\n\t" + "\n\t"), F1: 4294967295}
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int32 = (v9 / ((v9) | 1))
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 int8 = 0
            v20 := GenS2{F0: (((("hello" + "\n\t") + ("a" + "hello")) + "hello") + "hello"), F1: 1}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    switch "" {
    case "a":
        var v21 uint8 = genFn0(2, v3, 743)
        _ = v21
    default:
        var v22 uint8 = 128
        _ = v22
    }
    if false {
        if false {
            v9 = -10
            if true {
                v1[4294967295] = true
            } else {
                v23 := GenS3{F0: ("hello" + (("\n\t" + "") + "foo bar"))}
                _ = v23
            }
        } else {
            if (false && (true || (!(true && ((v2 * (v2 ^ v2)) == v2))))) {
                fmt.Println(v3)
                var v24 float64 = 3.14159
                _ = v24
            } else {
                if false {
                    fmt.Println(v3)
                }
            }
        }
        switch "a" {
        case "hello":
            switch ((("\n\t" + "\n\t") + "") + "") {
            case "foo bar":
                var v25 uint16 = 65535
                _ = v25
            }
        default:
            v26 := GenS2{F0: "", F1: 2}
            _ = v26
        }
    }
    _ = len(v1)
    fmt.Println(v2, v3, v2)
    fmt.Println(v3, v3, v2)
    var v27 uint32 = (v2 ^ (v2 ^ (v2 / ((v2) | 1))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v9
}
