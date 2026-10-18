package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint; F2 bool }

func main() {
    var v1 uint32 = 634
    var v2 uint16 = 1
    var v3 uint = 65535
    var v4 int32 = -361
    var v5 string = "\n\t"
    var v6 float32 = 0.0
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 string = v5
        v9 := []int{fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint = (^(v3))
    v11 := GenS1{F0: v3, F1: 4294967295}
    if (false && false) {
        v12 := []int{128}
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 15; fi13++ {
            if (869 < 42) {
                var v14 int = fi13
                _ = v14
            } else {
                v15 := GenS2{F0: fi13, F1: (v1 & 190)}
                switch fi13 {
                case 164:
                    var v16 uint16 = 776
                    _ = v16
                }
                _ = v15
            }
            if (false || (2 >= 290)) {
                v1 = v1
                v3 = uint(v1)
            } else {
                v17 := []int{(fi13 ^ fi13), (^(fi13))}
                _ = v17
            }
            _ = fi13
        }
        switch "" {
        case "foo bar":
            v18 := GenS1{F0: (v3 | 573), F1: 290}
            _ = v18
        }
        if (!true) {
            var v19 uint = 918
            _ = v19
        }
    } else {
        v20 := GenS1{F0: v3, F1: 18446744073709551615}
        v21 := []int{32767, -9223372036854775808}
        if (!false) {
            var v22 bool = true
            for fi23 := 0; fi23 < 17; fi23++ {
                var v24 bool = v22
                v21 = append(v21, (fi23 + (fi23 / ((fi23) | 1))))
                var v25 int16 = 42
                _ = fi23
                _ = v24
                _ = v25
            }
            var v26 float32 = 1.5
            _ = v22
            _ = v26
        }
        _ = v20
        _ = v21
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        v4 = v4
        _ = fi27
    }
    v28 := GenS1{F0: v3, F1: 429}
    if false {
        v29 := GenS3{F0: 255, F1: 183, F2: true}
        var v30 uint64 = 376
        _ = v29
        _ = v30
    } else {
        fmt.Println(v1, v2, v3)
    }
    var v31 uint32 = 0
    fmt.Println(v5, v5)
    var v32 float32 = ((v6 * float32(1e10)) - ((v6 * 0.0) - v6))
    switch v5 {
    case "a":
        v33 := []int{607}
        _ = v33
    case "foo bar":
        fmt.Println(v32, v2, v1)
    case "hello":
        v6 = 0.0
    default:
        v6 = v32
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        v35 := []int{-9223372036854775808, int(v2), (fi34 + fi34)}
        var v36 int = -9223372036854775808
        switch v36 {
        case -1:
            if len(v35) >= 2 {
                v35 = v35[1:len(v35)]
            }
        case 563:
            switch v5 {
            case "foo bar":
                if (("hello" > (v5 + v5)) || true) {
                    v35 = append(v35, (-(v36)))
                    v37 := GenS2{F0: 9223372036854775807, F1: (v31 | v31)}
                    var v38 int16 = 1000
                    _ = v37
                    _ = v38
                } else {
                    fmt.Println(v2, v36)
                }
            }
        case 0:
            if false {
                if len(v35) > 0 {
                    v35[0] = (fi34 - fi34)
                }
            } else {
                fmt.Println(fi34)
            }
        default:
            v1 = (v1 & 0)
        }
        _ = fi34
        _ = v35
        _ = v36
    }
    if ((^(v31)) < 846) {
        var v39 uint32 = 131
        _ = v39
    }
    var v40 bool = false
    v31 -= v31
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
