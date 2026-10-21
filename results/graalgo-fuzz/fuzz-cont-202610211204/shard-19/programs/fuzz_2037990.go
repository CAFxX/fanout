package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int32; F2 uint }
func genFn0(p0 uint16, p1 uint32, p2 uint16) uint {
    fmt.Println(p0)
    var v7 float32 = (3.14159 * (-2.25 / -2.25))
    switch 160 {
    case -1:
        v8 := []int{-2147483648, 127}
        _ = v8
    default:
        var v9 float32 = v7
        _ = v9
    }
    var v10 int32 = 349
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 670
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]uint32)
    var v3 int64 = 549
    v4 := []int{9223372036854775807, 128}
    var v5 int = -2147483648
    v6 := make(map[int]string)
    v3 += v3
    v5 += (-(v5))
    delete(v2, 424)
    switch ("" + "\n\t") {
    case "foo bar":
        v3 = v3
    case "":
        for fi11 := 0; fi11 < 12; fi11++ {
            _ = len(v6)
            v3 = (v3 - (^(v3)))
            _ = fi11
        }
    case "hello":
        for fi12 := 0; fi12 < 10; fi12++ {
            v3 = (v3 ^ 957)
            var v13 bool = false
            _ = len(v1)
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        v5 -= fi14
        _ = fi14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 bool = (4294967295 <= 65535)
        _ = fi15
        _ = v16
    }
    if (!(0 < 1000)) {
        if false {
            var v17 bool = true
            var v18 string = (("\n\t" + "hello") + "hello")
            _ = v17
            _ = v18
        } else {
            v5 -= v5
            v5 += (-(v5))
        }
        switch v5 {
        case -2147483648:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        v3 = 1
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        v20 := GenS1{F0: "hello"}
        _ = fi19
        _ = v20
    }
    if len(v4) > 0 {
        v4[0] = (^(v5))
    }
    v5 += -655
    switch (v5 % (((-(v5))) | 1)) {
    case -9223372036854775808:
        var v21 uint64 = 4294967295
        _ = v21
    case 583:
        switch 2147483647 {
        case 220:
            v22 := GenS1{F0: "hello"}
            _ = v22
        case -854:
            var v23 uint16 = 0
            _ = v23
        case 9223372036854775807:
            fmt.Println(v3, v5)
        default:
            v5 -= v5
        }
    }
    switch -179 {
    case 9223372036854775807:
        v24 := GenS2{F0: 65535, F1: 551, F2: genFn0(2, 1, 902)}
        _ = v24
    }
    var v25 bool = false
    fmt.Println(v3, v5)
    if ((0.0 * float32(1e10)) <= (((0.1 - -2.25) - 0.1) / (((float32(((((1.5 * 3.14159) + 1e10) + (-2.25 * ((float64(3.14159) - 0.0) / (1.5 + 0.1)))) + 1.5)) - 1.5) / 1.0) / 1e10))) {
        fmt.Println(v25, v3)
        if (1 <= 0) {
            var v26 bool = (497 != 620)
            var v27 int8 = 42
            v25 = (!true)
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 12; fi28++ {
                for fi29 := 0; fi29 < 6; fi29++ {
                    if (127 == 127) {
                        v5 = (v5 / (((v5 + fi29)) | 1))
                    }
                    _ = fi29
                }
                fmt.Println(v5, v3, v5)
                _ = fi28
            }
        }
        if ("" <= (("" + "") + "foo bar")) {
            fmt.Println(v5)
        }
    }
    var v30 int16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
