package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int8 }
type GenS2 struct { F0 uint64; F1 int16 }
func genFn0(p0 bool, p1 int) int32 {
    p0 = (float64(1.5) != 1e10)
    if (!p0) {
        if (42 == 42) {
            if ((("foo bar" + "") + "a") >= ("hello" + (("a" + "hello") + "\n\t"))) {
                var v5 int64 = -9223372036854775808
                var v6 float64 = (3.14159 + 1.5)
                for fi7 := 0; fi7 < 14; fi7++ {
                    v8 := GenS1{F0: "", F1: true, F2: 127}
                    _ = fi7
                    _ = v8
                }
                _ = v5
                _ = v6
            }
            switch "a" {
            case "foo bar":
                switch p1 {
                case -125:
                    for fi9 := 0; fi9 < 18; fi9++ {
                        for fi10 := 0; fi10 < 18; fi10++ {
                            p1 += 127
                            v11 := GenS1{F0: "a", F1: true, F2: -30}
                            var v12 int32 = 332
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                        var v13 int32 = 1
                        _ = fi9
                        _ = v13
                    }
                default:
                    v14 := []int{(p1 & p1), p1, p1}
                    _ = v14
                }
            default:
                v15 := GenS2{F0: 4294967295, F1: -32768}
                _ = v15
            }
        }
        fmt.Println(p0, p0)
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]int)
    var v3 float32 = 1e10
    var v4 uint32 = 659
    var v16 int32 = genFn0(false, 127)
    var v17 uint32 = (v4 + v4)
    switch 32767 {
    case -9223372036854775808:
        var v18 uint32 = v17
        _ = v18
    }
    var v19 bool = (true || true)
    v20 := v2[32768]
    v16 = (v16 >> ((811) & 63))
    for fi21 := 0; fi21 < 19; fi21++ {
        switch 128 {
        case 1:
            var v22 uint32 = (v17 - (v4 / ((1) | 1)))
            _ = v22
        default:
            var v23 int16 = 1000
            _ = v23
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := v2[(fi21 >> ((608) & 63))]
            var v26 float32 = v3
            var v27 string = ""
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        _ = fi21
    }
    var v28 bool = v19
    var v29 uint64 = 0
    v28 = false
    var v30 uint32 = 2
    v2[int(v29)] = v20
    if true {
        var v31 int16 = -32768
        if (!true) {
            var v32 uint16 = 131
            _ = v32
        }
        _ = v31
    } else {
        var v33 int32 = 0
        for fi34 := 0; fi34 < 17; fi34++ {
            fmt.Println(v16)
            for fi35 := 0; fi35 < 4; fi35++ {
                v19 = (((v33 != genFn0((!(((-2.25 + (v3 / v3)) <= ((v3 + ((v3 * 0.1) / v3)) + ((0.0 / v3) * v3))) && v19)), (fi34 + (v20 & v20)))) && (v28 || false)) || (((v3 - v3) <= ((v3 * 1e10) * (v3 * (1.5 - (float32(-2.25) / 3.14159))))) && true))
                v20 += (fi34 >> ((4294967295) & 63))
                _ = fi35
            }
            v36 := GenS1{F0: "\n\t", F1: true, F2: -128}
            _ = fi34
            _ = v36
        }
        _ = v33
    }
    var v37 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v4
}
