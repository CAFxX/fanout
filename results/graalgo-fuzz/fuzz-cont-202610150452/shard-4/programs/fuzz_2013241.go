package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int16 }

func main() {
    var v1 uint8 = 42
    var v2 uint16 = 2
    v3 := make(map[string]int)
    if (42 >= -2147483648) {
        var v4 float64 = ((float64((0.1 / 1.0)) - 0.1) * ((1e10 * 0.0) / (-2.25 * -2.25)))
        v4 += 3.14159
        v5 := []int{4294967295, 32767, 4294967295, -1}
        _ = v4
        _ = v5
    } else {
        var v6 uint32 = 4294967295
        _ = v6
    }
    if true {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 42}
            if ((!true) || false) {
                fmt.Println(v1)
                if (0.1 != (0.1 / 1.0)) {
                    for fi9 := 0; fi9 < 12; fi9++ {
                        v10 := []int{1, fi9, fi7}
                        _ = fi9
                        _ = v10
                    }
                    v2 = v2
                    v11 := []int{fi7, fi7, fi7, (fi7 | (fi7 & (fi7 << ((108) & 63))))}
                    _ = v11
                }
                var v12 uint64 = 4294967295
                _ = v12
            }
            if false {
                for fi13 := 0; fi13 < 14; fi13++ {
                    v14 := GenS1{F0: 63}
                    _ = fi13
                    _ = v14
                }
                if ((!((!false) || (false && true))) && false) {
                    var v15 float32 = (0.1 / (-2.25 - 1e10))
                    _ = v15
                }
                if false {
                    v1 = (^(v1))
                } else {
                    var v16 uint = 18446744073709551615
                    _ = v16
                }
            } else {
                var v17 int32 = 2147483647
                fmt.Println(fi7)
                for fi18 := 0; fi18 < 16; fi18++ {
                    fmt.Println(v2)
                    var v19 int32 = v17
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
            _ = fi7
            _ = v8
        }
    }
    v2 += 345
    var v20 uint8 = 255
    v21 := GenS2{F0: v2, F1: 937, F2: 0}
    switch -415 {
    case -1:
        if false {
            var v22 uint = 31
            v23 := GenS2{F0: v2, F1: 42, F2: 32767}
            v2 = 255
            _ = v22
            _ = v23
        }
    default:
        var v24 bool = (false && true)
        _ = v24
    }
    if true {
        v25 := []int{0, int(v2)}
        _ = v25
    } else {
        fmt.Println(v20, v1)
        if (!true) {
            for fi26 := 0; fi26 < 15; fi26++ {
                v1 = (v20 - 2)
                var v27 uint64 = 18446744073709551615
                _ = fi26
                _ = v27
            }
        }
        var v28 int8 = -1
        _ = v28
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        if false {
            v30 := v3["\n\t"]
            v31 := GenS2{F0: v2, F1: 2, F2: 1000}
            for fi32 := 0; fi32 < 10; fi32++ {
                v1 = 255
                _ = fi32
            }
            _ = v30
            _ = v31
        }
        _ = fi29
    }
    _ = len(v3)
    v33 := GenS2{F0: (v2 / ((v2) | 1)), F1: 1, F2: 1000}
    v34 := GenS3{F0: -535, F1: 0}
    if false {
        switch ("" + "a") {
        case "foo bar":
            var v35 int = 127
            _ = v35
        }
        v2 = v2
        v1 = uint8(v2)
    }
    var v36 uint = 2
    fmt.Println(v2, v20)
    v1 += v20
    for fi37 := 0; fi37 < 7; fi37++ {
        var v38 uint8 = v1
        v39 := GenS4{F0: 589}
        v40 := v3[(("hello" + "hello") + ("" + (("\n\t" + "a") + "hello")))]
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    var v41 uint = (v36 >> ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v41
}
