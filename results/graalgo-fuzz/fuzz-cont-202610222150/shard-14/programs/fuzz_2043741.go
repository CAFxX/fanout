package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]bool)
    var v3 float32 = 1.5
    var v4 uint = (v1 - 42)
    if (32767 < 32767) {
        fmt.Println(v1, v1)
        _ = len(v2)
        if false {
            var v5 int16 = -1
            _ = v5
        }
    }
    v4 = 824
    for fi6 := 0; fi6 < 13; fi6++ {
        v1 -= 126
        fmt.Println(v3, fi6, v1)
        _ = fi6
    }
    v7 := GenS1{F0: 18446744073709551615}
    fmt.Println(v3)
    v1 = (v4 | (v1 + v1))
    v8 := []int{}
    var v9 uint32 = 0
    v10 := GenS1{F0: 18446744073709551615}
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := v2[("foo bar" + ("a" + "foo bar"))]
            _ = fi12
            _ = v13
        }
        switch 9223372036854775807 {
        case 32768:
            for fi14 := 0; fi14 < 2; fi14++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi14
            }
        case 9223372036854775807:
            fmt.Println(v9, fi11)
        case 974:
            v4 -= v4
        default:
            if true {
                for fi15 := 0; fi15 < 6; fi15++ {
                    v4 += v4
                    v3 = ((v3 + ((0.1 + 0.0) * 1.5)) + (-2.25 + v3))
                    v1 -= (v1 / (((v4 * v4)) | 1))
                    _ = fi15
                }
                if (0 > 1) {
                    for fi16 := 0; fi16 < 6; fi16++ {
                        v17 := GenS1{F0: 0}
                        _ = fi16
                        _ = v17
                    }
                }
            }
        }
        _ = fi11
    }
    v2[""] = (true && false)
    var v18 string = ((("\n\t" + ("" + "a")) + "") + ("" + ((("hello" + (("\n\t" + "\n\t") + "a")) + ("foo bar" + "hello")) + "\n\t")))
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS2{F0: 875}
        _ = fi19
        _ = v20
    }
    v9 = (v9 + v9)
    for fi21 := 0; fi21 < 13; fi21++ {
        v22 := GenS1{F0: 2}
        _ = fi21
        _ = v22
    }
    fmt.Println(v9)
    fmt.Println(v3)
    for fi23 := 0; fi23 < 1; fi23++ {
        delete(v2, ("hello" + v18))
        var v24 int = (fi23 - (fi23 / (((-(fi23))) | 1)))
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
