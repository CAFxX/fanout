package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int32 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := []int{-9223372036854775808}
    v2 := make(map[uint32]int)
    var v3 int = 1
    v4 := make(map[int]int)
    v5 := []string{}
    v6 := make(map[string]string)
    v3 -= (v3 * -9223372036854775808)
    if (2 >= 146) {
        for fi7 := 0; fi7 < 1; fi7++ {
            fmt.Println(fi7, fi7)
            if false {
                delete(v4, v3)
                for fi8 := 0; fi8 < 16; fi8++ {
                    switch "" {
                    case "\n\t":
                        switch fi7 {
                        case 630:
                            v9 := GenS1{F0: 255, F1: (127 >= 0), F2: 0}
                            _ = v9
                        case 996:
                            var v10 int16 = -1
                            _ = v10
                        case -2147483648:
                            var v11 int8 = 127
                            _ = v11
                        default:
                            var v12 uint16 = 42
                            _ = v12
                        }
                    default:
                        var v13 int16 = 1000
                        _ = v13
                    }
                    _ = fi8
                }
            }
            _ = fi7
        }
    } else {
        switch "\n\t" {
        case "\n\t":
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 float32 = float32(3.14159)
                _ = fi14
                _ = v15
            }
        }
        if true {
            v16 := GenS2{F0: 127, F1: v3}
            v3 += (v3 * v3)
            v3 = -726
            _ = v16
        } else {
            var v17 int8 = 77
            var v18 uint8 = 24
            _ = len(v1)
            _ = v17
            _ = v18
        }
        var v19 uint16 = 255
        _ = v19
    }
    v20 := GenS1{F0: 160, F1: (32767 == 0), F2: 1}
    if (980 != 32767) {
        if len(v5) > 0 {
            v5[0] = "a"
        }
    }
    if false {
        if (!(118 >= -2147483648)) {
            fmt.Println(v3)
        } else {
            switch (("hello" + "\n\t") + "foo bar") {
            case "a":
                if (2147483647 >= 2) {
                    var v21 uint = 821
                    _ = v21
                }
            }
            var v22 int16 = 0
            var v23 uint8 = 0
            _ = v22
            _ = v23
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v3 += (v3 >> ((18446744073709551615) & 63))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v24 := GenS2{F0: 42, F1: -1}
    if (!false) {
        v25 := v2[80]
        v26 := GenS3{F0: 18446744073709551615}
        _ = v25
        _ = v26
    }
    v27 := v6[("hello" + "a")]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
