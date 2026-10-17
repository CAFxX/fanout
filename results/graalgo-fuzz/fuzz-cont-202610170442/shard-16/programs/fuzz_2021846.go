package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int }
type GenS2 struct { F0 int32; F1 uint; F2 int64 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int8, p1 uint64) uint16 {
    if (((!true) || true) || true) {
        var v4 int32 = 42
        fmt.Println(v4, v4)
        _ = v4
    }
    var v5 bool = (!(!(!(false || (!(4294967295 == 2))))))
    if v5 {
        var v6 uint16 = 255
        v7 := GenS1{F0: v6, F1: 1, F2: -2147483648}
        v8 := []int{127, 0, 128, -447}
        _ = v6
        _ = v7
        _ = v8
    } else {
        fmt.Println(p0, p0, p1)
        if (255 != 42) {
            p0 += 100
            fmt.Println(v5, p1)
            switch "" {
            case "":
                p1 = p1
            case "\n\t":
                v5 = (!(!v5))
            }
        } else {
            switch 128 {
            case 9223372036854775807:
                v9 := []int{386, 1}
                _ = v9
            case 0:
                v10 := []int{-323, -9223372036854775808}
                _ = v10
            case 4294967295:
                p1 -= p1
            default:
                for fi11 := 0; fi11 < 5; fi11++ {
                    if false {
                        var v12 uint16 = 0
                        p0 = (^(p0))
                        _ = v12
                    } else {
                        p1 = 477
                        var v13 int8 = (^(p0))
                        _ = v13
                    }
                    v14 := GenS1{F0: 65535, F1: 2147483647, F2: fi11}
                    if (1 > 2) {
                        v15 := GenS1{F0: 0, F1: 0, F2: fi11}
                        for fi16 := 0; fi16 < 5; fi16++ {
                            var v17 int = fi11
                            _ = fi16
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 1; fi18++ {
                            v5 = (!false)
                            if (p1 == 18446744073709551615) {
                                var v19 uint8 = 84
                                v20 := []int{}
                                _ = v19
                                _ = v20
                            } else {
                                v21 := GenS2{F0: -316, F1: 607, F2: -9223372036854775808}
                                var v22 bool = (!((^(p0)) >= (-(p0))))
                                _ = v21
                                _ = v22
                            }
                            p0 = (p0 / (((p0 + 76)) | 1))
                            _ = fi18
                        }
                        _ = v15
                    } else {
                        v23 := []int{fi11, fi11, (fi11 + fi11), fi11}
                        for fi24 := 0; fi24 < 9; fi24++ {
                            p1 -= p1
                            fmt.Println(p1)
                            _ = fi24
                        }
                        _ = v23
                    }
                    _ = fi11
                    _ = v14
                }
            }
        }
        if false {
            v25 := []int{-881, 9223372036854775807, 678, 2147483647}
            fmt.Println(p1, p1, p0)
            _ = v25
        }
    }
    v26 := []int{-1, 32768, 32767}
    _ = p0
    _ = p1
    _ = v26
    _ = v5
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    var v3 int8 = -128
    fmt.Println(v3)
    var v27 string = "foo bar"
    var v28 uint32 = 1
    fmt.Println(v3, v27, v3)
    if (798 <= 1) {
        v29 := []int{}
        var v30 uint = 230
        var v31 int = 809
        _ = v29
        _ = v30
        _ = v31
    } else {
        var v32 uint8 = 128
        _ = v32
    }
    v28 = 1
    _ = len(v2)
    if false {
        fmt.Println(v3, v28, v27)
        v28 = 2
        var v33 uint16 = 36
        _ = v33
    }
    v3 += (v3 * v3)
    v34 := GenS1{F0: 276, F1: 984, F2: 9223372036854775807}
    v35 := GenS3{F0: (v28 | v28)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
}
