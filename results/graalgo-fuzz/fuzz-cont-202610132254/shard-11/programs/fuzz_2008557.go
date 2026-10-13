package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int8 }
func genFn0() bool {
    var v6 uint64 = 0
    var v7 int8 = 1
    _ = v6
    _ = v7
    return false
}

func main() {
    var v1 int8 = -78
    var v2 uint64 = 2
    var v3 int64 = -351
    v4 := make(map[string]bool)
    v5 := []string{(("foo bar" + "hello") + "a"), ("foo bar" + (("hello" + ("a" + "hello")) + (("\n\t" + "") + "hello")))}
    var v8 uint32 = 210
    v9 := GenS1{F0: "hello", F1: -326, F2: 0}
    if (!(255 < 255)) {
        for fi10 := 0; fi10 < 18; fi10++ {
            fmt.Println(v1)
            _ = fi10
        }
        fmt.Println(v1, v3)
    }
    _ = len(v4)
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            switch (^(fi11)) {
            case 2147483647:
                fmt.Println(v1)
            }
            delete(v4, "foo bar")
            if (uint(v8) < 255) {
                v13 := v4[("\n\t" + "\n\t")]
                delete(v4, ("foo bar" + "hello"))
                _ = v13
            } else {
                v1 = 103
                if true {
                    v14 := GenS2{F0: v8}
                    var v15 int16 = 0
                    _ = v14
                    _ = v15
                }
                switch fi11 {
                case 883:
                    v3 = int64(v2)
                case 32768:
                    if false {
                        v4["foo bar"] = false
                        fmt.Println(v2, fi11, fi11)
                        var v16 float64 = -2.25
                        _ = v16
                    }
                case 128:
                    var v17 uint64 = (v2 - v2)
                    _ = v17
                default:
                    v2 -= (v2 << ((255) & 63))
                }
            }
            if ((((v8 | (^(v8))) >= (v8 ^ 4294967295)) && ((201 < 255) && true)) || false) {
                var v18 float32 = 0.0
                _ = v18
            }
            _ = fi12
        }
        v19 := GenS2{F0: (v8 | v8)}
        v1 = 0
        _ = fi11
        _ = v19
    }
    var v20 uint8 = 97
    v1 = -63
    var v21 uint64 = (v2 << ((65535) & 63))
    v4[("a" + (((("" + "") + "foo bar") + (("\n\t" + ("foo bar" + "")) + ("\n\t" + ("hello" + ("hello" + "a"))))) + "hello"))] = (true || (!false))
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 string = ""
        v24 := GenS3{F0: -1, F1: (v1 - (-(v1)))}
        v25 := GenS2{F0: (v8 + v8)}
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    var v26 string = (("\n\t" + "") + "")
    v27 := v4["a"]
    v28 := v4["foo bar"]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
