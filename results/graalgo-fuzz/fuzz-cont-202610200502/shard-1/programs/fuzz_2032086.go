package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 bool }
func genFn0(p0 uint32) int64 {
    fmt.Println(p0)
    var v5 uint8 = 0
    for fi6 := 0; fi6 < 17; fi6++ {
        switch fi6 {
        case 2147483647:
            switch ("hello" + "\n\t") {
            case "":
                for fi7 := 0; fi7 < 10; fi7++ {
                    var v8 uint16 = 32767
                    v9 := []int{}
                    v9 = append(v9, 1)
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            case "a":
                if true {
                    var v10 int16 = -32768
                    _ = v10
                } else {
                    fmt.Println(p0)
                    v5 += (v5 << ((uint(p0)) & 63))
                }
            case "\n\t":
                p0 += p0
            default:
                v5 = v5
            }
        case 0:
            v11 := []int{fi6}
            _ = v11
        default:
            v12 := []int{fi6, (fi6 << ((uint(p0)) & 63)), fi6, 42}
            _ = v12
        }
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 int64 = -9223372036854775808
            _ = fi13
            _ = v14
        }
        _ = fi6
    }
    v15 := GenS1{F0: true}
    v5 = v5
    _ = p0
    _ = v15
    _ = v5
    return -1
}

func main() {
    var v1 float64 = (1.5 / 3.14159)
    v2 := make(map[int]uint32)
    var v3 int8 = -57
    var v4 string = "\n\t"
    v16 := v2[32767]
    fmt.Println(v4)
    v3 = v3
    switch 127 {
    case 617:
        delete(v2, 42)
    case -9223372036854775808:
        if (2 > 127) {
            var v17 float64 = v1
            v17 = (((1e10 + -2.25) - (((float64(1e10) * 0.0) / v17) + float64((0.0 * 1e10)))) - (v1 + 3.14159))
            v4 = v4
            _ = v17
        } else {
            v2[4294967295] = v16
        }
    }
    v18 := GenS2{F0: (^(v16))}
    switch (v4 + "") {
    case "hello":
        var v19 float64 = 0.0
        _ = v19
    case "a":
        v20 := GenS3{F0: 0, F1: false}
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        switch "hello" {
        case "hello":
            var v22 int32 = 0
            _ = v22
        case "":
            var v23 int32 = -1
            _ = v23
        }
        delete(v2, (fi21 & 4294967295))
        v24 := []int{}
        _ = fi21
        _ = v24
    }
    if ((^(v16)) == 999) {
        for fi25 := 0; fi25 < 10; fi25++ {
            v26 := []int{fi25, fi25}
            for fi27 := 0; fi27 < 10; fi27++ {
                switch fi25 {
                case 198:
                    v3 += v3
                case 2147483647:
                    v1 = v1
                case -9223372036854775808:
                    v16 -= 42
                default:
                    v28 := GenS1{F0: true}
                    _ = v28
                }
                _ = len(v26)
                if ("\n\t" == "hello") {
                    var v29 float64 = ((v1 - (0.0 / float64(1e10))) - (v1 - v1))
                    v1 += (v1 - v29)
                    _ = v29
                } else {
                    v30 := GenS3{F0: 598, F1: (("" + (v4 + v4)) <= v4)}
                    _ = v30
                }
                _ = fi27
            }
            _ = fi25
            _ = v26
        }
        v31 := []int{}
        _ = v31
    } else {
        for fi32 := 0; fi32 < 9; fi32++ {
            var v33 float64 = v1
            _ = fi32
            _ = v33
        }
        for fi34 := 0; fi34 < 6; fi34++ {
            v35 := GenS1{F0: true}
            _ = fi34
            _ = v35
        }
        var v36 bool = (0 <= -1)
        _ = v36
    }
    var v37 int16 = 32767
    v38 := GenS1{F0: ((v37 | -881) > (v37 ^ 42))}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
