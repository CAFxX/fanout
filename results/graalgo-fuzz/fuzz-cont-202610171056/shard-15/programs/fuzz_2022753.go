package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS2 struct { F0 bool; F1 int8; F2 string }
func genFn0(p0 int32, p1 int32) uint16 {
    p1 += p1
    for fi7 := 0; fi7 < 5; fi7++ {
        if (-2.25 > float64(1.5)) {
            var v8 uint = 65535
            _ = v8
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        p1 -= -373
        p1 = p1
        _ = fi9
    }
    _ = p0
    _ = p1
    return 469
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (false || true)
    v3 := make(map[uint32]bool)
    var v4 float32 = (3.14159 + (0.1 / -2.25))
    var v5 int = 1
    v6 := make(map[int]uint32)
    v2 = v2
    v2 = (65535 > 65535)
    v5 = v5
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            fmt.Println(v5, fi11, fi11)
            if ("hello" == "foo bar") {
                fmt.Println(v4)
                switch (fi11 % ((-1) | 1)) {
                case -900:
                    fmt.Println(v4, fi11, v4)
                case 909:
                    v12 := GenS1{F0: 42, F1: 530, F2: 0}
                    _ = v12
                }
                switch "a" {
                case "":
                    fmt.Println(v5, v4)
                case "hello":
                    fmt.Println(v5)
                default:
                    fmt.Println(v5, fi10)
                }
            }
            _ = fi11
        }
        _ = fi10
    }
    v13 := []int{v5, v5, -9223372036854775808, 0}
    var v14 uint32 = 1
    v4 -= v4
    if (uint8(v14) <= uint8(v14)) {
        v13 = append(v13, (v5 ^ int(v14)))
        v5 -= v5
        var v15 int32 = -1
        _ = v15
    } else {
        v2 = ("" == "hello")
    }
    fmt.Println(v4)
    v3[v14] = v2
    if (((v4 < v4) && true) || v2) {
        v16 := GenS2{F0: true, F1: 100, F2: "foo bar"}
        v3[(^(v14))] = (0 > -1)
        v17 := v1[("a" + "")]
        _ = v16
        _ = v17
    } else {
        switch (((((("\n\t" + "\n\t") + ((("hello" + "hello") + ("foo bar" + ("foo bar" + "foo bar"))) + "foo bar")) + "foo bar") + "a") + "foo bar") + "hello") {
        case "hello":
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        case "foo bar":
            for fi18 := 0; fi18 < 15; fi18++ {
                fmt.Println(v5, v14, v4)
                for fi19 := 0; fi19 < 2; fi19++ {
                    _ = len(v13)
                    _ = fi19
                }
                var v20 int8 = 100
                _ = fi18
                _ = v20
            }
        case "":
            fmt.Println(v2, v4, v5)
        }
        var v21 bool = (!v2)
        fmt.Println(v21)
        _ = v21
    }
    v13 = append(v13, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
