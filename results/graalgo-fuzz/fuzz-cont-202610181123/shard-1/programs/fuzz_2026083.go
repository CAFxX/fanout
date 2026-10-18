package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint32; F2 uint64 }
func genFn0(p0 int64) string {
    v7 := []int{244, -1}
    var v8 int64 = -9223372036854775808
    v9 := GenS1{F0: "\n\t"}
    if len(v7) > 0 {
        v7[0] = 0
    }
    v10 := GenS2{F0: 957, F1: 2, F2: 42}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return "a"
}
func genFn1(p0 bool) int32 {
    p0 = true
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := []int{(fi12 * fi12), fi11, (-(fi12)), 128}
            p0 = (!(!true))
            switch 0 {
            case -2147483648:
                var v14 bool = (18446744073709551615 == 4294967295)
                _ = v14
            default:
                if ((true || true) || (!(p0 || p0))) {
                    var v15 int16 = 651
                    _ = v15
                } else {
                    switch (((("a" + "\n\t") + (("a" + "\n\t") + (((("hello" + "foo bar") + "a") + ("" + "")) + "a"))) + "\n\t") + "hello") {
                    case "a":
                        switch fi11 {
                        case -246:
                            v13 = append(v13, (-(fi11)))
                            fmt.Println(p0, fi11, fi11)
                            v16 := GenS2{F0: -1, F1: 2, F2: 42}
                            _ = v16
                        case 2147483647:
                            p0 = (255 <= 0)
                        case 32768:
                            fmt.Println(p0, fi12, p0)
                            var v17 bool = true
                            _ = v17
                        }
                    case "foo bar":
                        var v18 int16 = -32768
                        _ = v18
                    default:
                        var v19 string = ""
                        _ = v19
                    }
                    if len(v13) > 0 {
                        v13[0] = fi11
                    }
                    var v20 float32 = -2.25
                    _ = v20
                }
            }
            _ = fi12
            _ = v13
        }
        p0 = ("foo bar" >= ("foo bar" + ("hello" + "foo bar")))
        var v21 float64 = float64((1e10 / 1e10))
        _ = fi11
        _ = v21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        switch "foo bar" {
        case "hello":
            var v23 int16 = 0
            _ = v23
        }
        switch "" {
        case "a":
            v24 := []int{(fi22 / ((-548) | 1)), 851}
            _ = v24
        case "":
            for fi25 := 0; fi25 < 20; fi25++ {
                p0 = (p0 || (!("\n\t" <= "")))
                _ = fi25
            }
        }
        _ = fi22
    }
    var v26 int16 = 32767
    _ = p0
    _ = v26
    return 369
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 2147483647
    var v3 uint16 = 1
    var v4 float64 = (0.0 - 0.0)
    v5 := make(map[int]int)
    v6 := make(map[uint32]uint32)
    for fi27 := 0; fi27 < 6; fi27++ {
        v3 += v3
        for fi28 := 0; fi28 < 19; fi28++ {
            v2 = (v2 / ((0) | 1))
            v3 = (v3 ^ v3)
            fmt.Println(v3)
            _ = fi28
        }
        _ = fi27
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        for fi30 := 0; fi30 < 4; fi30++ {
            var v31 float64 = v4
            _ = fi30
            _ = v31
        }
        var v32 int8 = 1
        var v33 int64 = -9223372036854775808
        _ = fi29
        _ = v32
        _ = v33
    }
    v2 -= -2147483648
    for fi34 := 0; fi34 < 6; fi34++ {
        for fi35 := 0; fi35 < 19; fi35++ {
            fmt.Println(v4, v4)
            _ = fi35
        }
        v36 := GenS2{F0: -1, F1: 4294967295, F2: 18446744073709551615}
        _ = fi34
        _ = v36
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        var v38 uint = 255
        v38 += v38
        var v39 float32 = ((((0.1 / 3.14159) * -2.25) * ((((0.0 * ((((((1.5 - (3.14159 + (-2.25 / (1.5 + -2.25)))) * 3.14159) - float32(1e10)) / 0.1) * (0.0 + (1.5 + ((-2.25 * -2.25) * ((1e10 * 1e10) / 0.1))))) - 0.0)) / 1.0) * 1e10) * 1.5)) * 3.14159)
        _ = fi37
        _ = v38
        _ = v39
    }
    _ = len(v1)
    var v40 int64 = -479
    v41 := GenS1{F0: ("\n\t" + ("hello" + ("foo bar" + "foo bar")))}
    var v42 string = "hello"
    _ = len(v1)
    v40 = (^(v40))
    var v43 float32 = (0.1 * (0.0 / (3.14159 * 1.5)))
    if (!(v43 >= 1.5)) {
        var v44 float32 = v43
        var v45 int = 804
        _ = len(v1)
        _ = v44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
