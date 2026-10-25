package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int }
type GenS2 struct { F0 int64; F1 uint8; F2 uint8 }
func genFn0(p0 uint16, p1 uint, p2 string) string {
    p1 -= p1
    if (-9223372036854775808 >= 730) {
        if false {
            v7 := GenS1{F0: p0, F1: "a", F2: -9223372036854775808}
            _ = v7
        } else {
            fmt.Println(p2, p2, p1)
            p0 -= p0
        }
        p1 = (p1 & 42)
        p1 = p1
    }
    var v8 uint = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 792
    var v4 uint8 = 42
    v5 := []uint32{}
    var v6 uint32 = uint32(v3)
    v6 = v6
    switch 4294967295 {
    case 32768:
        var v9 int8 = 100
        _ = v9
    default:
        v10 := GenS2{F0: 441, F1: v4, F2: (v4 * 255)}
        _ = v10
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            v4 += 128
            var v13 float64 = 1e10
            _ = len(v5)
            if (v13 >= v13) {
                fmt.Println(fi12)
                v14 := v1[("" + "\n\t")]
                v15 := GenS1{F0: v3, F1: "hello", F2: fi12}
                _ = v14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
        _ = len(v2)
        if (!(!true)) {
            switch "hello" {
            case "hello":
                var v16 int32 = 42
                _ = v16
            case "foo bar":
                v17 := v1[""]
                _ = v17
            }
            switch "a" {
            case "\n\t":
                v18 := GenS1{F0: 255, F1: ("hello" + ""), F2: fi11}
                _ = v18
            case "a":
                for fi19 := 0; fi19 < 17; fi19++ {
                    v6 = v6
                    _ = fi19
                }
            case "foo bar":
                v6 -= 0
            }
        }
        _ = fi11
    }
    v5 = append(v5, (v6 | v6))
    for fi20 := 0; fi20 < 3; fi20++ {
        if false {
            if false {
                _ = len(v2)
                fmt.Println(v6)
                _ = len(v5)
            }
            for fi21 := 0; fi21 < 3; fi21++ {
                if (true && ((1.5 * 3.14159) != float32(0.1))) {
                    v22 := GenS1{F0: v3, F1: "foo bar", F2: fi20}
                    _ = v22
                }
                v6 = 2147483647
                v4 = (^(v4))
                _ = fi21
            }
        } else {
            v3 -= v3
            if ((-(fi20)) <= fi20) {
                var v23 int16 = 1000
                v3 += v3
                _ = v23
            }
            v4 += 127
        }
        fmt.Println(fi20)
        if false {
            v24 := GenS1{F0: 325, F1: "\n\t", F2: (-(fi20))}
            fmt.Println(v6, v6)
            var v25 string = ("hello" + ((("a" + "foo bar") + ("" + "foo bar")) + "foo bar"))
            _ = v24
            _ = v25
        }
        _ = fi20
    }
    var v26 uint64 = uint64(v4)
    v27 := GenS1{F0: (v3 + (v3 & v3)), F1: "foo bar", F2: 128}
    var v28 uint64 = (v26 + (v26 / (((v26 & uint64(v4))) | 1)))
    v3 = 65535
    if true {
        if true {
            fmt.Println(v28, v6)
            v4 -= (^(v4))
        }
    } else {
        v3 = v3
    }
    var v29 uint32 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
