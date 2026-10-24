package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int; F1 int }
func genFn0() bool {
    v4 := []int{-1, 1, 32768, 127}
    var v5 uint64 = 929
    fmt.Println(v5)
    _ = v4
    _ = v5
    return (("\n\t" + "foo bar") <= "a")
}
func genFn1(p0 int64) uint {
    p0 = p0
    if false {
        fmt.Println(p0)
        if ("a" > ("a" + ("\n\t" + "hello"))) {
            v6 := GenS1{F0: 4294967295, F1: 248}
            _ = v6
        } else {
            var v7 int64 = (p0 >> ((130) & 63))
            _ = v7
        }
    }
    _ = p0
    return 82
}

func main() {
    v1 := []uint32{42, 0, 42, 2147483647}
    v2 := make(map[string]int)
    var v3 float32 = 1.5
    var v8 int64 = 815
    switch "" {
    case "foo bar":
        if true {
            if (168 < 2) {
                if (9223372036854775807 >= 32767) {
                    if false {
                        v9 := GenS1{F0: 2, F1: 123}
                        v3 -= v3
                        fmt.Println(v8)
                        _ = v9
                    }
                } else {
                    fmt.Println(v3, v3)
                }
                v10 := GenS1{F0: 804, F1: 0}
                v8 = 1
                _ = v10
            } else {
                v8 += v8
            }
            var v11 uint32 = 65535
            v11 = v11
            _ = v11
        } else {
            for fi12 := 0; fi12 < 17; fi12++ {
                for fi13 := 0; fi13 < 15; fi13++ {
                    v2["\n\t"] = fi12
                    _ = fi13
                }
                if (!true) {
                    v14 := GenS1{F0: 1, F1: 216}
                    if false {
                        for fi15 := 0; fi15 < 2; fi15++ {
                            var v16 uint16 = 65535
                            _ = fi15
                            _ = v16
                        }
                        v8 = v8
                    } else {
                        if (false || (true || true)) {
                            for fi17 := 0; fi17 < 20; fi17++ {
                                v8 += (v8 | (v8 & v8))
                                fmt.Println(v3)
                                _ = fi17
                            }
                            var v18 uint8 = 128
                            v8 = v8
                            _ = v18
                        } else {
                            _ = len(v1)
                        }
                        v3 += (((float32(1e10) / ((1e10 * 1e10) * 1e10)) + (float32((-2.25 - (1.5 - 1e10))) - -2.25)) / (-2.25 * 3.14159))
                    }
                    _ = v14
                } else {
                    v3 = v3
                    v19 := GenS2{F0: (fi12 | fi12), F1: -9223372036854775808}
                    var v20 int32 = 42
                    _ = v19
                    _ = v20
                }
                _ = fi12
            }
            fmt.Println(v8, v3)
            var v21 uint8 = 255
            _ = v21
        }
    case "hello":
        v8 -= (-(v8))
    default:
        v3 = v3
    }
    var v22 int8 = 42
    if (false || (false || false)) {
        for fi23 := 0; fi23 < 2; fi23++ {
            _ = len(v2)
            for fi24 := 0; fi24 < 12; fi24++ {
                if false {
                    var v25 uint = 65535
                    _ = v25
                } else {
                    switch fi23 {
                    case 472:
                        v26 := GenS2{F0: fi24, F1: (fi24 % ((fi23) | 1))}
                        _ = v26
                    case 2147483647:
                        var v27 int8 = v22
                        _ = v27
                    }
                }
                v8 -= v8
                _ = fi24
            }
            _ = fi23
        }
        switch ("a" + "a") {
        case "":
            fmt.Println(v22, v22)
        case "foo bar":
            var v28 string = "foo bar"
            _ = v28
        }
        if (!true) {
            v3 += v3
            v22 += -75
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v3 = ((0.1 * 0.0) * float32(3.14159))
        }
    }
    fmt.Println(v3, v8, v22)
    var v29 bool = (!false)
    switch "a" {
    case "":
        switch 0 {
        case 4294967295:
            var v30 bool = v29
            _ = v30
        case -758:
            v1 = append(v1, 2147483648)
        default:
            v29 = v29
        }
    case "foo bar":
        v3 -= -2.25
    case "a":
        switch "" {
        case "":
            var v31 bool = v29
            _ = v31
        case "\n\t":
            var v32 uint8 = 127
            _ = v32
        default:
            v33 := GenS1{F0: 630, F1: 127}
            _ = v33
        }
    }
    switch "hello" {
    case "hello":
        for fi34 := 0; fi34 < 15; fi34++ {
            v3 += (3.14159 * v3)
            _ = fi34
        }
    case "\n\t":
        _ = len(v1)
    default:
        var v35 int16 = 1
        _ = v35
    }
    var v36 uint16 = 65535
    if len(v1) > 0 {
        v1[0] = 65535
    }
    var v37 uint8 = 31
    v22 += (v22 & 127)
    v37 = (v37 - (v37 + 1))
    fmt.Println(v22, v8)
    if v29 {
        v3 = 1e10
        v38 := GenS1{F0: 254, F1: 2}
        v39 := v2[""]
        _ = v38
        _ = v39
    } else {
        fmt.Println(v8)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v8
}
