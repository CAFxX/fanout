package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 42
    v3 := make(map[int]bool)
    fmt.Println(v2)
    var v4 bool = (false || true)
    fmt.Println(v2, v4)
    for fi5 := 0; fi5 < 16; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 bool = (231 < 2147483647)
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    switch ((((("foo bar" + ("hello" + (("" + ("hello" + "a")) + "hello"))) + (("a" + ("hello" + ("" + "foo bar"))) + "\n\t")) + "a") + "\n\t") + (("foo bar" + ("" + "\n\t")) + "foo bar")) {
    case "":
        v8 := []int{}
        _ = v8
    case "foo bar":
        switch 128 {
        case 42:
            fmt.Println(v2)
        case 2147483647:
            for fi9 := 0; fi9 < 18; fi9++ {
                fmt.Println(v4, v4, v4)
                _ = fi9
            }
        }
    default:
        v10 := GenS1{F0: "a", F1: v2}
        _ = v10
    }
    v2 += (v2 - v2)
    v11 := GenS1{F0: "\n\t", F1: (v2 / (((v2 * (v2 >> ((813) & 63)))) | 1))}
    var v12 int = 0
    fmt.Println(v2)
    var v13 uint8 = 127
    if ("\n\t" >= "\n\t") {
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS1{F0: "\n\t", F1: (v2 | v2)}
            v12 += 127
            v16 := GenS1{F0: "foo bar", F1: 1}
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        if v4 {
            v17 := GenS1{F0: ("" + "hello"), F1: -1}
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := []int{}
                _ = fi18
                _ = v19
            }
            if (!true) {
                if true {
                    v20 := GenS2{F0: 2147483648}
                    for fi21 := 0; fi21 < 7; fi21++ {
                        var v22 uint = 115
                        var v23 uint64 = 336
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    v24 := v1[0]
                    _ = v20
                    _ = v24
                } else {
                    v12 -= v12
                }
                v13 -= 100
                v4 = v4
            }
            _ = v17
        }
        var v25 uint16 = 0
        if v4 {
            var v26 float64 = (1e10 + 0.1)
            switch ("a" + "foo bar") {
            case "a":
                v12 += 662
            default:
                _ = len(v3)
            }
            v25 += (v25 ^ 0)
            _ = v26
        } else {
            var v27 int16 = 0
            for fi28 := 0; fi28 < 10; fi28++ {
                for fi29 := 0; fi29 < 4; fi29++ {
                    _ = len(v1)
                    v25 += 468
                    _ = fi29
                }
                _ = fi28
            }
            _ = v27
        }
        _ = v25
    }
    var v30 float32 = 1e10
    var v31 int16 = 967
    v32 := GenS3{F0: 0, F1: v13}
    for fi33 := 0; fi33 < 19; fi33++ {
        var v34 uint64 = 505
        var v35 uint8 = v13
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
    _ = v4
}
