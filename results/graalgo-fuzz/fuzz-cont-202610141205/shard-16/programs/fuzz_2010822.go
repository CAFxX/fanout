package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint8; F2 string }
func genFn0(p0 int8, p1 int32, p2 int64) uint16 {
    v6 := GenS1{F0: false, F1: 32767, F2: 0}
    if (("foo bar" + ("a" + "")) <= "hello") {
        if (false || true) {
            fmt.Println(p0, p1, p1)
        }
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        p2 -= (p2 << ((2) & 63))
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := []int{fi8, (fi8 * (fi7 - 9223372036854775807))}
            switch "hello" {
            case "hello":
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            case "\n\t":
                p0 += p0
            case "a":
                if false {
                    v10 := GenS1{F0: true, F1: 1000, F2: 42}
                    _ = v10
                } else {
                    p0 += p0
                    fmt.Println(p0)
                }
            default:
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            }
            var v11 uint = 65535
            p1 = -791
            p0 -= int8(p2)
            _ = fi8
            _ = v11
            _ = v9
        }
        _ = fi7
    }
    var v12 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    return 65535
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []int{32767, 0}
    v3 := make(map[string]string)
    var v4 string = ""
    v5 := make(map[int]int)
    if true {
        v4 = v4
        if true {
            v4 = "foo bar"
            var v13 float32 = float32(-2.25)
            _ = v13
        }
        v4 = "hello"
    }
    var v14 uint64 = 34
    v15 := GenS1{F0: true, F1: 1000, F2: 196}
    var v16 int = 4294967295
    var v17 int64 = 0
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 uint32 = 0
        var v20 uint8 = 130
        _ = fi18
        _ = v19
        _ = v20
    }
    if true {
        v17 += (v17 & -841)
        var v21 int8 = int8(v14)
        var v22 uint32 = (^(v1))
        _ = v21
        _ = v22
    } else {
        v23 := GenS2{F0: v1, F1: v16, F2: -128}
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 int64 = v17
            _ = len(v3)
            for fi26 := 0; fi26 < 8; fi26++ {
                v14 = 97
                _ = fi26
            }
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    if (0 > -108) {
        _ = len(v2)
        for fi27 := 0; fi27 < 11; fi27++ {
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 string = (v4 + "foo bar")
                _ = fi28
                _ = v29
            }
            fmt.Println(v1, v14, v16)
            if ((^(v16)) >= (fi27 >> ((255) & 63))) {
                var v30 uint = 453
                var v31 float32 = 1.5
                _ = v30
                _ = v31
            }
            _ = fi27
        }
    }
    v2 = append(v2, v16)
    v32 := GenS3{F0: (v14 >> ((1) & 63)), F1: 0, F2: ""}
    var v33 string = ""
    for fi34 := 0; fi34 < 11; fi34++ {
        v35 := GenS2{F0: 0, F1: v16, F2: 1}
        _ = fi34
        _ = v35
    }
    v17 += v17
    var v36 uint16 = genFn0(-1, -849, 907)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
}
