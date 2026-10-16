package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint8, p1 int32) int8 {
    if false {
        p1 -= -1
        if false {
            switch 780 {
            case -9223372036854775808:
                var v5 float32 = ((3.14159 - -2.25) / 1.5)
                _ = v5
            case 32767:
                v6 := []int{9223372036854775807, 1}
                _ = v6
            }
            if (743 == 32767) {
                p1 = 1
            }
            var v7 int8 = 1
            _ = v7
        }
    } else {
        p1 = (p1 % ((42) | 1))
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        p0 = (p0 | p0)
        _ = fi8
    }
    _ = p0
    _ = p1
    return 127
}
func genFn1(p0 uint32, p1 uint32, p2 uint8) int32 {
    var v9 uint = 1
    fmt.Println(p2, p0)
    v10 := GenS1{F0: ("a" + ("a" + "\n\t"))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -86
}

func main() {
    var v1 string = ((("foo bar" + "foo bar") + ("\n\t" + "")) + "a")
    var v2 int64 = 85
    v3 := make(map[int]uint32)
    var v4 float32 = (0.0 + 1.5)
    v2 = (v2 >> ((255) & 63))
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            fmt.Println(v2)
            if (false || true) {
                v13 := GenS2{F0: v1, F1: 62, F2: v2}
                _ = v13
            } else {
                switch "hello" {
                case "\n\t":
                    v14 := []int{fi12, (fi11 >> ((4294967295) & 63)), fi12, fi12}
                    _ = v14
                case "a":
                    var v15 int16 = -32768
                    _ = v15
                }
                v16 := GenS2{F0: "\n\t", F1: 0, F2: v2}
                v2 += v2
                _ = v16
            }
            if false {
                v17 := GenS1{F0: "foo bar"}
                v2 = 250
                _ = v17
            }
            _ = fi12
        }
        v2 = v2
        _ = fi11
    }
    v4 = v4
    fmt.Println(v4)
    switch 42 {
    case 479:
        for fi18 := 0; fi18 < 8; fi18++ {
            delete(v3, 9223372036854775807)
            v2 = v2
            _ = fi18
        }
    case -578:
        v4 -= v4
    }
    v19 := GenS3{F0: 42}
    if true {
        v20 := GenS1{F0: "hello"}
        var v21 int64 = (v2 * (v2 % ((9223372036854775807) | 1)))
        v4 -= (3.14159 / (0.1 / (float32(float64(v4)) * (v4 + 1e10))))
        _ = v20
        _ = v21
    }
    v22 := GenS3{F0: 192}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
