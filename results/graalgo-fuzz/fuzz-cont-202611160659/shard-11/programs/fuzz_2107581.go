package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint; F1 uint16; F2 int8 }
type GenS4 struct { F0 string; F1 uint; F2 int16 }

func main() {
    var v1 int16 = 32767
    var v2 float64 = ((float64(-2.25) + 3.14159) / (1e10 - 0.1))
    var v3 bool = (false || false)
    v4 := make(map[uint32]bool)
    v5 := GenS1{F0: "a"}
    v1 = -815
    var v6 string = ("\n\t" + "hello")
    if ((390 != -365) && (!v3)) {
        for fi7 := 0; fi7 < 15; fi7++ {
            switch fi7 {
            case -654:
                var v8 int64 = 1
                _ = v8
            case 215:
                var v9 string = v6
                _ = v9
            case 127:
                v3 = (true || (!v3))
            }
            var v10 int16 = v1
            var v11 bool = ((fi7 << ((75) & 63)) != (-(fi7)))
            _ = fi7
            _ = v10
            _ = v11
        }
        v1 = v1
    } else {
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS1{F0: "hello"}
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: ("foo bar" + v6)}
        if v3 {
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := []int{575, fi15, fi15}
                _ = fi15
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 uint16 = 42
                v19 := GenS1{F0: "foo bar"}
                _ = fi17
                _ = v18
                _ = v19
            }
            switch v6 {
            case "hello":
                v3 = true
            case "":
                var v20 uint64 = 4294967295
                _ = v20
            }
            v6 = v6
        }
        _ = v14
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        v6 = v6
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 uint16 = 255
            v24 := []int{fi22, (fi21 / ((0) | 1)), (fi22 % ((fi21) | 1))}
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 uint64 = 2
        _ = fi21
        _ = v25
    }
    fmt.Println(v1)
    var v26 int = -2147483648
    if v3 {
        if (2147483647 != 0) {
            var v27 uint64 = 484
            v27 = (^(v27))
            v28 := GenS2{F0: v3, F1: (!(0 <= 1))}
            _ = v27
            _ = v28
        } else {
            switch v26 {
            case 4294967295:
                var v29 uint8 = 42
                _ = v29
            default:
                v1 = (v1 | (v1 * v1))
            }
            fmt.Println(v26)
        }
        v30 := GenS3{F0: 0, F1: 173, F2: -1}
        _ = v30
    } else {
        v31 := GenS4{F0: v6, F1: 1, F2: (-(v1))}
        v1 = (v1 + v1)
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
