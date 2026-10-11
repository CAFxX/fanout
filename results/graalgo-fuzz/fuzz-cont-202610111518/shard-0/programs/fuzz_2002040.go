package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 string) int16 {
    if (!(!(9223372036854775807 != 360))) {
        var v5 int = -1
        var v6 string = "hello"
        var v7 bool = true
        _ = v5
        _ = v6
        _ = v7
    } else {
        v8 := []int{32767}
        var v9 float64 = -2.25
        fmt.Println(v9, p0)
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            p0 = ("foo bar" + "")
            _ = fi11
        }
        var v12 uint = 2
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint64 = 42
            _ = fi13
            _ = v14
        }
        _ = fi10
        _ = v12
    }
    _ = p0
    return 1
}

func main() {
    var v1 int = 9223372036854775807
    var v2 int32 = -671
    v3 := make(map[uint32]string)
    var v4 uint = 0
    if true {
        var v15 int16 = -803
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := GenS1{F0: false}
            if false {
                for fi18 := 0; fi18 < 4; fi18++ {
                    v19 := GenS1{F0: ((((1e10 * (((0.0 - (((1e10 - 1e10) * (-2.25 * -2.25)) * ((((-2.25 + ((1.5 / 3.14159) - 0.1)) / (0.1 + 0.0)) / float32(1.5)) - 0.0))) / ((3.14159 * 1.5) + ((float32(-2.25) / (-2.25 + (1e10 - 3.14159))) * 1.5))) * 0.0)) != (0.0 * ((((3.14159 + 1e10) - 0.0) / (((1.5 * 1.5) + (1.5 + (1.5 / (0.0 + 0.1)))) / (0.1 * 1.5))) - 1.5))) && (-107 > 1)) || false)}
                    _ = fi18
                    _ = v19
                }
                v20 := GenS1{F0: (!true)}
                _ = v20
            } else {
                fmt.Println(v2)
                _ = len(v3)
                v4 = v4
            }
            var v21 int16 = (v15 << ((2) & 63))
            _ = fi16
            _ = v17
            _ = v21
        }
        _ = len(v3)
        _ = v15
    }
    v2 += 0
    for fi22 := 0; fi22 < 8; fi22++ {
        if (true || (!(255 > 128))) {
            v23 := []int{}
            _ = v23
        }
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 uint16 = 994
            _ = fi24
            _ = v25
        }
        _ = fi22
    }
    v1 = (v1 % ((v1) | 1))
    var v26 int64 = -799
    var v27 uint8 = 158
    for fi28 := 0; fi28 < 10; fi28++ {
        if false {
            switch "" {
            case "a":
                v29 := GenS1{F0: true}
                _ = v29
            default:
                fmt.Println(fi28)
            }
            var v30 int16 = 0
            for fi31 := 0; fi31 < 16; fi31++ {
                v26 -= (v26 * 42)
                var v32 uint16 = 0
                v1 = fi31
                v33 := v3[4294967295]
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = v30
        }
        for fi34 := 0; fi34 < 18; fi34++ {
            v1 += v1
            _ = fi34
        }
        _ = fi28
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        v2 = v2
        fmt.Println(v4, fi35, v4)
        var v36 int16 = -333
        _ = fi35
        _ = v36
    }
    v2 -= (-(v2))
    var v37 int32 = 0
    _ = len(v3)
    fmt.Println(v27, v1, v27)
    v27 += 156
    var v38 uint16 = 122
    v39 := []int{v1, (v1 * (v1 + (^(v1))))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
