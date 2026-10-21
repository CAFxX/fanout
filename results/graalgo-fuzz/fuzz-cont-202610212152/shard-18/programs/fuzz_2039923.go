package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint32, p1 string, p2 int32) uint8 {
    fmt.Println(p0, p2, p1)
    switch 9223372036854775807 {
    case 9223372036854775807:
        v6 := []int{32768, -715, 469}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}
func genFn1(p0 int8, p1 uint16, p2 int8) uint64 {
    p1 = p1
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint64 = 1
        v9 := GenS1{F0: 42, F1: (p1 - (^(p1)))}
        _ = fi7
        _ = v8
        _ = v9
    }
    p2 -= 6
    if true {
        if (0.1 <= ((((0.1 / (-2.25 * (1.5 * ((((((1e10 + 0.0) - 0.0) + 1.5) - (0.0 / -2.25)) * -2.25) * float32((float64(3.14159) - 0.0)))))) / 3.14159) - 1.5) / float32(1.5))) {
            v10 := []int{}
            if len(v10) > 0 {
                v10[0] = -693
            }
            _ = v10
        } else {
            p0 += p0
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{}
    v3 := []uint32{42, 42, 42}
    var v4 uint = 996
    var v5 string = "hello"
    var v11 string = (v5 + (v5 + v5))
    for fi12 := 0; fi12 < 7; fi12++ {
        switch (fi12 - fi12) {
        case 42:
            v3 = append(v3, 719)
        case -9223372036854775808:
            v3 = append(v3, 244)
        }
        v5 = "foo bar"
        _ = len(v3)
        _ = fi12
    }
    switch -809 {
    case 33:
        if true {
            _ = len(v3)
            if ((1 > 18446744073709551615) || (!(!false))) {
                v13 := GenS1{F0: 32767, F1: 2}
                var v14 uint32 = 42
                v1[(v14 + v14)] = (v14 > (v14 << (((v4 >> ((v4) & 63))) & 63)))
                _ = v13
                _ = v14
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v5 = v11
            for fi15 := 0; fi15 < 20; fi15++ {
                v5 = v11
                var v16 int16 = -1
                v17 := GenS1{F0: v16, F1: uint16(v4)}
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    case 32768:
        v18 := v1[65535]
        _ = v18
    case -9223372036854775808:
        for fi19 := 0; fi19 < 10; fi19++ {
            for fi20 := 0; fi20 < 15; fi20++ {
                fmt.Println(v5, fi20)
                _ = fi20
            }
            for fi21 := 0; fi21 < 3; fi21++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v22 bool = (!(float64(0.0) == float64(-2.25)))
                var v23 float32 = (float32(float64((1e10 / 0.1))) - 1.5)
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 string = v5
            _ = fi19
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        v26 := GenS2{F0: 4294967295, F1: 747}
        _ = fi25
        _ = v26
    }
    v5 = ("hello" + (v5 + v5))
    if true {
        v5 = v5
        var v27 uint32 = 2
        delete(v1, (v27 * (^(v27))))
        _ = v27
    }
    var v28 int16 = -32768
    var v29 uint8 = 131
    v30 := GenS1{F0: -1, F1: 32767}
    v28 = (^(v28))
    for fi31 := 0; fi31 < 6; fi31++ {
        var v32 float32 = ((1.5 + 3.14159) / 1.5)
        _ = fi31
        _ = v32
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    for fi33 := 0; fi33 < 9; fi33++ {
        v28 += v28
        v28 -= (v28 & v28)
        var v34 uint64 = 0
        _ = fi33
        _ = v34
    }
    if (!false) {
        var v35 int64 = -1
        _ = v35
    } else {
        var v36 float64 = -2.25
        _ = v36
    }
    var v37 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v5
}
