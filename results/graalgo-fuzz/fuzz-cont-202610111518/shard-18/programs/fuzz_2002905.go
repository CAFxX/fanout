package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 bool; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint8, p2 string) int16 {
    var v7 string = p2
    var v8 uint = 2
    v8 = (v8 & 42)
    var v9 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -746
}

func main() {
    v1 := []int{508}
    v2 := []string{"foo bar", ((("" + (((("foo bar" + ("a" + "")) + ("foo bar" + "hello")) + ("foo bar" + (("" + ("a" + "hello")) + "foo bar"))) + "hello")) + "") + ("\n\t" + "foo bar")), "a", ((((("a" + "hello") + "a") + "foo bar") + ("hello" + "foo bar")) + "foo bar")}
    var v3 int16 = 872
    var v4 uint = 255
    v5 := []uint32{}
    var v6 int32 = 2147483647
    v3 -= (v3 / ((genFn0(2, 127, (("" + "a") + "hello"))) | 1))
    var v10 float32 = float32((0.1 - -2.25))
    fmt.Println(v10)
    switch 9223372036854775807 {
    case -567:
        if true {
            if ((255 <= 255) && (488 < -594)) {
                var v11 uint8 = 1
                if true {
                    var v12 uint16 = 42
                    for fi13 := 0; fi13 < 8; fi13++ {
                        v3 = v3
                        for fi14 := 0; fi14 < 16; fi14++ {
                            v10 += 0.1
                            var v15 int = fi13
                            _ = fi14
                            _ = v15
                        }
                        var v16 uint8 = v11
                        _ = fi13
                        _ = v16
                    }
                    _ = v12
                }
                v1 = append(v1, -2147483648)
                _ = v11
            }
            var v17 uint64 = 18446744073709551615
            _ = v17
        }
    case 128:
        var v18 uint16 = 2
        _ = v18
    }
    var v19 int64 = -961
    var v20 float32 = -2.25
    _ = len(v5)
    var v21 uint64 = 2
    if ((true || (((v19 << ((v4) & 63)) < v19) || true)) && (410 == 32767)) {
        for fi22 := 0; fi22 < 15; fi22++ {
            var v23 uint = v4
            for fi24 := 0; fi24 < 2; fi24++ {
                v25 := GenS1{F0: true}
                v26 := GenS1{F0: false}
                fmt.Println(v3, v3)
                _ = fi24
                _ = v25
                _ = v26
            }
            fmt.Println(v6, v23)
            _ = fi22
            _ = v23
        }
        switch -1 {
        case -339:
            v27 := GenS1{F0: (2 < 0)}
            _ = v27
        case 1:
            v19 = v19
        case 127:
            if (false || ((569 > 549) && true)) {
                v10 = v20
                var v28 int64 = v19
                v6 += v6
                _ = v28
            } else {
                if true {
                    v29 := GenS2{F0: v6, F1: (!false), F2: 1}
                    _ = v29
                } else {
                    if len(v1) > 0 {
                        v1[0] = int(v21)
                    }
                }
            }
        default:
            switch 4294967295 {
            case 127:
                v6 = v6
            case 79:
                v20 = (v20 / (v10 + v10))
            case 128:
                var v30 int64 = v19
                _ = v30
            default:
                v31 := GenS1{F0: false}
                _ = v31
            }
        }
        if false {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        } else {
            v3 = (v3 ^ 1)
            v32 := GenS1{F0: false}
            _ = v32
        }
    }
    fmt.Println(v4, v10)
    v10 -= (float32(1e10) + v10)
    if true {
        v4 += uint(v21)
        for fi33 := 0; fi33 < 5; fi33++ {
            for fi34 := 0; fi34 < 18; fi34++ {
                v35 := GenS1{F0: false}
                v36 := GenS3{F0: v19}
                _ = fi34
                _ = v35
                _ = v36
            }
            v37 := GenS3{F0: v19}
            _ = fi33
            _ = v37
        }
        if false {
            v38 := GenS4{F0: v19}
            _ = len(v2)
            _ = v38
        } else {
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
            if (2147483648 == 65535) {
                var v39 int64 = v19
                v40 := GenS3{F0: v39}
                v5 = append(v5, 2147483647)
                _ = v39
                _ = v40
            } else {
                v10 -= 3.14159
                if false {
                    if (false || (!true)) {
                        fmt.Println(v21, v6)
                        var v41 uint32 = 943
                        v3 = v3
                        _ = v41
                    }
                    v21 += 18446744073709551615
                    v4 = 1
                }
            }
        }
    } else {
        v42 := GenS1{F0: (true || (genFn0(42, 209, ("\n\t" + "")) == v3))}
        _ = v42
    }
    switch -132 {
    case 1:
        fmt.Println(v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
