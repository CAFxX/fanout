package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 string }

func main() {
    var v1 uint8 = 2
    var v2 int16 = 42
    var v3 uint32 = 65535
    var v4 float32 = (float32(0.1) * (-2.25 - 0.0))
    v2 += (v2 / (((v2 * (v2 << ((1) & 63)))) | 1))
    var v5 float32 = v4
    var v6 int8 = 100
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: v3, F1: "\n\t", F2: (((("foo bar" + (("\n\t" + "") + "\n\t")) + "") + ("a" + "hello")) + "foo bar")}
        for fi9 := 0; fi9 < 7; fi9++ {
            fmt.Println(v5, v2, v6)
            if false {
                v10 := []int{(^(fi7))}
                var v11 int32 = 42
                _ = len(v10)
                fmt.Println(v2, v1, fi7)
                var v12 uint = 464
                _ = v10
                _ = v11
                _ = v12
            } else {
                v13 := GenS2{F0: ("" + "foo bar"), F1: 0}
                if true {
                    v14 := []int{fi7}
                    var v15 int8 = (v6 << ((42) & 63))
                    v16 := GenS1{F0: v3, F1: "\n\t", F2: ("foo bar" + "")}
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint16 = 974
                    v18 := GenS1{F0: (v3 | (^(v3))), F1: "foo bar", F2: ("a" + ("foo bar" + ("foo bar" + "\n\t")))}
                    v19 := []int{fi7, -212, fi9}
                    fmt.Println(v1)
                    v1 += (v1 ^ v1)
                    var v20 float32 = ((v5 / v4) / v5)
                    if (0 > -1) {
                        fmt.Println(v5, v20, fi9)
                        var v21 uint32 = v3
                        var v22 float32 = (v4 / 1e10)
                        _ = v21
                        _ = v22
                    } else {
                        switch (fi9 / ((fi7) | 1)) {
                        case 692:
                            fmt.Println(v5, v6)
                        case 42:
                            var v23 int32 = 0
                            _ = v23
                        }
                        v19 = append(v19, 846)
                    }
                    v24 := GenS3{F0: int32(v3)}
                    v6 += 127
                    v20 = v4
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                    _ = v24
                }
                v25 := GenS3{F0: 2147483647}
                _ = v13
                _ = v25
            }
            _ = fi9
        }
        if (false || false) {
            v6 += v6
        } else {
            fmt.Println(v2, fi7)
            v26 := []int{1, fi7}
            fmt.Println(v4)
            _ = v26
        }
        _ = fi7
        _ = v8
    }
    if (!false) {
        if (true || (v2 < 0)) {
            v3 += v3
        }
    }
    v27 := GenS2{F0: "hello", F1: 65535}
    switch 640 {
    case 912:
        switch 2147483647 {
        case 2147483647:
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 float32 = v5
                _ = fi28
                _ = v29
            }
        case 951:
            for fi30 := 0; fi30 < 17; fi30++ {
                for fi31 := 0; fi31 < 17; fi31++ {
                    var v32 bool = false
                    if v32 {
                        v1 = v1
                        fmt.Println(fi31)
                    } else {
                        v33 := []int{}
                        switch "\n\t" {
                        case "a":
                            var v34 uint8 = v1
                            _ = v34
                        case "":
                            v35 := GenS4{F0: 0, F1: "\n\t"}
                            _ = v35
                        case "foo bar":
                            v6 -= v6
                        default:
                            var v36 uint = 1
                            _ = v36
                        }
                        v32 = (v6 <= (v6 - (v6 * v6)))
                        _ = v33
                    }
                    _ = fi31
                    _ = v32
                }
                var v37 float32 = v5
                _ = fi30
                _ = v37
            }
        }
    case 32768:
        v38 := GenS1{F0: (v3 / (((^(v3))) | 1)), F1: "a", F2: "\n\t"}
        _ = v38
    }
    v39 := []int{0, 2147483647, -386, 1}
    if len(v39) > 0 {
        v39[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
