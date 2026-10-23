package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int8 }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
func genFn0(p0 bool, p1 uint32) bool {
    v4 := []int{32768, 29, 0, -1}
    fmt.Println(p0, p1)
    var v5 int = -1
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 uint8, p1 string) int {
    for fi6 := 0; fi6 < 11; fi6++ {
        if (4294967295 < 461) {
            v7 := GenS1{F0: 414, F1: -128, F2: 0}
            if true {
                fmt.Println(p1)
            } else {
                for fi8 := 0; fi8 < 14; fi8++ {
                    fmt.Println(fi8, p1)
                    _ = fi8
                }
                for fi9 := 0; fi9 < 17; fi9++ {
                    var v10 int16 = -334
                    fmt.Println(p0, fi9)
                    switch (fi6 % ((2147483647) | 1)) {
                    case 9223372036854775807:
                        var v11 float32 = 1e10
                        _ = v11
                    case -9223372036854775808:
                        v10 = -880
                    }
                    _ = fi9
                    _ = v10
                }
            }
            _ = v7
        } else {
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := GenS1{F0: 133, F1: 100, F2: 127}
                _ = fi12
                _ = v13
            }
            var v14 int = (fi6 | 128)
            _ = v14
        }
        v15 := []int{(^(fi6))}
        var v16 uint64 = 0
        _ = fi6
        _ = v15
        _ = v16
    }
    switch -68 {
    case 32767:
        var v17 float32 = 1.5
        _ = v17
    case 9223372036854775807:
        var v18 string = p1
        _ = v18
    }
    v19 := []int{1, 128, 4294967295}
    if true {
        v19 = append(v19, 32767)
    }
    _ = p0
    _ = p1
    _ = v19
    return -328
}

func main() {
    var v1 string = "foo bar"
    v2 := []int{128, 2147483647, -2147483648, 321}
    var v3 int8 = -1
    var v20 int8 = 124
    var v21 bool = ((!(42 <= 945)) && (!false))
    if (328 <= 42) {
        _ = len(v2)
        for fi22 := 0; fi22 < 7; fi22++ {
            switch ((v1 + v1) + (v1 + v1)) {
            case "\n\t":
                var v23 uint = 255
                _ = v23
            }
            v24 := GenS1{F0: 4294967295, F1: v20, F2: -128}
            v21 = true
            _ = fi22
            _ = v24
        }
        v1 = (v1 + "a")
    }
    fmt.Println(v1, v1)
    for fi25 := 0; fi25 < 5; fi25++ {
        _ = len(v2)
        if (170 > 255) {
            v26 := GenS1{F0: 1, F1: 42, F2: -1}
            _ = v26
        }
        switch (fi25 >> ((1) & 63)) {
        case -2147483648:
            for fi27 := 0; fi27 < 5; fi27++ {
                fmt.Println(v20)
                if v21 {
                    v28 := GenS2{F0: 42, F1: 42, F2: 856}
                    fmt.Println(fi25, v21)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v28
                }
                v21 = v21
                _ = fi27
            }
        }
        _ = fi25
    }
    var v29 uint32 = 1
    var v30 float64 = (float64(-2.25) - float64(-2.25))
    for fi31 := 0; fi31 < 5; fi31++ {
        for fi32 := 0; fi32 < 8; fi32++ {
            var v33 int8 = v3
            _ = fi32
            _ = v33
        }
        var v34 int32 = -2147483648
        _ = fi31
        _ = v34
    }
    var v35 int16 = -728
    v36 := GenS2{F0: 65535, F1: -9223372036854775808, F2: -767}
    if len(v2) > 0 {
        v2[0] = 127
    }
    v37 := GenS2{F0: 65535, F1: 9223372036854775807, F2: 14}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
}
