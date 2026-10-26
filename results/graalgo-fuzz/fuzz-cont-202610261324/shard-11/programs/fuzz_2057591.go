package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8; F1 int8; F2 uint8 }
func genFn0() string {
    var v6 int32 = 42
    if (!(v6 < (v6 / ((2147483647) | 1)))) {
        v6 += 1
        var v7 uint = 1
        _ = v7
    }
    _ = v6
    return "hello"
}
func genFn1(p0 uint16, p1 int16) uint {
    var v8 string = "\n\t"
    if true {
        fmt.Println(p0, v8, p0)
    } else {
        p1 = (p1 / ((p1) | 1))
    }
    v9 := GenS1{F0: 42}
    p1 -= (p1 ^ -1)
    switch (v8 + ("" + v8)) {
    case "\n\t":
        var v10 uint = 922
        _ = v10
    case "foo bar":
        v11 := []int{9223372036854775807, -248, 127, 32768}
        _ = v11
    case "":
        v12 := GenS2{F0: 1}
        _ = v12
    default:
        if false {
            if true {
                var v13 int16 = 32767
                _ = v13
            } else {
                var v14 uint32 = 1
                var v15 uint32 = v14
                if (337 <= 1) {
                    switch v8 {
                    case "hello":
                        for fi16 := 0; fi16 < 17; fi16++ {
                            fmt.Println(v8, v14, v15)
                            switch (fi16 % ((fi16) | 1)) {
                            case 42:
                                v17 := []int{961, fi16, -463}
                                _ = v17
                            default:
                                p0 += 65535
                            }
                            v8 = ((v8 + (("\n\t" + v8) + v8)) + (v8 + "foo bar"))
                            _ = fi16
                        }
                    default:
                        v18 := []int{-9223372036854775808}
                        _ = v18
                    }
                    p0 = 937
                    switch 0 {
                    case -1:
                        v19 := []int{821, 116, 128, 5}
                        _ = v19
                    default:
                        v20 := []int{-477, 32767, -392, 128}
                        _ = v20
                    }
                }
                _ = v14
                _ = v15
            }
            v21 := GenS1{F0: 185}
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 964
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{180}
    var v3 uint16 = 0
    var v4 uint64 = 42
    v5 := make(map[string]bool)
    var v22 string = "foo bar"
    v23 := GenS2{F0: 2147483647}
    v24 := GenS3{F0: 82}
    fmt.Println(v4, v4)
    v25 := GenS4{F0: 543}
    if false {
        var v26 int = 32767
        v4 = v4
        for fi27 := 0; fi27 < 5; fi27++ {
            switch v26 {
            case 760:
                v28 := GenS5{F0: 0}
                _ = v28
            case 538:
                v3 -= (v3 ^ (v3 - 65535))
            case 9223372036854775807:
                if len(v2) > 0 {
                    v2[0] = fi27
                }
            default:
                var v29 uint = 255
                _ = v29
            }
            var v30 float64 = (0.0 * -2.25)
            _ = fi27
            _ = v30
        }
        _ = v26
    }
    switch 1 {
    case -2147483648:
        for fi31 := 0; fi31 < 5; fi31++ {
            switch (-(fi31)) {
            case -49:
                v32 := GenS6{F0: (v3 ^ v3)}
                _ = v32
            default:
                var v33 float64 = (3.14159 + 0.0)
                _ = v33
            }
            v4 = (v4 % ((v4) | 1))
            _ = fi31
        }
    case 4294967295:
        v22 = "foo bar"
    }
    var v34 uint64 = (v4 % ((v4) | 1))
    v22 = v22
    if true {
        for fi35 := 0; fi35 < 10; fi35++ {
            fmt.Println(v34)
            _ = fi35
        }
        delete(v5, v22)
    }
    _ = len(v2)
    switch (v22 + v22) {
    case "hello":
        var v36 uint16 = (v3 | v3)
        _ = v36
    default:
        switch 9223372036854775807 {
        case 1:
            if false {
                v34 = v4
            }
        default:
            if (true || (100 == -128)) {
                _ = len(v5)
                for fi37 := 0; fi37 < 1; fi37++ {
                    for fi38 := 0; fi38 < 17; fi38++ {
                        v39 := GenS7{F0: uint8(v3), F1: -97, F2: 158}
                        v3 += 32767
                        v40 := GenS1{F0: (^(v4))}
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    var v41 int16 = 42
                    var v42 float64 = (((1.5 - -2.25) * 1.5) * -2.25)
                    _ = fi37
                    _ = v41
                    _ = v42
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
