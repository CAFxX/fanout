package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint32; F1 int32; F2 int64 }

func main() {
    var v1 uint8 = 1
    var v2 float32 = 0.1
    var v3 int32 = -31
    var v4 string = ("hello" + "foo bar")
    v5 := make(map[string]string)
    var v6 int = -171
    v2 = ((0.0 - 1e10) / (-2.25 / v2))
    v2 = 1.5
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(v6, fi7, v1)
        var v8 int32 = (v3 >> ((uint(v1)) & 63))
        v6 += fi7
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 255, F1: 1}
    if true {
        delete(v5, ((((v4 + ("\n\t" + v4)) + v4) + "") + ""))
        v4 = "a"
    } else {
        var v10 float32 = v2
        switch (v4 + "foo bar") {
        case "a":
            v1 += (v1 & 255)
        case "foo bar":
            v2 = (0.1 * v10)
        default:
            v4 = "\n\t"
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            v1 = 255
            _ = fi11
        }
        _ = v10
    }
    v12 := GenS1{F0: 65535, F1: 0}
    if true {
        var v13 int32 = 1
        v4 = (v4 + "")
        v14 := v5[(v4 + ("foo bar" + (v4 + (v4 + v4))))]
        _ = v13
        _ = v14
    }
    fmt.Println(v1, v6)
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := GenS2{F0: 4294967295, F1: -293, F2: -1}
        if true {
            var v17 uint32 = 65535
            if ((v4 + ((((v4 + v4) + (v4 + v4)) + ("" + ("\n\t" + v4))) + "hello")) != "hello") {
                var v18 uint32 = v17
                if (255 < 397) {
                    switch (v6 / ((fi15) | 1)) {
                    case 42:
                        if false {
                            fmt.Println(v4)
                        } else {
                            var v19 int8 = 42
                            _ = v19
                        }
                    case -9223372036854775808:
                        if false {
                            if (false || ((v3 != (v3 | 42)) || false)) {
                                delete(v5, "foo bar")
                            } else {
                                v20 := GenS2{F0: v17, F1: (v3 << ((2) & 63)), F2: 1}
                                _ = v20
                            }
                            v21 := []int{127}
                            v17 = (^(v18))
                            _ = v21
                        }
                    case -64:
                        var v22 uint64 = 945
                        _ = v22
                    }
                    v23 := []int{(v6 >> ((4294967295) & 63)), (v6 * (v6 - -941))}
                    var v24 uint8 = (v1 >> ((18446744073709551615) & 63))
                    _ = v23
                    _ = v24
                }
                var v25 string = v4
                _ = v18
                _ = v25
            } else {
                var v26 uint64 = 18446744073709551615
                _ = v26
            }
            for fi27 := 0; fi27 < 7; fi27++ {
                var v28 int = v6
                _ = fi27
                _ = v28
            }
            _ = v17
        } else {
            for fi29 := 0; fi29 < 9; fi29++ {
                var v30 uint16 = 42
                v6 = v6
                var v31 uint32 = 65535
                _ = fi29
                _ = v30
                _ = v31
            }
            delete(v5, "hello")
        }
        _ = fi15
        _ = v16
    }
    switch v6 {
    case 2147483647:
        v6 -= (v6 + int(v3))
    case 128:
        var v32 int16 = 1000
        _ = v32
    case 4294967295:
        v2 -= -2.25
    default:
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 bool = true
            v34 = ((100 >= 1) && false)
            fmt.Println(v3)
            _ = fi33
            _ = v34
        }
    }
    var v35 uint = 981
    v1 -= 127
    fmt.Println(v2, v35)
    var v36 string = v4
    v3 = (v3 << ((v35) & 63))
    v37 := []int{741, -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
