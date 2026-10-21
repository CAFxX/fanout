package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string }
func genFn0(p0 int8, p1 int32, p2 uint8) int64 {
    fmt.Println(p0, p0, p0)
    v6 := []int{4294967295, 0, -747}
    if true {
        fmt.Println(p0, p2)
    } else {
        v7 := GenS1{F0: 587, F1: 2, F2: ((3.14159 * -2.25) != 0.0)}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return int64(p1)
}

func main() {
    v1 := []uint32{1}
    var v2 uint32 = 604
    var v3 uint16 = 1
    var v4 string = (("\n\t" + "\n\t") + ("a" + "hello"))
    v5 := []int{}
    _ = len(v1)
    v3 += (^(v3))
    v2 = v2
    if ((("foo bar" > ("foo bar" + (v4 + "foo bar"))) && false) || false) {
        v8 := GenS2{F0: 9223372036854775807}
        var v9 uint8 = 128
        _ = v8
        _ = v9
    } else {
        v10 := GenS3{F0: -94}
        if false {
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 int32 = 1
                switch v4 {
                case "hello":
                    switch (v4 + "") {
                    case "a":
                        var v13 uint8 = 232
                        _ = v13
                    default:
                        var v14 int8 = 42
                        _ = v14
                    }
                case "\n\t":
                    var v15 float64 = 0.0
                    _ = v15
                case "":
                    v16 := GenS4{F0: (v12 + (v12 * -779))}
                    _ = v16
                }
                _ = fi11
                _ = v12
            }
        }
        _ = v10
    }
    switch -1 {
    case 614:
        v3 = (v3 / (((v3 << ((65535) & 63))) | 1))
    }
    if (false || (!true)) {
        var v17 int16 = 1
        for fi18 := 0; fi18 < 19; fi18++ {
            if len(v1) > 0 {
                v1[0] = 453
            }
            fmt.Println(fi18, fi18, v2)
            var v19 int16 = 1000
            _ = fi18
            _ = v19
        }
        v3 -= (^(v3))
        _ = v17
    }
    if true {
        fmt.Println(v4, v2)
        fmt.Println(v2, v2, v3)
    } else {
        if (!(0 < 42)) {
            v20 := GenS5{F0: (v4 + v4)}
            _ = v20
        } else {
            for fi21 := 0; fi21 < 17; fi21++ {
                switch fi21 {
                case 9223372036854775807:
                    for fi22 := 0; fi22 < 6; fi22++ {
                        v1 = append(v1, uint32(v3))
                        if (!false) {
                            if len(v1) > 0 {
                                v1[0] = (^(v2))
                            }
                            switch fi21 {
                            case 0:
                                var v23 int64 = 9223372036854775807
                                _ = v23
                            }
                            var v24 int64 = 176
                            _ = v24
                        } else {
                            v5 = append(v5, int(v3))
                            v1 = append(v1, (v2 ^ v2))
                            var v25 uint8 = 0
                            _ = v25
                        }
                        _ = fi22
                    }
                }
                _ = fi21
            }
            v4 = v4
            _ = len(v5)
        }
        var v26 bool = (!(!(!(!true))))
        _ = v26
    }
    v2 -= v2
    for fi27 := 0; fi27 < 2; fi27++ {
        v4 = v4
        _ = fi27
    }
    v4 = ""
    fmt.Println(v3, v3)
    for fi28 := 0; fi28 < 4; fi28++ {
        switch v4 {
        case "hello":
            v4 = v4
        case "a":
            fmt.Println(v4, fi28)
        }
        _ = fi28
    }
    var v29 int16 = -457
    v5 = append(v5, 9223372036854775807)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
