package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
type GenS2 struct { F0 int64; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 uint }
func genFn0(p0 string) int16 {
    var v7 uint32 = 15
    v8 := GenS1{F0: 65535, F1: 0, F2: (!(false || (((1e10 - ((0.0 - (0.1 - 0.0)) - 1e10)) > ((((((((1e10 - 3.14159) * ((3.14159 - -2.25) / 0.1)) - 0.1) - 1e10) / 1.0) + ((0.0 / 1e10) * 0.1)) * (1e10 * (float32(1e10) + (1e10 - -2.25)))) / ((((1e10 - float32(float64(0.1))) + 3.14159) / 3.14159) + -2.25))) || (true || (100 != 1)))))}
    v7 = v7
    v9 := GenS1{F0: 42, F1: 30, F2: (false && ((v7 / ((v7) | 1)) >= v7))}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1() uint {
    v10 := []int{-398}
    var v11 uint8 = 101
    switch (("foo bar" + "hello") + "foo bar") {
    case "hello":
        v12 := GenS2{F0: -1, F1: (2147483647 >= 448), F2: 32767}
        _ = v12
    case "foo bar":
        for fi13 := 0; fi13 < 1; fi13++ {
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 uint8 = (v11 & (^(v11)))
                v15 += (v15 / ((v11) | 1))
                _ = fi14
                _ = v15
            }
            v11 = (v11 & 127)
            _ = fi13
        }
    }
    _ = v10
    _ = v11
    return 65535
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint8 = 0
    v3 := []int{1, 40}
    var v4 int8 = 1
    v5 := make(map[uint32]bool)
    var v6 uint16 = 65535
    var v16 float64 = 0.1
    v3 = append(v3, 127)
    v3 = append(v3, -740)
    if true {
        v17 := GenS3{F0: 575}
        v5[2] = (3.14159 != float32((v16 + float64(3.14159))))
        v6 = (v6 >> (((v1 - (^(v1)))) & 63))
        _ = v17
    }
    var v18 int8 = 40
    v4 += (v4 << ((358) & 63))
    fmt.Println(v16)
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 19; fi20++ {
            switch "" {
            case "a":
                v18 = v18
            case "":
                v5[982] = true
            default:
                v4 = (v4 & v18)
            }
            if false {
                _ = len(v3)
                var v21 int8 = 100
                if false {
                    v21 = v21
                    v3 = append(v3, (^(fi19)))
                }
                _ = v21
            }
            var v22 uint32 = 18
            _ = fi20
            _ = v22
        }
        fmt.Println(v18)
        var v23 uint32 = 2147483647
        _ = fi19
        _ = v23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        fmt.Println(fi24, fi24, v16)
        var v25 uint8 = v2
        for fi26 := 0; fi26 < 1; fi26++ {
            v16 += v16
            switch "" {
            case "\n\t":
                var v27 float64 = ((v16 * (0.0 - v16)) - v16)
                _ = v27
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v28 := GenS4{F0: -103, F1: 924}
                _ = v28
            }
            v1 += 131
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    v29 := GenS5{F0: genFn0(("" + "hello"))}
    v30 := v5[42]
    var v31 bool = (42 == v2)
    v5[4294967295] = (42 == -628)
    fmt.Println(v18, v6)
    if (0.1 == (0.1 - 0.1)) {
        switch "foo bar" {
        case "a":
            fmt.Println(v4)
        }
        v32 := GenS4{F0: 32767, F1: 1000}
        for fi33 := 0; fi33 < 18; fi33++ {
            switch -1 {
            case 2147483647:
                v3 = append(v3, (fi33 << ((255) & 63)))
            case -191:
                v34 := GenS6{F0: -1, F1: 47}
                _ = v34
            }
            if ((9223372036854775807 == 42) || v31) {
                var v35 int32 = 1
                var v36 int8 = -1
                v35 -= (^(v35))
                _ = v35
                _ = v36
            } else {
                for fi37 := 0; fi37 < 3; fi37++ {
                    fmt.Println(v4, v31, v4)
                    fmt.Println(v31, v31)
                    _ = fi37
                }
            }
            _ = fi33
        }
        _ = v32
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        fmt.Println(v2)
        _ = fi38
    }
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
