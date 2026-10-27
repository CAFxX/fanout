package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint }
type GenS2 struct { F0 uint16; F1 uint8; F2 int8 }
func genFn0(p0 int32) uint64 {
    p0 += (^(p0))
    p0 = (p0 | (p0 & p0))
    var v6 int = 42
    _ = p0
    _ = v6
    return 42
}

func main() {
    v1 := []int{130, 32767, 2147483647, 1}
    v2 := []string{(("hello" + "hello") + "foo bar"), "hello", ("foo bar" + "a")}
    v3 := make(map[int]string)
    var v4 int64 = 466
    v5 := make(map[string]int)
    delete(v3, -741)
    fmt.Println(v4)
    v4 += v4
    v4 = (-(v4))
    v4 = v4
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(fi7, fi7)
        var v8 int64 = -9223372036854775808
        _ = fi7
        _ = v8
    }
    fmt.Println(v4)
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            if len(v2) > 0 {
                v2[0] = ("\n\t" + "\n\t")
            }
            var v11 int = int(v4)
            v4 = v4
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        switch "hello" {
        case "foo bar":
            var v13 uint32 = 174
            _ = v13
        default:
            for fi14 := 0; fi14 < 6; fi14++ {
                v4 = (v4 + (v4 & (v4 ^ (^(v4)))))
                var v15 uint = 2
                _ = fi14
                _ = v15
            }
        }
        var v16 uint8 = 255
        delete(v3, int(v4))
        _ = fi12
        _ = v16
    }
    var v17 bool = (65535 != 2147483647)
    v18 := GenS1{F0: (("foo bar" + "\n\t") + "foo bar"), F1: 2, F2: 648}
    switch 1 {
    case 32768:
        v17 = (false || (223 >= 0))
    case 9223372036854775807:
        var v19 int16 = -358
        _ = v19
    case 42:
        v20 := GenS1{F0: "foo bar", F1: 255, F2: 18446744073709551615}
        _ = v20
    default:
        if (!(-1 != -1)) {
            v1 = append(v1, -832)
            for fi21 := 0; fi21 < 13; fi21++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi21
            }
        } else {
            v22 := GenS2{F0: 42, F1: 42, F2: 42}
            v23 := GenS2{F0: 255, F1: 0, F2: 61}
            _ = v22
            _ = v23
        }
    }
    var v24 float32 = float32(1.5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
