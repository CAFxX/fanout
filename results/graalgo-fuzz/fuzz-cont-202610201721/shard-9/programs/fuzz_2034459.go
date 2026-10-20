package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 string; F1 uint8; F2 uint }
type GenS3 struct { F0 uint64; F1 int; F2 uint64 }
func genFn0(p0 int64, p1 int16) int8 {
    p0 += 9223372036854775807
    if false {
        if true {
            if (100 >= -123) {
                var v7 uint32 = 4294967295
                var v8 uint16 = 0
                v9 := GenS1{F0: -1, F1: p1}
                _ = v7
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS1{F0: int32(fi10), F1: 1000}
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 4; fi12++ {
                v13 := GenS2{F0: "foo bar", F1: 127, F2: 4294967295}
                p0 -= p0
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := []int{}
                    if ((false || false) && ((1.5 - (3.14159 * (1.5 - ((-2.25 + -2.25) + 0.1)))) > 1.5)) {
                        fmt.Println(fi14, fi14)
                        var v16 uint = 42
                        _ = v16
                    } else {
                        if len(v15) > 0 {
                            v15[0] = (^(fi14))
                        }
                        p1 = -614
                        var v17 string = ""
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
        }
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 uint64 = 42
        fmt.Println(fi18)
        var v20 uint64 = (v19 & 352)
        _ = fi18
        _ = v19
        _ = v20
    }
    p0 = -17
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int16) bool {
    if true {
        p0 += (^(p0))
    }
    v21 := GenS2{F0: "\n\t", F1: 128, F2: 0}
    _ = p0
    _ = v21
    return false
}

func main() {
    var v1 int32 = -29
    var v2 uint64 = 843
    v3 := []string{}
    var v4 uint32 = 0
    v5 := make(map[string]string)
    v6 := []string{"\n\t"}
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 uint8 = 0
        _ = fi22
        _ = v23
    }
    if true {
        for fi24 := 0; fi24 < 11; fi24++ {
            switch fi24 {
            case -1:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v25 := v5[""]
                _ = v25
            }
            _ = fi24
        }
        fmt.Println(v1, v1, v2)
        v4 = 1
    }
    v2 = v2
    v4 += v4
    if true {
        switch ("hello" + "a") {
        case "a":
            v26 := GenS2{F0: "foo bar", F1: 68, F2: 2}
            _ = v26
        default:
            fmt.Println(v2, v4)
        }
    }
    if ("\n\t" < ((("foo bar" + ("a" + "hello")) + ("foo bar" + "hello")) + "a")) {
        if false {
            v27 := GenS1{F0: 2147483647, F1: 66}
            _ = v27
        }
        v28 := GenS3{F0: (v2 / ((v2) | 1)), F1: -206, F2: (^(v2))}
        var v29 string = "a"
        _ = v28
        _ = v29
    }
    var v30 int = -216
    fmt.Println(v2, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
