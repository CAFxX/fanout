package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }
type GenS3 struct { F0 uint; F1 string; F2 uint16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint32) int32 {
    p0 += p0
    p0 = p0
    var v5 uint32 = 42
    var v6 float64 = 0.0
    _ = p0
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    var v1 int8 = 1
    v2 := []uint32{2}
    var v3 int64 = -1
    v4 := []string{""}
    v7 := GenS1{F0: -2147483648, F1: 134}
    v1 += int8(v3)
    var v8 float64 = 0.1
    if (2 >= 4294967295) {
        v9 := GenS1{F0: genFn0(1), F1: -822}
        _ = v9
    } else {
        v3 = v3
        var v10 int8 = v1
        if true {
            if false {
                for fi11 := 0; fi11 < 19; fi11++ {
                    v12 := GenS2{F0: 957, F1: 190, F2: "foo bar"}
                    _ = fi11
                    _ = v12
                }
            }
            switch "\n\t" {
            case "\n\t":
                var v13 float32 = ((((3.14159 * (0.0 / (3.14159 + (0.0 - -2.25)))) / 1e10) / (((-2.25 + 0.0) - 0.1) * ((0.0 - 1.5) / 0.1))) * float32((v8 - -2.25)))
                _ = v13
            }
            for fi14 := 0; fi14 < 15; fi14++ {
                v1 = v10
                _ = fi14
            }
        } else {
            var v15 uint32 = 2
            switch 753 {
            case 32767:
                var v16 bool = ((2 < 86) || true)
                _ = v16
            case 269:
                var v17 uint = 4294967295
                _ = v17
            case 9223372036854775807:
                switch "a" {
                case "\n\t":
                    var v18 uint64 = 485
                    _ = v18
                case "":
                    v2 = append(v2, 65535)
                }
            default:
                v19 := GenS3{F0: 2, F1: "hello", F2: 32767}
                _ = v19
            }
            v1 = 127
            _ = v15
        }
        _ = v10
    }
    v3 += (^(v3))
    v8 = v8
    var v20 int = 42
    v21 := GenS2{F0: 714, F1: 2, F2: ((("a" + "") + ("hello" + (("a" + "foo bar") + (("hello" + "foo bar") + ("" + "\n\t"))))) + "foo bar")}
    switch "foo bar" {
    case "a":
        v8 += v8
    case "foo bar":
        var v22 uint64 = 628
        _ = v22
    default:
        v8 = (0.1 / 1e10)
    }
    v4 = append(v4, "a")
    if (false && (-2.25 < -2.25)) {
        if false {
            switch (^(v20)) {
            case 2147483647:
                v20 += v20
            default:
                v23 := GenS1{F0: genFn0(4294967295), F1: -32768}
                _ = v23
            }
        } else {
            fmt.Println(v3)
            var v24 float64 = v8
            if len(v2) > 0 {
                v2[0] = 331
            }
            _ = v24
        }
    }
    v20 += v20
    v25 := GenS4{F0: 618}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
