package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = float64(3.14159)
    var v3 int32 = -2147483648
    var v4 uint32 = 2147483647
    var v5 bool = (((v3 ^ -799) <= v3) || false)
    v3 += (v3 + v3)
    var v6 int = 2147483647
    v2 += v2
    if v5 {
        fmt.Println(v6, v4)
        if (!v5) {
            var v7 int16 = -32768
            _ = v7
        }
        var v8 uint64 = 264
        _ = v8
    } else {
        if v5 {
            fmt.Println(v2, v6, v6)
            if v5 {
                for fi9 := 0; fi9 < 5; fi9++ {
                    v2 = v2
                    switch ("\n\t" + ("foo bar" + ("\n\t" + (("foo bar" + "") + (("\n\t" + "foo bar") + "hello"))))) {
                    case "":
                        var v10 bool = false
                        _ = v10
                    default:
                        v5 = v5
                    }
                    if (0 >= 42) {
                        fmt.Println(v4, v2)
                        var v11 uint8 = 1
                        _ = v11
                    } else {
                        v4 += (^(v4))
                        fmt.Println(fi9, v6, v5)
                    }
                    _ = fi9
                }
                fmt.Println(v5, v5, v2)
            } else {
                if true {
                    var v12 float64 = -2.25
                    for fi13 := 0; fi13 < 19; fi13++ {
                        var v14 uint = 0
                        v12 += v2
                        _ = fi13
                        _ = v14
                    }
                    delete(v1, v4)
                    _ = v12
                } else {
                    v15 := GenS1{F0: 100, F1: -827}
                    switch ("\n\t" + ("" + ((("" + "a") + ("a" + "foo bar")) + "hello"))) {
                    case "hello":
                        var v16 int64 = 9223372036854775807
                        _ = v16
                    case "":
                        v3 += 371
                    case "\n\t":
                        v17 := []int{}
                        _ = v17
                    default:
                        v2 = -2.25
                    }
                    for fi18 := 0; fi18 < 1; fi18++ {
                        v4 = (v4 - 2)
                        _ = fi18
                    }
                    _ = v15
                }
                var v19 uint32 = 2
                for fi20 := 0; fi20 < 19; fi20++ {
                    v1[v4] = (v3 <= -2147483648)
                    fmt.Println(v6, v5, v5)
                    _ = fi20
                }
                _ = v19
            }
            var v21 uint64 = 18446744073709551615
            _ = v21
        } else {
            v22 := GenS1{F0: 127, F1: -159}
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 bool = (!((true || (v5 || v5)) && (-36 <= -68)))
                v5 = v24
                var v25 bool = false
                _ = fi23
                _ = v24
                _ = v25
            }
            v1[(v4 & v4)] = v5
            _ = v22
        }
    }
    var v26 float32 = (0.1 + 1e10)
    v3 = (v3 | -2147483648)
    var v27 int8 = -1
    v3 = v3
    switch 1 {
    case 1:
        var v28 int = v6
        _ = v28
    case 32767:
        for fi29 := 0; fi29 < 19; fi29++ {
            if v5 {
                var v30 uint16 = 692
                _ = v30
            }
            v6 = v6
            var v31 int8 = 42
            _ = fi29
            _ = v31
        }
    }
    v27 += (v27 - (v27 - (v27 + (v27 - 42))))
    for fi32 := 0; fi32 < 9; fi32++ {
        var v33 float64 = 0.0
        var v34 int16 = -238
        v26 = -2.25
        _ = fi32
        _ = v33
        _ = v34
    }
    v26 = ((0.0 / ((v26 + v26) - float32(v2))) + v26)
    var v35 float64 = v2
    switch v6 {
    case 128:
        switch (-(v6)) {
        case 127:
            var v36 string = ""
            _ = v36
        default:
            if v5 {
                v37 := GenS1{F0: v27, F1: -1}
                if false {
                    v1[(^(v4))] = v5
                } else {
                    for fi38 := 0; fi38 < 7; fi38++ {
                        v39 := GenS1{F0: v27, F1: -1}
                        _ = fi38
                        _ = v39
                    }
                }
                _ = v37
            }
        }
    default:
        for fi40 := 0; fi40 < 17; fi40++ {
            var v41 int = (fi40 << ((1) & 63))
            for fi42 := 0; fi42 < 17; fi42++ {
                v5 = (!(!(2147483647 != (v4 >> ((858) & 63)))))
                v41 = 9223372036854775807
                v2 = 1e10
                _ = fi42
            }
            delete(v1, (v4 << ((819) & 63)))
            _ = fi40
            _ = v41
        }
    }
    v35 -= v2
    var v43 bool = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
