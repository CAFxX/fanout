package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint; F1 int16; F2 uint32 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 bool = false
    v3 = (!((1.5 / (3.14159 * float32(-2.25))) > (((1.5 * (-2.25 - (float32((-2.25 / 1.0)) - float32(1e10)))) - ((1.5 + (float32(-2.25) + ((0.1 * 1e10) * (0.0 * float32(((float64(0.1) * 3.14159) / float64(3.14159))))))) + float32(0.0))) + 1.5)))
    for fi4 := 0; fi4 < 8; fi4++ {
        if v3 {
            switch ("a" + "hello") {
            case "a":
                var v5 uint64 = 1
                _ = v5
            }
            var v6 int16 = 1
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 int8 = 0
                var v9 int64 = 935
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = v6
        }
        switch ("" + "hello") {
        case "hello":
            var v10 uint8 = 42
            _ = v10
        default:
            if ("a" > "hello") {
                v3 = (2 < 0)
                v11 := GenS1{F0: 810, F1: 139}
                _ = v11
            }
        }
        v1[fi4] = "\n\t"
        _ = fi4
    }
    v12 := GenS2{F0: -2147483648, F1: 56, F2: 0}
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            switch ("foo bar" + "") {
            case "hello":
                v3 = true
            case "a":
                v3 = (!v3)
            case "\n\t":
                v3 = false
            }
            var v15 int64 = 9223372036854775807
            v3 = v3
            _ = fi14
            _ = v15
        }
        v3 = false
        v16 := GenS3{F0: 2, F1: -32768, F2: 4294967295}
        _ = fi13
        _ = v16
    }
    fmt.Println(v3)
    var v17 float32 = -2.25
    for fi18 := 0; fi18 < 14; fi18++ {
        v19 := GenS1{F0: 18446744073709551615, F1: 42}
        var v20 uint = 0
        var v21 string = ((("" + ("\n\t" + "hello")) + "hello") + "a")
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    v3 = (("" + "hello") > "foo bar")
    if ((42 > 32767) || v3) {
        v3 = (0 >= -32768)
        if (float32(0.0) > 1e10) {
            if v3 {
                switch (("\n\t" + "\n\t") + "a") {
                case "hello":
                    v3 = v3
                case "":
                    var v22 int32 = -1
                    _ = v22
                default:
                    var v23 int8 = 0
                    _ = v23
                }
                v17 -= v17
            }
            v24 := v1[16]
            _ = v24
        }
    } else {
        var v25 int8 = 100
        _ = v25
    }
    if v3 {
        v26 := GenS1{F0: 18446744073709551615, F1: 42}
        var v27 uint8 = 42
        _ = v26
        _ = v27
    } else {
        switch 0 {
        case 2147483647:
            if false {
                for fi28 := 0; fi28 < 8; fi28++ {
                    v17 = v17
                    _ = fi28
                }
                var v29 float32 = v17
                v3 = (!("\n\t" != "a"))
                _ = v29
            } else {
                v17 = (float32(1.5) / 1.5)
                var v30 int32 = -772
                _ = v30
            }
        }
        switch "" {
        case "hello":
            var v31 uint32 = 1
            _ = v31
        case "":
            v32 := GenS2{F0: 2147483647, F1: 970, F2: 0}
            _ = v32
        }
        v17 -= v17
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        for fi34 := 0; fi34 < 13; fi34++ {
            if false {
                v35 := []int{fi33, (fi33 & (fi33 & fi33)), 2147483647}
                var v36 float32 = 0.1
                switch fi33 {
                case 2147483647:
                    v3 = (v3 && false)
                default:
                    var v37 int64 = 1
                    _ = v37
                }
                _ = v35
                _ = v36
            }
            v38 := v2[42]
            _ = fi34
            _ = v38
        }
        v3 = (255 <= 255)
        v39 := []int{fi33, fi33, (^(fi33)), -306}
        _ = fi33
        _ = v39
    }
    v40 := GenS4{F0: -64}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v40
}
