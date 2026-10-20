package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 bool; F1 bool; F2 uint32 }
func genFn0() uint16 {
    var v4 int16 = 781
    var v5 int8 = 1
    _ = v4
    _ = v5
    return 32767
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := []int{fi6, (fi6 & (fi6 * -9223372036854775808))}
        switch "" {
        case "hello":
            for fi8 := 0; fi8 < 7; fi8++ {
                if false {
                    v1 = float64(float32(-2.25))
                }
                v1 = ((1.5 / v1) - v1)
                var v9 uint64 = 136
                _ = fi8
                _ = v9
            }
        case "foo bar":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        _ = fi6
        _ = v7
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            v12 := GenS1{F0: 127, F1: genFn0()}
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            if ("\n\t" < ("a" + "hello")) {
                v14 := GenS1{F0: -88, F1: 42}
                var v15 int8 = 42
                v16 := GenS1{F0: (v15 * (v15 & -91)), F1: 42}
                fmt.Println(v1, fi13)
                _ = v14
                _ = v15
                _ = v16
            }
            v1 -= (v1 * (v1 + float64((-2.25 / float32(v1)))))
            v1 += float64(0.0)
            _ = fi13
        }
        _ = fi10
    }
    v2["hello"] = 4294967295
    if true {
        v1 = (v1 * float64(0.0))
        v1 -= v1
        if true {
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 int16 = -32768
                _ = fi17
                _ = v18
            }
            var v19 string = "hello"
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        _ = len(v2)
        _ = fi20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS1{F0: -13, F1: 32767}
        fmt.Println(fi21, fi21)
        if ((-2.25 <= ((float64(1.5) / v1) + (v1 - ((float64(3.14159) - v1) - v1)))) && (18446744073709551615 <= 174)) {
            var v23 int8 = 71
            v24 := GenS1{F0: v23, F1: 1}
            _ = v23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    fmt.Println(v1)
    if (0 != 63) {
        v1 = (v1 / ((0.0 - (0.0 - 1e10)) + v1))
        fmt.Println(v1)
    } else {
        for fi25 := 0; fi25 < 12; fi25++ {
            switch (fi25 + fi25) {
            case 32768:
                var v26 float32 = 1.5
                _ = v26
            }
            v1 += -2.25
            v27 := []int{fi25, fi25}
            _ = fi25
            _ = v27
        }
        for fi28 := 0; fi28 < 5; fi28++ {
            fmt.Println(fi28)
            v29 := []int{fi28, fi28}
            _ = fi28
            _ = v29
        }
        v1 += -2.25
    }
    switch 32767 {
    case 32767:
        fmt.Println(v1)
    }
    v30 := GenS2{F0: (true || (1 == 42)), F1: (!false), F2: 705}
    v2["hello"] = 0
    if (581 <= 480) {
        if false {
            if false {
                for fi31 := 0; fi31 < 8; fi31++ {
                    var v32 bool = true
                    var v33 uint = 65535
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                v34 := GenS2{F0: true, F1: (308 >= 1), F2: 2147483647}
                for fi35 := 0; fi35 < 17; fi35++ {
                    v1 = (v1 + ((0.0 + -2.25) + v1))
                    _ = fi35
                }
                _ = v34
            } else {
                var v36 string = "hello"
                _ = v36
            }
        }
    }
    switch 657 {
    case 184:
        v1 += v1
    case 42:
        var v37 int = 127
        _ = v37
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
}
