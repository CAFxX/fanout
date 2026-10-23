package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 bool; F2 uint32 }
func genFn0(p0 string, p1 bool, p2 string) uint16 {
    v7 := []int{0, 32768, 32768, 9223372036854775807}
    fmt.Println(p1, p1)
    p0 = "hello"
    if p1 {
        switch "a" {
        case "foo bar":
            p1 = (2 == 348)
        default:
            if (((0.0 + (((1.5 - (0.1 - float64(3.14159))) / 1.0) + -2.25)) * 1.5) != ((0.0 + 0.0) * (0.0 - 3.14159))) {
                fmt.Println(p0, p0)
            }
        }
        switch 2147483647 {
        case 127:
            v8 := GenS1{F0: 255}
            _ = v8
        case 32767:
            var v9 int32 = 557
            _ = v9
        case 9223372036854775807:
            for fi10 := 0; fi10 < 2; fi10++ {
                p0 = "\n\t"
                v7 = append(v7, 704)
                for fi11 := 0; fi11 < 7; fi11++ {
                    var v12 uint8 = 233
                    var v13 float32 = float32(0.0)
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi10
            }
        }
        _ = len(v7)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    var v3 float64 = (1.5 * (0.0 * float64((0.0 + (0.1 / (3.14159 - (0.0 * (3.14159 / 1.0))))))))
    v4 := make(map[string]uint32)
    var v5 uint8 = 0
    var v6 string = "\n\t"
    for fi14 := 0; fi14 < 3; fi14++ {
        if true {
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 uint16 = genFn0(v6, false, (v6 + ""))
                v17 := []int{42, int(v5), fi14}
                v1[v6] = uint32(v5)
                v16 += genFn0("", false, v6)
                _ = fi15
                _ = v16
                _ = v17
            }
            fmt.Println(v3, v3, v3)
        } else {
            var v18 string = (v6 + v6)
            v6 = ""
            _ = v18
        }
        v5 -= v5
        v3 += v3
        _ = fi14
    }
    v19 := GenS1{F0: 42}
    for fi20 := 0; fi20 < 15; fi20++ {
        for fi21 := 0; fi21 < 10; fi21++ {
            if false {
                v22 := []int{(fi20 ^ (fi20 & 396)), (fi21 ^ -9223372036854775808)}
                var v23 int = fi20
                v22 = append(v22, (fi21 >> ((1) & 63)))
                _ = v22
                _ = v23
            }
            switch v6 {
            case "a":
                var v24 int32 = -2147483648
                _ = v24
            case "":
                v25 := GenS2{F0: 32767, F1: (-9223372036854775808 == -1), F2: 272}
                _ = v25
            case "foo bar":
                fmt.Println(v5)
            default:
                var v26 uint16 = uint16(v5)
                _ = v26
            }
            if false {
                switch "" {
                case "hello":
                    v3 += v3
                case "foo bar":
                    v27 := v2[793]
                    _ = v27
                case "a":
                    var v28 uint8 = 128
                    _ = v28
                default:
                    v5 += (^(v5))
                }
                switch (-(fi21)) {
                case -38:
                    v29 := GenS1{F0: 65535}
                    _ = v29
                case -756:
                    var v30 float32 = 1.5
                    _ = v30
                }
            }
            _ = fi21
        }
        v6 = ((v6 + "foo bar") + ((v6 + v6) + v6))
        _ = fi20
    }
    var v31 uint8 = 157
    var v32 uint8 = 127
    v33 := v2[2147483647]
    var v34 int8 = 0
    var v35 uint8 = 62
    var v36 int16 = int16(v34)
    fmt.Println(v31, v32)
    if true {
        var v37 int8 = v34
        v32 -= 127
        v36 = v36
        _ = v37
    }
    v33 -= (v33 & (v33 * 4294967295))
    var v38 uint16 = 341
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
