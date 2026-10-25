package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint8 }

func main() {
    v1 := []int{4294967295, -2147483648, -82, 451}
    var v2 uint32 = 2147483648
    var v3 int16 = 0
    var v4 float32 = 1e10
    v5 := []int{424, 790, -691, 32768}
    var v6 int16 = v3
    if (true || (!false)) {
        if false {
            for fi7 := 0; fi7 < 7; fi7++ {
                if (42 > fi7) {
                    if len(v5) > 0 {
                        v5[0] = 749
                    }
                } else {
                    v3 = (v6 * v6)
                    var v8 bool = false
                    v2 += (v2 ^ (v2 >> ((65535) & 63)))
                    _ = v8
                }
                v9 := GenS1{F0: false, F1: true, F2: 255}
                _ = fi7
                _ = v9
            }
            var v10 int = 639
            var v11 int16 = -162
            _ = v10
            _ = v11
        }
        v6 -= (v3 << ((905) & 63))
        v2 = v2
    }
    var v12 int16 = v3
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v12, v2, v4)
    fmt.Println(v4)
    v3 = -892
    var v13 uint32 = v2
    fmt.Println(v6)
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS1{F0: (1e10 != 0.0), F1: (0.0 <= (3.14159 / -2.25)), F2: 164}
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 uint8 = 1
            var v18 int8 = 0
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        switch fi19 {
        case 1:
            fmt.Println(v2, v3, v4)
        case 42:
            if (194 <= 0) {
                fmt.Println(fi19)
                v20 := GenS1{F0: (9223372036854775807 < 1), F1: (true && (!(!false))), F2: 2}
                _ = v20
            } else {
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 string = ""
                    _ = fi21
                    _ = v22
                }
            }
        default:
            if false {
                fmt.Println(v4, v6)
                if len(v1) > 0 {
                    v1[0] = 128
                }
                switch (fi19 << ((42) & 63)) {
                case -932:
                    fmt.Println(v4, v13, v13)
                default:
                    var v23 bool = false
                    _ = v23
                }
            } else {
                if len(v1) > 0 {
                    v1[0] = fi19
                }
            }
        }
        var v24 uint16 = 147
        v12 -= 1000
        _ = fi19
        _ = v24
    }
    fmt.Println(v2, v2)
    v3 = (v6 / (((v3 | -1)) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
