package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 bool }
type GenS2 struct { F0 uint8; F1 uint32; F2 string }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 uint16, p1 int32) bool {
    var v4 float32 = (1e10 - (1.5 / 1.0))
    var v5 float64 = 3.14159
    v6 := GenS1{F0: 255, F1: 368, F2: true}
    if false {
        v4 -= v4
        v7 := []int{42, int(p1)}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return (831 >= 32768)
}
func genFn1(p0 int64, p1 int, p2 uint64) string {
    v8 := []int{(-(p1)), p1}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v9 int64 = p0
    if (false && false) {
        var v10 int16 = 970
        _ = v10
    }
    fmt.Println(v9, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    var v1 uint8 = 2
    var v2 string = "\n\t"
    var v3 int = 9223372036854775807
    var v11 uint64 = 76
    v2 = "\n\t"
    var v12 int = v3
    if true {
        switch int(v11) {
        case 127:
            for fi13 := 0; fi13 < 4; fi13++ {
                if true {
                    v14 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: (!(((0.1 - -2.25) - 0.0) >= float64(-2.25)))}
                    _ = v14
                } else {
                    v15 := GenS2{F0: v1, F1: 1, F2: v2}
                    var v16 uint16 = 255
                    for fi17 := 0; fi17 < 9; fi17++ {
                        v18 := GenS1{F0: 685, F1: 1, F2: false}
                        var v19 float32 = 3.14159
                        var v20 uint64 = v11
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    _ = v15
                    _ = v16
                }
                _ = fi13
            }
        default:
            var v21 bool = false
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := []int{(fi22 / (((v3 - (v12 - 1))) | 1))}
        for fi24 := 0; fi24 < 2; fi24++ {
            v3 = (fi22 - v12)
            _ = fi24
        }
        v1 += 250
        _ = fi22
        _ = v23
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := []int{}
        var v27 int16 = 1000
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        v29 := []int{2147483647, (v12 & v12)}
        v2 = v2
        var v30 float64 = 0.1
        _ = fi28
        _ = v29
        _ = v30
    }
    v31 := GenS2{F0: v1, F1: 0, F2: v2}
    fmt.Println(v2)
    v32 := GenS1{F0: 1, F1: 1, F2: (!false)}
    v3 -= (^(v3))
    v33 := GenS3{F0: 32767, F1: 42}
    for fi34 := 0; fi34 < 13; fi34++ {
        v11 = (v11 * (v11 % ((4294967295) | 1)))
        _ = fi34
    }
    var v35 float64 = 0.1
    var v36 bool = ((65535 == uint32(v1)) || ((v2 + (v2 + v2)) >= ("" + v2)))
    for fi37 := 0; fi37 < 13; fi37++ {
        var v38 float32 = 3.14159
        v38 -= float32((v35 - 0.0))
        switch v3 {
        case -9223372036854775808:
            v39 := []int{fi37, fi37, (-(fi37))}
            _ = v39
        default:
            v11 = (v11 * 1)
        }
        _ = fi37
        _ = v38
    }
    for fi40 := 0; fi40 < 10; fi40++ {
        switch ((v2 + "foo bar") + v2) {
        case "a":
            v41 := []int{}
            _ = v41
        case "hello":
            switch v3 {
            case -833:
                var v42 uint32 = uint32(v11)
                _ = v42
            case -1:
                v11 = v11
            case 2147483647:
                switch v2 {
                case "":
                    v43 := []int{}
                    _ = v43
                case "\n\t":
                    v44 := GenS1{F0: 255, F1: 803, F2: v36}
                    _ = v44
                }
            default:
                v45 := []int{9223372036854775807, int(v11)}
                _ = v45
            }
        case "":
            fmt.Println(v35, v35)
        }
        if v36 {
            fmt.Println(fi40)
            var v46 uint32 = 4294967295
            _ = v46
        }
        _ = fi40
    }
    v1 -= 244
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v36
}
