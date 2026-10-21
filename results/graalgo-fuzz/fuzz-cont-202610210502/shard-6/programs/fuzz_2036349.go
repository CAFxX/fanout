package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int32; F1 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 1
    var v3 int32 = 2147483647
    var v4 int64 = 9223372036854775807
    var v5 string = ("\n\t" + "hello")
    if (2 > 0) {
        var v6 int64 = v4
        if ((false || false) && false) {
            v7 := GenS1{F0: v2, F1: true}
            _ = v7
        }
        v5 = v5
        _ = v6
    }
    v8 := []int{1}
    switch -667 {
    case 42:
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint8 = 0
            for fi11 := 0; fi11 < 9; fi11++ {
                if (255 != 65535) {
                    var v12 int32 = -2147483648
                    fmt.Println(v12)
                    _ = v12
                } else {
                    v5 = ((v5 + v5) + v5)
                    var v13 uint8 = v10
                    _ = v13
                }
                fmt.Println(fi9)
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 string = "hello"
        switch -237 {
        case 1:
            if false {
                v15 = ("\n\t" + (v15 + (((v15 + "a") + "hello") + "foo bar")))
            }
        case -9223372036854775808:
            if (255 >= 255) {
                for fi16 := 0; fi16 < 15; fi16++ {
                    v17 := GenS1{F0: (v2 & 0), F1: (!(!(42 > 128)))}
                    v18 := v1[((v15 + v15) + v5)]
                    var v19 uint = 65535
                    v20 := GenS1{F0: 2, F1: (!((v19 / (((^(v19))) | 1)) != 0))}
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                var v21 int16 = 217
                _ = v21
            }
        }
        _ = fi14
        _ = v15
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 uint = 0
        var v24 uint32 = 65535
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        v3 = (v3 & 2147483647)
        v26 := GenS2{F0: -1, F1: v5}
        for fi27 := 0; fi27 < 2; fi27++ {
            if true {
                v28 := GenS3{F0: 1}
                if (!(4294967295 != 65535)) {
                    var v29 float64 = float64(float32((1e10 / (-2.25 - 0.0))))
                    var v30 int32 = v3
                    if false {
                        v8 = append(v8, fi25)
                    }
                    _ = v29
                    _ = v30
                }
                _ = v28
            } else {
                var v31 uint32 = 514
                switch fi25 {
                case 32767:
                    v32 := GenS4{F0: (^(v3)), F1: false}
                    v31 += v31
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = v32
                case 567:
                    var v33 uint8 = 0
                    _ = v33
                }
                var v34 int32 = v3
                _ = v31
                _ = v34
            }
            if (!true) {
                v35 := GenS4{F0: 2147483647, F1: true}
                _ = v35
            } else {
                switch fi27 {
                case -9223372036854775808:
                    var v36 float32 = 3.14159
                    _ = v36
                }
                var v37 uint = 743
                var v38 int8 = 40
                _ = v37
                _ = v38
            }
            switch v5 {
            case "foo bar":
                fmt.Println(v3, v4)
            case "a":
                v39 := GenS5{F0: (v3 - (v3 - (v3 << ((969) & 63)))), F1: 127}
                _ = v39
            }
            var v40 int16 = -32768
            fmt.Println(v5, v3)
            _ = fi27
            _ = v40
        }
        _ = fi25
        _ = v26
    }
    fmt.Println(v2, v5)
    v2 = 0
    fmt.Println(v5, v3)
    var v41 bool = (-89 < -128)
    v42 := GenS2{F0: 0, F1: v5}
    fmt.Println(v41, v3)
    var v43 bool = ((v4 & (v4 ^ v4)) == -240)
    v2 = (v2 % (((v2 << ((316) & 63))) | 1))
    for fi44 := 0; fi44 < 2; fi44++ {
        for fi45 := 0; fi45 < 12; fi45++ {
            var v46 float32 = 1e10
            _ = fi45
            _ = v46
        }
        _ = fi44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v8
}
