package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint; F2 uint64 }
func genFn0(p0 bool) int16 {
    switch "a" {
    case "a":
        p0 = p0
    }
    fmt.Println(p0)
    p0 = ((!(!true)) || p0)
    _ = p0
    return 32767
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]bool)
    var v3 float64 = 1e10
    v1 += (v1 - 375)
    fmt.Println(v1, v3)
    if true {
        fmt.Println(v1, v1)
    } else {
        if (v1 != v1) {
            var v4 float64 = 0.0
            _ = v4
        }
        v1 = v1
    }
    var v5 uint32 = 2147483647
    v6 := GenS1{F0: 474, F1: ""}
    var v7 int16 = -1
    if (42 <= -586) {
        v8 := GenS1{F0: 42, F1: "a"}
        v3 = (v3 / (((v3 + (v3 - ((v3 - v3) * (-2.25 - ((float64(((-2.25 + float32(0.0)) + (1e10 + 0.0))) / v3) - (v3 * 1.5)))))) / v3) + 3.14159))
        _ = v8
    }
    switch 9223372036854775807 {
    case 32768:
        if true {
            var v9 int64 = 502
            _ = v9
        }
    case -9223372036854775808:
        for fi10 := 0; fi10 < 1; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                if true {
                    var v12 uint32 = 1
                    var v13 int8 = 15
                    var v14 bool = true
                    _ = v12
                    _ = v13
                    _ = v14
                }
                if false {
                    var v15 uint32 = 1
                    v3 = 1.5
                    fmt.Println(v3)
                    v16 := []int{(fi11 - fi10), fi11, 127, 42}
                    fmt.Println(v1, v7)
                    _ = v15
                    _ = v16
                }
                var v17 uint = 2
                var v18 int8 = -1
                _ = fi11
                _ = v17
                _ = v18
            }
            _ = fi10
        }
    case 42:
        v19 := GenS2{F0: -2147483648, F1: (v5 | (v5 >> ((71) & 63)))}
        _ = v19
    default:
        if (((v5 + (v5 >> ((uint(v1)) & 63))) <= (v5 & v5)) && false) {
            if (!false) {
                fmt.Println(v1, v7)
                v3 = float64((1e10 + 1.5))
                v20 := GenS2{F0: -751, F1: 178}
                _ = v20
            }
            if ((4294967295 == 966) && true) {
                if (!true) {
                    v5 += v5
                    v21 := GenS2{F0: 257, F1: 0}
                    _ = v21
                }
                for fi22 := 0; fi22 < 1; fi22++ {
                    if (!(true || (0 >= 4294967295))) {
                        for fi23 := 0; fi23 < 17; fi23++ {
                            var v24 uint8 = 255
                            _ = fi23
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 20; fi25++ {
                            var v26 int16 = 195
                            _ = fi25
                            _ = v26
                        }
                        v27 := GenS1{F0: 255, F1: "hello"}
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 2; fi28++ {
                            v1 -= v1
                            if ((!false) && true) {
                                v3 -= -2.25
                                var v29 uint32 = v5
                                v30 := GenS1{F0: 942, F1: ""}
                                v31 := []int{fi28, fi22, (fi22 * 1)}
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            _ = fi28
                        }
                        switch 9223372036854775807 {
                        case 2147483647:
                            var v32 float64 = (1.5 - float64(1.5))
                            _ = v32
                        case 292:
                            if true {
                                delete(v2, "")
                            }
                        }
                        var v33 uint = 560
                        _ = v33
                    }
                    _ = fi22
                }
            }
            v34 := []int{350, 32767, 128, -1}
            _ = v34
        } else {
            v35 := GenS3{F0: false, F1: 961, F2: 4294967295}
            if false {
                v36 := []int{412, -1, -9223372036854775808}
                v3 = v3
                _ = v36
            }
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
