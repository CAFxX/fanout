package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int; F1 int16; F2 bool }
func genFn0(p0 int16) int8 {
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 uint16 = 0
        p0 -= (^(p0))
        _ = fi5
        _ = v6
    }
    fmt.Println(p0)
    v7 := []int{61}
    if (27 == 42) {
        var v8 int32 = 1
        _ = v8
    }
    var v9 float32 = (0.1 / 1.0)
    _ = p0
    _ = v7
    _ = v9
    return 33
}
func genFn1() int16 {
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint8 = 255
        var v12 bool = true
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        var v13 uint32 = 42
        _ = v13
    } else {
        if false {
            var v14 uint = 0
            v15 := []int{2147483647, 32768, 19}
            _ = v14
            _ = v15
        }
    }
    var v16 uint8 = 200
    v17 := []int{-736}
    v16 += v16
    _ = v16
    _ = v17
    return 0
}

func main() {
    var v1 uint = 65535
    var v2 int = 127
    var v3 uint32 = 2147483647
    var v4 bool = (float64(0.0) == ((-2.25 + float64((0.0 * 1.5))) / (-2.25 - 1e10)))
    v18 := GenS1{F0: -996, F1: 42}
    var v19 uint = 249
    if (-2147483648 > -2147483648) {
        if (-128 <= 42) {
            var v20 int32 = 733
            var v21 uint64 = 42
            for fi22 := 0; fi22 < 19; fi22++ {
                v23 := []int{(fi22 * 0), -1, (fi22 >> (((^(v19))) & 63)), (fi22 * 0)}
                v2 += v2
                _ = fi22
                _ = v23
            }
            _ = v20
            _ = v21
        } else {
            if (((((float64((((1e10 + (((float32(-2.25) / ((3.14159 + (0.0 + -2.25)) - -2.25)) * 0.1) * ((-2.25 + (1.5 + 1e10)) - 1.5))) * -2.25) + (0.0 * (3.14159 / -2.25)))) - (0.0 / 0.1)) / 1.0) - (3.14159 * (0.0 / 1.0))) / ((3.14159 / 1.5) / (1e10 * ((0.1 * float64(1e10)) * 1.5)))) == float64(float32((1.5 * ((((0.1 - 3.14159) * ((0.0 / (1e10 + 1e10)) / (1e10 - 0.0))) * (float64(1.5) / 1.0)) * (0.1 / 1.0)))))) {
                switch "" {
                case "":
                    v24 := []int{}
                    _ = v24
                case "\n\t":
                    var v25 uint16 = 1
                    _ = v25
                }
                if ((v19 - (v19 * (v1 * (v1 << (((v1 + uint(v3))) & 63))))) >= (v19 ^ uint(v3))) {
                    var v26 float32 = 3.14159
                    if false {
                        v27 := []int{v2}
                        _ = v27
                    }
                    _ = v26
                }
                v19 += (v19 << (((v19 / (((^(v1))) | 1))) & 63))
            }
        }
        v28 := []int{2147483647, v2}
        if v4 {
            fmt.Println(v3, v19, v2)
            switch v2 {
            case 4294967295:
                _ = len(v28)
            case -9223372036854775808:
                v29 := GenS1{F0: (v2 / ((v2) | 1)), F1: 0}
                _ = v29
            case 2147483647:
                v28 = append(v28, v2)
            }
            var v30 uint32 = v3
            _ = v30
        }
        _ = v28
    } else {
        if v4 {
            v3 = 1
            var v31 uint32 = (v3 + v3)
            _ = v31
        }
        var v32 uint64 = 63
        switch ("\n\t" + "a") {
        case "foo bar":
            for fi33 := 0; fi33 < 13; fi33++ {
                v4 = (v4 || false)
                _ = fi33
            }
        }
        _ = v32
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        for fi35 := 0; fi35 < 14; fi35++ {
            v36 := []int{}
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    v37 := GenS1{F0: 0, F1: 0}
    v19 += v1
    v3 -= (v3 + 2147483647)
    v38 := GenS2{F0: 42, F1: 42, F2: v4}
    if v4 {
        v4 = (0.0 < -2.25)
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
