package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int64 }
func genFn0(p0 uint64, p1 int32) int32 {
    for fi4 := 0; fi4 < 2; fi4++ {
        p1 = (p1 / ((p1) | 1))
        _ = fi4
    }
    fmt.Println(p0, p1)
    p1 += 910
    v5 := GenS1{F0: ("hello" + "foo bar"), F1: -1}
    _ = p0
    _ = p1
    _ = v5
    return p1
}
func genFn1(p0 bool) uint32 {
    if (!(916 >= 298)) {
        p0 = (255 <= 46)
        switch (("foo bar" + ("" + ("" + "foo bar"))) + (("foo bar" + "a") + ("hello" + "a"))) {
        case "foo bar":
            switch "\n\t" {
            case "a":
                for fi6 := 0; fi6 < 11; fi6++ {
                    var v7 uint16 = 65535
                    var v8 int8 = 0
                    _ = fi6
                    _ = v7
                    _ = v8
                }
            case "foo bar":
                var v9 uint64 = 307
                _ = v9
            case "\n\t":
                var v10 int16 = 0
                _ = v10
            }
        case "\n\t":
            v11 := []int{}
            _ = v11
        case "hello":
            for fi12 := 0; fi12 < 19; fi12++ {
                switch ("hello" + "\n\t") {
                case "":
                    v13 := GenS1{F0: ("\n\t" + "foo bar"), F1: 42}
                    _ = v13
                case "hello":
                    v14 := GenS2{F0: -907, F1: 32767}
                    _ = v14
                }
                p0 = p0
                var v15 int = fi12
                _ = fi12
                _ = v15
            }
        default:
            var v16 int64 = -1
            _ = v16
        }
        var v17 int64 = -9223372036854775808
        _ = v17
    }
    var v18 string = "a"
    v18 = (v18 + v18)
    _ = p0
    _ = v18
    return 65535
}

func main() {
    var v1 int64 = 1
    var v2 float32 = 1e10
    v3 := make(map[string]bool)
    v1 = v1
    v19 := []int{-400, -9223372036854775808}
    switch -2147483648 {
    case 127:
        var v20 uint32 = 1
        _ = v20
    case -2147483648:
        v3[("a" + "foo bar")] = (2 > 226)
    }
    v1 -= (^(v1))
    for fi21 := 0; fi21 < 12; fi21++ {
        v22 := GenS3{F0: 719, F1: 0, F2: v1}
        v1 -= (-(v1))
        _ = fi21
        _ = v22
    }
    var v23 float64 = (3.14159 * 1.5)
    var v24 uint8 = 42
    var v25 uint16 = 0
    var v26 uint64 = 42
    v26 = (v26 << ((681) & 63))
    v26 += (v26 - v26)
    if (true || true) {
        var v27 int16 = 0
        v28 := GenS3{F0: v26, F1: (^(v25)), F2: (v1 + (v1 + v1))}
        _ = v27
        _ = v28
    } else {
        delete(v3, "")
        if (!("" < "a")) {
            v29 := GenS1{F0: "a", F1: (-(v1))}
            _ = v29
        } else {
            for fi30 := 0; fi30 < 9; fi30++ {
                v26 += v26
                _ = fi30
            }
            v25 = 300
            v26 = 18446744073709551615
        }
    }
    for fi31 := 0; fi31 < 19; fi31++ {
        var v32 float64 = (v23 - (0.1 / (float64(v2) / float64(1.5))))
        v1 -= v1
        _ = fi31
        _ = v32
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
