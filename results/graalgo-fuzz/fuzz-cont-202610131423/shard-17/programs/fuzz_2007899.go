package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int8; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint8 }
func genFn0(p0 uint64) uint {
    var v6 int16 = 1
    v7 := []int{}
    if ((v6 * v6) > (v6 % ((-32768) | 1))) {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint8 = 254
            _ = fi8
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: 1}
            v6 = v6
            _ = len(v7)
            _ = fi10
            _ = v11
        }
        v6 += (^(v6))
        v6 += v6
    }
    _ = p0
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := []int{32767}
    var v2 uint64 = 4294967295
    v3 := []string{"\n\t", "a", "foo bar", "a"}
    v4 := []uint32{65535, 527}
    var v5 string = ("hello" + ("a" + "\n\t"))
    var v12 uint64 = (v2 ^ 4294967295)
    for fi13 := 0; fi13 < 20; fi13++ {
        fmt.Println(v5, fi13, fi13)
        var v14 uint64 = v12
        switch fi13 {
        case -455:
            var v15 int = (fi13 - fi13)
            _ = v15
        case 42:
            switch v5 {
            case "a":
                if ((311 > v12) || (!false)) {
                    v16 := GenS2{F0: 1, F1: 42, F2: 0}
                    _ = v16
                } else {
                    if (-3 > -128) {
                        if true {
                            v12 -= (v2 / ((1) | 1))
                        } else {
                            fmt.Println(v5, v12)
                        }
                    }
                }
            case "":
                for fi17 := 0; fi17 < 9; fi17++ {
                    v18 := GenS3{F0: 32767, F1: 42, F2: 2}
                    v14 = (v12 / ((v12) | 1))
                    _ = fi17
                    _ = v18
                }
            case "foo bar":
                v2 -= (v2 | (v14 | (v12 >> ((genFn0(208)) & 63))))
            }
        case -2147483648:
            for fi19 := 0; fi19 < 15; fi19++ {
                v1 = append(v1, (fi19 + (fi19 - fi13)))
                var v20 uint = 65535
                _ = fi19
                _ = v20
            }
        }
        _ = fi13
        _ = v14
    }
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    v2 -= 18446744073709551615
    if len(v4) > 0 {
        v4[0] = 42
    }
    var v21 int64 = 1
    v2 = (v2 + 1)
    if (v2 == (v2 & (^(v12)))) {
        for fi22 := 0; fi22 < 8; fi22++ {
            fmt.Println(v12, fi22)
            v12 += v2
            if true {
                if true {
                    var v23 uint8 = 255
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 5; fi24++ {
                        switch (-(fi22)) {
                        case 523:
                            var v25 int = 128
                            var v26 float32 = (-2.25 / 0.1)
                            _ = v25
                            _ = v26
                        default:
                            var v27 float64 = (3.14159 + 0.1)
                            _ = v27
                        }
                        _ = fi24
                    }
                }
                switch (fi22 / (((-(fi22))) | 1)) {
                case 128:
                    if len(v4) > 0 {
                        v4[0] = 345
                    }
                case 127:
                    for fi28 := 0; fi28 < 16; fi28++ {
                        switch -270 {
                        case 32768:
                            v29 := GenS3{F0: 481, F1: 42, F2: 255}
                            _ = v29
                        case 4294967295:
                            v2 -= v2
                        }
                        _ = fi28
                    }
                }
            }
            _ = fi22
        }
        v5 = ((v5 + (v5 + "\n\t")) + v5)
    } else {
        switch ((v5 + "") + "foo bar") {
        case "hello":
            var v30 uint64 = (v2 | v12)
            _ = v30
        case "\n\t":
            v2 -= 317
        case "a":
            var v31 uint16 = 42
            _ = v31
        }
        v32 := GenS1{F0: 0}
        _ = v32
    }
    v5 = (v5 + ("hello" + "hello"))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch v5 {
    case "hello":
        v2 = 2
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = 58
        }
    }
    v33 := GenS3{F0: 833, F1: 165, F2: 255}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
