package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
func genFn0(p0 int, p1 int8, p2 bool) uint {
    p0 -= p0
    var v7 int64 = 9223372036854775807
    v8 := GenS1{F0: 1}
    fmt.Println(p0, v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 196
}

func main() {
    var v1 string = ("a" + "\n\t")
    var v2 uint8 = 1
    var v3 uint64 = 2
    var v4 int8 = 0
    v5 := []string{("\n\t" + v1)}
    var v6 uint16 = 0
    var v9 float32 = (float32((3.14159 * ((1.5 / 1.0) * 0.1))) + (3.14159 / float32(float64(-2.25))))
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: 4294967295}
        _ = fi10
        _ = v11
    }
    v3 += (v3 % (((v3 >> ((556) & 63))) | 1))
    var v12 uint = 723
    switch v1 {
    case "":
        v13 := GenS1{F0: 2147483647}
        _ = v13
    default:
        for fi14 := 0; fi14 < 11; fi14++ {
            if true {
                if false {
                    fmt.Println(v6, v12)
                    v15 := GenS1{F0: 2147483647}
                    v5 = append(v5, "a")
                    _ = v15
                }
            } else {
                switch (fi14 / (((fi14 | (fi14 + fi14))) | 1)) {
                case 96:
                    switch "foo bar" {
                    case "foo bar":
                        var v16 int8 = v4
                        _ = v16
                    }
                default:
                    var v17 uint8 = (v2 + 249)
                    _ = v17
                }
            }
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 uint32 = 419
                v20 := GenS1{F0: (v19 / ((v19) | 1))}
                v19 -= (v19 & v19)
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi14
        }
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        fmt.Println(fi21, v12)
        v9 = v9
        fmt.Println(v6)
        _ = fi21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        switch fi22 {
        case 32768:
            if len(v5) > 0 {
                v5[0] = ("\n\t" + (v1 + v1))
            }
        default:
            v23 := GenS2{F0: v1}
            _ = v23
        }
        switch "hello" {
        case "foo bar":
            var v24 string = "a"
            _ = v24
        case "hello":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "":
            switch (v1 + v1) {
            case "a":
                v2 -= v2
            case "":
                var v25 uint64 = 1
                _ = v25
            default:
                var v26 int = (fi22 | (fi22 & -9223372036854775808))
                _ = v26
            }
        }
        _ = fi22
    }
    v27 := GenS2{F0: (v1 + ("hello" + "\n\t"))}
    switch ((v1 + v1) + v1) {
    case "\n\t":
        switch 373 {
        case 0:
            v5 = append(v5, "\n\t")
        }
    case "a":
        var v28 int32 = 0
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
