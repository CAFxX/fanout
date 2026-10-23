package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0() bool {
    var v7 uint64 = 948
    v8 := []int{4294967295, 650, -186}
    v9 := GenS1{F0: "\n\t", F1: 1}
    if (-563 == 1000) {
        for fi10 := 0; fi10 < 6; fi10++ {
            for fi11 := 0; fi11 < 1; fi11++ {
                switch fi11 {
                case 730:
                    v7 -= v7
                }
                var v12 int8 = -1
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    } else {
        v7 = v7
        if false {
            v8 = append(v8, 1)
            v7 = (v7 * 468)
            var v13 int64 = 42
            _ = v13
        } else {
            if false {
                var v14 int32 = 0
                v15 := GenS1{F0: (("a" + "a") + ""), F1: 18446744073709551615}
                var v16 int = -2147483648
                _ = v14
                _ = v15
                _ = v16
            } else {
                if len(v8) > 0 {
                    v8[0] = 0
                }
                var v17 int32 = 2147483647
                _ = v17
            }
        }
    }
    var v18 uint32 = 4294967295
    _ = v18
    _ = v7
    _ = v8
    _ = v9
    return (true || false)
}
func genFn1(p0 string) int64 {
    if true {
        var v19 int8 = 4
        var v20 float64 = ((((0.0 + float64(((3.14159 - 1.5) / 1.5))) / 1.0) * (1e10 / 1.0)) - (1e10 * (0.1 * float64(1e10))))
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: (p0 + "hello"), F1: 65535}
    if (42 != 0) {
        v22 := GenS1{F0: "\n\t", F1: 583}
        _ = v22
    }
    p0 = (p0 + "\n\t")
    fmt.Println(p0)
    _ = p0
    _ = v21
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v4 uint8 = 0
    var v5 float32 = (((1e10 * float32((-2.25 + (0.1 + (-2.25 + (1e10 - (3.14159 + 3.14159))))))) - -2.25) * ((-2.25 / (1e10 / 1.5)) + -2.25))
    v6 := make(map[uint32]string)
    var v23 uint32 = 84
    v24 := GenS2{F0: -233, F1: -9223372036854775808}
    var v25 int32 = -1
    var v26 int16 = 42
    fmt.Println(v26, v26)
    if true {
        var v27 uint32 = 1
        v28 := GenS1{F0: ("\n\t" + "hello"), F1: 42}
        _ = v27
        _ = v28
    }
    var v29 uint32 = (v23 & 2147483647)
    v25 -= 1
    if (false || true) {
        v30 := v2[-1]
        v4 = (v4 & v4)
        v31 := GenS2{F0: -886, F1: -1}
        _ = v30
        _ = v31
    }
    if false {
        fmt.Println(v5)
        v23 -= (v29 << ((uint(v23)) & 63))
    } else {
        v32 := []int{}
        _ = v32
    }
    fmt.Println(v4, v25)
    for fi33 := 0; fi33 < 12; fi33++ {
        var v34 string = ""
        v29 = v29
        _ = fi33
        _ = v34
    }
    switch ("" + "foo bar") {
    case "hello":
        var v35 int8 = 100
        _ = v35
    case "\n\t":
        var v36 int = int(v4)
        _ = v36
    case "a":
        delete(v2, 895)
    default:
        fmt.Println(v26, v4, v5)
    }
    var v37 string = "\n\t"
    var v38 int16 = v26
    v23 = (v23 << ((0) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
