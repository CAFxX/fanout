package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int8; F1 int16; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -632
    var v3 int64 = -840
    var v4 float64 = (1e10 - (1e10 - float64(0.0)))
    if false {
        for fi5 := 0; fi5 < 9; fi5++ {
            var v6 uint64 = 39
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 20; fi7++ {
            v3 += (v3 % ((42) | 1))
            _ = fi7
        }
    }
    var v8 string = "\n\t"
    var v9 uint = 75
    if true {
        var v10 uint64 = 18446744073709551615
        v11 := []int{-9223372036854775808}
        v12 := GenS1{F0: 319, F1: 4294967295, F2: 127}
        _ = v10
        _ = v11
        _ = v12
    } else {
        if false {
            v13 := GenS2{F0: 100, F1: int16(v9), F2: 127}
            v4 -= float64((3.14159 / (0.0 + ((0.1 * 1e10) - 1.5))))
            v14 := GenS3{F0: 100}
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        fmt.Println(v4, fi15)
        for fi16 := 0; fi16 < 3; fi16++ {
            fmt.Println(fi16)
            _ = fi16
        }
        _ = fi15
    }
    var v17 int16 = 177
    v2 = (-(v2))
    v18 := GenS4{F0: 0, F1: 0}
    if true {
        var v19 float32 = 1e10
        var v20 uint = v9
        for fi21 := 0; fi21 < 19; fi21++ {
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 int64 = v3
                var v24 bool = true
                _ = fi22
                _ = v23
                _ = v24
            }
            for fi25 := 0; fi25 < 1; fi25++ {
                v4 -= v4
                fmt.Println(v3, v3)
                _ = fi25
            }
            _ = fi21
        }
        _ = v19
        _ = v20
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        if (255 != 32767) {
            v27 := []int{}
            var v28 string = ((v8 + ("\n\t" + v8)) + v8)
            fmt.Println(fi26)
            _ = v27
            _ = v28
        }
        var v29 int64 = v3
        if (v29 > v3) {
            v17 += v17
            switch v8 {
            case "foo bar":
                v4 = v4
            case "":
                switch (((((((v8 + (v8 + (("hello" + "hello") + v8))) + (("hello" + ((v8 + v8) + v8)) + ("" + "a"))) + "") + "hello") + v8) + (v8 + v8)) + ("" + (((((v8 + "hello") + (v8 + v8)) + v8) + v8) + (("hello" + "foo bar") + "\n\t")))) {
                case "\n\t":
                    v30 := []int{int(v2), 212, fi26, (fi26 - 394)}
                    _ = v30
                default:
                    var v31 string = "\n\t"
                    _ = v31
                }
            case "a":
                v32 := GenS4{F0: -36, F1: 0}
                _ = v32
            default:
                v8 = "foo bar"
            }
            if false {
                fmt.Println(v29, v3)
                v1[428] = v8
            }
        }
        _ = fi26
        _ = v29
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        v1[0] = v8
        v34 := GenS4{F0: 42, F1: 256}
        switch (^(fi33)) {
        case -489:
            switch (-(fi33)) {
            case 32767:
                var v35 int32 = v2
                _ = v35
            case 32768:
                v1[249] = "a"
            case 83:
                var v36 int16 = v17
                _ = v36
            }
        }
        _ = fi33
        _ = v34
    }
    v37 := GenS2{F0: 100, F1: (^(v17)), F2: 255}
    v4 = v4
    v2 -= v2
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v8
    _ = v9
}
