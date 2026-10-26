package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 bool; F2 int32 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int64 = -1
    var v2 uint = 2
    var v3 uint8 = 128
    v4 := make(map[int]string)
    var v5 uint = v2
    v6 := GenS1{F0: 1}
    var v7 uint32 = 467
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := v4[(fi8 % ((fi8) | 1))]
        var v10 float64 = 3.14159
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        if true {
            v2 = (v5 << ((1) & 63))
        } else {
            v3 = (v3 >> ((v5) & 63))
            delete(v4, (fi11 - fi11))
        }
        v12 := GenS2{F0: (fi11 - (fi11 % ((fi11) | 1))), F1: false, F2: 42}
        if ("\n\t" > "") {
            switch 9223372036854775807 {
            case -2147483648:
                switch "hello" {
                case "\n\t":
                    var v13 int32 = 1
                    _ = v13
                case "foo bar":
                    for fi14 := 0; fi14 < 15; fi14++ {
                        var v15 float64 = 0.1
                        _ = fi14
                        _ = v15
                    }
                }
            }
        }
        _ = fi11
        _ = v12
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := []int{32767, (fi16 & (fi16 | (fi16 ^ -9223372036854775808)))}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := GenS2{F0: (fi18 >> ((v2) & 63)), F1: (v1 >= (v1 ^ (^(v1)))), F2: 42}
            fmt.Println(v1, v2)
            var v21 uint = (v2 >> ((v5) & 63))
            _ = fi19
            _ = v20
            _ = v21
        }
        if (-32768 > 0) {
            if ((fi18 * (-(fi18))) >= (fi18 * -2147483648)) {
                var v22 uint32 = v7
                _ = len(v4)
                _ = v22
            }
        } else {
            var v23 uint64 = 0
            _ = v23
        }
        switch "" {
        case "hello":
            v1 = 1
        case "":
            fmt.Println(fi18, v1, fi18)
        case "a":
            fmt.Println(v5, fi18, v7)
        default:
            var v24 int8 = 0
            _ = v24
        }
        _ = fi18
    }
    v25 := v4[-1]
    delete(v4, 127)
    v7 -= v7
    var v26 float64 = ((((-2.25 - float64(1e10)) / float64((float32(3.14159) * 1e10))) * 1e10) * (float64((1e10 - 0.1)) - (1e10 * ((0.0 * (float64(float32(-2.25)) + 1.5)) * 3.14159))))
    v1 -= v1
    v4[1] = ((("foo bar" + "a") + v25) + v25)
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 float64 = v26
        v29 := GenS3{F0: -1, F1: ("\n\t" + ("\n\t" + v25))}
        _ = fi27
        _ = v28
        _ = v29
    }
    v1 = (v1 & 42)
    for fi30 := 0; fi30 < 5; fi30++ {
        if true {
            for fi31 := 0; fi31 < 1; fi31++ {
                if false {
                    var v32 uint = v2
                    _ = v32
                } else {
                    var v33 uint = 18446744073709551615
                    _ = v33
                }
                _ = fi31
            }
            for fi34 := 0; fi34 < 1; fi34++ {
                var v35 uint64 = 0
                v26 = (v26 * v26)
                if false {
                    var v36 uint = v5
                    _ = v36
                }
                _ = fi34
                _ = v35
            }
        } else {
            v37 := GenS4{F0: 362, F1: v3}
            _ = v37
        }
        switch v25 {
        case "a":
            var v38 int64 = v1
            _ = v38
        }
        _ = fi30
    }
    var v39 int32 = -2147483648
    v1 -= (v1 & 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
