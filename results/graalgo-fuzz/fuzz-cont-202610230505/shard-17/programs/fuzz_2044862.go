package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 int8; F2 uint32 }

func main() {
    var v1 string = ("hello" + (((("foo bar" + "foo bar") + "") + (((("\n\t" + ("foo bar" + "a")) + "foo bar") + ("a" + "hello")) + "foo bar")) + "\n\t"))
    var v2 int64 = 0
    var v3 uint64 = 4294967295
    v4 := []int{4294967295, 351, 32767, 32768}
    var v5 int8 = 72
    var v6 int64 = v2
    var v7 int32 = -2147483648
    v8 := GenS1{F0: 100, F1: (^(v3)), F2: 739}
    v9 := GenS1{F0: v5, F1: (v3 >> ((255) & 63)), F2: -1}
    v4 = append(v4, 0)
    v4 = append(v4, 128)
    if (!((float64(0.1) + 1e10) >= (-2.25 / (-2.25 * ((((float64((3.14159 * (0.0 - 3.14159))) + 0.0) - float64(1.5)) - -2.25) / 1.5))))) {
        var v10 float64 = 0.0
        _ = v10
    }
    v1 = ("foo bar" + "\n\t")
    v6 += (v6 - (v6 >> ((0) & 63)))
    v11 := GenS1{F0: v5, F1: v3, F2: 339}
    for fi12 := 0; fi12 < 11; fi12++ {
        if (((v2 ^ (v6 - 0)) > (v2 + (^(v6)))) && (0.1 >= -2.25)) {
            fmt.Println(v2, v3, v3)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v13 := GenS1{F0: int8(v6), F1: (v3 / ((v3) | 1)), F2: -2147483648}
            _ = v13
        } else {
            switch fi12 {
            case -1:
                v4 = append(v4, fi12)
            case 1:
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    var v14 int8 = 42
                    fmt.Println(v14, v6, v14)
                    _ = v14
                }
            case -2147483648:
                v3 = (v3 & v3)
            }
            var v15 int8 = 100
            _ = v15
        }
        v16 := GenS1{F0: -1, F1: 18446744073709551615, F2: fi12}
        _ = fi12
        _ = v16
    }
    var v17 uint16 = 2
    v4 = append(v4, -814)
    if (!true) {
        for fi18 := 0; fi18 < 19; fi18++ {
            _ = len(v4)
            v17 -= v17
            _ = fi18
        }
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(v17, fi19, v3)
            v20 := GenS2{F0: true}
            _ = fi19
            _ = v20
        }
        if len(v4) > 0 {
            v4[0] = 1
        }
    }
    _ = len(v4)
    if false {
        fmt.Println(v7, v6)
        for fi21 := 0; fi21 < 10; fi21++ {
            switch ("foo bar" + ((v1 + v1) + v1)) {
            case "":
                if true {
                    v22 := GenS2{F0: false}
                    var v23 uint16 = v17
                    switch v1 {
                    case "a":
                        v24 := GenS3{F0: 0}
                        _ = v24
                    case "\n\t":
                        v7 += -2147483648
                    default:
                        var v25 int8 = (v5 >> ((42) & 63))
                        _ = v25
                    }
                    _ = v22
                    _ = v23
                } else {
                    v26 := GenS4{F0: -673, F1: v5, F2: 648}
                    var v27 int8 = -27
                    _ = v26
                    _ = v27
                }
            }
            v3 = 0
            _ = fi21
        }
    }
    fmt.Println(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
