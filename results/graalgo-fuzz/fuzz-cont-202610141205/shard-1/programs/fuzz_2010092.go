package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint; F1 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 127
    v3 := make(map[uint32]bool)
    for fi4 := 0; fi4 < 2; fi4++ {
        v2 -= (v2 >> ((42) & 63))
        v5 := GenS1{F0: 0}
        _ = fi4
        _ = v5
    }
    var v6 string = ("a" + "hello")
    v6 = ""
    v7 := GenS2{F0: 42, F1: 42}
    var v8 uint64 = 18446744073709551615
    if (true || true) {
        v9 := GenS3{F0: -128, F1: false}
        v10 := []int{v2, 32768, v2, v2}
        switch (v6 + v6) {
        case "foo bar":
            _ = len(v10)
        case "hello":
            v1[(v6 + v6)] = 42
        }
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(v2, v6, fi11)
        v12 := []int{(fi11 * -423), (fi11 >> ((18446744073709551615) & 63)), (v2 | 32767)}
        _ = fi11
        _ = v12
    }
    var v13 int64 = 9223372036854775807
    for fi14 := 0; fi14 < 11; fi14++ {
        v13 += 42
        var v15 bool = false
        switch 495 {
        case -2147483648:
            if v15 {
                switch "\n\t" {
                case "":
                    var v16 uint16 = 42
                    _ = v16
                case "foo bar":
                    if (151 > 0) {
                        var v17 uint8 = 1
                        v17 = (v17 * 0)
                        v18 := GenS1{F0: 401}
                        _ = v17
                        _ = v18
                    } else {
                        v19 := GenS1{F0: 18446744073709551615}
                        v20 := v1[("foo bar" + v6)]
                        _ = v19
                        _ = v20
                    }
                case "\n\t":
                    v6 = v6
                }
                var v21 uint32 = uint32(v8)
                _ = v21
            }
        }
        _ = fi14
        _ = v15
    }
    v8 += v8
    var v22 uint64 = (v8 ^ 4294967295)
    if true {
        for fi23 := 0; fi23 < 3; fi23++ {
            for fi24 := 0; fi24 < 12; fi24++ {
                if false {
                    delete(v3, 1)
                    v25 := GenS2{F0: 65535, F1: 42}
                    v26 := GenS4{F0: 292, F1: 42}
                    _ = v25
                    _ = v26
                }
                _ = fi24
            }
            _ = fi23
        }
    } else {
        v27 := []int{}
        var v28 float32 = float32((0.1 * (1.5 - ((-2.25 * -2.25) / (1e10 - float64(((1e10 - -2.25) - (0.0 - 1e10))))))))
        _ = v27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
