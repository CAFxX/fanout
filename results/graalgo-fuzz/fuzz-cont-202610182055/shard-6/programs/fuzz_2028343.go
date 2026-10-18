package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int; F2 int64 }
type GenS4 struct { F0 int }
func genFn0(p0 uint8) bool {
    p0 = (p0 << ((18446744073709551615) & 63))
    fmt.Println(p0)
    if true {
        if (!((1 <= 42) || false)) {
            var v6 int16 = 295
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 int32 = 658
                _ = fi7
                _ = v8
            }
            var v9 string = "a"
            _ = v6
            _ = v9
        }
    }
    _ = p0
    return (205 > 1)
}

func main() {
    var v1 int32 = 1
    var v2 float64 = ((1e10 + float64(1e10)) - float64((1.5 - -2.25)))
    var v3 int8 = -1
    var v4 int = -238
    var v5 bool = (("\n\t" == "a") && (!((-2.25 >= 0.1) || true)))
    switch "foo bar" {
    case "hello":
        v3 += v3
    case "a":
        v3 += v3
    default:
        if v5 {
            for fi10 := 0; fi10 < 14; fi10++ {
                fmt.Println(fi10)
                fmt.Println(v4)
                v11 := GenS1{F0: (v4 + (^(fi10))), F1: 1}
                _ = fi10
                _ = v11
            }
            switch "a" {
            case "a":
                v12 := GenS2{F0: 170}
                _ = v12
            case "hello":
                var v13 int32 = v1
                _ = v13
            case "foo bar":
                v3 = (v3 << ((0) & 63))
            }
            for fi14 := 0; fi14 < 17; fi14++ {
                v15 := []int{fi14, 4294967295}
                v2 -= ((1e10 / (3.14159 / (0.0 * (v2 + v2)))) - (1.5 + float64(0.1)))
                _ = fi14
                _ = v15
            }
        } else {
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 int32 = (v1 ^ v1)
                switch (fi16 + v4) {
                case 32767:
                    v18 := []int{fi16, (v4 & fi16), 42}
                    _ = v18
                case -398:
                    v4 += 1
                default:
                    v19 := []int{v4}
                    _ = v19
                }
                v20 := GenS2{F0: 4294967295}
                _ = fi16
                _ = v17
                _ = v20
            }
        }
    }
    v3 = v3
    for fi21 := 0; fi21 < 20; fi21++ {
        if v5 {
            var v22 int = 9223372036854775807
            _ = v22
        } else {
            var v23 uint = 1
            for fi24 := 0; fi24 < 15; fi24++ {
                v25 := []int{fi21, (-(v4)), fi21}
                var v26 int = 127
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = v23
        }
        _ = fi21
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        fmt.Println(v4)
        _ = fi27
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 uint32 = 2
        v29 += 4294967295
        _ = fi28
        _ = v29
    }
    var v30 float64 = v2
    v31 := GenS3{F0: 65535, F1: (-(v4)), F2: -9223372036854775808}
    v2 += (-2.25 + (v30 + ((v2 / v30) / (v2 - (((((v30 / ((v30 + v30) / 1.5)) / float64((-2.25 / (1.5 * ((-2.25 / 1.5) * 1e10))))) * 0.1) + float64(0.1)) / (v2 / ((v2 / float64(((((0.0 / 1.0) * 3.14159) - 3.14159) * float32(v2)))) + float64(1.5))))))))
    for fi32 := 0; fi32 < 10; fi32++ {
        v2 = float64(float32(v30))
        v33 := []int{}
        _ = fi32
        _ = v33
    }
    v34 := GenS4{F0: v4}
    for fi35 := 0; fi35 < 19; fi35++ {
        v1 += 834
        v36 := []int{}
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
