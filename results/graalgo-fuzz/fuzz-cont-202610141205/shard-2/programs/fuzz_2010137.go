package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
func genFn0(p0 uint, p1 string, p2 uint8) uint8 {
    var v4 float64 = 0.0
    v5 := []int{32768, -2147483648}
    if (!false) {
        p2 = p2
        switch -408 {
        case 128:
            var v6 float32 = ((float32(v4) - -2.25) / (0.1 + 0.1))
            _ = v6
        default:
            var v7 uint8 = p2
            _ = v7
        }
        if false {
            var v8 uint8 = 154
            v4 += 3.14159
            _ = v8
        } else {
            p1 = "foo bar"
            switch p1 {
            case "hello":
                v9 := GenS1{F0: ((809 != 0) || (329 > 0)), F1: 0}
                _ = v9
            }
        }
    }
    p0 = 304
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 255
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    v3 := []uint32{}
    _ = len(v3)
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: true, F1: -9223372036854775808}
        var v12 uint64 = 1
        if (849 > 0) {
            v12 = (v12 + (v12 + v12))
            v13 := GenS1{F0: true, F1: -9223372036854775808}
            _ = v13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    switch 616 {
    case -683:
        v14 := v2[-1]
        _ = v14
    }
    if (float64(1.5) != 0.0) {
        if len(v3) > 0 {
            v3[0] = 976
        }
    }
    if false {
        if (1 <= -1) {
            for fi15 := 0; fi15 < 7; fi15++ {
                switch ("\n\t" + "\n\t") {
                case "foo bar":
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v17 := GenS1{F0: true, F1: 0}
                        var v18 int32 = 2147483647
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                default:
                    v19 := GenS1{F0: (148 > 0), F1: -1}
                    _ = v19
                }
                if (-32768 >= -32768) {
                    var v20 uint32 = 42
                    _ = v20
                } else {
                    var v21 uint64 = 2
                    if (2 > 217) {
                        var v22 uint16 = 222
                        var v23 uint16 = 787
                        _ = v22
                        _ = v23
                    }
                    _ = v21
                }
                _ = fi15
            }
            delete(v1, 4294967295)
            var v24 float32 = -2.25
            _ = v24
        }
        if (2147483647 == 0) {
            switch "foo bar" {
            case "\n\t":
                v25 := GenS1{F0: (42 != 1), F1: 9223372036854775807}
                _ = v25
            }
            var v26 float32 = 3.14159
            for fi27 := 0; fi27 < 17; fi27++ {
                v28 := v2[fi27]
                _ = fi27
                _ = v28
            }
            _ = v26
        } else {
            if (true || false) {
                if len(v3) > 0 {
                    v3[0] = 163
                }
                v3 = append(v3, 172)
                v29 := GenS1{F0: true, F1: 1}
                _ = v29
            }
        }
        switch 0 {
        case 32768:
            var v30 int8 = -128
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        var v32 int64 = 42
        switch "a" {
        case "":
            v33 := GenS1{F0: (((3.14159 - 0.0) - 1e10) >= (float64(3.14159) - 1.5)), F1: (v32 << ((255) & 63))}
            _ = v33
        case "a":
            if ((0 >= 42) && (true || true)) {
                fmt.Println(v32)
                v32 = v32
                _ = len(v3)
            }
        }
        for fi34 := 0; fi34 < 2; fi34++ {
            v32 = v32
            var v35 uint32 = 1
            var v36 uint16 = 894
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi31
        _ = v32
    }
    var v37 float64 = (0.0 - 3.14159)
    fmt.Println(v37)
    for fi38 := 0; fi38 < 12; fi38++ {
        switch "foo bar" {
        case "a":
            for fi39 := 0; fi39 < 15; fi39++ {
                _ = len(v2)
                fmt.Println(fi38, fi39)
                _ = fi39
            }
        default:
            if (true || ((!false) || (true && false))) {
                var v40 float64 = v37
                fmt.Println(fi38, v40, v37)
                var v41 float32 = 3.14159
                _ = v40
                _ = v41
            }
        }
        _ = fi38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
}
