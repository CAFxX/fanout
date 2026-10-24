package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int }
func genFn0() uint16 {
    var v4 uint64 = 2
    var v5 uint = 2
    v6 := GenS1{F0: 65535}
    v5 = v5
    v4 -= 18446744073709551615
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 uint8) string {
    v7 := GenS2{F0: 42, F1: (p0 + (^(p0)))}
    var v8 string = "hello"
    _ = p0
    _ = v7
    _ = v8
    return (((((("\n\t" + (v8 + v8)) + "\n\t") + ((("a" + "a") + ("hello" + ("hello" + v8))) + v8)) + ("a" + v8)) + v8) + "")
}

func main() {
    var v1 int64 = 497
    var v2 uint = 629
    v3 := make(map[int]bool)
    var v9 string = "\n\t"
    if true {
        var v10 int8 = 127
        _ = v10
    } else {
        if (!(128 <= 42)) {
            v11 := GenS1{F0: 255}
            _ = v11
        } else {
            if ((("\n\t" + "hello") + "hello") != v9) {
                v12 := GenS3{F0: 311, F1: 127}
                v13 := []int{}
                if len(v13) > 0 {
                    v13[0] = -324
                }
                _ = v12
                _ = v13
            } else {
                v14 := GenS1{F0: 453}
                fmt.Println(v1)
                if true {
                    v1 -= (v1 ^ (-(v1)))
                    var v15 int32 = 598
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        v1 = (v1 ^ 26)
                        if ((-32768 < -255) && (!(128 >= 149))) {
                            v3[(^(fi16))] = true
                            for fi17 := 0; fi17 < 7; fi17++ {
                                var v18 float64 = 1e10
                                if (935 <= 0) {
                                    v2 -= v2
                                }
                                v18 = v18
                                _ = fi17
                                _ = v18
                            }
                        }
                        _ = fi16
                    }
                    if true {
                        v19 := GenS1{F0: 843}
                        _ = v19
                    }
                }
                _ = v14
            }
            v20 := GenS2{F0: 522, F1: 0}
            _ = v20
        }
        var v21 uint16 = 737
        var v22 float32 = 3.14159
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        switch fi23 {
        case 2147483647:
            var v24 int8 = 46
            _ = v24
        default:
            switch fi23 {
            case 33:
                var v25 float64 = float64(3.14159)
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 9; fi26++ {
            if (fi26 > fi23) {
                var v27 int64 = (v1 & v1)
                _ = v27
            }
            _ = fi26
        }
        v1 = (^(v1))
        _ = fi23
    }
    v28 := GenS2{F0: uint32(v2), F1: 255}
    var v29 int8 = 1
    switch "" {
    case "a":
        var v30 string = v9
        _ = v30
    case "hello":
        v31 := GenS1{F0: 42}
        _ = v31
    case "\n\t":
        fmt.Println(v29, v9)
    default:
        fmt.Println(v1)
    }
    switch (v9 + "foo bar") {
    case "hello":
        var v32 uint32 = 1
        _ = v32
    case "a":
        var v33 int32 = 1
        _ = v33
    default:
        var v34 string = "a"
        _ = v34
    }
    v35 := GenS3{F0: 0, F1: 624}
    v36 := GenS3{F0: uint64(v2), F1: 127}
    fmt.Println(v2, v29, v1)
    fmt.Println(v1, v9, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v9
}
