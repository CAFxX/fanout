package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int32; F1 bool; F2 uint32 }
func genFn0(p0 int64, p1 int64) int16 {
    v5 := []int{4294967295, 895, -612}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 := GenS1{F0: 1, F1: 4294967295}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -32768
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{}
    v3 := []int{128, -600}
    v4 := make(map[string]bool)
    if false {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS2{F0: ("hello" + "foo bar"), F1: 0}
            _ = fi7
            _ = v8
        }
    } else {
        var v9 uint32 = 42
        _ = v9
    }
    if (777 < -1) {
        var v10 string = "foo bar"
        _ = v10
    }
    var v11 uint32 = 42
    v12 := GenS1{F0: 210, F1: 0}
    for fi13 := 0; fi13 < 17; fi13++ {
        if (3.14159 > 0.1) {
            delete(v4, "hello")
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if false {
                fmt.Println(v11)
            } else {
                fmt.Println(fi13)
                var v14 float32 = 0.1
                _ = v14
            }
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 uint32 = v11
                _ = fi15
                _ = v16
            }
        }
        _ = fi13
    }
    switch 43 {
    case 892:
        for fi17 := 0; fi17 < 12; fi17++ {
            v11 = 2147483648
            _ = fi17
        }
    case 701:
        if (!(true || false)) {
            for fi18 := 0; fi18 < 17; fi18++ {
                for fi19 := 0; fi19 < 15; fi19++ {
                    fmt.Println(v11, v11)
                    switch fi18 {
                    case 1:
                        v2 = append(v2, "foo bar")
                        v11 -= 2147483648
                        var v20 uint8 = 78
                        _ = v20
                    case -315:
                        var v21 string = "a"
                        _ = v21
                    case -9223372036854775808:
                        v11 -= v11
                    }
                    var v22 uint8 = 127
                    _ = fi19
                    _ = v22
                }
                _ = fi18
            }
        }
    case -2147483648:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println(v11)
    if true {
        var v23 int32 = -1
        v24 := GenS3{F0: v23, F1: ((v23 * v23) > v23), F2: v11}
        fmt.Println(v23)
        _ = v23
        _ = v24
    }
    v25 := GenS2{F0: "foo bar", F1: 42}
    for fi26 := 0; fi26 < 18; fi26++ {
        for fi27 := 0; fi27 < 1; fi27++ {
            fmt.Println(v11)
            v11 += (v11 >> ((18446744073709551615) & 63))
            var v28 int64 = 190
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    var v29 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
}
