package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int8; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 int32; F2 string }
type GenS5 struct { F0 uint8; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 1.5
    var v3 int = 127
    var v4 uint = 74
    v5 := make(map[string]uint32)
    v6 := GenS1{F0: -128}
    var v7 float32 = v2
    v7 = v2
    for fi8 := 0; fi8 < 1; fi8++ {
        switch v3 {
        case 0:
            v9 := v1["\n\t"]
            _ = v9
        default:
            if (407 == 2147483647) {
                if true {
                    v10 := []int{}
                    var v11 float64 = float64(0.1)
                    var v12 string = "\n\t"
                    _ = v10
                    _ = v11
                    _ = v12
                }
            } else {
                v3 += (v3 - -2147483648)
            }
        }
        fmt.Println(v2)
        var v13 int64 = 1
        _ = fi8
        _ = v13
    }
    v4 += v4
    v4 += v4
    switch (^(v3)) {
    case -272:
        switch "a" {
        case "hello":
            for fi14 := 0; fi14 < 11; fi14++ {
                v5["a"] = 522
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 int64 = 42
                    _ = fi15
                    _ = v16
                }
                var v17 int64 = 800
                _ = fi14
                _ = v17
            }
        case "a":
            var v18 uint8 = 127
            _ = v18
        default:
            v3 += -153
        }
    }
    v19 := GenS1{F0: 0}
    v2 = 1.5
    v20 := v5[""]
    v21 := GenS2{F0: false, F1: -128, F2: 0}
    for fi22 := 0; fi22 < 3; fi22++ {
        v1[("\n\t" + (("\n\t" + "\n\t") + "\n\t"))] = (v20 * v20)
        v2 = 0.0
        var v23 int32 = int32(v3)
        _ = fi22
        _ = v23
    }
    if true {
        v24 := v1[(("a" + ("a" + "\n\t")) + "foo bar")]
        _ = v24
    } else {
        fmt.Println(v7, v4)
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 uint64 = 1
        if (380 <= 428) {
            if false {
                v27 := GenS3{F0: v20}
                v28 := GenS4{F0: (v20 >> ((0) & 63)), F1: 1, F2: ("hello" + "a")}
                if (2 != 78) {
                    for fi29 := 0; fi29 < 8; fi29++ {
                        v4 += v4
                        fmt.Println(v2)
                        v7 = (float32(((float64(3.14159) * 0.0) - ((1e10 - (3.14159 / 1.5)) * (0.0 - 0.0)))) / v7)
                        _ = fi29
                    }
                    v30 := GenS3{F0: v20}
                    _ = v30
                }
                _ = v27
                _ = v28
            } else {
                switch "\n\t" {
                case "\n\t":
                    switch (fi25 - (-(v3))) {
                    case 2147483647:
                        for fi31 := 0; fi31 < 9; fi31++ {
                            v26 -= uint64(v20)
                            _ = fi31
                        }
                    }
                }
                v20 += 0
                v5["a"] = (v20 & v20)
            }
            v2 -= ((1e10 + v7) - (1e10 * v2))
            var v32 int64 = -899
            _ = v32
        } else {
            v2 = v7
        }
        _ = fi25
        _ = v26
    }
    v20 = 126
    delete(v1, "")
    v33 := GenS5{F0: 127, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
