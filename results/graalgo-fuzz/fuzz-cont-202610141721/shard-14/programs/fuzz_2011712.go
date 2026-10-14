package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
func genFn0(p0 uint16, p1 uint16, p2 int) uint16 {
    fmt.Println(p2)
    var v7 int8 = 127
    p1 -= p1
    var v8 uint32 = 301
    v9 := GenS1{F0: 61, F1: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]uint32)
    var v3 int = 32768
    var v4 uint8 = 2
    v5 := make(map[uint32]bool)
    var v6 int64 = 9223372036854775807
    _ = len(v5)
    v10 := []int{v3}
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := GenS2{F0: -32768}
        fmt.Println(v6, v1, fi11)
        var v13 uint32 = 0
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 int32 = -470
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: v1}
    v1 -= -231
    for fi17 := 0; fi17 < 19; fi17++ {
        _ = len(v5)
        _ = len(v5)
        var v18 float64 = -2.25
        _ = fi17
        _ = v18
    }
    fmt.Println(v3, v1)
    var v19 int64 = (v6 >> ((508) & 63))
    var v20 int = (^(v3))
    fmt.Println(v1)
    if true {
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 int = (fi21 & fi21)
            var v23 uint16 = 2
            if true {
                for fi24 := 0; fi24 < 11; fi24++ {
                    if true {
                        v1 = (v1 % ((v1) | 1))
                        var v25 int8 = 100
                        var v26 int = (fi21 >> ((374) & 63))
                        var v27 int = (v26 | (v3 + 0))
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        var v28 bool = (127 == v20)
                        var v29 int64 = (v6 - v19)
                        if ((-2.25 / ((1.5 + ((float32(3.14159) * -2.25) - 1e10)) - 0.1)) <= 3.14159) {
                            v30 := GenS3{F0: "hello"}
                            _ = v30
                        } else {
                            var v31 uint16 = (v23 | (v23 ^ (v23 << ((18446744073709551615) & 63))))
                            _ = v31
                        }
                        _ = v28
                        _ = v29
                    }
                    v10 = append(v10, fi24)
                    fmt.Println(v3)
                    _ = fi24
                }
                var v32 float32 = 3.14159
                for fi33 := 0; fi33 < 4; fi33++ {
                    switch (((("" + ("foo bar" + ("" + "hello"))) + "") + "\n\t") + "foo bar") {
                    case "hello":
                        v34 := v2["hello"]
                        _ = v34
                    }
                    _ = fi33
                }
                _ = v32
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        for fi35 := 0; fi35 < 13; fi35++ {
            fmt.Println(v6)
            if ((!true) || false) {
                for fi36 := 0; fi36 < 13; fi36++ {
                    if (v4 == v4) {
                        v20 = fi36
                        v4 += v4
                        v37 := GenS4{F0: (("\n\t" + "\n\t") + "a")}
                        _ = v37
                    }
                    _ = fi36
                }
            }
            var v38 uint32 = 0
            _ = fi35
            _ = v38
        }
        v3 += 4294967295
    } else {
        for fi39 := 0; fi39 < 11; fi39++ {
            v3 = v20
            _ = fi39
        }
        var v40 float32 = 3.14159
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
