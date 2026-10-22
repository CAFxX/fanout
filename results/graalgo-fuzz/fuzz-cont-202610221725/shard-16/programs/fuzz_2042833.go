package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint }
type GenS2 struct { F0 bool; F1 uint }
func genFn0() uint16 {
    v5 := GenS1{F0: 628, F1: 759, F2: 0}
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := []int{fi7, 128}
            var v9 bool = true
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 2, F1: 2, F2: 1}
        _ = fi6
        _ = v10
    }
    _ = v5
    return 65535
}
func genFn1(p0 int16, p1 int8, p2 uint32) bool {
    v11 := GenS1{F0: 588, F1: 1, F2: 4294967295}
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(fi12)
        v13 := []int{4294967295, 127}
        _ = fi12
        _ = v13
    }
    if (-9223372036854775808 < 0) {
        v14 := []int{-744, 306, 0, 32767}
        _ = v14
    } else {
        v15 := GenS1{F0: 4294967295, F1: 2, F2: 394}
        _ = v15
    }
    fmt.Println(p2)
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return ((("\n\t" + ((("foo bar" + "hello") + "hello") + (("hello" + "") + "\n\t"))) + "hello") >= "\n\t")
}

func main() {
    var v1 int16 = 173
    v2 := make(map[uint32]int)
    var v3 string = ((("a" + "foo bar") + "\n\t") + "")
    var v4 float32 = (-2.25 * (((1.5 - float32((float64(float32(3.14159)) - 1.5))) - 0.1) / 1e10))
    var v16 float32 = v4
    var v17 int16 = (v1 ^ (v1 * -32768))
    for fi18 := 0; fi18 < 6; fi18++ {
        fmt.Println(v3, v3, v3)
        _ = fi18
    }
    var v19 int8 = 0
    v1 = (v17 + (v17 / ((-32768) | 1)))
    if (true && (25 > 42)) {
        delete(v2, 2147483647)
        v20 := []int{int(v17), 1, -832, 9223372036854775807}
        switch 1 {
        case 2147483647:
            if ((202 != 128) || (true || true)) {
                if ((v1 != v1) || (-964 != 42)) {
                    var v21 int8 = 100
                    var v22 int32 = -2147483648
                    var v23 uint8 = 127
                    _ = v21
                    _ = v22
                    _ = v23
                }
                for fi24 := 0; fi24 < 10; fi24++ {
                    v19 += v19
                    _ = len(v2)
                    v17 += v17
                    _ = fi24
                }
                switch 128 {
                case 1:
                    for fi25 := 0; fi25 < 16; fi25++ {
                        delete(v2, 550)
                        _ = fi25
                    }
                case -1:
                    fmt.Println(v4, v3)
                case 9223372036854775807:
                    var v26 float64 = (((1.5 + 0.0) * 1e10) * 1e10)
                    _ = v26
                }
            }
        case 613:
            var v27 int32 = -865
            _ = v27
        }
        _ = v20
    } else {
        var v28 uint32 = 892
        _ = v28
    }
    switch 127 {
    case 32767:
        if false {
            if (!true) {
                v29 := GenS2{F0: false, F1: 77}
                _ = v29
            }
        }
    }
    v4 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
