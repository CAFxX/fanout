package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint16; F1 string; F2 uint32 }
func genFn0(p0 int32, p1 uint64) string {
    if (102 == 127) {
        v6 := []int{-1, 42}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v6
    } else {
        p1 -= (p1 * p1)
        v7 := []int{32768, 128, 42}
        for fi8 := 0; fi8 < 1; fi8++ {
            switch "\n\t" {
            case "foo bar":
                p0 += (p0 - (p0 << ((18446744073709551615) & 63)))
            }
            p1 -= 1
            if ((!true) && (32767 != 65535)) {
                var v9 string = (("foo bar" + "") + ("a" + "foo bar"))
                fmt.Println(p0, v9, p0)
                _ = v9
            } else {
                v10 := GenS1{F0: 18446744073709551615, F1: false}
                if false {
                    v7 = append(v7, 4294967295)
                    p0 = p0
                    if false {
                        fmt.Println(p0)
                        p0 = p0
                        for fi11 := 0; fi11 < 8; fi11++ {
                            p1 = 18446744073709551615
                            var v12 string = "foo bar"
                            var v13 int16 = 1
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    }
                } else {
                    var v14 uint = 808
                    p0 += (p0 | 2147483647)
                    _ = v14
                }
                _ = v10
            }
            _ = fi8
        }
        _ = v7
    }
    switch "hello" {
    case "\n\t":
        var v15 int64 = 42
        _ = v15
    default:
        if true {
            p1 -= p1
            var v16 uint16 = 32767
            fmt.Println(p1)
            _ = v16
        }
    }
    _ = p0
    _ = p1
    return (("\n\t" + "a") + "")
}

func main() {
    var v1 bool = ((0 > 9223372036854775807) || (2147483647 == 2147483647))
    var v2 int = 1
    var v3 int32 = 2147483647
    var v4 uint64 = 835
    var v5 float64 = 1.5
    if (v1 || v1) {
        if true {
            switch "\n\t" {
            case "\n\t":
                v17 := []int{v2, (v2 ^ 986), (-(v2))}
                _ = v17
            default:
                fmt.Println(v2, v3)
            }
            v18 := GenS1{F0: 2, F1: true}
            _ = v18
        }
    }
    v19 := []int{4294967295, 32767, -315, 128}
    var v20 int = v2
    if len(v19) > 0 {
        v19[0] = (-(v2))
    }
    var v21 bool = (float64(float32(1.5)) == (v5 * (1e10 / v5)))
    var v22 int32 = (v3 - v3)
    v4 -= (^(v4))
    if false {
        switch "\n\t" {
        case "hello":
            fmt.Println(v5)
        }
        fmt.Println(v21)
        if v21 {
            switch ("foo bar" + "hello") {
            case "a":
                var v23 uint = 4294967295
                _ = v23
            default:
                switch "" {
                case "hello":
                    for fi24 := 0; fi24 < 3; fi24++ {
                        var v25 int8 = 100
                        v26 := GenS1{F0: 4294967295, F1: (v4 < v4)}
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                case "":
                    v22 = v3
                default:
                    if len(v19) > 0 {
                        v19[0] = (v20 | v2)
                    }
                }
            }
        } else {
            for fi27 := 0; fi27 < 10; fi27++ {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                _ = fi27
            }
            var v28 int64 = -1
            if v1 {
                v29 := GenS2{F0: 2, F1: "", F2: 42}
                _ = v29
            }
            _ = v28
        }
    }
    if (827 != -32768) {
        var v30 uint32 = 1
        _ = v30
    }
    v20 = 42
    for fi31 := 0; fi31 < 5; fi31++ {
        if v1 {
            var v32 int32 = v22
            fmt.Println(v4, v20, v4)
            _ = v32
        }
        if v1 {
            var v33 int16 = -927
            var v34 int32 = v3
            _ = v33
            _ = v34
        }
        _ = fi31
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        v36 := GenS1{F0: 33, F1: ((-104 != -122) || v1)}
        _ = fi35
        _ = v36
    }
    var v37 int64 = -1
    var v38 uint64 = (v4 % ((v4) | 1))
    v2 -= v20
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
