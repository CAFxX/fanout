package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }

func main() {
    v1 := make(map[string]bool)
    var v2 bool = ((-2.25 / 0.1) < 3.14159)
    var v3 uint16 = 255
    if (-906 > -15) {
        if (1 < -128) {
            v2 = false
            var v4 int16 = 0
            for fi5 := 0; fi5 < 1; fi5++ {
                v6 := []int{(fi5 / ((42) | 1))}
                v4 -= (v4 >> ((4294967295) & 63))
                var v7 int64 = -1
                _ = fi5
                _ = v6
                _ = v7
            }
            _ = v4
        } else {
            switch "a" {
            case "a":
                var v8 int32 = -109
                _ = v8
            }
            v3 = (v3 + (v3 ^ (v3 / (((^(v3))) | 1))))
        }
        if (1.5 == 3.14159) {
            v9 := []int{42}
            _ = len(v9)
            var v10 bool = true
            _ = v10
            _ = v9
        } else {
            var v11 uint = 4294967295
            _ = v11
        }
        _ = len(v1)
    }
    v3 = (v3 ^ 685)
    if v2 {
        switch 107 {
        case 128:
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 int16 = 1
                switch ("foo bar" + "") {
                case "a":
                    v2 = ((2 != 2) && (-873 != 1))
                case "hello":
                    v13 = (^(v13))
                }
                if v2 {
                    v14 := GenS1{F0: 4294967295, F1: (fi12 + 539)}
                    _ = v14
                }
                _ = fi12
                _ = v13
            }
        case -808:
            var v15 int16 = -793
            _ = v15
        }
    }
    v3 = (v3 >> ((255) & 63))
    var v16 int16 = 42
    if v2 {
        v3 = (^(v3))
        v2 = true
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            v18 := GenS1{F0: 1, F1: (fi17 & fi17)}
            v3 = v3
            v3 = (v3 & (^(v3)))
            _ = fi17
            _ = v18
        }
        fmt.Println(v16, v2, v16)
    }
    fmt.Println(v2, v2)
    v3 = 605
    if v2 {
        fmt.Println(v16, v2)
        fmt.Println(v3, v16)
    }
    v19 := GenS1{F0: 951, F1: 2147483647}
    fmt.Println(v16, v16)
    var v20 float64 = 1.5
    if false {
        for fi21 := 0; fi21 < 19; fi21++ {
            if (460 != 653) {
                _ = len(v1)
                for fi22 := 0; fi22 < 6; fi22++ {
                    v16 = v16
                    v3 += 461
                    _ = fi22
                }
                v23 := []int{}
                _ = v23
            }
            if (1.5 >= ((float32(float64(0.0)) / 1.0) + (1.5 - 0.1))) {
                var v24 int32 = 11
                _ = v24
            } else {
                v16 -= (v16 ^ (v16 ^ v16))
            }
            _ = fi21
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
