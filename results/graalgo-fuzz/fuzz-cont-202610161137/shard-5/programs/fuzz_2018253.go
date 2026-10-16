package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 int; F2 uint64 }
type GenS6 struct { F0 bool }

func main() {
    var v1 int8 = -1
    var v2 string = ""
    var v3 float64 = (1.5 / (float64(0.1) / (((((-2.25 / (float64(1e10) * 0.1)) / 3.14159) - 0.0) / 1.5) + 1e10)))
    v4 := []int{816, 77, 0, -138}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS1{F0: v1, F1: 2}
        _ = fi5
        _ = v6
    }
    if ((128 >= 42) && true) {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 float64 = float64(1e10)
            var v9 bool = false
            if v9 {
                _ = len(v4)
            } else {
                v10 := GenS1{F0: v1, F1: 128}
                _ = v10
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            v1 = 127
            var v13 bool = (true || true)
            if false {
                v3 -= (v3 * ((1e10 + v3) + 3.14159))
                if (42 != 810) {
                    fmt.Println(v1, fi12)
                }
            } else {
                var v14 uint32 = 356
                _ = v14
            }
            v15 := GenS2{F0: 128}
            v16 := GenS2{F0: 42}
            if len(v4) > 0 {
                v4[0] = int(v1)
            }
            var v17 uint = 2
            _ = fi12
            _ = v13
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi11
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        if len(v4) > 0 {
            v4[0] = (-(fi18))
        }
        v19 := GenS3{F0: 459, F1: v2}
        switch v2 {
        case "\n\t":
            if (183 < -430) {
                v20 := GenS4{F0: 127}
                for fi21 := 0; fi21 < 14; fi21++ {
                    v22 := GenS5{F0: (-(v1)), F1: (fi21 * fi18), F2: 1}
                    var v23 bool = (976 != 550)
                    if true {
                        v24 := GenS6{F0: (false && true)}
                        _ = v24
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = v20
            }
        case "":
            var v25 string = ("foo bar" + "a")
            _ = v25
        case "foo bar":
            var v26 float64 = v3
            _ = v26
        }
        _ = fi18
        _ = v19
    }
    var v27 int32 = 42
    v1 = v1
    var v28 uint64 = 740
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
