package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint16; F2 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 100
    v3 := make(map[string]uint32)
    var v4 uint8 = 171
    var v5 int = 4294967295
    _ = len(v1)
    v2 = v2
    if true {
        switch ("foo bar" + (("hello" + "a") + "hello")) {
        case "foo bar":
            var v6 float64 = (1.5 + (3.14159 - 1.5))
            _ = v6
        case "hello":
            v4 += v4
        }
        v4 = (v4 << ((2) & 63))
        v5 += (v5 ^ v5)
    } else {
        for fi7 := 0; fi7 < 14; fi7++ {
            _ = len(v1)
            var v8 uint32 = 324
            _ = fi7
            _ = v8
        }
        switch v5 {
        case -113:
            var v9 int16 = -371
            _ = v9
        case -1:
            var v10 bool = false
            _ = v10
        }
        if false {
            if (v4 != 255) {
                for fi11 := 0; fi11 < 7; fi11++ {
                    var v12 bool = false
                    _ = fi11
                    _ = v12
                }
                fmt.Println(v5)
                var v13 float32 = ((1e10 * float32(1.5)) * ((float32(((0.0 / 3.14159) * float64(-2.25))) + 0.1) - (float32(3.14159) * ((0.0 * float32(float64((1e10 - (0.1 / 1e10))))) / 3.14159))))
                _ = v13
            }
        } else {
            var v14 uint64 = 2
            v15 := GenS1{F0: 65535, F1: (v4 & (v4 << ((18446744073709551615) & 63)))}
            _ = v14
            _ = v15
        }
    }
    v16 := GenS1{F0: 255, F1: v4}
    v5 += v5
    fmt.Println(v5)
    var v17 int32 = 579
    v4 += v4
    v2 = (v2 * -128)
    v18 := GenS2{F0: ("" + "hello")}
    for fi19 := 0; fi19 < 8; fi19++ {
        if (false || (!((v17 >= v17) || (false && false)))) {
            if true {
                v17 -= (v17 - (-(v17)))
            } else {
                var v20 float32 = 0.0
                _ = v20
            }
            var v21 uint8 = v4
            _ = v21
        }
        _ = fi19
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        v23 := GenS3{F0: "a", F1: 523, F2: v2}
        _ = fi22
        _ = v23
    }
    if (("foo bar" <= ((("\n\t" + ("\n\t" + ("foo bar" + ""))) + "a") + ("hello" + "a"))) || ((false || false) && false)) {
        v2 += int8(v17)
        v4 = 255
    } else {
        for fi24 := 0; fi24 < 16; fi24++ {
            v25 := []int{v5, fi24, v5}
            _ = fi24
            _ = v25
        }
        var v26 float32 = (1.5 - (1.5 + 0.1))
        _ = v26
    }
    v17 = v17
    switch ("" + "hello") {
    case "foo bar":
        v27 := []int{}
        _ = v27
    case "\n\t":
        v28 := []int{}
        _ = v28
    }
    v4 = v4
    fmt.Println(v4, v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
