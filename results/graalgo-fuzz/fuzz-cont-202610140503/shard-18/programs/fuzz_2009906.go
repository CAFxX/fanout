package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 761
    v3 := make(map[string]bool)
    if false {
        if (false && (390 != 1)) {
            v2 += (v2 ^ v2)
            v4 := GenS1{F0: -605}
            v2 += v2
            _ = v4
        } else {
            v2 = (v2 & 781)
            var v5 uint64 = 2
            _ = v5
        }
        var v6 uint = 42
        if (((0.1 * (((1e10 / (1e10 + float64(1.5))) * (-2.25 - 3.14159)) / 1.5)) - 0.0) == 3.14159) {
            v2 = (v2 << (((v6 * v6)) & 63))
            if false {
                if (161 <= 2) {
                    var v7 float64 = float64((1.5 * -2.25))
                    if true {
                        v7 = (((v7 / 3.14159) * v7) + 1.5)
                        v8 := GenS1{F0: v2}
                        _ = v8
                    } else {
                        v9 := GenS1{F0: v2}
                        _ = v9
                    }
                    _ = v7
                }
                var v10 uint64 = 784
                _ = v10
            } else {
                var v11 float64 = 1e10
                _ = v11
            }
            v2 = (v2 >> (((v6 | v6)) & 63))
        }
        _ = v6
    } else {
        fmt.Println(v2)
    }
    v2 = (v2 & v2)
    var v12 bool = (3.14159 >= 0.1)
    switch 32767 {
    case 2147483647:
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := []int{-2147483648}
            v12 = (-933 != 9223372036854775807)
            _ = fi13
            _ = v14
        }
    case 0:
        v2 = 2147483647
    case 32768:
        v15 := GenS1{F0: v2}
        _ = v15
    }
    v12 = ((^(v2)) > v2)
    v12 = (v12 && (293 > 1))
    if true {
        var v16 int64 = 0
        if (2 == 390) {
            switch -1 {
            case 32768:
                if v12 {
                    v17 := []int{-840, 128}
                    v18 := GenS2{F0: 42}
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v12 = ((42 == 0) && (!(!("\n\t" <= ("foo bar" + "a")))))
                        switch "\n\t" {
                        case "foo bar":
                            var v20 uint64 = 805
                            _ = v20
                        case "\n\t":
                            fmt.Println(v12, v2, v12)
                        case "a":
                            var v21 uint8 = 0
                            _ = v21
                        default:
                            var v22 int8 = -101
                            _ = v22
                        }
                        switch (fi19 + fi19) {
                        case 42:
                            v23 := v3[(("foo bar" + "\n\t") + "a")]
                            _ = v23
                        case 432:
                            v24 := GenS1{F0: (v2 << ((2) & 63))}
                            _ = v24
                        case -743:
                            if ((^(fi19)) < int(v16)) {
                                fmt.Println(v16, v12)
                                v2 -= (v2 / (((v2 ^ 2147483647)) | 1))
                            }
                        }
                        _ = fi19
                    }
                    _ = v17
                    _ = v18
                } else {
                    var v25 int64 = (-(v16))
                    v26 := []int{2147483647, 127, 128}
                    _ = v25
                    _ = v26
                }
            case 128:
                if ((!v12) || true) {
                    v2 -= v2
                    v27 := []int{672, 4294967295}
                    v16 += v16
                    _ = v27
                }
            default:
                fmt.Println(v2, v2)
            }
            _ = len(v1)
            fmt.Println(v16)
        } else {
            var v28 bool = v12
            var v29 bool = true
            var v30 bool = v29
            _ = v28
            _ = v29
            _ = v30
        }
        _ = v16
    } else {
        v2 = (v2 / ((-542) | 1))
        v31 := v1[42]
        var v32 uint64 = 4294967295
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        switch 42 {
        case -2147483648:
            if (1e10 > float64(0.1)) {
                var v34 float64 = ((0.1 - (0.0 - (-2.25 + 1.5))) - 1.5)
                _ = v34
            }
        case 1:
            v35 := v1[245]
            _ = v35
        }
        _ = fi33
    }
    var v36 uint = 473
    var v37 uint64 = 1
    v38 := GenS1{F0: (v2 >> ((42) & 63))}
    for fi39 := 0; fi39 < 20; fi39++ {
        var v40 int64 = 1
        _ = fi39
        _ = v40
    }
    v41 := v3["a"]
    v2 += 42
    if v41 {
        switch (("a" + ("a" + "\n\t")) + "a") {
        case "a":
            v41 = (28 == 1000)
        default:
            v42 := []int{}
            _ = v42
        }
    }
    _ = len(v1)
    v1[65535] = (9223372036854775807 != 163)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v41
}
