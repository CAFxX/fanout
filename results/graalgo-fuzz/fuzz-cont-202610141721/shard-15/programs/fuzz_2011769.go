package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := []string{("hello" + (("hello" + "") + "")), "hello", "\n\t"}
    var v2 uint16 = 627
    var v3 float32 = 0.0
    v4 := []uint32{}
    switch "" {
    case "foo bar":
        for fi5 := 0; fi5 < 19; fi5++ {
            v4 = append(v4, 208)
            for fi6 := 0; fi6 < 9; fi6++ {
                fmt.Println(fi5, fi6, fi5)
                _ = fi6
            }
            switch fi5 {
            case 1:
                switch (fi5 >> ((58) & 63)) {
                case 9223372036854775807:
                    for fi7 := 0; fi7 < 5; fi7++ {
                        v3 = ((v3 * v3) * 1.5)
                        v2 -= v2
                        var v8 int64 = 0
                        _ = fi7
                        _ = v8
                    }
                }
            case -640:
                v3 += 0.0
            default:
                var v9 int64 = 42
                _ = v9
            }
            _ = fi5
        }
    case "\n\t":
        v3 -= v3
    case "a":
        var v10 string = (("hello" + "\n\t") + "a")
        _ = v10
    default:
        v3 = 0.0
    }
    var v11 int64 = 1
    for fi12 := 0; fi12 < 10; fi12++ {
        switch "a" {
        case "\n\t":
            for fi13 := 0; fi13 < 4; fi13++ {
                v11 -= (-(v11))
                _ = fi13
            }
        }
        v3 += 1e10
        _ = fi12
    }
    var v14 int16 = -32768
    switch "" {
    case "\n\t":
        var v15 int = 2147483647
        _ = v15
    case "foo bar":
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 uint16 = v2
            v18 := GenS1{F0: "a", F1: 255}
            v14 += v14
            _ = fi16
            _ = v17
            _ = v18
        }
    case "":
        var v19 uint = 65535
        _ = v19
    }
    var v20 uint64 = 688
    for fi21 := 0; fi21 < 20; fi21++ {
        v2 = v2
        var v22 int = fi21
        v23 := GenS1{F0: "a", F1: 1}
        _ = fi21
        _ = v22
        _ = v23
    }
    if false {
        v24 := GenS1{F0: "", F1: 1}
        if false {
            var v25 float64 = (1.5 * 1e10)
            _ = len(v4)
            fmt.Println(v11, v2, v25)
            _ = v25
        }
        v26 := GenS1{F0: ("a" + ("a" + "foo bar")), F1: 255}
        _ = v24
        _ = v26
    } else {
        for fi27 := 0; fi27 < 10; fi27++ {
            if (true || (-9223372036854775808 <= (v11 | int64(fi27)))) {
                var v28 int = fi27
                _ = v28
            }
            _ = fi27
        }
    }
    var v29 float32 = (v3 - 0.1)
    if true {
        var v30 float32 = v29
        v31 := GenS1{F0: "", F1: 153}
        for fi32 := 0; fi32 < 7; fi32++ {
            v33 := GenS2{F0: 1, F1: 0}
            var v34 uint64 = (^(v20))
            switch "\n\t" {
            case "foo bar":
                fmt.Println(fi32, v2)
            default:
                switch (("" + "a") + "\n\t") {
                case "\n\t":
                    var v35 int = 9223372036854775807
                    _ = v35
                case "":
                    v2 += v2
                case "hello":
                    v29 = ((float32(((-2.25 / 1.5) * 0.1)) - (v29 - 1e10)) + (1e10 / 0.1))
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v30
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
}
