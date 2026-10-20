package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint8, p1 uint64) int16 {
    var v5 uint8 = 0
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint = 42
        v8 := []int{-9223372036854775808}
        v9 := GenS1{F0: v7, F1: -243}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    if true {
        for fi10 := 0; fi10 < 10; fi10++ {
            if ("foo bar" != ("hello" + "\n\t")) {
                for fi11 := 0; fi11 < 17; fi11++ {
                    v12 := GenS2{F0: ((42 > 32767) || (uint(p0) > 2)), F1: 65535}
                    v13 := GenS1{F0: 935, F1: 688}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v14 := GenS2{F0: true, F1: 0}
                v15 := []int{0, fi10}
                _ = v14
                _ = v15
            }
            fmt.Println(fi10)
            var v16 bool = (-1 <= 1)
            _ = fi10
            _ = v16
        }
        if false {
            var v17 int32 = 42
            _ = v17
        } else {
            switch 9223372036854775807 {
            case 42:
                v18 := []int{-316, -1, 0, 4294967295}
                _ = v18
            }
            for fi19 := 0; fi19 < 14; fi19++ {
                v5 = p0
                var v20 float32 = (1e10 * (0.0 / float32(((3.14159 - 0.1) / 1.5))))
                v21 := GenS1{F0: 255, F1: -501}
                _ = fi19
                _ = v20
                _ = v21
            }
            v22 := []int{-864, 271, -1}
            _ = v22
        }
    } else {
        p1 += p1
        p0 += (^(p0))
        fmt.Println(p0, v5, v5)
    }
    v5 += (^(p0))
    _ = p0
    _ = p1
    _ = v5
    return 32767
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint64 = (v1 | (v1 % (((v1 ^ (^(v1)))) | 1)))
    var v3 int = 0
    var v4 uint32 = 11
    if false {
        var v23 int8 = -55
        _ = v23
    }
    var v24 int16 = 1
    v24 = genFn0(1, (v2 % ((v1) | 1)))
    var v25 uint = 664
    for fi26 := 0; fi26 < 19; fi26++ {
        if ((true && true) && (((3.14159 / ((float64((0.1 - 0.0)) * 1e10) / 1.0)) >= ((1.5 - (3.14159 / (0.1 / (1.5 / 1e10)))) - (1.5 * 1.5))) || true)) {
            var v27 int32 = int32(v3)
            _ = v27
        } else {
            var v28 int8 = 52
            for fi29 := 0; fi29 < 20; fi29++ {
                var v30 bool = (65535 >= v4)
                if (v1 != v2) {
                    var v31 float64 = 1e10
                    _ = v31
                }
                v24 = v24
                _ = fi29
                _ = v30
            }
            v32 := GenS2{F0: false, F1: 42}
            _ = v28
            _ = v32
        }
        v33 := []int{(v3 % ((-297) | 1))}
        _ = fi26
        _ = v33
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        for fi35 := 0; fi35 < 10; fi35++ {
            if ((v1 - (v2 % (((v1 - v2)) | 1))) <= 864) {
                if false {
                    fmt.Println(v4, fi35)
                    fmt.Println(v24)
                } else {
                    var v36 float32 = (1e10 - float32((((((1.5 + (1e10 / 3.14159)) / ((0.1 / 1.0) - 3.14159)) + 1e10) * 1.5) + 1e10)))
                    var v37 float32 = (((0.1 + v36) + ((float32((1.5 * 1e10)) * 1.5) / v36)) * ((float32(1e10) * ((1.5 + (v36 * v36)) + v36)) / v36))
                    _ = v36
                    _ = v37
                }
                v38 := GenS1{F0: (^(v25)), F1: -504}
                if false {
                    var v39 string = "hello"
                    var v40 int = fi34
                    _ = v39
                    _ = v40
                } else {
                    v4 = v4
                    v41 := GenS3{F0: -1}
                    v42 := []int{}
                    _ = v41
                    _ = v42
                }
                _ = v38
            } else {
                switch ("" + "foo bar") {
                case "a":
                    var v43 uint64 = 0
                    _ = v43
                case "foo bar":
                    v1 -= (^(v2))
                case "":
                    switch -2147483648 {
                    case 0:
                        var v44 uint32 = v4
                        v2 = v2
                        _ = v44
                    case 699:
                        var v45 uint64 = v2
                        _ = v45
                    default:
                        if (false || true) {
                            var v46 int64 = -1
                            var v47 float64 = ((1e10 / 1.5) + 0.0)
                            v1 = (v2 % (((^(v2))) | 1))
                            v48 := []int{v3, 866, (fi35 >> ((102) & 63)), fi35}
                            v2 += (v1 - (v1 * (v1 >> (((^(v25))) & 63))))
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                    }
                    v49 := GenS4{F0: 42, F1: 229}
                    var v50 int = (fi34 / ((1) | 1))
                    _ = v49
                    _ = v50
                default:
                    var v51 uint = v25
                    _ = v51
                }
                v1 = 106
            }
            _ = fi35
        }
        v3 -= -77
        _ = fi34
    }
    v4 -= 65535
    fmt.Println(v3, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
