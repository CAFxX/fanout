package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint; F2 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 int; F2 uint64 }

func main() {
    var v1 int8 = 100
    var v2 int32 = 42
    var v3 float64 = float64(0.1)
    for fi4 := 0; fi4 < 4; fi4++ {
        fmt.Println(v2, v3)
        var v5 string = (("a" + "hello") + "\n\t")
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 255, F1: 543}
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS2{F0: 35, F1: 950}
        for fi9 := 0; fi9 < 6; fi9++ {
            v3 = (float64((0.0 + (3.14159 * (float32(1.5) * 1e10)))) - (v3 * (v3 - ((float64(((-2.25 + -2.25) + 3.14159)) * v3) * 3.14159))))
            switch fi9 {
            case 1:
                v3 += (v3 / (3.14159 / ((v3 / (3.14159 * v3)) - (0.0 + float64(0.1)))))
            case 2147483647:
                var v10 uint = 18446744073709551615
                _ = v10
            }
            if false {
                v11 := GenS2{F0: 32767, F1: 42}
                switch fi9 {
                case 0:
                    var v12 uint8 = 165
                    _ = v12
                }
                _ = v11
            }
            _ = fi9
        }
        var v13 uint32 = 460
        _ = fi7
        _ = v13
        _ = v8
    }
    v14 := []int{2147483647, -1, -9223372036854775808, 9223372036854775807}
    v15 := GenS3{F0: -83, F1: 255, F2: ("a" + "hello")}
    fmt.Println(v2, v1)
    v2 = v2
    if (42 == 588) {
        switch 32768 {
        case 32768:
            v3 = -2.25
        case -112:
            var v16 int32 = (-(v2))
            _ = v16
        }
    }
    var v17 uint = 4294967295
    if (3.14159 != ((-2.25 - ((1.5 * 0.1) / 3.14159)) * 1e10)) {
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 uint8 = 0
            var v20 int16 = -1
            _ = fi18
            _ = v19
            _ = v20
        }
    } else {
        fmt.Println(v2, v17)
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 uint16 = 2
            var v23 uint8 = 128
            v24 := GenS3{F0: int8(fi21), F1: (v17 + 0), F2: "a"}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
    }
    if false {
        var v25 string = "a"
        fmt.Println(v2, v25)
        _ = v25
    } else {
        v3 += (float64(-2.25) + v3)
        if true {
            for fi26 := 0; fi26 < 9; fi26++ {
                v27 := GenS4{F0: v1}
                if false {
                    var v28 int16 = -245
                    var v29 string = "foo bar"
                    _ = v28
                    _ = v29
                }
                _ = len(v14)
                _ = fi26
                _ = v27
            }
        } else {
            v1 += (v1 + 127)
            v30 := GenS2{F0: 0, F1: 1}
            if (v17 > (v17 << ((772) & 63))) {
                var v31 int16 = 1
                if len(v14) > 0 {
                    v14[0] = 239
                }
                switch 2147483647 {
                case -2147483648:
                    v17 -= 965
                }
                _ = v31
            } else {
                v32 := GenS2{F0: 860, F1: 255}
                v2 = -2147483648
                _ = v32
            }
            _ = v30
        }
        v33 := GenS5{F0: -36, F1: 1, F2: 4294967295}
        _ = v33
    }
    switch "a" {
    case "":
        var v34 uint16 = 2
        _ = v34
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v3 = v3
    v17 -= v17
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v6
}
