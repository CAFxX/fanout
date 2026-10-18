package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int16 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int64; F1 bool; F2 int }
type GenS4 struct { F0 uint64; F1 int64 }

func main() {
    var v1 bool = (!false)
    var v2 int16 = 901
    var v3 int16 = v2
    switch "" {
    case "a":
        fmt.Println(v3)
    case "foo bar":
        v4 := GenS1{F0: 2147483647, F1: v2, F2: v2}
        _ = v4
    default:
        v5 := []int{-2147483648, -214, 101}
        _ = v5
    }
    fmt.Println(v2, v3, v2)
    var v6 int16 = (v2 % ((v2) | 1))
    v3 = v2
    var v7 int64 = 0
    v8 := []int{-2147483648, 845, 303}
    fmt.Println(v7, v6)
    v9 := GenS2{F0: 113, F1: 65535, F2: 4294967295}
    v8 = append(v8, -2147483648)
    if v1 {
        switch -9223372036854775808 {
        case 942:
            for fi10 := 0; fi10 < 7; fi10++ {
                _ = len(v8)
                v3 += v3
                _ = fi10
            }
        case 42:
            v2 += (^(v6))
        case 128:
            var v11 string = "a"
            _ = v11
        default:
            v7 = v7
        }
        if true {
            switch 42 {
            case 4294967295:
                v12 := GenS3{F0: v7, F1: (v7 >= v7), F2: 32767}
                _ = v12
            case 9223372036854775807:
                switch 703 {
                case 504:
                    fmt.Println(v3, v7, v7)
                }
            }
            v13 := GenS2{F0: 4294967295, F1: 32767, F2: 65535}
            v14 := GenS2{F0: 0, F1: 65535, F2: 0}
            _ = v13
            _ = v14
        }
        var v15 bool = (!v1)
        _ = v15
    }
    if ((float32(float64((3.14159 / 0.1))) / -2.25) != 0.1) {
        switch 4294967295 {
        case 42:
            var v16 int64 = (v7 + v7)
            _ = v16
        case 32768:
            var v17 int8 = 6
            _ = v17
        case 0:
            var v18 float64 = float64(float32(1e10))
            _ = v18
        }
    }
    if (2147483647 != 2147483647) {
        if v1 {
            var v19 float32 = 0.1
            if (int32(v7) > -2147483648) {
                fmt.Println(v6)
            }
            var v20 float64 = 0.0
            _ = v19
            _ = v20
        } else {
            switch 128 {
            case -2147483648:
                if v1 {
                    v21 := GenS3{F0: v7, F1: v1, F2: 42}
                    fmt.Println(v7, v7, v3)
                    if len(v8) > 0 {
                        v8[0] = 412
                    }
                    _ = v21
                }
            case -7:
                switch ("\n\t" + "foo bar") {
                case "a":
                    var v22 int = -1
                    _ = v22
                default:
                    var v23 float64 = 3.14159
                    _ = v23
                }
            case 128:
                v3 += v2
            default:
                var v24 int32 = -1
                _ = v24
            }
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 int64 = v7
                v3 -= v3
                _ = fi25
                _ = v26
            }
            v1 = (v1 && (!(v1 && (v1 && v1))))
        }
        switch ((("\n\t" + ((("\n\t" + "hello") + "hello") + "a")) + "\n\t") + ("a" + "\n\t")) {
        case "\n\t":
            v8 = append(v8, 672)
        case "a":
            _ = len(v8)
        case "foo bar":
            for fi27 := 0; fi27 < 6; fi27++ {
                v3 -= v3
                _ = fi27
            }
        default:
            for fi28 := 0; fi28 < 10; fi28++ {
                v29 := GenS4{F0: 2, F1: (v7 | (^(v7)))}
                v1 = false
                for fi30 := 0; fi30 < 17; fi30++ {
                    var v31 int8 = 60
                    v31 = 1
                    var v32 uint = 775
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = fi28
                _ = v29
            }
        }
        var v33 string = ("" + "hello")
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
