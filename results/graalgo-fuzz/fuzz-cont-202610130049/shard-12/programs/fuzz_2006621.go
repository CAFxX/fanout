package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int32; F2 bool }
type GenS4 struct { F0 int32 }
func genFn0() uint {
    var v7 float64 = 1e10
    if false {
        v7 = 3.14159
    } else {
        switch -1 {
        case 2147483647:
            v8 := GenS1{F0: 1, F1: 0}
            _ = v8
        case -937:
            if (4294967295 == 0) {
                var v9 int16 = 1
                fmt.Println(v9, v7)
                _ = v9
            }
        default:
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 int = (fi10 << ((1) & 63))
                v11 += v11
                var v12 uint64 = 42
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    }
    v13 := GenS1{F0: 533, F1: 42}
    v14 := GenS2{F0: 42}
    _ = v13
    _ = v14
    _ = v7
    return 65535
}
func genFn1() int {
    if true {
        if true {
            var v15 int16 = -172
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 uint64 = 42
                switch "" {
                case "a":
                    switch "a" {
                    case "":
                        v17 = v17
                    case "hello":
                        v17 = v17
                    }
                default:
                    v15 = (v15 / ((v15) | 1))
                }
                var v18 float32 = 3.14159
                _ = fi16
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 float64 = (1.5 * (1e10 + 3.14159))
                var v21 float32 = (3.14159 * (1e10 - 1.5))
                v22 := GenS3{F0: ("a" + "hello"), F1: -680, F2: (-119 != 127)}
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            _ = v15
        }
    } else {
        v23 := GenS2{F0: 252}
        _ = v23
    }
    v24 := []int{}
    _ = v24
    return 128
}

func main() {
    var v1 uint8 = 7
    v2 := make(map[uint32]int)
    var v3 uint64 = 602
    var v4 int = 0
    v5 := make(map[string]string)
    v6 := make(map[int]bool)
    fmt.Println(v3, v1, v3)
    v25 := GenS3{F0: "", F1: 89, F2: (3.14159 >= float64(0.0))}
    if true {
        var v26 int16 = 1
        _ = v26
    } else {
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 float32 = 3.14159
            v29 := GenS3{F0: ((("\n\t" + "hello") + "a") + ("foo bar" + "")), F1: -92, F2: (0 <= 65535)}
            _ = fi27
            _ = v28
            _ = v29
        }
        if (1 < -1) {
            var v30 float64 = float64(0.0)
            _ = v30
        }
        v4 = v4
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        v32 := GenS1{F0: 102, F1: 156}
        v1 = v1
        _ = fi31
        _ = v32
    }
    delete(v5, "")
    if (253 >= (v1 / ((0) | 1))) {
        var v33 uint64 = 20
        for fi34 := 0; fi34 < 7; fi34++ {
            var v35 float64 = float64(1.5)
            v36 := []int{0, v4, 997}
            fmt.Println(v1, v33)
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = v33
    } else {
        v37 := GenS3{F0: "a", F1: -1, F2: (!true)}
        v1 -= 0
        _ = v37
    }
    var v38 bool = ((42 < -2147483648) && true)
    var v39 int16 = -1
    var v40 int32 = -988
    var v41 uint16 = 2
    if v38 {
        var v42 uint = 468
        v43 := GenS4{F0: v40}
        _ = v42
        _ = v43
    } else {
        v44 := v2[2]
        v39 += -174
        v38 = ((0.0 - -2.25) < 0.0)
        _ = v44
    }
    switch "" {
    case "hello":
        if (v38 && (!v38)) {
            v45 := []int{878, 127, v4}
            v3 = 42
            _ = v45
        }
    case "a":
        fmt.Println(v38)
    case "":
        v46 := GenS3{F0: "foo bar", F1: (v40 >> ((18446744073709551615) & 63)), F2: (2 != 300)}
        _ = v46
    default:
        v39 = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
