package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint }
type GenS2 struct { F0 uint16; F1 uint; F2 string }
type GenS3 struct { F0 uint; F1 uint32 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{127}
    var v3 uint8 = 128
    var v4 int64 = v1
    fmt.Println(v1, v1)
    var v5 uint8 = (v3 >> ((255) & 63))
    var v6 string = ""
    v4 = v4
    if (932 != 211) {
        for fi7 := 0; fi7 < 1; fi7++ {
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                v5 += v5
                _ = len(v2)
            }
            var v8 uint64 = 42
            _ = fi7
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS1{F0: 42, F1: v1, F2: 42}
            for fi11 := 0; fi11 < 7; fi11++ {
                switch fi9 {
                case 492:
                    if len(v2) > 0 {
                        v2[0] = (fi11 << ((855) & 63))
                    }
                }
                v12 := GenS2{F0: 2, F1: 2, F2: "\n\t"}
                _ = fi11
                _ = v12
            }
            if len(v2) > 0 {
                v2[0] = (fi9 >> ((799) & 63))
            }
            _ = fi9
            _ = v10
        }
        v13 := GenS3{F0: 2, F1: 401}
        _ = v13
    }
    v2 = append(v2, 9223372036854775807)
    for fi14 := 0; fi14 < 1; fi14++ {
        if (1.5 <= (((0.1 - 1.5) + float32(-2.25)) - (1e10 * 0.1))) {
            for fi15 := 0; fi15 < 10; fi15++ {
                fmt.Println(fi15, fi15)
                switch (fi15 % ((fi15) | 1)) {
                case 42:
                    fmt.Println(v5)
                case 32767:
                    v16 := GenS1{F0: 811, F1: (-(v4)), F2: 824}
                    _ = v16
                case 310:
                    var v17 int = (fi14 % ((-720) | 1))
                    _ = v17
                }
                _ = fi15
            }
        }
        v5 = (^(v5))
        _ = fi14
    }
    v1 += (^(v4))
    switch 278 {
    case 815:
        if true {
            v3 -= (v3 % ((v5) | 1))
            v3 += 0
        } else {
            var v18 int64 = -9223372036854775808
            fmt.Println(v3)
            _ = v18
        }
    case -402:
        var v19 int8 = 110
        _ = v19
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
