package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 float64 = 3.14159
    var v4 int16 = 1000
    v5 := []uint32{65535, 2, 0}
    v1["a"] = ("\n\t" + ("hello" + (("\n\t" + "hello") + "hello")))
    if (((v3 * v3) / ((v3 + (float64((-2.25 / 1.0)) * v3)) - v3)) < (v3 / float64(0.1))) {
        var v6 int16 = 0
        _ = v6
    } else {
        v7 := GenS1{F0: (-(v4)), F1: 5}
        var v8 uint = 303
        _ = v7
        _ = v8
    }
    if false {
        v3 -= (0.0 + (v3 + v3))
        v3 = 3.14159
    } else {
        v3 = v3
        if false {
            var v9 uint64 = 2
            fmt.Println(v9, v4, v4)
            _ = v9
        }
    }
    var v10 uint16 = 65535
    if (v3 == v3) {
        switch 2147483647 {
        case 42:
            var v11 bool = false
            _ = v11
        }
        if true {
            v10 = v10
        } else {
            v12 := GenS2{F0: -119}
            var v13 int64 = -263
            _ = v12
            _ = v13
        }
    }
    v4 -= (v4 & v4)
    if (1e10 < (v3 - v3)) {
        var v14 uint16 = (v10 >> ((18446744073709551615) & 63))
        var v15 string = "hello"
        _ = v14
        _ = v15
    }
    _ = len(v5)
    switch 127 {
    case -1:
        v16 := GenS2{F0: 0}
        _ = v16
    case 490:
        fmt.Println(v10, v3)
    case 627:
        _ = len(v1)
    }
    v17 := GenS1{F0: (v4 - v4), F1: 0}
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 uint32 = 1
        v2[("" + "\n\t")] = ""
        switch fi18 {
        case 2147483647:
            v3 -= (float64(1e10) + v3)
        }
        _ = fi18
        _ = v19
    }
    var v20 uint8 = 228
    if true {
        v1["a"] = ("a" + (("\n\t" + ("a" + "")) + "a"))
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        switch (((("foo bar" + ("hello" + ("" + "a"))) + "\n\t") + "\n\t") + "\n\t") {
        case "foo bar":
            if (true || false) {
                v3 = (v3 - 3.14159)
                fmt.Println(v3, v20)
                for fi21 := 0; fi21 < 20; fi21++ {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        if len(v5) > 0 {
                            v5[0] = uint32(v10)
                        }
                        _ = fi22
                    }
                    _ = fi21
                }
            }
        default:
            v20 += 101
        }
    }
    delete(v1, ("foo bar" + "foo bar"))
    v20 -= 118
    v3 = (v3 - (1e10 / 0.1))
    v23 := GenS3{F0: -1}
    for fi24 := 0; fi24 < 11; fi24++ {
        if (true && false) {
            fmt.Println(fi24)
            var v25 int32 = int32(v10)
            if ((v10 << ((971) & 63)) <= (v10 - 42)) {
                v4 -= -795
                var v26 float32 = 1.5
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v26
            }
            _ = v25
        }
        v27 := GenS4{F0: (v20 * (v20 >> ((578) & 63))), F1: 42}
        v10 = v10
        _ = fi24
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
