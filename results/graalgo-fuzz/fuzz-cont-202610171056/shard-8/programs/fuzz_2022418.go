package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64; F1 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 42
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    v5 := make(map[int]bool)
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(v2)
        _ = fi7
    }
    if (2147483648 <= 42) {
        v8 := GenS1{F0: 2147483647, F1: -572}
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 string = ("" + (((((("\n\t" + ("\n\t" + (("\n\t" + "foo bar") + "a"))) + ("a" + "a")) + ("" + ((("foo bar" + "hello") + ((("\n\t" + "foo bar") + "") + "a")) + ""))) + "a") + "") + "foo bar"))
            var v11 uint8 = 255
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 int32 = -247
        v13 = (v13 * (v13 % (((v13 - (^(v13)))) | 1)))
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 0, F1: 42}
    var v15 uint = 1
    v2 += (^(v2))
    v15 = v15
    fmt.Println(v15, v15)
    delete(v3, int(v15))
    switch "a" {
    case "foo bar":
        switch -170 {
        case 4294967295:
            v16 := GenS2{F0: 42, F1: 32767}
            _ = v16
        case -60:
            for fi17 := 0; fi17 < 17; fi17++ {
                if true {
                    switch "" {
                    case "\n\t":
                        v18 := GenS1{F0: -1, F1: 42}
                        _ = v18
                    case "a":
                        v2 = int8(v15)
                    }
                }
                _ = fi17
            }
        }
    case "\n\t":
        switch ("hello" + "") {
        case "foo bar":
            v2 = -78
        }
    }
    v19 := GenS2{F0: -334, F1: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
