package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint64; F2 bool }
func genFn0() bool {
    var v5 uint16 = 42
    switch "a" {
    case "a":
        v5 -= 65535
    default:
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: 42, F1: 0}
            _ = fi6
            _ = v7
        }
    }
    var v8 int32 = 1
    var v9 int = 482
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: (v8 << ((231) & 63)), F1: 0}
        if false {
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := []int{}
                switch "" {
                case "foo bar":
                    var v14 float32 = 3.14159
                    _ = v14
                case "":
                    v15 := GenS2{F0: 0}
                    _ = v15
                case "hello":
                    if (3.14159 <= 0.0) {
                        switch "a" {
                        case "foo bar":
                            switch (fi10 * fi12) {
                            case 0:
                                fmt.Println(fi10, v5, fi10)
                            case -221:
                                if len(v13) > 0 {
                                    v13[0] = int(v8)
                                }
                            case 388:
                                v8 += (^(v8))
                            }
                        case "\n\t":
                            var v16 int16 = 1
                            _ = v16
                        default:
                            var v17 uint16 = (^(v5))
                            _ = v17
                        }
                        switch "" {
                        case "a":
                            fmt.Println(v5)
                        }
                        var v18 uint8 = 120
                        _ = v18
                    } else {
                        v19 := GenS1{F0: v8, F1: 1000}
                        v8 = (v8 / ((v8) | 1))
                        _ = v19
                    }
                }
                fmt.Println(fi12, fi12)
                fmt.Println(fi12)
                var v20 int = fi12
                _ = fi12
                _ = v13
                _ = v20
            }
            var v21 string = "foo bar"
            _ = v21
        } else {
            switch ("hello" + "\n\t") {
            case "foo bar":
                v22 := []int{(fi10 % (((v9 % (((fi10 ^ -745)) | 1))) | 1))}
                _ = v22
            case "hello":
                if false {
                    v23 := GenS2{F0: -1}
                    _ = v23
                }
            }
            var v24 string = ("" + ("" + "hello"))
            v8 -= (v8 | -328)
            _ = v24
        }
        var v25 uint = 18446744073709551615
        _ = fi10
        _ = v11
        _ = v25
    }
    _ = v5
    _ = v8
    _ = v9
    return (true || ((v9 - (v9 | (-(v9)))) != 1))
}

func main() {
    var v1 float64 = ((float64(float32(0.1)) * ((1.5 + (3.14159 * 3.14159)) * 1.5)) / 1.5)
    var v2 int64 = -9223372036854775808
    var v3 bool = ((!(-2.25 <= 0.1)) || (!(127 > 208)))
    var v4 int32 = -165
    v26 := GenS2{F0: 0}
    for fi27 := 0; fi27 < 14; fi27++ {
        v28 := GenS3{F0: 1, F1: 4294967295, F2: (!v3)}
        v3 = true
        _ = fi27
        _ = v28
    }
    v29 := []int{int(v2)}
    fmt.Println(v4)
    switch 0 {
    case 0:
        for fi30 := 0; fi30 < 12; fi30++ {
            var v31 int = fi30
            var v32 uint32 = 1
            _ = fi30
            _ = v31
            _ = v32
        }
    case 32767:
        var v33 uint32 = 2
        _ = v33
    case -2147483648:
        for fi34 := 0; fi34 < 11; fi34++ {
            v1 = v1
            var v35 bool = v3
            v2 = 1
            _ = fi34
            _ = v35
        }
    }
    var v36 int32 = v4
    var v37 uint8 = 128
    fmt.Println(v36, v36)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
