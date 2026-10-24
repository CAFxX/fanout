package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int64 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 string = "hello"
    v4 := GenS1{F0: 1, F1: 699, F2: 523}
    fmt.Println(v3)
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            v3 = "hello"
            _ = fi6
        }
        _ = len(v2)
        var v7 uint = 117
        _ = fi5
        _ = v7
    }
    v3 = (v3 + v3)
    var v8 uint64 = 2
    v9 := GenS2{F0: -2147483648, F1: 37}
    for fi10 := 0; fi10 < 19; fi10++ {
        delete(v2, 2)
        v11 := v1[v3]
        _ = fi10
        _ = v11
    }
    if (false || true) {
        fmt.Println(v8)
        switch (v3 + v3) {
        case "":
            if (2147483648 <= 2147483647) {
                v3 = v3
            }
        case "\n\t":
            v12 := []int{-160}
            _ = v12
        case "foo bar":
            var v13 int16 = 32767
            _ = v13
        }
        switch 32767 {
        case -2147483648:
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 string = (v3 + (v3 + v3))
                var v16 string = "\n\t"
                v17 := GenS3{F0: 65535, F1: 541}
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        case 32768:
            if false {
                var v18 bool = (!false)
                _ = v18
            }
        case 91:
            if false {
                var v19 uint8 = uint8(v8)
                _ = v19
            }
        default:
            v20 := GenS1{F0: -817, F1: -2147483648, F2: 1}
            _ = v20
        }
    }
    v21 := GenS3{F0: 2, F1: 42}
    _ = len(v1)
    var v22 float64 = 1e10
    fmt.Println(v3, v8, v3)
    if false {
        for fi23 := 0; fi23 < 10; fi23++ {
            if false {
                v22 -= v22
                var v24 float64 = (v22 - v22)
                v25 := GenS4{F0: (v8 & 0), F1: -128}
                _ = v24
                _ = v25
            }
            if (true && (2147483647 > 65535)) {
                var v26 uint8 = 168
                for fi27 := 0; fi27 < 14; fi27++ {
                    switch ((v3 + v3) + "foo bar") {
                    case "\n\t":
                        v28 := []int{128, fi23}
                        _ = v28
                    default:
                        v1["foo bar"] = 987
                    }
                    _ = fi27
                }
                _ = v26
            } else {
                v29 := GenS2{F0: 42, F1: 42}
                _ = v29
            }
            _ = fi23
        }
        for fi30 := 0; fi30 < 17; fi30++ {
            v31 := []int{-919, fi30}
            v32 := GenS5{F0: 103}
            _ = fi30
            _ = v31
            _ = v32
        }
        var v33 uint32 = 385
        _ = v33
    } else {
        for fi34 := 0; fi34 < 2; fi34++ {
            if (32767 != -841) {
                v8 -= v8
                for fi35 := 0; fi35 < 1; fi35++ {
                    v36 := []int{(fi34 | 42), fi35, -2147483648, fi34}
                    _ = fi35
                    _ = v36
                }
                var v37 uint16 = 593
                _ = v37
            }
            _ = fi34
        }
        if true {
            if ((1e10 * (0.1 * 1.5)) <= -2.25) {
                v22 = 3.14159
            } else {
                if false {
                    for fi38 := 0; fi38 < 14; fi38++ {
                        for fi39 := 0; fi39 < 11; fi39++ {
                            var v40 int32 = 42
                            fmt.Println(v22, fi39)
                            v41 := []int{42, -502}
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                        _ = fi38
                    }
                }
                var v42 int8 = 127
                fmt.Println(v3, v8, v22)
                _ = v42
            }
            v22 = 0.0
            var v43 int = -1
            _ = v43
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
