package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 uint8 }
func genFn0(p0 uint64, p1 int32, p2 uint64) uint8 {
    p1 = 2147483647
    for fi4 := 0; fi4 < 17; fi4++ {
        fmt.Println(p1)
        v5 := GenS1{F0: ("foo bar" + "foo bar")}
        _ = fi4
        _ = v5
    }
    var v6 int16 = -32768
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 250
}

func main() {
    var v1 int = 1
    var v2 uint = 42
    v3 := make(map[string]string)
    if false {
        if (true && (!("\n\t" < "a"))) {
            var v7 string = "hello"
            v8 := []int{v1, (v1 - v1), v1, 42}
            _ = v7
            _ = v8
        } else {
            _ = len(v3)
            var v9 int32 = 42
            v10 := GenS1{F0: "hello"}
            _ = v10
            _ = v9
        }
        var v11 int64 = 42
        for fi12 := 0; fi12 < 8; fi12++ {
            v11 -= (v11 / (((^(v11))) | 1))
            for fi13 := 0; fi13 < 11; fi13++ {
                fmt.Println(v2, v11)
                _ = fi13
            }
            for fi14 := 0; fi14 < 18; fi14++ {
                v15 := GenS1{F0: "hello"}
                _ = fi14
                _ = v15
            }
            _ = fi12
        }
        _ = v11
    } else {
        v1 += (v1 + v1)
        v16 := []int{v1, (v1 - (v1 * 128))}
        _ = v16
    }
    var v17 float64 = -2.25
    var v18 int = v1
    if false {
        v19 := GenS2{F0: 127}
        var v20 float32 = (-2.25 - 1.5)
        _ = v19
        _ = v20
    } else {
        switch "a" {
        case "a":
            v2 += (v2 ^ (v2 | v2))
        case "":
            for fi21 := 0; fi21 < 14; fi21++ {
                for fi22 := 0; fi22 < 13; fi22++ {
                    v23 := GenS3{F0: 42, F1: 2147483648, F2: 255}
                    fmt.Println(v18, v1)
                    _ = fi22
                    _ = v23
                }
                _ = fi21
            }
        default:
            if true {
                fmt.Println(v18, v2)
                if ((3.14159 / 1.0) < 1e10) {
                    var v24 int8 = 100
                    switch v1 {
                    case 32768:
                        var v25 int32 = 42
                        _ = v25
                    case 9223372036854775807:
                        v2 = v2
                    }
                    _ = v24
                }
                switch 1 {
                case 127:
                    var v26 int64 = 9223372036854775807
                    _ = v26
                case 652:
                    v17 += v17
                }
            } else {
                switch v1 {
                case 2147483647:
                    var v27 uint32 = 4294967295
                    _ = v27
                }
            }
        }
        if false {
            if false {
                v2 = 185
            }
            var v28 bool = (((v18 >> (((v2 << ((v2) & 63))) & 63)) == 284) || false)
            _ = v28
        } else {
            if true {
                v17 -= v17
            } else {
                var v29 int32 = 42
                v29 += (v29 - (v29 | (v29 + v29)))
                _ = v29
            }
            v1 -= 596
            v30 := GenS4{F0: ("\n\t" + "\n\t"), F1: false}
            _ = v30
        }
        var v31 bool = (false && (-2.25 >= 1e10))
        _ = v31
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        fmt.Println(v1)
        _ = fi32
    }
    fmt.Println(v17)
    fmt.Println(v17, v17, v1)
    fmt.Println(v1, v1)
    var v33 bool = ((v1 / (((-(v18))) | 1)) != (v18 % (((v1 % ((v1) | 1))) | 1)))
    delete(v3, "foo bar")
    v34 := GenS5{F0: 1}
    v35 := GenS6{F0: (v1 - v18), F1: 1}
    var v36 int32 = 42
    for fi37 := 0; fi37 < 20; fi37++ {
        v36 = 0
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
