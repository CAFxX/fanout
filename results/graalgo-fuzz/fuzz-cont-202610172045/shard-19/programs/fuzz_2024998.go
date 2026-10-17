package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint32; F2 uint }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 int64) int32 {
    fmt.Println(p0)
    fmt.Println(p0)
    var v6 int16 = -32768
    _ = p0
    _ = v6
    return 2147483647
}

func main() {
    var v1 uint16 = 95
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 100
    var v7 float64 = ((float64(3.14159) * float64(-2.25)) + ((3.14159 + 0.1) / 1e10))
    var v8 string = "foo bar"
    var v9 uint32 = uint32(v5)
    if (291 != (v9 / ((v9) | 1))) {
        var v10 int32 = genFn0(0)
        v8 = "\n\t"
        _ = v10
    }
    var v11 float64 = v7
    v12 := GenS1{F0: 2}
    if true {
        var v13 string = v8
        for fi14 := 0; fi14 < 8; fi14++ {
            v8 = "hello"
            var v15 uint8 = v5
            _ = fi14
            _ = v15
        }
        if (true || true) {
            var v16 int16 = 727
            _ = v16
        }
        _ = v13
    }
    var v17 int16 = 1000
    for fi18 := 0; fi18 < 9; fi18++ {
        v19 := []int{}
        switch (^(fi18)) {
        case 4294967295:
            v20 := GenS1{F0: v5}
            _ = v20
        case 234:
            v1 -= uint16(v9)
        }
        if (!false) {
            var v21 uint16 = 42
            switch (v8 + (((("\n\t" + v8) + "") + v8) + "foo bar")) {
            case "":
                _ = len(v19)
            case "a":
                switch ("a" + "a") {
                case "foo bar":
                    switch v8 {
                    case "\n\t":
                        for fi22 := 0; fi22 < 9; fi22++ {
                            v23 := GenS1{F0: 255}
                            v24 := GenS2{F0: 128, F1: (v9 >> ((189) & 63)), F2: 57}
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    case "hello":
                        v1 += v1
                    case "a":
                        if (!((^(v17)) < -32768)) {
                            v17 = v17
                        }
                    default:
                        v9 = (v9 >> ((255) & 63))
                    }
                case "":
                    var v25 float64 = 0.0
                    _ = v25
                case "\n\t":
                    v26 := GenS1{F0: v5}
                    _ = v26
                }
            }
            if (true || (!(((v8 + "foo bar") >= "foo bar") && (v5 <= (v5 / (((v5 % ((v5) | 1))) | 1)))))) {
                v21 += (v21 | (v21 + 0))
            }
            _ = v21
        } else {
            v9 = v9
        }
        _ = fi18
        _ = v19
    }
    v27 := GenS3{F0: 42, F1: 128}
    var v28 int32 = int32(v9)
    v7 -= (3.14159 / v11)
    v29 := GenS3{F0: 0, F1: 809}
    v30 := GenS1{F0: v5}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
