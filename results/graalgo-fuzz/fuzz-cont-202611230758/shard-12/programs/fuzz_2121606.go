package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int8, p1 uint32) int {
    switch "a" {
    case "a":
        var v6 bool = false
        _ = v6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(p1, fi7)
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    return -181
}
func genFn1() uint32 {
    var v8 int32 = 0
    v9 := []int{1, 32767, 42, 42}
    _ = v8
    _ = v9
    return 2147483648
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = -2.25
    var v3 uint = 39
    var v4 float32 = ((((1e10 / 1.0) + (3.14159 + 1e10)) - -2.25) * 3.14159)
    var v5 int16 = 509
    var v10 uint = 1
    v4 += (v4 / (v4 + (v4 + (v4 / (v4 + v4)))))
    fmt.Println(v5)
    var v11 float32 = v4
    switch int(v5) {
    case -2147483648:
        if false {
            v12 := GenS1{F0: 65535, F1: 1}
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 string = ""
                for fi15 := 0; fi15 < 3; fi15++ {
                    v14 = v14
                    var v16 int64 = 649
                    _ = fi15
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
            _ = v12
        } else {
            for fi17 := 0; fi17 < 4; fi17++ {
                switch "foo bar" {
                case "":
                    v10 -= v3
                case "hello":
                    if (127 <= -84) {
                        v18 := GenS1{F0: 839, F1: 2}
                        _ = v18
                    }
                }
                _ = fi17
            }
            var v19 uint8 = 128
            _ = v19
        }
    case 9223372036854775807:
        delete(v1, 65535)
    default:
        if (1 == 0) {
            v5 = (v5 & (v5 % (((^(v5))) | 1)))
        }
    }
    v20 := v1[32]
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 uint8 = 70
            v3 = (v3 | uint(v22))
            _ = fi21
            _ = v22
        }
    case "\n\t":
        if (v2 <= v2) {
            if (-2.25 == v2) {
                if false {
                    v23 := []int{v20, (v20 * genFn0(0, 42))}
                    v3 = (v3 / ((2) | 1))
                    _ = v23
                }
                v5 = v5
            }
        }
    default:
        var v24 int16 = (^(v5))
        _ = v24
    }
    if true {
        v25 := GenS1{F0: 428, F1: 957}
        _ = v25
    }
    v20 -= v20
    v26 := GenS1{F0: 37, F1: 4294967295}
    for fi27 := 0; fi27 < 2; fi27++ {
        switch (v20 / (((^(v20))) | 1)) {
        case 32767:
            if (!(771 < 18446744073709551615)) {
                v28 := []int{v20, fi27}
                _ = v28
            } else {
                v29 := GenS1{F0: 2147483648, F1: 544}
                _ = v29
            }
        }
        v4 = ((v11 + v11) + (1e10 / (v4 - v11)))
        var v30 uint64 = 4294967295
        _ = fi27
        _ = v30
    }
    switch genFn0(100, 685) {
    case 4294967295:
        var v31 bool = false
        _ = v31
    case 9:
        for fi32 := 0; fi32 < 6; fi32++ {
            if ((556 > 0) || (-9223372036854775808 < 0)) {
                for fi33 := 0; fi33 < 17; fi33++ {
                    var v34 int32 = 594
                    v10 = 85
                    _ = fi33
                    _ = v34
                }
                var v35 uint16 = 32767
                var v36 int16 = (v5 / (((^(v5))) | 1))
                _ = v35
                _ = v36
            } else {
                if false {
                    var v37 int = (fi32 >> ((42) & 63))
                    _ = v37
                }
            }
            _ = fi32
        }
    case -9223372036854775808:
        v38 := []int{genFn0(42, 242)}
        _ = v38
    default:
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
