package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 float32 = 3.14159
    var v2 string = "\n\t"
    var v3 string = ""
    var v4 uint = 65535
    var v5 uint = 2
    var v6 float64 = ((1.5 * 1e10) * (0.1 + (3.14159 - ((0.0 + (0.1 + 0.0)) + ((float64(v1) + -2.25) - 3.14159)))))
    v1 = (v1 / float32((v6 * (((0.1 * 1e10) - v6) / v6))))
    if true {
        v7 := []int{}
        switch 1 {
        case 4294967295:
            switch "a" {
            case "hello":
                fmt.Println(v2, v3)
            case "\n\t":
                var v8 uint = (v4 + (v4 | v5))
                _ = v8
            default:
                switch v2 {
                case "":
                    for fi9 := 0; fi9 < 19; fi9++ {
                        var v10 float32 = v1
                        if false {
                            var v11 uint16 = 65535
                            _ = v11
                        } else {
                            for fi12 := 0; fi12 < 14; fi12++ {
                                if len(v7) > 0 {
                                    v7[0] = 42
                                }
                                v4 = 1
                                _ = fi12
                            }
                            fmt.Println(fi9, fi9)
                        }
                        _ = fi9
                        _ = v10
                    }
                case "\n\t":
                    _ = len(v7)
                }
            }
        case 32767:
            var v13 int8 = 1
            _ = v13
        case 32768:
            if len(v7) > 0 {
                v7[0] = 42
            }
        default:
            v14 := GenS1{F0: -456, F1: 1}
            _ = v14
        }
        _ = v7
    } else {
        if true {
            v15 := GenS2{F0: -1, F1: 1000}
            _ = v15
        }
        v16 := []int{63, 9223372036854775807}
        _ = v16
    }
    var v17 int32 = 1
    var v18 int = 1
    var v19 uint8 = 128
    if true {
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := []int{fi20, v18}
            var v22 int16 = 42
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 uint8 = (v19 | (v19 ^ (v19 ^ (v19 & v19))))
                _ = fi23
                _ = v24
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        v6 = (v6 * v6)
    }
    if true {
        var v25 int8 = -87
        for fi26 := 0; fi26 < 6; fi26++ {
            switch ("a" + v3) {
            case "foo bar":
                fmt.Println(v2)
            case "a":
                for fi27 := 0; fi27 < 9; fi27++ {
                    v25 = -1
                    _ = fi27
                }
            }
            var v28 int16 = -965
            for fi29 := 0; fi29 < 13; fi29++ {
                v3 = ("\n\t" + v3)
                _ = fi29
            }
            _ = fi26
            _ = v28
        }
        _ = v25
    } else {
        var v30 float64 = ((v6 - 3.14159) + ((v6 / (v6 / 3.14159)) + -2.25))
        _ = v30
    }
    v31 := []int{(v18 | v18), v18}
    var v32 string = v2
    v33 := GenS2{F0: 42, F1: 792}
    var v34 bool = false
    switch (("hello" + "") + ((v3 + v32) + v32)) {
    case "\n\t":
        var v35 int64 = 0
        _ = v35
    case "hello":
        v34 = (!v34)
    case "":
        for fi36 := 0; fi36 < 14; fi36++ {
            var v37 float64 = -2.25
            v34 = (65535 < 105)
            _ = fi36
            _ = v37
        }
    default:
        if (!v34) {
            v6 = ((v6 - float64(v1)) * float64(v1))
        } else {
            if len(v31) > 0 {
                v31[0] = (v18 | (v18 % (((v18 & v18)) | 1)))
            }
        }
    }
    var v38 int32 = 996
    switch v18 {
    case 433:
        v32 = ""
    case 2147483647:
        v6 = (v6 - -2.25)
    case 1:
        v38 += v17
    default:
        for fi39 := 0; fi39 < 16; fi39++ {
            fmt.Println(v6)
            if len(v31) >= 2 {
                v31 = v31[1:len(v31)]
            }
            v40 := GenS2{F0: 42, F1: 1}
            _ = fi39
            _ = v40
        }
    }
    v41 := GenS3{F0: 148}
    if len(v31) > 0 {
        v31[0] = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
