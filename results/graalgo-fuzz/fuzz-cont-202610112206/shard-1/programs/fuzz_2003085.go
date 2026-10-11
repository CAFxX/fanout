package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 string = "hello"
    v2 := []int{}
    var v3 string = "foo bar"
    var v4 string = (v3 + v1)
    var v5 int = 128
    v6 := GenS1{F0: 2, F1: 42}
    v2 = append(v2, v5)
    v2 = append(v2, (v5 * v5))
    v2 = append(v2, 1)
    fmt.Println(v1)
    fmt.Println(v5, v1, v5)
    if true {
        v5 = 32768
        if (!(737 > 150)) {
            v7 := GenS1{F0: 283, F1: 2}
            for fi8 := 0; fi8 < 10; fi8++ {
                v4 = v4
                var v9 uint32 = 399
                v3 = v4
                _ = fi8
                _ = v9
            }
            _ = v7
        }
        if (42 <= 2) {
            var v10 string = v4
            var v11 int = (v5 & v5)
            _ = v10
            _ = v11
        } else {
            if len(v2) > 0 {
                v2[0] = (v5 / ((v5) | 1))
            }
            if len(v2) > 0 {
                v2[0] = (v5 | v5)
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    } else {
        var v12 bool = (((127 != 67) && (18446744073709551615 <= 4294967295)) || (3.14159 > 1.5))
        var v13 int64 = 42
        _ = v12
        _ = v13
    }
    var v14 float32 = ((1e10 - (0.0 * float32(-2.25))) + 3.14159)
    v15 := GenS1{F0: 2, F1: 139}
    for fi16 := 0; fi16 < 12; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            v5 += 42
            _ = len(v2)
            v1 = v3
            _ = fi17
        }
        var v18 bool = (502 < 255)
        var v19 int64 = -1
        _ = fi16
        _ = v18
        _ = v19
    }
    if false {
        v2 = append(v2, 278)
        switch 142 {
        case 1:
            v5 -= 127
        case 32767:
            v3 = (("hello" + v1) + "")
        case 9223372036854775807:
            switch "a" {
            case "a":
                if false {
                    var v20 int32 = 0
                    for fi21 := 0; fi21 < 2; fi21++ {
                        for fi22 := 0; fi22 < 20; fi22++ {
                            v23 := GenS1{F0: 42, F1: 1}
                            var v24 uint32 = 4294967295
                            var v25 float32 = v14
                            _ = fi22
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        _ = fi21
                    }
                    _ = v20
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (562 <= 556) {
                        v2 = append(v2, -1)
                    }
                    switch v1 {
                    case "foo bar":
                        v5 -= v5
                    }
                }
            default:
                v14 = 0.1
            }
        }
        for fi26 := 0; fi26 < 18; fi26++ {
            for fi27 := 0; fi27 < 14; fi27++ {
                v28 := GenS2{F0: 1}
                if len(v2) > 0 {
                    v2[0] = (^(fi26))
                }
                _ = fi27
                _ = v28
            }
            v29 := GenS2{F0: 42}
            _ = fi26
            _ = v29
        }
    }
    var v30 int16 = 1
    fmt.Println(v30)
    var v31 int = (v5 ^ (v5 + (v5 | (v5 + (^(v5))))))
    var v32 bool = ((-873 < 220) && (v4 != ""))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
