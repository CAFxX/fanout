package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 bool; F1 bool; F2 int32 }
type GenS4 struct { F0 int64; F1 int16; F2 int32 }

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]uint32)
    var v3 int = -1
    var v4 uint8 = 1
    switch (("" + ("foo bar" + ("\n\t" + ("foo bar" + ("\n\t" + ""))))) + "a") {
    case "a":
        var v5 int64 = 673
        _ = v5
    }
    switch 32768 {
    case 127:
        v1 = v1
    }
    var v6 uint16 = v1
    v6 = 662
    if true {
        switch "\n\t" {
        case "a":
            var v7 float64 = float64((1e10 + 1.5))
            _ = v7
        default:
            var v8 float64 = ((float64(0.0) - (0.0 * 3.14159)) / 1.0)
            _ = v8
        }
        v9 := v2[(("\n\t" + "") + (("foo bar" + "\n\t") + ("hello" + "")))]
        _ = v9
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        fmt.Println(v6)
    }
    fmt.Println(v6)
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS1{F0: 824, F1: (-1 > 42)}
        _ = fi10
        _ = v11
    }
    v4 += v4
    switch v3 {
    case -1:
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 uint64 = 1
            _ = fi12
            _ = v13
        }
    case -640:
        if (!(127 <= 100)) {
            var v14 int16 = 430
            _ = v14
        }
    }
    var v15 uint8 = 128
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 int64 = 9223372036854775807
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: -1, F1: 4294967295}
    if true {
        var v19 uint8 = v15
        if (1 >= 963) {
            v20 := []int{128, (v3 % ((v3) | 1)), v3}
            if true {
                v21 := GenS2{F0: 575, F1: 0}
                switch (v3 >> ((468) & 63)) {
                case 587:
                    v6 = v1
                case 32768:
                    var v22 int8 = 32
                    _ = v22
                case 4294967295:
                    v23 := GenS3{F0: true, F1: ((-2.25 * 3.14159) < float32(1e10)), F2: 166}
                    _ = v23
                }
                v3 = (-(v3))
                _ = v21
            } else {
                if true {
                    for fi24 := 0; fi24 < 13; fi24++ {
                        for fi25 := 0; fi25 < 13; fi25++ {
                            v26 := GenS3{F0: true, F1: (false && true), F2: -2147483648}
                            _ = fi25
                            _ = v26
                        }
                        _ = fi24
                    }
                }
                var v27 uint32 = 2
                for fi28 := 0; fi28 < 12; fi28++ {
                    v29 := GenS3{F0: true, F1: (v15 > (v19 | 208)), F2: 42}
                    var v30 int16 = 1
                    for fi31 := 0; fi31 < 11; fi31++ {
                        _ = len(v2)
                        var v32 uint32 = uint32(v1)
                        v1 = 32767
                        _ = fi31
                        _ = v32
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = v27
            }
            _ = v20
        } else {
            _ = len(v2)
            delete(v2, ("hello" + "a"))
        }
        _ = v19
    } else {
        var v33 int8 = -128
        v34 := GenS4{F0: int64(v15), F1: 368, F2: 0}
        v35 := GenS1{F0: 1, F1: false}
        _ = v33
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
