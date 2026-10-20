package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int = -1
    var v3 int64 = -553
    v4 := GenS1{F0: 507}
    v5 := GenS2{F0: true}
    delete(v1, "foo bar")
    if (517 <= 2) {
        switch ("" + "\n\t") {
        case "foo bar":
            v6 := []int{v2, (v2 / ((v2) | 1)), (v2 >> ((72) & 63))}
            _ = v6
        }
        v3 += (-(v3))
    } else {
        fmt.Println(v2, v2)
        for fi7 := 0; fi7 < 7; fi7++ {
            for fi8 := 0; fi8 < 6; fi8++ {
                if (!true) {
                    v2 -= 42
                    var v9 int8 = 127
                    _ = v9
                } else {
                    v10 := GenS1{F0: 2}
                    v11 := GenS2{F0: true}
                    switch "foo bar" {
                    case "":
                        v3 = -317
                    }
                    _ = v10
                    _ = v11
                }
                if (42 >= -675) {
                    v12 := GenS1{F0: 1}
                    switch int(v3) {
                    case 128:
                        v3 += v3
                    case 4294967295:
                        v13 := []int{0, (fi8 >> ((255) & 63))}
                        _ = v13
                    default:
                        v14 := []int{(fi7 - (^(fi8))), 128, v2}
                        _ = v14
                    }
                    _ = v12
                }
                if (true || true) {
                    if (103 != 32767) {
                        v15 := GenS3{F0: false}
                        _ = v15
                    } else {
                        var v16 uint16 = 42
                        if ((1e10 * -2.25) <= 1e10) {
                            v2 += v2
                            switch fi8 {
                            case 351:
                                var v17 float32 = 1.5
                                _ = v17
                            default:
                                v3 = -804
                                v18 := GenS4{F0: 2147483647}
                                if false {
                                    delete(v1, "a")
                                }
                                _ = v18
                            }
                            var v19 bool = (!(fi8 <= (v2 / ((fi7) | 1))))
                            var v20 int16 = 620
                            _ = v19
                            _ = v20
                        }
                        _ = v16
                    }
                    v3 -= (-(v3))
                    v3 = (v3 | v3)
                }
                v2 = (^(fi7))
                var v21 uint8 = 241
                _ = fi8
                _ = v21
            }
            var v22 float64 = float64(-2.25)
            v22 = (1.5 * ((v22 * 1e10) * (v22 - v22)))
            _ = fi7
            _ = v22
        }
        _ = len(v1)
    }
    v3 = (v3 * v3)
    var v23 int16 = 1
    var v24 int32 = 42
    v2 = (v2 - (v2 - 299))
    switch (v2 >> ((65535) & 63)) {
    case 543:
        v3 = (v3 / ((1) | 1))
    case 2147483647:
        v25 := GenS3{F0: (2147483647 > 2)}
        _ = v25
    }
    v26 := GenS3{F0: true}
    v3 = (v3 * v3)
    switch "a" {
    case "foo bar":
        for fi27 := 0; fi27 < 10; fi27++ {
            for fi28 := 0; fi28 < 15; fi28++ {
                v29 := []int{v2, 9223372036854775807, 904}
                v30 := GenS1{F0: 0}
                _ = fi28
                _ = v29
                _ = v30
            }
            v3 += v3
            v3 += (v3 - -904)
            _ = fi27
        }
    default:
        v1["a"] = (!false)
    }
    v23 -= (v23 - (v23 << ((2) & 63)))
    v31 := []int{9223372036854775807, 753, 0}
    v24 = 1
    v32 := GenS4{F0: 0}
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
