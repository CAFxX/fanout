package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 int32) uint {
    fmt.Println(p0)
    var v6 uint32 = 4294967295
    _ = p0
    _ = v6
    return 65535
}

func main() {
    v1 := []string{"", "foo bar", "\n\t", ("a" + "")}
    var v2 string = "hello"
    var v3 uint32 = 390
    var v4 int16 = -32768
    v5 := make(map[uint32]int)
    for fi7 := 0; fi7 < 5; fi7++ {
        v2 = ("a" + (v2 + "\n\t"))
        v8 := GenS1{F0: -2147483648}
        for fi9 := 0; fi9 < 1; fi9++ {
            v3 = (v3 / ((4294967295) | 1))
            var v10 float32 = -2.25
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v3, v2)
    v11 := GenS1{F0: 1}
    if false {
        switch 1 {
        case 4294967295:
            switch -9223372036854775808 {
            case 127:
                v3 -= v3
            case -2147483648:
                v4 -= 42
            case 32767:
                fmt.Println(v2, v4, v2)
            }
        case -740:
            switch v2 {
            case "foo bar":
                for fi12 := 0; fi12 < 8; fi12++ {
                    var v13 float64 = (-2.25 * ((1e10 - 1.5) - (((-2.25 + (0.1 / 1.5)) + 3.14159) * (-2.25 - (-2.25 * (((1.5 + (float64(1e10) + (float64(1e10) - (1.5 - 1e10)))) / 1.5) - float64(float32(0.1))))))))
                    var v14 uint8 = 127
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case "\n\t":
                var v15 int16 = v4
                _ = v15
            }
        default:
            v2 = "a"
        }
    } else {
        v4 = 32767
    }
    var v16 uint64 = 71
    v3 = (v3 * (v3 ^ v3))
    var v17 string = v2
    var v18 int = -2147483648
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 int64 = 9223372036854775807
        v17 = (("a" + "foo bar") + v17)
        var v21 uint64 = v16
        _ = fi19
        _ = v20
        _ = v21
    }
    v3 += 830
    _ = len(v1)
    v5[v3] = 127
    var v22 float64 = (0.1 - 3.14159)
    var v23 int8 = 1
    for fi24 := 0; fi24 < 7; fi24++ {
        if false {
            for fi25 := 0; fi25 < 12; fi25++ {
                fmt.Println(fi24)
                var v26 uint = 0
                _ = fi25
                _ = v26
            }
        } else {
            delete(v5, v3)
            if (-981 == 0) {
                if true {
                    if false {
                        delete(v5, v3)
                    }
                    if (1 > 2) {
                        var v27 int64 = 0
                        v3 = v3
                        var v28 uint64 = 2
                        _ = v27
                        _ = v28
                    }
                    var v29 string = v17
                    _ = v29
                }
            } else {
                var v30 uint = 18446744073709551615
                for fi31 := 0; fi31 < 6; fi31++ {
                    fmt.Println(v16, v30, v23)
                    switch (^(v18)) {
                    case 281:
                        v1 = append(v1, "\n\t")
                        var v32 uint = (v30 * v30)
                        _ = v32
                    case 787:
                        var v33 uint32 = (v3 | v3)
                        _ = v33
                    default:
                        v2 = (v2 + ("" + (("a" + (((v2 + ((v2 + ("" + (v2 + ("foo bar" + ("" + (v17 + "\n\t")))))) + v17)) + v2) + (("foo bar" + v2) + v17))) + v2)))
                    }
                    fmt.Println(v4, v3, v22)
                    _ = fi31
                }
                var v34 int64 = -1
                _ = v30
                _ = v34
            }
            v35 := GenS2{F0: v4, F1: 0}
            _ = v35
        }
        var v36 int32 = 42
        _ = fi24
        _ = v36
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = "hello"
    if true {
        switch (v18 + (v18 + (v18 ^ v18))) {
        case 4294967295:
            if (int32(v18) == 0) {
                var v37 uint32 = (v3 | v3)
                v38 := GenS2{F0: v4, F1: 948}
                fmt.Println(v18, v37, v17)
                _ = v37
                _ = v38
            } else {
                v39 := v5[(v3 | 65535)]
                _ = v39
            }
        default:
            v1 = append(v1, (v2 + ((v2 + "hello") + "\n\t")))
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
