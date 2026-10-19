package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 int8) bool {
    if false {
        p0 = p0
        v4 := []int{0, 817, -595}
        _ = v4
    } else {
        if false {
            v5 := []int{}
            v6 := GenS1{F0: 922, F1: 32767, F2: 2}
            _ = v5
            _ = v6
        } else {
            fmt.Println(p0)
            for fi7 := 0; fi7 < 8; fi7++ {
                p0 = p0
                _ = fi7
            }
        }
        if (42 > 65535) {
            switch "hello" {
            case "\n\t":
                p0 = 100
            }
        }
    }
    switch 32767 {
    case -359:
        if (32767 <= 9223372036854775807) {
            v8 := GenS1{F0: 510, F1: 2, F2: 1}
            var v9 int32 = -480
            var v10 bool = (2 > 668)
            _ = v10
            _ = v8
            _ = v9
        } else {
            v11 := []int{}
            v12 := GenS1{F0: 2, F1: 1, F2: 0}
            _ = v11
            _ = v12
        }
    case 127:
        v13 := GenS1{F0: 4294967295, F1: 42, F2: 78}
        _ = v13
    default:
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 string = ((("" + "a") + "\n\t") + "\n\t")
                if true {
                    var v17 string = v16
                    _ = v17
                } else {
                    v18 := GenS2{F0: 1, F1: 65}
                    v19 := []int{(fi14 >> ((65535) & 63)), -470, (fi14 >> ((18446744073709551615) & 63))}
                    var v20 bool = true
                    _ = v18
                    _ = v19
                    _ = v20
                }
                var v21 uint64 = 0
                _ = fi15
                _ = v16
                _ = v21
            }
            v22 := []int{fi14, fi14}
            _ = fi14
            _ = v22
        }
    }
    _ = p0
    return true
}
func genFn1(p0 int64, p1 string) int {
    var v23 uint64 = 0
    v24 := GenS3{F0: 255, F1: -1}
    v25 := GenS1{F0: 255, F1: 0, F2: 127}
    var v26 string = ("\n\t" + "hello")
    if (0 <= 1) {
        var v27 float64 = -2.25
        for fi28 := 0; fi28 < 19; fi28++ {
            if (42 < 1) {
                var v29 int32 = -2147483648
                var v30 bool = true
                _ = v29
                _ = v30
            }
            v31 := GenS4{F0: v26, F1: (true || (!false))}
            _ = fi28
            _ = v31
        }
        _ = v27
    } else {
        for fi32 := 0; fi32 < 17; fi32++ {
            var v33 int64 = (p0 / (((p0 % (((p0 * (^(p0)))) | 1))) | 1))
            p1 = p1
            _ = fi32
            _ = v33
        }
        switch -394 {
        case 2147483647:
            var v34 int32 = int32(p0)
            _ = v34
        }
    }
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    return 9223372036854775807
}

func main() {
    v1 := []string{}
    var v2 float32 = -2.25
    var v3 uint8 = 0
    var v35 uint8 = v3
    var v36 int16 = 772
    v37 := GenS1{F0: 65535, F1: 0, F2: v35}
    v2 -= v2
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    var v38 uint = 42
    fmt.Println(v3, v35, v36)
    v1 = append(v1, ("foo bar" + ("foo bar" + "a")))
    fmt.Println(v36, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
