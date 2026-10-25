package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint; F1 int16; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 bool; F2 int32 }
type GenS5 struct { F0 int; F1 uint8; F2 int }
type GenS6 struct { F0 bool; F1 string; F2 int32 }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS10 struct { F0 int16 }
func genFn0(p0 int) int8 {
    v5 := []int{452, -798}
    if ((float32(((float64(float32((-2.25 * -2.25))) + 0.1) / 1.0)) != 0.0) && true) {
        if false {
            if len(v5) > 0 {
                v5[0] = 32767
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            if (!(((1e10 + float64((0.0 * 0.0))) + 1.5) > 0.1)) {
                var v6 uint32 = 4294967295
                v6 -= v6
                var v7 int = p0
                _ = v6
                _ = v7
            }
            var v8 int32 = -1
            if false {
                if (42 != 65535) {
                    v8 -= (v8 / ((-2147483648) | 1))
                    if (!false) {
                        _ = len(v5)
                        var v9 uint64 = 0
                        _ = v9
                    }
                } else {
                    if false {
                        fmt.Println(v8, p0)
                    }
                }
                var v10 uint32 = 833
                v11 := GenS1{F0: ("foo bar" + "foo bar"), F1: 0, F2: 88}
                _ = v10
                _ = v11
            } else {
                var v12 int8 = 0
                v13 := GenS2{F0: 450, F1: 1000, F2: 151}
                p0 = (^(p0))
                _ = v12
                _ = v13
            }
            _ = v8
        }
        switch ("hello" + "\n\t") {
        case "foo bar":
            v14 := GenS3{F0: 80}
            _ = v14
        }
        var v15 int8 = 120
        _ = v15
    } else {
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 uint8 = 2
            var v18 bool = true
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 string = "hello"
        var v20 int8 = 1
        _ = v19
        _ = v20
    }
    if true {
        var v21 int64 = 42
        var v22 int64 = v21
        var v23 uint16 = 1
        _ = v21
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: "", F1: 255, F2: -1}
    v25 := GenS4{F0: 1, F1: true, F2: 2147483647}
    _ = p0
    _ = v24
    _ = v25
    _ = v5
    return -57
}
func genFn1() int8 {
    if true {
        if (!(true || true)) {
            var v26 uint = 4294967295
            v27 := GenS5{F0: 548, F1: 71, F2: 2147483647}
            v26 += v26
            _ = v26
            _ = v27
        } else {
            v28 := GenS6{F0: ((false && false) || false), F1: "a", F2: 647}
            _ = v28
        }
        v29 := GenS7{F0: -317, F1: 4294967295}
        _ = v29
    }
    v30 := GenS6{F0: false, F1: "foo bar", F2: -1}
    _ = v30
    return 127
}

func main() {
    var v1 float64 = (float64(1.5) / (-2.25 + 1e10))
    var v2 int = 32767
    var v3 uint32 = 2
    v4 := []string{""}
    switch "\n\t" {
    case "a":
        var v31 uint16 = 437
        _ = v31
    case "\n\t":
        v2 += v2
    case "hello":
        v2 -= 9223372036854775807
    }
    var v32 int64 = -166
    for fi33 := 0; fi33 < 4; fi33++ {
        for fi34 := 0; fi34 < 20; fi34++ {
            v35 := GenS8{F0: ""}
            _ = fi34
            _ = v35
        }
        var v36 int16 = 42
        _ = fi33
        _ = v36
    }
    v37 := GenS9{F0: (v3 * v3), F1: 1, F2: 127}
    if ((^(v2)) >= 9223372036854775807) {
        if false {
            v32 = v32
            v1 -= v1
        } else {
            v1 = ((((0.1 + v1) + (v1 - float64(1.5))) * v1) - v1)
            var v38 string = ("\n\t" + "a")
            if (18446744073709551615 < 42) {
                var v39 uint = 525
                _ = v39
            } else {
                var v40 int8 = 27
                _ = v40
            }
            _ = v38
        }
        var v41 uint64 = 18446744073709551615
        v1 -= 1.5
        _ = v41
    } else {
        switch v2 {
        case 127:
            v32 += (^(v32))
        case 4294967295:
            v3 = v3
        }
        var v42 uint32 = 2147483647
        var v43 bool = false
        _ = v42
        _ = v43
    }
    v3 -= v3
    for fi44 := 0; fi44 < 10; fi44++ {
        if len(v4) > 0 {
            v4[0] = ""
        }
        var v45 uint64 = 0
        _ = fi44
        _ = v45
    }
    fmt.Println(v32, v3, v3)
    for fi46 := 0; fi46 < 3; fi46++ {
        if true {
            fmt.Println(v3, v2, v2)
            fmt.Println(v1, v32)
            var v47 int8 = 33
            _ = v47
        } else {
            v4 = append(v4, "")
        }
        _ = fi46
    }
    fmt.Println(v2, v1)
    if (v32 < v32) {
        v48 := GenS5{F0: -1, F1: 128, F2: (v2 - v2)}
        _ = v48
    }
    switch (v2 + (v2 & v2)) {
    case 0:
        for fi49 := 0; fi49 < 14; fi49++ {
            var v50 int64 = 361
            v51 := GenS10{F0: 0}
            for fi52 := 0; fi52 < 6; fi52++ {
                var v53 float64 = ((v1 / 3.14159) * (v1 / ((3.14159 * v1) - float64(1.5))))
                _ = fi52
                _ = v53
            }
            _ = fi49
            _ = v50
            _ = v51
        }
    case 42:
        var v54 bool = false
        _ = v54
    case -2147483648:
        var v55 string = ""
        _ = v55
    default:
        v3 = (^(v3))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v4
}
