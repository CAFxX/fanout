package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint }
func genFn0(p0 uint32, p1 int16) uint64 {
    if ((p1 > p1) && true) {
        fmt.Println(p1)
        var v5 int64 = 42
        _ = v5
    } else {
        for fi6 := 0; fi6 < 6; fi6++ {
            p1 = -32768
            p1 += p1
            fmt.Println(p1, p1, p1)
            _ = fi6
        }
    }
    v7 := GenS1{F0: 616, F1: 42}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    var v1 bool = true
    v2 := []string{}
    var v3 int8 = 0
    v4 := make(map[string]uint32)
    var v8 int8 = v3
    switch -1 {
    case -11:
        v8 += -1
    case 564:
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    case -76:
        var v9 uint8 = 127
        _ = v9
    }
    delete(v4, ("a" + ("hello" + ("a" + "foo bar"))))
    switch 812 {
    case -9223372036854775808:
        v8 += (v3 << ((0) & 63))
    case 127:
        if (1 < 18446744073709551615) {
            v8 = v3
            fmt.Println(v8, v8)
        } else {
            var v10 float64 = (((-2.25 / 3.14159) / float64(float32(-2.25))) * (1e10 * (0.1 / float64(1.5))))
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            _ = v10
        }
    case -1:
        for fi11 := 0; fi11 < 3; fi11++ {
            if false {
                switch "hello" {
                case "foo bar":
                    if v1 {
                        for fi12 := 0; fi12 < 8; fi12++ {
                            _ = len(v2)
                            v13 := GenS1{F0: 65535, F1: 1}
                            _ = fi12
                            _ = v13
                        }
                        v8 += (v3 ^ (v8 * v8))
                    }
                }
            } else {
                v4[(("foo bar" + "") + "a")] = 1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v14 := GenS1{F0: 2, F1: 1}
            _ = fi11
            _ = v14
        }
    default:
        fmt.Println(v1)
    }
    var v15 uint = 255
    fmt.Println(v8, v3)
    v1 = false
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 float32 = ((-2.25 + 3.14159) - float32(0.1))
        var v18 int32 = 0
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v8 += 0
        _ = fi19
    }
    if ((v15 - (v15 << ((0) & 63))) == v15) {
        v20 := GenS2{F0: 1, F1: (v15 % ((v15) | 1))}
        _ = v20
    } else {
        if false {
            for fi21 := 0; fi21 < 7; fi21++ {
                v22 := GenS1{F0: 1, F1: 18446744073709551615}
                var v23 uint16 = 0
                _ = fi21
                _ = v22
                _ = v23
            }
        } else {
            var v24 uint8 = 255
            switch 42 {
            case 9223372036854775807:
                var v25 int = 252
                _ = v25
            case 42:
                switch "" {
                case "":
                    var v26 bool = v1
                    _ = v26
                default:
                    v27 := GenS1{F0: 2147483647, F1: 0}
                    _ = v27
                }
            case 4294967295:
                var v28 uint16 = 1
                _ = v28
            }
            fmt.Println(v8)
            _ = v24
        }
        for fi29 := 0; fi29 < 11; fi29++ {
            var v30 bool = (((0 == -686) || false) && v1)
            _ = fi29
            _ = v30
        }
        var v31 float32 = 0.0
        _ = v31
    }
    var v32 int64 = 42
    switch (("foo bar" + "") + "\n\t") {
    case "a":
        if (true && (true && (v1 || true))) {
            v8 += v8
            delete(v4, "hello")
        }
    case "foo bar":
        var v33 bool = true
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v8
}
