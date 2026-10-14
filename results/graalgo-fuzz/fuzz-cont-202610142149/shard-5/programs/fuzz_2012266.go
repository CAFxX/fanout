package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint16) string {
    p0 = (0 == 42)
    v7 := []int{-1, -9223372036854775808, 32767}
    var v8 int32 = 2147483647
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return ("\n\t" + "\n\t")
}
func genFn1(p0 uint, p1 uint, p2 bool) int32 {
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            p1 -= p1
            v11 := GenS1{F0: "foo bar"}
            fmt.Println(fi10, p2)
            _ = fi10
            _ = v11
        }
        fmt.Println(fi9)
        _ = fi9
    }
    v12 := []int{788, -798, -1}
    var v13 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return -953
}

func main() {
    var v1 int16 = 1000
    var v2 int16 = 32767
    var v3 string = "foo bar"
    var v4 int16 = v2
    var v5 int = 0
    v6 := make(map[int]int)
    v4 -= 32767
    if false {
        for fi14 := 0; fi14 < 3; fi14++ {
            for fi15 := 0; fi15 < 7; fi15++ {
                _ = len(v6)
                _ = fi15
            }
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 uint16 = 427
                var v18 bool = true
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi14
        }
        for fi19 := 0; fi19 < 10; fi19++ {
            if true {
                delete(v6, v5)
                switch ("" + v3) {
                case "foo bar":
                    var v20 uint64 = 0
                    _ = v20
                default:
                    var v21 int = (^(v5))
                    _ = v21
                }
                var v22 uint16 = 65535
                _ = v22
            }
            _ = fi19
        }
    }
    v23 := GenS1{F0: (v3 + (v3 + v3))}
    v24 := v6[(v5 - v5)]
    if (!(!false)) {
        fmt.Println(v3, v3, v24)
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 string = "foo bar"
        _ = fi25
        _ = v26
    }
    switch (v24 - 42) {
    case 1:
        v5 = 4294967295
    case 238:
        for fi27 := 0; fi27 < 14; fi27++ {
            switch (fi27 & (^(v24))) {
            case 42:
                var v28 int8 = -1
                _ = v28
            case 128:
                if true {
                    fmt.Println(v5, v5, v1)
                    v24 -= (^(v5))
                    v29 := v6[(v24 / (((fi27 % ((128) | 1))) | 1))]
                    _ = v29
                } else {
                    switch "" {
                    case "foo bar":
                        v30 := GenS1{F0: ("foo bar" + "foo bar")}
                        _ = v30
                    case "a":
                        var v31 int8 = 1
                        _ = v31
                    case "\n\t":
                        var v32 int8 = -128
                        _ = v32
                    default:
                        var v33 bool = (!((v1 >> ((151) & 63)) != (v1 + (-(v2)))))
                        _ = v33
                    }
                }
            }
            _ = fi27
        }
    default:
        v34 := GenS1{F0: ((((v3 + v3) + "hello") + (v3 + v3)) + v3)}
        _ = v34
    }
    if false {
        if (0.0 <= float32(-2.25)) {
            v2 -= 0
        } else {
            if (!(!false)) {
                if true {
                    v4 -= (v1 / (((v1 ^ -348)) | 1))
                    for fi35 := 0; fi35 < 5; fi35++ {
                        v36 := GenS1{F0: v3}
                        v37 := GenS2{F0: v3, F1: 1, F2: 100}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                } else {
                    v4 = -487
                    var v38 int8 = 11
                    var v39 int32 = -691
                    _ = v38
                    _ = v39
                }
                for fi40 := 0; fi40 < 8; fi40++ {
                    v41 := GenS3{F0: 2}
                    v3 = v3
                    _ = fi40
                    _ = v41
                }
                var v42 uint64 = 2
                _ = v42
            }
            var v43 int32 = -2147483648
            _ = v43
        }
        var v44 uint64 = 4294967295
        _ = v44
    }
    v45 := []int{v5, (v24 ^ v5)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
