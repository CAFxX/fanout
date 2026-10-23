package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint32; F2 int16 }

func main() {
    var v1 bool = (!(!((0.0 >= 0.0) || false)))
    var v2 uint32 = 374
    var v3 int64 = 9223372036854775807
    var v4 int64 = v3
    var v5 int16 = 1000
    switch "\n\t" {
    case "":
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int64 = (v3 & (v4 >> ((65535) & 63)))
            var v8 float32 = 0.0
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    switch (("hello" + "a") + "a") {
    case "":
        var v9 float64 = 3.14159
        _ = v9
    case "foo bar":
        var v10 uint64 = 2
        _ = v10
    case "\n\t":
        v3 = -9223372036854775808
    default:
        if v1 {
            v2 = v2
        } else {
            for fi11 := 0; fi11 < 5; fi11++ {
                for fi12 := 0; fi12 < 9; fi12++ {
                    var v13 int8 = 0
                    switch (-(fi11)) {
                    case 9223372036854775807:
                        var v14 uint64 = 404
                        _ = v14
                    }
                    v15 := GenS1{F0: 1, F1: "\n\t"}
                    _ = fi12
                    _ = v13
                    _ = v15
                }
                _ = fi11
            }
        }
    }
    if v1 {
        v3 = 42
        fmt.Println(v5, v1)
        switch (("\n\t" + ("a" + ("hello" + "\n\t"))) + "hello") {
        case "":
            if (v5 >= v5) {
                var v16 int16 = 682
                v4 = v4
                if false {
                    v17 := []int{622, 2147483647}
                    _ = v17
                } else {
                    var v18 float64 = 3.14159
                    fmt.Println(v3, v1)
                    v2 += (v2 | 42)
                    _ = v18
                }
                _ = v16
            } else {
                for fi19 := 0; fi19 < 6; fi19++ {
                    v4 -= v4
                    if (0 <= 42) {
                        var v20 int32 = -2147483648
                        v21 := []int{fi19, 17, (^(fi19))}
                        var v22 float64 = (3.14159 - 1.5)
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        var v23 uint16 = 633
                        var v24 int64 = v4
                        _ = v23
                        _ = v24
                    }
                    _ = fi19
                }
                v5 += v5
            }
        case "hello":
            var v25 bool = (0 >= 32767)
            _ = v25
        default:
            var v26 int32 = 1
            _ = v26
        }
    } else {
        switch "\n\t" {
        case "":
            v2 = (v2 << ((515) & 63))
        default:
            if (!(!true)) {
                switch -2147483648 {
                case 9223372036854775807:
                    var v27 uint8 = 42
                    _ = v27
                case 0:
                    var v28 uint32 = 1
                    _ = v28
                case 32768:
                    fmt.Println(v5)
                default:
                    fmt.Println(v1)
                }
                v29 := GenS2{F0: v1, F1: -32768}
                _ = v29
            } else {
                var v30 float64 = ((0.0 / (0.1 + 0.0)) / -2.25)
                fmt.Println(v1, v30)
                _ = v30
            }
        }
        v31 := GenS2{F0: false, F1: (v5 - (v5 - v5))}
        v32 := GenS3{F0: 463}
        _ = v31
        _ = v32
    }
    if v1 {
        fmt.Println(v2, v5, v4)
    } else {
        v33 := GenS2{F0: v1, F1: (v5 | v5)}
        _ = v33
    }
    v3 = int64(v5)
    v5 = 1
    var v34 uint = 1
    v1 = v1
    if true {
        v4 = 9223372036854775807
        for fi35 := 0; fi35 < 10; fi35++ {
            var v36 int64 = v4
            _ = fi35
            _ = v36
        }
        for fi37 := 0; fi37 < 15; fi37++ {
            v4 += v3
            _ = fi37
        }
    }
    v38 := []int{991, 965}
    v1 = true
    if (v34 != (v34 & v34)) {
        v3 -= v3
        switch 1 {
        case 0:
            fmt.Println(v34, v34)
        }
        if v1 {
            fmt.Println(v5, v3, v3)
            v39 := GenS4{F0: (v2 * (v2 & v2)), F1: (v2 * (v2 | v2)), F2: v5}
            v38 = append(v38, -915)
            _ = v39
        } else {
            if v1 {
                v38 = append(v38, 9223372036854775807)
                switch (("a" + "\n\t") + "a") {
                case "":
                    v5 = v5
                case "hello":
                    v3 += (-(v4))
                case "foo bar":
                    if v1 {
                        v34 += v34
                        fmt.Println(v1, v4)
                    }
                }
                var v40 int64 = v3
                _ = v40
            }
            var v41 bool = false
            var v42 uint32 = (^(v2))
            _ = v41
            _ = v42
        }
    } else {
        var v43 uint64 = 2
        v4 += v4
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v5
}
